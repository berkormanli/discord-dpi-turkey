//go:build windows

package services

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

type windowsServiceManager struct {
	mgr *mgr.Mgr
}

func newWindowsServiceManager() (ServiceManager, error) {
	m, err := mgr.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to service manager: %w", err)
	}

	return &windowsServiceManager{mgr: m}, nil
}

func (w *windowsServiceManager) Install(name, displayName, description, executable string, args []string) error {
	s, err := w.mgr.OpenService(name)
	if err == nil {
		s.Close()
		return fmt.Errorf("service %s already exists", name)
	}

	config := mgr.Config{
		DisplayName: displayName,
		Description: description,
		StartType:   mgr.StartAutomatic,
	}

	s, err = w.mgr.CreateService(name, executable, config, args...)
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}
	defer s.Close()

	return nil
}

func (w *windowsServiceManager) Uninstall(name string) error {
	s, err := w.mgr.OpenService(name)
	if err != nil {
		return fmt.Errorf("failed to open service: %w", err)
	}
	defer s.Close()

	// Try to stop the service first
	status, err := s.Query()
	if err == nil && status.State != svc.Stopped {
		_, err = s.Control(svc.Stop)
		if err != nil {
			// Continue even if stop fails
		}
		// Wait for service to stop
		time.Sleep(2 * time.Second)
	}

	err = s.Delete()
	if err != nil {
		return fmt.Errorf("failed to delete service: %w", err)
	}

	return nil
}

func (w *windowsServiceManager) Start(name string) error {
	s, err := w.mgr.OpenService(name)
	if err != nil {
		return fmt.Errorf("failed to open service: %w", err)
	}
	defer s.Close()

	err = s.Start()
	if err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}

	return nil
}

func (w *windowsServiceManager) Stop(name string) error {
	s, err := w.mgr.OpenService(name)
	if err != nil {
		return fmt.Errorf("failed to open service: %w", err)
	}
	defer s.Close()

	status, err := s.Control(svc.Stop)
	if err != nil {
		return fmt.Errorf("failed to stop service: %w", err)
	}

	// Wait for service to stop
	timeout := time.Now().Add(10 * time.Second)
	for status.State != svc.Stopped {
		if time.Now().After(timeout) {
			return fmt.Errorf("timeout waiting for service to stop")
		}
		time.Sleep(300 * time.Millisecond)
		status, err = s.Query()
		if err != nil {
			return fmt.Errorf("failed to query service status: %w", err)
		}
	}

	return nil
}

func (w *windowsServiceManager) Status(name string) (ServiceStatus, error) {
	s, err := w.mgr.OpenService(name)
	if err != nil {
		return StatusUnknown, fmt.Errorf("failed to open service: %w", err)
	}
	defer s.Close()

	status, err := s.Query()
	if err != nil {
		return StatusUnknown, fmt.Errorf("failed to query service: %w", err)
	}

	switch status.State {
	case svc.Stopped:
		return StatusStopped, nil
	case svc.Running:
		return StatusRunning, nil
	case svc.StartPending:
		return StatusStarting, nil
	case svc.StopPending:
		return StatusStopping, nil
	default:
		return StatusUnknown, nil
	}
}

func (w *windowsServiceManager) List() ([]ServiceInfo, error) {
	services, err := w.mgr.ListServices()
	if err != nil {
		return nil, fmt.Errorf("failed to list services: %w", err)
	}

	var result []ServiceInfo
	for _, serviceName := range services {
		s, err := w.mgr.OpenService(serviceName)
		if err != nil {
			continue
		}

		config, err := s.Config()
		status, statusErr := s.Query()
		s.Close()

		if err != nil {
			continue
		}

		var serviceStatus ServiceStatus
		if statusErr == nil {
			switch status.State {
			case svc.Stopped:
				serviceStatus = StatusStopped
			case svc.Running:
				serviceStatus = StatusRunning
			case svc.StartPending:
				serviceStatus = StatusStarting
			case svc.StopPending:
				serviceStatus = StatusStopping
			default:
				serviceStatus = StatusUnknown
			}
		}

		result = append(result, ServiceInfo{
			Name:        serviceName,
			DisplayName: config.DisplayName,
			Description: config.Description,
			Status:      serviceStatus,
		})
	}

	return result, nil
}
