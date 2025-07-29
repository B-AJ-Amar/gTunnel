package cmd

import (
	"fmt"

	"github.com/B-AJ-Amar/gTunnel/internal/server"
	"github.com/spf13/cobra"
)

var port int

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the gTunnel server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting server on port %d...\n", port)
		server.StartServer(fmt.Sprintf(":%d", port))
	},
}

func init() {
	startCmd.Flags().IntVarP(&port, "port", "p", 5780, "Port to run the server on")
	
}
