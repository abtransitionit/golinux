package file

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

func CopyArtifactToDest(hostName string, artifact string, dst string, artifactType string, logger logx.Logger) error {
	// get cli
	cli := cliForCopyArtifact(artifact, dst, artifactType)
	// // play cli
	// if out, err := run.RunCli(hostName, cli, logger); err != nil {
	// 	return fmt.Errorf("%s > downloading file from url %s > err > %w > out:%s", hostName, url, err, out)
	// }
	// log
	logger.Debugf("%s > copy %s to %s (cli: %.15s)", hostName, artifact, dst, cli)
	// handle success
	return nil
}

func cliForCopyArtifact(artefact, dst, artefactType string) string {
	// define cli
	var clis = []string{
		fmt.Sprintf(`do the job with %s, %s, %s`, artefact, dst, artefactType),
	}
	cli := strings.Join(clis, " && ")

	// handle success
	return cli
}
