package cmd

import (
	"fmt"
	"log"

	"github.com/B-AJ-Amar/gTunnel/internal/server/repositories"
	"github.com/spf13/cobra"
)

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var (
	showConfig bool
	setToken   string
	setPort    int
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
			log.Fatalf("Failed to initialize config: %v", err)
		}

		// Handle different operations
		if setToken != "" {
			if err := configRepo.UpdateAccessToken(setToken); err != nil {
				log.Fatalf("Failed to update access token: %v", err)
			}
			fmt.Println("Access token updated successfully")
			return
		}

		if setPort != 0 {
			if err := configRepo.UpdatePort(setPort); err != nil {
				log.Fatalf("Failed to update port: %v", err)
			}
			fmt.Printf("Server port updated to: %d\n", setPort)
			return
		}

		// Show configuration (default behavior)
		config, err := configRepo.Load()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}

		fmt.Printf("Configuration file: %s\n", configRepo.GetConfigPath())
		fmt.Printf("Server Port: %d\n", config.Port)
		if config.AccessToken != "" {
			fmt.Printf("Access Token: %s...\n", config.AccessToken[:min(len(config.AccessToken), 10)])
		} else {
			fmt.Println("Access Token: (not set)")
		}
	},
}

func init() {
	configCmd.Flags().BoolVarP(&showConfig, "show", "s", false, "Show current configuration")
	configCmd.Flags().StringVarP(&setToken, "set-token", "t", "", "Set the access token")
	configCmd.Flags().IntVarP(&setPort, "set-port", "p", 0, "Set the server port (default: 7205)")
}
