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
│ ├── client/ # Client binary
│ └── server/ # Server binary
├── client/ # Client-side tunnel logic
├── server/ # Public server logic
├── protocol/ # Shared message format & protocol
├── go.mod
└── README.md
```

##  Run/Deploy the Public Tunnel Server

 comming soon

## Run the Tunnel Client (on your local machine)

comming soon

## Features

- [x] TCP tunnel

- []HTTP subdomain routing

- []Authentication (API token)

- []TLS encryption (optional)

- []WebSocket support

- []Multiplexing (multiple streams per client)

- []Admin dashboard (via Chi)

## Contributing

All contributions are welcome! Whether it's bug fixes, new features, documentation improvements, or ideas — feel free to open issues or submit pull requests. Let’s build gTunnel together!