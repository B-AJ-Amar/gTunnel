# gTunnel

**gTunnel** is a fast, lightweight tunneling solution written in Go that allows you to expose local services to the internet through a secure tunnel. Similar to [ngrok](https://ngrok.com/) or [VSCode Port-Forwarding](https://code.visualstudio.com/docs/debugtest/port-forwarding), GTunnel provides a simple way to share your local development server with the world.

![License](https://img.shields.io/github/license/B-AJ-Amar/gTunnel)
![Go Version](https://img.shields.io/github/go-mod/go-version/B-AJ-Amar/gTunnel)
![Release](https://img.shields.io/github/v/release/B-AJ-Amar/gTunnel)
<!--![Downloads](https://img.shields.io/github/downloads/B-AJ-Amar/gTunnel/total) -->

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

### Quick Installation

```bash
# Install client (recommended for most users)
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash

# Install server for self-hosting
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s server
```

### Alternative Methods

- **[Download Binaries](https://github.com/B-AJ-Amar/gTunnel/releases)** - Direct downloads for all platforms
- **[Docker Images](https://ghcr.io/b-aj-amar/gtunnel-client)** - Container deployment
- **Package Managers** - DEB, RPM, APK packages

📖 **For detailed installation instructions, see our [Installation Guide](https://b-aj-amar.github.io/gTunnel/docs/getting-started/installation)**

## 📖 Usage

### Basic Client Usage

```bash
# Connect to public server (replace with your server URL)
gtc connect -u https://gtunnel-server-1i1b.onrender.com -e /my-app 3000

# Your app is now available at: https://gtunnel-server-1i1b.onrender.com/my-app
```

### Server Usage

```bash
# Start your own server
gts start --bind-address 0.0.0.0:7205
```

📖 **For complete usage examples and configuration, see our [Documentation](https://b-aj-amar.github.io/gTunnel/docs/introduction)**

## 🌐 Deployment

### Quick Deploy (Recommended)

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy?repo=https://github.com/B-AJ-Amar/gTunnel)

Deploy your own gTunnel server with one click - includes free HTTPS, automatic SSL, and health monitoring.

📖 **For self-hosted options and advanced deployment, see our [Deployment Guide](https://b-aj-amar.github.io/gTunnel/docs/deployment)**

## 📚 Documentation

Our comprehensive documentation covers everything you need:

- 🚀 **[Quick Start](https://b-aj-amar.github.io/gTunnel/docs/getting-started/quick-start)** - Get running in 5 minutes
- 📦 **[Installation](https://b-aj-amar.github.io/gTunnel/docs/getting-started/installation)** - All installation methods
- ⚙️ **[Configuration](https://b-aj-amar.github.io/gTunnel/docs/configuration)** - Detailed configuration options
- 🌐 **[Deployment](https://b-aj-amar.github.io/gTunnel/docs/deployment)** - Production deployment guides
- 🛠️ **[CLI Reference](https://b-aj-amar.github.io/gTunnel/docs/cli-reference)** - Complete command reference
- ❓ **[FAQ](https://b-aj-amar.github.io/gTunnel/docs/faq)** - Frequently asked questions

**Visit [Documentation Site](https://b-aj-amar.github.io/gTunnel/) for the complete guide.**

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
