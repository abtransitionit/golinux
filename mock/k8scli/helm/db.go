package helm

import (
	_ "embed"
	"path/filepath"
)

// -----------------------------------------
// ------ define file location -------------
// -----------------------------------------

//go:embed db.repo.yaml
var yamlListRepo []byte // automatically cache the raw yaml file in this var
//go:embed db.conf.yaml
var yamlConf []byte // automatically cache the raw yaml file in this var

var registryCfgRelPath = filepath.Join("wkspc", ".config", "registry", "credential.yaml")
var artifactCfgRelPath = filepath.Join("wkspc", ".config", "artifact", "cfg.yaml")
