package helm

import (
	"fmt"
	"strings"

	// "github.com/abtransitionit/gocore/filex"
	"github.com/abtransitionit/gocore/mock/filex"
)

// -----------------------------------------
// ------ get YAML file into a var     -----
// -----------------------------------------

//	func GetYamlRepo(hostName string) (*MapYaml, error) {
//		// 1 - get local auto cached (embedded) file into a struct
//		mapYaml, err := filex.LoadYamlIntoStruct[MapYaml](yamlListRepo)
//		if err != nil {
//			return nil, fmt.Errorf("%s > loading yaml: %w", hostName, err)
//		}
//		// handle success
//		return mapYaml, nil
//	}
func GetYamlRepo() (*MapYaml, error) {
	// 1 - get local auto cached (embedded) file into a struct
	mapYaml, err := filex.LoadYamlIntoStruct[MapYaml](yamlListRepo)
	if err != nil {
		return nil, fmt.Errorf("loading yaml: %w", err)
	}
	// handle success
	return mapYaml, nil
}
func GetYamlConf(hostName string) (*Cfg, error) {
	// 1 - get local auto cached (embedded) file into a struct
	confYaml, err := filex.LoadYamlIntoStruct[Cfg](yamlConf)
	if err != nil {
		return nil, fmt.Errorf("%s > loading yaml: %w", hostName, err)
	}
	// handle success
	return confYaml, nil
}
func GetHelmHost(hostName string) (string, error) {
	// 1 - get the yaml file into a var/struct
	YamlStruct, err := GetYamlConf(hostName)
	if err != nil {
		return "", fmt.Errorf("%s > getting the yaml > %w", hostName, err)
	}
	// handle success
	return YamlStruct.Conf.Helm.Host, nil
}

func (i *MapYaml) ConvertToString() string {
	var sb strings.Builder
	sb.WriteString("NAME\tDESCRIPTION\tURL\n") // header

	for _, r := range i.List {
		sb.WriteString(fmt.Sprintf(
			"%s\t%s\t%s\n",
			r.Name,
			r.Desc,
			r.Url,
		))
	}

	return sb.String()
}
