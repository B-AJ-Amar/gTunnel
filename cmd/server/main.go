package main

import (
	"fmt"
	"github.com/B-AJ-Amar/gTunnel/server"
)

func main() {
	fmt.Println("Hello, gTunnel server!")
	server.StartServer(":8080")
}