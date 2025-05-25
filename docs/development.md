# Development Guide

This document covers setting up a development environment for **kicli** and contributing to the project.

---

## Table of Contents

- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Project Structure](#project-structure)
- [Building and Running](#building-and-running)
- [Testing](#testing)
- [Code Style](#code-style)
- [Development Workflow](#development-workflow)
- [Debugging](#debugging)

---

## Prerequisites

- **Go 1.21+** (pure Go dependencies, no CGO required)
- **Git** for version control
- A terminal that supports modern TUI features

---

## Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/semidark/kicli.git
   cd kicli
   ```

2. **Download dependencies:**
   ```bash
   go mod tidy
   ```

3. **Build the project:**
   ```bash
   go build ./cmd/kicli
   ```

4. **Set up sample configuration (optional for development):**
   ```bash
   mkdir -p ~/.config/kicli
   cp assets/config.yaml ~/.config/kicli/config.yaml
   # Edit the config file to add your AI API credentials
   ```

---

## Project Structure

```
kicli/
├── cmd/kicli/           # Main application entry point
├── internal/
│   ├── app/             # Main Bubbletea model and application logic
│   │   ├── model.go     # Core KicliModel struct
│   │   └── messages.go  # Message type definitions
│   ├── tui/
│   │   ├── views/       # UI rendering components
│   │   └── styles/      # Lipgloss styling definitions
│   ├── ptyhandler/      # PTY shell management
│   ├── aiclient/        # AI/LLM client implementation
│   ├── configmanager/   # Configuration loading and management
│   ├── storage/         # SQLite history backend
│   └── util/            # Shared utilities
├── assets/              # Sample configuration files
├── docs/                # Documentation
└── go.mod               # Go module definition
```

For detailed architecture information, see [docs/architecture.md](architecture.md).

---

## Building and Running

### Development Build
```bash
go build ./cmd/kicli
./kicli
```

### Run Without Building
```bash
go run ./cmd/kicli
```

### Build for Production
```bash
go build -ldflags="-s -w" ./cmd/kicli
```

---

## Testing

### Run All Tests
```bash
go test ./...
```

### Run Tests with Coverage
```bash
go test -cover ./...
```

### Run Tests for Specific Package
```bash
go test ./internal/configmanager
```

---

## Code Style

### Formatting
- Use `go fmt` to format all code
- Run `go vet` to catch common issues
- Follow [Effective Go](https://go.dev/doc/effective_go) conventions

### Linting
```bash
# Install golangci-lint if not already installed
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run
```

### Import Organization
- Standard library imports first
- Third-party imports second  
- Local project imports last
- Use `goimports` to organize automatically

---

## Development Workflow

### 1. **Start with Documentation**
- Review [docs/interfaces.md](interfaces.md) for component contracts
- Check [docs/implementation-plan.md](implementation-plan.md) for current priorities

### 2. **Create Feature Branch**
```bash
git checkout -b feature/your-feature-name
```

### 3. **Implement Changes**
- Follow the interface definitions in `docs/interfaces.md`
- Add tests for new functionality
- Update documentation as needed

### 4. **Test Your Changes**
```bash
go test ./...
go build ./cmd/kicli
./kicli  # Manual testing
```

### 5. **Submit Pull Request**
- Ensure all tests pass
- Follow the PR template
- Reference relevant issues

---

## Debugging

### Using Delve Debugger
```bash
# Start debugger in headless mode
dlv debug --headless --api-version=2 --listen=127.0.0.1:43000 ./cmd/kicli

# In another terminal, connect to debugger
dlv connect 127.0.0.1:43000
```

### Logging During Development
Since kicli controls stdout/stderr, use file logging for debugging:

```go
if len(os.Getenv("DEBUG")) > 0 {
    f, err := tea.LogToFile("debug.log", "debug")
    if err != nil {
        fmt.Println("fatal:", err)
        os.Exit(1)
    }
    defer f.Close()
}
```

Then run with debugging enabled:
```bash
DEBUG=1 ./kicli
# In another terminal:
tail -f debug.log
```

---

## Current Development Status

As of the current version, the following components are implemented:

- ✅ **Project structure and build system**
- ✅ **Basic Bubbletea model skeleton**
- ✅ **Message type definitions**
- ✅ **Configuration type definitions**
- ⏳ **Configuration loading** (next priority)
- ⏳ **PTY handler implementation**
- ⏳ **AI client implementation**
- ⏳ **Storage backend**
- ⏳ **TUI layout and styling**

See [docs/implementation-plan.md](implementation-plan.md) for detailed roadmap.

---

## Getting Help

- Check existing [issues](https://github.com/semidark/kicli/issues)
- Review [documentation](../README.md)
- Start a [discussion](https://github.com/semidark/kicli/discussions)

---

**Happy coding!** 🚀 