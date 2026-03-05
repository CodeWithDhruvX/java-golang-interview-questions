# Git Automation, Scripting & Hooks (Product-Based Companies)

This document explores how Git can be programmed and automated to enforce workflows, validate code, and integrate seamlessly with external enterprise systems.

## 1. How do you automate repetitive Git tasks?

Repetitive tasks like checking out specific branches, rebasing against `main`, and resolving standard semantic version tags can be automated using several methods:
*   **Git Aliases:** You can configure custom shortcuts in your `~/.gitconfig` file. 
    *   Example: `git config --global alias.co checkout` or `git config --global alias.cleanup "!git branch --merged | grep -v '*' | xargs -n 1 git branch -d"`. The bang `!` allows you to execute arbitrary bash commands.
*   **Shell Scripts/Makefiles:** Writing Bash/Python scripts that wrap multiple Git commands sequentially. Useful for complex release procedures.
*   **Git Hooks:** Automatically executing specific scripts triggered by standard Git events (commit, push, checkout).
*   **Server-side Webhooks:** The modern enterprise standard. GitHub/GitLab fires JSON payloads to external CI/CD orchestration servers (like Jenkins or GitHub Actions) to fully automate builds, tests, and semantic deployments.

## 2. How to write a Git hook to enforce commit messages format?

Organizations often strictly enforce **Conventional Commits** (e.g., `feat: API added`, `fix: correct layout crash`).
To enforce this locally before a commit is even created, you use a **`commit-msg` hook**:
1.  Navigate to `.git/hooks/`.
2.  Create a file named precisely `commit-msg` (no file extension) and make it executable (`chmod +x commit-msg`).
3.  Write a script (usually Bash or Node.js) that reads the temporary file containing the developer's commit message (passed as the first argument, `$1`).
4.  The script uses a Regular Expression (Regex) to validate the first line of the message.
5.  If the Regex fails to match the `type: subject` format, the script explicitly prints an error message (`echo "Error: Commit message doesn't follow conventions"`) and exits with a non-zero status code (`exit 1`), completely aborting the commit.
6.  *Modern Approach:* Instead of manually writing bash scripts, teams use tools like `Husky` (for Node.js) or `pre-commit` (for Python) to centrally manage and automatically install these hooks across the entire team's local environments.

## 3. How do you integrate Git with CI/CD pipelines?

The foundational integration mechanism is the **Webhook**.
1.  **Configuration:** A repository administrator configures a Webhook URL within the GitHub/GitLab repository settings, pointing it to the CI/CD server (e.g., `https://jenkins.company.com/github-webhook/`).
2.  **Trigger:** A developer pushes code or opens a Pull Request.
3.  **Payload:** GitHub fires an HTTP POST request containing an enormous JSON payload to Jenkins. The payload describes the exact event type (`push`), branch name (`feature/login`), commit hashes, and author details.
4.  **Action:** Jenkins intercepts the Webhook, verifies the cryptographic signature (using a shared secret configuration to prevent spoofing), and automatically clones the specifically mentioned Git commit hash.
5.  **Execution:** Jenkins then sequentially runs the testing, linting, and build scripts defined in the repository's `Jenkinsfile` or YAML configuration, before deploying the artifact.

## 4. How to automate tagging and releases via scripts?

In mature CI/CD pipelines, release management shouldn't involve developers manually typing `git tag v2.0.0` or manually updating `CHANGELOG.md` files. This is mostly solved by **Semantic Release** tools.
*   **The Workflow:**
    1.  Developers merge code into the `main` branch continuously.
    2.  A CI pipeline kicks off. It analyzes all new commit messages since the last successful release tag.
    3.  If it detects commits labeled `fix:`, it bumps the PATCH version (e.g., `1.0.0 -> 1.0.1`). If it detects `feat:`, it bumps MINOR (`1.0.0 -> 1.1.0`). If it detects `BREAKING CHANGE:`, it bumps MAJOR (`1.0.0 -> 2.0.0`).
    4.  The pipeline script automatically generates release notes inside a `CHANGELOG.md` file based directly on the commit messages.
    5.  It automatically creates a Git tag mapped to the new version and performs a `git push origin --tags` back to the repository using a machine service account token, finalizing the automated release.

## 5. How to run a batch rebase or batch merge using scripting?

Large open-source projects or enterprise squads might receive dozens of Pull Requests daily from external contributors that need to be rebased against a rapidly moving `main` branch before being tested.
*   **The Problem:** Rebasing 30 branches manually is exceptionally tedious and error-prone.
*   **The Scripting Solution:** A repository maintainer can build a Bash script to automate it:
    1.  Fetch all active PR branches locally: `git fetch origin`.
    2.  Write a loop that sequentially checks out each PR branch: `for branch in $(git branch -r | grep 'feature/'); do`
    3.  Execute `git rebase origin/main` inside the loop.
    4.  If the shell returns an exit code of `0` (success), forcefully push the newly rebased branch back: `git push -f origin HEAD`.
    5.  If it returns a non-zero exit code (meaning merge conflicts occurred), automatically run `git rebase --abort` and print a warning for manual intervention before moving onto the next branch.
