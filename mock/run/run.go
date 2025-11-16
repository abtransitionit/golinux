package run

import (
	"fmt"
	"os/exec"

	"github.com/abtransitionit/gocore/logx"
)

// Description: executes a CLI locally or remotly (via SSH) and returns its output as a string
//
// Notes:
// - choice to is made based on the targetName parameter

func RunCli(targetName, cli string, logger logx.Logger) (string, error) {
	if targetName == "local" {
		return RunOnLocal(cli, logger)
	}
	return RunOnRemote(targetName, cli, logger)
}

// Description: executes a CLI locally and returns its output as a string
func RunOnLocal(cde string, logger logx.Logger) (string, error) {
	// define CLI
	cli := exec.Command("sh", "-c", cde)
	// log
	// logger.Debugf("executed local CLI > %s ", cli)
	// run CLI
	output, err := cli.CombinedOutput()
	// handle system error
	if err != nil {
		return string(output), fmt.Errorf("running cli locally: %v, output: %s", err, string(output))
	}
	// handle success
	return string(output), nil
}

// Description: executes a CLI remotely via SSH and returns its output as a string
func RunOnRemote(vm string, cde string, logger logx.Logger) (string, error) {
	// define CLI - Build the SSH command: ssh <vm> "<cli>"
	sshCmd := fmt.Sprintf("ssh %s \"%s\"", vm, cde)
	cli := exec.Command("sh", "-c", sshCmd)
	// log
	// logger.Debugf("executed remote CLI: %s > %s", vm, cli)
	// run CLI
	output, err := cli.CombinedOutput()
	// handle system error
	if err != nil {
		return string(output), fmt.Errorf("running cli remotely on %s: %v, output: %s", vm, err, string(output))
	}
	// handle success
	return string(output), nil
}

// ExecuteCli executes a CLI command locally and returns its output as a string.
// func ExecuteCli(cli string, logger logx.Logger) (string, error) {
// 	// Split command into name and args (simple approach, works for basic commands)
// 	parts := strings.Fields(cli)
// 	if len(parts) == 0 {
// 		return "", nil
// 	}

// 	cmd := exec.Command(parts[0], parts[1:]...)
// 	var out bytes.Buffer
// 	var stderr bytes.Buffer
// 	cmd.Stdout = &out
// 	cmd.Stderr = &stderr

// 	err := cmd.Run()
// 	if err != nil {
// 		return "", fmt.Errorf("%v: %s", err, stderr.String())
// 	}

// 	// Trim whitespace and return
// 	return strings.TrimSpace(out.String()), nil
// }
