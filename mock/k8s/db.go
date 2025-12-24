package k8s

import (
	_ "embed"
)

// -----------------------------------------
// ------ define file location -------------
// -----------------------------------------

//go:embed db.cfg.yaml
var yamlCfg []byte // automatically cache the raw yaml file in this var
