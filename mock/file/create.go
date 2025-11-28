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
