package helm

import (
	"fmt"

	"github.com/abelgacem/lucg/luccore"
)

// cli := fmt.Sprintf(`helm list -A`)
// output,cerr,err := config.PlayQueryHelm(cli)
// if err != nil { fmt.Fprintln(os.Stderr, cerr)}
// fmt.Println(output)

func PlayQueryHelm(HelmQuery string) (result string, customErr string, err error) {
	// Play request
	output, _, cerrSrc, errSrc := luccore.RunCLIOnVM(HelmHost, HelmQuery)
	// handle FAILURE
	if errSrc != nil {
		customErr = fmt.Sprintf("❌ Helm Command failed > %v > Output > %s", cerrSrc, output)
		return "", customErr, errSrc
	}
	// handle SUCCESS
	return output, "", nil
}
