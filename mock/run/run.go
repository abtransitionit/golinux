package run

import (
	"fmt"
	"os/exec"

	"github.com/abtransitionit/gocore/logx"
)

func RunOnLocal(cli string, logger logx.Logger) (string, error) {
	cmd := exec.Command("sh", "-c", cli)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("running cli locally: %v, output: %s", err, string(output))
	}
	return string(output), nil
}

// ExecuteCli executes a CLI command locally and returns its output as a string.
// func ExecuteCli(cli string, logger logx.Logger) (string, error) {
// 	// Split command into name and args (simple approach, works for basic commands)
// 	parts := strings.Fields(cli)
// 	if len(parts) == 0 {
// 		return "", nil
// 	}

// 	cmd := exec.Command(parts[0], parts[1:]...)
// 	var out bytes.Buffer
// 	var stderr bytes.Buffer
// 	cmd.Stdout = &out
// 	cmd.Stderr = &stderr

// 	err := cmd.Run()
// 	if err != nil {
// 		return "", fmt.Errorf("%v: %s", err, stderr.String())
// 	}

// 	// Trim whitespace and return
// 	return strings.TrimSpace(out.String()), nil
// }
