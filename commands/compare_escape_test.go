package commands

import (
	"path/filepath"
	"testing"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
	"benchspotter/repository"
)

func writeEscapeAnalysis(t *testing.T, storage billy.Filesystem, sessionID, data string) {
	t.Helper()
	require.NoError(t, util.WriteFile(storage,
		filepath.Join("sessions", sessionID, engine.EscapeFilename),
		[]byte(data), 0644))
}

func TestCompareEscape(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		t.Run("statuses", func(t *testing.T) {
			storage := memfs.New()
			baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
			newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)

			writeEscapeAnalysis(t, storage, baseID,
				"foo.go:10:5: x escapes to heap\n"+
					"foo.go:20:3: y escapes to heap\n")
			writeEscapeAnalysis(t, storage, newID,
				"foo.go:10:5: x escapes to heap\n"+ // same
					"foo.go:30:2: z escapes to heap\n") // added; y removed

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON

			err := runCompareEscape(env, compareAnalysisOptions{baseSession: baseID, newSession: newID, all: true})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, `"added"`)
			assert.Contains(t, out, `"removed"`)
			assert.Contains(t, out, `"same"`)
			assert.Contains(t, out, "z escapes")
			assert.Contains(t, out, "y escapes")
			assert.Contains(t, out, "x escapes")
		})

		t.Run("heap_escape_field", func(t *testing.T) {
			storage := memfs.New()
			baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
			newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)

			writeEscapeAnalysis(t, storage, baseID, "")
			writeEscapeAnalysis(t, storage, newID, "foo.go:10:5: x escapes to heap\n")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON

			err := runCompareEscape(env, compareAnalysisOptions{baseSession: baseID, newSession: newID})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, `"heap_escape": true`)
		})

		t.Run("base_new_lines_for_same_shifted", func(t *testing.T) {
			storage := memfs.New()
			baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
			newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)

			writeEscapeAnalysis(t, storage, baseID, "foo.go:10:5: x escapes to heap\n")
			writeEscapeAnalysis(t, storage, newID, "foo.go:15:5: x escapes to heap\n")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON

			err := runCompareEscape(env, compareAnalysisOptions{baseSession: baseID, newSession: newID, all: true})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, `"same"`)
			assert.Contains(t, out, `"base_line": 10`)
			assert.Contains(t, out, `"new_line": 15`)
		})
	})

	t.Run("raw", func(t *testing.T) {
		storage := memfs.New()
		baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
		newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)

		writeEscapeAnalysis(t, storage, baseID,
			"foo.go:10:5: x escapes to heap\n"+
				"foo.go:20:3: y escapes to heap\n")
		writeEscapeAnalysis(t, storage, newID,
			"foo.go:10:5: x escapes to heap\n"+ // same — must not appear
				"foo.go:30:2: z escapes to heap\n") // added; y removed

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatRaw

		err := runCompareEscape(env, compareAnalysisOptions{baseSession: baseID, newSession: newID})
		require.NoError(t, err)

		out := env.Out.String()
		assert.Contains(t, out, "+ foo.go:30:2: z escapes to heap")
		assert.Contains(t, out, "- foo.go:20:3: y escapes to heap")
		assert.NotContains(t, out, "x escapes")
	})

	t.Run("text", func(t *testing.T) {
		t.Run("summary_counts", func(t *testing.T) {
			storage := memfs.New()
			baseID := createTestSession(t, storage, "session-base", []string{"BenchmarkFoo"}, "", false)
			newID := createTestSession(t, storage, "session-new", []string{"BenchmarkFoo"}, "", false)

			writeEscapeAnalysis(t, storage, baseID, "foo.go:10:5: x escapes to heap\n")
			writeEscapeAnalysis(t, storage, newID, "foo.go:20:3: y escapes to heap\n")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatText

			err := runCompareEscape(env, compareAnalysisOptions{baseSession: baseID, newSession: newID})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, "1 new, 1 fixed")
			assert.Contains(t, out, "session-base")
			assert.Contains(t, out, "session-new")
		})

		t.Run("no_changes", func(t *testing.T) {
			storage := memfs.New()
			baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
			newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)

			sameData := "foo.go:10:5: x escapes to heap\n"
			writeEscapeAnalysis(t, storage, baseID, sameData)
			writeEscapeAnalysis(t, storage, newID, sameData)

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatText

			err := runCompareEscape(env, compareAnalysisOptions{baseSession: baseID, newSession: newID})
			require.NoError(t, err)
			assert.Contains(t, env.Out.String(), "(no changes)")
		})

		t.Run("source_rendering", func(t *testing.T) {
			const commit = "abc1234def5678abc1234def5678abc1234def56"
			gitSrc := repository.MapGitSource{
				commit + ":foo.go": []byte(`package foo

func BenchmarkFoo(b *testing.B) {
	x := make([]byte, 1024)
	_ = x
}
`),
			}
			storage := memfs.New()
			baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, commit, false)
			newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, commit, false)

			writeEscapeAnalysis(t, storage, baseID, "")
			writeEscapeAnalysis(t, storage, newID, "foo.go:4:7: make([]byte, 1024) escapes to heap\n")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, gitSrc))
			env.Format = execenv.FormatText

			err := runCompareEscape(env, compareAnalysisOptions{baseSession: baseID, newSession: newID})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, "escapes to heap")
		})
	})

	t.Run("session_without_escape_data", func(t *testing.T) {
		storage := memfs.New()
		baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
		newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)
		writeEscapeAnalysis(t, storage, baseID, "foo.go:10:5: x escapes to heap\n")
		// newID has no escape.txt → HasProfile returns false

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		err := runCompareEscape(env, compareAnalysisOptions{baseSession: baseID, newSession: newID})
		assert.ErrorContains(t, err, "no analysis of this type")
	})

	t.Run("session_not_found", func(t *testing.T) {
		storage := memfs.New()
		newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)
		writeEscapeAnalysis(t, storage, newID, "")

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		err := runCompareEscape(env, compareAnalysisOptions{baseSession: "nonexistent", newSession: newID})
		assert.ErrorContains(t, err, `"nonexistent" not found`)
	})
}
