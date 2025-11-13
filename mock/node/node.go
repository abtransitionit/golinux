package node

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// Description: check a ssh target is reachable
//
// Notes:
// - a ssh target is the localhost, a remote VM or a container
func IsSshReachable(targetName string, logger logx.Logger) (bool, error) {

	// 1 - Check node is configured.
	isConfigured, err := IsSshConfigured(targetName, logger)
	if err != nil {
		return false, fmt.Errorf("checking ssh configuration: %w", err)
	}
	if !isConfigured {
		return false, nil
	}

	// 2 -  Now that we know the node is configured, check if it's reachable.
	cli := fmt.Sprintf("ssh -o BatchMode=yes -o ConnectTimeout=5 %s 'exit'", targetName)

	_, err = run.RunOnLocal(cli, logger)
	if err != nil {
		// If an error is returned, it means the SSH connection failed
		// This is the expected behavior for a non-reachable host.
		return false, nil
	}

	return true, nil

}

// Description: check a ssh target is configured.
//
// Notes:
// - a ssh target is the localhost, a remote VM or a container

func IsSshConfigured(targetName string, logger logx.Logger) (bool, error) {
	// build CLI command
	cli := fmt.Sprintf("ssh -G %s | grep '^hostname ' | cut -d' ' -f2", targetName)

	// run command
	resolvedHostname, err := run.RunOnLocal(cli, logger)
	if err != nil {
		return false, fmt.Errorf("getting resolved hostname: %w", err)
	}

	// trim spaces/newlines
	resolvedHostname = strings.TrimSpace(resolvedHostname)
	// logger.Debugf("resolvedHostname for %s > %q", targetName, resolvedHostname)

	// logic: if hostname is empty or identical, it's not configured
	isConfigured := resolvedHostname != "" && resolvedHostname != targetName

	return isConfigured, nil
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
