package cilium

import (
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/mock/tpl"
	"github.com/abtransitionit/golinux/mock/k8scli/kubectl"
)

func GetValueFile(param map[string]string, logger logx.Logger) ([]byte, error) {
	var varPlaceHolder map[string]map[string]string

	// get placeholder
	varPlaceHolder, err := getPlaceholder(param, logger)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	// load parameter file
	yamlCfgRenderAsByte, err := tpl.LoadTplFile(YamlBasicCfg, "", varPlaceHolder)
	if err != nil {
		return nil, fmt.Errorf("loading config template file: %v", err)
	}

	// handle success
	return yamlCfgRenderAsByte, nil
}

func getPlaceholder(param map[string]string, logger logx.Logger) (varPlaceHolder map[string]map[string]string, err error) {

	// 1 - define helm host
	helmHost := "o1u"

	// 2 - define var placeholder if any
	// 21 - get PodCidr
	podCidr, ok := param["podcidr"]
	if !ok {
		return nil, fmt.Errorf("podcidr not found in the map")
	}

	// 22 - get ApiServerIp
	i := kubectl.Resource{Type: kubectl.ResApiServer}
	apiServerIp, err := i.GetIp("local", helmHost, logger)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	varPlaceHolder = map[string]map[string]string{
		"Cluster": {
			"PodCidr":     strings.TrimSpace(podCidr),
			"ApiServerIp": strings.TrimSpace(apiServerIp),
		},
	}

	// // 23 - get the resolved value file as byte[]
	// cfgAsbyte, err := cilium.GetValueFile(cilium.YamlBasicCfg, varPlaceHolder, logger)
	// if err != nil {
	// 	return fmt.Errorf("%s:%s:%s > getting value file > %w", hostName, helmHost, i.Name, err)
	// }
	return varPlaceHolder, nil
}
