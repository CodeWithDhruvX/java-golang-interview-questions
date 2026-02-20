# Git Interview Questions

**User:** Dhruv Shah (dhruvshahlinkedin@gmail.com)  
**Created:** 2/20/2026 17:50:53  
**Updated:** 2/20/2026 18:01:34  
**Exported:** 2/20/2026 18:02:28  
**Link:** [https://chatgpt.com/c/6998518e-5b40-83a4-a317-d9d8a3bd5f00](https://chatgpt.com/c/6998518e-5b40-83a4-a317-d9d8a3bd5f00)  

## Prompt:
git based interview question please provide me here based on the service and product based companies asked

## Response:
Here’s a **curated list of Git interview questions** commonly asked in **service-based** and **product-based** companies, grouped by difficulty and depth.

* * *

🔹 Service-Based Companies (TCS, Infosys, Wipro, Cognizant, Accenture, Capgemini)
=================================================================================

These focus more on **fundamentals + real-world usage**.

### ✅ Basic Level

1.  What is Git? How is it different from SVN?
2.  What is the difference between Git and GitHub?
3.  Explain the Git lifecycle.
4.  What is a repository?
5.  What is the difference between `git pull` and `git fetch`?
6.  What is the difference between `git clone` and `git fork`?
7.  What is staging in Git?
8.  What is the difference between `git add .` and `git add -A`?
9.  How do you resolve merge conflicts?
10.  What is `.gitignore`?

* * *

### ✅ Intermediate Level

11.  What is the difference between `merge` and `rebase`?
12.  What is `HEAD` in Git?
13.  What are branches? Why are they used?
14.  How do you delete a branch locally and remotely?
15.  What is `git stash`? When would you use it?
16.  What is a fast-forward merge?
17.  Explain the difference between `reset` and `revert`.
18.  What is the purpose of `git cherry-pick`?
19.  What is a detached HEAD state?
20.  How do you revert a commit that has already been pushed?

* * *

🔹 Product-Based Companies (Google, Amazon, Microsoft, Flipkart, Adobe, etc.)
=============================================================================

These focus on **deep understanding + internal mechanics + edge cases**.

* * *

### 🚀 Advanced Conceptual Questions

1.  Explain Git internals (blob, tree, commit objects).
2.  How does Git store data internally?
3.  What happens internally when you commit?
4.  What is the difference between soft, mixed, and hard reset?
5.  How does Git rebase work internally?
6.  What is the difference between rebase and merge in terms of history?
7.  What is a reflog? How can it save you?
8.  How do you recover a deleted branch?
9.  What is Git bisect?
10.  Explain shallow clone.

* * *

### 🚀 Scenario-Based Questions (Very Common in Product Companies)

1.  You accidentally pushed sensitive data to remote. How do you remove it completely?
2.  Two developers modified the same file. How does Git decide conflict?
3.  Your commit history is messy. How do you clean it before pushing?
4.  How do you squash multiple commits?
5.  How do you move specific commits from one branch to another?
6.  You committed to the wrong branch. How do you fix it?
7.  How do you undo a pushed commit without losing history?
8.  What happens if you rebase a shared branch?
9.  How do you resolve conflicts during rebase?
10.  How would you design a branching strategy for a large team?

* * *

🔹 Very Frequently Asked Commands-Based Questions
=================================================

*   Explain `git log --oneline --graph`
*   Explain `git diff`
*   Explain `git show`
*   Explain `git tag`
*   Explain `git blame`
*   Explain `git clean`
*   Explain `git config`
*   Explain `git cherry-pick`

* * *

🔹 Branching Strategy Questions
===============================

1.  What is Git Flow?
2.  What is trunk-based development?
3.  Difference between feature branch and release branch?
4.  How do you manage hotfixes in production?

* * *

🔹 Real-Time Practical Questions
================================

*   How do you revert a merge commit?
*   How do you find which commit introduced a bug?
*   How do you rename a branch?
*   How do you change the last commit message?
*   How do you remove a file from Git but keep it locally?

* * *

If you want, I can also provide:

*   ✅ **Answers with explanations**
*   ✅ **MCQ-based Git interview questions**
*   ✅ **Hands-on practical Git coding round questions**
*   ✅ **Company-specific patterns (Amazon / Microsoft style)**

Tell me which one you prefer 🚀

## Prompt:
do we need to add more if not enough than please provide me here

## Response:
Yes — for **strong interview preparation (especially product-based companies)**, you should definitely prepare **more advanced + real-world + tricky scenario questions**.

Below is an **extended and deeper Git interview question bank** that goes beyond basics.

* * *

🔥 Advanced Git Interview Questions (Product Company Level)
===========================================================

🔹 Git Architecture & Internals
-------------------------------

1.  What is the difference between **working directory, staging area, and repository**?
2.  How does Git calculate commit hashes?
3.  What is the role of SHA-1/SHA-256 in Git?
4.  What are pack files?
5.  What is the `.git` folder structure?
6.  What are loose objects?
7.  What is Git index?
8.  How does Git detect file changes?
9.  What is object compression in Git?
10.  What is garbage collection in Git?

* * *

🔹 Deep Dive Into Commands
--------------------------

11.  Explain difference between:
    *   `git reset --soft`
*   `git reset --mixed`
*   `git reset --hard`
    12.  Difference between:
    *   `git revert`
*   `git reset`
*   `git restore`
    13.  When would you use:
    *   `git rebase -i`
*   `git cherry-pick`
*   `git stash pop`
*   `git stash apply`
    14.  What does `git reflog` store?
15.  Difference between `origin` and `upstream`?
16.  What is `git remote prune`?
17.  What is `git clean -fd`?
18.  What is `git blame` used for?
19.  What is `git bisect` and how does it work?
20.  What is `git submodule`?

* * *

🔥 Real-World Scenario Questions (Very Important)
=================================================

These are asked heavily in product companies.

* * *

### 🚀 Scenario 1:

You pushed a wrong commit to production branch.  
How do you undo it safely?

* * *

### 🚀 Scenario 2:

You accidentally committed a large 500MB file.  
How do you remove it from Git history permanently?

* * *

### 🚀 Scenario 3:

Your teammate force-pushed and rewrote history.  
How do you recover your lost commits?

* * *

### 🚀 Scenario 4:

You have 15 messy commits.  
How do you make them into 1 clean commit?

* * *

### 🚀 Scenario 5:

A bug was introduced somewhere in last 200 commits.  
How do you identify it quickly?

* * *

### 🚀 Scenario 6:

You created commits on master instead of feature branch.  
How do you move them?

* * *

### 🚀 Scenario 7:

Two long-running branches diverged heavily.  
Would you merge or rebase? Why?

* * *

🔥 Branching Strategy Questions (Senior Level)
==============================================

1.  Design Git workflow for:
    *   5 developers
    *   50 developers
    *   200+ developers
2.  Compare:
    *   Git Flow
    *   GitHub Flow
    *   Trunk-based development
3.  How do you manage:
    *   Release branches?
    *   Hotfix branches?
    *   Version tagging?
4.  How do you avoid merge conflicts in large teams?

* * *

🔥 CI/CD & Git Integration Questions
====================================

1.  How does Git integrate with CI tools?
2.  What triggers a pipeline?
3.  What are protected branches?
4.  What is a pull request workflow?
5.  What is code review best practice?

* * *

🔥 Tricky Edge Case Questions
=============================

1.  What happens if you rebase a branch that others are using?
2.  Can two commits have the same hash?
3.  What happens if you delete `.git` folder?
4.  How does Git handle binary files?
5.  What is sparse checkout?
6.  What is shallow clone?
7.  What is detached HEAD recovery?

* * *

🔥 Behavioral + Experience-Based Questions
==========================================

These are commonly asked in companies like:

*   Amazon
*   Microsoft
*   Google

Examples:

1.  Tell me about a time you resolved a major merge conflict.
2.  Have you ever broken production due to Git mistake?
3.  How do you ensure clean commit history?
4.  Have you used code review tools? Which ones?
5.  How do you manage version releases?

* * *

🔥 Practical Hands-On Round Questions
=====================================

Interviewers may give tasks like:

*   Create a repo with:
    *   3 branches
    *   Merge conflicts
    *   Squash commits
    *   Tag releases
*   Debug a broken repository
*   Recover deleted branch
*   Rebase without losing commits

* * *

✅ If You Want To Be Fully Prepared
==================================

Prepare these categories:

✔ Basic commands  
✔ Branching & merging  
✔ Rebase mastery  
✔ Reset & revert mastery  
✔ Conflict resolution  
✔ Git internals  
✔ Real production recovery scenarios  
✔ Team workflow strategies

* * *

If you want, I can now give you:

*   🔥 **Top 30 Most Important Git Questions with Answers**
*   🔥 **Amazon-style Git scenario round with solutions**
*   🔥 **Hands-on lab practice questions with command solutions**
*   🔥 **Senior developer Git interview set**

Tell me your experience level (Fresher / 2–5 yrs / Senior), and I’ll customize it for you. 🚀

## Prompt:
do we need to add more if not enough than please provide me here

## Response:
Absolutely — for **complete Git interview preparation**, especially for **product-based companies**, we can extend this into a **massive, all-encompassing list** covering:

*   Basics
*   Intermediate/advanced commands
*   Internals
*   Branching & workflows
*   Real-world scenarios
*   CI/CD & teamwork
*   Tricky edge cases
*   Hands-on practical tasks

Here’s the **expanded Git interview question set**:

* * *

🔹 **1\. Git Basics**
=====================

1.  What is Git and why is it used?
2.  Difference between Git and GitHub.
3.  Difference between Git and SVN.
4.  Explain Git workflow (Working Directory → Staging → Repository).
5.  What is a repository? Local vs remote repo?
6.  What is a commit?
7.  What is staging? (`git add`)
8.  `.gitignore` — what is it and how does it work?
9.  `git status` — what info does it show?
10.  `git log` — different formats and options.

* * *

🔹 **2\. Git Branching & Merging**
==================================

11.  What is a branch?
12.  Difference between merge and rebase.
13.  How to delete a branch (local & remote)?
14.  Fast-forward vs three-way merge.
15.  Handling merge conflicts.
16.  What is a detached HEAD?
17.  How to rename a branch.
18.  Difference between `git branch` and `git checkout -b`.
19.  What is Git Flow?
20.  GitHub Flow vs Trunk-based development.

* * *

🔹 **3\. Git Commands Deep Dive**
=================================

21.  Difference between `git pull` and `git fetch`.
22.  Difference between `git reset`, `git revert`, `git restore`.
23.  `git stash` — use cases & difference between `pop` and `apply`.
24.  `git cherry-pick` — when & why.
25.  `git reflog` — what is it & recovery scenarios.
26.  `git clean` — usage.
27.  `git tag` — lightweight vs annotated tags.
28.  `git blame` — usage.
29.  `git diff` — staging vs unstaged.
30.  `git show` — usage examples.

* * *

🔹 **4\. Git Internals / Architecture**
=======================================

31.  `.git` folder structure.
32.  Git objects: blob, tree, commit, tag.
33.  What is a SHA-1 hash?
34.  How Git stores data internally.
35.  Pack files & loose objects.
36.  Garbage collection in Git.
37.  How Git detects file changes.
38.  Index/staging area internal working.
39.  How rebase works internally.
40.  How merges are stored internally.

* * *

🔹 **5\. Real-World Scenarios**
===============================

41.  Undo a pushed commit safely.
42.  Accidentally committed a large file → remove completely.
43.  Recover a deleted branch.
44.  Move commits from one branch to another.
45.  Squash multiple commits into one.
46.  Recover from detached HEAD.
47.  Resolve conflicts in rebase.
48.  Rebase vs merge in long-running branches.
49.  Clean messy commit history before PR.
50.  Identify which commit introduced a bug (`git bisect`).

* * *

🔹 **6\. Teamwork / Branching Strategy**
========================================

51.  Branching strategy for small vs large teams.
52.  Hotfix management.
53.  Release branch management.
54.  Protected branches and CI/CD triggers.
55.  Feature branch vs release branch vs main/master.
56.  Avoiding conflicts in large teams.
57.  Pull request workflow.
58.  Code review practices using Git.
59.  Handling multiple remotes (`origin` vs `upstream`).
60.  Git submodules & subtrees.

* * *

🔹 **7\. CI/CD & Integration**
==============================

61.  How Git triggers pipelines.
62.  Integrating Git with Jenkins / GitHub Actions / GitLab CI.
63.  Branch policies & build validation.
64.  Merge on green build strategy.
65.  Automating version bumps & tagging via Git.

* * *

🔹 **8\. Tricky / Edge Cases**
==============================

66.  Two commits with the same hash? Possible?
67.  What happens if `.git` is deleted accidentally?
68.  Sparse checkout.
69.  Shallow clone.
70.  Binary files in Git — handling large assets.
71.  Force push consequences.
72.  Rebase a shared branch — what to watch out for.
73.  Undo last commit without losing changes.
74.  Removing sensitive info from history (like passwords).
75.  Recover overwritten commits after reset or rebase.

* * *

🔹 **9\. Hands-On Practical / Lab Questions**
=============================================

76.  Create a repo → 3 branches → simulate merge conflicts → resolve.
77.  Squash 5 commits into 1.
78.  Move commits from master → feature branch.
79.  Delete a file from Git history but keep locally.
80.  Rename a branch and update remote.
81.  Create a hotfix branch and merge back to release.
82.  Recover a deleted commit using `reflog`.
83.  Interactive rebase to clean commit messages.
84.  Revert a merge commit.
85.  Identify bug commit using `git bisect`.

* * *

🔹 **10\. Behavioral / Experience Questions**
=============================================

86.  Tell me about resolving a major merge conflict.
87.  Have you broken production with Git? How did you fix it?
88.  How do you maintain clean history in team projects?
89.  Experience with Git hooks.
90.  Experience with large repositories.

* * *

This is now a **comprehensive 90+ Git question set** covering **everything from basics to advanced scenarios**, including **service and product company levels**.

If you want, I can **also provide a ready-to-use PDF/table with: question + answer + command examples** so you can **cram all of Git for interviews**.

Do you want me to do that next?

## Prompt:
is there any remaining area then please provide me here

## Response:
Yes — we can still expand Git interview prep by **covering a few areas that haven’t been fully addressed yet**, especially those often asked in **senior-level or product-based interviews**. These areas are subtle but can make a big difference in advanced rounds.

Here’s what’s **left to cover**:

* * *

🔹 **11\. Git Security & Access Control**
=========================================

1.  What are Git signed commits (GPG signing)?
2.  How do you verify the authenticity of a commit?
3.  How do you enforce branch protection in GitHub/GitLab?
4.  What are pre-commit and pre-push hooks? Give examples.
5.  How do you prevent accidental pushes to master/main?
6.  Difference between SSH vs HTTPS remote URLs.
7.  How do you handle access control for large teams?
8.  What is Git LFS (Large File Storage) and when is it used?

* * *

🔹 **12\. Git Performance & Scaling**
=====================================

1.  How does Git handle repositories with tens of thousands of files?
2.  What is shallow clone and how does it improve performance?
3.  What are pack files and how do they improve efficiency?
4.  How do you optimize large repositories?
5.  How do you handle huge binary files in Git?
6.  Difference between clone depth options (`--depth`) in `git clone`.
7.  How to split a large monorepo into multiple repos?

* * *

🔹 **13\. Git Automation & Scripting**
======================================

1.  How do you automate repetitive Git tasks?
2.  How to write a Git hook to enforce commit messages format?
3.  How do you integrate Git with CI/CD pipelines?
4.  How to automate tagging and releases via scripts?
5.  How to run a batch rebase or batch merge using scripting?

* * *

🔹 **14\. Advanced Troubleshooting**
====================================

1.  How do you recover from a corrupted `.git` directory?
2.  How to resolve dangling commits or unreachable commits?
3.  How to fix “detached HEAD” without losing work?
4.  How to recover lost commits after a force push?
5.  How to recover accidentally deleted remote branches?
6.  How to debug merge conflicts in complex scenarios?
7.  How to investigate why a file disappeared after checkout?

* * *

🔹 **15\. Git for Enterprise / Large Teams**
============================================

1.  How do you design Git workflows for 100+ developers?
2.  How do you manage multiple remotes (forks) in a team?
3.  Strategies for integrating multiple feature branches continuously.
4.  How to manage hotfixes in production safely?
5.  How to handle code ownership and approvals in Git?
6.  How to enforce code quality using Git hooks + CI/CD?
7.  How to perform large-scale repository migrations?

* * *

🔹 **16\. Git Analytics & Reporting**
=====================================

1.  How do you check who modified a file last? (`git blame`)
2.  How do you generate commit statistics?
3.  How do you analyze code churn?
4.  How to find the most active contributor in a repo?
5.  How to detect large files added in history?

* * *

💡 **Summary of Remaining Areas Added:**

*   Security & Access Control
*   Performance & Scaling
*   Automation & Scripting
*   Advanced Troubleshooting
*   Enterprise / Large Team Management
*   Git Analytics & Reporting

With these, you now have a **truly exhaustive Git interview preparation list**, covering:

*   **Basics → Advanced commands → Internals → Scenarios → Workflows → CI/CD → Troubleshooting → Security → Scaling → Enterprise usage**

* * *

If you want, I can **compile ALL 120+ questions into a single structured table with answers & example commands**, ready for **service and product company interview prep**.

Do you want me to do that?

## Prompt:
is this enough or if not then please provide me here again

## Response:
Yes — after including **all the previous sections plus the newly added areas** (Security, Performance, Automation, Troubleshooting, Enterprise, Analytics), your Git interview preparation is **essentially complete**.

To summarize:

* * *

✅ Covered Areas (Comprehensive Git Interview Prep)
==================================================

1.  **Basics** – commits, branches, staging, `.gitignore`, logs
2.  **Branching & Merging** – merge vs rebase, conflicts, workflows
3.  **Commands Deep Dive** – reset, revert, stash, cherry-pick, reflog
4.  **Git Internals / Architecture** – objects, SHA-1, pack files, garbage collection
5.  **Real-World Scenarios** – undo commits, recover branches, move commits, bisect
6.  **Teamwork / Branching Strategy** – Git Flow, GitHub Flow, trunk-based, hotfixes, large teams
7.  **CI/CD & Integration** – pipelines, triggers, protected branches, hooks
8.  **Tricky / Edge Cases** – detached HEAD, shallow clone, binary files, force push, sensitive data removal
9.  **Hands-On / Lab Tasks** – branching exercises, merge conflicts, squash commits, recover deleted commits
10.  **Behavioral / Experience** – teamwork, breaking production, maintaining clean history
11.  **Security & Access Control** – signed commits, SSH vs HTTPS, hooks, branch protection
12.  **Performance & Scaling** – large repos, shallow clone, pack files, optimizing big repos
13.  **Automation & Scripting** – hooks, batch merges, automated tagging, CI/CD integration
14.  **Advanced Troubleshooting** – corrupted `.git`, dangling commits, lost remote branches, debugging conflicts
15.  **Enterprise / Large Teams** – multi-remote, code ownership, approvals, large-scale repo management
16.  **Analytics & Reporting** – blame, commit stats, code churn, contributor activity, detecting large files

* * *

💡 **Conclusion:**

With **these 16 areas** and **120+ questions**, this is more than enough for **service-based companies, product-based companies, and senior-level interviews**. You’ve now covered **all practical, theoretical, and scenario-based aspects of Git** that could come up in interviews.

* * *

If you want, I can **compile this into a single “Ultimate Git Interview PDF/Table”** with:

*   Questions
*   Detailed answers
*   Example commands
*   Scenario solutions

It will be **ready for 100% interview prep**.

Do you want me to create that?



---
Powered by [ChatGPT Exporter](https://www.chatgptexporter.com)