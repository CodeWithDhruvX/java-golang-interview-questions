## 🚢 Docker Real-World Debugging & Troubleshooting (Questions 401-410)

### Question 401: How do you debug a stuck container that won’t exit?

**Answer:**
1.  **Check Status:** `docker ps`. Is it `Up`, `Paused`, or `Exited`?
2.  **Force Stop:** `docker kill <id>` (sends `SIGKILL`).
3.  **Zombie:** If `kill` fails, it might be a zombie process in the host kernel. Check host `dmesg` for IO errors. You might need to restart the Docker daemon or the Host.

---

### Question 402: How do you inspect a container’s open ports?

**Answer:**
1.  `docker port <container_id>`: Shows Host->Container mapping.
2.  `docker exec <id> netstat -tuln`: Shows ports listening inside.
3.  `docker inspect <id>`: Look for `NetworkSettings.Ports`.

---

### Question 403: What is the impact of high CPU usage inside a container on the host?

**Answer:**
Since containers share the kernel, one container using 100% CPU can starve other containers and host processes (if no limits set).
**Mitigation:** Always set `--cpu-shares` or `--cpus`.

---

### Question 404: How do you detect memory leaks inside a container?

**Answer:**
1.  **Monitor:** Watch `docker stats`. If MEM USAGE grows indefinitely.
2.  **OOM:** Check `docker inspect` for `OOMKilled: true`.
3.  **Profile:** Use `pprof` (Go) or Heap Dumps (Java) inside the container.

---

### Question 405: How do you troubleshoot network latency between Docker containers?

**Answer:**
1.  **Ping:** Check basic connectivity.
2.  **Iperf:** Measure bandwidth (`iperf -s` on A, `iperf -c A` on B).
3.  **Trace:** `traceroute` (to see if routing hops are correct).
4.  **MTU:** Check for MTU mismatch between Host/Docker bridge/Overlay.

---

### Question 406: How do you debug a container that exits immediately?

**Answer:**
1.  **Logs:** `docker logs <id>`. (Most likely app crash).
2.  **Foreground:** `docker run -it ... /bin/sh`. Run the entrypoint manually.
3.  **Exit Code:** `docker inspect` (Code 0 = Script finished, Code 1 = Error, Code 127 = Cmd not found).

---

### Question 407: How can you identify a misconfigured Dockerfile layer?

**Answer:**
1.  **Build History:** `docker history <image>`. Look for huge jumps in size.
2.  **Dive:** Use the `dive` tool to inspect exactly which files were added in that layer.

---

### Question 408: What tools can be used to trace syscalls from within a container?

**Answer:**
Use `strace`.
Requires `--cap-add=SYS_PTRACE`.
`docker exec -it --user root --cap-add=SYS_PTRACE <id> strace -p 1`

---

### Question 409: What does the `docker events` command show and how is it used?

**Answer:**
Streams real-time events from the daemon (create, die, oom, attach).
**Usage:** Debugging restart loops or triggering automations (like log shipping on start).

---

### Question 410: How do you inspect dead containers?

**Answer:**
Containers in `Dead` status are partially removed but failed to cleanup.
1.  Check Daemon logs (`journalctl -u docker`).
2.  Usually indicates a "Device or resource busy" error (filesystem mount didn't detach).
3.  Manual fix: Remove `/var/lib/docker/containers/<id>`.

---

## 🧱 Linux Kernel + Docker Deep Dive (Questions 411-420)

### Question 411: How do containers differ from chroot jails?

**Answer:**
- **Chroot:** Only isolates the File System root. Process can still see all other processes/network.
- **Containers:** Use Chroot (rootfs) + Namespaces (process/net isolation) + Cgroups (limits).

---

### Question 412: How does Docker leverage union filesystems like OverlayFS?

**Answer:**
Starts with empty `merged` directory.
Mounts:
- `lowerdir`: Image layers (RO).
- `upperdir`: Container layer (RW).
- `workdir`: Internal atomic work.
Result: The container sees a unified FS.

---

### Question 413: How does the kernel handle multiple PID 1s in containers?

**Answer:**
Kernel supports **PID Namespaces**.
The kernel sees the process as PID 24567.
The container namespace sees it as PID 1.
The kernel maintains this mapping.

---

### Question 414: How does a container's network stack differ from the host?

**Answer:**
It has its own slice of the Network Stack:
- Own Interface (`eth0`).
- Own Routing Table.
- Own Iptables chains.
- Own Loopback (`lo`).
Connected via VETH pair to Host Bridge.

---

### Question 415: How do Linux namespaces isolate containers at the kernel level?

**Answer:**
Kernel data structures (like process list) are tagged with a Namespace ID.
When a process asks "List PIDs", Kernel filters the list: "Show only PIDs matching this Namespace ID".

---

### Question 416: What are control group (cgroup v2) differences from v1?

**Answer:**
- **v1:** Hierarchies were per-controller (cpu tree distinct from memory tree). Complex.
- **v2:** A single unified hierarchy. Safer design, better support for rootless containers.

---

### Question 417: How does Docker limit disk IOPS using cgroups?

**Answer:**
`--device-read-iops /dev/sda:1000`.
Docker instructs the BlkIO controller to throttle requests from that cgroup to the specific block device.

---

### Question 418: How does a container know its resource constraints?

**Answer:**
It usually **doesn't** (Standard tools `free`, `top` show Host Limits).
**LXCFS:** A fuse filesystem that overlays `/proc/meminfo` to show Container Limits instead of Host Limits.

---

### Question 419: What happens when a container reaches its memory limit?

**Answer:**
1.  Kernel attempts to reclaim memory (Cache eviction).
2.  If fails, OOM Killer wakes up.
3.  OOM Killer identifies the process with highest score (usually the app) and `SIGKILLs` it.

---

### Question 420: How does Docker use capabilities like `CAP_NET_BIND_SERVICE`?

**Answer:**
This specific capability allows a non-root process to bind to ports < 1024.
Docker allows this by default so web servers can listen on port 80.

---

## 🔧 Advanced Docker CLI & API Usage (Questions 421-430)

### Question 421: How do you use the Docker REST API for automation?

**Answer:**
Enable TCP socket or use Unix socket.
`GET /images/json`.
Useful for building custom dashboards or CI integrations.

---

### Question 422: How do you create a container without starting it?

**Answer:**
`docker create --name my-app nginx`.
Status: `Created`.
Start later: `docker start my-app`.

---

### Question 423: What is the use of `docker wait`?

**Answer:**
Blocks until a container stops, then prints its exit code.
**Use Case:** Scripting.
`docker start app; code=$(docker wait app); echo "Finished with $code"`

---

### Question 424: What does `docker pause` and `unpause` do?

**Answer:**
Uses the Cgroup Freezer.
The process is suspended in memory (CPU usage drops to 0). It is NOT killed.
Useful for snapshotting.

---

### Question 425: How do you attach to a background-running container?

**Answer:**
`docker attach <id>`.
Connects your Stdin/Stdout to the main process.

---

### Question 426: How do you use `docker update` to change resource limits?

**Answer:**
Updates container configuration runtime.
`docker update --cpus 2 my-running-container`.
*(Does not require restart)*.

---

### Question 427: How do you copy files from a running container?

**Answer:**
`docker cp <id>:/path/file ./local`.
Note: The tar stream is generated by the daemon, checking the filesystem.

---

### Question 428: How do you see which containers are using a specific volume?

**Answer:**
`docker ps --filter volume=my-vol`.
Lists containers correctly mounting that volume.

---

### Question 429: What’s the difference between `docker top` and `docker stats`?

**Answer:**
- **Stats:** Resource usage metrics (CPU % / Mem).
- **Top:** List of processes running inside (PID, User, Time).

---

### Question 430: How do you tail logs in JSON format?

**Answer:**
Docker logs command doesn't output JSON structure, it outputs the raw log line.
If logging driver is json-file, you can inspect the file directly on host:
`/var/lib/docker/containers/<id>/<id>-json.log`.

---

## 📦 Docker in Build Pipelines (CI/CD) (Questions 431-440)

### Question 431: How do you use Docker multi-stage builds in CI/CD pipelines?

**Answer:**
Simply use a Dockerfile with `FROM ... AS ...`.
CI run `docker build`.
Benefit: The final image pushed to registry is small, keeping CD fast.

---

### Question 432: What are benefits of Docker caching in CI runners?

**Answer:**
Speeds up build.
Runners usually start clean. Use `--cache-from` (remote registry) or local persistence volume (if self-hosted) to reuse layers.

---

### Question 433: How do you reduce build times using GitHub Actions + Docker?

**Answer:**
Use `docker/setup-buildx-action` and `cache-from: type=gha`.
It saves cache blobs to GitHub's cache storage.

---

### Question 434: How do you manage secrets in CI/CD Docker builds?

**Answer:**
Use `--secret id=mytoken,src=env` in buildx.
In CI, populate the env var from repo secrets.

---

### Question 435: What happens if your CI pipeline gets interrupted during `docker build`?

**Answer:**
The intermediate layers already built successfully are cached locally (on that agent).
Failed/Partial layer is discarded.

---

### Question 436: How do you create reproducible builds across CI agents?

**Answer:**
Pin Everything.
- Base Image SHA.
- System Packages versions.
- App Dependency versions.
- Build Tool versions.

---

### Question 437: What’s the difference between `docker build` and `docker-compose build`?

**Answer:**
- **docker build:** Builds a single Dockerfile.
- **docker-compose build:** Builds *all* services defined in compose.yml that have a `build:` key.

---

### Question 438: How do you use Docker-in-Docker in CI?

**Answer:**
Service definition (GitLab/GitHub):
```yaml
services:
  docker:
    image: docker:dind
```
The job connects to this service (via TCP or Socket) to run docker commands.

---

### Question 439: How do you push to a registry from a CI job securely?

**Answer:**
1.  `docker login -u $USER -p $TOKEN registry.com`
2.  `docker push ...`
3.  `docker logout` (Cleanup).

---

### Question 440: How do you use cache mounts in `RUN` commands?

**Answer:**
```dockerfile
RUN --mount=type=cache,target=/root/.m2 mvn package
```
In CI, this cache is effective only if the cache mount directory is persisted between jobs (Docker BuildKit handles this if configured).

---

## 🧩 Docker Compose in Depth (Questions 441-450)

### Question 441: What is the difference between `depends_on` and true service readiness?

**Answer:**
`depends_on` only waits for the container **state** to be "Running".
It does NOT wait for the app (e.g., DB listening on port).
Use `condition: service_healthy` for true readiness.

---

### Question 442: How do you configure networks in Docker Compose?

**Answer:**
Top-level `networks` key.
```yaml
networks:
  frontend:
  backend:
    internal: true
```
Assign servies via `networks: [backend]`.

---

### Question 443: What are named volumes and how are they reused across services?

**Answer:**
Top-level `volumes` key.
```yaml
volumes:
  db_data:
```
Multiple services can mount `db_data`.

---

### Question 444: How do you share environment variables across multiple services?

**Answer:**
1.  **Duplicate:** Copy-paste in YAML (Bad).
2.  **YAML Anchors:** Define `&common-env` and `<<: *common-env`.
3.  **Extends:** Use `extends` keyword (Compose v2).

---

### Question 445: How do you define build contexts per service in Compose?

**Answer:**
```yaml
services:
  web:
    build:
      context: ./web
      dockerfile: Dockerfile.prod
```

---

### Question 446: How do you debug Compose service dependencies?

**Answer:**
Remove `detach` (`-d`).
`docker-compose up` (Foreground).
Watch the order of logs. See if App crashes before DB is ready.

---

### Question 447: How do you override Compose files for production?

**Answer:**
`docker-compose -f docker-compose.yml -f docker-compose.prod.yml up`
Prod file sets restarting policies, resource limits, and specific image tags.

---

### Question 448: What is the `.env` file's role in Compose projects?

**Answer:**
Variable substitution for the YAML pre-processor.
`DEBUG=${DEBUG_MODE}`.

---

### Question 449: How do you scale a service in Compose?

**Answer:**
`docker-compose up --scale worker=5`.
(Only works if no host port conflicts).

---

### Question 450: How do you set container labels from Compose?

**Answer:**
```yaml
services:
  web:
    labels:
      - "traefik.enable=true"
```
Useful for reverse proxies and monitoring discovery.
