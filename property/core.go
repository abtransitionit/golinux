// file: golinux/property/core.go
package property

import (
	"fmt"
	"runtime"
	"strings"

	coreproperty "github.com/abtransitionit/gocore/property"
	"github.com/abtransitionit/gocore/run"
)

// Name: PropertyHandler
//
// Description: retrieves a system property.
type PropertyHandler func(...string) (string, error)

// GetProperty handles local (cross-platform or Linux) and remote properties.
func GetProperty(vmName, property string, params ...string) (string, error) {
	// Remote property
	if vmName != "" {
		cmd := fmt.Sprintf("goluc property %s %s", property, strings.Join(params, " "))
		// does CLI goluc exist on the remote
		output, err := run.RunCliSsh(vmName, cmd)
		if err != nil {
			return "", fmt.Errorf("failed to get remote property '%s' from '%s': %w", property, vmName, err)
		}
		return strings.TrimSpace(output), nil
	}

	// Local property: try cross-platform first
	if val, err := coreproperty.GetProperty(property, params...); err == nil {
		return val, nil
	}

	// Local property: try linux-platform then
	if runtime.GOOS == "linux" {
		if handler, ok := linuxProperties[property]; ok {
			val, err := handler(params...)
			if err != nil {
				return "", fmt.Errorf("error getting %s: %w", property, err)
			}
			return strings.TrimSpace(val), nil
		}
	}

	// Local property: unknown property
	return "", fmt.Errorf("unknown property requested: %s", property)
}
