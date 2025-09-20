package repository

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
	"golang.org/x/sys/execabs"

	"benchspotter/repository/locate"
)

// Repository is an abstracted access to files and data in a project.
type Repository struct {
	path string
}

// AutoDetect tries to detect the repository root directory. It does, in order:
// 1. uses the BENCHSPOTTER_PATH environment variable if it exists
// 2. search upwards from the current directory for a git repository
// 3. uses "go env GOMOD" to find the module root
// 4. uses the current directory
func AutoDetect() (*Repository, error) {
	if p, ok := os.LookupEnv("BENCHSPOTTER_PATH"); ok {
		if fi, err := os.Lstat(p); err == nil && fi.IsDir() {
			return &Repository{path: filepath.Clean(p)}, nil
		}
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current directory: %v", err)
	}

	cmd := execabs.Command("go", "env", "GOMOD")
	if out, err := cmd.Output(); err == nil {
		return &Repository{path: filepath.Dir(string(out))}, nil
	}

	if gitDir, err := detectGitPath(cwd, 0); err == nil {
		return &Repository{path: filepath.Dir(gitDir)}, nil
	}

	return &Repository{path: cwd}, nil
}

func (repo *Repository) Benchmarks(ctx context.Context) ([]locate.BenchInfo, error) {
	return locate.Benchmarks(ctx, repo.path)
}

type LocalStorage interface {
	billy.Filesystem
	RemoveAll(path string) error
}

func (repo *Repository) Storage() LocalStorage {
	return billyLocalStorage{Filesystem: osfs.New(repo.path)}
}
