package commands

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/util"
	"github.com/google/pprof/profile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
	"benchspotter/repository"
)

// profileFilename mirrors the encoding used by engine.RunProfile.
var profileReplacer = strings.NewReplacer("/", "ᚋ", ".", "ᚗ", "-", "ᚑ", "~", "א")

func profileFilename(pkg, name string) string {
	return profileReplacer.Replace(pkg+"."+name) + ".profile"
}

func writeProfile(t *testing.T, storage billy.Filesystem, sessionID, profileType, pkg, benchName string, prof *profile.Profile) {
	t.Helper()
	dir := filepath.Join("sessions", sessionID, profileType)
	require.NoError(t, storage.MkdirAll(dir, 0755))
	var buf bytes.Buffer
	require.NoError(t, prof.Write(&buf))
	require.NoError(t, util.WriteFile(storage, filepath.Join(dir, profileFilename(pkg, benchName)), buf.Bytes(), 0644))
}

func setupProfileSession(t *testing.T) (billy.Filesystem, string) {
	t.Helper()
	storage := memfs.New()
	id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "abc1234", false)
	prof := makeCPUProfile([]string{"runtime.mallocgc", "mypackage.HotFunc", "mypackage.CallerFunc"})
	writeProfile(t, storage, id, "cpu", ".", "BenchmarkFoo", prof)
	return storage, id
}

func setupMultiBenchProfileSession(t *testing.T) (billy.Filesystem, string) {
	t.Helper()
	storage := memfs.New()
	id := createTestSession(t, storage, "multi-bench", []string{"BenchmarkFoo", "BenchmarkBar"}, "abc1234", false)
	fooProf := makeCPUProfile([]string{"mypackage.FooFunc"})
	barProf := makeCPUProfile([]string{"mypackage.BarFunc"})
	writeProfile(t, storage, id, "cpu", ".", "BenchmarkFoo", fooProf)
	writeProfile(t, storage, id, "cpu", ".", "BenchmarkBar", barProf)
	return storage, id
}

func makeCPUProfile(funcNames []string) *profile.Profile {
	funcs := make([]*profile.Function, len(funcNames))
	for i, name := range funcNames {
		funcs[i] = &profile.Function{ID: uint64(i + 1), Name: name}
	}
	locs := make([]*profile.Location, len(funcNames))
	for i, fn := range funcs {
		locs[i] = &profile.Location{
			ID:   uint64(i + 1),
			Line: []profile.Line{{Function: fn}},
		}
	}

	p := &profile.Profile{
		SampleType: []*profile.ValueType{
			{Type: "samples", Unit: "count"},
			{Type: "cpu", Unit: "nanoseconds"},
		},
		Function: funcs,
		Location: locs,
		Sample: []*profile.Sample{
			{Location: locs, Value: []int64{10, int64(10 * time.Millisecond)}},
		},
		TimeNanos:     time.Now().UnixNano(),
		DurationNanos: int64(100 * time.Millisecond),
		Period:        int64(10 * time.Millisecond),
		PeriodType:    &profile.ValueType{Type: "cpu", Unit: "nanoseconds"},
	}
	return p
}

func TestShowProfile(t *testing.T) {
	t.Run("cpu", func(t *testing.T) {
		t.Run("text", func(t *testing.T) {
			storage, id := setupProfileSession(t)
			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))

			err := runShowProfile(t.Context(), env, showProfileOptions{
				session:     id,
				profileType: engine.ProfileCPU,
				top:         20,
			})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, "FLAT")
			assert.Contains(t, out, "FUNCTION")
			assert.Contains(t, out, "HotFunc")
			assert.Contains(t, out, "mallocgc")
		})

		t.Run("json", func(t *testing.T) {
			storage, id := setupProfileSession(t)
			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON

			err := runShowProfile(t.Context(), env, showProfileOptions{
				session:     id,
				profileType: engine.ProfileCPU,
				top:         20,
			})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, `"name"`)
			assert.Contains(t, out, `"flat_ns"`)
			assert.Contains(t, out, `"flat_pct"`)
			assert.Contains(t, out, "HotFunc")
		})

		t.Run("raw", func(t *testing.T) {
			storage, id := setupProfileSession(t)
			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatRaw

			err := runShowProfile(t.Context(), env, showProfileOptions{
				session:     id,
				profileType: engine.ProfileCPU,
				top:         20,
			})
			require.NoError(t, err)

			_, err = profile.ParseData(env.Out.Bytes())
			assert.NoError(t, err, "raw output should be a valid pprof profile")
		})

		t.Run("no_profile", func(t *testing.T) {
			storage := memfs.New()
			id := createTestSession(t, storage, "no-profiles", []string{"BenchmarkFoo"}, "abc1234", false)
			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))

			err := runShowProfile(t.Context(), env, showProfileOptions{
				session:     id,
				profileType: engine.ProfileCPU,
				top:         20,
			})
			assert.ErrorContains(t, err, "no cpu profile")
		})

		t.Run("bench_filter", func(t *testing.T) {
			storage, id := setupMultiBenchProfileSession(t)
			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON

			err := runShowProfile(t.Context(), env, showProfileOptions{
				session:     id,
				bench:       "BenchmarkFoo",
				profileType: engine.ProfileCPU,
				top:         20,
			})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, "FooFunc")
			assert.NotContains(t, out, "BarFunc")
		})

		t.Run("multi_bench_requires_bench", func(t *testing.T) {
			storage, id := setupMultiBenchProfileSession(t)
			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON

			err := runShowProfile(t.Context(), env, showProfileOptions{
				session:     id,
				bench:       "",
				profileType: engine.ProfileCPU,
				top:         20,
			})
			assert.ErrorContains(t, err, "specify --bench")
		})

		t.Run("bench_unknown", func(t *testing.T) {
			storage, id := setupMultiBenchProfileSession(t)
			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))

			err := runShowProfile(t.Context(), env, showProfileOptions{
				session:     id,
				bench:       "BenchmarkNonExistent",
				profileType: engine.ProfileCPU,
				top:         20,
			})
			assert.ErrorContains(t, err, "no cpu profile found for benchmark")
		})
	})
}
