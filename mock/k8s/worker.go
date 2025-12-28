package k8s

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

// Description: Add a Node (that is a node) to an existing cluster
//
// Notes:
//
//   - Worker represents a Kubernetes worker node.
//   - the node is not reset before being added to the cluster.
func (worker *Worker) Add(logger logx.Logger) error {
	// 1 - get join cli
	joinCli := "todo"

	// 1 - get cli
	cli := worker.cliForAdd(joinCli)
	// log
	logger.Infof("%s > will play cli: %s ", worker.Name, cli)
	// handle success
	return nil
}

// Description: Add a Node (that is a node) to an existing cluster
//
// Notes:
//
//   - Worker represents a Kubernetes worker node.
//   - the node is not reset before being added to the cluster.
func (worker *Worker) cliForAdd(joinCli string) string {

	// build the CLI
	var clis = []string{
		fmt.Sprintf(`sudo %s`, joinCli),
	}
	cli := strings.Join(clis, " && ")

	// return
	return cli
}

// Description: Add a Node (that is a node) to an existing cluster
//
// Notes:
//
//   - Worker represents a Kubernetes worker node.
//   - the node is being reset before being added to the cluster.
func (worker *Worker) AddWithReset(joinCli string, logger logx.Logger) (string, error) {
	// func AddWorkerWithReset(joinCli string) (string, error) {

	// build the CLI
	var clis = []string{
		`sudo kubeadm reset --force`,
		fmt.Sprintf(`sudo %s`, joinCli),
	}
	cli := strings.Join(clis, " && ")

	// return
	return cli, nil
}
