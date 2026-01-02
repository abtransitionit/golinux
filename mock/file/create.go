package file

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// description: creates a file from the specified multiline string in a sudo location
func SudoCreateFileFromString(filePath, content string) string {
	// escapedContent := strings.ReplaceAll(content, `'`, `'\''`)
	// Wrap in single quotes and use printf to preserve newlines
	var cmds = []string{
		fmt.Sprintf("printf '%s' | sudo tee %q > /dev/null", content, filePath),
	}
	cli := strings.Join(cmds, " && ")

	return cli
}

func SudoCreateGpgFileFromUrl(url string, filePath string) string {
	var cmds = []string{
		// fmt.Sprintf(`sudo install -d -m 0755  $(dirname %s)`, filePath),
		"set -o pipefail",
		fmt.Sprintf(`curl -fsSL %s | gpg --dearmor | sudo tee %s`, url, filePath),
		fmt.Sprintf(`sudo chmod 0644  %s`, filePath),
	}
	cli := strings.Join(cmds, " && ")
	return cli
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

// description: CLI to create an empty file
//
// Notes:
//
// - if force = true : create the file even if it exists
// - if force = false: create the file only if it not exists
func (i *File) cliToCreateEmptyFile(force bool) string {
	var firstCmd string

	if force {
		firstCmd = fmt.Sprintf("echo > %s", i.FullPath)
	} else {
		firstCmd = fmt.Sprintf("touch %s", i.FullPath)
	}

	clis := []string{
		firstCmd,
		fmt.Sprintf(`echo %s`, i.FullPath),
	}
	cli := strings.Join(clis, " && ")
	return cli

}

// description: TODO: create a new empty RC file if it not exists. If file exists: emit an error
func (i *File) CreateRcFile(hostName string, nodeName string, logger logx.Logger) error {
	return nil
}
