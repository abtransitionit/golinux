package filex

import (
	"fmt"
)

func GetFileFromUrl(url string, filePath string, isRootPath bool) string{
	sudoPrefix := ""
	if isRootPath {
		sudoPrefix = "sudo "
	}

	var cmds = []string{
  	fmt.Sprintf(`curl -fsSL %s |  %stee %s`, url, sudoPrefix, filePath)
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

func GetGpgFileFromUrlAsSudo(url string, filePath string) string{
	var cmds = []string{
  	fmt.Sprintf(`curl -fsSL %s | gpg --dearmor | sudo tee %s`, url, filePath),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
