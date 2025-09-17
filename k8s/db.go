package k8s

// define the cluster configuration file
var configFileTpl = `
		---
		apiVersion: kubeadm.k8s.io/v1beta4
		kind: InitConfiguration
		nodeRegistration:
		  ignorePreflightErrors:
		    - NumCPU
		
		---
		
		apiVersion: kubeadm.k8s.io/v1beta4
		kind: ClusterConfiguration
		kubernetesVersion: {{.K8sVersion}}
		networking:
		  podSubnet: {{.K8sPodCidr}}
		  serviceSubnet: {{.K8sServiceCidr}}
		
		---
		
		apiVersion: kubelet.config.k8s.io/v1beta1
		kind: KubeletConfiguration
		containerRuntimeEndpoint: {{.CrSocketName}}
		failSwapOn: false
		
		---
		
		apiVersion: kubeadm.k8s.io/v1beta4
		kind: JoinConfiguration
		nodeRegistration:    
		  ignorePreflightErrors:
		    - NumCPU
		---
		
		apiVersion: kubeadm.k8s.io/v1beta4
		kind: ResetConfiguration
		cleanupTmpDir: true
		`

// , config.K8sVersion, config.KbePodCidr, config.KbeServiceCidr, config.CrioSocketName)
