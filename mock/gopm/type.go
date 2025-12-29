package gopm

// define types
type Cli struct {
	Name    string // custom name
	Type    string // ie. tgz, zip, exe
	Version string
	Url     string
	Doc     []string
}

// defrine slices
type CliSlice []Cli

// define getters
func GetCli(cli Cli) *Cli {
	return &Cli{
		Name:    cli.Name,
		Version: cli.Version,
	}
}

// -------------------------------------------------------
// -------	 struct for YAML Cli List
// -------------------------------------------------------

// Description: represents the organization's repository db for go CLI(s)
//
// Notes:
//   - Manage the YAML repo file
//   - denotes the whitelist
type MapYaml struct {
	List map[string]Cli
}
