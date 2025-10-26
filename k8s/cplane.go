package k8s

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/filex"
	"github.com/abtransitionit/gocore/tpl"
)

// func getK8sConfigFilePath() string {
// 	return filepath.Join("/tmp", "config.yaml")
// }

func InitCPlane(k8sConf K8sConf) (string, error) {

	// get the resolved configuration file
	config, err := getConfig(k8sConf)
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

func InitCPlaneWithReset(k8sConf K8sConf) (string, error) {

	// get the resolved configuration file
	config, err := getConfig(k8sConf)
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

func GetJoinCli() string {

	// build the CLI
	var clis = []string{
		`sudo kubeadm token create --print-join-command`,
	}
	cli := strings.Join(clis, " && ")

	// return
	return cli
}

// Name: GetConfig
//
// Description: resolve the templated config file and return it as a YamlString
func getConfig(k8sConf K8sConf) (string, error) {

	// define the structure that holds the vars that will be used to resolve the templated file
	k8sConfigFileTplVar := K8sConf{
		K8sVersion:     k8sConf.K8sVersion,
		K8sPodCidr:     k8sConf.K8sPodCidr,
		K8sServiceCidr: k8sConf.K8sServiceCidr,
		CrSocketName:   k8sConf.CrSocketName,
	}

	// resolve the templated file
	K8sConfigFile, err := tpl.ResolveTplConfig(configFileTpl, k8sConfigFileTplVar)
	if err != nil {
		return "", fmt.Errorf("faild to resolve templated repo file, for the file: %s", configFileTpl)
	}

	// resturn the YamlString
	return K8sConfigFile, nil

}
