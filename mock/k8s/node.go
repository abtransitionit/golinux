package k8s

import (
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

// Description: Reset a node of a cluster
func (node *Node) Reset(logger logx.Logger) error {
	// get cli
	cli := node.cliForReset()

	// play  cli
	logger.Infof("%s > Resetting Node with cli: %s ", node.Name, cli)

	// handle success
	return nil
}
func (node *Node) cliForReset() string {
	// define cli
	var clis = []string{
		`sudo kubeadm reset --force`,
	}
	cli := strings.Join(clis, " && ")

	// handle success
	return cli
}
