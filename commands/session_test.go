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

func TestSession(t *testing.T) {
	t.Run("ls", func(t *testing.T) {
		t.Run("text", func(t *testing.T) {
			storage := memfs.New()
			createTestSession(t, storage, "first-session", []string{"BenchmarkFoo", "BenchmarkBar"}, "abc1234def5678abc1234def5678abc1234def56", false)
			createTestSession(t, storage, "second-session", []string{"BenchmarkBaz"}, "xyz9876fed5432xyz9876fed5432xyz9876fed54", true)

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			err := runSessionLs(env, sessionLsOptions{tag: ""})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, "first-session")
			assert.Contains(t, out, "second-session")
			assert.Contains(t, out, "abc1234")
			assert.Contains(t, out, "xyz9876±")    // has git diff
			assert.NotContains(t, out, "abc1234±") // no diff, no marker
			assert.Contains(t, out, "2")           // first-session has 2 benchmarks
			assert.Contains(t, out, "1")           // second-session has 1
		})

		t.Run("profile_indicators", func(t *testing.T) {
			storage := memfs.New()
			id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)
			dir := filepath.Join("sessions", id, "cpu")
			require.NoError(t, storage.MkdirAll(dir, 0755))
			f, err := storage.Create(filepath.Join(dir, "ᚗᚗBenchmarkFoo.profile"))
			require.NoError(t, err)
			_ = f.Close()

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			require.NoError(t, runSessionLs(env, sessionLsOptions{tag: ""}))
			assert.Contains(t, env.Out.String(), "✓")
		})

		t.Run("tags", func(t *testing.T) {
			storage := memfs.New()
			id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)
			addTagToSession(t, storage, id, "baseline")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			require.NoError(t, runSessionLs(env, sessionLsOptions{tag: ""}))
			assert.Contains(t, env.Out.String(), "baseline")
		})

		t.Run("filter_by_tag", func(t *testing.T) {
			storage := memfs.New()
			id1 := createTestSession(t, storage, "tagged", []string{"BenchmarkFoo"}, "", false)
			createTestSession(t, storage, "untagged", []string{"BenchmarkBar"}, "", false)
			addTagToSession(t, storage, id1, "baseline")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			require.NoError(t, runSessionLs(env, sessionLsOptions{tag: "baseline"}))

			out := env.Out.String()
			assert.Contains(t, out, "tagged")
			assert.NotContains(t, out, "untagged")
		})

		t.Run("filter_by_bench", func(t *testing.T) {
			t.Run("exact", func(t *testing.T) {
				storage := memfs.New()
				createTestSession(t, storage, "has-foo", []string{"BenchmarkFoo", "BenchmarkBar"}, "", false)
				createTestSession(t, storage, "has-bar", []string{"BenchmarkBar"}, "", false)
				createTestSession(t, storage, "has-baz", []string{"BenchmarkBaz"}, "", false)

				env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
				require.NoError(t, runSessionLs(env, sessionLsOptions{bench: "BenchmarkFoo"}))

				out := env.Out.String()
				assert.Contains(t, out, "has-foo")
				assert.NotContains(t, out, "has-bar")
				assert.NotContains(t, out, "has-baz")
			})

			t.Run("partial_regex", func(t *testing.T) {
				storage := memfs.New()
				createTestSession(t, storage, "has-foo", []string{"BenchmarkFoo"}, "", false)
				createTestSession(t, storage, "has-bar", []string{"BenchmarkBar"}, "", false)

				env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
				require.NoError(t, runSessionLs(env, sessionLsOptions{bench: "Foo"}))

				out := env.Out.String()
				assert.Contains(t, out, "has-foo")
				assert.NotContains(t, out, "has-bar")
			})

			t.Run("multi_session_match", func(t *testing.T) {
				storage := memfs.New()
				createTestSession(t, storage, "has-foo", []string{"BenchmarkFoo"}, "", false)
				createTestSession(t, storage, "has-bar", []string{"BenchmarkBar"}, "", false)
				createTestSession(t, storage, "has-baz", []string{"BenchmarkBaz"}, "", false)

				env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
				require.NoError(t, runSessionLs(env, sessionLsOptions{bench: "Foo|Bar"}))

				out := env.Out.String()
				assert.Contains(t, out, "has-foo")
				assert.Contains(t, out, "has-bar")
				assert.NotContains(t, out, "has-baz")
			})

			t.Run("no_match", func(t *testing.T) {
				storage := memfs.New()
				createTestSession(t, storage, "has-foo", []string{"BenchmarkFoo"}, "", false)

				env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
				require.NoError(t, runSessionLs(env, sessionLsOptions{bench: "Zzz"}))

				assert.NotContains(t, env.Out.String(), "has-foo")
			})

			t.Run("invalid_regex", func(t *testing.T) {
				storage := memfs.New()
				createTestSession(t, storage, "any", []string{"BenchmarkFoo"}, "", false)

				env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
				err := runSessionLs(env, sessionLsOptions{bench: "["})
				require.ErrorContains(t, err, "invalid --bench pattern")
			})
		})

		t.Run("json", func(t *testing.T) {
			storage := memfs.New()
			id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "abc1234def5678abc1234def5678abc1234def56", true)
			addTagToSession(t, storage, id, "baseline")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON
			err := runSessionLs(env, sessionLsOptions{tag: ""})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, `"human_name"`)
			assert.Contains(t, out, `"my-session"`)
			assert.Contains(t, out, `"has_diff": true`)
			assert.Contains(t, out, `"BenchmarkFoo"`)
			assert.Contains(t, out, `"abc1234def5678abc1234def5678abc1234def56"`)
			assert.Contains(t, out, `"baseline"`)
		})

		t.Run("notes", func(t *testing.T) {
			storage := memfs.New()
			id := createTestSession(t, storage, "noted", []string{"BenchmarkFoo"}, "", false)
			setNoteOnSession(t, storage, id, "this is my note")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			require.NoError(t, runSessionLs(env, sessionLsOptions{}))
			assert.Contains(t, env.Out.String(), "this is my note")
		})

		t.Run("notes_truncated", func(t *testing.T) {
			storage := memfs.New()
			id := createTestSession(t, storage, "noted", []string{"BenchmarkFoo"}, "", false)
			long := "this note is way too long and should be truncated in the table view"
			setNoteOnSession(t, storage, id, long)

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			require.NoError(t, runSessionLs(env, sessionLsOptions{}))
			out := env.Out.String()
			assert.NotContains(t, out, long)
			assert.Contains(t, out, "…")
		})

		t.Run("notes_newlines_collapsed", func(t *testing.T) {
			storage := memfs.New()
			id := createTestSession(t, storage, "noted", []string{"BenchmarkFoo"}, "", false)
			setNoteOnSession(t, storage, id, "line one\nline two")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			require.NoError(t, runSessionLs(env, sessionLsOptions{}))
			out := env.Out.String()
			assert.NotContains(t, out, "\n\n") // no raw newline breaking the table row
			assert.Contains(t, out, "line one line two")
		})

		t.Run("json_notes", func(t *testing.T) {
			storage := memfs.New()
			id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)
			setNoteOnSession(t, storage, id, "important context")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON
			require.NoError(t, runSessionLs(env, sessionLsOptions{}))
			assert.Contains(t, env.Out.String(), `"important context"`)
		})
	})

	t.Run("tag", func(t *testing.T) {
		storage := memfs.New()
		id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		require.NoError(t, runSessionTag(env, []string{id, "baseline"}))

		meta := readSessionMeta(t, storage, id)
		assert.Equal(t, []interface{}{"baseline"}, meta["tags"])
	})

	t.Run("untag", func(t *testing.T) {
		storage := memfs.New()
		id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)
		addTagToSession(t, storage, id, "baseline")
		addTagToSession(t, storage, id, "fast")

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		require.NoError(t, runSessionUntag(env, []string{id, "baseline"}))

		meta := readSessionMeta(t, storage, id)
		assert.Equal(t, []interface{}{"fast"}, meta["tags"])
	})

	t.Run("note", func(t *testing.T) {
		t.Run("set", func(t *testing.T) {
			storage := memfs.New()
			id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			require.NoError(t, runSessionNote(env, []string{id, "my note text"}))

			meta := readSessionMeta(t, storage, id)
			assert.Equal(t, "my note text", meta["notes"])
		})

		t.Run("clear", func(t *testing.T) {
			storage := memfs.New()
			id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)
			setNoteOnSession(t, storage, id, "existing note")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			require.NoError(t, runSessionNote(env, []string{id, ""}))

			meta := readSessionMeta(t, storage, id)
			assert.Empty(t, meta["notes"])
		})
	})

	t.Run("rename", func(t *testing.T) {
		storage := memfs.New()
		id := createTestSession(t, storage, "old-name", []string{"BenchmarkFoo"}, "", false)

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		require.NoError(t, runSessionRename(env, []string{id, "new-name"}))

		meta := readSessionMeta(t, storage, id)
		assert.Equal(t, "new-name", meta["name"])
	})

	t.Run("rm", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			storage := memfs.New()
			id := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			require.NoError(t, runSessionRm(env, []string{id}, sessionRmOptions{skipConfirmation: true}))

			_, err := storage.Stat(filepath.Join("sessions", id))
			assert.Error(t, err)
		})

		t.Run("not_found", func(t *testing.T) {
			storage := memfs.New()
			createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			err := runSessionRm(env, []string{"00000000-0000-0000-0000-000000000000"}, sessionRmOptions{skipConfirmation: true})
			assert.ErrorContains(t, err, "not found")
		})
	})
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

func setNoteOnSession(t *testing.T, storage billy.Filesystem, id, note string) {
	t.Helper()
	metaPath := filepath.Join("sessions", id, "meta.json")
	f, err := storage.Open(metaPath)
	require.NoError(t, err)
	var meta map[string]interface{}
	require.NoError(t, json.NewDecoder(f).Decode(&meta))
	_ = f.Close()

	meta["notes"] = note

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
