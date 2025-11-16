package git

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// Description: git merge branch dev to main and push to github
func MergeDevToMain(targetName string, repoName string, logger logx.Logger) (bool, error) {

	// define var
	const repoFolder = "/Users/max/wkspc/git"
	repoPath := fmt.Sprintf("%s/%s", repoFolder, repoName)

	// 1 - build CLI to merge dev to main
	cmds := []string{
		fmt.Sprintf("cd %s", repoPath),
		"git add .",
		"git diff --cached --quiet || git commit -m 'update'",
		"git push origin dev",
		"git checkout main",
		"git merge --no-edit dev",
		"git push origin main",
		"git checkout dev",
	}
	cli := strings.Join(cmds, " && ")

	// 2 - run CLI
	_, err := run.RunCli(targetName, cli, logger)

	// 3 - handle system error
	if err != nil {
		return false, fmt.Errorf("target: %s > node: %s > system error > merging and pushing dev to main: %w", targetName, repoName, err)
	}

	// 4 - handle success
	return true, nil
}
