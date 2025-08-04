package server

import (
	"net/http"
	"sync"

	"github.com/B-AJ-Amar/gTunnel/internal/logger"
	"github.com/B-AJ-Amar/gTunnel/internal/server/handlers"
	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
	"github.com/B-AJ-Amar/gTunnel/internal/server/sec"
	"github.com/B-AJ-Amar/gTunnel/internal/server/utils"
	"github.com/go-chi/chi/v5"
)

var (
	// authenticating is a map to keep track of connections that are in the process of authenticating
	// after authentication, they will be moved to the connections map
	authenticating = make(map[string]*models.ServerTunnelConn)
	authMu         sync.Mutex

	connections = make(map[string]*models.ServerTunnelConn)
	connMu      sync.Mutex
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Generate a unique ID for this connection

	// Establish WebSocket connection
	conn, err := handlers.EstablishWSConn(w, r)
	if err != nil {
		return
	}

	tunnel := handlers.SaveTunnel(conn, authenticating, &connMu)
	id := tunnel.ID

	success, err := sec.HandleWSAuth(tunnel, r, authenticating, &authMu, connections, &connMu)

	if err != nil {
		logger.Errorf("Authentication error: %v", err)
		conn.Close()
		return
	}
	if !success {
		logger.Warnf("Authentication failed for tunnel: %s", id)
		conn.Close()
		return
	}
	logger.Info("Authentication successful")

	handlers.HandleWSMessages(tunnel)

	handlers.TunnelCleanup(id, conn, connections, &connMu)()
}

func httpToWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: add more routers later
	handlers.HTTPToWebSocketHandler(w, r, utils.PathTunnelRouter, connections)
}

func StartServer(addr string) {
	r := chi.NewRouter()
	r.Get("/___gTl___/ws", wsHandler)
	r.NotFound(httpToWebSocketHandler)

	logger.Infof("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
