package filex

import (
	"fmt"
	"os"

	"github.com/abtransitionit/gocore/errorx"
	"github.com/abtransitionit/golinux/executor"
	"github.com/abtransitionit/golinux/user"
)

// Name: TouchAsSudo
//
// Description:
//
// Touch a file using `sudo`.
// Creates a file at the given path by executing the `touch` command with `sudo` privileges.
// This function first verifies if the current user has `sudo` permissions without requiring
// a password and then proceeds with the file creation. It is a Linux-specific operation.
//
// Parameters:
//
//  filePath: string: The absolute path to to be created.
//
// Returns:
//
// - bool: `true` if the file was successfully created, `false` otherwise.
// - error: An error if the user lacks privileges or the command fails for another reason.
//
// Notes:
//
// - It is designed for creating files in privileged locations, such as `/usr/local/bin` or `/etc`.

func TouchAsSudo(filePath string) (bool, error) {
	// Step 1: Check if the user has administrative privileges.
	canSudo, err := user.CanBeSudoAndIsNotRoot()
	if err != nil {
		// handle generic error explicitly: unexpected failure
		return false, errorx.Wrap(err, "failed to check for sudo privileges")
	}
	if !canSudo {
		// handle specific error explicitly: expected outcome
		return false, errorx.New("current user does not have sudo privileges")
	}

	// Step 2: Attempt to create the file
	command := fmt.Sprintf("sudo touch %s", filePath)
	_, err = executor.RunCliLocal(command)
	if err != nil {
		// handle generic error explicitly: unexpected failure (e.g., file path is invalid, permissions issues, etc.).
		return false, errorx.Wrap(err, "failed to create file at %s as sudo", filePath)
	}

	// success
	return true, nil
} // Name: DeleteAsSudo
//
// Description: Deletes a file at the specified path using `sudo`.
//
// Parameters:
// - filePath: string: The full path to the file to be deleted.
//
// Returns:
// - bool: `true` if the file was successfully deleted or does not exist, `false` otherwise.
// - error: An error if the user lacks privileges or the command fails for another reason.
//
// Notes:
//
// - It first verifies if the user has sudo privileges before attempting the operation.
//

func DeleteAsSudo(filePath string) (bool, error) {
	// Step 1: Check if the file exists. If it doesn't, there's nothing to do.
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return true, nil
	}

	// Step 2: Check if the user has administrative privileges.
	canSudo, err := user.CanBeSudoAndIsNotRoot()
	if err != nil {
		// handle generic error explicitly: unexpected failure
		return false, errorx.Wrap(err, "failed to check for sudo privileges")
	}
	if !canSudo {
		// handle specific error explicitly: expected outcome
		return false, errorx.New("current user does not have sudo privileges")
	}

	// Step 3: Attempt to delete the file using `sudo rm`.
	command := fmt.Sprintf("sudo rm -f %s", filePath)
	_, err = executor.RunCliLocal(command)
	if err != nil {
		// handle generic error explicitly: unexpected failure (e.g., file path is invalid, permissions issues, etc.).
		return false, errorx.Wrap(err, "failed to delete file at %s as sudo", filePath)
	}

	// success
	return true, nil
}
