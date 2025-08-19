package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/B-AJ-Amar/gTunnel/internal/client/repositories"
	"github.com/B-AJ-Amar/gTunnel/internal/logger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	verbose bool
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display client connection status",
	Long: `Display the current connection status to the gTunnel server.

Examples:
  gtc status           # Show basic connection status
  gtc status -v        # Show detailed connection information`,
	Run: func(cmd *cobra.Command, args []string) {
		configRepo := repositories.NewClientConfigRepo()

		// Initialize config before loading
		if err := configRepo.InitConfig(); err != nil {
			logger.Fatalf("Failed to initialize config: %v", err)
		}

		// Load config to get server URL
		config, err := configRepo.Load()
		if err != nil {
			logger.Fatalf("Failed to load config: %v", err)
		}

		if config.ServerURL == "" {
			printError("Not configured")
			fmt.Println("Use 'gtc config --set-url <server-url>' to configure the server URL")
			return
		}

		if verbose {
			fmt.Printf("Server URL: %s\n", config.ServerURL)
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		}

		// Check server health
		status, isConnected := checkServerHealth(config.ServerURL)
		if isConnected {
			printSuccess(status)
		} else {
			printError(status)
		}
	},
}

func checkServerHealth(serverURL string) (string, bool) {
	// Build health check URL - convert ws/wss to http/https
	healthURL := buildHealthURL(serverURL)
	
	if verbose {
		fmt.Printf("Checking server health at: %s\n", healthURL)
	}
	
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Send GET request to health endpoint
	start := time.Now()
	resp, err := client.Get(healthURL)
	duration := time.Since(start)
	
	if err != nil {
		if verbose {
			fmt.Printf("Response time: %v\n", duration)
		}
		return fmt.Sprintf("Not connected (Error: %v)", err), false
	}
	defer resp.Body.Close()

	if verbose {
		fmt.Printf("Response time: %v\n", duration)
		fmt.Printf("HTTP Status: %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// Check response status
	if resp.StatusCode == http.StatusOK {
		if verbose {
			return fmt.Sprintf("Connected (Server is healthy) - Response time: %v", duration), true
		}
		return "Connected (Server is healthy)", true
	}
	
	return fmt.Sprintf("Not connected (Server returned status: %d)", resp.StatusCode), false
}

func buildHealthURL(serverURL string) string {
	// Convert WebSocket URL to HTTP URL for health check
	httpURL := serverURL
	
	// Add protocol if not present
	if serverURL != "" && !hasProtocol(serverURL) {
		httpURL = "http://" + serverURL
	}
	
	// Convert ws:// to http:// and wss:// to https://
	if len(httpURL) >= 5 && httpURL[:5] == "ws://" {
		httpURL = "http://" + httpURL[5:]
	} else if len(httpURL) >= 6 && httpURL[:6] == "wss://" {
		httpURL = "https://" + httpURL[6:]
	}
	
	// Add health endpoint
	if httpURL[len(httpURL)-1] != '/' {
		httpURL += "/"
	}
	httpURL += "___gTl___/health"
	
	return httpURL
}

func hasProtocol(url string) bool {
	return len(url) > 7 && (url[:7] == "http://" || url[:8] == "https://" || url[:5] == "ws://" || url[:6] == "wss://")
}

// Color printing functions
func printSuccess(message string) {
	green := color.New(color.FgGreen, color.Bold)
	fmt.Print("Client connection status: ")
	green.Println(message)
}

func printError(message string) {
	red := color.New(color.FgRed, color.Bold)
	fmt.Print("Client connection status: ")
	red.Println(message)
}

func init() {
	statusCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show detailed connection information")
}
