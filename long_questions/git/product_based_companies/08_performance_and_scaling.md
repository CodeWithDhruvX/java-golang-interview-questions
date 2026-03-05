# Git Performance & Scaling (Product-Based Companies)

This document addresses advanced strategies for managing massive Git repositories, mono-repos, and ensuring high performance for enterprise teams (e.g., at FAANG companies).

## 1. How does Git handle repositories with tens of thousands of files?

Git was originally designed to manage the huge Linux Kernel, so it handles large numbers of files well natively. However, as scaling pushes into the millions of files (Google/Facebook scale), native Git slows down because operations like `git status` scan the entire working tree and `git clone` downloads massive histories.
To mitigate this, companies use:
*   Sparse Checkouts (reducing the working tree size)
*   Shallow Clones (reducing downloaded history)
*   File System Monitors (FSMonitor daemon) to skip full-tree scans during `git status`.
*   Virtual File Systems (like Microsoft's VFS for Git / Scalar or Meta's Sapling) that lazily load files only when an IDE or compiler explicitly opens them.

## 2. What is Shallow Clone and how does it improve performance?

**Shallow Clone** (`git clone --depth <N> <URL>`) drastically reduces the history pulled from a remote server. By limiting `<N>` to `1`, you only download the absolute latest commit snapshot on the specified branch, rather than gigabytes of trailing history data.
*   **Performance Gain:** It saves massive amounts of network bandwidth and disk space.
*   **Use Cases:** It is strictly used in CI/CD pipelines where historical context is unneeded for building applications, and on constrained embedded systems.

## 3. What are Pack Files and how do they improve efficiency?

**Packfiles** (`.pack` and `.idx` files inside `.git/objects/pack/`) are Git's primary space-optimization mechanism. 
*   Instead of keeping thousands of individual, loose `zlib`-compressed objects (files, trees, commits), Git periodically triggers "Garbage Collection" (`git gc`).
*   The GC process bundles all these thousands of loose loose files into a single, dense binary packfile.
*   **Delta Compression:** Inside the packfile, Git analyzes similar files over time. It stores only one base version completely, and stores tiny "diffs" or "deltas" for historical versions, drastically compressing the overall size of the repository.

## 4. How do you optimize large repositories?

If a local repository becomes sluggish over time, a developer should manually run:
1.  **`git gc --prune=now --aggressive`**: Forces Git to immediately sweep unreachable orphaned objects and aggressively search for better delta-compression opportunities to create tighter packfiles.
2.  **`git repack -a -d`**: Drops old packfiles and actively repacks all reachable objects into a single, new packfile.
3.  **`git config core.fsmonitor true`**: Enables a background daemon that watches the OS file system, making commands like `git status` nearly instant because they no longer have to scan the entire project folder.

## 5. How do you handle huge binary files in Git?

Storing 50MB videos or 1GB AI models in Git rapidly bloats the `.git` folder because delta compression fails on binaries.
*   **Git LFS (Large File Storage):** This open-source Git extension is the industry standard. It intercepts massive binaries, uploads them to an external storage server (like S3), and commits a tiny text pointer file in their place within the Git repository. When checking out a branch, LFS acts as a filter driver to transparently download the actual massive file locally.

## 6. What is the difference between clone depth options (`--depth`) in `git clone`?

*   `--depth 1` creates a shallow clone with only the latest commit. It implies a single branch (usually `main`).
*   `--shallow-since="1 month ago"` allows you to clone the repo's full history only within a specific temporal window, rather than a hard commit count.
*   `--shallow-exclude=v1.0` clones the repo down to a specific commit or tag but excludes everything before it.
*   **Deepening:** You can later convert a shallow clone to a full clone (or pull more history) using `git fetch --unshallow` or `git fetch --depth 10`.

## 7. How do you split a large monorepo into multiple repos?

If a single massive repository is causing intolerable pain and the company decides to break monolithic projects into poly-repos (microservices), it must be done carefully to preserve commit history.
*   **Modern Tool:** Use `git filter-repo`.
*   **Process:** 
    1. Clone the massive repo.
    2. Run `git filter-repo --path specific-microservice-folder/ --to-subdirectory-filter ""` (This rewrites history, isolating *only* the commits related to that folder and promoting it to the root).
    3. Add a new `origin` remote for the new microservice repository.
    4. Provide a `git push -u origin main` to upload the newly extracted project's tailored history.

## 8. What is Scalar and VFS for Git?

When standard Git customizations fail to handle Windows OS scale (tens of millions of files, hundreds of GBs of history), Microsoft developed specialized tools.
*   **VFS for Git** (Virtual File System) was a custom filesystem driver that presented a virtual view of a massive repo to developers. It only downloaded files from the server the exact microsecond an IDE or build tool demanded them (lazy loading).
*   **Scalar** is the modern evolution. Written in C# and integrated deeply into native Git. It utilizes Sparse Checkout, background fetching, and file system monitors to provide near-instant performance for the largest repositories in the world without requiring a complex custom OS filesystem driver.
