package oskernel

import (
	"github.com/abtransitionit/gocore/logx"
)

func (module Module) Add(logger logx.Logger) (string, error) {
	// log
	logger.Infof("Adding Kernel module: %s", module.Name)
	// handle success
	return "", nil
}

func AddKModule(phaseName, hostName string, paramList [][]any, logger logx.Logger) (bool, error) {
	// log
	logger.Info("AddKModule called with param:")
	// handle success
	return true, nil
}
