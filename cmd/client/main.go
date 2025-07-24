package main

import (
	"fmt"
	"net/url"

	"github.com/B-AJ-Amar/gTunnel/internal/client"
)

func main() {
	fmt.Println("Hello, gTunnel client!")
	u := url.URL{
		Scheme:   "ws",
		Host:     "localhost:8080",
		Path:     "/___gTl___/ws",
	}
	q := u.Query()
	q.Set("baseUrl", "/base_url")
	u.RawQuery = q.Encode()
	client.StartClient(u)
}