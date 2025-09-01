package oservice

func EnableLinger() string {
	return "loginctl enable-linger"
}

// Disable for the current user services to runs after a logout
func DissableLinger() string {
	return "loginctl disable-linger"
}
