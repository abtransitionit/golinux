// file golinux/run/run.go
package run

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetVmOsFamily detects the Linux OS family of a remote VM
func GetVmOsFamily(vmName string) (string, error) {
	// Get ID first
	cmdID := exec.Command("ssh", vmName, "grep ^ID= /etc/os-release | cut -d= -f2 | tr -d '\"'")
	outputID, err := cmdID.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get ID for VM %s: %v", vmName, err)
	}

	osID := strings.TrimSpace(string(outputID))
	switch osID {
	case "debian", "ubuntu":
		return "debian", nil
	case "fedora":
		return "fedora", nil
	case "rhel", "centos", "almalinux", "rocky":
		return "rhel", nil
	}

	// Fallback: check ID_LIKE if it exists
	cmdLike := exec.Command("ssh", vmName, "grep ^ID_LIKE= /etc/os-release | cut -d= -f2 | tr -d '\"'")
	outputLike, err := cmdLike.CombinedOutput()
	if err == nil {
		idLike := strings.Fields(strings.TrimSpace(string(outputLike)))
		if len(idLike) > 0 {
			switch idLike[0] {
			case "debian", "ubuntu":
				return "debian", nil
			case "rhel", "centos", "almalinux", "rocky":
				return "rhel", nil
			case "fedora":
				return "fedora", nil
			}
		}
	}

	return "", fmt.Errorf("unsupported Linux OS Family for VM %s: %s", vmName, osID)
}
