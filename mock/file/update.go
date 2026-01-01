package file

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// description: sudo copy a file from source to destination using if different orn if not exists on dest

func (i *File) AddStringOnce(hostName, nodeName, content string, logger logx.Logger) error {
	// log
	// get cli
	cli := i.cliToEnsureLineExists(content)
	// logger.Infof("%s:%s > Add string %s to file : %s", hostName, nodeName, content, i.FullPath)
	logger.Infof("%s:%s > play cli : %s", hostName, nodeName, cli)
	// play cli
	_, err := run.RunCli(nodeName, cli, logger)
	if err != nil {
		return fmt.Errorf("%s:%s > adding line to file %s > %w", hostName, nodeName, i.FullPath, err)
	}
	// handle success
	return nil
}

func (i *File) cliToEnsureLineExists(line string) string {
	var cmds = []string{
		fmt.Sprintf(`grep -qxF %q %q || echo %q >> %q`, line, i.FullPath, line, i.FullPath),
	}

	cli := strings.Join(cmds, " && ")
	return cli
}
