package commands

import (
	"encoding/json"
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

func setupCompareStorage(t *testing.T) (billy.Filesystem, string, string) {
	t.Helper()
	storage := osfs.New(t.TempDir())
	id1 := createTestSession(t, storage, "before", []string{"BenchmarkFoo"}, "abc1234def5678abc1234def5678abc1234def56", false)
	id2 := createTestSession(t, storage, "after", []string{"BenchmarkFoo"}, "xyz9876fed5432xyz9876fed5432xyz9876fed54", false)
	writeBenchResults(t, storage, id1, "BenchmarkFoo-8\t1000000\t1200 ns/op\t64 B/op\t2 allocs/op\n")
	writeBenchResults(t, storage, id2, "BenchmarkFoo-8\t1000000\t1000 ns/op\t64 B/op\t2 allocs/op\n")
	return storage, id1, id2
}

func compareBenchOpts(id1, id2 string) compareBenchOptions {
	return compareBenchOptions{
		sessions:   []string{id1, id2},
		table:      ".config",
		row:        ".fullname",
		col:        ".file",
		filter:     "*",
		confidence: 0.95,
	}
}

func TestCompareBench(t *testing.T) {
	t.Run("text", func(t *testing.T) {
		// benchfmt.Files reads from real disk paths via BenchFullPath(), so storage
		// must be backed by a real directory rather than memfs.
		storage, id1, id2 := setupCompareStorage(t)
		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))

		err := runCompareBench(env, compareBenchOpts(id1, id2))
		require.NoError(t, err)

		out := env.Out.String()
		assert.Contains(t, out, "Foo")
		assert.Contains(t, out, "before")
		assert.Contains(t, out, "after")
	})

	t.Run("json", func(t *testing.T) {
		storage, id1, id2 := setupCompareStorage(t)
		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		err := runCompareBench(env, compareBenchOpts(id1, id2))
		require.NoError(t, err)

		out := env.Out.String()
		assert.Contains(t, out, `"unit"`)
		assert.Contains(t, out, `"benchmarks"`)
		assert.Contains(t, out, `"Foo`)
		assert.Contains(t, out, `"before"`)
		assert.Contains(t, out, `"after"`)
		assert.Contains(t, out, `"center"`)
		assert.Contains(t, out, `"delta"`)
	})

	t.Run("raw", func(t *testing.T) {
		storage, id1, id2 := setupCompareStorage(t)
		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatRaw

		err := runCompareBench(env, compareBenchOpts(id1, id2))
		require.NoError(t, err)

		out := env.Out.String()
		assert.Contains(t, out, "BenchmarkFoo-8")
		assert.Contains(t, out, "before")
		assert.Contains(t, out, "after")
	})

	t.Run("by name", func(t *testing.T) {
		storage := osfs.New(t.TempDir())
		id1 := createTestSession(t, storage, "before", []string{"BenchmarkFoo"}, "", false)
		id2 := createTestSession(t, storage, "after", []string{"BenchmarkFoo"}, "", false)
		writeBenchResults(t, storage, id1, "BenchmarkFoo-8\t1000000\t1200 ns/op\n")
		writeBenchResults(t, storage, id2, "BenchmarkFoo-8\t1000000\t1000 ns/op\n")

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		err := runCompareBench(env, compareBenchOptions{
			sessions:   []string{"before", "after"}, // names, not IDs
			table:      ".config",
			row:        ".fullname",
			col:        ".file",
			filter:     "*",
			confidence: 0.95,
		})
		require.NoError(t, err)

		out := env.Out.String()
		assert.Contains(t, out, `"before"`)
		assert.Contains(t, out, `"after"`)
	})

	t.Run("baseline", func(t *testing.T) {
		storage := osfs.New(t.TempDir())
		id1 := createTestSession(t, storage, "first", []string{"BenchmarkFoo"}, "", false)
		id2 := createTestSession(t, storage, "second", []string{"BenchmarkFoo"}, "", false)
		id3 := createTestSession(t, storage, "third", []string{"BenchmarkFoo"}, "", false)
		writeBenchResults(t, storage, id1, "BenchmarkFoo-8\t1000000\t1200 ns/op\n")
		writeBenchResults(t, storage, id2, "BenchmarkFoo-8\t1000000\t1000 ns/op\n")
		writeBenchResults(t, storage, id3, "BenchmarkFoo-8\t1000000\t900 ns/op\n")

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		// Pass sessions in order first, second, third but designate third as baseline.
		err := runCompareBench(env, compareBenchOptions{
			sessions:   []string{id1, id2, id3},
			baseline:   "third", // name lookup
			table:      ".config",
			row:        ".fullname",
			col:        ".file",
			filter:     "*",
			confidence: 0.95,
		})
		require.NoError(t, err)

		// The first column in the JSON output is the baseline. Verify "third" appears
		// before the others by checking column order in the first benchmark entry.
		var tables []compareBenchJSONTable
		require.NoError(t, json.Unmarshal([]byte(env.Out.String()), &tables))
		require.NotEmpty(t, tables)
		require.NotEmpty(t, tables[0].Benchmarks)

		sessions := tables[0].Benchmarks[0].Sessions
		require.NotEmpty(t, sessions)
		assert.Equal(t, "third", sessions[0].Name)
	})
}

func writeBenchResults(t *testing.T, storage billy.Filesystem, id, content string) {
	t.Helper()
	err := util.WriteFile(storage, filepath.Join("sessions", id, "results.bench"), []byte(content), 0644)
	require.NoError(t, err)
}
