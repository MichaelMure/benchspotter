package engine

import (
	"path/filepath"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStripGOMAXPROCS(t *testing.T) {
	cases := []struct{ in, want string }{
		{"BenchmarkFoo-8", "BenchmarkFoo"},
		{"BenchmarkFoo-16", "BenchmarkFoo"},
		{"BenchmarkFoo/N=100-8", "BenchmarkFoo/N=100"},
		{"BenchmarkFoo", "BenchmarkFoo"},
		{"BenchmarkFoo-", "BenchmarkFoo-"},
		{"BenchmarkFoo-abc", "BenchmarkFoo-abc"},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.want, stripGOMAXPROCS(tc.in), "input %q", tc.in)
	}
}

func TestLoadTrendData(t *testing.T) {
	s1 := makeTrendSession(t, "session-1", `BenchmarkFoo-8   1000000   100 ns/op   64 B/op   2 allocs/op
BenchmarkFoo-8   1000000   102 ns/op   64 B/op   2 allocs/op
`)
	s2 := makeTrendSession(t, "session-2", `BenchmarkFoo-8   1000000   120 ns/op   64 B/op   2 allocs/op
BenchmarkFoo-8   1000000   118 ns/op   64 B/op   2 allocs/op
`)
	s3 := makeTrendSession(t, "session-3", `BenchmarkFoo-8   1000000    90 ns/op   64 B/op   2 allocs/op
BenchmarkBar-8   1000000   200 ns/op
`)

	data, err := LoadTrendData([]*SessionInfo{s1, s2, s3}, 0.95)
	require.NoError(t, err)

	assert.Equal(t, []string{"BenchmarkBar", "BenchmarkFoo"}, data.BenchNames)
	assert.Equal(t, []string{"ns/op", "B/op", "allocs/op"}, data.Units)

	fooNs := data.Points["BenchmarkFoo"]["ns/op"]
	require.Len(t, fooNs, 3)
	assert.InDelta(t, 101.0, fooNs[0].Center, 5.0) // median of 100, 102
	assert.InDelta(t, 119.0, fooNs[1].Center, 5.0) // median of 118, 120
	assert.InDelta(t, 90.0, fooNs[2].Center, 1.0)  // single value
	assert.Equal(t, 2, fooNs[0].N)
	assert.Equal(t, 1, fooNs[2].N)

	barNs := data.Points["BenchmarkBar"]["ns/op"]
	require.Len(t, barNs, 1)
	assert.InDelta(t, 200.0, barNs[0].Center, 1.0)

	assert.Empty(t, data.Points["BenchmarkBar"]["B/op"])
}

func TestLoadTrendData_noResults(t *testing.T) {
	s := makeTrendSession(t, "empty", "")
	data, err := LoadTrendData([]*SessionInfo{s}, 0.95)
	require.NoError(t, err)
	assert.Empty(t, data.BenchNames)
}

func TestPreferredUnitsFirst(t *testing.T) {
	units := []string{"custom/op", "allocs/op", "ns/op", "B/op"}
	preferredUnitsFirst(units)
	assert.Equal(t, []string{"ns/op", "B/op", "allocs/op", "custom/op"}, units)
}

// makeTrendSession creates a minimal SessionInfo backed by its own memfs storage.
func makeTrendSession(t *testing.T, name, content string) *SessionInfo {
	t.Helper()
	storage := memfs.New()

	uid, err := uuid.NewV7()
	require.NoError(t, err)
	id := uid.String()
	dir := filepath.Join(sessionDir, id)
	require.NoError(t, storage.MkdirAll(dir, 0755))
	require.NoError(t, util.WriteFile(storage, filepath.Join(dir, metaFilename),
		[]byte(`{"name":"`+name+`"}`), 0644))
	if content != "" {
		require.NoError(t, util.WriteFile(storage, filepath.Join(dir, benchFilename),
			[]byte(content), 0644))
	}

	sessions, err := LocateSessions(storage)
	require.NoError(t, err)
	require.Len(t, sessions, 1)
	return sessions[0]
}
