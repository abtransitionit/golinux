package openebs

import (
	_ "embed"
)

// -----------------------------------------
// ------ define file location -------------
// -----------------------------------------

//go:embed db.cfg.basic.yaml
var YamlBasicCfg []byte // automatically cache the raw yaml file in this var

//go:embed db.cfg.hostpath.yaml
var YamlHostPathCfg []byte // automatically cache the raw yaml file in this var

//go:embed db.cfg.localpv.yaml
var YamlLocalPvCfg []byte // automatically cache the raw yaml file in this var
