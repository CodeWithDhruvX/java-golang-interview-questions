## 🟢 Real-World Implementation & Security (Questions 151-200)

### Question 151: How do you use `makeExecutableSchema`?

**Answer:**
Combines type definitions and resolvers manually.
Useful when not using Apollo Server's auto-config or when stitching.

```javascript
const schema = makeExecutableSchema({
  typeDefs,
  resolvers,
});
```

---

### Question 152: What is `mergeTypeDefs` and `mergeResolvers`?

**Answer:**
Helpers to merge arrays of types/resolvers.
Allows keeping `user.graphql` and `product.graphql` separate.

---

### Question 153: How do you introspect a remote schema?

**Answer:**
Run the standard Introspection Query against the URL.
Tools like `graphql-codegen` do this to generate SDKs for 3rd party APIs (like GitHub API).

---

### Question 154: What is graphql-yoga and how is it different from Apollo Server?

**Answer:**
Built on **Envelop** plugin system.
Framework agnostic (works in Next.js API routes, Cloudflare Workers, Node).
Lightweight.

---

### Question 155: What are some popular GraphQL IDE plugins?

**Answer:**
VSCode: "GraphQL" (Grammar, Autocomplete). "Apollo GraphQL" (Stats from Studio).

---

### Question 156: What is Hasura and how does it relate to GraphQL?

**Answer:**
Postgres-to-GraphQL engine.
You define DB tables. Hasura generates GraphQL Query/Mutation API instantly.
Great for prototyping or CRUD-heavy apps.

---

### Question 157: What is StepZen?

**Answer:**
Now part of IBM.
Declarative approach. Define `.graphql` files that map queries to REST endpoints via directives (`@rest`).
Deploys to serverless.

---

### Question 158: What is GraphQL Modules?

**Answer:**
Framework for dependency injection in GraphQL.
Modules encapsulate Schema + Resolvers + Providers.
Better for massive monoliths than raw file splitting.

---

### Question 159: What is graphql-request and how is it used?**

**Answer:**
Minimal client.
`request(endpoint, query, variables)`.
Promise based. No cache. Good for scripts.

---

### Question 160: How do you compare Prisma with Hasura?

**Answer:**
*   **Prisma:** ORM library used *inside* your Node.js code to build a custom GraphQL server.
*   **Hasura:** Standalone Server that connects to DB and *is* the GraphQL API.

---

### Question 161: How do you migrate a REST API to GraphQL?

**Answer:**
**Wrapper approach.**
Create GraphQL Schema.
Resolvers call existing REST (via `fetch`).
Once frontend moves to GraphQL, replace Resolver logic with direct DB calls.

---

### Question 162: How do you combine REST and GraphQL in a hybrid architecture?

**Answer:**
Use GraphQL for "Complex Reads" (Dashboards).
Keep REST for "Simple Writes" or Webhooks or File Uploads.
Or use **Sofa** to auto-generate REST endpoints from GraphQL Schema.

---

### Question 163: How do you handle file uploads with Apollo Server?

**Answer:**
Enable `uploads: true` (Uses `graphql-upload`).
Resolver receives `Promise<FileUpload>`.
`file.createReadStream().pipe(s3Stream)`.

---

### Question 164: How do you implement pagination using cursor-based approach?

**Answer:**
Encode `cursor` (usually base64 of ID or Timestamp).
Query `WHERE id > decoded_cursor LIMIT 10`.
Return `nextCursor`.

---

### Question 165: What are some common pagination strategies in GraphQL?

**Answer:**
1.  **Offset:** `page: 1` (Simple DB limit/offset).
2.  **Cursor:** `after: "xyz"` (Scalable, Real-time safe).

---

### Question 166: How do you control access at the field level in GraphQL?

**Answer:**
Directives: `email: String @auth(role: ADMIN)`.
Directives wrap the resolver. If User is not Admin, return null or Error.

---

### Question 167: How do you support multi-tenancy in GraphQL?

**Answer:**
Extract `x-tenant-id` header in Context.
Pass `tenantId` to every DB call.
`db.users.find({ tenantId, ... })`.

---

### Question 168: How do you maintain backwards compatibility in your schema?

**Answer:**
**Never delete.**
Rename old field `address` -> `addressLegacy` (or keep name and allow new structure alongside).
Use Union types if return shape changes drastically.

---

### Question 169: How do you implement rate-limiting per user?

**Answer:**
Redis Token Bucket.
Middleware runs before resolver.
`key = user:${userId}`.
Decrement token. If 0, throw "Too Many Requests".

---

### Question 170: How do you integrate GraphQL with microservices?**

**Answer:**
**Federation.**
Gateway talks to 5 services.
Service 1 (Users) resolves `User`.
Service 2 (Reviews) extends `User { reviews: [] }`.
Gateway merges them.

---

### Question 171: What is GraphQL over gRPC?

**Answer:**
Using gRPC for backend-to-backend (Gateway to Service).
Gateway converts GQL Request -> gRPC Call.
Faster than JSON over HTTP.

---

### Question 172: How do you set up a GraphQL API Gateway?

**Answer:**
`ApolloGateway`.
Provide `serviceList` (URLs).
Gateway introspects them and builds Query Plan.

---

### Question 173: What is the role of Redis in GraphQL APIs?

**Answer:**
1.  **Response Cache:** Cache full GQL response.
2.  **Dataloader Cache:** Shared cache for DB entities.
3.  **PubSub:** Distribute Subscription events.

---

### Question 174: How do you track which clients use which fields in a schema?

**Answer:**
Apollo Trace Reporting.
Sends metrics to Studio.
"Field `User.age` used by `iOS-App-v1` 500 times".

---

### Question 175: How do you handle large payloads efficiently?

**Answer:**
`@stream` directive.
Send first 5 items of list.
Keep connection open. Send rest as chunks.
Frontend updates list progressively.

---

### Question 176: How do you version your GraphQL schema?

**Answer:**
Field evolution.
`fieldName` -> `fieldNameV2`.
Clients migrate. Remove `fieldName` later.

---

### Question 177: What is a GraphQL schema registry?

**Answer:**
Repo for your schemas (e.g., Apollo Studio).
Validates pushed schemas (CI).
Gateway downloads latest schema from Registry (High Availability).

---

### Question 178: What are breaking changes in GraphQL and how do you prevent them?

**Answer:**
**Breaking:** Removing field, changing Type (Int->String), Adding required arg.
**Prevent:** `graphql-inspector diff`. Fails build if breaking change detected.

---

### Question 179: What’s the best way to handle long-running mutations?

**Answer:**
Async pattern.
Mutation returns `JobId`.
Query `jobStatus(id)`.
Or Subscription `jobFinished`.

---

### Question 180: How do you handle transactions in GraphQL mutations?

**Answer:**
DB Transaction in Resolver.
`await db.transaction(async t => { ... })`.
GraphQL itself doesn't have transactions across multiple mutations in one request (they run sequentially but independently).

---

### Question 181: How do you protect against GraphQL injection?

**Answer:**
Use Variables.
Never interpolate string into query.
`db.query("SELECT * FROM users WHERE id = " + args.id)` is bad.
Use Prepared Statements (ORM).

---

### Question 182: What is depth analysis and why is it important?

**Answer:**
Blocks queries that are too deep.
`friends { friends { friends ... } }`.
Prevents Stack Overflow / DB thrashing.

---

### Question 183: How do you prevent overly nested queries?

**Answer:**
`graphql-depth-limit`. Pass to validation rules.

---

### Question 184: How do you use GraphQL Armor for security?

**Answer:**
Package `graphql-armor`.
Configures alias limits, depth limits, and disables stack traces automatically.

---

### Question 185: How do you prevent information leakage in GraphQL?

**Answer:**
Disable Suggestions.
If user types `password` (incorrectly), GraphQL suggests "Did you mean passwordHash?".
Disable this in production.

---

### Question 186: What is a security audit checklist for GraphQL APIs?

**Answer:**
1.  Introspection Disabled?
2.  Depth Limit?
3.  Debug mode off?
4.  Auth on all fields?

---

### Question 187: What is the purpose of query whitelisting?

**Answer:**
Allow only known queries (Persisted Queries).
Attacker cannot craft shape to scrape data. They can only run valid app queries.

---

### Question 188: How do you enforce scopes or roles at the schema level?

**Answer:**
`directive @scope(scopes: ["read:user"])`.
Middleware validates `context.user.scopes` has required scope.

---

### Question 189: How do you block introspection queries in production?**

**Answer:**
Validation Rule `NoIntrospection`.
If query contains `__schema`, return Error.

---

### Question 190: How can you control query cost by complexity?

**Answer:**
`graphql-query-complexity`.
Define: `friends: 10`, `nested: 10`.
Multiply by list arguments (`first: 10` * 10 = 100).
Max Cost: 500.

---

### Question 191: How do you authenticate users in subscriptions?

**Answer:**
Standard HTTP headers don't exist in WS after connect.
Send `authToken` in `payload` of `connection_init` message.
Verify in `onConnect` server callback.

---

### Question 192: What are the best practices for securing GraphQL endpoints?

**Answer:**
Treat it like a database exposed to web.
Validate everything. Limit everything.

---

### Question 193: What is the risk of `__schema` access and how to mitigate it?

**Answer:**
Risk: Automating attacks (finding types).
Mitigate: Disable introspection.

---

### Question 194: How do you enforce field-level authorization using middleware?

**Answer:**
**GraphQL Shield.**
Code-first permissions.
`rule()` function returns true/false.
Map rules to Schema Map.

---

### Question 195: What are the security implications of third-party schema stitching?

**Answer:**
If the remote schema adds a malicious field, your gateway exposes it.
Always validate/filter remote schemas.

---

### Question 196: How do you handle authorization for deeply nested fields?

**Answer:**
Auth should cascade.
If `User` is private, `User.email` is implicitly private.
But usually, checks are per-resolver for granularity.

---

### Question 197: How do you isolate public vs internal schemas?

**Answer:**
Directives `@internal`.
Tools `graphql-transform` filter schema.
Generate `public-schema.graphql` (stripping internal) for public Gateway.

---

### Question 198: How do you audit client usage in GraphQL?

**Answer:**
Require `apollographql-client-name` header.
Log every request with client name.

---

### Question 199: What is the risk of auto-generated queries?

**Answer:**
Scrapers can generate massive queries to dump DB.
Prevent using Complexity Analysis.

---

### Question 200: How do you implement RBAC (role-based access control) in GraphQL?

**Answer:**
Schema Directives are best.
`@auth(requires: ADMIN)`.
Declarative and visible in SDL.
