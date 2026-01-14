package kubectl

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

func (i Resource) Describe(hostName, helmHost string, logger logx.Logger) (string, error) {
	return runKubectl(hostName, helmHost, "described "+i.Type.String(), i.CliToDesc(), logger)
}
func (i Resource) ListEvent(hostName, helmHost string, logger logx.Logger) (string, error) {
	return runKubectl(hostName, helmHost, "described "+i.Type.String(), i.CliToListEvent(), logger)
}

func List(resType ResType, hostName, helmHost string, logger logx.Logger) (string, error) {
	return runKubectl(hostName, helmHost, "listed "+resType.String(), CliToList(resType), logger)
}

func CliToList(resType ResType) string {
	switch resType {
	case ResNS:
		return `kubectl get namespaces`
	case ResNode:
		return `kubectl get nodes -Ao wide | awk '{print $1,$8,$(NF-1),$6,$2,$4,$3}' | column -t`
	case ResPod:
		return `kubectl get pods -Ao wide | awk '{print $1,$2,$4,$6,$8,$7}'| column -t`
	case ResSA:
		return `kubectl get sa -Ao wide`
	case ResCM:
		return `kubectl get cm -Ao wide`
	default:
		panic("unsupported resource type: " + resType)
	}
}
func (i Resource) CliToListEvent() string {
	switch i.Type {
	case ResPod:
		return fmt.Sprintf(`kubectl get events -n %s --field-selector involvedObject.name=%s`, i.Ns, i.Name)
	default:
		panic("unsupported resource type: " + i.Type)
	}
}

func (i Resource) CliToDesc() string {
	switch i.Type {
	case ResNode:
		return fmt.Sprintf(`kubectl describe node %s`, i.Name)
	case ResPod:
		return fmt.Sprintf(`kubectl describe pod %s -n %s`, i.Name, i.Ns)
	case ResNS:
		return fmt.Sprintf(`kubectl describe ns %s`, i.Name)
	case ResCM:
		return fmt.Sprintf(`kubectl describe cm %s`, i.Name)
	case ResSA:
		return fmt.Sprintf(`kubectl describe sa %s -n %s`, i.Name, i.Ns)
	default:
		panic("unsupported resource type: " + i.Type)
	}
}
