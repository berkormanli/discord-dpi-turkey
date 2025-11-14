//go:build !darwin

package services

import "fmt"

func newDarwinServiceManager() (ServiceManager, error) {
	return nil, fmt.Errorf("macOS service manager not available on this platform")
}
