package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the gTunnel client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Configuring client settings...")
		// TODO: Call actual client config logic from internal/client
	},
}

func init() {
	configCmd.Flags().StringVarP(&serverHost, "host", "H", "localhost", "Default server host")
	configCmd.Flags().IntVarP(&serverPort, "port", "p", 8080, "Default server port")
}
