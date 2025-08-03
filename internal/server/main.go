package server

import (
	"log"
	"net/http"
	"sync"

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
		log.Println("Authentication error:", err)
		conn.Close()
		return
	}
	if !success {
		log.Println("Authentication failed for tunnel:", id)
		conn.Close()
		return
	}
	log.Println("authentication SUCCESS")

	// Handle WebSocket messages
	handlers.HandleWSMessages(tunnel)

	handlers.TunnelCleanup(id, conn, connections, &connMu)()
}

func httpToWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// i will add more routers later
	handlers.HTTPToWebSocketHandler(w, r, utils.PathTunnelRouter, connections)
}

func StartServer(addr string) {
	r := chi.NewRouter()
	r.Get("/___gTl___/ws", wsHandler)
	r.NotFound(httpToWebSocketHandler)

	log.Println("Server listening on", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
