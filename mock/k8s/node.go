package k8s

import (
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

// Description: Reset a node of a cluster
func (node *Node) Reset(logger logx.Logger) (bool, error) {
	// get cli
	cli := node.cliForReset()

	// play  cli
	logger.Infof("%s > will play cli: %s ", node.Name, cli)

	// handle success
	return true, nil
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

// // Description: Initializing a node of a cluster
// func (node *Node) Init(logger logx.Logger) (string, error) {
// 	// log
// 	logger.Info("Initializing Node")
// 	// handle success
// 	return "", nil
// }
