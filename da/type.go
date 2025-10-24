package da

type Repo struct {
	Name string
	Url  RepoUrl
	Mgr  Manager
}
type RepoUrl struct {
	Repo string
	Gpg  string
}
type Package struct {
	Name    string
	Version string
	Mgr     Manager
}

type Manager interface {
	CliList() (string, error)
	CliAdd() (string, error)
	CliDelete() (string, error)
}

// APT Manager (manage for both Repo and Package)
type AptManager struct {
	Repo   *Repo
	Pkg    *Package
	Distro string // e.g., "fedora", "rhel", "rocky", "alma"
}

// DNF Manager (manage for both Repo and Package)
type DnfManager struct {
	Repo   *Repo
	Pkg    *Package
	Distro string
}

// a map of repo
type MapRepo map[string]Repo

// // Todo in next version

// type Repo struct {
// 	Name    string
// 	URL     string
// 	GPGKey  string
// 	manager RepoManager // OS-specific manager: dnf or apt
// 	Enabled bool
// }

// // DNF implementation
// type DnfManager struct {
// 	Repo Repo
// }

// // APT implementation
// type AptManager struct {
// 	Repo Repo
// }

// // type Repository struct {
// // 	Name     string // logical name
// // 	FileName string // Os file name
// // 	Version  string // the version of the package repository
// // 	UrlRepo  string
// // 	UrlGpg   string
// // }

// type Package struct {
// 	Name    string // logical name
// 	CName   string // canonical name
// 	Version string // the version of the package
// }

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

// // func DetectOS() OSType

// // var repo RepositoryManager
// // var pkg PackageManager

// // switch DetectOS() {
// // case OS_APT:
// //     repo = NewAPTRepositoryManager()
// //     pkg = NewAPTPackageManager()
// // case OS_DNF:
// //     repo = NewDNFRepositoryManager()
// //     pkg = NewDNFPackageManager()
// // default:
// //     // Handle unknown OS
// // }
