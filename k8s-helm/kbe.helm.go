package helm

// Helm configuration
type HelmRepoProperty struct {
	Name        string
	Description string
	Url         string
	FlagVar     *bool
}

var ListHelmRepo = []HelmRepoProperty{
	{KDashbHelmRepoName, "The standard kubernetes dashboard Helm repository", KDashbHelmRepoUrl, new(bool)},
	{CiliumHelmRepoName, "The cilium   Helm repository", CiliumHelmRepoUrl, new(bool)},
	{"bitnami", "The bitnami  Helm repository", "https://charts.bitnami.com/bitnami", new(bool)},
	{IngressNginxControllerHelmRepoName, "The nginx ingress controler Helm repository", IngressNginxControllerHelmRepoUrl, new(bool)},
}

const ()

var HelmHost = KbeCplaneNode
