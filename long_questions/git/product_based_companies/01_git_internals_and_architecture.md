# Git Internals & Architecture (Product-Based Companies)

This document covers deep, conceptual Git architecture and internals questions that are heavily focused on during interviews at top product-based companies (Google, Microsoft, Amazon, Atlassian, etc.).

## 1. What is the difference between working directory, staging area, and repository?

*   **Working Directory:** The local filesystem directory where you checkout a specific version of your project and edit files. Files here can be modified but aren't necessarily tracked by Git yet.
*   **Staging Area (Index):** A specialized data structure (a binary file usually located at `.git/index`) that acts as a caching layer. It holds the snapshot of the files that will go into the *next* commit.
*   **Repository (The `.git` directory):** The database where Git permanently stores all metadata and object databases containing the complete history of your project. When you commit, Git takes the snapshot from the staging area and stores it here permanently.

## 2. How does Git calculate commit hashes?

Git calculates commit hashes using a cryptographic hash function called **SHA-1** (though newer versions are migrating to SHA-256). 
When you make a commit, Git hashes the following metadata:
1.  The hash of the top-level **Tree** object (which captures the directory structure and file hashes).
2.  The hash of the **Parent Commit(s)**.
3.  The **Author** and **Committer** information (name, email, timestamp).
4.  The **Commit Message**.

Any change to even a single byte of a file, timestamp, or author name completely changes the resulting 40-character hexadecimal SHA-1 string. This makes Git's history cryptographically secure and immutable.

## 3. What is the role of SHA-1/SHA-256 in Git?

The SHA-1 (and subsequently SHA-256) algorithm acts as a content-addressable storage mechanism. 
*   **Data Integrity:** Because the hash is generated from the content itself, there is no way to alter the content of a file or commit without Git instantly detecting a hash mismatch.
*   **Deduplication:** If two files have the exact same content, they will have the exact same SHA hash. Git will only store one copy of that blob in its database, automatically deduplicating the repository.

## 4. What are pack files?

**Packfiles** are how Git compresses and optimizes its object database to save disk space and improve network transfer speeds (like during `git fetch` or `clone`).
*   Instead of keeping thousands of individual loose objects (blobs, trees, commits), Git periodically runs a garbage collection process (`git gc`).
*   This process packs all those loose objects into a single binary "packfile" (and an accompanying `.idx` index file).
*   Inside the packfile, Git uses **delta compression**. It stores one base version of a file entirely, and for subsequent versions, it only stores the *deltas* (the differences/changes) from the base version.

## 5. What is the `.git` folder structure?

The `.git/` directory contains all metadata and object databases. Key components include:
*   **`objects/`**: The core database storing all content (blobs, trees, commits, tags) mapped by their SHA-1 hashes.
*   **`refs/`**: Contains pointers to commit objects. `refs/heads/` stores local branches, `refs/remotes/` stores remote tracking branches, and `refs/tags/` stores tags.
*   **`HEAD`**: A text file pointing to the currently checked-out branch (e.g., `ref: refs/heads/main`).
*   **`index`**: The binary file representing the staging area.
*   **`config`**: Project-specific configuration settings.
*   **`logs/`**: Keeps a history of where branch tips have pointed (this is what powers the `reflog`).

## 6. What are loose objects?

When you initially add or commit files, Git creates individual, zlib-compressed files inside the `.git/objects/` directory. These individual files are called **loose objects**. Over time, having too many loose objects degrades performance, which is why Git periodically packs them into **packfiles**.

## 7. What is Git Index?

The Git Index is another name for the **Staging Area**. Technically, it is a single, large binary file (`.git/index`) that lists all the files in the current repository along with their timestamps, file permissions, and the SHA-1 hashes of the file contents (blobs) that are staged for the next commit. It bridges the gap between the working directory and the commit history.

## 8. How does Git detect file changes?

Git does *not* read every file character-by-character to detect changes. It relies heavily on the **Index** and **file system metadata**.
When you run `git status`, Git performs a fast stat() system call on files in the working directory and checks their metadata (size, modification timestamp, inode). 
*   If the metadata matches the cache in the Index, Git skips the file.
*   If the metadata differs, Git considers the file potentially modified. It then actually reads the file content, hashes it via SHA-1, and compares this new hash with the hash stored in the Index to confirm a real change.

## 9. What is object compression in Git?

Object compression in Git happens at two levels:
1.  **Zlib Compression:** Every newly created object (blob, tree, commit) is immediately compressed using the standard zlib compression algorithm before being stored as a loose object.
2.  **Delta Compression:** When Git runs garbage collection to create a packfile, it finds similar files (e.g., v1, v2, and v3 of `app.js`) and stores only one full version (usually the most recent) and compressed binary "deltas" for the historical versions.

## 10. What is garbage collection in Git?

Garbage collection (`git gc`) is a maintenance command that Git runs automatically in the background (or can be triggered manually). 
It performs several optimization tasks:
*   Compresses loose objects into highly efficient packfiles.
*   Removes "unreachable" objects (e.g., commits that were abandoned via `git reset --hard` or deleted branches) that are older than an expiration period (default is 2 weeks).
*   Packs references (`refs/`) into a single `packed-refs` file to speed up repository loading.
