package filex

import (
	"fmt"
	"strings"
)

func CreateFileFromUrl(url string, filePath string, isRootPath bool) string {
	sudoPrefix := ""
	if isRootPath {
		sudoPrefix = "sudo "
	}

	var cmds = []string{
		fmt.Sprintf(`curl -fsSL %s |  %stee %s`, url, sudoPrefix, filePath),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

func CreateGpgFileFromUrlAsSudo(url string, filePath string) string {
	var cmds = []string{
		// fmt.Sprintf(`sudo install -d -m 0755  $(dirname %s)`, filePath),
		"set -o pipefail",
		fmt.Sprintf(`curl -fsSL %s | gpg --dearmor | sudo tee %s`, url, filePath),
		fmt.Sprintf(`sudo chmod 0644  %s`, filePath),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
