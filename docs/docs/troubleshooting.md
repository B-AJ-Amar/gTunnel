---
sidebar_position: 25
---

# Troubleshooting

Common issues and solutions when using gTunnel.

## Connection Issues

### Cannot Connect to Server

**Symptoms:**
- Client fails to establish connection
- "Connection refused" errors
- Timeout errors

**Solutions:**

1. **Check server URL:**
   ```bash
   # Verify server is running
   curl -I https://tunnel.example.com/health
   ```

2. **Verify authentication:**
   ```yaml
   server:
     auth_token: "correct-token-here"
   ```

3. **Check firewall settings:**
   - Ensure server port is open
   - Check both client and server firewalls

4. **Test network connectivity:**
   ```bash
   # Test basic connectivity
   telnet tunnel.example.com 8080
   ```

### Frequent Disconnections

**Symptoms:**
- Client repeatedly disconnects
- Unstable tunnel connections
- Intermittent errors

**Solutions:**

1. **Increase timeout values:**
   ```yaml
   server:
     timeout: 60s
     retry_interval: 10s
     max_retries: 20
   ```

2. **Check network stability:**
   - Test with different networks
   - Check for proxy interference

3. **Enable compression:**
   ```yaml
   server:
     compression: true
   ```

## Authentication Problems

### Invalid Token Errors

**Symptoms:**
- "Authentication failed" messages
- 401 Unauthorized responses

**Solutions:**

1. **Verify token format:**
   ```yaml
   server:
     auth_token: "your-actual-token-here"
   ```

2. **Check token expiration (JWT):**
   - Generate new JWT token
   - Verify token hasn't expired

3. **Environment variable issues:**
   ```bash
   # Check if environment variable is set
   echo $GTUNNEL_AUTH_TOKEN
   ```

## Performance Issues

### Slow Tunnel Performance

**Symptoms:**
- High latency
- Slow response times
- Timeouts

**Solutions:**

1. **Check server resources:**
   - Monitor CPU and memory usage
   - Check network bandwidth

2. **Optimize configuration:**
   ```yaml
   server:
     compression: true
     ping_interval: 30s
   ```

3. **Network optimization:**
   - Use server closer to clients
   - Check for network bottlenecks

### High Memory Usage

**Symptoms:**
- Server consuming excessive memory
- Out of memory errors

**Solutions:**

1. **Adjust connection limits:**
   ```yaml
   limits:
     max_connections: 500
     max_tunnels_per_client: 5
   ```

2. **Enable request size limits:**
   ```yaml
   limits:
     max_request_size: "10MB"
     max_bandwidth_per_client: "100MB/s"
   ```

## Configuration Issues

### Configuration File Not Found

**Symptoms:**
- "Config file not found" errors
- Using default settings unexpectedly

**Solutions:**

1. **Specify config path explicitly:**
   ```bash
   gtunnel-client --config /path/to/config.yml connect
   ```

2. **Check file permissions:**
   ```bash
   ls -la ~/.gtunnel/client.yml
   chmod 600 ~/.gtunnel/client.yml
   ```

3. **Create default config:**
   ```bash
   gtunnel-client config init
   ```

### Invalid Configuration Syntax

**Symptoms:**
- YAML parsing errors
- "Invalid configuration" messages

**Solutions:**

1. **Validate YAML syntax:**
   ```bash
   # Use online YAML validator or
   python -c "import yaml; yaml.safe_load(open('config.yml'))"
   ```

2. **Check indentation:**
   - Use spaces, not tabs
   - Maintain consistent indentation

3. **Validate with built-in command:**
   ```bash
   gtunnel-server config validate
   ```

## TLS/SSL Issues

### Certificate Verification Failed

**Symptoms:**
- SSL certificate errors
- "Certificate verification failed"

**Solutions:**

1. **For development only:**
   ```yaml
   server:
     insecure: true  # Skip certificate verification
   ```

2. **Fix certificate issues:**
   - Ensure certificate is valid
   - Check certificate chain
   - Verify domain matches certificate

3. **Let's Encrypt issues:**
   ```yaml
   server:
     tls:
       auto_cert: true
       cache_dir: "/var/cache/gtunnel"
   ```

## Subdomain Issues

### Subdomain Not Working

**Symptoms:**
- 404 errors on subdomain
- Subdomain doesn't resolve

**Solutions:**

1. **Check DNS configuration:**
   ```bash
   dig myapp.tunnel.example.com
   nslookup myapp.tunnel.example.com
   ```

2. **Verify server domain config:**
   ```yaml
   subdomains:
     enabled: true
     wildcard_domain: "*.tunnel.example.com"
   ```

3. **Check subdomain patterns:**
   ```yaml
   subdomains:
     allowed_patterns:
       - "^[a-z0-9-]+$"  # Ensure pattern matches your subdomain
   ```

### Reserved Subdomain Error

**Symptoms:**
- "Subdomain reserved" errors
- Cannot use desired subdomain

**Solutions:**

1. **Check reserved list:**
   ```yaml
   subdomains:
     reserved:
       - "www"
       - "api"
       - "admin"  # Your subdomain might be here
   ```

2. **Use different subdomain:**
   - Choose non-reserved name
   - Follow allowed patterns

## Local Service Issues

### Service Not Reachable

**Symptoms:**
- 502 Bad Gateway errors
- Connection refused to local service

**Solutions:**

1. **Verify local service is running:**
   ```bash
   curl http://localhost:3000
   netstat -tlnp | grep :3000
   ```

2. **Check binding address:**
   - Ensure service binds to localhost or 0.0.0.0
   - Some services only bind to 127.0.0.1

3. **Verify port configuration:**
   ```yaml
   tunnels:
     - local_port: 3000  # Must match actual service port
   ```

## Docker Issues

### Container Won't Start

**Symptoms:**
- Docker container exits immediately
- Configuration errors in container

**Solutions:**

1. **Check container logs:**
   ```bash
   docker logs gtunnel-server
   docker logs gtunnel-client
   ```

2. **Verify volume mounts:**
   ```bash
   docker run -v /path/to/config:/etc/gtunnel gtunnel/server
   ```

3. **Check environment variables:**
   ```bash
   docker run -e GTUNNEL_AUTH_TOKEN=token gtunnel/client
   ```

## Debugging

### Enable Debug Logging

**Client:**
```yaml
logging:
  level: "debug"
  format: "text"
```

**Server:**
```yaml
logging:
  level: "debug"
  format: "json"
  file: "/var/log/gtunnel-debug.log"
```

### Analyze Logs

**Common log patterns to look for:**

1. **Connection attempts:**
   ```
   INFO: Client connecting from 192.168.1.100
   ERROR: Authentication failed for client
   ```

2. **Tunnel creation:**
   ```
   INFO: Tunnel created: myapp.tunnel.example.com -> localhost:3000
   ERROR: Port 3000 not reachable
   ```

3. **Performance issues:**
   ```
   WARN: High memory usage: 90%
   WARN: Rate limit exceeded for client
   ```

### Network Diagnostics

**Test connectivity:**
```bash
# Test WebSocket connection
wscat -c wss://tunnel.example.com

# Test HTTP endpoint
curl -v https://myapp.tunnel.example.com

# Check DNS resolution
dig +trace myapp.tunnel.example.com
```

**Monitor traffic:**
```bash
# Monitor network traffic
tcpdump -i any port 8080

# Check connections
ss -tulpn | grep gtunnel
```

## Getting Help

If you're still experiencing issues:

1. **Check existing issues:**
   - [GitHub Issues](https://github.com/B-AJ-Amar/gTunnel/issues)

2. **Gather diagnostic information:**
   - gTunnel version
   - Operating system
   - Configuration files (remove sensitive data)
   - Error messages and logs

3. **Create detailed issue report:**
   - Describe expected vs actual behavior
   - Include steps to reproduce
   - Attach relevant logs

4. **Community support:**
   - [GitHub Discussions](https://github.com/B-AJ-Amar/gTunnel/discussions)

For more help, see:
<!-- - [FAQ](./faq.md)
- [Configuration Guide](./configuration.md)
- [Security Guide](./security.md) -->
