package k8s

import (
	"strings"
)

// Name: ResetNode
//
// Description: return a CLI to reset a worker or a control plane
func ResetNode() (string, error) {

	// build the CLI
	var clis = []string{
		`sudo kubeadm reset --force`,
	}
	cli := strings.Join(clis, " && ")

	// return
	return cli, nil
}
