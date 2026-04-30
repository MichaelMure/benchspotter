package commands

import (
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
)

func newShowCommand(env *execenv.Env) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Display raw benchmarking results",
		Long: `Display the data recorded for a session.

Each subcommand targets a specific artifact type. Only sessions that actually
contain the requested artifact are offered in the interactive selector.`,
	}

	cmd.AddCommand(
		newShowDiffCommand(env),
		newShowCPUCommand(env),
		newShowMemCommand(env),
		newShowBlockCommand(env),
		newShowMutexCommand(env),
		newShowEscapeCommand(env),
		newShowInlineCommand(env),
	)

	return cmd
}
