package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// TODO : Move Tunnel FUnctions to the repo
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func EstablishWSConn(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return nil, err
	}
	return conn, nil
}

func SaveTunnel(conn *websocket.Conn, authenticating map[string]*models.ServerTunnelConn, connMu *sync.Mutex) *models.ServerTunnelConn {
	id := uuid.New().String()
	tunnel := &models.ServerTunnelConn{
		ID:         id,
		Conn:       conn,
		ResponseCh: make(chan []byte),
		BaseURL:    "",
	}

	connMu.Lock()
	authenticating[id] = tunnel
	connMu.Unlock()

	log.Printf("New connection established: %s", id)
	return tunnel
}

func MoveTunnelToConnections(id string, authenticating map[string]*models.ServerTunnelConn, connections map[string]*models.ServerTunnelConn, connMu *sync.Mutex) {
	connMu.Lock()
	defer connMu.Unlock()

	if tunnel, ok := authenticating[id]; ok {
		delete(authenticating, id)
		connections[id] = tunnel
		log.Printf("Tunnel moved to connections: %s", id)
	}
}

func TunnelCleanup(id string, conn *websocket.Conn, connections map[string]*models.ServerTunnelConn, connMu *sync.Mutex) func() {
	return func() {
		connMu.Lock()
		delete(connections, id)
		connMu.Unlock()
		conn.Close()
		log.Printf("Connection closed: %s", id)
	}
}

func HandleWSMessages(tunnel *models.ServerTunnelConn) {
	for {
		_, message, err := tunnel.Conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Printf("[%s] Received: %s", tunnel.ID, message)

		// Send to response channel (non-blocking)
		select {
		case tunnel.ResponseCh <- message:
		default:
			log.Printf("[%s] WARNING: Dropping message - no listener waiting", tunnel.ID)
		}
	}
}
