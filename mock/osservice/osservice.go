package osservice

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

// Description: Enables the service to start at VM boot time
func (service *Service) Enable(hostName string, logger logx.Logger) (string, error) {
	// log
	logger.Infof("%s > Enabling OS service: %s", hostName, service.Name)
	// handle success
	return "", nil
}

// Description: Starts the service immediately for the current session
func (service *Service) Start(hostName string, logger logx.Logger) (string, error) {
	// log
	logger.Infof("%s > starting OS service: %s", hostName, service.Name)

	cmds := []string{
		"sudo systemctl daemon-reload",
		fmt.Sprintf("sudo systemctl enable %s", service.Name),
		fmt.Sprintf("sudo systemctl start %s", service.Name),
	}
	cli := strings.Join(cmds, " && ")

	// handle success
	return cli, nil
}

// Description: Install the service
func (service *Service) Install(hostName string, logger logx.Logger) (string, error) {
	// log
	logger.Infof("%s > Installing OS service: %s", hostName, service.Name)

	// // define cli
	// var cmds = []string{
	// 	"sudo systemctl daemon-reload",
	// 	fmt.Sprintf("sudo systemctl enable %s", service.Name),
	// 	fmt.Sprintf("sudo systemctl start %s", service.Name),
	// }
	// cli := strings.Join(cmds, " && ")

	// handle success
	return "", nil
}
