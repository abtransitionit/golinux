package kubectl

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

// description: add a pvc
//
// Notes:
//
// - create a default StorageClass and allow PVCs.
// - create a simple local-path-storage
func (i Resource) AddPvc(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "added pvc as local-path-storage", i.cliToAddPvc(), logger)
}

func (i Resource) cliToAddPvc() string {
	// check
	if i.Type != ResPvc {
		panic("resource type not supported for this cli: " + i.Type)
	}
	yaml := "https://raw.githubusercontent.com/rancher/local-path-provisioner/master/deploy/local-path-storage.yaml"
	return fmt.Sprintf(`kubectl apply -f %s`, yaml)
}
