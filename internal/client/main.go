package client

import (
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/B-AJ-Amar/gTunnel/internal/client/handlers"
	"github.com/B-AJ-Amar/gTunnel/internal/client/models"
	"github.com/B-AJ-Amar/gTunnel/internal/client/repositories"
	"github.com/B-AJ-Amar/gTunnel/internal/logger"
	"github.com/B-AJ-Amar/gTunnel/internal/protocol"
	"github.com/gorilla/websocket"
)

var (
	connections = make(map[string]*models.ClientTunnelConn)
	connMu      sync.Mutex
)

func authenticate(wsURL url.URL, accessToken, baseURL string) (*models.ClientTunnelConn, error) {
	
	conn, _, err := websocket.DefaultDialer.Dial(wsURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("dial error: %w", err)
	}

	authRequest := protocol.AuthRequestMessage{
		AccessToken: accessToken,
		BaseURL:     baseURL,
	}

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

	logger.Info("Authentication request sent, waiting for response...")

	_, message, err := conn.ReadMessage()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to read auth response: %w", err)
	}

	var socketMessage protocol.SocketMessage
	err = protocol.DeserializeMessage(message, &socketMessage)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to deserialize auth response: %w", err)
	}

	if socketMessage.Type != protocol.MessageTypeAuthResponse {
		conn.Close()
		return nil, fmt.Errorf("expected auth response, got message type: %d", socketMessage.Type)
	}

	var authResponse protocol.AuthResponseMessage
	err = protocol.DeserializeMessage(socketMessage.Payload, &authResponse)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to parse auth response: %w", err)
	}

	if !authResponse.Success {
		conn.Close()
		return nil, fmt.Errorf("authentication failed: %s", authResponse.Message)
	}

	if authResponse.ID == nil {
		conn.Close()
		return nil, fmt.Errorf("authentication succeeded but no ID provided")
	}

	tunnel := &models.ClientTunnelConn{
		ID:   *authResponse.ID,
		Conn: conn,
	}

	httpURL := fmt.Sprintf("http://%s%s", wsURL.Host, baseURL)
	logger.Infof("Authentication successful. Connection ID: %s", *authResponse.ID)
	logger.Infof("Tunnel URL: %s", httpURL)
	return tunnel, nil
}

func WsClientHandler(tunnel *models.ClientTunnelConn, tunnelHost, tunnelPort string) {
	conn := tunnel.Conn
	id := tunnel.ID

	tunnel.Host = tunnelHost
	tunnel.Port = tunnelPort

	logger.Infof("Starting WebSocket handler for connection: %s", id)

	connMu.Lock()
	connections[id] = tunnel
	connMu.Unlock()

	defer func() {
		connMu.Lock()
		delete(connections, id)
		connMu.Unlock()
		conn.Close()
		logger.Infof("Connection closed: %s", id)
	}()

	conn.SetPongHandler(func(appData string) error {
		logger.Debugf("Received pong: %s", appData)
		return nil
	})

	// Ping loop
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			<-ticker.C
			logger.Debugf("Sending ping to %s", id)
			err := conn.WriteMessage(websocket.PingMessage, []byte("ping"))
			if err != nil {
				logger.Errorf("Ping failed, closing connection: %v", err)
				conn.Close()
				return
			}
		}
	}()

	// WebSocket read loop
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			logger.Errorf("Read error: %v", err)
			break
		}

		logger.Debugf("[%s] Received: %s", id, message)

		var socketMessage protocol.SocketMessage
		err = protocol.DeserializeMessage(message, &socketMessage)
		if err != nil {
			logger.Errorf("[%s] Error deserializing message: %v", id, err)
			continue
		}
		logger.Debugf("[%s] Message type: %d", id, socketMessage.Type)

		switch socketMessage.Type {

		case protocol.MessageTypeHTTPRequest:
			err := handlers.ClientHTTPRequestHandler(socketMessage, tunnel)
			if err != nil {
				logger.Errorf("[%s] Error handling HTTP request: %v", id, err)
				continue
			}
			logger.Debugf("[%s] HTTP response sent successfully", id)

		default:
			logger.Warnf("[%s] Unknown message type: %d", id, socketMessage.Type)
			continue
		}
	}
}

func StartClient(wsURL url.URL, tunnelHost, tunnelPort string, baseURL string) {
	configRepo := repositories.NewClientConfigRepo()
	if err := configRepo.InitConfig(); err != nil {
		logger.Warnf("Failed to initialize config: %v", err)
	}
	
	config, err := configRepo.Load()
	if err != nil {
		logger.Warnf("Failed to load config: %v", err)
	}
	
	accessToken := ""
	if config != nil {
		accessToken = config.AccessToken
	}
	
	tunnel, err := authenticate(wsURL, accessToken, baseURL)
	if err != nil {
		logger.Fatalf("Authentication failed: %v", err)
	}
	
	WsClientHandler(tunnel, tunnelHost, tunnelPort)
}