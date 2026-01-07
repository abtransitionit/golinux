package k8s

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// Description: Add a Node (that is a node) to an existing cluster
//
// Notes:
//
//   - Worker represents a Kubernetes worker node.
//   - the node is not reset before being added to the cluster.
func (i *Worker) Add(controlPlaneName string, logger logx.Logger) error {
	// 1 - get instance of control plane
	controlPlane := GetCplane(controlPlaneName)

	// 1 - get the join cli
	joinCli, err := controlPlane.GetJoinCli(i.Name, logger)
	if err != nil {
		return fmt.Errorf("%s > getting join cli from control plane %s: %w", i.Name, controlPlaneName, err)
	}

	// get and play cli to add the worker
	cli := i.cliToAdd(joinCli)
	logger.Debugf("%s > cli is : %s", i.Name, cli)
	out, err := run.RunCli(i.Name, cli, logger)
	if err != nil {
		return fmt.Errorf("%s > adding worker node to the cluster > %w", i.Name, err)
	}
	// log output
	logger.Debug("--- BEGIN: add output ---")
	scanner := bufio.NewScanner(bytes.NewReader([]byte(out)))
	for scanner.Scan() {
		logger.Debugf("%s > %s", i.Name, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Warnf("%s > error while logging init output: %v", i.Name, err)
	}
	logger.Debug("--- END ---")

	// handle success
	return nil
}

// Description: Add a Node (that is a node) to an existing cluster
//
// Notes:
//
//   - Worker represents a Kubernetes worker node.
//   - the node is not reset before being added to the cluster.
func (i *Worker) cliToAdd(joinCli string) string {

	// build the CLI
	var clis = []string{
		fmt.Sprintf(`sudo %s `, strings.TrimSpace(joinCli)),
	}
	cli := strings.Join(clis, " && ")

	// return
	return cli
}

// // build the CLI
// encoded := base64.StdEncoding.EncodeToString(ClusterConfAsByte)
// var clis = []string{
// 	fmt.Sprintf(`echo '%s' | base64 -d | sudo kubeadm init --config /dev/stdin`, encoded),
// }
// cli := strings.Join(clis, " && ")
// return cli

// Description: Add a Node (that is a node) to an existing cluster
//
// Notes:
//
//   - Worker represents a Kubernetes worker node.
//   - the node is being reset before being added to the cluster.
func (i *Worker) AddWithReset(joinCli string, logger logx.Logger) (string, error) {
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

// func getJoinCli() string {
// 	return cliToGetJoinCli()
// }

// func (i *Worker) GetJoinCli(controlPlane string, logger logx.Logger) error {
// 	// get instance of control plane

// }
