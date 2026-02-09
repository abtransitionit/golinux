package kubectl

// -------------------------------------------------------
// -------	 generic k8s resource
// -------------------------------------------------------
type ResType string

const (
	ResCM        ResType = "cm"
	ResEp        ResType = "ep"
	ResDeploy    ResType = "deploy"
	ResDs        ResType = "ds"
	ResNode      ResType = "node"
	ResNS        ResType = "ns"
	ResPod       ResType = "pod"
	ResPv        ResType = "pv"  // related to SC
	ResPvc       ResType = "pvc" // related to SC
	ResSA        ResType = "sa"
	ResSecret    ResType = "secret"
	ResSC        ResType = "sc"
	ResSvc       ResType = "svc"
	ResCRD       ResType = "crd"
	ResApiServer ResType = "api"
	ResRes       ResType = "res"      // api-resources
	ResManifest  ResType = "manifest" // manifest file
)

func (t ResType) String() string {
	return string(t)
}

type Resource struct {
	Name     string
	Kind     string  // the k8s kind (e.g., Deployment, DaemonSet, Service, etc.)
	Type     ResType // node, pod, ns, cm, sa, secret, mnf,
	Ns       string
	UserName string            // for Secret only
	Desc     string            // for manifest only
	Url      string            // for manifest only
	Doc      []string          // for manifest only
	Param    map[string]string // list of placeholders for manifest only
	Port     string            // for svc only
}

// define slice
type SliceResource []Resource

// -------------------------------------------------------
// -------	 private methods
// -------------------------------------------------------

func (i *Resource) isClusterScoped() bool {
	switch i.Type {
	case ResNode, ResNS, ResPv, ResSC:
		return true
	default:
		return false
	}
}

// -------------------------------------------------------
// -------	 struct for YAML
// -------------------------------------------------------

// -------	 struct for YAML Conf
// Description: represents the organization's repository db for Helm host
type Cfg struct {
	Conf *Conf
}
type Conf struct {
	Kubectl struct {
		Host string
	}
}

// -------	 struct for YAML list Manifest
// Description: represents the organization's repository db for manifest file that can be applied
type MapYamlManifest struct {
	List map[string]Resource
}
