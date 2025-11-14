//go:build darwin

package utils

import (
	"fmt"
	"strings"
)

func (d *DNSManager) setDNSDarwin(primary, secondary string) error {
	// Get list of network services
	output, err := RunCommand("networksetup", "-listallnetworkservices")
	if err != nil {
		return err
	}

	// Parse services (skip first line which is a header)
	lines := strings.Split(output, "\n")
	var services []string
	for i, line := range lines {
		if i == 0 || strings.TrimSpace(line) == "" {
			continue
		}
		// Skip disabled services (marked with *)
		if !strings.HasPrefix(line, "*") {
			services = append(services, strings.TrimSpace(line))
		}
	}

	// Set DNS for each active service
	for _, service := range services {
		_, err := RunCommand("networksetup", "-setdnsservers", service, primary, secondary)
		if err != nil {
			// Continue even if one service fails
			continue
		}
	}

	return nil
}

func (d *DNSManager) resetDNSDarwin() error {
	// Get list of network services
	output, err := RunCommand("networksetup", "-listallnetworkservices")
	if err != nil {
		return err
	}

	// Parse services
	lines := strings.Split(output, "\n")
	var services []string
	for i, line := range lines {
		if i == 0 || strings.TrimSpace(line) == "" {
			continue
		}
		if !strings.HasPrefix(line, "*") {
			services = append(services, strings.TrimSpace(line))
		}
	}

	// Reset DNS to DHCP for each service
	for _, service := range services {
		_, err := RunCommand("networksetup", "-setdnsservers", service, "empty")
		if err != nil {
			continue
		}
	}

	return nil
}
