---
sidebar_position: 8
---

# Deployment Guide

This guide covers deploying gTunnel in various environments.

## Quick Deployment

### Using Docker

```bash
# Run server
docker run -p 8080:8080 gtunnel/server

# Run client
docker run gtunnel/client --server ws://localhost:8080
```

### Using Docker Compose

```yaml
version: '3.8'
services:
  gtunnel-server:
    image: gtunnel/server
    ports:
      - "8080:8080"
    environment:
      - AUTH_TOKEN=your-secret-token
```

## Production Deployment

More detailed production deployment instructions coming soon.

## Cloud Providers

Instructions for deploying on AWS, GCP, Azure, and other cloud providers.
