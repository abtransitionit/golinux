package cli

// define types
type Kubectl struct {
	CplaneNdeName   string
	InstallNodeName string
}

// define getters
func GetKubectl(cplaneNodeName string, installNodeName string) *Kubectl {
	return &Kubectl{
		CplaneNdeName:   cplaneNodeName,
		InstallNodeName: installNodeName,
	}
}
