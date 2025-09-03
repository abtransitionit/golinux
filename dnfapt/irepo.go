package dnfapt

import (
	"context"
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

// Name: InstallPackage
//
// Description: install a dnfapt package on a Linux distro
func InstallDaRepository(ctx context.Context, logger logx.Logger, osFamily string, repoName string) error {

	if osFamily != "rhel" && osFamily != "fedora" && osFamily != "debian" {
		return fmt.Errorf("this function only supports Linux (rhel, fedora, debian), but found: %s", osFamily)
	}

	// logic for installtion
	switch osFamily {
	case "rhel", "fedora":
	case "debian":
	}

	return nil
}
