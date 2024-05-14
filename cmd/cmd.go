package cmd

import (
	"github.com/spf13/cobra"

	"github.com/simplifyd-systems/buildman/internal/commands"
)

// NewBuildmanCommand generates a Buildman command
//
//nolint:staticcheck
func NewBuildmanCommand() (*cobra.Command, error) {
	cobra.EnableCommandSorting = false

	rootCmd := &cobra.Command{
		Use:   "buildman",
		Short: "CLI for building apps using Buildman",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

		},
	}

	rootCmd.Flags().Bool("version", false, "Show current 'buildman' version")

	commands.AddHelpFlag(rootCmd, "buildman")

	rootCmd.AddCommand(commands.Build())
	rootCmd.AddCommand(commands.Plan())

	rootCmd.AddCommand(commands.CompletionCommand())

	return rootCmd, nil
}
