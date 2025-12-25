package osservice

// define types
type Service struct {
	Name string   // service name
	File struct { // service conf file
		Name    string
		Path    string
		Content string
	}
}

// define slices
type ServiceSlice []Service

// define getters
func GetService(name string) *Service {
	return &Service{
		Name: name,
	}
}
