---
sidebar_position: 5
---

# Server Configuration

Comprehensive guide to configuring the gTunnel server component.

## Configuration Overview

The gTunnel server configuration defines how the server operates, handles client connections, and manages tunnels. Configuration is done through YAML files.

## Basic Configuration

Minimal server configuration:

```yaml
server:
  port: 8080
  domain: "tunnel.example.com"
  
auth:
  required: false
```

## Server Settings

Configure the core server behavior.

### Network Configuration

```yaml
server:
  port: 8080                         # Server listening port
  bind_address: "0.0.0.0"           # Bind address (0.0.0.0 for all interfaces)
  domain: "tunnel.example.com"       # Server domain
  public_url: "https://tunnel.example.com"  # Public URL (optional)
```

### TLS Configuration

```yaml
server:
  tls:
    enabled: true                    # Enable TLS
    cert_file: "/etc/ssl/certs/server.crt"
    key_file: "/etc/ssl/private/server.key"
    auto_cert: false                 # Use Let's Encrypt auto-certificates
    cache_dir: "/var/cache/gtunnel"  # Certificate cache directory
```

### Timeouts and Limits

```yaml
server:
  read_timeout: 30s                 # Read timeout
  write_timeout: 30s                # Write timeout
  idle_timeout: 120s                # Idle connection timeout
  max_header_bytes: 1048576         # Max header size (1MB)
```

## Authentication Configuration

Control who can create tunnels.

### No Authentication

```yaml
auth:
  required: false
```

### Token Authentication

```yaml
auth:
  required: true
  provider: "token"
  tokens:
    - "super-secret-token-1"
    - "another-secret-token"
    - "client-specific-token"
```

### JWT Authentication

```yaml
auth:
  required: true
  provider: "jwt"
  jwt:
    secret: "your-jwt-secret-key"
    issuer: "gtunnel"
    audience: "clients"
    expiration: "24h"
    algorithm: "HS256"
```

## Connection Limits

Control resource usage and prevent abuse.

### Basic Limits

```yaml
limits:
  max_connections: 1000             # Maximum concurrent connections
  max_tunnels_per_client: 10        # Tunnels per client
  connection_timeout: "5m"          # New connection timeout
```

### Advanced Limits

```yaml
limits:
  max_connections: 1000
  max_tunnels_per_client: 10
  max_bandwidth_per_client: "100MB/s"
  max_request_size: "10MB"
  connection_timeout: "5m"
  idle_timeout: "30m"
  
  # Rate limiting
  rate_limit:
    enabled: true
    requests_per_minute: 100
    burst: 20
    block_duration: "5m"
    whitelist_ips:
      - "192.168.1.0/24"
      - "10.0.0.0/8"
```

## Subdomain Configuration

Control subdomain assignment and validation.

### Basic Subdomain Settings

```yaml
subdomains:
  enabled: true                     # Enable subdomain support
  wildcard_domain: "*.tunnel.example.com"  # Wildcard domain
```

### Advanced Subdomain Control

```yaml
subdomains:
  enabled: true
  wildcard_domain: "*.tunnel.example.com"
  
  # Reserved subdomains (not available to clients)
  reserved:
    - "www"
    - "api"
    - "admin"
    - "mail"
    - "ftp"
    - "blog"
    - "app"
    
  # Allowed subdomain patterns
  allowed_patterns:
    - "^[a-z0-9-]{3,20}$"          # 3-20 chars, alphanumeric + hyphens
    - "^test-[a-z0-9]+$"           # Test prefix pattern
    
  # Blocked patterns
  blocked_patterns:
    - ".*admin.*"
    - ".*root.*"
    - ".*system.*"
```

## Custom Domain Support

Allow clients to use their own domains.

### Basic Custom Domains

```yaml
custom_domains:
  enabled: true                     # Enable custom domain support
  verification_required: false     # Require domain verification
```

### Advanced Custom Domain Control

```yaml
custom_domains:
  enabled: true
  verification_required: true      # Require DNS verification
  dns_verification: true
  verification_record: "_gtunnel-verification"
  verification_timeout: "5m"
  
  # Allowed domain patterns
  allowed_domains:
    - "*.mycompany.com"
    - "*.example.org"
    - "app.example.net"
    
  # Blocked domains
  blocked_domains:
    - "*.google.com"
    - "*.facebook.com"
    - "localhost"
```

## Logging Configuration

Configure logging and monitoring.

### Basic Logging

```yaml
logging:
  level: "info"                     # debug, info, warn, error
  format: "json"                    # text or json
```

### Advanced Logging

```yaml
logging:
  level: "info"
  format: "json"
  
  # Log files
  file: "/var/log/gtunnel-server.log"
  access_log: "/var/log/gtunnel-access.log"
  audit_log: "/var/log/gtunnel-audit.log"
  
  # Log rotation
  rotate:
    enabled: true
    max_size: "100MB"
    max_files: 30
    max_age: "90d"
    compress: true
```

## Monitoring Configuration

Enable metrics and health checks.

```yaml
monitoring:
  enabled: true
  metrics_port: 9090               # Prometheus metrics port
  metrics_path: "/metrics"         # Metrics endpoint path
  health_check_path: "/health"     # Health check endpoint
  
  # Health check configuration
  health_check:
    interval: "30s"
    timeout: "5s"
    
  # Prometheus configuration
  prometheus:
    enabled: true
    namespace: "gtunnel"
    labels:
      environment: "production"
      region: "us-east-1"
```

## Database Configuration

Configure persistent storage (optional).

### SQLite (Default)

```yaml
database:
  type: "sqlite"
  connection: "/var/lib/gtunnel/gtunnel.db"
```

### PostgreSQL

```yaml
database:
  type: "postgres"
  connection: "postgres://user:password@localhost/gtunnel"
  max_connections: 10
  ssl_mode: "require"
```

### MySQL

```yaml
database:
  type: "mysql"
  connection: "user:password@tcp(localhost:3306)/gtunnel"
  max_connections: 10
```

## Complete Example

Production-ready server configuration:

```yaml
server:
  port: 443
  bind_address: "0.0.0.0"
  domain: "tunnel.mycompany.com"
  
  tls:
    enabled: true
    auto_cert: true
    cache_dir: "/var/cache/gtunnel/certs"
    
  read_timeout: 30s
  write_timeout: 30s
  idle_timeout: 300s

auth:
  required: true
  provider: "jwt"
  jwt:
    secret: "${JWT_SECRET}"
    issuer: "gtunnel-prod"
    audience: "clients"
    expiration: "24h"

limits:
  max_connections: 5000
  max_tunnels_per_client: 20
  max_bandwidth_per_client: "500MB/s"
  connection_timeout: "10m"
  idle_timeout: "60m"
  
  rate_limit:
    enabled: true
    requests_per_minute: 1000
    burst: 100
    block_duration: "10m"

subdomains:
  enabled: true
  wildcard_domain: "*.tunnel.mycompany.com"
  reserved:
    - "www"
    - "api"
    - "admin"
    - "status"
  allowed_patterns:
    - "^[a-z0-9-]{3,30}$"
  blocked_patterns:
    - ".*admin.*"
    - ".*root.*"

custom_domains:
  enabled: true
  verification_required: true
  dns_verification: true
  allowed_domains:
    - "*.mycompany.com"
    - "*.mydomain.org"

logging:
  level: "info"
  format: "json"
  file: "/var/log/gtunnel-server.log"
  access_log: "/var/log/gtunnel-access.log"
  audit_log: "/var/log/gtunnel-audit.log"
  rotate:
    enabled: true
    max_size: "100MB"
    max_files: 60
    max_age: "90d"

monitoring:
  enabled: true
  metrics_port: 9090
  health_check_path: "/health"
  prometheus:
    enabled: true
    namespace: "gtunnel"
    labels:
      environment: "production"

database:
  type: "postgres"
  connection: "${DATABASE_URL}"
  max_connections: 50
  ssl_mode: "require"
```

## Environment Variables

Override configuration with environment variables:

### Server Variables

- `GTUNNEL_PORT`: Server port
- `GTUNNEL_DOMAIN`: Server domain
- `GTUNNEL_TLS_CERT`: TLS certificate file
- `GTUNNEL_TLS_KEY`: TLS private key file
- `GTUNNEL_CONFIG_PATH`: Configuration file path

### Authentication Variables

- `GTUNNEL_AUTH_REQUIRED`: Require authentication (true/false)
- `GTUNNEL_JWT_SECRET`: JWT secret key
- `GTUNNEL_AUTH_TOKENS`: Comma-separated auth tokens

### Database Variables

- `DATABASE_URL`: Database connection string
- `DATABASE_TYPE`: Database type (sqlite, postgres, mysql)

## Configuration Validation

The server validates configuration on startup:

### Common Validation Errors

**Invalid Port:**
```
Error: port must be between 1 and 65535
```

**Missing TLS Files:**
```
Error: TLS certificate file not found
```

**Invalid Domain:**
```
Error: domain must be a valid hostname
```

**JWT Configuration:**
```
Error: JWT secret is required when using JWT authentication
```

## Management Commands

### Initialize Configuration

```bash
gtunnel-server config init
```

### Show Configuration

```bash
gtunnel-server config show
```

### Validate Configuration

```bash
gtunnel-server config validate
```

### Test Configuration

```bash
gtunnel-server config test --dry-run
```

For more information, see:
<!-- - [Authentication & Security](./security.md)
- [Production Deployment](./production.md)
- [Monitoring Guide](./monitoring.md) -->
