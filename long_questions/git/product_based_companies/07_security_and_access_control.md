# Git Security & Access Control (Product-Based Companies)

This document covers high-level security concepts related to Git repositories, code integrity, and supply chain management for enterprise environments.

## 1. What are Git Signed Commits (GPG signing) and why are they used?

**Signed Commits** use cryptographic keys (like GPG or SSH keys) to provide strict proof of identity for every single commit.
*   **The Problem:** By default, you can configure `git config --global user.email "ceo@yourcompany.com"` and Git will happily attach that fake email to all your commits. A malicious actor could spoof their identity and push backdoors to a repository, framing someone else.
*   **The Solution:** You generate a private/public key pair. You tell Git to sign every commit using your private key (`git commit -S`). You upload the public key to GitHub/GitLab. The server automatically verifies the cryptographically signed commit against your public key. If the signature is valid, it adds a green "Verified" badge.
*   *Enterprise Security:* Many strict enterprise environments fully block unsigned commits from being pushed or merged into `main`.

## 2. How do you verify the authenticity of a commit?

To verify a commit locally in the terminal rather than relying on the GitHub UI:
*   Run `git log --show-signature`. 
*   This command checks the GPG signature embedded against the committer's public key stored in your local keyring and explicitly states `Good signature from "User Name (Email)"` or `Bad signature` if the underlying object hash was tampered with.

## 3. How do you enforce branch protection in GitHub/GitLab?

Branch protection rules are server-side configurations that strictly guard critical branches (like `main`, `master`, or `release/v1`) from accidental or malicious modification. Common enforcements include:
*   **Require Pull Request reviews before merging:** Prevents any developer from pushing code directly to the branch. It must pass through a PR and receive at least 'N' approvals.
*   **Require status checks to pass before merging:** Mandates that automated CI/CD pipelines (unit tests, security scans, linters) must run and return a "Success" status before the "Merge" button is clickable.
*   **Require signed commits:** Blocks any commits the server cannot cryptographically verify.
*   **Include administrators:** Forces repository admins/owners to also follow all the above rules instead of bypassing them.
*   **Restrict who can push to matching branches:** Limits merge access strictly to specific teams (e.g., only the 'Release Engineering' team can push to `main`).

## 4. Difference between SSH vs HTTPS remote URLs?

*   **HTTPS (`https://github.com/user/repo.git`):**
    *   Authenticates over standard port 443. 
    *   Usually requires a Personal Access Token (PAT) rather than a password for interaction.
    *   Easier to set up initially on new machines or behind strict corporate firewalls that block all non-standard ports.
*   **SSH (`git@github.com:user/repo.git`):**
    *   Authenticates securely using an asymmetric RSA/Ed25519 key pair over port 22.
    *   Requires generating keys (`ssh-keygen`) and pasting the public portion into your GitHub profile.
    *   Once configured, you never have to type a username or token again for subsequent clones, pulls, and pushes.
    *   *Security Context:* Most senior developers strictly prefer SSH. It is more secure (no tokens sitting in environment files) and explicitly relies on the machine's private key.

## 5. How do you handle access control for large teams?

Access control is managed at the organizational level on the Git hosting platform (GitHub/GitLab), not within core Git commands.
1.  **RBAC (Role-Based Access Control):** Users are assigned to specific Identity Provider (IdP) groups (like Okta/Azure AD) rather than added individually. Those groups map directly to GitHub Teams.
2.  **Granular Permissions:**
    *   *Read/Pull:* All developers can clone the code.
    *   *Write/Push:* Only specific feature teams can push branches to specific repositories.
    *   *Maintain/Admin:* Only senior leads can configure branch protections or delete the repository.
3.  **Code Owners (`CODEOWNERS` file):** A special file that defines strictly which individuals or groups *must* review and approve Pull Requests that modify specific directories (e.g., `* @security-team` or `/database/* @db-admins`). It automates routing and prevents unapproved modifications to critical infrastructure code.

## 6. How do you prevent accidental pushes to `master`/`main`?

*   **Server-Side:** Ensure the branch is "Protected" via GitHub/GitLab UI. This is the only bulletproof way, as it forces the server to outright reject direct `git push origin main` attempts from any user.
*   **Client-Side (Local safety nets):** 
    1.  Use a `pre-push` Git Hook script locally that checks `if branch == 'main'`. If true, the script exits with an error code, aborting the push before it ever talks to the network.
    2.  Set `git config --global push.default simple` (the modern default) requiring your current branch name to perfectly match the remote branch.

## 7. What is git-crypt and how does it secure secrets?

To avoid accidentally committing passwords to public databases, teams use tools like **`git-crypt`**. 
*   It enables transparent encryption and decryption of specific files in a Git repository using GPG keys or a shared symmetric key.
*   You configure a `.gitattributes` file to mark sensitive configuration files (e.g., `config/secrets.yml filter=git-crypt diff=git-crypt`). 
*   When you commit those files, `git-crypt` silently encrypts the text locally. The Git database and GitHub only see binary, AES-256 encrypted gibberish.
*   When an authorized developer pulls the repository and runs `git-crypt unlock`, the tool transparently decrypts the files back into plain text locally on their machine.

## 8. What are pre-commit and pre-push hooks specifically for security?

Security-focused Git hooks run locally to catch common errors before they leave the developer's machine:
*   **Secret Scanning:** Hooks like `trufflehog` or `git-secrets` scan every line of code being committed for regex patterns matching API keys, AWS tokens, or passwords. If found, the commit is aborted instantly.
*   **Dependency Auditing:** Hooks that run `npm audit` or `safety check` (Python) to ensure newly added open-source dependencies do not have known CVE vulnerabilities before allowing the commit or push.
