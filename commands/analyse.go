package commands

import (
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
)

func newAnalyzeCommand(env *execenv.Env) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "analyze",
		Aliases: []string{"analyse"},
		Short:   "Analyse results in various ways",
	}

	cmd.AddCommand(newAnalyseCompareCommand(env))

	return cmd
}
