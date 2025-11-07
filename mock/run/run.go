package run

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

// ExecuteCli executes a CLI command locally and returns its output as a string.
func ExecuteCli(cli string, logger logx.Logger) (string, error) {
	// Split command into name and args (simple approach, works for basic commands)
	parts := strings.Fields(cli)
	if len(parts) == 0 {
		return "", nil
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%v: %s", err, stderr.String())
	}

	// Trim whitespace and return
	return strings.TrimSpace(out.String()), nil
}
