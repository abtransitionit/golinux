package distro

import (
	"fmt"
	"runtime"
)

// Name: CheckOsIsLinux
// Description: ensures the current operating system is Linux.
func CheckOsIsLinux() error {
	if runtime.GOOS != "linux" {
		// return errorx.WithStack(fmt.Errorf("this library only supports Linux, but found: %s", runtime.GOOS))
		return fmt.Errorf("this library only supports Linux, but found: %s", runtime.GOOS)
	}
	return nil
}

// Name: GetPackageManager
// Return:
// string: the name of the package manager for the current distro.
func GetPackageManager() string {
	// In a real implementation, you would check for /etc/os-release etc.
	// We'll just return a placeholder for this example.
	return "systemd-manager"
}
