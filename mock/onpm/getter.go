package onpm

import (
	// _ "embed"
	"sync"

	"fmt"

	"github.com/abtransitionit/gocore/filex"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/mock/yamlx"
)

// -----------------------------------------
// ------ define caching parameters --------
// -----------------------------------------

// Description: used to cache the resolved YAML file for the same couple (osFamily, osDistro)
var (
	configMgrCache = make(map[string]*ManagerConfig)
	cacheMgrMu     sync.Mutex
)

// Description: used to cache the resolved YAML file for the same couple (osFamily, osDistro)
var (
	configRepoCache = make(map[string]*RepoConfig)
	cacheRepoMu     sync.Mutex
)

// -----------------------------------------
// ------ get manager ----------------------
// -----------------------------------------

// Description: returns a system manager that implement the SysCli interface
//
// Notes:
//
// - the manager has access to only to a part of the YAML (that is related to him)
func GetSysMgr(osFamily, osDistro string) (SysCli, error) {
	// 1 - get local auto cached (embedded) file into a struct
	cfg, err := filex.LoadYamlIntoStruct[ManagerConfig](yamlCfgGlobal)
	if err != nil {
		return nil, err
	}

	// 2 - retrun the package manager
	switch osFamily {
	case "debian":
		return &AptSysManager{Cfg: cfg.Apt}, nil
	case "rhel", "fedora":
		return &DnfSysManager{Cfg: cfg.Dnf}, nil
	default:
		return nil, fmt.Errorf("unsupported OS family: %s", osFamily)
	}
}

// Description: returns a package manager that implement the PackageCli interface
//
// Notes:
// - the manager has access to only to a part of the YAML (that is related to him)
func GetPkgMgr(osFamily, osDistro string) (PackageCli, error) {

	// 1 - get local auto cached (embedded) file into a struct
	cfg, err := filex.LoadYamlIntoStruct[ManagerConfig](yamlCfgGlobal)
	if err != nil {
		return nil, err
	}

	// 2 - retrun the package manager
	switch osFamily {
	case "debian":
		return &AptPkgManager{Cfg: cfg.Apt}, nil
	case "rhel", "fedora":
		return &DnfPkgManager{Cfg: cfg.Dnf}, nil
	default:
		return nil, fmt.Errorf("unsupported OS family : %s", osFamily)
	}
}

// Description: returns a repository manager that implement the RepoCli interface
//
// Notes:
// - the manager has access to only to a part of the YAML (that is related to him)
func GetRepoMgr(osFamily, osDistro string, logger logx.Logger) (RepoCli, error) {
	// 1 - get local auto cached (embedded) file into a struct
	cfg, err := filex.LoadYamlIntoStruct[ManagerConfig](yamlCfgGlobal)
	if err != nil {
		return nil, err
	}

	// 2 - retrun the package manager
	switch osFamily {
	case "debian":
		// logger.Debugf("%s:%s: cfg.apt is %v", osFamily, osDistro, cfg.Apt)
		return &AptRepoManager{Cfg: cfg.Apt}, nil
	case "rhel", "fedora":
		// logger.Debugf("%s:%s: cfg.dnf is %v", osFamily, osDistro, cfg.Dnf)
		return &DnfRepoManager{Cfg: cfg.Dnf}, nil
	default:
		return nil, fmt.Errorf("unsupported OS family: %s", osFamily)
	}
}

// -----------------------------------------
// ------ get YAML file --------------------
// -----------------------------------------

// Description: returns the YAML repository db
//
// Notes:
// - The placeholders {{ .XXX }} in the YAML are reolved
// - TODO: find a better structured data or solution for the resolution
func getRepoConfig(repoVersion, pkgType, gpgUrlExt, osDistro string) (*RepoConfig, error) {
	cacheKey := fmt.Sprintf("%s-%s", osDistro)

	cacheRepoMu.Lock()
	theYaml, found := configRepoCache[cacheKey]
	cacheRepoMu.Unlock()

	if found {
		return theYaml, nil
	}

	// Not in cache → resolve YAML
	// 1. define the data structure to resolve the YAML placeholders
	varPlaceholder := map[string]map[string]string{
		"Repo": {
			"Tag": repoVersion,
			"Pkg": pkgType,
			"Gpg": gpgUrlExt,
		},
	}
	// 2 - get resolved YAML into struct
	theYaml, err := yamlx.LoadTplYamlFileEmbed[RepoConfig](yamlRepoList, varPlaceholder)
	if err != nil {
		return nil, fmt.Errorf("getting YAML config file in package: %w", err)
	}

	// Store in cache
	cacheRepoMu.Lock()
	configRepoCache[cacheKey] = theYaml
	cacheRepoMu.Unlock()

	return theYaml, nil
}

// ####### of Pkg List #######

func getPkgList() (*PkgYamlList, error) {
	theYaml, err := filex.LoadYamlIntoStruct[PkgYamlList](yamlPkgList)
	if err != nil {
		return nil, fmt.Errorf("getting YAML config file in package: %w", err)
	}
	return theYaml, nil
}

// ####### of Repo Content #######

func getRepoContentConfig(repoName, repoUrl, gpgUrl, gpgFilepath string) (*RepoContentConfig, error) {

	// 1. define the data structure to resolve the YAML placeholders
	varPlaceholder := map[string]map[string]string{
		"Repo": {
			"Name": repoName,
			"Url":  repoUrl,
		},
		"Gpg": {
			"Filepath": gpgFilepath,
			"Url":      gpgUrl,
		},
	}
	// 2 - get resolved YAML into struct
	theYaml, err := yamlx.LoadTplYamlFileEmbed[RepoContentConfig](yamlCfgRepo, varPlaceholder)
	if err != nil {
		return nil, fmt.Errorf("getting YAML config file in package: %w", err)
	}
	return theYaml, nil
}
