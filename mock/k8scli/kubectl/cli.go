package kubectl

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

func (i *Resource) Describe(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "described "+i.Type.String(), i.cliToDesc(), logger)
}
func (i *Resource) Create(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	logger.Debugf("%s:%s > CLI %s", hostName, kubectlHost, i.cliToCreate())
	return play(hostName, kubectlHost, "created "+i.Type.String(), i.cliToCreate(), logger)
}
func (i *Resource) GetYaml(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "got yaml for "+i.Type.String(), i.cliToGetYaml(), logger)
}
func (i *Resource) ListEvent(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "listed event for "+i.Type.String(), i.cliToListEvent(), logger)
}
func (i *Resource) GetIp(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "got ip for "+i.Type.String(), i.cliToGetIp(), logger)
}

func List(resType ResType, hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "listed "+resType.String(), cliToList(resType), logger)
}
func ListNoNs(resType ResType, hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "listed "+resType.String(), cliToListNoNs(resType), logger)
}
func ListNs(resType ResType, hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "listed "+resType.String(), cliToListNs(resType), logger)
}

func cliToListNoNs(resType ResType) string {
	switch resType {
	case ResRes:
		return `kubectl api-resources --namespaced=false  | sort `
	default:
		panic("unsupported resource type: " + resType)
	}
}
func cliToListNs(resType ResType) string {
	switch resType {
	case ResRes:
		return `kubectl api-resources --namespaced=true  | sort`
	default:
		panic("unsupported resource type: " + resType)
	}
}

func cliToList(resType ResType) string {
	switch resType {
	case ResRes:
		return `kubectl api-resources | sort`
	case ResSC:
		return `kubectl get sc`
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
	case ResCRD:
		return `kubectl get crd -Ao wide`
	default:
		panic("unsupported resource type: " + resType)
	}
}
func (i *Resource) cliToListEvent() string {
	switch i.Type {
	case ResPod:
		return fmt.Sprintf(`kubectl get events -n %s --field-selector involvedObject.name=%s`, i.Ns, i.Name)
	default:
		panic("unsupported resource type: " + i.Type)
	}
}
func (i *Resource) cliToGetIp() string {
	switch i.Type {
	case RestApiServer:
		return `kubectl config view --minify | yq -r '.clusters[0].cluster.server' | tr -d '/' | cut -d: -f2`
	default:
		panic("unsupported resource type: " + i.Type)
	}
}

func (i *Resource) cliToDesc() string {
	switch i.Type {
	case ResRes:
		return fmt.Sprintf(`kubectl explain %s`, i.Name)
	case ResSC:
		return fmt.Sprintf(`kubectl describe sc %s`, i.Name)
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
func (i *Resource) cliToCreate() string {
	switch i.Type {
	case ResNS:
		return fmt.Sprintf(`kubectl create namespace %s`, i.Name)
	default:
		panic("unsupported resource type: " + i.Type)
	}
}
func (i *Resource) cliToGetYaml() string {
	switch i.Type {
	case ResSC:
		return fmt.Sprintf("kubectl get sc %s -o yaml", i.Name)
		// return fmt.Sprintf("kubectl get node %s -o yaml | yq '.status.nodeInfo.kubeletVersion'", i.Name)
	case ResNode:
		return fmt.Sprintf("kubectl get node %s -o yaml", i.Name)
		// return fmt.Sprintf("kubectl get node %s -o yaml | yq '.status.nodeInfo.kubeletVersion'", i.Name)
	case ResPod:
		return fmt.Sprintf("kubectl get pod %s -n %s -o yaml", i.Name, i.Ns)
	case ResNS:
		return fmt.Sprintf("kubectl get ns %s -o yaml", i.Name)
	case ResCM:
		return fmt.Sprintf("kubectl get cm %s -n %s -o yaml", i.Name, i.Ns)
	case ResSA:
		return fmt.Sprintf("kubectl get sa %s -n %s -o yaml", i.Name, i.Ns)
	default:
		panic("unsupported resource type: " + string(i.Type))
	}
}
