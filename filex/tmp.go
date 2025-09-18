package filex

import (
	"strings"
)

func DeleteTmp() string {
	var cmds = []string{
		`sudo delete -rf /tmp/*`,
		`sudo delete -rf /var/tmp/*`,
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
