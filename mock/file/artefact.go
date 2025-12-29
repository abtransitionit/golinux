package file

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

func CopyArtifactToDest(hostName string, artifactPath string, dstFullPath string, artifactType string, logger logx.Logger) error {
	// get cli
	cli, err := cliForCopyArtifact(artifactPath, dstFullPath, artifactType)
	if err != nil {
		return err
	}
	// // play cli
	// if out, err := run.RunCli(hostName, cli, logger); err != nil {
	// 	return fmt.Errorf("%s > downloading file from url %s > err > %w > out:%s", hostName, url, err, out)
	// }
	// log
	logger.Debugf("%s >%s ", hostName, cli)
	// handle success
	return nil
}

func cliForCopyArtifact(artefact, dst, artefactType string) (string, error) {
	var clis []string

	switch artefactType {
	case "exe":
		clis = []string{
			fmt.Sprintf(`sudo copy exe %s to %s`, artefact, dst),
		}
	case "tgz":
		clis = []string{
			fmt.Sprintf(`sudo copy tgz %s to %s`, artefact, dst),
		}
	default:
		return "", fmt.Errorf("not supported artefact type: %s (artefact: %s)", artefactType, artefact)
	}

	// define cli
	cli := strings.Join(clis, " && ")

	return cli, nil
}
