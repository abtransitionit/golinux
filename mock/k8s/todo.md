# Best practice


# joining the worker the right way

this cli `sudo kubeadm token create --print-join-command` bypass the manifest `kind: JoinConfiguration`


# Step 1
Instead of asking the Control Plane for a string to copy, your Go tool should gather the three "secrets" required for a join:

1. **The API Server Endpoint:** (e.g., `51.210.10.195:6443`)
2. **The Token:** (e.g., `abcdef.1234567890abcdef`)
3. **The CA Cert Hash:** (e.g., `sha256:7c2...`)

# Step 2

Your `JoinConfiguration` template should look like this to be fully functional. Note the `discovery` block—this is what replaces the CLI arguments:

```yaml
apiVersion: kubeadm.k8s.io/v1beta4
kind: JoinConfiguration
discovery:
  bootstrapToken:
    apiServerEndpoint: "{{ .Cluster.ControlPlaneIP }}:6443"
    token: "{{ .Cluster.Token }}"
    caCertHashes:
      - "{{ .Cluster.CertHash }}"
nodeRegistration:    
  name: {{ .Node.CustomName }} # Now this will finally work!
  ignorePreflightErrors:
    - NumCPU
---
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
failSwapOn: false

```

# Step 3

1. **On Control Plane:** Run `kubeadm token create` and `openssl x509 -pubkey...` (or use the Go SDK) to get the token and hash.
2. **Locally:** Render your template with these values + the specific `CustomName` for that worker (e.g., `worker-01`).
3. **On Worker:** Upload the rendered file to `/etc/kubernetes/join-config.yaml`.
4. **On Worker:** Execute:
```bash
sudo kubeadm join --config /etc/kubernetes/join-config.yaml

```




# Why this is a "level up"

By using the config file, you gain **Granular Control**.
For example, if you wanted the Fedora node to use a different container runtime path than the Debian node, you could handle that logic entirely inside your Go template based on the OS, and `kubeadm` would respect it because it's reading your file, not its own defaults.

To get those custom names working, you need to extract the "connection secrets" from your control plane and feed them into your Go template.


# Get the Token and Hash

Run these on your **control plane** (Node 3 in your list) to get the data for your template:

**The Token:**

```bash
kubeadm token create
# Output example: abcdef.1234567890abcdef

```

**The CA Cert Hash:**

```bash
openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | \
openssl rsa -pubin -outform der 2>/dev/null | \
openssl dgst -sha256 -hex | sed 's/^.* //'
# Output example: 7c2... (a long hex string)

```

# 2. Update your `goluc` Template

Now, merge these into your YAML structure. This ensures that when you run `kubeadm join --config`, the worker knows exactly where to go and what its name should be.

```yaml
apiVersion: kubeadm.k8s.io/v1beta4
kind: JoinConfiguration
discovery:
  bootstrapToken:
    apiServerEndpoint: "51.210.10.195:6443"
    token: "{{ .Cluster.Token }}"                # From 'token create'
    caCertHashes: 
      - "sha256:{{ .Cluster.CertHash }}"         # From 'openssl' command
nodeRegistration:    
  name: {{ .Node.CustomName }}                   # e.g., "worker-fedora"
  ignorePreflightErrors:
    - NumCPU
---
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
failSwapOn: false

```

# 3. The "Impact" Check

When you re-run your process using `sudo kubeadm join --config /path/to/rendered.yaml`, you will see:

1. **Custom Names:** `kubectl get nodes` will show `worker-fedora`, `worker-rocky`, etc.
2. **Cilium Cleanliness:** Cilium will automatically use these names for its internal node-to-node health checks.
3. **No More "None" Roles:** If you want to label them as workers during join, you can even add a `taints` or `labels` section to that same YAML.

# One quick tip for your multi-distro setup:

Since you are using **Fedora** and **AlmaLinux**, watch out for **SELinux**. If it's set to `Enforcing`, you might need to add a few lines to your Go tool to run `setenforce 0` or install the `container-selinux` policies before joining, otherwise the Kubelet might struggle to talk to the container runtime.

