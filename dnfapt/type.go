package dnfapt

type DaRepo struct {
	Name     string // logical name
	FileName string // Os file name
	Version  string // the version of the package repository
	UrlRepo  string
	UrlGpg   string
}

type DaRepoReference struct {
	Folder string
	Ext    string
	Pack   string
	Gpg    string
}

type DaOsRepoCte struct {
	Folder    string
	Ext       string
	Pack      string
	Gpg       string
	GpgFolder string
	GpgExt    string
}

type SliceDaRepo []DaRepo
type MapDaRepo map[string]DaRepo
type MapDaRepoCte map[string]DaOsRepoCte

func (list SliceDaRepo) GetListName() []string {
	names := make([]string, 0, len(list))
	for _, s := range list {
		names = append(names, s.Name)
	}
	return names
}

// Structure to resolve the content of a DaRepoTplFileContent
type RepoFileContentVar struct {
	RepoName    string
	UrlRepo     string
	UrlGpg      string
	GpgFilePath string
}

// template: data structure - parse template with this structure - Execute the template with this structure
