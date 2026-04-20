package repository

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
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

// GitSource retrieves file content at a specific git commit.
type GitSource interface {
	FileAtCommit(commit, relPath string) ([]byte, error)
}

// goGitSource is the real GitSource backed by a go-git repository.
type goGitSource struct {
	repo *gogit.Repository
}

func (g *goGitSource) FileAtCommit(commit, relPath string) ([]byte, error) {
	commitObj, err := g.repo.CommitObject(plumbing.NewHash(commit))
	if err != nil {
		return nil, fmt.Errorf("commit %s: %w", commit, err)
	}
	tree, err := commitObj.Tree()
	if err != nil {
		return nil, err
	}
	file, err := tree.File(relPath)
	if err != nil {
		return nil, err
	}
	contents, err := file.Contents()
	if err != nil {
		return nil, err
	}
	return []byte(contents), nil
}
