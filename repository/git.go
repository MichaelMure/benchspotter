package repository

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/execabs"
)

func detectGitPath(path string, depth int) (string, error) {
	if depth >= 10 {
		return "", fmt.Errorf("gitdir loop detected")
	}

	// normalize the path
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	for {
		fi, err := os.Stat(filepath.Join(path, ".git"))
		if err == nil {
			if !fi.IsDir() {
				// See if our .git item is a dotfile that holds a submodule reference
				dotfile, err := os.Open(filepath.Join(path, fi.Name()))
				if err != nil {
					// Can't open error
					return "", fmt.Errorf(".git exists but is not a directory or a readable file: %w", err)
				}
				// We aren't going to defer the dotfile.Close, because we might keep looping, so we have to be sure to
				// clean up before returning an error
				reader := bufio.NewReader(io.LimitReader(dotfile, 2048))
				line, _, err := reader.ReadLine()
				_ = dotfile.Close()
				if err != nil {
					return "", fmt.Errorf(".git exists but is not a directory and cannot be read: %w", err)
				}
				dotContent := string(line)
				if strings.HasPrefix(dotContent, "gitdir:") {
					// This is a submodule parent path link. Strip the prefix, clean the string of whitespace just to
					// be safe, and return
					dotContent = strings.TrimSpace(strings.TrimPrefix(dotContent, "gitdir: "))
					p, err := detectGitPath(dotContent, depth+1)
					if err != nil {
						return "", fmt.Errorf(".git gitdir error: %w", err)
					}
					return p, nil
				}
				return "", fmt.Errorf(".git exist but is not a directory or module/workspace file")
			}
			return filepath.Join(path, ".git"), nil
		}
		if !os.IsNotExist(err) {
			// unknown error
			return "", err
		}

		// detect bare repo
		ok, err := isGitDir(path)
		if err != nil {
			return "", err
		}
		if ok {
			return path, nil
		}

		if parent := filepath.Dir(path); parent == path {
			return "", fmt.Errorf(".git not found")
		} else {
			path = parent
		}
	}
}

func isGitDir(path string) (bool, error) {
	markers := []string{"HEAD", "objects", "refs"}

	for _, marker := range markers {
		_, err := os.Stat(filepath.Join(path, marker))
		if err == nil {
			continue
		}
		if !os.IsNotExist(err) {
			// unknown error
			return false, err
		} else {
			return false, nil
		}
	}

	return true, nil
}

// GitSource retrieves file content and metadata from a git repository.
type GitSource interface {
	// FileAtCommit returns the content of relPath (relative to the sources root)
	// at the given git commit hash. Returns an error if not a git repository or
	// if the commit or path cannot be found.
	FileAtCommit(ctx context.Context, commit, relPath string) ([]byte, error)
	// HeadCommit returns the hash of the current HEAD commit.
	HeadCommit(ctx context.Context) (string, error)
	// Diff returns a unified diff of all uncommitted changes vs HEAD.
	// Returns nil, nil when the working tree is clean.
	Diff(ctx context.Context) ([]byte, error)
}

// execGitSource is a GitSource backed by the git binary.
type execGitSource struct {
	root string
}

func (g *execGitSource) FileAtCommit(ctx context.Context, commit, relPath string) ([]byte, error) {
	cmd := execabs.CommandContext(ctx, "git", "show", commit+":"+relPath)
	cmd.Dir = g.root
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git show %s:%s: %w", commit, relPath, err)
	}
	return out, nil
}

func (g *execGitSource) HeadCommit(ctx context.Context) (string, error) {
	cmd := execabs.CommandContext(ctx, "git", "rev-parse", "HEAD")
	cmd.Dir = g.root
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func (g *execGitSource) Diff(ctx context.Context) ([]byte, error) {
	cmd := execabs.CommandContext(ctx, "git", "diff", "HEAD")
	cmd.Dir = g.root
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, nil
	}
	return out, nil
}
