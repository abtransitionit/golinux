package kubectl

import (
	"fmt"
	"strings"
)

func (i *MapYamlManifest) string() string {
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
