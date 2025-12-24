package k8s

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/mock/tpl"
)

func getFullClusterConf(embeddedTplFile []byte, clusterParam ClusterParam, logger logx.Logger) ([]byte, error) {
	var yamlCfgRenderAsByte []byte
	var err error

	// define var placeholder
	varPlaceHolder := map[string]map[string]string{
		"Cluster": {
			"Version":     clusterParam.K8sVersion,
			"PodCidr":     clusterParam.PodCidr,
			"ServiceCidr": clusterParam.ServiceCidr,
		},
		"CRuntime": {
			"SocketName": clusterParam.CrSocketName,
		},
	}
	// load file
	if yamlCfgRenderAsByte, err = tpl.LoadTplFile(embeddedTplFile, "", varPlaceHolder); err != nil {
		return nil, fmt.Errorf("loading config template file: %v", err)
	}
	// // log
	// logger.Debug("--- BEGIN:Rendered kubeadm config  ---")
	// scanner := bufio.NewScanner(bytes.NewReader(yamlCfgRenderAsByte))
	// for scanner.Scan() {
	// 	logger.Debug(scanner.Text())
	// }
	// if err := scanner.Err(); err != nil {
	// 	logger.Warnf("error while logging rendered kubeadm config: %v", err)
	// }
	// logger.Debug("--- END ---")

	// handle success
	return yamlCfgRenderAsByte, nil
}

// type ClusterVars struct {
// 	Version     string
// 	PodCidr     string
// 	ServiceCidr string
// }

// type CRuntimeVars struct {
// 	SocketName string
// }

// type TemplateVars struct {
// 	Cluster  ClusterVars
// 	CRuntime CRuntimeVars
// }

// varPlaceHolder := TemplateVars{
// 	Cluster: ClusterVars{
// 		Version:     clusterParam.K8sVersion,
// 		PodCidr:     clusterParam.PodCidr,
// 		ServiceCidr: clusterParam.ServiceCidr,
// 	},
// 	CRuntime: CRuntimeVars{
// 		SocketName: clusterParam.CrSocketName,
// 	},
// }
