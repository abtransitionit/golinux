package other

// define types
type Manifest struct {
	Name string // custom name
	Desc string // custom name
	Url  string
	Doc  []string
}

// defrine slices
type ManifestSlice []Manifest

// define getters
func GetCli(i Manifest) *Manifest {
	return &Manifest{
		Name: i.Name,
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
	List map[string]Manifest
}
