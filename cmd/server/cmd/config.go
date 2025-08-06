package cmd

import (
	"fmt"

	"github.com/B-AJ-Amar/gTunnel/internal/logger"
	"github.com/B-AJ-Amar/gTunnel/internal/server/repositories"
	"github.com/spf13/cobra"
)

var (
	showConfig bool
	setToken   string
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the gTunnel server",
	Long: `Configure the gTunnel server settings.

Examples:
  gts config                           # Show current configuration
  gts config --show                    # Show current configuration  
  gts config --set-token abc123        # Set access token
  gts config --set-port 8080           # Set server port`,
	Run: func(cmd *cobra.Command, args []string) {
		configRepo := repositories.NewServerConfigRepo()

		// Initialize config if it doesn't exist
		if err := configRepo.InitConfig(); err != nil {
			logger.Fatalf("Failed to initialize config: %v", err)
		}

		// Handle different operations
		if setToken != "" {
			if err := configRepo.UpdateAccessToken(setToken); err != nil {
				logger.Fatalf("Failed to update access token: %v", err)
			}
			fmt.Println("Access token updated successfully")
			return
		}

		// Show configuration (default behavior)
		config, err := configRepo.Load()
		if err != nil {
			logger.Fatalf("Failed to load config: %v", err)
		}

		fmt.Printf("Configuration file: %s\n", configRepo.GetConfigPath())
		if config.AccessToken != "" {
			fmt.Printf("Access Token: %s*****\n", config.AccessToken[:min(len(config.AccessToken), 3)])
		} else {
			fmt.Println("Access Token: (not set)")
		}
	},
}

func init() {
	configCmd.Flags().BoolVarP(&showConfig, "show", "s", false, "Show current configuration")
	configCmd.Flags().StringVarP(&setToken, "set-token", "t", "", "Set the access token")
}
