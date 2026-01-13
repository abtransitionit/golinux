package cilium

import (
	_ "embed"
)

// -----------------------------------------
// ------ define file location -------------
// -----------------------------------------

//go:embed db.cfg.basic.yaml
var YamlBasicCfg []byte // automatically cache the raw yaml file in this var

//go:embed db.cfg.all.yaml
var YamlAllCfg []byte // automatically cache the raw yaml file in this var

//go:embed db.cfg.ingress.yaml
var YamlIngressCfg []byte // automatically cache the raw yaml file in this var
