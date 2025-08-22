# gTunnel

**gTunnel** is a fast, lightweight tunneling solution written in Go that allows you to expose local services to the internet through a secure tunnel. Similar to [ngrok](https://ngrok.com/) or [VSCode Port-Forwarding](https://code.visualstudio.com/docs/debugtest/port-forwarding), GTunnel provides a simple way to share your local development server with the world.

![License](https://img.shields.io/github/license/B-AJ-Amar/gTunnel)
![Go Version](https://img.shields.io/github/go-mod/go-version/B-AJ-Amar/gTunnel)
![Release](https://img.shields.io/github/v/release/B-AJ-Amar/gTunnel)
![Downloads](https://img.shields.io/github/downloads/B-AJ-Amar/gTunnel/total)

## ✨ Features

- 🚀 **Fast & Lightweight** - Written in Go for optimal performance
- 🔒 **Secure** - Token-based authentication and HTTPS/WSS support
- 🌐 **Cross-Platform** - Works on Windows, macOS, and Linux
- 🐳 **Docker Ready** - Pre-built Docker images for easy deployment
- 📦 **Easy Installation** - Multiple installation methods (binaries, packages, Docker)
- ⚡ **Zero Configuration** - Works out of the box with sensible defaults
- 🔧 **Flexible** - CLI tools for both client and server management

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Local App     │    │  gTunnel Client │    │  gTunnel Server │
│ localhost:3000  │◄──►│      (gtc)      │◄──►│      (gts)      │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                                       ▲
                                                       │
                                               ┌───────▼───────┐
                                               │  Public URL   │
                                               │ tunnel.me:443 │
                                               └───────────────┘
```

## 📋 Table of Contents

- [Quick Start](#-quick-start)
- [Installation](#-installation)
- [Usage](#-usage)
- [Deployment](#-deployment)
- [Configuration](#-configuration)
- [CLI Reference](#-cli-reference)
- [Development](#-development)
- [Contributing](#-contributing)

## 🚀 Quick Start

### 1. Install the Client

**Download from releases:**
```bash
# Linux/macOS
curl -L https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel-client_linux_amd64.tar.gz | tar -xz
sudo mv gtc /usr/local/bin/

# Or use our installation script (supports client, server, or both)
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash

# Install server only
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s server

# Install both client and server
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s both
```

### 2. Connect to Public Server

```bash
# Expose your local port 3000
# https://gtunnel-server-1i1b.onrender.com is the default server for the version v0.0.0
gtc connect -u https://gtunnel-server-1i1b.onrender.com -e /my-app 3000

# Your app is now available at: https://gtunnel-server-1i1b.onrender.com/my-app
```

### 3. Deploy Your Own Server (Optional)

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy?repo=https://github.com/B-AJ-Amar/gTunnel)

## 📦 Installation

### Client Installation

#### Option 1: Download Binary (Recommended)

**Linux/macOS:**
```bash
# Download and install
curl -L https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel-client_linux_amd64.tar.gz | tar -xz
sudo mv gtc /usr/local/bin/

# Verify installation
gtc version
```

**Windows:**
```powershell
# Download from GitHub releases
Invoke-WebRequest -Uri "https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel-client_windows_amd64.zip" -OutFile "gtunnel-client.zip"
Expand-Archive gtunnel-client.zip
```

#### Option 2: Package Managers

**Linux (DEB/RPM/APK):**
```bash
# Debian/Ubuntu
wget https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel_linux_amd64.deb
sudo dpkg -i gtunnel_linux_amd64.deb

# Red Hat/CentOS/Fedora
wget https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel-1.0.0-1.x86_64.rpm
sudo rpm -i gtunnel-1.0.0-1.x86_64.rpm

# Alpine Linux
wget https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel-1.0.0-r0.apk
sudo apk add --allow-untrusted gtunnel-1.0.0-r0.apk
```

#### Option 3: Installation Script (Flexible)

The installation script supports installing client, server, or both components with a single command:

```bash
# Install client only (default)
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash

# Install server only
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s server

# Install both client and server
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s both

# Install specific version
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s -- --version v1.0.0 client

# Custom installation directory
INSTALL_DIR=~/.local/bin bash <(curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh) both

# Get help and see all options
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s -- --help
```

**Features:**

- ✅ **Cross-platform** - Works on Linux, macOS, and Windows
- ✅ **Architecture detection** - Automatically detects amd64, arm64, 386
- ✅ **Version selection** - Install latest or specific version
- ✅ **Component selection** - Choose client, server, or both
- ✅ **Custom install directory** - Install to any directory
- ✅ **Verification** - Automatically verifies installation

#### Option 4: Build from Source

```bash
# Prerequisites: Go 1.19+
git clone https://github.com/B-AJ-Amar/gTunnel.git
cd gTunnel

# Build both client and server
make build

# Install to system
sudo make install
```

#### Option 5: Docker

For containerized environments and microservices:

```bash
# Run client in Docker
docker run -d --name gtunnel-client \
  -e GTUNNEL_SERVER_URL="https://your-server.com" \
  -e GTUNNEL_ACCESS_TOKEN="your-token" \
  -e GTUNNEL_TARGET_HOST="app" \
  -e GTUNNEL_TARGET_PORT="3000" \
  --network your-network \
  ghcr.io/b-aj-amar/gtunnel-client:latest

# With Docker Compose
cat > docker-compose.yml << EOF
version: '3.8'
services:
  app:
    image: your-app:latest
    ports:
      - "3000:3000"
    networks:
      - app-network

  gtunnel-client:
    image: ghcr.io/b-aj-amar/gtunnel-client:latest
    environment:
      - GTUNNEL_SERVER_URL=https://your-server.com
      - GTUNNEL_ACCESS_TOKEN=your-token
      - GTUNNEL_TARGET_HOST=app
      - GTUNNEL_TARGET_PORT=3000
      - GTUNNEL_BASE_ENDPOINT=/my-app
      - GTUNNEL_DEBUG=false
    depends_on:
      - app
    networks:
      - app-network
    restart: unless-stopped

networks:
  app-network:
    driver: bridge
EOF

docker-compose up -d
```

**Environment Variables:**

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `GTUNNEL_SERVER_URL` | gTunnel server URL | `""` | ✅ |
| `GTUNNEL_ACCESS_TOKEN` | Authentication token | `""` | ❌ |
| `GTUNNEL_TARGET_HOST` | Target host to tunnel | `localhost` | ❌ |
| `GTUNNEL_TARGET_PORT` | Target port to tunnel | `3000` | ❌ |
| `GTUNNEL_BASE_ENDPOINT` | Base endpoint path | `""` | ❌ |
| `GTUNNEL_DEBUG` | Enable debug logging | `false` | ❌ |

### Server Installation

#### Option 1: Installation Script (Recommended)
```bash
# Install server only
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s server

# Install both client and server
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s both

# Install specific version
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s -- --version v1.0.0 server
```

#### Option 2: Docker (Cloud Deployment)
```bash
# Run server with Docker
docker run -d --name gtunnel-server 
  -p 7205:7205 
  -e GTUNNEL_ACCESS_TOKEN="your-secret-token" 
  ghcr.io/b-aj-amar/gtunnel-server:latest

# Or with Docker Compose
cat > docker-compose.yml << EOF
version: '3.8'
services:
  gtunnel-server:
    image: ghcr.io/b-aj-amar/gtunnel-server:latest
    ports:
      - "7205:7205"
    environment:
      - GTUNNEL_ACCESS_TOKEN=your-secret-token
      - GTUNNEL_USE_ENV=true
    restart: unless-stopped
EOF

docker-compose up -d
```

## 📖 Usage

### Basic Client Usage

**Expose a local service:**
```bash
# Expose localhost:3000
gtc connect 3000

# Expose specific host:port
gtc connect myapp.local:8080

# Connect to custom server
gtc connect -u https://your-server.com 3000

# Connect with custom base path
gtc connect -e /my-app 3000
```

**Configuration:**
```bash
# Set default server URL
gtc config --set-url https://your-server.com

# Set access token
gtc config --set-token your-token

# View current config
gtc config show
```

### Docker Client Usage

**Run client in a container:**
```bash
# Basic usage
docker run --rm \
  -e GTUNNEL_SERVER_URL="https://your-server.com" \
  -e GTUNNEL_TARGET_HOST="host.docker.internal" \
  -e GTUNNEL_TARGET_PORT="3000" \
  ghcr.io/b-aj-amar/gtunnel-client:latest

# With authentication
docker run --rm \
  -e GTUNNEL_SERVER_URL="https://your-server.com" \
  -e GTUNNEL_ACCESS_TOKEN="your-token" \
  -e GTUNNEL_TARGET_HOST="host.docker.internal" \
  -e GTUNNEL_TARGET_PORT="8080" \
  -e GTUNNEL_BASE_ENDPOINT="/api" \
  ghcr.io/b-aj-amar/gtunnel-client:latest

# Debug mode
docker run --rm \
  -e GTUNNEL_DEBUG="true" \
  -e GTUNNEL_SERVER_URL="https://your-server.com" \
  -e GTUNNEL_TARGET_HOST="host.docker.internal" \
  -e GTUNNEL_TARGET_PORT="3000" \
  ghcr.io/b-aj-amar/gtunnel-client:latest
```

### Basic Server Usage

**Start the server:**
```bash
# Start with default settings
gts start

# Start on custom address
gts start --bind-address 0.0.0.0:8080

# Start with debug logging
gts start --debug
```

**Configuration:**
```bash
# Set access token
gts config set access-token your-secret-token

# View server status
gts status

# View version
gts version
```

## 🌐 Deployment

### Cloud Platforms

#### Render (Free Tier)
[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy?repo=https://github.com/B-AJ-Amar/gTunnel)

**What you get:**
- ✅ Free HTTPS URL (e.g., `https://your-app.onrender.com`)
- ✅ Automatic SSL certificates
- ✅ Health checks and auto-restart
- ✅ Auto-deploy on git push
- ✅ Environment variable management

<!-- #### Railway
```bash
# Install Railway CLI and deploy
railway login
railway init
railway up
```

#### Heroku
```bash
# Deploy to Heroku
heroku create your-gtunnel-server
heroku container:push web
heroku container:release web
``` -->

### Self-Hosted

#### VPS/Dedicated Server

**Using systemd (Linux):**
```bash
# Download and install
wget https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel-server_linux_amd64.tar.gz
tar -xzf gtunnel-server_linux_amd64.tar.gz
sudo mv gts /usr/local/bin/

# Create service user
sudo useradd --system --shell /bin/false gtunnel

# Create config
sudo mkdir -p /etc/gtunnel
echo "GTUNNEL_ACCESS_TOKEN=your-secret-token" | sudo tee /etc/gtunnel/server.env

# Create systemd service
sudo tee /etc/systemd/system/gtunnel.service > /dev/null << EOF
[Unit]
Description=gTunnel Server
After=network.target

[Service]
Type=simple
User=gtunnel
ExecStart=/usr/local/bin/gts start --bind-address 0.0.0.0:7205
EnvironmentFile=/etc/gtunnel/server.env
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

# Start and enable service
sudo systemctl daemon-reload
sudo systemctl enable gtunnel
sudo systemctl start gtunnel
```

#### Docker Compose

**Production setup:**
```yaml
version: '3.8'
services:
  gtunnel-server:
    image: ghcr.io/b-aj-amar/gtunnel-server:latest
    ports:
      - "443:7205"
    environment:
      - GTUNNEL_ACCESS_TOKEN=${GTUNNEL_ACCESS_TOKEN}
      - GTUNNEL_USE_ENV=true
      - GTUNNEL_LOG_LEVEL=info
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:7205/___gTl___/health"]
      interval: 30s
      timeout: 10s
      retries: 3
    networks:
      - gtunnel

  # Optional: Reverse proxy
  nginx:
    image: nginx:alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf
      - ./ssl:/etc/nginx/ssl
    depends_on:
      - gtunnel-server
    networks:
      - gtunnel

networks:
  gtunnel:
    driver: bridge
```

## ⚙️ Configuration

### Environment Variables

**Server:**
| Variable | Description | Default |
|----------|-------------|---------|
| `GTUNNEL_ACCESS_TOKEN` | Authentication token | `""` |
| `GTUNNEL_USE_ENV` | Use environment variables | `false` |
| `GTUNNEL_LOG_LEVEL` | Log level (debug/info/warn/error) | `info` |
| `PORT` | Server port (for cloud platforms) | `7205` |

**Client:**
| Variable | Description | Default |
|----------|-------------|---------|
| `GTUNNEL_SERVER_URL` | Default server URL | `""` |
| `GTUNNEL_ACCESS_TOKEN` | Authentication token | `""` |

### Configuration Files

**Client config** (`~/.config/gtunnel/config.yaml`):
```yaml
server_url: "https://your-server.com"
access_token: "your-token"
```

**Server config** (`~/.config/gtunnel/config.yaml`):
```yaml
access_token: "your-secret-token"
```

## 🛠️ CLI Reference

### gtc (Client)

```bash
# Connection
gtc connect [flags] <port|host:port>
  -u, --server-url string     Server URL (e.g., example.com:443)
  -e, --base-endpoint string  Base endpoint path
  -d, --debug                 Enable debug logging

# Configuration
gtc config <command>
  set <key> <value>          Set configuration value
  get <key>                  Get configuration value
  show                       Show all configuration

# Utility
gtc version                  Show version information
gtc completion <shell>       Generate completion script
```

### gts (Server)

```bash
# Server management
gts start [flags]
  --bind-address string      Address to bind (default "0.0.0.0:7205")
  -d, --debug               Enable debug logging

# Status and monitoring
gts status                   Show server status
gts version                  Show version information

# Configuration
gts config <command>
  set <key> <value>          Set configuration value
  show                       Show configuration

# Utility
gts completion <shell>       Generate completion script
```

## 🧪 Development

### Development Setup

```bash
# Clone repository
git clone https://github.com/B-AJ-Amar/gTunnel.git
cd gTunnel

# Install dependencies
go mod download

# Run tests
make test

# Build for development
make build-dev

# Build with version info
make VERSION=dev build
```

### Available Make Targets

| Command | Description |
|---------|-------------|
| `make build` | Build both client and server with version info |
| `make build-client` | Build only the client |
| `make build-server` | Build only the server |
| `make build-dev` | Build without version info |
| `make install` | Install binaries to GOPATH/bin |
| `make test` | Run all tests |
| `make clean` | Clean build directory |
| `make release` | Create a release with GoReleaser |
| `make help` | Show all available targets |

### Project Structure

```
gTunnel/
├── cmd/
│   ├── client/              # Client CLI application
│   └── server/              # Server CLI application
├── internal/
│   ├── client/              # Client-side logic
│   ├── server/              # Server-side logic
│   ├── protocol/            # Shared protocol definitions
│   ├── logger/              # Logging utilities
│   └── version/             # Version information
├── scripts/                 # Build and deployment scripts
├── docs/                    # Documentation
├── Dockerfile.server        # Server Docker image
├── render.yaml              # Render deployment config
├── .goreleaser.yml          # Release configuration
└── Makefile                 # Build automation
```

## 🤝 Contributing

We welcome contributions! Here's how you can help:

### Getting Started

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Make your changes**
4. **Run tests**: `make test`
5. **Commit your changes**: `git commit -m 'Add amazing feature'`
6. **Push to the branch**: `git push origin feature/amazing-feature`
7. **Open a Pull Request**

### Development Guidelines

- **Go style**: Follow Go conventions and use `gofmt`
- **Tests**: Add tests for new features
- **Documentation**: Update README and code comments
- **Commits**: Use clear, descriptive commit messages

### Roadmap

- [x] **Basic tunneling** - HTTP/HTTPS tunneling through WebSocket
- [x] **Authentication** - Token-based authentication (basic)
- [x] **CLI tools** - Complete command-line interface
- [x] **Docker support** - Pre-built Docker images
- [x] **Multi-platform** - Windows, macOS, and Linux support
- [ ] **WebSocket support** - Native WebSocket tunneling for real-time applications
- [ ] **Multiple tunnels** - Multiple tunnels per client
- [ ] **Multicasting** - Broadcast same message to multiple clients (for webhooks)
- [ ] **Advanced client management** - Profiles, advanced auth, permissions ...
- [ ] **VS Code extension** - Integrated tunneling directly in VS Code
- [ ] **Documentation site** - Comprehensive docs with Docusaurus
- [ ] **Admin dashboard** - Web interface for server management
- [ ] **Request queuing** - Queue requests for offline clients
- [ ] **.gtunnel config file** - Project-specific config file for easier connection management
- [ ] **Subdomain routing** - Custom subdomains for tunnels

## 📄 License

This project is licensed under the Apache-2.0 license - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with [Go](https://golang.org/), [Chi](https://github.com/go-chi/chi), [Gorilla WebSocket](https://github.com/gorilla/websocket), and [Cobra](https://github.com/spf13/cobra)
- Thanks to all contributors and users!

---

**Made with ❤️ by [@B-AJ-Amar](https://github.com/B-AJ-Amar)**
