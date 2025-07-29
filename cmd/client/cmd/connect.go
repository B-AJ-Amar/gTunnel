package cmd

import (
	"fmt"
	"net/url"

	"github.com/B-AJ-Amar/gTunnel/internal/client"
	"github.com/spf13/cobra"
)

var (
	serverHost string
	serverPort int
	basePath   string
	baseUrl    string
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a gTunnel server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Connecting to server at %s:%d...\n", serverHost, serverPort)
		
		u := url.URL{
			Scheme: "ws",
			Host:   fmt.Sprintf("%s:%d", serverHost, serverPort),
			Path:   basePath,
		}
		
		if baseUrl != "" {
			q := u.Query()
			q.Set("baseUrl", baseUrl)
			u.RawQuery = q.Encode()
		}
		
		client.StartClient(u)
	},
}

func init() {
	connectCmd.Flags().StringVarP(&serverHost, "host", "H", "localhost", "Server host to connect to")
	connectCmd.Flags().IntVarP(&serverPort, "port", "p", 8080, "Server port to connect to")
	connectCmd.Flags().StringVarP(&basePath, "path", "P", "/___gTl___/ws", "WebSocket path on the server")
	connectCmd.Flags().StringVarP(&baseUrl, "base-url", "b", "/base_url", "Base URL for the tunnel")
}
