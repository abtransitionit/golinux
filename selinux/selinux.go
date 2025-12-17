package selinux

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/property"
)

func ConfigureSelinux(hostName string, logger logx.Logger) (string, error) {

	// 1 - get host:property
	osFamily, err := property.GetProperty(logger, hostName, "osFamily")
	if err != nil {
		return "", fmt.Errorf("error 01: %v ", err)
	}

	// 2 - skip if Os:family not rhel or not fedora
	if osFamily != "rhel" && osFamily != "fedora" {
		// logger.Debugf("%s:%s 🅐 Skipping selinux configuration", vmName, osFamily)
		return "", nil
	}

	// 3 - here: osFamily == "rhel" || "fedora" => DoTheJob

	// call ConfigureSelinuxAtRuntime
	// cliRuntime := ConfigureSelinuxAtRuntime()

	// call ConfigureSelinuxAtStartup
	// cliStartup := ConfigureSelinuxAtStartup()
	// log
	logger.Infof("%s:%s > ConfigureSelinux called", hostName, osFamily)
	// handle success
	return osFamily, nil
}

func ConfigureSelinuxAtRuntime() string {
	var cmds = []string{
		"sudo setenforce 0",
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

func ConfigureSelinuxAtStartup() string {
	var cmds = []string{
		fmt.Sprintf(`sudo sed -i 's/^SELINUX=.*/SELINUX=permissive/' %s`, selinuxCfgFilePath),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
