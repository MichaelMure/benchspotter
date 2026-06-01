package commands

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/engine"
	"charm.land/huh/v2"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
)

type benchOptions struct {
	profiles   []engine.Profile
	allProfile bool
	bench      []string
	allBench   bool
	name       string
	count      int
	printID    bool
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

	flags.BoolVar(&options.allProfile, "all-profile", false, "Run all profiling modes")
	flags.StringArrayVarP(&options.bench, "bench", "b", []string{}, "Run only benchmarks matching `regexp` (repeatable, patterns are OR-ed)")
	flags.BoolVarP(&options.allBench, "all-bench", "a", false, "Run all benchmarks")
	flags.StringVarP(&options.name, "name", "n", unsetStringMarker, "A name for the benchmark session, for the user to record what is being tested")
	flags.IntVarP(&options.count, "count", "c", -1, "Run benchmarks `n` times")
	flags.BoolVar(&options.printID, "print-id", false, "Print only the session UUID to stdout (useful for scripting)")

	return cmd
}

func runBench(env *execenv.Env, options benchOptions) error {
	var err error

	if options.allProfile {
		options.profiles = []engine.Profile{
			engine.ProfileBench, engine.ProfileCPU, engine.ProfileMem,
			engine.ProfileMutex, engine.ProfileBlock, engine.ProfileEscape, engine.ProfileInline,
		}
	} else if len(options.profiles) == 0 {
		const recallKey = "bench_profiles"
		preSelected := env.Repo.GetRecalls(recallKey)
		selected := func(p engine.Profile) bool { return slices.Contains(preSelected, strconv.Itoa(int(p))) }

		err = env.FormSingle("-p/--profile (or --all-profile)", huh.NewMultiSelect[engine.Profile]().
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
	if !options.allBench && len(options.bench) == 0 {
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
		var patterns []*regexp.Regexp
		for _, pat := range options.bench {
			re, err := regexp.Compile(pat)
			if err != nil {
				return fmt.Errorf("invalid --bench pattern %q: %w", pat, err)
			}
			patterns = append(patterns, re)
		}

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
			if options.allBench {
				selection = append(selection, info)
				continue
			}
			for _, re := range patterns {
				if re.MatchString(info.Name) {
					selection = append(selection, info)
					break
				}
			}
		}
		if len(selection) == 0 {
			return fmt.Errorf("no benchmarks found matching: %v", options.bench)
		}
	}

	if options.name == unsetStringMarker {
		err = env.FormSingle("-n/--name", huh.NewInput().
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
		err = env.FormSingle("-c/--count", huh.NewInput().
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

	// Pre-compute display widths shared across all printers.
	nameWidth := 0
	for _, info := range selection {
		if w := len(displayBenchName(info.Name)); w > nameWidth {
			nameWidth = w
		}
	}
	nameWidth++ // ensure at least one space after the longest name

	profileKinds := []struct {
		p    engine.Profile
		kind string
	}{
		{engine.ProfileCPU, "cpu"},
		{engine.ProfileMem, "mem"},
		{engine.ProfileMutex, "mutex"},
		{engine.ProfileBlock, "block"},
	}

	wantEscape := slices.Contains(options.profiles, engine.ProfileEscape)
	wantInline := slices.Contains(options.profiles, engine.ProfileInline)

	kindWidth := len("bench")
	for _, pk := range profileKinds {
		if slices.Contains(options.profiles, pk.p) {
			kindWidth = max(kindWidth, len(pk.kind))
		}
	}
	if wantEscape && wantInline {
		kindWidth = max(kindWidth, len("escape+inline"))
	} else if wantEscape {
		kindWidth = max(kindWidth, len("escape"))
	} else if wantInline {
		kindWidth = max(kindWidth, len("inline"))
	}

	if slices.Contains(options.profiles, engine.ProfileBench) {
		it := engine.RunBenches(env.Ctx, env.Repo.Storage(), id, selection, options.count)
		bp := newBenchPrinter(env, options.count, nameWidth, kindWidth)
		for {
			res, iterErr := bp.Next(env.Ctx, it)
			if iterErr != nil {
				return iterErr
			}
			if res == nil {
				break
			}
			bp.Add(res)
		}
		bp.Flush()
	}

	for _, prof := range profileKinds {
		if !slices.Contains(options.profiles, prof.p) {
			continue
		}
		it := engine.RunProfile(env.Ctx, env.Repo.Storage(), id, selection, prof.p)
		pp := newProfilePrinter(env, prof.kind, len(selection), nameWidth, kindWidth)
		for _, info := range selection {
			start := time.Now()
			err = env.Spinner().Title(prof.kind + " " + info.Name).ActionWithErr(func(ctx context.Context) error {
				return it()
			}).Context(env.Ctx).Run()
			if err != nil {
				return err
			}
			pp.Done(info.Name, time.Since(start))
		}
	}

	if wantEscape || wantInline {
		kind := "escape+inline"
		if wantEscape && !wantInline {
			kind = "escape"
		} else if wantInline && !wantEscape {
			kind = "inline"
		}
		start := time.Now()
		err = env.Spinner().Title("Compiler analysis").ActionWithErr(func(ctx context.Context) error {
			return engine.RecordCompilerAnalysis(ctx, env.Repo.Sources().Root(), env.Repo.Storage(), id, wantEscape, wantInline)
		}).Context(env.Ctx).Run()
		if err != nil {
			return err
		}
		styledKind := env.Style.Accent(kind)
		kindPad := strings.Repeat(" ", max(0, kindWidth-len(kind)))
		elapsed := time.Since(start).Truncate(10 * time.Millisecond)
		env.Err.Printf("  %s%s  %s\n", styledKind, kindPad, env.Style.TonedDown(elapsed.String()))
	}

	switch env.Format {
	case execenv.FormatText:
		if options.printID {
			env.Out.Println(id)
		}
	case execenv.FormatJSON:
		if options.printID {
			return fmt.Errorf("--print-id is only available with text format")
		}
		type benchRunJSON struct {
			ID         string   `json:"id"`
			Name       string   `json:"name,omitempty"`
			Profiles   []string `json:"profiles"`
			Benchmarks []string `json:"benchmarks"`
		}
		benchmarks := make([]string, len(selection))
		for i, b := range selection {
			benchmarks[i] = b.Name
		}
		profiles := make([]string, 0, len(options.profiles))
		for _, p := range options.profiles {
			profiles = append(profiles, profileIds[p][0])
		}
		return env.Out.PrintJSON(benchRunJSON{
			ID:         id,
			Name:       options.name,
			Profiles:   profiles,
			Benchmarks: benchmarks,
		})
	default:
		return fmt.Errorf("unsupported format: %v", env.Format)
	}
	return nil
}
