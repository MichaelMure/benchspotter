package repository

import (
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
)

const recallDir = "recall"

// Repository is an abstracted access to sources and BenchSpotter's data in a project.
type Repository struct {
	sources billy.Filesystem
	storage billy.Filesystem
	git     GitSource // nil if not a git repository
}

// AutoDetect tries to detect the repository root directory. It does, in order:
//  1. uses the BENCHSPOTTER_PATH and BENCHSPOTTER_SOURCES environment
//     variables if they exist
//  2. search upwards from the current directory for a git repository
//  3. uses "go env GOMOD" to find the module root
//  4. uses the current directory
func AutoDetect() (*Repository, error) {
	const repoDir = ".benchspotter"

	var repo Repository

	if p, ok := os.LookupEnv("BENCHSPOTTER_PATH"); ok {
		if fi, err := os.Lstat(p); err == nil && fi.IsDir() {
			repo.storage = osfs.New(p, osfs.WithBoundOS())
		} else {
			return nil, fmt.Errorf("BENCHSPOTTER_PATH is not a directory")
		}
	}
	if p, ok := os.LookupEnv("BENCHSPOTTER_SOURCES"); ok {
		if fi, err := os.Lstat(p); err == nil && fi.IsDir() {
			repo.sources = osfs.New(p, osfs.WithBoundOS())
		} else {
			return nil, fmt.Errorf("BENCHSPOTTER_SOURCES is not a directory")
		}
	}
	if repo.sources != nil && repo.storage == nil {
		repo.storage, _ = repo.sources.Chroot(repoDir)
		return &repo, repo.init()
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current directory: %v", err)
	}

	if gitDir, err := detectGitPath(cwd, 0); err == nil {
		repo.sources = osfs.New(filepath.Dir(gitDir), osfs.WithBoundOS())
		if repo.storage == nil {
			repo.storage, _ = repo.sources.Chroot(repoDir)
		}
		return &repo, repo.init()
	}

	cmd := execabs.Command("go", "env", "GOMOD")
	if out, err := cmd.Output(); err == nil {
		repo.sources = osfs.New(filepath.Dir(string(out)), osfs.WithBoundOS())
		if repo.storage == nil {
			repo.storage, _ = repo.sources.Chroot(repoDir)
		}
		return &repo, repo.init()
	}

	repo.sources = osfs.New(filepath.Join(cwd, repoDir), osfs.WithBoundOS())
	if repo.storage == nil {
		repo.storage, _ = repo.sources.Chroot(repoDir)
	}

	return &repo, repo.init()
}

func (repo *Repository) init() error {
	const perm = 0755

	// create root storage directory
	err := repo.storage.MkdirAll("/", perm)
	if err != nil {
		return err
	}

	// create root .gitignore
	if _, err := repo.storage.Stat(".gitignore"); os.IsNotExist(err) {
		content := "/" + recallDir + "/"

		err = util.WriteFile(repo.storage, ".gitignore", []byte(content), 0644)
		if err != nil {
			return err
		}
	}

	if _, err := detectGitPath(repo.sources.Root(), 0); err == nil {
		repo.git = &execGitSource{root: repo.sources.Root()}
	}

	return nil
}

func (repo *Repository) Sources() billy.Filesystem {
	return repo.sources
}

// Storage returns the storage space dedicated to benchspotter
func (repo *Repository) Storage() billy.Filesystem {
	return repo.storage
}

// FileAtCommit returns the content of relPath (relative to the sources root)
// at the given git commit hash. Returns an error if not a git repository or
// if the commit or path cannot be found.
func (repo *Repository) FileAtCommit(ctx context.Context, commit, relPath string) ([]byte, error) {
	return repo.git.FileAtCommit(ctx, commit, relPath)
}

// HeadCommit returns the hash of the current HEAD commit.
func (repo *Repository) HeadCommit(ctx context.Context) (string, error) {
	return repo.git.HeadCommit(ctx)
}

// Diff returns a unified diff of all uncommitted changes vs HEAD.
// Returns nil, nil when the working tree is clean or there is no git repository.
func (repo *Repository) Diff(ctx context.Context) ([]byte, error) {
	return repo.git.Diff(ctx)
}

// GetRecall returns a single value from the "recall" storage for a better UX
// in a command.
func (repo *Repository) GetRecall(name string) string {
	data, err := util.ReadFile(repo.storage, filepath.Join(recallDir, name))
	if err != nil {
		return ""
	}
	return string(data)
}

// SetRecall stores a single value in a "recall" storage for a better UX
// in a command.
func (repo *Repository) SetRecall(name string, value string) error {
	return util.WriteFile(repo.storage, filepath.Join(recallDir, name), []byte(value), 0644)
}

// GetRecalls returns multiple values from the "recall" storage for a better UX
// in a command.
func (repo *Repository) GetRecalls(name string) []string {
	data, err := util.ReadFile(repo.storage, filepath.Join(recallDir, name))
	if err != nil {
		return nil
	}
	return strings.Split(string(data), "\n")
}

// SetRecalls stores multiple values in a "recall" storage for a better UX
// in a command.
func (repo *Repository) SetRecalls(name string, values iter.Seq[string]) error {
	var data strings.Builder
	first := true
	for str := range values {
		if !first {
			data.WriteString("\n")
		}
		first = false
		data.WriteString(str)
	}

	return util.WriteFile(repo.storage, filepath.Join(recallDir, name), []byte(data.String()), 0644)
}
