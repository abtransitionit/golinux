package gopm

import (
	"fmt"

	"github.com/abtransitionit/gocore/filex"
)

func getYaml(hostName string) (*MapYaml, error) {
	// 1 - get local auto cached (embedded) file into a struct
	yamlAsStruct, err := filex.LoadYamlIntoStruct[MapYaml](yamlList)
	if err != nil {
		return nil, fmt.Errorf("%s > loading config: %w", hostName, err)
	}
	// handle success
	return yamlAsStruct, nil
}

// description: returns the cli info from the DB
//
// Notes:
// - the info are VM specific
// func GeRawtUrlFromYaml(cliName string) (string, error) {
// 	// // 1 - get local auto cached (embedded) file into a struct
// 	// yamlFile, err := filex.LoadYamlIntoStruct[CliYaml](yamlList)
// 	// if err != nil {
// 	// 	return "", fmt.Errorf("loading config: %w", err)
// 	// }
// 	yamlAsStruct, err := getYaml(cliName)
// 	if err != nil {
// 		return "", err
// 	}

// 	// 2 - look up the requested RAWCLI by name
// 	cli, ok := yamlAsStruct.List[cliName]
// 	if !ok {
// 		return "", fmt.Errorf("CLI %q not found in YAML", cliName)
// 	}
// 	// handle success
// 	return cli.Url, nil
// }

// // 3 - resolve the cli:url

// cli.Url, err = ResolveURL(nil, cli, osFamily, osDistro, "")
// if err != nil {
// 	return nil, err
// }

// handle success

// // 2 - retrun the package manager
// switch osFamily {
// case "debian":
// 	return &AptPkgManager{Cfg: cfg.Apt}, nil
// case "rhel", "fedora":
// 	return &DnfPkgManager{Cfg: cfg.Dnf}, nil
// default:
// 	return nil, fmt.Errorf("unsupported OS family : %s", osFamily)
// }
