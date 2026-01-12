package dns

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/property"
	"github.com/abtransitionit/golinux/mock/run"
)

func FixDns(hostName string, logger logx.Logger) error {

	// 1 - get host:property
	osFamily, err := property.GetProperty(logger, hostName, "osFamily")
	if err != nil {
		return fmt.Errorf("error 01: %v ", err)
	}

	// 1 - get host:property
	osDistro, err := property.GetProperty(logger, hostName, "osDistro")
	if err != nil {
		return fmt.Errorf("error 01: %v ", err)
	}

	// 2 - skip if Os:family not rhel or fedora
	if osFamily != "rhel" {
		return nil
	}

	// 3 - Here: osFamily in ["rhel"] => almalinux, rocky
	// 31 - get and play cli
	if _, err := run.RunCli(hostName, cliToFixDns(), logger); err != nil {
		return fmt.Errorf("%s:%s:%s > configuring selinux for current session > %v", hostName, osFamily, osDistro, err)
	}

	// handle success
	logger.Infof("%s:%s:%s > fixied missing resolv.conf file", hostName, osFamily, osDistro)
	return nil
}

func cliToFixDns() string {
	// 1 - create the release
	var cmds = []string{
		`sudo mkdir -p /run/systemd/resolve > /dev/null`,
		`sudo ln -sf /etc/resolv.conf /run/systemd/resolve/resolv.conf > /dev/null`,
		`echo fixing resolv.conf`,
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

// func FixResolv(hostName, helmHost string, logger logx.Logger) error {

// 	// 2 - get and play cli
// 	if _, err := run.RunCli(helmHost, i.cliToFixResolv(), logger); err != nil {
// 		return fmt.Errorf("%s:%s:%s > fixing the revolv.conf file > %w", hostName, helmHost, err)
// 	}

// 	// handle success
// 	logger.Debugf("%s:%s:%s > fixed resolv.conf file on the node", hostName, helmHost)
// 	return nil

// }
