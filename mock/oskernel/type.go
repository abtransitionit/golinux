package oskernel

// defrine types for YAML
type Conf struct {
	folder struct {
		module string
		param  string
	}
}

// defrine types
type Module struct {
	Name        string // kernel module name
	CfgFilePath string // kernel module config file
}

type Parameter struct {
	Name        string // eg. net.ipv4.ip_forward
	Value       string // eg. 1
	CfgFilePath string // kernel parameter config file
}

// defrine slices
type ModuleSlice []Module
type ParameterSlice []Parameter

// define getters
func GetModule(name string, cfgFilePath string) *Module {
	return &Module{
		Name:        name,
		CfgFilePath: cfgFilePath,
	}
}
func GetParameter(name string, value string, cfgFilePath string) *Parameter {
	return &Parameter{
		Name:        name,
		Value:       value,
		CfgFilePath: cfgFilePath,
	}
}
