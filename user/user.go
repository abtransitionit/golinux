package user

import (
	"os"
	"os/user"
	"strings"

	"github.com/abtransitionit/gocore/errorx"
	gc_props "github.com/abtransitionit/gocore/properties"
)

// Name: CanBeSudo
//
// Description:
//
// Checks if the current user has administrative privileges by verifying group membership
// based on the operating system type.
//
// Returns:
//
//   - bool:  Returns `true` if the user can use `sudo` (Linux) or `admin` (Darwin), false otherwise.
//   - error: Returns an error if the user or group information cannot be retrieved.
func CanBeSudo() (bool, error) {
	// Step 1: Quick check if already running as root ⚡️
	if os.Geteuid() == 0 {
		return true, nil
	}

	// Step 2: Get the current user
	usr, err := user.Current()
	if err != nil {
		return false, errorx.Wrap(err, "cannot determine current user")
	}

	// Step 3: Get the OS type to determine which group to check.
	osType, err := gc_props.GetPropertyLocal("ostype")
	if err != nil {
		return false, errorx.Wrap(err, "failed to determine os type")
	}

	// Step 4: Check for the appropriate administrative groups based on the OS.
	var adminGroupsToCheck []string
	switch osType {
	case "linux":
		adminGroupsToCheck = []string{"sudo", "wheel"}
	case "darwin":
		adminGroupsToCheck = []string{"admin", "wheel"}
	default:
		// handle specific error explicitly: expected outcome : If the OS is not Linux or Darwin we do not managed
		return false, errorx.New("not yet supported os type: %s", osType)
	}

	// Get the user's group IDs to check against our list of administrative groups.
	groupIDs, err := usr.GroupIds()
	if err != nil {
		// handle generic error explicitly: unexpected failure
		return false, errorx.Wrap(err, "cannot get user groups")
	}

	// Iterate through each of the user's group IDs.
	for _, gid := range groupIDs {
		// Look up the group name for each ID.
		group, err := user.LookupGroupId(gid)
		if err != nil {
			// A failure to look up a group can be an expected occurrence on some systems. We can safely skip it.
			continue
		}
		// Compare the looked-up group name with the list of administrative groups for the current OS.
		for _, adminGroup := range adminGroupsToCheck {
			// If the group name matches, the user has admin privileges. We use EqualFold for a case-insensitive comparison.
			if strings.EqualFold(group.Name, adminGroup) {
				// Found a match, so the user has admin privileges.
				return true, nil
			}
		}
	}
	// If we've iterated through all the groups and found no match, the user does not have admin privileges.
	return false, nil
}
