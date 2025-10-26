package da

import (
	"fmt"
	"strings"
)

var MapRepoReference = MapRepo{
	"crio": {
		Name: "crio",
		Url: Url{
			Repo: "https://download.opensuse.org/repositories/isv:/cri-o:/stable:/v$TAG/$PACK/",
			Gpg:  "https://download.opensuse.org/repositories/isv:/cri-o:/stable:/v$TAG/$PACK/$GPG",
		},
	},
	"k8s": {
		Name: "k8s",
		Url: Url{
			Repo: "https://pkgs.k8s.io/core:/stable:/v$TAG/$PACK/",
			Gpg:  "https://pkgs.k8s.io/core:/stable:/v$TAG/$PACK/$GPG",
		},
	},
}

// Method to convert MapRepo to tab-separated string
func (m MapRepo) ConvertToString() string {
	var sb strings.Builder
	sb.WriteString("Name\tRepo\tGpg\tMgr\n") // header

	for _, repo := range m {
		sb.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n",
			repo.Name, repo.Url.Repo, repo.Url.Gpg, repo.Cbd))
	}

	return sb.String()
}

func (m MapRepo) ConvertToStringTruncated() string {
	var sb strings.Builder
	sb.WriteString("Name\tRepo\tGpg\tMgr\n") // header

	truncate := func(s string, n int) string {
		if len(s) <= n {
			return s
		}
		return s[:n] + "…" // add ellipsis to indicate truncation
	}

	for _, repo := range m {
		sb.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n",
			repo.Name,
			truncate(repo.Url.Repo, 50),
			truncate(repo.Url.Gpg, 50),
			repo.Cbd,
		))
	}

	return sb.String()
}

var configFileTpl = `
  apt:
    pkg: deb
    ext: .list
    folder:
      repo: "/etc/apt/sources.list.d"
      gpgKey: "/usr/share/keyrings"
  dnf:
    pkg: rpm
    ext: .repo
    folder:
      repo: "/etc/yum.repos.d"
    os:
      family: {{.Os.Family}}
      distro: {{.Os.Distro}}
`

// var MapDaRepoCteReference = MapDaRepoCte{
// 	"debian": {
// 		Folder:    "/etc/apt/sources.list.d",
// 		Ext:       ".list",
// 		Pack:      "deb",
// 		Gpg:       "Release.key",
// 		GpgFolder: "/etc/apt/keyrings", //
// 		GpgExt:    "-apt-keyring.gpg",  // /usr/share/keyrings/<repo>.gpg and reference in .list
// 	},
// 	"rhel": {
// 		Folder: "/etc/yum.repos.d",
// 		Ext:    ".repo",
// 		Pack:   "rpm",
// 		Gpg:    "repodata/repomd.xml.key",
// 	},
// }

// var MapDaRepoTplFileContent = map[string]string{
// 	"rhel": `
// 		[{{.RepoName}}]
// 		enabled=1
// 		name={{.RepoName}}
// 		gpgcheck=1
// 		baseurl={{.UrlRepo}}
// 		gpgkey={{.UrlGpg}}
// 	`,

// 	"debian": `
// 		deb [signed-by={{.GpgFilePath}}] {{.UrlRepo}} /
// 	`,
// }

// func init() {
// 	// lookup fedora → lookuprhel
// 	MapDaRepoCteReference["fedora"] = MapDaRepoCteReference["rhel"]
// 	MapDaRepoTplFileContent["fedora"] = MapDaRepoTplFileContent["rhel"]
// }

// // Todo in next version
