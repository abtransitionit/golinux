package property

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func GetProperty(logger logx.Logger, hostName string, propertyName string, propertyParams ...string) (string, error) {
	// define var
	const goAgentPath = "luca" // in /usr/local/bin
	var output string
	var err error

	if hostName == "local" {
		// 1 -  local execution - get handler
		// 11 - get handler
		fnHandler, ok := PropertyMap[propertyName]
		// 12 - handle error
		if !ok {
			return "", fmt.Errorf("unknown property requested: %s", propertyName)
		}
		// 13 - play function
		output, err = fnHandler(propertyParams...)
	} else {
		// 2 - remote execution
		// 21 - define CLI
		quotedParams := make([]string, len(propertyParams))
		for i, param := range propertyParams {
			quotedParams[i] = fmt.Sprintf("%q", param)
		}
		PropertyAsFlag := fmt.Sprintf("%s%s", "--", propertyName)
		cli := fmt.Sprintf("%s property %s %s", goAgentPath, PropertyAsFlag, strings.Join(quotedParams, " "))
		// logger.Debugf("CLI is : %s", cli)
		// 22 - run CLI on remote
		output, err = run.RunOnRemote(hostName, cli, nil)
	}

	// ---------  common to local and remote execution ---------

	// 3 - handle system error
	if err != nil {
		return "", fmt.Errorf("%s > getting property %q > %v > output: %s", hostName, propertyName, err, output)

	}
	// 24 - handle success
	return strings.TrimSpace(output), nil
}
