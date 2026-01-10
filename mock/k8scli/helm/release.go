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

	// 1 - lookup this release's chart (does this chart exists ?)
	_, err := i.getChartFromRelease(hostName, helmHost, logger)
	if err != nil {
		return fmt.Errorf("%s:%s > getting chart: maybe it is not in the helm whitelist or the release has bad config:%w", hostName, helmHost, err) // maybe it is not in the whitelist:%w", hostName, helmHost, err)
	}

	// handle success
	logger.Debugf("%s:%s:%s > installed helm release", hostName, helmHost, i.Name)
	return nil
}

func (i *Release) getChartFromRelease(hostName, helmHost string, logger logx.Logger) (*Chart, error) {
	return nil, nil
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
