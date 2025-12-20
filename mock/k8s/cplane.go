package k8s

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/filex"
	"github.com/abtransitionit/gocore/logx"
)

// func getK8sConfigFilePath() string {
// 	return filepath.Join("/tmp", "config.yaml")
// }

// Description: Initialize the control plane of a Kubernetes cluster
// Notes:
//
//   - CPlane represents a Kubernetes control plane node.
//	 - the control plane is not being reset before being initialized.

func (cplane CPlane) Init(clusterConf ClusterConf, logger logx.Logger) (string, error) {
	// func InitCPlane(k8sConf K8sConf) (string, error) {

	// get the resolved configuration file
	config, err := getConfig(clusterConf)
	if err != nil {
		return "", err
	}

	// build the CLI
	var clis = []string{
		fmt.Sprintf(`printf '%%s\n' '%s'  | sudo kubeadm init --config /dev/stdin`, filex.DeleteLeftTab(config)),
	}
	cli := strings.Join(clis, " && ")

	// return
	return cli, nil
}

// Description: Initialize the control plane of a Kubernetes cluster
//
// Notes:
//
//   - CPlane represents a Kubernetes control plane node.
//	 - the control plane is being reset before being initialized.

func (cplane CPlane) InitCPlaneWithReset(clusterConf ClusterConf, logger logx.Logger) (string, error) {
	//func InitCPlaneWithReset(k8sConf K8sConf) (string, error) {

	// get the resolved configuration file
	config, err := getConfig(clusterConf)
	if err != nil {
		return "", err
	}

	// build the CLI
	var clis = []string{
		`sudo kubeadm reset --force`,
		fmt.Sprintf(`printf '%%s\n' '%s'  | sudo kubeadm init --config /dev/stdin`, filex.DeleteLeftTab(config)),
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
func getConfig(k8sConf K8sConf) (string, error) {

	// define the structure that holds the vars that will be used to resolve the templated file
	k8sConfigFileTplVar := K8sConf{
		K8sVersion:   k8sConf.K8sVersion,
		PodCidr:      k8sConf.PodCidr,
		ServiceCidr:  k8sConf.ServiceCidr,
		CrSocketName: k8sConf.CrSocketName,
	}

	// resolve the templated file
	K8sConfigFile, err := tpl.ResolveTplConfig(configFileTpl, k8sConfigFileTplVar)
	if err != nil {
		return "", fmt.Errorf("faild to resolve templated repo file, for the file: %s", configFileTpl)
	}

	// resturn the YamlString
	return K8sConfigFile, nil

}
