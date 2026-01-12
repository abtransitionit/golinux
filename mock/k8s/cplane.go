package k8s

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// Description: Initialize the control plane of a Kubernetes cluster
//
// Notes:
//
//   - the control plane is not being reset before being initialized.
func (i *CPlane) Init(clusterParam ClusterParam, logger logx.Logger) error {
	// get the resolved configuration file as byte[]
	fullConfigAsbyte, err := getFullClusterConf(yamlCfg, clusterParam, logger)
	if err != nil {
		return fmt.Errorf("%s > getting full cluster config > %w", i.Name, err)
	}

	// log entire file
	logger.Debug("--- BEGIN: cluster FULL config  ---")
	scanner := bufio.NewScanner(bytes.NewReader(fullConfigAsbyte))
	for scanner.Scan() {
		logger.Debug(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Warnf("error while logging rendered kubeadm config: %v", err)
	}
	logger.Debug("--- END ---")

	// get and play cli
	out, err := run.RunCli(i.Name, i.cliToInit(fullConfigAsbyte), logger)
	if err != nil {
		return fmt.Errorf("%s > initializing node > %w", i.Name, err)
	}

	// log output
	logger.Debug("--- BEGIN: init output ---")
	scanner = bufio.NewScanner(bytes.NewReader([]byte(out)))
	for scanner.Scan() {
		logger.Debug(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Warnf("%s > error while logging init output: %v", i.Name, err)
	}
	logger.Debug("--- END ---")

	// handle success
	// logger.Infof("%s > initialized the control plane", i.Name)
	return nil
}

func (i *CPlane) cliToInit(ClusterConfAsByte []byte) string {

	// build the CLI
	encoded := base64.StdEncoding.EncodeToString(ClusterConfAsByte)
	var clis = []string{
		fmt.Sprintf(`echo '%s' | base64 -d | sudo kubeadm init --config /dev/stdin`, encoded),
	}
	cli := strings.Join(clis, " && ")
	return cli
}

// Description: Initialize the control plane of a Kubernetes cluster
//
// Notes:
//
//   - the control plane is being reset before being initialized.
func (i *CPlane) InitWithReset(clusterParam ClusterParam, logger logx.Logger) (string, error) {

	// get the resolved configuration file as byte[]
	fullConfigAsbyte, err := getFullClusterConf(yamlCfg, clusterParam, logger)
	if err != nil {
		return "", err
	}
	// log
	logger.Debug("--- BEGIN:Rendered kubeadm config  ---")
	scanner := bufio.NewScanner(bytes.NewReader(fullConfigAsbyte))
	for scanner.Scan() {
		logger.Debug(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Warnf("error while logging rendered kubeadm config: %v", err)
	}
	logger.Debug("--- END ---")

	// build the CLI
	var clis = []string{
		`sudo kubeadm reset --force`,
		// fmt.Sprintf(`printf '%%s\n' '%s'  | sudo kubeadm init --config /dev/stdin`, filex.DeleteLeftTab(config)),
		// fmt.Sprintf(`printf '%%s\n' '%s'  | sudo kubeadm init --config /dev/stdin`, filex.DeleteLeftTab(fullConfigAsbyte)),
	}
	cli := strings.Join(clis, " && ")

	// return
	return cli, nil
}
func (i *CPlane) GetJoinCli(WorkerName string, logger logx.Logger) (string, error) {
	// get and play cli
	out, err := run.RunCli(i.Name, i.cliToGetJoinCli(), logger)
	if err != nil {
		return "", fmt.Errorf("%s > getting join cli from control plane %s > %w", WorkerName, i.Name, err)
	}

	// handle success
	logger.Infof("%s > got the join cli from the control plane %s", WorkerName, i.Name)
	return out, nil
}
func (i *CPlane) cliToGetJoinCli() string {

	// build the CLI
	var clis = []string{
		`sudo kubeadm token create --print-join-command`,
	}
	cli := strings.Join(clis, " && ")

	// return
	return cli
}

// 1 - get the initial config from the cplane ([]byte)
// 2 - install it on the kubectl node (~/.kube/config)

// ssh o1u "cat /etc/kubernetes/admin.conf"  > ~/.kube/config
// sudo cat /etc/kubernetes/admin.conf > ~/.kube/config
// chmod 600 ~/.kube/config
// export KUBECONFIG=~/.kube/config

// // Name: GetConfig
// //
// // Description: resolve the templated config file and return it as a YamlString
// func getConfig(k8sConf ClusterConf) (string, error) {

// 	// define the structure that holds the vars that will be used to resolve the templated file
// 	k8sConfigFileTplVar := ClusterConf{
// 		K8sVersion:   k8sConf.K8sVersion,
// 		PodCidr:      k8sConf.PodCidr,
// 		ServiceCidr:  k8sConf.ServiceCidr,
// 		CrSocketName: k8sConf.CrSocketName,
// 	}

// 	// resolve the templated file
// 	K8sConfigFile, err := tpl.ResolveTplConfig(configFileTpl, k8sConfigFileTplVar)
// 	if err != nil {
// 		return "", fmt.Errorf("faild to resolve templated repo file, for the file: %s", configFileTpl)
// 	}

// 	// resturn the YamlString
// 	return K8sConfigFile, nil

// }

// func getK8sConfigFilePath() string {
// 	return filepath.Join("/tmp", "config.yaml")
// }

// fmt.Sprintf(`printf '%%s\n' '%s'  | sudo kubeadm init --config /dev/stdin`, filex.DeleteLeftTab(config)),
// fmt.Sprintf(`printf '%%s\n' '%s'  | sudo kubeadm init --config /dev/stdin`, fmt.Println(string(fullConfigAsbyte))),
