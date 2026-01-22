package other

import (
	_ "embed"
)

// -----------------------------------------
// ------ define file location -------------
// -----------------------------------------

//go:db.manifest.yaml
var yamlList []byte // automatically cache the raw yaml file in this var
