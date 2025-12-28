package selinux

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/property"
)

func (i *Selinux) Configure(hostName string, logger logx.Logger) (string, error) {

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
	// 31 - session configure
	// cli := i.configureForSession()
	logger.Infof("%s:%s > configuring selinux for current session", hostName, osFamily)

	// 32 -  startup configure
	// cli = i.configureAfterReboot()
	logger.Infof("%s:%s > configuring selinux at startup", hostName, osFamily)

	// handle success
	return "", nil
}

// description: configures selinux for the current session (at runtime)
func (selinux *Selinux) configureForSession() string {
	var cmds = []string{
		"sudo setenforce 0",
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

// description: configures selinux after a host reboot (at startup)
func (selinux *Selinux) configureAfterReboot() string {

	var cmds = []string{
		fmt.Sprintf(`sudo sed -i 's/^SELINUX=.*/SELINUX=permissive/' %s`, CfgFilePath),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
