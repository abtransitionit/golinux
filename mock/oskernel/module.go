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
func (s ModuleSet) Load(hostName string, logger logx.Logger) error {

	// check
	if len(s.ModuleSlice) == 0 {
		return fmt.Errorf("%s > Module list is empty", hostName)
	}

	// 1 - load for session
	if err := s.loadForSession(hostName, logger); err != nil {
		return fmt.Errorf("%s : loading Kernel modules for the current session > %w", hostName, err)
	}

	// 2 - load after a reboot (at startup)
	if err := s.loadAfterReboot(hostName, logger); err != nil {
		return fmt.Errorf("%s : loading Kernel modules for the current session > %w", hostName, err)
	}

	// handle success
	return nil

}

// description: loads a list of kernel module for the current session (at runtime)
func (s ModuleSet) loadForSession(hostName string, logger logx.Logger) error {
	// 1 - get cli
	cli := s.cliToLoadForSession()
	// 2 - play cli
	if _, err := run.RunCli(hostName, cli, logger); err != nil {
		return err
	}
	// 3 - log
	logger.Debugf("%s > loaded Kernel modules for current session", hostName)
	// handle success
	return nil
}

func (s ModuleSet) loadAfterReboot(hostName string, logger logx.Logger) error {
	// 1 - get yaml that contains kernel conf
	cfg, err := getCfg()
	if err != nil {
		return fmt.Errorf("getting kernel conf: %w", err)
	}
	// 2 - define the file that will contain the module list
	cfgFilePath := filepath.Join(cfg.Conf.Folder.Module, s.CfgFileName)

	// 3 - define the file content of the file - 1 module per line
	var b strings.Builder
	for _, m := range s.ModuleSlice {
		b.WriteString(m.Name)
		b.WriteString("\n")
	}
	content := b.String()

	// 4 - create this file as sudo with this content exactly
	cli := file.SudoCreateFileFromStringBase64(cfgFilePath, content)
	if _, err := run.RunCli(hostName, cli, logger); err != nil {
		return fmt.Errorf("%s: creating kernel module file %s > %w", hostName, cfgFilePath, err)
	}

	// log
	// logger.Debugf("%s > file: %s", hostName, kernelModuleFilePath)
	// logger.Debugf("%s > slice: %v", hostName, s.ModuleSlice)

	// handle success
	logger.Debugf("%s > configured the loading of Kernel modules after a reboot", hostName)
	return nil

}

// description: loads a list of kernel module for the current session (at runtime)
func (s ModuleSet) cliToLoadForSession() string {

	// 1 - create cli
	cmds := make([]string, 0, len(s.ModuleSlice))
	for _, m := range s.ModuleSlice {
		cmds = append(cmds,
			fmt.Sprintf("sudo modprobe %s || true", m.Name),
		)
	}

	// handle success
	cli := strings.Join(cmds, " && ")
	return cli
}

// linux cli to check if module is loaded
// for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "lsmod | sort | grep -w overlay"; done

// linux cli to check if module is configure to be loaded after reboot
// for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "cat /etc/modules-load.d/99*"; done
