package inputs

import (
	"context"
	"fmt"
	"slices"

	"charm.land/huh/v2"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
)

func SelectBenchmark(env *execenv.Env, preSelect string) (engine.BenchInfo, error) {
	var benchs []engine.BenchInfo

	err := env.Spinner().Title("Finding benchmarks").
		ActionWithErr(func(ctx context.Context) error {
			var err error
			benchs, err = engine.LocateBenchmarks(ctx, env.Repo.Sources())
			return err
		}).Context(env.Ctx).Run()
	if err != nil {
		return engine.BenchInfo{}, fmt.Errorf("failed to discover benchmarks: %w", err)
	}

	if len(benchs) == 0 {
		return engine.BenchInfo{}, fmt.Errorf("no benchmarks found")
	}

	var selected engine.BenchInfo
	err = env.FormSingle(huh.NewSelect[engine.BenchInfo]().
		Title("Select benchmark").
		OptionsFunc(func() []huh.Option[engine.BenchInfo] {
			options := make([]huh.Option[engine.BenchInfo], len(benchs))
			for i, info := range benchs {
				label := info.Name + env.Style.TonedDown(" — "+info.Package)
				options[i] = huh.NewOption(label, info).
					Selected(preSelect == info.Name)
			}
			return options
		}, nil).
		Value(&selected)).
		RunWithContext(env.Ctx)
	if err != nil {
		return engine.BenchInfo{}, err
	}

	return selected, nil
}

func SelectBenchmarks(env *execenv.Env, preSelect []string) ([]engine.BenchInfo, error) {
	var benchs []engine.BenchInfo

	err := env.Spinner().Title("Finding benchmarks").
		ActionWithErr(func(ctx context.Context) error {
			var err error
			benchs, err = engine.LocateBenchmarks(ctx, env.Repo.Sources())
			return err
		}).Context(env.Ctx).Run()
	if err != nil {
		return nil, fmt.Errorf("failed to discover benchmarks: %w", err)
	}

	if len(benchs) == 0 {
		return nil, fmt.Errorf("no benchmarks found")
	}

	var selection []engine.BenchInfo

	err = env.FormSingle(huh.NewMultiSelect[engine.BenchInfo]().
		Title("Select benchmarks (/ to filter)").
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
		RunWithContext(env.Ctx)
	if err != nil {
		return nil, err
	}

	return selection, nil
}
