package commands

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
	"benchspotter/execute"
	"benchspotter/repository/locate"
)

func newBenchCommand(env *execenv.Env) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bench",
		Short:   "Run benchmarks in various ways",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runBench(cmd.Context(), env)
		},
	}

	return cmd
}

func runBench(ctx context.Context, env *execenv.Env) error {
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

	var selection []locate.BenchInfo

	err = env.FormSingle(huh.NewMultiSelect[locate.BenchInfo]().
		Title("Select benchmarks").
		OptionsFunc(func() []huh.Option[locate.BenchInfo] {
			opts := make([]huh.Option[locate.BenchInfo], len(benchs))
			for i, info := range benchs {
				opts[i] = huh.Option[locate.BenchInfo]{
					Key:   info.Name,
					Value: info,
				}
			}
			return opts
		}, nil).
		Value(&selection)).
		RunWithContext(ctx)

	return execute.RunBenches(ctx, env, selection)
}
