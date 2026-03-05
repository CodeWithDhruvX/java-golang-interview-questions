# Git Workflows, Scaling & Security (Product-Based Companies)

This document covers high-level, architectural Git questions asked for senior developer and architect roles at product companies, focusing on team collaboration, scaling repositories, and CI/CD pipelines.

## 1. Compare Git Flow, GitHub Flow, and Trunk-Based Development

*   **Git Flow:** A strict branching model utilizing two long-lived branches (`master` for production releases, `develop` for integration) and several short-lived branches (`feature/`, `release/`, `hotfix/`). Best for strict release cycles (e.g., mobile apps, desktop software) but can be heavy and slow for agile web teams.
*   **GitHub Flow:** A simplified model utilizing a single long-lived branch (`main`), where developers create short-lived feature branches, collaborate via Pull Requests, merge into `main`, and deploy continuously to production. Best for continuous deployment (CD) workflows.
*   **Trunk-Based Development:** All developers commit directly to a single branch (`trunk` or `main`), often multiple times a day. Feature branches are strictly short-lived (less than a day) or entirely nonexistent (using feature flags instead). Essential for true Continuous Integration (CI) and large-scale engineering teams.

## 2. How do you manage hotfixes in production safely?

In most branching strategies (like Git Flow):
1.  Immediately branch off the exact commit/tag currently running in production (e.g., `git checkout -b hotfix/payment-bug v1.2.0`).
2.  Fix the critical bug directly in that branch and test it rapidly.
3.  Deploy the hotfix branch to production.
4.  Merge the `hotfix` branch back into `master`/`main` *and* cross-merge it back into the active development branch (`develop`) to ensure the bug doesn’t regress in the next release.

## 3. What are Git Hooks? Give examples of pre-commit and pre-push hooks.

Git Hooks are custom scripts (bash, Python, Node) that Git automatically executes before or after events such as committing or pushing. They live in `.git/hooks/`.
*   **`pre-commit`:** Runs locally *before* a commit is created. Often used to enforce code formatting (e.g., Prettier, Black), run fast linters (ESLint), or ensure commit messages follow a formal structure (Conventional Commits).
*   **`pre-push`:** Runs locally *before* commits are pushed to the remote server. Often used to run a fast suite of unit tests. If the tests fail, the hook blocks the push, saving CI pipeline minutes and keeping the remote branch clean.

## 4. How does Git integrate with CI/CD pipelines?

Git is the foundation of modern CI/CD. When a developer pushes a commit or opens a Pull Request to a remote Git hosting service (GitHub/GitLab), the service emits a **Webhook** payload to the CI server (Jenkins, Actions).
The CI engine reads the payload, clones the specific repository commit securely, and sequentially executes the pipeline configuration (build, test, lint, deploy statuses). It then posts the status (Success/Failure) back to the Git provider's Pull Request UI to block or allow merges.

## 5. How do you optimize large repositories that are slow to clone and operate?

*   **Bloated History (Pack files):** Manually trigger `git gc --prune=now` to garbage collect and compress loose objects into optimal packfiles.
*   **Shallow Clone:** If a CI pipeline only needs to build the latest code without historical context, use `git clone --depth 1 <URL>`. This fetches only the absolute latest snapshot, dropping gigabytes of historical blobs.
*   **Single Branch Clone:** `git clone --single-branch --branch main <URL>` fetches history for only one specific branch instead of the entire repo.
*   **Sparse Checkout:** Used in massive monorepos. It allows developers to configure Git to only checkout a specific subdirectory of a massive repository into their Working Directory, ignoring thousands of irrelevant folders.

## 6. How do you handle huge binary files (videos, models, PSDs) in Git?

Git is terrible at tracking large binaries. Since it delta compresses text, changing a 500MB video file results in Git storing an entirely new 500MB blob, rapidly bloating the `.git` directory history.
*   **Solution: Git LFS (Large File Storage).** LFS replaces the massive binary files in your repository with tiny text "pointers." The actual 500MB file content is hosted on a separate remote LFS storage server. When you checkout the repo locally, the LFS client intercepts the pointer and securely downloads the large file on demand.

## 7. What are Git signed commits (GPG signing)?

A standard Git commit only relies on an email string (`dhruv@example.com`), meaning anyone can push a commit pretending to be someone else.
**GPG Signed Commits** use public-key cryptography. You configure Git with a private GPG key to cryptographically sign every commit you create. The remote server (e.g., GitHub) uses your public key to verify the signature. If a malicious actor spoofs your email address, their commits will display an "Unverified" badging. Product companies use this to tighten supply chain security.

## 8. How do you handle access control and protected branches in a large team?

At an enterprise product company:
1.  **Branch Protection Rules:** The `main` or `release` branches are locked via GitHub/GitLab UI. Developers are strictly blocked from pushing directly to them.
2.  **Pull Request Requirements:** Merging code into a protected branch requires an opened PR, at minimum 1 or 2 approving reviews from Code Owners, and a passing CI pipeline (green build and tests).
3.  **Linear History Constraints:** Force the PR merges to strictly "Squash and Merge" or "Rebase and Merge" to ensure the main history graph never contains messy merge commits.

## 9. What is Git sparse checkout?

`git sparse-checkout` is an advanced optimization for massive monorepos (like Google or Microsoft scale). While a normal clone checks out every single file across thousands of microservices into your working directory, a sparse checkout allows you to define a specific cone of directories (e.g., only `services/payment/` and `libs/shared/`) that you care about. Git will populate your local working directory with *only* those nested files, significantly boosting IDE speeds and search times while ignoring the rest of the colossal codebase.

## 10. How do you search the entire project history for when a specific string of code was added/deleted?

If you need to find the specific commit where an API key or a specific variable name (`calculateTax`) was introduced or removed, use the "Pickaxe" search option:
`git log -S "calculateTax"`
This will traverse the entire history graph and return only the lists of commits that specifically added or removed that exact string from any file.
