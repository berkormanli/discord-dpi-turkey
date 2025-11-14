//go:build !windows

package utils

import "fmt"

func (d *DNSManager) setDNSWindows(primary, secondary string) error {
	return fmt.Errorf("Windows DNS management not available on this platform")
}

func (d *DNSManager) resetDNSWindows() error {
	return fmt.Errorf("Windows DNS management not available on this platform")
}

func (d *DNSManager) enableDoHWindows() error {
	return fmt.Errorf("Windows DoH management not available on this platform")
}

func (d *DNSManager) disableDoHWindows() error {
	return fmt.Errorf("Windows DoH management not available on this platform")
}

func (d *DNSManager) getActiveInterfaceWindows() (string, error) {
	return "", fmt.Errorf("Windows interface detection not available on this platform")
}
