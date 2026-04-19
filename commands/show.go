package commands

import (
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
)

func newShowCommand(env *execenv.Env) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Display raw benchmarking results",
	}

	cmd.AddCommand(
		newShowDiffCommand(env),
		newShowCPUCommand(env),
		newShowMemCommand(env),
		newShowBlockCommand(env),
		newShowMutexCommand(env),
		newShowEscapeCommand(env),
	)

	return cmd
}
