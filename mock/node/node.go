package node

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/property"
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
	if isConfigured, err := IsSshConfigured(hostName, nodeName, logger); err != nil {
		return false, fmt.Errorf("host: %s > node: %s > system error > checking ssh configuration: %w", hostName, nodeName, err)
	} else if !isConfigured {
		return false, fmt.Errorf("host: %s > node: %s > is not SSH configured", hostName, nodeName)
	}

	// 2 - SSH is configured - Build CLI to test reachability
	cli := fmt.Sprintf("ssh -o BatchMode=yes -o ConnectTimeout=5 %s 'exit'", nodeName)

	// 3 - Run CLI locally or remotely
	_, err := run.RunCli(hostName, cli, logger)

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
func IsSshOnline(hostName string, nodeName string, delayMax string, logger logx.Logger) (bool, error) {

	// 1 - check parameters - Convert delayMax provided as secondAsString into to seconds
	maxDelayInSec, err := strconv.Atoi(delayMax)
	if err != nil {
		return false, fmt.Errorf("invalid delayMax value '%s': %w", delayMax, err)
	}
	// 2 - check node is configured first
	if isConfigured, err := IsSshConfigured(hostName, nodeName, logger); err != nil {
		return false, fmt.Errorf("host: %s > node: %s > system error > checking ssh configuration: %w", hostName, nodeName, err)
	} else if !isConfigured {
		return false, fmt.Errorf("host: %s > node: %s > is not SSH configured", hostName, nodeName)
	}

	// 3 - Build CLI to test reachability
	cli := fmt.Sprintf("ssh -o BatchMode=yes -o ConnectTimeout=5 %s 'exit' 2>/dev/null && echo true || echo false", nodeName)
	// logger.Debugf("play cli: %s", cli)

	// 4 - loop until delayMax is expired
	start := time.Now()
	timeout := time.Duration(maxDelayInSec) * time.Second

	for {
		// 41 - Run CLI locally or remotely - return true or false as string
		ok, err := run.RunCli(hostName, cli, logger)

		// 42 - system error ? → if yes → finish
		if err != nil {
			return false, fmt.Errorf("host: %s > node: %s > system error > getting ssh reachability: %w", hostName, nodeName, err)
		}
		// 43 - SSH reachable ? → if yes → finish
		if strings.TrimSpace(ok) == "true" {
			// get osKernelVersion
			osKernelVersion, err := property.GetProperty(logger, nodeName, "osKernelVersion")
			if err != nil {
				return false, err
			}
			// logè
			logger.Infof("%s/%s > KVersion: %s", hostName, nodeName, osKernelVersion)
			return true, nil
		}

		// 44 - Not SSH reachable → are we outside delayMax ? → if yes → finish
		if time.Since(start) > timeout {
			return false, nil // Not reachable within delayMax
		}

		// 44 - inside delayMax → wait before retry
		time.Sleep(1 * time.Second) // Wait before retry
	}

}

// Description: reboot a host if needed
func RebootIfNeeded(hostName string, logger logx.Logger) (string, error) {
	// 1 - get host:property
	osFamily, err := property.GetProperty(logger, hostName, "osFamily")
	if err != nil {
		return "", err
	}
	osDistro, err := property.GetProperty(logger, hostName, "osDistro")
	if err != nil {
		return "", err
	}
	needReboot, err := property.GetProperty(logger, hostName, "needReboot")
	if err != nil {
		return "", err
	}
	osKernelVersion, err := property.GetProperty(logger, hostName, "osKernelVersion")
	if err != nil {
		return "", err
	}

	if strings.TrimSpace(needReboot) == "true" {
		// 2 - log
		logger.Debugf("%s:%s:%s > kVersion: %s. Rebooting", hostName, osFamily, osDistro, osKernelVersion, needReboot)
		// 3 - get play and play it
		out, err := run.RunCli(hostName, "sudo systemctl reboot", logger)
		if err != nil {
			return "", fmt.Errorf("%s > %s:%s > %w > out:%s", hostName, osFamily, osDistro, err, out)
		}
	}
	// handle success
	return "", nil
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
