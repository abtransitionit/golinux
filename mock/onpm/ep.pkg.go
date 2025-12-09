package onpm

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/property"
	"github.com/abtransitionit/golinux/mock/run"
)

// Description: add a set of native os packages
func AddPkg(hostName string, pkgName string, logger logx.Logger) (string, error) {
	var cli, osFamily, osDistro string
	var err error
	// 1 - get host:property
	osFamily, err = property.GetProperty(logger, hostName, "osFamily")
	if err != nil {
		return "", err
	}
	osDistro, err = property.GetProperty(logger, hostName, "osDistro")
	if err != nil {
		return "", err
	}

	// if hostName == "o1u" {
	// 	osFamily = "debian"
	// 	osDistro = "ubuntu"
	// } else if hostName == "o2a" {
	// 	osFamily = "rhel"
	// 	osDistro = "almalinux"
	// } else if hostName == "o3r" {
	// 	osFamily = "rhel"
	// 	osDistro = "rocky"
	// } else if hostName == "o4f" {
	// 	osFamily = "fedora"
	// 	osDistro = "fedora"
	// }

	// 2 - get a manager (dnf or apt)
	pkgMgr, err := GetPkgMgr(osFamily, osDistro)
	if err != nil {
		return "", err
	}

	// 1 - get yaml:resolved organization's repository list
	pkgYamlList, err := getPkgList()
	if err != nil {
		return "", fmt.Errorf("getting YAML repo config file: %w", err)
	}
	// 2 - is there an entry for our package (that denote a different pkg name)
	pkgNameFormal := pkgYamlList.Package[pkgName]
	if pkgNameFormal == "" {
		pkgNameFormal = pkgName
	} else {
		logger.Debugf("%s:%s > package name overridden: %s", hostName, pkgName, pkgNameFormal)
	}

	// 3 - get the cli
	cli, err = pkgMgr.Add(pkgNameFormal, logger)
	if err != nil {
		return "", err
	}
	// log
	logger.Infof("%s > %s:%s > %s", hostName, osFamily, osDistro, cli)

	// play the cli
	out, err := run.RunCli(hostName, cli, logger)
	if err != nil {
		return "", fmt.Errorf("%s > %s:%s > %w > out:%s", hostName, osFamily, osDistro, err, out)
	}
	// handle success
	return "", nil
}
