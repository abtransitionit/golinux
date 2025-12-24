package oskernel

import (
	_ "embed"
)

//go:embed db.yaml
var yamlCfg []byte // automatically cache the raw yaml file in this var
