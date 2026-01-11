package run

import (
	"encoding/base64"
	"fmt"
	"os/exec"

	"github.com/abtransitionit/gocore/logx"
)

// Description: executes a CLI locally or remotly (via SSH) and returns its output as a string
//
// Notes:
// - choice to is based on the hostName parameter
func RunCli(hostName, cli string, logger logx.Logger) (string, error) {
	if hostName == "local" {
		return RunOnLocal(cli, logger)
	}
	return RunOnRemote(hostName, cli, logger)
}

// Description: executes a CLI locally and returns its output as a string
func RunOnLocal(cde string, logger logx.Logger) (string, error) {
	// 1 - define CLI
	cli := exec.Command("sh", "-c", cde)

	// if logger != nil {
	// 	logger.Debugf("local: command executed > %s", cli.String())
	// } else {
	// 	fmt.Printf("local: command executed > %s\n", cli.String())
	// }

	// 2 - run CLI
	output, err := cli.CombinedOutput()

	// 3 - handle system error
	if err != nil {
		return string(output), fmt.Errorf("running cli locally: %v", err)
	}

	// 4 - handle success
	return string(output), nil
}

// Description: executes a CLI remotely via SSH and returns its output as a string
func RunOnRemote(hostName string, cde string, logger logx.Logger) (string, error) {

	// 1 - Base64 encode the input to handle complex quoting and special characters.
	cliEncoded := base64.StdEncoding.EncodeToString([]byte(cde))
	// 2 - define the SSH command
	// sshCmd := fmt.Sprintf(`ssh -o BatchMode=yes -o ConnectTimeout=5%s "echo '%s' | base64 --decode | $SHELL -l"`, hostName, cliEncoded)
	sshCmd := fmt.Sprintf(`ssh %s "echo '%s' | base64 --decode | $SHELL -l"`, hostName, cliEncoded)
	// 3 - define Local command that launch the SSH command
	cli := exec.Command("sh", "-c", sshCmd)

	// if logger != nil {
	// 	logger.Debugf("%s: remote command executed > %s", hostName, cli.String())
	// } else {
	// 	fmt.Printf("%s: remote command executed > %s\n", hostName, cli.String())
	// }

	// 2 - run CLI
	output, err := cli.CombinedOutput()

	// 3 - handle system error
	if err != nil {
		return string(output), fmt.Errorf("running cli remotely on %s: %v, output: %s", hostName, err, string(output))
	}

	// 4 - handle success
	return string(output), nil
}

func RunCliQuery(hostName, cli string, logger logx.Logger) (string, error) {
	if hostName == "local" {
		return RunOnLocal(cli, logger)
	}
	return RunOnRemote(hostName, cli, logger)
}

func ExecuteCliQuery(cli string, logger logx.Logger, isLocal bool, remoteHost string) (string, error) {
	var output string
	var err error

	// 1. Determine execution environment and run the command
	if isLocal {
		logger.Debugf("running on local: %s", cli)
		output, err = RunCli("local", cli, logger)
	} else {
		output, err = RunCli(remoteHost, cli, logger)
	}

	if err != nil {
		return output, fmt.Errorf("failed to run command: %s: %w", cli, err)
	}
	return output, nil
}
