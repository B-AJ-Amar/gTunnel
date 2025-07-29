package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gts", // gtunnel-server
	Short: "gTunnel server CLI for managing and running the tunnel server",
	Long:  `A command-line tool to run and manage the gTunnel server instance.`,
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
}
