package models

import (
	"github.com/gorilla/websocket"
)

type ServerTunnelConn struct {
	ID         string
	Conn       *websocket.Conn
	ResponseCh chan []byte
	BaseURL    string // ? for the first version , base url should be only one level deep , e.g /app-1 , // later we can make it more complex
}
