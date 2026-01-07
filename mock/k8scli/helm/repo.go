package helm

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func (i *Repo) Add(hostName, helmNode string, logger logx.Logger) error {
	// 1 - get the yaml
	YamlStruct, err := getYaml(hostName)
	if err != nil {
		return fmt.Errorf("%s > %w", hostName, err)
	}
	// 2 - get the repo from the yaml
	repo, err := i.getYamlRepo(hostName, YamlStruct, logger)
	if err != nil {
		return fmt.Errorf("%s:%s > getting repo: %w", hostName, helmNode, err)
	}

	// log
	logger.Debugf("%s:%s > repo is %+v", hostName, helmNode, repo)
	// 3 - set the URL from the yaml
	i.Url = repo.Url

	// log
	logger.Debugf("%s:%s > instance url is %s", hostName, helmNode, i.Url)

	// 1 - get and play cli
	i.cliToAdd()
	logger.Debugf("CLI  > %s", i.cliToAdd())
	if _, err := run.RunCli(helmNode, i.cliToAdd(), logger); err != nil {
		return err
	}
	// handle success
	logger.Debugf("%s:%s:%s > added helm repo", hostName, i.Name, helmNode)
	return nil
}

// description: get the raw url of a cli from the yaml
func (i *Repo) getYamlRepo(hostName string, yaml *MapYaml, logger logx.Logger) (Repo, error) {

	// 2 - look up the requested Repo by name
	repo, ok := yaml.List[i.Name]
	if !ok {
		return Repo{}, fmt.Errorf("%s > repo %q not found in YAML", hostName, i.Name)
	}
	// handle success
	return repo, nil

}

func (i *Repo) List(hostName, helmNode string, logger logx.Logger) error {
	// 1 - get and play cli
	if _, err := run.RunCli(hostName, i.cliToList(), logger); err != nil {
		return err
	}
	// handle success
	logger.Debugf("%s:%s > listed all helm repos", hostName, helmNode)
	return nil
}
func (i *Repo) DeleteAll(hostName, helmNode string, logger logx.Logger) error {
	// 1 - get and play cli
	if _, err := run.RunCli(hostName, i.cliToDeleteAll(), logger); err != nil {
		return err
	}
	// handle success
	logger.Debugf("%s:%s > deleted all helm repos", hostName, helmNode)
	return nil
}

func (i *Repo) Delete(hostName, helmNode string, logger logx.Logger) error {
	// 1 - get and play cli
	if _, err := run.RunCli(hostName, i.cliToDelete(), logger); err != nil {
		return err
	}
	// handle success
	logger.Debugf("%s:%s > deleted helm repo: %s", hostName, helmNode, i.Name)
	return nil
}

func (i *Repo) ListChart(hostName, helmNode string, logger logx.Logger) error {
	// 1 - get and play cli
	if _, err := run.RunCli(hostName, i.cliToListChart(), logger); err != nil {
		return err
	}
	// handle success
	logger.Debugf("%s:%s > listed all charts of repo : %s", hostName, i.Name, helmNode)
	return nil
}

func (i *Repo) cliToAdd() string {
	var cmds = []string{
		`. ~/.profile`,
		fmt.Sprintf(`helm repo add %s %s`, i.Name, i.Url),
		`helm repo update`,
	}
	cli := strings.Join(cmds, " && ")
	return cli

}

func (i *Repo) cliToDelete() string {
	var cmds = []string{
		fmt.Sprintf(`helm repo remove %s`, i.Name),
	}
	cli := strings.Join(cmds, " && ")
	return cli

}

func (i *Repo) cliToDeleteAll() string {
	var cmds = []string{
		`rm -rf ~/.config/helm/repositories.yaml`,
		`rm -rf ~/.cache/helm/repository`,
		`rm -rf ~/.local/share/helm`,
	}
	cli := strings.Join(cmds, " && ")
	return cli

}

// Returns the cli to list all repositories
func (i *Repo) cliToList() string {
	var cmds = []string{
		`helm repo list`,
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

// Returns the cli to list the chart in a repo
func (i *Repo) cliToListChart() string {
	var cmds = []string{
		fmt.Sprintf(`helm search repo %s`, i.Name),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
