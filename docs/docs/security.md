---
sidebar_position: 15
---

# Security & Authentication

Learn how to secure your gTunnel deployments with authentication, encryption, and access controls.

## Authentication Methods

gTunnel supports multiple authentication methods to secure tunnel connections.

### Token-Based Authentication

The simplest authentication method uses pre-shared tokens.

**Server Configuration:**

```yaml
auth:
  required: true
  provider: "token"
  tokens:
    - "super-secret-token-1"
    - "another-secret-token"
```

**Client Configuration:**

```yaml
server:
  url: "wss://tunnel.example.com"
  auth_token: "super-secret-token-1"
```

### JWT Authentication

For more advanced scenarios, use JWT tokens with expiration and custom claims.

**Server Configuration:**

```yaml
auth:
  required: true
  provider: "jwt"
  jwt:
    secret: "your-jwt-secret-key"
    issuer: "gtunnel"
    audience: "clients"
    expiration: "24h"
```

**Client Configuration:**

```yaml
server:
  url: "wss://tunnel.example.com" 
  auth_token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

## TLS/SSL Encryption

Secure your tunnels with TLS encryption to protect data in transit.

### Server TLS Configuration

**Using Certificate Files:**

```yaml
server:
  port: 443
  tls:
    enabled: true
    cert_file: "/etc/ssl/certs/tunnel.crt"
    key_file: "/etc/ssl/private/tunnel.key"
```

**Using Let's Encrypt (Auto-Cert):**

```yaml
server:
  port: 443
  domain: "tunnel.example.com"
  tls:
    enabled: true
    auto_cert: true
    cache_dir: "/var/cache/gtunnel/certs"
```

### Client TLS Configuration

**Secure Connection:**

```yaml
server:
  url: "wss://tunnel.example.com"  # Use wss:// for TLS
  insecure: false  # Verify server certificates
```

**Development Mode (Skip Verification):**

```yaml
server:
  url: "wss://tunnel.example.com"
  insecure: true  # Skip certificate verification (NOT for production)
```

## Access Controls

### IP Address Restrictions

Limit access to specific IP addresses or networks.

**Server Configuration:**

```yaml
auth:
  required: true
  ip_restrictions:
    enabled: true
    allowed_ips:
      - "192.168.1.0/24"
      - "10.0.0.0/8"
      - "203.0.113.0/24"
    denied_ips:
      - "192.168.1.100"
```

### Rate Limiting

Protect against abuse with rate limiting.

```yaml
limits:
  rate_limit:
    enabled: true
    requests_per_minute: 100
    burst: 20
    block_duration: "5m"
```

### Connection Limits

Control resource usage with connection limits.

```yaml
limits:
  max_connections: 1000
  max_tunnels_per_client: 10
  max_bandwidth_mbps: 100
  idle_timeout: "30m"
```

## Subdomain Security

### Reserved Subdomains

Prevent clients from using reserved subdomains.

```yaml
subdomains:
  enabled: true
  reserved:
    - "www"
    - "api"
    - "admin"
    - "mail"
    - "ftp"
    - "blog"
```

### Subdomain Patterns

Enforce naming conventions for subdomains.

```yaml
subdomains:
  enabled: true
  allowed_patterns:
    - "^[a-z0-9-]{3,20}$"  # Alphanumeric and hyphens, 3-20 chars
  blocked_patterns:
    - ".*admin.*"
    - ".*root.*"
```

## Custom Domain Security

### Domain Verification

Require DNS verification for custom domains.

```yaml
custom_domains:
  enabled: true
  verification_required: true
  dns_verification: true
  verification_record: "_gtunnel-verification"
```

**Verification Process:**

1. Client requests custom domain: `myapp.example.com`
2. Server generates verification token: `abc123def456`
3. Client adds DNS TXT record: `_gtunnel-verification.myapp.example.com` = `abc123def456`
4. Server verifies DNS record before allowing domain

### Domain Ownership

Restrict custom domains to verified owners.

```yaml
custom_domains:
  enabled: true
  allowed_domains:
    - "*.mycompany.com"
    - "*.mydomain.org"
  blocked_domains:
    - "*.example.com"
```

## Monitoring & Auditing

### Access Logging

Log all tunnel connections and requests for security monitoring.

```yaml
logging:
  level: "info"
  access_log: "/var/log/gtunnel-access.log"
  audit_log: "/var/log/gtunnel-audit.log"
  format: "json"
```

**Log Format Example:**

```json
{
  "timestamp": "2025-08-13T10:30:00Z",
  "client_ip": "203.0.113.100",
  "tunnel_id": "abc123",
  "subdomain": "myapp",
  "method": "GET",
  "path": "/api/users",
  "status": 200,
  "bytes": 1024,
  "duration_ms": 45
}
```

### Security Events

Monitor and alert on security-related events.

```yaml
security:
  events:
    failed_auth: true
    suspicious_activity: true
    rate_limit_exceeded: true
  alerts:
    webhook_url: "https://alerts.example.com/webhook"
    email: "security@example.com"
```

## Best Practices

### Authentication Security

1. **Use Strong Tokens**: Generate cryptographically random tokens
   ```bash
   # Generate secure token
   openssl rand -base64 32
   ```

2. **Rotate Tokens Regularly**: Change tokens periodically
3. **Use JWT for Advanced Features**: Expiration, custom claims, etc.
4. **Store Tokens Securely**: Use environment variables or secret managers

### Network Security

1. **Always Use TLS**: Encrypt all tunnel traffic
2. **Firewall Configuration**: Limit server access to necessary ports
3. **VPN Integration**: Consider VPN for additional security layers
4. **Network Segmentation**: Isolate tunnel servers from sensitive systems

### Application Security

1. **Service Authentication**: Secure your local services independently
2. **Input Validation**: Validate all data in your applications
3. **HTTPS Termination**: Use HTTPS for end-to-end encryption
4. **CORS Configuration**: Configure CORS policies appropriately

### Operational Security

1. **Regular Updates**: Keep gTunnel updated to latest versions
2. **Monitor Logs**: Review access and audit logs regularly
3. **Incident Response**: Have procedures for security incidents
4. **Backup Configurations**: Securely backup configuration files

## Security Checklist

### Server Security

- [ ] TLS encryption enabled
- [ ] Strong authentication configured
- [ ] Rate limiting enabled
- [ ] IP restrictions configured (if needed)
- [ ] Reserved subdomains protected
- [ ] Access logging enabled
- [ ] Regular security updates
- [ ] Firewall properly configured

### Client Security

- [ ] Secure token storage
- [ ] TLS certificate verification enabled
- [ ] Local service authentication
- [ ] Network access controls
- [ ] Regular client updates

### Infrastructure Security

- [ ] Server hardening completed
- [ ] Network segmentation implemented
- [ ] Monitoring and alerting configured
- [ ] Backup and recovery procedures
- [ ] Incident response plan
- [ ] Regular security assessments

## Common Security Issues

### Weak Authentication

**Problem**: Using simple or predictable tokens
**Solution**: Generate strong, random tokens and rotate regularly

### Unencrypted Traffic

**Problem**: Using HTTP instead of HTTPS
**Solution**: Always enable TLS encryption

### Overprivileged Access

**Problem**: Giving broader access than necessary
**Solution**: Apply principle of least privilege

### Missing Monitoring

**Problem**: No visibility into tunnel usage
**Solution**: Enable comprehensive logging and monitoring

For more security guidance, see:
<!-- - [Production Deployment](./production.md)
- [Monitoring Guide](./monitoring.md)
- [Troubleshooting](./troubleshooting.md) -->
