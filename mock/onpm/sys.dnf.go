package onpm

import (
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

func (d *DnfSysManager) Clean() (string, error) {
	return "sudo dnf clean all", nil
}

func (d *DnfSysManager) Update() (string, error) {
	return "sudo dnf update -q -y", nil
}

func (d *DnfSysManager) Upgrade(logger logx.Logger) string {
	cmds := []string{
		"sudo dnf update -q -y",
		"sudo dnf upgrade -q -y",
		"sudo dnf clean all",
	}
	logger.Infof("pkg is: %s", d.Cfg.Pkg)
	return strings.Join(cmds, " && ")
}
