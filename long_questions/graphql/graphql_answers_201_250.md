## 🟢 Architecture, Scalability & Data Modeling (Questions 201-250)

### Question 201: How do you serve a GraphQL API over WebSocket?

**Answer:**
Use `graphql-ws` (standard).
Server defines `SubscriptionServer` or uses `useServer` with `ws` lib.
Client uses `GraphQLWsLink`.
Handles `connection_init`, `ping/pong`, and streaming updates.

---

### Question 202: What is a GraphQL Gateway?

**Answer:**
Entry point that aggregates subgraphs.
Clients talk to Gateway. Gateway talks to User/Product services.
Hides microservice complexity.

---

### Question 203: How do you handle file uploads in a serverless GraphQL environment?

**Answer:**
Avoid binary payload in Lambda (Limit 6MB).
**Pre-signed URL Pattern:**
1.  Mutation `getUploadUrl(filename)`.
2.  Server returns S3 Signed URL.
3.  Client PUTs file to S3.
4.  Mutation `confirmUpload(key)`.

---

### Question 204: What is the difference between a Graph and a DAG in GraphQL context?**

**Answer:**
*   **Graph:** The Schema (Types reference each other cyclically). `User -> Post -> User`.
*   **DAG/Tree:** The Query Execution. It must terminate (acyclic) or it loops forever.

---

### Question 205: How do you optimize cold starts for GraphQL on AWS Lambda?

**Answer:**
1.  Use `swc/esbuild` to minify.
2.  Use `graphql-yoga` (lighter than Apollo).
3.  Avoid heavy schema stitching at runtime (bundle it).
4.  Use Provisioned Concurrency.

---

### Question 206: How do you protect GraphQL APIs from DDoS attacks?

**Answer:**
WAF (AWS WAF).
Rate Limit by IP.
Complexity Limit (stop heavy queries).
Persisted Queries (stop random queries).

---

### Question 207: What is query flattening and when is it used?

**Answer:**
Optimization.
Converting nested GQL selection set `users { posts }` into optimized SQL `JOIN`.
`join-monster` library does this.

---

### Question 208: How do you handle multi-region data replication in GraphQL?

**Answer:**
GraphQL is stateless.
Deploy GraphQL servers in all regions behind GeoDNS.
Connect them to local Read Replicas (Aurora Global DB).
Write to Primary (Cross-region latency for mutations).

---

### Question 209: How do you implement a BFF (Backend For Frontend) with GraphQL?

**Answer:**
One Schema per Client? No.
Federation is better.
But valid BFF pattern: A lightweight Node GQL server wrapping legacy REST APIs specifically for the Mobile App.

---

### Question 210: What is the role of a schema registry in a CI/CD pipeline?

**Answer:**
Gatekeeper.
CI pushes schema candidate.
Registry checks: "Does this break Prod?".
If yes, fail CI. If no, register version.

---

### Question 211: How do you monitor individual field usage?

**Answer:**
Trace logging.
Apollo Studio offers "Field Usage" heatmap.
Identify fields safe to delete.

---

### Question 212: How do you handle timeouts for specific resolvers?

**Answer:**
`Promise.race([ dbCall, timeout(5000) ])`.
If timeout wins, throw `GraphQLError("Timeout")`.
Server-level: `httpServer.timeout`.

---

### Question 213: What is the N+1 problem in Distributed Systems?

**Answer:**
Resolving a list of `Orders`. For each order, calling `User Service` via HTTP to get `User`.
100 Orders = 101 HTTP calls.
**Fix:** Batch HTTP calls (Dataloader for Services).

---

### Question 214: How do you version a Federated Graph?**

**Answer:**
Subgraph schema changes are composed by Gateway.
Gateway validates compatibility.
You don't version the "Graph", you evolve it continuously.

---

### Question 215: How do you implement caching in a GraphQL Gateway?

**Answer:**
Query Plan Caching (saving the execution tree).
Subgraph Caching (caching REST responses from sub-services).

---

### Question 216: What is "Schema Stewardship"?

**Answer:**
Governance.
Defining who owns `User` type. Review process for new fields. Avoiding duplication.

---

### Question 217: How would you design a GraphQL API for a chat application?

**Answer:**
`Conversation`, `Message`, `User`.
Subscription `messageAdded(conversationId)`.
Mutation `sendMessage`.
Cursor Pagination for `messages(last: 50, before: cursor)`.

---

### Question 218: How do you handle data consistency in distributed mutations?

**Answer:**
Saga Pattern.
Mutation `createOrder` -> Calls OrderService.
If success, triggers InventoryService.
If Inventory fails, trigger `compensateOrder` (Cancel).

---

### Question 219: How do you secure a public-facing GraphQL API?

**Answer:**
Strict Rate Limiting.
Disable Introspection.
Whitelisted Queries ONLY (Persisted).
No free-form queries allowed.

---

### Question 220: What are "supergraphs"?

**Answer:**
The unified graph composed of all subgraphs in Federation.

---

### Question 221: How do you solve the N+1 problem without DataLoader?

**Answer:**
**Look-ahead.**
Inspect `info` argument.
See that `posts` are requested.
Modify SQL: `SELECT * FROM users LEFT JOIN posts ...`.
Return nested structure.

---

### Question 222: What is the difference between server-side and client-side caching?

**Answer:**
*   **Server:** Caches *Responses* (Redis/CDN). Shared across users.
*   **Client:** Caches *Entities* (Normalized Store). Shared across components.

---

### Question 223: How do you use HTTP caching with GraphQL?

**Answer:**
Use GET.
Query params: `?query={...}&variables={...}`.
Response Header: `Cache-Control: max-age=60`.
CDN caches key based on URL.

---

### Question 224: What is a persisted query whitelist?

**Answer:**
A dictionary on server: `hash -> query string`.
Server rejects any request body that isn't a known hash ID.

---

### Question 225: How do you optimize resolvers that call slow 3rd party APIs?

**Answer:**
Cache the 3rd party response (TTL 5 min).
Use `@defer` to return the rest of the query first, then stream the slow field.

---

### Question 226: What is the `@defer` directive?

**Answer:**
Client sends: `query { slowField @defer }`.
Server sends multipart HTTP response.
Chunk 1: Initial data.
Chunk 2: Path to `slowField` + data.

---

### Question 227: What is the `@stream` directive?

**Answer:**
For Lists. `items @stream(initialCount: 5)`.
Server sends first 5. Then streams remaining items one by one (or batches).

---

### Question 228: How do you implement Edge Caching for GraphQL?

**Answer:**
**GraphCDN** (or Stellate).
Specialized CDN that understands GQL Cache Headers.
Purges specific types (`User:1`) globally when Mutation occurs.

---

### Question 229: What is "Over-fetching" in the context of DB performance?

**Answer:**
Fetching 20 columns from SQL when GraphQL only asked for 2.
Optimize by mapping GraphQL selection set to SQL `SELECT name, email`.

---

### Question 230: How do you handle massive schema sizes (10k+ types)?**

**Answer:**
Tooling slows down.
Split into Federation.
Lazy-load types in IDE plugins.

---

### Question 231: How do you load-balance GraphQL subscriptions?

**Answer:**
Sticky Sessions (Hash by UserID).
Or Redis PubSub Backplane. Server A receives event, publishes to Redis. Redis wakes up Server B, C, D to broadcast to their WS clients.

---

### Question 232: What is memoization in resolvers?

**Answer:**
Caching function results.
`_.memoize(fetchUserById)`.
Scopes usually to Request Context (Dataloader) to avoid stale data across requests.

---

### Question 233: How do you profile a slow GraphQL request?

**Answer:**
Identify "Critical Path".
The longest chain of serial resolvers.
Optimize that chain (parallelize or optimize DB).

---

### Question 234: How do you minimize payload size?

**Answer:**
Use Short Aliases.
`u: user { n: name }`.
Or standard compression (Gzip).

---

### Question 235: How do you handle "Thundering Herd" with GraphQL?

**Answer:**
Cache Stampede.
Use `DataLoader` (coalesces requests).
Or Redis "Lock" when regenerating cache.

---

### Question 236: What is GraphQL JIT (Just-In-Time) compilation?

**Answer:**
`graphql-jit`.
Compiles the query plan into a JS function.
Skips AST traversal overhead for subsequent requests of same query.

---

### Question 237: How do you optimize for mobile networks?

**Answer:**
Persisted Queries (Small Upload).
Normalized Cache (Less fetching).
Defer/Stream (Fast Time-to-Interactive).

---

### Question 238: How do you handle large lists in GraphQL?

**Answer:**
Hard Limit. `first` arg max 100.
Pagination.
Never return array without limit.

---

### Question 239: What is the cost of Schema Stitching vs Federation?

**Answer:**
Stitching: High Maintenance (Glue code). central bottleneck.
Federation: Distributed complexity. network overhead.

---

### Question 240: How do you stress test a GraphQL server?

**Answer:**
`k6` or `Artillery`.
Send high-complexity queries.
Send subscriptions bomb.

---

### Question 241: What is a "node" interface?

**Answer:**
Relay Standard.
`id: ID!` global unique.
Allows `node(id: "User:1") { ... on User { name } }` to fetch anything.

---

### Question 242: How do you design specific mutations vs generic CRUD?

**Answer:**
**RPC Style:** `publishPost(id: 1)`. (Intent).
**CRUD Style:** `updatePost(id: 1, status: PUBLISHED)`. (Data).
RPC is better for complex business logic.

---

### Question 243: How do you handle polymorphic data?

**Answer:**
`Interface` (Shared fields).
`Union` (Disjoint types).
`search { ... on User {} ... on Post {} }`.

---

### Question 244: What is the best practice for naming lists?

**Answer:**
Plural `users`.
Or `usersConnection` if using Relay edges.

---

### Question 245: How do you model Many-to-Many relationships?

**Answer:**
`User.groups` -> `[Group]`.
`Group.users` -> `[User]`.
Resolvers query joint table.

---

### Question 246: When should you use a Custom Scalar?

**Answer:**
For domain types like `Date`, `Email`, `URL`, `JSON`.
Ensures validation at input/output level.

---

### Question 247: How do you handle "Input Unions" (Polymorphic Input)?

**Answer:**
Not supported natively by spec.
Workaround: `OneOf` Input Object Directive (RFC).
`input Media { image: ImageInput, video: VideoInput }`. Enforce only one is set.

---

### Question 248: How do you model "permissions" in the schema?

**Answer:**
`type Permissions { canEdit: Boolean, canDelete: Boolean }`.
`viewer { permissions }`.
Allows frontend to disable buttons conditionally.

---

### Question 249: How do you handle Date/Time?

**Answer:**
Use `graphql-scalars` package.
`DateTime` scalar. Serializes to ISO String `2023-01-01T00:00:00Z`.

---

### Question 250: What is the "Payload" pattern in mutations?

**Answer:**
Return an object, not the entity directly.
```graphql
type CreateUserPayload {
  user: User
  errors: [UserError]
  success: Boolean
}
```
Allows extensibility.
