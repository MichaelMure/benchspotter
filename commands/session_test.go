package commands

import (
	"context"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"benchspotter/commands/execenv"
	"benchspotter/repository"
)

func TestSessionLs(t *testing.T) {
	storage := memfs.New()

	createTestSession(t, storage, "first-session", []string{"BenchmarkFoo", "BenchmarkBar"}, "abc1234def5678abc1234def5678abc1234def56", false)
	createTestSession(t, storage, "second-session", []string{"BenchmarkBaz"}, "xyz9876fed5432xyz9876fed5432xyz9876fed54", true)

	env := execenv.NewTestEnv(repository.New(memfs.New(), storage))

	err := runSessionLs(context.Background(), env)
	require.NoError(t, err)

	out := env.Out.String()
	assert.Contains(t, out, "first-session")
	assert.Contains(t, out, "second-session")
	assert.Contains(t, out, "BenchmarkFoo, BenchmarkBar")
	assert.Contains(t, out, "BenchmarkBaz")
	assert.Contains(t, out, "abc1234")
	assert.Contains(t, out, "xyz9876±") // has git diff
	assert.NotContains(t, out, "abc1234±") // no diff, no marker
}

// createTestSession writes a minimal session into the given storage filesystem
// and returns the session ID.
func createTestSession(t *testing.T, storage billy.Filesystem, name string, benches []string, commit string, hasDiff bool) string {
	t.Helper()

	uid, err := uuid.NewV7()
	require.NoError(t, err)
	id := uid.String()

	dir := filepath.Join("sessions", id)
	err = storage.MkdirAll(dir, 0755)
	require.NoError(t, err)

	f, err := storage.Create(filepath.Join(dir, "meta.json"))
	require.NoError(t, err)
	err = json.NewEncoder(f).Encode(map[string]interface{}{
		"name":       name,
		"benchs":     benches,
		"git_commit": commit,
	})
	require.NoError(t, err)
	_ = f.Close()

	if hasDiff {
		f, err = storage.Create(filepath.Join(dir, "git.diff"))
		require.NoError(t, err)
		_, err = f.Write([]byte("diff --git a/foo.go b/foo.go\n+// a change\n"))
		require.NoError(t, err)
		_ = f.Close()
	}

	return id
}
