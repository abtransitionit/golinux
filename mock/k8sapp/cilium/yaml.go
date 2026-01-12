package cilium

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/mock/tpl"
)

func GetValueFile(embeddedTplFile []byte, varPlaceHolder map[string]map[string]string, logger logx.Logger) ([]byte, error) {
	var yamlCfgRenderAsByte []byte
	var err error

	// load file
	if yamlCfgRenderAsByte, err = tpl.LoadTplFile(embeddedTplFile, "", varPlaceHolder); err != nil {
		return nil, fmt.Errorf("loading config template file: %v", err)
	}

	// handle success
	return yamlCfgRenderAsByte, nil
}
