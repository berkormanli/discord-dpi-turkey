@echo off
REM Build script for SplitWire-Turkey (Go version)

echo ============================================
echo SplitWire-Turkey Build Script
echo ============================================
echo.

REM Check if Go is installed
go version >nul 2>&1
if errorlevel 1 (
    echo ERROR: Go is not installed or not in PATH
    echo Please install Go from https://golang.org/dl/
    pause
    exit /b 1
)

echo Go version:
go version
echo.

echo Downloading dependencies...
go mod download
if errorlevel 1 (
    echo ERROR: Failed to download dependencies
    pause
    exit /b 1
)

echo.
echo Building SplitWire-Turkey for Windows (64-bit)...
go build -ldflags="-s -w" -o splitwire-turkey.exe
if errorlevel 1 (
    echo ERROR: Build failed
    pause
    exit /b 1
)

echo.
echo ============================================
echo Build successful!
echo Output: splitwire-turkey.exe
echo ============================================
echo.
echo To run the application, right-click splitwire-turkey.exe
echo and select "Run as Administrator"
echo.
pause
