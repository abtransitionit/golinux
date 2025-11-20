package node

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// Description: check if a node is SSH configured on a host.
//
// Notes:
// - a node is a remote VM, the localhost, a container or a remote container
// - a host is a node from which the ssh command is executed
func IsSshConfigured(hostName string, nodeName string, logger logx.Logger) (bool, error) {
	// 1 - build CLI to test ssh configuration
	cli := fmt.Sprintf("ssh -T -G %s | grep '^hostname ' | cut -d' ' -f2", nodeName)

	// 2 - run CLI
	resolvedHostname, err := run.RunCli(hostName, cli, logger)

	// 3 - handle system error
	if err != nil {
		return false, fmt.Errorf("host: %s > node: %s > system error > getting resolved hostname: %w", hostName, nodeName, err)
	}

	// 4 - handle logic error - if hostname is empty or identical, it's not configured
	resolvedHostname = strings.TrimSpace(resolvedHostname)
	if resolvedHostname == "" || resolvedHostname == nodeName {
		return false, nil // // SSH not configured - NOT a system error
	}
	// 5 - handle success
	return true, nil
}

// Description: check if a node is SSH reachable from a host.
//
// Notes:
// - a node is a remote VM, the localhost, a container or a remote container
// - a host is a node from which the ssh command is executed
func IsSshReachable(hostName string, nodeName string, logger logx.Logger) (bool, error) {

	// 1 - check node is configured first
	isConfigured, err := IsSshConfigured(hostName, nodeName, logger)
	if err != nil {
		return false, fmt.Errorf("host: %s > node: %s > system error > checking ssh configuration: %w", hostName, nodeName, err)
	}
	if !isConfigured {
		logger.Debugf("host: %s > node: %s > is not SSH configured", hostName, nodeName)
		return false, nil // SSH not configured, not a system error
	}
	//
	// 2 - SSH is configured - Build CLI to test reachability
	cli := fmt.Sprintf("ssh -o BatchMode=yes -o ConnectTimeout=5 %s 'exit'", nodeName)

	// 3 - Run CLI locally or remotely
	_, err = run.RunCli(hostName, cli, logger)

	// 4 - If an error is returned, it means the SSH connection failed. This is the expected behavior for a non-reachable host.
	if err != nil {
		// return false, fmt.Errorf("host: %s > node: %s > is not SSH reachable", hostName, nodeName)
		return false, nil
	}
	// handle success
	return true, nil
}

// Description: check if a node is SSH reachable (within a delayMax) from a host.
//
// Parameters:
// - delay: delay in seconds after which the ssh check failed
//
// Prerequisites:
// - the node is supposed to be ssh configured (this is tested)

// Returns:
// - (-,err) if a system error occurred or host is not ssh configured (this is a prerequisite)
// - (true,nil) if the node is SSH online (within delayMax)
// - (false,nil) if the node was not SSH online (within delayMax)
//
// Notes:
// - a node is a remote VM, the localhost, a container or a remote container
// - a host is a node from which the ssh command is executed
// - the check consist of a SSH connection test every X seconds until the delay is reached
// - the purpose of the check is to test availability for example after a reboot
func IsSshOnline(hostName string, nodeName string, logger logx.Logger) (bool, error) {

	// 1 - check node is configured first
	if isConfigured, err := IsSshConfigured(hostName, nodeName, logger); err != nil {
		return false, fmt.Errorf("host: %s > node: %s > system error > checking ssh configuration: %w", hostName, nodeName, err)
	} else if !isConfigured {
		return false, fmt.Errorf("host: %s > node: %s > is not SSH configured", hostName, nodeName)
	}

	// 1 - Build CLI to test reachability
	cli := fmt.Sprintf("ssh -o BatchMode=yes -o ConnectTimeout=5 %s 'exit' 2>/dev/null && echo true || echo false", nodeName)

	// 2 - Run CLI locally or remotely - return true or false as string
	ok, err := run.RunCli(hostName, cli, logger)

	// 3 - handle system error
	if err != nil {
		return false, fmt.Errorf("host: %s > node: %s > system error > getting ssh reachability: %w", hostName, nodeName, err)
	}

	// 4 - handle logic for success evaluates to true if ok == "true". else evaluates to false
	return ok == "true", nil

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
