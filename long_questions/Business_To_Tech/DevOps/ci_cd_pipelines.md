# CI/CD Pipelines

## 1. What is CI/CD?
*   **CI (Continuous Integration)**: Developers merge their changes back to the main branch as often as possible. Automated builds and tests run to ensure the code works.
*   **CD (Continuous Delivery/Deployment)**:
    *   *Delivery*: Code is automatically built/tested and ready to release (Click button to deploy).
    *   *Deployment*: Code is automatically deployed to production without manual intervention.

## 2. Anatomy of a Pipeline
A typical pipeline (e.g., Jenkins, GitLab CI, GitHub Actions) has these stages:

### Stage 1: Build & Lint
*   **Checkout**: Clone the repo.
*   **Lint**: Check coding standards (e.g., Checkstyle, ESLint).
*   **Build**: Compile the code (e.g., `mvn package`, `npm build`).
*   *Output*: An "Artifact" (JAR file, Dist folder).

### Stage 2: Test
*   **Unit Tests**: Run tests that verify small components (`junit`).
*   **Integration Tests**: Verify interactions between components (In-memory DB).
*   *Gate*: If tests fail, pipeline stops immediately (Fast Feedback Loop).

### Stage 3: Dockerize (Containerize)
*   **Build Image**: `docker build -t my-app:v1 .`
*   **Push Registry**: `docker push my-registry/my-app:v1`.

### Stage 4: Deploy (CD)
*   **Deploy Dev**: Deploy to a Dev environment.
*   **E2E Tests**: Run Selenium/Cypress automation against Dev.
*   **Approval**: (Optional) Manual approval via Slack/Email.
*   **Deploy Prod**: `kubectl set image deployment/my-app my-app=v1`.

## 3. Best Practices
1.  **Fail Fast**: Run fastest tests (Unit) first. Don't wait 1 hour for E2E tests just to find a syntax error.
2.  **Immutability**: Build the artifact/image ONCE. Deploy the *same* image to Dev, QA, and Prod. Don't rebuild for Prod (code might change).
3.  **Infrastructure as Code (IaC)**: Define the pipeline in a file (`Jenkinsfile`, `.gitlab-ci.yml`) inside the repo.

## 4. Interview Questions
1.  **What is the difference between Continuous Delivery and Continuous Deployment?**
    *   *Ans*:
        *   **Delivery**: You *can* release at any time (the artifact is ready), but human presses the button.
        *   **Deployment**: Every commit that passes tests goes straight to Production automatically.
2.  **How do you handle secrets (DB passwords) in CI/CD?**
    *   *Ans*: NEVER check them into Git. Use "Secrets Management" tools (GitHub Secrets, Jenkins Credentials, HashiCorp Vault) and inject them as Environment Variables during the build/deploy phase.
