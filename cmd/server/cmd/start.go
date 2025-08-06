package cmd

import (
	"github.com/B-AJ-Amar/gTunnel/internal/logger"
	"github.com/B-AJ-Amar/gTunnel/internal/server"
	"github.com/B-AJ-Amar/gTunnel/internal/server/repositories"
	"github.com/spf13/cobra"
)

var (
	bindAddress string
	debug       bool
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

		logger.Infof("Starting server on %s...", bindAddress)

		if config != nil && config.AccessToken == "" {
			logger.Critical("Access token is not set in the config. Please set it for secure access.")
		}

		server.StartServer(bindAddress)
	},
}

func init() {
	startCmd.Flags().StringVar(&bindAddress, "bind-address", "0.0.0.0:7205", "Address to bind the server to (e.g., 0.0.0.0:8080)")
	startCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
}
