package client

import (
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/B-AJ-Amar/gTunnel/internal/client/handlers"
	"github.com/B-AJ-Amar/gTunnel/internal/client/models"
	"github.com/B-AJ-Amar/gTunnel/internal/protocol"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)
var (
    connections = make(map[string]*models.ClientTunnelConn)
    connMu      sync.Mutex
)


func WsClientHandler(WsUrl url.URL) {
    // Connect to the WebSocket server
	id := uuid.New().String()
	// Add the id as a query param to the WebSocket URL
	q := WsUrl.Query()
	q.Set("id", id)
	WsUrl.RawQuery = q.Encode()
	log.Println("Connecting to WebSocket server at:", WsUrl.String())
	conn, _, err := websocket.DefaultDialer.Dial(WsUrl.String(), nil)
    if err != nil {
        log.Fatal("Dial error:", err)
    }

	tunnel := &models.ClientTunnelConn{
		ID:         id,
		Conn:       conn,
		Port:       "3000", // for test
		Host:       "localhost",
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

		// Deserialize the message
		var socketMessage protocol.SocketMessage
		err = protocol.DeserializeMessage(message, &socketMessage)
		if err != nil {
			log.Printf("[%s] Error deserializing message: %v", id, err)
			continue
		}
		log.Printf("[%s] Message type: %d", id, socketMessage.Type)

		// Handle the message based on its type
		switch socketMessage.Type {

			case protocol.MessageTypeHTTPRequest:
				err := handlers.ClientHTTPRequestHandler(socketMessage, tunnel) 
				if err != nil {
					log.Printf("[%s] Error handling HTTP request: %v", id, err)
					continue
				}
				log.Printf("[%s] HTTP response Sent Successfuly", id)

			default:
				log.Printf("[%s] Unknown message type: %d", id, socketMessage.Type)
				continue
		}


	}


}

func StartClient(WsUrl url.URL) {
	WsClientHandler(WsUrl)
}
