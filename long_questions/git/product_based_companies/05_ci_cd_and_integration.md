# Git CI/CD & Integration (Product-Based Companies)

This document covers questions related to how Git integrates with Continuous Integration and Continuous Deployment (CI/CD) pipelines, a critical area for product engineering roles.

## 1. How does Git trigger pipelines?

Git triggers CI/CD pipelines primarily through **Webhooks**.
*   When an event occurs in the Git repository (e.g., a `push`, a new `tag` is created, or a Pull Request is opened/merged), the Git hosting server (GitHub, GitLab, Bitbucket) automatically sends an HTTP POST request containing a JSON payload to a pre-configured URL endpoint on the CI/CD server (like Jenkins, GitHub Actions, or CircleCI).
*   The CI server parses this payload to determine the branch name, the commit hash, and the author, and then kicks off the corresponding pipeline defined for that specific event.

## 2. Integrating Git with Jenkins / GitHub Actions / GitLab CI

*   **GitHub Actions:** Integration is native. You simply create YAML files inside the `.github/workflows/` directory of your repository. GitHub automatically detects these files and runs the pipelines on its own runners based on the events defined in the `on:` trigger map.
*   **GitLab CI:** Native integration. You create a `.gitlab-ci.yml` file in the root of the project. GitLab Runners pick up the jobs defined here whenever code is pushed.
*   **Jenkins:** Requires adding a "Jenkinsfile" to the repository. In Jenkins, you configure a Multibranch Pipeline job and provide it with repository credentials and the webhook URL. When Jenkins receives a webhook from Git, it reads the Jenkinsfile to execute the pipeline.

## 3. What are Branch Policies and Build Validation?

Branch policies (or Branch Protections) are rules enforced by the Git hosting platform to ensure code quality before it can be merged into critical branches (like `main` or `release`).
*   **Build Validation:** A specific branch policy that mandates that a CI pipeline *must* run and successfully pass (turn green) against the Pull Request before the "Merge" button becomes clickable. This guarantees that broken code or failing tests are never merged into the main codebase.

## 4. What is a "Merge on Green Build" strategy?

Also known as a "Merge Queue" or "Auto-Merge."
In high-velocity teams where dozens of Pull Requests are approved daily, a developer can click "Auto-Merge" on an approved PR. The system will add the PR to a queue, automatically run the slow integration tests, and if the build stays green, it merges the code without further human intervention. If the build fails, the PR is kicked out of the queue, and the author is notified.

## 5. How do you automate version bumps and tagging via Git in CI/CD?

In mature product teams, developers rarely manually tag releases. It is automated through standard commit conventions and CI/CD:
1.  **Conventional Commits:** Developers write commit messages following a strict standard (e.g., `fix: resolve login bug`, `feat: add payment gateway`, `BREAKING CHANGE: update API`).
2.  **Semantic Release:** When code is merged to `main`, a tool in the CI pipeline (like `semantic-release`) parses the commit history since the last release.
3.  **Automation:** Based on the commit types, it calculates the new Semantic Version number (e.g., bumps from v1.2.3 to v1.3.0 for a `feat`).
4.  **Git Tagging:** The pipeline automatically creates a new Git Tag (using `git tag -a v1.3.0 -m "Release v1.3.0"`) and pushes it back to the repository (`git push --tags`). This tag often triggers a subsequent deployment pipeline.
