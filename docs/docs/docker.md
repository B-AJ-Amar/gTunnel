---
sidebar_position: 12
---

# Docker Deployment

Complete guide for deploying gTunnel using Docker and Docker Compose.

## Overview

gTunnel provides official Docker images for both client and server components, making deployment simple and consistent across different environments.

## Docker Images

### Available Images

- `gtunnel/server:latest` - gTunnel server
- `gtunnel/client:latest` - gTunnel client
- `gtunnel/server:v1.0.0` - Specific version tags
- `gtunnel/client:v1.0.0` - Specific version tags

### Image Architecture

Multi-architecture support:
- linux/amd64
- linux/arm64
- linux/arm/v7

## Server Deployment

### Basic Server Container

```bash
# Run gTunnel server
docker run -d \
  --name gtunnel-server \
  -p 8080:8080 \
  -p 9090:9090 \
  -e GTUNNEL_DOMAIN=tunnel.example.com \
  -e GTUNNEL_AUTH_REQUIRED=true \
  -e GTUNNEL_AUTH_TOKENS=token1,token2 \
  gtunnel/server:latest
```

### Server with Configuration File

```bash
# Create configuration directory
mkdir -p ./config

# Create server configuration
cat > ./config/server.yml << EOF
server:
  port: 8080
  domain: "tunnel.example.com"
  
auth:
  required: true
  provider: "token"
  tokens:
    - "secure-token-1"
    - "secure-token-2"

logging:
  level: "info"
  format: "json"
EOF

# Run with configuration
docker run -d \
  --name gtunnel-server \
  -p 8080:8080 \
  -v $(pwd)/config:/etc/gtunnel \
  gtunnel/server:latest
```

### Server with TLS

```bash
# Create certificate directory
mkdir -p ./certs

# Copy your certificates
cp server.crt ./certs/
cp server.key ./certs/

# Run with TLS
docker run -d \
  --name gtunnel-server \
  -p 443:443 \
  -v $(pwd)/certs:/etc/ssl/certs \
  -v $(pwd)/config:/etc/gtunnel \
  -e GTUNNEL_TLS_ENABLED=true \
  -e GTUNNEL_TLS_CERT=/etc/ssl/certs/server.crt \
  -e GTUNNEL_TLS_KEY=/etc/ssl/certs/server.key \
  gtunnel/server:latest
```

## Client Deployment

### Basic Client Container

```bash
# Run gTunnel client
docker run -d \
  --name gtunnel-client \
  -e GTUNNEL_SERVER_URL=wss://tunnel.example.com \
  -e GTUNNEL_AUTH_TOKEN=your-token \
  -e GTUNNEL_LOCAL_PORT=3000 \
  -e GTUNNEL_SUBDOMAIN=myapp \
  --network host \
  gtunnel/client:latest
```

### Client with Configuration

```bash
# Create client configuration
cat > ./config/client.yml << EOF
server:
  url: "wss://tunnel.example.com"
  auth_token: "your-secure-token"

tunnels:
  - name: "web-app"
    local_port: 3000
    subdomain: "myapp"
    
  - name: "api"
    local_port: 8080
    subdomain: "api"

logging:
  level: "info"
  format: "text"
EOF

# Run with configuration
docker run -d \
  --name gtunnel-client \
  -v $(pwd)/config:/etc/gtunnel \
  --network host \
  gtunnel/client:latest
```

## Docker Compose

### Basic Setup

```yaml
# docker-compose.yml
version: '3.8'

services:
  gtunnel-server:
    image: gtunnel/server:latest
    ports:
      - "8080:8080"
      - "9090:9090"  # Metrics
    environment:
      - GTUNNEL_DOMAIN=tunnel.example.com
      - GTUNNEL_AUTH_REQUIRED=true
      - GTUNNEL_AUTH_TOKENS=token1,token2
    volumes:
      - ./logs:/var/log/gtunnel
    restart: unless-stopped
```

### Development Environment

```yaml
version: '3.8'

services:
  # Your application
  app:
    build: .
    ports:
      - "3000:3000"
    environment:
      - NODE_ENV=development
    volumes:
      - .:/app
      - /app/node_modules
    command: npm run dev
    
  # gTunnel client
  gtunnel-client:
    image: gtunnel/client:latest
    depends_on:
      - app
    environment:
      - GTUNNEL_SERVER_URL=wss://tunnel.example.com
      - GTUNNEL_AUTH_TOKEN=${GTUNNEL_TOKEN}
      - GTUNNEL_LOCAL_PORT=3000
      - GTUNNEL_SUBDOMAIN=dev-${USER}
    network_mode: "service:app"
    restart: unless-stopped
```

### Production Stack

```yaml
version: '3.8'

services:
  gtunnel-server:
    image: gtunnel/server:latest
    ports:
      - "443:443"
      - "9090:9090"
    environment:
      - GTUNNEL_DOMAIN=tunnel.mycompany.com
      - GTUNNEL_TLS_ENABLED=true
      - GTUNNEL_TLS_AUTO_CERT=true
      - JWT_SECRET=${JWT_SECRET}
      - DATABASE_URL=${DATABASE_URL}
    volumes:
      - ./config:/etc/gtunnel
      - ./certs:/var/cache/gtunnel/certs
      - ./logs:/var/log/gtunnel
    restart: unless-stopped
    
  postgres:
    image: postgres:15
    environment:
      - POSTGRES_DB=gtunnel
      - POSTGRES_USER=gtunnel
      - POSTGRES_PASSWORD=${DB_PASSWORD}
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: unless-stopped
    
  prometheus:
    image: prom/prometheus:latest
    ports:
      - "9091:9090"
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
      - prometheus_data:/prometheus
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'
      - '--web.console.libraries=/etc/prometheus/console_libraries'
      - '--web.console.templates=/etc/prometheus/consoles'
    restart: unless-stopped
    
  grafana:
    image: grafana/grafana:latest
    ports:
      - "3001:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=${GRAFANA_PASSWORD}
    volumes:
      - grafana_data:/var/lib/grafana
    restart: unless-stopped

volumes:
  postgres_data:
  prometheus_data:
  grafana_data:
```

### Multi-Service Application

```yaml
version: '3.8'

services:
  # Frontend application
  frontend:
    build: ./frontend
    ports:
      - "3000:3000"
    environment:
      - REACT_APP_API_URL=https://api.myapp.tunnel.example.com
    
  # Backend API
  backend:
    build: ./backend
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://user:pass@postgres:5432/myapp
      - REDIS_URL=redis://redis:6379
    depends_on:
      - postgres
      - redis
      
  # Database
  postgres:
    image: postgres:15
    environment:
      - POSTGRES_DB=myapp
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=pass
    volumes:
      - postgres_data:/var/lib/postgresql/data
      
  # Cache
  redis:
    image: redis:7
    
  # gTunnel for frontend
  gtunnel-frontend:
    image: gtunnel/client:latest
    depends_on:
      - frontend
    environment:
      - GTUNNEL_SERVER_URL=wss://tunnel.example.com
      - GTUNNEL_AUTH_TOKEN=${GTUNNEL_TOKEN}
      - GTUNNEL_LOCAL_PORT=3000
      - GTUNNEL_SUBDOMAIN=myapp
    network_mode: "service:frontend"
    
  # gTunnel for backend
  gtunnel-backend:
    image: gtunnel/client:latest
    depends_on:
      - backend
    environment:
      - GTUNNEL_SERVER_URL=wss://tunnel.example.com
      - GTUNNEL_AUTH_TOKEN=${GTUNNEL_TOKEN}
      - GTUNNEL_LOCAL_PORT=8080
      - GTUNNEL_SUBDOMAIN=api-myapp
    network_mode: "service:backend"

volumes:
  postgres_data:
```

## Environment Variables

### Server Environment Variables

- `GTUNNEL_PORT` - Server port (default: 8080)
- `GTUNNEL_DOMAIN` - Server domain
- `GTUNNEL_BIND_ADDRESS` - Bind address (default: 0.0.0.0)
- `GTUNNEL_TLS_ENABLED` - Enable TLS (true/false)
- `GTUNNEL_TLS_CERT` - TLS certificate file path
- `GTUNNEL_TLS_KEY` - TLS private key file path
- `GTUNNEL_TLS_AUTO_CERT` - Enable Let's Encrypt auto-certificates
- `GTUNNEL_AUTH_REQUIRED` - Require authentication (true/false)
- `GTUNNEL_AUTH_TOKENS` - Comma-separated auth tokens
- `JWT_SECRET` - JWT secret key
- `DATABASE_URL` - Database connection string
- `GTUNNEL_LOG_LEVEL` - Log level (debug, info, warn, error)

### Client Environment Variables

- `GTUNNEL_SERVER_URL` - Server WebSocket URL
- `GTUNNEL_AUTH_TOKEN` - Authentication token
- `GTUNNEL_LOCAL_PORT` - Local port to tunnel
- `GTUNNEL_LOCAL_HOST` - Local host (default: localhost)
- `GTUNNEL_SUBDOMAIN` - Requested subdomain
- `GTUNNEL_CUSTOM_DOMAIN` - Custom domain
- `GTUNNEL_INSECURE` - Skip TLS verification (true/false)
- `GTUNNEL_LOG_LEVEL` - Log level

## Docker Networking

### Host Network Mode

```yaml
services:
  gtunnel-client:
    image: gtunnel/client:latest
    network_mode: "host"
    environment:
      - GTUNNEL_SERVER_URL=wss://tunnel.example.com
      - GTUNNEL_LOCAL_PORT=3000
```

### Custom Networks

```yaml
services:
  app:
    build: .
    networks:
      - app-network
      
  gtunnel-client:
    image: gtunnel/client:latest
    networks:
      - app-network
    environment:
      - GTUNNEL_LOCAL_HOST=app
      - GTUNNEL_LOCAL_PORT=3000

networks:
  app-network:
    driver: bridge
```

### Service Discovery

```yaml
services:
  web:
    image: nginx:alpine
    
  api:
    build: ./api
    
  gtunnel-web:
    image: gtunnel/client:latest
    environment:
      - GTUNNEL_LOCAL_HOST=web
      - GTUNNEL_LOCAL_PORT=80
      - GTUNNEL_SUBDOMAIN=web
      
  gtunnel-api:
    image: gtunnel/client:latest
    environment:
      - GTUNNEL_LOCAL_HOST=api
      - GTUNNEL_LOCAL_PORT=8080
      - GTUNNEL_SUBDOMAIN=api
```

## Persistent Storage

### Configuration Persistence

```yaml
services:
  gtunnel-server:
    image: gtunnel/server:latest
    volumes:
      - ./config:/etc/gtunnel:ro
      - gtunnel_data:/var/lib/gtunnel
      - ./logs:/var/log/gtunnel

volumes:
  gtunnel_data:
```

### Certificate Management

```yaml
services:
  gtunnel-server:
    image: gtunnel/server:latest
    volumes:
      - ./certs:/etc/ssl/certs:ro
      - cert_cache:/var/cache/gtunnel/certs
      
volumes:
  cert_cache:
```

## Health Checks

```yaml
services:
  gtunnel-server:
    image: gtunnel/server:latest
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
      
  gtunnel-client:
    image: gtunnel/client:latest
    healthcheck:
      test: ["CMD", "gtunnel-client", "status"]
      interval: 30s
      timeout: 10s
      retries: 3
```

## Troubleshooting

### Common Issues

**Container Won't Start:**
```bash
# Check logs
docker logs gtunnel-server
docker logs gtunnel-client

# Check configuration
docker exec gtunnel-server cat /etc/gtunnel/server.yml
```

**Network Connectivity:**
```bash
# Test from container
docker exec gtunnel-client ping tunnel.example.com
docker exec gtunnel-client nslookup tunnel.example.com
```

**Permission Issues:**
```bash
# Fix volume permissions
sudo chown -R 1000:1000 ./config ./logs
```

For more information, see:
<!-- - [Configuration Guide](./configuration.md)
- [Production Deployment](./production.md)
- [Monitoring Guide](./monitoring.md) -->
