# Git Advanced Troubleshooting & Reporting (Product-Based Companies)

This document groups heavy analytical and debugging scenarios where developers must dig through Git's internals to fix catastrophic failures or generate intricate reports.

## 1. How do you recover from a corrupted `.git` directory?

Repository corruption happens if a machine crashes while writing a commit, or if a failing hard drive corrupts a packfile object.
*   **Identify Corruption:** Run `git fsck --full`. This performs an integrity check traversing every object and link in the database. It outputs orphaned nodes or corrupted SHA-1 hashes.
*   **The Fix:** 
    1. If a specific loose object is corrupted, you might be able to manually delete that object from `.git/objects/` and then run `git pull` or `git fetch` to re-download the uncorrupted version from the remote server.
    2. If the `index` file is corrupted (it says "bad index file"), it only means the staging area is dead. You simply delete `.git/index` and run `git reset` to rebuild it from the current `HEAD` commit.
    3. *Nuclear Option:* If the repository is fully unsalvageable and heavily corrupted beyond repair, rename the broken folder, re-clone the clean project from the remote server, and manually copy all your modified working directory files from the broken folder into the clean clone.

## 2. How to resolve dangling commits or unreachable commits?

*   **Dangling Commit:** A valid commit object exists in the `.git/objects/` database, but absolutely no branch, tag, or reference points to it. This happens normally when you `git commit --amend`, delete branches, or aggressively `git reset`.
*   **Dangling Blob:** A file was staged (`git add`), but you physically deleted the file or abandoned the changes before ever creating an actual commit linking to it.
*   **Resolution:** 
    1. Usually, do nothing. It's a non-issue. Git's internal garbage collector (`git gc`) will automatically permanently delete unreachable objects older than 2 weeks on its next background run.
    2. If you are desperately low on disk space, you can immediately obliterate them by running `git gc --prune=now`. 

## 3. How to investigate why a file mysteriously disappeared after checkout?

**Scenario:** You check out the `main` branch, and `config.yml` simply vanishes from the file system.
*   **Is it ignored?** Check `.gitignore`. Someone might have added it.
*   **Who deleted it?** Run `git log --all --full-history -- <path-to-file>`. This command searches every branch in history to find all commits that ever touched that file, ultimately showing you the exact commit name and author where the file was deleted. You can then `git restore --source=<commit-before-deletion> -- <file>` to resurrect it.

## 4. How do you check who modified a file last? (`git blame`)

`git blame <filename>` is the standard analytical tool to track down exactly which developer last modified each specific line of code in a file, displaying the commit hash alongside the line.
*   **Advanced usage:** If a massive indentation fix or formatting change (like Prettier) touched every line and ruined `git blame`, you can pass the `-w` flag (`git blame -w <file>`) to tell Git to completely ignore whitespace modifications and find the actual, original developer who wrote the logic.
*   You can also specify line ranges: `git blame -L 50,75 <file>`.

## 5. How do you generate commit statistics and analyze code churn?

Engineering managers often want to track developer velocity or identify highly fragile ("churning") files that break frequently.
*   **Basic Stats:** `git shortlog -sn` neatly prints a leaderboard of all contributors ranked by their total number of commits.
*   **Churn Analysis:** You can find the most frequently modified files (often a sign of tech debt) by piping Git logs through bash utilities:
    `git log --name-only --oneline | grep -v ' ' | sort | uniq -c | sort -nr | head -10`
    This prints the top 10 files with the highest number of historical commits attached to them.
*   **Line Stats:** `git log --author="Dhruv Shah" --stat` prints the total number of lines inserted and deleted across all your commits.
