package dnfapt

type DaRepository struct {
	Name        string // logical name
	FileName    string // Os file name
	Description string
	Version     string // the version of the package repository
	UrlRepo     string
	UrlGpg      string
}

type SliceDaRepository []DaRepository
type MapDaRepository map[string]DaRepository

func (list SliceDaRepository) GetListName() []string {
	names := make([]string, 0, len(list))
	for _, s := range list {
		names = append(names, s.Name)
	}
	return names
}
