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

var MapDaRepoTplFileContent = map[string]string{
	"rhel": `
		[{{.RepoName}}]
		enabled=1
		name={{.RepoName}}
		gpgcheck=1
		baseurl={{.UrlRepo}}
		gpgkey={{..UrlGpg}}
	`,

	"debian": `
		deb [signed-by={{.GpgFilePath}}] %{{UrlRepo}} /
	`,
}

// func init() {
// 	// alias fedora → rhel
// 	DaRepoTplFileName["fedora"] = DaRepoTplFileName["rhel"]
// }
