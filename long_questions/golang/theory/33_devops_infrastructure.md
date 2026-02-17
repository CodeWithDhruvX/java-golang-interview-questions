# 🟢 Go Theory Questions: 641–660 Go for DevOps & Infrastructure

## 641. How do you create a custom Kubernetes operator in Go?

**Answer:**
We use **Kubebuilder** or **Operator SDK**.
1.  `kubebuilder init`.
2.  Define CRD struct `type MyResource struct { Spec ... Status ... }`.
3.  Implement the **Reconcile Loop**.
`Reconcile(req)`: Fetch the resource. Compare "Desired State" (Spec) vs "Actual State" (Cluster). Apply fixes (create missing pods).
It basically puts the Logic of a human operator into a Go binary loop.

---

## 642. How do you write a Helm plugin in Go?

**Answer:**
A Helm plugin is just a binary installed in `~/.helm/plugins`.
We write a Go CLI (using Cobra).
Example: `helm s3-push`.
The plugin `plugin.yaml` points to our binary.
The binary executes logic (e.g., uploading the Chart tarball to an S3 bucket) using the AWS SDK, extending Helm's capabilities with native Go power.

---

## 643. How do you use Go for infrastructure automation?

**Answer:**
Go is the language of Cloud (Terraform, K8s, Docker are Go).
We use **Pulumi (Go SDK)** or **Terraform CDK**.
Instead of HCL/YAML, we write actual Go code:
`s3.NewBucket(ctx, "my-bucket", ...)`
This gives us loops, functions, compiler checks, and testing for our infrastructure definitions, which static YAML cannot provide.

---

## 644. How do you write a CLI for managing AWS/GCP resources?

**Answer:**
We embed the Cloud SDKs (`aws-sdk-go-v2`).
We define a Command Pattern.
`myapp create-vm --size large`.
Internally makes `ec2.RunInstances` call.
Crucially, we handle **Long Running Ops**. We poll the API: `waiter := ec2.NewInstanceRunningWaiter(...)`. This makes the CLI synchronous and scriptable ("Wait until VM is up, then SSH").

---

## 645. How do you use Go to write Terraform providers?

**Answer:**
We use `hashicorp/terraform-plugin-sdk`.
We map Go functions to CRUD operations:
`Create: resourceCreate`
`Read:   resourceRead`
`Update: resourceUpdate`
`Delete: resourceDelete`
The Terraform Core calls our binary (via RPC) to "Apply" changes. This is how you allow Terraform to manage a custom internal API or a new SaaS product.

---

## 646. How do you build a dynamic inventory script in Go for Ansible?

**Answer:**
Ansible expects a script to output JSON:
`{ "webservers": ["10.0.0.1", "10.0.0.2"] }`.
We write a Go binary that queries our CMDB (or generic internal API/consul).
`./inventory-go --list`.
Go is perfect here because it's a single binary (no dependency hell on the Ansible controller) and runs instantly compared to Python scripts that might need to import heavy libraries.

---

## 647. How do you parse and generate YAML in Go?

**Answer:**
We use `gopkg.in/yaml.v3`.
It works exactly like `encoding/json`.
`yaml.Unmarshal(bytes, &configStruct)`.
We use mapstructure tags `yaml:"port"` to map fields.
Note: YAML is strict about indentation. The library handles this, but we often treat config as "Read in YAML, process in struct, Write out JSON" for logging consistency.

---

## 648. How do you interact with Docker API in Go?

**Answer:**
`github.com/docker/docker/client`.
`cli, _ := client.NewClientWithOpts(...)`.
`cli.ContainerList(ctx, types.ContainerListOptions{})`.
This allows us to write "Sidecars" or custom agents that orchestrate containers on the host, monitor their stats, or reap zombies, avoiding the need to shell out to the `docker` CLI binary.

---

## 649. How do you manage Kubernetes CRDs in Go?

**Answer:**
CRD (Custom Resource Definition) extends the K8s API.
In Go, we define the types in `apis/v1alpha1/types.go`.
We run `controller-gen` (part of Kubebuilder) to generate the deepcopy methods and the actual YAML manifest for the CRD.
The client then treats `MyResource` just like `Pod` or `Deployment`, with full type safety.

---

## 650. How do you write Go code to scale deployments in K8s?

**Answer:**
We use the `client-go` library to patch the `Replicas` field.
1.  Get Deployment: `clientset.AppsV1().Deployments(ns).Get(...)`.
2.  Update Replicas: `dep.Spec.Replicas = pointer.Int32(5)`.
3.  Update: `clientset.AppsV1().Deployments(ns).Update(...)`.
This allows building custom Autoscalers that react to metrics K8s HPA doesn't support (e.g., "Twitter Sentiment" or "Inventory Stock Level").

---

## 651. How do you tail logs from containers using Go?

**Answer:**
`cli.ContainerLogs(ctx, id, options)`.
It returns an `io.ReadCloser`.
We must handle the **Docker Multiplexing Format**: The stream contains [Header | Payload]. The header says if it's STDOUT or STDERR.
We use `stdcopy.StdCopy(os.Stdout, os.Stderr, reader)` to demultiplex this correctly, otherwise the logs look like garbage characters.

---

## 652. How do you manage service discovery in Go apps?

**Answer:**
(See Q 368).
If running in K8s, we use the K8s API to find Endpoints.
`endpoints, _ := k8s.CoreV1().Endpoints(ns).Get("myservice")`.
This gives us the list of Pod IPs backing a service. We can client-side load balance (Round Robin) directly to these IPs, bypassing the Service VIP (ClusterIP) for potentially lower latency (headless service).

---

## 653. How do you build a Kubernetes admission controller in Go?

**Answer:**
An Admission Controller is a Webhook (HTTP Server).
K8s API Server sends `AdmissionReview` JSON to us **before** creating a pod.
We check logic: "Does this image start with `mycompany.com/`?".
If no, return `Allowed: false`.
In Go, this is just a standard HTTP handler that parses the JSON, validates the spec, and returns a JSON response allowing or denying the request.

---

## 654. How do you build a metrics exporter for Prometheus in Go?

**Answer:**
If we have a "Black Box" (like a legacy fancy Appliance) that doesn't speak Prometheus.
We write a Go Exporter.
1.  Scrape/Query the Appliance (e.g., via SNMP or SOAP).
2.  Convert values to `prometheus.Gauge`.
3.  Expose `/metrics`.
This "Adapter Pattern" allows us to monitor non-cloud-native stuff using standard cloud-native tools.

---

## 655. How do you set up health checks for a Go microservice?

**Answer:**
Middleware pattern.
`/healthz` -> Returns 200 OK.
We can implement a `Checker` interface.
`health.Register("db", db.Ping)`.
`health.Register("redis", redis.Ping)`.
The handler runs all checkers parallel. If any fail, return 500. This gives the Load Balancer a signal to take the node out of rotation.

---

## 656. How do you build a custom load balancer in Go?

**Answer:**
`httputil.ReverseProxy` + `Director` modification.
We maintain a list of backends `[]*url.URL`.
In the `Director` function, we pick one (Round Robin or Random).
`req.URL.Host = selectedBackend.Host`.
Go handles the rest (piping the request and response).
We can add Health Checking logic: If a backend fails 3 times, remove it from the slice.

---

## 657. How do you implement graceful shutdown with Kubernetes SIGTERM?

**Answer:**
K8s sends SIGTERM. It waits (default 30s) before SIGKILL.
In Go:
1.  Listen for signal: `signal.Notify(c, syscall.SIGTERM)`.
2.  Block until signal received: `<-c`.
3.  Call `server.Shutdown(ctx)`. This stops accepting new connections but allows existing requests to finish.
This is mandatory to prevent 502 errors during rolling updates.

---

## 658. How do you use Go with Envoy/Consul service mesh?

**Answer:**
Envoy can be extended using WASM or External Authorization servers.
We can write a Go gRPC service that acts as an **External Authz**.
Envoy asks our Go service: "Can User X access Path Y?".
Go checks OPA policy or DB.
Replie: Yes/No.
This centralizes policy enforcement out of the services and into the mesh layer.
