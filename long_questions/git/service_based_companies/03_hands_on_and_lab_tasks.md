# Git Hands-On Practical & Lab Tasks (General / Service-Based Companies)

This document contains practical, hands-on interview tasks often given as 10-15 minute paired programming exercises to prove you actually know the command line, not just the theory.

## Task 1: Create a repo with 3 branches, simulate a merge conflict, and resolve it.

**The Scenario:** The interviewer wants to see your basic workflow and conflict resolution skills without a GUI.
**The Execution:**
```bash
# 1. Initialize repository
mkdir git-lab && cd git-lab
git init
echo "Line 1" > file.txt
git add file.txt
git commit -m "Initial commit on main"
# Note: You are now on 'main' by default.

# 2. Create Branch A and modify the file
git checkout -b feature-a
echo "Modified by Feature A" > file.txt
git commit -am "Feature A modification"

# 3. Create Branch B from main (diverge) and modify the exact same line
git checkout main
git checkout -b feature-b
echo "Modified by Feature B instead" > file.txt
git commit -am "Feature B modification"

# 4. Simulate the conflict by merging both into main
git checkout main
git merge feature-a  # Merges cleanly (Fast-forward)
git merge feature-b  # CONFLICT (content): Merge conflict in file.txt

# 5. Resolve
# Open file.txt in your editor (nano/vim/VSCode)
# You will see <<<<<<< HEAD ... ======= ... >>>>>>> feature-b
# Delete the markers and choose the correct line (e.g., keep Feature B's line).
git add file.txt
git commit -m "Resolve merge conflict between feature-a and feature-b"
```

## Task 2: Squash 5 messy commits into 1 clean commit.

**The Scenario:** You have `WIP 1`, `typo fix`, `actually working now`, `more fixes`, and `final ready`. The interviewer wants you to make it 1 commit before merging.
**The Execution:**
```bash
# Assuming you have 5 dirty commits on your current branch since you branched off main

# Start an interactive rebase referencing the last 5 commits
git rebase -i HEAD~5

# Git opens your default text editor (usually Vim or Nano).
# It lists the 5 commits from oldest (top) to newest (bottom).
# It will look like:
# pick 1234567 WIP 1
# pick 89abcde typo fix
# pick f012345 actually working now
# pick 6789abc more fixes
# pick def0123 final ready

# Change the word 'pick' to 'squash' (or 's') for the bottom 4 commits.
# Leave the top commit as 'pick'.
# pick 1234567 WIP 1
# squash 89abcde typo fix
# squash f012345 actually working now
# squash 6789abc more fixes
# squash def0123 final ready

# Save and exit the editor.
# Git opens a SECOND editor window prompting you to create a brand new unifying commit message.
# Delete the ugly WIP messages and write: "feat: implemented user login completely"
# Save and exit. The 5 commits are perfectly squashed into one linear commit.
```

## Task 3: Move specific commits from `master` to a `feature` branch.

**The Scenario:** You made 2 awesome commits directly on `master` but realized you were supposed to put them in `feature/new-login`.
**The Execution:**
```bash
# 1. You are currently on master and it has the 2 new commits.
# Note the hash of the latest commit: git log --oneline

# 2. Create the feature branch pointing to this exact spot, but DON'T switch to it.
git branch feature/new-login

# 3. Forcibly rewind master backward by 2 commits, permanently removing them from master.
git reset --hard HEAD~2

# 4. Now checkout your feature branch. It still points forward where the 2 commits exist.
git checkout feature/new-login
```

## Task 4: Recover a deleted commit using `reflog`.

**The Scenario:** You accidentally ran `git reset --hard HEAD~1` and wiped out an hour of work. The commit is gone from `git log`. Prove you can revive it.
**The Execution:**
```bash
# 1. View Git's secret chronological movement journal:
git reflog

# Output looks something like:
# 3b4c5d6 (HEAD -> main) HEAD@{0}: reset: moving to HEAD~1
# a1b2c3d HEAD@{1}: commit: add complex database query   <-- This is the deleted commit!

# 2. Create a temporary branch to resurrect the lost code at that exact hash.
git branch recovered-code a1b2c3d

# 3. Checkout the branch to verify the code is back safely.
git checkout recovered-code
```

## Task 5: Identify a bug using `git bisect` (Simulated).

**The Scenario:** Production (`HEAD`) is broken. The `v1.0` tag from a month ago works perfectly. Find the exact commit hash that broke it.
**The Execution:**
```bash
# Start the binary search tool
git bisect start

# Tell Git that the current commit is broken
git bisect bad

# Tell Git that the v1.0 tag (or an older #hash) was good
git bisect good v1.0

# Git calculates the halfway commit and automatically checks it out.
# > "Bisecting: 50 revisions left to test after this (roughly 6 steps)"

# You manually run your app test (e.g., `npm test` or `mvn clean install`).
# If the test FAILS:
git bisect bad

# If the test PASSES:
git bisect good

# Git halves the remaining commits again and checks out the new middle point.
# You repeat Good/Bad until Git outputs exactly:
# > "a1b2c3d4 is the first bad commit"

# Exit the tool and return to where you started
git bisect reset
```
