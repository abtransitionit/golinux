package osservice

type Service struct {
	Name     string // service name
	FileName string // service file conf file
}

type ServiceSlice []Service
