package utils

import (
	"net/http"

	"github.com/B-AJ-Amar/gTunnel/internal/logger"
	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
)

// PathTunnelRouter routes the HTTP request to the appropriate ServerTunnelConn based on the base URL.
// It expects the URL path to be in the format /app-123/a/c/b/..., extracts "app-123" and looks up the connection.
func PathTunnelRouter(r *http.Request, connections map[string]*models.ServerTunnelConn) (*models.ServerTunnelConn, string, string) {
	appID, endpoint, err := ExtractPath(r.URL.Path)

	logger.Debug("connections : ")
	for _, conn := range connections {
		logger.Debugf("PathTunnelRouter: Connection BaseURL: %s", conn.BaseURL)
		logger.Debugf("PathTunnelRouter: Connection ID: %s", conn.ID)
	}

	if err != nil {
		logger.Errorf("PathTunnelRouter: Error extracting path: %v", err)
		return nil, "", ""
	}

	logger.Debug("PathTunnelRouter: Received request with path:", r.URL.Path)
	logger.Debugf("PathTunnelRouter: Extracted appID: %s", appID)

	for _, conn := range connections {
		if conn.BaseURL == appID {
			return conn, appID, endpoint
		}
	}
	return nil, "", ""
}

// TODO : subdomain router
