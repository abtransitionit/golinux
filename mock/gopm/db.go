package gopm

import (
	_ "embed"
)

// -----------------------------------------
// ------ define file location -------------
// -----------------------------------------

//go:embed db.yaml
var yamlList []byte // automatically cache the raw yaml file in this var
