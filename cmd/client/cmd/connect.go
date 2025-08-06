package cmd

import (
	"net/url"
	"strings"

	"github.com/B-AJ-Amar/gTunnel/internal/client"
	"github.com/B-AJ-Amar/gTunnel/internal/client/repositories"
	"github.com/B-AJ-Amar/gTunnel/internal/logger"
	"github.com/spf13/cobra"
)

var (
	serverURL string
	baseURL   string
	debug     bool
)

// buildWebSocketURL constructs the complete WebSocket URL with endpoint
func buildWebSocketURL(serverURL string) (string, error) {
	// Check if it's an HTTPS URL and convert to wss with port 443
	if strings.HasPrefix(serverURL, "https://") {
		serverURL = strings.TrimPrefix(serverURL, "https://")
		serverURL = strings.TrimSuffix(serverURL, "/")
		// Add port 443 if no port is specified
		if !strings.Contains(serverURL, ":") {
			serverURL = serverURL + ":443"
		}
		return serverURL + "/___gTl___/ws", nil
	}

	// Remove any existing protocol
	serverURL = strings.TrimPrefix(serverURL, "ws://")
	serverURL = strings.TrimPrefix(serverURL, "wss://")
	serverURL = strings.TrimPrefix(serverURL, "http://")

	// Remove trailing slash
	serverURL = strings.TrimSuffix(serverURL, "/")

	// Append the WebSocket endpoint
	return serverURL + "/___gTl___/ws", nil
}

var connectCmd = &cobra.Command{
	Use:   "connect <port|host:port>",
	Short: "Connect to a gTunnel server",
	Long: `Connect to a gTunnel server and tunnel traffic to a local service.

You can specify the tunnel target in two ways:
  - Port only: connect 3000 (defaults to localhost:3000)
  - Host and port: connect myapp.local:3000

The server URL is loaded from configuration. Use 'gtc config --set-url <url>' to set it.
The WebSocket endpoint (/___gTl___/ws) is automatically appended.
For HTTPS URLs, port 443 is automatically used if no port is specified.

Examples:
  gtc connect 3000                                              # Tunnel to localhost:3000
  gtc connect api.example.com:8080                              # Tunnel to api.example.com:8080
  gtc connect -u https://example.com 3000                       # Uses port 443 automatically
  gtc connect -u example.com:9000 3000                          # Override server URL for this connection`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Show banner
		logger.ShowBanner("client")

		// Initialize logger with debug level if debug flag is set
		logLevel := logger.LevelInfo
		if debug {
			logLevel = logger.LevelDebug
		}
		logger.Init(logLevel, true)

		target := args[0]

		configRepo := repositories.NewClientConfigRepo()
		if err := configRepo.InitConfig(); err != nil {
			logger.Fatalf("Failed to initialize config: %v", err)
		}

		var finalServerURL string
		if serverURL != "" {
			finalServerURL = serverURL
		} else {
			config, err := configRepo.Load()
			if err != nil {
				logger.Fatalf("Failed to load config: %v", err)
			}
			finalServerURL = config.ServerURL
			if finalServerURL == "" {
				logger.Fatal("No server URL provided. Use --server-url flag or set it in config with 'gtc config --set-url <url>'")
			}
		}

		// Build the complete WebSocket URL with endpoint
		wsURL, err := buildWebSocketURL(finalServerURL)
		if err != nil {
			logger.Fatalf("Failed to build WebSocket URL: %v", err)
		}

		var tunnelHost, tunnelPort string
		if strings.Contains(target, ":") {
			parts := strings.SplitN(target, ":", 2)
			tunnelHost = parts[0]
			tunnelPort = parts[1]
		} else {
			tunnelHost = "localhost"
			tunnelPort = target
		}

		logger.Infof("Tunneling %s:%s ...", tunnelHost, tunnelPort)

		u, err := url.Parse("wss://" + wsURL)
		if err != nil {
			logger.Fatalf("Failed to parse WebSocket URL: %v", err)
		}

		client.StartClient(*u, tunnelHost, tunnelPort, baseURL)
	},
}

func init() {
	connectCmd.Flags().StringVarP(&serverURL, "server-url", "u", "", "Server URL (without WebSocket endpoint, e.g., example.com:443)")
	connectCmd.Flags().StringVarP(&baseURL, "base-endpoint", "e", "", "Base endpoint path to route the tunneled app")
	connectCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
}
