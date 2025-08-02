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
		configRepo := repositories.NewServerConfigRepo()

		config, err := configRepo.Load()
		if err != nil {
			log.Printf("WARNING: Failed to load config, That means that authentication is not Supported: %v", err)
		}

		// Determine which port to use: command line flag > config file > default (7205)
		finalPort := port
		if !cmd.Flags().Changed("port") && config != nil && config.Port != 0 {
			// Port flag not explicitly set, use config value if available
			finalPort = config.Port
		}

		fmt.Printf("Starting server on port %d...\n", finalPort)

		if config != nil && config.AccessToken == "" {
			fmt.Println("Warning: Access token is not set in the config. Please set it for secure access.")
		}

		server.StartServer(fmt.Sprintf(":%d", finalPort))
	},
}

func init() {
	startCmd.Flags().IntVarP(&port, "port", "p", 7205, "Port to run the server on")

}
