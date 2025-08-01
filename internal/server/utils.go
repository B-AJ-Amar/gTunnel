package server

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
)

// ValidateAndExtractParams validates the request parameters and extracts id and baseURL
func ValidateAndExtractParams(r *http.Request) (string, string, error) {
	id := r.URL.Query().Get("id")
	baseURL := r.URL.Query().Get("base_url")

	log.Printf("WebSocket connection request: id=%s, baseURL=%s", id, baseURL)

	if id == "" {
		return "", "", &ValidationError{Message: "Missing 'id' query parameter", StatusCode: http.StatusBadRequest}
	}

	return id, baseURL, nil
}

// GenerateBaseURL generates a base URL if one is not provided
func GenerateBaseURL(baseURL, id string) string {
	if baseURL == "" {
		baseURL = "/app-" + strings.Split(id, "-")[4] // simple example, can be improved
		// baseURL = "/app"// !! JUST FOR DEV
		log.Printf("Generated base URL: %s", baseURL)
	}
	return baseURL
}

// ValidateBaseURLAvailability checks if the baseURL is already in use
func ValidateBaseURLAvailability(baseURL string, connections map[string]*models.ServerTunnelConn, connMu *sync.Mutex) error {
	connMu.Lock()
	defer connMu.Unlock()

	for _, t := range connections {
		if t.BaseURL == baseURL {
			return &ValidationError{Message: "Base URL already in use", StatusCode: http.StatusConflict}
		}
	}
	return nil
}

func ExtractPath(urlPath string) (string, string, error) {
	parts := strings.Split(urlPath, "/")
	log.Println("ExtractPath: Received path:", urlPath)
	if len(parts) < 2 {
		return "", "", errors.New("invalid path")
	}
	appID := "/" + parts[1]
	rest := "/" + strings.Join(parts[2:], "/")
	return appID, rest, nil
}

// PathTunnelRouter routes the HTTP request to the appropriate ServerTunnelConn based on the base URL.
// It expects the URL path to be in the format /app-123/a/c/b/..., extracts "app-123" and looks up the connection.
func PathTunnelRouter(r *http.Request, connections map[string]*models.ServerTunnelConn) (*models.ServerTunnelConn, string, string) {
	appID, endpoint, err := ExtractPath(r.URL.Path)

	if err != nil {
		log.Printf("PathTunnelRouter: Error extracting path: %v", err)
		return nil, "", ""
	}

	log.Println("PathTunnelRouter: Received request with path:", r.URL.Path)
	log.Printf("PathTunnelRouter: Extracted appID: %s", appID)

	for _, conn := range connections {
		if conn.BaseURL == appID {
			return conn, appID, endpoint
		}
	}
	return nil, "", ""
}

// TODO : subdomain router
