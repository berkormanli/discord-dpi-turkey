//go:build !windows

package services

import "fmt"

func newWindowsServiceManager() (ServiceManager, error) {
	return nil, fmt.Errorf("Windows service manager not available on this platform")
}
