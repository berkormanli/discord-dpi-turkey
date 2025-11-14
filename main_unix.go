//go:build !windows

package main

import (
	"os"
)

func checkAdminPrivileges() bool {
	// On Unix-like systems, check if running as root (UID 0)
	return os.Geteuid() == 0
}
