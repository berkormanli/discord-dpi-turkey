//go:build linux

package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type linuxServiceManager struct{}

func newLinuxServiceManager() (ServiceManager, error) {
	// Check if systemd is available
	if _, err := exec.LookPath("systemctl"); err != nil {
		return nil, fmt.Errorf("systemd not found: %w", err)
	}
	return &linuxServiceManager{}, nil
}

func (l *linuxServiceManager) Install(name, displayName, description, executable string, args []string) error {
	// Create systemd service file
	serviceContent := fmt.Sprintf(`[Unit]
Description=%s
After=network.target

[Service]
Type=simple
ExecStart=%s %s
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
`, description, executable, strings.Join(args, " "))

	servicePath := filepath.Join("/etc/systemd/system", name+".service")
	if err := os.WriteFile(servicePath, []byte(serviceContent), 0644); err != nil {
		return fmt.Errorf("failed to write service file: %w", err)
	}

	// Reload systemd
	if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
		return fmt.Errorf("failed to reload systemd: %w", err)
	}

	// Enable service
	if err := exec.Command("systemctl", "enable", name).Run(); err != nil {
		return fmt.Errorf("failed to enable service: %w", err)
	}

	return nil
}

func (l *linuxServiceManager) Uninstall(name string) error {
	// Stop service first
	exec.Command("systemctl", "stop", name).Run()

	// Disable service
	if err := exec.Command("systemctl", "disable", name).Run(); err != nil {
		// Continue even if disable fails
	}

	// Remove service file
	servicePath := filepath.Join("/etc/systemd/system", name+".service")
	if err := os.Remove(servicePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove service file: %w", err)
	}

	// Reload systemd
	exec.Command("systemctl", "daemon-reload").Run()

	return nil
}

func (l *linuxServiceManager) Start(name string) error {
	if err := exec.Command("systemctl", "start", name).Run(); err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}
	return nil
}

func (l *linuxServiceManager) Stop(name string) error {
	if err := exec.Command("systemctl", "stop", name).Run(); err != nil {
		return fmt.Errorf("failed to stop service: %w", err)
	}
	return nil
}

func (l *linuxServiceManager) Status(name string) (ServiceStatus, error) {
	out, err := exec.Command("systemctl", "is-active", name).Output()
	if err != nil {
		// Service might not exist or be inactive
		return StatusStopped, nil
	}

	status := strings.TrimSpace(string(out))
	switch status {
	case "active":
		return StatusRunning, nil
	case "activating":
		return StatusStarting, nil
	case "deactivating":
		return StatusStopping, nil
	case "inactive", "failed":
		return StatusStopped, nil
	default:
		return StatusUnknown, nil
	}
}

func (l *linuxServiceManager) List() ([]ServiceInfo, error) {
	// List all splitwire-related services
	out, err := exec.Command("systemctl", "list-units", "--type=service", "--all", "--no-pager").Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list services: %w", err)
	}

	var result []ServiceInfo
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if !strings.Contains(strings.ToLower(line), "splitwire") &&
			!strings.Contains(strings.ToLower(line), "wiresock") &&
			!strings.Contains(strings.ToLower(line), "byedpi") &&
			!strings.Contains(strings.ToLower(line), "zapret") &&
			!strings.Contains(strings.ToLower(line), "goodbyedpi") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		serviceName := strings.TrimSuffix(fields[0], ".service")
		status, _ := l.Status(serviceName)

		result = append(result, ServiceInfo{
			Name:        serviceName,
			DisplayName: serviceName,
			Description: "",
			Status:      status,
		})
	}

	return result, nil
}
