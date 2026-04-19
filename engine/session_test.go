package engine

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func writeTestMeta(t *testing.T, fs billy.Filesystem, sessionPath string, meta sessionMeta) {
	t.Helper()
	require.NoError(t, fs.MkdirAll(sessionPath, 0755))
	data, err := json.Marshal(meta)
	require.NoError(t, err)
	require.NoError(t, util.WriteFile(fs, filepath.Join(sessionPath, metaFilename), data, 0644))
}

func readMeta(t *testing.T, fs billy.Filesystem, sessionPath string) sessionMeta {
	t.Helper()
	f, err := fs.Open(filepath.Join(sessionPath, metaFilename))
	require.NoError(t, err)
	defer f.Close()
	var meta sessionMeta
	require.NoError(t, json.NewDecoder(f).Decode(&meta))
	return meta
}

func TestTagSession(t *testing.T) {
	fs := memfs.New()
	path := "sessions/test-id"
	writeTestMeta(t, fs, path, sessionMeta{Name: "my-session"})

	require.NoError(t, TagSession(fs, path, "baseline"))
	require.NoError(t, TagSession(fs, path, "fast"))

	assert.Equal(t, []string{"baseline", "fast"}, readMeta(t, fs, path).Tags)
}

func TestTagSession_idempotent(t *testing.T) {
	fs := memfs.New()
	path := "sessions/test-id"
	writeTestMeta(t, fs, path, sessionMeta{Name: "my-session"})

	require.NoError(t, TagSession(fs, path, "baseline"))
	require.NoError(t, TagSession(fs, path, "baseline"))

	assert.Equal(t, []string{"baseline"}, readMeta(t, fs, path).Tags)
}

func TestUntagSession(t *testing.T) {
	fs := memfs.New()
	path := "sessions/test-id"
	writeTestMeta(t, fs, path, sessionMeta{Name: "my-session", Tags: []string{"baseline", "fast"}})

	require.NoError(t, UntagSession(fs, path, "baseline"))

	assert.Equal(t, []string{"fast"}, readMeta(t, fs, path).Tags)
}

func TestUntagSession_notPresent(t *testing.T) {
	fs := memfs.New()
	path := "sessions/test-id"
	writeTestMeta(t, fs, path, sessionMeta{Name: "my-session", Tags: []string{"fast"}})

	require.NoError(t, UntagSession(fs, path, "missing"))

	assert.Equal(t, []string{"fast"}, readMeta(t, fs, path).Tags)
}

func TestRenameSession(t *testing.T) {
	fs := memfs.New()
	path := "sessions/test-id"
	writeTestMeta(t, fs, path, sessionMeta{Name: "old-name"})

	require.NoError(t, RenameSession(fs, path, "new-name"))

	assert.Equal(t, "new-name", readMeta(t, fs, path).Name)
}

func TestRenameSession_preservesTags(t *testing.T) {
	fs := memfs.New()
	path := "sessions/test-id"
	writeTestMeta(t, fs, path, sessionMeta{Name: "old-name", Tags: []string{"baseline"}})

	require.NoError(t, RenameSession(fs, path, "new-name"))

	meta := readMeta(t, fs, path)
	assert.Equal(t, "new-name", meta.Name)
	assert.Equal(t, []string{"baseline"}, meta.Tags)
}

func TestRemoveSession(t *testing.T) {
	fs := memfs.New()
	path := "sessions/test-id"
	writeTestMeta(t, fs, path, sessionMeta{Name: "my-session"})
	require.NoError(t, fs.MkdirAll(filepath.Join(path, "cpu"), 0755))
	require.NoError(t, util.WriteFile(fs, filepath.Join(path, "cpu", "bench.profile"), []byte("data"), 0644))

	require.NoError(t, RemoveSession(fs, path))

	_, err := fs.Stat(path)
	assert.Error(t, err)
}
