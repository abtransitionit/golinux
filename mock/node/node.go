package node

import (
	"fmt"

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

	_, err = run.ExecuteCli(cli, logger)
	if err != nil {
		// If an error is returned, it means the SSH connection failed
		// This is the expected behavior for a non-reachable host.
		return false, nil
	}

	return true, nil

}

func IsSshConfigured(vmName string, logger logx.Logger) (bool, error) {

	// define cli
	cli := fmt.Sprintf("ssh -G %s | grep 'hostname ' | cut -d' ' -f2", vmName)

	// play cli
	resolvedHostname, err := run.ExecuteCli(cli, logger)
	if err != nil {
		return false, fmt.Errorf("getting resolved hostname: %w", err)
	}
	// The logic : compare the resolved hostname to the alias.
	// If they are not the same, the VM is configured. We also check for an empty
	// resolved hostname, which would indicate no match was found.
	isConfigured := resolvedHostname != vmName && resolvedHostname != ""

	// return response
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
