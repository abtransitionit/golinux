package selinux

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/property"
)

func (selinux *Selinux) Configure(hostName string, logger logx.Logger) (string, error) {

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

	// Here: osFamily in ["rhel", "fedora"] => DoTheJob
	// 3 - configure selinux
	// 31 - runtime configuration
	// cli = selinux.configureAtRuntime(logger)
	// if err != nil {
	// 	return "", err
	// }

	// 32 -  startup configuration
	// cli = selinux.configureAtStartup(logger)
	// if err != nil {
	// 	return "", err
	// }

	// log
	logger.Infof("%s:%s > Configuring Selinux called", hostName, osFamily)
	// handle success
	return osFamily, nil
}

func (selinux *Selinux) ConfigureAtRuntime() string {
	var cmds = []string{
		"sudo setenforce 0",
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

func (selinux *Selinux) ConfigureAtStartup() string {

	var cmds = []string{
		fmt.Sprintf(`sudo sed -i 's/^SELINUX=.*/SELINUX=permissive/' %s`, CfgFilePath),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
