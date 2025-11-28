package file

import (
	"fmt"
	"strings"
)

// description: creates a file from the specified multiline string in a sudo location
func SudoCreateFileFromString(filePath, content string) string {
	// escapedContent := strings.ReplaceAll(content, `'`, `'\''`)
	// Wrap in single quotes and use printf to preserve newlines
	var cmds = []string{
		fmt.Sprintf("printf '%s' | sudo tee %q > /dev/null", content, filePath),
	}
	cli := strings.Join(cmds, " && ")

	return cli
}

func SudoCreateGpgFileFromUrl(url string, filePath string) string {
	var cmds = []string{
		// fmt.Sprintf(`sudo install -d -m 0755  $(dirname %s)`, filePath),
		"set -o pipefail",
		fmt.Sprintf(`curl -fsSL %s | gpg --dearmor | sudo tee %s`, url, filePath),
		fmt.Sprintf(`sudo chmod 0644  %s`, filePath),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
