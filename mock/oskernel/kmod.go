package oskernel

import (
	"fmt"
	"path/filepath"

	"github.com/abtransitionit/gocore/logx"
)

func (module *Module) Add(hostName string, logger logx.Logger) (string, error) {

	// 1 - get global kernel conf
	cfg, err := getKernelConf()
	if err != nil {
		return "", fmt.Errorf("getting kernel conf: %w", err)
	}
	// 2 - define var
	cfgFilepath := filepath.Join(cfg.Conf.Folder.Module, module.CfgFilePath)

	// log
	logger.Infof("%s > Adding Kernel module: %s to conf file: %s", hostName, module.Name, cfgFilepath)

	// handle success
	return "", nil
}
