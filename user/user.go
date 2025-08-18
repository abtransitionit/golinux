package user

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

// Name: CanBeSudoAndIsNotRoot
//
// Description: Checks if the current user can be sudo AND is not root
//
// Returns:
//
//   - bool: `true` only if the user can use `sudo` AND is not root, `false` otherwise.
//   - error: if the user or group information cannot be retrieved.
//
// Notes:
//
// - On Darwin (macOS), this function will give an incorrect result. Use CanBeSudoAndIsNotRootExtend instead.
func CanBeSudoAndIsNotRoot() (bool, error) {
	// Step 1: Quick check if user is root
	if os.Geteuid() == 0 {
		return false, nil
	}

	// Step 2: Get the current user
	usr, err := user.Current()
	if err != nil {
		// handle generic error explicitly: unexpected failure
		return false, fmt.Errorf("cannot determine current user: %w", err)
	}

	// Step 3: Get the current user's group IDs
	groupIDs, err := usr.GroupIds()
	if err != nil {
		return false, fmt.Errorf("cannot get user groups: %w", err)
	}

	// Step 4: Check if the user is a member of the 'sudo' or 'wheel' group
	for _, gid := range groupIDs {
		group, err := user.LookupGroupId(gid)
		if err != nil {
			// A failure to look up a group can be an expected occurrence on some systems. We can safely skip it.
			continue
		}
		// Compare the looked-up group name directly.
		switch strings.EqualFold(group.Name, "sudo") || strings.EqualFold(group.Name, "wheel") {
		case true:
			// The user is a member of the 'sudo' or 'wheel' group
			return true, nil
		}
	}

	// If we've iterated through all the groups and found no match does not belongs to the 'sudo' or 'wheel' group,
	return false, nil
}

// Name: CanBeSudoAndIsNotRootExtend
//
// Description: Checks if the current user can be sudo AND is not root
//
// Returns:
//
//   - bool: `true` only if the user can use `sudo` AND is not root, `false` otherwise.
//   - error: if the user or group information cannot be retrieved.
//
// Notes:
//
// - works on Linux and Darwin (macOS - admin group)
func CanBeSudoAndIsNotRootExtend() (bool, error) {
	// Step 1: Quick check if user is root
	if os.Geteuid() == 0 {
		return false, nil
	}

	// Step 2: Get the current user
	usr, err := user.Current()
	if err != nil {
		// handle generic error explicitly: unexpected failure
		return false, fmt.Errorf("cannot determine current user: %w", err)
	}

	// Step 3: Get the current user's group IDs
	groupIDs, err := usr.GroupIds()
	if err != nil {
		return false, fmt.Errorf("cannot get user groups: %w", err)
	}

	// Step 4: Check if the user is a member of the 'sudo' or 'wheel' group
	for _, gid := range groupIDs {
		group, err := user.LookupGroupId(gid)
		if err != nil {
			// A failure to look up a group can be an expected occurrence on some systems. We can safely skip it.
			continue
		}
		// Compare the looked-up group name directly.
		switch strings.EqualFold(group.Name, "sudo") || strings.EqualFold(group.Name, "wheel") || strings.EqualFold(group.Name, "admin") {
		case true:
			// The user is a member of the 'sudo' or 'wheel' group
			return true, nil
		}
	}

	// If we've iterated through all the groups and found no match does not belongs to the 'sudo' or 'wheel' group,
	return false, nil
}
