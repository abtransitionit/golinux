module github.com/abtransitionit/golinux

go 1.24.2

// prod mode
require (
	github.com/abtransitionit/gocore v1.0.0
)

// dev mode
replace github.com/abtransitionit/gocore => ../gocore
