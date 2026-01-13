package kubectl

import "strings"

func (i *Namespace) cliToList() string {
	var cmds = []string{
		`kubectl get namespaces`,
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
