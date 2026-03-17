# Priority Category 4: Containerization & Kubernetes

### 1. How do you containerize Go applications efficiently?
**Your Response:**
"Efficiency in Docker is about size and security.
I use **Multi-Stage Builds**. 
In Step 1, I use `golang:alpine` to build the binary using `go build -ldflags="-s -w"` (to strip debug symbols).
In Step 2, I copy *only* that binary into a `scratch` or `distroless` image.
This takes an image from 800MB down to maybe 15MB. It's faster to pull, faster to start, and since there's no shell in the final image, it's significantly harder for a hacker to exploit."

### 2. Explain multi-stage Docker builds for Go services.
**Your Response:**
"Multi-stage builds allow us to separate our build environment from our runtime environment.
In the first stage, we have all our tools: the Go compiler, git for fetching modules, etc.
Once the static binary is created, we discard that whole environment. We start a fresh, empty image and just `COPY --from=builder` the binary. 
The result is a production container that contains *nothing* except the code it needs to run. It's the 'Platinum Standard' for Go deployments."

### 3. How do you deploy Go microservices on Kubernetes?
**Your Response:**
"I use a standard deployment manifest that defines:
1. **Resources**: Specific CPU/Memory requests and limits so K8s can schedule precisely.
2. **Probes**: Liveness to detect crashes, and Readiness to ensure the app is ready for traffic.
3. **Graceful Termination**: Setting `terminationGracePeriodSeconds` to give my Go app enough time to finish in-flight requests during a rollout. 
I bundle these into a **Helm Chart** to make the deployment repeatable and version-controlled."

### 4. What are the best practices for Go applications in containers?
**Your Response:**
- "Always set `CGO_ENABLED=0` to ensure the binary is statically linked and doesn't need external libraries.
- Use **Environment Variables** for configuration so the same container can run in any environment.
- Implement **SIGTERM handling** in Go to allow for graceful shutdown.
- Never run as **root**; use a non-privileged user inside the Dockerfile to follow the Principle of Least Privilege."

### 5. How do you handle configuration management in Kubernetes?
**Your Response:**
"I use a combination of **ConfigMaps** and **Secrets**.
- **ConfigMaps**: For things like log levels, feature flags, or service URLs.
- **Secrets**: For database passwords or API keys.
I inject these as Environment Variables into the Go service. For a truly production-grade system, I also implement a watcher in Go that detects if a ConfigMap changes and reloads the configuration without needing to restart the whole pod."
