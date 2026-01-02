package util

import (
	"github.com/abtransitionit/gocore/logx"
)

// description: build and return a sequence of path from a root folder
//
// Notes:
//
// - the sequence is separated by the char':'
func GetTreePath(rootFolder string, logger logx.Logger) (string, error) {
	// define var
	path := ""
	// log
	logger.Debugf("getTreePath called with rootFolder: %s", rootFolder)
	// handle success
	return path, nil

}
