package cilium

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func (ciliumService) DisplayStatus(hostName, helmHost string, logger logx.Logger) (string, error) {

	// play cli
	output, err := run.RunCli(helmHost, `cilium status`, logger)
	if err != nil {
		return "", fmt.Errorf("failed to display cilium status > %w", err)
	}
	return output, nil
}

// description: delete a helm release from a K8s cluster
func (ciliumService) InstallIngress(hostName, helmHost string, releaseName string, logger logx.Logger) error {

	// // 1 - get instance from release name
	// release := helm.GetRelease(releaseName, "", "", "", nil)
	// // 2 - get the play the cli
	// _, err := run.RunCli(helmHost, CilumSvc.cliToAddIngress(release, YamlIngressCfg), logger)
	// if err != nil {
	// 	return fmt.Errorf("%s:%s:%s > upgrading the cilium release in the cluster with the cilium ingress> %w", hostName, helmHost, release.Name, err) // TODO: add the release helm release from the cluster > %w", hostName, helmHost, i.Name, err)
	// }

	// handle success
	logger.Debugf("%s:%s:%s > upgraded the helm cilium release with the cilium ingress", hostName, helmHost, releaseName)
	return nil
}

// func (ciliumService) cliToAddIngress(i *helm.Release, cfgAsbyte []byte) string {
// 	// 31 - build cli
// 	encoded := base64.StdEncoding.EncodeToString(cfgAsbyte)
// 	var cmds = []string{
// 		fmt.Sprintf(
// 			`printf '%s' | base64 -d | helm upgrade % %s --reuse-values --atomic --wait --namespace %s -f -`,
// 			encoded,
// 			i.Name,
// 			i.Chart.QName,
// 			i.Namespace,
// 		),
// 	}
// 	cli := strings.Join(cmds, " && ")
// 	return cli
// }
