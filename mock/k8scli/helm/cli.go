package helm

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/mock/filex"
	"github.com/abtransitionit/golinux/mock/k8sapp/cilium"
	"github.com/abtransitionit/golinux/mock/k8sapp/openebs"
)

func (i *Resource) Logout(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "logged out from "+i.Type.String(), i.cliToLogout(), logger)
}
func (i *Resource) Login(hostName, helmHost string, logger logx.Logger) error {
	return i.StepToLogin(hostName, helmHost, logger)
}

func (i *Resource) Install(hostName, helmHost string, logger logx.Logger) error {
	return i.StepToInstall(hostName, helmHost, logger)
}

//	func (i *Resource) Build(hostName, helmHost string, logger logx.Logger) (string, error) {
//		return play(hostName, helmHost, "builded "+i.Type.String(), i.cliToBuild(), logger)
//	}
func (i *Resource) Push(hostName, helmHost string, logger logx.Logger) error {
	return i.StepToPush(hostName, helmHost, logger)
}
func (i *Resource) Build(hostName, helmHost string, logger logx.Logger) error {
	return i.StepToBuild(hostName, helmHost, logger)
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
	return i.StepToAdd(hostName, helmHost, logger)
}

func (i *Resource) ListAuth(hostName, helmHost string, logger logx.Logger) (string, error) {
	return i.StepToListAuth()
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
	case ResChart:
		switch i.SType {
		case STypeChartBuild:
			// 1 - resolve the user yaml file
			userHome, err := os.UserHomeDir()
			if err != nil {
				panic(fmt.Sprintf("pbs getting user home to get file %q", artifactCfgRelPath))
			}
			artifactCfgFullPath := filepath.Join(userHome, artifactCfgRelPath)
			// 2 - build the cli
			var cmds = []string{
				fmt.Sprintf(`file=%s`, artifactCfgFullPath),
				`echo name`,
				fmt.Sprintf(`cat $file | yq e -r '.artifact."%s" | keys | .[]' -`, string(STypeChartBuild)),
			}
			cli := strings.Join(cmds, " && ")
			return cli
		case "", STypeChartStd:
			if i.Repo == "" {
				return `helm search repo`
			}
			return fmt.Sprintf(`helm search repo %s`, i.Repo)
		default:
			panic(fmt.Sprintf("unsupported subtype %q for resource type %q", i.SType, i.Type))
		}
	case ResRegistry:
		return `helm repo list`
	case ResRepo:
		return `helm repo list`
	case ResRelease:
		return `helm list -A`
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}

// description: delete a release and all its resources from the cluster
func (i *Resource) CliToDelete() string {
	switch i.Type {
	case ResRelease:
		var cmds = []string{
			fmt.Sprintf(`helm uninstall %s     -n %s`, i.Name, i.Namespace),
			// fmt.Sprintf(`helm uninstall %s     -n %s  --keep-history=false`, i.Name, i.Namespace),
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
func (i *Resource) StepToListAuth() (string, error) {
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
func (i *Resource) StepToAdd(hostName, helmHost string, logger logx.Logger) (string, error) {
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

// description: list authorized OCI chart
func (i *Resource) ListOci(hostName string, logger logx.Logger) error {
	// 1 - check
	if i.Type != ResChart {
		return fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}
	// handle success
	return nil
}

// description: return a registry (with all it properties) from a local user credential file
func (i *Resource) GetRegistry(hostName string, logger logx.Logger) (*Registry, error) {
	// 1 - check
	if i.Type != ResRegistry {
		return nil, fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}
	// 2 - resolve the registry credential file path
	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to resolve home directory > %w", err)
	}
	registryCredentialFullPath := filepath.Join(userHome, registryCredentialRelPath)

	// 3 - get the yaml
	registrySliceCfg, err := filex.LoadExternalYamlIntoStruct[RegistryCfg](registryCredentialFullPath)
	if err != nil {
		return nil, fmt.Errorf("%s > loading registry yaml credential file: %w", hostName, err)
	}

	// 4 - get the registry
	registry, ok := registrySliceCfg.Registry[i.Name]
	if !ok {
		return nil, fmt.Errorf("registry %q not found in registry credential file", i.Name)
	}
	// handle success
	return &registry, nil
}

// description: return an artifact (with all it properties) from a local well known rel file path
func (i *Resource) GetArtifact(hostName string, logger logx.Logger) (*Artifact, error) {
	// 1 - check
	if i.Type != ResChart {
		return nil, fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}

	// 2 - get the absolute file path
	artifactCfgFullPath, err := filex.GetUserFilePath(artifactCfgRelPath)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	// 3 - get the yaml
	artifactSliceCfg, err := filex.LoadExternalYamlIntoStruct[ArtifactCfg](artifactCfgFullPath)
	if err != nil {
		return nil, fmt.Errorf("%s > loading registry yaml credential file: %w", hostName, err)
	}
	// 4 - get a set
	artifactSet, ok := artifactSliceCfg.Artifact[string(STypeChartBuild)]
	if !ok {
		return nil, fmt.Errorf("artifact group %q not found", i.SType)
	}

	// 5 - get the artifact
	artifact, ok := artifactSet.Items[i.Name]
	if !ok {
		return nil, fmt.Errorf("artifact %q not found in group %q", i.Name, i.Type)
	}

	// handle success
	return &artifact, nil
}

func (i *Resource) StepToLogin(hostName, helmHost string, logger logx.Logger) error {
	// 1 - check
	if i.Type != ResRegistry {
		return fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}
	registry, err := i.GetRegistry(hostName, logger)
	if err != nil {
		return fmt.Errorf("%s:%s > getting registry credential yaml > %w", hostName, helmHost, err)
	}
	// operate on instance whith field
	_, err = play(hostName, helmHost, "listed "+i.Type.String(), i.cliToLogin(registry), logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > login to registry %s with provided token > %w", hostName, helmHost, i.Name, i.Name, err)
	}
	// handle success
	logger.Debugf("%s:%s:%s > logged in Helm to registry", hostName, helmHost, i.Name)
	return nil

}
func (i *Resource) StepToInstall(hostName, helmHost string, logger logx.Logger) error {
	// 1 - check
	if i.Type != ResRelease {
		return fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}
	// 12 - check chart exist
	// 121 - get instance and operate
	chart := Resource{Type: ResChart, QName: i.QName, Version: i.Version}
	out, err := play(hostName, helmHost, "listed "+i.Type.String(), chart.cliToCheckExistence(), logger)
	outBool := map[string]bool{"true": true, "false": false}[strings.TrimSpace(out)]
	if err != nil {
		return fmt.Errorf("%s:%s:%s > checking chart existence > %w", hostName, helmHost, chart.QName, err)
	} else if outBool != true {
		return fmt.Errorf("%s:%s:%s > chart %s does not exist on the helm client", hostName, helmHost, i.Name, chart.QName)
	}
	// 122 - check the chart version exists
	chartVersion := Resource{Type: ResChartVersion, QName: i.QName, Version: i.Version}
	out, err = play(hostName, helmHost, "listed "+i.Type.String(), chartVersion.cliToCheckExistence(), logger)
	outBool = map[string]bool{"true": true, "false": false}[strings.TrimSpace(out)]
	if err != nil {
		return fmt.Errorf("%s:%s:%s > checking chart version existence > %w", hostName, helmHost, chart.QName, err)
	} else if outBool != true {
		return fmt.Errorf("%s:%s:%s > chart version %s does not exist on the helm client", hostName, helmHost, i.Name, chart.QName)
	}
	// 2 - Get value file
	valueFileAsbyte, err := i.GetValueFile(logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > getting value file > %w", hostName, helmHost, i.Name, err)
	}

	// log
	logger.Debug("--- BEGIN:Rendered Value file  ---")
	scanner := bufio.NewScanner(bytes.NewReader(valueFileAsbyte))
	for scanner.Scan() {
		logger.Debug(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.Warnf("error while logging rendered kubeadm config: %v", err)
	}
	logger.Debug("--- END ---")

	// 3 - install
	// 31 - get instance and operate
	release := Resource{Type: ResRelease, QName: i.QName, Version: i.Version, Name: i.Name, Namespace: i.Namespace}
	out, err = play(hostName, helmHost, "listed "+i.Type.String(), release.cliToInstall(valueFileAsbyte), logger)
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
	case strings.Contains(i.Name, "openebs"):
		return openebs.GetValueFile(i.Param, logger)

	default:
		return nil, fmt.Errorf("no value provider for release %s", i.Name)
	}
}

// Description: install a release for the first time or upgrade an existing one
//
// Note:
//
// - the resource (ie. a chart must define a value file)
func (i *Resource) cliToInstall(valueFile []byte) string {
	// 1 - check
	if i.Type != ResRelease {
		panic("resource type not supported for this action: %s" + i.Type)
	}
	// 2 - build
	encodedValueFile := base64.StdEncoding.EncodeToString(valueFile)
	// TODO - if chart is local path change this - for now works only for chart in a repo
	i.Repo = strings.Split(i.QName, "/")[0]
	var cmds = []string{
		fmt.Sprintf(
			`printf '%s' | base64 -d | helm upgrade --install %s --labels "repoName=%s" %s --atomic --wait --create-namespace --timeout 10m --namespace %s %s -f -`,
			encodedValueFile,
			i.Name,
			i.Repo,
			i.QName,
			i.Namespace,
			i.versionFlag()),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
func (i *Resource) cliToBuild(artifact *Artifact) string {
	var cli string
	switch i.Type {
	case ResChart:
		var cmds = []string{
			fmt.Sprintf(`helm package %s --destination %s`,
				artifact.FolderSrc,
				artifact.FolderDst,
			),
		}
		cli = strings.Join(cmds, " && ")

	default:
		panic("unsupported resource type for this action: " + i.Type)
	}

	return cli
}
func (i *Resource) cliToBuild2() string {
	// 1 - check
	if i.Type != ResArtifact {
		panic("resource type not supported for this action: %s" + i.Type)
	}

	var cmds = []string{
		fmt.Sprintf(`helm package %s --destination %s`,
			filepath.Join(i.Param["folderSrcRoot"], i.Name),
			i.Param["folderDst"],
		),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}
func (i *Resource) StepToBuild(hostName, helmHost string, logger logx.Logger) error {
	// 1 - check
	if i.Type != ResChart {
		return fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}
	artifact, err := i.GetArtifact(hostName, logger)
	if err != nil {
		return fmt.Errorf("%s:%s > getting artifact yaml > %w", hostName, helmHost, err)
	}
	// cli to play
	_, err = play(hostName, helmHost, "builded "+i.Type.String(), i.cliToBuild(artifact), logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > login to registry %s with provided token > %w", hostName, helmHost, i.Name, i.Name, err)
	}
	// handle success
	// logger.Debugf("%s:%s:%s > builded Helm chart's artifact %s", hostName, helmHost, i.Name, artifact.FolderDst)
	return nil

}

// todo: an artifact is pushed in a registry
func (i *Resource) StepToPush(hostName, helmHost string, logger logx.Logger) error {
	// 1 - check
	if i.Type != ResArtifact {
		return fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}
	artifact, err := i.GetArtifact(hostName, logger)
	if err != nil {
		return fmt.Errorf("%s:%s > getting artifact yaml > %w", hostName, helmHost, err)
	}
	// cli to play
	_, err = play(hostName, helmHost, "builded "+i.Type.String(), i.cliToBuild(artifact), logger)
	if err != nil {
		return fmt.Errorf("%s:%s:%s > login to registry %s with provided token > %w", hostName, helmHost, i.Name, i.Name, err)
	}
	// handle success
	// logger.Debugf("%s:%s:%s > builded Helm chart's artifact %s", hostName, helmHost, i.Name, artifact.FolderDst)
	return nil

}

func (i *Resource) cliToLogin(registry *Registry) string {
	// 1 - check
	if i.Type != ResRegistry {
		panic("resource type not supported for this action: %s" + i.Type)
	}
	// 2 - base64 encode the token to avoid it displays in logs
	encodedRegistryAccessToken := base64.StdEncoding.EncodeToString([]byte(registry.Param.AccessToken))
	var cmds = []string{
		fmt.Sprintf(`echo %s | base64 -d | helm registry login %s -u %s --password-stdin`,
			encodedRegistryAccessToken,
			registry.Param.DnsOrIp,
			registry.Param.User,
		),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

func (i *Resource) cliToLogout() string {
	// 1 - check
	if i.Type != ResRegistry {
		panic("resource type not supported for this action: %s" + i.Type)
	}

	var cmds = []string{

		fmt.Sprintf(`helm registry logout %s`,
			i.Param["DnsOrIp"],
		),
	}
	cli := strings.Join(cmds, " && ")
	return cli
}

// --labels "repoName=%s"

func (i *Resource) versionFlag() string {
	// 1 - check
	if i.Type != ResRelease {
		panic("resource type not supported for this action: %s" + i.Type)
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

// // description: build a chart artifact from a chart folder
// func (i *Resource) StepToBuild(hostName, helmHost string, logger logx.Logger) error {
// 	// 1 - check
// 	if i.Type != ResChart {
// 		return fmt.Errorf("resource type not supported for this action: %s", i.Type)
// 	}
// 	_, err := play(hostName, helmHost, "builded "+i.Type.String(), i.cliToBuild(), logger)
// 	if err != nil {
// 		return fmt.Errorf("%s:%s:%s > building helm artifact from chart %s > %w", hostName, helmHost, i.Name, i.Name, err)
// 	}

//		// handle success
//		logger.Debugf("%s:%s:%s > builded chart artifact into  into %s", hostName, helmHost, i.Name, i.Param["folderDst"])
//		return nil
//	}

//	func (i *Resource) GetFromYaml(hostName, helmHost string, logger logx.Logger) (*Resource, error) {
//		return i.getFromYaml(hostName)
//	}
//
//	func (i *Resource) GetFromYaml(hostName, helmHost string, logger logx.Logger) (*Resource, error) {
//		return i.getFromYaml(hostName)
//	}

// // description: get a set of artifacts (with their properties) from a local user file
// func (i *Resource) Get(hostName string, logger logx.Logger) (Resource, error) {
// 	// 1 - check
// 	if i.Type != ResArtifact {
// 		return Resource{}, fmt.Errorf("resource type not supported for this action: %s", i.Type)
// 	}
// 	// 2 - resolve the artifact file path
// 	userHome, err := os.UserHomeDir()
// 	if err != nil {
// 		return Resource{}, fmt.Errorf("failed to resolve home directory > %w", err)
// 	}
// 	artifactCfgFullPath := filepath.Join(userHome, artifactCfgRelPath)

// 	// 3 - get the yaml
// 	artifactSliceConfig, err := filex.LoadExternalYamlIntoStruct[Config](artifactCfgFullPath)
// 	if err != nil {
// 		return Resource{}, fmt.Errorf("%s > loading artifact yaml file: %w", hostName, err)
// 	}

//		// 4 - get the instance
//		artifact, ok := artifactSliceConfig.Registry[i.Name]
//		if !ok {
//			return Resource{}, fmt.Errorf("registry %q not found in registry credential file", i.Name)
//		}
//		// handle success
//		return artifact, nil
//	}
// func (i *Resource) StepToBuild2(hostName, helmHost string, logger logx.Logger) error {
// 	// 1 - check
// 	if i.Type != ResArtifact {
// 		return fmt.Errorf("resource type not supported for this action: %s", i.Type)
// 	}
// 	artifact, err := i.Get(hostName, logger)
// 	if err != nil {
// 		return fmt.Errorf("%s:%s > getting registry credential yaml > %w", hostName, helmHost, err)
// 	}
// 	// operate on instance whith field
// 	_, err = play(hostName, helmHost, "builded "+i.Type.String(), i.cliToBuild(artifact), logger)
// 	if err != nil {
// 		return fmt.Errorf("%s:%s:%s > login to registry %s with provided token > %w", hostName, helmHost, i.Name, i.Name, err)
// 	}
// 	// handle success
// 	logger.Debugf("%s:%s:%s > logged in Helm to registry", hostName, helmHost, i.Name)
// 	return nil

// }
