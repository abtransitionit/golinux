package gopm

import (
	"fmt"

	"github.com/abtransitionit/gocore/filex"
)

// description: returns the cli info from the DB
//
// Notes:
// - the info are VM specific
func GetCliFromYaml(osFamily, osDistro, cliName string) (*Cli, error) {
	// 1 - get local auto cached (embedded) file into a struct
	yamlFile, err := filex.LoadYamlIntoStruct[CliYaml](yamlList)
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}

	// 2 - look up the requested CLI by name
	cli, ok := yamlFile.List[cliName]
	if !ok {
		return nil, fmt.Errorf("CLI %q not found in YAML", cliName)
	}

	// handle success
	return &cli, nil
}

// // 2 - retrun the package manager
// switch osFamily {
// case "debian":
// 	return &AptPkgManager{Cfg: cfg.Apt}, nil
// case "rhel", "fedora":
// 	return &DnfPkgManager{Cfg: cfg.Dnf}, nil
// default:
// 	return nil, fmt.Errorf("unsupported OS family : %s", osFamily)
// }
