---
sidebar_position: 8
---

# Deployment Guide

Deploy your own gTunnel server to share and manage your tunnels. Choose from cloud platforms for quick setup or self-hosted solutions for more control.

## 🚀 Quick Deployment with Render

The easiest way to deploy a gTunnel server is using Render's one-click deployment:

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy?repo=https://github.com/B-AJ-Amar/gTunnel)

**What you get with Render:**

- ✅ **Free HTTPS URL** (e.g., `https://your-app.onrender.com`)
- ✅ **Automatic SSL certificates** - Secure by default
- ✅ **Health checks and auto-restart** - High availability
- ✅ **Auto-deploy on git push** - Continuous deployment
- ✅ **Environment variable management** - Easy configuration
- ✅ **Zero setup required** - Just click and deploy!

:::tip Easy Setup
Render deployment includes all necessary configuration out of the box. Your server will be ready to use immediately after deployment.
:::




## VPS/Dedicated Server

**Using systemd (Linux):**

```bash
# Download and install
wget https://github.com/B-AJ-Amar/gTunnel/releases/latest/download/gtunnel-server_linux_amd64.tar.gz
tar -xzf gtunnel-server_linux_amd64.tar.gz
sudo mv gts /usr/local/bin/

# Create service user
sudo useradd --system --shell /bin/false gtunnel

# Create config
sudo mkdir -p /etc/gtunnel
echo "GTUNNEL_ACCESS_TOKEN=your-secret-token" | sudo tee /etc/gtunnel/server.env

# Create systemd service
sudo tee /etc/systemd/system/gtunnel.service > /dev/null << EOF
[Unit]
Description=gTunnel Server
After=network.target

[Service]
Type=simple
User=gtunnel
ExecStart=/usr/local/bin/gts start --bind-address 0.0.0.0:7205
EnvironmentFile=/etc/gtunnel/server.env
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

# Start and enable service
sudo systemctl daemon-reload
sudo systemctl enable gtunnel
sudo systemctl start gtunnel
```

#### Other Options

For additional deployment methods, see our **[Installation Guide](./getting-started/installation.md)**:



## 🆘 Need Help?

If you encounter issues during deployment:

1. Check our **[FAQ](./faq.md)** for common solutions
2. Review the **[CLI Reference](./cli-reference.md)** for command details
3. Visit our **[GitHub Issues](https://github.com/B-AJ-Amar/gTunnel/issues)** for support
