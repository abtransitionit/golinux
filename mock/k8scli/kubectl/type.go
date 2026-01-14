package kubectl

// -------------------------------------------------------
// -------	 generic k8s resource
// -------------------------------------------------------
type ResType string

const (
	ResNode ResType = "node"
	ResPod  ResType = "pod"
	ResNS   ResType = "ns"
	ResCM   ResType = "cm"
	ResSA   ResType = "sa"
)

func (t ResType) String() string {
	return string(t)
}

type Resource struct {
	Name string
	Type ResType // node, pod, ns, cm, sa
	Ns   string
}

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

// -------------------------------------------------------
// -------	 statefull services
// -------------------------------------------------------
// type Node struct {
// 	Type string // set to node
// 	Name string // eg. o1u, o2a
// 	Ns   string // eg. kube-system, default
// }
// type Pod struct {
// 	Type string // set to pod
// 	Name string
// 	Ns   string // eg. kube-system, default
// }
// type Namespace struct {
// 	Type string // set to ns
// 	Name string // eg. kube-system, default
// }
// type ConfigMap struct {
// 	Type string // set to cm
// 	Name string // eg. kube-system, default
// }
// type ServiceAccount struct {
// 	Type string // set to sa
// 	Name string
// 	Ns   string // eg. kube-system, default
// }

// -------------------------------------------------------
// -------	 stateless services it fake instance
// -------------------------------------------------------
// type resService struct{}

// type nodeService struct{}

// type podService struct{}
// type nsService struct{}

// type saService struct{}
// type cmService struct{}

// var NodeSvc = nodeService{}

// var PodSvc = podService{}
// var NsSvc = nsService{}

// var SaSvc = saService{}
// var CmSvc = cmService{}
// var resSvc = resService{}
