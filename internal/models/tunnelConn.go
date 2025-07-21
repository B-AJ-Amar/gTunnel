package models

import (
	"github.com/gorilla/websocket"
)

type ServerTunnelConn struct {
	ID         string
	Conn       *websocket.Conn
	ResponseCh chan []byte
	BaseURL    string
}


type ClientTunnelConn struct {
	ID         string
	Conn       *websocket.Conn
	ResponseCh chan []byte
	Port       string
}

