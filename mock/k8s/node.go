package k8s

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// Description: Reset a node of a cluster
func (i *Node) Reset(logger logx.Logger) (string, error) {
	var out string
	var err error

	// get and play cli
	if out, err = run.RunCli(i.Name, i.cliForReset(), logger); err != nil {
		return "", fmt.Errorf("%s > resetting node: %v", i.Name, err)
	}
	// log output
	logger.Debug("--- BEGIN: Reset output ---")
	scanner := bufio.NewScanner(bytes.NewReader([]byte(out)))
	for scanner.Scan() {
		logger.Debugf("%s > %s", i.Name, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Warnf("%s > error while logging reset output: %v", i.Name, err)
	}
	logger.Debug("--- END ---")
	// handle success
	return out, nil
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
