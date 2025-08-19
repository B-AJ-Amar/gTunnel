---
sidebar_position: 4
---

# CLI Reference

Complete command-line interface reference for gTunnel client and server components.

## Client Commands

### gtc (gTunnel Client)

The gTunnel client connects to a gTunnel server and forwards local traffic through the tunnel.

#### Basic Usage

```bash
gtc [command]
```

#### Global Flags

- `--help`, `-h`: Show help information
- `--version`: Show version information

### Available Commands

#### connect

Connect to a gTunnel server and start tunneling.

```bash
gtc connect <port|host:port> [flags]
```

**Arguments:**
- `<port|host:port>`: The local service to tunnel
  - Port only: `3000` (defaults to localhost:3000)
  - Host and port: `myapp.local:3000`

**Flags:**
- `--server-url`, `-u`: Server URL (without WebSocket endpoint, e.g., example.com:443)
- `--base-endpoint`, `-e`: Base endpoint path to route the tunneled app
- `--debug`, `-d`: Enable debug logging

**Examples:**
```bash
# Tunnel to localhost:3000
gtc connect 3000

# Tunnel to a specific host and port
gtc connect api.example.com:8080

# Override server URL for this connection
gtc connect -u example.com:9000 3000

# Use HTTPS URL (automatically uses port 443)
gtc connect -u https://example.com 3000

# Enable debug logging
gtc connect -d 3000
```

:::note
- The WebSocket endpoint (`/___gTl___/ws`) is automatically appended
- For HTTPS URLs, port 443 is automatically used if no port is specified
- Server URL is loaded from configuration if not provided via flag
:::

#### config

Manage client configuration settings.

```bash
gtc config [flags]
```

**Flags:**
- `--show`, `-s`: Show current configuration
- `--set-url`, `-u`: Set the server URL
- `--set-token`, `-t`: Set the access token

**Examples:**
```bash
# Show current configuration
gtc config
gtc config --show

# Set server URL (protocol and endpoint will be cleaned)
gtc config --set-url example.com:8080
gtc config --set-url ws://example.com:8080/ws

# Set access token
gtc config --set-token abc123def456
```

:::tip
- URLs are automatically cleaned (protocols and endpoints removed)
- Configuration is stored in `~/.config/gtunnel/config.yaml`
:::

#### status

Display client connection status to the gTunnel server.

```bash
gtc status [flags]
```

**Flags:**
- `--verbose`, `-v`: Show detailed connection information

**Examples:**
```bash
# Show basic connection status
gtc status

# Show detailed information including response time and HTTP status
gtc status -v
```

:::info Output
- **Basic**: Shows if connected or not
- **Verbose**: Includes server URL, response time, and HTTP status details
:::

#### version

Show version information.

```bash
gtc version [flags]
```

**Flags:**
- `--output`, `-o`: Output format. One of: `default|json|short`

**Examples:**
```bash
# Show default version info
gtc version

# Show JSON format
gtc version -o json

# Show only version number
gtc version -o short
```

#### completion

Generate shell autocompletion script.

```bash
gtc completion [bash|zsh|fish|powershell]
```

**Examples:**
```bash
# Generate bash completion
gtc completion bash

# Generate zsh completion
gtc completion zsh
```

## Server Commands

### gts (gTunnel Server)

The gTunnel server handles incoming tunnel connections and routes traffic.

#### Basic Usage

```bash
gts [command]
```

#### Global Flags

- `--help`, `-h`: Show help information
- `--version`: Show version information

### Available Commands

#### start

Start the gTunnel server.

```bash
gts start [flags]
```

**Flags:**
- `--bind-address`: Address to bind the server to (default: `0.0.0.0:7205`)
- `--debug`, `-d`: Enable debug logging

**Examples:**
```bash
# Start server with default settings
gts start

# Start server on specific address and port
gts start --bind-address 0.0.0.0:8080

# Start server with debug logging
gts start -d
```

#### config

Manage server configuration settings.

```bash
gts config [flags]
```

**Flags:**
- `--show`, `-s`: Show current configuration
- `--set-token`, `-t`: Set the access token

**Examples:**
```bash
# Show current configuration
gts config
gts config --show

# Set access token
gts config --set-token abc123def456
```

:::note
- Configuration is stored in `~/.config/gtunnel/config.yaml`
- Access token is required for secure connections
:::

#### status

Show server status.

```bash
gts status
```

:::warning
Currently shows a mock status. Real status implementation coming soon.
:::

#### version

Show version information.

```bash
gts version [flags]
```

**Flags:**
- `--output`, `-o`: Output format. One of: `default|json|short`

#### completion

Generate shell autocompletion script.

```bash
gts completion [bash|zsh|fish|powershell]
```

## Configuration

For detailed configuration information including file-based and environment variable configuration, Docker setup, and Kubernetes deployment, see the [Configuration documentation](./configuration.md).

### Quick Configuration Reference

**Client:**
```bash
gtc config --set-url example.com:8080
gtc config --set-token your-token
```

**Server:**
```bash
gts config --set-token your-secure-token
```

**Environment Variables (Server):**
```bash
export GTUNNEL_USE_ENV=true
export GTUNNEL_ACCESS_TOKEN=your-token
```



## Environment Variables

### Server Environment Variables

For containerized or serverless deployments, you can use environment variables instead of configuration files:

- `GTUNNEL_USE_ENV`: Set to `"true"` to enable environment variable configuration mode
- `GTUNNEL_ACCESS_TOKEN`: Server access token (when `GTUNNEL_USE_ENV=true`)

:::note Environment Configuration Mode
When `GTUNNEL_USE_ENV=true` is set, the server will:
- Load configuration from environment variables instead of files
- Disable config file creation and modification commands
- Work seamlessly in containerized and serverless environments
:::

### Client Environment Variables

Currently, gTunnel client uses configuration files. Environment variable support for the client may be added in future versions.

## Troubleshooting

### Common Issues

1. **"Not configured" error**: Run `gtc config --set-url <server-url>` to set the server URL
2. **Connection failures**: Use `gtc status -v` to check server connectivity  
3. **Authentication errors**: Ensure access token is set with `gtc config --set-token <token>`

:::tip Troubleshooting Steps
1. Check configuration: `gtc config`
2. Test connectivity: `gtc status -v`
3. Verify server is running: `gts status`
4. Enable debug mode: `gtc connect -d <port>`
:::

### Debug Mode

Enable debug logging with the `-d` flag to see detailed connection information:

```bash
# Client debug mode
gtc connect -d 3000

# Server debug mode
gts start -d
```
