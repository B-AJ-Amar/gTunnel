package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gtc", // gtunnel-client
	Short: "gTunnel client CLI for connecting to tunnel servers",
	Long:  `A command-line tool to connect to and manage gTunnel client connections.`,
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
}
