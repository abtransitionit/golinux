package oskernel

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/file"
	"github.com/abtransitionit/golinux/mock/run"
)

// description: loads a list of kernel module
//
// TODO:
//
//   - check module exists on thye OS before modprobe with "modinfo overlay >/dev/null 2>&1 || true"
func (i ModuleSet) Load(hostName string, logger logx.Logger) error {

	// check
	if len(i.ModuleSlice) == 0 {
		return fmt.Errorf("%s > Module list is empty", hostName)
	}

	// 1 - load for session
	if err := i.loadForSession(hostName, logger); err != nil {
		return fmt.Errorf("%s : loading Kernel modules for the current session > %w", hostName, err)
	}

	// 2 - load after a reboot (at startup)
	if err := i.loadAfterReboot(hostName, logger); err != nil {
		return fmt.Errorf("%s : loading Kernel modules for the current session > %w", hostName, err)
	}

	// handle success
	return nil

}

// description: loads a list of kernel module for the current session (at runtime)
func (i ModuleSet) loadForSession(hostName string, logger logx.Logger) error {
	// 1 - get cli and play it
	if _, err := run.RunCli(hostName, i.cliToLoadForSession(), logger); err != nil {
		return err
	}
	// handle success
	logger.Debugf("%s > loaded Kernel modules for current session", hostName)
	return nil
}

func (i ModuleSet) loadAfterReboot(hostName string, logger logx.Logger) error {
	// 1 - get yaml that contains kernel conf
	cfg, err := getCfg()
	if err != nil {
		return fmt.Errorf("getting kernel conf: %w", err)
	}
	// 2 - define the file that will contain the module list
	cfgFilePath := filepath.Join(cfg.Conf.Folder.Module, i.CfgFileName)

	// 3 - define the file content of the file - 1 module per line
	var b strings.Builder
	for _, m := range i.ModuleSlice {
		b.WriteString(m.Name)
		b.WriteString("\n")
	}
	content := b.String()

	// 4 - create this file as sudo with this content exactly
	cli := file.SudoCreateFileFromStringBase64(cfgFilePath, content)
	if _, err := run.RunCli(hostName, cli, logger); err != nil {
		return fmt.Errorf("%s: creating kernel module file %s > %w", hostName, cfgFilePath, err)
	}

	// handle success
	// logger.Debugf("%s > file: %s", hostName, cfgFilePath)
	// logger.Debugf("%s > slice: %v", hostName, i.ModuleSlice)
	logger.Debugf("%s > configured the loading of Kernel modules after a reboot", hostName)
	return nil

}

// description: loads a list of kernel module for the current session (at runtime)
func (i ModuleSet) cliToLoadForSession() string {
	cmds := make([]string, 0, len(i.ModuleSlice))
	for _, m := range i.ModuleSlice {
		cmds = append(cmds,
			fmt.Sprintf("sudo modprobe %s || true", m.Name),
		)
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

// linux cli to check if module is loaded
// for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "lsmod | sort | grep -w overlay"; done

// linux cli to check if module is configure to be loaded after reboot
// for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "cat /etc/modules-load.d/99*"; done
