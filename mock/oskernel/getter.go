package oskernel

import (
	"fmt"

	"github.com/abtransitionit/gocore/mock/filex"
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
