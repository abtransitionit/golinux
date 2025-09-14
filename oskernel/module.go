package oskernel

import (
	"path/filepath"
	"strings"
)

func LoadOsKModule(listKModule []string) (string, error) {

	// create string from slice
	// stringContent := list.GetStringWithSepFromSlice(listKModule, " ")
	var cliTmp strings.Builder
	for _, mod := range listKModule {
		cliTmp.WriteString("sudo modprobe ")
		cliTmp.WriteString(mod)
		cliTmp.WriteString("\n")
	}

	var cmds = []string{
		cliTmp.String(),
	}
	cli := strings.Join(cmds, " && ")
	return cli, nil
}

func GetKModuleFilePath(kernelFilename string) string {
	return filepath.Join(MapOsKernelReference["KModuleFolder"], kernelFilename)
}

// apply changes at runtime for each module => cli = fmt.Sprintf(`sudo modprobe %s`, module)

// // loop over each cli
// for _, osKModule := range listOsKModule {

// 	// Get the CLI to activate one OS core module
// 	cli, err := oskernel.LoadOsKModule(osKModule)
// 	if err != nil {
// 		return "", err
// 	}

// 	// // play the CLI
// 	logger.Debugf("%s:%s loadding OS kernel module: %s,%s", vmName, osKModule, cli)
// 	// _, err = run.RunOnVm(vmName, cli)
// 	// if err != nil {
// 	// 	return "", err
// 	// }

// }
