package commands

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/perf/benchfmt"
	"golang.org/x/perf/benchmath"
	"golang.org/x/perf/benchproc"

	"benchspotter/commands/benchstat/benchtab"
	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/engine"
)

type analyseCompareOptions struct {
	sessions   []string
	thresholds benchmath.Thresholds
	table      string
	row        string
	col        string
	ignore     string
	filter     string
	confidence float64
	format     string
}

func newAnalyseCompareCommand(env *execenv.Env) *cobra.Command {
	options := analyseCompareOptions{
		thresholds: benchmath.DefaultThresholds,
	}

	cmd := &cobra.Command{
		Use:     "compare",
		Short:   "Compare benchmark results with x/perf/cmd/benchstat",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAnalyseCompare(cmd.Context(), env, options)
		},
	}

	flags := cmd.Flags()

	flags.StringVar(&options.table, "table", ".config", "split results into tables by distinct values of `projection`")
	flags.StringVar(&options.row, "row", ".fullname", "split results into rows by distinct values of `projection`")
	flags.StringVar(&options.col, "col", ".file", "split results into columns by distinct values of `projection`")
	flags.StringVar(&options.ignore, "ignore", "", "ignore variations in `keys`")
	flags.StringVar(&options.filter, "filter", "*", "use only benchmarks matching benchfilter `query`")
	flags.Float64Var(&options.thresholds.CompareAlpha, "alpha",
		options.thresholds.CompareAlpha, "consider change significant if p < `α`")
	flags.Float64Var(&options.confidence, "confidence", 0.95, "confidence `level` for ranges")
	flags.StringVar(&options.format, "format", "text", "print results in `format`:\n  text - plain text\n  csv  - comma-separated values (warnings will be written to stderr)\n")

	return cmd
}

func runAnalyseCompare(ctx context.Context, env *execenv.Env, options analyseCompareOptions) error {
	// Note: largely taken from golang.org/x/perf/cmd/benchstat/main.go
	// at revision v0.0.0-20250909190841-7e13e04d9366/

	var err error
	var selection []*engine.SessionInfo

	if len(options.sessions) == 0 {
		const recallKey = "analyse_compare_sessions"
		preSelected := env.Repo.GetRecall(recallKey)

		selection, err = inputs.SelectSessions(ctx, env, preSelected)
		if err != nil {
			return err
		}

		err = env.Repo.SetRecall(recallKey, func(yield func(string) bool) {
			for _, info := range selection {
				if !yield(info.Id) {
					return
				}
			}
		})
		if err != nil {
			return err
		}
	}

	filter, err := benchproc.NewFilter(options.filter)
	if err != nil {
		return fmt.Errorf("parsing -filter: %s", err)
	}

	var parser benchproc.ProjectionParser
	var parseErr error
	mustParse := func(name, val string, unit bool) *benchproc.Projection {
		var proj *benchproc.Projection
		var err error
		if unit {
			proj, _, err = parser.ParseWithUnit(val, filter)
		} else {
			proj, err = parser.Parse(val, filter)
		}
		if err != nil && parseErr == nil {
			parseErr = fmt.Errorf("parsing %s: %s", name, err)
		}
		return proj
	}
	tableBy := mustParse("-table", options.table, true)
	rowBy := mustParse("-row", options.row, false)
	colBy := mustParse("-col", options.col, false)
	mustParse("-ignore", options.ignore, false)
	residue := parser.Residue()
	if parseErr != nil {
		return parseErr
	}

	if options.thresholds.CompareAlpha < 0 || options.thresholds.CompareAlpha > 1 {
		return fmt.Errorf("-alpha must be in range [0, 1]")
	}
	if options.confidence < 0 || options.confidence > 1 {
		return fmt.Errorf("-confidence must be in range [0, 1]")
	}

	stat := benchtab.NewBuilder(tableBy, rowBy, colBy, residue)

	paths := make([]string, len(selection))
	for i, info := range selection {
		paths[i] = fmt.Sprintf("%s=%s", info.HumanName, info.BenchFullPath())
	}

	files := benchfmt.Files{Paths: paths, AllowStdin: true, AllowLabels: true}
	for files.Scan() {
		switch rec := files.Result(); rec := rec.(type) {
		case *benchfmt.SyntaxError:
			// Non-fatal result parse error. Warn
			// but keep going.
			fmt.Fprintln(env.Err, rec)
		case *benchfmt.Result:
			if ok, err := filter.Apply(rec); !ok {
				if err != nil {
					// Print the reason we rejected this result.
					fmt.Fprintln(env.Err, err)
				}
				continue
			}

			stat.Add(rec)
		}
	}
	if err := files.Err(); err != nil {
		return err
	}

	tables := stat.ToTables(benchtab.TableOpts{
		Confidence: options.confidence,
		Thresholds: &options.thresholds,
		Units:      files.Units(),
	})

	viewport, runFn := env.Viewport(ctx)
	err = tables.ToText(viewport, false)
	if err != nil {
		return err
	}
	return runFn()
}
