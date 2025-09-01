package filex

import (
	"fmt"
	"strings"
)

func CreateFileFromStringAsSudo(filePath, content string) string {
	var cmds = []string{
		fmt.Sprintf(`cleaned_content=$(echo %q | sed 's/^[ \t]*//')`, content),
		fmt.Sprintf(`echo "$cleaned_content" | sudo tee %q > /dev/null`, filePath),
	}

	cli := strings.Join(cmds, " && ")
	return cli
}

func CreateTmpFileFromString(filePath, content string) string {
	var cmds = []string{
		"tmpfile=$(mktemp)",
		fmt.Sprintf(`echo %q | tee $tmpfile > /dev/null`, content),
		"echo $tmpfile",
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
