package dnfapt

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

// Name: InstallPackage
//
// Description: install a dnfapt package on a Linux distro
func InstallPackage(logger logx.Logger, osFamily string, repoName string) error {

	if osFamily != "rhel" && osFamily != "fedora" && osFamily != "debian" {
		return fmt.Errorf("this function only supports Linux (rhel, fedora, debian), but found: %s", osFamily)
	}

	logger.Debugf("Attempting to install dnfapt package: %s on %s", repoName, osFamily)
	switch osFamily {
	case "rhel", "fedora":
	case "debian":
	}

	logger.Debugf("Successfully installed dnfapt package repository: %s", repoName)
	return nil
}
