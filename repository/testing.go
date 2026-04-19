package repository

import "github.com/go-git/go-billy/v5"

// New constructs a Repository from explicit filesystems. Intended for tests.
func New(sources, storage billy.Filesystem) *Repository {
	return &Repository{sources: sources, storage: storage}
}
