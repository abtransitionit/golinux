package oskernel

// defrine types for YAML
type Cfg struct {
	Conf *Conf
}
type Conf struct {
	Folder struct {
		Module string
		Param  string
	}
}
type FolderConfig struct {
	Module string
	Param  string
}

// defrine types
type Module struct {
	Name string // kernel module name
	// CfgFileName string // kernel module config file name
}

type Parameter struct {
	Name  string // eg. net.ipv4.ip_forward
	Value string // eg. 1
	// CfgFileName string // kernel parameter config file
}

// defrine slices
type ModuleSlice []Module
type ParameterSlice []Parameter

// defrine sets

type ModuleSet struct {
	ModuleSlice []Module
	CfgFileName string
}

type ParameterSet struct {
	ParameterSlice []Parameter
	CfgFileName    string
}

// define getters
//
//	func GetModule(name string, cfgFileName string) *Module {
//		return &Module{
//			Name:        name,
//			CfgFileName: cfgFileName,
//		}
//	}
func GetModuleSet(moduleSlice []Module, cfgFileName string) *ModuleSet {

	return &ModuleSet{
		ModuleSlice: moduleSlice,
		CfgFileName: cfgFileName,
	}
}

func GetParameterSet(parameterSlice []Parameter, cfgFileName string) *ParameterSet {

	return &ParameterSet{
		ParameterSlice: parameterSlice,
		CfgFileName:    cfgFileName,
	}
}

// func GetParameter(name string, value string, cfgFileName string) *Parameter {
// 	return &Parameter{
// 		Name:        name,
// 		Value:       value,
// 		CfgFileName: cfgFileName,
// 	}
// }
