# gTunnel

**gTunnel** is a lightweight Go-based tunneling system similar to [ngrok](https://ngrok.com/) or [VS Code dev tunnels], allowing you to expose a local port to the

- 🌐 **Tunnel Server**: Listens publicly and relays traffic
- 🧑‍💻 **Tunnel Client**: Forwards local traffic to the server
<!-- - 🔌 **Custom Protocol**: Efficient socket-based communication -->
- ⚡ Written in Go

---

## 📦 Project Structure

```bash
gTunnel/
├── cmd/
│ ├── client/   # Client cli
│ └── server/   # Server cli
├── internal/   # Client-side tunnel logic
│ ├── client/   # Client-side tunnel logic
│ ├── server/   # Public server logic
│ └── protocol/ # Shared message format & protocol
├── go.mod
└── README.md
```

## 🚀 Installation & Build

### Prerequisites

- **Go 1.19+** installed on your system
- **Git** for cloning the repository

### Quick Installation

#### Option 1: Build from Source

```bash
# Clone the repository
git clone https://github.com/B-AJ-Amar/gTunnel.git
cd gTunnel

# Build both client and server
make build

# The binaries will be available in the build/ directory
./build/gtc --help     # Client
./build/gts --help     # Server
```

#### Option 2: Install to GOPATH

```bash
# Clone and install directly to your GOPATH/bin
git clone https://github.com/B-AJ-Amar/gTunnel.git
cd gTunnel
make install

# Now you can use the commands globally
gtc --help
gts --help
```

#### Option 3: Development Build

```bash
# For development (builds without version info)
make build-dev
```

### Build Options

#### Available Make Targets

| Command         | Description                                 |
| --------------- | ------------------------------------------- |
| `make build`    | Build both client and server with version  |
| `make build-client` | Build only the client                  |
| `make build-server` | Build only the server                  |
| `make build-dev`    | Build without version info (development) |
| `make install`      | Install binaries to GOPATH/bin         |
| `make test`         | Run all tests                           |
| `make clean`        | Clean build directory                   |
| `make help`         | Show all available targets              |

#### Custom Version Building

```bash
# Build with custom version
make VERSION=1.0.0 build

# Build with custom build info
make VERSION=1.0.0 GIT_COMMIT=abc123 BUILD_DATE=2024-01-01 build
```

#### Manual Build (without Make)

```bash
# Build client
go build -o gtc ./cmd/client

# Build server  
go build -o gts ./cmd/server

# Build with version info manually
go build -ldflags "-X github.com/B-AJ-Amar/gTunnel/internal/version.Version=1.0.0" -o gtc ./cmd/client
```

### Verify Installation

```bash
# Check version
gtc version
gts --version

# Get detailed version info
gtc version --output json
```

## Run/Deploy the Public Tunnel Server

 comming soon

## Run the Tunnel Client (on your local machine)

comming soon

## Features

- [x] TCP tunnel
- []Authentication (API token)
  - in the first version i will use one token for all clients
- [x] CLI
- []subdomain/path routing (multiple tunnels per client)
- []Multiplexing (multiple streams per client)
- []TLS encryption (optional)
- []WebSocket support
- []Admin dashboard (via Chi)
- []double client server
  - gtunnel client to gtunnel server
  - client management server (dashboard , settings ..etc)
- []Queue for request in case of disconnected clients (for webhooks)

## CLI

TODO : add viper for config

### gtunnel-server (gts)

| Command      | Description                                    |
| ------------ | ---------------------------------------------- |
| `start`      | Starts the tunnel server                       |
| `status`     | Shows server health & active connections       |
| `users`      | Manage connected clients / auth users          |
| `config`     | Manage server configuration                    |
| `logs`       | Show recent connection logs or tunnel activity |
| `version`    | Print server version info                      |
| `completion` | Generate shell autocompletion script           |

### gtunnel-client (gtc)

| Command      | Description                                 |
| ------------ | ------------------------------------------- |
| `connect`    | Connects to the tunnel server               |
| `status`     | Shows current connection status             |
| `disconnect` | Gracefully closes the tunnel connection     |
| `config`     | Manage client config (`set`, `get`, `show`) |
| `version`    | Print client version info                   |
| `completion` | Generate shell autocompletion script        |


## Contributing

All contributions are welcome! Whether it's bug fixes, new features, documentation improvements, or ideas — feel free to open issues or submit pull requests. Let’s build gTunnel together!