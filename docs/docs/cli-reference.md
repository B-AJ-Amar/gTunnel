---
sidebar_position: 10
---

# CLI Reference

Complete command-line interface reference for gTunnel client and server components.

## Client Commands

### gtunnel-client

The gTunnel client connects to a gTunnel server and forwards local traffic through the tunnel.

#### Basic Usage

```bash
gtunnel-client connect [flags]
```

#### Global Flags

- `--config`, `-c`: Configuration file path (default: `~/.gtunnel/client.yml`)
- `--help`, `-h`: Show help information
- `--version`, `-v`: Show version information

#### Connect Command

Connect to a gTunnel server and start tunneling.

```bash
gtunnel-client connect --server wss://tunnel.example.com --port 8080
```

**Flags:**
- `--server`, `-s`: Server WebSocket URL (required)
- `--port`, `-p`: Local port to tunnel (required)
- `--host`: Local host to tunnel (default: `localhost`)
- `--subdomain`: Requested subdomain
- `--auth-token`: Authentication token
- `--insecure`: Skip TLS verification

#### Configuration Command

Manage client configuration.

```bash
gtunnel-client config [subcommand]
```

**Subcommands:**
- `init`: Initialize configuration file
- `show`: Show current configuration
- `edit`: Edit configuration file

#### Status Command

Show client status and active tunnels.

```bash
gtunnel-client status
```

#### Version Command

Show version information.

```bash
gtunnel-client version
```

## Server Commands

### gtunnel-server

The gTunnel server handles incoming tunnel connections and routes traffic.

#### Basic Usage

```bash
gtunnel-server start [flags]
```

#### Global Flags

- `--config`, `-c`: Configuration file path (default: `/etc/gtunnel/server.yml`)
- `--help`, `-h`: Show help information
- `--version`, `-v`: Show version information

#### Start Command

Start the gTunnel server.

```bash
gtunnel-server start --port 8080 --domain tunnel.example.com
```

**Flags:**
- `--port`, `-p`: Server port (default: `8080`)
- `--domain`, `-d`: Server domain
- `--tls-cert`: TLS certificate file
- `--tls-key`: TLS private key file
- `--auth-required`: Require authentication
- `--max-connections`: Maximum client connections

#### Configuration Command

Manage server configuration.

```bash
gtunnel-server config [subcommand]
```

**Subcommands:**
- `init`: Initialize configuration file
- `show`: Show current configuration
- `edit`: Edit configuration file

#### Status Command

Show server status and connected clients.

```bash
gtunnel-server status
```

#### Version Command

Show version information.

```bash
gtunnel-server version
```

## Configuration Files

### Client Configuration

Default location: `~/.gtunnel/client.yml`

```yaml
server:
  url: "wss://tunnel.example.com"
  auth_token: "your-auth-token"
  insecure: false

tunnels:
  - name: "web-dev"
    local_port: 3000
    subdomain: "myapp"
  - name: "api-dev"
    local_port: 8080
    subdomain: "api"
```

### Server Configuration

Default location: `/etc/gtunnel/server.yml`

```yaml
server:
  port: 8080
  domain: "tunnel.example.com"
  tls:
    cert_file: "/path/to/cert.pem"
    key_file: "/path/to/key.pem"

auth:
  required: true
  tokens:
    - "token-1"
    - "token-2"

limits:
  max_connections: 100
  max_tunnels_per_client: 5
```

## Environment Variables

### Client

- `GTUNNEL_SERVER_URL`: Default server URL
- `GTUNNEL_AUTH_TOKEN`: Default authentication token
- `GTUNNEL_CONFIG_PATH`: Configuration file path

### Server

- `GTUNNEL_PORT`: Server port
- `GTUNNEL_DOMAIN`: Server domain
- `GTUNNEL_TLS_CERT`: TLS certificate file
- `GTUNNEL_TLS_KEY`: TLS private key file
- `GTUNNEL_CONFIG_PATH`: Configuration file path

## Exit Codes

- `0`: Success
- `1`: General error
- `2`: Configuration error
- `3`: Connection error
- `4`: Authentication error
- `5`: Permission error

## Examples

### Quick Start

```bash
# Start a tunnel for local development server
gtunnel-client connect --server wss://tunnel.example.com --port 3000

# Start server with custom domain
gtunnel-server start --port 8080 --domain tunnel.example.com
```

### Advanced Usage

```bash
# Connect with authentication
gtunnel-client connect \
  --server wss://tunnel.example.com \
  --port 8080 \
  --auth-token your-token \
  --subdomain myapp

# Start server with TLS
gtunnel-server start \
  --port 443 \
  --domain tunnel.example.com \
  --tls-cert /path/to/cert.pem \
  --tls-key /path/to/key.pem
```

<!-- For more examples and use cases, see our [examples documentation](./examples.md). -->
