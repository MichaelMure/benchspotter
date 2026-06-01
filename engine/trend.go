package engine

import (
	"math"
	"sort"

	"golang.org/x/perf/benchfmt"
	"golang.org/x/perf/benchmath"
)

// TrendComparison controls which reference session each data point is compared against.
type TrendComparison int

const (
	ComparisonSequential TrendComparison = iota // each session vs the previous one
	ComparisonBaseline                          // each session vs the first session
	trendComparisonCount
)

var trendComparisonNames = [trendComparisonCount]string{"sequential", "baseline"}

func (c TrendComparison) String() string        { return trendComparisonNames[c] }
func (c TrendComparison) Next() TrendComparison { return (c + 1) % trendComparisonCount }

// TrendPoint is one session's summarised result for a benchmark/unit pair.
type TrendPoint struct {
	Session *SessionInfo
	Center  float64 // median
	Lo      float64 // confidence interval low
	Hi      float64 // confidence interval high
	N       int     // number of measurements
}

// TrendData holds all trend points loaded from a set of sessions.
type TrendData struct {
	Sessions   []*SessionInfo
	BenchNames []string                           // sorted unique names (GOMAXPROCS stripped)
	Units      []string                           // sorted, ns/op first when present
	Points     map[string]map[string][]TrendPoint // [bench][unit] in session order
}

// LoadTrendData reads each session file once and returns a TrendData covering
// all benchmarks and units found.
func LoadTrendData(sessions []*SessionInfo, confidence float64) (*TrendData, error) {
	filtered := sessions[:0:0]
	for _, s := range sessions {
		if s.HasBench() {
			filtered = append(filtered, s)
		}
	}
	sessions = filtered

	type accumKey struct {
		bench string
		unit  string
		idx   int
	}
	accum := make(map[accumKey][]float64)
	seenBench := make(map[string]struct{})
	seenUnit := make(map[string]struct{})

	for idx, s := range sessions {
		f, err := s.OpenFile(benchFilename)
		if err != nil {
			continue
		}
		reader := benchfmt.NewReader(f, s.Id)
		for reader.Scan() {
			res, ok := reader.Result().(*benchfmt.Result)
			if !ok {
				continue
			}
			bench := "Benchmark" + stripGOMAXPROCS(string(res.Name))
			seenBench[bench] = struct{}{}
			for _, v := range res.Values {
				unit, value := origUnitValue(v)
				if math.IsNaN(value) || math.IsInf(value, 0) {
					continue
				}
				seenUnit[unit] = struct{}{}
				k := accumKey{bench: bench, unit: unit, idx: idx}
				accum[k] = append(accum[k], value)
			}
		}
		_ = f.Close()
	}

	benchNames := sortedMapKeys(seenBench)
	units := sortedMapKeys(seenUnit)
	preferredUnitsFirst(units)

	td := &TrendData{
		Sessions:   sessions,
		BenchNames: benchNames,
		Units:      units,
		Points:     make(map[string]map[string][]TrendPoint),
	}

	thresholds := benchmath.DefaultThresholds
	for _, bench := range benchNames {
		td.Points[bench] = make(map[string][]TrendPoint)
		for _, unit := range units {
			for idx, s := range sessions {
				vals := accum[accumKey{bench: bench, unit: unit, idx: idx}]
				if len(vals) == 0 {
					continue
				}
				sample := benchmath.NewSample(vals, &thresholds)
				sum := benchmath.AssumeNothing.Summary(sample, confidence)
				lo, hi := sum.Lo, sum.Hi
				if math.IsInf(lo, 0) || math.IsNaN(lo) {
					lo = sum.Center
				}
				if math.IsInf(hi, 0) || math.IsNaN(hi) {
					hi = sum.Center
				}
				td.Points[bench][unit] = append(td.Points[bench][unit], TrendPoint{
					Session: s,
					Center:  sum.Center,
					Lo:      lo,
					Hi:      hi,
					N:       len(vals),
				})
			}
		}
	}

	return td, nil
}

func sortedMapKeys(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// preferredUnitsFirst reorders units so ns/op, B/op, allocs/op come first.
func preferredUnitsFirst(units []string) {
	order := map[string]int{"ns/op": 0, "B/op": 1, "allocs/op": 2}
	sort.SliceStable(units, func(i, j int) bool {
		oi, oki := order[units[i]]
		oj, okj := order[units[j]]
		if oki && okj {
			return oi < oj
		}
		if oki {
			return true
		}
		if okj {
			return false
		}
		return units[i] < units[j]
	})
}

// origUnitValue returns the human-readable unit and value, reversing benchfmt's
// normalization (e.g. sec/op → ns/op with the original nanosecond value).
func origUnitValue(v benchfmt.Value) (string, float64) {
	if v.OrigUnit != "" {
		return v.OrigUnit, v.OrigValue
	}
	return v.Unit, v.Value
}

// stripGOMAXPROCS removes the trailing "-N" GOMAXPROCS suffix from a benchmark name.
func stripGOMAXPROCS(name string) string {
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '-' && i < len(name)-1 {
			return name[:i]
		}
		if name[i] < '0' || name[i] > '9' {
			break
		}
	}
	return name
}
