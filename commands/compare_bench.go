package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/perf/benchfmt"
	"golang.org/x/perf/benchmath"
	"golang.org/x/perf/benchproc"

	"benchspotter/commands/benchstat/benchtab"
	"benchspotter/commands/execenv"
	"benchspotter/commands/inputs"
	"benchspotter/engine"
)

type compareBenchOptions struct {
	sessions         []string
	thresholds       benchmath.Thresholds
	table            string
	row              string
	col              string
	ignore           string
	filter           string
	confidence       float64
	skipMachineCheck bool
}

func newCompareBenchCommand(env *execenv.Env) *cobra.Command {
	options := compareBenchOptions{
		thresholds: benchmath.DefaultThresholds,
	}

	cmd := &cobra.Command{
		Use:     "bench",
		Short:   "Compare benchmark results with x/perf/cmd/benchstat",
		PreRunE: execenv.LoadRepo(env),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCompareBench(env, options)
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
	flags.BoolVar(&options.skipMachineCheck, "skip-machine-check", false, "suppress cross-machine warning")

	return cmd
}

func runCompareBench(env *execenv.Env, options compareBenchOptions) error {
	// Note: largely taken from golang.org/x/perf/cmd/benchstat/main.go
	// at revision v0.0.0-20250909190841-7e13e04d9366/

	var err error
	var selection []*engine.SessionInfo

	if len(options.sessions) == 0 {
		const recallKey = "compare_bench_sessions"
		preSelected := env.Repo.GetRecalls(recallKey)

		selection, err = inputs.SelectSessions(env, preSelected)
		if err != nil {
			return err
		}
		if len(selection) == 0 {
			return fmt.Errorf("no sessions selected")
		}

		err = env.Repo.SetRecalls(recallKey, func(yield func(string) bool) {
			for _, info := range selection {
				if !yield(info.Id) {
					return
				}
			}
		})
		if err != nil {
			return err
		}
	} else {
		all, err := engine.LocateSessions(env.Repo.Storage())
		if err != nil {
			return err
		}
		for _, id := range options.sessions {
			for _, s := range all {
				if s.Id == id {
					selection = append(selection, s)
					break
				}
			}
		}
		if len(selection) == 0 {
			return fmt.Errorf("no matching sessions found")
		}
	}

	if !options.skipMachineCheck {
		if mc := engine.NewMachineContext(selection); mc != nil {
			fmt.Fprintln(env.Err, "Warning: comparing sessions from different machines:")
			for i, e := range mc.Entries {
				fmt.Fprintf(env.Err, "  ⚙%d  %s\n", i+1, engine.FormatMachineLine(&e.Machine, e.GoVersion))
			}
			fmt.Fprintln(env.Err)
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
			fmt.Fprintln(env.Err, rec)
		case *benchfmt.Result:
			if ok, err := filter.Apply(rec); !ok {
				if err != nil {
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

	tableOpts := benchtab.TableOpts{
		Confidence: options.confidence,
		Thresholds: &options.thresholds,
		Units:      files.Units(),
	}

	switch env.Format {
	case execenv.FormatText:
		tables := stat.ToTables(tableOpts)
		viewport, runFn := env.Viewport()
		if err := tables.ToText(viewport, &env.Style); err != nil {
			return err
		}
		return runFn()
	case execenv.FormatJSON:
		tables := stat.ToTables(tableOpts)
		return env.Out.PrintJSON(compareBenchToJSON(tables))
	case execenv.FormatRaw:
		for _, info := range selection {
			fmt.Fprintf(env.Out, "file: %s\n", info.HumanName)
			f, err := os.Open(info.BenchFullPath())
			if err != nil {
				return err
			}
			_, copyErr := io.Copy(env.Out, f)
			_ = f.Close()
			if copyErr != nil {
				return copyErr
			}
		}
		return nil
	default:
		return fmt.Errorf("unsupported format %v for compare bench (text, json, raw)", env.Format)
	}
}

type compareBenchJSONTable struct {
	Unit       string                      `json:"unit"`
	Benchmarks []compareBenchJSONBenchmark `json:"benchmarks"`
}

type compareBenchJSONBenchmark struct {
	Name     string                    `json:"name"`
	Sessions []compareBenchJSONSession `json:"sessions"`
}

type compareBenchJSONSession struct {
	Name   string  `json:"name"`
	Center float64 `json:"center"`
	Range  string  `json:"range,omitempty"`
	Delta  string  `json:"delta,omitempty"`
	Stats  string  `json:"stats,omitempty"`
}

func compareBenchToJSON(tables *benchtab.Tables) []compareBenchJSONTable {
	result := make([]compareBenchJSONTable, len(tables.Tables))
	for i, t := range tables.Tables {
		jTable := compareBenchJSONTable{
			Unit: t.Unit,
		}
		for _, row := range t.Rows {
			jBench := compareBenchJSONBenchmark{
				Name: row.StringValues(),
			}
			for _, col := range t.Cols {
				cell, ok := t.Cells[benchtab.TableKey{Row: row, Col: col}]
				if !ok {
					continue
				}
				jSess := compareBenchJSONSession{
					Name:   col.StringValues(),
					Center: cell.Summary.Center,
					Range:  cell.Summary.PctRangeString(),
				}
				if cell.Baseline != nil {
					jSess.Delta = cell.Comparison.FormatDelta(cell.Baseline.Summary.Center, cell.Summary.Center)
					jSess.Stats = cell.Comparison.String()
				}
				jBench.Sessions = append(jBench.Sessions, jSess)
			}
			jTable.Benchmarks = append(jTable.Benchmarks, jBench)
		}
		result[i] = jTable
	}
	return result
}
