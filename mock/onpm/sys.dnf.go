package onpm

import (
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

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
