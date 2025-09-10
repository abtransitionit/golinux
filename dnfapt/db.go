package dnfapt

var MapDaRepoReference = MapDaRepo{
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

var MapDaRepoCteReference = MapDaRepoCte{
	"debian": {
		Folder:    "/etc/apt/sources.list.d",
		Ext:       ".list",
		Pack:      "deb",
		Gpg:       "Release.key",
		GpgFolder: "/etc/apt/keyrings",
		GpgExt:    "-apt-keyring.gpg",
	},
	"rhel": {
		Folder: "/etc/yum.repos.d",
		Ext:    ".repo",
		Pack:   "rpm",
		Gpg:    "repodata/repomd.xml.key",
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

func init() {
	// lookup fedora → lookuprhel
	MapDaRepoCteReference["fedora"] = MapDaRepoCteReference["rhel"]
	MapDaRepoTplFileContent["fedora"] = MapDaRepoTplFileContent["rhel"]
}
