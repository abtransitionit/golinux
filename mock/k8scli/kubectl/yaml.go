package kubectl

import (
	"fmt"

	// "github.com/abtransitionit/gocore/filex"
	"github.com/abtransitionit/gocore/mock/filex"
)

// -----------------------------------------
// ------ get YAML file into a var     -----
// -----------------------------------------

func GetConfFromYaml(hostName string) (*Cfg, error) {
	// 1 - get local auto cached (embedded) file into a struct
	cfg, err := filex.LoadYamlIntoStruct[Cfg](yamlConf)
	if err != nil {
		return nil, fmt.Errorf("%s > loading yaml: %w", hostName, err)
	}
	// handle success
	return cfg, nil
}
func GetHost(hostName string) (string, error) {
	// 1 - get the yaml file into a var/struct
	cfg, err := GetConfFromYaml(hostName)
	if err != nil {
		return "", fmt.Errorf("%s > getting the yaml > %w", hostName, err)
	}
	// handle success
	return cfg.Conf.Kubectl.Host, nil
}
