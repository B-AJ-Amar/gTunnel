package server

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/B-AJ-Amar/gTunnel/internal/protocol"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type TunnelConn struct {
	ID   string
	Conn *websocket.Conn
}

var (
	connections = make(map[string]*TunnelConn)
	connMu      sync.Mutex
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	id := uuid.New().String()
	tunnel := &TunnelConn{ID: id, Conn: conn}

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

	// Echo loop
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("[%s] Received: %s", id, message)
		if err := conn.WriteMessage(mt, message); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func httpToWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// For now, just use the first available connection
	connMu.Lock()
	var tunnel *TunnelConn
	for _, t := range connections {
		tunnel = t
		break
	}
	connMu.Unlock()

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

	reqMsg := protocol.HTTPRequestMessage{
		Method:  r.Method,
		URL:     r.URL.String(),
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

	_, responseData, err := tunnel.Conn.ReadMessage()
	if err != nil {
		http.Error(w, "Tunnel read failed", http.StatusBadGateway)
		return
	}

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
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("testHandler called")
	connMu.Lock()
	if len(connections) == 0 {
		log.Println("No active connections")
	} else {
		for id, tunnel := range connections {
			log.Printf("Sending test message to: %s", id)
			tunnel.Conn.WriteMessage(websocket.TextMessage, []byte("Hello from testHandler"))
		}
	}
	connMu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"TEST"}`))
}

func StartServer(addr string) {
	r := chi.NewRouter()
	r.Get("/ws", wsHandler)
	// r.NotFound(testHandler)
	r.NotFound(httpToWebSocketHandler)

	log.Println("Server listening on", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
