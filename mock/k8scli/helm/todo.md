# todo

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