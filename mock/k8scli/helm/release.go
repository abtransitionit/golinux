package helm

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// description: install an instance of a helm chart into a K8s cluster's namespace
//
// notes:
// - it create a sets of K8s resources (svc, deploy, ingress, etc) from the chart
func (i *Release) Install(hostName, helmHost string, logger logx.Logger) error {

	// 1 - check this chart exist

	// 11 - create a chart instance
	chart := GetChart(i.Name, i.CQName, i.Version)

	// 12 - check the chart existence
	out, err := chart.Exists(hostName, helmHost, logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > checking chart existence > %w", hostName, helmHost, chart.QName, err) // maybe it is not in the whitelist:%w", hostName, helmHost, err)
	} else if out != true {
		return fmt.Errorf("%s:%s:%s > chart %s does not exist on the helm client", hostName, helmHost, i.Name, chart.QName)
	}

	// 13 - check the chart version existence
	out, err = chart.VersionExists(hostName, helmHost, logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > checking chart version existence > %w", hostName, helmHost, chart.QName, err) // maybe it is not in the whitelist:%w", hostName, helmHost, err)
	} else if out != true {
		return fmt.Errorf("%s:%s:%s > chart version %s does not exist on the helm client", hostName, helmHost, i.Name, chart.QName)
	}

	// handle success
	logger.Debugf("%s:%s:%s > installed helm release from chart %s", hostName, helmHost, i.Name, i.CQName)
	return nil
}

func (releaseService) List(hostName string, helmHost string, logger logx.Logger) (string, error) {
	// 1 - get and play cli
	out, err := run.RunCli(helmHost, ReleaseSvc.cliToList(), logger)
	if err != nil {
		return "", fmt.Errorf("%s:%s > listing helm repos > %w", hostName, helmHost, err)
	}
	// 	// handle success
	return out, nil
}

func (releaseService) cliToList() string {
	var cmds = []string{
		`. ~/.profile`,
		"helm list -A", //  list releases in all namespace
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
func (releaseService) cliToInstall() string {
	var cmds = []string{
		`. ~/.profile`,
		"helm list -A", //  list releases in all namespace
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
