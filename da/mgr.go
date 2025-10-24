package da

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Description: helper function to normalize the OS type (in case name differ between OS).
//
// Todo: move to a nother package
func normalizeOsType(osType string) string {
	s := strings.ToLower(strings.TrimSpace(osType))
	switch s {
	case "debian", "Debian":
		return "debian"
	case "rhel", "RHEL":
		return "rhel"
	default:
		return ""
	}
}

// Name:GetManager
//
// Description: factory function that returns a PackageManager based on the OS type (rhel/dnf or debian/apt)
func getManager(osType string, repo *Repo, pkg *Package) (Manager, error) {
	family := normalizeOsType(osType)
	if family == "" {
		return nil, fmt.Errorf("unsupported OS: %s", osType)
	}

	switch family {
	case "rhel":
		mgr := &DnfManager{
			Repo:   repo,
			Pkg:    pkg,
			Distro: osType, // keep original distro string for later use
		}
		return mgr, nil
	case "debian":
		mgr := &AptManager{
			Repo:   repo,
			Pkg:    pkg,
			Distro: osType, // keep original distro string for later use
		}
		return mgr, nil
	default:
		return nil, fmt.Errorf("unsupported OS family: %s", family)
	}
}

// convenience method
func (pkg *Package) GetManager(osType string, repo *Repo) (Manager, error) {
	return getManager(osType, repo, pkg)
}

// convenience method
func (repo *Repo) GetManager(osType string) (Manager, error) {
	return getManager(osType, repo, nil)
}

// LoadConfig loads the YAML configuration file into a Config struct.
func LoadConfig(path string) (*Config, error) {
	// convert file into memory data
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// convert memory data into struct
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
