# 📦 **Container Storage & Volumes Advanced (181–190)**

---

### 181. What is the difference between `tmpfs` and volume mounts?
"**tmpfs** stores data in **host RAM** — it's fast (no disk I/O) but ephemeral (gone when container stops) and shared with the host's memory.

`docker run --tmpfs /tmp:rw,size=100m myapp`

**Volume mounts** store data on **disk** (managed by Docker or bind-mounted), persisting across container restarts.

Use `tmpfs` for: sensitive data that shouldn't touch disk (session caches, secrets during processing), logs that don't need persistence, and high-speed temporary scratch space. Avoid for: anything that needs to survive container restarts."

#### In-depth
`tmpfs` mounts bypass the copy-on-write filesystem entirely — reads and writes go directly to memory. This makes them ideal for high-frequency write workloads (Redis-style caches, socket files, session stores) that would cause excessive copy-on-write overhead on the overlay filesystem. Many security-focused deployments mount `/tmp` as tmpfs to prevent sensitive temp file residue from persisting on disk.

---

### 182. How do you persist database state in Docker?
"Always use a **named volume** for database data directories:

```yaml
services:
  postgres:
    image: postgres:15
    volumes:
      - pgdata:/var/lib/postgresql/data
    environment:
      POSTGRES_PASSWORD: secret

volumes:
  pgdata:
```

The named volume persists when you `docker compose down` (but not `docker compose down -v`). Run `docker volume ls` to see it. Backup: `docker run --volumes-from postgres_container alpine tar czf - /var/lib/postgresql/data > backup.tar.gz`."

#### In-depth
Database persistence in Docker has a common gotcha: the postgres image initializes data only if `/var/lib/postgresql/data` is empty. If you volume-mount a pre-existing directory with other files, postgres may fail to initialize. Always use named Docker volumes for databases — Docker creates the volume directory clean, and the DB initializes correctly. For production databases, I strongly recommend managed cloud databases (RDS, Cloud SQL) over containerized databases — they include automated backups, replication, and failover built-in.

---

### 183. How do you backup and restore Docker volumes?
"**Backup a volume**:
```bash
docker run --rm \
  -v myvolume:/source:ro \
  -v $(pwd):/backup \
  alpine tar czf /backup/myvolume_backup.tar.gz -C /source .
```

**Restore a volume**:
```bash
docker volume create myvolume
docker run --rm \
  -v myvolume:/target \
  -v $(pwd):/backup \
  alpine tar xzf /backup/myvolume_backup.tar.gz -C /target
```

The pattern: spin up a temporary alpine container with the volume mounted, use tar to backup/restore. The running database container doesn't need to stop for read-consistent backups (use `pg_dump` for postgres instead)."

#### In-depth
For production database backups, application-aware backup methods (pg_dump, mysqldump, mongodump) are far superior to filesystem-level tar backups. Application-aware backups: guarantee consistency (no partial transactions), support point-in-time recovery, are smaller (compressed, no docker metadata), and are portable across Docker versions. Use filesystem backups only for stateful non-database volumes where no application-level backup tool exists.

---

### 184. How do you use NFS with Docker volumes?
"NFS (Network File System) volumes enable shared storage across multiple Docker hosts:

```bash
docker volume create \
  --driver local \
  --opt type=nfs \
  --opt o=addr=nfs-server.example.com,rw \
  --opt device=:/exported/path \
  nfs-volume

docker run -v nfs-volume:/data myapp
```

In Compose:
```yaml
volumes:
  nfs-volume:
    driver: local
    driver_opts:
      type: nfs
      o: addr=nfs-server.example.com,rw,nolock
      device: ':/exported/path'
```"

#### In-depth
NFS volumes in Docker have performance limitations: NFS adds network latency to every file operation. For database workloads, NFS is generally unsuitable — the latency and locking semantics cause poor performance. For shared file storage (user uploads, shared assets), NFS or cloud-native equivalents (Amazon EFS, Azure Files, Google Filestore) work well. In Kubernetes, these translate to PersistentVolumes with NFS or EFS CSI drivers.

---

### 185. What are CSI (Container Storage Interface) drivers?
"CSI is an **industry-standard specification** for storage system plugins in container orchestration platforms (Kubernetes, Mesos, Nomad).

CSI drivers allow storage vendors (AWS EFS, GCP Filestore, NetApp, Rook/Ceph) to implement a plugin that works with any CSI-compatible orchestrator — without modifying the orchestrator itself.

In Kubernetes: CSI drivers appear as StorageClasses. You create PersistentVolumeClaims (PVCs) that reference a StorageClass, and Kubernetes provisions the appropriate storage automatically.

In Docker/Swarm: `docker plugin install` installs volume plugins — the precursor to standardized CSI."

#### In-depth
CSI drivers implement three operations: **CreateVolume** (provision a new PV), **AttachVolume** (attach an EBS volume to an EC2 node), **MountVolume** (mount the attached volume into the container). The separation of attach and mount enables: dynamic provisioning (create the volume when the pod is scheduled), multi-attach (share EFS across nodes), and topology awareness (provision volumes in the same AZ as nodes). AWS EFS CSI driver (`efs.csi.aws.com`) is the most common for shared persistent storage in EKS.

---

### 186. How do you share volumes across containers?
"Multiple containers can mount the same named volume:

```yaml
services:
  writer:
    image: writer-app
    volumes:
      - shared:/data

  reader:
    image: reader-app
    volumes:
      - shared:/data:ro  # read-only for the reader

volumes:
  shared:
```

Or at runtime: `docker run -v shared-vol:/data container1 & docker run -v shared-vol:/data:ro container2`.

The `--volumes-from` flag (legacy): `docker run --volumes-from container1 container2` — mounts all volumes from container1 into container2."

#### In-depth
Sharing writable volumes between containers without file locking can lead to data corruption. Two services writing to the same file simultaneously without coordination will interleave their writes. Solutions: use one writer with one or more readers (common for log files), implement file locking in the application, or use a message queue/database instead of shared volumes for coordination. The read-only `:ro` flag prevents unintended writes from the reader service.

---

### 187. How do you inspect data inside a volume?
"Mount the volume into a temporary container and explore:

```bash
# Interactive inspection
docker run --rm -it -v myvolume:/inspect alpine sh

# List files non-interactively
docker run --rm -v myvolume:/inspect alpine ls -la /inspect

# Copy data out
docker run --rm -v myvolume:/inspect -v $(pwd):/output alpine cp -r /inspect /output/volume-contents
```

For Postgres: `docker exec postgres psql -U user -c '\l'` (use the running container directly).

`docker volume inspect myvolume` gives metadata: name, driver, mount point on host (`/var/lib/docker/volumes/myvolume/_data`)."

#### In-depth
Volumes live at `/var/lib/docker/volumes/<name>/_data` on the host. On Linux, root can access these files directly. On macOS/Windows with Docker Desktop, they live inside the Linux VM — not directly accessible from the macOS filesystem. The temporary-container approach (`docker run --rm -v vol:/inspect alpine`) works cross-platform. For debugging volume contents in production, a short-lived debug container avoids touching production service containers.

---

### 188. How does Docker handle read-only volumes?
"Append `:ro` to the volume mount: `docker run -v myconfig:/app/config:ro myapp`.

Inside the container, the file system at `/app/config` is mounted read-only. Any write attempt returns `EROFS (Read-only file system)`.

For the full container filesystem: `docker run --read-only myapp`. This makes the entire root filesystem read-only. Apps that need writable temp space: `docker run --read-only --tmpfs /tmp myapp`.

In Compose: `volumes: - ./config:/app/config:ro`."

#### In-depth
`--read-only` is one of the most important security controls — it prevents an attacker who gains container code execution from modifying the container filesystem (installing backdoors, editing scripts). Combined with `--tmpfs /tmp` and `--tmpfs /run`, the container has writable scratch space where needed without exposing the application filesystem to modification. Kubernetes equivalent: `readOnlyRootFilesystem: true` in the container security context.

---

### 189. What's the impact of file permission issues in volume mounts?
"File permission issues are among the most common Docker pain points, especially with bind mounts.

The problem: the user running the process inside the container (UID 1000) may differ from the host user (UID 1001) or the file owner on the host (root/UID 0). Result: `Permission denied`.

Solutions:
1. **Match UIDs**: `docker run --user $(id -u):$(id -g) myapp`
2. **chown in entrypoint**: `chown -R appuser:appgroup /data && exec gosu appuser app`
3. **Named volumes**: Docker sets ownership to the container's default user automatically
4. **`fixuid`**: a small binary that remaps the container UID to match the host user dynamically"

#### In-depth
The UID mismatch is particularly acute on Linux. On macOS/Windows, Docker Desktop runs a Linux VM and translates host file access — permissions appear as the host user regardless of container UIDs. On Linux (native Docker), UIDs are real — a mismatch causes real permission errors. Production containers running as non-root with volume mounts need careful UID planning. The cleanest solution: use named volumes with Docker managing the permissions, and avoid bind mounts in production entirely.

---

### 190. How do you manage volume lifecycles?
"Volume lifecycle management:

**Create**: `docker volume create myvolume` or implicitly in `docker run -v myvolume:/data`

**Inspect**: `docker volume inspect myvolume` — shows driver, mount path, labels

**List**: `docker volume ls -f dangling=true` — shows unused volumes

**Remove**: `docker volume rm myvolume` — fails if attached to a container

**Prune**: `docker volume prune` — removes all unused volumes (those not mounted by any container)

**Automatic cleanup**: `docker run --rm -v tempvol:/data` — the `--rm` flag removes anonymous volumes created for that run."

#### In-depth
Volume pruning is **irreversible** — removed volume data is gone. Unlike image pruning (images can be re-pulled), volume data cannot be recovered without a backup. Set up lifecycle policies: in production, volumes should either have explicit backup jobs or be recreatable from a database restore. Use labels (`docker volume create --label env=prod myvolume`) to categorize volumes and avoid accidentally pruning production data with a broad `docker volume prune` command. Always use `--filter label=env=temp` to scope pruning.

---
