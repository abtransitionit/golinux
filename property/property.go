// file: golinux/property/property.go
package property

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/abtransitionit/gocore/errorx"
	"github.com/abtransitionit/gocore/run"
	"github.com/opencontainers/selinux/go-selinux"
	"github.com/shirou/gopsutil/v3/host"
)

var linuxProperties = map[string]PropertyHandler{
	"uuid":          getUuid,   // code change from original
	"uname":         getUnameM, // code change from original
	"osdistro":      getOsDistro,
	"osfamily":      getOsFamily,
	"pathtree":      getPathTree,
	"rcfilepath":    getRcFilePath,
	"selinuxStatus": getSelinuxStatus,
	"selinuxMode":   getSelinuxMode,
}

func getSelinuxMode(_ ...string) (string, error) {
	switch selinux.EnforceMode() {
	case selinux.Enforcing:
		return "enforcing", nil
	case selinux.Permissive:
		return "permissive", nil
	case selinux.Disabled:
		return "disabled", nil
	default:
		return "unknown", nil
	}
}

func getSelinuxStatus(_ ...string) (string, error) {
	if selinux.GetEnabled() {
		return "enabled", nil
	}
	return "disabled", nil
}

func getRcFilePath(params ...string) (string, error) {
	return "$HOME/.profile", nil
}
func getPathTree(params ...string) (string, error) {
	if len(params) < 1 {
		return "", fmt.Errorf("base path name required")
	}

	// get input
	basePath := params[0]

	// play code

	cli := fmt.Sprintf(`find %s -type d | sort | paste -sd\:`, basePath)
	path, err := run.RunOnLocal(cli)
	if err != nil {
		return "", err
	}
	return path, nil
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
