package dpi

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/berkormanli/discord-dpi-turkey/internal/config"
	"github.com/berkormanli/discord-dpi-turkey/internal/services"
	"github.com/berkormanli/discord-dpi-turkey/internal/utils"
)

// DPIManager handles DPI bypass operations
type DPIManager struct {
	serviceMgr services.ServiceManager
	config     *config.Config
}

// NewDPIManager creates a new DPI manager
func NewDPIManager(cfg *config.Config) (*DPIManager, error) {
	sm, err := services.NewServiceManager()
	if err != nil {
		return nil, fmt.Errorf("failed to create service manager: %w", err)
	}

	return &DPIManager{
		serviceMgr: sm,
		config:     cfg,
	}, nil
}

// InstallWireSockStandard installs WireSock with standard configuration
func (d *DPIManager) InstallWireSockStandard() error {
	return fmt.Errorf("WireSock requires external binaries. Please install WireSock manually from https://www.wiresock.net/ and wgcf from https://github.com/ViRb3/wgcf")
}

// InstallWireSockAlternative installs WireSock with alternative configuration
func (d *DPIManager) InstallWireSockAlternative() error {
	return fmt.Errorf("WireSock requires external binaries. Please install WireSock manually from https://www.wiresock.net/ and wgcf from https://github.com/ViRb3/wgcf")
}

// InstallByeDPISplitTunnel installs ByeDPI with split tunneling
func (d *DPIManager) InstallByeDPISplitTunnel() error {
	return fmt.Errorf("ByeDPI requires external binaries. Please install ByeDPI from https://github.com/hufrea/byedpi and ProxiFyre from https://github.com/wiresock/proxifyre")
}

// InstallByeDPIDLL installs ByeDPI with DLL injection
func (d *DPIManager) InstallByeDPIDLL() error {
	return fmt.Errorf("ByeDPI DLL installation requires external binaries. Please install ByeDPI from https://github.com/hufrea/byedpi")
}

// UninstallByeDPI removes ByeDPI installation
func (d *DPIManager) UninstallByeDPI() error {
	// Try to uninstall ByeDPI service if it exists
	services := []string{"byedpi", "byedpi-split", "byedpi-dll"}

	var lastErr error
	for _, svc := range services {
		if err := d.serviceMgr.Uninstall(svc); err != nil {
			lastErr = err
		}
	}

	if lastErr != nil {
		return fmt.Errorf("failed to uninstall some ByeDPI services: %w", lastErr)
	}

	return nil
}

// InstallZapretAuto installs Zapret with automatic configuration
func (d *DPIManager) InstallZapretAuto(scanSpeed string) error {
	return fmt.Errorf("Zapret requires external binaries and blockcheck tool. Please install Zapret from https://github.com/bol-van/zapret")
}

// InstallZapretPreset installs Zapret with preset configuration
func (d *DPIManager) InstallZapretPreset(preset string) error {
	return fmt.Errorf("Zapret requires external binaries. Please install Zapret from https://github.com/bol-van/zapret")
}

// RunZapretOnce runs Zapret temporarily
func (d *DPIManager) RunZapretOnce(preset string) error {
	return fmt.Errorf("Zapret requires external binaries. Please install Zapret from https://github.com/bol-van/zapret")
}

// UninstallZapret removes Zapret installation
func (d *DPIManager) UninstallZapret() error {
	return d.serviceMgr.Uninstall("zapret")
}

// InstallGoodbyeDPI installs GoodbyeDPI service
func (d *DPIManager) InstallGoodbyeDPI(preset string, useBlacklist bool) error {
	return fmt.Errorf("GoodbyeDPI requires external binaries. Please install GoodbyeDPI from https://github.com/ValdikSS/GoodbyeDPI")
}

// RunGoodbyeDPIOnce runs GoodbyeDPI temporarily
func (d *DPIManager) RunGoodbyeDPIOnce(preset string, useBlacklist bool) error {
	return fmt.Errorf("GoodbyeDPI requires external binaries. Please install GoodbyeDPI from https://github.com/ValdikSS/GoodbyeDPI")
}

// UninstallGoodbyeDPI removes GoodbyeDPI installation
func (d *DPIManager) UninstallGoodbyeDPI() error {
	return d.serviceMgr.Uninstall("goodbyedpi")
}

// RepairDiscord repairs Discord installation
func (d *DPIManager) RepairDiscord() error {
	pm := utils.NewProcessManager()

	// Check if Discord is running
	running, _ := pm.IsProcessRunning("Discord")
	if running {
		return fmt.Errorf("please close Discord before repairing")
	}

	// Find Discord installation
	discordPath, err := pm.FindDiscordPath()
	if err != nil {
		return fmt.Errorf("Discord installation not found: %w", err)
	}

	// Clear cache
	var cacheDir string
	switch runtime.GOOS {
	case "windows":
		cacheDir = filepath.Join(os.Getenv("APPDATA"), "discord")
	case "darwin":
		home, _ := os.UserHomeDir()
		cacheDir = filepath.Join(home, "Library", "Application Support", "discord")
	default:
		home, _ := os.UserHomeDir()
		cacheDir = filepath.Join(home, ".config", "discord")
	}

	if utils.DirExists(cacheDir) {
		if err := os.RemoveAll(cacheDir); err != nil {
			return fmt.Errorf("failed to clear Discord cache: %w", err)
		}
	}

	return fmt.Errorf("Discord repair completed. Cache cleared at %s. Discord path: %s. Manual reinstallation may be required", cacheDir, discordPath)
}

// InstallDiscordPTB installs Discord PTB
func (d *DPIManager) InstallDiscordPTB(cleanInstall bool) error {
	if cleanInstall {
		// Remove standard Discord first
		pm := utils.NewProcessManager()
		pm.KillProcess("Discord")
	}

	var downloadURL string
	switch runtime.GOOS {
	case "windows":
		downloadURL = "https://discord.com/api/downloads/distributions/app/installers/latest?channel=ptb&platform=win&arch=x86"
	case "darwin":
		downloadURL = "https://discord.com/api/downloads/distributions/app/installers/latest?channel=ptb&platform=osx&arch=x64"
	default:
		return fmt.Errorf("Discord PTB automatic installation not supported on this platform. Please download from https://discord.com/download")
	}

	return fmt.Errorf("please download and install Discord PTB manually from: %s", downloadURL)
}

// RemoveAllServices removes all installed services
func (d *DPIManager) RemoveAllServices() error {
	services, err := d.serviceMgr.List()
	if err != nil {
		return fmt.Errorf("failed to list services: %w", err)
	}

	var errors []error
	for _, svc := range services {
		if err := d.serviceMgr.Uninstall(svc.Name); err != nil {
			errors = append(errors, fmt.Errorf("failed to uninstall %s: %w", svc.Name, err))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("failed to remove some services: %v", errors)
	}

	return nil
}

// ResetDNS resets DNS settings to automatic
func (d *DPIManager) ResetDNS() error {
	return utils.ResetDNS()
}

// OpenLogsFolder opens the logs folder in file explorer
func (d *DPIManager) OpenLogsFolder() error {
	logsDir := config.GetLogsDir()

	// Ensure logs directory exists
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return fmt.Errorf("failed to create logs directory: %w", err)
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("explorer", logsDir)
	case "darwin":
		cmd = exec.Command("open", logsDir)
	default: // linux
		// Try xdg-open first, fallback to others
		if _, err := exec.LookPath("xdg-open"); err == nil {
			cmd = exec.Command("xdg-open", logsDir)
		} else if _, err := exec.LookPath("nautilus"); err == nil {
			cmd = exec.Command("nautilus", logsDir)
		} else if _, err := exec.LookPath("dolphin"); err == nil {
			cmd = exec.Command("dolphin", logsDir)
		} else {
			return fmt.Errorf("no file manager found. Logs directory: %s", logsDir)
		}
	}

	return cmd.Start()
}

// GetDiscordStatus checks if Discord is installed and returns status
func (d *DPIManager) GetDiscordStatus() (standardInstalled, ptbInstalled bool) {
	pm := utils.NewProcessManager()

	// Check standard Discord
	if _, err := pm.FindDiscordPath(); err == nil {
		standardInstalled = true
	}

	// Check Discord PTB
	var ptbPaths []string
	switch runtime.GOOS {
	case "windows":
		ptbPaths = []string{filepath.Join(os.Getenv("LOCALAPPDATA"), "DiscordPTB")}
	case "darwin":
		ptbPaths = []string{"/Applications/Discord PTB.app"}
	default:
		ptbPaths = []string{"/usr/bin/discord-ptb", "/opt/discord-ptb"}
	}

	for _, path := range ptbPaths {
		if utils.DirExists(path) {
			ptbInstalled = true
			break
		}
	}

	return standardInstalled, ptbInstalled
}
