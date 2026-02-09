package kubectl

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
)

func (i *Resource) ListAuth(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return i.StepToListAuth()
}

func (i *Resource) ListResName(hostName, helmHost string, logger logx.Logger) (string, error) {
	return play(hostName, helmHost, "list resource type/name to be created for "+i.Name, i.CliToListResName(), logger)
}

func (i *Resource) ListResKind(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return play(hostName, kubectlHost, "list resource type to be created for "+i.Name, i.cliToListResKind(), logger)
}

func (i *Resource) Apply(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	return i.StepTopply(hostName, kubectlHost, logger)
	// return play(hostName, kubectlHost, "applied manifest "+i.Name, i.cliToApply(), logger)
}

func (i *Resource) StepToListAuth() (string, error) {
	// 1 - check
	if i.Type != ResManifest {
		return "", fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}
	// 2 - get the yaml file into a var/struct
	YamlStruct, err := getYamlListManifest()
	if err != nil {
		return "", fmt.Errorf("getting the yaml > %w", err)
	}

	// handle success
	return YamlStruct.string(), nil
}

func (i *Resource) cliToApply() string {
	switch i.Type {
	case ResManifest:
		return fmt.Sprintf(`kubectl apply -f %s`, i.Url)
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}

func (i *Resource) cliToListResKind() string {
	switch i.Type {
	case ResManifest:
		return fmt.Sprintf(`echo -e "Res Kind\tNb" && kubectl apply -f %s --dry-run=server -o yaml | yq -r '.items[].kind' | sort | uniq -c | awk '{print $2 "\t" $1}'`, i.Url)
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}

func (i *Resource) CliToListResName() string {
	switch i.Type {
	case ResManifest:
		return fmt.Sprintf(`echo -e "Res Kind\tName\tNamespace" && kubectl apply -f %s --dry-run=server -o yaml | yq -r ".items[]  | [.kind, .metadata.name, .metadata.namespace] | @tsv" | sort | awk '{print $1 "\t" $2 "\t" $3}'`, i.Url)
	default:
		panic("unsupported resource type for this action: " + i.Type)
	}
}

// kubectl apply -f %s --dry-run=client -o yaml | yq -r '.items[] | [.kind, .metadata.name, .metadata.namespace] | @tsv' | sort | awk '{print $1 "\t" $2 "\t" ($3 == "null" ? "-" : $3)}'

func (i *Resource) StepTopply(hostName, kubectlHost string, logger logx.Logger) (string, error) {
	// 1 - check
	if i.Type != ResManifest {
		return "", fmt.Errorf("resource type not supported for this action: %s", i.Type)
	}
	// 2 - lookup this manifest into the yaml
	manifest, err := i.getFromYaml(hostName)
	if err != nil {
		return "", fmt.Errorf("%s:%s > getting repo: maybe it is not in the whitelist:%w", hostName, kubectlHost, err)
	}

	// 3 - set an instance property extracted from the yaml
	i.Url = manifest.Url

	// 4 - get and play cli
	return play(hostName, kubectlHost, "applied manifest "+i.Name, i.cliToApply(), logger)
}
