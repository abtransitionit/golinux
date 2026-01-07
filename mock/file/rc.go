package file

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/property"
	"github.com/abtransitionit/golinux/mock/run"
	"github.com/abtransitionit/golinux/mock/util"
)

func ForceCreateRcFile(hostName, nodeName, customRcFileName string, logger logx.Logger) error {
	// 11 - get an instance
	i := GetFile(customRcFileName, "~", "")

	// 12 - operate
	if err := i.ForceCreateRc(hostName, nodeName, logger); err != nil {
		return fmt.Errorf("creating rc file %s > %w", i.FullPath, err)
	}

	return nil
}

// description: create a new empty RC file even if it exists
func (i *File) ForceCreateRc(hostName string, nodeName string, logger logx.Logger) error {

	// 1 - get cli
	cli := i.cliToCreateEmptyFile(true)

	// 2 - play CLI - that return the created file name
	CreatedFilePath, err := run.RunCli(nodeName, cli, logger)
	if err != nil {
		return fmt.Errorf("%s:%s > creating rc file %s > %w", hostName, nodeName, i.FullPath, err)
	}

	// 3 - get user std rc file from it
	// rcStdFile := GetFile("", "", filepath.Dir(CreatedFilePath)+"/.profile")
	rcStdFile := GetFile(".profile", filepath.Dir(CreatedFilePath), "")

	// // 4 - operate
	content := fmt.Sprintf("source %s", strings.TrimSpace(CreatedFilePath))
	if err := rcStdFile.AddStringOnce(hostName, nodeName, content, logger); err != nil {
		return fmt.Errorf("%s:%s > adding line to file %s > %w", hostName, nodeName, rcStdFile.FullPath, err)
	}
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
	var (
		treePath string
		path     string
		newPath  string
		err      error
	)

	// 1 - get tree path
	if treePath, err = property.GetProperty(logger, nodeName, "pathTree", folderRootPath); err != nil {
		return fmt.Errorf("%s:%s : getting tree path from %s > %w", hostName, nodeName, folderRootPath, err)
	}

	// 2 - get PATH envar
	if path, err = property.GetProperty(logger, nodeName, "envar", "PATH"); err != nil {
		return fmt.Errorf("%s:%s : getting tree path from %s > %w", hostName, nodeName, folderRootPath, err)
	}
	logger.Debugf("%s:%s : current PATH is %s", hostName, nodeName, path)

	// 3 - get new PATH
	cli := util.CliToFusionString(path, treePath, ":")
	if newPath, err = run.RunCli(nodeName, cli, logger); err != nil {
		return fmt.Errorf("%s:%s : creating new PATH from %s and %s > %w", hostName, nodeName, path, treePath, err)
	}

	// 4 - persist this new PATH into the user's custom RC file
	// 41 - get an instance
	i := GetFile(customRcFileName, "~", "")

	content := fmt.Sprintf(`export PATH=%s`, newPath)
	if err := i.AddStringOnce(hostName, nodeName, content, logger); err != nil {
		return fmt.Errorf("%s:%s > adding line to file %s > %w", hostName, nodeName, i.FullPath, err)
	}

	// log
	logger.Debugf("%s:%s : persisted new path (%s) to user's custom rc file: %s", hostName, nodeName, newPath, i.FullPath)

	// handle success
	return nil
}
