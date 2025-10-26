package da

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
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
// Description: factory function that returns a PackageManager based on the OS type (rhel/dnf or debian/apt)
func (cbd CliBuilderFactory) get(osFamily string, osDistro string, repo *Repo, pkg *Package) (CliBuilder, error) {
	// normalize input
	family := normalizeOsFamily(osFamily)

	if family == "" {
		return nil, fmt.Errorf("unsupported OS family: %s", osFamily)
	}

	// create object from property
	data := ConfigFileData{
		Os: OsInfo{
			Distro: osDistro,
			Family: osFamily,
		},
	}
	fmt.Println("hello")
	resolvedYaml, err := getConfig(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("helli")

	fmt.Println(resolvedYaml)

	// get the package manager
	switch family {
	case "rhel", "fedora":
		mgr := &DnfManager{
			Repo: repo,
			Pkg:  pkg,
			Conf: cbd.Conf.Dnf,
		}
		return mgr, nil
	case "debian":
		mgr := &AptManager{
			Repo: repo,
			Pkg:  pkg,
			Conf: cbd.Conf.Apt,
		}
		return mgr, nil
	default:
		return nil, fmt.Errorf("unsupported OS family: %s", family)
	}
}

// convenience method
func (pkg *Package) GetCliBuilder(osFamily string, osDistro string, repo *Repo) (CliBuilder, error) {
	return CliBuilderFactory{}.get(osFamily, osDistro, repo, pkg)
}

// convenience method
func (repo *Repo) GetCliBuilder(osFamily string, osDistro string) (CliBuilder, error) {
	return CliBuilderFactory{}.get(osFamily, osDistro, repo, nil)
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
	configFile, err := resolveTplConfig(configFileTpl, c)
	if err != nil {
		return "", fmt.Errorf("faild to resolve templated file: %s", configFileTpl)
	}

	// resturn the YamlString
	return configFile, nil

}

func resolveTplConfig(tplFile string, vars ConfigFileData) (string, error) {
	tpl, err := template.New("repo").Parse(tplFile)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, vars); err != nil {
		return "", err
	}

	return buf.String(), nil
}
