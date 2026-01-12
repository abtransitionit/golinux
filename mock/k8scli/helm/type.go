package helm

// define types
type Release struct {
	Name      string            // eg. kbe-cilium, kbe-kdashb
	CQName    string            // the chart qualified name. eg. RepoName/ChartName or /tmp/chart/ChartName
	Version   string            // the version of the chart
	Namespace string            // the targeteted k8s namespace
	ValueFile string            // TODO : in test
	Param     map[string]string // list of placeholders
	// Cluster   string // the targeteted k8s cluster
}
type Repo struct {
	Name string // eg. cilium, kdashb
	Desc string
	Url  string
	Doc  []string
}
type Chart struct {
	QName   string // is qualified or not ie. RepoName/ChartName or /tmp/chart/ChartName
	Name    string // ie. ChartName or /tmp/chart/ChartName
	Version string
	Desc    string
	Repo    *Repo
}

// define stateless services with their fake instances
type repoService struct{}
type releaseService struct{}

var RepoSvc = repoService{}
var ReleaseSvc = releaseService{}

// defrine slices
type RepoSlice []Repo
type ReleaseSlice []Release

// define getters
func GetRepo(name, url string) *Repo {
	i := &Repo{Name: name}
	if url != "" {
		i.Url = url
	}
	return i
}
func GetChart(name, qName, version string) *Chart {
	i := &Chart{QName: qName}
	if name != "" {
		i.Name = name
	}
	if version != "" {
		i.Version = version
	}
	return i
}
func GetRelease(name, cqName, version, namespace string, param map[string]string) *Release {
	i := &Release{}
	if name != "" {
		i.Name = name
	}
	if version != "" {
		i.Version = version
	}
	if cqName != "" {
		i.CQName = cqName
	}
	if namespace != "" {
		i.Namespace = namespace
	}
	if len(param) > 0 {
		i.Param = param
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
