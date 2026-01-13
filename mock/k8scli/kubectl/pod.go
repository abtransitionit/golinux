package kubectl

import (
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func (podService) List(hostName, helmHost string, logger logx.Logger) (string, error) {
	// get and play cli
	result, err := run.RunCli(helmHost, PodSvc.cliToList(), logger)
	if err != nil {
		return "", err
	}
	// handle success
	logger.Debug("listed items")
	return result, nil
}

func (podService) cliToList() string {
	var cmds = []string{
		`kubectl get pods -Ao wide | awk '{print $1,$2,$4,$7, $8}' | column -t`,
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
