package k8s

// define types
type Node struct {
	Name string // cluster node name
}
type Pvc struct {
	Name string // cluster node name
}

type Worker struct {
	Name string // cluster node name
}

type CPlane struct {
	Name string // cluster node name
}
type ClusterParam struct {
	K8sVersion   string
	PodCidr      string
	ServiceCidr  string
	CrSocketName string
}
type Kubectl struct {
	CplaneNdeName   string
	InstallNodeName string
}

// define slices
type NodeSlice []Node

// define getters
func GetNode(name string) *Node {
	return &Node{
		Name: name,
	}
}
func GetCplane(name string) *CPlane {
	return &CPlane{
		Name: name,
	}
}
func GetWorker(name string) *Worker {
	return &Worker{
		Name: name,
	}
}

func GetKubectl(cplaneNodeName string, installNodeName string) *Kubectl {
	return &Kubectl{
		CplaneNdeName:   cplaneNodeName,
		InstallNodeName: installNodeName,
	}
}

// desscription: return a pointer to the struct
func GetClusterParam(clusterParam ClusterParam) *ClusterParam {
	return &ClusterParam{
		K8sVersion:   clusterParam.K8sVersion,
		PodCidr:      clusterParam.PodCidr,
		ServiceCidr:  clusterParam.ServiceCidr,
		CrSocketName: clusterParam.CrSocketName,
	}
}
