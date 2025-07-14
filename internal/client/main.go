package main

import (
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/B-AJ-Amar/gTunnel/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)
var (
    connections = make(map[string]*models.ClientTunnelConn)
    connMu      sync.Mutex
)

func WsClientHandler(WsUrl url.URL) {
    // Connect to the WebSocket server
    conn, _, err := websocket.DefaultDialer.Dial(WsUrl.String(), nil)
    if err != nil {
        log.Fatal("Dial error:", err)
    }

    id := uuid.New().String()
	tunnel := &models.ClientTunnelConn{
		ID:         id,
		Conn:       conn,
		ResponseCh: make(chan []byte),
	}
    log.Printf("New connection established: %s", id)

	connMu.Lock()
	connections[id] = tunnel
	connMu.Unlock()

    conn.WriteMessage(websocket.TextMessage, []byte("Hello from client!"))
    
    defer func() {
		connMu.Lock()
		delete(connections, id)
		connMu.Unlock()
		conn.Close()
		log.Printf("Connection closed: %s", id)
	}()

    conn.SetPongHandler(func(appData string) error {
		log.Println("Received pong:", appData)
		return nil
	})

	// Ping loop
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			<-ticker.C
			log.Println("Sending ping to", id)
			err := conn.WriteMessage(websocket.PingMessage, []byte("ping"))
			if err != nil {
				log.Println("Ping failed, closing connection:", err)
				conn.Close()
				return
			}
		}
	}()
    

    // WebSocket read loop
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Printf("[%s] Received: %s", id, message)

		// Send to response channel (non-blocking)
		select {
		case tunnel.ResponseCh <- message:
		default:
			log.Printf("[%s] WARNING: Dropping message - no listener waiting", id)
		}
	}


}

func main() {
    var WsUrl = url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}

    WsClientHandler(WsUrl)
}
