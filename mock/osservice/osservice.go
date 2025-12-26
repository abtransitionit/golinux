package osservice

import (
	"github.com/abtransitionit/gocore/logx"
)

// Description: Starts the service for the current session (at runtime)
func (i *Service) Start(hostName string, logger logx.Logger) (string, error) {
	// log
	logger.Infof("%s > starting OS service for the current session: %s", hostName, i.Name)

	// cmds := []string{
	// 	"sudo systemctl daemon-reload",
	// 	fmt.Sprintf("sudo systemctl enable %s", i.Name),
	// 	fmt.Sprintf("sudo systemctl start %s", i.Name),
	// }
	// cli := strings.Join(cmds, " && ")

	// handle success
	cli := ""
	return cli, nil
}

// Description: enables a service to start after at host reboot (at startup)
func (i *Service) Enable(hostName string, logger logx.Logger) (string, error) {
	// log
	logger.Infof("%s > enabling OS service after a host reboot: %s ", hostName, i.Name)

	// handle success
	return "", nil
}

// Description: installs a service using its configuration file
func (i *Service) Install(hostName string, logger logx.Logger) (string, error) {
	// log
	logger.Infof("%s > installing OS service: %s", hostName, i.Name)

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
