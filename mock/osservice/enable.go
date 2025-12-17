package osservice

import (
	"github.com/abtransitionit/gocore/logx"
)

// Description: Enables the service to start at VM boot time
func (service Service) Enable(logger logx.Logger) (string, error) {
	// log
	logger.Info("Enabling Service")
	// handle success
	return "", nil
}

// Description: Starts the service immediately for the current session
func (service Service) Start(logger logx.Logger) (string, error) {
	// log
	logger.Info("Startinging Service")
	// handle success
	return "", nil
}
