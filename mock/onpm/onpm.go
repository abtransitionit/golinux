package onpm

import (
	_ "embed"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/property"
)

// //go:embed db.conf.yaml
// var confData []byte

func UpgradeOs(targetName string, logger logx.Logger) (string, error) {
	// 1 - get os:family
	osFamily, err := property.GetProperty(logger, targetName, "osFamily")
	if err != nil {
		return "", err
	}
	// 2 - get os:type
	osDistro, err := property.GetProperty(logger, targetName, "osDistro")
	if err != nil {
		return "", err
	}

	// 3 - get a manager
	sysMgr, err := GetSysMgr(osFamily, osDistro)
	if err != nil {
		return "", err
	}
	// 4 - get CLI
	cli := sysMgr.Upgrade(logger)
	logger.Infof("%s > %s:%s > %s", targetName, osFamily, osDistro, cli)
	// 5 - run CLI

	// handle success
	return "", nil
}
