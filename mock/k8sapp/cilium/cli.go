package cilium

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

func (ciliumService) GetStatus(hostName, helmHost string, logger logx.Logger) (string, error) {
	out, err := play(hostName, helmHost, "", CilumSvc.cliToGetFeatureStatus(), logger)
	if err != nil {
		return "", err
	}
	msg := fmt.Sprintf(`ssh %s "cat %s" | tee %[2]s > /dev/null; code %[2]s`, strings.TrimSpace(helmHost), strings.TrimSpace(out))
	fmt.Println(msg)
	return play(hostName, helmHost, " got status", CilumSvc.cliToGetStatus(), logger)
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
		`echo "$tmp"`,
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
