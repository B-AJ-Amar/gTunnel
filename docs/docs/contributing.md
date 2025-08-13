---
sidebar_position: 30
---

# Contributing

Learn how to contribute to the gTunnel project.

## Getting Started

We welcome contributions to gTunnel! Whether you're fixing bugs, adding features, or improving documentation, your help is appreciated.

### Ways to Contribute

- Report bugs and issues
- Suggest new features
- Submit code improvements
- Improve documentation
- Help with testing

## Development Setup

### Prerequisites

- Go 1.21 or later
- Git
- Docker (optional, for testing)

### Clone Repository

```bash
git clone https://github.com/B-AJ-Amar/gTunnel.git
cd gTunnel
```

### Build from Source

```bash
# Build client
go build -o bin/gtunnel-client ./cmd/client

# Build server
go build -o bin/gtunnel-server ./cmd/server
```

### Run Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./internal/client/...
```

## Code Guidelines

### Go Style

Follow standard Go conventions:
- Use `gofmt` for formatting
- Follow effective Go guidelines
- Write comprehensive tests
- Add meaningful comments

### Project Structure

```
gTunnel/
├── cmd/                 # Command-line applications
│   ├── client/         # Client CLI
│   └── server/         # Server CLI
├── internal/           # Internal packages
│   ├── client/         # Client implementation
│   ├── server/         # Server implementation
│   └── shared/         # Shared utilities
├── pkg/                # Public packages
├── docs/               # Documentation
└── scripts/            # Build and deployment scripts
```

## Submitting Changes

### Before You Start

1. Check existing issues and PRs
2. Create an issue to discuss major changes
3. Fork the repository
4. Create a feature branch

### Development Workflow

1. **Create branch:**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make changes:**
   - Write code following style guidelines
   - Add tests for new functionality
   - Update documentation as needed

3. **Test changes:**
   ```bash
   go test ./...
   go vet ./...
   golangci-lint run
   ```

4. **Commit changes:**
   ```bash
   git add .
   git commit -m "Add: brief description of changes"
   ```

5. **Push and create PR:**
   ```bash
   git push origin feature/your-feature-name
   ```

### Pull Request Guidelines

- Use clear, descriptive titles
- Include detailed description of changes
- Reference related issues
- Ensure all tests pass
- Update documentation if needed

## Reporting Issues

### Bug Reports

Include the following information:

- gTunnel version
- Operating system
- Go version
- Steps to reproduce
- Expected vs actual behavior
- Error messages/logs
- Configuration files (remove sensitive data)

### Feature Requests

- Clearly describe the proposed feature
- Explain the use case and benefits
- Provide examples if possible
- Consider implementation complexity

## Documentation

### Writing Documentation

- Use clear, concise language
- Include code examples
- Test all examples
- Follow existing documentation style

### Documentation Structure

```
docs/
├── src/
│   ├── pages/          # Landing pages
│   └── components/     # React components
├── docs/               # Main documentation
├── blog/               # Blog posts
└── static/             # Static assets
```

## Testing

### Test Categories

1. **Unit Tests:** Test individual functions/methods
2. **Integration Tests:** Test component interactions
3. **End-to-End Tests:** Test complete workflows

### Writing Tests

```go
func TestClientConnect(t *testing.T) {
    // Arrange
    client := NewClient(Config{...})
    
    // Act
    err := client.Connect()
    
    // Assert
    assert.NoError(t, err)
    assert.True(t, client.IsConnected())
}
```

### Running Tests

```bash
# Unit tests
go test ./internal/...

# Integration tests
go test -tags=integration ./test/...

# All tests
make test
```

## Release Process

### Version Numbers

gTunnel follows semantic versioning (SemVer):
- MAJOR.MINOR.PATCH
- Example: 1.2.3

### Release Checklist

1. Update version numbers
2. Update CHANGELOG.md
3. Run full test suite
4. Build release binaries
5. Create GitHub release
6. Update documentation
7. Announce release

## Community

### Communication Channels

- [GitHub Issues](https://github.com/B-AJ-Amar/gTunnel/issues) - Bug reports and feature requests
- [GitHub Discussions](https://github.com/B-AJ-Amar/gTunnel/discussions) - General discussions
- [Documentation](./intro.md) - Project documentation

### Code of Conduct

Please be respectful and constructive in all interactions. We follow the [Go Community Code of Conduct](https://golang.org/conduct).

## Recognition

Contributors are recognized in:
- CONTRIBUTORS.md file
- Release notes
- Project documentation

Thank you for contributing to gTunnel!
