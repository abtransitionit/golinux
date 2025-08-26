// file: golinux/properties/properties.go
package property

import (
	"os/exec"
	"strings"

	"github.com/abtransitionit/gocore/errorx"
	"github.com/shirou/gopsutil/v3/host"
)

// PropertyHandler is a function that retrieves a system property.
type PropertyHandler func(...string) (string, error)

var linuxProperties = map[string]PropertyHandler{
	"uuid":     getUuid,   // code change from original
	"uname":    getUnameM, // code change from original
	"osdistro": getOsDistro,
	"osfamily": getOsFamily,
}

// // Cross-platform property
// val, err := gocore.GetPropertyLocal("ostype")

// // Linux-only property (called only on Linux)
// if runtime.GOOS == "linux" {
//     val, err := golinux.GetPropertyLocal("uuid")
// }

// GetLinuxPropertyMap exposes the map of Linux-specific properties to external callers.
func GetLinuxPropertyMap() map[string]PropertyHandler {
	return linuxProperties
}

// getUuid retrieves the system's UUID.
func getUuid(_ ...string) (string, error) {
	// This command is specific to Linux systems.
	cmd := "sudo cat /sys/class/dmi/id/product_uuid"
	output, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", errorx.Wrap(err, "getUuid failed")
	}
	return strings.TrimSpace(string(output)), nil
}

// getUnameM retrieves the machine architecture from `uname -m`.
func getUnameM(_ ...string) (string, error) {
	// `uname -m` is a standard Unix command.
	output, err := exec.Command("uname", "-m").CombinedOutput()
	if err != nil {
		return "", errorx.Wrap(err, "uname failed")
	}
	return strings.TrimSpace(string(output)), nil
}

// getOsDistro retrieves the OS distribution.
func getOsDistro(_ ...string) (string, error) {
	// `gopsutil` retrieves this, but the concept is primarily Linux-specific.
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return info.Platform, nil
}

// getOsFamily retrieves the OS family.
func getOsFamily(_ ...string) (string, error) {
	// `gopsutil` retrieves this, but the concept is primarily Linux-specific.
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return info.PlatformFamily, nil
}
