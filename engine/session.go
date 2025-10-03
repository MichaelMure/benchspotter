package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/go-git/go-billy/v5"
	"github.com/google/uuid"
	"github.com/mergestat/timediff"
	"golang.org/x/sys/execabs"

	"benchspotter/commands/execenv"
)

const sessionDir = "sessions"
const metaFilename = "meta.json"

// sessionMeta is a metadata record in the session folder
type sessionMeta struct {
	TimeMillis int64  `json:"timeMillis"`
	Name       string `json:"name,omitempty"`
}

func PrepareSession(ctx context.Context, env *execenv.Env, name string) (string, error) {
	uid, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("failed to generate session id: %w", err)
	}
	id := uid.String()

	err = recordDiffIfAvailable(ctx, env, id)
	if err != nil {
		return "", err
	}

	err = recordMeta(env, id, name)
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

func recordMeta(env *execenv.Env, id string, name string) error {
	filename := filepath.Join(sessionDir, id, metaFilename)
	f, err := env.Repo.Storage().Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create meta file: %w", err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(&sessionMeta{
		TimeMillis: time.Now().UnixMilli(),
		Name:       name,
	})
	if err != nil {
		return fmt.Errorf("failed to encode meta file: %w", err)
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
	Name string
	Root string // root of the storage
	Path string // path in storage
	Time time.Time
}

// HumanName returns the given name if available, or make one up if not.
func (s SessionInfo) HumanName() string {
	if len(s.Name) > 0 {
		return s.Name
	}
	return timediff.TimeDiff(s.Time)
}

func LocateSessions(fs billy.Filesystem) ([]SessionInfo, error) {
	dirs, err := fs.ReadDir(sessionDir)
	if err != nil {
		return nil, err
	}

	res := make([]SessionInfo, len(dirs))
	for i, dir := range dirs {
		f, err := fs.Open(filepath.Join(sessionDir, dir.Name(), metaFilename))
		if err != nil {
			return nil, fmt.Errorf(`failed to open meta file for session "%s": %w`, dir.Name(), err)
		}
		meta := &sessionMeta{}
		err = json.NewDecoder(f).Decode(meta)
		if err != nil {
			return nil, fmt.Errorf(`failed to decode meta file for session "%s": %w`, dir.Name(), err)
		}

		res[i] = SessionInfo{
			Id:   dir.Name(),
			Name: meta.Name,
			Root: fs.Root(),
			Path: filepath.Join(sessionDir, dir.Name()),
			Time: time.UnixMilli(meta.TimeMillis),
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
