package kubectl

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// func runKubectl(hostName, kubectlHost, action string, cli string, logger logx.Logger) (string, error) {

// 	result, err := run.RunCli(kubectlHost, cli, logger)
// 	if err != nil {
// 		return "", err
// 	}

//		logger.Debugf("%s:%s > %s (%s)", hostName, kubectlHost, action, cli)
//		return result, nil
//	}
func play(hostName, kubectlHost, action string, cli string, logger logx.Logger) (string, error) {

	result, err := run.RunCli(kubectlHost, cli, logger)
	if err != nil {
		return "", err
	}

	logger.Debugf("%s:%s > %s (%s)", hostName, kubectlHost, action, cli)
	return result, nil
}
