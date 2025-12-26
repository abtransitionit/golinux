package oskernel

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

// description: loads a list of kernel module
func (s ModuleSet) Load(hostName string, logger logx.Logger) (string, error) {

	// 3 - Load
	// 31 - session load - get CLI
	_, err := s.LoadForSession(hostName, logger)
	if err != nil {
		return "", fmt.Errorf("loading Kernel modules for the current session: %w", err)
	}
	// play CLI

	// 32 - startup load
	_, err = s.LoadAfterReboot(hostName, logger)
	if err != nil {
		return "", fmt.Errorf("setting Kernel modules loading at startup: %w", err)
	}

	// handle success
	return "", nil

}

// description: loads a list of kernel module for the current session (at runtime)
func (s ModuleSet) LoadForSession(hostName string, logger logx.Logger) (string, error) {

	// 1 - check if the slice is not empty
	if len(s.ModuleSlice) == 0 {
		return "", fmt.Errorf("%s > Module list is empty", hostName)
	}
	// 2 - create string from slice
	var cliTmp strings.Builder
	for _, item := range s.ModuleSlice {
		cliTmp.WriteString("sudo modprobe ")
		cliTmp.WriteString(item.Name)
		cliTmp.WriteString("\n")
	}
	var cmds = []string{
		cliTmp.String(),
	}
	cli := strings.Join(cmds, " && ")

	// log
	logger.Debugf("%s > loading Kernel modules for current session", hostName)

	// handle success
	return cli, nil
}

// description: loads a list of kernel module after a host restart (at startup)
func (s ModuleSet) LoadAfterReboot(hostName string, logger logx.Logger) (string, error) {
	// 1 - get global kernel conf
	cfg, err := getKernelConf()
	if err != nil {
		return "", fmt.Errorf("getting kernel conf: %w", err)
	}
	// 2 - define var
	kernelParameterFilePath := filepath.Join(cfg.Conf.Folder.Module, s.CfgFileName)

	// log
	logger.Debugf("%s > configuring the loading of Kernel modules at startup: (%s)", hostName, kernelParameterFilePath)

	// 1 - create file as sudo from a list of module:name
	// 2 - save the file
	// - create content from slice
	// stringContent := list.GetStringWithSepFromSlice(listOsKModule, "\n")
	// - define the kernel file path to write this content
	// filePath := oskernel.GetKModuleFilePath(kernelFilename)
	// - define the cli
	// cli = filex.CreateFileFromStringAsSudo(filePath, stringContent)
	// - play the cli
	// _, err = run.RunCliSsh(vmName, cli)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to play cli %s on vm '%s': %w", cli, vmName, err)
	// }

	// success
	// logger.Debugf("%s: 🅑 persisted kernel module(s) in file : %s", vmName, filePath)
	return "", nil

}

// // description: loads a kernel module
// func (i *Module) Load(hostName string, logger logx.Logger) (string, error) {
// 	// log
// 	logger.Infof("%s > loading Kernel module: %s", hostName, i.Name)

// 	// handle success
// 	return "", nil
// }

// // description: loads a kernel module for the current session (at runtime)
// func (i *Module) LoadForSession(hostName string, logger logx.Logger) (string, error) {

// 	// log
// 	logger.Infof("%s > loading Kernel module for the current session: %s", hostName, i.Name)

// 	// handle success
// 	return "", nil
// }

// description: loads a kernel module after a reboot (at startup)
// func (i *Module) LoadAfterReboot(hostName string, logger logx.Logger) (string, error) {

// 	// 1 - get global kernel conf
// 	cfg, err := getKernelConf()
// 	if err != nil {
// 		return "", fmt.Errorf("getting kernel conf: %w", err)
// 	}
// 	// 2 - define var
// 	cfgFilepath := filepath.Join(cfg.Conf.Folder.Module, i.CfgFileName)

// 	// log
// 	logger.Infof("%s > loading Kernel module after a reboot: %s (%s)", hostName, i.Name, cfgFilepath)

// 	// handle success
// 	return "", nil
// }
