---
sidebar_position: 3
---

# Configuration

Learn how to configure gTunnel client and server components for your specific needs.

## Overview

gTunnel uses YAML configuration files for both client and server components. Configuration can be managed through:

- Configuration files
- Environment variables  
- Command-line flags

## Client Configuration

### Configuration File Location

The client looks for configuration files in the following order:

1. Path specified by `--config` flag
2. `GTUNNEL_CONFIG_PATH` environment variable
3. `~/.gtunnel/client.yml` (default)
4. `./client.yml` (current directory)

### Basic Client Configuration

Create a basic client configuration file:

```yaml
# ~/.gtunnel/client.yml
server:
  url: "wss://tunnel.example.com"
  auth_token: "your-auth-token"
  timeout: 30s
  
tunnels:
  - name: "web-dev"
    local_port: 3000
    local_host: "localhost"
    subdomain: "myapp"
  - name: "api-dev"  
    local_port: 8080
    subdomain: "api"
```

### Advanced Client Configuration

```yaml
server:
  url: "wss://tunnel.example.com"
  auth_token: "your-auth-token"
  timeout: 30s
  retry_interval: 5s
  max_retries: 10
  insecure: false

tunnels:
  - name: "web-app"
    local_port: 3000
    local_host: "localhost" 
    subdomain: "myapp"
    custom_domain: "myapp.example.com"
    auth:
      basic:
        username: "user"
        password: "pass"
  - name: "database"
    local_port: 5432
    protocol: "tcp"
    subdomain: "db"

logging:
  level: "info"
  format: "json"
  file: "/var/log/gtunnel-client.log"
```

## Server Configuration

### Configuration File Location

The server looks for configuration files in the following order:

1. Path specified by `--config` flag
2. `GTUNNEL_CONFIG_PATH` environment variable
3. `/etc/gtunnel/server.yml` (default)
4. `./server.yml` (current directory)

### Basic Server Configuration

```yaml
# /etc/gtunnel/server.yml
server:
  port: 8080
  domain: "tunnel.example.com"
  
auth:
  required: false
```

### Advanced Server Configuration

```yaml
server:
  port: 8080
  domain: "tunnel.example.com"
  bind_address: "0.0.0.0"
  read_timeout: 30s
  write_timeout: 30s
  idle_timeout: 120s
  
  tls:
    enabled: true
    cert_file: "/etc/ssl/certs/tunnel.crt"
    key_file: "/etc/ssl/private/tunnel.key"
    auto_cert: false
    
auth:
  required: true
  provider: "token"
  tokens:
    - "super-secret-token-1"
    - "super-secret-token-2"
  # Alternative: JWT authentication
  # provider: "jwt"
  # jwt:
  #   secret: "jwt-secret-key"
  #   issuer: "gtunnel"
  #   audience: "clients"

limits:
  max_connections: 1000
  max_tunnels_per_client: 10
  rate_limit:
    requests_per_minute: 100
    burst: 20
  bandwidth:
    upload_mbps: 100
    download_mbps: 100

subdomains:
  enabled: true
  allowed_patterns:
    - "^[a-z0-9-]+$"
  reserved:
    - "www"
    - "api"
    - "admin"
    - "mail"

custom_domains:
  enabled: true
  verification_required: true
  dns_verification: true

logging:
  level: "info"
  format: "json"
  file: "/var/log/gtunnel-server.log"
  access_log: "/var/log/gtunnel-access.log"

monitoring:
  metrics_enabled: true
  metrics_port: 9090
  health_check_path: "/health"
```

## Environment Variables

### Client Environment Variables

- `GTUNNEL_SERVER_URL`: Server WebSocket URL
- `GTUNNEL_AUTH_TOKEN`: Authentication token
- `GTUNNEL_CONFIG_PATH`: Configuration file path
- `GTUNNEL_LOG_LEVEL`: Logging level (debug, info, warn, error)

### Server Environment Variables

- `GTUNNEL_PORT`: Server port
- `GTUNNEL_DOMAIN`: Server domain
- `GTUNNEL_TLS_CERT`: TLS certificate file path
- `GTUNNEL_TLS_KEY`: TLS private key file path
- `GTUNNEL_CONFIG_PATH`: Configuration file path
- `GTUNNEL_LOG_LEVEL`: Logging level

## Configuration Validation

gTunnel validates configuration on startup. Common validation errors:

### Client Validation

- Missing or invalid server URL
- Invalid port numbers
- Conflicting tunnel names
- Invalid subdomain patterns

### Server Validation  

- Invalid port numbers
- Missing TLS certificates when TLS enabled
- Invalid domain names
- Conflicting authentication settings

## Configuration Management

### Initialize Configuration

Generate default configuration files:

```bash
# Client configuration
gtunnel-client config init

# Server configuration  
gtunnel-server config init
```

### Show Current Configuration

Display the active configuration:

```bash
# Show client config
gtunnel-client config show

# Show server config
gtunnel-server config show
```

### Validate Configuration

Validate configuration without starting:

```bash
# Validate client config
gtunnel-client config validate

# Validate server config
gtunnel-server config validate
```

## Next Steps

<!-- - [Client Configuration](./client-config) - Detailed client configuration options
- [Server Configuration](./server-config) - Detailed server configuration options  
- [Security](./security) - Authentication and security configuration
- [CLI Reference](./cli-reference) - Command-line interface reference -->
