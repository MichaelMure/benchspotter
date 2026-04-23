package commands

import (
	"encoding/json"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
	"benchspotter/repository"
)

func makeOptimizeResultCh(points []engine.EvaluatedPoint) <-chan engine.EvaluatedPoint {
	ch := make(chan engine.EvaluatedPoint, len(points))
	for _, p := range points {
		ch <- p
	}
	close(ch)
	return ch
}

func optimizeTestFixtures() (results []engine.EvaluatedPoint, bench engine.BenchInfo, strategy engine.Strategy) {
	inp := engine.IntType{}
	bench = engine.BenchInfo{Name: "BenchmarkFoo", Package: "."}
	strategy = engine.NewRandomStrategy([]engine.InputInfo{inp})
	results = []engine.EvaluatedPoint{
		{
			Candidate: engine.Candidate{{Input: inp, Value: "1"}},
			Metrics:   map[string]float64{"ns/op": 200, "B/op": 64},
		},
		{
			Candidate: engine.Candidate{{Input: inp, Value: "2"}},
			Metrics:   map[string]float64{"ns/op": 100, "B/op": 32},
		},
		{
			Candidate: engine.Candidate{{Input: inp, Value: "3"}},
			Metrics:   map[string]float64{"ns/op": 150, "B/op": 48},
		},
	}
	return
}

//go:fix inline
func newOptimizeTestEnv(t *testing.T) *execenv.Env {
	return execenv.NewTestEnv(repository.NewForTesting(memfs.New(), memfs.New(), nil))
}

func TestOptimize(t *testing.T) {
	t.Run("json/shape", func(t *testing.T) {
		results, bench, strategy := optimizeTestFixtures()
		env := newOptimizeTestEnv(t)
		env.Format = execenv.FormatJSON

		require.NoError(t, runOptimizeJSON(env, bench, strategy, makeOptimizeResultCh(results), "ns/op", true))

		var out optimizeJSONOutput
		require.NoError(t, json.Unmarshal([]byte(env.Out.String()), &out))

		assert.Equal(t, "BenchmarkFoo", out.Benchmark)
		assert.Equal(t, "ns/op", out.Metric)
		assert.True(t, out.Minimize)
		assert.Len(t, out.Trials, 3)
	})

	t.Run("json/best", func(t *testing.T) {
		results, bench, strategy := optimizeTestFixtures()
		env := newOptimizeTestEnv(t)
		env.Format = execenv.FormatJSON

		require.NoError(t, runOptimizeJSON(env, bench, strategy, makeOptimizeResultCh(results), "ns/op", true))

		var out optimizeJSONOutput
		require.NoError(t, json.Unmarshal([]byte(env.Out.String()), &out))

		// Trial 2 has the lowest ns/op (100).
		require.NotNil(t, out.Best)
		assert.Equal(t, 2, out.Best.N)
		assert.InDelta(t, 100.0, out.Best.Metrics["ns/op"], 1e-9)
	})

	t.Run("json/correlation_present", func(t *testing.T) {
		results, bench, strategy := optimizeTestFixtures()
		env := newOptimizeTestEnv(t)
		env.Format = execenv.FormatJSON

		require.NoError(t, runOptimizeJSON(env, bench, strategy, makeOptimizeResultCh(results), "ns/op", true))

		var out optimizeJSONOutput
		require.NoError(t, json.Unmarshal([]byte(env.Out.String()), &out))

		assert.NotEmpty(t, out.Correlation, "expected correlation for 3+ results")
	})

	t.Run("json/error_trial", func(t *testing.T) {
		inp := engine.IntType{}
		bench := engine.BenchInfo{Name: "BenchmarkBar", Package: "."}
		strategy := engine.NewRandomStrategy([]engine.InputInfo{inp})
		results := []engine.EvaluatedPoint{
			{Candidate: engine.Candidate{{Input: inp, Value: "1"}}, Err: assert.AnError},
			{Candidate: engine.Candidate{{Input: inp, Value: "2"}}, Metrics: map[string]float64{"ns/op": 50}},
		}

		env := newOptimizeTestEnv(t)
		env.Format = execenv.FormatJSON

		require.NoError(t, runOptimizeJSON(env, bench, strategy, makeOptimizeResultCh(results), "ns/op", true))

		var out optimizeJSONOutput
		require.NoError(t, json.Unmarshal([]byte(env.Out.String()), &out))

		assert.Len(t, out.Trials, 2)
		assert.NotEmpty(t, out.Trials[0].Error)
		assert.Empty(t, out.Trials[0].Metrics)
		assert.Empty(t, out.Trials[1].Error)
		assert.NotEmpty(t, out.Trials[1].Metrics)
	})

	t.Run("plain/streaming", func(t *testing.T) {
		results, _, _ := optimizeTestFixtures()
		env := newOptimizeTestEnv(t)

		require.NoError(t, runOptimizePlain(env, makeOptimizeResultCh(results), "ns/op", true, 0))

		out := env.Out.String()
		assert.Contains(t, out, "Trial 1")
		assert.Contains(t, out, "Trial 2")
		assert.Contains(t, out, "Trial 3")
		assert.Contains(t, out, "← best")
		assert.Contains(t, out, "Best:")
	})

	t.Run("plain/max_trials_prefix", func(t *testing.T) {
		results, _, _ := optimizeTestFixtures()
		env := newOptimizeTestEnv(t)

		require.NoError(t, runOptimizePlain(env, makeOptimizeResultCh(results), "ns/op", true, 10))

		assert.Contains(t, env.Out.String(), "Trial 1/10")
	})

	t.Run("plain/maximize", func(t *testing.T) {
		results, _, _ := optimizeTestFixtures()
		env := newOptimizeTestEnv(t)

		require.NoError(t, runOptimizePlain(env, makeOptimizeResultCh(results), "ns/op", false, 0))

		out := env.Out.String()
		// Trial 1 has the highest ns/op (200) → best when maximizing.
		assert.Contains(t, out, "Best:")
		assert.Contains(t, out, "200")
	})

	t.Run("report/best_and_correlation", func(t *testing.T) {
		results, _, _ := optimizeTestFixtures()
		env := newOptimizeTestEnv(t)

		printOptimizeReport(env, results, "ns/op", true)

		out := env.Out.String()
		assert.Contains(t, out, "Best:")
		assert.Contains(t, out, "← best")
		assert.Contains(t, out, "correlation")
		assert.Contains(t, out, "ns/op")
	})

	t.Run("report/empty", func(t *testing.T) {
		env := newOptimizeTestEnv(t)
		printOptimizeReport(env, nil, "ns/op", true)
		assert.Empty(t, env.Out.String())
	})
}
