package commands

import (
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
)

func newCompareCommand(env *execenv.Env) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compare",
		Short: "Compare results across sessions",
	}

	cmd.AddCommand(newCompareBenchCommand(env))

	return cmd
}
