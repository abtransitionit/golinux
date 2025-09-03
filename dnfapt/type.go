package dnfapt

type DaRepository struct {
	Name        string // logical name
	FileName    string // Os file name
	Version     string // the version of the package repository
	Description string
}

type SliceDaRepository []DaRepository
type MapDaRepository map[string]DaRepository
