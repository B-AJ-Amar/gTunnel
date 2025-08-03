package utils

import (
	"log"
	"net/http"

	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
)

// PathTunnelRouter routes the HTTP request to the appropriate ServerTunnelConn based on the base URL.
// It expects the URL path to be in the format /app-123/a/c/b/..., extracts "app-123" and looks up the connection.
func PathTunnelRouter(r *http.Request, connections map[string]*models.ServerTunnelConn) (*models.ServerTunnelConn, string, string) {
	appID, endpoint, err := ExtractPath(r.URL.Path)

	log.Println("connections : ")
	for _, conn := range connections {
		log.Printf("PathTunnelRouter: Connection BaseURL: %s", conn.BaseURL)
		log.Printf("PathTunnelRouter: Connection ID: %s", conn.ID)
	}

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