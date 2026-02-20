# 🟡 **21–40: Intermediate & Commands Deep Dive**

### 21. Difference between `git pull` and `git fetch`.
"Both commands retrieve updates from a remote repository, but their actions are vastly different.

`git fetch` is safe. It downloads new commits, branches, and tags from the remote repository to my local `.git` directory, but it explicitly does **not** modify my current working tree. It just updates my remote-tracking branches (like `origin/main`). I use it constantly to see what my colleagues have done without touching my active code.

`git pull`, on the other hand, is a two-step command. It first runs `git fetch`, and then immediately runs `git merge` (or `git rebase` if configured) to integrate those downloaded changes into my currently checked-out branch. I use this when my branch is clean and I explicitly want to consume the latest code."

#### Indepth
If you have uncommitted changes in your working directory, `git fetch` will succeed, but `git pull` will likely fail to merge and throw an error to prevent overwriting your unsaved work. You would need to stash or commit them first.

---

### 22. Difference between `git reset`, `git revert`, `git restore`.
"These three commands all 'undo' things, but they have distinct targets.

`git revert` creates a **new commit** that perfectly reverses the changes introduced by a previous commit. This is what I use for public history because it doesn't rewrite the past.

`git reset` moves the branch pointer backward, effectively erasing commits from history. I only use it for local commits that haven't been pushed. `--soft` keeps my files staged, `--hard` completely destroys the changes.

`git restore` is specifically for unmodifying files in my working tree or un-staging them from the index without affecting commits. It’s the modern, safer way to say 'discard my local changes to this file'."

#### Indepth
Before `git restore` was introduced in Git 2.23, you had to use `git checkout -- file` to discard changes and `git reset HEAD file` to unstage them. `git restore` simplifies this, separating branch operations from file operations.

---

### 23. `git stash` — use cases & difference between `pop` and `apply`.
"`git stash` is my way of temporarily shelving incomplete work. 

When I am halfway through a feature but suddenly need to switch branches to fix a critical bug, I run `git stash`. This saves my modifiedtracked files and cleans my working directory so I can switch safely.

When I return, I have two choices. `git stash apply` brings the changes back into my working directory but keeps the stash in the internal list. `git stash pop` applies the changes and immediately deletes the stash from the list. I exclusively use `pop` unless I plan to apply that same stash to multiple different branches."

#### Indepth
By default, `git stash` does not save untracked files (newly created files). To stash those as well, you must use `git stash -u` (include untracked). Stashes are local to your machine and are not pushed to remotes.

---

### 24. `git cherry-pick` — when & why.
"`git cherry-pick` allows me to grab a specific commit from one branch and apply it to my current branch.

I use this when an important bug fix was committed to a feature branch, but I need that exact fix immediately on `main` without waiting for the entire feature to be finalized. I just find the commit hash and run `git cherry-pick <hash>`.

It is incredibly useful, but I use it sparingly because it essentially duplicates the commit (it gets a new SHA-1) which can theoretically lead to merge conflicts later when the original branch is merged."

#### Indepth
You can cherry-pick multiple commits sequentially by providing a range: `git cherry-pick A^..B`. If you want the file changes but not an automatic commit, you use the `-n` (no-commit) flag, staging the changes so you can modify them further.

---

### 25. `git reflog` — what is it & recovery scenarios.
"The `git reflog` is my ultimate safety net. It is a local, chronological log of everywhere my `HEAD` pointer has been.

Unlike `git log`, which only shows commits in the active branch history, the reflog tracks every time I check out a branch, reset, rebase, or commit. 

If I accidentally run a `git reset --hard` and lose a bunch of commits, or if I botch a rebase, I just look at `git reflog`, find the SHA of my branch before the mistake, and run `git reset --hard <old-SHA>`. It practically makes it impossible to lose data locally."

#### Indepth
The reflog automatically clears old entries (usually after 90 days), and it explicitly does not track changes to the working directory that were never committed or stashed. It is strictly local and never shared over the network.

---

### 26. `git clean` — usage.
"`git clean` completely removes untracked files from the working directory.

If I run a build script that generates dozens of temporary compile files that aren't in `.gitignore`, running `git status` becomes a mess. I use `git clean -fd` to forcefully (`-f`) remove untracked files and untracked directories (`-d`).

It is a destructive command—you cannot undo a `git clean` (it doesn't use the recycling bin). Therefore, I always run `git clean -n` (dry run) first to see exactly what Git intends to delete."

#### Indepth
If you have untracked files listed in `.gitignore` (like compiled binaries), `git clean -fd` will actually ignore them! If you genuinely want to wipe the directory to a pristine cloned state, including ignored files, you must use `git clean -fdx`.

---

### 27. `git tag` — lightweight vs annotated tags.
"A **tag** is an immutable pointer to a specific commit, usually marking release points like `v1.0.0`.

A **lightweight tag** is just like a branch that doesn't move. It’s strictly a pointer to a commit hash. I create one with `git tag v1.0`.

An **annotated tag** is a full Git object stored in the `.git/objects` database. It contains a tagging message, date, my author info, and it can be GPG signed to prove authenticity. I create them using `git tag -a v1.0 -m "Release 1.0"`. I *only* use annotated tags for public releases because they contain important metadata."

#### Indepth
By default, running `git push` does not transfer tags to remote servers. You have to be explicit and run `git push origin <tagname>` or `git push origin --tags`.

---

### 28. `git blame` — usage.
"`git blame` determines exactly who made the last modification to each line of a file, and in which commit.

When I find a bizarre piece of legacy code and I can't figure out why it exists. Instead of guessing, I run `git blame file.go` to find the exact commit that introduced the change. This lets me read the commit message or ask the original author directly.

It’s an invaluable tool for understanding the context behind technical decisions."

#### Indepth
`git blame` is sensitive to trivial whitespace changes. If someone runs a code formatter over the entire file, they will become the "author" of every line. To bypass this, always use `git blame -w`, which ignores whitespace modifications and points to the actual logical author.

---

### 29. `git diff` — staging vs unstaged.
"`git diff` shows the exact line-by-line changes in the code.

Running `git diff` naked only shows changes in my **working directory** that I have *not yet* staged (meaning, differences between my files and the index). 

If I have already run `git add`, naked `git diff` shows nothing! To see what I'm actually about to commit (differences between the index and the last commit), I must run `git diff --staged` or `--cached`. I do this obsessively during code reviews of my own work before hitting commit."

#### Indepth
If you want to compare your current local branch against a remote branch to see what will happen if you pull, you would use `git diff master...origin/master`.

---

### 30. `git show` — usage examples.
"`git show` is a simple command to inspect various Git objects.

Most commonly, I use it to view the details of a specific commit: `git show <commit_hash>`. It prints the commit metadata (author, date, message) immediately followed by the patch (the exact diff of what was changed).

It's also great for viewing files at older points in time without checking them out. `git show main:src/main.go` will print the contents of that file exactly as it exists on the `main` branch, even if I'm currently on a feature branch."

#### Indepth
`git show` combined with relative references is extremely powerful. `git show HEAD~1` quickly prints what you did in the second-to-last commit, which is excellent for verifying history before interactive rebasing.

---

### 31. How does Git integrate with CI tools?
"Continuous Integration (CI) tools like Jenkins or GitHub Actions listen for Git events.

When I push code to a remote repository, the Git server (GitHub/GitLab) fires a **webhook**—an HTTP POST request—to the CI server. The webhook contains a JSON payload detailing the branch, the user, and the new commit hashes.

The CI tool parses this payload, clones the repository at that exact commit, and executes the defined pipeline scripts (linting, compiling, testing). If a test fails, the CI tool uses the Git server's API to mark the commit or Pull Request as 'failed', blocking me from merging broken code."

#### Indepth
Git itself has powerful hooks locally (`.git/hooks/`), but CI tools rely entirely on server-side webhooks. GitHub specifically abstracts this away with GitHub Actions, which essentially behaves like a CI tool embedded directly into the Git hosting service.

---

### 32. What triggers a pipeline?
"Pipelines are highly configurable, but in a standard project, they are triggered by specific Git events.

The most common trigger is **pushing to a PR branch**. Every time I update my Pull Request with a new commit, tests run to ensure I haven't broken anything. 

Another trigger is merging into `main` or pushing a **tag**. When a tag like `v2.0` is pushed, it usually triggers a completely different pipeline—a Continuous Deployment (CD) pipeline that builds the production artifact and deploys it to staging or production servers."

#### Indepth
You can also configure pipelines to run on cron schedules (e.g., nightly security scans) or manual dispatches. Branch logic is crucial: you generally don't want to run expensive e2e deployment tasks on every single feature branch push.

---

### 33. What are protected branches?
"Protected branches are rules enforced by the Git hosting server (like GitHub or GitLab) to protect critical branches like `main` or `production`.

I configure these to strictly prevent force pushing (`git push -f`) and to prevent branch deletion. But more importantly, I use them to enforce workflows. 

For instance, I can require that nobody pushes directly to `main`; all code must come through a Pull Request, must have at least one approval from a code reviewer, and must pass all CI status checks before the 'Merge' button unlocks."

#### Indepth
This is strictly a feature of the hosting provider, not the Git client itself. If you try to bypass a branch protection rule via terminal, the remote Git server will simply reject your push with a HTTP 403 Forbidden or similar pre-receive hook error.

---

### 34. What is a pull request workflow?
"The Pull Request (PR) workflow is the standard for code collaboration.

I fork the repository (or create a feature branch), make my changes locally, and push my branch to the remote server. Then, I open a Pull Request against the target branch (usually `main`).

This creates a dedicated web page for my change. It runs automated CI tests to ensure the build passes. Then, my colleagues leave comments on specific lines of code. I push new commits addressing their feedback. Once the team approves it and tests pass, a maintainer merges the PR, closing the feature."

#### Indepth
The term 'Pull Request' originated from the literal concept of asking the project maintainer to "pull" your changes into their repository from yours. GitLab refers to this exact exact process more accurately as a "Merge Request".

---

### 35. Automating version bumps & tagging via Git.
"Automating versioning removes human error from the release process. 

I usually hook Git into CI. When a PR is merged into `main`, a script reads the commit messages. If I enforce **Conventional Commits** (like `feat: added login` or `fix: header color`), the script automatically determines the next semantic version. `fix` triggers a patch bump (v1.0.1), `feat` triggers a minor bump (v1.1.0).

The CI server automatically modifies `package.json`, commits it, pushes an annotated Git tag (`v1.1.0`), and builds the release. I never tag manually."

#### Indepth
Tools like `semantic-release` or Standard Version handle this heavily. They fundamentally rely on Git's `log` to analyze what has happened since the previous tag, generate a `CHANGELOG.md` autonomously, and push the final artifacts.

---

### 36. Scenario: Squash 5 commits into 1.
"If I have made 5 messy commits on my feature branch (like 'WIP', 'fixing typo', 'actually fixing typo'), I squash them into one clean commit before opening a Pull Request.

I use **interactive rebase**: `git rebase -i HEAD~5`.

This opens a text editor showing my 5 commits from oldest to newest. I leave the word `pick` next to the very first commit, but I change `pick` to `squash` (or `s`) for the other four. When I save and close, Git rolls all their file changes into the first commit and prompts me to write one clean, unified commit message."

#### Indepth
If you have already pushed those 5 commits to your remote feature branch, rewriting history locally means your local branch and remote branch now diverge entirely. You must aggressively overwrite the remote using `git push -f origin my-branch`.

---

### 37. Scenario: Move commits from master → feature branch.
"If I accidentally made 3 local commits directly onto `main` instead of creating a feature branch, I can easily move them.

First, I create the correct branch where I am right now: `git branch my-feature`. (I don't check it out yet). This ensures `my-feature` captures my newly written commits.

Next, I need to roll `main` back 3 commits to its actual unaltered state: `git reset --hard HEAD~3`. 

Finally, I check out my newly created branch: `git checkout my-feature`. The commits are now exactly where they belong, and `main` is clean."

#### Indepth
This works flawlessly because creating a branch in Git doesn't move any files; it just slaps a post-it note on the current commit hash. Hard resetting `main` removes its pointer back three steps, leaving the `my-feature` pointer exactly where you left it.

---

### 38. Scenario: Recover a deleted branch using reflog.
"If I accidentally delete an unmerged feature branch, the `git log` won't show it anymore, but the commits still exist locally.

First, I run `git reflog` to view my absolute history. I scan the log to find the SHA hash where my deleted branch's `HEAD` was pointing (usually labeled as something like 'checkout: moving from my-feature...').

Once I find the hash (e.g., `a1b2c3d`), I simply recreate the branch pointing directly to that hash: `git checkout -b recovered-branch a1b2c3d`. The branch, and all its commits, are instantly restored."

#### Indepth
Git won't garbage-collect unreachable commits immediately; it waits a set period (usually around 30 days for detached objects). As long as you catch the deletion within that window, recovery is instantaneous and 100% complete.

---

### 39. Scenario: Delete a file from Git history but keep locally.
"Occasionally, I accidentally commit an environment file or personal config that I shouldn't have tracked, but I don't want to actually delete from my hard drive.

To remove it from the repository index but leave the physical file intact on my OS, I use the `--cached` flag:
`git rm --cached .env`

After doing this, I must immediately add `.env` to my `.gitignore` file. Then I commit the deletion: `git commit -m "Removed .env from tracking"`. The file is now safely ignored by Git but stays on my computer."

#### Indepth
This removes the file from future snapshots, but if someone views an *old* commit, the `.env` file is still there in the history. If the file contained an actual password, `.cached` is not enough; you must deeply rewrite the history using a tool like `git filter-repo` or BFG to shred it entirely.

---

### 40. Scenario: Rename a branch and update remote.
"If my branch is named poorly and I need to rename it both locally and on GitHub, it requires a sequence of steps.

First, I rename my local branch:
`git branch -m old-name new-name`

Because remote branches cannot simply be 'renamed', I have to push my new local branch and simultaneously instruct the server to delete the old one.
Push new: `git push origin -u new-name`
Delete old: `git push origin --delete old-name`

The branch is now flawlessly transitioned for both me and anyone else checking the remote repository."

#### Indepth
If other developers already pulled the `old-name` branch, they will hit errors when trying to push. They will manually need to pull the `new-name` branch and reset their upstream tracking branch using `git branch -u origin/new-name`.
