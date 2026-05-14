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

func writeInlineAnalysis(t *testing.T, storage billy.Filesystem, sessionID, data string) {
	t.Helper()
	require.NoError(t, util.WriteFile(storage,
		filepath.Join("sessions", sessionID, engine.InlineFilename),
		[]byte(data), 0644))
}

func TestCompareInline(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		t.Run("statuses", func(t *testing.T) {
			storage := memfs.New()
			baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
			newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)

			writeInlineAnalysis(t, storage, baseID,
				"foo.go:10:1: cannot inline Foo: too complex\n"+
					"foo.go:20:1: inlining call to Bar\n")
			writeInlineAnalysis(t, storage, newID,
				"foo.go:10:1: cannot inline Foo: too complex\n"+ // same
					"foo.go:30:1: cannot inline Baz: function too large\n") // added; Bar removed

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON

			err := runCompareInline(env, compareAnalysisOptions{baseSession: baseID, newSession: newID})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, `"added"`)
			assert.Contains(t, out, `"removed"`)
			assert.Contains(t, out, `"same"`)
			assert.Contains(t, out, "Baz")
			assert.Contains(t, out, "Bar")
			assert.Contains(t, out, "Foo")
		})

		t.Run("kind_field", func(t *testing.T) {
			storage := memfs.New()
			baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
			newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)

			writeInlineAnalysis(t, storage, baseID, "")
			writeInlineAnalysis(t, storage, newID,
				"foo.go:10:1: cannot inline Foo: too complex\n"+
					"foo.go:20:1: inlining call to Bar\n"+
					"foo.go:30:1: can inline Baz\n")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON

			err := runCompareInline(env, compareAnalysisOptions{baseSession: baseID, newSession: newID})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, `"cannot_inline"`)
			assert.Contains(t, out, `"inlining_call"`)
			assert.Contains(t, out, `"can_inline"`)
		})

		t.Run("base_new_lines_for_same_shifted", func(t *testing.T) {
			storage := memfs.New()
			baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
			newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)

			writeInlineAnalysis(t, storage, baseID, "foo.go:10:1: cannot inline Foo: too complex\n")
			writeInlineAnalysis(t, storage, newID, "foo.go:12:1: cannot inline Foo: too complex\n")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON

			err := runCompareInline(env, compareAnalysisOptions{baseSession: baseID, newSession: newID})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, `"same"`)
			assert.Contains(t, out, `"base_line": 10`)
			assert.Contains(t, out, `"new_line": 12`)
		})
	})

	t.Run("text", func(t *testing.T) {
		t.Run("summary_counts", func(t *testing.T) {
			storage := memfs.New()
			baseID := createTestSession(t, storage, "session-base", []string{"BenchmarkFoo"}, "", false)
			newID := createTestSession(t, storage, "session-new", []string{"BenchmarkFoo"}, "", false)

			writeInlineAnalysis(t, storage, baseID, "foo.go:10:1: inlining call to Bar\n")
			writeInlineAnalysis(t, storage, newID, "foo.go:10:1: cannot inline Foo: too complex\n")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatText

			err := runCompareInline(env, compareAnalysisOptions{baseSession: baseID, newSession: newID})
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

			sameData := "foo.go:10:1: cannot inline Foo: too complex\n"
			writeInlineAnalysis(t, storage, baseID, sameData)
			writeInlineAnalysis(t, storage, newID, sameData)

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatText

			err := runCompareInline(env, compareAnalysisOptions{baseSession: baseID, newSession: newID})
			require.NoError(t, err)
			assert.Contains(t, env.Out.String(), "(no changes)")
		})

		t.Run("source_rendering", func(t *testing.T) {
			const commit = "abc1234def5678abc1234def5678abc1234def56"
			gitSrc := repository.MapGitSource{
				commit + ":foo.go": []byte(`package foo

func BenchmarkFoo(b *testing.B) {
	callHelper()
}

func callHelper() {}
`),
			}
			storage := memfs.New()
			baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, commit, false)
			newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, commit, false)

			writeInlineAnalysis(t, storage, baseID, "")
			writeInlineAnalysis(t, storage, newID, "foo.go:4:2: inlining call to callHelper\n")

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, gitSrc))
			env.Format = execenv.FormatText

			err := runCompareInline(env, compareAnalysisOptions{baseSession: baseID, newSession: newID})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, "callHelper")
			assert.Contains(t, out, "inlining call")
		})
	})

	t.Run("session_without_inline_data", func(t *testing.T) {
		storage := memfs.New()
		baseID := createTestSession(t, storage, "base", []string{"BenchmarkFoo"}, "", false)
		newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)
		writeInlineAnalysis(t, storage, baseID, "foo.go:10:1: cannot inline Foo\n")
		// newID has no inline.txt → HasProfile returns false

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		err := runCompareInline(env, compareAnalysisOptions{baseSession: baseID, newSession: newID})
		assert.ErrorContains(t, err, "no analysis of this type")
	})

	t.Run("session_not_found", func(t *testing.T) {
		storage := memfs.New()
		newID := createTestSession(t, storage, "new", []string{"BenchmarkFoo"}, "", false)
		writeInlineAnalysis(t, storage, newID, "")

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		err := runCompareInline(env, compareAnalysisOptions{baseSession: "nonexistent", newSession: newID})
		assert.ErrorContains(t, err, `"nonexistent" not found`)
	})
}
