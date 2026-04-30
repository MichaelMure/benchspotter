package engine

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"iter"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-billy/v5"
	"github.com/google/pprof/profile"
	"golang.org/x/sys/execabs"
)

var replacer = strings.NewReplacer(
	"/", "ᚋ",
	".", "ᚗ",
	"-", "ᚑ",
	"~", "א",
)

var reverseReplacer = strings.NewReplacer(
	"ᚋ", "/",
	"ᚗ", ".",
	"ᚑ", "-",
	"א", "~",
)

func RunProfile(ctx context.Context, storage billy.Filesystem, id string, benches []BenchInfo, profile Profile) func() error {
	next, _ := iter.Pull(func(yield func(error) bool) {
		for _, infos := range benches {
			var dir, filename, option string
			switch profile {
			case ProfileCPU:
				dir, filename, option = "cpu", replacer.Replace(infos.Package+"."+infos.Name)+".profile", "-cpuprofile"
			case ProfileMem:
				dir, filename, option = "mem", replacer.Replace(infos.Package+"."+infos.Name)+".profile", "-memprofile"
			case ProfileMutex:
				dir, filename, option = "mutex", replacer.Replace(infos.Package+"."+infos.Name)+".profile", "-mutexprofile"
			case ProfileBlock:
				dir, filename, option = "block", replacer.Replace(infos.Package+"."+infos.Name)+".profile", "-blockprofile"
			default:
				panic("invalid profile")
			}

			subDir := storage.Join(sessionDir, id, dir)
			if err := storage.MkdirAll(subDir, 0755); err != nil {
				yield(err)
				return
			}
			outFilename := storage.Join(storage.Root(), subDir, filename)

			cmd := execabs.CommandContext(ctx, "go", "test",
				"-bench", infos.Regex(),
				option+"="+outFilename,
				"-run", "^$", ".")
			cmd.Dir = infos.Package

			if out, err := cmd.CombinedOutput(); err != nil {
				yield(fmt.Errorf("failed to profile: %w, %s", err, out))
				return
			}
			if !yield(nil) {
				return
			}
		}
	})
	return func() error {
		err, _ := next()
		return err
	}
}

type ProfileFunc struct {
	Name       string
	File       string // source file (absolute path from pprof data)
	StartLine  int64  // function definition line (1-based)
	Flat       time.Duration
	Cumulative time.Duration
	FlatPct    float64
	CumPct     float64
}

// ProfileDir returns the storage subdirectory for a pprof profile type.
// Panics for ProfileBench and ProfileEscape, which are not pprof profiles.
func ProfileDir(p Profile) string {
	switch p {
	case ProfileCPU:
		return "cpu"
	case ProfileMem:
		return "mem"
	case ProfileMutex:
		return "mutex"
	case ProfileBlock:
		return "block"
	default:
		panic("invalid profile")
	}
}

// ProfileBenchSummary summarises a single benchmark's profile for display.
// For CPU, Block and Mutex, Primary holds nanoseconds. For Mem, Primary holds
// inuse_space bytes and AllocBytes holds alloc_space bytes.
type ProfileBenchSummary struct {
	Name       string
	Primary    int64
	AllocBytes int64 // Mem only
}

// ListProfileBenchmarkSummaries is like ListProfileBenchmarks but also reads
// each profile to compute a quick total for display in interactive selectors.
func ListProfileBenchmarkSummaries(fs billy.Filesystem, sessionPath string, p Profile) ([]ProfileBenchSummary, error) {
	names, err := ListProfileBenchmarks(fs, sessionPath, p)
	if err != nil {
		return nil, err
	}
	summaries := make([]ProfileBenchSummary, len(names))
	for i, name := range names {
		prof, err := readProfile(fs, sessionPath, p, name)
		if err != nil {
			return nil, err
		}
		s := ProfileBenchSummary{Name: name, Primary: sumProfileValues(prof, PickValueIndex(prof, p))}
		if p == ProfileMem {
			s.AllocBytes = sumProfileValues(prof, MemMetricIndex(prof, MetricAllocSpace))
		}
		summaries[i] = s
	}
	return summaries, nil
}

func sumProfileValues(prof *profile.Profile, idx int) int64 {
	var total int64
	for _, s := range prof.Sample {
		if len(s.Value) > idx {
			total += s.Value[idx]
		}
	}
	return total
}

// ListProfileBenchmarks returns the benchmark names that have profiles stored
// for the given session and profile type.
func ListProfileBenchmarks(fs billy.Filesystem, sessionPath string, p Profile) ([]string, error) {
	dir := filepath.Join(sessionPath, ProfileDir(p))
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("no %s profile found for this session", ProfileDir(p))
	}
	var names []string
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".profile") {
			continue
		}
		decoded := reverseReplacer.Replace(strings.TrimSuffix(e.Name(), ".profile"))
		// decoded is "package.BenchmarkName" — benchmark name is after the last dot
		if idx := strings.LastIndex(decoded, "."); idx >= 0 {
			names = append(names, decoded[idx+1:])
		}
	}
	return names, nil
}

// ReadProfileFunctions parses the pprof file for the given session, profile
// type, and benchmark name, and returns the functions sorted by flat time
// descending. Use ListProfileBenchmarks to enumerate valid bench values.
func ReadProfileFunctions(fs billy.Filesystem, sessionPath string, p Profile, bench string) ([]ProfileFunc, error) {
	prof, err := readProfile(fs, sessionPath, p, bench)
	if err != nil {
		return nil, err
	}
	return AggregateFuncs(prof, PickValueIndex(prof, p), SortFlat), nil
}

// ReadParsedProfile parses the pprof file for the given session, profile type,
// and benchmark, returning the raw profile for custom aggregation with
// AggregateFuncs and MemMetricIndex / PickValueIndex.
func ReadParsedProfile(fs billy.Filesystem, sessionPath string, p Profile, bench string) (*profile.Profile, error) {
	return readProfile(fs, sessionPath, p, bench)
}

// ReadProfileRaw parses the pprof file for the given session, profile type, and
// benchmark name, and returns the re-serialized binary suitable for piping to
// go tool pprof. Use ListProfileBenchmarks to enumerate valid bench values.
func ReadProfileRaw(fs billy.Filesystem, sessionPath string, p Profile, bench string) ([]byte, error) {
	prof, err := readProfile(fs, sessionPath, p, bench)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := prof.Write(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ProfileFilePath returns the absolute OS path to the pprof file for the given
// session, profile type, and benchmark. The storage filesystem must be an OS
// filesystem (osfs), which is always the case in production.
func ProfileFilePath(fs billy.Filesystem, sessionPath string, p Profile, bench string) (string, error) {
	dir := filepath.Join(sessionPath, ProfileDir(p))
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("no %s profile found for this session", ProfileDir(p))
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".profile") {
			continue
		}
		decoded := reverseReplacer.Replace(strings.TrimSuffix(e.Name(), ".profile"))
		name := decoded
		if idx := strings.LastIndex(decoded, "."); idx >= 0 {
			name = decoded[idx+1:]
		}
		if name != bench {
			continue
		}
		return filepath.Join(fs.Root(), dir, e.Name()), nil
	}
	return "", fmt.Errorf("no %s profile found for benchmark %q in this session", ProfileDir(p), bench)
}

// readProfile opens the single pprof file for bench and parses it.
// bench must be a non-empty name returned by ListProfileBenchmarks.
func readProfile(fs billy.Filesystem, sessionPath string, p Profile, bench string) (*profile.Profile, error) {
	dir := filepath.Join(sessionPath, ProfileDir(p))
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("no %s profile found for this session", ProfileDir(p))
	}

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".profile") {
			continue
		}
		decoded := reverseReplacer.Replace(strings.TrimSuffix(e.Name(), ".profile"))
		name := decoded
		if idx := strings.LastIndex(decoded, "."); idx >= 0 {
			name = decoded[idx+1:]
		}
		if name != bench {
			continue
		}
		f, err := fs.Open(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, err
		}
		defer f.Close()
		data, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}
		return profile.ParseData(data)
	}

	return nil, fmt.Errorf("no %s profile found for benchmark %q in this session", ProfileDir(p), bench)
}

// AggregateFuncs aggregates pprof samples into per-function flat/cumulative
// totals using the given value index. Flat counts a function only when it
// appears at the top of the call stack (i.e. it was on-CPU or holding the lock
// at sample time). Cumulative counts it whenever it appears anywhere in the
// stack, deduplicated per sample. Percentages are relative to the total value
// across all samples.
//
// This reimplements the core of github.com/google/pprof/internal/report, which
// is not importable outside the pprof module.
func AggregateFuncs(prof *profile.Profile, valueIdx int, order SortOrder) []ProfileFunc {

	type entry struct {
		name      string
		file      string
		startLine int64
		flat      int64
		cum       int64
	}
	byFunc := make(map[string]*entry)

	var total int64
	for _, s := range prof.Sample {
		if len(s.Value) <= valueIdx {
			continue
		}
		v := s.Value[valueIdx]
		total += v

		if len(s.Location) > 0 {
			for _, line := range s.Location[0].Line {
				name := line.Function.Name
				e := byFunc[name]
				if e == nil {
					e = &entry{name: name, file: line.Function.Filename, startLine: line.Function.StartLine}
					byFunc[name] = e
				}
				e.flat += v
			}
		}
		seen := make(map[string]bool)
		for _, loc := range s.Location {
			for _, line := range loc.Line {
				name := line.Function.Name
				if seen[name] {
					continue
				}
				seen[name] = true
				e := byFunc[name]
				if e == nil {
					e = &entry{name: name}
					byFunc[name] = e
				}
				e.cum += v
			}
		}
	}

	result := make([]ProfileFunc, 0, len(byFunc))
	for _, e := range byFunc {
		var flatPct, cumPct float64
		if total > 0 {
			flatPct = float64(e.flat) / float64(total) * 100
			cumPct = float64(e.cum) / float64(total) * 100
		}
		result = append(result, ProfileFunc{
			Name:       e.name,
			File:       e.file,
			StartLine:  e.startLine,
			Flat:       time.Duration(e.flat),
			Cumulative: time.Duration(e.cum),
			FlatPct:    flatPct,
			CumPct:     cumPct,
		})
	}

	switch order {
	case SortCumulative:
		sort.Slice(result, func(i, j int) bool {
			if result[i].Cumulative != result[j].Cumulative {
				return result[i].Cumulative > result[j].Cumulative
			}
			if result[i].Flat != result[j].Flat {
				return result[i].Flat > result[j].Flat
			}
			return result[i].Name < result[j].Name
		})
	case SortName:
		sort.Slice(result, func(i, j int) bool {
			return result[i].Name < result[j].Name
		})
	default: // SortFlat
		sort.Slice(result, func(i, j int) bool {
			if result[i].Flat != result[j].Flat {
				return result[i].Flat > result[j].Flat
			}
			if result[i].Cumulative != result[j].Cumulative {
				return result[i].Cumulative > result[j].Cumulative
			}
			return result[i].Name < result[j].Name
		})
	}

	return result
}

// MemMetricIndex returns the SampleType index for the given MemMetric, or 0
// if not found.
func MemMetricIndex(prof *profile.Profile, m MemMetric) int {
	for i, st := range prof.SampleType {
		if st.Type == m.String() {
			return i
		}
	}
	return 0
}

// PickValueIndex returns the index into prof.SampleType whose Type we want to
// aggregate. A pprof file stores multiple parallel value columns per sample;
// which column is meaningful depends on the profile type:
//
//	CPU:         [samples/count, cpu/nanoseconds]                                       → cpu
//	Mem:         [alloc_objects/count, alloc_space/bytes, inuse_objects/count, inuse_space/bytes] → inuse_space
//	Block/Mutex: [contentions/count, delay/nanoseconds]                                → delay
//
// Falls back to index 0 if the preferred type is not found.
func PickValueIndex(prof *profile.Profile, p Profile) int {
	preferred := map[Profile]string{
		ProfileCPU:   "cpu",
		ProfileMem:   MetricInuseSpace.String(),
		ProfileBlock: "delay",
		ProfileMutex: "delay",
	}
	want, ok := preferred[p]
	if ok {
		for i, st := range prof.SampleType {
			if st.Type == want {
				return i
			}
		}
	}
	return 0
}
