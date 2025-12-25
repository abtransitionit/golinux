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

	// 3 - Here: osFamily in ["rhel", "fedora"] => configure selinux
	// 31 - at runtime
	cli := selinux.configureAtRuntime()
	if err != nil {
		return "", fmt.Errorf("configring selinux at runtime > %w ", err)
	}
	logger.Infof("%s:%s > runtime > %s", hostName, osFamily, cli)
	// 32 -  at startup
	cli = selinux.configureAtStartup()
	if err != nil {
		return "", fmt.Errorf("configring selinux at startup > %w ", err)
	}
	logger.Infof("%s:%s > startup > %s", hostName, osFamily, cli)

	// handle success
	return osFamily, nil
}

func (selinux *Selinux) configureAtRuntime() string {
	var cmds = []string{
		"sudo setenforce 0",
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

func (selinux *Selinux) configureAtStartup() string {

	var cmds = []string{
		fmt.Sprintf(`sudo sed -i 's/^SELINUX=.*/SELINUX=permissive/' %s`, CfgFilePath),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
