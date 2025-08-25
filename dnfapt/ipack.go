package dnfapt

import (
	"fmt"
	"runtime"

	"github.com/abtransitionit/gocore/logx"
)

// Name: InstallPackage
// Description: install a dnfapt package on a Linux distro
func InstallPackage(packageName string) error {
	logx.Init()
	logx.Infof("Attempting to install d dnfapt package: %s", packageName)

	os := runtime.GOOS
	if os != "linux" {
		// Use the errorx package from gocore to return a professional error
		// return errorx.WithStack(fmt.Errorf("this function only supports Linux, but found: %s", os))
		// fmt.Sprintf("this function only supports Linux, but found: %s", os)
		return fmt.Errorf("this function only supports Linux, but found: %s", os)
	}

	// logic
	fmt.Printf("Using a single primitive to install %s on a Linux system.\n", packageName)

	// success
	logx.Infof("Successfully installed dnfapt package: %s", packageName)
	return nil
}
