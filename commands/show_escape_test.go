package commands

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
	"benchspotter/repository"
)

func TestShowEscape(t *testing.T) {
	t.Run("split_subject", func(t *testing.T) {
		cases := []struct {
			msg                     string
			prefix, subject, suffix string
		}{
			{"data escapes to heap", "", "data", " escapes to heap"},
			{"&buf escapes to heap", "", "&buf", " escapes to heap"},
			{"moved to heap: result", "moved to heap: ", "result", ""},
			{"leaking param: buf", "leaking param: ", "buf", ""},
			{"leaking param content: r", "leaking param content: ", "r", ""},
			{"can inline Foo", "", "can", " inline Foo"},
			{"singleword", "", "singleword", ""},
		}
		for _, tc := range cases {
			p, s, suf := splitEscapeSubject(tc.msg)
			assert.Equal(t, tc.prefix, p, "prefix for %q", tc.msg)
			assert.Equal(t, tc.subject, s, "subject for %q", tc.msg)
			assert.Equal(t, tc.suffix, suf, "suffix for %q", tc.msg)
		}
	})

	t.Run("is_project_file", func(t *testing.T) {
		root := "/home/user/myproject"
		cases := []struct {
			file string
			want bool
		}{
			{"engine/foo.go", true},
			{"./engine/foo.go", true},
			{"commands/bar.go", true},
			{"/home/user/myproject/engine/foo.go", true},
			{"../../.asdf/installs/golang/1.25/go/src/fmt/format.go", false},
			{"../../go/pkg/mod/github.com/foo/bar@v1.0.0/baz.go", false},
			{"/usr/local/go/src/fmt/format.go", false},
			{"/home/user/go/pkg/mod/github.com/foo/bar@v1.0.0/baz.go", false},
		}
		for _, tc := range cases {
			assert.Equal(t, tc.want, isProjectFile(root, tc.file), "isProjectFile(%q)", tc.file)
		}
	})

	t.Run("json", func(t *testing.T) {
		t.Run("default", func(t *testing.T) {
			storage := memfs.New()
			sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

			escapeData := `engine/foo.go:10:5: bar escapes to heap
engine/foo.go:20:3: leaking param: buf
engine/foo.go:30:1: can inline Baz
`
			dir := filepath.Join("sessions", sessionID)
			require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.EscapeFilename), []byte(escapeData), 0644))

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON

			err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, `"heap_escape": true`)
			assert.Contains(t, out, `"leaking_param": true`)
			assert.NotContains(t, out, "can inline")
		})

		t.Run("all", func(t *testing.T) {
			storage := memfs.New()
			sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

			escapeData := `engine/foo.go:10:5: bar escapes to heap
engine/foo.go:30:1: can inline Baz
`
			dir := filepath.Join("sessions", sessionID)
			require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.EscapeFilename), []byte(escapeData), 0644))

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
			env.Format = execenv.FormatJSON

			err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID, all: true})
			require.NoError(t, err)

			assert.Contains(t, env.Out.String(), "can inline")
		})

		t.Run("project_filter", func(t *testing.T) {
			storage := memfs.New()
			sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

			escapeData := `engine/foo.go:10:5: bar escapes to heap
../../.asdf/installs/golang/1.25/go/src/fmt/format.go:123:4: x escapes to heap
`
			dir := filepath.Join("sessions", sessionID)
			require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.EscapeFilename), []byte(escapeData), 0644))

			t.Run("default", func(t *testing.T) {
				env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
				env.Format = execenv.FormatJSON

				err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID})
				require.NoError(t, err)
				out := env.Out.String()
				assert.Contains(t, out, "engine/foo.go")
				assert.NotContains(t, out, "fmt/format.go")
			})

			t.Run("include_deps", func(t *testing.T) {
				env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
				env.Format = execenv.FormatJSON

				err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID, includeDeps: true})
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
	var s []byte
	for i := 0; i < b.N; i++ {
		s = make([]byte, 1024)
	}
	_ = s
}
`),
			}

			storage := memfs.New()
			sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, commit, false)

			escapeData := "engine/foo.go:4:7: make([]byte, 1024) escapes to heap\n"
			dir := filepath.Join("sessions", sessionID)
			require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.EscapeFilename), []byte(escapeData), 0644))

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, gitSrc))
			env.Format = execenv.FormatText

			err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, "BenchmarkFoo")
			assert.Contains(t, out, "make([]byte, 1024)")
			assert.Contains(t, out, "escapes to heap")
		})

		t.Run("git_diff_applied", func(t *testing.T) {
			const commit = "abc1234def5678abc1234def5678abc1234def56"

			committedSrc := `package engine

func BenchmarkFoo(b *testing.B) {
	var s []byte
	for i := 0; i < b.N; i++ {
		s = make([]byte, 512)
	}
	_ = s
}
`
			diff := `diff --git a/engine/foo.go b/engine/foo.go
--- a/engine/foo.go
+++ b/engine/foo.go
@@ -6,1 +6,1 @@
-		s = make([]byte, 512)
+		s = make([]byte, 1024)
`

			gitSrc := repository.MapGitSource{
				commit + ":engine/foo.go": []byte(committedSrc),
			}

			storage := memfs.New()
			sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, commit, false)

			dir := filepath.Join("sessions", sessionID)
			require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.GitDiffFilename), []byte(diff), 0644))
			escapeData := "engine/foo.go:6:7: make([]byte, 1024) escapes to heap\n"
			require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.EscapeFilename), []byte(escapeData), 0644))

			env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, gitSrc))
			env.Format = execenv.FormatText

			err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID})
			require.NoError(t, err)

			out := env.Out.String()
			assert.Contains(t, out, "1024")
			assert.NotContains(t, out, "512")
		})
	})

	t.Run("no_escape_file", func(t *testing.T) {
		storage := memfs.New()
		sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID})
		assert.ErrorContains(t, err, "no escape analysis recorded")
	})

	t.Run("session_not_found", func(t *testing.T) {
		storage := memfs.New()
		createTestSession(t, storage, "some-session", []string{"BenchmarkFoo"}, "", false)

		env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, nil))
		env.Format = execenv.FormatJSON

		err := runShowEscape(t.Context(), env, showEscapeOptions{session: "nonexistent-id"})
		assert.ErrorContains(t, err, `"nonexistent-id" not found`)
	})

}

func TestFuzzyMatch(t *testing.T) {
	cases := []struct {
		query, target string
		want          bool
	}{
		{"", "anything", true},
		{"", "", true},
		{"abc", "abc", true},
		{"abc", "axbxc", true},
		{"abc", "ABC", true}, // case-insensitive
		{"abc", "ab", false},
		{"abc", "acb", false}, // order matters
		{"foo", "foobar", true},
		{"bar", "foobar", true},
		{"baz", "foobar", false},
		{"eng", "engine/foo.go", true},
		{"bfoo", "BenchmarkFoo", true},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.want, fuzzyMatch(tc.query, tc.target),
			"fuzzyMatch(%q, %q)", tc.query, tc.target)
	}
}

func TestTruncateLeft(t *testing.T) {
	cases := []struct {
		s      string
		maxLen int
		want   string
	}{
		{"hello", 10, "hello"},
		{"hello", 5, "hello"},
		{"hello world", 8, "…o world"},
		{"hello", 1, "…"},
		{"αβγδε", 3, "…δε"},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.want, truncateLeft(tc.s, tc.maxLen),
			"truncateLeft(%q, %d)", tc.s, tc.maxLen)
	}
}

func TestSearchState(t *testing.T) {
	makeBase := func(headers []headerEntry) *sourceViewBase {
		b := &sourceViewBase{headers: headers}
		return b
	}

	t.Run("ctrl_s_enters_search", func(t *testing.T) {
		b := makeBase(nil)
		consumed := b.handleSearchKey("ctrl+s")
		assert.True(t, consumed)
		assert.True(t, b.searchMode)
		assert.Equal(t, "", b.searchQuery)
	})

	t.Run("esc_exits_search", func(t *testing.T) {
		b := makeBase(nil)
		b.handleSearchKey("ctrl+s")
		consumed := b.handleSearchKey("esc")
		assert.True(t, consumed)
		assert.False(t, b.searchMode)
		assert.Equal(t, "", b.searchQuery)
	})

	t.Run("typing_builds_query", func(t *testing.T) {
		b := makeBase(nil)
		b.handleSearchKey("ctrl+s")
		b.handleSearchKey("f")
		b.handleSearchKey("o")
		b.handleSearchKey("o")
		assert.Equal(t, "foo", b.searchQuery)
	})

	t.Run("backspace_removes_char", func(t *testing.T) {
		b := makeBase(nil)
		b.handleSearchKey("ctrl+s")
		b.handleSearchKey("f")
		b.handleSearchKey("o")
		b.handleSearchKey("backspace")
		assert.Equal(t, "f", b.searchQuery)
	})

	t.Run("non_search_key_not_consumed", func(t *testing.T) {
		b := makeBase(nil)
		consumed := b.handleSearchKey("a")
		assert.False(t, consumed)
		assert.False(t, b.searchMode)
	})

	t.Run("status_line_empty_when_inactive", func(t *testing.T) {
		b := makeBase(nil)
		assert.Equal(t, "", b.searchStatusLine())
	})

	t.Run("status_line_no_query", func(t *testing.T) {
		b := makeBase(nil)
		b.handleSearchKey("ctrl+s")
		s := b.searchStatusLine()
		assert.Contains(t, s, "search:")
		assert.NotContains(t, s, "no match")
		assert.NotContains(t, s, "/")
	})

	t.Run("status_line_no_match", func(t *testing.T) {
		headers := []headerEntry{
			{file: "engine/foo.go", fn: "BenchmarkFoo"},
		}
		b := makeBase(headers)
		b.handleSearchKey("ctrl+s")
		b.handleSearchKey("z")
		b.handleSearchKey("z")
		b.handleSearchKey("z")
		s := b.searchStatusLine()
		assert.Contains(t, s, "no match")
	})

	t.Run("status_line_shows_match_count", func(t *testing.T) {
		headers := []headerEntry{
			{file: "engine/foo.go", fn: "BenchmarkFoo"},
			{file: "engine/bar.go", fn: "BenchmarkBar"},
		}
		b := makeBase(headers)
		b.handleSearchKey("ctrl+s")
		b.handleSearchKey("e") // matches both (engine/)
		s := b.searchStatusLine()
		assert.Contains(t, s, "1/2")
	})
}

func TestCountBadgeInTextOutput(t *testing.T) {
	const commit = "abc1234def5678abc1234def5678abc1234def56"
	gitSrc := repository.MapGitSource{
		commit + ":engine/foo.go": []byte(`package engine

func BenchmarkFoo(b *testing.B) {
	var s []byte
	s = make([]byte, 1024)
	t := new(int)
	_ = s
	_ = t
}
`),
	}
	storage := memfs.New()
	sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, commit, false)

	escapeData := "engine/foo.go:5:7: make([]byte, 1024) escapes to heap\nengine/foo.go:6:7: new(int) escapes to heap\n"
	dir := filepath.Join("sessions", sessionID)
	require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.EscapeFilename), []byte(escapeData), 0644))

	env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), storage, gitSrc))
	env.Format = execenv.FormatText

	err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID})
	require.NoError(t, err)

	out := env.Out.String()
	// Two sites in same function → badge should say (2)
	assert.Contains(t, out, "(2)")
}

func TestFlowChainRendering(t *testing.T) {
	site := engine.EscapeSite{
		File:      "engine/foo.go",
		Line:      5,
		Col:       7,
		Message:   "x escapes to heap",
		FlowChain: []string{"flow: from parameter to heap", "flow: assigned to interface"},
	}

	env := execenv.NewTestEnv(t.Context(), repository.NewForTesting(memfs.New(), memfs.New(), nil))

	model := &escapeViewModel{
		sourceViewBase: sourceViewBase{
			env:      env,
			rawCache: make(map[string][]byte),
		},
	}

	var sb strings.Builder
	model.showFlow = false
	model.renderEscapeAnnotations(&sb, []engine.EscapeSite{site})
	assert.NotContains(t, sb.String(), "flow: from parameter")

	sb.Reset()
	model.showFlow = true
	model.renderEscapeAnnotations(&sb, []engine.EscapeSite{site})
	assert.Contains(t, sb.String(), "flow: from parameter")
	assert.Contains(t, sb.String(), "flow: assigned to interface")
}
