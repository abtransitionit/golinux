package onpm

import (
	"github.com/abtransitionit/gocore/logx"
)

// Description: add a set of native os packages
func AddPkg(hostName string, pkg Pkg2, logger logx.Logger) (string, error) {
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
	pkgMgr, err := GetPkgMgr(osFamily, osDistro)
	if err != nil {
		return "", err
	}

	// 3 - get CLI
	cli := pkgMgr.Add(pkg)

	// log
	logger.Infof("%s > %s:%s > will install package with %v", hostName, osFamily, osDistro, cli)

	// // 4 - run CLI
	// out, err := run.RunCli(hostName, cli, logger)
	// if err != nil {
	// 	return "", fmt.Errorf("%s > %s:%s > %w > out:%s", hostName, osFamily, osDistro, err, out)
	// }

	// handle success
	return "", nil
}
