package main

import (
	"fmt"
	"net/url"

	"github.com/B-AJ-Amar/gTunnel/internal/client"
)

func main() {
	fmt.Println("Hello, gTunnel client!")
	client.StartClient(url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"})
}