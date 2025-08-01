package handlers

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/B-AJ-Amar/gTunnel/internal/protocol"
	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
	"github.com/gorilla/websocket"
)

// HTTPToWebSocketHandler handles HTTP requests and forwards them through WebSocket handler Using channels
func HTTPToWebSocketHandler(w http.ResponseWriter, r *http.Request, pathTunnelRouter func(*http.Request, map[string]*models.ServerTunnelConn) (*models.ServerTunnelConn, string, string), connections map[string]*models.ServerTunnelConn) {
	tunnel, appID, _ := pathTunnelRouter(r, connections)

	if tunnel == nil {
		http.Error(w, "No tunnel connected", http.StatusServiceUnavailable)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Read error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	new_url := strings.Replace(r.URL.String(), appID, "", 1)
	reqMsg := protocol.HTTPRequestMessage{
		Method:  r.Method,
		URL:     new_url,
		Headers: map[string]string{},
		Body:    body,
	}
	for name, values := range r.Header {
		if len(values) > 0 {
			reqMsg.Headers[name] = values[0]
		}
	}

	payload, err := protocol.SerializeMessage(reqMsg)
	if err != nil {
		http.Error(w, "Serialization error", http.StatusInternalServerError)
		return
	}

	fullMsg := protocol.SocketMessage{
		Type:    protocol.MessageTypeHTTPRequest,
		Payload: payload,
	}

	encoded, err := protocol.SerializeMessage(fullMsg)
	if err != nil {
		http.Error(w, "Message encoding failed", http.StatusInternalServerError)
		return
	}

	if err := tunnel.Conn.WriteMessage(websocket.TextMessage, encoded); err != nil {
		http.Error(w, "Tunnel write failed", http.StatusBadGateway)
		return
	}
	log.Println("HTTPToWebSocketHandler: message sent to tunnel")

	// Wait for response (with timeout)
	select {
	case responseData := <-tunnel.ResponseCh:
		var responseMsg protocol.SocketMessage
		if err := protocol.DeserializeMessage(responseData, &responseMsg); err != nil {
			http.Error(w, "Invalid tunnel response", http.StatusInternalServerError)
			return
		}

		if responseMsg.Type != protocol.MessageTypeHTTPResponse {
			http.Error(w, "Unexpected message type", http.StatusInternalServerError)
			return
		}

		var httpResp protocol.HTTPResponseMessage
		if err := protocol.DeserializeMessage(responseMsg.Payload, &httpResp); err != nil {
			http.Error(w, "Invalid response payload", http.StatusInternalServerError)
			return
		}

		for name, value := range httpResp.Headers {
			w.Header().Set(name, value)
		}
		w.WriteHeader(httpResp.StatusCode)
		w.Write(httpResp.Body)

	case <-time.After(10 * time.Second):
		http.Error(w, "Tunnel response timeout", http.StatusGatewayTimeout)
	}
}
