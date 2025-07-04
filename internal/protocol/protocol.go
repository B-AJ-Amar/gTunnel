package protocol

import (
	"encoding/json"
)

// MessageType defines the type of message sent over the websocket.
type MessageType int

const (
	MessageTypeHTTPRequest  MessageType = 1
	MessageTypeHTTPResponse MessageType = 2
)

// SocketMessage is a generic message that wraps different message types based on Type.
type SocketMessage struct {
	Type    MessageType       `json:"type"`
	Payload json.RawMessage   `json:"payload"`
}

// HTTPRequestMessage represents an HTTP request sent from client to server.
type HTTPRequestMessage struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    []byte            `json:"body"`
}

// HTTPResponseMessage represents an HTTP response sent from server to client.
type HTTPResponseMessage struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       []byte            `json:"body"`
}


func SerializeMessage(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func DeserializelMessage(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}