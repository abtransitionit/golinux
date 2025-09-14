package oskernel

import (
	"path/filepath"
	"strings"
)

func LoadOsKParam() (string, error) {

	var cmds = []string{
		// if needed sudo sysctl --system -p /etc/sysctl.d/99-kbe.conf
		"sudo sysctl --system",
	}
	cli := strings.Join(cmds, " && ")
	return cli, nil
}

func GetKParamFilePath(kernelFilename string) string {
	return filepath.Join(MapOsKernelReference["KParamFolder"], kernelFilename)
}

func (s SliceOsKParam) GetContent() string {
	var b strings.Builder
	for _, param := range s {
		b.WriteString("# ")
		b.WriteString(param.Description)
		b.WriteString("\n")
		b.WriteString(param.Kvp)
		b.WriteString("\n")
	}
	return b.String()
}

// apply changes at runtime => cli = fmt.Sprintf(`sudo /sbin/sysctl --system`)

// KParamConf := fmt.Sprintf(`
// 	# Enable IP forwarding - kernel parameter
// 	net.ipv4.ip_forward = 1

// 	# Allow bridged IPv4 traffic to go through iptables - br_netfilter module parameter
// 	# Enable filtering of bridged traffic
// 	net.bridge.bridge-nf-call-iptables = 1

// 	# Allow bridged IPv6 traffic to go through iptables - br_netfilter module parameter
// 	# Enable filtering of IPv6 traffic on bridged interfaces
// 	net.bridge.bridge-nf-call-ip6tables = 1

// 	# Enable filtering of bridged traffic for all interfaces
// 	# net.bridge.bridge-nf-call = 1
// `)
