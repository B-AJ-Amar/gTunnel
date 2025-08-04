package cmd

import (
	"fmt"

	"github.com/B-AJ-Amar/gTunnel/internal/logger"
	"github.com/B-AJ-Amar/gTunnel/internal/server"
	"github.com/B-AJ-Amar/gTunnel/internal/server/repositories"
	"github.com/spf13/cobra"
)

var port int

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the gTunnel server",
	Run: func(cmd *cobra.Command, args []string) {
		// Show banner
		logger.ShowBanner("server")
		
		// Initialize logger
		logger.Init(logger.LevelInfo, true)
		
		configRepo := repositories.NewServerConfigRepo()

		config, err := configRepo.Load()
		if err != nil {
			logger.Criticalf("Failed to load config, authentication is not supported: %v", err)
		}

		// Determine which port to use: command line flag > config file > default (7205)
		finalPort := port
		if !cmd.Flags().Changed("port") && config != nil && config.Port != 0 {
			// Port flag not explicitly set, use config value if available
			finalPort = config.Port
		}

		logger.Infof("Starting server on port %d...", finalPort)

		if config != nil && config.AccessToken == "" {
			logger.Critical("Access token is not set in the config. Please set it for secure access.")
		}

		server.StartServer(fmt.Sprintf(":%d", finalPort))
	},
}

func init() {
	startCmd.Flags().IntVarP(&port, "port", "p", 7205, "Port to run the server on")

}
