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

var DaRepoTplFileContent = map[string]string{
	"rhel": `
		[%data.CName]
		name=%data.CName
		enabled=1
		gpgcheck=1
		baseurl=%data.UrlRepo
		gpgkey=%data.UrlGpg
	`,

	"debian": `
		deb [signed-by=%data.GpgFilePath] %data.UrlRepo /
	`,
}

// func init() {
// 	// alias fedora → rhel
// 	DaRepoTplFileName["fedora"] = DaRepoTplFileName["rhel"]
// }
