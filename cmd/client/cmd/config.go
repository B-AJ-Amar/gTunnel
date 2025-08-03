package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/B-AJ-Amar/gTunnel/internal/client/repositories"
	"github.com/spf13/cobra"
)

var (
	showConfig bool
	setURL     string
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
		configRepo := repositories.NewClientConfigRepo()

		// Initialize config if it doesn't exist
		if err := configRepo.InitConfig(); err != nil {
			log.Fatalf("Failed to initialize config: %v", err)
		}

		// Handle different operations
		if setURL != "" {
			// Ensure URL starts with ws:// or wss://
			if !strings.HasPrefix(setURL, "ws://") && !strings.HasPrefix(setURL, "wss://") {
				setURL = "ws://" + setURL
			}

			if err := configRepo.UpdateServerURL(setURL); err != nil {
				log.Fatalf("Failed to update server URL: %v", err)
			}
			fmt.Printf("Server URL updated to: %s\n", setURL)
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

func init() {
	configCmd.Flags().BoolVarP(&showConfig, "show", "s", false, "Show current configuration")
	configCmd.Flags().StringVarP(&setURL, "set-url", "u", "", "Set the server WebSocket URL")
	configCmd.Flags().StringVarP(&setToken, "set-token", "t", "", "Set the access token")
	// TODO: set a config file directly e.g .gtunnle file
}
