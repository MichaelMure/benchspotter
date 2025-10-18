package engine

import (
	"context"
	"iter"

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

func RunCPUProfile(ctx context.Context, storage billy.Filesystem, id string, benches []BenchInfo, profiles []Profile) func() error {
	next, _ := iter.Pull(func(yield func(error) bool) {
		for _, infos := range benches {
			for _, profile := range profiles {
				var filename, option string
				switch profile {
				case ProfileCPU:
					filename, option = "cpu.profile", "cpuprofile"
				case ProfileMem:
					filename, option = "mem.profile", "memprofile"
				case ProfileMutex:
					filename, option = "mutex.profile", "mutexprofile"
				case ProfileBlock:
					filename, option = "block.profile", "blockprofile"
				default:
					panic("invalid profile")
				}

				out := storage.Join(storage.Root(), sessionDir, id, filename)

				cmd := execabs.CommandContext(ctx, "go", "test", "-bench",
					"^\\Q"+infos.Name+"\\E$", option+"="+out,
					"-run", "^$", ".")
				cmd.Dir = infos.Package

				if err := cmd.Run(); err != nil {
					yield(err)
					return
				}
				if !yield(nil) {
					return
				}
			}
		}
	})
	return func() error {
		err, _ := next()
		return err
	}
}
