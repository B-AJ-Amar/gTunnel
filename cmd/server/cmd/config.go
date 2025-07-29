package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config the gTunnel server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("config server ")
		// TODO: Call actual server conf logic from internal/server
	},
}

func init() {
	configCmd.Flags().IntVarP(&port, "port", "p", 12345, "Port to run the server on")
}
