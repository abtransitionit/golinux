package oskernel

import (
	"fmt"

	"github.com/abtransitionit/gocore/filex"
)

// -----------------------------------------
// ------ get YAML file --------------------
// -----------------------------------------
func getKernelConf() (*ConfigYaml, error) {
	// 1 - get local auto cached (embedded) file into a struct
	cfg, err := filex.LoadYamlIntoStruct[ConfigYaml](yamlCfg)
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}

	// handle success
	return cfg, nil
}
