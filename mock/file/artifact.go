package file

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func CopyArtifactToDest(hostName string, artifactPath string, dstFullPath string, artifactType string, logger logx.Logger) error {
	// define var
	var out string
	// get cli
	cli, err := cliForCopyArtifact(artifactPath, dstFullPath, artifactType, logger)
	if err != nil {
		return err
	}
	// play cli
	if out, err := run.RunCli(hostName, cli, logger); err != nil {
		return fmt.Errorf("%s > copying artifact %s to %s > err : %w > out:%s", hostName, artifactPath, dstFullPath, err, out)
	}
	// log
	logger.Debugf("%s > cli is %s > out is -%s-", hostName, cli, out)
	// handle success
	return nil
}

func cliForCopyArtifact(artifactPath, dstArtifact, artifactType string, logger logx.Logger) (string, error) {
	var clis []string
	var src = strings.TrimSpace(artifactPath)
	var dst = strings.TrimSpace(dstArtifact)

	switch artifactType {
	case "exe":
		clis = []string{
			fmt.Sprintf(`sudo cp %s %s > /dev/null`, src, dst),
			fmt.Sprintf(`sudo chmod +x %s`, dst),
		}

	case "tgz":
		// log
		clis = []string{
			// fmt.Sprintf(`nbFolder=\$(tar -tzf %q | awk -F/ 'NF>1 {print \$1"/"}' | sort -u | wc -l)`, src),
			// fmt.Sprintf(`folderDepth=1`),
			fmt.Sprintf(`folderDepth=\$(tar -tzf %s | head -1 | awk -F/ '{print NF-1}')`, src),
			fmt.Sprintf(`sudo mkdir -p %s`, dst),
			fmt.Sprintf(`sudo tar -xvzf %s -C %s --strip-components=\$folderDepth`, src, dst),
			`echo nbFolder=\${nbFolder}`,
		}

	default:
		return "", fmt.Errorf("not supported artefact type: %s (artefact: %s)", artifactType, artifactPath)
	}

	// define cli
	cli := strings.Join(clis, " && ")

	return cli, nil
}

func (i *Artifact) Download(hostName string, logger logx.Logger) (string, error) {
	// get cli and play it
	cli := i.cliToDownload()
	logger.Debugf("%s > executed cli is : %s", hostName, cli)
	artifactFullPath, err := run.RunCli(hostName, cli, logger)
	if err != nil {
		return "", fmt.Errorf("%s > downloading artifact from url %s > %w ", hostName, i.Url, err)
	}
	// set artifact full path
	i.FullPath = strings.TrimSpace(artifactFullPath)
	// handle success
	return artifactFullPath, nil
}
func DownloadArtifact(hostName string, url string, prefix string, extension string, logger logx.Logger) (string, error) {
	// get cli and play it
	cli := cliForDownload(url, prefix, extension)
	// play cli
	artifactFullPath, err := run.RunCli(hostName, cli, logger)
	if err != nil {
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
func (i *Artifact) cliToDownload() string {
	// define cli
	var clis = []string{
		fmt.Sprintf(`tmp=$(mktemp /tmp/%s-XXXXXX)`, i.Name),
		`touch "\\$tmp"`,
		// fmt.Sprintf(`mv $tmp $tmp.%s`, i.Type),
		// fmt.Sprintf(`curl -fL %s -o $tmp &> /dev/null`, i.Url),
		// `echo $tmp`,
	}
	cli := strings.Join(clis, " && ")

	// handle success
	return cli
}

// fmt.Sprintf(`tmp=$(mktemp /tmp/%s-XXXXXX)`, i.Name),
// fmt.Sprintf(`mv $tmp $tmp.%s`, i.Type),
// fmt.Sprintf(`curl -fL %s -o $tmp &> /dev/null`, i.Url),
