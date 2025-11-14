package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

// Config holds application configuration
type Config struct {
	Language      string   `json:"language"`       // TR, EN, RU
	IsDarkMode    bool     `json:"isDarkMode"`
	LastUpdated   string   `json:"lastUpdated"`
	Version       string   `json:"version"`
	CustomFolders []string `json:"customFolders"`  // For custom tunneling
	BrowserTunnel bool     `json:"browserTunnel"`  // Tunnel browsers too
}

// Default returns default configuration
func Default() *Config {
	return &Config{
		Language:      "EN",
		IsDarkMode:    false,
		Version:       "2.0.0",
		CustomFolders: []string{},
		BrowserTunnel: false,
	}
}

// Load reads configuration from disk
func Load() (*Config, error) {
	configPath := getConfigPath()
	
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return Default(), nil
		}
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Save writes configuration to disk
func (c *Config) Save() error {
	configPath := getConfigPath()
	
	// Ensure config directory exists
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// getConfigPath returns platform-specific config file path
func getConfigPath() string {
	var configDir string

	switch runtime.GOOS {
	case "windows":
		configDir = filepath.Join(os.Getenv("LOCALAPPDATA"), "SplitWire-Turkey")
	case "darwin":
		home, _ := os.UserHomeDir()
		configDir = filepath.Join(home, "Library", "Application Support", "SplitWire-Turkey")
	default: // linux and others
		home, _ := os.UserHomeDir()
		configDir = filepath.Join(home, ".config", "splitwire-turkey")
	}

	return filepath.Join(configDir, "config.json")
}

// GetLogsDir returns platform-specific logs directory
func GetLogsDir() string {
	var logsDir string

	switch runtime.GOOS {
	case "windows":
		logsDir = filepath.Join(os.Getenv("LOCALAPPDATA"), "SplitWire-Turkey", "Logs")
	case "darwin":
		home, _ := os.UserHomeDir()
		logsDir = filepath.Join(home, "Library", "Logs", "SplitWire-Turkey")
	default: // linux and others
		home, _ := os.UserHomeDir()
		logsDir = filepath.Join(home, ".local", "share", "splitwire-turkey", "logs")
	}

	return logsDir
}
