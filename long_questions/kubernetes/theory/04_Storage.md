# 🟠 Storage: Volumes, PV, PVC & CSI

---

### 1. What is a Volume in Kubernetes?

"A Volume is a **storage abstraction** that allows containers in a Pod to access shared, persistent data.

Unlike Docker volumes that are scoped to a container, K8s Volumes are scoped to a **Pod**. Multiple containers in the same Pod can share a volume, and the data survives container restarts (but not pod deletion, unless using persistent volumes).

I use `emptyDir` for ephemeral scratch space between containers, and PersistentVolumes for anything that must outlive a pod."

#### In Depth
K8s supports many volume types: `emptyDir`, `hostPath`, `configMap`, `secret`, `persistentVolumeClaim`, `projected`, `nfs`, and many CSI-based volumes. The volume lifecycle is tied to the Pod lifecycle unless it's backed by a PV (PersistentVolume), which has its own independent lifecycle.

---

### 2. Difference between emptyDir and hostPath?

"`emptyDir` is created fresh when a Pod starts and is deleted when the Pod stops. It's stored on the **node's disk or in memory** (`medium: Memory`). Good for temporary data shared between containers.

`hostPath` mounts **a specific directory from the host node's filesystem** into the container. Persistent across pod restarts (as long as the pod stays on the same node).

I use `emptyDir` for log buffering between a main container and a log-shipper sidecar. I avoid `hostPath` in production — it creates tight coupling to specific nodes, breaks pod mobility, and creates security risks."

#### In Depth
`hostPath` volumes are a common source of security vulnerabilities — a misconfigured `hostPath: /` gives a container root access to the host filesystem. PSA (Pod Security Admission) `restricted` profile blocks hostPath usage. Use it only for DaemonSets that legitimately need host access (e.g., log collectors reading `/var/log`).

---

### 3. What is PersistentVolume (PV)?

"A PersistentVolume is a **cluster-level storage resource** that has been provisioned by an administrator or dynamically by a StorageClass.

It's completely decoupled from Pods — a PV has its own lifecycle (Created → Bound → Released → Available). It persists beyond the Pod that uses it.

Think of a PV like a physical disk drive that can be attached and detached from different machines (pods) over time."

#### In Depth
PVs have three access modes: `ReadWriteOnce` (RWO, one node can mount read-write), `ReadOnlyMany` (ROX, many nodes can mount read-only), `ReadWriteMany` (RWX, many nodes can mount read-write). Not all storage backends support all access modes — AWS EBS only supports RWO, NFS supports RWX. Check your storage backend's capabilities during architecture design.

---

### 4. What is PersistentVolumeClaim (PVC)?

"A PVC is a **user's request for storage**. It's the abstraction between the application and the actual storage implementation.

In a PVC, you specify the access mode, storage class, and capacity you need. K8s binds it to a matching PV.

This separation is powerful: developers write PVCs, platform teams provision PVs. The developer doesn't need to know about the underlying storage (EBS, NFS, Ceph, etc.)."

#### In Depth
PVC binding is based on: access modes match, requested capacity ≤ PV capacity, StorageClass matches (or is empty if using static provisioning), and any label selectors you define. Once bound, a PVC is exclusive to a PV — even if the PVC only used 1Gi of a 100Gi PV, no other PVC can use that PV. Size your PVs appropriately.

---

### 5. What is StorageClass?

"StorageClass defines a **template for dynamic PV provisioning**. When a PVC references a StorageClass, K8s automatically creates a PV on the backing storage system.

Parameters vary per provisioner: for AWS EBS, you'd specify `type: gp3`, `iopsPerGB`, etc. For GKE Persistent Disks, you'd set `type: pd-ssd`.

In production, I create multiple StorageClasses: `standard` for HDD-backed storage, `fast` for SSD, and `ultra-fast` for NVMe — letting teams choose based on their IOPS requirements."

#### In Depth
StorageClass has a `reclaimPolicy`: `Delete` (default — PV and underlying disk deleted when PVC is deleted) or `Retain` (PV and disk kept, must be manually cleaned up). For production data, I always use `Retain` to prevent accidental data deletion. The `volumeBindingMode: WaitForFirstConsumer` delays PV creation until a pod using the PVC is scheduled — critical for topology-aware volumes like zone-specific EBS.

---

### 6. What is dynamic volume provisioning?

"Dynamic provisioning means K8s **automatically creates a PV** when a PVC is created with a StorageClass reference.

Without dynamic provisioning, an admin must manually create PVs before users can claim them. That doesn't scale.

In all cloud environments (GKE, EKS, AKS), dynamic provisioning works out of the box. A PVC with `storageClassName: gp3` on EKS automatically provisions an EBS gp3 volume."

#### In Depth
Dynamic provisioning is implemented by **external provisioners** running as controllers in the cluster. The CSI driver for your storage backend acts as the provisioner. When a PVC is created, the CSI provisioner calls the storage backend's API to create the underlying volume and then creates a PV object referencing it.

---

### 7. What is the reclaim policy of a PersistentVolume?

"The reclaim policy determines what happens to a PV when its PVC is deleted:

- **Retain**: PV moves to `Released` state. Underlying storage retained. Manual cleanup required. **Use for production data.**
- **Delete**: PV and underlying storage both deleted automatically. Default for dynamically provisioned PVs. **Dangerous for important data.**
- **Recycle** (deprecated): Data wiped, PV made available again. Not supported by most CSI drivers.

I've seen teams lose data because they used `Delete` policy in production. Always use `Retain` for databases and critical storage."

#### In Depth
When a PV's reclaim policy is `Retain` and the PVC is deleted, the PV's status shows `Released` not `Available`. You cannot directly bind a new PVC to a `Released` PV — you must manually remove the `claimRef` from the PV spec to make it `Available` again. This is an intentional safety gate.

---

### 8. How do StatefulSets handle persistent storage?

"StatefulSets use `volumeClaimTemplates` to automatically create a **unique PVC per pod replica**.

Pod `database-0` gets PVC `data-database-0`, pod `database-1` gets `data-database-1`. If database-0 is rescheduled to a different node, it reattaches to its own PVC — state is preserved.

This is what makes StatefulSets suitable for databases like Postgres, Cassandra, or Kafka. The pod identity (name, PVC, DNS name) is stable across restarts."

#### In Depth
When you scale down a StatefulSet (e.g., from 5 to 3 replicas), the PVCs for the removed pods are **NOT deleted**. The data survives. If you scale back up, the same PVCs are reused. This is by design for safety — but it means you need to manually delete orphaned PVCs when decommissioning. Automate this with lifecycle hooks or cleanup jobs.

---

### 9. What is a VolumeSnapshot?

"VolumeSnapshot creates a **point-in-time backup of a PVC**.

It's analogous to an EBS snapshot or a ZFS snapshot. You can then restore from a snapshot by creating a PVC from it using `dataSource`.

I use VolumeSnapshots in CI pipelines: take a snapshot of the production database clone, create a new PVC from it for staging, and run tests against real data without touching production."

#### In Depth
VolumeSnapshot requires the **CSI snapshot controller** and the storage backend's CSI driver to support the snapshot feature. Not all CSI drivers implement it. The objects involved are: `VolumeSnapshotClass` (like StorageClass for snapshots), `VolumeSnapshot` (the user-created request), and `VolumeSnapshotContent` (the actual snapshot resource, analogous to PV).

---

### 10. What is CSI (Container Storage Interface)?

"CSI is the **standard interface** between Kubernetes and external storage systems.

Before CSI, storage drivers were compiled directly into the K8s binary — updating a storage driver required a K8s release. CSI decouples this: storage vendors ship their driver as a container that runs in the cluster.

CSI drivers implement capabilities like provisioning, attaching, mounting, resizing, snapshotting, and cloning volumes. AWS EBS CSI, GCE Persistent Disk CSI, NetApp Trident, and Longhorn are all CSI implementations."

#### In Depth
A CSI driver typically consists of: **Node Driver** (DaemonSet — mounts/unmounts on the node), **Controller Driver** (Deployment — creates/deletes volumes via storage API), and **External Sidecars** (attacher, provisioner, resizer, snapshotter). The sidecars watch K8s resources and call CSI driver gRPC APIs. This is the clean separation that allows any storage vendor to plug into K8s.

---

### 11. What is the difference between ReadWriteOnce and ReadWriteMany?

"`ReadWriteOnce (RWO)`: The volume can be mounted read-write by **exactly one node**. Multiple pods on the same node can use it. Block storage like EBS, Azure Disk, GCE PD.

`ReadWriteMany (RWX)`: The volume can be mounted read-write by **multiple nodes simultaneously**. This requires network file systems like NFS, CephFS, Azure Files, or EFS (on AWS).

I've seen teams burn themselves by trying to mount EBS (RWO only) across multiple nodes for distributed caching — it doesn't work. For shared storage across pods on different nodes, you must use RWX-capable storage."

#### In Depth
There's also `ReadOnlyMany (ROX)` — multiple nodes can mount read-only. Useful for shared config files or read-only datasets. The access modes declared in PVC and PV must be **compatible**: the PVC's requested mode must be a subset of what the PV supports. A PV supporting RWX can satisfy a PVC requesting RWO, RWX, or ROX.

---

### 12. How do you resize a PVC?

"PVC resizing is done by editing the PVC's `spec.resources.requests.storage` — increase the value.

The storage must support volume expansion (the StorageClass must have `allowVolumeExpansion: true`), and the CSI driver must support online or offline expansion.

I've expanded EBS volumes in production this way without any downtime — the CSI driver calls the AWS API to resize the EBS volume and then expands the filesystem inside the container automatically."

#### In Depth
PVC size can only be **increased**, never decreased. For some volumes, expansion is online (pod keeps running). For others, you need to stop the pod to unmount the volume first (offline expansion). After the resize API call, three things happen: (1) underlying storage is resized, (2) the volume is resized on the node, (3) the filesystem inside the container is expanded. Step 3 may require a pod restart for some filesystems.

---

### 13. What happens if a pod uses a PVC that no longer exists?

"If the PVC that a Pod references doesn't exist, the Pod will stay in **`Pending` state** with the event: `persistentvolumeclaim "my-pvc" not found`.

It will not start until the PVC is created and bound.

I've seen this happen after namespace migrations where PVCs were deleted before the manifests were applied. Always apply PVCs before the pods that use them, or use a dependency management tool like Helm's lifecycle hooks."

#### In Depth
This is also a critical failure mode during **disaster recovery**. If you restore workloads but forget to restore PVCs, all pods that need the PVCs will be stuck in pending. Tools like Velero handle this by backing up and restoring PVCs alongside other resources and ensuring ordering.

---

### 14. What is ephemeral storage?

"Ephemeral storage is **temporary storage** that exists only for the lifetime of a pod.

Types: `emptyDir`, `ConfigMap`, `Secret` volumes, the container's writable layer. It's stored on the node's local disk.

You can set resource limits on ephemeral storage: `resources.limits.ephemeral-storage: 1Gi`. If a pod exceeds this, it's evicted. This prevents a misbehaving pod from filling the node's disk and bringing down all other pods."

#### In Depth
Ephemeral storage tracking is done by the kubelet via the `du` command or inotify watches. The limit covers the container's writable layer plus any `emptyDir` volumes. Logs written by the container to stdout/stderr are stored in the node's log directory and count against the node's storage, not the pod's ephemeral limit. Use log rotation (`--container-log-max-size` on kubelet) to prevent log accumulation.

---
