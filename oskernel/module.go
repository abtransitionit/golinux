package oskernel

import (
	"path/filepath"
	"strings"
)

func LoadOsKModule(listKModule []string) (string, error) {

	// create string from slice
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
