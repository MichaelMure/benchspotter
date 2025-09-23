package commands

import (
	"github.com/spf13/cobra"

	"benchspotter/commands/execenv"
)

const RootCmdName = "benchspotter"

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   RootCmdName,
		Short: "", // TODO
		Long:  ``, // TODO

		SilenceUsage:      true,
		DisableAutoGenTag: true,
	}

	const mainGroup = "main"
	const utilitiesGroup = "utilities"

	cmd.AddGroup(&cobra.Group{ID: mainGroup, Title: "Main commands"})
	cmd.AddGroup(&cobra.Group{ID: utilitiesGroup, Title: "Utilities"})

	addCmdWithGroup := func(child *cobra.Command, groupID string) {
		cmd.AddCommand(child)
		child.GroupID = groupID
	}

	env := execenv.NewEnv()

	addCmdWithGroup(newBenchCommand(env), mainGroup)
	addCmdWithGroup(newVersionCommand(env), utilitiesGroup)
	cmd.SetHelpCommandGroupID(utilitiesGroup)
	cmd.SetCompletionCommandGroupID(utilitiesGroup)

	return cmd
}
