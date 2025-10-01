package engine

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	"github.com/google/uuid"
	"golang.org/x/sys/execabs"

	"benchspotter/commands/execenv"
)

const sessionDir = "sessions"

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
	filename := filepath.Join(sessionDir, id, "git.diff")
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

type SessionInfo struct {
	Id   string
	Root string
	Path string // in storage
}

func LocateSessions(fs billy.Filesystem) ([]SessionInfo, error) {
	dirs, err := fs.ReadDir(sessionDir)
	if err != nil {
		return nil, err
	}

	res := make([]SessionInfo, len(dirs))
	for i, dir := range dirs {
		res[i] = SessionInfo{
			Id:   dir.Name(),
			Root: fs.Root(),
			Path: filepath.Join(sessionDir, dir.Name()),
		}
	}

	return res, nil
}

func (s SessionInfo) BenchPath() string {
	return filepath.Join(s.Path, benchFilename)
}

func (s SessionInfo) BenchFullPath() string {
	return filepath.Join(s.Root, s.Path, benchFilename)
}
