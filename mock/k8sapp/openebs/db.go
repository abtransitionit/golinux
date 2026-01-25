package openebs

import (
	_ "embed"
)

// -----------------------------------------
// ------ define file location -------------
// -----------------------------------------

//go:embed db.cfg.basic.yaml
var YamlBasicCfg []byte // automatically cache the raw yaml file in this var
