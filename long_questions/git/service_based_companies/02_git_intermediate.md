# Git Intermediate (Service-Based Companies)

This document covers intermediate Git interview questions frequently asked in service-based companies (TCS, Infosys, Wipro, Cognizant, Accenture, Capgemini).

## 11. What is the difference between `merge` and `rebase`?

Both commands are used to integrate changes from one branch into another, but they do it very differently:
*   **`git merge`:** Combines the histories of both branches by creating a new, dedicated "merge commit". It preserves the complete history and chronological order of all commits but can lead to a messy, non-linear history when many people are working simultaneously.
*   **`git rebase`:** Moves or "replays" your entire branch to begin on the tip of the target branch. It creates a perfectly linear, clean project history. However, it rewrites the commit history (creating new commit hashes), which can cause serious problems if used on a shared, public branch.

## 12. What is `HEAD` in Git?

`HEAD` is a reference pointer in Git that tells you exactly where you are "currently looking" inside the repository. Usually, `HEAD` points to the latest commit of the branch you have checked out (e.g., pointing to `master` or `feature/login`). If you check out a specific older commit rather than a branch name, you enter a "detached HEAD" state.

## 13. What are branches? Why are they used?

A branch in Git is an independent line of development. By default, you work in the `master` or `main` branch. 
*   **Why they are used:** Branches allow you to isolate your work (e.g., developing a new feature or fixing a bug) without affecting the stable, main codebase. Once you are certain your code works, you merge your branch back into the main branch.

## 14. How do you delete a branch locally and remotely?

*   **Locally:** First, make sure you are not on the branch you want to delete.
    *   Safe delete (if fully merged): `git branch -d branch_name`
    *   Force delete (even if unmerged): `git branch -D branch_name`
*   **Remotely:**
    *   Push a deletion to the remote server: `git push origin --delete branch_name`

## 15. What is `git stash`? When would you use it?

`git stash` takes the modified tracked files and staged changes in your working directory and saves them on a stack of unfinished changes that you can reapply at any time.
*   **When to use it:** You are working on a feature, and suddenly a critical bug is reported that needs fixing in another branch. Your current work isn't ready to commit. You can `git stash` your changes, switch to the bug-fix branch, fix the issue, return to your original branch, and then run `git stash pop` to restore your unfinished work.

## 16. What is a fast-forward merge?

A fast-forward merge happens when there is a straight path from the current branch's tip to the target branch. Instead of creating a special "merge commit," Git simply moves (fast-forwards) the branch pointer forward to point to the incoming commit. It looks as if all your changes happened sequentially in a single line.

## 17. Explain the difference between `reset` and `revert`.

Both undo changes, but differently:
*   **`git revert <commit>`:** Creates a *new* commit that correctly undoes the exact changes made in the specified earlier commit. It does not alter history, making it completely **safe to use on shared branches**.
*   **`git reset <commit>`:** Moves the `HEAD` and branch pointer backward to a specific past commit, effectively rewriting history. 
    *   `--soft`: keeps changes mapped in the staging area.
    *   `--mixed` (default): keeps changes mapped in the working directory.
    *   `--hard`: entirely deletes the changes and history after that commit. Extremely dangerous on shared branches.

## 18. What is the purpose of `git cherry-pick`?

`git cherry-pick <commit-hash>` enables you to pick an arbitrary, specific commit from one branch and apply it to the branch you are currently on. 
*   **Purpose:** Imagine a coworker created a feature branch with three commits, but you only strictly need the single commit that contains a specific utility function. You can "cherry-pick" just that one commit to your branch without merging their entire feature.

## 19. What is a detached HEAD state?

A "detached HEAD" state occurs when `HEAD` points directly to a specific commit (via its specific hash) instead of pointing to a branch name (like `main`). If you create new commits in a detached HEAD state, they do not belong to any named branch. If you switch to another branch later without creating a new branch to save them, those new commits will be orphaned and eventually garbaged collected by Git.

## 20. How do you revert a commit that has already been pushed?

You should never use `git reset` if the commit has been pushed and shared with other developers, as it rewrites history.
*   **The solution:** Use `git revert <commit-hash>`. This will generate a brand new commit that inversely undoes changes from the bad commit, leaving an intact public history. Then, push the new revert commit to the remote.
