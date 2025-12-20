package k8s

import (
	"github.com/abtransitionit/gocore/logx"
)

// Description: Reset a node of a cluster
func (node Node) Reset(logger logx.Logger) (string, error) {
	// log
	logger.Info("Resetting Node with cli: sudo kubeadm reset --force")
	// handle success
	return "", nil
}

// Description: Initializing a node of a cluster
func (node Node) Init(logger logx.Logger) (string, error) {
	// log
	logger.Info("Initializing Node")
	// handle success
	return "", nil
}
