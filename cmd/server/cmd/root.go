package cmd

import (
	"github.com/B-AJ-Amar/gTunnel/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gts", // gtunnel-server
	Short:   "gTunnel server CLI for managing and running the tunnel server",
	Long:    `A command-line tool to run and manage the gTunnel server instance.`,
	Version: version.GetVersion(),
	RunE: func(cmd *cobra.Command, args []string) error {
		// If version flag is used, it will be handled automatically by cobra
		// Otherwise, show help
		return cmd.Help()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(versionCmd)
}
