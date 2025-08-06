#!/bin/bash

# gTunnel Installation Script
# This script installs the latest version of gTunnel client (gtc)
# Usage: curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
REPO="B-AJ-Amar/gTunnel"
INSTALL_DIR="/usr/local/bin"
TMP_DIR="/tmp/gtunnel-install"

# Default installation mode
INSTALL_MODE="client"  # Options: client, server, both

# Platform detection
detect_platform() {
    local os=$(uname -s | tr '[:upper:]' '[:lower:]')
    local arch=$(uname -m)
    
    case $os in
        linux*)
            OS="linux"
            ;;
        darwin*)
            OS="darwin"
            ;;
        mingw*|msys*|cygwin*)
            OS="windows"
            ;;
        *)
            echo -e "${RED}Error: Unsupported operating system: $os${NC}"
            exit 1
            ;;
    esac
    
    case $arch in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        i386|i686)
            ARCH="386"
            ;;
        *)
            echo -e "${RED}Error: Unsupported architecture: $arch${NC}"
            exit 1
            ;;
    esac
    
    echo -e "${BLUE}Detected platform: ${OS}_${ARCH}${NC}"
}

# Get latest release version
get_latest_version() {
    echo -e "${BLUE}Fetching latest release information...${NC}"
    
    # Try to get version from GitHub API
    if command -v curl >/dev/null 2>&1; then
        VERSION=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    elif command -v wget >/dev/null 2>&1; then
        VERSION=$(wget -qO- "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
    else
        echo -e "${RED}Error: Neither curl nor wget is available${NC}"
        exit 1
    fi
    
    if [ -z "$VERSION" ]; then
        echo -e "${YELLOW}Warning: Could not fetch latest version, using 'latest'${NC}"
        VERSION="latest"
    else
        echo -e "${GREEN}Latest version: ${VERSION}${NC}"
    fi
}

# Download and extract binary
download_binary() {
    local component=$1  # client or server
    local binary_name=""
    
    if [ "$component" = "client" ]; then
        binary_name="gtc"
        archive_name="gtunnel-client"
    elif [ "$component" = "server" ]; then
        binary_name="gts"
        archive_name="gtunnel-server"
    else
        echo -e "${RED}Error: Invalid component: $component${NC}"
        exit 1
    fi
    
    echo -e "${BLUE}Downloading gTunnel ${component}...${NC}"
    
    # Create temporary directory
    mkdir -p "$TMP_DIR"
    cd "$TMP_DIR"
    
    # Construct download URL
    if [ "$VERSION" = "latest" ]; then
        DOWNLOAD_URL="https://github.com/${REPO}/releases/latest/download/${archive_name}_${OS}_${ARCH}.tar.gz"
    else
        DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${VERSION}/${archive_name}_${OS}_${ARCH}.tar.gz"
    fi
    
    echo -e "${BLUE}Download URL: ${DOWNLOAD_URL}${NC}"
    
    # Download the archive
    if command -v curl >/dev/null 2>&1; then
        curl -L -o "${component}.tar.gz" "$DOWNLOAD_URL"
    elif command -v wget >/dev/null 2>&1; then
        wget -O "${component}.tar.gz" "$DOWNLOAD_URL"
    else
        echo -e "${RED}Error: Neither curl nor wget is available${NC}"
        exit 1
    fi
    
    # Extract the archive
    echo -e "${BLUE}Extracting ${component} archive...${NC}"
    tar -xzf "${component}.tar.gz"
    
    # Find the binary (it might be in a subdirectory)
    BINARY_PATH=$(find . -name "$binary_name" -type f | head -n 1)
    
    if [ -z "$BINARY_PATH" ]; then
        echo -e "${RED}Error: ${component} binary not found in archive${NC}"
        exit 1
    fi
    
    echo -e "${GREEN}${component} binary found: ${BINARY_PATH}${NC}"
    
    # Return the binary path and name for installation
    echo "$BINARY_PATH|$binary_name"
}

# Install binary
install_binary() {
    local binary_path=$1
    local binary_name=$2
    
    echo -e "${BLUE}Installing ${binary_name}...${NC}"
    
    # Check if we need sudo
    if [ -w "$INSTALL_DIR" ]; then
        SUDO=""
    else
        SUDO="sudo"
        echo -e "${YELLOW}Administrator privileges required for installation${NC}"
    fi
    
    # Copy binary to install directory
    $SUDO cp "$binary_path" "$INSTALL_DIR/$binary_name"
    $SUDO chmod +x "$INSTALL_DIR/$binary_name"
    
    echo -e "${GREEN}${binary_name} installed to ${INSTALL_DIR}/${binary_name}${NC}"
}

# Verify installation
verify_installation() {
    echo -e "${BLUE}Verifying installation...${NC}"
    local success=true
    
    case "$INSTALL_MODE" in
        "client")
            verify_component "gtc" "client"
            ;;
        "server")
            verify_component "gts" "server"
            ;;
        "both")
            verify_component "gtc" "client"
            verify_component "gts" "server"
            ;;
    esac
    
    if [ "$success" = true ]; then
        echo ""
        echo -e "${BLUE}Quick start:${NC}"
        if [ "$INSTALL_MODE" = "client" ] || [ "$INSTALL_MODE" = "both" ]; then
            echo -e "  gtc connect 3000                           # Expose localhost:3000"
            echo -e "  gtc connect -u https://server.com 3000     # Use custom server"
        fi
        if [ "$INSTALL_MODE" = "server" ] || [ "$INSTALL_MODE" = "both" ]; then
            echo -e "  gts start                                  # Start tunnel server"
            echo -e "  gts start --bind-address 0.0.0.0:8080     # Start on custom address"
        fi
        echo -e "  gtc/gts version                            # Show version"
        echo -e "  gtc/gts --help                             # Show help"
    fi
}

# Verify individual component
verify_component() {
    local binary_name=$1
    local component_name=$2
    
    if command -v "$binary_name" >/dev/null 2>&1; then
        VERSION_OUTPUT=$("$binary_name" version 2>/dev/null || echo "version command not available")
        echo -e "${GREEN}✓ ${component_name} installation successful!${NC}"
        echo -e "${GREEN}✓ Binary location: $(which $binary_name)${NC}"
        echo -e "${GREEN}✓ Version: ${VERSION_OUTPUT}${NC}"
    else
        echo -e "${RED}✗ ${component_name} installation verification failed${NC}"
        echo -e "${YELLOW}You may need to restart your shell or add ${INSTALL_DIR} to your PATH${NC}"
        success=false
    fi
}

# Cleanup
cleanup() {
    echo -e "${BLUE}Cleaning up temporary files...${NC}"
    rm -rf "$TMP_DIR"
}

# Handle errors
handle_error() {
    echo -e "${RED}Installation failed!${NC}"
    cleanup
    exit 1
}

# Main installation function
main() {
    echo -e "${GREEN}"
    echo "  ____  _____                        _ "
    echo " / ___|_   _|   _ _ __  _ __   ___| |"
    echo "| |  _  | || | | | '_ \| '_ \ / _ \ |"
    echo "| |_| | | || |_| | | | | | | |  __/ |"
    echo " \____| |_| \__,_|_| |_|_| |_|\___|_|"
    echo ""
    echo "Fast, lightweight tunneling solution"
    echo -e "${NC}"
    echo -e "${GREEN}gTunnel Installation Script${NC}"
    echo -e "${GREEN}============================${NC}"
    echo ""
    
    # Set error trap
    trap handle_error ERR
    
    # Check for required tools
    if ! command -v tar >/dev/null 2>&1; then
        echo -e "${RED}Error: tar is required but not installed${NC}"
        exit 1
    fi
    
    # Detect platform
    detect_platform
    
    # Get latest version
    get_latest_version
    
    # Download and install components based on mode
    case "$INSTALL_MODE" in
        "client")
            result=$(download_binary "client")
            binary_path=$(echo "$result" | cut -d'|' -f1)
            binary_name=$(echo "$result" | cut -d'|' -f2)
            install_binary "$binary_path" "$binary_name"
            ;;
        "server")
            result=$(download_binary "server")
            binary_path=$(echo "$result" | cut -d'|' -f1)
            binary_name=$(echo "$result" | cut -d'|' -f2)
            install_binary "$binary_path" "$binary_name"
            ;;
        "both")
            # Install client
            result=$(download_binary "client")
            binary_path=$(echo "$result" | cut -d'|' -f1)
            binary_name=$(echo "$result" | cut -d'|' -f2)
            install_binary "$binary_path" "$binary_name"
            
            # Install server
            result=$(download_binary "server")
            binary_path=$(echo "$result" | cut -d'|' -f1)
            binary_name=$(echo "$result" | cut -d'|' -f2)
            install_binary "$binary_path" "$binary_name"
            ;;
        *)
            echo -e "${RED}Error: Invalid installation mode: $INSTALL_MODE${NC}"
            exit 1
            ;;
    esac
    
    # Verify
    verify_installation
    
    # Cleanup
    cleanup
    
    echo ""
    echo -e "${GREEN}🎉 gTunnel installation completed successfully!${NC}"
    
    case "$INSTALL_MODE" in
        "client")
            echo -e "${GREEN}Installed: gTunnel Client (gtc)${NC}"
            ;;
        "server")
            echo -e "${GREEN}Installed: gTunnel Server (gts)${NC}"
            ;;
        "both")
            echo -e "${GREEN}Installed: gTunnel Client (gtc) and Server (gts)${NC}"
            ;;
    esac
    
    echo -e "${BLUE}Documentation: https://github.com/${REPO}${NC}"
    echo -e "${BLUE}Issues: https://github.com/${REPO}/issues${NC}"
}

# Handle command line arguments
case "${1:-}" in
    --help|-h)
        echo "gTunnel Installation Script"
        echo ""
        echo "Usage: $0 [options] [component]"
        echo ""
        echo "Components:"
        echo "  client         Install only the client (gtc) - default"
        echo "  server         Install only the server (gts)"
        echo "  both           Install both client and server"
        echo ""
        echo "Options:"
        echo "  --help, -h     Show this help message"
        echo "  --version, -v  Install specific version (e.g., --version v1.0.0)"
        echo ""
        echo "Environment variables:"
        echo "  INSTALL_DIR    Installation directory (default: /usr/local/bin)"
        echo ""
        echo "Examples:"
        echo "  # Install client only (default)"
        echo "  curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash"
        echo ""
        echo "  # Install server only"
        echo "  curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s server"
        echo ""
        echo "  # Install both client and server"
        echo "  curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s both"
        echo ""
        echo "  # Install specific version"
        echo "  curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh | bash -s -- --version v1.0.0 client"
        echo ""
        echo "  # Custom install directory"
        echo "  INSTALL_DIR=~/.local/bin bash <(curl -sSL https://raw.githubusercontent.com/B-AJ-Amar/gTunnel/main/scripts/install.sh) both"
        exit 0
        ;;
    --version|-v)
        if [ -z "${2:-}" ]; then
            echo -e "${RED}Error: --version requires a version argument${NC}"
            exit 1
        fi
        VERSION="$2"
        shift 2
        echo -e "${BLUE}Installing specific version: ${VERSION}${NC}"
        ;;
    client|server|both)
        INSTALL_MODE="$1"
        echo -e "${BLUE}Installation mode: ${INSTALL_MODE}${NC}"
        ;;
    "")
        # Default to client if no argument provided
        INSTALL_MODE="client"
        ;;
    *)
        echo -e "${RED}Error: Invalid argument '$1'${NC}"
        echo -e "${YELLOW}Valid options: client, server, both${NC}"
        echo -e "${YELLOW}Use --help for more information${NC}"
        exit 1
        ;;
esac

# Handle remaining arguments (in case version was specified with component)
if [ -n "${1:-}" ]; then
    case "$1" in
        client|server|both)
            INSTALL_MODE="$1"
            echo -e "${BLUE}Installation mode: ${INSTALL_MODE}${NC}"
            ;;
    esac
fi

# Override install directory if specified
if [ -n "${INSTALL_DIR:-}" ]; then
    echo -e "${BLUE}Using custom install directory: ${INSTALL_DIR}${NC}"
fi

# Run main installation
main
