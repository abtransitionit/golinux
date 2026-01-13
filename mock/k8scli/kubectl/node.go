package kubectl

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func (i *Node) Describe(hostName, helmHost string, logger logx.Logger) (string, error) {
	// get and play cli
	result, err := run.RunCli(helmHost, i.cliToDescribe(), logger)
	if err != nil {
		return "", err
	}
	// handle success
	logger.Debug("listed node")
	return result, nil
}

func (i *Node) cliToDescribe() string {
	var cmds = []string{
		fmt.Sprintf(`kubectl describe nodes %s`, i.Name),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

func (nodeService) List(hostName, helmHost string, logger logx.Logger) (string, error) {
	// get and play cli
	result, err := run.RunCli(helmHost, NodeSvc.cliToList(), logger)
	if err != nil {
		return "", err
	}
	// handle success
	logger.Debug("listed items")
	return result, nil
}

func (nodeService) cliToList() string {
	var cmds = []string{
		`kubectl get nodes -Ao wide | awk '{print $1,$8,$(NF-1),$6,$2,$4,$3}' | column -t`,
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
