package file

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func DownloadArtifact(hostName string, url string, prefix string, extension string, logger logx.Logger) (string, error) {
	var artifactFullPath string
	var err error
	// get cli
	cli := cliForDownload(url, prefix, extension)
	// play cli
	if artifactFullPath, err = run.RunCli(hostName, cli, logger); err != nil {
		return "", fmt.Errorf("%s > downloading file from url %s > err > %w > out:%s", hostName, url, err, artifactFullPath)
	}
	// handle success
	return artifactFullPath, nil
}

func cliForDownload(url, prefix, extension string) string {
	// define cli
	var clis = []string{
		fmt.Sprintf(`tmp=\$(mktemp /tmp/%s-XXXXXX.%s)`, prefix, extension),
		fmt.Sprintf(`curl -fL %q -o "\$tmp" &> /dev/null`, url),
		`echo "\$tmp"`,
	}
	cli := strings.Join(clis, " && ")

	// handle success
	return cli
}
