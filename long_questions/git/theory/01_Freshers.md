# 🟢 **1–20: Basics & Branching**

### 1. What is Git and why is it used?
"Git is a **distributed version control system (VCS)** that helps developers track changes in their source code over time.

I use it every day to collaborate with my team. It allows multiple people to work on the same project simultaneously without overwriting each other's changes. It keeps a complete history of every modification, which means I can always revert to a previous working state if something breaks.

For me, it’s an essential tool because it acts as a safety net and a central source of truth for the codebase, whether I'm working solo or in a large distributed team."

#### Indepth
Unlike older centralized systems, Git is distributed. This means every developer has a full local copy of the repository, including its entire history. This architecture makes operations like branching, merging, and logging incredibly fast because they don't require network requests.

---

### 2. Difference between Git and GitHub.
"Git is the **tool**, while GitHub is a **service**.

Git is the local version control software I install on my machine to track code changes. GitHub is a cloud-based hosting platform built around Git, owned by Microsoft.

I use Git to make commits and manage branches locally. Then, I use GitHub to host those repositories securely in the cloud, review my colleagues' pull requests, and manage CI/CD pipelines. Other alternatives to GitHub include GitLab and Bitbucket, but they all use Git under the hood."

#### Indepth
GitHub adds a profound collaboration layer on top of Git. Features like Pull Requests (PRs), Issues, Actions, and team access controls are entirely GitHub's concepts, not native Git functionalities. Git purely handles the versioning of files.

---

### 3. Difference between Git and SVN.
"The main difference is their architecture: **Git is distributed**, while **SVN is centralized**.

With SVN, there is only one central repository on a server. When I want to see history or commit, I need a network connection. With Git, I have the entire repository history cloned locally on my machine. 

I prefer Git because it allows me to work completely offline. Branching and merging in Git are also incredibly fast and lightweight compared to SVN, where branching essentially copies the entire directory structure."

#### Indepth
Git stores data as a series of snapshots (using a Merkle tree structure), whereas SVN historically stored differences (deltas) between files. This makes Git incredibly efficient at switching branches, as it primarily updates the working tree to reflect a specific commit's snapshot.

---

### 4. Explain Git workflow (Working Directory → Staging → Repository).
"The Git workflow has three main states that my files transition through.

First is the **Working Directory**: These are the actual files I'm editing on my disk. Second is the **Staging Area (Index)**: When I finish a change and run `git add`, the file moves here. It’s a preparatory area where I group related changes before committing them. Finally, there is the **Repository (.git directory)**: When I run `git commit`, the staged changes are permanently saved as a snapshot in Git's internal database.

I love this three-tree architecture because it allows me to carefully craft my commits. I can modify 10 files but only stage and commit 2 of them, keeping the logical history clean."

#### Indepth
The staging area is a unique concept in Git. Technically, it is a single file (usually `.git/index`) that contains directory tree snapshot information. When you `git commit`, Git just wraps the state of the index into a commit object.

---

### 5. What is a repository? Local vs remote repo?
"A **repository** (or repo) is essentially a project folder that Git is tracking. It contains all the project files alongside their entire version history in the hidden `.git` folder.

A **local repository** lives on my personal computer. I can commit, branch, and merge entirely within it without internet access. A **remote repository** is a version of my project hosted on the internet or a network server (like GitHub or GitLab).

The workflow involves me doing my work in the local repo and then syncing it with the remote repo using `git push` to share my changes, or `git pull` to get updates from my team."

#### Indepth
Technically, Git treats every repository equally—there's no inherent "master" repo. "Remote" is just a named reference (like `origin`) pointing to a URL. When you push or fetch, Git simply synchronizes its object database with the remote one.

---

### 6. What is a commit?
"A **commit** is a snapshot of my repository at a specific point in time. 

When I make a commit, I am permanently recording the changes I've staged into the Git history. Each commit gets a unique ID (a SHA-1 hash) and requires a commit message explaining what was changed.

I think of commits as 'save points' in a video game. If I introduce a bug, I can always look at the commit history and revert the repository back to a previous healthy commit."

#### Indepth
Internally, a commit object in Git is tiny. It contains a pointer to the main tree object (`snapshot`), pointers to immediate parent commits, the author's name and email, a timestamp, and the commit message. It doesn't store the file diffs directly, just pointers to the completely hashed file states (blobs).

---

### 7. What is staging? (\`git add\`)
"**Staging** is the process of selecting which modified files I want to include in my next commit using the `git add` command.

If I edit three files, I might realize that two of them belong to a feature, and one is just a typo fix. Instead of committing all three together, I can use `git add` on the two feature files, commit them, and then stage the typo fix for a separate commit.

This intermediate step is crucial for maintaining a clean and logical project history instead of a messy dump of all current changes."

#### Indepth
Using `git add -p` (patch mode) is a powerful technique. It allows you to stage parts of a file interactively. This is extremely useful if you made multiple logical changes within a single file and want to split them into separate, atomic commits.

---

### 8. \`.gitignore\` — what is it and how does it work?
"The `.gitignore` file is a plain text file where I define patterns for files and directories that Git should completely ignore.

I use it constantly to prevent things like compiled binaries, logs, secrets (`.env` files), and IDE-specific folders (like `.idea/` or `node_modules/`) from being accidentally committed. 

Once a pattern is in this file, running `git status` won't even show those files as untracked, and `git add .` won't stage them. It keeps the repository lightweight and secure."

#### Indepth
If a file has already been committed to the repository, adding it to `.gitignore` will *not* remove it. Git will continue tracking it. To ignore it, you must first untrack it using `git rm --cached <file>` and then commit that deletion, keeping the file on your local disk but removing it from Git's index.

---

### 9. \`git status\` — what info does it show?
"`git status` is the command I use most frequently to check the current state of my working directory and staging area.

It tells me which branch I am currently on and whether it is up to date with the remote branch. It lists files in three categories: **Untracked** (new files Git doesn't know about), **Modified** (changed but not staged), and **Staged** (ready to be committed).

I run this command almost obsessively before every `git add` and `git commit` to ensure I'm exactly sure of what I am about to save."

#### Indepth
The output of `git status` is highly optimized. It quickly scans the working tree and compares file metadata (size and modification time) against the index to detect changes without rehashing the entire file contents.

---

### 10. \`git log\` — different formats and options.
"`git log` displays the chronological commit history of the repository.

By default, it shows the commit hash, author, date, and message. However, I almost always use formatted options. My favorite is `git log --oneline --graph`, which condenses each commit to a single line and draws a nice ASCII graph of the branches on the left.

If I'm debugging, I might use `git log -p` to see the actual code diffs introduced in each commit, or `git log --author="Dhruv"` to filter by a specific person."

#### Indepth
You can define aliases in `~/.gitconfig` for complex log formats. For example, `git config --global alias.lg "log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"` creates a beautiful Custom tree log view you can trigger with simply `git lg`.

---

### 11. What is a branch?
"In Git, a **branch** is simply an independent line of development. 

When I want to add a new feature or fix a bug, I create a new branch. This allows me to experiment and make commits without affecting the main, stable codebase (usually called `main` or `master`). 

For me, branches are lightweight and cheap. They encapsulate my work. Once the feature is complete and tested, I merge my branch back into the main line."

#### Indepth
Internally, a branch is virtually weightless. It is just a lightweight, movable file containing a single 40-character SHA-1 hash pointing to a specific commit. Creating a branch does not copy files; it just creates a new pointer. This is why branching in Git is instantaneous.

---

### 12. Difference between merge and rebase.
"Both commands integrate changes from one branch into another, but they do it differently.

**Merge** combines the histories of both branches, creating a new "merge commit" that ties them together. It preserves the exact chronological history. 
**Rebase** takes my branch's commits and replays them one by one on top of the target branch. This rewrites the commit history to look perfectly linear.

I usually use **rebase** when I'm updating my local feature branch from `main` to keep my commits clean. I use **merge** when a feature is complete and ready to be integrated into `main` permanently."

#### Indepth
The golden rule of Git is "Never rebase commits that exist outside your repository and that people may have based work on." Because rebase rewrites history (generating new SHAs for the commits), doing it on a shared branch will cause massive synchronization issues for other developers.

---

### 13. How to delete a branch (local & remote)?
"To delete a branch, I use two separate commands because local and remote branches are distinct.

To delete a local branch, I ensure I'm not checked into it, and I run `git branch -d branch_name`. If the branch hasn't been merged and I want to force delete it, I use a capital `-D`.

To delete it from the remote repository (like GitHub), I push an empty reference to the remote branch using `git push origin --delete branch_name`."

#### Indepth
The `-d` flag is a safety check; Git will refuse to delete the branch if its HEAD is not reachable from the current branch or its upstream. After a remote delete, other team members will still see their remote tracking branches until they run `git fetch --prune`.

---

### 14. Fast-forward vs three-way merge.
"A **fast-forward merge** happens when the target branch hasn't received any new commits since my feature branch was created. Git simply moves the target branch pointer forward to my latest commit. No new merge commit is created, keeping history strictly linear.

A **three-way merge** happens when both branches have diverged (both have new commits). Git must combine the changes. It finds the common ancestor, looks at both branch tips, and creates a formal 'merge commit' to tie them together.

I often enforce `--no-ff` (no fast-forward) when merging major features into `main` because the explicit merge commit explicitly groups the feature's history."

#### Indepth
A Three-way merge algorithm uses the common ancestor as the base. Fast-forwards can be prevented if you want to retain the context that a group of commits belonged together as a feature, which is the default behavior in GitHub Pull Requests.

---

### 15. Handling merge conflicts.
"A merge conflict happens when Git cannot automatically merge changes—usually because two people modified the same line in a file.

When this happens, Git stops the merge and marks the conflicted files. I open those files in my editor, where I will see conflict markers (`<<<<<<<`, `=======`, `>>>>>>>`). My job is to manually edit the file to look exactly how it should, removing the markers. 

After resolving it, I stage the resolved file with `git add` and finalize the process by running `git commit`."

#### Indepth
Modern IDEs like VS Code have powerful built-in 3-way conflict resolution tools that let you easily accept "Current Change", "Incoming Change", or "Both". Under the hood, Git pauses the merge state in the index, waiting for you to resolve the conflict and complete the commit.

---

### 16. What is a detached HEAD?
"A **detached HEAD** state occurs when I checkout a specific commit hash or a remote tag instead of a local branch name. 

The `HEAD` pointer normally points to a branch reference, which points to a commit. In a detached state, `HEAD` points directly to a commit hash.

It’s great for looking around history or compiling old versions. However, if I make changes and commit in this state, those commits won't be attached to any branch. If I switch away, they become orphaned and can be garbage collected. To save the work, I just need to create a new branch from there using `git checkout -b new-branch`."

#### Indepth
Technically, when `HEAD` is detached, Git writes the commit SHA directly into the `.git/HEAD` file instead of writing `ref: refs/heads/branch_name`. Git's garbage collection (`git gc`) will eventually delete dangling commits that aren't reachable by any reference.

---

### 17. How to rename a branch.
"Renaming a branch is very straightforward.

If I am currently on the branch I want to rename, I just run `git branch -m new_branch_name`. 
If I'm not on it, I specify the old name: `git branch -m old_name new_name`.

If I have already pushed the old branch to a remote repository, renaming is a bit trickier. I have to delete the old remote branch (`git push origin --delete old_name`) and push the newly renamed branch (`git push origin -u new_branch_name`)."

#### Indepth
The `-m` stands for move. This is consistent with Unix conventions where renaming is just moving a file path. In Git's internals, renaming a branch just renames the file inside `.git/refs/heads/`.

---

### 18. Difference between \`git branch\` and \`git checkout -b\`.
"`git branch <name>` simply **creates** a new branch pointing to the current commit, but my `HEAD` stays on the current branch. I'm still working in the old branch.

`git checkout -b <name>` performs two actions in one command: it **creates** the new branch and then immediately **switches** my working directory to it.

I use `git checkout -b` (or the newer `git switch -c`) almost exclusively because 99% of the time I create a branch, I immediately intend to start working on it."

#### Indepth
Go version 2.23 introduced `git switch` specifically to separate branch changing from file restoration, clarifying the heavily overloaded `git checkout` command. The modern equivalent of `checkout -b` is `git switch -c <name>`.

---

### 19. What is Git Flow?
"**Git Flow** is a robust, structured branching model designed around project releases.

It defines strict roles for branches. We have long-running branches: `master` (always production-ready) and `develop` (the active integration branch). Then there are short-lived branches: `feature` branches (for new work), `release` branches (for QA/prep), and `hotfix` branches (to quickly patch production).

I've used it in enterprise environments where we had scheduled deployment cycles. It’s excellent because it provides a very clear history, though it can feel a bit heavy and bureaucratic for fast-moving startups."

#### Indepth
Git Flow is governed by strict rules. For example, a `hotfix` branch is created from `master`, fixed, and then must be merged back into *both* `master` and `develop` to ensure the fix isn't lost in the next release.

---

### 20. GitHub Flow vs Trunk-based development.
"Both are agile alternatives to Git Flow, designed for Continuous Deployment.

In **GitHub Flow**, the `main` branch is always deployable. I create a feature branch, work on it, open a Pull Request, get it reviewed, merge it into `main`, and then deploy. It's very simple and relies heavily on PRs and testing.

In **Trunk-based development**, developers push directly to `main` (the trunk) very frequently, often multiple times a day. Feature branches are practically non-existent or extremely short-lived (hours, not days). It relies on 'feature flags' to hide incomplete work in production.

I prefer Trunk-based development for highly mature CI/CD teams because it practically eliminates "merge hell," whereas GitHub Flow is great for open-source and standard product teams."

#### Indepth
Trunk-based development strictly requires profound automated testing. If developers are constantly pushing to `main`, a broken commit immediately halts the entire team. It is the cornerstone practice for true Continuous Integration (CI).
