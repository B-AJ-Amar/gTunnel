# Makefile for gTunnel

# Version information
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GIT_COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GO_VERSION ?= $(shell go version | awk '{print $$3}')

# Build flags
LDFLAGS = -X github.com/B-AJ-Amar/gTunnel/internal/pkg.Version=$(VERSION) \
          -X github.com/B-AJ-Amar/gTunnel/internal/pkg.GitCommit=$(GIT_COMMIT) \
          -X github.com/B-AJ-Amar/gTunnel/internal/pkg.BuildDate=$(BUILD_DATE) \
          -X github.com/B-AJ-Amar/gTunnel/internal/pkg.GoVersion=$(GO_VERSION)

# Build directory
BUILD_DIR = build

# Binary names
CLIENT_BINARY = gtc
SERVER_BINARY = gts
CLIENT_LONG_NAME = gtunnel-client
SERVER_LONG_NAME = gtunnel-server

.PHONY: all build build-client build-server build-links clean test help install
# .PHONY: all build build-client build-server build-links clean test help install release release-snapshot goreleaser-check

# Default target
all: build

# Build both client and server with links
build: build-client build-server build-links

# Build client
build-client:
	@echo "Building client..."
	@mkdir -p $(BUILD_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(CLIENT_BINARY) ./cmd/client

# Build server
build-server:
	@echo "Building server..."
	@mkdir -p $(BUILD_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(SERVER_BINARY) ./cmd/server

# Create symbolic links for long names
build-links:
	@echo "Creating symbolic links..."
	@mkdir -p $(BUILD_DIR)
	@cd $(BUILD_DIR) && ln -sf $(CLIENT_BINARY) $(CLIENT_LONG_NAME)
	@cd $(BUILD_DIR) && ln -sf $(SERVER_BINARY) $(SERVER_LONG_NAME)
	@echo "Available commands:"
	@echo "  ./$(BUILD_DIR)/$(CLIENT_BINARY) or ./$(BUILD_DIR)/$(CLIENT_LONG_NAME)"
	@echo "  ./$(BUILD_DIR)/$(SERVER_BINARY) or ./$(BUILD_DIR)/$(SERVER_LONG_NAME)"

# Build for development (without version info)
build-dev:
	@echo "Building for development..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(CLIENT_BINARY) ./cmd/client
	go build -o $(BUILD_DIR)/$(SERVER_BINARY) ./cmd/server
	@$(MAKE) build-links

# Install binaries to GOPATH/bin
install: build
	@echo "Installing binaries..."
	go install -ldflags "$(LDFLAGS)" ./cmd/client
	go install -ldflags "$(LDFLAGS)" ./cmd/server
	@echo "Creating symbolic links in GOPATH/bin..."
	@cd $(shell go env GOPATH)/bin && ln -sf $(CLIENT_BINARY) $(CLIENT_LONG_NAME)
	@cd $(shell go env GOPATH)/bin && ln -sf $(SERVER_BINARY) $(SERVER_LONG_NAME)
	@echo "Installed commands:"
	@echo "  $(CLIENT_BINARY) or $(CLIENT_LONG_NAME)"
	@echo "  $(SERVER_BINARY) or $(SERVER_LONG_NAME)"

# Run tests
test:
	go test ./...

# Clean build directory
clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)

# Show help
help:
	@echo "Available targets:"
	@echo "  all          - Build both client and server with links (default)"
	@echo "  build        - Build both client and server with links"
	@echo "  build-client - Build only the client"
	@echo "  build-server - Build only the server"
	@echo "  build-links  - Create symbolic links for long command names"
	@echo "  build-dev    - Build without version info (development)"
	@echo "  install      - Install binaries to GOPATH/bin with links"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build directory"
	@echo "  help         - Show this help message"
	@echo "  release      - Create a release using GoReleaser"
	@echo "  release-snapshot - Create a snapshot release using GoReleaser"
	@echo "  goreleaser-check - Check GoReleaser configuration"
	@echo ""
	@echo "Available commands after build:"
	@echo "  Short names: gtc, gts"
	@echo "  Long names:  gtunnel-client, gtunnel-server"
	@echo ""
	@echo "Environment variables:"
	@echo "  VERSION      - Override version (default: git describe)"
	@echo "  GIT_COMMIT   - Override git commit (default: git rev-parse)"
	@echo "  BUILD_DATE   - Override build date (default: current date)"

# GoReleaser targets
goreleaser-check:
	@echo "Checking GoReleaser configuration..."
	goreleaser check

release-snapshot:
	@echo "Creating snapshot release..."
	goreleaser release --snapshot --clean --skip=publish

release:
	@echo "Creating release..."
	goreleaser release --clean
