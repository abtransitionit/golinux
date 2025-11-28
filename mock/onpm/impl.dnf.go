package onpm

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/file"
	"github.com/abtransitionit/golinux/mock/run"
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

func (mgr *DnfRepoManager) Add(hostName string, repo Repo2, logger logx.Logger) (string, error) {
	// 1 - get variables
	// 11 - get resolved repo:filepath
	repoFilePath := filepath.Join(mgr.Cfg.Folder.Repo, repo.Filename+mgr.Cfg.Ext.Repo)
	// 12 - get resolved organization:repo:list
	repoYamlCfg, err := getRepoConfig(repo.Version, mgr.Cfg.Pkg.Type, mgr.Cfg.Ext.Gpg.Url, "rhel")
	if err != nil {
		return "", fmt.Errorf("getting YAML repo config file: %w", err)
	}
	// 13 - get repo file content with placeholders resolved
	repoFileContent, err := getRepoContentConfig(repo.Name, repoYamlCfg.Repository[repo.Name].Url.Repo, repoYamlCfg.Repository[repo.Name].Url.Gpg, "")
	if err != nil {
		return "", fmt.Errorf("getting repo file content: %w", err)
	}
	// log
	// logger.Debugf("repo:name >   (%s)   %v", mgr.Cfg.Pkg.Type, repoYamlCfg.Repository[repo.Name].Name)
	// logger.Debugf("repo:url:repo (%s) > %v", mgr.Cfg.Pkg.Type, repoYamlCfg.Repository[repo.Name].Url.Repo)
	// logger.Debugf("repo:url:gpg  (%s) > %v", mgr.Cfg.Pkg.Type, repoYamlCfg.Repository[repo.Name].Url.Gpg)
	// logger.Debugf("repo filecontent : %s", repoFileContent.Dnf)
	logger.Debugf("repo:filepath     > (%s) %s", mgr.Cfg.Pkg.Type, repoFilePath)
	// fmt.Printf("%s", repoFileContent.Dnf)
	// logger.Debugf("TODO: CreateFileFromStringAsSudo for repo file")

	// 2 - save content to file - GPG key url is included in the repo file
	logger.Debugf("DOING: CreateFileFromStringAsSudo for repo file")
	cli := file.SudoCreateFileFromString("/usr/local/bin/mxtest", repoFileContent.Dnf)
	_, err = run.RunCli(hostName, cli, logger)
	if err != nil {
		return "", fmt.Errorf("%s creating repo file with cli %s : %w", hostName, cli, err)
	}

	// other
	cli = "dnf config-manager --add-repo <repo>"
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
