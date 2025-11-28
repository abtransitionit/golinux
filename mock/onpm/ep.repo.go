package onpm

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/property"
)

// Description: add a native os package repository
func AddRepo(hostName string, repo Repo2, logger logx.Logger) (string, error) {
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
	repoMgr, err := GetRepoMgr(repo, osFamily, osDistro)
	if err != nil {
		return "", err
	}

	// 3 - get CLI
	cli, err := repoMgr.Add(hostName, repo, logger)
	if err != nil {
		return "", err
	}

	// log
	logger.Infof("%s/%s > %s:%s > add repo >  %v", hostName, repo.Name, osFamily, osDistro, cli)

	// // 4 - run CLI
	// out, err := run.RunCli(hostName, cli, logger)
	// if err != nil {
	// 	return "", fmt.Errorf("%s > %s:%s > %w > out:%s", hostName, osFamily, osDistro, err, out)
	// }

	// handle success
	return "", nil
}

// 1 -
