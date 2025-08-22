---
sidebar_position: 2
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Installation Guide

Get gTunnel up and running on your system with multiple installation options. Choose the method that works best for your environment.

## 🚀 Client Installation

The gTunnel client (`gtc`) connects your local services to a gTunnel server for public access.

### 📥 Option 1: Download Binary (Recommended)

The fastest way to get started - download pre-built binaries for your platform.

<Tabs groupId="operating-systems">
<TabItem value="linux-macos" label="Linux/macOS">

```bash
# Download and install
curl -L https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel-client_linux_amd64.tar.gz | tar -xz
sudo mv gtc /usr/local/bin/

# Verify installation
gtc version
```

</TabItem>
<TabItem value="windows" label="Windows">

```bash
# Download from GitHub releases
Invoke-WebRequest -Uri "https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel-client_windows_amd64.zip" -OutFile "gtunnel-client.zip"
Expand-Archive gtunnel-client.zip

# Add to PATH or move to desired location
Move-Item .\gtunnel-client\gtc.exe C:\Windows\System32\
```

</TabItem>
</Tabs>

### 📦 Option 2: Package Managers

Install using your system's package manager for automatic updates and easy removal.

<Tabs groupId="package-managers">
<TabItem value="debian" label="Debian/Ubuntu">

```bash
# Download and install DEB package
wget https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel_linux_amd64.deb
sudo dpkg -i gtunnel_linux_amd64.deb
```

</TabItem>
<TabItem value="redhat" label="Red Hat/CentOS/Fedora">

```bash
# Download and install RPM package
wget https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel-1.0.0-1.x86_64.rpm
sudo rpm -i gtunnel-1.0.0-1.x86_64.rpm
```

</TabItem>
<TabItem value="alpine" label="Alpine Linux">

```bash
# Download and install APK package
wget https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel-1.0.0-r0.apk
sudo apk add --allow-untrusted gtunnel-1.0.0-r0.apk
```

</TabItem>
</Tabs>

### 🛠️ Option 3: Installation Script (Flexible)

Our smart installation script automatically detects your system and installs the right version.

```bash title="Basic Installation"
# Install client only (default)
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash
```

<details>
<summary><strong>Advanced Installation Options</strong></summary>

```bash
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

</details>

**Script Features:**

- ✅ **Cross-platform** - Works on Linux, macOS, and Windows
- ✅ **Architecture detection** - Automatically detects amd64, arm64, 386
- ✅ **Version selection** - Install latest or specific version
- ✅ **Component selection** - Choose client, server, or both
- ✅ **Custom install directory** - Install to any directory
- ✅ **Verification** - Automatically verifies installation

### 🔨 Option 4: Build from Source

For developers who want the latest features or need to modify the code.

```bash title="Build from Source"
# Prerequisites: Go 1.19+
git clone https://github.com/B-AJ-Amar/gTunnel.git
cd gTunnel

# Build both client and server
make build

# Install to system
sudo make install
```

:::note Prerequisites
Make sure you have Go 1.19 or later installed on your system. You can download it from [golang.org](https://golang.org/dl/).
:::

### 🐳 Option 5: Docker

Perfect for containerized environments and microservices architectures.

<Tabs groupId="docker-methods">
<TabItem value="docker-run" label="Docker Run">

```bash title="Run Client Container"
docker run -d --name gtunnel-client \
  -e GTUNNEL_SERVER_URL="https://your-server.com" \
  -e GTUNNEL_ACCESS_TOKEN="your-token" \
  -e GTUNNEL_TARGET_HOST="app" \
  -e GTUNNEL_TARGET_PORT="3000" \
  --network your-network \
  ghcr.io/b-aj-amar/gtunnel-client:latest
```

</TabItem>
<TabItem value="docker-compose" label="Docker Compose">

```yaml title="docker-compose.yml"
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
```

```bash
# Start the services
docker-compose up -d
```

</TabItem>
</Tabs>

**Docker Environment Variables:**

| Variable | Description | Default | Required |
|----------|-------------|---------|:--------:|
| `GTUNNEL_SERVER_URL` | gTunnel server URL | `""` | ✅ |
| `GTUNNEL_ACCESS_TOKEN` | Authentication token | `""` | ❌ |
| `GTUNNEL_TARGET_HOST` | Target host to tunnel | `localhost` | ❌ |
| `GTUNNEL_TARGET_PORT` | Target port to tunnel | `3000` | ❌ |
| `GTUNNEL_BASE_ENDPOINT` | Base endpoint path | `""` | ❌ |
| `GTUNNEL_DEBUG` | Enable debug logging | `false` | ❌ |

## 🖥️ Server Installation

The gTunnel server (`gts`) receives connections from clients and exposes them to the public internet.

### 🛠️ Option 1: Installation Script (Recommended)

The easiest way to install the server component.

```bash title="Install Server"
# Install server only
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s server

# Install both client and server
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s both

# Install specific version
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s -- --version v1.0.0 server
```

### 🐳 Option 2: Docker Deployment

Ideal for cloud platforms and containerized infrastructure.

<Tabs groupId="server-docker">
<TabItem value="docker-run-server" label="Docker Run">

```bash title="Run Server Container"
docker run -d --name gtunnel-server \
  -p 7205:7205 \
  -e GTUNNEL_ACCESS_TOKEN="your-secret-token" \
  ghcr.io/b-aj-amar/gtunnel-server:latest
```

</TabItem>
<TabItem value="docker-compose-server" label="Docker Compose">

```yaml title="docker-compose.yml"
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
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:7205/___gTl___/health"]
      interval: 30s
      timeout: 10s
      retries: 3
```

```bash
# Start the server
docker-compose up -d
```

</TabItem>
</Tabs>

:::tip Cloud Deployment
For quick cloud deployment, check out our [Deployment Guide](/deployment.md) which includes one-click deployment options.
:::

## ✅ Verify Installation

After installation, verify that gTunnel is working correctly:

```bash
# Check client version
gtc version

# Check server version (if installed)
gts version

# Test client configuration
gtc config show
```

## 🆘 Troubleshooting

### Common Issues

<details>
<summary><strong>Command not found after installation</strong></summary>

**Solution:** Add the installation directory to your PATH:

```bash
# Add to ~/.bashrc or ~/.zshrc
export PATH=$PATH:/usr/local/bin

# Or for custom installation directory
export PATH=$PATH:~/.local/bin
```

</details>

<details>
<summary><strong>Permission denied errors</strong></summary>

**Solution:** Ensure you have proper permissions:

```bash
# Make binary executable
chmod +x /path/to/gtc

# Or install with proper permissions
sudo mv gtc /usr/local/bin/
sudo chmod +x /usr/local/bin/gtc
```

</details>

<details>
<summary><strong>Docker container won't start</strong></summary>

**Solution:** Check the logs and environment variables:

```bash
# Check container logs
docker logs gtunnel-client

# Verify environment variables
docker inspect gtunnel-client
```

</details>

## 📚 Next Steps

Now that you have gTunnel installed:

1. **[Quick Start](/getting-started/quick-start.md)** - Get your first tunnel running
2. **[Configuration](/configuration.md)** - Set up authentication and preferences  
3. **[CLI Reference](/cli-reference.md)** - Learn all available commands
4. **[Deployment](/deployment.md)** - Deploy your own server

---

Need help? Check our **[FAQ](/faq.md)** or visit our **[GitHub Issues](https://github.com/B-AJ-Amar/gTunnel/issues)** for support.
