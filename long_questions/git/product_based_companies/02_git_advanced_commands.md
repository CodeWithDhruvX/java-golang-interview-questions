# Git Advanced Commands (Product-Based Companies)

This document covers deep dives into advanced Git commands frequently asked in product-based companies (Google, Microsoft, Amazon, Atlassian, etc.).

## 1. Explain the difference between `git reset --soft`, `--mixed`, `--hard`

`git reset` rewinds history by moving HEAD and the current branch pointer backward. The modes dictate what happens to the changes that were in the commits you are rewinding through:
*   **`--soft`:** Moves HEAD backward but keeps the Working Directory and the Staging Area unchanged. All changes from the rewound commits are still staged and ready to be committed again. (Useful for squashing commits).
*   **`--mixed` (Default):** Moves HEAD and clears the Staging Area, but leaves the Working Directory unchanged. Changes are "unstaged" but still exist in your code.
*   **`--hard`:** Extremely destructive. Moves HEAD, clears the Staging Area, and completely overwrites the Working Directory to perfectly match the target commit. All uncommitted changes are permanently lost.

## 2. Difference between `git revert`, `git reset`, and `git restore`

*   **`git reset <commit>`:** Rewrites history by forcibly moving HEAD backward. Never use this on a shared branch because it deletes public history.
*   **`git revert <commit>`:** Safely undoes a previous commit by creating a brand-new commit that reverses the exact changes. It preserves public history, making it the only safe way to undo shared commits.
*   **`git restore <file>`:** Replaces the version of a file in the Working Directory or Staging Area with the version from a specific commit (or HEAD). It does explicitly *not* move HEAD or rewrite commit history. It is purely for discarding uncommitted changes to files.

## 3. When would you use: `git rebase -i`, `git cherry-pick`, `git stash pop`, `git stash apply`?

*   **`git rebase -i <commit>` (Interactive Rebase):** Used to clean up your local messy commit history before pushing. You can rewrite, squash, split, reorder, or delete entire commits.
*   **`git cherry-pick <commit>`:** Used when you need a specific commit (e.g., a hotfix) from another branch but *do not* want to merge the entire branch.
*   **`git stash pop`:** Used to restore temporarily saved (stashed) changes and simultaneously remove that stash from the stash list. 
*   **`git stash apply`:** Used to restore stashed changes but keep them in the stash list (useful if you want to apply the same stash to multiple different branches).

## 4. What does `git reflog` store?

`git reflog` (Reference Log) is Git's safety net. It strictly records every single time the `HEAD` pointer moves in your local repository. 
*   Even if you perform a disastrous `git reset --hard` and rewrite history, the reflog remembers exactly what commit hash HEAD pointed to before the reset. 
*   You can recover "lost" commits by checking the reflog and resetting back to the old hash. It tracks changes for 90 days by default.

## 5. Difference between `origin` and `upstream`?

*   **`origin`:** The default name Git assigns to the remote repository you cloned *from*. Typically, this is your own fork of a project on GitHub.
*   **`upstream`:** The conventional name developers give to the original, central repository that they forked the project from. You fetch from `upstream` to get the latest changes made by other contributors, and you push to `origin` (your fork).

## 6. What is `git remote prune`?

When a team merges a pull request on GitHub and deletes the remote branch, your local Git configuration still remembers that the remote branch exists in `refs/remotes/origin/`.
Running `git remote prune origin` (or `git fetch --prune`) cleans up your local tracking branches by deleting any references to branches that have been deleted on the specified remote server.

## 7. What is `git clean -fd`?

`git clean` removes *untracked* files from the Working Directory (files Git doesn't know about yet).
*   `-f` (force) is required by default to actually delete files (as a safety measure).
*   `-d` tells Git to also remove untracked directories.
It is often paired with `git reset --hard` to completely wipe a working area clean and revert everything exactly to a specific commit.

## 8. What is `git blame` used for?

`git blame <filename>` is an analytics tool that displays the contents of a file along with metadata for each line, showing the exact commit hash and the author who last modified that specific line. It is primarily used to track down the person who introduced a bug or to understand the context of a tricky line of code.

## 9. What is `git bisect` and how does it work?

`git bisect` is a binary search algorithm for finding exactly which commit introduced a bug in a large codebase.
*   You start by telling Git a "bad" commit (e.g., HEAD) where the bug is present, and a older "good" commit where you know the code worked perfectly.
*   Git calculates the middle commit between the two and checks it out.
*   You test the code and type `git bisect good` or `git bisect bad`.
*   Git halves the search space again until it pinpoints the exact offending commit in O(log N) steps.

## 10. What is `git submodule`?

A Python, C++, or Node project might rely on another independent Git repository (e.g., a shared C library). 
A `git submodule` allows you to embed one Git repository completely inside another as a subdirectory, while keeping their commit histories completely separate and independent. It points to a specific commit of the embedded repository, freezing the dependency version.
