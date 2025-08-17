package filex

import (
	"os"
	"os/exec"

	"github.com/abtransitionit/gocore/errorx"
)

// Name: TouchAsSudo
//
// Description:
// Creates a file at the given path by executing the `touch` command with `sudo` privileges.
// This function first verifies if the current user has `sudo` permissions without requiring
// a password and then proceeds with the file creation. It is a Linux-specific operation.
//
// Parameters:
//
//	filePath: The absolute path where the file should be created.
//
// Returns:
//
//   - bool:  Returns `true` if the file was successfully created, `false` otherwise.
//   - error: Returns an error if the user lacks `sudo` privileges or if the command fails.
//
// Notes:
//
//   - This function is a wrapper for a system command and relies on the `sudo` utility being available and properly configured.
//   - It is designed for creating files in privileged locations, such as `/usr/local/bin` or `/etc`.
//   - The function will fail if `sudo` credentials are required, as it uses the `-n` flag to avoid prompting for a password.
func TouchAsSudo(filePath string) (bool, error) {
	// 1. Check for sudo privileges.
	// The `sudo -n -v` command is used to check for cached sudo credentials without prompting for a password.
	cmd := exec.Command("sudo", "-n", "-v")
	if err := cmd.Run(); err != nil {
		// If the command fails, it means the user either doesn't have sudo privileges or a password is required.
		return false, errorx.Wrap(err, "current user does not have passwordless sudo privileges")
	}

	// 2. Attempt to create the file with sudo.
	// We use the `touch` command, a standard Linux utility for file creation.
	sudoCmd := exec.Command("sudo", "touch", filePath)
	if err := sudoCmd.Run(); err != nil {
		return false, errorx.Wrap(err, "failed to create file at %s as sudo", filePath)
	}

	// 3. Verify file existence.
	// We check for the file's existence to confirm the sudo command's success.
	if _, err := os.Stat(filePath); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, errorx.New("file not found after sudo touch at %s", filePath)
	}

	// Catch-all for other unexpected errors during verification.
	return false, errorx.Wrap(err, "failed to verify file existence at %s", filePath)
}
