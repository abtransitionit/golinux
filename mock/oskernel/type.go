package oskernel

type Module struct {
	Name string // kernel module name
}

type Parameter struct {
	Name  string // eg. net.ipv4.ip_forward
	Value string // eg. 1
}

type ModuleSlice []Module
type ParameterSlice []Parameter
