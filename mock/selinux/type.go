package selinux

// Selinux
type Selinux struct {
	CfgFilePath string
}

func GetSelinux() *Selinux {
	return &Selinux{
		CfgFilePath: "/etc/selinux/config",
	}
}
