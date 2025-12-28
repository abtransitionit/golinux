package k8s

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

// Description: Initialize the control plane of a Kubernetes cluster
//
// Notes:
//
//   - the control plane is not being reset before being initialized.
func (cplane *CPlane) Init(clusterParam ClusterParam, logger logx.Logger) error {
	// get cli
	cli, err := cplane.cliForInit(clusterParam, logger)
	if err != nil {
		return fmt.Errorf("%s > getting cli: %v", cplane.Name, err)
	}

	// play  cli
	logger.Infof("%s > will play cli: %s ", cplane.Name, cli)

	// handle success
	return nil
}

func (cplane *CPlane) cliForInit(clusterParam ClusterParam, logger logx.Logger) (string, error) {

	// get the resolved configuration file as byte[]
	fullConfigAsbyte, err := getFullClusterConf(yamlCfg, clusterParam, logger)
	if err != nil {
		return "", err
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

	// build the CLI
	encoded := base64.StdEncoding.EncodeToString(fullConfigAsbyte)
	var clis = []string{
		// fmt.Sprintf(`printf '%%s\n' '%s'  | sudo kubeadm init --config /dev/stdin`, filex.DeleteLeftTab(config)),
		// fmt.Sprintf(`printf '%%s\n' '%s'  | sudo kubeadm init --config /dev/stdin`, fmt.Println(string(fullConfigAsbyte))),
		fmt.Sprintf(`echo '%s' | base64 -d | sudo kubeadm init --config /dev/stdin`, encoded),
	}
	cli := strings.Join(clis, " && ")

	// return
	return cli, nil
}

// Description: Initialize the control plane of a Kubernetes cluster
//
// Notes:
//
//   - the control plane is being reset before being initialized.
func (cplane *CPlane) InitWithReset(clusterParam ClusterParam, logger logx.Logger) (string, error) {

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

// func GetJoinCli() string {

// 	// build the CLI
// 	var clis = []string{
// 		`sudo kubeadm token create --print-join-command`,
// 	}
// 	cli := strings.Join(clis, " && ")

// 	// return
// 	return cli
// }

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
