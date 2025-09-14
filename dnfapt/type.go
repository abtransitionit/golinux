package dnfapt

type DaRepo struct {
	Name     string // logical name
	FileName string // Os file name
	Version  string // the version of the package repository
	UrlRepo  string
	UrlGpg   string
}

type DaPack struct {
	Name    string // logical name
	CName   string // canonical name
	Version string // the version of the package
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

// Structure to resolve the content of a DaRepoTplFileContent
type RepoFileContentVar struct {
	RepoName    string
	UrlRepo     string
	UrlGpg      string
	GpgFilePath string
}

type SliceDaRepo []DaRepo
type SliceDaPack []DaPack
type MapDaRepo map[string]DaRepo
type MapDaPack map[string]DaPack
type MapDaRepoCte map[string]DaOsRepoCte

func (list SliceDaRepo) GetListName() []string {
	names := make([]string, 0, len(list))
	for _, s := range list {
		names = append(names, s.Name)
	}
	return names
}

func (list SliceDaPack) GetListName() []string {
	names := make([]string, 0, len(list))
	for _, s := range list {
		names = append(names, s.Name)
	}
	return names
}

// template: data structure - parse template with this structure - Execute the template with this structure
