package release

import (
	"context"
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

func listRelease(ctx context.Context, logger logx.Logger) (string, error) {
	logger.Info("To implement: Manage Helm Releases")
	logger.Info(`A [Helm] release is an instanciation of a [Helm] chart. It consists of:
	→ a release name.
	→ a K8s namespave.
	→ a set of values`)
	fmt.Sprintf(`helm list -A`)
	return "", nil
}
