package osservice

// define types
type Service struct {
	Name     string // service name
	FileName string // service file conf file
}

// define slices
type ServiceSlice []Service

// define getters
func GetService(name string) *Service {
	return &Service{
		Name: name,
	}
}
