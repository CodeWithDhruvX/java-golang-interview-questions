## 🟢 Internals, Integrations & DX (Questions 301-350)

### Question 301: How does GraphQL parse and execute a query under the hood?

**Answer:**
1.  **Lexing/Parsing:** Converts string to AST (Abstract Syntax Tree).
2.  **Validation:** Checks AST against Schema (types existence, argument validity).
3.  **Execution:** Traverses AST (BFS/DFS). Calls resolvers. Coerces results.
4.  **Response:** Shapes JSON output.

---

### Question 302: What is the GraphQL type system made of internally?**

**Answer:**
Objects representing Types (`GraphQLObjectType`, `GraphQLScalarType`).
Config objects defining Fields, Thunks (lazy evaluation for recursive types), and Interfaces.

---

### Question 303: How does GraphQL resolve a deeply nested query step-by-step?

**Answer:**
Root Query Resolver -> Returns `User` object.
Engine passes `User` to `User.posts` resolver -> Returns `[Post]`.
Engine passes each `Post` to `Post.comments` resolver -> Returns `[Comment]`.
Merges up the tree.

---

### Question 304: What is the role of the execution context in a GraphQL operation?

**Answer:**
Singleton object passed to `execute()`.
Contains: `schema`, `fragments`, `rootValue`, `contextValue`, `operation`, `variableValues`.
Passed down to every resolver.

---

### Question 305: How is a GraphQL schema represented in memory?

**Answer:**
A `GraphQLSchema` instance.
Contains maps of `typeMap` (all types by name), `directives`, `queryType`, `mutationType`.
Lookup is O(1).

---

### Question 306: What are abstract syntax trees (AST) in GraphQL?

**Answer:**
JSON-like tree structure of the query.
`{ kind: "OperationDefinition", operation: "query", selectionSet: { ... } }`.
Resolvers receive `info` which contains the node's AST.

---

### Question 307: How does `graphql-js` implement validation rules?

**Answer:**
Visitor Pattern.
It walks the AST.
Visitors (Rules) triggers on `enter/leave` node.
e.g., `KnownTypeNames` visitor checks if type exists in Schema.

---

### Question 308: How does GraphQL prevent infinite recursion in queries?

**Answer:**
It *doesn't* automatically.
If schema allows `A -> B -> A`, client can write infinite query.
Server must implement **Max Depth Analysis** to reject it.

---

### Question 309: What are execution phases in GraphQL?

**Answer:**
1.  **Coerce Variables.**
2.  **Execute Operation.**
3.  **Complete Value** (resolve fields, handle nulls, handle lists).
4.  **Serialize.**

---

### Question 310: How does GraphQL treat circular fragments internally?

**Answer:**
Validator detects cycles.
`fragment A on User { ...B } fragment B on User { ...A }`.
Parser/Validator checks Fragment Spreads and throws validation error "Cycle detected".

---

### Question 311: How are variables resolved and substituted during execution?

**Answer:**
Engine reads `variableDefinitions` from Operation.
Coerces provided JSON variables to expected Types.
Passes coerced values to `execute`.
Resolvers access them via `args`.

---

### Question 312: What is the difference between execution phase and resolver logic?

**Answer:**
*   **Execution Phase:** Generic framework code (handling lists, null bubbling, merging fields).
*   **Resolver Logic:** User code (fetching data).

---

### Question 313: How is a GraphQL request batched internally by Apollo Server?

**Answer:**
Apollo supports "Batch HTTP" (Array of Queries in POST body).
Apollo loops over array, executes them in parallel (Promise.all), returns JSON array.
It is *not* one single execution context.

---

### Question 314: What is `graphql-parse-resolve-info` and how is it useful?

**Answer:**
Library to parse `info` argument.
Tells you *exactly* what sub-fields the client asked for.
Useful for Look-ahead optimization (SQL JOINs).

---

### Question 315: How can you hook into the execution flow of GraphQL?

**Answer:**
`extensions` class in `graphql-js`.
Or Plugins in Apollo (`requestDidStart`, `executionDidStart`, `willResolveField`).

---

### Question 316: What is the difference between schema-first and resolver-first approaches?

**Answer:**
*   **Schema-First:** String parsing. Fast iteration on API design.
*   **Resolver-First (Code-First):** Programmatic schema construction. Better for Refactoring/Typescript integration.

---

### Question 317: What are the disadvantages of schema-first design?**

**Answer:**
Stringly typed strings.
Resolvers might not match Schema (Runtime error vs Compile time).
IDE support inside string templates is sometimes limited.

---

### Question 318: How do directives impact the execution engine?

**Answer:**
Standard directives (`@skip`) are handled by the executor.
Custom directives often wrap the resolver function at *Schema Construction Time*, adding logic (middleware).

---

### Question 319: What is a GraphQL execution strategy?

**Answer:**
`executeSerial` (Mutations): Wait for prev to finish before starting next.
`executeFields` (Queries): Start all promises, wait for all (Parallel).

---

### Question 320: What are synchronous vs asynchronous resolvers in detail?

**Answer:**
*   **Sync:** `return user.name`. AST traversal continues immediately.
*   **Async:** `return fetch(...)`. AST traversal pauses branch, continues others.

---

### Question 321: How do you integrate GraphQL with PostgreSQL?

**Answer:**
Use Prisma or TypeORM in Resolvers.
Map GQL Args -> SQL WHERE clause.
Return SQL Rows -> Match GQL Type keys.

---

### Question 322: How does Hasura auto-generate GraphQL from Postgres?

**Answer:**
Reads `information_schema`.
Tables -> Types.
Foreign Keys -> Nesting/Relationships.
Columns -> Fields.
Builds executable schema in memory.

---

### Question 323: How do you integrate MongoDB with GraphQL?

**Answer:**
Mongoose Models return Promises.
GraphQL handles Promises nicely.
`_id` (ObjectId) needs a custom scalar or toString() mapping to `ID`.

---

### Question 324: What is Prisma’s role in GraphQL backends?

**Answer:**
Type-safe Database Client.
Replaces boilerplate SQL in resolvers.
`db.user.findMany()` returns objects matching the GraphQL Type definition usually.

---

### Question 325: How do you expose REST APIs via GraphQL Mesh?

**Answer:**
Mesh config:
```yaml
sources:
  - name: Wiki
    handler:
      openapi:
        source: wiki-swagger.json
```
Mesh generates Resolvers that call the REST endpoints.

---

### Question 326: How do you integrate Elasticsearch with GraphQL?

**Answer:**
Types: `type SearchResult { hits: [Hit] }`.
Resolver: `client.search({ q: args.query })`.
Allows typesafe search interface.

---

### Question 327: What is Neo4j GraphQL Library and how does it work?

**Answer:**
Augments schema with `@cypher` directives.
Transpiles GraphQL selection set -> Cypher Query.
Fetches graph data efficiently.

---

### Question 328: How do you consume GraphQL in Flutter?

**Answer:**
`graphql_flutter` package.
`Query(options: QueryOptions(...), builder: (result) => ...)`
Widgets-based approach (like React hooks).

---

### Question 329: How do you implement GraphQL with FastAPI or Django?

**Answer:**
**Strawberry** (Python).
Data classes with decorators `@strawberry.type`.
Generates schema. Serves via ASGI.

---

### Question 330: How do you use gRPC services inside GraphQL resolvers?

**Answer:**
Node gRPC client.
Query `user(id: 1)` -> Resolver calls `grpcUserClient.GetUser({ id })`.
Maps Protobuf response to GraphQL Object.

---

### Question 331: What are pros/cons of wrapping SOAP in GraphQL?

**Answer:**
**Pro:** Clean JSON API for frontend.
**Con:** SOAP complexity (XML namespaces) makes wrapper logic brittle.

---

### Question 332: How do you expose Kafka streams via GraphQL Subscriptions?

**Answer:**
`AsyncIterator`.
On Kafka Message -> `pubsub.publish("TOPIC", data)`.
Resolver: `subscribe: () => pubsub.asyncIterator("TOPIC")`.

---

### Question 333: How do you use GraphQL with Firebase?

**Answer:**
Apollo Server in Cloud Functions.
Resolver: `admin.firestore().collection('...').get()`.

---

### Question 334: What are the trade-offs of using GraphQL with FaunaDB?

**Answer:**
Native GQL support means no resolver code (Fast).
But custom logic requires FQL (Learning curve).

---

### Question 335: How do you expose GraphQL APIs from an AWS Lambda?

**Answer:**
API Gateway (Proxy) -> Lambda.
Lambda runs Apollo Server.
Cold starts can be an issue.

---

### Question 336: How do you integrate Auth0 with GraphQL APIs?

**Answer:**
Auth0 validates User. Issues Access Token.
GraphQL API validates Token (Middleware).
Permissions claim in Token -> Directives in Schema.

---

### Question 337: How do you use GraphQL with GitHub’s API?

**Answer:**
Use their public GraphQL Endpoint.
Or schema stitching to merge it into your company graph (`viewer { githubRepos { ... } }`).

---

### Question 338: How can GraphQL federate multiple SaaS APIs?

**Answer:**
Each SaaS wrapped in a "Subgraph" (Adapter).
Stripe Subgraph, Contentful Subgraph.
Gateway merges them.

---

### Question 339: What is the use of OpenAPI-to-GraphQL converters?

**Answer:**
Quick migration.
Take Swagger JSON. Get GQL Schema + Resolvers automatically.
Good for wrapping legacy apps.

---

### Question 340: How do you expose IoT or edge data using GraphQL?

**Answer:**
Subscription-heavy.
Device -> MQTT -> GraphQL Server -> Subscription -> Frontend.

---

### Question 341: How do you set up a GraphQL monorepo for multiple teams?

**Answer:**
Workspaces (Yarn/NPM).
`packages/schema-user`, `packages/schema-product`.
Gateway package imports them.
Shared `eslint-config-graphql`.

---

### Question 342: How do you implement GitHub Actions for GraphQL linting?**

**Answer:**
Action runs `graphql-inspector`.
Checks for breaking changes vs `main`.
Checks for lint rules.
Comments on PR if errors found.

---

### Question 343: What are good ESLint rules for GraphQL in frontend code?

**Answer:**
`graphql/template-strings`.
Validates queries against local schema file.
Catches typos `useQuery(gql'{ user { nmae } }')` instantly in IDE.

---

### Question 344: How do you set up Git pre-commit hooks for GraphQL validation?

**Answer:**
`lint-staged`.
`*.graphql`: `graphql-schema-linter`.
Prevents committing bad SDL.

---

### Question 345: How do you build a local schema explorer for devs?

**Answer:**
Embed `GraphiQL` in the dev server.
Or host `Spectaql` docs locally.

---

### Question 346: What is the best way to mock remote GraphQL schemas for local dev?

**Answer:**
Download introspection JSON.
`graphql-tools` -> `addMocksToSchema`.
Start local server. Frontend devs code against mock without backend running.

---

### Question 347: How do you auto-document GraphQL queries in frontend PRs?

**Answer:**
GitHub Bot.
Extracts GQL operations.
Runs `graphql-inspector`.
Posts comment: "This query fetches 20 fields. Complexity: 50."

---

### Question 348: What CLI tools are essential for GraphQL developers?

**Answer:**
*   `rover` (Apollo).
*   `graphql-codegen` (Types).
*   `get-graphql-schema` (Download schema).

---

### Question 349: What is `graphql-cli` and how do you use it?

**Answer:**
Unified toolset. `graphql init`, `graphql diff`.
Configured via `.graphqlconfig`.

---

### Question 350: How do you use `.graphqlconfig` in multi-project setups?

**Answer:**
Define projects:
```yaml
projects:
  app:
    schema: app.graphql
  admin:
    schema: admin.graphql
```
IDE plugins use this to autocomplete correctly per file.
