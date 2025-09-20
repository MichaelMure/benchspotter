package main

import (
	"os"

	"benchspotter/commands"
)

func main() {
	root := commands.NewRootCommand()
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
