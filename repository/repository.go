package repository

import (
	"cmp"
	"context"
	"fmt"
	"iter"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-billy/v5/util"
	"golang.org/x/sys/execabs"

	"benchspotter/repository/locate"
)

// Repository is an abstracted access to files and data in a project.
type Repository struct {
	// TODO: make those being billy.Filesystem to facilitate mocking
	sourcePath string
	path       string
}

// AutoDetect tries to detect the repository root directory. It does, in order:
//  1. uses the BENCHSPOTTER_PATH and BENCHSPOTTER_SOURCES environment
//     variables if they exist
//  2. search upwards from the current directory for a git repository
//  3. uses "go env GOMOD" to find the module root
//  4. uses the current directory
func AutoDetect() (*Repository, error) {
	const repoDir = ".benchspotter"
	const perm = 0755

	var repo Repository

	if p, ok := os.LookupEnv("BENCHSPOTTER_PATH"); ok {
		if fi, err := os.Lstat(p); err == nil && fi.IsDir() {
			repo.path = filepath.Clean(p)
		} else {
			return nil, fmt.Errorf("BENCHSPOTTER_PATH is not a directory")
		}
	}
	if p, ok := os.LookupEnv("BENCHSPOTTER_SOURCES"); ok {
		if fi, err := os.Lstat(p); err == nil && fi.IsDir() {
			repo.sourcePath = filepath.Clean(p)
		} else {
			return nil, fmt.Errorf("BENCHSPOTTER_SOURCES is not a directory")
		}
	}
	if repo.sourcePath != "" && repo.path == "" {
		repo.path = filepath.Join(repo.sourcePath, repoDir)
		return &repo, os.MkdirAll(repo.path, perm)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current directory: %v", err)
	}

	if gitDir, err := detectGitPath(cwd, 0); err == nil {
		repo.sourcePath = filepath.Dir(gitDir)
		repo.path = cmp.Or(repo.path, filepath.Join(repo.sourcePath, repoDir))
		return &repo, os.MkdirAll(repo.path, perm)
	}

	cmd := execabs.Command("go", "env", "GOMOD")
	if out, err := cmd.Output(); err == nil {
		repo.sourcePath = filepath.Dir(string(out))
		repo.path = cmp.Or(repo.path, filepath.Join(repo.sourcePath, repoDir))
		return &repo, os.MkdirAll(repo.path, perm)
	}

	repo.sourcePath = filepath.Join(cwd, repoDir)
	repo.path = cmp.Or(repo.path, filepath.Join(repo.sourcePath, repoDir))
	return &repo, os.MkdirAll(repo.path, perm)
}

func (repo *Repository) Benchmarks(ctx context.Context) ([]locate.BenchInfo, error) {
	return locate.Benchmarks(ctx, repo.sourcePath)
}

type LocalStorage interface {
	billy.Filesystem
	RemoveAll(path string) error
}

// Storage returns the storage space dedicated to benchspotter
func (repo *Repository) Storage() LocalStorage {
	return billyLocalStorage{Filesystem: osfs.New(repo.path)}
}

const recallDir = "recall"

// GetRecall returns the "recall" data for a better UX in a command.
func (repo *Repository) GetRecall(name string, or []string) []string {
	data, err := util.ReadFile(repo.Storage(), filepath.Join(recallDir, name))
	if err != nil {
		return or
	}
	return strings.Split(string(data), "\n")
}

// SetRecall stores a "recall" data for a better UX in a command.
func (repo *Repository) SetRecall(name string, values iter.Seq[string]) error {
	if _, err := repo.Storage().Stat(filepath.Join(recallDir, ".gitignore")); os.IsNotExist(err) {
		err = util.WriteFile(repo.Storage(), filepath.Join(recallDir, ".gitignore"), []byte("*"), 0644)
		if err != nil {
			return err
		}
	}

	var data strings.Builder
	first := true
	for str := range values {
		if !first {
			data.WriteString("\n")
		}
		first = false
		data.WriteString(str)
	}

	return util.WriteFile(repo.Storage(), filepath.Join(recallDir, name), []byte(data.String()), 0644)
}
