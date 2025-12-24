package oskernel

import (
	"github.com/abtransitionit/gocore/logx"
)

func (param *Parameter) Add(hostname string, logger logx.Logger) (string, error) {
	// log
	logger.Debugf("%s > Setting Kernel parameter %s to %s", hostname, param.Name, param.Value)
	// handle success
	return "", nil
}
