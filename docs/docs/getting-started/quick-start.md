---
sidebar_position: 1
---
# Quick Start

Get up and running with gTunnel in minutes! This guide will walk you through installing and setting up a basic HTTP tunnel.

## Installation

Run our automated installation script:

```bash
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash
```

This installs the client component only (default). For specific components:

```bash
# Client only (default)
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash

# Server only  
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s server

# Both client and server
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s both
```

### Manual Installation

Download the latest release from [GitHub Releases](https://github.com/B-AJ-Amar/gTunnel/releases):

```bash
# Download and extract (replace VERSION with latest)
wget https://github.com/B-AJ-Amar/gTunnel/releases/download/v0.0.0/gtunnel-client_linux_amd64.tar.gz
tar -xzf gtunnel-client_linux_amd64.tar.gz
sudo mv gtc /usr/local/bin/
```

## Basic Usage

### 1. Start the Server

First, start a gTunnel server:

```bash
gts start --bind-address 0.0.0.0:7205
```

The server will start and display:

```text
🚀 gTunnel Server started on 0.0.0.0:7205
```

### 2. Start Your Local Service

In another terminal, start the service you want to expose. For example, a simple HTTP server:

```bash
# Python
python -m http.server 3000

# Node.js  
npx serve -p 3000

# Go
cd your-app && go run main.go
```

### 3. Connect the Client

Connect your local service to the gTunnel server:

```bash
gtc connect -u localhost:7205 3000
```

You'll see output like:

```text
✅ Connected to gTunnel server
🌐 Your service is now accessible at: http://localhost:7205/tunnel/abc123
```

### 4. Access Your Service

Your local service is now accessible through the tunnel! Open the provided URL in your browser or share it with others.

## Docker Usage

### Using Docker Compose

Create a `docker-compose.yml`:

```yaml
version: '3.8'
services:
  gtunnel-server:
    image: ghcr.io/b-aj-amar/gtunnel-server:latest
    ports:
      - "7205:7205"
    command: ["start", "--bind-address", "0.0.0.0:7205"]

  gtunnel-client:
    image: ghcr.io/b-aj-amar/gtunnel-client:latest
    depends_on:
      - gtunnel-server
    environment:
      - GTUNNEL_SERVER_URL=http://gtunnel-server:7205
      - GTUNNEL_TARGET_HOST=your-app
      - GTUNNEL_TARGET_PORT=3000
```

Run with:

```bash
docker-compose up
```

### Individual Docker Containers

```bash
# Start server
docker run -p 7205:7205 ghcr.io/b-aj-amar/gtunnel-server:latest start --bind-address 0.0.0.0:7205

# Start client
docker run ghcr.io/b-aj-amar/gtunnel-client:latest connect -u host.docker.internal:7205 3000
```

## Common Use Cases

### Expose Local Development Server

Perfect for sharing your work-in-progress with team members:

```bash
# Start your dev server
npm run dev  # Usually runs on localhost:3000

# Tunnel it
gtc connect -u your-server:7205 3000
```

### Testing Webhooks

Test webhooks locally during development:

```bash
# Your webhook endpoint
gtc connect -u tunnel.yourcompany.com:7205 4000
```

### Remote Access to Services

Access services running on remote machines:

```bash
# On remote machine
gtc connect -u tunnel-server:7205 8080
```

## Next Steps

- 📖 Read the [full documentation](../introduction.md) for advanced configuration
- 🚀 Deploy your own [production server](../deployment)
- 🎯 Check out [cli references](../cli-reference.md)

## Need Help?

- 📋 Check the [FAQ](../faq.md)
- 🐛 [Report issues on GitHub](https://github.com/B-AJ-Amar/gTunnel/issues)
- 💬 [Join our discussions](https://github.com/B-AJ-Amar/gTunnel/discussions)

---

**Ready to dive deeper?** Explore our comprehensive [documentation](../introduction.md) for advanced features and configuration options.
