package onpm

type MapPack map[string]Pack
type MapRepo map[string]Repo3

type Repo3 struct {
	Name     string // logical name
	FileName string // Os file name
	Version  string // the version of the package repository
	UrlRepo  string
	UrlGpg   string
}

type Pack struct {
	Name    string // logical name
	CName   string // canonical name
	Version string // the version of the package
}

var MapDaPackReference = MapPack{
	"crio": {
		CName: "cri-o",
	},
	"kubeadm": {
		CName: "kubeadm",
	},
	"kubelet": {
		CName: "kubelet",
	},
	"kubectl": {
		CName: "kubectl",
	},
}
var MapDaRepoReference = MapRepo{
	"crio": {
		Name:    "crio",
		UrlRepo: "https://download.opensuse.org/repositories/isv:/cri-o:/stable:/v$TAG/$PACK/",
		UrlGpg:  "https://download.opensuse.org/repositories/isv:/cri-o:/stable:/v$TAG/$PACK/$GPG",
	},
	"k8s": {
		Name:    "k8s",
		UrlRepo: "https://pkgs.k8s.io/core:/stable:/v$TAG/$PACK/",
		UrlGpg:  "https://pkgs.k8s.io/core:/stable:/v$TAG/$PACK/$GPG",
	},
}

var MapDaRepoTplFileContent = map[string]string{
	"rhel": `
		[{{.RepoName}}]
		enabled=1
		name={{.RepoName}}
		gpgcheck=1
		baseurl={{.UrlRepo}}
		gpgkey={{.UrlGpg}}
	`,

	"debian": `
		deb [signed-by={{.GpgFilePath}}] {{.UrlRepo}} /
	`,
}

// # debian:
// # 	Folder:    "/etc/apt/sources.list.d",
// # 	Gpg:       "Release.key",
// # 	GpgExt:    "-apt-keyring.gpg",  // /usr/share/keyrings/<repo>.gpg and reference in .list
// ## 	Ext:       ".list",
// ## 	Pack:      "deb",
// ## 	GpgFolder: "/etc/apt/keyrings", //
// # },
// # "rhel": {
// # 	Folder: "/etc/yum.repos.d",
// # 	Gpg:    "repodata/repomd.xml.key",
// ## 	Ext:    ".repo",
// ## 	Pack:   "rpm",
