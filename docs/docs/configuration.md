---
sidebar_position: 3
---

# Configuration

Learn how to configure gTunnel client and server components for your specific needs.

## Overview

gTunnel supports multiple configuration methods:

- **Configuration files** - YAML files stored in standard locations
- **Environment variables** - For containerized and serverless deployments
- **Command-line flags** - For one-time overrides and testing

## Client Configuration

### Configuration File

**Default location:** `~/.config/gtunnel/config.yaml`

The client configuration file is automatically created when you first use the `gtc config` command.

```yaml
access_token: "your-auth-token"
server_url: "example.com:8080"
```

### Configuration Fields

| Field | Type | Description | Example |
|-------|------|-------------|---------|
| `access_token` | string | Authentication token for server access | `"abc123def456"` |
| `server_url` | string | Server URL in host:port format (without protocol/endpoints) | `"tunnel.example.com:8080"` |

### Managing Client Configuration

```bash
# Show current configuration
gtc config

# Set server URL (protocols and endpoints are automatically cleaned)
gtc config --set-url tunnel.example.com:8080
gtc config --set-url ws://tunnel.example.com:8080/ws  # Same result

# Set access token
gtc config --set-token your-secure-token

# Show detailed configuration
gtc config --show
```

:::tip URL Processing
gTunnel automatically processes server URLs to ensure consistency:
- Removes protocols (`ws://`, `wss://`, `http://`, `https://`)
- Removes endpoints (anything after `/`)
- Stores clean host:port format
:::

## Server Configuration

### File-Based Configuration

**Default location:** `~/.config/gtunnel/config.yaml`

```yaml
access_token: "your-secure-server-token"
```

### Configuration Fields

| Field | Type | Description | Example |
|-------|------|-------------|---------|
| `access_token` | string | Required token for client authentication | `"secure-server-token-123"` |

### Managing Server Configuration

```bash
# Show current configuration
gts config

# Set access token
gts config --set-token your-secure-token

# Show detailed configuration
gts config --show
```

:::warning Security Notice
Access token is mandatory for secure client-server communication. Ensure you set this before starting the server.
:::

## Environment Variable Configuration

For containerized deployments, serverless platforms, or CI/CD environments, gTunnel server supports environment variable configuration.

### Enabling Environment Mode

Set the `GTUNNEL_USE_ENV` environment variable to enable this mode:

```bash
export GTUNNEL_USE_ENV=true
```

### Server Environment Variables

| Variable | Description | Example | Default |
|----------|-------------|---------|---------|
| `GTUNNEL_USE_ENV` | Enable environment variable configuration | `"true"` | `false` |
| `GTUNNEL_ACCESS_TOKEN` | Server access token | `"secure-token-123"` | - |
| `GTUNNEL_PORT` | Server port (Docker only) | `"8080"` | `7205` |

### Client Environment Variables (Docker)

| Variable | Description | Example | Default |
|----------|-------------|---------|---------|
| `GTUNNEL_SERVER_URL` | Server URL | `"tunnel.example.com:8080"` | - |
| `GTUNNEL_ACCESS_TOKEN` | Access token | `"client-token-123"` | - |
| `GTUNNEL_DEBUG` | Enable debug logging | `"true"` | `"false"` |

### Environment Configuration Example

```bash
# Server with environment variables
export GTUNNEL_USE_ENV=true
export GTUNNEL_ACCESS_TOKEN=secure-server-token

# Start server
gts start --bind-address 0.0.0.0:8080
```

:::note Environment Mode Behavior
When `GTUNNEL_USE_ENV=true` is set:
- Configuration is loaded from environment variables instead of files
- Config file creation and modification commands are disabled
- Perfect for containerized and serverless environments
:::

## Docker Configuration

### Server Container

**Basic deployment:**

```bash
# Using environment variables (recommended)
docker run -d \
  -p 8080:7205 \
  -e GTUNNEL_USE_ENV=true \
  -e GTUNNEL_ACCESS_TOKEN=your-secure-token \
  ghcr.io/b-aj-amar/gtunnel-server:latest
```

**With custom port:**

```bash
docker run -d \
  -p 8080:8080 \
  -e GTUNNEL_USE_ENV=true \
  -e GTUNNEL_ACCESS_TOKEN=your-secure-token \
  -e GTUNNEL_PORT=8080 \
  ghcr.io/b-aj-amar/gtunnel-server:latest
```

### Client Container

```bash
# Basic client setup
docker run -d \
  -e GTUNNEL_SERVER_URL=tunnel.example.com:8080 \
  -e GTUNNEL_ACCESS_TOKEN=your-token \
  -e GTUNNEL_TARGET_PORT=3000 \
  --network your-app-network \
  ghcr.io/b-aj-amar/gtunnel-client:latest
```

