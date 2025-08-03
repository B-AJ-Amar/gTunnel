package protocol

import (
	"encoding/json"
)

type MessageType int

const (
	MessageTypeHTTPRequest  MessageType = 1
	MessageTypeHTTPResponse MessageType = 2

	MessageTypeAuthRequest  MessageType = 3
	MessageTypeAuthResponse MessageType = 4

	MessaageTypeConfigRequest MessageType = 5
	MessageTypeConfigResponse MessageType = 6
	MessageTypeError          MessageType = 7
)

type SocketMessage struct {
	// ID      string            `json:"ID"` // ? to add : connection id
	Type    MessageType     `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type HTTPRequestMessage struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    []byte            `json:"body"`
}

type HTTPResponseMessage struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       []byte            `json:"body"`
}

type AuthRequestMessage struct {
	AccessToken string `json:"access_token"`
	BaseURL     string `json:"base_url"`
}

type AuthResponseMessage struct {
	ID      *string `json:"id,omitempty"`
	Success bool    `json:"success,omitempty"`
	Message string  `json:"error,omitempty"` // Optional error message if success is false
}

func NewHTTPRequestMessage(id, method, url string, headers map[string]string, body []byte) (*SocketMessage, error) {
	httpReq := HTTPRequestMessage{
		Method:  method,
		URL:     url,
		Headers: headers,
		Body:    body,
	}

	return NewSocketMessage(id, MessageTypeHTTPRequest, httpReq)
}

func NewHTTPResponseMessage(id string, statusCode int, headers map[string]string, body []byte) (*SocketMessage, error) {
	httpResp := HTTPResponseMessage{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       body,
	}

	return NewSocketMessage(id, MessageTypeHTTPResponse, httpResp)
}

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

func SerializeMessage(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func DeserializeMessage(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
