#!/bin/bash

# Build script for SplitWire-Turkey (Go version)

set -e

echo "============================================"
echo "SplitWire-Turkey Build Script"
echo "============================================"
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "ERROR: Go is not installed or not in PATH"
    echo "Please install Go from https://golang.org/dl/"
    exit 1
fi

echo "Go version:"
go version
echo ""

echo "Downloading dependencies..."
go mod download

echo ""
echo "Building SplitWire-Turkey..."
go build -ldflags="-s -w" -o splitwire-turkey

echo ""
echo "============================================"
echo "Build successful!"
echo "Output: splitwire-turkey"
echo "============================================"
echo ""
echo "To run the application, use:"
echo "  sudo ./splitwire-turkey"
echo ""
