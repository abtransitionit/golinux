package k8s

type Node struct {
	Name string // cluster node name
}
type Worker struct {
	Name string // cluster node name
}
type CPlane struct {
	Name string // cluster node name
}
type ClusterConf struct {
	K8sVersion   string
	PodCidr      string
	ServiceCidr  string
	CrSocketName string
}

type NodeSlice []Node
