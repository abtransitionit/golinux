package onpm

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

// -----------------------------------------
// ------ implementation - manage pkg ------
// -----------------------------------------

func (d *DnfPkgManager) List() string {
	cli := "dnf list installed"
	return cli
}

func (d *DnfPkgManager) Add(pkg Pkg2) string {
	cmds := []string{
		fmt.Sprintf("sudo dnf install -q -y %s > /dev/null", pkg.Name),
	}
	// logger.Infof("pkg is: %s", d.Cfg.Pkg)

	return strings.Join(cmds, " && ")
}

func (d *DnfPkgManager) Remove() string {
	cli := "dnf remove <pkg>"
	return cli
}

// -----------------------------------------
// ------ implementation - manage repo ------
// -----------------------------------------

func (d *DnfRepoManager) List() string {
	cli := "dnf list repos"
	return cli
}

func (d *DnfRepoManager) Add(repo Repo2) string {
	cli := "dnf config-manager --add-repo <repo>"
	return cli
}

func (d *DnfRepoManager) Remove() string {
	cli := "dnf config-manager --remove-repo <repo>"
	return cli
}

// -----------------------------------------
// ------ implementation - manage sys ------
// -----------------------------------------

func (d *DnfSysManager) NeedReboot(logger logx.Logger) string {
	cmds := []string{
		"command -v needs-restarting >/dev/null && needs-restarting -r | grep -q 'Reboot is required' && echo true || echo false",
	}
	return strings.Join(cmds, " && ")
}
func (d *DnfSysManager) Upgrade(logger logx.Logger) string {
	cmds := []string{
		"sudo dnf update -q -y",
		"sudo dnf upgrade -q -y",
		"sudo dnf clean all",
	}
	// logger.Infof("pkg is: %s", d.Cfg.Pkg)
	return strings.Join(cmds, " && ")
}

func (d *DnfSysManager) Update(logger logx.Logger) string {
	logger.Infof("pkg is: %s", d.Cfg.Pkg)

	// cmds := []string{
	// 	"sudo dnf update -q -y",
	// 	"sudo dnf upgrade -q -y",
	// 	"sudo dnf clean all",
	// }
	// return strings.Join(cmds, " && ")
	return ""
}
