package client

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/B-AJ-Amar/gTunnel/internal/client/handlers"
	"github.com/B-AJ-Amar/gTunnel/internal/client/models"
	"github.com/B-AJ-Amar/gTunnel/internal/client/repositories"
	"github.com/B-AJ-Amar/gTunnel/internal/protocol"
	"github.com/gorilla/websocket"
)

var (
	connections = make(map[string]*models.ClientTunnelConn)
	connMu      sync.Mutex
)

// authenticate handles the authentication process with the server
func authenticate(wsURL url.URL, accessToken, baseURL string) (*models.ClientTunnelConn, error) {
	log.Println("Connecting to WebSocket server at:", wsURL.String())
	
	// Connect to the WebSocket server (no query parameters needed)
	conn, _, err := websocket.DefaultDialer.Dial(wsURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("dial error: %w", err)
	}

	

	// Create authentication request
	authRequest := protocol.AuthRequestMessage{
		AccessToken: accessToken,
		BaseURL:     baseURL,
	}

	// Send authentication request
	authMessage, err := protocol.NewSocketMessage("", protocol.MessageTypeAuthRequest, authRequest)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to create auth message: %w", err)
	}

	authData, err := protocol.SerializeMessage(authMessage)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to serialize auth message: %w", err)
	}

	err = conn.WriteMessage(websocket.TextMessage, authData)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to send auth request: %w", err)
	}

	log.Println("Authentication request sent, waiting for response...")

	// Wait for authentication response
	_, message, err := conn.ReadMessage()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to read auth response: %w", err)
	}

	// Deserialize the response
	var socketMessage protocol.SocketMessage
	err = protocol.DeserializeMessage(message, &socketMessage)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to deserialize auth response: %w", err)
	}

	// Check if it's an auth response
	if socketMessage.Type != protocol.MessageTypeAuthResponse {
		conn.Close()
		return nil, fmt.Errorf("expected auth response, got message type: %d", socketMessage.Type)
	}

	// Parse auth response
	var authResponse protocol.AuthResponseMessage
	err = protocol.DeserializeMessage(socketMessage.Payload, &authResponse)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to parse auth response: %w", err)
	}

	// Check if authentication was successful
	if !authResponse.Success {
		conn.Close()
		return nil, fmt.Errorf("authentication failed: %s", authResponse.Message)
	}

	if authResponse.ID == nil {
		conn.Close()
		return nil, fmt.Errorf("authentication succeeded but no ID provided")
	}

	// Create tunnel connection
	tunnel := &models.ClientTunnelConn{
		ID:   *authResponse.ID,
		Conn: conn,
	}

	log.Printf("Authentication successful. Connection ID: %s", *authResponse.ID)
	return tunnel, nil
}

// WsClientHandler handles the WebSocket communication after authentication
func WsClientHandler(tunnel *models.ClientTunnelConn, tunnelHost, tunnelPort string) {
	conn := tunnel.Conn
	id := tunnel.ID

	// Set tunnel host and port
	tunnel.Host = tunnelHost
	tunnel.Port = tunnelPort

	log.Printf("Starting WebSocket handler for connection: %s", id)

	// Store the connection
	connMu.Lock()
	connections[id] = tunnel
	connMu.Unlock()

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
			log.Printf("[%s] HTTP response Sent Successfully", id)

		default:
			log.Printf("[%s] Unknown message type: %d", id, socketMessage.Type)
			continue
		}
	}
}

// StartClient initiates the client connection with authentication
func StartClient(wsURL url.URL, tunnelHost, tunnelPort string, baseURL string) {
	// Load access token from client config
	configRepo := repositories.NewClientConfigRepo()
	if err := configRepo.InitConfig(); err != nil {
		log.Printf("Warning: Failed to initialize config: %v", err)
	}
	
	config, err := configRepo.Load()
	if err != nil {
		log.Printf("Warning: Failed to load config: %v", err)
	}
	
	accessToken := ""
	if config != nil {
		accessToken = config.AccessToken
	}
	
	// Authenticate and get tunnel connection
	tunnel, err := authenticate(wsURL, accessToken, baseURL)
	if err != nil {
		log.Fatal("Authentication failed:", err)
	}
	
	// Start the WebSocket handler
	WsClientHandler(tunnel, tunnelHost, tunnelPort)
}
