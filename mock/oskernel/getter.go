package oskernel

import (
	"fmt"

	"github.com/abtransitionit/gocore/filex"
)

// -----------------------------------------
// ------ get YAML file --------------------
// -----------------------------------------
func getCfg() (*Cfg, error) {
	// 1 - get local auto cached (embedded) file into a struct
	yamlAsStruct, err := filex.LoadYamlIntoStruct[Cfg](yamlCfg)
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}

	// handle success
	return yamlAsStruct, nil
}

// func getYaml(hostName string) (*MapYaml, error) {
// 	// 1 - get local auto cached (embedded) file into a struct
// 	yamlAsStruct, err := filex.LoadYamlIntoStruct[MapYaml](yamlList)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s > loading config: %w", hostName, err)
// 	}
// 	// handle success
// 	return yamlAsStruct, nil
// }
