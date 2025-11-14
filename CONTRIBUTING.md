# Contributing to SplitWire-Turkey (Go Version)

Thank you for your interest in contributing to SplitWire-Turkey! This document provides guidelines for contributing to the Go version of the project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [How to Contribute](#how-to-contribute)
- [Code Guidelines](#code-guidelines)
- [Testing](#testing)
- [Documentation](#documentation)
- [Pull Request Process](#pull-request-process)

## Code of Conduct

This project follows a simple code of conduct:

- Be respectful and inclusive
- Be collaborative and constructive
- Focus on what is best for the community
- Show empathy towards others

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/discord-dpi-turkey.git
   cd discord-dpi-turkey
   ```
3. **Add upstream remote**:
   ```bash
   git remote add upstream https://github.com/berkormanli/discord-dpi-turkey.git
   ```

## Development Setup

### Prerequisites

- **Go 1.21 or later**: [Download](https://golang.org/dl/)
- **C compiler** (for CGO):
  - Windows: MinGW-w64
  - Linux: gcc
  - macOS: Xcode Command Line Tools
- **Platform-specific libraries**:
  - Linux: `libgl1-mesa-dev xorg-dev`
  - macOS: Included in Xcode

### Install Dependencies

```bash
go mod download
```

### Build

```bash
# For your platform
make build

# Or manually
go build -o splitwire-turkey
```

### Run

```bash
# Must run with admin/root privileges
sudo ./splitwire-turkey  # Linux/macOS
# Or run as administrator on Windows
```

## How to Contribute

### Reporting Bugs

1. **Check existing issues** to avoid duplicates
2. **Use the issue template** if available
3. **Include**:
   - OS and version
   - Go version
   - Steps to reproduce
   - Expected vs actual behavior
   - Logs from `~/.local/share/splitwire-turkey/logs/` (or equivalent)

### Suggesting Features

1. **Open an issue** with the "enhancement" label
2. **Describe**:
   - The problem it solves
   - Proposed solution
   - Alternatives considered
   - Implementation notes (optional)

### Code Contributions

1. **Pick an issue** or create one for discussion
2. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. **Make your changes** following the guidelines below
4. **Test your changes**
5. **Commit with clear messages**
6. **Push to your fork**
7. **Open a Pull Request**

## Code Guidelines

### Go Style

Follow standard Go conventions:

```bash
# Format code
go fmt ./...

# Run linters
go vet ./...

# Use staticcheck if available
staticcheck ./...
```

### Project Structure

- **`main.go`**: Application entry point only
- **`internal/`**: Internal packages (not importable by other projects)
  - `config/`: Configuration management
  - `gui/`: GUI components
  - `i18n/`: Translations
  - `services/`: Service management
  - `utils/`: Utility functions
- **Platform-specific files**: Use build tags

### Build Tags

Use build tags for platform-specific code:

```go
//go:build windows

package mypackage

// Windows-specific code
```

```go
//go:build !windows

package mypackage

// Non-Windows stub
```

### Error Handling

Always handle errors explicitly:

```go
// Good
result, err := someFunction()
if err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}

// Bad
result, _ := someFunction()  // Never ignore errors
```

### Comments

- **Exported functions** must have comments
- **Complex logic** should be explained
- **Platform-specific behavior** should be documented

```go
// SetDNS configures DNS servers for the primary network interface.
// On Windows, this uses netsh commands.
// On Linux, this modifies systemd-resolved or resolv.conf.
// On macOS, this uses networksetup.
func SetDNS(primary, secondary string) error {
    // Implementation
}
```

### Naming Conventions

- **Packages**: lowercase, single word
- **Files**: lowercase, underscores for separation
- **Types**: PascalCase
- **Functions**: camelCase (exported: PascalCase)
- **Constants**: PascalCase or UPPER_CASE for exported

### Internationalization

When adding UI strings:

1. Add to **`internal/i18n/translations_en.go`**
2. Add translations to **`translations_tr.go`** and **`translations_ru.go`**
3. Use the key in code: **`i18n.T("my_key")`**

Example:

```go
// translations_en.go
"my_new_button": "Click Me",

// translations_tr.go
"my_new_button": "Tıkla",

// translations_ru.go
"my_new_button": "Нажми",

// In code
button := widget.NewButton(i18n.T("my_new_button"), callback)
```

## Testing

### Running Tests

```bash
# All tests
go test ./...

# With coverage
go test -cover ./...

# With race detector
go test -race ./...
```

### Writing Tests

- **File naming**: `*_test.go`
- **Test functions**: `func TestXxx(t *testing.T)`
- **Table-driven tests** for multiple cases

Example:

```go
func TestSomething(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {"case1", "input1", "output1", false},
        {"case2", "input2", "output2", false},
        {"error case", "bad", "", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := SomeFunction(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("SomeFunction() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("SomeFunction() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Documentation

### Code Documentation

- Update **README_GO.md** for user-facing changes
- Update **MIGRATION.md** for migration-related changes
- Update **CHANGELOG.md** for all changes
- Add inline comments for complex code

### Commit Messages

Use clear, descriptive commit messages:

```
Add DNS management for Linux

- Implement systemd-resolved support
- Add fallback to resolv.conf
- Add error handling for missing permissions
```

Format:
```
<type>: <subject>

<body>

<footer>
```

Types: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`

## Pull Request Process

1. **Update documentation** if needed
2. **Add tests** for new functionality
3. **Update CHANGELOG.md**
4. **Ensure all tests pass**:
   ```bash
   go test ./...
   go vet ./...
   go fmt ./...
   ```
5. **Create PR** with clear description:
   - What changed?
   - Why was it changed?
   - How was it tested?
   - Related issues?

### PR Review Process

- At least one approval required
- CI checks must pass
- Code must follow guidelines
- Documentation must be updated

### After Approval

- **Squash commits** if requested
- **Update branch** if needed
- Maintainer will merge

## Platform-Specific Contributions

### Windows
- Test on Windows 10 and 11
- Consider both admin and non-admin scenarios
- Test with different antivirus software

### Linux
- Test on Ubuntu/Debian and Fedora/RHEL
- Consider both X11 and Wayland
- Test with and without systemd

### macOS
- Test on Intel and Apple Silicon
- Consider both admin and non-admin scenarios
- Test on recent macOS versions

## Questions?

- Open a **GitHub Discussion** for questions
- Check **existing issues** and documentation
- Ask in PR comments for review-specific questions

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to SplitWire-Turkey! 🎉
