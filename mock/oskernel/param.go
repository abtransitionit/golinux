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
func (s *ParameterSet) Load(hostName string, logger logx.Logger) error {

	// check
	if len(s.ParameterSlice) == 0 {
		return fmt.Errorf("%s > Parameter list is empty", hostName)
	}

	// 1 - load for session
	if err := s.loadForSession(hostName, logger); err != nil {
		return fmt.Errorf("%s : loading Kernel module parameters for the current session > %w", hostName, err)
	}
	// 2 - load after a reboot (at startup)
	if err := s.loadAfterReboot(hostName, logger); err != nil {
		return fmt.Errorf("%s : loading Kernel module parameters for the current session > %w", hostName, err)
	}

	// // 12 - startup load
	// if err := s.LoadAfterReboot2(hostName, logger); err != nil {
	// 	return fmt.Errorf("setting Kernel module parameters loading at startup: %w", err)
	// }
	// handle success
	return nil
}

// description: loads a list of kernel module parameters for the current session (at runtime)
func (s ParameterSet) loadForSession(hostName string, logger logx.Logger) error {
	// 1 - get cli
	cli := s.cliToLoadForSession()
	// 2 - play cli
	if _, err := run.RunCli(hostName, cli, logger); err != nil {
		return err
	}
	// 3 - log
	logger.Debugf("%s > loaded Kernel module parameters for current session", hostName)
	// handle success
	return nil
}

// description: load a list of kernel module parameter for the current session
func (s ParameterSet) cliToLoadForSession() string {
	var cmds []string
	for _, item := range s.ParameterSlice {
		cmds = append(cmds,
			fmt.Sprintf("sudo sysctl -w %s=%s || true", item.Name, item.Value),
		)
	}

	cli := strings.Join(cmds, " && ")
	return cli
}

// func (s ParameterSet) cliToLoadForSession(hostName string, logger logx.Logger) string {

// 	// 2 - define cli
// 	var cliTmp strings.Builder
// 	for _, item := range s.ParameterSlice {
// 		cliTmp.WriteString("sysctl -w ")
// 		cliTmp.WriteString(item.Name)
// 		cliTmp.WriteString("=")
// 		cliTmp.WriteString(item.Value)
// 		cliTmp.WriteString("\n")
// 	}
// 	var cmds = []string{
// 		cliTmp.String(),
// 	}
// 	cli := strings.Join(cmds, " && ")

// 	// handle success
// 	// logger.Debugf("%s > play cli: %s", hostName, cli)
// 	return cli
// }

func (s ParameterSet) loadAfterReboot(hostName string, logger logx.Logger) error {
	// 1 - get yaml that contains kernel conf
	cfg, err := getCfg()
	if err != nil {
		return fmt.Errorf("getting kernel conf: %w", err)
	}
	// 2 - define the file that will contain the module parameter list
	cfgFilePath := filepath.Join(cfg.Conf.Folder.Param, s.CfgFileName)

	// 3 - define the file content of the file - 1 parameter per line
	var b strings.Builder
	for _, m := range s.ParameterSlice {
		b.WriteString(fmt.Sprintf("%s = %s\n", m.Name, m.Value))
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

// linux cli to check the value of a kernel module parameter
// for a in o1u o2a o3r o4f o5d; do echo $a; ssh "$a" "sudo sysctl -n net.ipv4.ip_forward"; done

// // description: load a kernel module parameter
// func (i *Parameter) Load(hostname string, logger logx.Logger) (string, error) {

// 	// log
// 	logger.Debugf("%s > loading Kernel module parameter %s (%s)", hostname, i.Name, i.Value)

// 	// handle success
// 	return "", nil
// }

// // description: load a kernel module parameter for the current session
// func (i *Parameter) LoadForSession(hostname string, logger logx.Logger) (string, error) {

// 	// log
// 	logger.Debugf("%s > setting Kernel parameter %s to %s in file %s", hostname, i.Name, i.Value)

// 	// handle success
// 	return "", nil
// }
