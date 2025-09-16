package oskernel

import (
	"path/filepath"
	"strings"
)

func LoadOsKParam() (string, error) {

	var cmds = []string{
		// if needed sudo sysctl --system -p /etc/sysctl.d/99-kbe.conf
		"sudo sysctl --system",
	}
	cli := strings.Join(cmds, " && ")
	return cli, nil
}

func GetKParamFilePath(kernelFilename string) string {
	return filepath.Join(MapOsKernelReference["KParamFolder"], kernelFilename)
}

func (s SliceOsKParam) GetContent() string {
	var b strings.Builder
	for _, param := range s {
		b.WriteString("# ")
		b.WriteString(param.Description)
		b.WriteString("\n")
		b.WriteString(param.Kvp)
		b.WriteString("\n")
	}
	return b.String()
}
