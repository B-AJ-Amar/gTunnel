---
sidebar_position: 20
---

# FAQ

Common questions and answers about gTunnel.

## General Questions

<details>
<summary><strong>What is gTunnel?</strong></summary>

gTunnel is a fast, secure, and lightweight HTTP tunneling solution written in Go. It allows you to expose local services to the internet securely through WebSocket connections.
</details>

<details>
<summary><strong>How does gTunnel work?</strong></summary>

gTunnel consists of two components:
- **Server**: Runs on a public server and handles incoming HTTP requests
- **Client**: Runs locally and forwards traffic from the server to your local services

The client establishes a WebSocket connection to the server, and all HTTP traffic is tunneled through this connection.
</details>

<details>
<summary><strong>Is gTunnel free to use?</strong></summary>

Yes, gTunnel is open source and free to use. You can run your own server or use community-provided servers.
</details>

<details>
<summary><strong>Does gTunnel support WebSockets?</strong></summary>

not yet, we are working on it.
</details>

## Installation & Setup

<details>
<summary><strong>How do I install gTunnel?</strong></summary>

See our [Installation Guide](./getting-started/installation.md) for detailed instructions. Quick options:
</details>

<details>
<summary><strong>Do I need to run my own server?</strong></summary>

You can either:
- Run your own gTunnel server for full control
- Use a community-provided server (if available)
- Use a cloud service that supports gTunnel
</details>

<details>
<summary><strong>How do I configure authentication?</strong></summary>

gTunnel (v0.0.0) basic token-based authentication:

```yaml
# Server config
auth_token: "your-secret-token"

# Client config  
auth_token: "your-secret-token"
```
</details>

<details>
<summary><strong>Should I expose sensitive services?</strong></summary>

Best practices:
- Only expose services that need external access
- Use authentication on both gTunnel and your services
- Monitor access logs regularly
- Use HTTPS for sensitive data
</details>

<details>
<summary><strong>Docker deployment</strong></summary>

gTunnel provides official Docker images:

```bash
# Run server
docker run -p 8080:8080 gtunnel/server

# Run client
docker run gtunnel/client connect --server wss://tunnel.example.com --port 3000
```
</details>

## Contributing & Community

<details>
<summary><strong>Where can I get support?</strong></summary>

- [GitHub Issues](https://github.com/B-AJ-Amar/gTunnel/issues) - Bug reports and feature requests
- [GitHub Discussions](https://github.com/B-AJ-Amar/gTunnel/discussions) - Community discussions
- [Documentation](./getting-started/installation.md) - Comprehensive guides
</details>

<details>
<summary><strong>How do I report bugs?</strong></summary>

1. Check existing issues first
2. Gather relevant information:
   - gTunnel version
   - Operating system
   - Configuration files (remove sensitive data)
   - Error messages and logs
3. Create a detailed issue report
</details>

<details>
<summary><strong>How can I contribute?</strong></summary>

See our [Contributing Guide](./contributing.md) for information about:
- Code contributions
- Documentation improvements
- Bug reports
- Feature suggestions
</details>

<details>
<summary><strong>How many tunnels can I run?</strong></summary>

as many as your server can handle, depending on its resources and configuration.
</details>


Still have questions? Ask in [GitHub Discussions](https://github.com/B-AJ-Amar/gTunnel/discussions).
