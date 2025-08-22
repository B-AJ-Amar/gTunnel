---
sidebar_position: 1
---

# Introduction

Welcome to the comprehensive gTunnel documentation! This guide will help you understand, configure, and deploy gTunnel for your tunneling needs.

## What is gTunnel?

gTunnel is a fast, secure, and lightweight HTTP tunneling solution written in Go. It allows you to expose local services to the internet securely, making it perfect for:

- **Local Development**: Share your work-in-progress with team members
- **Webhook Testing**: Test webhooks locally during development  
- **Remote Access**: Access services running on remote machines
- **Demos & Presentations**: Quickly expose local apps for demonstrations

## Key Features

- 🚀 **High Performance**: Built with Go for speed and efficiency
- 🔒 **Secure**: WebSocket-based tunneling with authentication support
- 🐳 **Docker Ready**: Full Docker support with multi-arch images
- 🎯 **Simple**: Easy installation and minimal configuration
- 🔧 **Flexible**: Support for HTTP, HTTPS, and custom protocols
- 📦 **Self-Hosted**: Deploy your own tunnel infrastructure

## Architecture

gTunnel consists of two main components:

### Server Component
- Handles incoming tunnel connections
- Routes traffic between clients and external requests
- Manages authentication and authorization
- Provides monitoring and logging

### Client Component  
- Connects to gTunnel servers
- Forwards local traffic through the tunnel
- Handles reconnection and error recovery
- Supports multiple tunnel configurations

## Quick Navigation

### Getting Started
- [Quick Start Guide](./getting-started/quick-start) - Get running in 5 minutes
- [Installation](./getting-started/installation) - Detailed installation options
- [Basic Configuration](./configuration) - Essential configuration

### Core Features
- [CLI Reference](./cli-reference) - Complete command reference
- [Configuration Guide](./configuration) - Detailed configuration options
- [Deployment Guide](./deployment) - Deployment and hosting options

### Advanced Topics
- [FAQ](./faq) - Frequently asked questions and troubleshooting

## Community & Support

- 📋 [FAQ](./faq) - Frequently asked questions
- 🐛 [Issue Tracker](https://github.com/B-AJ-Amar/gTunnel/issues) - Report bugs and request features  
- 💬 [Discussions](https://github.com/B-AJ-Amar/gTunnel/discussions) - Community discussions
- 📖 [Contributing](./contributing) - How to contribute to the project

## Version Information

This documentation covers gTunnel v0.0.0 and later. For older versions, please refer to the [changelog](./changelog).

---

**New to gTunnel?** Start with our [Quick Start Guide](./getting-started/quick-start) to get up and running in minutes!

**Need specific help?** Use the navigation menu to find detailed information about any aspect of gTunnel.
