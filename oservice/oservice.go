package oservice

import (
	"fmt"
	"strings"
)

func EnableLinger() string {
	return "loginctl enable-linger"
}

// Disable for the current user services to runs after a logout
func DissableLinger() string {
	return "loginctl disable-linger"
}

func StartService(serviceCName string) string {

	var cmds = []string{
		"sudo systemctl daemon-reload",
		fmt.Sprintf("sudo systemctl enable %s", serviceCName),
		fmt.Sprintf("sudo systemctl start %s", serviceCName),
	}
	cli := strings.Join(cmds, " && ")
	return cli

}

func (s OsService) Start() string {
	cmds := []string{
		"sudo systemctl daemon-reload",
		fmt.Sprintf("sudo systemctl enable %s", s.CName),
		fmt.Sprintf("sudo systemctl start %s", s.CName),
	}
	return strings.Join(cmds, " && ")
}
