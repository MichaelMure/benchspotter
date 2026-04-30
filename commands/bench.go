package commands

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"time"
	"unicode/utf8"

	"charm.land/huh/v2"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
	"golang.org/x/perf/benchfmt"

	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/engine"
)

type benchOptions struct {
	profiles   []engine.Profile
	benchmarks []string
	name       string
	count      int
}

var profileIds = map[engine.Profile][]string{
	engine.ProfileBench:  {"bench"},
	engine.ProfileCPU:    {"cpu"},
	engine.ProfileMem:    {"mem"},
	engine.ProfileMutex:  {"mutex"},
	engine.ProfileBlock:  {"block"},
	engine.ProfileEscape: {"escape"},
	engine.ProfileInline: {"inline"},
}

var profileHelp = map[engine.Profile]string{
	engine.ProfileBench:  "Run benchmarks",
	engine.ProfileCPU:    "Run CPU profiling",
	engine.ProfileMem:    "Run memory profiling",
	engine.ProfileMutex:  "Run mutex profiling",
	engine.ProfileBlock:  "Run blocking profiling",
	engine.ProfileEscape: "Capture escape analysis",
	engine.ProfileInline: "Capture inlining decisions",
}

// unsetStringMarker is a value marking a string not being set in a string flag.
// It's an invalid utf8 string.
const unsetStringMarker = "\x80"

func newBenchCommand(env *execenv.Env) *cobra.Command {
	options := benchOptions{}

	cmd := &cobra.Command{
		Use:   "bench",
		Short: "Run benchmarks in various ways",
		Long: `Run benchmarks and record the results as a named session.

Each invocation walks an interactive wizard:
  1. Choose one or more profiling modes to run simultaneously.
  2. Select which benchmarks to execute.
  3. Give the session a name (optional, for your own reference).
  4. If running "bench" mode, choose how many times to repeat each benchmark.

Profiling modes:

  bench   — Runs 'go test -bench -benchmem -count N' and records the numeric
             output (ns/op, B/op, allocs/op, and any custom metrics). Use this
             when you want to compare performance numbers across code changes.

  cpu     — Captures a CPU profile via 'go test -cpuprofile'. Shows where
             wall-clock time is spent. Use it to find hot code paths.

  mem     — Captures a memory profile via 'go test -memprofile'. Shows heap
             allocation sites. Use it to find what is allocating most.

  mutex   — Captures a mutex-contention profile via 'go test -mutexprofile'.
             Shows which sync.Mutex locks are held the longest. Useful when
             goroutine scheduling is a bottleneck.

  block   — Captures a blocking profile via 'go test -blockprofile'. Shows
             where goroutines spend time waiting (channels, locks, syscalls).

  escape  — Passes -gcflags='-m' to the compiler and records which variables
             escape to the heap. Use it to reduce unintended allocations.

  inline  — Passes -gcflags='-m' to the compiler and records which functions
             the compiler was unable to inline, and why.

Results are stored under .benchspotter/sessions/<uuid>/. All prompts remember
your last selection and pre-populate it the next time you run this command.
Any interactive prompt can be bypassed with the corresponding flags.`,
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runBench(env, options)
		},
	}

	flags := cmd.Flags()

	profileEnum := enumflag.NewSlice(&options.profiles, "profile", profileIds, enumflag.EnumCaseInsensitive)
	flags.VarP(profileEnum, "profile", "p", "Profiling mode(s) to run")
	err := profileEnum.RegisterCompletion(cmd, "profile", profileHelp)
	if err != nil {
		panic(err)
	}

	flags.StringSliceVarP(&options.benchmarks, "benchmarks", "b", []string{}, "Benchmarks to run")
	flags.StringVarP(&options.name, "name", "n", unsetStringMarker, "A name for the benchmark session, for the user to record what is being tested")
	flags.IntVarP(&options.count, "count", "c", -1, "Run benchmarks `n` times")

	return cmd
}

func runBench(env *execenv.Env, options benchOptions) error {
	var err error

	if len(options.profiles) == 0 {
		const recallKey = "bench_profiles"
		preSelected := env.Repo.GetRecalls(recallKey)
		selected := func(p engine.Profile) bool { return slices.Contains(preSelected, strconv.Itoa(int(p))) }

		err = env.FormSingle(huh.NewMultiSelect[engine.Profile]().
			Title("Benchmarking profile(s)").
			Options(
				huh.NewOption("Benchmark", engine.ProfileBench).Selected(selected(engine.ProfileBench)),
				huh.NewOption("CPU", engine.ProfileCPU).Selected(selected(engine.ProfileCPU)),
				huh.NewOption("Memory", engine.ProfileMem).Selected(selected(engine.ProfileMem)),
				huh.NewOption("Mutex", engine.ProfileMutex).Selected(selected(engine.ProfileMutex)),
				huh.NewOption("Blocking", engine.ProfileBlock).Selected(selected(engine.ProfileBlock)),
				huh.NewOption("Escape analysis", engine.ProfileEscape).Selected(selected(engine.ProfileEscape)),
				huh.NewOption("Inlining decisions", engine.ProfileInline).Selected(selected(engine.ProfileInline)),
			).
			Value(&options.profiles)).
			RunWithContext(env.Ctx)
		if err != nil {
			return err
		}

		err = env.Repo.SetRecalls(recallKey, func(yield func(string) bool) {
			for _, p := range options.profiles {
				if !yield(strconv.Itoa(int(p))) {
					return
				}
			}
		})
		if err != nil {
			return err
		}
	}

	var selection []engine.BenchInfo
	if len(options.benchmarks) == 0 {
		const recallKey = "bench_benchmarks"
		preSelected := env.Repo.GetRecalls(recallKey)

		selection, err = inputs.SelectBenchmarks(env, preSelected)
		if err != nil {
			return err
		}

		err = env.Repo.SetRecalls(recallKey, func(yield func(string) bool) {
			for _, info := range selection {
				if !yield(info.Name) {
					return
				}
			}
		})
		if err != nil {
			return err
		}
	} else {
		var all []engine.BenchInfo
		err = env.Spinner().Title("Finding benchmarks").ActionWithErr(func(ctx context.Context) error {
			var locErr error
			all, locErr = engine.LocateBenchmarks(ctx, env.Repo.Sources())
			return locErr
		}).Context(env.Ctx).Run()
		if err != nil {
			return err
		}
		for _, info := range all {
			if slices.Contains(options.benchmarks, info.Name) {
				selection = append(selection, info)
			}
		}
		if len(selection) == 0 {
			return fmt.Errorf("no benchmarks found matching: %v", options.benchmarks)
		}
	}

	if options.name == unsetStringMarker {
		err = env.FormSingle(huh.NewInput().
			Title("Name of the session (optional)").
			Validate(func(s string) error {
				if !utf8.ValidString(s) {
					return fmt.Errorf("invalid UTF-8 string %q", s)
				}
				return nil
			}).
			Value(&options.name)).
			RunWithContext(env.Ctx)
		if err != nil {
			return err
		}
	}

	if slices.Contains(options.profiles, engine.ProfileBench) && options.count == -1 {
		var value string
		err = env.FormSingle(huh.NewInput().
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
				return nil
			}).
			Value(&value)).
			RunWithContext(env.Ctx)
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

	id, err := engine.PrepareSession(env, options.name, selection)
	if err != nil {
		return err
	}

	if slices.Contains(options.profiles, engine.ProfileBench) {
		it := engine.RunBenches(env.Ctx, env.Repo.Storage(), id, selection, options.count)
		for {
			start := time.Now()
			var res *benchfmt.Result
			err = env.Spinner().Title("Running benchmarks...").ActionWithErr(func(ctx context.Context) error {
				res, err = it()
				return err
			}).Run()
			if err != nil {
				return err
			}
			if res == nil {
				break
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

	if slices.Contains(options.profiles, engine.ProfileCPU) {
		it := engine.RunProfile(env.Ctx, env.Repo.Storage(), id, selection, engine.ProfileCPU)
		for _, info := range selection {
			start := time.Now()
			err = env.Spinner().Title("CPU " + info.Name).ActionWithErr(func(ctx context.Context) error {
				return it()
			}).Context(env.Ctx).Run()
			if err != nil {
				return err
			}
			env.Out.Printf("> CPU profile for %s done in %v\n", info.Name, time.Since(start).Truncate(100*time.Millisecond))
		}
	}

	if slices.Contains(options.profiles, engine.ProfileMem) {
		it := engine.RunProfile(env.Ctx, env.Repo.Storage(), id, selection, engine.ProfileMem)
		for _, info := range selection {
			start := time.Now()
			err = env.Spinner().Title("Memory " + info.Name).ActionWithErr(func(ctx context.Context) error {
				return it()
			}).Context(env.Ctx).Run()
			if err != nil {
				return err
			}
			env.Out.Printf("> Memory profile for %s done in %v\n", info.Name, time.Since(start).Truncate(100*time.Millisecond))
		}
	}

	if slices.Contains(options.profiles, engine.ProfileMutex) {
		it := engine.RunProfile(env.Ctx, env.Repo.Storage(), id, selection, engine.ProfileMutex)
		for _, info := range selection {
			start := time.Now()
			err = env.Spinner().Title("Mutex " + info.Name).ActionWithErr(func(ctx context.Context) error {
				return it()
			}).Context(env.Ctx).Run()
			if err != nil {
				return err
			}
			env.Out.Printf("> Mutex profile for %s done in %v\n", info.Name, time.Since(start).Truncate(100*time.Millisecond))
		}
	}

	if slices.Contains(options.profiles, engine.ProfileBlock) {
		it := engine.RunProfile(env.Ctx, env.Repo.Storage(), id, selection, engine.ProfileBlock)
		for _, info := range selection {
			start := time.Now()
			err = env.Spinner().Title("Block " + info.Name).ActionWithErr(func(ctx context.Context) error {
				return it()
			}).Context(env.Ctx).Run()
			if err != nil {
				return err
			}
			env.Out.Printf("> Block profile for %s done in %v\n", info.Name, time.Since(start).Truncate(100*time.Millisecond))
		}
	}

	wantEscape := slices.Contains(options.profiles, engine.ProfileEscape)
	wantInline := slices.Contains(options.profiles, engine.ProfileInline)
	if wantEscape || wantInline {
		title := "Compiler analysis"
		if wantEscape && !wantInline {
			title = "Escape analysis"
		} else if wantInline && !wantEscape {
			title = "Inlining decisions"
		}
		start := time.Now()
		err = env.Spinner().Title(title).ActionWithErr(func(ctx context.Context) error {
			return engine.RecordCompilerAnalysis(ctx, env.Repo.Sources().Root(), env.Repo.Storage(), id, wantEscape, wantInline)
		}).Context(env.Ctx).Run()
		if err != nil {
			return err
		}
		env.Out.Printf("> %s done in %v\n", title, time.Since(start).Truncate(100*time.Millisecond))
	}

	return nil
}
