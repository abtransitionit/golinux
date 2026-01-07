package oskernel

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/file"
	"github.com/abtransitionit/golinux/mock/run"
)

// description: load a kernel module parameter
func (i *ParameterSet) Load(hostName string, logger logx.Logger) error {

	// check
	if len(i.ParameterSlice) == 0 {
		return fmt.Errorf("%s > Parameter list is empty", hostName)
	}

	// 1 - load for session
	if err := i.loadForSession(hostName, logger); err != nil {
		return fmt.Errorf("%s : loading Kernel module parameters for the current session > %w", hostName, err)
	}
	// 2 - load after a reboot (at startup)
	if err := i.loadAfterReboot(hostName, logger); err != nil {
		return fmt.Errorf("%s : loading Kernel module parameters for the current session > %w", hostName, err)
	}

	// handle success
	return nil
}

// description: loads a list of kernel module parameters for the current session (at runtime)
func (i ParameterSet) loadForSession(hostName string, logger logx.Logger) error {
	// 1 - get and play cli
	cli := i.cliToLoadForSession()
	if _, err := run.RunCli(hostName, cli, logger); err != nil {
		return err
	}
	// handle success
	logger.Debugf("%s > loaded Kernel module parameters for current session", hostName)
	return nil
}

// description: load a list of kernel module parameter for the current session
func (i ParameterSet) cliToLoadForSession() string {
	var cmds []string
	for _, item := range i.ParameterSlice {
		cmds = append(cmds,
			fmt.Sprintf("sudo sysctl -w %s=%s || true", item.Name, item.Value),
		)
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

func (i ParameterSet) loadAfterReboot(hostName string, logger logx.Logger) error {
	// 1 - get yaml that contains kernel conf
	cfg, err := getCfg()
	if err != nil {
		return fmt.Errorf("getting kernel conf: %w", err)
	}
	// 2 - define the file that will contain the module parameter list
	cfgFilePath := filepath.Join(cfg.Conf.Folder.Param, i.CfgFileName)

	// 3 - define the file content of the file - 1 parameter per line
	var b strings.Builder
	for _, m := range i.ParameterSlice {
		b.WriteString(fmt.Sprintf("%s = %s\n", m.Name, m.Value))
	}
	content := b.String()

	// 4 - create this file as sudo with this content exactly
	cli := file.SudoCreateFileFromStringBase64(cfgFilePath, content)
	if _, err := run.RunCli(hostName, cli, logger); err != nil {
		return fmt.Errorf("%s: creating kernel module file %s > %w", hostName, cfgFilePath, err)
	}

	// handle success
	// logger.Debugf("%s > file: %s", hostName, cfgFilePath)
	// logger.Debugf("%s > slice: %v", hostName, s.ModuleSlice)
	logger.Debugf("%s > configured the loading of Kernel module parameters after a reboot", hostName)
	return nil

}

// linux cli to check the value of a kernel module parameter
// for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "sudo sysctl -n net.ipv4.ip_forward"; done

// linux cli to check if module is configure to be loaded after reboot
// for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "cat /etc/sysctl.d/99*"; done
