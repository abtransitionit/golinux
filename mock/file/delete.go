package file

import (
	"fmt"
	"path"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func DeleteFileWithSudo(hostName string, nodeName string, fileProperty FileProperty, logger logx.Logger) (string, error) {
	// define var
	var dstFile string

	// 1 - logic
	if strings.HasSuffix(fileProperty.Dst, "/") {
		// if folder append Src:FileName
		dstFile = path.Join(fileProperty.Dst, path.Base(fileProperty.Src))
	} else {
		// if file just copy Dst
		dstFile = fileProperty.Dst
	}

	// 2 - define CLI
	cmds := []string{
		fmt.Sprintf(`cat %s | ssh %s 'sudo tee %s'`, fileProperty.Src, nodeName, dstFile),
		// fmt.Sprintf(`sudo chmod +x %s`, fileProperty.Dst),
	}
	cli := strings.Join(cmds, " && ")
	// log
	// logger.Infof("%s/%s > sudo copy to %s : %v", hostName, nodeName, dstType, cli)

	// // 3 - run CLI
	output, err := run.RunCli(hostName, cli, nil)
	// handle system error
	if err != nil {
		return "", fmt.Errorf("sudo scping to %s > %w", dstFile, err)

	}

	// handle success
	// return "", nil
	return output, nil
}
