package selinux

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/property"
	"github.com/abtransitionit/golinux/mock/run"
)

func (i *Selinux) Configure(hostName string, logger logx.Logger) (string, error) {

	// 1 - get host:property
	osFamily, err := property.GetProperty(logger, hostName, "osFamily")
	if err != nil {
		return "", fmt.Errorf("error 01: %v ", err)
	}

	// 2 - skip if Os:family not rhel or fedora
	if osFamily != "rhel" && osFamily != "fedora" {
		// logger.Debugf("%s:%s 🅐 Skipping selinux configuration", vmName, osFamily)
		return "", nil
	}

	// 3 - Here: osFamily in ["rhel", "fedora"] => configure selinux
	// 31 - configure for session
	if err := i.configureForSession(hostName, logger); err != nil {
		return "", fmt.Errorf("%s:%s > configuring selinux for current session > %v", hostName, osFamily, err)
	}
	logger.Infof("%s:%s > configured selinux for current session", hostName, osFamily)
	// 32 - configure after reboot
	if err := i.configureAfterReboot(hostName, logger); err != nil {
		return "", fmt.Errorf("%s:%s > configuring selinux after a reboot > %v", hostName, osFamily, err)
	}
	logger.Infof("%s:%s > configured selinux after a reboot", hostName, osFamily)

	// handle success
	return "", nil
}

func (i *Selinux) configureForSession(hostName string, logger logx.Logger) error {
	// 1 - get and play cli
	if _, err := run.RunCli(hostName, i.cliToConfigureForSession(), logger); err != nil {
		return err
	}
	// handle success
	return nil
}
func (i *Selinux) configureAfterReboot(hostName string, logger logx.Logger) error {
	// 1 - get and play cli
	if _, err := run.RunCli(hostName, i.cliToConfigureAfterReboot(), logger); err != nil {
		return err
	}
	// handle success
	return nil
}

// description: configures selinux for the current session (at runtime)
func (i *Selinux) cliToConfigureForSession() string {
	var cmds = []string{
		"sudo setenforce 0",
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

// description: configures selinux after a host reboot (at startup)
func (i *Selinux) cliToConfigureAfterReboot() string {
	var cmds = []string{
		fmt.Sprintf(`sudo sed -i 's/^SELINUX=.*/SELINUX=permissive/' %s`, i.CfgFilePath),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

// linux cli to check the status of selinux
// for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "getenforce 2>/dev/null || echo 'getenforce not installed'"; done
// for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "[ -e /sys/fs/selinux/enforce ] && echo \"SELinux: $([ \"$(cat /sys/fs/selinux/enforce)\" = 1 ] && echo Enforcing || echo Permissive)\" || echo \"SELinux: Disabled\""; done
// for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" '[ -e /sys/fs/selinux/enforce ] && [ "$(cat /sys/fs/selinux/enforce)" = 1 ] && echo "SELinux: Enforcing" || { [ -e /sys/fs/selinux/enforce ] && echo "SELinux: Permissive" || echo "SELinux: Disabled"; }'; done
// for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "sestatus 2>/dev/null | head -n 3 || echo 'sestatus not available'"; done
// for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "command -v sestatus >/dev/null && sestatus | head -n 3 || echo 'SELinux not installed'"; done

// setenforce 1 → Enforcing
// setenforce 0 → Permissive
