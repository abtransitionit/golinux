package oskernel

import (
	"github.com/abtransitionit/gocore/logx"
)

func (module *Module) Add(hostName string, logger logx.Logger) (string, error) {
	// log
	logger.Infof("%s > Adding Kernel module: %s", hostName, module.Name)
	// handle success
	return "", nil
}
