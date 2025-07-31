package cmd

import (
	"log"
	"net/url"
	"strings"

	"github.com/B-AJ-Amar/gTunnel/internal/client"
	"github.com/B-AJ-Amar/gTunnel/internal/client/repositories"
	"github.com/spf13/cobra"
)

var (
	serverURL string
	baseURL   string
)

var connectCmd = &cobra.Command{
	Use:   "connect <port|host:port>",
	Short: "Connect to a gTunnel server",
	Long: `Connect to a gTunnel server and tunnel traffic to a local service.

You can specify the tunnel target in two ways:
  - Port only: connect 3000 (defaults to localhost:3000)
  - Host and port: connect myapp.local:3000

The server URL is loaded from configuration. Use 'gtc config --set-url <url>' to set it.

Examples:
  gtc connect 3000                                              # Tunnel to localhost:3000
  gtc connect api.example.com:8080                              # Tunnel to api.example.com:8080
  gtc connect -u ws://example.com:9000/tunnel/ws 3000           # Override server URL for this connection`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]

		// Initialize config repository
		configRepo := repositories.NewConfigRepo()
		if err := configRepo.InitConfig(); err != nil {
			log.Fatalf("Failed to initialize config: %v", err)
		}

		// Determine server URL - use flag if provided, otherwise load from config
		var finalServerURL string
		if serverURL != "" {
			finalServerURL = serverURL
		} else {
			config, err := configRepo.Load()
			if err != nil {
				log.Fatalf("Failed to load config: %v", err)
			}
			finalServerURL = config.ServerURL
			if finalServerURL == "" {
				log.Fatal("No server URL provided. Use --server-url flag or set it in config with 'gtc config --set-url <url>'")
			}
		}

		// Parse the target argument to extract host and port
		var tunnelHost, tunnelPort string
		if strings.Contains(target, ":") {
			// host:port format
			parts := strings.SplitN(target, ":", 2)
			tunnelHost = parts[0]
			tunnelPort = parts[1]
		} else {
			// port only format
			tunnelHost = "localhost"
			tunnelPort = target
		}

		// Parse the server URL
		u, err := url.Parse(finalServerURL)
		if err != nil {
			log.Fatalf("Invalid server URL: %v", err)
		}

		log.Printf("Connecting to server at %s...\n", finalServerURL)
		log.Printf("Tunneling %s:%s...\n", tunnelHost, tunnelPort)

		if baseURL != "" {
			q := u.Query()
			q.Set("baseURL", baseURL)
			u.RawQuery = q.Encode()
		}

		client.StartClient(*u, tunnelHost, tunnelPort)
	},
}

func init() {
	// url should not have a default value , thats a temp solution untill i setup the config file
	// connectCmd.Flags().StringVarP(&serverURL, "url", "u", "localhost:8080/___gTl___/ws", "Server WebSocket URL to connect to")
	connectCmd.Flags().StringVarP(&serverURL, "server-url", "u", "", "Server WebSocket URL to connect to")
	connectCmd.Flags().StringVarP(&baseURL, "route-base-endpoint", "r", "", "Base endpoint path to route the tunneled app")
}
