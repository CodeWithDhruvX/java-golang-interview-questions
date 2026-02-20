# 🔒 **Docker Security & Compliance Advanced (121–130)**

---

### 121. What is a Docker security profile (seccomp)?
"**seccomp** (Secure Computing Mode) is a Linux kernel feature that **filters which system calls a process can make**.

By default, Docker applies a seccomp profile that blocks 44+ dangerous syscalls (like `kexec_load`, `ptrace`, `reboot`) while allowing the ~300+ needed for normal operations. This prevents a compromised container from calling kernel APIs it shouldn't need.

Docker's default profile is stored at `/etc/docker/seccomp.json`. You can apply a custom profile: `docker run --security-opt seccomp=myprofile.json myapp`. Or disable entirely (not recommended): `--security-opt seccomp=unconfined`."

#### In-depth
seccomp profiles are key for defense-in-depth. Even if an attacker exploits an app vulnerability, they can't use dangerous syscalls to escape. Custom profiles should follow the principle of least privilege: start with `unconfined`, use `strace` to determine which syscalls your app actually uses, then deny everything else. Tools like `oci-seccomp-bpf-hook` can automatically generate profiles by recording syscalls during testing.

---

### 122. What is AppArmor and how does it work with Docker?
"**AppArmor** (Application Armor) is a Linux Mandatory Access Control (MAC) security module that **confines programs to a set of resources** defined in a profile.

Each profile is a list of allowed files, capabilities, and network access. Docker applies the `docker-default` AppArmor profile to all containers automatically (on Ubuntu/Debian). It prevents containers from accessing certain host paths and operations.

Custom profile: `docker run --security-opt apparmor=my-profile myapp`. Disable: `--security-opt apparmor=unconfined`."

#### In-depth
AppArmor and seccomp are complementary: seccomp restricts which kernel syscalls are available; AppArmor restricts what files/resources those syscalls can operate on. Think seccomp as "you can't call `open()`" vs AppArmor as "you can call `open()` but only on these paths". Together they form a strong security boundary. AppArmor profiles are loaded by the kernel, making them resistant to container-level manipulation.

---

### 123. What is SELinux and how is it used with Docker?
"**SELinux** (Security-Enhanced Linux) is an alternative MAC system to AppArmor, primarily used on RHEL/CentOS/Fedora systems. It uses **labels** on files and processes to enforce access policies.

Docker containers get a SELinux label like `svirt_lxc_net_t`. It restricts what host filesystem paths containers can access. Volumes need the `:z` (shared) or `:Z` (private) flag to be relabeled for container access: `docker run -v /data:/app:Z myapp`.

SELinux in **enforcing** mode (not permissive) is required for meaningful security. Check: `getenforce`."

#### In-depth
The `:z` and `:Z` volume suffixes are critical on SELinux systems. Without them, bind-mounted directories retain their host SELinux context and containers (labeled differently) receive `Permission denied` on access. `:z` applies a shared label (multiple containers can access), `:Z` applies a private label (only this container can access). This trips up developers new to RHEL-based Docker deployments.

---

### 124. What is the purpose of capabilities in Docker?
"Linux **capabilities** divide root's omnipotent privileges into distinct units. Instead of all-or-nothing root, capabilities give fine-grained control.

Examples: `CAP_NET_BIND_SERVICE` (bind to ports <1024), `CAP_SYS_ADMIN` (broad admin access), `CAP_NET_ADMIN` (configure network interfaces), `CAP_SYS_PTRACE` (use ptrace).

By default, Docker grants containers a subset of capabilities. Drop all and add only necessary ones: `docker run --cap-drop=ALL --cap-add=NET_BIND_SERVICE myapp`. This ensures a container can bind to port 443 without having any other elevated privileges."

#### In-depth
The full list of Linux capabilities is ~40. Docker's default set (described in the Docker docs) includes commonly needed ones but excludes dangerous ones like `CAP_SYS_ADMIN`. `CAP_SYS_ADMIN` is so broad it's nicknamed 'the new root' — it includes 200+ different privilege operations. Never add it unless absolutely necessary, and if you do, audit what specific operations in `CAP_SYS_ADMIN` you actually need and look for a more specific capability.

---

### 125. How do you restrict Linux capabilities for a container?
"Two complementary approaches:

**Drop all, add back minimal**: `--cap-drop=ALL --cap-add=CAP_NET_BIND_SERVICE --cap-add=CAP_SETUID`

**Drop specific dangerous ones**: `--cap-drop=SYS_ADMIN --cap-drop=NET_ADMIN`

List capabilities: inspect with `docker exec container grep CapEff /proc/self/status` and decode with `capsh --decode=<hex>`.

In Compose: `cap_drop: [ALL]` + `cap_add: [NET_BIND_SERVICE]`. In Kubernetes: `securityContext.capabilities.drop: [ALL]` + `add: [NET_BIND_SERVICE]`."

#### In-depth
Getting capability requirements right takes testing. I typically run the app first with `--cap-add=ALL` to confirm it works, then systematically drop capabilities one by one to find the minimum set. After identifying needs, I document them in the Dockerfile/Compose as comments explaining WHY each capability is needed. This makes security reviews easier and prevents future contributors from blindly adding `--privileged`.

---

### 126. How do you enable Docker Content Trust (DCT)?
"Set the environment variable: `export DOCKER_CONTENT_TRUST=1`.

With DCT enabled: pulls only succeed for signed images, pushes automatically sign images, and unsigned images are rejected.

For CI: `export DOCKER_CONTENT_TRUST=1` in the pipeline env, provide a passphrase for the signing key via `DOCKER_CONTENT_TRUST_REPOSITORY_PASSPHRASE`. Keys are stored in `~/.docker/trust/`.

To sign an existing image: `docker trust sign myimage:v1.0`."

#### In-depth
DCT uses **Notary** (TUF implementation) by default. For production, consider **cosign** (from Sigstore) as a more modern alternative — it integrates with OIDC-based keyless signing in CI/CD (GitHub Actions, Google Cloud). Cosign signatures are stored as OCI artifacts in the registry alongside the image, making key management simpler than Notary's traditional key store approach.

---

### 127. What is the `--security-opt` flag used for?
"`--security-opt` applies security profiles and settings to a container. Options:

- `seccomp=profile.json` — apply a seccomp syscall filter profile
- `apparmor=my-profile` — apply an AppArmor MAC profile
- `label=type:svirt_sbox_t` — set a custom SELinux label
- `no-new-privileges` — prevent privilege escalation via setuid binaries
- `seccomp=unconfined` / `apparmor=unconfined` — disable the respective layer

`no-new-privileges` is the most universally useful: it prevents the container process from gaining more privileges than its parent, even if a setuid binary is present."

#### In-depth
`--security-opt=no-new-privileges` is critically important. Without it, if a setuid binary exists in the container (e.g., `ping`, `sudo`, `su`), a non-root user could potentially use it to gain root inside the container. With `no-new-privileges`, the kernel guarantees that `execve()` cannot result in elevated privileges — seccomp and AppArmor profiles cannot be bypassed via setuid. Enable it for all production containers.

---

### 128. How do you isolate sensitive data from logs in containers?
"Several layers of protection:

1. **Don't log sensitive data in the app** — use structured logging and explicitly control what fields are logged
2. **Redact in log pipeline** — use Fluentd/Logstash filters to redact patterns (credit cards, tokens, SSNs)
3. **Avoid environment variables for secrets** — they appear in `docker inspect` and `/proc/<pid>/environ`. Use file-based secrets (Docker secrets) instead
4. **Use ephemeral log retention** — configure `--log-opt max-size=10m max-file=3` to limit log history
5. **Audit log access** — restrict who can `docker logs` in production"

#### In-depth
Environment variables are particularly dangerous for secrets. They're visible in: `docker inspect container`, `docker exec env`, `/proc/<pid>/environ`, and can be dumped by any process inside the container. The Kubernetes secrets-as-env-vars problem is well-documented. File-based injection (secrets mounted as files, readable only by the running process) is significantly more secure — the OS controls access via normal file permissions.

---

### 129. How does Docker protect against container escape?
"Docker's defense-in-depth against container escape:

1. **Namespaces**: isolate PID, network, mount, UTS, IPC views
2. **cgroups**: limit resource usage, prevent resource-based attacks
3. **Capabilities**: drop unnecessary kernel privileges
4. **seccomp**: block dangerous syscalls
5. **AppArmor/SELinux**: MAC policies restricting file and resource access
6. **rootless Docker**: run daemon without root, entire privilege tree is shifted
7. **gVisor/Kata Containers**: hardware-level or additional kernel isolation

No single layer is sufficient — container escape is possible (CVEs exist), but each layer makes it significantly harder."

#### In-depth
Notable container escape CVEs: CVE-2019-5736 (runc file descriptor vulnerability, affected all Docker versions < 18.09.2), CVE-2020-15257 (containerd API exposure via abstract Unix socket accessible from container). Mitigation: keep Docker/containerd up to date, subscribe to Docker security advisories, run vulnerability scanners on infrastructure. The Kubernetes ecosystem uses OPA/Gatekeeper policies to enforce security standards at admission control level.

---

### 130. What are common CVEs related to Docker?
"Key historical Docker CVEs to know:

**CVE-2019-5736**: runc vulnerability allowing overwrite of the `runc` binary from inside a privileged container, leading to host command execution. Patched in runc 1.0-rc6+.

**CVE-2019-13139**: Docker build command injection via maliciously crafted repository URLs.

**CVE-2020-15257**: containerd exposure of an abstract Unix socket accessible from containers, allowing privilege escalation.

**CVE-2021-21284**: Docker rootless mode privilege escalation via `--userns-remap`.

**CVE-2022-0492**: Linux kernel cgroups escape affecting containers."

#### In-depth
The pattern across all container escape CVEs: they exploit the **shared kernel boundary** or **privileged APIs** accessible from containers. The fundamental limitation of Linux containers (as opposed to VMs) is that they share the host kernel — a kernel vulnerability is a container escape vulnerability. The defense: use a hardened kernel with `grsecurity`/`PAX` patches (enterprise), run workloads in gVisor/Kata for additional kernel-level isolation, and maintain aggressive patch cycles.

---
