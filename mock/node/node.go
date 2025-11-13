package node

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func IsSshReachable(nodeName string, logger logx.Logger) (bool, error) {

	// 1 - Check node is configured.
	isConfigured, err := IsSshConfigured(nodeName, logger)
	if err != nil {
		return false, fmt.Errorf("checking ssh configuration")
	}
	if !isConfigured {
		return false, nil
	}

	// 2 -  Now that we know the node is configured, check if it's reachable.
	cli := fmt.Sprintf("ssh -o BatchMode=yes -o ConnectTimeout=5 %s 'exit'", nodeName)

	_, err = run.RunOnLocal(cli, logger)
	if err != nil {
		// If an error is returned, it means the SSH connection failed
		// This is the expected behavior for a non-reachable host.
		return false, nil
	}

	return true, nil

}

// Description: check if a VM is SSH configured.
func IsSshConfigured(vmName string, logger logx.Logger) (bool, error) {
	// build CLI command
	cli := fmt.Sprintf("ssh -G %s | grep '^hostname ' | cut -d' ' -f2", vmName)

	// run command
	resolvedHostname, err := run.RunOnLocal(cli, logger)
	if err != nil {
		return false, fmt.Errorf("getting resolved hostname: %w", err)
	}

	// trim spaces/newlines
	resolvedHostname = strings.TrimSpace(resolvedHostname)
	// logger.Debugf("resolvedHostname for %s > %q", vmName, resolvedHostname)

	// logic: if hostname is empty or identical, it's not configured
	isConfigured := resolvedHostname != "" && resolvedHostname != vmName

	return isConfigured, nil
}

func IsVm() bool {
	return true
}

func IsLocal() bool {
	return true
}

func IsContainer() bool {
	return true
}
