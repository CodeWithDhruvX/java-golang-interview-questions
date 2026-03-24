# Configuration & Environment Management Interview Questions (149-153)

## Configuration Patterns

### 149. How do you manage configs across environments?
"I separate configuration from code.

I use `application.yml` for default values (local dev).
For environments (Dev, QA, Prod), I override these values using **Environment Variables** or a **Config Server**.

For example, the DB URL is `jdbc:mysql://localhost:3306` in `application.yml`. But in Prod, Kubernetes injects `SPRING_DATASOURCE_URL` as an environment variable pointing to a real RDS instance. This way, the artifact (JAR/Docker image) is identical across all environments. Only the runtime config changes."

**Spoken Format:**
"Managing configurations across environments is like having different settings for your house, car, and office.

**Code configuration** is like having default settings built into your devices - your phone has default ringtone, your car has default seat positions.

**Environment variables** are like having different remote controls - at home you use one remote, at office another, in production yet another.

The beauty is that:
- Your application (the device) is the same everywhere
- Only the configuration (the settings) changes based on where you're using it
- This prevents committing secrets to Git and allows the same deployment to work in multiple environments

It's like having one universal remote that works with your TV at home, car, and office, but each location has its own preferred settings!"

### 150. What is 12-factor app methodology?
"It’s a set of best practices for building modern, cloud-native (SaaS) applications.

The key ones I follow religiously are:
1.  **Codebase**: One repo, many deploys.
2.  **Dependencies**: Explicitly declare them (Maven/Gradle). No expected system-wide libraries.
3.  **Config**: Store config in the environment, not code.
4.  **Backing Services**: Treat DBs/Queues as attached resources.
5.  **Build, Release, Run**: Strict separation of stages.
6. **Stateless Processes**: Sticky sessions are evil. 7. **Port Binding**: App exports its own port (starts its own Tomcat)."

**Spoken Format:**
"The 12-factor app methodology is like having 12 golden rules for building reliable cloud applications.

Think of it like building a house that can withstand earthquakes:

**Codebase** is like having one set of blueprints - you don't have multiple versions scattered around.

**Dependencies** are like using standardized building materials - everyone knows exactly what bricks and pipes to use.

**Config** is like having electrical switches on the outside - you can change lighting without rewiring the whole house.

**Backing Services** is like having utility connections - you don't build your own generator, you connect to the city's power grid.

**Build, Release, Run** is like having separate construction, inspection, and occupation phases - you don't live in the house while it's being built.

**Stateless** is like having hotel rooms that don't remember guests - each check-in is fresh, no confusion about previous stays.

**Port Binding** is like having your own front door - you don't need to coordinate with neighbors about which door number to use.

These principles ensure your application is reliable, scalable, and maintainable in the cloud!"

### 151. Difference between application.yml and bootstrap.yml?
"`bootstrap.yml` is loaded *before* `application.yml` by the Spring Cloud Context.

It is used specifically for locating external configuration. If you use **Spring Cloud Config Server**, you put the URL of the config server in `bootstrap.yml` so the app knows where to fetch the real config from during startup.

However, in newer Spring Boot versions (2.4+), `bootstrap.yml` is deprecated in favor of `spring.config.import` in `application.yml`."

**Spoken Format:**
"The difference between bootstrap and application config is like the difference between a map and a destination.

**Bootstrap.yml** is like having a map that tells you how to get to the treasure - it shows you where to find the configuration server.

**Application.yml** is like the actual treasure chest - it contains all the real configuration values your application needs.

The process works like this:
1. First, you look at the map (bootstrap) to find out where the treasure is buried
2. Then, you use the map's directions to get to the actual treasure (application config)
3. Finally, you open the treasure chest and use what's inside

In modern Spring Boot, you can put the map location directly in the treasure chest, making the process simpler.

The key insight: Bootstrap is for finding configuration, application is for using configuration!"

### 152. How do you handle secrets securely?
"Never commit them to Git. That’s rule #1.

In local dev, I use environment variables or a local `.env` file (which is gitignored).

In production (Kubernetes), we use **K8s Secrets** or an external Vault like **AWS Parameter Store** or **HashiCorp Vault**. The app fetches the secret at runtime using an IAM role, so the actual password never touches a disk or a config file."

**Spoken Format:**
"Handling secrets securely is like protecting the keys to your kingdom.

**Never commit to Git** is like never writing your kingdom's keys on a public scroll that everyone can read.

For development, you might use a local keychain (environment variables or .env file) - it's like keeping keys in your personal safe.

For production, you use a secure vault (Kubernetes Secrets, AWS Parameter Store, HashiCorp Vault) - it's like having a magical, protected vault that only authorized people can access.

The application fetches secrets at runtime using proper authentication, so:
- Keys never appear in logs or config files
- No one accidentally commits secrets to version control
- Different environments can use different keys

It's like having a royal treasury that protects the kingdom's most valuable secrets!"

### 153. What is feature flagging?
"Feature flags (or toggles) allow us to modify system behavior without changing code.

It looks like: `if (featureManager.isActive("NEW_CHECKOUT")) { useNewCheckout(); } else { useOldCheckout(); }`

This lets us:
1.  **Merge unfinished code** to main without breaking production (it's toggled off).
2.  **Canary Deploy**: Enable new checkout for only 1% of users to test stability.
3.  **Kill Switch**: If the new feature has a bug, we turn it off instantly without rolling back to deployment."

**Spoken Format:**
"Feature flags are like having dimmer switches for your house lights.

Instead of rewiring the entire house to install new lights, you just flip a switch to turn on new features.

The power of this approach:

**Merge unfinished code** - You can wire up the new lights but keep them turned off until you're ready to test.

**Canary Deploy** - You turn on the new lights in just one room to see if they work properly before lighting up the whole house.

**Kill Switch** - If the new lights start flickering or causing problems, you can instantly turn them off everywhere with one switch.

This gives you:
- Safe testing without affecting production
- Instant rollback capability
- Ability to test features with real users
- No need for emergency deployments when something goes wrong

It's like having a smart home system where you can control everything from a central panel!"
