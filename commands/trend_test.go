package commands

import (
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/MichaelMure/benchspotter/commands/execenv"
	"github.com/MichaelMure/benchspotter/engine"
	"github.com/MichaelMure/benchspotter/repository"
)

// setupTrendStorage builds three sessions:
//
//	s1 = 100+102 ns/op, 64 B/op, 2 allocs/op  (baseline; median ~101 ns)
//	s2 =  80+82  ns/op, 48 B/op, 1 allocs/op  (improvement vs s1)
//	s3 =  90+88  ns/op, 56 B/op, 2 allocs/op  (still better than s1, but regression vs s2)
//
// In sequential mode s3 shows ↑ (regression vs s2).
// In baseline mode  s3 shows ↓ (improvement vs s1).
func setupTrendStorage(t *testing.T) *repository.Repository {
	t.Helper()
	storage := memfs.New()
	id1 := createTestSession(t, storage, "s1", []string{"BenchmarkFoo"}, "", false)
	id2 := createTestSession(t, storage, "s2", []string{"BenchmarkFoo"}, "", false)
	id3 := createTestSession(t, storage, "s3", []string{"BenchmarkFoo"}, "", false)
	writeBenchResults(t, storage, id1,
		"BenchmarkFoo-8\t1000000\t100 ns/op\t64 B/op\t2 allocs/op\n"+
			"BenchmarkFoo-8\t1000000\t102 ns/op\t64 B/op\t2 allocs/op\n")
	writeBenchResults(t, storage, id2,
		"BenchmarkFoo-8\t1000000\t80 ns/op\t48 B/op\t1 allocs/op\n"+
			"BenchmarkFoo-8\t1000000\t82 ns/op\t48 B/op\t1 allocs/op\n")
	writeBenchResults(t, storage, id3,
		"BenchmarkFoo-8\t1000000\t90 ns/op\t56 B/op\t2 allocs/op\n"+
			"BenchmarkFoo-8\t1000000\t88 ns/op\t56 B/op\t2 allocs/op\n")
	return repository.NewForTesting(memfs.New(), storage, nil)
}

func TestTrend(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		repo := setupTrendStorage(t)
		env := execenv.NewTestEnv(t.Context(), repo)
		env.Format = execenv.FormatJSON

		sessions, err := engine.LocateSessions(repo.Storage())
		require.NoError(t, err)
		data, err := engine.LoadTrendData(sessions, 0.95)
		require.NoError(t, err)

		require.NoError(t, renderTrendJSON(env, data, nil))

		out := env.Out.String()
		assert.Contains(t, out, `"benchmarks"`)
		assert.Contains(t, out, `"BenchmarkFoo"`)
		assert.Contains(t, out, `"ns/op"`)
		assert.Contains(t, out, `"sessions"`)
		assert.Contains(t, out, `"101.0ns"`)
	})

	t.Run("json_single_bench", func(t *testing.T) {
		repo := setupTrendStorage(t)
		env := execenv.NewTestEnv(t.Context(), repo)
		env.Format = execenv.FormatJSON

		sessions, err := engine.LocateSessions(repo.Storage())
		require.NoError(t, err)
		data, err := engine.LoadTrendData(sessions, 0.95)
		require.NoError(t, err)

		require.NoError(t, renderTrendJSON(env, data, []string{"BenchmarkFoo"}))

		assert.Contains(t, env.Out.String(), `"BenchmarkFoo"`)
	})

	t.Run("text", func(t *testing.T) {
		t.Run("overview", func(t *testing.T) {
			repo := setupTrendStorage(t)
			env := execenv.NewTestEnv(t.Context(), repo)

			sessions, err := engine.LocateSessions(repo.Storage())
			require.NoError(t, err)
			data, err := engine.LoadTrendData(sessions, 0.95)
			require.NoError(t, err)

			require.NoError(t, renderTrendOverviewText(env, data, trendOptions{}))

			out := env.Out.String()
			assert.Contains(t, out, "Foo")
			assert.Contains(t, out, "s1")
			assert.Contains(t, out, "s2")
			assert.Contains(t, out, "s3")
			assert.Contains(t, out, "ns")
		})

		t.Run("detail", func(t *testing.T) {
			repo := setupTrendStorage(t)
			env := execenv.NewTestEnv(t.Context(), repo)

			sessions, err := engine.LocateSessions(repo.Storage())
			require.NoError(t, err)
			data, err := engine.LoadTrendData(sessions, 0.95)
			require.NoError(t, err)

			require.NoError(t, renderTrendDetailText(env, data, trendOptions{bench: "BenchmarkFoo"}))

			out := env.Out.String()
			assert.Contains(t, out, "BenchmarkFoo")
			assert.Contains(t, out, "s1")
			assert.Contains(t, out, "s2")
			assert.Contains(t, out, "s3")
			assert.Contains(t, out, "ns/op")
			assert.Contains(t, out, "B/op")
		})
	})

	t.Run("comparison", func(t *testing.T) {
		t.Run("overview", func(t *testing.T) {
			t.Run("sequential", func(t *testing.T) {
				repo := setupTrendStorage(t)
				env := execenv.NewTestEnv(t.Context(), repo)

				sessions, err := engine.LocateSessions(repo.Storage())
				require.NoError(t, err)
				data, err := engine.LoadTrendData(sessions, 0.95)
				require.NoError(t, err)

				require.NoError(t, renderTrendOverviewText(env, data, trendOptions{comparison: engine.ComparisonSequential}))

				// s3 (~89ns) vs s2 (~81ns) → ratio ≈1.10 → ↑ (regression)
				assert.Contains(t, env.Out.String(), "↑")
			})

			t.Run("baseline", func(t *testing.T) {
				repo := setupTrendStorage(t)
				env := execenv.NewTestEnv(t.Context(), repo)

				sessions, err := engine.LocateSessions(repo.Storage())
				require.NoError(t, err)
				data, err := engine.LoadTrendData(sessions, 0.95)
				require.NoError(t, err)

				require.NoError(t, renderTrendOverviewText(env, data, trendOptions{comparison: engine.ComparisonBaseline}))

				// s3 (~89ns) vs s1 (~101ns) → ratio ≈0.88 → ↓ (improvement); no ↑ expected
				out := env.Out.String()
				assert.Contains(t, out, "↓")
				assert.NotContains(t, out, "↑")
			})
		})

		t.Run("detail", func(t *testing.T) {
			t.Run("sequential", func(t *testing.T) {
				repo := setupTrendStorage(t)
				env := execenv.NewTestEnv(t.Context(), repo)

				sessions, err := engine.LocateSessions(repo.Storage())
				require.NoError(t, err)
				data, err := engine.LoadTrendData(sessions, 0.95)
				require.NoError(t, err)

				require.NoError(t, renderTrendDetailText(env, data, trendOptions{bench: "BenchmarkFoo", comparison: engine.ComparisonSequential}))

				// s3 (~89ns) vs s2 (~81ns) → ratio ≈1.10 → ↑ (regression)
				assert.Contains(t, env.Out.String(), "↑")
			})

			t.Run("baseline", func(t *testing.T) {
				repo := setupTrendStorage(t)
				env := execenv.NewTestEnv(t.Context(), repo)

				sessions, err := engine.LocateSessions(repo.Storage())
				require.NoError(t, err)
				data, err := engine.LoadTrendData(sessions, 0.95)
				require.NoError(t, err)

				require.NoError(t, renderTrendDetailText(env, data, trendOptions{bench: "BenchmarkFoo", comparison: engine.ComparisonBaseline}))

				// s3 (~89ns) vs s1 (~101ns) → ratio ≈0.88 → ↓ (improvement); no ↑ expected
				out := env.Out.String()
				assert.Contains(t, out, "↓")
				assert.NotContains(t, out, "↑")
			})
		})
	})

	t.Run("bench flag", func(t *testing.T) {
		setup := func(t *testing.T) *repository.Repository {
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
			return repository.NewForTesting(memfs.New(), storage, nil)
		}

		t.Run("json_filters_to_matching", func(t *testing.T) {
			env := execenv.NewTestEnv(t.Context(), setup(t))
			env.Format = execenv.FormatJSON

			require.NoError(t, runTrend(env, trendOptions{bench: "Foo", confidence: 0.95}))

			out := env.Out.String()
			assert.Contains(t, out, `"BenchmarkFoo"`)
			assert.NotContains(t, out, `"BenchmarkBar"`)
		})

		t.Run("json_multi_match_all_included", func(t *testing.T) {
			env := execenv.NewTestEnv(t.Context(), setup(t))
			env.Format = execenv.FormatJSON

			require.NoError(t, runTrend(env, trendOptions{bench: "Benchmark", confidence: 0.95}))

			out := env.Out.String()
			assert.Contains(t, out, `"BenchmarkBar"`)
			assert.Contains(t, out, `"BenchmarkFoo"`)
		})

		t.Run("json_no_match_error", func(t *testing.T) {
			env := execenv.NewTestEnv(t.Context(), setup(t))
			env.Format = execenv.FormatJSON

			err := runTrend(env, trendOptions{bench: "Zzz", confidence: 0.95})
			require.ErrorContains(t, err, "no benchmarks match")
		})

		t.Run("text_opens_detail_for_first_match", func(t *testing.T) {
			// BenchNames are sorted, so BenchmarkBar < BenchmarkFoo: first match for "Benchmark" is BenchmarkBar.
			env := execenv.NewTestEnv(t.Context(), setup(t))

			require.NoError(t, runTrend(env, trendOptions{bench: "Benchmark", confidence: 0.95}))

			out := env.Out.String()
			assert.Contains(t, out, "BenchmarkBar")
			assert.NotContains(t, out, "BenchmarkFoo")
		})

		t.Run("text_no_match_error", func(t *testing.T) {
			env := execenv.NewTestEnv(t.Context(), setup(t))

			err := runTrend(env, trendOptions{bench: "Zzz", confidence: 0.95})
			require.ErrorContains(t, err, "no benchmarks match")
		})

		t.Run("invalid_regex_error", func(t *testing.T) {
			env := execenv.NewTestEnv(t.Context(), setup(t))

			err := runTrend(env, trendOptions{bench: "[", confidence: 0.95})
			require.ErrorContains(t, err, "invalid --bench pattern")
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

		require.NoError(t, renderTrendJSON(env, data, nil))
		assert.Contains(t, env.Out.String(), `"BenchmarkFoo"`)
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
