package helm

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func (i *Resource) getFromYaml(hostName string) (Repo, error) {
	// 1 - check
	if i.Type != ResRepo {
		return Repo{}, fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}

	// 2 - get the yaml file into a var/struct
	yaml, err := GetYamlRepo()
	if err != nil {
		return Repo{}, fmt.Errorf("getting the yaml > %w", hostName, err)
	}

	// 2 - look up the requested Repo by name
	repo, ok := yaml.List[i.Name]
	if !ok {
		return Repo{}, fmt.Errorf("%s > repo %q not found in YAML", hostName, i.Name)
	}
	// handle success
	return repo, nil

}

// func (i *Repo) Add(hostName, helmHost string, logger logx.Logger) error {

// 	// 1 - lookup this repo into the yaml
// 	repo, err := i.getRepoFromYaml(hostName)
// 	if err != nil {
// 		return fmt.Errorf("%s:%s > getting repo: maybe it is not in the whitelist:%w", hostName, helmHost, err)
// 	}

// 	// 2 - set an instance property extracted from the yaml
// 	i.Url = repo.Url

// 	// 3 - get and play cli
// 	if _, err := run.RunCli(helmHost, i.cliToAdd(), logger); err != nil {
// 		return err
// 	}
// 	// handle success
// 	logger.Debugf("%s:%s:%s > added helm repo", hostName, helmHost, i.Name)
// 	return nil
// }

// description: get a repo in the yaml
func (i *Repo) getRepoFromYaml(hostName string) (Repo, error) {

	// 2 - get the yaml file into a var/struct
	yaml, err := GetYamlRepo()
	if err != nil {
		return Repo{}, fmt.Errorf("getting the yaml > %w", hostName, err)
	}

	// 2 - look up the requested Repo by name
	repo, ok := yaml.List[i.Name]
	if !ok {
		return Repo{}, fmt.Errorf("%s > repo %q not found in YAML", hostName, i.Name)
	}
	// handle success
	return repo, nil

}

// description: get the printable string of the whitelist yaml
func (repoService) GetWhitelist(hostName string) (string, error) {
	// 1 - get the yaml file into a var/struct
	YamlStruct, err := GetYamlRepo()
	if err != nil {
		return "", fmt.Errorf("%s > getting the yaml > %w", hostName, err)
	}
	// 2 - return it as a printable string
	return YamlStruct.ConvertToString(), nil
}

//	func (repoService) List(hostName string, helmHost string, logger logx.Logger) (string, error) {
//		// 1 - get and play cli
//		out, err := run.RunCli(helmHost, RepoSvc.cliToList(), logger)
//		if err != nil {
//			return "", fmt.Errorf("%s:%s > listing helm repos > %w", hostName, helmHost, err)
//		}
//		// 	// handle success
//		return out, nil
//	}
func (i *Repo) DeleteAll(hostName, helmHost string, logger logx.Logger) error {
	// 1 - get and play cli
	if _, err := run.RunCli(hostName, i.cliToDeleteAll(), logger); err != nil {
		return err
	}
	// handle success
	logger.Debugf("%s:%s > deleted all helm repos", hostName, helmHost)
	return nil
}

// func (i *Repo) Delete(hostName, helmHost string, logger logx.Logger) (string, error) {
// 	// 1 - get and play cli
// 	output, err := run.RunCli(helmHost, i.cliToDelete(), logger)
// 	if err != nil {
// 		return "", fmt.Errorf("%s:%s:%s > deleting helm rpo from the helm client > %w", hostName, helmHost, i.Name, err)
// 	}
// 	// handle success
// 	logger.Debugf("%s:%s > deleted helm repo: %s", hostName, helmHost, i.Name)
// 	return output, nil
// }

func (i *Repo) cliToDeleteAll() string {
	var cmds = []string{
		`rm -rf ~/.config/helm/repositories.yaml`,
		`rm -rf ~/.cache/helm/repository`,
		`rm -rf ~/.local/share/helm`,
	}
	cli := strings.Join(cmds, " && ")
	return cli

}
