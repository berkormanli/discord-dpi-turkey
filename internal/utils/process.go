package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

// ProcessManager handles process operations
type ProcessManager struct{}

// NewProcessManager creates a new process manager
func NewProcessManager() *ProcessManager {
	return &ProcessManager{}
}

// IsProcessRunning checks if a process with the given name is running
func (p *ProcessManager) IsProcessRunning(name string) (bool, error) {
	switch runtime.GOOS {
	case "windows":
		return p.isProcessRunningWindows(name)
	case "linux", "darwin":
		return p.isProcessRunningUnix(name)
	default:
		return false, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

// KillProcess terminates a process by name
func (p *ProcessManager) KillProcess(name string) error {
	switch runtime.GOOS {
	case "windows":
		return p.killProcessWindows(name)
	case "linux", "darwin":
		return p.killProcessUnix(name)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

// FindDiscordPath finds the Discord installation path
func (p *ProcessManager) FindDiscordPath() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return p.findDiscordPathWindows()
	case "linux":
		return p.findDiscordPathLinux()
	case "darwin":
		return p.findDiscordPathDarwin()
	default:
		return "", fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

func (p *ProcessManager) isProcessRunningWindows(name string) (bool, error) {
	if !strings.HasSuffix(name, ".exe") {
		name += ".exe"
	}

	cmd := exec.Command("tasklist", "/FI", fmt.Sprintf("IMAGENAME eq %s", name))
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	return strings.Contains(string(output), name), nil
}

func (p *ProcessManager) isProcessRunningUnix(name string) (bool, error) {
	cmd := exec.Command("pgrep", "-x", name)
	err := cmd.Run()
	return err == nil, nil
}

func (p *ProcessManager) killProcessWindows(name string) error {
	if !strings.HasSuffix(name, ".exe") {
		name += ".exe"
	}

	cmd := exec.Command("taskkill", "/F", "/IM", name)
	return cmd.Run()
}

func (p *ProcessManager) killProcessUnix(name string) error {
	cmd := exec.Command("pkill", "-9", name)
	return cmd.Run()
}

func (p *ProcessManager) findDiscordPathWindows() (string, error) {
	// Common Discord installation paths on Windows
	paths := []string{
		`%LOCALAPPDATA%\Discord\app-*\Discord.exe`,
		`%LOCALAPPDATA%\DiscordPTB\app-*\DiscordPTB.exe`,
		`%LOCALAPPDATA%\DiscordCanary\app-*\DiscordCanary.exe`,
	}

	for _, path := range paths {
		expanded := os.ExpandEnv(path)
		// This would need globbing support in a real implementation
		if FileExists(expanded) {
			return expanded, nil
		}
	}

	return "", fmt.Errorf("Discord not found")
}

func (p *ProcessManager) findDiscordPathLinux() (string, error) {
	// Common Discord installation paths on Linux
	paths := []string{
		"/usr/bin/discord",
		"/usr/local/bin/discord",
		"/opt/discord/discord",
		"/snap/bin/discord",
	}

	for _, path := range paths {
		if FileExists(path) {
			return path, nil
		}
	}

	return "", fmt.Errorf("Discord not found")
}

func (p *ProcessManager) findDiscordPathDarwin() (string, error) {
	// Common Discord installation paths on macOS
	paths := []string{
		"/Applications/Discord.app/Contents/MacOS/Discord",
		"/Applications/Discord PTB.app/Contents/MacOS/Discord PTB",
		"/Applications/Discord Canary.app/Contents/MacOS/Discord Canary",
	}

	for _, path := range paths {
		if FileExists(path) {
			return path, nil
		}
	}

	return "", fmt.Errorf("Discord not found")
}

// StartProcess starts a process with the given command and arguments
func (p *ProcessManager) StartProcess(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	// Detach from parent process
	if runtime.GOOS != "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setpgid: true,
		}
	}

	return cmd.Start()
}
