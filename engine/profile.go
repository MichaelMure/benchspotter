package engine

import (
	"context"
	"fmt"
	"iter"
	"strings"

	"github.com/go-git/go-billy/v5"
	"golang.org/x/sys/execabs"
)

type Profile int

const (
	ProfileBench Profile = iota
	ProfileCPU
	ProfileMem
	ProfileMutex
	ProfileBlock
)

var replacer = strings.NewReplacer(
	"/", "ᚋ",
	".", "ᚗ",
	"-", "ᚑ",
	"~", "א",
)

func RunProfile(ctx context.Context, storage billy.Filesystem, id string, benches []BenchInfo, profile Profile) func() error {
	next, _ := iter.Pull(func(yield func(error) bool) {
		for _, infos := range benches {
			var dir, filename, option string
			switch profile {
			case ProfileCPU:
				dir, filename, option = "cpu", replacer.Replace(infos.Package+"."+infos.Name)+".profile", "-cpuprofile"
			case ProfileMem:
				dir, filename, option = "mem", replacer.Replace(infos.Package+"."+infos.Name)+".profile", "-memprofile"
			case ProfileMutex:
				dir, filename, option = "mutex", replacer.Replace(infos.Package+"."+infos.Name)+".profile", "-mutexprofile"
			case ProfileBlock:
				dir, filename, option = "block", replacer.Replace(infos.Package+"."+infos.Name)+".profile", "-blockprofile"
			default:
				panic("invalid profile")
			}

			subDir := storage.Join(sessionDir, id, dir)
			if err := storage.MkdirAll(subDir, 0755); err != nil {
				yield(err)
				return
			}
			outFilename := storage.Join(storage.Root(), subDir, filename)

			cmd := execabs.CommandContext(ctx, "go", "test",
				"-bench", infos.Regex(),
				option+"="+outFilename,
				"-run", "^$", ".")
			cmd.Dir = infos.Package

			if out, err := cmd.CombinedOutput(); err != nil {
				yield(fmt.Errorf("failed to profile: %w, %s", err, out))
				return
			}
			if !yield(nil) {
				return
			}
		}
	})
	return func() error {
		err, _ := next()
		return err
	}
}
