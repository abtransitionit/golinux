package onpm

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/property"
	"github.com/abtransitionit/golinux/mock/run"
)

// Description: upgrade the hostname:OS native pkgRepo and Pkg to version latest
func UpgradeOs(hostName string, logger logx.Logger) (string, error) {
	// 1 - get host:property
	osFamily, err := property.GetProperty(logger, hostName, "osFamily")
	if err != nil {
		return "", err
	}
	osDistro, err := property.GetProperty(logger, hostName, "osDistro")
	if err != nil {
		return "", err
	}
	osKVersion, err := property.GetProperty(logger, hostName, "osKernelVersion")
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

	// 2 - get a manager
	sysMgr, err := GetSysMgr(osFamily, osDistro)
	if err != nil {
		return "", err
	}

	// 3 - get CLI
	cli := sysMgr.Upgrade(logger)

	// log
	logger.Infof("%s > %s:%s > %s", hostName, osFamily, osDistro, osKVersion)

	// 4 - run CLI
	out, err := run.RunCli(hostName, cli, logger)
	if err != nil {
		return "", fmt.Errorf("%s > %s:%s > %w > out:%s", hostName, osFamily, osDistro, err, out)
	}

	// handle success
	return "", nil
}

// Description: add the hostname:OS native standard missing/required pkgRepo and Pkg
func UpdateOs(hostName string, logger logx.Logger) (string, error) {
	var osFamily, osDistro string
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
	// osKVersion, err := property.GetProperty(logger, hostName, "osKernelVersion")
	// if err != nil {
	// 	return "", err
	// }

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
	// } else if hostName == "o5d" {
	// 	osFamily = "debian"
	// 	osDistro = "debian"
	// }

	// 2 - get a system manager
	sysMgr, err := GetSysMgr(osFamily, osDistro)
	if err != nil {
		return "", err
	}

	// 3 - do the job
	cli, err := sysMgr.Update(hostName, osDistro, logger)
	if err != nil {
		return "", fmt.Errorf("%s > %s:%s > updating the OS with cli : %s> %v", hostName, osFamily, osDistro, cli, err)
	}

	// log

	// // 4 - run CLI
	// out, err := run.RunCli(hostName, cli, logger)
	// if err != nil {
	// 	return "", fmt.Errorf("%s > %s:%s > %w > out:%s", hostName, osFamily, osDistro, err, out)
	// }

	// handle success
	logger.Infof("%s > %s:%s > updated OS successfully with %s", hostName, osFamily, osDistro, cli)
	return "", nil
}
func NeedReboot(hostName string, logger logx.Logger) (string, error) {
	// 1 - get host:property
	osFamily, err := property.GetProperty(logger, hostName, "osFamily")
	if err != nil {
		return "", err
	}
	osDistro, err := property.GetProperty(logger, hostName, "osDistro")
	if err != nil {
		return "", err
	}
	// 2 - get a system manager
	sysMgr, err := GetSysMgr(osFamily, osDistro)
	if err != nil {
		return "", err
	}

	// 3 - get CLI
	cli := sysMgr.NeedReboot(logger)

	// log
	logger.Infof("%s > %s:%s > %s", hostName, osFamily, osDistro, cli)

	// 4 - run CLI
	out, err := run.RunCli(hostName, cli, logger)
	if err != nil {
		return "", fmt.Errorf("%s > %s:%s > %w > out:%s", hostName, osFamily, osDistro, err, out)
	}

	// handle success
	return out, nil
}
