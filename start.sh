#!/bin/sh

# Use PORT environment variable if set, otherwise default to 7205
PORT=${PORT:-7205}

echo "Starting gTunnel server on 0.0.0.0:$PORT"

# Start the server with the correct bind address
exec ./gtunnel-server start --bind-address "0.0.0.0:$PORT"
