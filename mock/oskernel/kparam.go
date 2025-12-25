package oskernel

import (
	"fmt"
	"path/filepath"

	"github.com/abtransitionit/gocore/logx"
)

func (param *Parameter) Add(hostname string, logger logx.Logger) (string, error) {

	// 1 - get global kernel conf
	cfg, err := getKernelConf()
	if err != nil {
		return "", fmt.Errorf("getting kernel conf: %w", err)
	}
	// 2 - define var
	cfgFilepath := filepath.Join(cfg.Conf.Folder.Module, param.CfgFilePath)

	// log
	logger.Debugf("%s > Setting Kernel parameter %s to %s in file %s", hostname, param.Name, param.Value, cfgFilepath)

	// handle success
	return "", nil
}
