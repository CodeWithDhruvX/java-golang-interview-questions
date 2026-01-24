## 🟢 Testing, Evolution & Governance (Questions 251-300)

### Question 251: How do you deprecate an enum value?

**Answer:**
Standard `@deprecated` directive applies to Enum Values.
```graphql
enum Role {
  USER
  ADMIN
  SUPER_ADMIN @deprecated(reason: "Use ADMIN with scope")
}
```

---

### Question 252: How do you design for nullability?

**Answer:**
**Best Practice:** Nullable by default.
Allows partial failure. If `User.avatar` service fails, we can still return `User.name`.
Only make fields Non-Null (`!`) if the object is meaningless without them (e.g., `id`).

---

### Question 253: How do you share types between Input and Output?

**Answer:**
You cannot. The spec separates `input` and `type`.
You must duplicate them (or use codegen tools to scaffold them).
`input UserInput { name: String }` vs `type User { name: String }`.

---

### Question 254: How do you group related fields?

**Answer:**
Namespace pattern.
Instead of `userAddressCity`, `userAddressZip`, use:
`user { address { city, zip } }`.
Keeps the root object clean.

---

### Question 255: How do you handle arguments on fields?

**Answer:**
Any field can have arguments.
`User.avatar(size: Int)`.
`User.posts(status: PUBLISHED)`.
Resolvers use these `args` to filter/transform the return data.

---

### Question 256: What is the Global Object Identification specification?

**Answer:**
(Relay).
Every object must have an `id: ID!`.
That ID must be globally unique across the entire graph (often base64 encoded `Type:DB_ID`).
Enables `node(id)` refetching.

---

### Question 257: How do you handle schema documentation?

**Answer:**
Use description strings (Markdown supported).
`"The user's public profile"` above the type definition.
Tools show this in tooltips.

---

### Question 258: How do you design for filterable lists?

**Answer:**
Use a rich Input Object `filter`.
`users(filter: { name_contains: "A", age_gt: 18 })`.
Allows extensible criteria without exploding argument count.

---

### Question 259: How do you model errors in the schema (Result Types)?

**Answer:**
Use Unions for outcomes.
```graphql
union CreateResult = User | InvalidInputError | DatabaseError
mutation { createUser { ... on User {} ... on Error {} } }
```

---

### Question 260: How do you handle circular dependencies in schema files?

**Answer:**
SDL files can reference types defined in other files freely.
In code (Resolvers), use `mergeTypeDefs` so load order doesn't matter.

---

### Question 261: What is "Schema Diffing"?

**Answer:**
Comparing Schema V1 vs V2.
Identifying changes: Added fields (Safe), Removed fields (Breaking).
`graphql-inspector` does this automatically in CI.

---

### Question 262: How do you mock GraphQL requests in frontend tests?

**Answer:**
**MockedProvider** (Apollo).
Define an array of `mocks`: `[{ request: { query: Q, variables: V }, result: { data: D } }]`.
Or `msw` handler to intercept the HTTP POST.

---

### Question 263: What implies a "Breaking Change" in GraphQL?

**Answer:**
1.  Removing a field `User.name`.
2.  Changing Scalar `Int` -> `String`.
3.  Adding a required argument `getUser(newArg: String!)`.
4.  Removing an enum value.

---

### Question 264: How do you verify backward compatibility?

**Answer:**
Automated Check.
Validate the New Schema against the Old Schema (Master branch).
If breaking changes found, block merge unless override approved.

---

### Question 265: How do you test permissions?

**Answer:**
Integration Tests.
1.  Set `context.user = NORMAL_USER`.
2.  Run Query `adminData`.
3.  Assert `errors[0].message` == "Forbidden" (or data is null).

---

### Question 266: What is a specialized GraphQL client for testing?

**Answer:**
You don't need a heavy client.
Use `supertest` to POST to `/graphql`.
Or execute the schema directly `graphql({ schema, source: query })` to bypass HTTP (faster).

---

### Question 267: How do you generate mock data from schema?

**Answer:**
`addMocksToSchema`.
Automatically fills `String` with "Hello World", `Int` with random numbers.
Custom mocks: `Int: () => 42`.

---

### Question 268: How do you test subscriptions?

**Answer:**
Use a WebSocket client in test.
Subscribe.
Trigger Mutation (via HTTP or separate client).
Assert message received on WS.

---

### Question 269: How do you validate query depth in tests?

**Answer:**
Write a test case with a highly nested query string.
Expect the server to return 400 or a specific ValidationError.

---

### Question 270: How do you test custom scalars?

**Answer:**
Unit test the Scalar implementation.
Input: `"invalid-date"` -> Expect parser to throw.
Input: `"2020-01-01"` -> Expect Date object.

---

### Question 271: How do you test N+1 problem detection?

**Answer:**
Count SQL queries.
Query 1 item. (Expect 2 SQL calls).
Query 10 items. (Expect 2 SQL calls).
If 11 calls, fail test.

---

### Question 272: What is property-based testing in GraphQL?

**Answer:**
Libraries like `fast-check`.
Generate random valid GraphQL queries.
Send to server.
Verify server doesn't crash (500).

---

### Question 273: How do you snapshot test a schema?

**Answer:**
Dump `print(schema)` to `schema.graphql` file.
Commit it.
If a PR changes the generated string, the snapshot test fails (shows diff). Requires review.

---

### Question 274: How do you test Federation?

**Answer:**
Spin up a local Gateway + Mocks for Subgraphs.
Verify the Gateway can compose the supergraph (no key violations).
Run queries that span multiple subgraphs.

---

### Question 275: How do you test error formatting?

**Answer:**
Trigger a forced error.
Assert response matches `{ errors: [{ message, extensions: { code } }] }`.
Ensure stack trace is hidden in "production" mode tests.

---

### Question 276: How do you test directive logic?

**Answer:**
Create a dummy schema using the directive: `query { hello @upper }`.
Execute.
Assert result is "HELLO".

---

### Question 277: How do you load test a subscription server?

**Answer:**
Tool `artillery-plugin-graphql`.
Ramp up 10k connections.
Broadcast.
Measure time-to-receive on last client.

---

### Question 278: How do you test resolver timeouts?

**Answer:**
Mock the database call to delay 10s.
Set server timeout to 2s.
Assert response contains "TimeoutError".

---

### Question 279: How do you ensure test isolation?

**Answer:**
Transaction Rollback.
Start Transaction -> Run Test -> Rollback.
Or Drop/Recreate In-Memory DB (SQLite) per test file.

---

### Question 280: How do you test for breaking changes in clients?**

**Answer:**
**Operation Registry Check.**
Check breaking change against *registered queries*.
If field `User.bio` is removed, but no registered query asks for `bio`, it's safe (technically "breaking" but impact-free).

---

### Question 281: What is the "Dangerous" change warning?

**Answer:**
Inspector warning for changes that *might* break clients.
e.g., Changing a nullable argument to nullable (Safe) vs adding a new nullable argument (Safe). But changing default values is Dangerous.

---

### Question 282: How do you manage schema ownership?

**Answer:**
Annotate types with `@owner(team: "identity")`.
Lint tools check that only Identity Team modifies `User.graphql`.

---

### Question 283: What is a Field Policy?

**Answer:**
(Apollo Client).
Defines how a field is read from cache.
`keyArgs: false` (all args share same cache).
`merge(existing, incoming)` (pagination logic).

---

### Question 284: How do you promote a schema from Dev to Prod?

**Answer:**
Schema Registry Pipelines.
`main` branch pushes to "Staging" graph.
After integration tests pass, push to "Prod" graph.

---

### Question 285: How do you handle "Zombie" fields (unused)?

**Answer:**
Apollo Studio "Unused Fields" report.
Filter by last 90 days.
If 0 requests, delete safely.

---

### Question 286: What is the "One Graph" principle?**

**Answer:**
Instead of `api.company/billing` and `api.company/users`, expose ONE endpoint `api.company/graphql` that covers everything.

---

### Question 287: How do you handle third-party vendor schemas?

**Answer:**
Don't stitch them directly (fragile).
Create an internal Type that maps to the Vendor API.
Your resolver calls Vendor API.
Decouples your schema from theirs.

---

### Question 288: How do you enforce naming conventions?

**Answer:**
Linter (`graphql-schema-linter`).
Rules: `fields-are-camel-case`, `types-are-pascal-case`, `enum-values-are-caps`.

---

### Question 289: How do you handle "God Objects" (e.g., massive User type)?

**Answer:**
Federation.
Identity Service defines `User { id username }`.
Billing Service extends `User { creditCards }`.
Notifications Service extends `User { preferences }`.
Code is distributed.

---

### Question 290: How do you coordinate cross-team schema changes?

**Answer:**
Design Reviews.
Propose SDL in a Pull Request.
Consumers (Frontend) review the shape before backend builds it.

---

### Question 291: What is "Schema-First" development in teams?

**Answer:**
Agree on Contract first.
Frontend mocks result. Backend builds impl.
Parallel work. No blocking.

---

### Question 292: How do you document "Why" a field exists?

**Answer:**
Schema Comments.
`""" Returns the user's avatar URL. If null, use default placeholder. """`

---

### Question 293: How do you handle multi-environment schemas?

**Answer:**
Feature Flags in Schema.
If env=PROD, strip `@experimental` fields from the introspection result given to clients.

---

### Question 294: What is the role of the "Graph Champion"?

**Answer:**
The person who ensures the graph stays clean, consistent, and performant. Stops "REST-over-GraphQL" anti-patterns.

---

### Question 295: How do you align GraphQL types with DB models?

**Answer:**
**Don't**.
DB: `user_tbl` (snake_case, foreign keys).
Graph: `User` (camelCase, object references).
Use mappers/connectors.

---

### Question 296: How do you handle sensitive PII in schema?

**Answer:**
Custom scalar `PIIString` or directive `@sensitive`.
Logs automatically redact fields with this type/directive.

---

### Question 297: How do you audit schema complexity over time?

**Answer:**
Track "Schema Size" (number of fields).
If growing exponentially, review governance. Maybe purge old fields.

---

### Question 298: How do you handle multiple "User" types (External, Internal, Admin)?

**Answer:**
Interface `User`.
Types `Customer`, `Employee`, `Admin` implement `User`.
Shared fields in Interface. Specific fields in concrete types.

---

### Question 299: What is the benefit of a "Linting" step for schema?

**Answer:**
Catches syntax errors early.
Enforces best practices (descriptions required).
Prevents trivial bike-shedding in PR reviews.

---

### Question 300: How do you ensure the graph remains "Product-Centric"?

**Answer:**
Ask "What does the UI need?" not "What columns are in the table?".
Group data by User Intent (`checkout`), not Table Name (`order_transaction_logs`).
