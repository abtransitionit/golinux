package cilium

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func (ciliumService) DisplayStatus(hostName, helmHost string, logger logx.Logger) (string, error) {

	// play cli
	output, err := run.RunCli(helmHost, `cilium status`, logger)
	if err != nil {
		return "", fmt.Errorf("failed to display cilium status > %w", err)
	}
	return output, nil
}

func play(hostName, helmHost, action string, cli string, logger logx.Logger) (string, error) {

	result, err := run.RunCli(helmHost, cli, logger)
	if err != nil {
		return "", err
	}

	logger.Debugf("%s:%s > %s (%s)", hostName, helmHost, action, cli)
	return result, nil
}
