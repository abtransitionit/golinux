# install cilium ingress
**define the 	value file**
```yaml
# Enable Envoy (required for Ingress)
envoy:
  enabled: true

# Enable Cilium Ingress Controller
ingressController:
  enabled: true
  default: true

# Recommended for modern clusters
kubeProxyReplacement: true

# Required capabilities
l7Proxy: true

# Optional but strongly recommended
operator:
  replicas: 2

# Enable NodePort for ingress traffic
nodePort:
  enabled: true
```

**install it**
```sh
helm upgrade --install cilium cilium/cilium \
  --namespace kube-system \
  -f values.yaml
```
# example to test cilium ingress
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: test
spec:
  ingressClassName: cilium
  rules:
  - host: test.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: my-service
            port:
              number: 80
```