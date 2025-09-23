package engine

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/google/uuid"
	"golang.org/x/sys/execabs"

	"benchspotter/commands/execenv"
)

func PrepareSession(ctx context.Context, env *execenv.Env) (string, error) {
	uid, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("failed to generate session id: %w", err)
	}
	id := uid.String()

	err = recordDiffIfAvailable(ctx, env, id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func recordDiffIfAvailable(ctx context.Context, env *execenv.Env, id string) error {
	filename := filepath.Join(id, "git.diff")
	diffFile, err := env.Repo.Storage().Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create diff file: %w", err)
	}

	cmd := execabs.CommandContext(ctx, "git", "diff")
	cmd.Stdout = &CountingWriter{writer: diffFile}
	err = cmd.Run()
	_ = diffFile.Close()
	if err != nil || cmd.Stdout.(*CountingWriter).bytesWritten == 0 {
		_ = env.Repo.Storage().Remove(filename)
	}
	return nil
}

type CountingWriter struct {
	writer       io.Writer
	bytesWritten int
}

func (c *CountingWriter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	c.bytesWritten += n
	return n, err
}
