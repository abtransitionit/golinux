package dnfapt

import (
	"fmt"
	"strings"
)

// Name: InstallPackage
//
// Description: define a CLI to install a dnfapt package
//
// Parameters:
//
// - osFamily: name of the OS
// - packageName: name of the package
// - logger: logger
//
// Returns:
// - the CLI
func UpgradeOs(osFamily string) (string, error) {

	if osFamily != "rhel" && osFamily != "fedora" && osFamily != "debian" {
		return "", fmt.Errorf("this function only supports Linux (rhel, fedora, debian), but found: %s", osFamily)
	}

	var cmds []string
	switch osFamily {
	case "debian":
		cmds = []string{
			"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' update -qq -y",
			"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' upgrade -qq -y",
			"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' clean -qq",
		}
	case "rhel", "fedora":
		cmds = []string{
			"sudo dnf update -q -y",
			"sudo dnf upgrade -q -y",
			"sudo dnf clean all",
		}
	}

	cli := strings.Join(cmds, " && ")
	return cli, nil
}
