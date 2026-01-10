# Todo

## The cli `helm show`

- display chart metadata - chart.yaml
```shell
helm show chart $chartName/$RepoName
```

- display chart metadata - values.yaml
```shell
helm show values $chartName/$RepoName
```

- Display user defined values about a relase
```shell
# xxx kbe-cilicium  kube-system
helm get values %s -n %s`, args[0], args[1]
```

```shell
## display chart metadata - chart.yaml
helm show chart $chartName/$RepoName

## display chart metadata - values.yaml
helm show values $chartName/$RepoName
```

## release's values
```
releaseValueShortDesc = "Display user defined values about a relase"
```
Example: 
- `xxx kbe-cilicium  kube-system`
- `cli := fmt.Sprintf(`helm get values %s -n %s`, args[0], args[1])`

## The file `db.rpo.yaml`
```
# // CiliumHelmRepoUrl   = https://helm.cilium.io/
# // CiliumHelmRepoName  = cilium
# // CiliumHelmChartName = cilium
# // CiliumHelmChartVersion = 1.17  // compatible with K8s 1.32.x-1.33.x
# // CiliumK8sNamespace  = kube-system
# // CiliumHelmReleaseName = kbe-cilium

# // // Kubernetes/Helm Dashboard conf
# // KDashbHelmChartVersion = 7.12  // works with K8s 1.32.x
# // KDashbK8sNamespace = kdashb
# // KDashbHelmReleaseName = kbe-kdash

# // // Kubernetes/Helm Nginx controller conf
# // IngressNginxControllerHelmRepoName = ingnginx
# // IngressNginxControllerHelmChartName = ingress-nginx
# // IngressNginxControllerHelmChartVersion = 4.12
# // IngressNginxControllerK8sNamespace = ingress-nginx
# // IngressNginxControllerHelmReleaseName = kbe-ingress-nginx
```