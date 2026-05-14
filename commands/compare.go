package commands

import (
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
)

func newCompareCommand(env *execenv.Env) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compare",
		Short: "Compare results across sessions",
		Long: `Compare benchmark results across two or more recorded sessions.

Select the sessions you want to compare and benchspotter produces statistical
summaries showing the central value, confidence interval, and percentage change
relative to the baseline (first) session, for each benchmark and metric.`,
	}

	cmd.AddCommand(newCompareBenchCommand(env))
	cmd.AddCommand(newCompareCPUCommand(env))
	cmd.AddCommand(newCompareMemCommand(env))
	cmd.AddCommand(newCompareMutexCommand(env))
	cmd.AddCommand(newCompareBlockCommand(env))
	cmd.AddCommand(newCompareInlineCommand(env))
	cmd.AddCommand(newCompareEscapeCommand(env))

	return cmd
}
