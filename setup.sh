#!/bin/bash

echo "🚀 DexBro Workshop Backend Setup"
echo "=================================="

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or higher."
    exit 1
fi

echo "✅ Go version: $(go version)"

# Check if .env exists
if [ ! -f .env ]; then
    echo "📝 Creating .env file from .env.example..."
    cp .env.example .env
    echo "✅ .env file created. Please update it with your MongoDB credentials."
else
    echo "✅ .env file already exists"
fi

# Install dependencies
echo "📦 Installing Go dependencies..."
go mod download
go mod tidy

echo ""
echo "✅ Setup complete!"
echo ""
echo "Next steps:"
echo "1. Update .env file with your MongoDB connection details"
echo "2. Make sure MongoDB is running"
echo "3. Run the server with: go run main.go"
echo "   Or use: make run"
echo ""
echo "Available commands:"
echo "  make run         - Run the server"
echo "  make build       - Build the binary"
echo "  make install     - Install dependencies"
echo "  make test        - Run tests"
echo ""
