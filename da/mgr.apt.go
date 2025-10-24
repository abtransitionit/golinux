package da

import (
	"fmt"
)

func (m *AptManager) CliList() (string, error) {
	// repo - cat /etc/apt/sources.list.d/*.sources
	if m.Repo != nil {
		return "cat /etc/apt/sources.list /etc/apt/sources.list.d/*.list /etc/apt/sources.list.d/*.sources 2>/dev/null | grep -E '^deb |^Components:'", nil
	}
	// package
	return fmt.Sprintf("dpkg -l | grep %s", m.Pkg.Name), nil
}

func (m *AptManager) CliAdd() (string, error) {
	// repo
	if m.Repo != nil {
		return fmt.Sprintf("add-apt-repository '%s'", m.Repo.Url), nil
	}
	// package
	return fmt.Sprintf("apt-get install -y %s", m.Pkg.Name), nil
}

func (m *AptManager) CliDelete() (string, error) {
	// repo
	if m.Repo != nil {
		return fmt.Sprintf("add-apt-repository --remove '%s'", m.Repo.Url), nil
	}
	// package
	return fmt.Sprintf("apt-get remove -y %s", m.Pkg.Name), nil
}

// func (m *AptManager) Add(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
// 	cli, _ := m.CliAdd()
// 	return run.ExecuteCliQuery(cli, logger, local, remoteHost, run.NoOpErrorHandler)
// }

// func (m *AptManager) Delete(ctx context.Context, local bool, remoteHost string, logger logx.Logger) (string, error) {
// 	cli, _ := m.CliDelete()
// 	return run.ExecuteCliQuery(cli, logger, local, remoteHost, run.NoOpErrorHandler)
// }
