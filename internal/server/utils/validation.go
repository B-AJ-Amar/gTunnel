package utils

import (
	"net/http"
	"strings"
	"sync"

	"github.com/B-AJ-Amar/gTunnel/internal/logger"
	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
)

func GenerateBaseURL(baseURL, id string) string {
	if baseURL == "" {
		baseURL = "app-" + strings.Split(id, "-")[4] // simple example, can be improved
		// baseURL = "app"// !! JUST FOR DEV
		logger.Debugf("Generated base URL: %s", baseURL)
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

	logger.Debugf("WebSocket connection request: id=%s, baseURL=%s", id, baseURL)

	if id == "" {
		return "", "", &ValidationError{Message: "Missing 'id' query parameter", StatusCode: http.StatusBadRequest}
	}

	return id, baseURL, nil
}

func ExtractPath(urlPath string) (string, string, error) {
	urlPath = strings.TrimPrefix(urlPath, "/")
	logger.Debug("Extracting path from URL:", urlPath)

	pathSegments := strings.SplitN(urlPath, "/", 2)
	if len(pathSegments) == 0 || pathSegments[0] == "" {
		return "", "", &ValidationError{Message: "Invalid path: missing appID", StatusCode: http.StatusBadRequest}
	}

	appID := pathSegments[0]
	remainingPath := ""
	if len(pathSegments) > 1 {
		remainingPath = "/" + pathSegments[1]
	}

	logger.Debugf("Extracted appID: %s, remainingPath: %s", appID, remainingPath)
	return appID, remainingPath, nil
}
