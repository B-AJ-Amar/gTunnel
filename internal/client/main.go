package client

import (
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/B-AJ-Amar/gTunnel/internal/models"
	"github.com/B-AJ-Amar/gTunnel/internal/protocol"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)
var (
    connections = make(map[string]*models.ClientTunnelConn)
    connMu      sync.Mutex
)


func ClientHTTPResponseHandler(message protocol.HTTPResponseMessage, conn *websocket.Conn){
	
}

func ClientHTTPRequestHandler(message protocol.HTTPRequestMessage, conn *websocket.Conn){

}

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
				var httpRequest protocol.HTTPRequestMessage
				err = protocol.DeserializeMessage(socketMessage.Payload, &httpRequest)
				if err != nil {
					log.Printf("[%s] Error deserializing HTTP request: %v", id, err)
					continue
				}
				log.Printf("[%s] HTTP Request: %s %s", id, httpRequest.Method, httpRequest.URL)
				ClientHTTPRequestHandler(httpRequest, conn)
			case protocol.MessageTypeHTTPResponse:
				var httpResponse protocol.HTTPResponseMessage
				err = protocol.DeserializeMessage(socketMessage.Payload, &httpResponse)
				if err != nil {
					log.Printf("[%s] Error deserializing HTTP response: %v", id, err)
					continue
				}
				log.Printf("[%s] HTTP Response: %d", id, httpResponse.StatusCode)
				ClientHTTPResponseHandler(httpResponse, conn)
			default:
				log.Printf("[%s] Unknown message type: %d", id, socketMessage.Type)
				continue
		}


		// Send to response channel (non-blocking)
		select {
		case tunnel.ResponseCh <- message:
		default:
			log.Printf("[%s] WARNING: Dropping message - no listener waiting", id)
		}
	}


}

func StartClient(WsUrl url.URL) {
	WsClientHandler(WsUrl)
}
