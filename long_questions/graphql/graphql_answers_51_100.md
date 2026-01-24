## 🟢 Advanced Features & Apollo (Questions 51-100)

### Question 51: What is DataLoader and how does it improve resolver performance?

**Answer:**
DataLoader is a utility that creates a Coalescing Batching mechanism.
It waits one "tick" of the event loop, collects all `ID`s requested by resolvers, and sends a single Batch Query to the DB. It also provides per-request caching.

**Code:**
```javascript
const userLoader = new DataLoader(keys => db.batchGetUsers(keys));
const user = await userLoader.load(1);
```

---

### Question 52: How do you handle authorization in resolvers?

**Answer:**
Logic inside the resolver.
Or wrapped in a Higher Order Function.
`requireAuth(resolver)`.

**Code:**
```javascript
const resolvers = {
  Query: {
    secret: (parent, args, context) => {
      if (!context.user) throw new AuthenticationError('Who are you?');
      return "Secret Data";
    }
  }
};
```

---

### Question 53: How can you reuse resolver logic across fields?

**Answer:**
Extract logic into Service functions (`UserService.getById(id)`).
Or use Resolver Composition tools (`graphql-resolvers`) to pipe small functions (`combine(isAuthenticated, isAdmin, getSecret)`).

---

### Question 54: How do you use middleware in resolvers?

**Answer:**
Middleware libraries (`graphql-middleware`) allow defining logic that runs before/after resolvers based on schema directives or types.

---

### Question 55: How do you test GraphQL resolvers?

**Answer:**
Unit Test them as pure functions.
Mock the `context` object (db, user).
Call the resolver `resolver(null, args, mockContext)` and assert result.

---

### Question 56: How does resolver priority work when multiple fields have the same name?

**Answer:**
1.  Is there a defined resolver for this field? Use it.
2.  If not, look for property on the `parent` object with same name.
3.  If not, return `null` (or undefined).

---

### Question 57: How do you handle caching in resolvers?

**Answer:**
Resolvers can check `Redis`.
`const cached = await redis.get(key); if(cached) return cached;`
Better: Use Data Sources (Apollo) which have built-in caching logic.

---

### Question 58: How do you resolve computed fields?

**Answer:**
Define a field `fullName` in schema.
In resolver:
```javascript
User: {
  fullName: (parent) => `${parent.firstName} ${parent.lastName}`
}
```

---

### Question 59: How do resolvers affect performance in large schemas?

**Answer:**
"Explosion of resolvers".
If fetching a List of 1000 items, and each item has a field resolving via DB, that's 1001 DB calls (N+1).
**Fix:** DataLoader is mandatory.

---

### Question 60: Can you use async/await in resolvers?

**Answer:**
Yes. Resolvers can return a Promise. GraphQL engine waits for it to resolve before moving to the next level of the tree.

---

### Question 61: What is GraphQL federation?

**Answer:**
A microservice architecture for GraphQL.
Instead of one Monolith schema, you have multiple "Subgraphs" (User Service, Product Service).
A "Gateway" composes them into one Supergraph.

---

### Question 62: What is Apollo Federation?

**Answer:**
The specific specification/libraries from Apollo to implement Federation.
Uses `@key`, `@shareable`, `@external` directives to define how entities link across multiple services.

---

### Question 63: What is Apollo Client and how does it work?

**Answer:**
A robust GraphQL client (React/Angular/Vue/iOS).
Features:
1.  **Intelligent Caching:** Normalizes data by ID.
2.  **Declarative Data Fetching:** Hooks (`useQuery`).
3.  **State Management:** Replaces Redux for server data.

---

### Question 64: How does Apollo Server differ from other GraphQL servers?

**Answer:**
It is "Batteries Included".
Includes: Playground/Explorer, Automatic Persisted Queries, Federation support, Tracing, Plugin system.

---

### Question 65: What is Relay in GraphQL?

**Answer:**
A Facebook-designed GraphQL client.
Focuses on High Performance and Scalability.
Enforces tight coupling between Component and Data Fragment (`createFragmentContainer`).
Uses strict Cursor Pagination (Connections).

---

### Question 66: What is the connection model in Relay?

**Answer:**
Standard for lists.
Structure:
```graphql
friends(first: 10, after: "cursor") {
  edges {
    cursor
    node { ... }
  }
  pageInfo { hasNextPage }
}
```

---

### Question 67: What are edges and nodes in Relay?

**Answer:**
*   **Node:** The actual Item (User).
*   **Edge:** Usage of the item in this list (includes Metadata like `addedAt` or Cursor).

---

### Question 68: What is GraphQL Mesh?

**Answer:**
A gateway that converts *non-GraphQL* sources (Swagger REST, gRPC, SOAP) into a GraphQL API automatically.

---

### Question 69: What are persisted queries?

**Answer:**
Security/Performance technique.
Client sends `hash("query { ... }")` instead of full string.
Server looks up hash.
Reduces Upload Bandwidth. Secure (whitelisting).

---

### Question 70: What are automatic persisted queries (APQ)?

**Answer:**
Zero-config persisted queries.
1.  Client sends Hash.
2.  Server says "Not Found".
3.  Client sends Hash + Full Query.
4.  Server caches Hash -> Query.
5.  Client sends Hash only in future.

---

### Question 71: What are GraphQL directives and how can they be extended?

**Answer:**
Directives modify behavior.
Custom Directives: `@upper`, `@auth`, `@cache`.
Implemented by wrapping schema/resolvers at startup.

---

### Question 72: What is client-side schema?

**Answer:**
Apollo Client feature.
Define fields in `client-schema.graphql` that only exist in Browser (e.g., `isLoggedIn`).
Query them mixed with server data: `user { name @client }`.

---

### Question 73: What is server-side schema?

**Answer:**
The actual API contract defined on the backend.

---

### Question 74: How does GraphQL support file uploads?

**Answer:**
Not in official spec.
Unofficial standard: **GraphQL Multipart Request Spec**.
Uses `Upload` scalar.
Payload is `multipart/form-data`. Map points file variables to binary parts.

---

### Question 75: What is graphql-upload?

**Answer:**
Node.js middleware that implements the Multipart Spec.
Processes the stream and injects `fs.ReadStream` into the resolver args.

---

### Question 76: How do you implement authentication in GraphQL?

**Answer:**
**Context!**
Parse JWT in the server setup (e.g., `server.use`).
Pass `user` object to context.
Resolvers use `context.user`.

**Code:**
```javascript
const context = ({ req }) => {
  const token = req.headers.authorization;
  const user = verify(token);
  return { user };
};
```

---

### Question 77: What is a context function in Apollo Server?

**Answer:**
(See Q76). Function that builds the context object for every request.

---

### Question 78: What is schema introspection and how do you disable it?

**Answer:**
Querying `__schema`.
Disable in production to prevent attackers from learning your graph.
`new ApolloServer({ introspection: false })`.

---

### Question 79: How do you rate limit queries in GraphQL?**

**Answer:**
Use `@rateLimit` directive libraries.
Or Cost Analysis (Assign points to fields).

---

### Question 80: What is the use of `@client` directive in Apollo?

**Answer:**
Marks a field as local-only. Apollo Client resolves it from cache/local resolvers, does not send to server.

---

### Question 81: What are the common security concerns with GraphQL?**

**Answer:**
1.  **DoS:** Deeply nested queries (`A.b.a.b...`).
2.  **Auth:** Leaking data in edges.
3.  **Introspection:** Exposing internal structure.

---

### Question 82: How do you prevent query overfetching?

**Answer:**
(Server side protection).
Prevent fetching *too much* data (Resource Exhaustion).
Limit `first: 100` max.

---

### Question 83: What is query complexity analysis in GraphQL?

**Answer:**
Assign cost.
`User: 1`. `posts: 5`. `nested: 10`.
Query Cost = Sum(fields).
If Cost > 1000, reject.

---

### Question 84: What is depth limiting and how is it used?

**Answer:**
Middleware checks indentation depth of query AST.
`graphql-depth-limit(10)`.
Rejects if query is too deep.

---

### Question 85: How do you prevent denial-of-service (DoS) in GraphQL?

**Answer:**
Combination of:
Complexity Limit + Depth Limit + Rate Limit + Timeouts.

---

### Question 86: How do you implement rate limiting in GraphQL?

**Answer:**
Sliding window based on IP or User.
Or "Token Bucket" based on Complexity points (User has 1000 points/min).

---

### Question 87: How do you log and monitor GraphQL queries?

**Answer:**
Apollo Studio.
Or `winston` logging plugin that logs `operationName`, `variables` (sanitized), and `duration`.

---

### Question 88: What is a GraphQL gateway?

**Answer:**
The single entry point in a Federated architecture.
It parses the query, breaks it into plans for subgraphs, executes sub-calls, merges results.

---

### Question 89: How do you handle breaking changes in GraphQL?

**Answer:**
Avoid them.
Use **Evolution**.
1.  Add new field `nameV2`.
2.  Deprecate `name`.
3.  Wait.
4.  Remove `name`.

---

### Question 90: How do you scale a GraphQL server?

**Answer:**
It is stateless (usually).
Run behind Nginx load balancer.
Scale horizontally (Add more Node pods).
Caching is key to reduce DB load.

---

### Question 91: How do you cache GraphQL responses?

**Answer:**
Difficult because it is POST.
**Apollo Server Cache Control:** Hints (`@cacheControl(maxAge: 60)`).
CDN or Gateway caches result based on Query Hash (GET requests preferred for caching).

---

### Question 92: What tools help debug GraphQL APIs?

**Answer:**
GraphiQL. Apollo Studio Explorer (Trace view is amazing).

---

### Question 93: What is Apollo Studio and how is it used?**

**Answer:**
SaaS for GraphQL.
Schema Registry. Usage Reporting. Performance Tracing. Schema Checks (CI).

---

### Question 94: How do you handle large nested queries efficiently?

**Answer:**
`@defer`.
Return root data fast. Stream nested data later.

---

### Question 95: What is schema federation vs schema stitching?

**Answer:**
*   **Stitching:** Manual. "Take Schema A and B and make C". Good for wrapping legacy APIs.
*   **Federation:** Declarative. "Service A extends User from Service B". Standard for microservices.

---

### Question 96: How do you modularize a GraphQL server codebase?

**Answer:**
Folders: `modules/user/`, `modules/product/`.
Each has `typeDefs.graphql` and `resolvers.js`.
Loader script imports all and arrays them for Schema constructor.

---

### Question 97: What are some common anti-patterns in GraphQL?

**Answer:**
1.  **God Object:** One type having links to everything.
2.  **No ID:** Returning objects without IDs (breaks caching).
3.  **JSON Scalar:** Returning generic JSON blob (defeats type safety).

---

### Question 98: How do you handle GraphQL subscriptions in production?

**Answer:**
Don't use In-Memory PubSub (doesn't scale across pods).
Use **Redis PubSub**.
Ensure Load Balancer supports long-lived WebSockets.

---

### Question 99: What’s the best way to paginate results in GraphQL?

**Answer:**
**Cursor Pagination** (Relay).
`first: 10, after: "cursor_hash"`.
Stable even if items inserted/deleted.

---

### Question 100: What are the benefits and drawbacks of using GraphQL in microservices?

**Answer:**
*   **Pros:** Frontend gets one API. No need to query 5 services manually.
*   **Cons:** Gateway is a single point of failure and bottleneck. Complex to manage distributed schema.
