package kubectl

import (
	_ "embed"
)

// -----------------------------------------
// ------ define file location -------------
// -----------------------------------------

//go:embed db.conf.yaml
var yamlConf []byte // automatically cache the raw yaml file in this var
