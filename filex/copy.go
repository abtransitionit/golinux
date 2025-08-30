package filex

import (
	"context"
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
//   - The function assume the user running this function is sudo on the remote
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

	command := fmt.Sprintf("cat %s | ssh %s 'sudo tee %s && sudo chmod +x %s'", source, vmName, vmFullPath, vmFullPath)

	l.Debugf("Initiating sudo SCP transfer from %s to %s", source, destination)

	output, err := run.RunOnLocal(command)
	if err != nil {
		// handle generic error explicitly: unexpected failure
		l.Error(strings.TrimSpace(output)) // Log the raw output for debugging
		return false, errorx.Wrap(err, "failed to run scp command with sudo")
	}

	l.Debugf("sudo SCP transfer completed successfully on %s", vmName)

	// success
	return true, nil
}

// Name: CpFileAsSudo
//
// Description: Copies a local file to a root location on the same host
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
//   - The function assume the user running this function is sudo
func CpAsSudo(ctx context.Context, l logx.Logger, source, destination string) (string, error) {
	var clis = []string{
		fmt.Sprintf("sudo cp %s %s", source, destination),
		fmt.Sprintf("sudo chmod +x %s", destination),
	}

	cli := strings.Join(clis, " && ")
	return cli, nil

}
func CpAsSudo2(ctx context.Context, l logx.Logger, source, destination string) (bool, error) {
	// Build the command
	// -p preserves mode, ownership, and timestamps
	// command := fmt.Sprintf("sudo cp -p %s %s", source, destination)
	command := fmt.Sprintf("sudo cp %s %s", source, destination)
	l.Debugf("Initiating sudo copy from %s to %s", source, destination)

	_, err := run.RunOnLocal(command)
	if err != nil {
		// l.Error(strings.TrimSpace(output)) // Log the raw output for debugging
		return false, errorx.Wrap(err, "failed to copy file with sudo")
	}

	l.Debugf("sudo copy completed successfully: %s -> %s", source, destination)

	return true, nil
}

// l.Debugf("Initiating sudo copy from %s to %s", source, destination)

// _, err := run.RunOnLocal(command)
// if err != nil {
// 	// l.Error(strings.TrimSpace(output)) // Log the raw output for debugging
// 	return false, errorx.Wrap(err, "failed to copy file with sudo")
// }

// l.Debugf("sudo copy completed successfully: %s -> %s", source, destination)

// return true, nil
