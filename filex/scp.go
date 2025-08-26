package filex

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/errorx"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/run"
)

// Name: ScpAsSudo
//
// Description: Scp file between local and a remote Vm using the `scp` command with sudo privileges.
//
// Parameters:
//
//	l: The logger to use for command output.
//	source: The source path (local or remote).
//	destination: The destination path (local or remote).
//
// Returns:
//
//	bool: True if the file transfer was successful.
//	error: An error if the transfer failed, including wrapped context.
//
// Prerequisites:
//
//	- The function assume the user running this function is sudo on the remote
//
// Notes:
//
// - It is designed for creating files in privileged locations, such as `/usr/local/bin` or `/etc`.
//
// Todos:
//
// - Create a function in gocore that check if a user is sudo on the remote
// - and uqe it here to avoid prerequisites.

func ScpAsSudo(l logx.Logger, source, destination string) (bool, error) {
	// // Step 1: Check if the user has administrative privileges.
	// canSudo, err := user.CanBeSudoAndIsNotRoot()
	// if err != nil {
	// 	// handle generic error explicitly: unexpected failure
	// 	return false, errorx.Wrap(err, "failed to check for sudo privileges")
	// }
	// if !canSudo {
	// 	// handle specific error explicitly: expected outcome
	// 	return false, errorx.New("current user does not have sudo privileges")
	// }

	// parsing the destination
	parts := strings.Split(destination, ":")
	if len(parts) != 2 {
		return false, errorx.New("invalid remote destination format: expected 'vmName:path'")
	}
	vmName := parts[0]
	vmFullPath := parts[1]

	// We use `sh -c` to ensure that `sudo` is applied to the entire command string.
	command := fmt.Sprintf("cat %s | ssh %s 'sudo tee %s && sudo chmod +x %s'", source, vmName, vmFullPath, vmFullPath)

	l.Infof("Initiating sudo SCP transfer from %s to %s", source, destination)

	output, err := run.RunOnLocal(command)
	if err != nil {
		// handle generic error explicitly: unexpected failure
		l.Error(strings.TrimSpace(output)) // Log the raw output for debugging
		return false, errorx.Wrap(err, "failed to run scp command with sudo")
	}

	l.Info("sudo SCP transfer completed successfully.")

	// success
	return true, nil
}
