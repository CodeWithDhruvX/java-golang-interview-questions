# 38. Spring Boot Deployment, Security & JPA

**Q: Spring Boot Profiles**
> "Profiles let you separate configuration for different environments.
> You don't want to connect to `localhost:3306` in Production.
>
> You define `application-dev.yml` and `application-prod.yml`.
> You activate one at runtime: `-Dspring.profiles.active=prod`.
> Spring automatically loads the correct file and ignores the others."

**Indepth:**
> **Multi-Doc**: `---`. You can put multiple profiles in a single `application.yml` separated by three dashes. This is useful for keeping "default" vs "dev" configs in one file while still separating them logically.


---

**Q: Graceful Shutdown**
> "In the past, if you killed the app, active requests just failed immediately.
>
> Now, you enable `server.shutdown=graceful`.
> When you send a *SIGTERM* (kill signal), Spring Boot stops accepting *new* requests but waits (e.g., 30s) for *active* requests to finish processing before actually shutting down. It prevents user errors during deployments."

**Indepth:**
> **Kubernetes**: `preStop` hook. In K8s, when a pod is terminated, it stops receiving traffic. But there's a race condition where the Load Balancer might still send a request. A graceful shutdown + a small `sleep` in the preStop hook ensures zero-downtime deployments.


---

**Q: Method Level Security (@PreAuthorize)**
> "Instead of securing URLs in a config file (which can get messy), you secure the Java methods directly.
>
> `@PreAuthorize("hasRole('ADMIN')")`
> `public void deleteUser(String id) { ... }`
>
> This is cleaner and safer. Even if you forget to secure the URL, the service method itself is protected."

**Indepth:**
> **PostAuthorize**: Less common but powerful. It runs *after* the method. You can check the *return value*. `@PostAuthorize("returnObject.owner == authentication.name")`. Use this to ensure a user can only read their *own* data.


---

**Q: Custom Authentication Provider**
> "If you aren't using standard Database or LDAP auth—maybe you verify users against a Legacy Mainframe or a 3rd Party API.
>
> You implement `AuthenticationProvider`.
> *   `authenticate()`: You write the logic to call the Mainframe.
> *   If success, return a valid `UsernamePasswordAuthenticationToken`.
> *   Spring Security plugs this specific logic into its general login flow."

**Indepth:**
> **Supports**: `supports()`. The `AuthenticationManager` loops through all providers. Your custom provider must implement `supports(Class<?> authentication)` to tell Spring "I know how to handle this specific type of token/login".


---

**Q: JPA N+1 Problem**
> "This is the most common performance killer.
> You fetch 10 Departments. Then for *each* Department, you fetch its Employees.
> *   1 Query for Departments.
> *   10 Queries for Employees.
> Total 11 queries.
>
> **Fix**: Use `JOIN FETCH` in JPQL: `SELECT d FROM Department d JOIN FETCH d.employees`. This fetches everything in **one** single query."

**Indepth:**
> **Statistics**: `Hibernate Statistics`. Enable `spring.jpa.properties.hibernate.generate_statistics=true` in tests. It prints "Session Metrics" at the end, showing exactly how many JDBC statements were executed. Ideally, it should be 1, not N+1.


---

**Q: Optimistic vs Pessimistic Locking**
> "**Optimistic**: You assume conflicts are rare. You add a `@Version` column. If two people save at the same time, the second one fails with `OptimisticLockException`.
>
> "**Pessimistic**: You assume conflicts are frequent. You lock the database row (`SELECT ... FOR UPDATE`). No one else can even *read* it until you are done. It's safer but slower."

**Indepth:**
> **Deadlocks**: Pessimistic Locking can cause deadlocks if two transactions lock resources in different orders. Always order your locks consistently (e.g., sort by ID before locking) to prevent this.


---

**Q: Cascade Types**
> "Cascading means: 'What happens to the Child when I do something to the Parent?'
>
> *   `CascadeType.PERSIST`: If I save the Parent, save the Children too.
> *   `CascadeType.REMOVE`: If I delete the Parent, delete all Children (orphans).
> *   `CascadeType.ALL`: Do everything."

**Indepth:**
> **Orphans**: `orphanRemoval=true`. This is different from Cascade DELETE. If you remove a child from the parent's list (`parent.getChildren().remove(child)`), `orphanRemoval` deletes the child from the DB. Cascade DELETE only works if you delete the *Parent*.


---

**Q: Flyway / Liquibase**
> "Never let Hibernate auto-generate your Production schema (`ddl-auto=update`). It's dangerous.
>
> Use **Flyway**.
> You write SQL scripts: `V1__init.sql`, `V2__add_column.sql`.
> On startup, Flyway checks the DB version. If it's at V1, it runs V2.
> This ensures your Database structure is always in sync with your Java code."

**Indepth:**
> **Baseline**: `baselineOnMigrate`. If you introduce Flyway to an existing populated database, you need to "baseline" it. This tells Flyway "Assume the DB is already at version 1, start running scripts from V2".


---

**Q: Logging SQL**
> "To see what Hibernate is doing:
> `spring.jpa.show-sql=true` (Basic).
>
> To see the *values* inside the `?`:
> `logging.level.org.hibernate.type.descriptor.sql=TRACE`.
>
> This is vital for debugging N+1 problems or wrong parameter bindings."

**Indepth:**
> **P6Spy**: `P6Spy` is a library that wraps the JDBC driver. It logs the *actual* SQL sent to the DB with all `?` replaced by real values, plus execution time. It is much better than Hibernate's built-in logging.


---

**Q: @Entity vs @Table**
> "**@Entity**: Makes it a JPA Object. The class name becomes the Table name by default.
>
> "**@Table**: Allows you to customize the DB table details.
> `@Table(name = "tbl_users", indexes = @Index(...))`
> You use `@Table` when the DB name doesn't match your Java class name."

**Indepth:**
> **Constraints**: Unique Constraints. You can define multi-column unique constraints inside `@Table(uniqueConstraints = @UniqueConstraint(columnNames = {"firstName", "lastName"}))`.

