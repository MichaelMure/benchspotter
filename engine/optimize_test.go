package engine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ── pearsonCorr ───────────────────────────────────────────────────────────────

func TestPearsonCorr(t *testing.T) {
	t.Run("perfect positive", func(t *testing.T) {
		xs := []float64{1, 2, 3, 4, 5}
		ys := []float64{2, 4, 6, 8, 10}
		assert.InDelta(t, 1.0, pearsonCorr(xs, ys), 1e-9)
	})
	t.Run("perfect negative", func(t *testing.T) {
		xs := []float64{1, 2, 3, 4, 5}
		ys := []float64{10, 8, 6, 4, 2}
		assert.InDelta(t, -1.0, pearsonCorr(xs, ys), 1e-9)
	})
	t.Run("zero when y is constant", func(t *testing.T) {
		xs := []float64{1, 2, 3, 4, 5}
		ys := []float64{7, 7, 7, 7, 7}
		assert.Equal(t, 0.0, pearsonCorr(xs, ys))
	})
	t.Run("zero when x is constant", func(t *testing.T) {
		xs := []float64{3, 3, 3, 3, 3}
		ys := []float64{1, 2, 3, 4, 5}
		assert.Equal(t, 0.0, pearsonCorr(xs, ys))
	})
}

// ── InputInfo.SpacedValues ────────────────────────────────────────────────────

func TestSpacedValues(t *testing.T) {
	assert.Equal(t, []float64{0, 1}, BoolType{}.SpacedValues(5)) // always 2

	assert.Len(t, IntType{min: 0, max: 4}.SpacedValues(5), 5)
	assert.Len(t, IntType{min: 0, max: 4}.SpacedValues(3), 3)

	assert.Len(t, IntLogType{min: 1, max: 1000}.SpacedValues(16), 16)
	assert.Len(t, FloatType{min: 0, max: 1}.SpacedValues(10), 10)
	assert.Len(t, FloatLogType{min: 0.001, max: 1}.SpacedValues(7), 7)
}

// ── BestPoint ─────────────────────────────────────────────────────────────────

func TestBestPoint(t *testing.T) {
	inp := IntType{name: "x", min: 0, max: 10}
	mkPoint := func(x int, v float64) EvaluatedPoint {
		return EvaluatedPoint{
			Candidate: Candidate{{Input: inp, Value: itoa(x)}},
			Metrics:   map[string]float64{"ns/op": v},
		}
	}

	results := []EvaluatedPoint{mkPoint(1, 100), mkPoint(2, 50), mkPoint(3, 200)}

	t.Run("minimize", func(t *testing.T) {
		idx, ok := BestPoint(results, "ns/op", true)
		require.True(t, ok)
		assert.Equal(t, 1, idx) // 50 ns/op
	})
	t.Run("maximize", func(t *testing.T) {
		idx, ok := BestPoint(results, "ns/op", false)
		require.True(t, ok)
		assert.Equal(t, 2, idx) // 200 ns/op
	})
	t.Run("missing metric", func(t *testing.T) {
		_, ok := BestPoint(results, "B/op", true)
		assert.False(t, ok)
	})
	t.Run("all errors", func(t *testing.T) {
		errResults := []EvaluatedPoint{{Candidate: Candidate{{Input: inp, Value: "1"}}, Err: assert.AnError}}
		_, ok := BestPoint(errResults, "ns/op", true)
		assert.False(t, ok)
	})
	t.Run("empty", func(t *testing.T) {
		_, ok := BestPoint(nil, "ns/op", true)
		assert.False(t, ok)
	})
}

// ── InputCorrelations ─────────────────────────────────────────────────────────

func TestInputCorrelations(t *testing.T) {
	inp := IntType{name: "size", min: 1, max: 10}

	// Perfect positive correlation: larger size → larger ns/op.
	var results []EvaluatedPoint
	for i := 1; i <= 5; i++ {
		results = append(results, EvaluatedPoint{
			Candidate: Candidate{assign(inp, float64(i))},
			Metrics:   map[string]float64{"ns/op": float64(i) * 10},
		})
	}

	corr := InputCorrelations(results, []string{"ns/op"})
	require.NotNil(t, corr)
	assert.InDelta(t, 1.0, corr["size"]["ns/op"], 1e-9)

	t.Run("too few points returns no entry", func(t *testing.T) {
		corr := InputCorrelations(results[:2], []string{"ns/op"})
		if corr != nil {
			assert.Empty(t, corr["size"])
		}
	})

	t.Run("nil for empty results", func(t *testing.T) {
		assert.Nil(t, InputCorrelations(nil, []string{"ns/op"}))
	})
}

// ── CoordDescentStrategy ──────────────────────────────────────────────────────

func TestCoordDescentStrategy(t *testing.T) {
	// midpointValue for [0,4] is 2; use "0" as the best so there IS an
	// improvement over the midpoint and the strategy runs a second pass.
	inp := IntType{name: "x", min: 0, max: 4} // 5 distinct values, midpoint=2

	s := NewCoordDescentStrategy([]InputInfo{inp}, "ns/op", true)

	steps := len(inp.SpacedValues(coordSweepSteps))

	// Collect all sweep candidates.
	var phase1 []Candidate
	for len(phase1) < steps {
		c := s.Next()
		require.NotNil(t, c, "unexpected nil after %d candidates", len(phase1))
		phase1 = append(phase1, c)
	}
	assert.Len(t, phase1, steps)

	// All dispatched: Next should return nil while waiting for results.
	assert.Nil(t, s.Next())

	// Feed results: best is at value "0" (not the midpoint "2").
	for _, c := range phase1 {
		val := 100.0
		if c[0].Value == "0" {
			val = 10.0
		}
		s.Observe(EvaluatedPoint{
			Candidate: c,
			Metrics:   map[string]float64{"ns/op": val},
		})
	}

	// After improvement, strategy starts a new sweep.
	c := s.Next()
	require.NotNil(t, c, "expected second sweep after improvement")
}

func TestCoordDescentConverges(t *testing.T) {
	// Single bool input → only two values; after one sweep with no improvement, converges.
	inp := BoolType{name: "flag"}
	s := NewCoordDescentStrategy([]InputInfo{inp}, "ns/op", true)

	const maxIter = 100
	for i := 0; i < maxIter; i++ {
		c := s.Next()
		if c == nil {
			return // converged
		}
		s.Observe(EvaluatedPoint{
			Candidate: c,
			Metrics:   map[string]float64{"ns/op": 42},
		})
	}
	t.Fatal("strategy did not converge within", maxIter, "iterations")
}

// ── RandomStrategy ────────────────────────────────────────────────────────────

func TestRandomStrategy(t *testing.T) {
	inputs := []InputInfo{
		BoolType{name: "flag"},
		IntType{name: "n", min: 1, max: 100},
		FloatType{name: "f", min: 0.0, max: 1.0},
	}
	s := NewRandomStrategy(inputs)

	for i := 0; i < 50; i++ {
		c := s.Next()
		require.NotNil(t, c)
		require.Len(t, c, len(inputs))
		assert.Contains(t, []string{"true", "false"}, c[0].Value)

		n, err := parseInt(c[1].Value)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, n, 1)
		assert.LessOrEqual(t, n, 100)

		f, err := parseFloat(c[2].Value)
		require.NoError(t, err)
		assert.GreaterOrEqual(t, f, 0.0)
		assert.Less(t, f, 1.0+1e-9)
	}
}

// ── helpers ───────────────────────────────────────────────────────────────────

func itoa(n int) string {
	return fmt.Sprintf("%d", n)
}

func parseInt(s string) (int, error) {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	return n, err
}

func parseFloat(s string) (float64, error) {
	var f float64
	_, err := fmt.Sscanf(s, "%g", &f)
	return f, err
}
