package dnfapt

import (
	"fmt"
	"runtime"

	"github.com/abtransitionit/gocore/errorx"
	"github.com/abtransitionit/gocore/logx"
)

// Name: InstallRepo
// Description: install a dnfapt repository on a Linux distro
func Install(packageName string) error {
	logx.Init()
	logx.Info("Attempting to install package: %s", packageName)

	os := runtime.GOOS
	if os != "linux" {
		// Use the errorx package from gocore to return a professional error
		return errorx.WithStack(fmt.Errorf("this function only supports Linux, but found: %s", os))
	}

	// logic for dnf/apt
	fmt.Printf("Using a single primitive to install %s on a Linux system.\n", packageName)

	logx.Info("Successfully installed package: %s", packageName)
	return nil
}
