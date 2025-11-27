package onpm

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

// -----------------------------------------
// ------ implementation - manage pkg ------
// -----------------------------------------

func (mgr *DnfPkgManager) List() string {
	cli := "dnf list installed"
	return cli
}

func (mgr *DnfPkgManager) Add(pkg Pkg2, logger logx.Logger) string {
	cmds := []string{
		fmt.Sprintf("sudo dnf install -q -y %s > /dev/null", pkg.Name),
	}
	// logger.Infof("pkg is: %s", d.Cfg.Pkg)
	logger.Debugf("yoyo")

	return strings.Join(cmds, " && ")
}

func (mgr *DnfPkgManager) Remove() string {
	cli := "dnf remove <pkg>"
	return cli
}

// -----------------------------------------
// ------ implementation - manage repo ------
// -----------------------------------------

func (mgr *DnfRepoManager) List() string {
	// 1 - GetRepoFilePath
	// 2 - GetGpgFilePath
	// 3 - GetUrlGpgResolved
	// 4 - GetRepoFileContent
	cli := "dnf list repos"
	return cli
}

func (mgr *DnfRepoManager) Add(repo Repo2, logger logx.Logger) string {
	// 1 - get os:repository file path
	repoFilePath := filepath.Join(mgr.Cfg.Folder.Repo, repo.Filename+mgr.Cfg.Ext)
	logger.Debugf("repo file path > %s", repoFilePath)
	// 1 - get organization's repoditory db
	theYaml, err := getRepoConfig(repo.Version, "rhel")
	if err != nil {
		return ""
	}
	logger.Debugf("repo db yaml > %v", theYaml)

	// fmt.Println("2 - GetRepoFileContent")
	// fmt.Println("3 - save the repo file") // CreateFileFromStringAsSudo(repoFilePath, repoFileContent)
	// fmt.Println("4 - GPG key url is included in the repo file and manage internally")
	cli := "dnf config-manager --add-repo <repo>"
	return cli
}

func (mgr *DnfRepoManager) Remove() string {
	cli := "dnf config-manager --remove-repo <repo>"
	return cli
}

// -----------------------------------------
// ------ implementation - manage sys ------
// -----------------------------------------

func (mgr *DnfSysManager) NeedReboot(logger logx.Logger) string {
	cmds := []string{
		"command -v needs-restarting >/dev/null && needs-restarting -r | grep -q 'Reboot is required' && echo true || echo false",
	}
	return strings.Join(cmds, " && ")
}
func (mgr *DnfSysManager) Upgrade(logger logx.Logger) string {
	cmds := []string{
		"sudo dnf update -q -y",
		"sudo dnf upgrade -q -y",
		"sudo dnf clean all",
	}
	// logger.Infof("pkg is: %s", d.Cfg.Pkg)
	return strings.Join(cmds, " && ")
}

func (mgr *DnfSysManager) Update(logger logx.Logger) string {
	logger.Infof("pkg is: %s", mgr.Cfg.Pkg)

	// cmds := []string{
	// 	"sudo dnf update -q -y",
	// 	"sudo dnf upgrade -q -y",
	// 	"sudo dnf clean all",
	// }
	// return strings.Join(cmds, " && ")
	return ""
}
