.PHONY: run build clean test install dev

# Run the application
run:
	go run main.go

# Build the application
build:
	go build -o dexbro-backend main.go

# Build for Windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o dexbro-backend.exe main.go

# Build for Linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o dexbro-backend main.go

# Build for macOS
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o dexbro-backend main.go

# Install dependencies
install:
	go mod download
	go mod tidy

# Clean build artifacts
clean:
	rm -f dexbro-backend dexbro-backend.exe

# Run tests
test:
	go test -v ./...

# Development mode with auto-reload (requires air)
dev:
	air

# Format code
fmt:
	go fmt ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Show help
help:
	@echo "Available commands:"
	@echo "  make run           - Run the application"
	@echo "  make build         - Build the application"
	@echo "  make build-windows - Build for Windows"
	@echo "  make build-linux   - Build for Linux"
	@echo "  make build-mac     - Build for macOS"
	@echo "  make install       - Install dependencies"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make test          - Run tests"
	@echo "  make dev           - Run with auto-reload"
	@echo "  make fmt           - Format code"
	@echo "  make lint          - Run linter"
