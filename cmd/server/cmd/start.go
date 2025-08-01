package cmd

import (
	"fmt"
	"log"

	"github.com/B-AJ-Amar/gTunnel/internal/server"
	"github.com/B-AJ-Amar/gTunnel/internal/server/repositories"
	"github.com/spf13/cobra"
)

var port int

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the gTunnel server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting server on port %d...\n", port)
		configRepo := repositories.NewServerConfigRepo()

		config, err := configRepo.Load()

		if err != nil {
			log.Printf("WARNING: Failed to load config, That means that authentication is not Supported: %v", err)
		}
		if config.AccessToken == "" {
			fmt.Println("Warning: Access token is not set in the config repo. Please set it for secure access.")
		}

		server.StartServer(fmt.Sprintf(":%d", port))
	},
}

func init() {
	startCmd.Flags().IntVarP(&port, "port", "p", 5780, "Port to run the server on")

}
