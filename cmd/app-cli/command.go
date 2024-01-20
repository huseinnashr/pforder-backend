package main

import (
	"context"

	"github.com/spf13/cobra"
)

func startCommand(ctx context.Context, commands ...*cobra.Command) error {
	rootCommand := &cobra.Command{Use: "app-cli"}
	rootCommand.AddCommand(commands...)

	if err := rootCommand.ExecuteContext(ctx); err != nil {
		return err
	}

	return nil
}
