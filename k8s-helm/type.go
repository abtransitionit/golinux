package helm

type HelmRepo struct {
	Name string // logical name
	Url  string
}

type MapHelmRepo map[string]HelmRepo
