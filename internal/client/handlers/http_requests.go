package handlers

import (
	"bytes"
	"encoding/base64"
	"io"
	"log"
	"net/http"

	"github.com/B-AJ-Amar/gTunnel/internal/client/models"
	"github.com/B-AJ-Amar/gTunnel/internal/protocol"
	"github.com/gorilla/websocket"
)

func ClientHTTPRequestHandler(socketMessage protocol.SocketMessage, tunnel *models.ClientTunnelConn) error {
	var httpRequest protocol.HTTPRequestMessage
	err := protocol.DeserializeMessage(socketMessage.Payload, &httpRequest)
	if err != nil {
		log.Printf("Error deserializing HTTP request: %v", err)
		return err
	}
	log.Printf("HTTP Request: %s %s", httpRequest.Method, httpRequest.URL)

	// Decode base64 body
	bodyBytes, err := base64.StdEncoding.DecodeString(string(httpRequest.Body))
	if err != nil {
		return err
	}

    // request
    req, err := http.NewRequest(httpRequest.Method, "http://localhost:"+ tunnel.Port  + httpRequest.URL, bytes.NewReader(bodyBytes))
    if err != nil {
        return err
    }

    // Set headers
    for key, value := range httpRequest.Headers {
        req.Header.Set(key, value)
    }

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()


	// Construct HTTPResponseMessage

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    respHeaders := make(map[string]string)
    for k, v := range resp.Header {
        if len(v) > 0 {
            respHeaders[k] = v[0]
        }
    }

    response := protocol.HTTPResponseMessage{
        StatusCode: resp.StatusCode,
        Headers:    respHeaders,
        Body:       respBody, 
    }

    responsePayload, err := protocol.SerializeMessage(response)
    if err != nil {
        return err
    }

    responseMsg := protocol.SocketMessage{
        // ID:      socketMessage.ID, // reply to the same ID
        Type:    protocol.MessageTypeHTTPResponse,
        Payload: responsePayload,
    }	


	// Serialize the response message
	responseBytes, err := protocol.SerializeMessage(responseMsg)

	if err != nil {
		return err
	}

	tunnel.Conn.WriteMessage(websocket.TextMessage, responseBytes)

    return nil
}
