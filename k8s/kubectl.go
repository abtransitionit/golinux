package k8s

import (
	"strings"
)

func ConfigureKubectlOnCPlane() string {

	// build the CLI
	var clis = []string{
		`sudo cat /etc/kubernetes/admin.conf | install -D -m 600 /dev/stdin ~/.kube/config`,
	}
	cli := strings.Join(clis, " && ")

	// return
	return cli
}

// 1 - get the initial config from the cplane ([]byte)
// 2 - install it on the kubectl node (~/.kube/config)

// ssh o1u "cat /etc/kubernetes/admin.conf"  > ~/.kube/config
// sudo cat /etc/kubernetes/admin.conf > ~/.kube/config
// chmod 600 ~/.kube/config
// export KUBECONFIG=~/.kube/config
