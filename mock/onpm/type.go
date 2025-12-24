package onpm

import "github.com/abtransitionit/gocore/logx"

// ---------------------------------------------------
// -------------------- interface --------------------
// ---------------------------------------------------

type SysCli interface {
	Update(hostName string, osDistro string, logger logx.Logger) (string, error)
	Upgrade(logx.Logger) string
	NeedReboot(logx.Logger) string
}

type PackageCli interface {
	List() string
	Add(pkgName string, logger logx.Logger) (string, error)
	Remove() string
}

type RepoCli interface {
	List() string
	Add(hostName string, repository Repo2, logger logx.Logger) (string, error)
	Remove() string
}

// -------------------------------------------------------
// -------	 struct for YAML Repo Cfg
// -------------------------------------------------------

// Description: represents the content of the repository file on the os
//
// Notes:
//   - Manage the YAML repo file content
type RepoContentConfig struct {
	Apt string
	Dnf string
}

// -------------------------------------------------------
// -------	 struct for YAML Repo List
// -------------------------------------------------------

// Description: represents the organization's repository db for repos
//
// Notes:
//   - Manage the YAML repo file
type RepoConfig struct {
	Repository map[string]RepoEntry
}
type RepoEntry struct {
	Name string
	Url  struct {
		Repo string
		Gpg  string
	}
}

// -------------------------------------------------------
// -------	 struct for YAML Pkg List
// -------------------------------------------------------

// Description: represents the organization's repository db for packages
//
// Notes:
//   - Manage the YAML repo file
type PkgYamlList struct {
	Package map[string]string
}

// -------------------------------------------------------
// -------	 struct for YAML Mgr Cfg
// -------------------------------------------------------

// Description: represents the YAML configuration file for the manager
//
// Notes:
//   - Manage the YAML configuration file
type ManagerConfig struct {
	Apt *AptConfig `yaml:"apt"`
	Dnf *DnfConfig `yaml:"dnf"`
}

// Description: represents a part of the YAML configuration file
//
// Notes:
//   - represents the part of the YAML configuration file for Apt manager
type AptConfig struct {
	Content string
	Pkg     struct {
		Type     string
		Required map[string][]string
	}
	Ext struct {
		Repo string
		Gpg  struct {
			Url  string
			File string
		}
	}
	Folder struct {
		Repo   string
		GpgKey string
	}
}

// Description: represents a part of the YAML configuration file
//
// Notes:
//   - represents the part of the YAML configuration file for Dnf manager
type DnfConfig struct {
	Content string
	Pkg     struct {
		Type     string
		Required map[string][]string
	}
	Ext struct {
		Repo string
		Gpg  struct {
			Url string
		}
	}
	Folder struct {
		Repo string
	}
	// OS struct {
	// 	Family string
	// 	Distro string
	// }
}

// ---------------------------------------------------
// -------------------- Other --------------
// ---------------------------------------------------

// Description: System manager implementations to manage non-repo and non-package actions
//
// Notes:
// - has access to the YAML configuration data
type AptSysManager struct{ Cfg *AptConfig }
type DnfSysManager struct{ Cfg *DnfConfig }

// Description: Package manager implementations
//
// Notes:
// - has access to the YAML configuration data
type AptPkgManager struct{ Cfg *AptConfig }
type DnfPkgManager struct{ Cfg *DnfConfig }

// Repo manager implementations
//
// Notes:
// - has access to the YAML configuration data
type AptRepoManager struct{ Cfg *AptConfig }
type DnfRepoManager struct{ Cfg *DnfConfig }

// Description: represents a package.
type Package struct {
	Name string
	Cbd  PackageCli // the CLI builder producing commands for this package - set by SetCliBuilder.
}

// Repo represents a repository to be managed.
type Repo struct {
	Name     string
	FileName string
	Version  string
	Url      string
	// Cbd      RepoCli // the CLI builder producing commands for this package - set by SetCliBuilder.
}

type Pkg2 struct {
	Name string
}

type PkgSlice []Pkg2
type Repo2 struct {
	Name     string
	Filename string
	Version  string
}

type RepoSlice []Repo2

// // Description: is a light factory for building CliBuilder instances.
// //
// // Later steps will give it configuration/loader dependencies (template path, resolver).
// type CliBuilderFactory struct {
// 	TplPath string // TplPath optionally holds the path to the templated config file.
// }

// type Config struct {
// 	Apt struct {
// 		Pkg    string `yaml:"pkg"`
// 		Ext    string `yaml:"ext"`
// 		Folder struct {
// 			Repo   string `yaml:"repo"`
// 			GpgKey string `yaml:"gpgKey"`
// 		} `yaml:"folder"`
// 	} `yaml:"apt"`

// 	Dnf struct {
// 		Pkg    string `yaml:"pkg"`
// 		Ext    string `yaml:"ext"`
// 		Folder struct {
// 			Repo string `yaml:"repo"`
// 		} `yaml:"folder"`
// 		Os struct {
// 			Family string `yaml:"family"`
// 			Distro string `yaml:"distro"`
// 		} `yaml:"os"`
// 	} `yaml:"dnf"`
// }
