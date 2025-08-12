---
slug: welcome-to-gtunnel
title: Welcome to gTunnel!
authors: [amar]
tags: [announcement, getting-started, wellcome]
---

We're excited to introduce **gTunnel** - a fast, secure, and lightweight HTTP tunneling solution built with Go! 🚀

<!--truncate-->

## What is gTunnel?

gTunnel makes it incredibly easy to expose your local development servers to the internet securely. Whether you're:

- 👥 **Collaborating with teammates** on a local development server
- 🔗 **Testing webhooks** from external services  
- 🎯 **Demoing your application** to clients or stakeholders
- 🧪 **Debugging integrations** with third-party APIs

gTunnel has you covered with a simple, fast, and secure tunneling solution.

## Key Features

### 🚀 High Performance

Built with Go, gTunnel delivers exceptional performance with minimal resource usage. Handle thousands of concurrent connections without breaking a sweat.

### 🔒 Security First

- WebSocket-based secure tunneling
- Authentication and authorization support
- No data logging or storage
- Self-hosted options for complete control

### 🐳 Docker Ready

Full Docker support with multi-architecture images available on GitHub Container Registry:

```bash
docker run ghcr.io/b-aj-amar/gtunnel-server:latest
```

### 🎯 Developer Friendly

- One-command installation
- Minimal configuration required
- Comprehensive CLI with helpful commands
- Detailed documentation and examples

## Quick Start

Get started in under 2 minutes:

```bash
# Install gTunnel (On the server)
curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash

# Start the server
gts start 

# on the client side ====================
# Connect your local service
gtc connect -u your-domain 3000
```

That's it! Your local service is now accessible through the tunnel.

**Want more detailed instructions?** Check out our comprehensive [Quick Start Guide](/quick-start) for step-by-step setup, Docker usage, and common use cases.


## Getting Involved

gTunnel is open source and we welcome contributions! Here's how you can get involved:

- ⭐ **Star us on GitHub** to show your support
- 🐛 **Report issues** or suggest features
- 💻 **Contribute code** with pull requests
- 📖 **Improve documentation** and examples
- 💬 **Join discussions** in our GitHub community

## Resources

- 📚 [Quick Start Guide](/quick-start) - Get running in 5 minutes
- 📖 [Full Documentation](/docs/intro) - Comprehensive guides
- 🔧 [GitHub Repository](https://github.com/B-AJ-Amar/gTunnel) - Source code
- 🐳 [Docker Images](https://ghcr.io/b-aj-amar/gtunnel-server) - Container images
- 💬 [Discussions](https://github.com/B-AJ-Amar/gTunnel/discussions) - Community support

## Thank You

Thank you for trying gTunnel! We're committed to making tunneling simple, secure, and accessible for all developers.

Have questions or feedback? We'd love to hear from you in our [GitHub discussions](https://github.com/B-AJ-Amar/gTunnel/discussions)!

Happy tunneling! 🚇✨
