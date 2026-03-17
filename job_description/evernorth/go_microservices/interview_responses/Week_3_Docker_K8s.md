# Week 3: Containerization & Kubernetes

### 🔹 Topic: Docker for Go Applications

**Interviewer:** "What is your approach to Dockerizing Go microservices?"

**Your Response:**
"I always use **multi-stage builds**. In the first stage (the build stage), I use a heavy image like `golang:alpine` to compile the binary. Then, in the second stage (the final image), I copy only that compiled binary and necessary config files into an extremely small image like `alpine` or even `scratch`. 

This keeps the final image size down to just a few megabytes, which makes deployments faster and reduces the security 'attack surface' since there are no unnecessary shells or tools in the production container."

### 🔹 Interview Focus: Docker & Kubernetes

**1. How do you minimize Go container image size?**
**Your Response:** "Aside from multi-stage builds, I make sure to turn off CGO during the build (using `CGO_ENABLED=0`) to create a fully static binary. I also use the `-ldflags="-s -w"` flag to strip debug information and symbol tables from the binary."

**2. Explain health checks and readiness probes.**
**Your Response:** "In Kubernetes, a **Liveness Probe** tells K8s if the container is alive; if it fails, K8s restarts it. A **Readiness Probe** tells K8s if the container is ready to serve traffic. I usually implement a `/health` endpoint that checks DB connectivity and internal state before returning a 200 OK."

**3. How do you handle secrets in Kubernetes?**
**Your Response:** "I never hardcode them. I use **Kubernetes Secrets** and inject them into the Go service as environment variables or mounted files. For more sensitive production environments, I prefer using a dedicated manager like HashiCorp Vault or AWS Secrets Manager."

**4. What is Horizontal Pod Autoscaling (HPA)?**
**Your Response:** "HPA automatically scales the number of pods in a deployment based on CPU or memory usage. If my Go service starts hitting 80% CPU because of high traffic, HPA will spin up more pods. Once traffic drops, it scales back down to save costs."

### 🔹 Week 3 Practice Problems: Spoken Walkthroughs

**1. Optimized Docker multi-stage build:**
"I'd start with `FROM golang AS builder`, set `ENV CGO_ENABLED=0`, and run `go build`. Then `FROM alpine:latest`, copy the binary from `builder`, and use `ENTRYPOINT` to run it. Simple, secure, and tiny."

**2. Kubernetes deployment with HPA:**
"I'd write a `deployment.yaml` defining my containers and a `hpa.yaml` targeting that deployment. I'd set a min/max pod count and a CPU utilization target. I'd also make sure my application's resource requests/limits are properly defined in the deployment."

**3. Helm chart for microservices:**
"I'd create a standard Helm structure with `values.yaml` for environment-specific configs. This allows me to deploy the same Go service to Dev, Staging, and Prod just by swapping out the values file, which is much cleaner than managing multiple manifests."
