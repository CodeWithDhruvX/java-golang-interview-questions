# Git Tricky & Edge Cases (Product-Based Companies)

This document covers difficult, edge-case, and deeply technical scenario questions frequently asked in senior product-based company interviews.

## 1. Two commits with the same hash? Is it possible?

Yes, but it is astronomically improbable. This is known as a **SHA-1 Collision**. A few years ago, Google researchers proved that it is technically possible to craft two different documents that hash to the exact same SHA-1 value (the SHAttered attack).
*   **Git's Defense:** Because of this theoretical vulnerability, Git has built-in detection that alerts developers and aborts operations if a SHA-1 collision is detected. Newer versions of Git are actively transitioning the internal cryptographic hash function from SHA-1 to the vastly more secure **SHA-256** to permanently eliminate this risk.

## 2. What happens if the `.git` folder is accidentally deleted?

The `.git/` folder *is* the entire repository. It contains all metadata, commit history, branches, branches, and object databases. Unseen to the casual observer, deleting this directory deletes the entire project history.
*   The project folder simply reverts to a normal, un-versioned directory. All files currently in the Working Directory are maintained as normal files.
*   *Recovery:* If the project was cloned from a remote origin, you simply have to `git clone` the project down again into a new folder to recover the full `.git` history, or run `git init` and push all current files as the initial commit if it was entirely a local project.

## 3. What is Sparse Checkout?

**Sparse Checkout** is a feature designed for gigantic monorepos to improve performance by restricting the working directory to only a subset of the repository.
*   By default, `git checkout` pulls down every single file across 100,000 directories in the project.
*   A backend developer working only in `services/payment/` can configure Git to sparsely checkout *only* that folder. The `services/user/` and `frontend/` folders won't physically exist on their hard drive in the working directory, saving immense time and disk space, even though the `.git` repository maintains the meta-history for the entire project.

## 4. What is a Shallow Clone (`--depth`)?

**Shallow Clone** is used to download only a limited, truncated history of a repository. Instead of pulling down a gigabyte packfile containing 10 years of commits, `git clone --depth 1 <URL>` pulls down only the very latest commit on the main branch.
*   This is incredibly fast and saves massive bandwidth. It is the gold standard for CI/CD pipelines that simply need the immediate code snapshot to run a build or execute unit tests but never need to traverse historical logs or run a `git blame`.

## 5. How do you handle Binary Files in Git?

Git is engineered for tracking text-based source code using line-by-line delta compression. Huge binary files (videos, compiled `.jar` files, textures, large images) do not compress well. Changing a single pixel in a 50MB image results in Git storing an entire new 50MB blob. This bloats repositories until they become unusable.
*   **Best Practice:** Use `.gitignore` to prevent any binary build artifacts from entering version control.
*   **For necessary assets (video games/marketing):** Use **Git LFS (Large File Storage)**. LFS replaces the massive binary files with tiny, lightweight text pointers inside the Git repository, while storing the actual massive files on a dedicated external file server.

## 6. What are the consequences of a Force Push?

`git push --force` (or `-f`) overwrites the remote repository's branch history to strictly match your local client's history. It is highly destructive.
*   **The danger:** If a colleague pushed 5 new commits to the remote branch while you were offline, and you perform a force push, you silently erase all 5 of their commits from the server.
*   **Safer Alternative:** Use `git push --force-with-lease`. This performs a safety check before pushing. It verifies that the remote branch has exactly the same commit hash as you thought it did when you last ran `git fetch`. If someone else added a new commit in the meantime, the push safely fails, alerting you to pull their changes first.

## 7. What happens if you rebase a shared branch — what to watch out for?

The golden rule of Git is "Never rebase a shared, public branch."
*   If you rebase the `main` branch locally, Git rewrites the history, creating entirely new SHA hashes for all the old commits.
*   If you then force-push this new history to the remote server, every other developer on your team will be completely out of sync. When they try to `git pull`, Git will try to merge their old, original history with the new, rewritten history, creating an apocalyptic storm of duplicate commits and unmanageable merge conflicts across the entire company.

## 8. How do you undo your last commit without losing your file changes?

**Scenario:** You committed 10 files with message "Fix login bug," but immediately realized you forgot to include a crucial CSS file.

**Solution:**
1.  **Fast way:** Simply stage the forgotten CSS file (`git add style.css`), and then run `git commit --amend --no-edit` (to keep the old message) or `git commit --amend` (to write a new message). This modifies the existing commit instead of creating a second "oops" commit.
2.  **Structural way:** You can rewind history but keep the changes in the staging/working area using `git reset --soft HEAD~1`. This removes the commit from history, leaving all 10 modified files sitting nicely in your staging area ready to be adjusted and re-committed.

## 9. How do you remove sensitive info from Git history (like passwords)?

If you accidentally committed a database connection string or AWS API key a month ago, it is permanently locked inside Git's historical objects database, even if you delete the file in a new commit today.
*   You must *rewrite history* completely to scrub it from every past commit it touched.
*   Use `git filter-repo --path sensitive-config.json --invert-paths` (which deletes the entire file from existence across all time).
*   Alternatively, use BFG Repo Cleaner or `git filter-branch` to rewrite commits and strip the specific password string from historical text blobs. And importantly, actively **rotate/revoke** that password or API key instantly before attempting Git fixes.

## 10. How do you recover overwritten commits after a reset or rebase?

**Scenario:** You ran a disastrous `git reset --hard HEAD~10` and wiped out 10 commits, or a branch got rewritten poorly in a rebase. Your commits aren't in your `git log` graph anymore.

**Solution:**
1.  Immediately run `git reflog`. This displays a chronological log of everywhere the `HEAD` pointer has moved on your local computer over the last 90 days.
2.  Scan down the list until you find the moment right before you ran the destructive command (e.g., `HEAD@{6}: commit: add new logging feature`).
3.  Note the hash next to that action (e.g., `d3f8a2c`).
4.  Run `git branch rescue-branch d3f8a2c` to instantly resurrect all those lost commits onto a safe new branch.
