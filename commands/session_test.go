package commands

import (
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

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))

	err := runSessionLs(t.Context(), env, "")
	require.NoError(t, err)

	out := env.Out.String()
	assert.Contains(t, out, "first-session")
	assert.Contains(t, out, "second-session")
	assert.Contains(t, out, "abc1234")
	assert.Contains(t, out, "xyz9876±")    // has git diff
	assert.NotContains(t, out, "abc1234±") // no diff, no marker
	// benchmark count, not names
	assert.Contains(t, out, "2") // first-session has 2 benchmarks
	assert.Contains(t, out, "1") // second-session has 1
}

func TestSessionLs_profileIndicators(t *testing.T) {
	storage := memfs.New()

	id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)
	// add a cpu profile directory with a file
	dir := filepath.Join("sessions", id, "cpu")
	require.NoError(t, storage.MkdirAll(dir, 0755))
	f, err := storage.Create(filepath.Join(dir, "ᚗᚗBenchmarkFoo.profile"))
	require.NoError(t, err)
	_ = f.Close()

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	require.NoError(t, runSessionLs(t.Context(), env, ""))

	out := env.Out.String()
	assert.Contains(t, out, "✓") // cpu present
}

func TestSessionLs_tags(t *testing.T) {
	storage := memfs.New()

	id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)
	addTagToSession(t, storage, id, "baseline")

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	require.NoError(t, runSessionLs(t.Context(), env, ""))

	assert.Contains(t, env.Out.String(), "baseline")
}

func TestSessionLs_filterByTag(t *testing.T) {
	storage := memfs.New()

	id1 := createTestSession(t, storage, "tagged", []string{"BenchmarkFoo"}, "", false)
	createTestSession(t, storage, "untagged", []string{"BenchmarkBar"}, "", false)
	addTagToSession(t, storage, id1, "baseline")

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	require.NoError(t, runSessionLs(t.Context(), env, "baseline"))

	out := env.Out.String()
	assert.Contains(t, out, "tagged")
	assert.NotContains(t, out, "untagged")
}

func TestSessionLsJSON(t *testing.T) {
	storage := memfs.New()
	id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "abc1234def5678abc1234def5678abc1234def56", true)
	addTagToSession(t, storage, id, "baseline")

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	env.Format = execenv.FormatJSON

	err := runSessionLs(t.Context(), env, "")
	require.NoError(t, err)

	out := env.Out.String()
	assert.Contains(t, out, `"human_name"`)
	assert.Contains(t, out, `"my-session"`)
	assert.Contains(t, out, `"has_diff": true`)
	assert.Contains(t, out, `"BenchmarkFoo"`)
	assert.Contains(t, out, `"abc1234def5678abc1234def5678abc1234def56"`)
	assert.Contains(t, out, `"baseline"`)
}

func TestSessionTag(t *testing.T) {
	storage := memfs.New()
	id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	require.NoError(t, runSessionTag(t.Context(), env, []string{id, "baseline"}))

	meta := readSessionMeta(t, storage, id)
	assert.Equal(t, []interface{}{"baseline"}, meta["tags"])
}

func TestSessionUntag(t *testing.T) {
	storage := memfs.New()
	id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)
	addTagToSession(t, storage, id, "baseline")
	addTagToSession(t, storage, id, "fast")

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	require.NoError(t, runSessionUntag(t.Context(), env, []string{id, "baseline"}))

	meta := readSessionMeta(t, storage, id)
	assert.Equal(t, []interface{}{"fast"}, meta["tags"])
}

func TestSessionRename(t *testing.T) {
	storage := memfs.New()
	id := createTestSession(t, storage, "old-name", []string{"BenchmarkFoo"}, "", false)

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	require.NoError(t, runSessionRename(t.Context(), env, []string{id, "new-name"}))

	meta := readSessionMeta(t, storage, id)
	assert.Equal(t, "new-name", meta["name"])
}

func TestSessionRm(t *testing.T) {
	storage := memfs.New()
	id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	require.NoError(t, runSessionRm(t.Context(), env, []string{id}, true))

	_, err := storage.Stat(filepath.Join("sessions", id))
	assert.Error(t, err)
}

func TestSessionRm_notFound(t *testing.T) {
	storage := memfs.New()
	createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	err := runSessionRm(t.Context(), env, []string{"00000000-0000-0000-0000-000000000000"}, true)
	assert.ErrorContains(t, err, "not found")
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

func addTagToSession(t *testing.T, storage billy.Filesystem, id, tag string) {
	t.Helper()
	metaPath := filepath.Join("sessions", id, "meta.json")
	f, err := storage.Open(metaPath)
	require.NoError(t, err)
	var meta map[string]interface{}
	require.NoError(t, json.NewDecoder(f).Decode(&meta))
	_ = f.Close()

	tags, _ := meta["tags"].([]interface{})
	meta["tags"] = append(tags, tag)

	out, err := storage.Create(metaPath)
	require.NoError(t, err)
	require.NoError(t, json.NewEncoder(out).Encode(meta))
	_ = out.Close()
}

func readSessionMeta(t *testing.T, storage billy.Filesystem, id string) map[string]interface{} {
	t.Helper()
	f, err := storage.Open(filepath.Join("sessions", id, "meta.json"))
	require.NoError(t, err)
	defer f.Close()
	var meta map[string]interface{}
	require.NoError(t, json.NewDecoder(f).Decode(&meta))
	return meta
}
