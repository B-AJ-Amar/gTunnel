package sec

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/B-AJ-Amar/gTunnel/internal/protocol"
	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
	"github.com/B-AJ-Amar/gTunnel/internal/server/repositories"
	"github.com/B-AJ-Amar/gTunnel/internal/server/utils"
)

// TODO : Handle BaseURL

func HandleAuthMessage(msg []byte, tunnel *models.ServerTunnelConn, connections map[string]*models.ServerTunnelConn, connMu *sync.Mutex, authenticating map[string]*models.ServerTunnelConn, authMu *sync.Mutex) (bool, error) {
	var socketMsg protocol.SocketMessage
	if err := protocol.DeserializeMessage(msg, &socketMsg); err != nil {
		log.Printf("Failed to deserialize auth message: %v", err)
		return false, err
	}

	switch socketMsg.Type {
	case protocol.MessageTypeAuthRequest:
		var authRequest protocol.AuthRequestMessage
		if err := protocol.DeserializeMessage(socketMsg.Payload, &authRequest); err != nil {
			// Failed to deserialize auth request
			return false, err
		}
		
		// Handle baseURL from auth request
		baseURL := authRequest.BaseURL
		if len(baseURL) > 0 && baseURL[0] == '/' {
			baseURL = baseURL[1:]
		}

		if baseURL == "" {
			// Generate base URL if not provided
			baseURL = utils.GenerateBaseURL("", tunnel.ID)
		}
		
		// Validate baseURL availability
		if err := utils.ValidateBaseURLAvailability(baseURL, connections, connMu); err != nil {
			log.Printf("BaseURL validation failed: %v", err)
			HandleAuthFailure(tunnel, authenticating, authMu)
			return false, err
		}
		
		// Set the baseURL in the tunnel
		tunnel.BaseURL = baseURL
		
		// Handle authentication request
		success, err := AuthenticateTunnel(&authRequest)
		if err != nil {
			log.Printf("Authentication failed: %v", err)
			HandleAuthFailure(tunnel, authenticating, authMu)
			return false, err
		}
		if success {
			HandleAuthSuccess(tunnel, connections, connMu, authenticating, authMu)
			return true, nil
		}
		HandleAuthFailure(tunnel, authenticating, authMu)
		return false, fmt.Errorf("authentication failed")
	default:
		log.Printf("Unknown auth message type: %v", socketMsg.Type)
	}
	return false, fmt.Errorf("unknown auth message type: %v", socketMsg.Type)
}

func AuthenticateTunnel(authReq *protocol.AuthRequestMessage) (bool, error) {
	configRepo := repositories.NewServerConfigRepo()

	config, err := configRepo.Load()
	if err != nil {
		// Return error instead of crashing the server
		return false, fmt.Errorf("failed to load config: %w", err)
	}
	log.Printf("SERVER Loaded config: %+v", config)

	// Compare with the expected token
	if authReq.AccessToken != config.AccessToken {
		return false, fmt.Errorf("invalid access_token")
	}

	// Authentication successful
	return true, nil
}
func HandleAuthSuccess(tunnel *models.ServerTunnelConn, connections map[string]*models.ServerTunnelConn, connMu *sync.Mutex, authenticating map[string]*models.ServerTunnelConn, authMu *sync.Mutex) {
	log.Printf("[%s] Authentication successful", tunnel.ID)

	// Add to connections map
	connMu.Lock()
	connections[tunnel.ID] = tunnel
	connMu.Unlock()

	// Remove from authenticating map - always use mutex
	authMu.Lock()
	delete(authenticating, tunnel.ID)
	authMu.Unlock()

	// Prepare and send the auth response
	authResponse := &protocol.AuthResponseMessage{
		ID:      &tunnel.ID,
		Success: true,
		Message: "Authentication successful",
	}

	serializedPayload, err := protocol.SerializeMessage(authResponse)
	if err != nil {
		log.Printf("[%s] Failed to serialize auth response: %v", tunnel.ID, err)
		return
	}

	socketMsg := &protocol.SocketMessage{
		Type:    protocol.MessageTypeAuthResponse,
		Payload: serializedPayload,
	}

	if err := tunnel.Conn.WriteJSON(socketMsg); err != nil {
		log.Printf("[%s] Failed to send auth success response: %v", tunnel.ID, err)
	}
}

func HandleAuthFailure(tunnel *models.ServerTunnelConn, authenticating map[string]*models.ServerTunnelConn, authMu *sync.Mutex) {
	authMu.Lock()
	delete(authenticating, tunnel.ID)
	authMu.Unlock()

	tunnel.Conn.Close()
	log.Printf("[%s] Connection closed due to authentication failure", tunnel.ID)
}

func HandleWSAuth(tunnel *models.ServerTunnelConn, r *http.Request, authenticating map[string]*models.ServerTunnelConn, authMu *sync.Mutex, connections map[string]*models.ServerTunnelConn, connMu *sync.Mutex) (bool, error) {
	done := make(chan struct{})
	var msg []byte
	var readErr error

	// Read auth message with timeout
	go func() {
		defer close(done)
		_, message, err := tunnel.Conn.ReadMessage()
		if err != nil {
			readErr = err
			return
		}
		msg = message
	}()

	select {
	case <-done:
		if readErr != nil {
			log.Printf("[%s] Read error during auth: %v", tunnel.ID, readErr)
			HandleAuthFailure(tunnel, authenticating, authMu)
			return false, readErr
		}

		log.Printf("[%s] Received auth message: %s", tunnel.ID, msg)
		success, err := HandleAuthMessage(msg, tunnel, connections, connMu, authenticating, authMu)
		if err != nil {
			log.Printf("[%s] Error handling auth message: %v", tunnel.ID, err)
			return false, err
		}

		return success, nil

	case <-time.After(10 * time.Second):
		log.Printf("[%s] Authentication timeout - no message received in 10 seconds", tunnel.ID)
		HandleAuthFailure(tunnel, authenticating, authMu)
		return false, fmt.Errorf("authentication timeout")
	}
}
