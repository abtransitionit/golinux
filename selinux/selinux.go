package selinux

import (
	"fmt"
	"strings"
)

func ConfSelinuxAtRuntime() string {
	var cmds = []string{
		"sudo setenforce 0",
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

func ConfSelinuxAtStartup() string {
	var cmds = []string{
		fmt.Sprintf(`sudo sed -i 's/^SELINUX=.*/SELINUX=permissive/' %s`, confSelinuxFilePath),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
