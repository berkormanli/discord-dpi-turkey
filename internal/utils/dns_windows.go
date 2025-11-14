//go:build windows

package utils

import (
	"fmt"
)

func (d *DNSManager) setDNSWindows(primary, secondary string) error {
	// Get active network interface
	iface, err := d.getActiveInterfaceWindows()
	if err != nil {
		return err
	}

	// Set primary DNS
	_, err = RunCommand("netsh", "interface", "ip", "set", "dns", iface, "static", primary)
	if err != nil {
		return fmt.Errorf("failed to set primary DNS: %w", err)
	}

	// Set secondary DNS
	if secondary != "" {
		_, err = RunCommand("netsh", "interface", "ip", "add", "dns", iface, secondary, "index=2")
		if err != nil {
			return fmt.Errorf("failed to set secondary DNS: %w", err)
		}
	}

	return nil
}

func (d *DNSManager) resetDNSWindows() error {
	iface, err := d.getActiveInterfaceWindows()
	if err != nil {
		return err
	}

	_, err = RunCommand("netsh", "interface", "ip", "set", "dns", iface, "dhcp")
	if err != nil {
		return fmt.Errorf("failed to reset DNS: %w", err)
	}

	return nil
}

func (d *DNSManager) enableDoHWindows() error {
	// Enable DoH for Google DNS (8.8.8.8)
	_, err := RunCommand("netsh", "dns", "add", "encryption", "server=8.8.8.8", "dohtemplate=https://dns.google/dns-query")
	if err != nil {
		return fmt.Errorf("failed to enable DoH: %w", err)
	}

	// Enable DoH for Quad9 (9.9.9.9)
	_, err = RunCommand("netsh", "dns", "add", "encryption", "server=9.9.9.9", "dohtemplate=https://dns.quad9.net/dns-query")
	if err != nil {
		// Continue even if second one fails
	}

	return nil
}

func (d *DNSManager) disableDoHWindows() error {
	_, err := RunCommand("netsh", "dns", "delete", "encryption", "server=8.8.8.8")
	if err != nil {
		// Continue even if it fails
	}

	_, err = RunCommand("netsh", "dns", "delete", "encryption", "server=9.9.9.9")
	if err != nil {
		// Continue even if it fails
	}

	return nil
}

func (d *DNSManager) getActiveInterfaceWindows() (string, error) {
	// Get the active network interface name
	// This is a simplified version - in production, you'd want to detect the actual active interface
	output, err := RunCommand("netsh", "interface", "show", "interface")
	if err != nil {
		return "", err
	}

	// Parse output to find connected interface
	// For now, use a common default
	_ = output

	// Common interface names
	commonNames := []string{
		"Ethernet",
		"Wi-Fi",
		"Local Area Connection",
		"Wireless Network Connection",
	}

	for _, name := range commonNames {
		// Try each name - if it works, use it
		_, err := RunCommand("netsh", "interface", "ip", "show", "config", name)
		if err == nil {
			return name, nil
		}
	}

	return "Ethernet", nil // Default fallback
}
