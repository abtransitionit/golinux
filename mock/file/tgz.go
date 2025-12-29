package file

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
)

func CliDetectTgzFolder(filePath string, logger logx.Logger) string {
	// define cli
	clis := []string{
		fmt.Sprintf(`tar -tzf %q | awk -F/ 'NF>1 {print $1"/"}' | sort -u | wc -l`, filePath),
	}
	cli := strings.Join(clis, " && ")
	// log
	logger.Debugf("CliDetectTgzFolder: %s", strings.Join(clis, " && "))
	// handle success
	return cli
}
