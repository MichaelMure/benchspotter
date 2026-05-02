package commands

import (
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
	"benchspotter/repository"
)

func setupTrendStorage(t *testing.T) billy.Filesystem {
	t.Helper()
	storage := memfs.New()
	id1 := createTestSession(t, storage, "baseline", []string{"BenchmarkFoo"}, "", false)
	id2 := createTestSession(t, storage, "after-opt", []string{"BenchmarkFoo"}, "", false)
	writeBenchResults(t, storage, id1, "BenchmarkFoo-8\t1000000\t100 ns/op\t64 B/op\t2 allocs/op\nBenchmarkFoo-8\t1000000\t102 ns/op\t64 B/op\t2 allocs/op\n")
	writeBenchResults(t, storage, id2, "BenchmarkFoo-8\t1000000\t80 ns/op\t48 B/op\t1 allocs/op\nBenchmarkFoo-8\t1000000\t82 ns/op\t48 B/op\t1 allocs/op\n")
	return storage
}

func TestTrend(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		storage := setupTrendStorage(t)
		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		sessions, err := engine.LocateSessions(storage)
		require.NoError(t, err)
		data, err := engine.LoadTrendData(sessions, 0.95)
		require.NoError(t, err)

		err = renderTrendJSON(env, data, nil)
		require.NoError(t, err)

		out := env.Out.String()
		assert.Contains(t, out, `"benchmarks"`)
		assert.Contains(t, out, `"BenchmarkFoo"`)
		assert.Contains(t, out, `"ns/op"`)
		assert.Contains(t, out, `"sessions"`)
		assert.Contains(t, out, `"101.0ns"`)
	})

	t.Run("json_single_bench", func(t *testing.T) {
		storage := setupTrendStorage(t)
		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		sessions, err := engine.LocateSessions(storage)
		require.NoError(t, err)
		data, err := engine.LoadTrendData(sessions, 0.95)
		require.NoError(t, err)

		err = renderTrendJSON(env, data, []string{"BenchmarkFoo"})
		require.NoError(t, err)

		out := env.Out.String()
		assert.Contains(t, out, `"BenchmarkFoo"`)
	})

	t.Run("text", func(t *testing.T) {
		t.Run("overview", func(t *testing.T) {
			storage := setupTrendStorage(t)
			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))

			sessions, err := engine.LocateSessions(storage)
			require.NoError(t, err)
			data, err := engine.LoadTrendData(sessions, 0.95)
			require.NoError(t, err)

			err = renderTrendOverviewText(env, data)
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, "Foo")
			assert.Contains(t, out, "baseline")
			assert.Contains(t, out, "after-opt")
			assert.Contains(t, out, "ns")
		})

		t.Run("detail", func(t *testing.T) {
			storage := setupTrendStorage(t)
			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))

			sessions, err := engine.LocateSessions(storage)
			require.NoError(t, err)
			data, err := engine.LoadTrendData(sessions, 0.95)
			require.NoError(t, err)

			err = renderTrendDetailText(env, data, "BenchmarkFoo")
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, "BenchmarkFoo")
			assert.Contains(t, out, "baseline")
			assert.Contains(t, out, "after-opt")
			assert.Contains(t, out, "ns/op")
			assert.Contains(t, out, "B/op")
		})
	})

	t.Run("tagFilter", func(t *testing.T) {
		storage := memfs.New()
		id1 := createTestSession(t, storage, "untagged", []string{"BenchmarkFoo"}, "", false)
		id2 := createTestSession(t, storage, "tagged", []string{"BenchmarkFoo"}, "", false)
		writeBenchResults(t, storage, id1, "BenchmarkFoo-8\t1000000\t100 ns/op\n")
		writeBenchResults(t, storage, id2, "BenchmarkFoo-8\t1000000\t90 ns/op\n")
		addTagToSession(t, storage, id2, "prod")

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		// "prod" matches id2; "other" matches nothing — any-match should still return id2.
		sessions, err := engine.LocateSessions(storage, "prod", "other")
		require.NoError(t, err)
		require.Len(t, sessions, 1)
		assert.Equal(t, "tagged", sessions[0].Name)

		data, err := engine.LoadTrendData(sessions, 0.95)
		require.NoError(t, err)

		err = renderTrendJSON(env, data, nil)
		require.NoError(t, err)
		out := env.Out.String()
		assert.Contains(t, out, `"BenchmarkFoo"`)
	})

	t.Run("session filter", func(t *testing.T) {
		storage := memfs.New()
		id1 := createTestSession(t, storage, "session-a", []string{"BenchmarkFoo"}, "", false)
		id2 := createTestSession(t, storage, "session-b", []string{"BenchmarkFoo"}, "", false)
		createTestSession(t, storage, "session-c", []string{"BenchmarkFoo"}, "", false)
		writeBenchResults(t, storage, id1, "BenchmarkFoo-8\t1000000\t100 ns/op\n")
		writeBenchResults(t, storage, id2, "BenchmarkFoo-8\t1000000\t90 ns/op\n")
		// session-c intentionally has no bench results to confirm it is excluded

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		err := runTrend(env, trendOptions{
			sessions:   []string{id1, "session-b"}, // mix of ID and name
			confidence: 0.95,
		})
		require.NoError(t, err)

		out := env.Out.String()
		assert.Contains(t, out, "session-a")
		assert.Contains(t, out, "session-b")
		assert.NotContains(t, out, "session-c")
	})
}

func setupTrendStorageMultiBench(t *testing.T) billy.Filesystem {
	t.Helper()
	storage := memfs.New()
	id1 := createTestSession(t, storage, "s1", []string{"BenchmarkBar", "BenchmarkFoo"}, "", false)
	id2 := createTestSession(t, storage, "s2", []string{"BenchmarkBar", "BenchmarkFoo"}, "", false)
	writeBenchResults(t, storage, id1,
		"BenchmarkBar-8\t1000000\t200 ns/op\n"+
			"BenchmarkFoo-8\t1000000\t100 ns/op\n")
	writeBenchResults(t, storage, id2,
		"BenchmarkBar-8\t1000000\t180 ns/op\n"+
			"BenchmarkFoo-8\t1000000\t90 ns/op\n")
	return storage
}

func TestTrendBenchFlag(t *testing.T) {
	t.Run("json_filters_to_matching", func(t *testing.T) {
		storage := setupTrendStorageMultiBench(t)
		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		require.NoError(t, runTrend(env, trendOptions{bench: "Foo", confidence: 0.95}))

		out := env.Out.String()
		assert.Contains(t, out, `"BenchmarkFoo"`)
		assert.NotContains(t, out, `"BenchmarkBar"`)
	})

	t.Run("json_multi_match_all_included", func(t *testing.T) {
		storage := setupTrendStorageMultiBench(t)
		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		require.NoError(t, runTrend(env, trendOptions{bench: "Benchmark", confidence: 0.95}))

		out := env.Out.String()
		assert.Contains(t, out, `"BenchmarkBar"`)
		assert.Contains(t, out, `"BenchmarkFoo"`)
	})

	t.Run("json_no_match_error", func(t *testing.T) {
		storage := setupTrendStorageMultiBench(t)
		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		err := runTrend(env, trendOptions{bench: "Zzz", confidence: 0.95})
		require.ErrorContains(t, err, "no benchmarks match")
	})

	t.Run("text_opens_detail_for_first_match", func(t *testing.T) {
		// BenchNames are sorted, so BenchmarkBar < BenchmarkFoo: first match for "Benchmark" is BenchmarkBar.
		storage := setupTrendStorageMultiBench(t)
		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))

		require.NoError(t, runTrend(env, trendOptions{bench: "Benchmark", confidence: 0.95}))

		out := env.Out.String()
		assert.Contains(t, out, "BenchmarkBar")
		assert.NotContains(t, out, "BenchmarkFoo")
	})

	t.Run("text_no_match_error", func(t *testing.T) {
		storage := setupTrendStorageMultiBench(t)
		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))

		err := runTrend(env, trendOptions{bench: "Zzz", confidence: 0.95})
		require.ErrorContains(t, err, "no benchmarks match")
	})

	t.Run("invalid_regex_error", func(t *testing.T) {
		storage := setupTrendStorageMultiBench(t)
		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))

		err := runTrend(env, trendOptions{bench: "[", confidence: 0.95})
		require.ErrorContains(t, err, "invalid --bench pattern")
	})
}

func TestFormatMetricValue(t *testing.T) {
	cases := []struct {
		v    float64
		unit string
		want string
	}{
		{500, "ns/op", "500.0ns"},
		{1500, "ns/op", "1.50µs"},
		{1500000, "ns/op", "1.50ms"},
		{2000000000, "ns/op", "2.00s"},
		{512, "B/op", "512B"},
		{2048, "B/op", "2.0KB"},
		{3145728, "B/op", "3.0MB"},
		{5, "allocs/op", "5"},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.want, formatMetricValue(tc.v, tc.unit), "v=%v unit=%s", tc.v, tc.unit)
	}
}
