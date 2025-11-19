package onpm

// Description: an object that model a repository for any native linux package manager
type Repo struct {
	Name string
	Url  Url
	Cbd  CliBuilder
}

type Url struct {
	Repo string
	Gpg  string
}

// Description: an object that model a package for any native linux package manager
type Package struct {
	Name    string
	Version string
	Cbd     CliBuilder
}

// Description: method to be implemented by each native package manager supported (apt, dnf)
type CliBuilder interface {
	CliList() (string, error)
	CliAdd() (string, error)
	CliDelete() (string, error)
}

// Description: factory method that return a CliBuilder
type CliBuilderFactory struct {
	Conf Config
}

// APT Manager (manage both Repo and Package)
type AptManager struct {
	Repo *Repo
	Pkg  *Package
	// Distro string // e.g., fedora, rhel, rocky, alma
	Conf AptConfig
}

// DNF Manager (manage both Repo and Package)
type DnfManager struct {
	Repo *Repo
	Pkg  *Package
	// Distro string // e.g., ubuntu
	Conf DnfConfig
}

// Notes:
// - define struct that contains data to be passed to the yaml
type ConfigFileData struct {
	Os OsObj
}

// Notes:
// - define struct to convert yaml (string of file) to memory struct
type Config struct {
	Apt AptConfig `yaml:"apt"`
	Dnf DnfConfig `yaml:"dnf"`
}

type AptConfig struct {
	Folder FolderObj `yaml:"folder"`
	Ext    string    `yaml:"ext"`
	Pkg    string    `yaml:"pkg"`
}

type DnfConfig struct {
	Folder FolderObj `yaml:"folder"`
	Ext    string    `yaml:"ext"`
	Pkg    string    `yaml:"pkg"`
	Os     OsObj     `yaml:"os"`
}
type FolderObj struct {
	Repo   string `yaml:"repo"`
	GpgKey string `yaml:"GpgKey,omitempty"`
}
type OsObj struct {
	Distro string
	Family string
}

// a map of repo
type MapRepo map[string]Repo

// // Todo in next version

// type OSType string

// const (
// 	OS_APT     OSType = "apt"
// 	OS_DNF     OSType = "dnf"
// 	OS_UNKNOWN OSType = "unknown"
// )

// // Load the YAML once at initialization:
// type ManagerConfig struct {
// 	RepoFolder string `yaml:"repoFolder"`
// 	KeyFolder  string `yaml:"keyFolder,omitempty"`
// 	GPGCheck   bool   `yaml:"gpgCheck,omitempty"`
// }

// // Load the YAML once at initialization:
// type Config struct {
// 	APT ManagerConfig `yaml:"apt"`
// 	DNF ManagerConfig `yaml:"dnf"`
// }

// // LoadConfig loads the manager.yaml file into a Config struct
// func LoadConfig(path string) (*Config, error) {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read config file: %v", err)
// 	}

// 	var config Config
// 	if err := yaml.Unmarshal(data, &config); err != nil {
// 		return nil, fmt.Errorf("failed to parse YAML: %v", err)
// 	}

// 	return &config, nil
// }

// // listFile := filepath.Join(config.APT.RepoFolder, "myrepo.list")
// // keyFile := filepath.Join(config.APT.KeyFolder, "myrepo.gpg")
// // repoFile := filepath.Join(config.DNF.RepoFolder, "myrepo.repo")
