package gopm

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/mock/yamlx"
	"github.com/abtransitionit/golinux/mock/file"
	"github.com/abtransitionit/golinux/mock/property"
)

// description: install a cli
func (i *Cli) Install(hostName, folderPath string, logger logx.Logger) error {

	// 1 - get the yaml
	YamlStruct, err := getYaml(hostName)
	if err != nil {
		return fmt.Errorf("%s > %w", hostName, err)
	}

	// 2 - get the raw cli
	rawCli, err := i.getRawCli(hostName, YamlStruct)
	if err != nil {
		return fmt.Errorf("getting raw url: %w", err)
	}
	// 3 - set the cli:resolved url
	if err = i.setUrl(hostName, rawCli, logger); err != nil {
		return fmt.Errorf("setting cli url: %w", err)
	}
	// 3 - set the cli:type
	if err = i.setType(hostName, rawCli, logger); err != nil {
		return fmt.Errorf("setting cli url: %w", err)
	}

	// 4 - donwload the artifact pointed by the url on a host temp file
	// 41 - get instance of artifact
	artifact := file.GetArtifact(i.Name, i.Url, i.Type)
	// artifactFullPath, err = file.DownloadArtifact(hostName, i.Url, i.Name, i.Type, logger)
	if _, err = artifact.Download(hostName, logger); err != nil {
		return fmt.Errorf("%s > donwloading url %s: %w", hostName, i.Url, err)
	}
	// // 5 - copy the host artifact temp file to the final destination on the host
	// dstFilePath := filepath.Join(folderPath, i.Name)
	// if err := file.CopyArtifactToDest(hostName, artifactFullPath, dstFilePath, i.Type, logger); err != nil {
	// 	return fmt.Errorf("copying file %s to %s: %w", artifactFullPath, dstFilePath, err)
	// }

	// log
	// logger.Debugf("%s > copy %s to %s", hostName, artifactFullPath, folderPath)
	// logger.Infof("%s > will do cli: %s", hostName, cli)
	// 5 - detect the donwload file type - ie. tar.gz, zip, exe, ...
	// 6 - move the file on the host

	// handle success
	logger.Debugf("%s:%s:%s > installed from url %s", hostName, i.Name, i.Version, i.Url)
	return nil
}

// description: get the raw url of a cli from the yaml
func (i *Cli) getRawCli(hostName string, yaml *MapYaml) (Cli, error) {

	// 2 - look up the requested RAWCLI by name
	cli, ok := yaml.List[i.Name]
	if !ok {
		return Cli{}, fmt.Errorf("%s > CLI %q not found in YAML", hostName, i.Name)
	}
	// handle success
	return cli, nil

}

// description: set the url of a cli from the raw url
func (i *Cli) setUrl(hostName string, rawCli Cli, logger logx.Logger) error {
	// define var
	var osType, osArch, uname string
	var err error
	var resolvedUrl []byte
	tpl := []byte(rawCli.Url)

	// 1 - get host:property
	if osType, err = property.GetProperty(logger, hostName, "osType"); err != nil {
		return err
	}
	if osArch, err = property.GetProperty(logger, hostName, "osArch"); err != nil {
		return err
	}
	if uname, err = property.GetProperty(logger, hostName, "uname"); err != nil {
		return err
	}

	// 2 - define  placeholder
	varPlaceholder := map[string]map[string]string{
		"Cli": {
			"Tag":  i.Version,
			"Name": i.Name,
		},
		"Os": {
			"Type":  osType,
			"Arch":  osArch,
			"Uname": uname,
		},
	}
	// 3 - resolve url
	if resolvedUrl, err = yamlx.ResolveTplConfig(tpl, varPlaceholder); err != nil {
		return fmt.Errorf("template resolve failed: %v", err)
	}

	// 4 - set url
	i.Url = string(resolvedUrl)

	// handle success
	return nil
}
func (i *Cli) setType(hostName string, rawCli Cli, logger logx.Logger) error {
	// define var

	// 4 - set type
	i.Type = rawCli.Type

	// handle success
	return nil
}

func (i *CliSlice) Install(logger logx.Logger) error {
	// log
	logger.Info("Install called")
	// handle success
	return nil
}
