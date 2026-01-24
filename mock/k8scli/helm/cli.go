package helm

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/k8sapp/cilium"
)

func (i *Resource) Install(hostName, helmHost string, logger logx.Logger) error {
	return i.ActionToInstall(hostName, helmHost, logger)
}

func (i *Resource) Detail(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "listed "+i.Type.String(), i.CliToDetail(), logger)
}
func (i *Resource) Describe(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "listed "+i.Type.String(), i.CliToDescribe(), logger)
}
func (i *Resource) GetReadme(hostName, helmHost string, logger logx.Logger) (string, error) {
	out, err := play(hostName, helmHost, "", i.CliToGetReadme(), logger)
	if err != nil {
		return "", err
	}
	msg := fmt.Sprintf(`ssh %s "cat %s" | tee %[2]s > /dev/null; code %[2]s`, strings.TrimSpace(helmHost), strings.TrimSpace(out))
	return msg, nil
}

//	func (i *Resource) GetFromYaml(hostName, helmHost string, logger logx.Logger) (*Resource, error) {
//		return i.getFromYaml(hostName)
//	}
//
//	func (i *Resource) GetFromYaml(hostName, helmHost string, logger logx.Logger) (*Resource, error) {
//		return i.getFromYaml(hostName)
//	}
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
	return i.ActionToAdd(hostName, helmHost, logger)
}

func (i *Resource) ListAuth(hostName, helmHost string, logger logx.Logger) (string, error) {
	return i.ActionToListAuth()
}
func (i *Resource) GetEnv(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "got env ", i.CliToGetEnv(), logger)
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
		return fmt.Sprintf(`echo -e "Res Kind\tNb" && helm template %s | yq -r '.kind' | sort | uniq -c | awk '{print $2 "\t" $1}'`, i.QName)
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
func (i *Resource) ActionToListAuth() (string, error) {
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
func (i *Resource) ActionToAdd(hostName, helmHost string, logger logx.Logger) (string, error) {
	// 1 - check
	if i.Type != ResRepo {
		return "", fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}
	// 2 - lookup this repo into the yaml
	repo, err := i.getFromYaml(hostName)
	if err != nil {
		return "", fmt.Errorf("%s:%s > getting repo: maybe it is not in the whitelist:%w", hostName, helmHost, err)
	}

	// 3 - set an instance property extracted from the yaml
	i.Url = repo.Url

	// 4 - get and play cli
	return play(hostName, helmHost, "listed "+i.Type.String(), i.CliToAdd(), logger)
}
func (i *Resource) ActionToInstall(hostName, helmHost string, logger logx.Logger) error {
	// 1 - check
	if i.Type != ResRelease {
		return fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}
	// 2 - check chart exist
	// 21 - get instance and operate
	chart := Resource{Type: ResChart, QName: i.QName, Version: i.Version}
	out, err := play(hostName, helmHost, "listed "+i.Type.String(), chart.cliToCheckExistence(), logger)
	outBool := map[string]bool{"true": true, "false": false}[strings.TrimSpace(out)]
	if err != nil {
		return fmt.Errorf("%s:%s:%s > checking chart existence > %w", hostName, helmHost, chart.QName, err)
	} else if outBool != true {
		return fmt.Errorf("%s:%s:%s > chart %s does not exist on the helm client", hostName, helmHost, i.Name, chart.QName)
	}
	// 22 - check the chart version exists
	chartVersion := Resource{Type: ResChartVersion, QName: i.QName, Version: i.Version}
	out, err = play(hostName, helmHost, "listed "+i.Type.String(), chartVersion.cliToCheckExistence(), logger)
	outBool = map[string]bool{"true": true, "false": false}[strings.TrimSpace(out)]
	if err != nil {
		return fmt.Errorf("%s:%s:%s > checking chart version existence > %w", hostName, helmHost, chart.QName, err)
	} else if outBool != true {
		return fmt.Errorf("%s:%s:%s > chart version %s does not exist on the helm client", hostName, helmHost, i.Name, chart.QName)
	}
	// 3 - Get value file
	cfgAsbyte, err := i.GetValueFile(logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > getting value file > %w", hostName, helmHost, i.Name, err)
	}
	// cfgAsbyte, err := cilium.GetValueFile(i.Param, logger)
	// if err != nil {
	// 	return fmt.Errorf("%s:%s:%s > getting value file > %w", hostName, helmHost, i.Name, err)
	// }

	// log
	logger.Debug("--- BEGIN:Rendered Value file  ---")
	scanner := bufio.NewScanner(bytes.NewReader(cfgAsbyte))
	for scanner.Scan() {
		logger.Debug(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Warnf("error while logging rendered kubeadm config: %v", err)
	}
	logger.Debug("--- END ---")

	// 4 - install
	// 41 - get instance and operate
	release := Resource{Type: ResRelease, QName: i.QName, Version: i.Version, Name: i.Name, Namespace: i.Namespace}
	out, err = play(hostName, helmHost, "listed "+i.Type.String(), release.cliToInstall(cfgAsbyte), logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > installing helm release from chart %s > %w", hostName, helmHost, i.Name, i.QName, err)
	}
	// handle success
	logger.Debugf("%s:%s:%s > installed helm release from chart %s:%s", hostName, helmHost, i.Name, i.QName, i.Version)
	return nil
}

func (i Resource) GetValueFile(logger logx.Logger) ([]byte, error) {
	// 1 - check
	if i.Type != ResRelease {
		return nil, fmt.Errorf("not a release resource")
	}
	// 2 - get
	switch {
	case strings.Contains(i.Name, "cilium"):
		return cilium.GetValueFile(i.Param, logger)

	default:
		return nil, fmt.Errorf("no value provider for release %s", i.Name)
	}
}

func (i *Resource) cliToInstall(cfg []byte) string {
	// 1 - check
	if i.Type != ResRelease {
		panic("resource type not supported for this cli: %s" + i.Type)
	}
	// 2 - build
	encoded := base64.StdEncoding.EncodeToString(cfg)
	var cmds = []string{
		fmt.Sprintf(
			`printf '%s' | base64 -d | helm install %s %s --atomic --wait --timeout 10m --namespace %s %s -f -`,
			encoded,
			i.Name,
			i.QName,
			i.Namespace,
			i.versionFlag()),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

func (i *Resource) versionFlag() string {
	// 1 - check
	if i.Type != ResRelease {
		panic("resource type not supported for this cli: %s" + i.Type)
	}
	// 2 - build
	if i.Version != "" {
		return fmt.Sprintf("--version %s", i.Version)
	}
	return ""
}

func (i *Resource) cliToCheckExistence() string {
	switch i.Type {
	case ResChart:
		return fmt.Sprintf(`helm show chart %s >/dev/null 2>&1 && echo "true" || echo "false"`, i.QName)
	case ResChartVersion:
		return fmt.Sprintf(`helm show chart --version %s %s >/dev/null 2>&1 && echo "true" || echo "false"`, i.Version, i.QName)

	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}

// TODO
func (i *Resource) CliToGetEnv() string {
	switch i.Type {
	case ResHelm:
		return `helm env`
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}
