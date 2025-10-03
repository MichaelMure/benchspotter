package commands

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"golang.org/x/perf/benchfmt"

	"benchspotter/commands/execenv"
	"benchspotter/engine"
)

type benchOptions struct {
	benchmarks []string
	name       string
	count      int
}

// unsetStringMarker is a value marking a string not being set in a string flag.
// It's an invalid utf8 string.
const unsetStringMarker = "\x80"

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
	flags.StringVarP(&options.name, "name", "n", unsetStringMarker, "A name for the benchmark session, for the user to record what is being tested")
	flags.IntVarP(&options.count, "count", "c", -1, "Run benchmarks `n` times")

	return cmd
}

func runBench(ctx context.Context, env *execenv.Env, options benchOptions) error {
	var selection []engine.BenchInfo

	if len(options.benchmarks) == 0 {
		var benchs []engine.BenchInfo

		err := env.Spinner().Title("Finding benchmarks").
			ActionWithErr(func(ctx context.Context) error {
				var err error
				benchs, err = engine.LocateBenchmarks(ctx, env.Repo.Sources())
				return err
			}).Context(ctx).Run()
		if err != nil {
			return fmt.Errorf("failed to discover benchmarks: %w", err)
		}

		if len(benchs) == 0 {
			return fmt.Errorf("no benchmarks found")
		}

		const recallKey = "bench_benchmarks"
		preSelected := env.Repo.GetRecall(recallKey)

		err = env.FormSingle(huh.NewMultiSelect[engine.BenchInfo]().
			Title("Select benchmarks").
			OptionsFunc(func() []huh.Option[engine.BenchInfo] {
				opts := make([]huh.Option[engine.BenchInfo], len(benchs))
				for i, info := range benchs {
					line := info.Name + env.Style.TonedDown(" - "+info.Package)
					opts[i] = huh.NewOption(line, info).
						Selected(slices.Contains(preSelected, info.Name))
				}
				return opts
			}, nil).
			Value(&selection)).
			RunWithContext(ctx)
		if err != nil {
			return err
		}

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
	}

	if options.name == unsetStringMarker {
		err := env.FormSingle(huh.NewInput().
			Title("Name of the session (optional)").
			Validate(func(s string) error {
				if !utf8.ValidString(s) {
					return fmt.Errorf("invalid UTF-8 string %q", s)
				}
				return nil
			}).
			Value(&options.name)).
			RunWithContext(ctx)
		if err != nil {
			return err
		}
	}

	if options.count == -1 {
		var value string
		err := env.FormSingle(huh.NewInput().
			Title("Run benchmarks `n` times").
			Placeholder("1").
			Validate(func(s string) error {
				if s == "" {
					return nil
				}
				_, err := strconv.Atoi(s)
				if err != nil {
					return fmt.Errorf("invalid integer")
				}
				return err
			}).
			Value(&value)).
			RunWithContext(ctx)
		if err != nil {
			return err
		}
		switch value {
		case "":
			options.count = 1
		default:
			options.count, _ = strconv.Atoi(value)
		}
	}

	id, err := engine.PrepareSession(ctx, env, options.name)
	if err != nil {
		return err
	}

	it := engine.RunBenches(ctx, env.Repo.Storage(), id, selection, options.count)
	for _, info := range selection {
		for range options.count {
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
	}

	return nil
}
