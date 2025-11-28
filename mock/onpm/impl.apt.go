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

func (mgr *AptPkgManager) List() string {
	cli := "apt list --installed"
	return cli
}

func (mgr *AptPkgManager) Add(pkg Pkg2, logger logx.Logger) (string, error) {
	cmds := []string{
		fmt.Sprintf("DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' install -qq -y %s > /dev/null", pkg.Name),
	}
	// logger.Infof("pkg is: %s", d.Cfg.Pkg)

	return strings.Join(cmds, " && "), nil
}

func (mgr *AptPkgManager) Remove() string {
	cli := "apt remove <pkg>"
	return cli
}

// -----------------------------------------
// ------ implementation - manage repo -----
// -----------------------------------------

func (mgr *AptRepoManager) List() string {
	cli := "apt list repos"
	return cli
}

func (mgr *AptRepoManager) Add(hostName string, repo Repo2, logger logx.Logger) (string, error) {
	var cli string
	// 1 - get variables
	// 11 - get resolved repo filepath
	repoFilepath := filepath.Join(mgr.Cfg.Folder.Repo, repo.Filename+mgr.Cfg.Ext.Repo)
	// 12 - get resolved organization's repository list
	repoYamlCfg, err := getRepoConfig(repo.Version, mgr.Cfg.Pkg.Type, mgr.Cfg.Ext.Gpg.Url, "rhel")
	if err != nil {
		return "", fmt.Errorf("getting YAML repo config file: %w", err)
	}
	// 13 - get resolved gpg filepath
	gpgFilepath := filepath.Join(mgr.Cfg.Folder.GpgKey, repo.Filename+mgr.Cfg.Ext.Gpg.File)
	// 14 - get resolved templated repo file content
	repoFileContent, err := getRepoContentConfig(repo.Name, repoYamlCfg.Repository[repo.Name].Url.Repo, repoYamlCfg.Repository[repo.Name].Url.Gpg, gpgFilepath)
	if err != nil {
		return "", fmt.Errorf("getting repo file content: %w", err)
	}
	// log
	// logger.Debugf("repo:name >   (%s)   %v", mgr.Cfg.Pkg.Type, repoYamlCfg.Repository[repo.Name].Name)
	// logger.Debugf("repo:url:repo (%s) > %v", mgr.Cfg.Pkg.Type, repoYamlCfg.Repository[repo.Name].Url.Repo)
	// logger.Debugf("repo:url:gpg  (%s) > %v", mgr.Cfg.Pkg.Type, repoYamlCfg.Repository[repo.Name].Url.Gpg)
	logger.Debugf("repo:filepath     > (%s) %s", mgr.Cfg.Pkg.Type, repoFilepath)
	logger.Debugf("repo:gpg:filepath > (%s) %s", mgr.Cfg.Pkg.Type, gpgFilepath)
	// logger.Debugf("repo file content : %s", repoFileContent.Apt)
	fmt.Printf("%s", repoFileContent.Apt)

	// 2 - create repo file - GPG key is saved separately
	cli = file.SudoCreateFileFromString(repoFilepath+".test", repoFileContent.Apt)
	_, err = run.RunCli(hostName, cli, logger)
	if err != nil {
		return "", fmt.Errorf("%s creating repo file with cli %s : %w", hostName, cli, err)
	}
	// 2 - create gpg file
	logger.Debugf("DOING: saving gpg key to /tmp/toto")
	cli = file.SudoCreateGpgFileFromUrl(repoYamlCfg.Repository[repo.Name].Url.Gpg, gpgFilepath+".test")
	// cli = file.SudoCreateGpgFileFromUrl(repoYamlCfg.Repository[repo.Name].Url.Repo, gpgFilepath)
	_, err = run.RunCli(hostName, cli, logger)
	if err != nil {
		return "", fmt.Errorf("%s creating repo file with cli %s : %w", hostName, cli, err)
	}

	return "", nil
}

func (mgr *AptRepoManager) Remove() string {
	cli := "remove-apt-repo <repo>"
	return cli
}

// -----------------------------------------
// ------ implementation - manage sys ------
// -----------------------------------------

func (mgr *AptSysManager) NeedReboot(logger logx.Logger) string {
	cmds := []string{
		"test -f /var/run/reboot-required && echo true || echo false",
	}
	// logger.Infof("pkg is: %s", d.Cfg.Pkg)

	return strings.Join(cmds, " && ")
}

func (mgr *AptSysManager) Upgrade(logger logx.Logger) string {
	cmds := []string{
		"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' update -qq -y",
		"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' upgrade -qq -y",
		"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' clean -qq",
	}
	// logger.Infof("pkg is: %s", d.Cfg.Pkg)

	return strings.Join(cmds, " && ")
}
func (mgr *AptSysManager) Update(logger logx.Logger) string {
	logger.Infof("pkg is: %s", mgr.Cfg.Pkg)
	// cmds := []string{
	// 	"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' update -qq -y",
	// 	"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' upgrade -qq -y",
	// 	"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' clean -qq",
	// }
	// // logger.Infof("pkg is: %s", d.Cfg.Pkg)

	// return strings.Join(cmds, " && ")
	return ""
}
