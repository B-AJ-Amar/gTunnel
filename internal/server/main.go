package server

import (
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/B-AJ-Amar/gTunnel/internal/protocol"
	"github.com/B-AJ-Amar/gTunnel/internal/server/models"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var (
	connections = make(map[string]*models.ServerTunnelConn)
	connMu      sync.Mutex
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	baseURL := r.URL.Query().Get("base_url")

	log.Printf("WebSocket connection request: id=%s, baseURL=%s", id, baseURL)
	if id == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	// if base url is empty , generate random base url e,.g /app-1
	if baseURL == "" {
		baseURL = "/app-" + strings.Split(id, "-")[4] // simple example, can be improved
		// baseURL = "/app"// !! JUST FOR DEV
		log.Printf("Generated base URL: %s", baseURL)
	}

	// check if baseURL is already in use
	connMu.Lock()
	for _, t := range connections {
		if t.BaseURL == baseURL {
			connMu.Unlock()
			http.Error(w, "Base URL already in use", http.StatusConflict)
			return
		}
	}
	connMu.Unlock()

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

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

	defer func() {
		connMu.Lock()
		delete(connections, id)
		connMu.Unlock()
		conn.Close()
		log.Printf("Connection closed: %s", id)
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

func httpToWebSocketHandler(w http.ResponseWriter, r *http.Request) {

	tunnel, appID, _ := PathTunnelRouter(r, connections)

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
	log.Println("httpToWebSocketHandler: message sent to tunnel")

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

func StartServer(addr string) {
	r := chi.NewRouter()
	r.Get("/___gTl___/ws", wsHandler)
	r.NotFound(httpToWebSocketHandler)

	log.Println("Server listening on", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
