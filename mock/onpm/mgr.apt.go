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

func (mgr *AptPkgManager) Add(pkgName string, logger logx.Logger) (string, error) {
	// 1 - is there an entry for this package in the organization's package list
	// 21 - get the list
	pkgYamlList, err := getPkgList()
	if err != nil {
		return "", fmt.Errorf("getting YAML repo config file: %w", err)
	}
	// 22 - check pkg exist in the list
	pkgNameFormal := pkgYamlList.Package[pkgName]
	if pkgNameFormal == "" {
		pkgNameFormal = pkgName
	} else {
		logger.Debugf("package %s > overridden: %s", pkgName, pkgNameFormal)
	}
	// 3 - define cli
	var cmds = []string{
		fmt.Sprintf("DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' install -qq -y %s > /dev/null", pkgName),
	}
	cli := strings.Join(cmds, " && ")

	// handle success
	return cli, nil
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

	return strings.Join(cmds, " && ")
}

func (mgr *AptSysManager) Upgrade(logger logx.Logger) string {
	cmds := []string{
		"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' update -qq -y",
		"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' upgrade -qq -y",
		"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' clean -qq",
	}

	return strings.Join(cmds, " && ")
}
func (mgr *AptSysManager) Update(hostName string, osDistro string, logger logx.Logger) (string, error) {
	var pkgMgr *AptPkgManager
	// 1 - get the section named required of the yaml:manager
	required := mgr.Cfg.Pkg.Required
	// logger.Debugf("distro = %s required: %v", osDistro, required)

	// 1 - loop over the map to get all the pkg to install
	var pkgToInstall []string
	for key, pkgs := range required {
		// 11 - if this key exists: add all the pkg to the list
		if key == "all" {
			pkgToInstall = append(pkgToInstall, pkgs...)
			continue
		}
		// 12 - if this key exists => add all the pkg to the list
		if key == osDistro {
			pkgToInstall = append(pkgToInstall, pkgs...)
		}
	}

	// 2 - exit if no pkg to install
	if len(pkgToInstall) == 0 {
		return "", nil
	}
	// log
	logger.Debugf("%s:%s > installing package(s): %v", hostName, osDistro, pkgToInstall)

	// 3 - install pcakge in the list
	for _, pkgName := range pkgToInstall {
		// 3 - get the cli
		cli, err := pkgMgr.Add(pkgName, logger)
		if err != nil {
			return "", fmt.Errorf("❌ %s:%s > installing package : %s with cli : %s : %w", hostName, osDistro, pkgName, cli, err)
		}
		// play the cli
		out, err := run.RunCli(hostName, cli, logger)
		if err != nil {
			return "", fmt.Errorf("❌ %s:%s > %w > out:%s", hostName, osDistro, err, out)
		}
		// log
		logger.Debugf("%s:%s > installed package: %s > out:%s", hostName, osDistro, pkgName, out)
	}
	// handle success
	return "", nil
}

// cmds := []string{
// 	"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' update -qq -y",
// 	"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' upgrade -qq -y",
// 	"DEBIAN_FRONTEND=noninteractive sudo apt-get -o Dpkg::Options::='--force-confdef' -o Dpkg::Options::='--force-confold' clean -qq",
// }
// // logger.Infof("pkg is: %s", d.Cfg.Pkg)

// return strings.Join(cmds, " && ")
