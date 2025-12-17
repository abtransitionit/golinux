package oskernel

import (
	"github.com/abtransitionit/gocore/logx"
)

func (param Parameter) Add(logger logx.Logger) (string, error) {
	// log
	logger.Info("Adding Kernel parameter")
	// handle success
	return "", nil
}

// func AddKParam(phaseName, hostName string, paramList [][]any, logger logx.Logger) (bool, error) {
// }
