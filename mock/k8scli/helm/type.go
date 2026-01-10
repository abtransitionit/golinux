package helm

// define types
type Release struct {
	Name      string // eg. kbe-cilium, kbe-kdashb
	Cluster   string // the targeteted k8s cluster
	Namespace string // the targeteted k8s namespace
	Chart     *Chart // the chart used
}
type Repo struct {
	Name string // eg. cilium, kdashb
	Desc string
	Url  string
	Doc  []string
}
type Chart struct {
	Name    string //ie. RepoName/ChartName or /tmp/chart/ChartName
	Qname   string //is qualified or not ie. RepoName/ChartName or /tmp/chart/ChartName
	Version string
	Desc    string
	Repo    *Repo
}

// define stateless services
type repoService struct{}
type releaseService struct{}

// define fake instances of each stateless services
var RepoSvc = repoService{}
var ReleaseSvc = releaseService{}

// defrine slices
type RepoSlice []Repo

// define getters
func GetRepo(name, url string) *Repo {
	i := &Repo{Name: name}
	if url != "" {
		i.Url = url
	}
	return i
}
func GetChart(name, qName, version string) *Chart {
	i := &Chart{Qname: qName}
	if name != "" {
		i.Name = name
	}
	if version != "" {
		i.Version = version
	}
	return i
}

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
