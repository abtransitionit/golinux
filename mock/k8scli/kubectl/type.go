package kubectl

// -------------------------------------------------------
// -------	 statefull services
// -------------------------------------------------------
type Node struct {
	Name string // eg. o1u, o2a
}
type Pod struct {
	Name string
}
type Namespace struct {
	Name string // eg. kube-system, default
}

// -------------------------------------------------------
// -------	 stateless services with their fake instances
// -------------------------------------------------------
type nodeService struct{}
type podService struct{}
type namespaceService struct{}

var NodeSvc = nodeService{}
var PodSvc = podService{}
var NamespaceSvc = namespaceService{}

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
