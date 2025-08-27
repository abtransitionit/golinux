// file: golinux/properties/core.go
package property

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/abtransitionit/gocore/properties"
	"github.com/abtransitionit/gocore/run"
)

// Name: GetPropertyLocal
//
// Description: retrieves a property from the Linux-specific set.
func GetPropertyLocal(property string, params ...string) (string, error) {
	// 1️⃣ Try cross-platform properties first
	if val, err := properties.GetPropertyLocal(property, params...); err == nil {
		return val, nil
	}

	// 2️⃣ If Linux, try Linux-specific properties
	if runtime.GOOS == "linux" {
		if val, err := GetPropertyLinuxLocal(property, params...); err == nil {
			return val, nil
		}
	}

	// 3️⃣ Unknown property
	return "", fmt.Errorf("unknown property requested: %s", property)
}

func GetPropertyLinuxLocal(property string, params ...string) (string, error) {
	fnPropertyHandler, ok := linuxProperties[property]
	if !ok {
		return "", fmt.Errorf("❌ unknown property requested: %s", property)
	}

	output, err := fnPropertyHandler(params...)
	if err != nil {
		return "", fmt.Errorf("❌ error getting %s: %w", property, err)
	}
	return strings.TrimSpace(output), nil
}

func GetPropertyRemote(vmName string, property string) (string, error) {
	// Build the CLI command to run remotely
	command := fmt.Sprintf("goluc prop %s --vm %s", property, vmName)

	// Use existing RunCliSsh function
	output, err := run.RunOnVm(vmName, command)
	if err != nil {
		return "", fmt.Errorf("failed to get remote property '%s' from '%s': %w", property, vmName, err)
	}

	return strings.TrimSpace(output), nil
}
