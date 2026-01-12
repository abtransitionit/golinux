package cilium

import (
	_ "embed"
)

// -----------------------------------------
// ------ define file location -------------
// -----------------------------------------

//go:embed db.cfg.yaml
var YamlCfg []byte // automatically cache the raw yaml file in this var
