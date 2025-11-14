//go:build !linux

package services

import "fmt"

func newLinuxServiceManager() (ServiceManager, error) {
	return nil, fmt.Errorf("Linux service manager not available on this platform")
}
