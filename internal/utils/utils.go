package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// FindExecutable searches for an executable in common locations
func FindExecutable(name string) (string, error) {
	// First try PATH
	if path, err := exec.LookPath(name); err == nil {
		return path, nil
	}

	// Platform-specific search paths
	var searchPaths []string
	
	switch runtime.GOOS {
	case "windows":
		searchPaths = []string{
			filepath.Join(os.Getenv("ProgramFiles"), name),
			filepath.Join(os.Getenv("ProgramFiles(x86)"), name),
			filepath.Join(os.Getenv("LOCALAPPDATA"), name),
		}
	case "linux":
		searchPaths = []string{
			"/usr/bin",
			"/usr/local/bin",
			"/opt/" + name,
			filepath.Join(os.Getenv("HOME"), ".local", "bin"),
		}
	case "darwin":
		searchPaths = []string{
			"/usr/local/bin",
			"/opt/homebrew/bin",
			"/Applications/" + name + ".app/Contents/MacOS",
			filepath.Join(os.Getenv("HOME"), "Applications", name+".app", "Contents", "MacOS"),
		}
	}

	// Search in common paths
	for _, base := range searchPaths {
		path := filepath.Join(base, name)
		if runtime.GOOS == "windows" && filepath.Ext(path) == "" {
			path += ".exe"
		}
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("executable %s not found", name)
}

// RunCommand executes a command and returns output
func RunCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

// RunCommandWithInput executes a command with stdin input
func RunCommandWithInput(input, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()
		stdin.Write([]byte(input))
	}()

	output, err := cmd.CombinedOutput()
	return string(output), err
}

// FileExists checks if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// DirExists checks if a directory exists
func DirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// EnsureDir creates a directory if it doesn't exist
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// CopyFile copies a file from src to dst
func CopyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}

// GetAppDataDir returns the application data directory
func GetAppDataDir() string {
	var appDataDir string
	
	switch runtime.GOOS {
	case "windows":
		appDataDir = filepath.Join(os.Getenv("LOCALAPPDATA"), "SplitWire-Turkey")
	case "darwin":
		home, _ := os.UserHomeDir()
		appDataDir = filepath.Join(home, "Library", "Application Support", "SplitWire-Turkey")
	default: // linux
		home, _ := os.UserHomeDir()
		appDataDir = filepath.Join(home, ".local", "share", "splitwire-turkey")
	}
	
	return appDataDir
}

// GetResourcesDir returns the resources directory
func GetResourcesDir() string {
	// Try to find resources relative to executable
	exe, err := os.Executable()
	if err == nil {
		exePath := filepath.Dir(exe)
		resPath := filepath.Join(exePath, "resources")
		if DirExists(resPath) {
			return resPath
		}
	}
	
	// Fallback to app data directory
	return filepath.Join(GetAppDataDir(), "resources")
}
