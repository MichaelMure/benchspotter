package inputs

import (
	"context"
	"fmt"
	"slices"

	"github.com/charmbracelet/huh"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
)

func SelectBenchmarks(ctx context.Context, env *execenv.Env, preSelect []string) ([]engine.BenchInfo, error) {
	var benchs []engine.BenchInfo

	err := env.Spinner().Title("Finding benchmarks").
		ActionWithErr(func(ctx context.Context) error {
			var err error
			benchs, err = engine.LocateBenchmarks(ctx, env.Repo.Sources())
			return err
		}).Context(ctx).Run()
	if err != nil {
		return nil, fmt.Errorf("failed to discover benchmarks: %w", err)
	}

	if len(benchs) == 0 {
		return nil, fmt.Errorf("no benchmarks found")
	}

	var selection []engine.BenchInfo

	err = env.FormSingle(huh.NewMultiSelect[engine.BenchInfo]().
		Title("Select benchmarks").
		OptionsFunc(func() []huh.Option[engine.BenchInfo] {
			opts := make([]huh.Option[engine.BenchInfo], len(benchs))
			for i, info := range benchs {
				line := info.Name + env.Style.TonedDown(" - "+info.Package)
				opts[i] = huh.NewOption(line, info).
					Selected(slices.Contains(preSelect, info.Name))
			}
			return opts
		}, nil).
		Value(&selection)).
		RunWithContext(ctx)
	if err != nil {
		return nil, err
	}

	return selection, nil
}
