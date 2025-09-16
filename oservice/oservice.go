package oservice

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/golinux/filex"
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

func (osService OsService) Start(osFamily string) (string, error) {

	// get service canonical name
	osServiceCName, err := OsServiceReference.GetCName(osService)
	if err != nil {
		return "", err
	}

	// logic
	install := false
	switch osServiceCName {
	case "apparmor.service":
		if osFamily == "debian" {
			install = true
		}
	default:
		install = true
	}

	// if nothing to isntall
	if !install {
		return "", nil
	}

	cmds := []string{
		"sudo systemctl daemon-reload",
		fmt.Sprintf("sudo systemctl enable %s", osServiceCName),
		fmt.Sprintf("sudo systemctl start %s", osServiceCName),
	}
	return strings.Join(cmds, " && "), nil
}

func (osService OsService) Install(osFamily string) (string, error) {

	// get service canonical name
	osServiceCName, err := OsServiceReference.GetCName(osService)
	if err != nil {
		return "", err
	}

	// logic
	install := false
	switch osServiceCName {
	case "apparmor.service":
		if osFamily == "debian" {
			install = true
		}
	}

	// if nothing to isntall
	if !install {
		return "", nil
	}

	// return the CLI to create a service file from string
	return filex.CreateFileFromStringAsSudo(osService.Path, osService.Content), nil
}
