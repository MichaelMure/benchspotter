package commands

import (
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
)

func newCompareCommand(env *execenv.Env) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "compare",
		Short: "Compare benchmark results across sessions",
	}

	cmd.AddCommand(newCompareStatCommand(env))

	return cmd
}
