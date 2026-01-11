package helm

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// description: delete a helm release from a K8s cluster
func (i *Release) Delete(hostName, helmHost string, logger logx.Logger) error {

	// 1 - check the release exists
	out, err := i.Exists(hostName, helmHost, logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > checking release existence > %w", hostName, helmHost, i.Name, err)
	} else if out != true {
		return fmt.Errorf("%s:%s:%s > releasez does not exist in the k8s cluster", hostName, helmHost, i.Name)
	}

	// 2 - get and play cli - delete the release
	_, err = run.RunCli(helmHost, i.cliToDelete(), logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > deleting helm release from the cluster > %w", hostName, helmHost, i.Name, err)
	}

	// handle success
	logger.Debugf("%s:%s:%s > deleted the helm release from the cluster", hostName, helmHost, i.Name)
	return nil
}

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

// description: check if a release exists in a k8s cluster
func (i *Release) Exists(hostName, helmHost string, logger logx.Logger) (bool, error) {

	// 1 - get and play cli
	output, err := run.RunCli(helmHost, i.cliToCheckExistence(), logger)
	if err != nil {
		return false, err
	}
	// handle success
	logger.Debugf("%s:%s:%s > checked release existence", hostName, helmHost, i.Name)
	boolean := map[string]bool{"true": true, "false": false}[strings.TrimSpace(output)]
	return boolean, nil
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
func (i *Release) cliToInstall() string {
	var cmds = []string{
		`. ~/.profile`,
		fmt.Sprintf(`helm install %s %s --namespace %s --version %s`, i.Name, i.CQName, i.Namespace, i.Version),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

func (i *Release) cliToDelete() string {
	var cmds = []string{
		`. ~/.profile`,
		fmt.Sprintf(`helm uninstall %s --namespace %s`, i.Name, i.Namespace),
	}
	cli := strings.Join(cmds, " && ")
	return cli

}
func (i *Release) cliToCheckExistence() string {

	var cmds = []string{
		`. ~/.profile`,
		fmt.Sprintf(
			`helm status %s --namespace %s >/dev/null 2>&1 && echo "true" || echo "false"`,
			i.Name,
			i.Namespace,
		),
	}
	return strings.Join(cmds, " && ")
}
