package other

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/mock/filex"
)

func getYaml() (*MapYaml, error) {
	// 1 - get local auto cached (embedded) file into a struct
	yamlAsStruct, err := filex.LoadYamlIntoStruct[MapYaml](yamlList)
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}
	return yamlAsStruct, nil
}

func (i *MapYaml) ConvertToString() string {
	var sb strings.Builder
	sb.WriteString("NAME\tDESCRIPTION\tURL\n") // header

	for _, r := range i.List {
		sb.WriteString(fmt.Sprintf(
			"%s\t%s\t%s\n",
			r.Name,
			r.Desc,
			r.Url,
		))
	}

	return sb.String()
}
