package onpm

import (
	_ "embed"
	"sync"

	"fmt"

	"github.com/abtransitionit/gocore/mock/yamlx"
)

// -----------------------------------------
// ------ define file location -------------
// -----------------------------------------

//go:embed db.mgr.yaml
var yamlMgr []byte // cache the raw yaml file in this var

//go:embed db.repo.list.yaml
var yamlRepo []byte // cache the raw yaml file in this var

//go:embed db.package.list.yaml
var yamlPkgList []byte // cache the raw yaml file in this var

//go:embed db.repo.content.yaml
var yamlRepoConent []byte // cache the raw yaml file in this var

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

// Description: returns a SysManager based on OS family
//
// Notes:
//
// - the PM has access to the part of the YAML configuration that relates to the package manager
func GetSysMgr(osFamily, osDistro string) (SysCli, error) {
	// 1 - load yaml mgr - Ensure it is loaded only once
	conf, err := getMgrConfig(osFamily, osDistro)
	if err != nil {
		return nil, err
	}

	// 2 - retrun the package manager
	switch osFamily {
	case "debian":
		return &AptSysManager{Cfg: conf.Apt}, nil
	case "rhel", "fedora":
		return &DnfSysManager{Cfg: conf.Dnf}, nil
	default:
		return nil, fmt.Errorf("unsupported OS family: %s", osFamily)
	}
}

// Description: returns a PackageManager based on OS family
//
// Notes:
// - the PM has access to the part of the YAML configuration that relates to the package manager
func GetPkgMgr(osFamily, osDistro string) (PackageCli, error) {

	// 1 - load config file
	conf, err := getMgrConfig(osFamily, osDistro)
	if err != nil {
		return nil, err
	}

	// 2 - retrun the package manager
	switch osFamily {
	case "debian":
		return &AptPkgManager{Cfg: conf.Apt}, nil
	case "rhel", "fedora":
		return &DnfPkgManager{Cfg: conf.Dnf}, nil
	default:
		return nil, fmt.Errorf("unsupported OS family: %s", osFamily)
	}
}

// Description: returns a RepoManager based on OS family
//
// Notes:
// - the PM has access to the part of the YAML configuration that relates to the package manager
func GetRepoMgr(osFamily, osDistro string) (RepoCli, error) {
	// 1 - load config file
	conf, err := getMgrConfig(osFamily, osDistro)
	if err != nil {
		return nil, err
	}

	// 2 - retrun the package manager
	switch osFamily {
	case "debian":
		return &AptRepoManager{Cfg: conf.Apt}, nil
	case "rhel", "fedora":
		return &DnfRepoManager{Cfg: conf.Dnf}, nil
	default:
		return nil, fmt.Errorf("unsupported OS family: %s", osFamily)
	}
}

// -----------------------------------------
// ------ get YAML file --------------------
// -----------------------------------------

// ####### of global config #######

// Description: returns the YAML configuration
//
// Notes:
// - The placeholders {{ .XXX }} in the YAML are reolved
// - TODO: find a better structured data or solution for the resolution
func getMgrConfig(osFamily, osDistro string) (*ManagerConfig, error) {
	cacheKey := fmt.Sprintf("%s-%s", osFamily, osDistro)

	cacheMgrMu.Lock()
	theYaml, found := configMgrCache[cacheKey]
	cacheMgrMu.Unlock()

	if found {
		return theYaml, nil
	}

	// Not in cache → resolve YAML
	// 1. define the data structure to resolve the YAML placeholders
	varPlaceholder := map[string]map[string]string{
		"Os": {
			"Family": osFamily,
			"Distro": osDistro,
		},
	}
	// 2 - get resolved YAML into struct
	theYaml, err := yamlx.LoadTplYamlFileEmbed[ManagerConfig](yamlMgr, varPlaceholder)
	if err != nil {
		return nil, fmt.Errorf("getting YAML config file in package: %w", err)
	}

	// Store in cache
	cacheMgrMu.Lock()
	configMgrCache[cacheKey] = theYaml
	cacheMgrMu.Unlock()

	// handle success
	return theYaml, nil
}

// ####### of Repo List #######

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
	theYaml, err := yamlx.LoadTplYamlFileEmbed[RepoConfig](yamlRepo, varPlaceholder)
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
	theYaml, err := yamlx.LoadTplYamlFileEmbed[PkgYamlList](yamlPkgList, "")
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
	theYaml, err := yamlx.LoadTplYamlFileEmbed[RepoContentConfig](yamlRepoConent, varPlaceholder)
	if err != nil {
		return nil, fmt.Errorf("getting YAML config file in package: %w", err)
	}
	return theYaml, nil
}
