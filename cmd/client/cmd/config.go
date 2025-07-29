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
		// TODO: ADD CONFIG FILE USING ZIPER + Call actual client config logic from internal/client
	},
}

func init() {
	configCmd.Flags().StringVarP(&serverUrl, "url", "u", "ws://localhost:8080/___gTl___/ws", "Default server WebSocket URL")
}
