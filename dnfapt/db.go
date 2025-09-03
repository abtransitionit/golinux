package dnfapt

var DaRepositoryReference = MapDaRepository{
	"crio": {
		Name:    "crio",
		UrlRepo: "https://download.opensuse.org/repositories/isv:/cri-o:/stable:/$TAG/$PACK/",
		UrlGpg:  "https://download.opensuse.org/repositories/isv:/cri-o:/stable:/$TAG/$PACK/$GPG",
	},
	"k8s": {
		Name:    "k8s",
		UrlRepo: "https://pkgs.k8s.io/core:/stable:/$TAG/$PACK/",
		UrlGpg:  "https://pkgs.k8s.io/core:/stable:/$TAG/$PACK/$GPG",
	},
}
