package kubectl

import (
	"fmt"

	"github.com/abtransitionit/gocore/mock/filex"
)

func GetConfFromYaml(hostName string) (*Cfg, error) {
	// 1 - get local auto cached (embedded) file into a struct
	cfg, err := filex.LoadYamlIntoStruct[Cfg](yamlConf)
	if err != nil {
		return nil, fmt.Errorf("%s > loading yaml: %w", hostName, err)
	}
	// handle success
	return cfg, nil
}

func getYamlListManifest() (*MapYamlManifest, error) {
	// 1 - get local auto cached (embedded) file into a struct
	yamlAsStruct, err := filex.LoadYamlIntoStruct[MapYamlManifest](yamlListManifest)
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}
	return yamlAsStruct, nil
}

func (i *Resource) getFromYaml(hostName string) (Resource, error) {
	// 1 - check
	if i.Type != ResManifest {
		return Resource{}, fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}

	// 2 - get the yaml file into a var/struct
	yaml, err := getYamlListManifest()
	if err != nil {
		return Resource{}, fmt.Errorf("getting the yaml > %w", hostName, err)
	}

	// 2 - look up the requested manifest by name
	resManifest, ok := yaml.List[i.Name]
	if !ok {
		return Resource{}, fmt.Errorf("%s > manifest %q not found in YAML", hostName, i.Name)
	}
	// handle success
	return resManifest, nil

}

// func GetHost(hostName string) (string, error) {
// 	// 1 - get the yaml file into a var/struct
// 	cfg, err := GetConfFromYaml(hostName)
// 	if err != nil {
// 		return "", fmt.Errorf("%s > getting the yaml > %w", hostName, err)
// 	}
// 	// handle success
// 	return cfg.Conf.Kubectl.Host, nil
// }
