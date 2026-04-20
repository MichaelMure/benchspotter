package commands

import (
	"path/filepath"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
	"benchspotter/repository"
)

func TestSplitSubject(t *testing.T) {
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
		p, s, suf := splitSubject(tc.msg)
		assert.Equal(t, tc.prefix, p, "prefix for %q", tc.msg)
		assert.Equal(t, tc.subject, s, "subject for %q", tc.msg)
		assert.Equal(t, tc.suffix, suf, "suffix for %q", tc.msg)
	}
}

func TestIsProjectFile(t *testing.T) {
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
}

func TestShowEscapeJSON(t *testing.T) {
	storage := memfs.New()
	sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

	escapeData := `engine/foo.go:10:5: bar escapes to heap
engine/foo.go:20:3: leaking param: buf
engine/foo.go:30:1: can inline Baz
`
	dir := filepath.Join("sessions", sessionID)
	require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.EscapeFilename), []byte(escapeData), 0644))

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	env.Format = execenv.FormatJSON

	err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID})
	require.NoError(t, err)

	out := env.Out.String()
	assert.Contains(t, out, `"heap_escape": true`)
	assert.Contains(t, out, `"leaking_param": true`)
	// "all" off by default: "can inline" should be excluded
	assert.NotContains(t, out, "can inline")
}

func TestShowEscapeJSON_all(t *testing.T) {
	storage := memfs.New()
	sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

	escapeData := `engine/foo.go:10:5: bar escapes to heap
engine/foo.go:30:1: can inline Baz
`
	dir := filepath.Join("sessions", sessionID)
	require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.EscapeFilename), []byte(escapeData), 0644))

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	env.Format = execenv.FormatJSON

	err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID, all: true})
	require.NoError(t, err)

	assert.Contains(t, env.Out.String(), "can inline")
}

func TestShowEscapeJSON_projectFilter(t *testing.T) {
	storage := memfs.New()
	sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

	escapeData := `engine/foo.go:10:5: bar escapes to heap
../../.asdf/installs/golang/1.25/go/src/fmt/format.go:123:4: x escapes to heap
`
	dir := filepath.Join("sessions", sessionID)
	require.NoError(t, util.WriteFile(storage, filepath.Join(dir, engine.EscapeFilename), []byte(escapeData), 0644))

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	env.Format = execenv.FormatJSON

	// default: project only
	err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID})
	require.NoError(t, err)
	out := env.Out.String()
	assert.Contains(t, out, "engine/foo.go")
	assert.NotContains(t, out, "fmt/format.go")

	// with deps included
	env2 := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	env2.Format = execenv.FormatJSON
	err = runShowEscape(t.Context(), env2, showEscapeOptions{session: sessionID, includeDeps: true})
	require.NoError(t, err)
	assert.Contains(t, env2.Out.String(), "fmt/format.go")
}

func TestShowEscapeText_gitSourceLoading(t *testing.T) {
	const commit = "abc1234def5678abc1234def5678abc1234def56"

	// Simulate a project file at the recorded commit with a known function.
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

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, gitSrc))
	env.Format = execenv.FormatText

	err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID})
	require.NoError(t, err)

	out := env.Out.String()
	// Source line from git content should appear.
	assert.Contains(t, out, "BenchmarkFoo")
	assert.Contains(t, out, "make([]byte, 1024)")
	assert.Contains(t, out, "escapes to heap")
}

func TestShowEscapeText_gitDiffApplied(t *testing.T) {
	const commit = "abc1234def5678abc1234def5678abc1234def56"

	// Committed version of the file.
	committedSrc := `package engine

func BenchmarkFoo(b *testing.B) {
	var s []byte
	for i := 0; i < b.N; i++ {
		s = make([]byte, 512)
	}
	_ = s
}
`
	// Diff that changes 512 → 1024 (the actual state when benchmarks were run).
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

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, gitSrc))
	env.Format = execenv.FormatText

	err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID})
	require.NoError(t, err)

	out := env.Out.String()
	// After applying the diff, line 6 should show the patched content (1024, not 512).
	assert.Contains(t, out, "1024")
	assert.NotContains(t, out, "512")
}

func TestShowEscape_noEscapeFile(t *testing.T) {
	storage := memfs.New()
	sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	env.Format = execenv.FormatJSON

	err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID})
	assert.ErrorContains(t, err, "no escape analysis recorded")
}

func TestShowEscape_sessionNotFound(t *testing.T) {
	storage := memfs.New()
	createTestSession(t, storage, "some-session", []string{"BenchmarkFoo"}, "", false)

	env := execenv.NewTestEnv(repository.NewForTesting(memfs.New(), storage, nil))
	env.Format = execenv.FormatJSON

	err := runShowEscape(t.Context(), env, showEscapeOptions{session: "nonexistent-id"})
	assert.ErrorContains(t, err, `"nonexistent-id" not found`)
}
