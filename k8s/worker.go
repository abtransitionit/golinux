package k8s

import (
	"fmt"
	"strings"
)

func AddWorker(joinCli string) (string, error) {

	// build the CLI
	var clis = []string{
		fmt.Sprintf(`sudo %s`, joinCli),
	}
	cli := strings.Join(clis, " && ")

	// return
	return cli, nil
}

func AddWorkerWithReset(joinCli string) (string, error) {

	// build the CLI
	var clis = []string{
		`sudo kubeadm reset --force`,
		fmt.Sprintf(`sudo %s`, joinCli),
	}
	cli := strings.Join(clis, " && ")

	// return
	return cli, nil
}
