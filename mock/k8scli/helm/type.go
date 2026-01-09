package helm

// define types
type Repo struct {
	Name string // eg. cilium, kdashb
	Desc string
	Url  string
	Doc  []string
}

// defrine slices
type RepoSlice []Repo

// define getters
func GetRepo(name, url string) *Repo {
	r := &Repo{Name: name}
	if url != "" {
		r.Url = url
	}
	return r
}

// func GetRepo(i Repo) *Repo {
// 	return &Repo{
// 		Name: i.Name,
// 	}
// }
// func GetRepoByName(name string) *Repo {
// 	return &Repo{
// 		Name: name,
// 	}
// }

// -------------------------------------------------------
// -------	 struct for YAML
// -------------------------------------------------------

// -------------------------------------------------------
// -------	 struct for YAML Repo List
// -------------------------------------------------------
//   - denotes the whitelist
type MapYaml struct {
	List map[string]Repo
}

// -------------------------------------------------------
// -------	 struct for YAML Conf
// -------------------------------------------------------
//   - denotes helm node

type Cfg struct {
	Conf *Conf
}
type Conf struct {
	Helm struct {
		Host string
	}
}

// type HelmRepo struct {
// 	Name string // logical name
// 	Desc string
// 	Url  string
// 	Doc  []string
// }
// type HelmChart struct {
// 	FullName string //ie. RepoName/ChartName or /tmp/chart/ChartName
// 	Version  string
// 	Desc     string
// 	Repo     HelmRepo
// }
// type HelmRelease struct {
// 	Name      string
// 	Repo      HelmRepo
// 	Chart     HelmChart
// 	Namespace string
// 	ValueFile string
// }

// type MapHelmRepo map[string]HelmRepo
