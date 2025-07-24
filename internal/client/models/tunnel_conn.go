package models

import (
	"github.com/gorilla/websocket"
)

type ClientTunnelConn struct {
	ID         string
	Conn       *websocket.Conn
	ResponseCh chan []byte
	Port       string
}

