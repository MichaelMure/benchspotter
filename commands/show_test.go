package commands

import (
	"context"
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

	err := runShowDiffCommand(context.Background(), env, showDiffOptions{session: session})
	require.NoError(t, err)

	assert.Contains(t, env.Out.String(), "a change")
}
