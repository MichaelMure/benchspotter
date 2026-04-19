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

	env := execenv.NewTestEnv(repository.New(memfs.New(), storage))
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

	env := execenv.NewTestEnv(repository.New(memfs.New(), storage))
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

	env := execenv.NewTestEnv(repository.New(memfs.New(), storage))
	env.Format = execenv.FormatJSON

	// default: project only
	err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID})
	require.NoError(t, err)
	out := env.Out.String()
	assert.Contains(t, out, "engine/foo.go")
	assert.NotContains(t, out, "fmt/format.go")

	// with deps included
	env2 := execenv.NewTestEnv(repository.New(memfs.New(), storage))
	env2.Format = execenv.FormatJSON
	err = runShowEscape(t.Context(), env2, showEscapeOptions{session: sessionID, includeDeps: true})
	require.NoError(t, err)
	assert.Contains(t, env2.Out.String(), "fmt/format.go")
}

func TestShowEscape_noEscapeFile(t *testing.T) {
	storage := memfs.New()
	sessionID := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "", false)

	env := execenv.NewTestEnv(repository.New(memfs.New(), storage))
	env.Format = execenv.FormatJSON

	err := runShowEscape(t.Context(), env, showEscapeOptions{session: sessionID})
	assert.ErrorContains(t, err, "no escape analysis recorded")
}

func TestShowEscape_sessionNotFound(t *testing.T) {
	storage := memfs.New()
	createTestSession(t, storage, "some-session", []string{"BenchmarkFoo"}, "", false)

	env := execenv.NewTestEnv(repository.New(memfs.New(), storage))
	env.Format = execenv.FormatJSON

	err := runShowEscape(t.Context(), env, showEscapeOptions{session: "nonexistent-id"})
	assert.ErrorContains(t, err, `"nonexistent-id" not found`)
}
