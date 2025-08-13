---
sidebar_position: 22
---

# Examples

Common use cases and implementation examples for gTunnel.

## Basic Examples

### Local Development Server

Expose a local development server to share with team members:

**Client Configuration:**

```yaml
server:
  url: "wss://tunnel.example.com"
  auth_token: "dev-token"

tunnels:
  - name: "dev-app"
    local_port: 3000
    subdomain: "myapp-dev"
```

**Command Line:**

```bash
gtunnel-client connect --server wss://tunnel.example.com --port 3000 --subdomain myapp-dev
```

### API Development

Expose a local API for webhook testing:

```yaml
server:
  url: "wss://tunnel.example.com"
  auth_token: "api-token"

tunnels:
  - name: "webhook-api"
    local_port: 8080
    subdomain: "webhooks"
    auth:
      basic:
        username: "api"
        password: "secure-password"
```

### Database Administration

Securely expose database admin tools:

```yaml
tunnels:
  - name: "db-admin"
    local_port: 5432
    subdomain: "dbadmin"
    auth:
      basic:
        username: "admin"
        password: "admin-password"
```

## Advanced Examples

### Multi-Service Development Environment

Expose multiple services for a microservices architecture:

```yaml
server:
  url: "wss://tunnel.example.com"
  auth_token: "microservices-token"

tunnels:
  - name: "frontend"
    local_port: 3000
    subdomain: "app"
    
  - name: "user-service"
    local_port: 8001
    subdomain: "users"
    
  - name: "payment-service"
    local_port: 8002
    subdomain: "payments"
    
  - name: "notification-service"
    local_port: 8003
    subdomain: "notifications"

logging:
  level: "debug"
  format: "json"
```

### Custom Domain Setup

Use your own domain with gTunnel:

```yaml
server:
  url: "wss://tunnel.mycompany.com"
  auth_token: "company-token"

tunnels:
  - name: "production-app"
    local_port: 8080
    custom_domain: "app.mycompany.com"
    
  - name: "staging-app"
    local_port: 3000
    custom_domain: "staging.mycompany.com"
```

### CI/CD Integration

Integrate gTunnel with CI/CD pipelines:

**GitHub Actions Example:**

```yaml
name: Test with gTunnel
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Start application
        run: |
          npm install
          npm start &
          sleep 10
          
      - name: Setup gTunnel
        run: |
          curl -L https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel-client-linux.tar.gz | tar xz
          ./gtunnel-client connect \
            --server wss://ci-tunnel.example.com \
            --auth-token ${{ secrets.GTUNNEL_TOKEN }} \
            --port 3000 \
            --subdomain pr-${{ github.event.number }} &
          
      - name: Run E2E tests
        run: |
          export TEST_URL="https://pr-${{ github.event.number }}.ci-tunnel.example.com"
          npm run test:e2e
```

## Docker Examples

### Client in Docker

**Dockerfile:**

```dockerfile
FROM alpine:latest

RUN apk add --no-cache curl

# Download gTunnel client
RUN curl -L https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel-client-linux.tar.gz \
    | tar xz -C /usr/local/bin/

COPY config.yml /etc/gtunnel/config.yml

EXPOSE 3000

CMD ["gtunnel-client", "--config", "/etc/gtunnel/config.yml", "connect"]
```

**Docker Compose:**

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "3000:3000"
    environment:
      - NODE_ENV=development
      
  gtunnel:
    image: gtunnel/client:latest
    depends_on:
      - app
    environment:
      - GTUNNEL_SERVER_URL=wss://tunnel.example.com
      - GTUNNEL_AUTH_TOKEN=${GTUNNEL_TOKEN}
      - GTUNNEL_LOCAL_PORT=3000
      - GTUNNEL_SUBDOMAIN=myapp
    network_mode: "service:app"
```

### Complete Stack

Full development environment with gTunnel:

```yaml
version: '3.8'

services:
  database:
    image: postgres:15
    environment:
      POSTGRES_DB: myapp
      POSTGRES_USER: dev
      POSTGRES_PASSWORD: dev
    ports:
      - "5432:5432"
      
  redis:
    image: redis:7
    ports:
      - "6379:6379"
      
  backend:
    build: ./backend
    depends_on:
      - database
      - redis
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://dev:dev@database:5432/myapp
      - REDIS_URL=redis://redis:6379
      
  frontend:
    build: ./frontend
    ports:
      - "3000:3000"
    environment:
      - REACT_APP_API_URL=https://api.myapp.tunnel.example.com
      
  gtunnel-api:
    image: gtunnel/client:latest
    depends_on:
      - backend
    environment:
      - GTUNNEL_SERVER_URL=wss://tunnel.example.com
      - GTUNNEL_AUTH_TOKEN=${GTUNNEL_TOKEN}
      - GTUNNEL_LOCAL_PORT=8080
      - GTUNNEL_SUBDOMAIN=api-myapp
    network_mode: "service:backend"
    
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
```

## Kubernetes Examples

### Basic Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gtunnel-client
spec:
  replicas: 1
  selector:
    matchLabels:
      app: gtunnel-client
  template:
    metadata:
      labels:
        app: gtunnel-client
    spec:
      containers:
      - name: app
        image: myapp:latest
        ports:
        - containerPort: 8080
        
      - name: gtunnel
        image: gtunnel/client:latest
        env:
        - name: GTUNNEL_SERVER_URL
          value: "wss://tunnel.example.com"
        - name: GTUNNEL_AUTH_TOKEN
          valueFrom:
            secretKeyRef:
              name: gtunnel-secret
              key: auth-token
        - name: GTUNNEL_LOCAL_PORT
          value: "8080"
        - name: GTUNNEL_SUBDOMAIN
          value: "k8s-app"
```

### Sidecar Pattern

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-with-tunnel
spec:
  template:
    spec:
      containers:
      - name: main-app
        image: myapp:latest
        ports:
        - containerPort: 8080
        
      - name: gtunnel-sidecar
        image: gtunnel/client:latest
        env:
        - name: GTUNNEL_CONFIG
          valueFrom:
            configMapKeyRef:
              name: gtunnel-config
              key: config.yaml
        volumeMounts:
        - name: gtunnel-config
          mountPath: /etc/gtunnel
          
      volumes:
      - name: gtunnel-config
        configMap:
          name: gtunnel-config
```

## Server Examples

### Basic Server Setup

```yaml
server:
  port: 8080
  domain: "tunnel.example.com"
  
auth:
  required: true
  provider: "token"
  tokens:
    - "dev-team-token"
    - "ci-cd-token"
    - "staging-token"

subdomains:
  enabled: true
  wildcard_domain: "*.tunnel.example.com"
  reserved:
    - "www"
    - "api"
    - "admin"

logging:
  level: "info"
  format: "json"
  file: "/var/log/gtunnel/server.log"
```

### Production Server

```yaml
server:
  port: 443
  domain: "tunnel.mycompany.com"
  tls:
    enabled: true
    auto_cert: true
    cache_dir: "/var/cache/gtunnel/certs"

auth:
  required: true
  provider: "jwt"
  jwt:
    secret: "${JWT_SECRET}"
    issuer: "gtunnel-prod"
    audience: "company-clients"
    expiration: "24h"

limits:
  max_connections: 5000
  max_tunnels_per_client: 20
  rate_limit:
    requests_per_minute: 1000
    burst: 100

database:
  type: "postgres"
  connection: "${DATABASE_URL}"
  max_connections: 50

monitoring:
  enabled: true
  metrics_port: 9090
  health_check_path: "/health"

logging:
  level: "info"
  format: "json"
  file: "/var/log/gtunnel/server.log"
  access_log: "/var/log/gtunnel/access.log"
  rotate:
    enabled: true
    max_size: "100MB"
    max_files: 30
```

## Integration Examples

### Webhook Testing

Test webhooks locally during development:

```bash
# Start your webhook handler
node webhook-handler.js &

# Expose it via gTunnel
gtunnel-client connect \
  --server wss://tunnel.example.com \
  --port 3000 \
  --subdomain webhooks

# Configure webhook URL in external service
# https://webhooks.tunnel.example.com/webhook
```

### API Testing

Test API integrations with external services:

```yaml
# config.yml
server:
  url: "wss://tunnel.example.com"
  auth_token: "api-testing-token"

tunnels:
  - name: "api-server"
    local_port: 8080
    subdomain: "api-test"
    headers:
      X-Forwarded-Proto: "https"
      X-Custom-Header: "api-test"
```

### Mobile App Development

Test mobile apps against local backends:

```bash
# Start local backend
npm run dev &

# Expose via gTunnel
gtunnel-client connect \
  --server wss://tunnel.example.com \
  --port 3000 \
  --subdomain mobile-backend

# Configure mobile app to use:
# https://mobile-backend.tunnel.example.com
```

For more examples and use cases, see:
<!-- - [Configuration Guide](./configuration.md)
- [CLI Reference](./cli-reference.md)
- [Production Deployment](./production.md) -->
