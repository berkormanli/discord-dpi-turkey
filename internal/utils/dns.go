package utils

import (
	"fmt"
	"runtime"
)

// DNSManager handles DNS and DoH configuration
type DNSManager struct{}

// NewDNSManager creates a new DNS manager
func NewDNSManager() *DNSManager {
	return &DNSManager{}
}

// SetDNS sets DNS servers for the primary network interface
func (d *DNSManager) SetDNS(primary, secondary string) error {
	switch runtime.GOOS {
	case "windows":
		return d.setDNSWindows(primary, secondary)
	case "linux":
		return d.setDNSLinux(primary, secondary)
	case "darwin":
		return d.setDNSDarwin(primary, secondary)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

// ResetDNS resets DNS to automatic (DHCP)
func (d *DNSManager) ResetDNS() error {
	switch runtime.GOOS {
	case "windows":
		return d.resetDNSWindows()
	case "linux":
		return d.resetDNSLinux()
	case "darwin":
		return d.resetDNSDarwin()
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

// EnableDoH enables DNS over HTTPS (Windows 11 only)
func (d *DNSManager) EnableDoH() error {
	if runtime.GOOS != "windows" {
		return fmt.Errorf("DoH configuration is only supported on Windows 11")
	}
	return d.enableDoHWindows()
}

// DisableDoH disables DNS over HTTPS
func (d *DNSManager) DisableDoH() error {
	if runtime.GOOS != "windows" {
		return fmt.Errorf("DoH configuration is only supported on Windows 11")
	}
	return d.disableDoHWindows()
}
