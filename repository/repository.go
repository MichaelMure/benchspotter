package repository

import (
	"context"
	"fmt"
	"iter"
	"os"
	"os/exec"
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

	return nil
}

func (repo *Repository) Sources() billy.Filesystem {
	return repo.sources
}

// Storage returns the storage space dedicated to benchspotter
func (repo *Repository) Storage() billy.Filesystem {
	return repo.storage
}

func (repo *Repository) Cmd(ctx context.Context, name string, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = repo.sources.Root()
	return cmd
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
