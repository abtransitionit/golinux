package helm

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
)

// description: returns the chart readme
func (i *Chart) ViewReadme(hostName, helmHost string, logger logx.Logger) (string, error) {

	// 1 - get and play cli
	output, err := run.RunCli(helmHost, i.cliToViewReadme(), logger)
	if err != nil {
		return "", err
	}
	// handle success
	logger.Debugf("%s:%s:%s > readme md file in %s", hostName, helmHost, i.Qname, output)
	return output, nil
}

// description: returns the list of all K8s resources kind the chart will create in the K8s cluster
func (i *Chart) ListResKind(hostName, helmHost string, logger logx.Logger) (string, error) {

	// 1 - get and play cli
	output, err := run.RunCli(helmHost, i.cliToListResKind(), logger)
	if err != nil {
		return "", err
	}
	// handle success
	logger.Debugf("%s:%s:%s > listed chart resources kind to be created in the cluster", hostName, helmHost, i.Name)
	return output, nil
}

// description: returns the list of the resources name with their kind to be created
func (i *Chart) ListRes(hostName, helmHost string, logger logx.Logger) (string, error) {

	// 1 - get and play cli
	output, err := run.RunCli(helmHost, i.cliToListRes(), logger)
	if err != nil {
		return "", err
	}

	// return response
	return output, nil
}

// func (i *Chart) Create() (string, error) {
// 	var cmds = []string{
// 		fmt.Sprintf(`helm create %s`, i.Qname),
// 	}
// 	cli := strings.Join(cmds, " && ")
// 	return cli, nil
// }

// description: returns a list of all the kind of templates a helm charts will create
func (i *Chart) cliToListResKind() string {

	var cmds = []string{
		`. ~/.profile`,
		fmt.Sprintf(`echo -e "Res Kind\tNb" && helm template %s | yq -r '.kind' | sort | uniq -c | awk "{print \$2 \"\t\" \$1}"`, i.Qname),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
func (i *Chart) cliToListRes() string {

	var cmds = []string{
		`. ~/.profile`,
		fmt.Sprintf(`echo -e "Res Kind\tName\tNamespace" && helm template %s | yq -r ". | select(.kind) | [.kind, .metadata.name, .metadata.namespace] | @tsv" | sort | awk "{print \$1 \"\t\" \$2 \"\t\" \$3}"`, i.Qname),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
func (i *Chart) cliToViewReadme() string {

	var cmds = []string{
		`. ~/.profile`,
		fmt.Sprintf(`tmp=$(mktemp /tmp/%s-XXXXXX)`, i.Name),
		fmt.Sprintf(`helm show readme %s > $tmp`, i.Qname),
		`echo $tmp`,
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
