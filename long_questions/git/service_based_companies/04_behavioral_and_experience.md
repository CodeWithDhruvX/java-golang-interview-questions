# Git Behavioral & Experience Questions (General / Service-Based Companies)

Behavioral questions assess how you handle team dynamics, production failures, and collaboration issues using version control over your past work experience.

## 1. Tell me about a time you resolved a major merge conflict.

**Good Answer Structure (STAR Method):**
*   **Situation:** "We were finalizing a major release, and two teams had been working on different sub-features of the massive `PaymentService.java` class for three weeks on separate branches."
*   **Task:** "When attempting to integrate into the `release` branch, Git threw massive conflict markers across 500 lines of code. It was my responsibility to safely merge the code without breaking either team's logic before the code freeze."
*   **Action:** "Instead of blindly accepting chunks, I immediately set up a 30-minute sync call with the lead developer from the other team. We opened a visual diff tool (like VSCode's 3-way merge editor). I explained my logic changes, and they explained theirs. We manually refactored the conflicting overlapping methods into two separate utility functions during the merge itself to satisfy both requirements."
*   **Result:** "The conflict was documented via a detailed commit message, the code compiled, CI tests passed perfectly, and we merged the PR two hours before the deadline. It taught me the importance of continuous integration over long-lived feature branches."

## 2. Have you ever broken production due to a Git mistake? How did you fix it?

**Good Answer Structure:**
*   **Honesty:** Admit a real mistake. "Yes, early in my career, I accidentally performed a `git push -f` to a shared development branch."
*   **The Problem:** "It wiped out a colleague's commits over the last two days, causing the nightly build to fail and panic in the team."
*   **The Fix:** "I immediately took ownership of the mistake in the team chat. I stopped anyone else from pushing or pulling. I asked the colleague if they still had their branch locally. Fortunately, they did. I guided them to push their local branch back up to the server, which restored the missing commits. To ensure it never happened again, I took the initiative to configure Branch Protection rules in GitLab, strictly disabling `-f` pushes on shared branches for the whole organization."
*   **Learning:** "It taught me that tooling should prevent human error, rather than relying on humans to never make errors."

## 3. How do you ensure a clean commit history in team projects?

**Good Answer Structure:**
*   **Standardization:** "I advocate for strict adherence to Conventional Commits (e.g., `feat:`, `fix:`, `docs:`) so the history reads like a proper changelog."
*   **Squashing:** "I enforce 'Squash and Merge' policies on Pull Requests. Even if a developer makes 20 WIP commits on their feature branch, when it merges into `main`, it condenses into a single, beautifully documented commit."
*   **Rebasing:** "For my personal local branches, I use interactive rebasing (`git rebase -i`) to clean up my own messy typo-fixes before I ever open a Pull Request for my team to review."
*   **Tooling:** "I implement `pre-commit` hooks (using Husky or similar) locally to lint commit messages before Git accepts them."

## 4. Have you used code review tools? Which ones and what is your process?

**Good Answer Structure:**
*   **Tools:** "I've extensively used GitHub Pull Requests, GitLab Merge Requests, and Bitbucket."
*   **The Process:** "When I review code, I don't just look for bugs. I check if the architecture aligns with our patterns, if tests are included, and if the commit messages are descriptive."
*   **Communication Style:** "I leave constructive, objective comments on specific lines. Most importantly, if a PR exceeds 400 lines of changes, I usually request a quick 5-minute sync with the author to walk through the logic live, as large async reviews often miss critical context and damage team morale."

## 5. What was the most complex Git problem you've debugged?

**Good Answer Structure:**
*   **Situation:** "A massive performance degradation was reported in production, but nobody knew which microservice update caused it over the last month."
*   **Action:** "I utilized `git bisect` on the backend repository. I set the 'good' commit to the last known stable release from 30 days ago, and 'bad' to the current `HEAD`."
*   **Execution:** "Through logarithmic binary searching, running performance unit tests at each midpoint, `git bisect` pinpointed the exact commit within 7 jumps out of 100 commits."
*   **Result:** "The commit was an innocent-looking SQL query change introduced by a junior developer that removed an index. I ran `git revert` on that specific commit hash, deploying the fix to production within the hour, restoring performance without rolling back the last month of other unrelated features."
