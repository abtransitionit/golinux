package dnfapt

import (
	"fmt"
	"strings"
)

// Name: InstallStdDaPackage
//
// Description: return a CLI to install a single STANDARD dnf or apt package on a Linux OS
//
// Parameters:
//
// - osFamily: name of the OS
// - osDistro: name of the distro
// - packageName: name of the package
//
// Returns:
// - the CLI or an ampty string if nothing is to be installed
// - an error
func InstallStdDaPackage(osFamily string, osDistro string, pkgName string) (string, error) {
	var install bool

	if osFamily != "rhel" && osFamily != "fedora" && osFamily != "debian" {
		return "", fmt.Errorf("this function only supports Linux Family (rhel, fedora, debian), but found: %s", osFamily)
	}

	// logic for installtion
	install = false
	switch pkgName {
	case "uidmap":
		if osFamily == "debian" {
			install = true
		}
	case "dbus-user-session":
		if osDistro == "debian" {
			install = true
		}
	case "gnupg":
		if osDistro == "debian" {
			install = true
		}
	}
	// case "almalinux":
	// 	packageList = "dnf-utils python3-dnf-plugin-versionlock"
	// 	// purpose     = "provision CLI needs-restarting versionlock"
	// case "rocky":
	// 	packageList = "python3-dnf-plugin-versionlock"
	// 	// purpose     = "provision CLI versionlock"  ;;
	// case "fedora":
	// 	packageList = "dnf-utils"
	// 	// purpose     = "provision CLI needs-restarting"  ;;

	// if nothing to install
	if !install {
		return "", nil
	}

	// logic
	var cmds []string
	switch osFamily {
	case "rhel", "fedora":
		cmds = []string{
			fmt.Sprintf("sudo dnf install -q -y %s > /dev/null", pkgName),
			"sudo dnf update -q -y > /dev/null",
		}
	case "debian":
		cmds = []string{
			fmt.Sprintf("DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' install -qq -y %s > /dev/null", pkgName),
			"sudo apt update -qq -y > /dev/null",
		}
	}

	cli := strings.Join(cmds, " && ")
	return cli, nil
}

// Name: InstallSingleDaPackage
//
// Description: return a CLI to install a single dnf or apt package on a Linux OS
//
// Parameters:
//
// - osFamily: name of the OS
// - packageName: name of the package
//
// Returns:
// - the CLI
func InstallDaPackage(osFamily string, daPack DaPack) (string, error) {

	if osFamily != "rhel" && osFamily != "fedora" && osFamily != "debian" {
		return "", fmt.Errorf("this function only supports Linux Family (rhel, fedora, debian), but found: %s", osFamily)
	}
	// define var - get the package canonical name
	packageCName, err := GetPackCName(daPack)
	if err != nil {
		return "", fmt.Errorf("%v : %s", err, packageCName)
	}

	// define var
	// packageVersion := daPack.Version

	// logic
	var cmds []string
	switch osFamily {
	case "debian":
		cmds = []string{
			fmt.Sprintf("DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' install -qq -y %s > /dev/null", packageCName),
		}
	case "rhel", "fedora":
		cmds = []string{
			fmt.Sprintf("sudo dnf install -q -y %s > /dev/null", packageCName),
		}
	}

	cli := strings.Join(cmds, " && ")
	return cli, nil
}

func listRepoPackage(osFamily string, repoName string) (string, error) {

	// dnf --disablerepo="*" --enablerepo="k8s" list available
	// dnf repoquery --repoid=crio

	return "", nil
}
