package cilium

// define types
type CiliumParam struct {
	PodCidr     string
	ApiServerIp string
}

// define stateless services with their fake instances
type ciliumService struct{}

var CilumSvc = ciliumService{}
