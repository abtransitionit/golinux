package oskernel

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

// description: load a kernel module parameter
func (s *ParameterSet) Load(hostName string, logger logx.Logger) error {

	// 1 - load
	// 11 - session load
	if err := s.LoadForSession(hostName, logger); err != nil {
		return fmt.Errorf("loading Kernel module parameters for the current session: %w", err)
	}
	// play CLI

	// 12 - startup load
	if err := s.LoadAfterReboot(hostName, logger); err != nil {
		return fmt.Errorf("setting Kernel module parameters loading at startup: %w", err)
	}
	// handle success
	return nil
}

// description: load a list of kernel module parameter for the current session
func (s ParameterSet) LoadForSession(hostName string, logger logx.Logger) error {

	// 1 - check if the slice is not empty
	if len(s.ParameterSlice) == 0 {
		return fmt.Errorf("%s > Module list is empty", hostName)
	}
	// 2 - define cli
	var cliTmp strings.Builder
	for _, item := range s.ParameterSlice {
		cliTmp.WriteString("sysctl -w ")
		cliTmp.WriteString(item.Name)
		cliTmp.WriteString("=")
		cliTmp.WriteString(item.Value)
		cliTmp.WriteString("\n")
	}
	var cmds = []string{
		cliTmp.String(),
	}
	cli := strings.Join(cmds, " && ")

	// play cli
	logger.Debugf("%s > play cli: %s", hostName, cli)

	// log
	logger.Debugf("%s > setting Kernel module parameters for current session", hostName)

	// handle success
	return nil

}

// description: load a list of kernel module parameter after a host restart
func (s ParameterSet) LoadAfterReboot(hostName string, logger logx.Logger) error {

	// 1 - get global kernel conf
	cfg, err := getKernelConf()
	if err != nil {
		return fmt.Errorf("getting kernel conf: %w", err)
	}
	// 2 - define var
	kernelParameterFilePath := filepath.Join(cfg.Conf.Folder.Param, s.CfgFileName)

	// log
	logger.Debugf("%s > configuring loading of Kernel module parameters at startup: (file: %s)", hostName, kernelParameterFilePath)

	// 1 - save the list to a file
	// 2 - save the file
	// save the content to the file
	// logger.Debugf("%s: persisting kernel module(s) to be applyed after rebbot in file : %s", vmName, filePath)
	// - create content from slice
	// stringContent := listOsKParam.GetContent()
	// - define the kernel file path to write this content
	// filePath := oskernel.GetKParamFilePath(kernelFilename)
	// - define the cli
	// cli = filex.CreateFileFromStringAsSudo(filePath, stringContent)
	// - play the cli
	// _, err = run.RunCliSsh(vmName, cli)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to play cli %s on vm '%s': %w", cli, vmName, err)
	// }

	// success
	// logger.Debugf("%s: 🅑 persisted kernel parameter(s) in file : %s", vmName, filePath)

	// handle success
	return nil
}

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
