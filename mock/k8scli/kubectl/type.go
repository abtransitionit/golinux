package kubectl

// -------------------------------------------------------
// -------	 generic k8s resource
// -------------------------------------------------------
type ResType string

const (
	ResNode      ResType = "node"
	ResPod       ResType = "pod"
	ResNS        ResType = "ns"
	ResCM        ResType = "cm"
	ResCRD       ResType = "crd"
	ResSA        ResType = "sa"
	ResApiServer ResType = "api"
	ResSC        ResType = "sc"
	ResPv        ResType = "pv"  // related to SC
	ResPvc       ResType = "pvc" // related to SC
	ResRes       ResType = "res" // api-resources
	ResSecret    ResType = "secret"
	ResManifest  ResType = "manifest" // manifest file
)

func (t ResType) String() string {
	return string(t)
}

type Resource struct {
	Name     string
	Type     ResType // node, pod, ns, cm, sa, secret, mnf,
	Ns       string
	UserName string   // for Secret only
	Desc     string   // for manifest only
	Url      string   // for manifest only
	Doc      []string // for manifest only
}

// define slice
type SliceResource []Resource

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
