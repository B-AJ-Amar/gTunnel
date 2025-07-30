package protocol

import (
	"encoding/json"
)

// MessageType defines the type of message sent over the websocket.
type MessageType int

const (
	MessageTypeHTTPRequest  MessageType = 1
	MessageTypeHTTPResponse MessageType = 2

	MessageTypeConnectionRequest  MessageType = 3
	MessageTypeConnectionResponse MessageType = 4

	MessaageTypeConfigRequest MessageType = 5
	MessageTypeConfigResponse MessageType = 6
	MessageTypeError          MessageType = 7
)

// SocketMessage is a generic message that wraps different message types based on Type.
type SocketMessage struct {
	// ID      string            `json:"ID"` // ? to add : connection id
	Type    MessageType     `json:"type"`
	Payload json.RawMessage `json:"payload"`
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

// NewHTTPRequestMessage creates a SocketMessage for an HTTP request
func NewHTTPRequestMessage(id, method, url string, headers map[string]string, body []byte) (*SocketMessage, error) {
	httpReq := HTTPRequestMessage{
		Method:  method,
		URL:     url,
		Headers: headers,
		Body:    body,
	}

	return NewSocketMessage(id, MessageTypeHTTPRequest, httpReq)
}

// NewHTTPResponseMessage creates a SocketMessage for an HTTP response
func NewHTTPResponseMessage(id string, statusCode int, headers map[string]string, body []byte) (*SocketMessage, error) {
	httpResp := HTTPResponseMessage{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       body,
	}

	return NewSocketMessage(id, MessageTypeHTTPResponse, httpResp)
}

// NewSocketMessage creates a new SocketMessage with the given type and payload
func NewSocketMessage(id string, msgType MessageType, payload interface{}) (*SocketMessage, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return &SocketMessage{
		// ID:      id,
		Type:    msgType,
		Payload: payloadBytes,
	}, nil
}

// GetPayload extracts and deserializes the payload into the provided interface
func (sm *SocketMessage) GetPayload(v interface{}) error {
	return json.Unmarshal(sm.Payload, v)
}

// ToBytes serializes the SocketMessage to bytes
func (sm *SocketMessage) ToBytes() ([]byte, error) {
	return json.Marshal(sm)
}

func SerializeMessage(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func DeserializeMessage(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
