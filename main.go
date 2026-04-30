package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"benchspotter/commands"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	root := commands.NewRootCommand(ctx)
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
