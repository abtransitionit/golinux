package osservice

import (
	"github.com/abtransitionit/gocore/logx"
)

// Description: Starts the service for the current session (at runtime)
func (i *Service) Start(hostName string, logger logx.Logger) error {
	// log
	logger.Infof("%s > starting OS service for the current session: %s", hostName, i.Name)

	// // define cli
	// cmds := []string{
	// 	"sudo systemctl daemon-reload",
	// 	fmt.Sprintf("sudo systemctl enable %s", i.Name),
	// 	fmt.Sprintf("sudo systemctl start %s", i.Name),
	// }
	// cli := strings.Join(cmds, " && ")

	return nil
}

// Description: enables a service to start after at host reboot (at startup)
func (i *Service) Enable(hostName string, logger logx.Logger) error {
	// log
	logger.Infof("%s > enabling OS service at startup: %s ", hostName, i.Name)
	// define cli
	// handle success
	return nil
}

// Description: installs a service using its configuration file
func (i *Service) Install(hostName string, logger logx.Logger) error {

	// // define cli
	// var cmds = []string{
	// 	"sudo systemctl daemon-reload",
	// 	fmt.Sprintf("sudo systemctl enable %s", i.Name),
	// 	fmt.Sprintf("sudo systemctl start %s", i.Name),
	// }
	// cli := strings.Join(cmds, " && ")

	// log
	logger.Infof("%s > installing OS service: %s", hostName, i.Name)

	// handle success
	return nil
}
