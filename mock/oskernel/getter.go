package oskernel

import (
	"fmt"

	"github.com/abtransitionit/gocore/filex"
)

// -----------------------------------------
// ------ get YAML file --------------------
// -----------------------------------------
func GetKernelConf() (*Conf, error) {
	// 1 - get local auto cached (embedded) file into a struct
	cfg, err := filex.LoadYamlIntoStruct[Conf](yamlCfg)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", cfg)

	// handle success
	return cfg, nil
}
