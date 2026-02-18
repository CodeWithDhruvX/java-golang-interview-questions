# Configuration & Environment Management Interview Questions (149-153)

## Configuration Patterns

### 149. How do you manage configs across environments?
"I separate configuration from code.

I use `application.yml` for default values (local dev).
For environments (Dev, QA, Prod), I override these values using **Environment Variables** or a **Config Server**.

For example, the DB URL is `jdbc:mysql://localhost:3306` in `application.yml`.
But in Prod, Kubernetes injects `SPRING_DATASOURCE_URL` as an environment variable pointing to the real RDS instance.

This way, the artifact (JAR/Docker image) is identical across all environments. Only the runtime config changes."

### 150. What is 12-factor app methodology?
"It’s a set of best practices for building modern, cloud-native (SaaS) applications.

The key ones I follow religiously are:
1.  **Codebase**: One repo, many deploys.
2.  **Dependencies**: Explicitly declare them (Maven/Gradle). No expected system-wide libraries.
3.  **Config**: Store config in the environment, not code.
4.  **Backing Services**: Treat DBs/Queues as attached resources.
5.  **Build, Release, Run**: Strict separation of stages.
6.  **Stateless Processes**: Sticky sessions are evil.
7.  **Port Binding**: App exports its own port (starts its own Tomcat)."

### 151. Difference between application.yml and bootstrap.yml?
"`bootstrap.yml` is loaded *before* `application.yml` by the Spring Cloud Context.

It is used specifically for locating external configuration. If you use **Spring Cloud Config Server**, you put the URL of the config server in `bootstrap.yml` so the app knows where to fetch the real config from during startup.

However, in newer Spring Boot versions (2.4+), `bootstrap.yml` is deprecated in favor of `spring.config.import` in `application.yml`."

### 152. How do you handle secrets securely?
"Never commit them to Git. That’s rule #1.

In local dev, I use environment variables or a local `.env` file (which is gitignored).

In production (Kubernetes), we use **K8s Secrets** or an external Vault like **AWS Parameter Store** or **HashiCorp Vault**. The app fetches the secret at runtime using an IAM role, so the actual password never touches a disk or a config file."

### 153. What is feature flagging?
"Feature flags (or toggles) allow us to modify system behavior without changing code.

It looks like: `if (featureManager.isActive("NEW_CHECKOUT")) { useNewCheckout(); } else { useOldCheckout(); }`

This lets us:
1.  **Merge unfinished code** to main without breaking production (it’s toggled off).
2.  **Canary Deploy**: Enable the new checkout for only 1% of users to test stability.
3.  **Kill Switch**: If the new feature has a bug, we turn it off instantly without rolling back the deployment."
