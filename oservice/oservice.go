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

func StartService(osService OsService) string {
	var cmds = []string{
		"sudo systemctl daemon-reload",
		fmt.Sprintf("sudo systemctl enable %s", osService.Name),
		fmt.Sprintf("sudo systemctl start %s", osService.Name),
	}
	cli := strings.Join(cmds, " && ")
	return cli

}
