# SplitWire-Turkey Makefile

# Application name
APP_NAME=splitwire-turkey

# Version
VERSION=2.0.0

# Build flags
LDFLAGS=-ldflags="-s -w -X main.Version=$(VERSION)"

# Default target
.PHONY: all
all: build

# Build for current platform
.PHONY: build
build:
	@echo "Building $(APP_NAME) for current platform..."
	go build $(LDFLAGS) -o $(APP_NAME)

# Build for Windows (amd64)
.PHONY: windows
windows:
	@echo "Building $(APP_NAME) for Windows (amd64)..."
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(APP_NAME).exe

# Build for Windows (386)
.PHONY: windows-32
windows-32:
	@echo "Building $(APP_NAME) for Windows (386)..."
	GOOS=windows GOARCH=386 go build $(LDFLAGS) -o $(APP_NAME)-32.exe

# Build for Linux (amd64)
.PHONY: linux
linux:
	@echo "Building $(APP_NAME) for Linux (amd64)..."
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(APP_NAME)-linux

# Build for Linux (arm64)
.PHONY: linux-arm
linux-arm:
	@echo "Building $(APP_NAME) for Linux (arm64)..."
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(APP_NAME)-linux-arm64

# Build for macOS (amd64)
.PHONY: darwin
darwin:
	@echo "Building $(APP_NAME) for macOS (amd64)..."
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(APP_NAME)-darwin

# Build for macOS (arm64 - Apple Silicon)
.PHONY: darwin-arm
darwin-arm:
	@echo "Building $(APP_NAME) for macOS (arm64)..."
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(APP_NAME)-darwin-arm64

# Build for all platforms
.PHONY: all-platforms
all-platforms: windows windows-32 linux linux-arm darwin darwin-arm
	@echo "Built for all platforms!"

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(APP_NAME) $(APP_NAME).exe $(APP_NAME)-*.exe $(APP_NAME)-linux* $(APP_NAME)-darwin*

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -cover -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run tests with race detector
.PHONY: test-race
test-race:
	@echo "Running tests with race detector..."
	go test -v -race ./...

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run linters
.PHONY: lint
lint:
	@echo "Running linters..."
	go vet ./...
	@if command -v staticcheck > /dev/null; then \
		staticcheck ./...; \
	else \
		echo "staticcheck not installed. Run: go install honnef.co/go/tools/cmd/staticcheck@latest"; \
	fi

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Run the application (requires admin/root)
.PHONY: run
run:
	@echo "Running $(APP_NAME)..."
	@if [ "$$(uname)" = "Linux" ] || [ "$$(uname)" = "Darwin" ]; then \
		sudo ./$(APP_NAME); \
	else \
		./$(APP_NAME).exe; \
	fi

# Help
.PHONY: help
help:
	@echo "SplitWire-Turkey Build System"
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build          Build for current platform"
	@echo "  windows        Build for Windows (64-bit)"
	@echo "  windows-32     Build for Windows (32-bit)"
	@echo "  linux          Build for Linux (64-bit)"
	@echo "  linux-arm      Build for Linux (ARM64)"
	@echo "  darwin         Build for macOS (Intel)"
	@echo "  darwin-arm     Build for macOS (Apple Silicon)"
	@echo "  all-platforms  Build for all platforms"
	@echo "  clean          Remove build artifacts"
	@echo "  test           Run tests"
	@echo "  test-coverage  Run tests with coverage"
	@echo "  test-race      Run tests with race detector"
	@echo "  fmt            Format code"
	@echo "  lint           Run linters"
	@echo "  deps           Install dependencies"
	@echo "  run            Run the application"
	@echo "  help           Show this help message"
