//go:build !linux

package utils

import "fmt"

func (d *DNSManager) setDNSLinux(primary, secondary string) error {
	return fmt.Errorf("Linux DNS management not available on this platform")
}

func (d *DNSManager) resetDNSLinux() error {
	return fmt.Errorf("Linux DNS management not available on this platform")
}

func (d *DNSManager) setDNSSystemdResolved(primary, secondary string) error {
	return fmt.Errorf("Linux DNS management not available on this platform")
}

func (d *DNSManager) resetDNSSystemdResolved() error {
	return fmt.Errorf("Linux DNS management not available on this platform")
}

func (d *DNSManager) setDNSResolvConf(primary, secondary string) error {
	return fmt.Errorf("Linux DNS management not available on this platform")
}
