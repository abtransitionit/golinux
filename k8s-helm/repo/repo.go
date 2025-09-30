package repo

import (
	"context"
	"fmt"

	"github.com/abtransitionit/luc/pkg/logx"
)

func addRepo(ctx context.Context, logger logx.Logger, repoName string) (string, error) {
	logger.Info("To implement: add repo")
	fmt.Sprintf(`helm repo add %s %s`, helmRepoName, helmRepoUrl)
	return "", nil
}
func deleteRepo(ctx context.Context, logger logx.Logger, repoName string) (string, error) {
	logger.Info("To implement: delete repo")
	fmt.Sprintf("helm repo remove %s", helmRepoName)
	return "", nil
}
func listRepo(ctx context.Context, logger logx.Logger) (string, error) {
	logger.Info("To implement: list repo")
	fmt.Sprintf(`helm repo list`)
	return "", nil
}
