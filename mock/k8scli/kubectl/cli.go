package kubectl

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

func (i *Resource) Delete(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	logger.Debugf("%s:%s > CLI %s", hostName, kubectlHost, i.cliToDelete())
	return play(hostName, kubectlHost, "deleted "+i.Type.String()+" "+i.Name, i.cliToDelete(), logger)
}

func (i *Resource) Describe(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "described "+i.Type.String(), i.cliToDescribe(), logger)
}
func (i *Resource) Create(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	logger.Debugf("%s:%s > CLI %s", hostName, kubectlHost, i.cliToCreate())
	return play(hostName, kubectlHost, "created "+i.Type.String()+":"+i.Name, i.cliToCreate(), logger)
}
func (i *Resource) Generate(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "generated "+i.Type.String()+":"+i.Name, i.cliToGenerate(), logger)
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
func (i *Resource) GetLog(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "got ip for "+i.Type.String(), i.cliToGetLog(), logger)
}
func (i *Resource) ListCachedImg(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "got ip for "+i.Type.String(), i.cliToListCachedImg(), logger)
}
func (i *Resource) ListResource(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "got ip for "+i.Type.String(), i.cliToListResource(), logger)
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

func (i *Resource) cliToListCachedImg() string {
	switch i.Type {
	case ResNode:
		return fmt.Sprintf(`kubectl get node %s -o jsonpath='{range .status.images[*]}{.names[0]}{"\t"}{.sizeBytes}{"\n"}{end}'  | sed 's/@sha256:/ /' | awk 'BEGIN{print "IMAGE\tSHA\tSIZE"} {print $1,substr($2,1,10),$3}' | column -t`, i.Name)
	default:
		panic("unsupported resource type: " + i.Type)
	}
}
func (i *Resource) cliToListResource() string {
	switch i.Type {
	case ResNS:
		return fmt.Sprintf(`kubectl get all -n %s`, i.Name)
	default:
		panic("unsupported resource type: " + i.Type)
	}
}

// Notes:
//
// - get column name from the yaml view of the resource (-o yaml)
func cliToList(resType ResType) string {
	switch resType {
	case ResCM, ResEp, ResCRD, ResSA:
		return fmt.Sprintf("kubectl get %s -Ao wide", resType.String())
	case ResDeploy:
		return `kubectl get deploy -Ao wide | awk '{$8=substr($8,1,35) "..."; print $1,$2,$3,$6,$7,$8}' | column -t`
	case ResDs:
		return `
		echo -e "Ns\tName\tApp\tCurrent\tAge" && \
    kubectl get ds -Ao yaml | yq -r '.items[] | [
	  .metadata.namespace, 
  	.metadata.name, 
  	(.spec.template.spec.serviceAccountName // "default"), 
		.status.currentNumberScheduled, 
    .metadata.creationTimestamp
		] | @tsv'
		`
		// return `echo -e "Ns\tName\tApp" && kubectl get ds -Ao yaml | yq -r ".items[] | [.metadata.namespace, .metadata.name, .spec.selector.matchLabels.app, .status.desiredNumberScheduled, .status.numberReady] | @tsv" `
		// return `kubectl get ds -Ao wide | awk '{print $1,$2,$3,$4,$5,$6,$7,$8,$9}' | column -t`
	case ResNode:
		return `kubectl get nodes -Ao wide | awk '{print $1,$8,$(NF-1),$6,$2,$4,$3}' | column -t`
	case ResNS:
		return `kubectl get namespaces`
	case ResPod:
		return `kubectl get pods -Ao wide | awk '{print $1,$2,$4,$6,$8,$7}'| column -t`
	case ResRes:
		return `kubectl api-resources| sort `
	case ResPv:

		return `kubectl get pv -A --no-headers -o custom-columns="
		:.spec.storageClassName,
		:.metadata.name,
		:.spec.capacity.storage,
		:.spec.accessModes[0],
		:.spec.persistentVolumeReclaimPolicy,
		:.status.phase,
		:.spec.claimRef.namespace,
		:.spec.claimRef.name,
		:.metadata.creationTimestamp" | awk 'BEGIN {print "SC\tNAME\tCAPACITY\tACCESS\tRECLAIM\tSTATUS\tCLAIM\tAGE"} {print $1,$2,$3,$4,$5,$6"/"$7,$8,$9}' | column -t
		`
	case ResPvc:
		return `kubectl get pvc -A --no-headers -o custom-columns="
		:.metadata.namespace,
		:.metadata.name,
		:.spec.storageClassName,
		:.status.phase,
		:.spec.volumeName,
		:.status.capacity.storage,
		:.spec.accessModes[0],
		:.metadata.creationTimestamp" | awk 'BEGIN {print "NAMESPACE\tNAME\tSC\tSTATUS\tPV\tCAPACITY\tACCESS\tAGE"} {print $1,$2,$3,$4,$5,$6,$7,$8}' | column -t
		`
	case ResSC:
		return `kubectl get sc --no-headers -o custom-columns="
		:.metadata.name,
		:.provisioner,
		:.reclaimPolicy,
		:.volumeBindingMode,
		:.allowVolumeExpansion,
		:.metadata.annotations.storageclass\.kubernetes\.io/is-default-class,
		:.metadata.creationTimestamp" | awk 'BEGIN {print "NAME\tPROVISIONER\tRECLAIM\tBINDING\tEXPAND\tDEFAULT\tAGE"} {print $1,$2,$3,$4,$5,$6,$7}' | column -t
		`
	case ResSecret:
		// return `kubectl get secrets -Ao wide`
		return `kubectl get secrets -Ao wide | awk '{print $1,$2,$3,$4}' | column -t`
	case ResSvc:
		return `kubectl get svc -Ao wide | awk '{print $1,$2,$3,$4,$5,$6,$7}' | column -t`
	default:
		panic("unsupported resource type: " + resType)
	}
}

func (i *Resource) cliToDelete() string {
	// 1. Handle special patterns first
	if i.Type == ResManifest {
		return fmt.Sprintf("kubectl delete -f %s --ignore-not-found", i.Url)
	}

	// 2. Handle standard patterns
	if i.isClusterScoped() {
		return fmt.Sprintf("kubectl delete %s %s", i.Type.String(), i.Name)
	}

	return fmt.Sprintf("kubectl delete %s %s -n %s", i.Type.String(), i.Name, i.Ns)
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
	case ResApiServer:
		return `kubectl config view --minify | yq -r '.clusters[0].cluster.server' | tr -d '/' | cut -d: -f2`
	default:
		panic("unsupported resource type: " + i.Type)
	}
}
func (i *Resource) cliToGetLog() string {
	switch i.Type {
	case ResPod:
		return fmt.Sprintf(`kubectl logs %s -n %s`, i.Name, i.Ns)
		// return fmt.Sprintf(`kubectl logs %s -n %s --previous`, i.Name, i.Ns)
	default:
		panic("unsupported resource type: " + i.Type)
	}
}
func (i *Resource) cliToDescribe() string {
	theType := i.Type.String()
	name := i.Name
	ns := i.Ns

	// 1. Handle special patterns first
	switch i.Type {
	case ResManifest:
		return fmt.Sprintf("kubectl get -f %s --ignore-not-found", i.Url)
	case ResRes:
		return fmt.Sprintf("kubectl explain %s", name)
	}

	// 2. Handle standard patterns
	if i.isClusterScoped() {
		return fmt.Sprintf("kubectl describe %s %s", theType, name)
	}

	return fmt.Sprintf("kubectl describe %s %s -n %s", theType, name, ns)
}
func (i *Resource) cliToGetYaml() string {
	// Handle standard patterns
	if i.isClusterScoped() {
		return fmt.Sprintf("kubectl get %s %s -o yaml", i.Type.String(), i.Name)
	}
	return fmt.Sprintf("kubectl get %s %s -n %s -o yaml", i.Type.String(), i.Name, i.Ns)
}

func (i *Resource) cliToCreate() string {
	switch i.Type {
	case ResNS:
		return fmt.Sprintf(`kubectl create namespace %s`, i.Name)
	case ResSecret:
		return fmt.Sprintf(`kubectl run htpasswd-gen -n %[1]s --quiet --restart=Never --rm -i --image=httpd:2.4-alpine -- sh -c 'apk add -q --no-cache apache2-utils && PASS=$(head -c 20 /dev/urandom | base64) && htpasswd -Bbn %[2]s $PASS' | kubectl create secret generic %[3]s -n %[1]s --from-file=auth=/dev/stdin --dry-run=client -o yaml | kubectl apply -f -`, i.Ns, i.UserName, i.Name)
		// return fmt.Sprintf(`kubectl run htpasswd-gen -n %[1]s --quiet --restart=Never --rm -i --image=httpd:2.4-alpine -- sh -c 'apk add -q --no-cache apache2-utils && PASS=$(head -c 20 /dev/urandom | base64) && htpasswd -Bbn %[2]s $PASS' | kubectl create secret generic %[3]s -n %[1]s --from-file=auth=/dev/stdin`,i.Ns, i.UserName, i.Name)
	default:
		panic("unsupported resource type: " + i.Type)
	}
}
func (i *Resource) cliToGenerate() string {
	switch i.Type {
	case ResSecret:
		return fmt.Sprintf(`kubectl run htpasswd-gen   -n %s   --quiet --restart=Never   --rm -i    --image=httpd:2.4-alpine   -- sh -c '
		apk add -q --no-cache apache2-utils && PASS=$(head -c 20 /dev/urandom | base64) &&  htpasswd -Bbn %s $PASS'`, i.Ns, i.UserName)
	default:
		panic("unsupported resource type: " + i.Type)
	}
}
