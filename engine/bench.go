package engine

import (
	"context"
	"iter"
	"path/filepath"

	"golang.org/x/perf/benchfmt"

	"benchspotter/commands/execenv"
	"benchspotter/repository/locate"
)

func RunBenches(ctx context.Context, env *execenv.Env, id string, benches []locate.BenchInfo) func() (*benchfmt.Result, error) {
	next, _ := iter.Pull2(func(yield func(*benchfmt.Result, error) bool) {
		out, err := env.Repo.Storage().Create(filepath.Join(id, "results.bench"))
		if err != nil {
			yield(nil, err)
			return
		}
		defer func() { _ = out.Close() }()
		w := benchfmt.NewWriter(out)

		for _, infos := range benches {
			cmd := env.ExecGo(ctx, "test", "-bench", "^\\Q"+infos.Name+"\\E$",
				"-benchmem", "-run", "^$", ".")
			cmd.Dir = infos.Package

			stdout, err := cmd.StdoutPipe()
			if err != nil {
				yield(nil, err)
				return
			}

			err = cmd.Start()
			if err != nil {
				yield(nil, err)
				return
			}

			r := benchfmt.NewReader(stdout, "")

			var res *benchfmt.Result
			for r.Scan() {
				line := r.Result()
				switch line := line.(type) {
				case *benchfmt.Result:
					res = line
					// 	// weird dance to set a non-internal config, the API is not meant for that
					// 	res.SetConfig("git-commit", "qsbdjkqsdnbkqqsd")
					// 	idx, _ := res.ConfigIndex("git-commit")
					// 	res.Config[idx].File = true // not internal, meaning it will print in the writer
				}

				err = w.Write(res)
				if err != nil {
					yield(nil, err)
					return
				}
			}

			if !yield(res, nil) {
				return
			}
		}
	})
	return func() (*benchfmt.Result, error) {
		res, err, _ := next()
		return res, err
	}
}
