module github.com/abtransitionit/golinux

// go toolchain version
go 1.24.2

// prod mode
require (
	github.com/abtransitionit/gocore v1.0.0
)

// used in dev mode - removes by CI at tag step
replace github.com/abtransitionit/gocore => ../gocore
