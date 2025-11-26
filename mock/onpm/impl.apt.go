package onpm

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

// -----------------------------------------
// ------ implementation - manage pkg ------
// -----------------------------------------

func (a *AptPkgManager) List() string {
	cli := "apt list --installed"
	return cli
}

func (a *AptPkgManager) Add(pkg Pkg2) string {
	cmds := []string{
		fmt.Sprintf("DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' install -qq -y %s > /dev/null", pkg.Name),
	}
	// logger.Infof("pkg is: %s", d.Cfg.Pkg)

	return strings.Join(cmds, " && ")
}

func (a *AptPkgManager) Remove() string {
	cli := "apt remove <pkg>"
	return cli
}

// -----------------------------------------
// ------ implementation - manage repo ------
// -----------------------------------------

func (a *AptRepoManager) List() string {
	cli := "apt list repos"
	return cli
}

func (a *AptRepoManager) Add(repo Repo2) string {
	cli := "add-apt-repo <repo>"
	return cli
}

func (a *AptRepoManager) Remove() string {
	cli := "remove-apt-repo <repo>"
	return cli
}

// -----------------------------------------
// ------ implementation - manage sys ------
// -----------------------------------------

func (d *AptSysManager) NeedReboot(logger logx.Logger) string {
	cmds := []string{
		"test -f /var/run/reboot-required && echo true || echo false",
	}
	// logger.Infof("pkg is: %s", d.Cfg.Pkg)

	return strings.Join(cmds, " && ")
}

func (d *AptSysManager) Upgrade(logger logx.Logger) string {
	cmds := []string{
		"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' update -qq -y",
		"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' upgrade -qq -y",
		"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' clean -qq",
	}
	// logger.Infof("pkg is: %s", d.Cfg.Pkg)

	return strings.Join(cmds, " && ")
}
func (d *AptSysManager) Update(logger logx.Logger) string {
	logger.Infof("pkg is: %s", d.Cfg.Pkg)
	// cmds := []string{
	// 	"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' update -qq -y",
	// 	"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' upgrade -qq -y",
	// 	"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' clean -qq",
	// }
	// // logger.Infof("pkg is: %s", d.Cfg.Pkg)

	// return strings.Join(cmds, " && ")
	return ""
}
