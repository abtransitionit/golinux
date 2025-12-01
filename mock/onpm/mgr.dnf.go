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

func (mgr *DnfPkgManager) Add(pkgName string, logger logx.Logger) (string, error) {
	// // 1 - get resolved organization's repository list
	// pkgYamlList, err := getPkgList()
	// if err != nil {
	// 	return "", fmt.Errorf("getting YAML repo config file: %w", err)
	// }
	// // 2 - is there an entry for our package (that denote a different pkg name)
	// pkgName := pkgYamlList.Package[pkg.Name]
	// if pkgName == "" {
	// 	pkgName = pkg.Name
	// } else {
	// 	logger.Debugf("%s/%s > package name overridden: %s", hostName, pkg.Name, pkgName)
	// }
	// 3 - define cli
	var cmds = []string{
		fmt.Sprintf("sudo dnf install -q -y %s > /dev/null", pkgName),
	}
	cli := strings.Join(cmds, " && ")

	// handle success
	return cli, nil
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
	// 11 - get resolved repo filepath
	repoFilepath := filepath.Join(mgr.Cfg.Folder.Repo, repo.Filename+mgr.Cfg.Ext.Repo)
	// 12 - get resolved organization's repository list
	repoYamlList, err := getRepoConfig(repo.Version, mgr.Cfg.Pkg.Type, mgr.Cfg.Ext.Gpg.Url, "rhel")
	if err != nil {
		return "", fmt.Errorf("getting YAML repo config file: %w", err)
	}
	// 13 - get resolved templated repo file content
	repoFileContent, err := getRepoContentConfig(repo.Name, repoYamlList.Repository[repo.Name].Url.Repo, repoYamlList.Repository[repo.Name].Url.Gpg, "")
	if err != nil {
		return "", fmt.Errorf("getting repo file content: %w", err)
	}
	// log
	// logger.Debugf("repo:name >   (%s)   %v", mgr.Cfg.Pkg.Type, repoYamlList.Repository[repo.Name].Name)
	// logger.Debugf("repo:url:repo (%s) > %v", mgr.Cfg.Pkg.Type, repoYamlList.Repository[repo.Name].Url.Repo)
	// logger.Debugf("repo:url:gpg  (%s) > %v", mgr.Cfg.Pkg.Type, repoYamlList.Repository[repo.Name].Url.Gpg)
	// logger.Debugf("repo filecontent : %s", repoFileContent.Dnf)
	logger.Debugf("repo:filepath     > (%s) %s", mgr.Cfg.Pkg.Type, repoFilepath)
	// fmt.Printf("%s", repoFileContent.Dnf)
	// logger.Debugf("TODO: CreateFileFromStringAsSudo for repo file")

	// 2 - create repo file - GPG key url is included as a parameter
	cli := file.SudoCreateFileFromString(repoFilepath+".test", repoFileContent.Dnf)
	_, err = run.RunCli(hostName, cli, logger)
	if err != nil {
		return "", fmt.Errorf("%s creating repo file with cli %s : %w", hostName, cli, err)
	}
	return "", nil
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
	return strings.Join(cmds, " && ")
}

func (mgr *DnfSysManager) Update(hostName string, osDistro string, logger logx.Logger) (string, error) {
	var pkgMgr *DnfPkgManager
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
	logger.Debugf("distro:%s > install package(s): %v", osDistro, pkgToInstall)

	// 3 - install pcakge in the list
	for _, pkg := range pkgToInstall {
		// 3 - get the cli
		cli, err := pkgMgr.Add(pkg, logger)
		if err != nil {
			return "", err
		}
		// play the cli
		out, err := run.RunCli(hostName, cli, logger)
		if err != nil {
			return "", fmt.Errorf("%s > %s > %w > out:%s", hostName, osDistro, err, out)
		}
	}
	// handle success
	return "", nil
}
