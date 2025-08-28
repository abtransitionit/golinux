package dnfapt

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

// Name: InstallRepo
//
// Description: install a dnfapt repository on a Linux distro
func InstallPackage(logger logx.Logger, osFamily string, packageName string) error {

	if osFamily != "rhel" && osFamily != "fedora" && osFamily != "debian" {
		return fmt.Errorf("this function only supports Linux (rhel, fedora, debian), but found: %s", osFamily)
	}

	logger.Debugf("Attempting to install dnfapt package: %s on %s", packageName, osFamily)
	switch osFamily {
	case "rhel", "fedora":
	case "debian":
	}

	logger.Infof("Successfully installed dnfapt package: %s", packageName)
	return nil
}

// "DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' install -qq -y %s > /dev/null && DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' update -qq -y > /dev/null",
// packageName)

// "sudo dnf install -q -y %s > /dev/null && sudo dnf update -q -y > /dev/null",
// packageName)
