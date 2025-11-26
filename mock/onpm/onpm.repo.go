package onpm

import (
	"github.com/abtransitionit/gocore/logx"
)

// Description: add a native os package repository
func AddRepo(hostName string, repo Repo2, logger logx.Logger) (string, error) {
	// // 1 - get host:property
	// osFamily, err := property.GetProperty(logger, hostName, "osFamily")
	// if err != nil {
	// 	return "", err
	// }
	// osDistro, err := property.GetProperty(logger, hostName, "osDistro")
	// if err != nil {
	// 	return "", err
	// }
	osFamily := "rhel"
	osDistro := "rhel"

	// 2 - get a system manager
	repoMgr, err := GetRepoMgr(osFamily, osDistro)
	if err != nil {
		return "", err
	}

	// 3 - get CLI
	cli := repoMgr.Add(repo)

	// log
	logger.Infof("%s/%s > %s:%s > will install repository with %v", hostName, repo.Name, osFamily, osDistro, cli)

	// // 4 - run CLI
	// out, err := run.RunCli(hostName, cli, logger)
	// if err != nil {
	// 	return "", fmt.Errorf("%s > %s:%s > %w > out:%s", hostName, osFamily, osDistro, err, out)
	// }

	// handle success
	return "", nil
}
