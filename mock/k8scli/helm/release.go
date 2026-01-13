package helm

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/k8sapp/cilium"
	"github.com/abtransitionit/golinux/mock/run"
)

func (i *Release) InstallIngressCilium(hostName, helmHost string, logger logx.Logger) error {

	// 1 - check the release exists
	out, err := i.Exists(hostName, helmHost, logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > checking release existence > %w", hostName, helmHost, i.Name, err)
	} else if out != true {
		return fmt.Errorf("%s:%s:%s > release does not exist in the k8s cluster", hostName, helmHost, i.Name)
	}

	// 2 - TODO: this cilium specific and shoub be place elsewhere
	// 21 - get the ApiServerIp
	cli := `kubectl config view --minify | yq -r '.clusters[0].cluster.server' | tr -d '/' | cut -d: -f2`
	output, err := run.RunCli(helmHost, cli, logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > getting k8s cluster api server ip > %w", hostName, helmHost, i.Name, err)
	}

	// 22 - define var placeholder for this chart/release
	varPlaceHolder := map[string]map[string]string{
		"Cluster": {
			"PodCidr":     strings.TrimSpace(i.Param["podcidr"]),
			"ApiServerIp": strings.TrimSpace(output),
		},
	}
	logger.Debugf("varPlaceHolder >  %+v", varPlaceHolder)
	// 23 - get the resolved value file as byte[]
	cfgAsbyte, err := cilium.GetValueFile(cilium.YamlBasicCfg, varPlaceHolder, logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > getting value file > %w", hostName, helmHost, i.Name, err)
	}

	// get and play cli
	_, err = run.RunCli(helmHost, i.cliToUpgrade(cfgAsbyte), logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > upgrading helm cilium release with cilium ingress > %w", hostName, helmHost, i.Name, err)
	}

	// handle success
	logger.Debugf("%s:%s:%s > upgraded the helm cilium release with cilium ingress", hostName, helmHost, i.Name)
	return nil
}

// description: delete a helm release from a K8s cluster
func (i *Release) Delete(hostName, helmHost string, logger logx.Logger) error {

	// 1 - check the release exists
	out, err := i.Exists(hostName, helmHost, logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > checking release existence > %w", hostName, helmHost, i.Name, err)
	} else if out != true {
		return fmt.Errorf("%s:%s:%s > release does not exist in the k8s cluster: %s", hostName, helmHost, i.Name, out)
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
	chart := GetChart(i.Name, i.Chart.QName, i.Chart.Version)

	// 12 - check the chart existence
	out, err := chart.Exists(hostName, helmHost, logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > checking chart existence > %w", hostName, helmHost, chart.QName, err)
	} else if out != true {
		return fmt.Errorf("%s:%s:%s > chart %s does not exist on the helm client", hostName, helmHost, i.Name, chart.QName)
	}

	// 13 - check the chart version existence
	out, err = chart.VersionExists(hostName, helmHost, logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > checking chart version existence > %w", hostName, helmHost, chart.QName, err)
	} else if out != true {
		return fmt.Errorf("%s:%s:%s > chart version %s does not exist on the helm client", hostName, helmHost, i.Name, chart.QName)
	}

	// 2 - TODO: this cilium specific and shoub be place elsewhere
	// 21 - get the ApiServerIp
	cli := `kubectl config view --minify | yq -r '.clusters[0].cluster.server' | tr -d '/' | cut -d: -f2`
	output, err := run.RunCli(helmHost, cli, logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > getting k8s cluster api server ip > %w", hostName, helmHost, i.Name, err)
	}

	// 22 - define var placeholder for this chart/release
	varPlaceHolder := map[string]map[string]string{
		"Cluster": {
			"PodCidr":     strings.TrimSpace(i.Param["podcidr"]),
			"ApiServerIp": strings.TrimSpace(output),
		},
	}
	logger.Debugf("varPlaceHolder >  %+v", varPlaceHolder)
	// 23 - get the resolved value file as byte[]
	cfgAsbyte, err := cilium.GetValueFile(cilium.YamlAllCfg, varPlaceHolder, logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > getting value file > %w", hostName, helmHost, i.Name, err)
	}

	// log
	logger.Debug("--- BEGIN:Rendered Value file  ---")
	scanner := bufio.NewScanner(bytes.NewReader(cfgAsbyte))
	for scanner.Scan() {
		logger.Debug(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Warnf("error while logging rendered kubeadm config: %v", err)
	}
	logger.Debug("--- END ---")

	// 3 - install
	// 31 - build cli
	_, err = run.RunCli(helmHost, i.cliToInstall(cfgAsbyte), logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > installing helm release from chart %s > %w", hostName, helmHost, i.Name, i.Chart.QName, err)
	}

	// handle success
	logger.Debugf("%s:%s:%s > installed helm release from chart %s:%s", hostName, helmHost, i.Name, i.Chart.QName, i.Chart.Version)
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
func (i *Release) cliToInstallV0() string {
	var cmds = []string{
		fmt.Sprintf(`
		helm install %s %s --atomic --wait --namespace %s %s -f %s 
		`,
			i.Name,
			i.Chart.QName,
			i.Namespace,
			i.versionFlag(),
			i.valueFlag()),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
func (i *Release) cliToInstall(cfg []byte) string {
	encoded := base64.StdEncoding.EncodeToString(cfg)
	var cmds = []string{
		fmt.Sprintf(
			`printf '%s' | base64 -d | helm upgrade %s %s --atomic --wait --namespace %s %s -f -`,
			encoded,
			i.Name,
			i.Chart.QName,
			i.Namespace,
			i.versionFlag()),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
func (i *Release) cliToAddUpgradeReuse(cfgAsbyte []byte) string {
	// encode config
	encoded := base64.StdEncoding.EncodeToString(cfgAsbyte)

	// build cli safely
	var cmds = []string{
		fmt.Sprintf(
			`printf '%s' | base64 -d | helm upgrade %s %s --reuse-values --atomic --wait --namespace %s -f -`,
			encoded,
			i.Name,
			i.Chart.QName,
			i.Namespace),
	}

	cli := strings.Join(cmds, " && ")
	return cli

}
func (i *Release) cliToUpgrade(cfgAsbyte []byte) string {
	// encode config
	encoded := base64.StdEncoding.EncodeToString(cfgAsbyte)

	// build cli safely
	var cmds = []string{
		fmt.Sprintf(
			`printf "%s" | base64 -d | helm upgrade %s %s  --atomic --wait --namespace %s -f -`,
			encoded,
			i.Name,
			i.Chart.QName,
			i.Namespace),
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
		fmt.Sprintf(
			`helm status %s --namespace %s >/dev/null 2>&1 && echo "true" || echo "false"`,
			i.Name,
			i.Namespace,
		),
	}
	return strings.Join(cmds, " && ")
}

// description : allows to install a release that do not have a value file
func (i *Release) valueFlag() string {
	if i.ValueFile != "" {
		return fmt.Sprintf("-f %s", i.ValueFile)
	}
	return ""
}
func (i *Release) versionFlag() string {
	if i.Chart.Version != "" {
		return fmt.Sprintf("--version %s", i.Chart.Version)
	}
	return ""
}
