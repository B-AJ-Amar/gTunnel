package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
)

// PathTunnelRouter routes the HTTP request to the appropriate ServerTunnelConn based on the base URL.
// It expects the URL path to be in the format /app-123/a/c/b/..., extracts "app-123" and looks up the connection.
func PathTunnelRouter(r *http.Request, connections map[string]*models.ServerTunnelConn) *models.ServerTunnelConn {
	parts := strings.Split(r.URL.Path, "/")
	log.Println("PathTunnelRouter: Received request with path:", r.URL.Path)
	if len(parts) < 2 {
		return nil
	}
	appID := parts[1]
	log.Printf("PathTunnelRouter: Extracted appID: %s", appID)

	// ? i will use an indexed db later
	for _, conn := range connections {
		if conn.BaseURL == appID {
			return conn
		}
	}
	return nil
}

// TODO : subdomain router