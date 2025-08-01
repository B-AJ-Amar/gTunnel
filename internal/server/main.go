package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/B-AJ-Amar/gTunnel/internal/server/handlers"
	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
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
	// Extract and validate parameters
	id, baseURL, err := ValidateAndExtractParams(r)
	if err != nil {
		if validationErr, ok := err.(*ValidationError); ok {
			http.Error(w, validationErr.Message, validationErr.StatusCode)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Generate base URL if needed
	baseURL = GenerateBaseURL(baseURL, id)

	if err := ValidateBaseURLAvailability(baseURL, connections, &connMu); err != nil {
		if validationErr, ok := err.(*ValidationError); ok {
			http.Error(w, validationErr.Message, validationErr.StatusCode)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Establish WebSocket connection
	conn, err := handlers.EstablishWSConn(w, r)
	if err != nil {
		return
	}

	tunnel := handlers.SaveTunnel(id, baseURL, conn, connections, &connMu)

	defer handlers.TunnelCleanup(id, conn, connections, &connMu)()

	// Handle WebSocket messages
	handlers.HandleWSMessages(tunnel)
}

func httpToWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	handlers.HTTPToWebSocketHandler(w, r, PathTunnelRouter, connections)
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
