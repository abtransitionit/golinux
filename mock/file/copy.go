package file

import (
	"fmt"
	"path"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// description: sudo copy a file from source to destination if different or if not exists on dest
func CopyFileWithSudo(hostName string, nodeName string, fileProperty FileProperty, logger logx.Logger) (string, error) {
	// define var
	var dstFile string
	var fileChmod string = fileProperty.Chmod // optional

	// 1 - logic
	if strings.HasSuffix(fileProperty.Dst, "/") {
		// if folder append Src:FileName
		dstFile = path.Join(fileProperty.Dst, path.Base(fileProperty.Src))
	} else {
		// if file just copy Dst
		dstFile = fileProperty.Dst
	}
	// 2 - logic - copy file only if different
	// 21 - get the SHA of the source file
	shaSrc, err := run.RunCli(hostName, "sha256sum "+fileProperty.Src, nil)
	if err != nil {
		return "", err
	}

	// 22 - get the SHA of the destination file if exists else difine a default value (needed for comparison)
	shaDst, err := run.RunCli(nodeName, "sha256sum "+dstFile, nil)
	if err != nil {
		shaDst = fmt.Sprintf("11111 %s", dstFile)
	}
	// 23 - extract only the hash (first field) and trim whitespace
	shaSrc = strings.Fields(strings.TrimSpace(shaSrc))[0]
	shaDst = strings.Fields(strings.TrimSpace(shaDst))[0]

	// 24 - return success if the files are the same
	if shaSrc == shaDst {
		// handle success
		return "", nil
	}

	// 3 - define CLI
	cmds := []string{
		fmt.Sprintf(`cat %s | ssh %s 'sudo tee %s > /dev/null'`, fileProperty.Src, nodeName, dstFile),
	}
	// 31 - add chmod
	if fileChmod != "" {
		cmds = append(cmds,
			fmt.Sprintf(`ssh %s 'sudo chmod %s %s'`, nodeName, fileChmod, dstFile),
		)
	}
	cli := strings.Join(cmds, " && ")

	// log
	// logger.Infof("%s/%s > sudo copy to %s : %v", hostName, nodeName, dstType, cli)

	// 4 - run CLI
	output, err := run.RunCli(hostName, cli, nil)
	// handle system error
	if err != nil {
		return "", fmt.Errorf("sudo scping file: %w", err)
	}
	// handle success
	return output, nil
}
