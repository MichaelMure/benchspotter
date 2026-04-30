package repository

import (
	"context"
	"fmt"

	"github.com/go-git/go-billy/v5"
)

// NewForTesting constructs a Repository from explicit filesystems. Intended for tests.
// Pass a MapGitSource (or any GitSource) to exercise git-backed source loading;
// pass nil when the test does not need git content.
func NewForTesting(sources, storage billy.Filesystem, git GitSource) *Repository {
	return &Repository{sources: sources, storage: storage, git: git}
}

// MapGitSource is a GitSource stub for tests. Keys are "commit:relPath".
type MapGitSource map[string][]byte

func (m MapGitSource) FileAtCommit(_ context.Context, commit, relPath string) ([]byte, error) {
	if data, ok := m[commit+":"+relPath]; ok {
		return data, nil
	}
	return nil, fmt.Errorf("commit %s: object not found: %s", commit, relPath)
}

func (m MapGitSource) HeadCommit(_ context.Context) (string, error) { return "", fmt.Errorf("no HEAD") }
func (m MapGitSource) Diff(_ context.Context) ([]byte, error)        { return nil, nil }
