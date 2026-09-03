package commands

import (
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/google/pprof/profile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/MichaelMure/benchspotter/commands/execenv"
	"github.com/MichaelMure/benchspotter/engine"
	"github.com/MichaelMure/benchspotter/repository"
)

func TestCompareProfile(t *testing.T) {
	t.Run("json/cpu", func(t *testing.T) {
		storage := memfs.New()
		baseID := createTestSession(t, storage, "session-base", []string{"BenchmarkFoo"}, "abc1234", false)
		newID := createTestSession(t, storage, "session-new", []string{"BenchmarkFoo"}, "xyz9876", false)
		writeProfile(t, storage, baseID, "cpu", ".", "BenchmarkFoo",
			makeCPUProfile([]string{"mypackage.HotFunc", "mypackage.OtherFunc"}))
		writeProfile(t, storage, newID, "cpu", ".", "BenchmarkFoo",
			makeCPUProfile([]string{"mypackage.HotFunc"}))

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		err := runCompareProfile(env, compareProfileOptions{
			baseSession: baseID,
			newSession:  newID,
			bench:       "BenchmarkFoo",
			profileType: engine.ProfileCPU,
			sort:        engine.DiffSortAbsFlat,
			top:         20,
		})
		require.NoError(t, err)

		out := env.Out.String()
		assert.Contains(t, out, `"name"`)
		assert.Contains(t, out, "HotFunc")
	})

	t.Run("json/auto_select_single_bench", func(t *testing.T) {
		storage := memfs.New()
		baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
		newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)
		writeProfile(t, storage, baseID, "cpu", ".", "BenchmarkFoo",
			makeCPUProfile([]string{"mypackage.Func"}))
		writeProfile(t, storage, newID, "cpu", ".", "BenchmarkFoo",
			makeCPUProfile([]string{"mypackage.Func"}))

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		// bench not specified; single common → auto-selected
		err := runCompareProfile(env, compareProfileOptions{
			baseSession: baseID,
			newSession:  newID,
			profileType: engine.ProfileCPU,
			sort:        engine.DiffSortAbsFlat,
			top:         20,
		})
		require.NoError(t, err)
		assert.Contains(t, env.Out.String(), `"name"`)
	})

	t.Run("raw/cpu_produces_valid_diff_pprof", func(t *testing.T) {
		storage := memfs.New()
		baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
		newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)
		writeProfile(t, storage, baseID, "cpu", ".", "BenchmarkFoo",
			makeCPUProfile([]string{"mypackage.BaseOnly", "mypackage.Both"}))
		writeProfile(t, storage, newID, "cpu", ".", "BenchmarkFoo",
			makeCPUProfile([]string{"mypackage.NewOnly", "mypackage.Both"}))

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatRaw

		err := runCompareProfile(env, compareProfileOptions{
			baseSession: baseID,
			newSession:  newID,
			bench:       "BenchmarkFoo",
			profileType: engine.ProfileCPU,
			sort:        engine.DiffSortAbsFlat,
			top:         20,
		})
		require.NoError(t, err)

		// Output must be a valid parseable pprof binary.
		data := []byte(env.Out.String())
		prof, err := profile.ParseData(data)
		require.NoError(t, err)

		// Collect all function names in the diff profile.
		names := make(map[string]bool)
		for _, fn := range prof.Function {
			names[fn.Name] = true
		}
		assert.True(t, names["mypackage.BaseOnly"] || names["mypackage.NewOnly"],
			"diff pprof should contain functions from at least one session")
	})

	t.Run("text/cpu", func(t *testing.T) {
		storage := memfs.New()
		baseID := createTestSession(t, storage, "session-base", []string{"BenchmarkFoo"}, "", false)
		newID := createTestSession(t, storage, "session-new", []string{"BenchmarkFoo"}, "", false)
		writeProfile(t, storage, baseID, "cpu", ".", "BenchmarkFoo",
			makeCPUProfile([]string{"mypackage.HotFunc"}))
		writeProfile(t, storage, newID, "cpu", ".", "BenchmarkFoo",
			makeCPUProfile([]string{"mypackage.HotFunc"}))

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatText

		err := runCompareProfile(env, compareProfileOptions{
			baseSession: baseID,
			newSession:  newID,
			bench:       "BenchmarkFoo",
			profileType: engine.ProfileCPU,
			sort:        engine.DiffSortAbsFlat,
			top:         20,
		})
		require.NoError(t, err)

		out := env.Out.String()
		assert.Contains(t, out, "session-base")
		assert.Contains(t, out, "session-new")
	})

	t.Run("multiple_common_benches_no_bench_flag", func(t *testing.T) {
		storage := memfs.New()
		baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo", "BenchmarkBar"}, "", false)
		newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo", "BenchmarkBar"}, "", false)
		for _, id := range []string{baseID, newID} {
			writeProfile(t, storage, id, "cpu", ".", "BenchmarkFoo",
				makeCPUProfile([]string{"mypackage.Func"}))
			writeProfile(t, storage, id, "cpu", ".", "BenchmarkBar",
				makeCPUProfile([]string{"mypackage.Func"}))
		}

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		err := runCompareProfile(env, compareProfileOptions{
			baseSession: baseID,
			newSession:  newID,
			profileType: engine.ProfileCPU,
			sort:        engine.DiffSortAbsFlat,
			top:         20,
		})
		assert.ErrorContains(t, err, "multiple common")
	})

	t.Run("no_common_benches", func(t *testing.T) {
		storage := memfs.New()
		baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
		newID := createTestSession(t, storage, "new", []string{"BenchmarkBar"}, "", false)
		writeProfile(t, storage, baseID, "cpu", ".", "BenchmarkFoo",
			makeCPUProfile([]string{"mypackage.Func"}))
		writeProfile(t, storage, newID, "cpu", ".", "BenchmarkBar",
			makeCPUProfile([]string{"mypackage.Func"}))

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		err := runCompareProfile(env, compareProfileOptions{
			baseSession: baseID,
			newSession:  newID,
			profileType: engine.ProfileCPU,
			sort:        engine.DiffSortAbsFlat,
			top:         20,
		})
		assert.ErrorContains(t, err, "no common")
	})

	t.Run("session_without_cpu_profile", func(t *testing.T) {
		storage := memfs.New()
		baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
		newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)
		writeProfile(t, storage, baseID, "cpu", ".", "BenchmarkFoo",
			makeCPUProfile([]string{"mypackage.Func"}))
		// newID has no profile

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		err := runCompareProfile(env, compareProfileOptions{
			baseSession: baseID,
			newSession:  newID,
			bench:       "BenchmarkFoo",
			profileType: engine.ProfileCPU,
		})
		assert.ErrorContains(t, err, "no cpu profile")
	})

	t.Run("session_not_found", func(t *testing.T) {
		storage := memfs.New()
		newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)
		writeProfile(t, storage, newID, "cpu", ".", "BenchmarkFoo",
			makeCPUProfile([]string{"mypackage.Func"}))

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		err := runCompareProfile(env, compareProfileOptions{
			baseSession: "nonexistent",
			newSession:  newID,
			profileType: engine.ProfileCPU,
		})
		assert.ErrorContains(t, err, `"nonexistent" not found`)
	})
}
