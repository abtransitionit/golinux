package k8s

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func (kubectl *Kubectl) PlayCli(cli string) error {
	return nil
}
func (i *Kubectl) Configure(hostName string, logger logx.Logger) error {
	// get cli and play it locally
	if _, err := run.RunCli("local", i.cliToInstallKubectlConfig(i.CplaneNdeName, i.InstallNodeName), logger); err != nil {
		return err
	}
	logger.Debugf("%s:%s > installed kubectl config from cplane: %s", hostName, i.InstallNodeName, i.CplaneNdeName)
	return nil
}

func (i *Kubectl) cliToInstallKubectlConfig(cplaneNodeName, kubectlNodeName string) string {

	// build the CLI
	var clis = []string{
		// `sudo cat /etc/kubernetes/admin.conf | install -D -m 600 /dev/stdin ~/.kube/config`,
		fmt.Sprintf(`ssh %s 'sudo cat /etc/kubernetes/admin.conf' | ssh %s 'install -D -m 600 /dev/stdin ~/.kube/config'`, cplaneNodeName, kubectlNodeName),
	}
	cli := strings.Join(clis, " && ")

	// return
	return cli
}
