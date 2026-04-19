package commands

import (
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"benchspotter/commands/execenv"
	"benchspotter/repository"
)

func TestShowDiff(t *testing.T) {
	storage := memfs.New()
	session := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "abc1234def5678abc1234def5678abc1234def56", true)
	env := execenv.NewTestEnv(repository.New(memfs.New(), storage))

	err := runShowDiffCommand(t.Context(), env, showDiffOptions{session: session})
	require.NoError(t, err)
	assert.Contains(t, env.Out.String(), "a change")
}

func TestShowDiffRaw(t *testing.T) {
	storage := memfs.New()
	session := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "abc1234def5678abc1234def5678abc1234def56", true)
	env := execenv.NewTestEnv(repository.New(memfs.New(), storage))
	env.Format = execenv.FormatRaw

	err := runShowDiffCommand(t.Context(), env, showDiffOptions{session: session})
	require.NoError(t, err)
	assert.Equal(t, "diff --git a/foo.go b/foo.go\n+// a change\n", env.Out.String())
}

func TestShowDiffJSON(t *testing.T) {
	storage := memfs.New()
	session := createTestSession(t, storage, "my-session", []string{"BenchmarkFoo"}, "abc1234def5678abc1234def5678abc1234def56", true)
	env := execenv.NewTestEnv(repository.New(memfs.New(), storage))
	env.Format = execenv.FormatJSON

	err := runShowDiffCommand(t.Context(), env, showDiffOptions{session: session})
	require.NoError(t, err)
	out := env.Out.String()
	assert.Contains(t, out, `"session_id"`)
	assert.Contains(t, out, `"my-session"`)
	assert.Contains(t, out, `"diff"`)
	assert.Contains(t, out, `a change`)
}
