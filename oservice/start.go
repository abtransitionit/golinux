package oservice

import (
	"fmt"
	"os/exec"

	"github.com/abtransitionit/gocore/errorx"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/distro"
)

// StartService starts a service using the appropriate service manager.
func StartService(serviceName string) error {
	if err := distro.CheckLinuxOS(); err != nil {
		return err
	}

	logx.Info("Attempting to start service: %s", serviceName)

	// ... rest of your code ...
	cmd := exec.Command("systemctl", "start", serviceName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return errorx.WithStack(fmt.Errorf("failed to start service '%s': %s, output: %s", serviceName, err, output))
	}

	logx.Info("Successfully started service: %s", serviceName)
	return nil
}
