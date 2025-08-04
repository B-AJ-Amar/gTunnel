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

		if !strings.HasPrefix(finalServerURL, "ws://") && !strings.HasPrefix(finalServerURL, "wss://") {
			finalServerURL = "ws://" + finalServerURL
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

		u, err := url.Parse(finalServerURL)
		if err != nil {
			logger.Fatalf("Invalid server URL: %v", err)
		}

		logger.Infof("Connecting to server at %s...", finalServerURL)
		logger.Infof("Tunneling %s:%s...", tunnelHost, tunnelPort)

		client.StartClient(*u, tunnelHost, tunnelPort, baseURL)
	},
}

func init() {
	connectCmd.Flags().StringVarP(&serverURL, "server-url", "u", "", "Server WebSocket URL to connect to")
	connectCmd.Flags().StringVarP(&baseURL, "base-endpoint", "e", "", "Base endpoint path to route the tunneled app")
	connectCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
}
