package util

import (
	"fmt"
	"strings"
)

func CliToFusionString(s1, s2 string, sep string) string {
	var cmds = []string{
		// fmt.Sprintf(`echo "%s%s%s" | tr ':' '\n' | sort -u | paste -sd ':'`, s1, sep, s2),
		// fmt.Sprintf(`echo "%s%s%s" | tr '%s' '\n' | awk '!seen[$0]++' | paste -sd '%s'`,s1, sep, s2, sep, sep),
		fmt.Sprintf(`echo %s%s%s`, s1, sep, s2),
	}

	cli := strings.Join(cmds, " && ")
	return cli
}
func CliToUniqSequence(sequence string, sep string) string {
	var cmds = []string{
		fmt.Sprintf(`for a in %s; do echo \$a; done`, sequence),
	}

	cli := strings.Join(cmds, " && ")
	return cli
}
