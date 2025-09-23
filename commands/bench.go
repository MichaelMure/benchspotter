package commands

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"golang.org/x/perf/benchfmt"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
	"benchspotter/repository/locate"
)

type benchOptions struct {
	benchmarks []string
}

func newBenchCommand(env *execenv.Env) *cobra.Command {
	options := benchOptions{}

	cmd := &cobra.Command{
		Use:     "bench",
		Short:   "Run benchmarks in various ways",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runBench(cmd.Context(), env, options)
		},
	}

	flags := cmd.Flags()

	flags.StringSliceVarP(&options.benchmarks, "benchmarks", "b", []string{}, "Benchmarks to run")

	return cmd
}

func runBench(ctx context.Context, env *execenv.Env, options benchOptions) error {
	var benchs []locate.BenchInfo

	err := env.Spinner().Title("Finding benchmarks").Context(ctx).
		ActionWithErr(func(ctx context.Context) error {
			var err error
			benchs, err = env.Repo.Benchmarks(ctx)
			return err
		}).Run()
	if err != nil {
		return fmt.Errorf("failed to discover benchmarks: %w", err)
	}

	if len(benchs) == 0 {
		return fmt.Errorf("no benchmarks found")
	}

	const recallKey = "bench_benchmarks"
	preSelected := env.Repo.GetRecall(recallKey, options.benchmarks)
	var selection []locate.BenchInfo

	err = env.FormSingle(huh.NewMultiSelect[locate.BenchInfo]().
		Title("Select benchmarks").
		OptionsFunc(func() []huh.Option[locate.BenchInfo] {
			opts := make([]huh.Option[locate.BenchInfo], len(benchs))
			for i, info := range benchs {
				opts[i] = huh.NewOption(info.Name, info).
					Selected(slices.Contains(preSelected, info.Name))
			}
			return opts
		}, nil).
		Value(&selection)).
		RunWithContext(ctx)

	err = env.Repo.SetRecall(recallKey, func(yield func(string) bool) {
		for _, info := range selection {
			if !yield(info.Name) {
				return
			}
		}
	})
	if err != nil {
		return err
	}

	id, err := engine.PrepareSession(ctx, env)
	if err != nil {
		return err
	}

	it := engine.RunBenches(ctx, env, id, selection)
	for _, info := range selection {
		start := time.Now()
		var res *benchfmt.Result
		err = env.Spinner().Title(info.Name).ActionWithErr(func(ctx context.Context) error {
			res, err = it()
			return err
		}).Run()
		if err != nil {
			return err
		}

		env.Out.Printf("> Benchmark%s (", res.Name)
		for i, value := range res.Values {
			if i > 0 {
				env.Out.Print(" | ")
			}
			if value.OrigUnit != "" {
				env.Out.Printf("%v %s", value.OrigValue, value.OrigUnit)
			} else {
				env.Out.Printf("%v %s", value.Value, value.Unit)
			}
		}
		env.Out.Printf(") done in %v\n", time.Since(start).Truncate(100*time.Millisecond))
	}

	return nil
}
