package cmd

import (
	"fmt"
	"log"

	"github.com/B-AJ-Amar/gTunnel/internal/client/repositories"
	"github.com/spf13/cobra"
)

var (
	showConfig bool
	setUrl     string
	setToken   string
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the gTunnel client",
	Long: `Configure the gTunnel client settings.

Examples:
  gtc config                                        # Show current configuration
  gtc config --show                                 # Show current configuration  
  gtc config --set-url ws://example.com:8080/ws    # Set server URL
  gtc config --set-token abc123                     # Set access token`,
	Run: func(cmd *cobra.Command, args []string) {
		configRepo := repositories.NewConfigRepo()
		
		// Initialize config if it doesn't exist
		if err := configRepo.InitConfig(); err != nil {
			log.Fatalf("Failed to initialize config: %v", err)
		}
		
		// Handle different operations
		if setUrl != "" {
			if err := configRepo.UpdateServerURL(setUrl); err != nil {
				log.Fatalf("Failed to update server URL: %v", err)
			}
			fmt.Printf("Server URL updated to: %s\n", setUrl)
			return
		}
		
		if setToken != "" {
			if err := configRepo.UpdateAccessToken(setToken); err != nil {
				log.Fatalf("Failed to update access token: %v", err)
			}
			fmt.Println("Access token updated successfully")
			return
		}
		
		// Show configuration (default behavior)
		config, err := configRepo.Load()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
		
		fmt.Printf("Configuration file: %s\n", configRepo.GetConfigPath())
		fmt.Printf("Server URL: %s\n", config.ServerURL)
		if config.AccessToken != "" {
			fmt.Printf("Access Token: %s...\n", config.AccessToken[:min(len(config.AccessToken), 10)])
		} else {
			fmt.Println("Access Token: (not set)")
		}
	},
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func init() {
	configCmd.Flags().BoolVarP(&showConfig, "show", "s", false, "Show current configuration")
	configCmd.Flags().StringVar(&setUrl, "set-url", "", "Set the server WebSocket URL")
	configCmd.Flags().StringVar(&setToken, "set-token", "", "Set the access token")
	// TODO: set a config file directly e.g .gtunnle file
}
