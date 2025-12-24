package osservice

import (
	"github.com/abtransitionit/gocore/logx"
)

// Description: Enables the service to start at VM boot time
func (service *Service) Enable(logger logx.Logger) (string, error) {
	// log
	logger.Infof("Enabling OS service: %s", service.Name)
	// handle success
	return "", nil
}

// Description: Starts the service immediately for the current session
func (service *Service) Start(logger logx.Logger) (string, error) {
	// log
	logger.Infof("starting OS service: %s", service.Name)
	// handle success
	return "", nil
}

// Description: Install the service
func (service *Service) Install(logger logx.Logger) (string, error) {
	// log
	logger.Infof("Installing OS service: %s", service.Name)
	// handle success
	return "", nil
}
