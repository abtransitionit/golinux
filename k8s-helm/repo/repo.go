package repo

import (
	"context"

	"github.com/abtransitionit/luc/pkg/logx"
)

func addRepo(ctx context.Context, logger logx.Logger, repoName string) (string, error) {
	logger.Info("To implement: add repo")
	return "", nil
}
func deleteRepo(ctx context.Context, logger logx.Logger, repoName string) (string, error) {
	logger.Info("To implement: delete repo")
	return "", nil
}
func listRepo(ctx context.Context, logger logx.Logger) (string, error) {
	logger.Info("To implement: list repo")
	return "", nil
}
