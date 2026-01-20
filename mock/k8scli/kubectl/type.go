package kubectl

// -------------------------------------------------------
// -------	 generic k8s resource
// -------------------------------------------------------
type ResType string

const (
	ResNode       ResType = "node"
	ResPod        ResType = "pod"
	ResNS         ResType = "ns"
	ResCM         ResType = "cm"
	ResCRD        ResType = "crd"
	ResSA         ResType = "sa"
	RestApiServer ResType = "api"
	ResSC         ResType = "sc"
	ResPvc        ResType = "pvc" // related to SC
	ResRes        ResType = "res" // api-resources
)

func (t ResType) String() string {
	return string(t)
}

type Resource struct {
	Name string
	Type ResType // node, pod, ns, cm, sa
	Ns   string
}

// define slice
type SliceResource []Resource

// -------------------------------------------------------
// -------	 struct for YAML
// -------------------------------------------------------

// -------	 struct for YAML Conf - denotes helm node
type Cfg struct {
	Conf *Conf
}
type Conf struct {
	Kubectl struct {
		Host string
	}
}
