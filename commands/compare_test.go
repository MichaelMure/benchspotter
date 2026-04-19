package commands

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-billy/v5/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"benchspotter/commands/execenv"
	"benchspotter/repository"
)

func TestCompareStat(t *testing.T) {
	// benchfmt.Files reads from real disk paths via BenchFullPath(), so storage
	// must be backed by a real directory rather than memfs.
	storage := osfs.New(t.TempDir())

	id1 := createTestSession(t, storage, "before", []string{"BenchmarkFoo"}, "abc1234def5678abc1234def5678abc1234def56", false)
	id2 := createTestSession(t, storage, "after", []string{"BenchmarkFoo"}, "xyz9876fed5432xyz9876fed5432xyz9876fed54", false)

	writeBenchResults(t, storage, id1, "BenchmarkFoo-8\t1000000\t1200 ns/op\t64 B/op\t2 allocs/op\n")
	writeBenchResults(t, storage, id2, "BenchmarkFoo-8\t1000000\t1000 ns/op\t64 B/op\t2 allocs/op\n")

	env := execenv.NewTestEnv(repository.New(memfs.New(), storage))

	err := runCompareStat(context.Background(), env, compareStatOptions{
		sessions:   []string{id1, id2},
		table:      ".config",
		row:        ".fullname",
		col:        ".file",
		filter:     "*",
		confidence: 0.95,
	})
	require.NoError(t, err)

	out := env.Out.String()
	assert.Contains(t, out, "Foo")
	assert.Contains(t, out, "before")
	assert.Contains(t, out, "after")
}

func writeBenchResults(t *testing.T, storage billy.Filesystem, id, content string) {
	t.Helper()
	err := util.WriteFile(storage, filepath.Join("sessions", id, "results.bench"), []byte(content), 0644)
	require.NoError(t, err)
}
