//go:build darwin

package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type darwinServiceManager struct{}

func newDarwinServiceManager() (ServiceManager, error) {
	// Check if launchctl is available
	if _, err := exec.LookPath("launchctl"); err != nil {
		return nil, fmt.Errorf("launchctl not found: %w", err)
	}
	return &darwinServiceManager{}, nil
}

func (d *darwinServiceManager) Install(name, displayName, description, executable string, args []string) error {
	// Create launchd plist file
	plistContent := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>%s</string>
    <key>ProgramArguments</key>
    <array>
        <string>%s</string>
`, name, executable)

	for _, arg := range args {
		plistContent += fmt.Sprintf("        <string>%s</string>\n", arg)
	}

	plistContent += `    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
    <key>StandardOutPath</key>
    <string>/tmp/` + name + `.log</string>
    <key>StandardErrorPath</key>
    <string>/tmp/` + name + `.err</string>
</dict>
</plist>
`

	plistPath := filepath.Join("/Library/LaunchDaemons", name+".plist")
	if err := os.WriteFile(plistPath, []byte(plistContent), 0644); err != nil {
		return fmt.Errorf("failed to write plist file: %w", err)
	}

	// Load the service
	if err := exec.Command("launchctl", "load", plistPath).Run(); err != nil {
		return fmt.Errorf("failed to load service: %w", err)
	}

	return nil
}

func (d *darwinServiceManager) Uninstall(name string) error {
	plistPath := filepath.Join("/Library/LaunchDaemons", name+".plist")

	// Unload service first
	exec.Command("launchctl", "unload", plistPath).Run()

	// Remove plist file
	if err := os.Remove(plistPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove plist file: %w", err)
	}

	return nil
}

func (d *darwinServiceManager) Start(name string) error {
	if err := exec.Command("launchctl", "start", name).Run(); err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}
	return nil
}

func (d *darwinServiceManager) Stop(name string) error {
	if err := exec.Command("launchctl", "stop", name).Run(); err != nil {
		return fmt.Errorf("failed to stop service: %w", err)
	}
	return nil
}

func (d *darwinServiceManager) Status(name string) (ServiceStatus, error) {
	out, err := exec.Command("launchctl", "list").Output()
	if err != nil {
		return StatusUnknown, fmt.Errorf("failed to list services: %w", err)
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, name) {
			fields := strings.Fields(line)
			if len(fields) >= 3 {
				// If PID is "-", service is not running
				if fields[0] == "-" {
					return StatusStopped, nil
				}
				return StatusRunning, nil
			}
		}
	}

	return StatusStopped, nil
}

func (d *darwinServiceManager) List() ([]ServiceInfo, error) {
	out, err := exec.Command("launchctl", "list").Output()
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
		if len(fields) < 3 {
			continue
		}

		serviceName := fields[2]
		status, _ := d.Status(serviceName)

		result = append(result, ServiceInfo{
			Name:        serviceName,
			DisplayName: serviceName,
			Description: "",
			Status:      status,
		})
	}

	return result, nil
}
