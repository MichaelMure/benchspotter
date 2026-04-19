package engine

import (
	"bytes"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/util"
	"github.com/google/pprof/profile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testSessionPath = "sessions/test-id"

func makeTestProfile(t *testing.T, fs billy.Filesystem, sessionPath, profileType, pkg, benchName string, prof *profile.Profile) {
	t.Helper()
	filename := replacer.Replace(pkg+"."+benchName) + ".profile"
	dir := filepath.Join(sessionPath, profileType)
	require.NoError(t, fs.MkdirAll(dir, 0755))
	var buf bytes.Buffer
	require.NoError(t, prof.Write(&buf))
	require.NoError(t, util.WriteFile(fs, filepath.Join(dir, filename), buf.Bytes(), 0644))
}

func simpleCPUProfile(funcNames []string) *profile.Profile {
	funcs := make([]*profile.Function, len(funcNames))
	for i, name := range funcNames {
		funcs[i] = &profile.Function{ID: uint64(i + 1), Name: name}
	}
	locs := make([]*profile.Location, len(funcNames))
	for i, fn := range funcs {
		locs[i] = &profile.Location{ID: uint64(i + 1), Line: []profile.Line{{Function: fn}}}
	}
	return &profile.Profile{
		SampleType:    []*profile.ValueType{{Type: "samples", Unit: "count"}, {Type: "cpu", Unit: "nanoseconds"}},
		Function:      funcs,
		Location:      locs,
		Sample:        []*profile.Sample{{Location: locs, Value: []int64{1, int64(time.Millisecond)}}},
		TimeNanos:     time.Now().UnixNano(),
		DurationNanos: int64(100 * time.Millisecond),
		Period:        int64(10 * time.Millisecond),
		PeriodType:    &profile.ValueType{Type: "cpu", Unit: "nanoseconds"},
	}
}

func TestListProfileBenchmarks(t *testing.T) {
	fs := memfs.New()
	makeTestProfile(t, fs, testSessionPath, "cpu", ".", "BenchmarkFoo", simpleCPUProfile([]string{"pkg.Func"}))
	makeTestProfile(t, fs, testSessionPath, "cpu", ".", "BenchmarkBar", simpleCPUProfile([]string{"pkg.Func"}))

	names, err := ListProfileBenchmarks(fs, testSessionPath, ProfileCPU)
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"BenchmarkFoo", "BenchmarkBar"}, names)
}

func TestListProfileBenchmarks_noDir(t *testing.T) {
	fs := memfs.New()
	_, err := ListProfileBenchmarks(fs, testSessionPath, ProfileCPU)
	assert.ErrorContains(t, err, "no cpu profile")
}

func TestReadProfileFunctions_flatVsCumulative(t *testing.T) {
	// Stack: Caller → Hot (Hot is at the top, so flat goes to Hot only)
	hot := &profile.Function{ID: 1, Name: "pkg.Hot"}
	caller := &profile.Function{ID: 2, Name: "pkg.Caller"}
	locHot := &profile.Location{ID: 1, Line: []profile.Line{{Function: hot}}}
	locCaller := &profile.Location{ID: 2, Line: []profile.Line{{Function: caller}}}

	prof := &profile.Profile{
		SampleType: []*profile.ValueType{
			{Type: "samples", Unit: "count"},
			{Type: "cpu", Unit: "nanoseconds"},
		},
		Function: []*profile.Function{hot, caller},
		Location: []*profile.Location{locHot, locCaller},
		Sample: []*profile.Sample{
			// Hot at top of stack, Caller below
			{Location: []*profile.Location{locHot, locCaller}, Value: []int64{10, int64(10 * time.Millisecond)}},
		},
		TimeNanos:     time.Now().UnixNano(),
		DurationNanos: int64(100 * time.Millisecond),
		Period:        int64(time.Millisecond),
		PeriodType:    &profile.ValueType{Type: "cpu", Unit: "nanoseconds"},
	}

	fs := memfs.New()
	makeTestProfile(t, fs, testSessionPath, "cpu", ".", "BenchmarkFoo", prof)

	funcs, err := ReadProfileFunctions(fs, testSessionPath, ProfileCPU, "BenchmarkFoo")
	require.NoError(t, err)
	require.Len(t, funcs, 2)

	byName := make(map[string]ProfileFunc)
	for _, f := range funcs {
		byName[f.Name] = f
	}

	// Hot is flat (top of stack) and cumulative
	assert.Equal(t, time.Duration(10*time.Millisecond), byName["pkg.Hot"].Flat)
	assert.Equal(t, time.Duration(10*time.Millisecond), byName["pkg.Hot"].Cumulative)
	assert.InDelta(t, 100.0, byName["pkg.Hot"].FlatPct, 0.01)

	// Caller is cumulative only (not at top of stack)
	assert.Equal(t, time.Duration(0), byName["pkg.Caller"].Flat)
	assert.Equal(t, time.Duration(10*time.Millisecond), byName["pkg.Caller"].Cumulative)
	assert.Equal(t, 0.0, byName["pkg.Caller"].FlatPct)
}

func TestReadProfileFunctions_notFound(t *testing.T) {
	fs := memfs.New()
	makeTestProfile(t, fs, testSessionPath, "cpu", ".", "BenchmarkFoo", simpleCPUProfile([]string{"pkg.Func"}))

	_, err := ReadProfileFunctions(fs, testSessionPath, ProfileCPU, "BenchmarkBar")
	assert.ErrorContains(t, err, `no cpu profile found for benchmark "BenchmarkBar"`)
}

func TestReadProfileRaw_validPprof(t *testing.T) {
	fs := memfs.New()
	makeTestProfile(t, fs, testSessionPath, "cpu", ".", "BenchmarkFoo", simpleCPUProfile([]string{"pkg.Func"}))

	data, err := ReadProfileRaw(fs, testSessionPath, ProfileCPU, "BenchmarkFoo")
	require.NoError(t, err)

	_, err = profile.ParseData(data)
	assert.NoError(t, err)
}

func TestPickValueIndex(t *testing.T) {
	cases := []struct {
		profileType Profile
		sampleTypes []*profile.ValueType
		wantIdx     int
	}{
		{
			ProfileCPU,
			[]*profile.ValueType{{Type: "samples", Unit: "count"}, {Type: "cpu", Unit: "nanoseconds"}},
			1,
		},
		{
			ProfileMem,
			// real heap profile types from runtime/pprof
			[]*profile.ValueType{
				{Type: "alloc_objects", Unit: "count"},
				{Type: "alloc_space", Unit: "bytes"},
				{Type: "inuse_objects", Unit: "count"},
				{Type: "inuse_space", Unit: "bytes"},
			},
			3,
		},
		{
			ProfileBlock,
			// real block profile types from runtime/pprof
			[]*profile.ValueType{{Type: "contentions", Unit: "count"}, {Type: "delay", Unit: "nanoseconds"}},
			1,
		},
		{
			ProfileCPU,
			[]*profile.ValueType{{Type: "samples", Unit: "count"}}, // no nanoseconds → fallback
			0,
		},
	}

	for _, tc := range cases {
		prof := &profile.Profile{SampleType: tc.sampleTypes}
		assert.Equal(t, tc.wantIdx, pickValueIndex(prof, tc.profileType))
	}
}
