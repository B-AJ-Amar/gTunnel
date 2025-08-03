package utils

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
)

func GenerateBaseURL(baseURL, id string) string {
	if baseURL == "" {
		baseURL = "app-" + strings.Split(id, "-")[4] // simple example, can be improved
		// baseURL = "app"// !! JUST FOR DEV
		log.Printf("Generated base URL: %s", baseURL)
	}
	return baseURL
}

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

func ValidateAndExtractParams(r *http.Request) (string, string, error) {
	id := r.URL.Query().Get("id")
	baseURL := r.URL.Query().Get("base_url")

	log.Printf("WebSocket connection request: id=%s, baseURL=%s", id, baseURL)

	if id == "" {
		return "", "", &ValidationError{Message: "Missing 'id' query parameter", StatusCode: http.StatusBadRequest}
	}

	return id, baseURL, nil
}

func ExtractPath(urlPath string) (string, string, error) {
	if strings.HasPrefix(urlPath, "/") {
		urlPath = urlPath[1:]
	}

	parts := strings.SplitN(urlPath, "?", 2)
	path := parts[0]
	query := ""
	if len(parts) > 1 {
		query = parts[1]
	}

	return path, query, nil
}
