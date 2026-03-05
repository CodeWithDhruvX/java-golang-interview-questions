# Git Basics (Service-Based Companies)

This document covers basic Git interview questions frequently asked in service-based companies (TCS, Infosys, Wipro, Cognizant, Accenture, Capgemini).

## 1. What is Git? How is it different from SVN?

**Git** is a Distributed Version Control System (DVCS) used for tracking changes in source code during software development. It allows multiple developers to work together non-linearly.

**Differences from SVN:**
*   **Architecture:** Git is Distributed (every developer has a full local copy of the repo history). SVN is Centralized (history is stored on a central server, and developers check out only the latest working copy).
*   **Speed:** Git is extremely fast because most operations are performed locally. SVN requires network access for most operations.
*   **Branching:** Branching and merging in Git are lightweight and easy. In SVN, branching is essentially creating a full directory copy, making it heavier and slower.
*   **Offline Work:** Git allows full committing, branching, and logging while offline. SVN requires a server connection to commit.

## 2. What is the difference between Git and GitHub?

*   **Git:** The fundamentally open-source version control software installed locally on your computer that manages the code history.
*   **GitHub:** A cloud-based hosting service by Microsoft that allows developers to host their Git repositories online, collaborate with others, review code via Pull Requests, and manage projects.

*Analogy:* Git is like the engine inside a car, whereas GitHub is a global network of roads and parking lots where your car can be driven and parked.

## 3. Explain the Git lifecycle.

The Git life cycle involves handling files across three distinct states/trees:
1.  **Working Directory:** Your local file system where you edit files. Files here can be tracked or untracked.
2.  **Staging Area (Index):** When you use `git add`, files are moved to the staging area. This is a preparatory zone where you assemble your next commit.
3.  **Local Repository (HEAD):** When you use `git commit`, the snapshot of the staging area is saved permanently to the local Git repository (.git directory).
4.  *(Optional but standard)* **Remote Repository:** When you use `git push`, the commits from your local repository are pushed to a remote server (like GitHub or GitLab).

## 4. What is a repository?

A **repository** (or "repo") is the central location/folder where Git stores all your project files along with the complete history of changes made to them.
*   **Local Repository:** Lives on your computer (.git folder inside your project directory).
*   **Remote Repository:** Hosted on a server (e.g., GitHub, BitBucket) for sharing and collaborating with a team.

## 5. What is the difference between `git pull` and `git fetch`?

*   **`git fetch`:** Only downloads the latest changes (commits, branches, tags) from the remote repository to your local repository. It *does not* integrate or merge those changes into your current working branch. It is a completely safe operation.
*   **`git pull`:** Does two things: it first runs `git fetch` to download the changes, and then immediately runs `git merge` to integrate those changes into your current active branch. It can result in merge conflicts if your local changes overlap with remote changes.

## 6. What is the difference between `git clone` and `git fork`?

*   **`git clone`:** A Git command used to create a direct local copy of a remote repository on your personal computer. It creates a `.git` folder locally, linking it to the remote source. 
*   **`git fork`:** Not a standard Git command, but a feature of platforms like GitHub/GitLab. It creates a server-side copy of a repository into your *own* GitHub account. This allows you to freely experiment with changes without affecting the original project. You usually "clone" your "fork" to work on it locally.

## 7. What is staging in Git?

Staging in Git is the step before committing. When you modify files, Git recognizes them as changed in the working directory. The **staging area (or index)** is a space where you gather and arrange exactly what changes you want to include in your next commit snapshot. You add files to the staging area using the `git add <file>` command. This allows you to intentionally pick partial changes in your code to form distinct logical commits.

## 8. What is the difference between `git add .` and `git add -A`?

Historically (before Git 2.0), there was a big difference:
*   `git add .`: Staged new and modified files in the current directory, but ignored deleted files.
*   `git add -A` (or `git add --all`): Staged new, modified, and deleted files across the entire repository.

*Note:* Since Git version 2.x, both `git add .` and `git add -A` behave almost identically from the root directory; however, `git add .` stages everything in the *current directory* and downwards, whereas `git add -A` stages everything in the *entire repository* regardless of the folder you are currently in.

## 9. How do you resolve merge conflicts?

A merge conflict occurs when Git cannot automatically resolve differences between two branches (e.g., two developers edited the same line of code).
Steps to resolve:
1.  **Identify:** Git will halt the merge and mark conflicts in the affected files. `git status` reveals which files are in conflict.
2.  **Open the file:** Look for Git's conflict markers (`<<<<<<<`, `=======`, `>>>>>>>`).
3.  **Edit:** Manually clean up the file by choosing which code to keep (Current Change, Incoming Change, or a mix of both), and removing the `<<<<<<<`, `=======`, `>>>>>>>` markers.
4.  **Stage:** Once resolved, use `git add <resolved-file>` to mark the conflict as resolved.
5.  **Commit:** Finally, run `git commit` to finalize the merge.

## 10. What is `.gitignore`?

`.gitignore` is a text file located in the root of your Git repository. It tells Git which files or directories to completely ignore and NOT track in version control.
Commonly ignored files include:
*   Compiled code or build artifacts (e.g., `*.class`, `/target/`, `/dist/`).
*   Dependency directories (e.g., `node_modules/`).
*   Environment files containing sensitive secrets (`.env`).
*   IDE config files (`.idea/`, `.vscode/`).
