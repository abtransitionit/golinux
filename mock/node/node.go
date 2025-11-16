package node

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// Description: check if a node iss SSH configured on a target.
//
// Notes:
// - a node is a remote VM, the localhost, a container or a remote container
// - a target is a node from which the ssh command is executed
func IsSshConfigured(targetName string, nodeName string, logger logx.Logger) (bool, error) {
	// 1 - build CLI to test ssh configuration
	cli := fmt.Sprintf("ssh -T -G %s | grep '^hostname ' | cut -d' ' -f2", nodeName)

	// 2 - run CLI
	resolvedHostname, err := run.RunCli(targetName, cli, logger)

	// 3 - handle system error
	if err != nil {
		return false, fmt.Errorf("target: %s > node: %s > system error > getting resolved hostname: %w", targetName, nodeName, err)
	}

	// 4 - handle logic error - if hostname is empty or identical, it's not configured
	resolvedHostname = strings.TrimSpace(resolvedHostname)
	if resolvedHostname == "" || resolvedHostname == nodeName {
		return false, nil // // SSH not configured - NOT a system error
	}
	// 5 - handle success
	return true, nil
}

// Description: check if a node is SSH reachable from a target.
//
// Notes:
// - a node is a remote VM, the localhost, a container or a remote container
// - a target is a node from which the ssh command is executed
func IsSshReachable(targetName string, nodeName string, logger logx.Logger) (bool, error) {

	// 1 - check node is configured first
	isConfigured, err := IsSshConfigured(targetName, nodeName, logger)
	if err != nil {
		return false, fmt.Errorf("target: %s > node: %s > system error > checking ssh configuration: %w", targetName, nodeName, err)
	}
	if !isConfigured {
		logger.Debugf("target: %s > node: %s > is not SSH configured", targetName, nodeName)
		return false, nil // SSH not configured, not a system error
	}
	//
	// 2 - SSH is configured - Build CLI to test reachability
	cli := fmt.Sprintf("ssh -o BatchMode=yes -o ConnectTimeout=5 %s 'exit'", nodeName)

	// 3 - Run CLI locally or remotely
	_, err = run.RunCli(targetName, cli, logger)

	// 4 - If an error is returned, it means the SSH connection failed. This is the expected behavior for a non-reachable host.
	if err != nil {
		// return false, fmt.Errorf("target: %s > node: %s > is not SSH reachable", targetName, nodeName)
		return false, nil
	}
	// handle success
	return true, nil
}

func IsRemoteVm() bool {
	return true
}

func IsLocalVm() bool {
	return true
}

func IsLocalHost() bool {
	return true
}

func IsContainer() bool {
	return true
}
func IsRemoteContainer() bool {
	return true
}
