## 🟢 Go for DevOps & Infrastructure (Questions 641-660)

### Question 641: How do you create a custom Kubernetes operator in Go?

**Answer:**
Use the **Operator SDK** (built on Kubebuilder).
1.  Initialize: `operator-sdk init`
2.  Create API: `operator-sdk create api --group=app --version=v1 --kind=MyApp`
3.  Implement `Reconcile` loop in the controller: Read declared state (CRD), Check actual state (Pods), Fix diff.

---

### Question 642: How do you write a Helm plugin in Go?

**Answer:**
A Helm plugin is just an executable binary.
1.  Structure: `plugin.yaml` and the binary.
2.  Logic: Go program that executes shell commands or manipulates files.
3.  Install: `helm plugin install https://github.com/my/plugin`.

---

### Question 643: How do you use Go for infrastructure automation?

**Answer:**
Go is the language of Cloud Native.
- **Terraform:** Write Providers (plugins) in Go.
- **Pulumi:** Define infrastructure (VPC, VM) using Go code (structs/methods) instead of YAML/HCL.

---

### Question 644: How do you write a CLI for managing AWS/GCP resources?

**Answer:**
Import the SDKs.
- `github.com/aws/aws-sdk-go-v2`
- `cloud.google.com/go`
Use these inside a **Cobra** CLI app to create custom workflows (e.g., "Reset Dev Env" -> Deletes EC2s, Clears S3).

---

### Question 645: How do you use Go to write Terraform providers?

**Answer:**
Implement the **Terraform Plugin SDK**.
Define `Resources` (Create/Read/Update/Delete methods) and `DataSources`.
This links HCL configuration blocks to real API calls of your service.

---

### Question 646: How do you build a dynamic inventory script in Go for Ansible?

**Answer:**
Ansible expects a script that outputs JSON with host groups.
Write a Go binary that:
1.  Queries AWS/Azure API for instances with tag `Role=Web`.
2.  Formats list into Ansible JSON structure.
3.  Prints to stdout.

---

### Question 647: How do you parse and generate YAML in Go?

**Answer:**
Use `gopkg.in/yaml.v3` (or v2).
Map structs to YAML with tags.
```go
type Config struct {
    Image string `yaml:"image"`
    Port  int    `yaml:"port"`
}
```
Crucial for manipulating K8s manifests programmatically.

---

### Question 648: How do you interact with Docker API in Go?

**Answer:**
Use the official client `github.com/docker/docker/client`.

```go
cli, _ := client.NewClientWithOpts(client.FromEnv)
containers, _ := cli.ContainerList(ctx, types.ContainerListOptions{})
for _, c := range containers {
    fmt.Println(c.ID, c.Image)
}
```

---

### Question 649: How do you manage Kubernetes CRDs i Go?

**Answer:**
Custom Resource Definitions (CRDs) extend the K8s API.
In Go, you define the Go struct (types.go) matching the CRD schema.
Use **controller-gen** to generate the DeepCopy methods and CRD YAML manifests automatically.

---

### Question 650: How do you write Go code to scale deployments in K8s?

**Answer:**
Use `client-go`.

```go
clientset, _ := kubernetes.NewForConfig(config)
deployments := clientset.AppsV1().Deployments("default")

scale, _ := deployments.GetScale(ctx, "my-app", metav1.GetOptions{})
scale.Spec.Replicas = 5
deployments.UpdateScale(ctx, "my-app", scale, metav1.UpdateOptions{})
```

---

### Question 651: How do you tail logs from containers using Go?

**Answer:**
Docker API or K8s API provides request methods that return an `io.ReadCloser` (Stream).
Stream this reader to `os.Stdout` to view logs in real-time within your tool.

---

### Question 652: How do you manage service discovery in Go apps?

**Answer:**
1.  **K8s DNS:** Just use hostname `http://my-service`. Go resolves it naturally.
2.  **Consul:** Use Consul API to register service on startup and query for peers.
3.  **Etcd:** Watch key updates for available nodes.

---

### Question 653: How do you build a Kubernetes admission controller in Go?

**Answer:**
It's a Webhook (HTTP Server).
K8s sends a `AdmissionReview` JSON payload (Pod creation request) to your Go server.
Your logic checks rules (e.g., "Image must come from internal registry").
Return `Allowed: true/false` JSON response.

---

### Question 654: How do you build a metrics exporter for Prometheus in Go?

**Answer:**
Write a Go app that queries a 3rd party system (e.g., a hardware raid controller) that has no monitoring.
Translate the values into Prometheus Metrics (`gauge.Set(val)`).
Expose `/metrics` HTTP endpoint.
Run as sidecar or cronjob.

---

### Question 655: How do you set up health checks for a Go microservice?

**Answer:**
(See Q436). Use `net/http` to expose endpoints.
For gRPC, implement the standard **GRPC Health Checking Protocol** (`google.golang.org/grpc/health`).

---

### Question 656: How do you build a custom load balancer in Go?

**Answer:**
Use `httputil.ReverseProxy`.
Implement a custom `Director`.
Store list of Backends. Use Round-Robin atomic counter to pick next backend.
Point the Proxy Director to that URL.

---

### Question 657: How do you implement graceful shutdown with Kubernetes SIGTERM?

**Answer:**
(See Q30 but specific to K8s).
1.  Catch `syscall.SIGTERM`.
2.  Set `readiness` probe to fail (stop traffic).
3.  Wait for existing requests to finish (`server.Shutdown(ctx)`).
4.  Exit.
This prevents dropping requests during Rolling Updates.

---

### Question 658: How do you use Go with Envoy/Consul service mesh?

**Answer:**
Go apps are usually unaware of the mesh (Envoy proxy runs as sidecar).
However, you can use the **xDS protocol** (Go implementation) to dynamically configure Envoy control planes yourself.

---

### Question 659: How do you configure Go apps for 12-factor principles?

**Answer:**
1.  **Config:** Connect via Env Vars (Viper).
2.  **Logs:** Write structured JSON to `Stdout` (no log files).
3.  **Port Binding:** Expose port via config.
4.  **Disposability:** Fast startup/shutdown.

---

### Question 660: How do you use Go for cloud automation scripts?

**Answer:**
Go is replacing Bash/Python for complex scripts.
Wait for resources (`helper.WaitForEC2`), parallel execution (`errgroup`), and strong typing prevent "script rot" where bash scripts fail silently on edge cases.

---
