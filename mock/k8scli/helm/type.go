package helm

// -------------------------------------------------------
// -------	 generic k8s resource
// -------------------------------------------------------
type ResType string
type ResSType string

const (
	ResRepo         ResType = "repo"
	ResChart        ResType = "chart"
	ResArtifact     ResType = "artifact"
	ResChartVersion ResType = "chartVersion"
	ResRelease      ResType = "release"
	ResRegistry     ResType = "registry"
	ResHelm         ResType = "helm"
)

// subtypes - chart
const (
	STypeChartStd   ResSType = "chart/std"
	STypeChartOCI   ResSType = "chart/oci"
	STypeChartBuild ResSType = "chart/helm"
)

// subtypes - artifact
const (
	STypeArtifactGo   ResSType = "artifact/go"
	STypeArtifactHelm ResSType = "artifact/helm"
)

func (t ResType) String() string {
	return string(t)
}

// the oci registry - begin
type RegistryCfg struct {
	Registry map[string]Registry `yaml:"registry"`
}

type Registry struct {
	Description string
	Type        string
	Param       RegistryParam
}

type RegistryParam struct {
	DnsOrIp     string `yaml:"dnsOrIp"`
	Org         string
	User        string
	Path        string
	AccessToken string `yaml:"accessToken"`
}

// the oci registry - end

// the artifact build from chart - begin
type ArtifactCfg struct {
	Artifact map[string]ArtifactSet `yaml:"artifact"`
}

type ArtifactSet struct {
	Items map[string]Artifact `yaml:",inline"`
}

type Artifact struct {
	FolderSrc string `yaml:"folderSrc"`
	FolderDst string `yaml:"folderDst"`
}

type ArtifactChartYaml struct {
	Name         string
	AppVersion   string `yaml:"appVersion"`
	ChartVersion string `yaml:"version"`
}

// the artifact build from chart - end

type Resource struct {
	Name      string
	Type      ResType
	SType     ResSType
	Url       string            // repo only
	Repo      string            // chart only
	QName     string            // chart only - eg. cilium/cilium;
	Namespace string            // release only
	Revision  string            // release only
	Version   string            // release and chart only - version of the chart to install if any
	Param     map[string]string // release, chart
	ValueFile []byte            // release only
}

// -------------------------------------------------------
// -------	 the old way
// -------------------------------------------------------

// define types

type Release struct {
	Name      string            // eg. kbe-cilium, kbe-kdashb
	Chart     *Chart            // the chart qualified name. eg. RepoName/ChartName or /tmp/chart/ChartName
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
}

// define stateless services with their fake instances
type repoService struct{}
type releaseService struct{}
type helmService struct{}

var RepoSvc = repoService{}
var ReleaseSvc = releaseService{}
var HelmSvc = helmService{}

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
func GetRelease(name, chartQName, chartVersion, namespace string, param map[string]string) *Release {
	i := &Release{
		Chart: &Chart{},
	}
	if name != "" {
		i.Name = name
	}
	if chartVersion != "" {
		i.Chart.Version = chartVersion
	}
	if chartQName != "" {
		i.Chart.QName = chartQName
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
