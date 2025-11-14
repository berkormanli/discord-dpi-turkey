//go:build !darwin

package utils

import "fmt"

func (d *DNSManager) setDNSDarwin(primary, secondary string) error {
	return fmt.Errorf("macOS DNS management not available on this platform")
}

func (d *DNSManager) resetDNSDarwin() error {
	return fmt.Errorf("macOS DNS management not available on this platform")
}
