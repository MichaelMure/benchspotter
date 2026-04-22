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
		env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		sessions, err := engine.LocateSessions(storage)
		require.NoError(t, err)
		data, err := engine.LoadTrendData(sessions, 0.95)
		require.NoError(t, err)

		err = renderTrendJSON(env, data, "")
		require.NoError(t, err)

		out := env.Out.String()
		assert.Contains(t, out, `"benchmarks"`)
		assert.Contains(t, out, `"BenchmarkFoo"`)
		assert.Contains(t, out, `"ns/op"`)
		assert.Contains(t, out, `"center"`)
		assert.Contains(t, out, `"session_id"`)
	})

	t.Run("json_single_bench", func(t *testing.T) {
		storage := setupTrendStorage(t)
		env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		sessions, err := engine.LocateSessions(storage)
		require.NoError(t, err)
		data, err := engine.LoadTrendData(sessions, 0.95)
		require.NoError(t, err)

		err = renderTrendJSON(env, data, "BenchmarkFoo")
		require.NoError(t, err)

		out := env.Out.String()
		assert.Contains(t, out, `"BenchmarkFoo"`)
	})

	t.Run("text", func(t *testing.T) {
		t.Run("overview", func(t *testing.T) {
			storage := setupTrendStorage(t)
			env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))

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
			env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))

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

		env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		// "prod" matches id2; "other" matches nothing — any-match should still return id2.
		sessions, err := engine.LocateSessions(storage, "prod", "other")
		require.NoError(t, err)
		require.Len(t, sessions, 1)
		assert.Equal(t, "tagged", sessions[0].Name)

		data, err := engine.LoadTrendData(sessions, 0.95)
		require.NoError(t, err)

		err = renderTrendJSON(env, data, "")
		require.NoError(t, err)
		out := env.Out.String()
		assert.Contains(t, out, `"BenchmarkFoo"`)
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
