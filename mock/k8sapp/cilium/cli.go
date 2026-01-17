package cilium

import (
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

func (ciliumService) GetStatus(hostName, helmHost string, logger logx.Logger) (string, error) {
	out, err := play(hostName, helmHost, "", CilumSvc.cliToGetFeatureStatus(), logger)
	if err != nil {
		return "", err
	}
	return play(hostName, helmHost, out+" got status", CilumSvc.cliToGetStatus(), logger)
}

func (ciliumService) cliToGetStatus() string {
	var cmds = []string{
		`cilium status`,
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
func (ciliumService) cliToGetFeatureStatus() string {
	var cmds = []string{
		`tmp=$(mktemp /tmp/cilium-feature-XXX.md)`,
		`cilium features status -o markdown > $tmp`,
		`echo "generate feature status in $tmp"`,
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
