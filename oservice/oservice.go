package oservice

func EnableLinger() string {
	return "loginctl enable-linger"
}

// Disable for the current user services to runs after a logout
func DissableLinger() string {
	return "loginctl disable-linger"
}

// func StartService(osService OsService) string {
// 	var cmds = []string{
// 		"sudo systemctl daemon-reload",
// 		fmt.Sprintf("sudo systemctl start %s", action, service)
// 		"echo $tmpfile",
// 	}
// 	cli := strings.Join(cmds, " && ")
// 	return cli

// }
