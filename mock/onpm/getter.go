package onpm

import (
	_ "embed"
	"sync"

	"fmt"

	"github.com/abtransitionit/gocore/yamlx"
)

//go:embed db.conf.yaml
var confData []byte

// cache the YAML configuration file  on call for the same couple (osFamily, osDistro)
var (
	configCache = make(map[string]*ManagerConfig)
	cacheMu     sync.Mutex
)

// Description: returns a SysManager based on OS family
//
// Notes:
// - the PM has access to the part of the YAML configuration that relates to the package manager
func GetSysMgr(osFamily, osDistro string) (SysCli, error) {
	// 1 - load config file - Ensure it is loaded only once
	conf, err := getConfig(osFamily, osDistro)
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
	conf, err := getConfig(osFamily, osDistro)
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
	conf, err := getConfig(osFamily, osDistro)
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

// Description: returns the YAML configuration
//
// Notes:
// - The YAML is reolved ({{.Os.Family}}, {{.Os.Distro}})
// - TODO: use a structured data instead of osXXX to do this resolution
// - TODO: this method should be private - for ease of use
func getConfig(osFamily, osDistro string) (*ManagerConfig, error) {
	cacheKey := fmt.Sprintf("%s-%s", osFamily, osDistro)

	cacheMu.Lock()
	cfg, found := configCache[cacheKey]
	cacheMu.Unlock()

	if found {
		return cfg, nil
	}

	// Not in cache → resolve YAML
	// 1. define the data structure to resolve the YAML placeholders
	ctx := map[string]map[string]string{
		"Os": {
			"Family": osFamily,
			"Distro": osDistro,
		},
	}
	// 2 - get resolved YAML into struct
	cfg, err := yamlx.LoadTplYamlFileEmbed[ManagerConfig](confData, ctx)
	if err != nil {
		return nil, fmt.Errorf("getting YAML config file in package: %w", err)
	}

	// Store in cache
	cacheMu.Lock()
	configCache[cacheKey] = cfg
	cacheMu.Unlock()

	return cfg, nil
}

// func getConfig(osFamily, osDistro string) (*ManagerConfig, error) {
// 	// cache the YAML configuration file for the same couple (osFamily, osDistro)
// 	cacheKey := fmt.Sprintf("%s-%s", osFamily, osDistro)
// 	cacheMu.Lock()
// 	defer cacheMu.Unlock()
// 	if cfg, ok := configCache[cacheKey]; ok {
// 		return cfg, nil
// 	}
// 	// 1. define the data structure to resolve the YAML placeholders
// 	ctx := map[string]map[string]string{
// 		"Os": {
// 			"Family": osFamily,
// 			"Distro": osDistro,
// 		},
// 	}

// 	// 2 - get resolved YAML into struct
// 	cfg, err := yamlx.LoadTplYamlFileEmbed[ManagerConfig](confData, ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting YAML config file in package")
// 	}

// 	return cfg, nil
// }
