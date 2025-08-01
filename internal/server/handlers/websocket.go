package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
	"github.com/gorilla/websocket"
)

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

func SaveTunnel(id, baseURL string, conn *websocket.Conn, connections map[string]*models.ServerTunnelConn, connMu *sync.Mutex) *models.ServerTunnelConn {
	tunnel := &models.ServerTunnelConn{
		ID:         id,
		Conn:       conn,
		ResponseCh: make(chan []byte),
		BaseURL:    baseURL,
	}

	connMu.Lock()
	connections[id] = tunnel
	connMu.Unlock()

	log.Printf("New connection established: %s", id)
	return tunnel
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
