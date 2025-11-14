package services

import (
	"fmt"
	"runtime"
)

// ServiceManager interface defines operations for managing system services
type ServiceManager interface {
	// Install creates and installs a new service
	Install(name, displayName, description, executable string, args []string) error

	// Uninstall removes a service
	Uninstall(name string) error

	// Start starts a service
	Start(name string) error

	// Stop stops a service
	Stop(name string) error

	// Status returns the status of a service
	Status(name string) (ServiceStatus, error)

	// List returns all managed services
	List() ([]ServiceInfo, error)
}

// ServiceStatus represents the current state of a service
type ServiceStatus int

const (
	StatusUnknown ServiceStatus = iota
	StatusStopped
	StatusRunning
	StatusStarting
	StatusStopping
)

func (s ServiceStatus) String() string {
	switch s {
	case StatusStopped:
		return "Stopped"
	case StatusRunning:
		return "Running"
	case StatusStarting:
		return "Starting"
	case StatusStopping:
		return "Stopping"
	default:
		return "Unknown"
	}
}

// ServiceInfo contains information about a service
type ServiceInfo struct {
	Name        string
	DisplayName string
	Description string
	Status      ServiceStatus
}

// NewServiceManager creates a platform-specific service manager
func NewServiceManager() (ServiceManager, error) {
	switch runtime.GOOS {
	case "windows":
		return newWindowsServiceManager()
	case "linux":
		return newLinuxServiceManager()
	case "darwin":
		return newDarwinServiceManager()
	default:
		return nil, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}
