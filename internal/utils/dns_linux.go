//go:build linux

package utils

import (
	"fmt"
	"os"
	"strings"
)

func (d *DNSManager) setDNSLinux(primary, secondary string) error {
	// Check if using systemd-resolved
	if FileExists("/etc/systemd/resolved.conf") {
		return d.setDNSSystemdResolved(primary, secondary)
	}

	// Fallback to /etc/resolv.conf
	return d.setDNSResolvConf(primary, secondary)
}

func (d *DNSManager) resetDNSLinux() error {
	// Check if using systemd-resolved
	if FileExists("/etc/systemd/resolved.conf") {
		return d.resetDNSSystemdResolved()
	}

	// For resolv.conf, we'll just remove our custom file
	if FileExists("/etc/resolv.conf.splitwire.bak") {
		return os.Rename("/etc/resolv.conf.splitwire.bak", "/etc/resolv.conf")
	}

	return nil
}

func (d *DNSManager) setDNSSystemdResolved(primary, secondary string) error {
	// Read current config
	content, err := os.ReadFile("/etc/systemd/resolved.conf")
	if err != nil {
		return err
	}

	// Backup original
	os.WriteFile("/etc/systemd/resolved.conf.splitwire.bak", content, 0644)

	// Parse and modify
	lines := strings.Split(string(content), "\n")
	var newLines []string
	dnsSet := false

	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "DNS=") {
			if !dnsSet {
				newLines = append(newLines, fmt.Sprintf("DNS=%s %s", primary, secondary))
				dnsSet = true
			}
		} else {
			newLines = append(newLines, line)
		}
	}

	if !dnsSet {
		// Add DNS line under [Resolve] section
		for i, line := range newLines {
			if strings.TrimSpace(line) == "[Resolve]" {
				newLines = append(newLines[:i+1], append([]string{fmt.Sprintf("DNS=%s %s", primary, secondary)}, newLines[i+1:]...)...)
				break
			}
		}
	}

	// Write new config
	err = os.WriteFile("/etc/systemd/resolved.conf", []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		return err
	}

	// Restart service
	_, err = RunCommand("systemctl", "restart", "systemd-resolved")
	return err
}

func (d *DNSManager) resetDNSSystemdResolved() error {
	if FileExists("/etc/systemd/resolved.conf.splitwire.bak") {
		err := os.Rename("/etc/systemd/resolved.conf.splitwire.bak", "/etc/systemd/resolved.conf")
		if err != nil {
			return err
		}
		_, err = RunCommand("systemctl", "restart", "systemd-resolved")
		return err
	}
	return nil
}

func (d *DNSManager) setDNSResolvConf(primary, secondary string) error {
	// Backup original
	if !FileExists("/etc/resolv.conf.splitwire.bak") {
		data, err := os.ReadFile("/etc/resolv.conf")
		if err != nil {
			return err
		}
		os.WriteFile("/etc/resolv.conf.splitwire.bak", data, 0644)
	}

	// Write new resolv.conf
	content := fmt.Sprintf("nameserver %s\nnameserver %s\n", primary, secondary)
	return os.WriteFile("/etc/resolv.conf", []byte(content), 0644)
}
