package file

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func ForceCreateRcFile(hostName, nodeName, customRcFileName string, logger logx.Logger) error {
	// 11 - get an instance
	i := GetFile(customRcFileName, "~", "")

	// 12 - operate
	if err := i.ForceCreateRcFile(hostName, nodeName, logger); err != nil {
		return fmt.Errorf("creating rc file %s > %w", i.FullPath, err)
	}

	return nil
}

// description: create a new empty RC file even if it exists
func (i *File) ForceCreateRcFile(hostName string, nodeName string, logger logx.Logger) error {

	// 1 - get cli
	cli := i.cliToCreateEmptyFile(true)

	// 2 - play CLI - that return the created file name
	CreatedFilePath, err := run.RunCli(nodeName, cli, logger)
	if err != nil {
		return fmt.Errorf("%s:%s > creating rc file %s > %w", hostName, nodeName, i.FullPath, err)
	}

	// 3 - get user std rc file from it
	rcStdFile := GetFile("", "", filepath.Dir(CreatedFilePath)+"/.profile")

	// // 4 - operate
	content := fmt.Sprintf("source %s", strings.TrimSpace(CreatedFilePath))
	rcStdFile.AddStringOnce(hostName, nodeName, content, logger)
	// log
	logger.Infof("%s:%s > added a line to user's rc file: %s", hostName, nodeName, rcStdFile.FullPath)

	// get cli
	// play cli
	// handle success
	return nil
}

// description: TODO: create a new empty RC file if it not exists. If file exists: emit an error
func (i *File) CreateRcFile(hostName string, nodeName string, logger logx.Logger) error {
	return nil
}

func RcAddPath(hostName, nodeName, folderRootPath, customRcFileName string, logger logx.Logger) error {
	// log
	// handle success
	return nil
}
