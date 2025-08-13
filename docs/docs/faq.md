---
sidebar_position: 20
---

# Frequently Asked Questions

Common questions and answers about gTunnel.

## General Questions

### What is gTunnel?

gTunnel is a fast, secure, and lightweight HTTP tunneling solution written in Go. It allows you to expose local services to the internet securely through WebSocket connections.

### How does gTunnel work?

gTunnel consists of two components:
- **Server**: Runs on a public server and handles incoming HTTP requests
- **Client**: Runs locally and forwards traffic from the server to your local services

The client establishes a WebSocket connection to the server, and all HTTP traffic is tunneled through this connection.

### Is gTunnel free to use?

Yes, gTunnel is open source and free to use. You can run your own server or use community-provided servers.

### How does gTunnel compare to other tunneling solutions?

gTunnel focuses on:
- High performance with Go's efficient networking
- Simple setup and configuration
- Self-hosted infrastructure control
- WebSocket-based communication for reliability

## Installation & Setup

### How do I install gTunnel?

<!-- See our [Installation Guide](./installation.md) for detailed instructions. Quick options: -->

```bash
# Using Go
go install github.com/B-AJ-Amar/gTunnel/cmd/client@latest
go install github.com/B-AJ-Amar/gTunnel/cmd/server@latest

# Using Docker
docker pull gtunnel/client
docker pull gtunnel/server
```

### Do I need to run my own server?

You can either:
- Run your own gTunnel server for full control
- Use a community-provided server (if available)
- Use a cloud service that supports gTunnel

### What ports does gTunnel use?

- **Server**: Typically runs on port 8080 (HTTP) or 443 (HTTPS)
- **Client**: Connects to any local port you specify
- **WebSocket**: Uses the same port as the HTTP server

## Configuration

### How do I configure authentication?

gTunnel supports token-based authentication:

```yaml
# Server config
auth:
  required: true
  tokens:
    - "your-secret-token"

# Client config  
server:
  auth_token: "your-secret-token"
```

### Can I use custom domains?

Yes, gTunnel supports both subdomains and custom domains:

```yaml
tunnels:
  - name: "my-app"
    local_port: 3000
    subdomain: "myapp"  # Creates myapp.yourdomain.com
    # OR
    custom_domain: "myapp.example.com"
```

### How do I enable HTTPS?

Configure TLS on the server:

```yaml
server:
  tls:
    enabled: true
    cert_file: "/path/to/cert.pem"
    key_file: "/path/to/key.pem"
```

## Usage & Features

### Can I tunnel multiple services simultaneously?

Yes, configure multiple tunnels in your client config:

```yaml
tunnels:
  - name: "web"
    local_port: 3000
    subdomain: "web"
  - name: "api"
    local_port: 8080
    subdomain: "api"
```

### Does gTunnel support WebSockets?

Yes, gTunnel supports WebSocket connections and properly forwards WebSocket upgrade requests.

### Can I use gTunnel for databases?

gTunnel is primarily designed for HTTP traffic. For database connections, consider:
- Using HTTP-based database APIs
- Setting up proper network security (VPN, SSH tunnels)
- Using database-specific tunneling solutions

### What protocols are supported?

Currently supported:
- HTTP/HTTPS
- WebSockets
- Any protocol that can be tunneled over HTTP

## Troubleshooting

### Connection keeps dropping

Common causes:
- Network instability
- Server overload
- Authentication issues
- Firewall blocking WebSocket connections

Solutions:
- Check network connectivity
- Verify authentication tokens
- Review firewall settings
- Check server logs

### "Connection refused" errors

This usually means:
- Local service isn't running
- Wrong local port configured
- Local service only accepts connections from specific addresses

### Slow performance

Performance can be affected by:
- Network latency between client and server
- Server resource limitations
- Local service performance
- Bandwidth limitations

### SSL/TLS certificate errors

For HTTPS tunnels:
- Ensure server has valid SSL certificates
- Check certificate chain completeness
- Verify domain names match certificates

## Security

### Is gTunnel secure?

gTunnel provides security through:
- WebSocket connections (can be encrypted with TLS)
- Token-based authentication
- Optional client IP restrictions
- Rate limiting capabilities

### Should I expose sensitive services?

Best practices:
- Only expose services that need external access
- Use authentication on both gTunnel and your services
- Monitor access logs regularly
- Use HTTPS for sensitive data
- Consider IP restrictions

### How do I monitor access?

Enable access logging in server configuration:

```yaml
logging:
  access_log: "/var/log/gtunnel-access.log"
```

## Advanced Usage

### Can I load balance multiple servers?

Currently, gTunnel doesn't include built-in load balancing, but you can:
- Use a reverse proxy (nginx, HAProxy) in front of multiple gTunnel servers
- Use DNS round-robin for simple load distribution
- Implement application-level load balancing

### Docker deployment

gTunnel provides official Docker images:

```bash
# Run server
docker run -p 8080:8080 gtunnel/server

# Run client
docker run gtunnel/client connect --server wss://tunnel.example.com --port 3000
```

### Kubernetes deployment

<!-- See our [deployment documentation](./deployment.md) for Kubernetes manifests and Helm charts. -->

## Getting Help

### Where can I get support?

<!-- - [GitHub Issues](https://github.com/B-AJ-Amar/gTunnel/issues) - Bug reports and feature requests
- [GitHub Discussions](https://github.com/B-AJ-Amar/gTunnel/discussions) - Community discussions
- [Documentation](./intro.md) - Comprehensive guides -->

### How do I report bugs?

1. Check existing issues first
2. Gather relevant information:
   - gTunnel version
   - Operating system
   - Configuration files (remove sensitive data)
   - Error messages and logs
3. Create a detailed issue report

### How can I contribute?

<!-- See our [Contributing Guide](./contributing.md) for information about: -->
- Code contributions
- Documentation improvements
- Bug reports
- Feature suggestions

## Performance & Limits

### What are the performance limits?

Performance depends on:
- Server hardware and network
- Client network connection
- Number of concurrent tunnels
- Traffic volume and patterns

### How many tunnels can I run?

Default limits (configurable):
- 10 tunnels per client
- 1000 concurrent connections per server
- Rate limiting: 100 requests/minute per client

### Can I increase the limits?

Yes, configure limits in server config:

```yaml
limits:
  max_connections: 5000
  max_tunnels_per_client: 50
  rate_limit:
    requests_per_minute: 1000
```

<!-- Still have questions? Check our [troubleshooting guide](./troubleshooting.md) or ask in [GitHub Discussions](https://github.com/B-AJ-Amar/gTunnel/discussions). -->
