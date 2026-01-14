package kubectl

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func runKubectl(hostName, helmHost, action string, cli string, logger logx.Logger) (string, error) {

	result, err := run.RunCli(helmHost, cli, logger)
	if err != nil {
		return "", err
	}

	logger.Debugf("%s:%s > %s (%s)", hostName, helmHost, action, cli)
	return result, nil
}
