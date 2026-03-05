# Git Real-World Scenarios (Product-Based Companies)

This document covers practical, scenario-based Git interview questions frequently asked in product-based companies. Interviewers use these to test your hands-on experience with troubleshooting and recovering from complex Git mistakes.

## 1. You pushed a wrong commit to the production branch. How do you undo it safely?

**Scenario:** You cannot use `git reset --hard` and force push (`git push -f`) because other developers might have already pulled the production branch, and rewriting public history is extremely dangerous.

**Solution:**
1.  Find the hash of the bad commit using `git log`.
2.  Run `git revert <bad-commit-hash>`. This automatically creates a new commit that perfectly negates the changes introduced by the bad commit.
3.  Run `git push` to publish the fix. The public history remains intact and forward-moving.

## 2. You accidentally committed a large 500MB file or a sensitive password. How do you remove it from Git history permanently?

**Scenario:** Simply deleting the file and committing again doesn't solve the problem—the file still exists in the repository's historical `.git/objects` database, bloating the repo size or exposing the password.

**Solution (The modern way):**
Use **`git filter-repo`** (the officially recommended replacement for the older `git filter-branch` command or tools like BFG Repo-Cleaner).
1.  Run `git filter-repo --invert-paths --path <path-to-large-or-sensitive-file>`. This rewrites the entire repository history, completely erasing the file from all commits.
2.  Force push the rewritten history to the remote: `git push origin --force --all`.
3.  *Note:* All team members must then freshly clone the new repository, as their local histories will suddenly be out of sync.

## 3. Your teammate force-pushed and rewrote history. How do you recover your lost commits?

**Scenario:** A teammate accidentally ran `git push -f` and overwrote the remote branch, erasing a week of your merged work.

**Solution:**
1.  As long as you (or someone) previously had those commits locally, the objects still exist on your machine.
2.  Run `git reflog` locally to find the `HEAD` pointer from right before the force push occurred.
3.  Once you find the hash of your old commit (e.g., `HEAD@{5}` or an exact hash), forcefully recreate a branch there: `git branch recovery-branch <hash>`.
4.  You can then merge or carefully force-push this recovered branch back to the remote server to restore the lost work.

## 4. You have 15 messy, work-in-progress commits. How do you clean them into 1 polished commit before a Pull Request?

**Scenario:** You've made dozens of "WIP," "fix," and "typo" commits, and you want to squash them for a clean, professional project history.

**Solution:**
Use an **Interactive Rebase**.
1.  Run `git rebase -i HEAD~15` (to interactively rebase the last 15 commits).
2.  An editor opens listing the commits.
3.  Keep the word `pick` for the very first (top) commit.
4.  Change the word `pick` to `squash` (or just `s`) for the remaining 14 commits.
5.  Save and close the editor. Git will prompt you to write a single, unified commit message for the new squashed commit.

## 5. You created commits on `master` instead of a new `feature` branch. How do you move them?

**Scenario:** You made 3 commits directly on `master` before realizing you were supposed to be on a separate branch. You haven't pushed yet.

**Solution:**
1.  While still on `master`, immediately create your new feature branch (but don't check it out yet): `git branch feature/new-login`.
    *   *Now both `master` and `feature/new-login` point to your latest work.*
2.  Move the `master` branch pointer backwards by 3 commits to where it originally was: `git reset --hard HEAD~3`.
3.  Checkout your feature branch: `git checkout feature/new-login`.
    *   *Now `master` is clean, and your feature branch retains the 3 commits.*

## 6. Two long-running branches diverged heavily. Would you merge or rebase? Why?

**Scenario:** A `feature` branch has been active for two months, while the `main` branch has received 500 new commits. 

**Solution:**
*   **A Merge (`git merge main` into `feature`):** Creating a merge commit is the safest default. If the branches have diverged heavily and conflicts exist across multiple files, merging allows you to resolve all conflicts exactly once during the merge commit. However, it creates a messy graph.
*   **A Rebase (`git rebase main`):** Rebasing rewrites the 2 months of history, individually replaying every single feature commit on top of the new `main`. This keeps history beautifully linear. *However*, if there are heavy conflicts, you will have to manually resolve conflicts for *every single one* of your feature commits, which can be a nightmare.
*   **Conclusion:** For extremely divergent, long-running branches, a **Merge** is significantly safer and easier to execute. Rebasing is strictly better for short-lived, frequently updated local branches.

## 7. How do you resolve conflicts during a rebase?

**Scenario:** During `git rebase main`, the process pauses because it encounters a merge conflict on a specific commit.

**Solution:**
1.  Rebase stops and Git tells you which files have conflicts.
2.  Open the files, find the `<<<<<<<` and `>>>>>>>` markers, and manually resolve the code.
3.  Stage the resolved files: `git add <resolved-file>`.
4.  *Crucially*, do **not** run `git commit`. Instead, tell Git to continue replaying the rest of the commits: `git rebase --continue`.
5.  Repeat the process if subsequent commits also conflict. (You can also abort the entire operation and go back to where you started using `git rebase --abort`).

## 8. Identify which commit introduced a bug (`git bisect`)

**Scenario:** A regression bug exists in production right now (`HEAD`), but it definitely wasn't there in a release tag from a month ago (`v1.5`).

**Solution:**
1.  Setup: `git bisect start`
2.  Mark current state as broken: `git bisect bad`
3.  Mark the old working state: `git bisect good v1.5`
4.  Git checks out the commit halfway between the two. You test the app.
    *   If the bug exists: `git bisect bad`
    *   If it works fine: `git bisect good`
5.  Git halves the list of suspects again. Repeat until Git prints the exact SHA hash of the bad commit. Finally, cleanly exit the tool with `git bisect reset`.

## 9. How do you recover from a detached HEAD state?

**Scenario:** You checked out an old commit hash (`git checkout 8a7c1b4`) to look around, then forgot and made 3 new commits. Those commits belong to no branch and will vanish if you checkout `main`.

**Solution:**
1.  While still in the detached HEAD state (where your 3 commits exist), forcibly create a new branch pointing to this exact spot: `git branch saved-work HEAD`.
2.  Now simply switch to that branch to cement standard behavior: `git checkout saved-work`.
3.  Your 3 commits are now safely attached to the `saved-work` branch.

## 10. How do you move specific commits from one branch to another?

**Scenario:** A `hotfix` branch contains 5 commits, but you only want the 2nd and 4th commits moved into the `release` branch.

**Solution:**
1.  Find the hashes of the 2nd and 4th commits using `git log`.
2.  Checkout your target `release` branch.
3.  Run `git cherry-pick <hash-of-2nd> <hash-of-4th>`. This copies those specific changes and creates new, corresponding commits on your current branch.
