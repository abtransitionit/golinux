package helm

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

func (i *Resource) Detail(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "listed "+i.Type.String(), i.CliToDetail(), logger)
}
func (i *Resource) Describe(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "listed "+i.Type.String(), i.CliToDescribe(), logger)
}
func (i *Resource) GetReadme(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "listed "+i.Type.String(), i.CliToGetReadme(), logger)
}
func (i *Resource) ListResName(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "listed "+i.Type.String(), i.CliToListResName(), logger)
}

func (i *Resource) ListResKind(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "listed "+i.Type.String(), i.CliToListResKind(), logger)
}

func (i *Resource) ListHistory(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "listed "+i.Type.String(), i.CliToListHistory(), logger)
}

func (i *Resource) List(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "listed "+i.Type.String(), i.CliToList(), logger)
}
func (i *Resource) Delete(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "deleted "+i.Type.String(), i.CliToDelete(), logger)
}
func (i *Resource) Add(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "listed "+i.Type.String(), i.CliToAdd(), logger)
}

func (i *Resource) ListPermit(hostName, helmHost string, logger logx.Logger) (string, error) {
	return i.ActionToListPermit()
}
func (i *Resource) Install(hostName, helmHost string, logger logx.Logger) (string, error) {
	return i.ActionToInstall()
}

func (i *Resource) CliToDetail() string {
	switch i.Type {
	case ResRelease:
		// return fmt.Sprintf(`helm get all %s --namespace %s --revision %s`, i.Name, i.Namespace, i.Revision)
		// return fmt.Sprintf(`helm get manifest %s --namespace %s --revision %s`, i.Name, i.Namespace, i.Revision)
		return fmt.Sprintf(`helm get values %s --namespace %s --revision %s`, i.Name, i.Namespace, i.Revision)
		// return fmt.Sprintf(`helm get hooks %s --namespace %s --revision %s`, i.Name, i.Namespace, i.Revision)
		// return fmt.Sprintf(`helm get notes %s --namespace %s --revision %s`, i.Name, i.Namespace, i.Revision)
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}
func (i *Resource) CliToDescribe() string {
	switch i.Type {
	case ResRelease:
		return fmt.Sprintf(`helm get manifest %s --namespace %s`, i.Name, i.Namespace)
		// return fmt.Sprintf(`helm get status %s --namespace %s`, i.Name, i.Namespace)
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}
func (i *Resource) CliToList() string {
	switch i.Type {
	case ResRepo:
		return `helm repo list`
	case ResChart:
		// list all charts of a all repo
		if i.Repo == "" {
			return `helm search repo`
		}
		// list charts of a specific repo
		return fmt.Sprintf(`helm search repo %s`, i.Repo)
	case ResRelease:
		return `helm list -A`
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}

func (i *Resource) CliToDelete() string {
	switch i.Type {
	case ResRelease:
		var cmds = []string{
			fmt.Sprintf(`helm uninstall %s     -n %s`, i.Name, i.Namespace),
			// fmt.Sprintf(`helm uninstall %s     -n %s  --keep-history=false`, i.Name, i.Namespace),
			fmt.Sprintf(`kubectl delete secret -n %s -l owner=helm,name=%s`, i.Namespace, i.Name),
			fmt.Sprintf(`kubectl delete all    -n %s -l owner=helm,name=%s`, i.Namespace, i.Name),
		}
		cli := strings.Join(cmds, " && ")
		return cli
	case ResRepo:
		return fmt.Sprintf(`helm repo remove %s`, i.Name)
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}
func (i *Resource) CliToListHistory() string {
	switch i.Type {
	case ResRelease:
		return fmt.Sprintf(`helm history %s -n %s`, i.Name, i.Namespace)
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}
func (i *Resource) CliToGetReadme() string {
	switch i.Type {
	case ResChart:
		var cmds = []string{
			fmt.Sprintf(`tmp=$(mktemp /tmp/%s-XXXXXX.md)`, i.Name),
			fmt.Sprintf(`helm show readme %s > $tmp`, i.QName),
			`echo $tmp`,
		}
		cli := strings.Join(cmds, " && ")
		return cli
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}
func (i *Resource) CliToListResKind() string {
	switch i.Type {
	case ResChart:
		return fmt.Sprintf(`echo -e "Res Kind\tNb" && helm template %s | yq -r '.kind' | sort | uniq -c | awk "{print \$2 \"\t\" \$1}"`, i.QName)
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}
func (i *Resource) CliToListResName() string {
	switch i.Type {
	case ResChart:
		return fmt.Sprintf(`echo -e "Res Kind\tName\tNamespace" && helm template %s | yq -r ". | select(.kind) | [.kind, .metadata.name, .metadata.namespace] | @tsv" | sort | awk "{print \$1 \"\t\" \$2 \"\t\" \$3}"`, i.QName)
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}

func (i *Resource) CliToAdd() string {
	switch i.Type {
	case ResRepo:
		var cmds = []string{
			fmt.Sprintf(`helm repo add %s %s`, i.Name, i.Url),
			`helm repo update`,
		}
		cli := strings.Join(cmds, " && ")
		return cli
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}
func (i *Resource) ActionToListPermit() (string, error) {
	// 1 - check
	if i.Type != ResRepo {
		return "", fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}
	// 2 - get the yaml file into a var/struct
	YamlStruct, err := GetYamlRepo()
	if err != nil {
		return "", fmt.Errorf("getting the yaml > %w", err)
	}
	// handle success
	return YamlStruct.ConvertToString(), nil
}
func (i *Resource) ActionToInstall() (string, error) {
	// 1 check
	// 11 - check resource type
	if i.Type != ResRelease {
		return "", fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}
	// 12 - TODO:check - all info are in the instance: i.Name, i.Namespace, i.QName
	// 2 - get the yaml file into a var/struct
	YamlStruct, err := GetYamlRepo()
	if err != nil {
		return "", fmt.Errorf("getting the yaml > %w", err)
	}
	// handle success
	return YamlStruct.ConvertToString(), nil
}
