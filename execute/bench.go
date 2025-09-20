package execute

import (
	"context"

	"golang.org/x/perf/benchfmt"

	"benchspotter/commands/execenv"
	"benchspotter/repository/locate"
)

func RunBenches(ctx context.Context, env *execenv.Env, benches []locate.BenchInfo) error {
	// batches := make(map[string][]locate.BenchInfo)
	// for _, bench := range benches {
	// 	batches[bench.Package] = append(batches[bench.Package], bench)
	// }
	// TODO: make sure to sanitize the bench name to avoid execution

	for _, infos := range benches {
		// var expr strings.Builder
		// expr.WriteString("^\\Q")
		// for _, info := range infos {
		// 	expr.WriteString(info.Name)
		// }
		// expr.WriteString("\\E$")

		cmd := env.ExecGo(ctx, "test", "-bench", "^\\Q"+infos.Name+"\\E$",
			"-benchmem", "-run", "^$", ".")
		cmd.Dir = infos.Package

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}

		out, err := env.Repo.Storage().Create(infos.Name + ".bench")
		if err != nil {
			return err
		}

		err = cmd.Start()
		if err != nil {
			return err
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
				return err
			}
		}
	}

	return nil
}
