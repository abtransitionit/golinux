package onpm

import (
	"fmt"
)

func (m *DnfManager) CliList() (string, error) {
	// repo  sudo dnf repolist -v
	if m.Repo != nil {
		switch m.Conf.Os.Distro {
		case "fedora":
			return "dnf repolist", nil
		default:
			return "dnf repolist -v", nil
		}
	}
	// package
	return fmt.Sprintf("dnf list installed %s", m.Pkg.Name), nil
}

func (m *DnfManager) CliAdd() (string, error) {
	// repo
	if m.Repo != nil {
		return fmt.Sprintf("dnf config-manager --add-repo %s", m.Repo.Url), nil
	}
	// package
	return fmt.Sprintf("dnf install -y %s", m.Pkg.Name), nil
}

func (m *DnfManager) CliDelete() (string, error) {
	// repo
	if m.Repo != nil {
		return fmt.Sprintf("dnf config-manager --remove-repo %s", m.Repo.Name), nil
	}
	// package
	return fmt.Sprintf("dnf remove -y %s", m.Pkg.Name), nil
}

// func (m *DnfManager) Add(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
// 	cli, _ := m.CliAdd()
// 	return run.ExecuteCliQuery(cli, logger, local, remoteHost, run.NoOpErrorHandler)
// }

// func (m *DnfManager) Delete(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
// 	cli, _ := m.CliDelete()
// 	return run.ExecuteCliQuery(cli, logger, local, remoteHost, run.NoOpErrorHandler)
// }

// func (m *DnfManager) List(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
// 	cli, _ := m.CliAdd()
// 	return run.ExecuteCliQuery(cli, logger, local, remoteHost, run.NoOpErrorHandler)
// }
