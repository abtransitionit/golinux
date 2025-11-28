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

func (mgr *AptRepoManager) Add(repo Repo2, logger logx.Logger) (string, error) {
	// 1 - get variables
	repoFilePath := filepath.Join(mgr.Cfg.Folder.Repo, repo.Filename+mgr.Cfg.Ext.Repo)
	gpgFilePath := filepath.Join(mgr.Cfg.Folder.GpgKey, repo.Filename+mgr.Cfg.Ext.Gpg.File)
	// 11 - get organization's repoditory db (from now a yaml file inside the package)
	repoYamlCfg, err := getRepoConfig(repo.Version, mgr.Cfg.Pkg.Type, mgr.Cfg.Ext.Gpg.Url, "rhel")
	if err != nil {
		return "", fmt.Errorf("getting YAML repo config file: %w", err)
	}
	logger.Debugf("repo:name >   (%s)   %v", mgr.Cfg.Pkg.Type, repoYamlCfg.Repository[repo.Name].Name)
	logger.Debugf("repo:url:repo (%s) > %v", mgr.Cfg.Pkg.Type, repoYamlCfg.Repository[repo.Name].Url.Repo)
	logger.Debugf("repo:url:gpg  (%s) > %v", mgr.Cfg.Pkg.Type, repoYamlCfg.Repository[repo.Name].Url.Gpg)
	logger.Debugf("repo:filepath     > (%s) %s", mgr.Cfg.Pkg.Type, repoFilePath)
	logger.Debugf("repo:gpg:filepath > (%s) %s", mgr.Cfg.Pkg.Type, gpgFilePath)
	logger.Debugf("TODO: CreateGpgFileFromUrlAsSudo")

	// 2 - save gpg key to destination file
	// cli = filex.CreateGpgFileFromUrlAsSudo(urlGpgResolved, GpgFilePath)

	// fmt.Println("2 - GetRepoFileContent")
	// fmt.Println("3 - save the repo file") // CreateFileFromStringAsSudo(repoFilePath, repoFileContent)
	// fmt.Println("4 - manage GPG key:  download GPG key - only for debian. For rhel: gpg key url is included in the repo file and manage internally")
	// download the key to dst file

	// _, err = run.RunCliSsh(vmName, cli)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to play cli %s on vm '%s': %w", cli, vmName, err)
	// }

	cli := "add-apt-repo <repo>"
	return cli, nil
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
