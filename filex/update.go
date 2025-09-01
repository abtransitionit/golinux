package filex

import (
	"fmt"
	"strings"
)

// func AddLineToFile(filePath, line string) string {
// 	var cmds = []string{
// 		fmt.Sprintf(`echo %q >> %q`, line, filePath),
// 		fmt.Sprintf(`echo %q`, filePath),
// 	}

// 	cli := strings.Join(cmds, " && ")
// 	return cli
// }

func EnsureLineInFile(filePath, line string) string {
	var cmds = []string{
		fmt.Sprintf(`grep -qxF %q %q || echo %q >> %q`, line, filePath, line, filePath),
	}

	cli := strings.Join(cmds, " && ")
	return cli
}
