package cli

import "github.com/abtransitionit/gocore/logx"

func (kubectl *Kubectl) PlayCli(cli string) error {
	return nil
}
func (i *Kubectl) Configure(hostName string, logger logx.Logger) error {

	// 2 - do the job
	logger.Debugf("%s > get the initial config from the cplane:  %s", hostName, i.CplaneNdeName) // TODO: as  ([]byte)
	logger.Debugf("%s > paste config on the kubectl node: %s", hostName, i.InstallNodeName)      // TODO: as  (~/.kube/config)

	// handle success
	return nil
}
