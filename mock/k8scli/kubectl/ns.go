package kubectl

import (
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func (namespaceService) List(hostName, helmHost string, logger logx.Logger) (string, error) {
	// get and play cli
	result, err := run.RunCli(helmHost, NamespaceSvc.cliToList(), logger)
	if err != nil {
		return "", err
	}
	// handle success
	logger.Debug("listed items")
	return result, nil
}

func (namespaceService) cliToList() string {
	var cmds = []string{
		`kubectl get namespaces`,
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
