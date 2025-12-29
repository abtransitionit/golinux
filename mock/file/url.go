package file

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

func DownloadArtifact(hostName string, url string, prefix string, extension string, logger logx.Logger) error {
	// get cli
	cli := cliForDownload(url, prefix, extension)
	// log
	logger.Infof("%s > will play cli %s", hostName, cli)
	// play cli
	out, err := run.RunCli(hostName, cli, logger)
	if err != nil {
		return fmt.Errorf("%s > downloading file from url %s > err > %w > out:%s", hostName, url, err, out)
	}
	// log
	// logger.Infof("%s > download %s with cli %s", hostName, url, cli)
	// install cli
	return nil
}

// func DownloadArtifact2(vmName string, url string, prefix string) (string, error) {
// 	// Remote download
// 	if vmName != "local" {
// 		cmd := fmt.Sprintf("goluc do download %s -p %s", url, prefix)
// 		fileName, err := run.RunCliSsh(vmName, cmd)
// 		if err != nil {
// 			return "", fmt.Errorf("failed to remote download file at URL: '%s' from '%s': %w", url, vmName, err)
// 		}
// 		return strings.TrimSpace(fileName), nil
// 	}

// 	// Local download
// 	// request the object pointed by the URL
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to GET %s: %w", url, err)
// 	}

// 	// clean before exit
// 	defer resp.Body.Close()

// 	// Check request status - Only 200 OK is considered successful - Any other code (404, 500, etc.) triggers an error.
// 	if resp.StatusCode != http.StatusOK {
// 		return "", fmt.Errorf("bad status %d when fetching %s", resp.StatusCode, url)
// 	}

// 	// Create a temporary file - "gocli-*" → default to a random chars to make it uniq
// 	tmpFile, err := os.CreateTemp("/tmp", prefix+"-*")
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create temp file: %w", err)
// 	}
// 	// clean before exit
// 	defer tmpFile.Close()

// 	// Stream response into the temp file
// 	_, err = io.Copy(tmpFile, resp.Body)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to write to temp file: %w", err)
// 	}

// 	return tmpFile.Name(), nil
// }

func cliForDownload(url, prefix, extension string) string {
	// define cli
	var clis = []string{
		fmt.Sprintf(`tmp=\$(mktemp /tmp/%s-XXXXXX)`, prefix),
		fmt.Sprintf(`curl -fL %q -o "\$tmp.%s"`, url, extension),
		// `echo "\$tmp"`,
	}
	cli := strings.Join(clis, " && ")

	// handle success
	return cli
}
