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

func (mgr *DnfPkgManager) Add(pkg Pkg2, logger logx.Logger) (string, error) {
	cmds := []string{
		fmt.Sprintf("sudo dnf install -q -y %s > /dev/null", pkg.Name),
	}
	// logger.Infof("pkg is: %s", d.Cfg.Pkg)
	logger.Debugf("yoyo")

	return strings.Join(cmds, " && "), nil
}

func (mgr *DnfPkgManager) Remove() string {
	cli := "dnf remove <pkg>"
	return cli
}

// -----------------------------------------
// ------ implementation - manage repo ------
// -----------------------------------------

func (mgr *DnfRepoManager) List() string {
	// 2 - GetGpgFilePath
	// 3 - GetUrlGpgResolved
	// 4 - GetRepoFileContent
	cli := "dnf list repos"
	return cli
}

func (mgr *DnfRepoManager) Add(repo Repo2, logger logx.Logger) (string, error) {
	// 1 - get variables
	repoFilePath := filepath.Join(mgr.Cfg.Folder.Repo, repo.Filename+mgr.Cfg.Ext.Repo)
	// 11 - get organization's repoditory db (from now a yaml file inside the package)
	repoYamlCfg, err := getRepoConfig(repo.Version, mgr.Cfg.Pkg.Type, mgr.Cfg.Ext.Gpg.Url, "rhel")
	if err != nil {
		return "", fmt.Errorf("getting YAML repo config file: %w", err)
	}
	logger.Debugf("repo:name >   (%s)   %v", mgr.Cfg.Pkg.Type, repoYamlCfg.Repository[repo.Name].Name)
	logger.Debugf("repo:url:repo (%s) > %v", mgr.Cfg.Pkg.Type, repoYamlCfg.Repository[repo.Name].Url.Repo)
	logger.Debugf("repo:url:gpg  (%s) > %v", mgr.Cfg.Pkg.Type, repoYamlCfg.Repository[repo.Name].Url.Gpg)
	logger.Debugf("repo:filepath     > (%s) %s", mgr.Cfg.Pkg.Type, repoFilePath)

	// fmt.Println("2 - GetRepoFileContent")
	// fmt.Println("3 - save the repo file") // CreateFileFromStringAsSudo(repoFilePath, repoFileContent)
	// fmt.Println("4 - GPG key url is included in the repo file and manage internally")
	cli := "dnf config-manager --add-repo <repo>"
	return cli, nil
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
