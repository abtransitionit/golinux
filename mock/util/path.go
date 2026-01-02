package util

import (
	"github.com/abtransitionit/gocore/logx"
)

// description: build and return a sequence of paths from a root folder
//
// Notes:
//
// - the sequence is separated by the char':'
func getTreePath(logger logx.Logger, rootFolder string) error {
	// log
	logger.Debugf("getTreePath called with rootFolder: %s", rootFolder)
	// handle success
	return nil

}
