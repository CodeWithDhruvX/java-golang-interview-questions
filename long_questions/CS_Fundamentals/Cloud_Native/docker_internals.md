# Docker Internals: Namespaces & Cgroups

## 1. What is a Container?
A container is **NOT** a virtualization technology like a VM. A container is just a normal Linux process, but with a restricted view of the system.
This restriction is achieved using two Linux Kernel features: **Namespaces** and **Cgroups**.

## 2. Linux Namespaces (Isolation)
Namespaces provide **Isolation**. They limit *what a process can see*.
When you run a container, the Kernel wraps the process in a set of namespaces so it thinks it is the only process on the machine.

### Key Namespaces:
1.  **PID Namespace**:
    *   *Effect*: The process sees itself as PID 1. It cannot see processes running on the host or in other containers.
2.  **MNT (Mount) Namespace**:
    *   *Effect*: The process has its own root filesystem (`/`). Modifications to `/etc` or `/bin` inside the container do not affect the host.
3.  **NET (Network) Namespace**:
    *   *Effect*: The process has its own IP address, routing table, and localhost interface.
4.  **UTS (Unix Timesharing) Namespace**:
    *   *Effect*: The process has its own Hostname.
5.  **USER Namespace**:
    *   *Effect*: The process can be `root` inside the container but a non-privileged user on the host (Security mapping).

## 3. Cgroups (Control Groups)
Cgroups provide **Resource Limiting**. They limit *how much a process can use*.
Even if a process is isolated, it could still crash the host by consuming 100% CPU or RAM. Cgroups prevent this.

### What Cgroups control:
1.  **CPU**: "This container can only use 0.5 cores".
2.  **Memory**: "This container is limited to 128MB RAM". (If it exceeds, OOM Killer kills it).
3.  **Disk I/O**: Limit Read/Write ops per second.
4.  **PIDs**: Limit max number of processes inside a container (prevents Fork Bombs).

## 4. Overlay Filesystem (UnionFS)
How do Docker images work efficiently?
They use a **Layered File System**.
*   **Image Layers**: Read-Only layers (e.g., Ubuntu Base -> Java Install -> App Code). Shared across containers.
*   **Container Layer**: A thin Read-Write layer on top. Any changes made by the running container happen here (Copy-on-Write).
*   *Benefit*: High storage efficiency and fast startup.

## 5. Interview Questions
1.  **Does Docker run on Windows/Mac?**
    *   *Ans*: Natively, No. Docker relies on *Linux* Kernel features. On Windows/Mac, Docker Desktop runs a lightweight Linux VM (using WSL2 or HyperKit) to host the containers.
2.  **Difference between VM and Container?**
    *   *Ans*:
        *   **VM**: Hardware Virtualization. Has a full OS kernel (Heavy, Slow boot).
        *   **Container**: OS Virtualization. Shares the Host Kernel (Light, Instant boot).
3.  **What is PID 1 in a container?**
    *   *Ans*: It is the entrypoint process (e.g., `java -jar app.jar`). If PID 1 dies, the container dies. PID 1 must also handle Unix Signals (SIGTERM) correctly to shut down gracefully.
