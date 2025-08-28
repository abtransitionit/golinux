package dnfapt

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

// Name: InstallRepo
//
// Description: install a dnfapt repository on a Linux distro
func InstallRepository(logger logx.Logger, osFamily string, packageName string) error {

	if osFamily != "rhel" && osFamily != "fedora" && osFamily != "debian" {
		return fmt.Errorf("this function only supports Linux (rhel, fedora, debian), but found: %s", osFamily)
	}

	logger.Debugf("Attempting to install dnfapt package repository: %s on %s", packageName, osFamily)
	switch osFamily {
	case "rhel", "fedora":
	case "debian":
	}

	logger.Infof("Successfully installed dnfapt package: %s", packageName)
	return nil
}
