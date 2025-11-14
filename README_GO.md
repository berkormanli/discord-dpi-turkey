# SplitWire-Turkey (Go Version)

**SplitWire-Turkey** is a DPI bypass and tunneling automation tool designed for internet users in Turkey. This is the Go version that provides cross-platform support for Windows, Linux, and macOS.

## Features

- **Cross-platform support**: Works on Windows, Linux, and macOS
- **Multiple DPI bypass methods**: WireSock, ByeDPI, Zapret, GoodbyeDPI
- **Split tunneling**: Tunnel specific applications (Discord, browsers, etc.)
- **Service management**: Automatic service installation and management
- **Multi-language support**: Turkish, English, Russian
- **Dark/Light theme**: Customizable UI theme

## Requirements

### Windows
- Windows 10/11
- Administrator privileges

### Linux
- X11 or Wayland
- systemd (for service management)
- Root privileges

### macOS
- macOS 10.14 or later
- Root privileges

## Building from Source

### Prerequisites

1. **Go 1.21 or later**: Download from [golang.org](https://golang.org/dl/)
2. **C compiler** (gcc or clang)
3. **Platform-specific dependencies**:

#### Windows
- MinGW-w64 (for CGO)

#### Linux
```bash
# Debian/Ubuntu
sudo apt-get install libgl1-mesa-dev xorg-dev

# Fedora/RHEL
sudo dnf install mesa-libGL-devel libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel

# Arch
sudo pacman -S libgl xorg-server-devel
```

#### macOS
- Xcode Command Line Tools:
```bash
xcode-select --install
```

### Build Instructions

```bash
# Clone the repository
git clone https://github.com/berkormanli/discord-dpi-turkey.git
cd discord-dpi-turkey

# Download dependencies
go mod download

# Build for your platform
go build -o splitwire-turkey

# Or build with optimizations
go build -ldflags="-s -w" -o splitwire-turkey
```

### Cross-Compilation

#### Build for Windows from Linux/macOS
```bash
GOOS=windows GOARCH=amd64 go build -o splitwire-turkey.exe
```

#### Build for Linux from Windows/macOS
```bash
GOOS=linux GOARCH=amd64 go build -o splitwire-turkey
```

#### Build for macOS from Linux/Windows
```bash
GOOS=darwin GOARCH=amd64 go build -o splitwire-turkey
```

## Installation

### Windows
1. Download `splitwire-turkey.exe`
2. Right-click and select "Run as Administrator"

### Linux
```bash
sudo ./splitwire-turkey
```

### macOS
```bash
sudo ./splitwire-turkey
```

## Usage

1. Launch the application with administrator/root privileges
2. Select your preferred language from the top bar
3. Choose a DPI bypass method:
   - **WireSock**: WireGuard-based tunneling
   - **ByeDPI**: DPI bypass with split tunneling
   - **Zapret**: Advanced DPI bypass with custom parameters
   - **GoodbyeDPI**: System-wide DPI bypass
4. Configure and install the service
5. The service will start automatically on system boot

## Configuration

Configuration is stored in platform-specific locations:
- **Windows**: `%LOCALAPPDATA%\SplitWire-Turkey\config.json`
- **Linux**: `~/.config/splitwire-turkey/config.json`
- **macOS**: `~/Library/Application Support/SplitWire-Turkey/config.json`

## Service Management

Services are managed through platform-specific managers:
- **Windows**: Windows Service Manager
- **Linux**: systemd
- **macOS**: launchd

## Logs

Logs are stored in:
- **Windows**: `%LOCALAPPDATA%\SplitWire-Turkey\Logs\`
- **Linux**: `~/.local/share/splitwire-turkey/logs/`
- **macOS**: `~/Library/Logs/SplitWire-Turkey/`

## Development

### Project Structure
```
discord-dpi-turkey/
├── main.go                 # Application entry point
├── main_windows.go         # Windows-specific code
├── main_unix.go            # Unix-specific code
├── internal/
│   ├── config/            # Configuration management
│   ├── gui/               # Fyne GUI implementation
│   ├── i18n/              # Internationalization
│   ├── services/          # Service management
│   └── utils/             # Utility functions
├── go.mod
└── go.sum
```

### Adding New Features

1. Follow Go best practices and conventions
2. Use the existing service abstraction for platform-specific code
3. Add translations to `internal/i18n/translations_*.go`
4. Update UI in `internal/gui/main_ui.go`

### Testing

```bash
# Run tests
go test ./...

# Run with race detector
go test -race ./...

# Run with coverage
go test -cover ./...
```

## Migrating from C# Version

The Go version maintains feature parity with the original C# version while adding cross-platform support. Key differences:

1. **GUI Framework**: Moved from WPF to Fyne for cross-platform support
2. **Service Management**: Abstracted to support Windows Services, systemd, and launchd
3. **Configuration**: JSON-based configuration compatible with original format
4. **Performance**: Generally faster startup and lower memory usage

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

Copyright © 2025 Çağrı Taşkın

This project is licensed under the MIT License - see the LICENSE file for details.

## Credits

- **Original C# Version**: [SplitWire-Turkey](https://github.com/cagritaskn/SplitWire-Turkey)
- **wgcf**: [ViRb3/wgcf](https://github.com/ViRb3/wgcf)
- **ProxiFyre**: [wiresock/proxifyre](https://github.com/wiresock/proxifyre)
- **ByeDPI**: [hufrea/byedpi](https://github.com/hufrea/byedpi)
- **WireSock**: [WireSock VPN](https://www.wiresock.net/)
- **GoodbyeDPI**: [ValdikSS/GoodbyeDPI](https://github.com/ValdikSS/GoodbyeDPI)
- **Zapret**: [bol-van/zapret](https://github.com/bol-van/zapret)
- **WinDivert**: [basil00/WinDivert](https://github.com/basil00/WinDivert)

## Disclaimer

**This software is for educational purposes.**

- This tool is for educational and personal use only
- Not suitable for commercial use
- The developer is not responsible for any damage arising from the use of this software
- Users are responsible for legal compliance in their jurisdiction

## Support

For issues, questions, or contributions:
- GitHub Issues: [Report an issue](https://github.com/berkormanli/discord-dpi-turkey/issues)
- GitHub Discussions: [Join the discussion](https://github.com/berkormanli/discord-dpi-turkey/discussions)
