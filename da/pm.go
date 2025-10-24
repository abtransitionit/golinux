package da

import "fmt"

// GetPM returns the right PackageManager for the specified OS (rhel/dnf or ubuntu/apt)
//
// Notes:
// - the call is done from a repo
func (pkg *Package) GetPM(osType string, repo *Repo) (Manager, error) {
	switch osType {
	case "rhel", "fedora":
		mgr := &DnfManager{Pkg: pkg}
		if repo != nil {
			mgr.Repo = repo
		}
		return mgr, nil
	case "ubuntu":
		mgr := &AptManager{Pkg: pkg}
		if repo != nil {
			mgr.Repo = repo
		}
		return mgr, nil
	default:
		return nil, fmt.Errorf("unsupported OS: %s", osType)
	}
}

// GetPM returns the right PackageManager for the specified OS (rhel/dnf or ubuntu/apt)
// Notes:
// - the call is done from a package

func (repo *Repo) GetPM(osType string) (Manager, error) {
	switch osType {
	case "rhel", "fedora":
		return &DnfManager{Repo: repo}, nil
	case "ubuntu":
		return &AptManager{Repo: repo}, nil
	default:
		return nil, fmt.Errorf("unsupported OS: %s", osType)
	}
}
