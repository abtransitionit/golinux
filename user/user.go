package user

import (
	"os"
	"os/user"
	"strings"

	"github.com/abtransitionit/gocore/errorx"
)

// Name: CanBeSudo
//
// Description:
// Checks if the current user is a member of the sudo or wheel group, which
// indicates they can run commands with `sudo` privileges. This function
// also checks if the program is already running as the root user.
//
// Returns:
//
//   - bool:  Returns `true` if the user is a member of the sudo group or is already running as root,
//     `false` otherwise.
//   - error: Returns an error if the user or group information cannot be retrieved from the system.
//
// Notes:
//
//   - This function is designed for Unix-like environments (Linux, macOS).
//   - It relies on system-level group information, which is a robust way to determine sudo capability.
func CanBeSudo() (bool, error) {
	// Quick check: running as root ⚡️
	// If the effective user ID is 0, the process is already running as root.
	if os.Geteuid() == 0 {
		return true, nil
	}

	// Get the current user
	usr, err := user.Current()
	if err != nil {
		return false, errorx.Wrap(err, "cannot determine current user")
	}

	// Check if the user is a member of the 'sudo' or 'wheel' group.
	// This is the standard way a user is granted sudo capability.
	groupIDs, err := usr.GroupIds()
	if err != nil {
		return false, errorx.Wrap(err, "cannot get user groups")
	}

	for _, gid := range groupIDs {
		group, err := user.LookupGroupId(gid)
		if err != nil {
			// Skip groups that can't be looked up, which is a normal occurrence on some systems.
			continue
		}
		if strings.EqualFold(group.Name, "sudo") || strings.EqualFold(group.Name, "wheel") {
			return true, nil
		}
	}

	return false, nil
}
