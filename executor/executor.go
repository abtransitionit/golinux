package executor

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os/exec"
	"strings"

	"github.com/abtransitionit/gocore/errorx"
)

// Name: RunCliSsh
// Description: Executes a command on a remote VM via SSH.
// Purpose: This function is the remote equivalent of RunLocal. It first performs a
// reachability check, then executes a command on the remote host, capturing the output.
// Inputs:
// - vmName:  string: The alias of the VM to connect to (e.g., "o1u").
// - command: string: The command string to be executed on the remote VM.
// Return:
// - string: The standard output and standard error from the remote command.
// - error: An error if the VM is not reachable, or if the SSH command fails.
// Notes:
// - The `-o BatchMode=yes` flag prevents interactive prompts and long waits.
// - The `-o ConnectTimeout=5` flag sets a 5-second timeout, so the function never hang indefinitely. Without this, if the remote VM is powered off or unreachable, the SSH command could hang for several minutes before timing out
// - The remote command is Base64 encoded to avoid issues with complex quotes and special characters.
func RunCliSsh(vmName, cli string) (string, error) {

	// step: check the VM is reachable
	isSshReachable, err := IsVmSshReachable(vmName)
	if err != nil {
		return "", errorx.WithStack(fmt.Errorf("failed to check VM SSH reachability: %w", err))

	}
	if !isSshReachable {
		return "", fmt.Errorf("vm '%s' is not SSH reachable", vmName)
	}

	// step: Base64 encode the input to handle complex quoting and special characters.
	cliEncoded := base64.StdEncoding.EncodeToString([]byte(cli))

	// step: Now that the VM is reachable, define the full SSH command to run - echo the encoded string, decode it, and execute it.
	// command := fmt.Sprintf("ssh -o BatchMode=yes -o ConnectTimeout=5 %s \"echo '%s' | base64 --decode | sh\"", vmName, cliEncoded)
	command := fmt.Sprintf(`ssh -o BatchMode=yes -o ConnectTimeout=5 %s "echo '%s' | base64 --decode | sh"`, vmName, cliEncoded)

	// step: Run the command
	output, err := RunCliLocal(command)

	// manage error
	if err != nil {
		return output, errorx.WithStack(fmt.Errorf("failed to run remote command on '%s': %w", vmName, err))
	}

	// success
	return output, nil
}

// Name: RunCLILocal
// Description: Executes a local command or complex CLI pipeline.
// Purpose: Provides a portable and safe way to execute local commands. It captures
// both standard output and standard error, returning them as a single string.
// This function is the local counterpart to RunSSH.
// Inputs:
// - command: string: The complete command string to be executed (e.g., "ls -la" or "ssh -G myhost | grep hostname").
// Return:
// - string: The combined standard output and standard error from the command, with leading/trailing whitespace removed.
// - error: An error if the command fails to run or exits with a non-zero status.
// ...
// Notes:
// - Uses `sh -c` to ensure complex commands with pipes and redirects execute correctly.
// - Captures both standard output and standard error.
// - Trims leading/trailing whitespace from the final output.
func RunCliLocal(command string) (string, error) {
	// cmd := exec.Command("bash", "-c", command)
	cmd := exec.Command("sh", "-c", command)

	// config: capture both standard output and standard error into a single buffer.
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	// Run the command and wait for it to finish: This will return a non-nil error if the command exits with a non-zero status.
	err := cmd.Run()
	output := strings.TrimSpace(out.String())

	// manage error
	if err != nil {
		return output, errorx.WithStack(fmt.Errorf("command failed: %w, output: %s", err, out.String()))
	}

	// success
	return output, nil
}
