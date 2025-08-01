package cmd

import (
	"github.com/B-AJ-Amar/gTunnel/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gtc", // gtunnel-client
	Short:   "gTunnel client CLI for connecting to tunnel servers",
	Long:    `A command-line tool to connect to and manage gTunnel client connections.`,
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
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(versionCmd)

}
