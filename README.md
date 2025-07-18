# gTunnel

**gTunnel** is a lightweight Go-based tunneling system similar to [ngrok](https://ngrok.com/) or [VS Code dev tunnels], allowing you to expose a local port to the public internet via a secure tunnel through a public relay server.

- 🌐 **Tunnel Server**: Listens publicly and relays traffic
- 🧑‍💻 **Tunnel Client**: Forwards local traffic to the server
<!-- - 🔌 **Custom Protocol**: Efficient socket-based communication -->
- ⚡ Written in Go

---

## 📦 Project Structure

```bash
gTunnel/
├── cmd/
│ ├── client/   # Client binary
│ └── server/   # Server binary
├── internal/   # Client-side tunnel logic
│ ├── client/   # Client-side tunnel logic
│ ├── server/   # Public server logic
│ ├── protocol/ # Shared message format & protocol
│ ├── models/   # Data models
│ └── handlers/ # Request handlers
├── go.mod
└── README.md
```

##  Run/Deploy the Public Tunnel Server

 comming soon

## Run the Tunnel Client (on your local machine)

comming soon

## Features

- [x] TCP tunnel
- []Authentication (API token)
- []subdomain/path routing (multiple tunnels per client)
- []Multiplexing (multiple streams per client)
- []TLS encryption (optional)
- []WebSocket support
- []Admin dashboard (via Chi)
- []double client server 
  - gtunnel client to gtunnel server
  - client management server (dashboard , settings ..etc)
- []Queue for request in case of disconnected clients


## Contributing

All contributions are welcome! Whether it's bug fixes, new features, documentation improvements, or ideas — feel free to open issues or submit pull requests. Let’s build gTunnel together!