package envar

type EnvVar struct {
	Name  string
	Value string
}

type SliceEnvVar []EnvVar

func (envVar EnvVar) GetName() string {
	return envVar.Name
}

func (envVar EnvVar) GetValue() string {
	return envVar.Value
}
