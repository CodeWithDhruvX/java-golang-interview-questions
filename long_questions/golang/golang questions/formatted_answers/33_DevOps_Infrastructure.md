# 🟢 **641–660: Go for DevOps & Infrastructure**

### 641. How do you create a custom Kubernetes operator in Go?
"I use the **Operator SDK** or **Kubebuilder**.
It scaffolds the project.
I define a CRD struct: `type MyDB struct { Spec MySpec }`.
I implement the `Reconcile` loop:
`func (r *Reconciler) Reconcile(ctx, req) (Result, error)`.
This loop reads desired state (CRD) vs actual state (Pods), and creates/updates K8s objects to match."

#### Indepth
**Level 1 vs Level 5**. Most operators start as "Basic Install" (Level 1). The goal is "Autopilot" (Level 5): handling upgrades, backups, failure recovery, and vertical scaling without human intervention. The SDK provides `Operator Lifecycle Manager (OLM)` capabilities to package these advanced features.

---

### 642. How do you write a Helm plugin in Go?
"A Helm plugin is just a binary executed by Helm.
I write a Go CLI (using Cobra).
I add a `plugin.yaml` describing the hooks (`helm myplugin install`).
My Go code accesses Helm's env vars (`HELM_BIN`, `HELM_NAMESPACE`) to interact with the cluster or decrypt secrets before the chart installs."

#### Indepth
**Downloader Plugins**. Helm supports custom protocol handlers (`s3://`, `git://`). You can write a Go plugin that registers itself as a downloader, allowing `helm install s3://my-bucket/chart.tgz` to work seamlessly. This is great for private chart repositories.

---

### 643. How do you use Go for infrastructure automation?
"I prefer Go over Bash for complex logic.
I use `os/exec` or native Cloud SDKs.
Go's strict error handling prevents the 'script continued after error' bugs common in Bash.
I compile it to a single binary `ops-tool` and distribute it to the team—no `pip install` or Ruby gems required."

#### Indepth
`Mage` is a popular Make-alternative written in Go. Instead of a `Makefile`, you write `magefile.go`. This gives you the full power of Go (imports, loops, type safety) for your build scripts, which is much more maintainable than complex Bash spaghetti.

---

### 644. How do you write a CLI for managing AWS/GCP resources?
"I use the official SDKs: `aws-sdk-go-v2` or `google-cloud-go`.
I organize commands by resource: `mycli ec2 list`.
I use `context` for timeouts—crucial for cloud APIs.
I use interfaces for the cloud client so I can mock AWS in tests, allowing me to verify my 'delete unused instances' logic safely."

#### Indepth
**Pagination**. Most Cloud APIs paginate results. A common bug is processing only the first page (e.g., first 50 buckets). The Go SDKs often proide "Paginators" (`NewListBucketsPaginator`). Always use them to interact with the collection, even if you think you'll only have 10 items.

---

### 645. How do you use Go to write Terraform providers?
"I use the **Terraform Plugin Framework**.
I define `Resources` (Create, Read, Update, Delete).
`func (r *resource) Create(ctx, req, resp)`.
My Go code maps the Terraform HCL state to API calls.
This allows me to manage my company's internal platform (which has an API but no official TF provider) via Terraform."

#### Indepth
**State Drift**. Your provider must implement `Read` correctly to detect drift. If someone manually changes a value in the UI, `terraform plan` should see the difference. The `Read` function essentially maps the *External API Response* back to the *Internal Terraform Schema*.

---

### 646. How do you build a dynamic inventory script in Go for Ansible?
"Ansible expects a JSON output.
I write a Go CLI that queries my Cloud Source (AWS/Consul).
I marshal the result to the specific JSON format: `{ "_meta": { "hostvars": ... } }`.
I compile it and point `ansible -i my-go-inventory` to the binary."

#### Indepth
**Dynamic Groups**. Your Go tool can contain logic: "If name starts with 'db-', add to 'databases' group". This centralizes grouping logic in code rather than in static INI files, making your Ansible playbooks automatically adapt to new instances without manual edits.

---

### 647. How do you parse and generate YAML in Go?
"I use `gopkg.in/yaml.v3`.
It supports comments and preserving order (unlike JSON).
`yaml.Unmarshal(data, &configStruct)`.
For K8s manifests, I use `sigs.k8s.io/yaml` which handles the specific JSON/YAML duality of Kubernetes.
I check for `KnownFields: true` to catch typos in config files."

#### Indepth
Use `yaml:",omitempty"` sparingly in Infrastructure-as-Code. If a user sets `replicas: 0`, and you use `omitempty`, the YAML generator omits the field, and K8s might default it to `1`. Explicitly serializing zero-values is often safer for declarative configs.

---

### 648. How do you interact with Docker API in Go?
"I use the official `docker/docker/client`.
`cli, _ := client.NewClientWithOpts(client.FromEnv)`.
`cli.ContainerList(ctx, options)`.
I can build custom sidecars that listen to Docker events (container died) and trigger alerts or cleanup tasks."

#### Indepth
**Stream Handling**. `ContainerAttach` or `ImagePull` returns a stream. You must handle this stream correctly (it often multiplexes Stdout/Stderr/Stdin). Failing to read the stream can cause the Docker daemon to block (backpressure) and hang the operation. use `stdcopy.StdCopy`.

---

### 649. How do you manage Kubernetes CRDs in Go?
"I define the struct with JSON tags.
`type MyResource struct { Spec MySpec ... }`.
I use `controller-gen` to generate the YAML CRD manifest from the Go struct tags.
In my code, I use the dynamic client (`client.Client`) to List/Watch these custom resources just like native Pods."

#### Indepth
**Validation Schemas**. CRDs allow you to define OpenAPI v3 schemas. This enforces types at the API server level (e.g., "replicas must be > 0"). Always define these schemas rigourously so your Go controller doesn't crash trying to parse invalid JSON from a user's CR.

---

### 650. How do you write Go code to scale deployments in K8s?
"I use the client-go `Scale` subresource.
`clientset.AppsV1().Deployments(ns).UpdateScale(...)`.
Or I calculate the desired replica count based on custom metrics (e.g., RabbitMQ queue depth) and patch the Deployment. This is how KEDA works."

#### Indepth
**Metrics Server**. To use the Horizontal Pod Autoscaler (HPA) with your custom metric, you must implement the `External Metrics API`. You write a Go adapter that translates "Queue Depth" into a format the K8s Metrics Server understands, allowing standard HPA resources to scale your app.

---

### 651. How do you tail logs from containers using Go?
"I use `cli.ContainerLogs(ctx, containerID, options)`.
It returns a multiplexed stream (stdout + stderr headers).
I must use `stdcopy.StdCopy(dstOut, dstErr, src)` to demultiplex the binary stream into readable text lines."

#### Indepth
**Follow Mode**. If you use `Follow: true`, the connection stays open. You must handle network interruptions. A robust log tailer loop monitors the error channel and *re-connects* if the stream breaks, possibly using the `Since` timestamp to resume from the last received log line.

---

### 652. How do you manage service discovery in Go apps?
"I avoid hardcoding IPs.
In K8s, I use DNS: `http://my-service.default.svc`.
For external service discovery (Consul), I use the Consul API client to watch for healthy nodes and update my internal load balancer dynamically."

#### Indepth
**Client-Side Load Balancing**. Instead of a VIP (Virtual IP) which is a bottleneck, the Go client knows all 50 backend IPs. It picks one directly. gRPC has this built-in. It reduces latency (1 less hop) but requires the client to be "smart" and aware of the topology.

---

### 653. How do you build a Kubernetes admission controller in Go?
"It’s a webhook (HTTP Server).
K8s POSTs the Pod definition to me *before* creating it.
I decode the `AdmissionReview`.
I inspect the Pod ('Does it have root privileges?').
I return `Allowed: false` or a JSON Patch to mutate it (e.g., inject a sidecar)."

#### Indepth
**Cert Management**. Webhooks require HTTPS. Managing these internal certs is painful. Use `cert-manager` with a `CAInjector` to automatically inject the CA bundle into your `ValidatingWebhookConfiguration`. This prevents the "x509: certificate signed by unknown authority" error during cluster upgrades.

---

### 654. How do you build a metrics exporter for Prometheus in Go?
"I write a 'Collector'.
`func (c *MyCollector) Collect(ch chan<- prometheus.Metric)`.
On every scrape, I query my target (e.g., a hardware device), convert values to Metrics, and send to channel.
Then I run `http.Handle("/metrics", promhttp.Handler())`."

#### Indepth
**Describe vs Collect**. Implementing `Describe` is optional but good practice. It sends metric metadata (Help string, Type) to Prometheus. `Collect` is the hot path. Ensure `Collect` doesn't block for too long, or the scraper will timeout. If data gathering is slow, do it asynchronously and return the last cached value in `Collect`.

---

### 655. How do you set up health checks for a Go microservice?
"I follow the K8s pattern.
`/healthz` (Liveness): Returns 200 if loop is running.
`/ready` (Readiness): Returns 200 if DB is reachable.
I use a library like `hellofresh/health-go` to check DB/Redis connectivity and aggregate the result."

#### Indepth
**Cascading Failures**. Be careful with Readiness checks. If your DB is slow, and all 100 pods return "Not Ready", K8s kills traffic to ALL of them. Now the DB recovers, but your pods are effectively offline or flapping. Sometimes returning "Healthy" even if the DB is down (Degraded Mode) is safer to keep the UI accessible.

---

### 656. How do you build a custom load balancer in Go?
"I use `httputil.ReverseProxy` with a custom `Director`.
I maintain a list of backends `[]*url.URL`.
The Director function picks the next backend (Round Robin) and updates `req.URL.Host`.
I handle errors: if the backend is down, I retry on another."

#### Indepth
**Active Health Checks**. Passive check (wait for error) is slow. A real LB runs a background loop `HEAD /health` on backends. If one fails, it removes it from the rotation *before* a user request hits it. The `Director` needs read-access to this dynamic "Healthy Backends" list.

---

### 657. How do you implement graceful shutdown with Kubernetes SIGTERM?
"I listen for signals.
`c := make(chan os.Signal, 1); signal.Notify(c, syscall.SIGTERM)`.
Block: `<-c`.
Start shutdown:
1.  Set `readiness` to 503.
2.  `server.Shutdown(ctx)` (waits for active requests).
3.  Wait for background jobs.
This ensures K8s stops sending traffic and existing requests complete before the pod dies."

#### Indepth
**PreStop Hook**. K8s updates endpoints *async*. Even after SIGTERM, network rules might route traffic to you for a few seconds. It's common to add a `preStop` hook: `sleep 5`. This ensures the pod stays alive long enough for K8s to propagate the "Remove me from Endpoints" command to all kube-proxies.

---

### 658. How do you use Go with Envoy/Consul service mesh?
"I usually don't interact with them directly; they are sidecars.
However, I can use the **xDS protocol** (Go control plane) to configure Envoy dynamically.
This allows me to tell Envoy 'Route 50% traffic to v2' programmatically."

#### Indepth
**Sidecar vs Library**. With gRPC, you can skip Envoy and stick the mesh logic *inside* the Go binary (proxyless service mesh). This reduces latency (no localhost hop) and complexity (no sidecar container), but requires all apps to be written in languages with rich SDKs (Go/Java/C++).

---

### 659. How do you configure Go apps for 12-factor principles?
"**Env Vars** via `os.Getenv` or `kelseyhightower/envconfig`.
**Logs** to Stdout (JSON).
**Backing Services** (DB) attached via connection string URL.
**Stateless**: No local files.
**Port Binding**: Listen on `$PORT`.
I stick to these rules so my app runs anywhere (Docker, Heroku, K8s) without changes."

#### Indepth
**Config Separation**. strict separation means the same Docker image is used for Dev, Staging, and Prod. The *only* difference is the Environment Variables. If you are baking `config.prod.json` into the image, you are violating 12-Factor and making rollbacks harder.

---

### 660. How do you use Go for cloud automation scripts?
"I replace Bash scripts with Go.
Benefit: **Parallelism**.
'Restart 100 VMs':
Bash: Loop 1..100 (slow).
Go: `for vm := range vms { go restart(vm) }` (fast).
I compile it to a static binary for CI/CD pipelines."

#### Indepth
**Cross Combination**. `GOOS=linux GOARCH=amd64 go build`. You can build the maintenance script on a Mac and ship it to a Linux server. This portability is why Go is the language of cloud infrastructure (Docker, K8s, Terraform, Vault are all Go).
