# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0] - 2025-11-14

### Added - Go Version (Cross-Platform)

#### Core Infrastructure
- Go module setup with proper dependencies
- Cross-platform build system with Makefile
- Build scripts for Windows (build.bat) and Unix (build.sh)
- Comprehensive .gitignore for Go projects

#### GUI (Fyne)
- Cross-platform GUI using Fyne framework
- All main tabs: WireSock, ByeDPI, Zapret, GoodbyeDPI, Repair, Advanced, About
- Dark/Light theme support with toggle
- Language selector with real-time switching

#### Internationalization
- Complete Turkish (TR) translations
- Complete English (EN) translations
- Complete Russian (RU) translations
- Runtime language switching without restart

#### Service Management
- Abstract service management interface
- Windows Services implementation (golang.org/x/sys)
- Linux systemd implementation
- macOS launchd implementation
- Platform-specific stubs for clean compilation

#### Configuration
- JSON-based configuration system
- Platform-specific config locations
  - Windows: %LOCALAPPDATA%\SplitWire-Turkey\config.json
  - Linux: ~/.config/splitwire-turkey/config.json
  - macOS: ~/Library/Application Support/SplitWire-Turkey/config.json
- Auto-save on application close

#### Utilities
- Cross-platform DNS management
  - Windows: netsh interface commands
  - Linux: systemd-resolved and resolv.conf
  - macOS: networksetup commands
- DoH (DNS over HTTPS) support for Windows 11
- Process management utilities
  - Process detection
  - Process termination
  - Discord installation detection
- File system utilities
  - Executable search in common locations
  - Directory operations
  - File copying

#### Platform Support
- Admin/root privilege checking
  - Windows: SID-based administrator check
  - Unix: UID-based root check
- Platform-specific code organization with build tags
- Graceful degradation on unsupported platforms

#### Documentation
- Comprehensive README_GO.md with:
  - Installation instructions
  - Build instructions for all platforms
  - Cross-compilation guide
  - Configuration reference
- MIGRATION.md detailing C# to Go migration
- Inline code documentation

### Changed

- Project structure reorganized for Go conventions
- Main README updated with Go version notice
- Build process simplified with Makefile

### Technical Details

- **Language**: Go 1.21+
- **GUI Framework**: Fyne v2.7.0
- **Platforms**: Windows, Linux, macOS
- **Binary Size**: ~30-40 MB (static)
- **Memory Usage**: ~40-60 MB
- **Startup Time**: ~1-2 seconds

### Known Limitations (To Be Implemented)

- WireSock integration (pending)
- ByeDPI integration (pending)
- Zapret integration (pending)
- GoodbyeDPI integration (pending)
- Discord repair functionality (pending)
- Service installation workflows (pending)

## [1.5.4] - 2025 (C# Version)

### Previous Version
This was the last C#/WPF version before the Go migration. See the main README for full C# version history.

---

## Migration Notes

### From C# (1.5.x) to Go (2.0.0)

**Breaking Changes:**
- Configuration location changed from Registry to JSON files
- Service names may differ between versions
- Installer format changed (no more InnoSetup)

**Migration Path:**
1. Uninstall C# version using built-in uninstaller
2. Install Go version
3. Reconfigure services as needed

**Compatibility:**
- Both versions can coexist on Windows
- Configuration is not automatically migrated
- Services need to be reinstalled

---

## Development Roadmap

### Version 2.1.0 (Planned)
- [ ] WireSock/WireGuard integration
- [ ] Complete service installation workflows
- [ ] Discord detection and repair
- [ ] Configuration import from C# version

### Version 2.2.0 (Planned)
- [ ] ByeDPI integration
- [ ] Split tunneling implementation
- [ ] Browser detection and configuration

### Version 2.3.0 (Planned)
- [ ] Zapret integration
- [ ] Automatic parameter detection
- [ ] Preset management

### Version 2.4.0 (Planned)
- [ ] GoodbyeDPI integration
- [ ] Blacklist management
- [ ] System-wide DPI bypass

### Version 2.5.0 (Planned)
- [ ] CLI interface
- [ ] Configuration profiles
- [ ] Auto-update mechanism

### Version 3.0.0 (Future)
- [ ] Plugin system
- [ ] Web-based configuration interface
- [ ] Multi-user support
- [ ] Telemetry and diagnostics (opt-in)

---

## Contributing

We welcome contributions! Please:

1. Check existing issues and pull requests
2. Follow Go best practices
3. Add tests for new features
4. Update documentation
5. Add changelog entries

## Support

- GitHub Issues: [Report a bug](https://github.com/berkormanli/discord-dpi-turkey/issues)
- GitHub Discussions: [Ask questions](https://github.com/berkormanli/discord-dpi-turkey/discussions)

## License

MIT License - see LICENSE file for details

Copyright © 2025 Çağrı Taşkın
