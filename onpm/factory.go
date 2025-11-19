package onpm

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/abtransitionit/gocore/tpl"
)

// Description: helper function to normalize the OS type (in case name differ between OS).
//
// Todo: move to a nother package
func normalizeOsFamily(osFamily string) string {
	s := strings.ToLower(strings.TrimSpace(osFamily))
	switch s {
	case "debian":
		return "debian"
	case "rhel", "fedora":
		return "rhel"
	default:
		return ""
	}
}

// Name:GetCliBuilder
//
// Description: returns a PackageManager based on the OS type (rhel/dnf or debian/apt)
func (cbd CliBuilderFactory) get(osFamily string, osDistro string, repo *Repo, pkg *Package) (CliBuilder, error) {
	// normalize input
	family := normalizeOsFamily(osFamily)

	if family == "" {
		return nil, fmt.Errorf("unsupported OS family: %s", osFamily)
	}
	if osDistro == "" {
		return nil, fmt.Errorf("unsupported OS family: %s", osFamily)
	}

	// create object from property
	data := ConfigFileData{
		Os: OsObj{
			Distro: osDistro,
			Family: osFamily,
		},
	}

	TemplatedConfFile, err := tpl.LoadFile("/Users/max/wkspc/git/golinux/da/conf.yaml")
	if err != nil {
		return nil, err
	}

	// Resolv templated conf into a Yaml string
	resolvedYamlAsString, err := tpl.ResolveTplConfig(TemplatedConfFile, data)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	// convert Yaml string into a struct
	var cfg Config // Declare the struct to hold the YAML
	err = yaml.Unmarshal([]byte(resolvedYamlAsString), &cfg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML: %v", err)
	}

	// get the package manager
	switch family {
	case "rhel", "fedora":
		mgr := &DnfManager{
			Repo: repo,
			Pkg:  pkg,
			Conf: cfg.Dnf,
		}
		return mgr, nil
	case "debian":
		mgr := &AptManager{
			Repo: repo,
			Pkg:  pkg,
			Conf: cfg.Apt,
		}
		return mgr, nil
	default:
		return nil, fmt.Errorf("unsupported OS family: %s", family)
	}
}

// convenience method
func (pkg *Package) SetCliBuilder(osFamily string, osDistro string, repo *Repo) error {
	cbd, err := CliBuilderFactory{}.get(osFamily, osDistro, repo, pkg)
	if err != nil {
		return err
	}
	pkg.Cbd = cbd
	return nil
}

// convenience method
func (repo *Repo) GetCliBuilder(osFamily string, osDistro string) error {
	cbd, err := CliBuilderFactory{}.get(osFamily, osDistro, repo, nil)
	if err != nil {
		return err
	}
	repo.Cbd = cbd
	return nil
}

// Name: GetConfig
//
// Description: resolve the templated config file and return it as a YamlString
func getConfig(c ConfigFileData) (string, error) {

	// // define the structure that holds the vars that will be used to resolve the templated file
	// configFileTplVar := ConfigFileData{
	// 	OsDistro: c.OsDistro,
	// }

	// resolve the templated file
	configFile, err := tpl.ResolveTplConfig(configFileTpl, c)
	if err != nil {
		return "", fmt.Errorf("faild to resolve templated file: %s", configFileTpl)
	}

	// resturn the YamlString
	return configFile, nil

}
