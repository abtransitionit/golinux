package other

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

// description: install a cli
func (i *Manifest) Apply(hostName, folderPath string, logger logx.Logger) error {

	// 1 - get the yaml
	YamlStruct, err := getYaml()
	if err != nil {
		return fmt.Errorf("%s > %w", hostName, err)
	}

	// 2 - get the manifest
	manifest, err := i.getManifest(hostName, YamlStruct)
	if err != nil {
		return fmt.Errorf("getting raw url: %w", err)
	}

	// handle success
	logger.Debugf("%s:%s > installed manifest url %+v", hostName, i.Name, manifest)
	return nil
}

func (i *Manifest) getManifest(hostName string, yaml *MapYaml) (Manifest, error) {

	// 2 - look up the requested RAWCLI by name
	manifest, ok := yaml.List[i.Name]
	if !ok {
		return Manifest{}, fmt.Errorf("%s > Manifest %q not found in YAML", hostName, i.Name)
	}
	// handle success
	return manifest, nil

}
