package commands

import (
	"path/filepath"
	"strings"
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

func TestShowInline(t *testing.T) {
	t.Run("split_subject", func(t *testing.T) {
		cases := []struct {
			msg                     string
			prefix, subject, suffix string
		}{
			{"cannot inline Foo: too complex", "cannot inline ", "Foo", ": too complex"},
			{"cannot inline Bar", "cannot inline ", "Bar", ""},
			{"inlining call to pkg.Baz", "inlining call to ", "pkg.Baz", ""},
			{"can inline Qux with cost 64 as:", "can inline ", "Qux", " with cost 64 as:"},
			{"can inline Bar", "can inline ", "Bar", ""},
			{"singleword", "", "singleword", ""},
		}
		for _, tc := range cases {
			p, s, suf := splitInlineSubject(tc.msg)
			assert.Equal(t, tc.prefix, p, "prefix for %q", tc.msg)
			assert.Equal(t, tc.subject, s, "subject for %q", tc.msg)
			assert.Equal(t, tc.suffix, suf, "suffix for %q", tc.msg)
		}
	})

	t.Run("json", func(t *testing.T) {
		inlineData := `engine/foo.go:10:5: cannot inline Foo: too complex
engine/foo.go:20:3: inlining call to bar.Baz
engine/foo.go:30:1: can inline Qux
`
		setup := func(t *testing.T) (string, billy.Filesystem) {
			t.Helper()
			storage := memfs.New()
			sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)
			dir := filepath.Join("sessions", sessionID)
			require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.InlineFilename), []byte(inlineData), 0644))
			return sessionID, storage
		}

		t.Run("default", func(t *testing.T) {
			sessionID, storage := setup(t)
			env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON

			err := runShowInline(t.Context(), env, showInlineOptions{session: sessionID})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, `"cannot_inline"`)
			assert.NotContains(t, out, `"inlining_call"`)
			assert.NotContains(t, out, `"can_inline"`)
		})

		t.Run("all", func(t *testing.T) {
			sessionID, storage := setup(t)
			env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON

			err := runShowInline(t.Context(), env, showInlineOptions{session: sessionID, all: true})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, `"cannot_inline"`)
			assert.Contains(t, out, `"inlining_call"`)
			assert.Contains(t, out, `"can_inline"`)
		})

		t.Run("project_filter", func(t *testing.T) {
			storage := memfs.New()
			sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)
			mixed := `engine/foo.go:10:5: cannot inline Foo: too complex
../../.asdf/installs/golang/1.25/go/src/fmt/format.go:123:4: cannot inline Fprintf: too complex
`
			dir := filepath.Join("sessions", sessionID)
			require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.InlineFilename), []byte(mixed), 0644))

			t.Run("default", func(t *testing.T) {
				env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
				env.Format = execenv.FormatJSON

				err := runShowInline(t.Context(), env, showInlineOptions{session: sessionID})
				require.NoError(t, err)
				out := env.Out.String()
				assert.Contains(t, out, "engine/foo.go")
				assert.NotContains(t, out, "fmt/format.go")
			})

			t.Run("include_deps", func(t *testing.T) {
				env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
				env.Format = execenv.FormatJSON

				err := runShowInline(t.Context(), env, showInlineOptions{session: sessionID, includeDeps: true})
				require.NoError(t, err)
				assert.Contains(t, env.Out.String(), "fmt/format.go")
			})
		})
	})

	t.Run("text", func(t *testing.T) {
		t.Run("git_source_loading", func(t *testing.T) {
			const commit = "abc1234def5678abc1234def5678abc1234def56"
			gitSrc := repository.MapGitSource{
				commit + ":engine/foo.go": []byte(`package engine

func BenchmarkFoo(b *testing.B) {
	callHelper()
}

func callHelper() {}
`),
			}
			storage := memfs.New()
			sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, commit, false)

			inlineData := "engine/foo.go:4:2: inlining call to callHelper\n"
			dir := filepath.Join("sessions", sessionID)
			require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.InlineFilename), []byte(inlineData), 0644))

			env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, gitSrc))
			env.Format = execenv.FormatText

			err := runShowInline(t.Context(), env, showInlineOptions{session: sessionID, all: true})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, "BenchmarkFoo")
			assert.Contains(t, out, "callHelper")
			assert.Contains(t, out, "inlining call to")
		})

		t.Run("git_diff_applied", func(t *testing.T) {
			const commit = "abc1234def5678abc1234def5678abc1234def56"
			committedSrc := `package engine

func BenchmarkFoo(b *testing.B) {
	oldHelper()
}

func oldHelper() {}
`
			diff := `diff --git a/engine/foo.go b/engine/foo.go
--- a/engine/foo.go
+++ b/engine/foo.go
@@ -4,1 +4,1 @@
-	oldHelper()
+	newHelper()
`
			gitSrc := repository.MapGitSource{
				commit + ":engine/foo.go": []byte(committedSrc),
			}
			storage := memfs.New()
			sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, commit, false)

			dir := filepath.Join("sessions", sessionID)
			require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.GitDiffFilename), []byte(diff), 0644))
			inlineData := "engine/foo.go:4:2: inlining call to newHelper\n"
			require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.InlineFilename), []byte(inlineData), 0644))

			env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, gitSrc))
			env.Format = execenv.FormatText

			err := runShowInline(t.Context(), env, showInlineOptions{session: sessionID, all: true})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, "newHelper")
			assert.NotContains(t, out, "oldHelper")
		})
	})

	t.Run("no_inline_file", func(t *testing.T) {
		storage := memfs.New()
		sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

		env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		err := runShowInline(t.Context(), env, showInlineOptions{session: sessionID})
		assert.ErrorContains(t, err, "no inline analysis recorded")
	})

	t.Run("session_not_found", func(t *testing.T) {
		storage := memfs.New()
		createTestSession(t, storage, "some-session", []string{"BenchmarkFoo"}, "", false)

		env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		err := runShowInline(t.Context(), env, showInlineOptions{session: "nonexistent-id"})
		assert.ErrorContains(t, err, `"nonexistent-id" not found`)
	})
}

func TestInlineSiteKind(t *testing.T) {
	cases := []struct {
		msg  string
		kind engine.InlineKind
	}{
		{"cannot inline Foo: too complex", engine.InlineCannotInline},
		{"cannot inline Bar", engine.InlineCannotInline},
		{"inlining call to pkg.Baz", engine.InlineInliningCall},
		{"can inline Qux", engine.InlineCanInline},
		{"can inline Qux with cost 64 as:", engine.InlineCanInline},
	}
	for _, tc := range cases {
		s := engine.InlineSite{Message: tc.msg}
		assert.Equal(t, tc.kind, s.Kind(), "Kind() for %q", tc.msg)
	}
}

func TestInlineCountBadge(t *testing.T) {
	const commit = "abc1234def5678abc1234def5678abc1234def56"
	gitSrc := repository.MapGitSource{
		commit + ":engine/foo.go": []byte(`package engine

func BenchmarkFoo(b *testing.B) {
	_ = helperA()
	_ = helperB()
}

func helperA() int { return 1 }
func helperB() int { return 2 }
`),
	}
	storage := memfs.New()
	sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, commit, false)

	inlineData := "engine/foo.go:4:6: cannot inline helperA: too complex\nengine/foo.go:5:6: cannot inline helperB: too complex\n"
	dir := filepath.Join("sessions", sessionID)
	require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.InlineFilename), []byte(inlineData), 0644))

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, gitSrc))
	env.Format = execenv.FormatText

	err := runShowInline(t.Context(), env, showInlineOptions{session: sessionID})
	require.NoError(t, err)

	out := env.Out.String()
	assert.Contains(t, out, "(2)")
}

func TestInlineAnnotationStyle(t *testing.T) {
	sites := []engine.InlineSite{
		{Message: "cannot inline Foo: too complex"},
		{Message: "inlining call to pkg.Bar"},
		{Message: "can inline Baz"},
	}
	var sb strings.Builder
	for _, s := range sites {
		sb.WriteString(inlineAnnotationStyle(s))
		sb.WriteByte('\n')
	}
	out := sb.String()
	// Each annotation carries the function name, preceded by the arrow prefix.
	assert.Contains(t, out, "↑")
	assert.Contains(t, out, "Foo")
	assert.Contains(t, out, "pkg.Bar")
	assert.Contains(t, out, "Baz")
}
