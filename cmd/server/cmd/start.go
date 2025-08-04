package cmd

import (
	"fmt"

	"github.com/B-AJ-Amar/gTunnel/internal/logger"
	"github.com/B-AJ-Amar/gTunnel/internal/server"
	"github.com/B-AJ-Amar/gTunnel/internal/server/repositories"
	"github.com/spf13/cobra"
)

var (
	port  int
	debug bool
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the gTunnel server",
	Run: func(cmd *cobra.Command, args []string) {
		logger.ShowBanner("server")

		logLevel := logger.LevelInfo
		if debug {
			logLevel = logger.LevelDebug
		}
		logger.Init(logLevel, true)

		configRepo := repositories.NewServerConfigRepo()

		config, err := configRepo.Load()
		if err != nil {
			logger.Criticalf("Failed to load config, authentication is not supported: %v", err)
		}

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
	startCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
}
