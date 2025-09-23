package engine

import (
	"context"
	"iter"

	"golang.org/x/perf/benchfmt"

	"benchspotter/commands/execenv"
	"benchspotter/repository/locate"
)

func RunBenches(ctx context.Context, env *execenv.Env, benches []locate.BenchInfo) func() (error, bool) {
	next, _ := iter.Pull(func(yield func(error) bool) {
		for _, infos := range benches {
			cmd := env.ExecGo(ctx, "test", "-bench", "^\\Q"+infos.Name+"\\E$",
				"-benchmem", "-run", "^$", ".")
			cmd.Dir = infos.Package

			stdout, err := cmd.StdoutPipe()
			if err != nil {
				yield(err)
				return
			}

			err = cmd.Start()
			if err != nil {
				yield(err)
				return
			}

			out, err := env.Repo.Storage().Create(infos.Name + ".bench")
			if err != nil {
				yield(err)
				return
			}

			w := benchfmt.NewWriter(out)
			r := benchfmt.NewReader(stdout, "")
			for r.Scan() {
				res := r.Result()
				// switch res := res.(type) {
				// case *benchfmt.Result:
				// 	// weird dance to set a non-internal config, the API is not meant for that
				// 	res.SetConfig("git-commit", "qsbdjkqsdnbkqqsd")
				// 	idx, _ := res.ConfigIndex("git-commit")
				// 	res.Config[idx].File = true // not internal, meaning it will print in the writer
				// }

				err = w.Write(res)
				if err != nil {
					_ = out.Close()
					yield(err)
					return
				}
			}

			err = out.Close()
			if err != nil {
				yield(err)
				return
			}

			if !yield(nil) {
				return
			}
		}
	})
	return next
}
