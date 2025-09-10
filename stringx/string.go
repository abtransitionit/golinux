package stringx

import (
	"fmt"
	"strings"
)

func EnsureFusionStringUniq(s1, s2 string, sep string) string {
	var cmds = []string{
		fmt.Sprintf(`echo "%s%s%s" | tr ':' '\n' | sort -u | paste -sd ':'`, s1, sep, s2),
	}

	cli := strings.Join(cmds, " && ")
	return cli
}

func SanitizeStringBeforeSaving(s string) string {
	var cmds = []string{
		fmt.Sprintf(`echo %q`, s),
	}

	cli := strings.Join(cmds, " && ")
	return cli
}
