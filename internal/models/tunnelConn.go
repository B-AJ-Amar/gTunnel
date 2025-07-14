package models

import (
	"github.com/gorilla/websocket"
)

type ServerTunnelConn struct {
	ID         string
	Conn       *websocket.Conn
	ResponseCh chan []byte
}

type ClientTunnelConn struct {
	ID         string
	Conn       *websocket.Conn
	ResponseCh chan []byte
	// TODO : add listen to  port and base url ( in case of routes based forwarding)
	// Port       int
	// BaseURL    string
}

