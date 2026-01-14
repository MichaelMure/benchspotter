package commands

import (
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
)

func newAnalyzeCommand(env *execenv.Env) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "analyze",
		Aliases: []string{"analyse"},
		Short:   "Analyse benchmarking results",
	}

	cmd.AddCommand(newAnalyseCompareCommand(env))
	cmd.AddCommand(newAnalyseAllocCommand(env))

	return cmd
}
