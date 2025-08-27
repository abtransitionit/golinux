package cli

import (
	"fmt"
	"os"

	"github.com/abtransitionit/gocore/errorx"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/filex"
)

// Name: DeployGoArtifact
//
// Description: Deploys a Go artifact to a rlocal or remote destination with sudo privileges.
//
// Parameters:
//
//	l: The logger to use.
//	artifactPath: The full path to the local executable to be deployed (e.g., "/tmp/goluc-linux").
//	remoteDestination: The remote destination path, including host and file (e.g., "o1u:/usr/local/bin/goluc").
//
// Returns:
//
//	bool: true if the deployment was successful.
//	error: An error if the deployment failed.
func DeployGoArtifactAsSudo(logger logx.Logger, artifactPath, remoteDestination string) (bool, error) {
	// check parameters
	if artifactPath == "" {
		return false, fmt.Errorf("artifact path is empty")
	}
	if remoteDestination == "" {
		return false, fmt.Errorf("remote destination is empty")
	}

	// check if the artifact exists.
	if _, err := os.Stat(artifactPath); os.IsNotExist(err) {
		return false, errorx.Wrap(err, "local artifact not found at path: %s", artifactPath)
	}

	// scp the artifact to a root remote location - if dst path is not root use filex.Scp
	copyOk, err := filex.ScpAsSudo(logger, artifactPath, remoteDestination)
	if err != nil {
		logger.Errorf("%v", err)
		return false, err
	}

	if !copyOk {
		logger.Error("artifact deployment failed")
		return false, err
	}

	// sucess
	// logger.Debugf("Successfully deployed artifact to %s", remoteDestination)
	return true, nil
}
