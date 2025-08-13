---
sidebar_position: 4
---

# Client Configuration

Detailed configuration options for the gTunnel client component.

## Configuration File

The client configuration file uses YAML format and defines how the client connects to servers and manages tunnels.

### Basic Structure

```yaml
server:
  # Server connection settings
  
tunnels:
  # Tunnel definitions
  
logging:
  # Logging configuration
```

## Server Configuration

Configure how the client connects to gTunnel servers.

### Required Settings

```yaml
server:
  url: "wss://tunnel.example.com"    # Server WebSocket URL
  auth_token: "your-auth-token"      # Authentication token
```

### Advanced Server Settings

```yaml
server:
  url: "wss://tunnel.example.com"
  auth_token: "your-auth-token"
  timeout: 30s                       # Connection timeout
  retry_interval: 5s                 # Retry interval on failure
  max_retries: 10                    # Maximum retry attempts
  insecure: false                    # Skip TLS verification (dev only)
  ping_interval: 30s                 # WebSocket ping interval
  compression: true                  # Enable WebSocket compression
```

## Tunnel Configuration

Define which local services to expose through tunnels.

### Basic Tunnel

```yaml
tunnels:
  - name: "web-app"
    local_port: 3000
    local_host: "localhost"
    subdomain: "myapp"
```

### Advanced Tunnel Options

```yaml
tunnels:
  - name: "web-app"
    local_port: 3000
    local_host: "localhost"
    subdomain: "myapp"
    protocol: "http"                 # Protocol: http, https, tcp
    custom_domain: "app.example.com" # Use custom domain instead
    
    # Authentication for this tunnel
    auth:
      basic:
        username: "user"
        password: "password"
      
    # Custom headers
    headers:
      X-Forwarded-Proto: "https"
      X-Custom-Header: "value"
      
    # Health check
    health_check:
      enabled: true
      path: "/health"
      interval: 30s
      timeout: 5s
```

## Authentication Configuration

### Token Authentication

```yaml
server:
  auth_token: "your-secret-token"
```

### JWT Authentication

```yaml
server:
  auth_token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  auth_type: "jwt"
```

### Environment Variable

```yaml
server:
  auth_token: "${GTUNNEL_AUTH_TOKEN}"  # Load from environment
```

## Logging Configuration

### Basic Logging

```yaml
logging:
  level: "info"                      # debug, info, warn, error
  format: "text"                     # text or json
```

### Advanced Logging

```yaml
logging:
  level: "info"
  format: "json"
  file: "/var/log/gtunnel-client.log"
  rotate:
    enabled: true
    max_size: "100MB"
    max_files: 10
    max_age: "30d"
```

## Multiple Tunnels

Configure multiple tunnels for different services:

```yaml
tunnels:
  - name: "frontend"
    local_port: 3000
    subdomain: "app"
    
  - name: "backend-api"
    local_port: 8080
    subdomain: "api"
    
  - name: "websocket-service"
    local_port: 9000
    subdomain: "ws"
    protocol: "http"
    
  - name: "database-admin"
    local_port: 5432
    subdomain: "dbadmin"
    auth:
      basic:
        username: "admin"
        password: "secure-password"
```

## Environment Variables

Override configuration with environment variables:

### Server Settings

- `GTUNNEL_SERVER_URL`: Server URL
- `GTUNNEL_AUTH_TOKEN`: Authentication token
- `GTUNNEL_INSECURE`: Skip TLS verification (true/false)

### Tunnel Settings

- `GTUNNEL_LOCAL_PORT`: Default local port
- `GTUNNEL_SUBDOMAIN`: Default subdomain

### Logging Settings

- `GTUNNEL_LOG_LEVEL`: Log level
- `GTUNNEL_LOG_FORMAT`: Log format (text/json)

## Configuration Examples

### Development Setup

```yaml
server:
  url: "ws://localhost:8080"
  insecure: true
  
tunnels:
  - name: "dev-server"
    local_port: 3000
    subdomain: "dev"
    
logging:
  level: "debug"
  format: "text"
```

### Production Setup

```yaml
server:
  url: "wss://tunnel.production.com"
  auth_token: "${GTUNNEL_AUTH_TOKEN}"
  timeout: 30s
  retry_interval: 5s
  max_retries: 20
  
tunnels:
  - name: "prod-api"
    local_port: 8080
    custom_domain: "api.mycompany.com"
    auth:
      basic:
        username: "${API_USERNAME}"
        password: "${API_PASSWORD}"
    health_check:
      enabled: true
      path: "/health"
      
logging:
  level: "info"
  format: "json"
  file: "/var/log/gtunnel-client.log"
  rotate:
    enabled: true
    max_size: "100MB"
    max_files: 30
```

### Multi-Service Setup

```yaml
server:
  url: "wss://tunnel.example.com"
  auth_token: "secure-token"
  
tunnels:
  - name: "web-frontend"
    local_port: 3000
    subdomain: "app"
    headers:
      X-Forwarded-Proto: "https"
      
  - name: "api-backend"
    local_port: 8080
    subdomain: "api"
    auth:
      basic:
        username: "api-user"
        password: "api-password"
        
  - name: "admin-panel"
    local_port: 9090
    custom_domain: "admin.mysite.com"
    auth:
      basic:
        username: "admin"
        password: "admin-password"
        
  - name: "webhook-receiver"
    local_port: 7000
    subdomain: "webhooks"
    
logging:
  level: "info"
  format: "json"
```

## Configuration Validation

The client validates configuration on startup. Common validation errors:

### Invalid Server URL

```
Error: invalid server URL format
```

**Solution**: Ensure URL uses `ws://` or `wss://` scheme

### Missing Authentication

```
Error: authentication token required
```

**Solution**: Provide `auth_token` in configuration

### Port Conflicts

```
Error: tunnel 'app1' and 'app2' use the same local port
```

**Solution**: Use different local ports for each tunnel

### Invalid Subdomain

```
Error: subdomain 'my_app' contains invalid characters
```

**Solution**: Use only alphanumeric characters and hyphens

## Command Line Overrides

Override configuration file settings with command line flags:

```bash
gtunnel-client connect \
  --config /path/to/config.yml \
  --server wss://tunnel.example.com \
  --auth-token your-token \
  --port 3000 \
  --subdomain myapp
```

## Configuration Management

### Initialize Configuration

Create a default configuration file:

```bash
gtunnel-client config init
```

### Show Configuration

Display current configuration:

```bash
gtunnel-client config show
```

### Validate Configuration

Check configuration without connecting:

```bash
gtunnel-client config validate
```

For more information, see:
<!-- - [Configuration Overview](./configuration.md)
- [Server Configuration](./server-config.md)
- [CLI Reference](./cli-reference.md) -->
