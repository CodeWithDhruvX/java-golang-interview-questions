# GraphQL Interview Questions & Answers

## 🔹 1. Basics & Core Concepts (1–20)

**Q1: What is GraphQL?**
GraphQL is a query language for APIs and a runtime for fulfilling those queries with your existing data. It allows clients to request exactly the data they need, making it efficient and flexible.

**Q2: How does GraphQL differ from REST?**
REST typically has multiple endpoints returning fixed data structures. GraphQL has a single endpoint where clients specify the shape of the data they need. GraphQL prevents over-fetching and under-fetching.

**Q3: What are the main features of GraphQL?**
Declarative data fetching, single endpoint, strongly typed schema, no over-fetching/under-fetching, and real-time updates via Subscriptions.

**Q4: Explain the GraphQL type system.**
The type system defines the capabilities of the GraphQL server. It includes scalars (Int, String), objects (custom types), interfaces, unions, enums, and input objects.

**Q5: What is a GraphQL schema?**
A schema is a contract between the client and server. It defines the types of data that can be queried and the relationships between them using Schema Definition Language (SDL).

**Q6: What are queries in GraphQL?**
Queries are requests made by the client to fetch data from the server. They are analogous to GET requests in REST but are more flexible.

**Q7: What is a mutation in GraphQL?**
Mutations are used to modify data on the server (create, update, delete). They are analogous to POST, PUT, DELETE requests in REST.

**Q8: What is a subscription in GraphQL?**
Subscriptions allow clients to listen for real-time updates from the server. They are typically implemented using WebSockets.

**Q9: How does GraphQL handle versioning?**
GraphQL generally avoids versioning. Since clients request only specific fields, new fields can be added without breaking existing queries. Deprecated fields can be marked using `@deprecated`.

**Q10: What are resolvers in GraphQL?**
Resolvers are functions responsible for fetching the data for a specific field in the schema. Every field in a GraphQL query is backed by a resolver.

**Q11: What is introspection in GraphQL?**
Introspection allows clients to query the schema itself to discover available types, fields, and queries. It is useful for tools like GraphiQL.

**Q12: What is the role of `__typename` in GraphQL?**
`__typename` is a meta-field that returns the name of the object type being queried. It is useful for client-side caching and handling fragments.

**Q13: How do you define custom scalar types in GraphQL?**
You can define custom scalars (e.g., `Date`, `JSON`) in the schema and provide serialization, parsing, and validation logic in the resolvers.

**Q14: What is the difference between input and output types?**
Output types (Object types) are returned in responses. Input types (`input`) are used as arguments in mutations or queries to pass complex objects.

**Q15: How are enums used in GraphQL?**
Enums defines a set of allowed values for a field. They ensure that the data adheres to a specific set of constants.

**Q16: Explain the concept of nullability in GraphQL.**
By default, fields in GraphQL are nullable. You can enforce a non-null value by adding an exclamation mark `!` (e.g., `String!`).

**Q17: What is the purpose of fragments?**
Fragments are reusable units of logic in queries. They allow you to define a set of fields once and reuse them in multiple queries.

**Q18: What is the difference between inline fragments and named fragments?**
Named fragments have a name and can be reused. Inline fragments are used directly within a query, often for Union or Interface types (e.g., `... on User`).

**Q19: Can you explain how aliases work in GraphQL?**
Aliases allow you to rename the result field of a query. This is useful when fetching the same field with different arguments in a single request.

**Q20: What are variables in GraphQL?**
Variables allow you to pass dynamic arguments to queries and mutations without modifying the query string itself. They are sent separately in a JSON object.

---

## 🔹 2. Schema & Type System (21–40)

**Q21: What is a root type in GraphQL?**
Root types are the entry points for the schema: `Query`, `Mutation`, and `Subscription`.

**Q22: How do you structure a GraphQL schema?**
Schemas are structured by defining types, relationships, and root operations. Large schemas are often split into modules or multiple files.

**Q23: What are object types in GraphQL?**
Object types represent a kind of object you can fetch from your service, and what fields it has (e.g., `type User { id: ID, name: String }`).

**Q24: How do you create relationships between types?**
Relationships are created by having a field in one object type return another object type (e.g., `type User { posts: [Post] }`).

**Q25: How are interfaces used in GraphQL?**
Interfaces define a list of fields that a type must include. Object types `implement` interfaces to guarantee they have those fields.

**Q26: What are unions in GraphQL?**
Unions allow a field to return one of several distinct object types. Unlike interfaces, unions do not guarantee common fields.

**Q27: How does GraphQL handle recursive types?**
GraphQL supports recursive types natively. A type can reference itself (e.g., `type Category { subCategories: [Category] }`).

**Q28: What are directives in GraphQL?**
Directives provide a way to modify the behavior of valid types or fields. Common directives include `@include`, `@skip`, and `@deprecated`.

**Q29: What is the `@include` and `@skip` directive?**
`@include(if: Boolean)` includes a field if the argument is true. `@skip(if: Boolean)` skips a field if the argument is true.

**Q30: What is the `@deprecated` directive and how do you use it?**
`@deprecated(reason: "...")` marks a field as deprecated, signaling to clients that it should not be used, without breaking types.

**Q31: How do you write custom directives?**
Custom directives are defined in the schema (e.g., `directive @auth on FIELD_DEFINITION`) and implemented in the server logic to modify execution.

**Q32: How do you ensure schema validation?**
GraphQL servers automatically validate the schema against the SDL rules at startup. Tools like `graphql-inspector` can also be used.

**Q33: How do you modularize a GraphQL schema?**
By splitting type definitions into separate files and merging them using tools like `graphql-tools` (`mergeTypeDefs`, `mergeResolvers`).

**Q34: What is schema stitching?**
Schema stitching is the process of combining multiple GraphQL schemas into a single executable schema.

**Q35: What is schema delegation?**
Delegation allows a resolver to forward a query to another schema or sub-schema, often used in stitching.

**Q36: What is the difference between schema-first and code-first approaches?**
Schema-first starts with writing SDL (`type.graphql`). Code-first generates the schema from programming language constructs (e.g., TypeScript classes).

**Q37: What tools are used to build GraphQL schemas?**
`graphql-js`, `Apollo Server`, `Nexus` (code-first), `TypeGraphQL`, `Hasura`.

**Q38: How do you document a GraphQL schema?**
You can add comments (strings inside `""`) or block strings (`"""`) above types and fields in the SDL, which show up in introspection tools.

**Q39: What is the difference between SDL and introspection schema?**
SDL is the human-readable string format. Introspection schema is the JSON result of a generic introspection query used by tools.

**Q40: How can you expose multiple versions of a schema?**
GraphQL encourages evolving a single version. If strictly needed, you can run separate endpoints for different schema versions.

---

## 🔹 3. Resolvers & Execution (41–50)

**Q41: How does a resolver function work?**
A resolver corresponds to a field. When the field is queried, the resolver function is called to fetch the value (from DB, API, etc.).

**Q42: What is the resolver signature?**
Typically `(parent, args, context, info)`.
- `parent`: The result of the previous resolver.
- `args`: Arguments provided to the field.
- `context`: Shared object for all resolvers (e.g., auth info).
- `info`: Execution state and schema details.

**Q43: What is the difference between field-level and type-level resolvers?**
Field-level resolvers handle specific fields (e.g., `User.name`). Type-level resolvers usually refer to the default behavior or the root object resolution.

**Q44: What is the parent (or root) argument in resolvers?**
It is the object returned by the parent field's resolver. For root queries, it is often `undefined` or a root value.

**Q45: How can you resolve nested fields?**
GraphQL executes resolvers top-down. The resolver for a parent field returns an object, and the resolvers for the child fields receive that object as `parent`.

**Q46: How can you handle errors in resolvers?**
Throwing an error in a resolver will result in a `null` value for that field (unless non-nullable) and an error entry in the `errors` array of the response.

**Q47: What are resolver chains and how do they work?**
The sequence of resolvers called to populate a deep query. Root Query -> Field A -> Field B. Each passes its result down.

**Q48: How do you access context in resolvers?**
The `context` argument is passed to every resolver. It's used to store global data like user authentication, database connections, or data loaders.

**Q49: How do you secure resolver logic?**
By checking permissions inside the resolver (using `context.user`) or using middleware/directives to authorize access before the resolver runs.

**Q50: How do you batch resolvers?**
Using **DataLoader**. It coalesces individual fetch requests from multiple resolvers into a single batch request to the database (solving the N+1 problem).

---

## 🔹 3. Resolvers & Execution (Continued) (51–60)

**Q51: What is DataLoader and how does it improve resolver performance?**
DataLoader is a utility that batches and caches database requests. It solves the N+1 problem by collecting keys from multiple resolver calls and fetching them in a single batch query.

**Q52: How do you handle authorization in resolvers?**
Authorization logic is often placed in the business logic layer or wrapped around resolvers (middleware/directives). You check if `context.user` has the required role to access the field.

**Q53: How can you reuse resolver logic across fields?**
You can extract common logic into separate functions or services and call them from multiple resolvers. Higher-order functions can also wrap resolvers.

**Q54: How do you use middleware in resolvers?**
Middleware (like `graphql-middleware`) runs before or after a resolver. It's useful for logging, authentication, and error handling across many fields.

**Q55: How do you test GraphQL resolvers?**
Resolvers are pure functions (mostly). You can unit test them by mocking the `context` and `args` and asserting the return value.

**Q56: How does resolver priority work when multiple fields have the same name?**
Field-level resolvers take precedence over trivial (default) resolution. If a resolver is defined for a field, it is executed; otherwise, the property of the same name on the parent object is returned.

**Q57: How do you handle caching in resolvers?**
Resolvers can check a cache (e.g., Redis) before hitting the DB. However, caching is often better handled at the DataLoader level (per-request) or the CDN/HTTP level (full response).

**Q58: How do you resolve computed fields?**
Computed fields (e.g., `fullName`) don't exist in the DB. The resolver combines data from the `parent` object (e.g., `parent.firstName + ' ' + parent.lastName`) to return the result.

**Q59: How do resolvers affect performance in large schemas?**
Deeply nested queries trigger many resolvers. If not batched (DFS execution), this leads to many database calls. Resolvers should remain lightweight.

**Q60: Can you use async/await in resolvers?**
Yes, resolvers handle Promises natively. You can make the resolver function `async` and `await` database calls or external API fetch requests.

---

## 🔹 4. Advanced Features (61–80)

**Q61: What is GraphQL federation?**
Federation allows you to divide your graph into multiple subgraphs (services) and compose them into a single supergraph. Clients query the gateway, which routes requests to appropriate services.

**Q62: What is Apollo Federation?**
An architecture and set of tools by Apollo for implementing federation. It uses `@key` and `@extends` directives to link types across services.

**Q63: What is Apollo Client and how does it work?**
A comprehensive state management library for JavaScript that enables you to manage both local and remote data with GraphQL. It fetches, caches, and modifies application data.

**Q64: How does Apollo Server differ from other GraphQL servers?**
Apollo Server is an opinionated, production-ready server that integrates well with the Apollo ecosystem (Federation, Studio, Client). It simplifies setup but is specific to Node.js.

**Q65: What is Relay in GraphQL?**
Relay is a JavaScript framework by Facebook for building data-driven React applications. It enforces specific schema patterns (Global Object ID, Connections) for performance and scalability.

**Q66: What is the connection model in Relay?**
A standard way to handle pagination using cursors. It uses `edges`, `node`, and `pageInfo` fields to traverse lists of data.

**Q67: What are edges and nodes in Relay?**
- `Node`: The actual object (data).
- `Edge`: A wrapper around the node that contains the `cursor` and the `node`.

**Q68: What is GraphQL Mesh?**
A tool that allows you to access data from any source (REST, gRPC, SOAP) as a GraphQL API without rewriting legacy services.

**Q69: What are persisted queries?**
A performance optimization where the query string is sent to the server once and saved. Clients then send a hash/ID instead of the full query string, reducing bandwidth.

**Q70: What are automatic persisted queries (APQ)?**
The client optimistically sends the hash. If the server doesn't know it, it asks for the full query once, caches it, and uses the hash for future requests. No build-time step required.

**Q71: What are GraphQL directives and how can they be extended?**
Directives decorate parts of the schema. You can define custom directives (`directive @log on FIELD`) and implement logic in the server to transform execution or the schema.

**Q72: What is client-side schema?**
Allows you to extend your server schema on the client side with local-only fields (using `@client`). Useful for managing local state (like `isLoggedIn`) alongside remote data.

**Q73: What is server-side schema?**
The authoritative schema defined on the backend that maps to your data sources.

**Q74: How does GraphQL support file uploads?**
The GraphQL spec doesn't natively support uploads. It requires the `Upload` scalar and a multipart request specification (often implemented via `graphql-upload`).

**Q75: What is graphql-upload?**
middleware that processes multipart/form-data requests and provides an efficient stream for file uploads in resolvers.

**Q76: How do you implement authentication in GraphQL?**
Authentication (identifying the user) is usually handled via HTTP headers (Bearer token) before GraphQL execution. The user object is then passed to the resolver context.

**Q77: What is a context function in Apollo Server?**
A function that runs for every request. It parses the request (e.g., headers) and returns the context object passed to all resolvers.

**Q78: What is schema introspection and how do you disable it?**
Introspection enables querying the schema structure. In production, it should often be disabled (using `introspection: false`) to prevent exposing internal API details to attackers.

**Q79: How do you rate limit queries in GraphQL?**
You can calculate query complexity (cost analysis) or depth limit. If a query exceeds the cost/depth threshold, reject it before execution.

**Q80: What is the use of `@client` directive in Apollo?**
It tells Apollo Client to resolve the field locally (from the cache or local resolvers) instead of sending it to the server.

---

## 🔹 5. Security, Performance, and Best Practices (81–100)

**Q81: What are the common security concerns with GraphQL?**
Over-fetching (DoS), deep nesting attacks, introspection exposure, leaky error messages, and broken authorization.

**Q82: How do you prevent query overfetching?**
This is a client-side benefit, but server-side you handle "DoS via large queries" by limiting complexity or depth.

**Q83: What is query complexity analysis in GraphQL?**
Assigning "points" to fields (e.g., Objects=5, Scalars=1). Before executing, calculate total points. If > Limit, reject query.

**Q84: What is depth limiting and how is it used?**
Limiting how deep a query can be nested (e.g., max depth 5). Prevents cyclic queries like `author { posts { author { posts ... } } }`.

**Q85: How do you prevent denial-of-service (DoS) in GraphQL?**
Timeouts, Depth Limiting, Complexity Analysis, Rate Limiting, and size limits on the request payload.

**Q86: How do you implement rate limiting in GraphQL?**
Unlike REST (per endpoint), GraphQL needs rate limiting based on query cost or simple request count per IP/User over time.

**Q87: How do you log and monitor GraphQL queries?**
Log unique operation names and hashes. Use tools like Apollo Studio or logging plugins to track latency, errors, and most used fields.

**Q88: What is a GraphQL gateway?**
An entry point server (like Apollo Gateway) that aggregates multiple underlying GraphQL services (subgraphs) into one unified API.

**Q89: How do you handle breaking changes in GraphQL?**
Avoid them. Use `@deprecated`. Add new fields instead of renaming. If removal is necessary, track usage and communicate with clients.

**Q90: How do you scale a GraphQL server?**
Horizontal scaling (stateless servers). For performance, use caching (CDN, Redis), DataLoaders, and efficient database indexing.

**Q91: How do you cache GraphQL responses?**
GraphQL is POST by default (hard to cache at HTTP level). You can use GET for queries (with APQ) or cache directives (e.g., `@cacheControl`) for CDN integration.

**Q92: What tools help debug GraphQL APIs?**
GraphiQL, GraphQL Playground, Apollo Studio Explorer, Postman.

**Q93: What is Apollo Studio and how is it used?**
A cloud platform for monitoring, managing, and collaborating on your graph. Provides metrics, schema checks, and an Explorer.

**Q94: How do you handle large nested queries efficiently?**
Use `DataLoader` to batch DB calls. Use `@defer` (experimental) to stream slow parts of the response. Apply depth limits.

**Q95: What is schema federation vs schema stitching?**
Federation is declarative (directives in schema). Stitching is imperative (code that merges schemas). Federation is the modern standard for microservices.

**Q96: How do you modularize a GraphQL server codebase?**
Organize by domain (User, Product) with their own TypeDefs and Resolvers. Merge them at the root.

**Q97: What are some common anti-patterns in GraphQL?**
- N+1 queries (not using DataLoader).
- Exposing DB structure 1:1.
- Giant root query type.
- Ignoring errors (swallowing them).
- Versioning via endpoints.

**Q98: How do you handle GraphQL subscriptions in production?**
Use a specialized PubSub system (Redis PubSub) instead of in-memory. Ensure sticky sessions if using WebSockets across multiple instances (or use a separate WS fleet).

**Q99: What’s the best way to paginate results in GraphQL?**
Cursor-based pagination (Relay style) is preferred over Offset-based because it handles real-time data changes better and is more performant for large sets.

**Q100: What are the benefits and drawbacks of using GraphQL in microservices?**
**Pros**: Unified API, decoupled frontend/backend.
**Cons**: Complexity of Federation, "Distributed Monolith" risk, difficult to trace requests across services.

---

## 🔹 6. Client-Side GraphQL (101–120)

**Q101: How do you integrate GraphQL with React?**
Using a client library like **Apollo Client** or **URQL**. providing a `ApolloProvider` wrapper and using hooks like `useQuery` and `useMutation`.

**Q102: How does Apollo Client cache data?**
It normalizes data by `__typename` and `id` (or `_id`). If the same object is fetched again or updated by a mutation, the cache updates automatically, reflecting across UI components.

**Q103: How do you invalidate or evict Apollo Client cache?**
Using `client.cache.evict({ id: '...' })` or `client.resetStore()` (clears everything). You can also use `garbageCollect()` to remove unreachable objects.

**Q104: What is optimistic UI in Apollo Client?**
A technique where the UI updates immediately with a predicted response *before* the server confirms it. If the server fails, the UI rolls back.

**Q105: How do you handle loading and error states with Apollo Client?**
The `useQuery` hook returns `{ data, loading, error }` objects, allowing you to conditionally render spinners or error messages effortlessly.

**Q106: How do you write local-only fields in Apollo Client?**
By adding a field with the `@client` directive in your query. You resolve it using local resolvers or Type Policies in the Apollo cache configuration.

**Q107: How do you manage state with Apollo Client?**
Apollo Client can act as a global state manager (replacing Redux). You can query local state and remote data in a single query using `@client`.

**Q108: What is Apollo Link and how does it work?**
Apollo Link is a middleware standard for modifying the control flow of GraphQL requests (e.g., adding headers, catching errors, retrying) before they hit the network.

**Q109: How do you use fragments with Apollo Client?**
Define fragments using `gql` and spread them into queries (`...FragmentName`). Useful for colocating data requirements with components.

**Q110: How do you execute a query manually using Apollo Client?**
Using the `useLazyQuery` hook which returns a trigger function, or directly calling `client.query({ ... })`.

**Q111: How do you handle polling in GraphQL queries?**
Pass the `pollInterval` option (in milliseconds) to `useQuery`. Apollo will re-fetch the query repeatedly at that interval.

**Q112: What is `useQuery` vs `useLazyQuery`?**
`useQuery` runs automatically when the component mounts. `useLazyQuery` returns a function to run the query on demand (e.g., on button click).

**Q113: How can you cancel a GraphQL request?**
Using `AbortController` signal passed to the fetch context, or using libraries that support cancellation triggers. Apollo Link can also handle cancellation.

**Q114: What’s the difference between Apollo Client and Relay?**
Apollo is flexible, easy to learn, and community-driven. Relay is highly optimized, opinionated (requires specific schema structure), and scales better for massive Facebook-scale apps.

**Q115: How do you use GraphQL in Angular applications?**
Using `Apollo Angular`, which provides services (`Apollo`) to execute queries and mutations using RxJS Observables.

**Q116: How do you integrate GraphQL with Vue.js?**
Using `Vue Apollo`. It integrates with Vue's reactivity system (Composition API `useQuery` or Options API `apollo` object).

**Q117: How do you manage multiple GraphQL endpoints in a frontend app?**
By initializing multiple `ApolloClient` instances or using `Apollo Link` split logic to route operations to different endpoints based on context.

**Q118: How do you debug client-side GraphQL issues?**
Using **Apollo Client DevTools** browser extension to inspect the cache, active queries, and mutations.

**Q119: How do you persist Apollo cache across page reloads?**
Using `apollo3-cache-persist`, which serializes the cache to `localStorage` or `AsyncStorage` and restores it on app boot.

**Q120: What is the purpose of Apollo DevTools?**
To visualize the normalized cache store, test queries in an embedded GraphiQL, and track active query states/errors within the browser.

---

## 🔹 7. Error Handling & Debugging (121–140)

**Q121: How do you define custom error codes in GraphQL?**
By adding an `extensions` field to the error object returned by the server, containing `code`, `timestamp`, or other metadata (e.g., `"code": "UNAUTHENTICATED"`).

**Q122: What is the format of a GraphQL error response?**
A JSON object with an `errors` array. Each error has `message`, `locations` (line/col), and `path` (field), plus optional `extensions`.

**Q123: How do you distinguish between GraphQL and network errors?**
Network errors (4xx/5xx) mean the server wasn't reached or crashed. GraphQL errors (200 OK) mean the query ran but specific fields failed.

**Q124: What is the best way to log GraphQL errors?**
Use a plugin (like `didEncounterErrors` in Apollo) to log errors to monitoring tools (Sentry, Datadog) including the operation name and variables (sanitized).

**Q125: How do you mask internal errors from clients?**
In production, catch errors in `formatError` and return generic messages ("Internal Server Error") while logging the full stack trace internally.

**Q126: How can you use extensions field in error responses?**
To pass machine-readable details like validation messages, error codes, or trace IDs without polluting the human-readable `message`.

**Q127: How do you create global error handlers for GraphQL?**
Using `Apollo Link` (specifically `onError` link) on the client side to intercept all errors (e.g., global logout on 401).

**Q128: What is GraphQL Shield and how is it used for auth/errors?**
A middleware library for permission validation. It throws authorization errors before resolvers run if rules aren't met.

**Q129: How do you debug circular references in resolvers?**
Use depth limiting or trace logging. Circular references usually result in "Maximum call stack exceeded" or infinite JSON structures.

**Q130: How do you catch field-level errors in resolvers?**
Wrap resolver logic in `try/catch`. If you catch it and return `null` (for nullable fields), the error is suppressed or added to the `errors` list partial response.

**Q131: How do you report validation errors in input objects?**
Throw a specific error type (e.g., `UserInputError`) containing a list of invalid fields. Apollo Server provides built-in `UserInputError`.

**Q132: How do you track slow queries in production?**
Enable tracing (Apollo Tracing) or use performance monitoring tools (Apollo Studio) to identify resolvers taking the most time.

**Q133: How do you debug failed subscriptions?**
Check WebSocket connection status, heartbeat (keep-alive) settings, and ensure the PubSub system (Redis) is reachable.

**Q134: How can error boundaries be implemented in a GraphQL frontend?**
In React, use Error Boundaries to catch crashes. Apollo's `useQuery` also provides an `error` object to render fallback UI for that specific part.

**Q135: How do you test error scenarios in GraphQL?**
Mock the resolver/server to throw specific errors and verify that the client handles them correctly (e.g., showing a toast notification).

**Q136: How can you log queries with sensitive data excluded?**
Sanitize variables (mask passwords/tokens) before logging. Also, consider logging only the query hash for persisted queries.

**Q137: How do you validate incoming queries for correctness?**
The GraphQL engine does this automatically against the Reference implementation (validate syntax, type matching).

**Q138: How do you detect unused types or fields in a schema?**
Using usage analytics (Apollo Studio) or tools like `graphql-inspector coverage` based on traffic logs or client codebases.

**Q139: How do you simulate backend errors during frontend testing?**
Using `MockedProvider` (Apollo) and passing an `error` result for a specific mock query.

**Q140: What is Apollo Server’s `formatError` function used for?**
To transform errors before sending them to the client. Used for masking sensitive info, formatting validation errors, or adding request IDs.

---

## 🔹 8. Tooling & Ecosystem (141–160)

**Q141: What is GraphiQL and how is it different from Playground?**
GraphiQL is the original lightweight IDE for GraphQL. Playground (by Prisma, later Apollo) added features like tabs, headers, and traces. (Apollo now uses "Explorer").

**Q142: What is GraphQL Voyager?**
A tool that visualizes your GraphQL schema as an interactive graph of nodes and connections. Great for exploring relationships.

**Q143: What is GraphQL Inspector used for?**
A tool for validating changes (diff), checking coverage, and finding similar types in your schema. Useful in CI/CD.

**Q144: What is GraphQL Code Generator?**
A CLI tool that generates code from your schema and operations (TypeScript types, React hooks, Resolvers) to ensure type safety.

**Q145: How do you generate TypeScript types from a schema?**
Using `graphql-codegen` with `typescript` and `typescript-resolvers` plugins. It reads schema.graphql and outputs types.ts.

**Q146: What is `graphql-tag` used for?**
A library to parse GraphQL query strings (using `gql` template literal) into AST (Abstract Syntax Tree) for the client.

**Q147: How do you use Postman with GraphQL?**
Postman has native GraphQL support. You can import schemas, write queries with auto-complete, and send variables.

**Q148: What is `graphql-tools` and how is it useful?**
A library for building, mocking, and stitching GraphQL schemas. It allows separating schema logic and merging it back.

**Q149: What is GraphQL Mesh used for?**
To query diverse sources (REST, Swagger, Databases) using GraphQL query language. It acts as a gateway that translates GraphQL to other protocols.

**Q150: What are some popular GraphQL servers in Node.js?**
Apollo Server, GraphQL Yoga, Mercurius (Fastify), Express-GraphQL.

**Q151: How do you use `makeExecutableSchema`?**
A function from `graphql-tools` that takes `typeDefs` and `resolvers` and combines them into a runnable GraphQL schema object.

**Q152: What is `mergeTypeDefs` and `mergeResolvers`?**
Utilities to combine multiple schema strings and resolver objects into one. Essential for modularizing large servers.

**Q153: How do you introspect a remote schema?**
By running an introspection query against the remote endpoint. Tools like `graphql-codegen` can do this to generate local types.

**Q154: What is graphql-yoga and how is it different from Apollo Server?**
Yoga is built on standard web APIs (Fetch), supports multiple runtimes (Node, Deno, Cloudflare Workers), and is generally lighter-weight than Apollo.

**Q155: What are some popular GraphQL IDE plugins?**
Apollo GraphQL (VS Code), GraphQL for VSCode. They provide syntax highlighting, auto-complete, and jump-to-definition.

**Q156: What is Hasura and how does it relate to GraphQL?**
An instant GraphQL engine on top of your database (Postgres). It auto-generates a CRUD GraphQL API with permissions.

**Q157: What is StepZen?**
A GraphQL-as-a-Service platform that helps you build a GraphQL API by configuring backends (REST, DBs) via configuration.

**Q158: What is GraphQL Modules?**
A framework for modularizing large schemas with dependency injection and scoped context management.

**Q159: What is graphql-request and how is it used?**
A minimal GraphQL client for Node and browsers. Simple wrapper around `fetch` for sending queries.

**Q160: How do you compare Prisma with Hasura?**
Prisma is an ORM (code-first, DB access layer). Hasura is an API engine (instant server). You assume write code with Prisma; you configure Hasura.

---

## 🔹 9. Real-world Implementation (161–180)

**Q161: How do you migrate a REST API to GraphQL?**
Wrap REST endpoints in GraphQL resolvers. Gradually deprecate REST endpoints as clients switch to the Graph.

**Q162: How do you combine REST and GraphQL in a hybrid architecture?**
Use GraphQL for new features/complex reads, keep REST for simple CRUD/legacy or file uploads. Or use GraphQL Gateway to wrap REST.

**Q163: How do you handle file uploads with Apollo Server?**
Using `graphql-upload` (multipart requests). However, best practice is generating Signed URLs (S3) via GraphQL and uploading directly from client to S3 (bypassing the Graph server).

**Q164: How do you implement pagination using cursor-based approach?**
Return `edges` (with `cursor` and `node`) and `pageInfo` (`endCursor`, `hasNextPage`). Client sends `after: "endCursor"` to get next page.

**Q165: What are some common pagination strategies in GraphQL?**
Offset-based (`skip`/`limit` - easy but performance issues) and Cursor-based (Relay style - scalable, stable).

**Q166: How do you control access at the field level in GraphQL?**
Apply directives (`@auth`) or checks in the resolver. If check fails, return `null` or throw unauthorized error for that field.

**Q167: How do you support multi-tenancy in GraphQL?**
Use `context` to inject `tenantId` (from headers). Resolvers/DB calls naturally filter data by this `tenantId`.

**Q168: How do you maintain backwards compatibility in your schema?**
Never remove fields. Mark them `@deprecated`. Add new optional arguments. Add new types for breaking changes.

**Q169: How do you implement rate-limiting per user?**
Use Redis to track request complexity/count per user ID or IP token. middleware checks this count before executing.

**Q170: How do you integrate GraphQL with microservices?**
Use Federation (Gateway). Each microservice serves a subgraph. Gateway composes them. Resolvers in service A can reference entities in service B.

**Q171: What is GraphQL over gRPC?**
Using gRPC for internal service-to-service communication (performance) but exposing GraphQL at the edge (flexibility) for clients.

**Q172: How do you set up a GraphQL API Gateway?**
Deploy an instance of Apollo Gateway or a similar router. Configure it with the service list (subgraphs). It handles query planning.

**Q173: What is the role of Redis in GraphQL APIs?**
Caching (APQ, Response Cache), PubSub for Subscriptions, and Rate Limiting counters.

**Q174: How do you track which clients use which fields in a schema?**
Using schema usage reporting (Apollo Studio). It analyzes incoming query shapes and maps them to client versions.

**Q175: How do you handle large payloads efficiently?**
Use pagination. Limit nesting depth. Implement `@defer` to stream data. Gzip/Brotli compression at HTTP layer.

**Q176: How do you version your GraphQL schema?**
You don't (typically). You evolve it. If you *must*, use global versioning in URL (`/v1/graphql`), but field-level evolution is preferred.

**Q177: What is a GraphQL schema registry?**
A central repository (like Apollo Studio) that stores the history of your schema, validates changes, and serves the current schema to the gateway.

**Q178: What are breaking changes in GraphQL and how do you prevent them?**
Removing types/fields or making optional args required. Prevent via CI checks (`graphql-inspector diff`) against the production schema.

**Q179: What’s the best way to handle long-running mutations?**
Return a "Job" or "Operation" ID immediately. Client subscribes to updates on that Job ID or polls for status.

**Q180: How do you handle transactions in GraphQL mutations?**
Handle it in the resolver/service layer using DB transactions. Ensure all DB writes succeeds or fail together before returning.

---

## 🔹 10. Security & Authorization (181–200)

**Q181: How do you protect against GraphQL injection?**
GraphQL is strongly typed and parameterized (variables), which naturally prevents SQL injection if using ORMs/Prepared statements. Validate input types stringently.

**Q182: What is depth analysis and why is it important?**
Analyzing the query tree depth. Important to stop malicious clients from crashing the server with `a { b { a { ... } } }`.

**Q183: How do you prevent overly nested queries?**
Use `graphql-depth-limit` middleware to reject queries deeper than N levels.

**Q184: How do you use GraphQL Armor for security?**
A middleware stack that automatically adds protections: cost limit, depth limit, block field suggestions, and block introspection.

**Q185: How do you prevent information leakage in GraphQL?**
Disable introspection in prod. Use custom error formatting to hide stack traces. disable field suggestions ("Did you mean X?").

**Q186: What is a security audit checklist for GraphQL APIs?**
Check AuthN/AuthZ, Introspection disabled, Rate limits enabled, timeout configured, cost analysis enabled, sensitive data masked in logs.

**Q187: What is the purpose of query whitelisting?**
Also called "Persisted Queries". Only allows specific, pre-approved queries to be executed. Everything else is rejected.

**Q188: How do you enforce scopes or roles at the schema level?**
Using directives: `type Mutation { deleteUser: User @auth(role: ADMIN) }`.

**Q189: How do you block introspection queries in production?**
In Apollo Server: `introspection: false`. Or block queries containing `__schema` or `__type` at the gateway/WAF level.

**Q190: How can you control query cost by complexity?**
Use `graphql-query-complexity`. Define costs for expensive fields (resolving lists, DB hits). Reject high-cost queries.

**Q191: How do you authenticate users in subscriptions?**
Pass the token in the `connectionParams` of the WebSocket handshake. Validate it in the `onConnect` callback.

**Q192: What are the best practices for securing GraphQL endpoints?**
Use HTTPS. Enforce Auth. Validate Inputs. Limit payload size. Monitor for anomalies.

**Q193: What is the risk of `__schema` access and how to mitigate it?**
Reveals entire business data model. Mitigate by disabling introspection publicly.

**Q194: How do you enforce field-level authorization using middleware?**
Use `graphql-shield` to define a permission tree that mirrors your schema. It runs before resolvers.

**Q195: What are the security implications of third-party schema stitching?**
Trust. If a remote schema is compromised or changes, it can break your gateway or inject malicious data. Validate remote schemas.

**Q196: How do you handle authorization for deeply nested fields?**
Pass auth context down. Each resolver checks permissions. OR fetch all allowed data at the root (optimization) if feasible.

**Q197: How do you isolate public vs internal schemas?**
Create two schemas. Or use Federation with a "Public" subgraph and "Private" subgraph. Or use libraries to filter schema fields based on user role.

**Q198: How do you audit client usage in GraphQL?**
Log `operationName`, `clientName`, `clientVersion` (headers). Apollo Studio does this automatically.

**Q199: What is the risk of auto-generated queries?**
They can be inefficient or massive. Whitelisting prevents arbitrary auto-generated queries from executing.

**Q200: How do you implement RBAC (role-based access control) in GraphQL?**
Store roles in JWT. In context creation, decode JWT. In resolvers/middleware, check `user.roles.includes(REQUIRED_ROLE)`.

---

## 🔹 11. Server Architecture & Infrastructure (201–220)

**Q201: How do you serve a GraphQL API over WebSocket?**
Using a library like `graphql-ws` or `subscriptions-transport-ws` (legacy). It establishes a persistent connection for real-time bidirectional communication.

**Q202: What is a GraphQL Gateway?**
A server that sits in front of your downstream services (subgraphs/microservices) and provides a single unified GraphQL API to clients.

**Q203: How do you handle file uploads in a serverless GraphQL environment?**
Accepting binary data in Lambdas is expensive. It's better to use a "Pre-signed URL" pattern: Mutation returns a secure upload URL (S3), client uploads directly to S3.

**Q204: What is the difference between a Graph and a DAG in GraphQL context?**
GraphQL schemas form a Graph (potentially cyclic). Queries must be trees (acyclic) to terminate.

**Q205: How do you optimize cold starts for GraphQL on AWS Lambda?**
Use lighter libraries (`graphql-yoga` instead of Apollo), bundle code with Webpack/Esbuild, and use Provisioned Concurrency.

**Q206: How do you protect GraphQL APIs from DDoS attacks?**
Rate limiting (IP-based), Query Cost Analysis, Timeouts, and WAF (Web Application Firewall) to filter malicious traffic.

**Q207: What is query flattening and when is it used?**
Rewriting a nested query into a flatter structure or SQL join to avoid N+1 queries. Tools like `join-monster` do this.

**Q208: How do you handle multi-region data replication in GraphQL?**
The GraphQL layer doesn't handle this. The underlying database (e.g., DynamoDB Global Tables, CockroachDB) handles replication. GraphQL just queries the local read replica.

**Q209: How do you implement a BFF (Backend For Frontend) with GraphQL?**
Create a specific GraphQL API for a specific client (e.g., Mobile App BFF) that wraps general-purpose microservices, tailored to that client's needs.

**Q210: What is the role of a schema registry in a CI/CD pipeline?**
It validates schemas before deployment, checks for breaking changes, and prevents invalid schemas from reaching production.

**Q211: How do you monitor individual field usage?**
Apollo Studio or similar tools track which fields are requested. This helps in safely deprecating unused fields.

**Q212: How do you handle timeouts for specific resolvers?**
Wrap the resolver logic in a `Promise.race` with a timer, or use server configuration to set execution time limits.

**Q213: What is the N+1 problem in Distributed Systems?**
Similar to DBs but worse: fetching a list of items and then making an HTTP call for each item to another microservice. Solved with DataLoaders or batch endpoints.

**Q214: How do you version a Federated Graph?**
You don't version the whole graph. You evolve individual subgraphs compatibly. The Gateway composes them into a "current" api.

**Q215: How do you implement caching in a GraphQL Gateway?**
Whole-query caching (APQ), or partial query caching (caching subgraph responses).

**Q216: What is "Schema Stewardship"?**
The practice of owning and maintaining the integrity of the graph, often by a platform team that reviews schema changes from other teams.

**Q217: How would you design a GraphQL API for a chat application?**
Heavy use of Subscriptions for new messages. Mutations for sending. Queries for history (with cursor-based pagination).

**Q218: How do you handle data consistency in distributed mutations?**
Difficult. Use Saga pattern or Two-Phase Commit (2PC) in the backend services. GraphQL just reports the final success/failure.

**Q219: How do you secure a public-facing GraphQL API?**
Strict rate limiting, complexity limits, disable introspection, require API keys or OAuth tokens.

**Q220: What are "supergraphs"?**
A term coined by Apollo for a graph composed of many subgraphs, creating a network of data across an organization.

---

## 🔹 12. Scalability, Caching & Optimization (221–240)

**Q221: How do you solve the N+1 problem without DataLoader?**
By using "Look-ahead" (inspecting `info` to see selected fields) and performing SQL JOINs to fetch everything in one query.

**Q222: What is the difference between server-side and client-side caching?**
Server-side (Redis/CDN) reduces load on the API. Client-side (Apollo Cache) reduces network requests and makes the UI snappy.

**Q223: How do you use HTTP caching with GraphQL?**
Use GET requests for queries. Send `Cache-Control` headers. CDNs interpret these headers to cache the JSON response.

**Q224: What is a persisted query whitelist?**
A security/performance list where the server only accepts query hashes that it knows. Rejects arbitrary new queries.

**Q225: How do you optimize resolvers that call slow 3rd party APIs?**
Cache the responses (Redis) with a TTL. Use timeouts. Return partial data if possible (making the slow field nullable).

**Q226: What is the `@defer` directive?**
Allows the server to return the initial part of the response immediately and stream the slow parts (fields marked `@defer`) later. Great for frontend performance.

**Q227: What is the `@stream` directive?**
Similar to `@defer` but for Lists. Return the first few items immediately and stream the rest as they become available.

**Q228: How do you implement Edge Caching for GraphQL?**
Using a CDN (Cloudflare/Fastly) that supports GraphQL (or passing query params in GET). Some CDNs parse GraphQL bodies to cache intelligently (`GraphCDN`).

**Q229: What is "Over-fetching" in the context of DB performance?**
Selecting `SELECT * FROM table` when the GQL query only asked for `name`. Optimize by selecting only requested columns.

**Q230: How do you handle massive schema sizes (10k+ types)?**
Split into subgraphs (Federation). Lazy load schema parts if using custom tooling. Ensure IDEs (IntelliJ/VSCode) can handle the large introspection result.

**Q231: How do you load-balance GraphQL subscriptions?**
Since WebSockets are stateful, you need a sticky session or a dedicated Subscription fleet that reads from a shared PubSub (Redis) to broadcast to the right instance.

**Q232: What is memoization in resolvers?**
Caching the function result based on arguments for the duration of the request (or longer). Dataloader essentially does request-level memoization.

**Q233: How do you profile a slow GraphQL request?**
Use Apollo Tracing or OpenTelemetry. Break down execution by resolver duration to find the bottleneck.

**Q234: How do you minimize payload size?**
Use Gzip/Brotli. Use short field aliases (extreme optimization). Use Persisted Queries (doesn't send query string).

**Q235: How do you handle "Thundering Herd" with GraphQL?**
If a cache key expires and 1000 requests hit it, use "Request Coalescing" (single flight) so only one hits the DB, others wait for that result.

**Q236: What is GraphQL JIT (Just-In-Time) compilation?**
An optimization (like `graphql-jit`) that compiles the query execution plan into native code/optimized functions at runtime, bypassing standard interpreter overhead.

**Q237: How do you optimize for mobile networks?**
Use Offline support (Apollo Cache persistence). Retry logic. Minimize payload. Use `@defer` to show content faster.

**Q238: How do you handle large lists in GraphQL?**
Pagination (Cursor). Never return unbounded lists. Set hard limits (e.g., `first: 100` max).

**Q239: What is the cost of Schema Stitching vs Federation?**
Stitching requires manual conflict resolution and imperative glue code. Federation distributes the responsibility, generally scaling better for teams.

**Q240: How do you stress test a GraphQL server?**
Use tools like `k6` or `JMeter`. Generate queries with varying complexity and depth to simulate real load.

---

## 🔹 13. Data Modeling & Schema Design (241–260)

**Q241: What is a "node" interface?**
A common pattern (Relay) where every object has a globally unique `id`. `interface Node { id: ID! }`. Allows refetching any object by ID.

**Q242: How do you design specific mutations vs generic CRUD?**
Prefer specific "Task-based" mutations (`registerUser`, `changePassword`) over generic (`updateUser`). It captures intent and business rules better.

**Q243: How do you handle polymorphic data?**
Use `Union` or `Interface`. E.g., `SearchResult = User | Post | Comment`.

**Q244: What is the best practice for naming lists?**
Pluralize (e.g., `users`). Or be explicit `userList`, `usersConnection` (for Relay).

**Q245: How do you model Many-to-Many relationships?**
Through connection types or list fields on both sides. `User.groups` and `Group.users`.

**Q246: When should you use a Custom Scalar?**
When you have a specific data format with validation rules (Date, Email, RGBColor, GeoPoint).

**Q247: How do you handle "Input Unions" (Polymorphic Input)?**
GraphQL doesn't support Input Unions natively (yet). Workaround: `input UserInput { email: String, phone: String }` (oneOf) or generic JSON scalar.

**Q248: How do you model "permissions" in the schema?**
You might expose a `viewer { permissions }` field. Or fields might just be null if unauthorized (implicit).

**Q249: How do you handle Date/Time?**
Use a custom scalar `DateTime` (ISO-8601 string). Standard `String` is too loose.

**Q250: What is the "Payload" pattern in mutations?**
Return a specific object for each mutation: `type UpdateUserPayload { user: User, errors: [Error] }`. Allows returning metadata/status.

**Q251: How do you deprecate an enum value?**
SDL supports `@deprecated` on enum values. `enum Role { ADMIN, @deprecated(reason: "Use SUPERADMIN") ROOT }`.

**Q252: How do you design for nullability?**
Default to nullable (resilience). Make non-null (`!`) only if you can guarantee it exists (IDs). If a non-null field fails, the whole parent becomes null.

**Q253: How do you share types between Input and Output?**
You can't. Input types and Object types are distinct. You must duplicate `UserInput` and `User` definitions often.

**Q254: How do you group related fields?**
Put them in a sub-object. `User { address { street, city } }` instead of `userCity`, `userStreet`.

**Q255: How do you handle arguments on fields?**
Use them for filtering, sorting, or formatting. `users(status: ACTIVE)`, `avatar(size: LARGE)`.

**Q256: What is the Global Object Identification specification?**
The Relay spec requirement that every object has a unique `id` and can be fetched via `node(id: ID!)`.

**Q257: How do you handle schema documentation?**
Write descriptions in the SDL strings (`""" desc """`). These appear in API docs automatically.

**Q258: How do you design for filterable lists?**
Use a specific input object: `users(filter: UserFilterInput)`. `input UserFilterInput { nameContains: String, ageGt: Int }`.

**Q259: How do you model errors in the schema (Result Types)?**
Use Unions for results: `union CreateUserResult = User | UserError`. explicitly models errors as data.

**Q260: How do you handle circular dependencies in schema files?**
In `typeDefs`, circular references are fine (SDL is declarative). In code (imports), use lazy loading or merge utilities.

---

## 🔹 14. Testing & QA (261–280)

**Q261: What is "Schema Diffing"?**
Comparing the previous schema with the new one to detect breaking changes before merging PRs.

**Q262: How do you mock GraphQL requests in frontend tests?**
Use `MockedProvider` (Apollo) or `msw` (Mock Service Worker) to intercept network requests and return mock JSON.

**Q263: What implies a "Breaking Change" in GraphQL?**
Removing a field/type, renaming a field, changing a type (Int -> String), or making an optional argument required.

**Q264: How do you verify backward compatibility?**
Run `graphql-inspector` against the production schema. Run regression tests.

**Q265: How do you test permissions?**
Write test cases for different user roles (Admin, User, Anon). Verify that unauthorized fields return null or throw errors.

**Q266: What is a specialized GraphQL client for testing?**
`supertest` (for HTTP), or simply calling the `execute()` function directly in integration tests to bypass HTTP layer.

**Q267: How do you generate mock data from schema?**
`graphql-tools` has `addMocksToSchema`. It auto-generates fake data (Lorem Ipsum) based on types.

**Q268: How do you test subscriptions?**
Connect a WebSocket client in tests, trigger a mutation, and assert that the correct message was received.

**Q269: How do you validate query depth in tests?**
Send a deeply nested query and assert the server returns an error (if depth limit is configured).

**Q270: How do you test custom scalars?**
Test the `parseValue`, `serialize`, and `parseLiteral` functions with valid and invalid inputs.

**Q271: How do you test N+1 problem detection?**
Count SQL queries generated by a test request. If it scales linearly with result size, fail the test.

**Q272: What is property-based testing in GraphQL?**
Generating random valid queries (based on schema) to crash the server or find edge cases.

**Q273: How do you snapshot test a schema?**
Save the SDL (`printSchema`) to a file. In tests, generate current SDL and compare with the file.

**Q274: How do you test Federation?**
Spin up all subgraphs locally + Gateway. Run queries against the Gateway to verify composition works.

**Q275: How do you test error formatting?**
Trigger an error and assert the response JSON structure matches the expected format (e.g., masking stack traces).

**Q276: How do you test directive logic?**
Create a test schema using the directive. Run a query. Assert the directive modified the behavior (e.g., upper-cased a string).

**Q277: How do you load test a subscription server?**
Open thousands of idle WS connections. Broadcast a message. Measure latency of receipt across clients.

**Q278: How do you test resolver timeouts?**
Mock a service to delay response. specific resolver timeout should return null/error.

**Q279: How do you ensure test isolation?**
Reset the DB or mock state between tests. Don't let side effects leak.

**Q280: How do you test for breaking changes in clients?**
Track client queries (Persisted Queries). Check if the schema change affects any known active query.

---

## 🔹 15. API Evolution & Governance (281–300)

**Q281: What is the "Dangerous" change warning?**
Some tools warn on optional changes that might be risky, like changing a default value or adding a nullable field to an input type.

**Q282: How do you manage schema ownership?**
CODEOWNERS file pointing to team definitions for specific schema files/modules.

**Q283: What is a Field Policy?**
A definition (often in Gateway or Client) of how to read/write a field, merge it, or cache it.

**Q284: How do you promote a schema from Dev to Prod?**
Schema Registry. CI pushes to "Staging" variant. Tests run. CD pushes to "Prod" variant.

**Q285: How do you handle "Zombie" fields (unused)?**
Identify with analytics. Mark `@deprecated`. Wait X months. Remove when usage checks show 0.

**Q286: What is the "One Graph" principle?**
The idea that an organization should have a single, unified graph instead of many scattered, disconnected API endpoints.

**Q287: How do you handle third-party vendor schemas?**
Wrap them. Do not expose them directly. Map their types to your internal domain model to decouple.

**Q288: How do you enforce naming conventions?**
Linting rules (`camelCase` for fields, `PascalCase` for types). `graphql-eslint`.

**Q289: How do you handle "God Objects" (e.g., massive User type)?**
Split them using Federation (extend type User in multiple services) so no single team owns the whole thing.

**Q290: How do you coordinate cross-team schema changes?**
Regular "Schema Guild" meetings, Design Reviews, and using Pull Requests on the schema repo.

**Q291: What is "Schema-First" development in teams?**
Teams agree on the SDL (contract) first. Then Frontend and Backend work in parallel using mocks.

**Q292: How do you document "Why" a field exists?**
Verbose descriptions in SDL. Link to JIRA tickets or Confluence pages in comments.

**Q293: How do you handle multi-environment schemas?**
Use environment variables or build configurations to toggle experimental fields (e.g., `@include(if: $isDev)`).

**Q294: What is the role of the "Graph Champion"?**
An internal advocate who promotes best practices, educates teams, and maintains the platform infrastructure.

**Q295: How do you align GraphQL types with DB models?**
You shouldn't align them 1:1. The Graph is a View. Use mapping layers (mappers) to translate DB entities to Graph types.

**Q296: How do you handle sensitive PII in schema?**
Annotate with `@pii` or `@sensitive`. Ensure logs mask these fields. Limit access via directives.

**Q297: How do you audit schema complexity over time?**
Track "Type Count", "Field Count", "Deprecation Count" metrics in your registry.

**Q298: How do you handle multiple "User" types (External, Internal, Admin)?**
Either separate types (`Customer`, `AdminUser`) or Interfaces (`interface User`).

**Q299: What is the benefit of a "Linting" step for schema?**
Consistency. Enforces description presence, naming rules, type safety best practices automatically.

**Q300: How do you ensure the graph remains "Product-Centric"?**
Design based on UI/Experience requirements, not Database tables. "Demand-Driven" schema design.

---

## 🔹 21. Multi-Client & Frontend Strategy (401–420)

**Q401: How do you design a GraphQL schema that supports both web and mobile apps?**
Use a shared core schema but leverage `@include/@skip` directives or client-specific fragments. Avoid "Backend for Frontend" (BFF) patterns per client unless absolutely necessary (keeps the graph unified).

**Q402: How do you version GraphQL queries across multiple clients?**
You version the *client's* queries (e.g., Persisted Queries v1, v2), not the schema. The schema evolves additively. Clients are upgraded to request new fields.

**Q403: What are named operations and how do they help in debugging across clients?**
Giving queries names like `query GetUserProfile { ... }` instead of anonymous `query { ... }`. Logs show `GetUserProfile`, making it easy to identify which client feature is failing.

**Q404: How do you reduce over-fetching in mobile GraphQL queries?**
Encourage mobile teams to write strict queries requesting *only* needed fields. Use `Persisted Queries` to ensure they don't accidentally send massive queries.

**Q405: How do you design for a React + iOS + Android + SPA setup with GraphQL?**
Treat the Schema as the "Contract". Use Codegen to generate native types (Swift/Kotlin/TS) from that schema. Ensure all clients use the same field naming conventions.

**Q406: How do you support multiple brands (white-label) via a single GraphQL API?**
Add a `brandId` header. Context filters data by brand. Or use Interfaces (`interface Product`) and implementations (`BrandAProduct`, `BrandBProduct`) if structures differ significantly.

**Q407: What are persisted queries and how do they benefit mobile clients?**
Mobile apps ship with query hashes. They send `id: "hash123"` instead of the full string. Saves bandwidth and latency, especially on flaky mobile networks.

**Q408: How do you analyze usage of GraphQL fields across multiple teams?**
Apollo Studio / GraphQL Hive. They break down usage by "Client Name" and "Client Version" (headers sent by Apollo Client/Relay).

**Q409: How do you handle breaking changes for multiple frontend teams?**
Deprecate the field. Monitor usage per client version. Reach out to teams using old versions. Only remove when usage drops to 0 (or acceptable risk).

**Q410: How do you allow experimentation for frontend teams using GraphQL?**
Expose "Beta" fields or types. Mark them with a comment `(Experimental)`. Or use a separate "Labs" endpoint.

**Q411: How do you build GraphQL clients for low-bandwidth environments?**
Aggressive caching (normalized cache). Use `@defer` to get critical data first. Persisted Queries to reduce request size.

**Q412: How do you manage fragment co-location across multiple clients?**
Each UI component defines its own fragment. The parent component composes them. This is the "Relay" way, now common in Apollo too.

**Q413: How do you allow dynamic UI configuration using GraphQL?**
Return "Server-Driven UI" structures: `type UIComponent { type: "CARD", title: String, actionUrl: String }`. The UI renders based on the type.

**Q414: How do you group operations by product team in schema analytics?**
Use a specific naming convention for operations: `TeamName_Feature_Query`. Analytics tools can regex match the operation name.

**Q415: How do you build feature flags using GraphQL query conditions?**
Pass flags as variables: `@include(if: $newFeatureEnabled)`. Or expose the feature flag state in `viewer { flags { newFeature } }`.

**Q416: How do you scale query delivery across multiple environments?**
CDNs. Edge Caching. The "Query" doesn't change, only the data/endpoint.

**Q417: How do you provide access-controlled documentation for client teams?**
Publish different documentation sites (using Spectaql) based on introspection of "Public" vs "Private" schema variants.

**Q418: How do you support legacy clients that rely on old schema patterns?**
Keep the old fields. Mark `@deprecated`. Map them to the new logic in resolvers. Never break them until the legacy client is sunset.

**Q419: What are GraphQL Operation Collections in Apollo?**
A registry of all allowed operations (whitelisting). You can browse them to see "What queries does the iOS app v3.5 run?".

**Q420: How do you detect unused operations across all consuming apps?**
Schema coverage tools. If an operation in the registry hasn't been executed in X days, it's unused.

---

## 🔹 22. Domain-Driven Design & Use Cases (421–440)

**Q421: How do you design a schema for a multi-tenant SaaS product?**
Inject `tenantId` into Context (via Header). All top-level queries implicitly filter by this ID. Types reflect tenant-agnostic concepts (`User`, `Invoice`).

**Q422: How do you model roles, permissions, and feature flags in GraphQL?**
`type User { permissions: [String] }` or `type User { canEditPost(postId: ID!): Boolean }`. Expose capability checks rather than raw roles.

**Q423: How do you expose product configuration via GraphQL?**
`type Config { theme: String, modules: [ModuleConfig] }`. Useful for initializing the app state.

**Q424: How do you model workflow states (like “Draft”, “Pending”, “Live”)?**
Enums: `enum PostStatus { DRAFT, PUBLISHED }`. Or Union types if the data shape changes based on state (`DraftPost`, `PublishedPost`).

**Q425: How do you structure schema for a digital marketplace?**
`type Marketplace { listings: [Listing] }`. `interface Listing`. `ProductListing`, `ServiceListing`. Cart and Order types.

**Q426: How do you model financial transactions with GraphQL?**
High precision. Use `Decimal` custom scalar (string), never Float. `type Transaction { amount: Decimal!, currency: Currency! }`.

**Q427: How do you design an audit log system with GraphQL queries?**
`type AuditLog { actor: User, action: String, timestamp: DateTime, target: Entity }`. Use Interfaces for `target`.

**Q428: How do you model e-commerce order states in GraphQL?**
`enum OrderStatus`. Mutations like `placeOrder`, `shipOrder`. Subscriptions for `orderUpdated`.

**Q429: How do you expose shipping or logistics data via GraphQL?**
Nested objects: `Order { shipment { trackingNumber, carrier, estimatedDelivery } }`.

**Q430: How do you handle schema evolution for compliance (e.g. GDPR)?**
Mark PII fields. Ability to return "Redacted" or null if the user rescinded consent.

**Q431: How do you expose A/B test variants via GraphQL?**
`viewer { experiments { name, variant } }`. The frontend uses this to decide which UI to show.

**Q432: How do you model timezones, locales, and formatting in schema?**
Arguments on fields: `createdAt(format: String, timezone: String)`. let the server handle formatting (moment/date-fns).

**Q433: How do you model document uploads and previews in GraphQL?**
`type Document { url: String, mimeType: String, thumbnail(size: Int): String }`. Get the upload URL via mutation first.

**Q434: How do you expose configuration-driven business rules via schema?**
`type ValidationRules { minPasswordLength: Int, requiredCampos: [String] }`.

**Q435: How do you represent user-to-user relationships in GraphQL?**
`User { friends: [User], followers: [User] }`. Connection pattern for pagination (`friendsConnection`).

**Q436: How do you structure content scheduling (like posts) in GraphQL?**
`Post { publishAt: DateTime, status: PostStatus }`. Query filters for `status: PUBLISHED` by default.

**Q437: How do you model event-driven workflows (e.g., triggers)?**
`type Trigger { id: ID, condition: String, action: Action }`.

**Q438: How do you expose email, SMS, or push templates via GraphQL?
`type NotificationTemplate { id: ID, subject: String, body: String (HTML), variables: [String] }`.

**Q439: How do you represent pricing, billing tiers, and discounts?**
`type Plan { price: Money, features: [String], discount(code: String): Money }`. Calculated fields are great here.

**Q440: How do you handle batch updates to related entities in GraphQL?**
`mutation updateManyTodos(input: [UpdateTodoInput!]!)`. Returns `[Todo]`. Transactional.

---

## 🔹 23. Security, Auth & Vulnerabilities (441–460)

**Q441: How do you implement row-level access control in GraphQL?**
In the Data Layer (ORM/SQL). Resolvers pass the `currentUser` to the Service/Repository, which adds `WHERE user_id = ?` to the query.

**Q442: How do you manage access tokens in GraphQL operations?**
Sent in Authorization Header (`Bearer ...`). Context parser validates it before execution.

**Q443: What are security implications of introspection in production?**
Leaks schema details, potential vulnerabilities, and unused fields. Attackers map the API. Disable it.

**Q444: How do you expose RBAC in GraphQL fields?**
Directives: `@auth(role: ADMIN)`. If user lacks role, field is null (and error added) or the whole query matches "Unauthorized".

**Q445: How do you ensure a mutation is only allowed by certain roles?**
Check `context.user.role` at the start of the resolver. Throw `ForbiddenError` if invalid.

**Q446: How do you prevent IDOR (Insecure Direct Object Reference) in GraphQL?**
Always check ownership: `if (document.ownerId !== user.id) throw Error`. Using UUIDs helps but isn't a replacement for permission checks.

**Q447: What are common GraphQL-specific attack vectors?**
DoS (Deep queries), Batching attacks (1000 alias fields), Introspection leak, Injection (if args passed to SQL raw).

**Q448: How do you prevent enumeration attacks in GraphQL?**
Don't return "User not found" distinct from "Password incorrect". Rate limit errors. Use UUIDs.

**Q449: How do you safely expose file URLs in a GraphQL schema?**
Return signed, short-lived URLs (S3 presigned). Don't return permanent public URLs if content is private.

**Q450: How do you integrate JWT authentication with resolvers?**
Middleware decodes JWT -> puts `user` in Context -> Resolver uses `context.user`.

**Q451: How do you enable field-level auth logic in federated schemas?**
Pass the `user` header/context to the subgraph. The subgraph resolver performs the check.

**Q452: How do you implement rate limiting per user in GraphQL?**
Redis counter based on User ID. Check before execution.

**Q453: How do you integrate OAuth2 or OpenID Connect with GraphQL?**
The Auth happens *before* GraphQL (Redirects/Tokens). GraphQL just validates the resulting Access Token.

**Q454: How do you block malicious queries based on behavior?**
Heuristics. "Too many errors". "Too deep". "Scanning for existing IDs". WAFs can detect these patterns.

**Q455: How do you validate request headers for GraphQL auth?**
Framework (Express/Fastify) middleware checks `req.headers`. If missing/invalid, return 401 before hitting GraphQL execution.

**Q456: How do you attach session context securely across federated services?**
Gateway verifies token. Passes a trusted "User-ID-Header" (signed or internal-only) to subgraphs. Subgraphs trust the Gateway.

**Q457: How do you avoid leaking sensitive data in GraphQL errors?**
Use `formatError`. Strip `exception.stacktrace` in production. Replace with "Internal Server Error" and a Reference ID.

**Q458: How do you implement attribute-based access control (ABAC) in resolvers?**
Check attributes: `if (resource.department == user.department && user.clearance >= resource.level)`. Detailed logic in domain layer.

**Q459: What tools can help scan GraphQL APIs for security risks?**
GraphQL Map, InQL (Burp Suite extension), Clairvoyance.

**Q460: How do you verify schema compliance with security policies?**
Linters (`graphql-inspector`). Custom rules: "All mutations must have @auth". "No fields named 'password'".

---

## 🔹 24. AI & Data Science Integrations (461–480)

**Q461: How can GraphQL be used to expose ML model predictions?**
`type Prediction { confidence: Float, label: String }`. Resolver calls the ML Model Service (e.g., Python/TensorFlow serve).

**Q462: How do you integrate GraphQL with vector databases like Pinecone?**
Resolver takes search text -> generates embedding (OpenAI API) -> queries Pinecone -> maps results to GraphQL types.

**Q463: How do you handle large language model APIs via GraphQL?**
Stream responses using `@defer` or Subscriptions. LLMs are slow; standard Query/Response might timeout.

**Q464: How can GraphQL support real-time analytics dashboards?**
Subscriptions on aggregated data streams. `subscription { revenueUpdated { total, daily } }`.

**Q465: How do you serve model inference pipelines via GraphQL?**
Start inference (Mutation -> Job ID). Poll or Subscribe for results. (Async pattern).

**Q466: How do you expose time-series anomaly detection in GraphQL APIs?**
`type Anomaly { timestamp: DateTime, score: Float }`. Query a time-series DB (Influx/Prometheus).

**Q467: How do you cache expensive ML predictions in resolvers?**
Redis. Key = Hash(Input features). Model inference is deterministic (mostly), so highly cacheable.

**Q468: How do you use GraphQL to fetch contextual embeddings for AI tasks?**
`query GetEmbeddings(text: String!)`. Returns vector array.

**Q469: How do you structure input types for prompt engineering via GraphQL?
`input PromptConfig { temperature: Float, templates: [String], vars: JSON }`.

**Q470: How do you paginate through ML-generated recommendations?**
Standard Cursor pagination. The "cursor" might be the index in the recommendation list.

**Q471: How do you expose tuning parameters via a GraphQL interface?**
Mutations to update model config. `updateHyperparams(...)`.

**Q472: What’s the best way to log model feedback via GraphQL mutations?**
`mutation ratePrediction(predictionId: ID!, rating: Int!)`. Feeds back into RLHF (Reinforcement Learning).

**Q473: How do you manage experiment tracking data with GraphQL?**
`type Experiment { runs: [Run] }`. `type Run { metrics: JSON, params: JSON }`. Great for ML Ops dashboards.

**Q474: How do you expose real-time metrics with GraphQL Subscriptions?**
Connect to Kafka/RabbitMQ. Stream metric updates to client.

**Q475: How can GraphQL help in AI-powered document search systems?**
One API to search distinct sources (Wiki, Jira, Slack). Search Resolver federates request to multiple search engines/vector DBs.

**Q476: How do you integrate OpenAI APIs into GraphQL resolvers?**
Call `openai.createCompletion` inside the resolver. Map the JSON response to your Schema types.

**Q477: How do you throttle usage of AI-based features via GraphQL?**
Token bucket limiting per user. Specific complexity cost for AI fields (e.g., `generateText` costs 100 points).

**Q478: How do you return explanations (XAI) via GraphQL queries?**
`type Prediction { result: String, explanation: String, featureImportance: [FeatureWeight] }`.

**Q479: How do you model semantic search layers in GraphQL schema?**
`type SearchResult { ... on Article { ... } ... on Video { ... } }`. Resolver handles the vector search logic.

**Q480: How can you use GraphQL to build an AI assistant interface?**
Mutations to send messages. Subscriptions to receive streaming character-by-character responses.

---

## 🔹 25. Edge Cases & Anti-Patterns (481–500)

**Q481: What are the dangers of exposing internal IDs in schema?**
Clients rely on them. Hard to migrate DBs later. Can leak count/velocity (auto-increment). Use Opaque IDs (base64 encoded) or UUIDs.

**Q482: How do you handle orphaned fields in evolving schemas?**
Visualizers. Fields connected to nothing. Remove them. They confuse devs.

**Q483: Why should you avoid dynamic field names in GraphQL?**
Breaks type safety. Clients can't request them explicitly. Use `JSON` scalar (carefully) or a List of Key-Value pairs.

**Q484: What are the risks of resolver-level side effects?**
Queries should be idempotent (Read-only). If a Query resolver writes to DB, caching breaks, and retries cause duplicate writes. Only Mutations should have side effects.

**Q485: What’s wrong with returning booleans for mutation success?**
Can't return the changed object (for cache update). Can't return detailed errors. Return a Payload object instead.

**Q486: What are problems with auto-generated schemas from databases?**
Leaks DB structure. Naming is usually DB-centric (snake_case). Hard to add computed fields or business logic later.

**Q487: Why is it bad to have a single “root” mutation field?**
`mutation { generic(action: "update") }`. Loses type safety. Hard to discover capabilities.

**Q488: What are the issues with overusing interfaces and unions?**
Client complexity. They have to write Fragments for every implementation. Large payloads if implementations diverge wildly.

**Q489: What problems arise from using scalar `JSON` types everywhere?**
Defeats the purpose of GraphQL (Type System). Client doesn't know what's inside. No validation.

**Q490: Why is schema nesting depth a scalability concern?**
`friends { friends { friends ... } }`. exponential complexity. DoS vector.

**Q491: What’s wrong with exposing file uploads as simple strings?**
Base64 strings bloat request size by 33%. Block the main thread during parsing. Stream it or use pre-signed URLs.

**Q492: What happens when too many clients share the same fragment?**
They over-fetch. If one client needs a new field, the fragment gets updated, and ALL clients now fetch that field, even if they don't need it.

**Q493: What causes memory leaks in GraphQL subscriptions?**
Not cleaning up listeners when client disconnects or when the server reloads. Infinite listener growth.

**Q494: What are the issues with exposing `__typename` to clients?**
Generally fine, but if class names are internal/sensitive, it leaks implementation details. (Usually considered a feature, not a bug, for caching).

**Q495: Why can using custom scalars lead to inconsistencies?**
If one team defines `Date` as Timestamp and another as ISO-String. Must standardize globally.

**Q496: What happens when you re-use input types across unrelated mutations?**
Coupling. Change the input for Mutation A, breaks Mutation B. Create dedicated inputs (`UpdateUserAddressInput`, `CreateUserInput`).

**Q497: How can poor nullability design lead to frontend bugs?**
Making everything non-null (`!`). One DB failure crashes the whole UI. Making everything nullable: Frontend code is full of `if (x)`. Balance is key.

**Q498: What’s the danger of relying only on introspection for docs?**
Missing "Why" context. Descriptions are often empty. Use a real Developer Portal with guides.

**Q499: Why should you avoid overly generic field names like `data` or `info`?**
Collisions. Ambiguity. `user.data` vs `product.data`. Be specific: `user.profile`, `product.details`.

**Q500: What is schema fatigue and how do you prevent it?**
Too many types, too many similar fields. Hard to find things. Prune deprecated fields. Organize into Modules/Domains.

---


## 🔹 16. Advanced Internals & Execution (301–320)

**Q301: How does GraphQL parse and execute a query under the hood?**
Lexing (breaking string into tokens) -> Parsing (creating AST) -> Validation (checking against schema) -> Execution (traversing AST and calling resolvers).

**Q302: What is the GraphQL type system made of internally?**
It's made of definitions (Scalar, Object, Interface, Union, Enum, InputObject) which are validated against the spec to form a Schema.

**Q303: How does GraphQL resolve a deeply nested query step-by-step?**
Breadth-first resolution (conceptually), but often executed via Depth-First Search (DFS) traversal where parent resolvers return data that triggers child resolvers.

**Q304: What is the role of the execution context in a GraphQL operation?**
It holds state shared across all resolvers for a single request, such as authentication results, database connections, and DataLoaders.

**Q305: How is a GraphQL schema represented in memory?**
As a graph of JavaScript objects (if using graphql-js) representing types and fields, enabling fast lookup during validation and execution.

**Q306: What are abstract syntax trees (AST) in GraphQL?**
A tree representation of the query string. Resolvers and tools work with the AST, not the raw string, to understand the query structure.

**Q307: How does `graphql-js` implement validation rules?**
It traverses the AST and checks it against a set of rules (e.g., "FieldsOnCorrectType", "KnownTypeNames") defined in the spec.

**Q308: How does GraphQL prevent infinite recursion in queries?**
It doesn't automatically preventing it (schema can be cyclic). You must implement **Maximum Depth** or **Complexity Analysis** to reject cyclic queries at runtime.

**Q309: What are execution phases in GraphQL?**
Parse -> Validate -> Execute. The Execute phase involves running resolvers and Coercing results.

**Q310: How does GraphQL treat circular fragments internally?**
Cyclic fragments are invalid. The validation phase detects cycles (e.g., Fragment A includes Fragment B, which includes A) and throws an error.

**Q311: How are variables resolved and substituted during execution?**
Variables are passed alongside the query. The executor replaces variable references in the AST with the actual values before running resolvers.

**Q312: What is the difference between execution phase and resolver logic?**
Execution phase is the engine's job (traversing the AST, handling nulls, merging errors). Resolver logic is your code (fetching the actual data).

**Q313: How is a GraphQL request batched internally by Apollo Server?**
Apollo Server doesn't auto-batch execution of separate requests by default unless you use Query Batching (array of queries). Internally, they are processed sequentially or in parallel depending on config.

**Q314: What is `graphql-parse-resolve-info` and how is it useful?**
A library to parse the complex `info` argument in resolvers into a cleaner object structure, making it easier to see which sub-fields are requested (Look-ahead).

**Q315: How can you hook into the execution flow of GraphQL?**
Using Plugins (Apollo) or Middleware (`graphql-middleware`). You can intercept Parse, Validate, or Execute phases.

**Q316: What is the difference between schema-first and resolver-first approaches?**
Schema-First: Write SDL, then write resolvers to match. Resolver-First (Code-First): Write code (classes/functions), generate SDL from it.

**Q317: What are the disadvantages of schema-first design?**
Resolver implementation can drift from SDL. Refactoring requires changing two places. String-based types are less type-safe in the backend code without codegen.

**Q318: How do directives impact the execution engine?**
They can alter execution at runtime. `@include`/`@skip` control field execution. Custom directives can wrap resolvers to inject logic.

**Q319: What is a GraphQL execution strategy?**
Logic defining how to map the schema and data to a result. Default is "Serial" for Mutations and "Parallel" (conceptually) for Queries.

**Q320: What are synchronous vs asynchronous resolvers in detail?**
Sync resolvers return a value immediately. Async resolvers return a Promise. The engine waits for Promises to resolve before continuing down that branch.

---

## 🔹 17. Third-Party Integrations (321–340)

**Q321: How do you integrate GraphQL with PostgreSQL?**
Use an ORM (Prisma, TypeORM, Sequelize) inside resolvers to query Postgres. Or use introspection tools like `PostGraphile` or `Hasura` to auto-generate the API.

**Q322: How does Hasura auto-generate GraphQL from Postgres?**
It inspects the Postgres catalog (tables, foreign keys) and builds a GraphQL schema that maps directly to SQL queries, handling all CRUD operations.

**Q323: How do you integrate MongoDB with GraphQL?**
Use Mongoose or the MongoDB Node driver in resolvers. Since Mongo documents are JSON-like, they map easily to GraphQL object types.

**Q324: What is Prisma’s role in GraphQL backends?**
Prisma acts as a strongly-typed database client (ORM replacement). It's popular in "Nexus" or "Code-First" setups because it provides type safety from DB to Resolver.

**Q325: How do you expose REST APIs via GraphQL Mesh?**
Define a "Source" (OpenAPI spec of the REST API). Mesh generates a GraphQL schema and converts incoming GraphQL queries into HTTP requests to the REST API.

**Q326: How do you integrate Elasticsearch with GraphQL?**
Create a `Search` type. The resolver constructs an Elasticsearch query DSL from the arguments and maps the ES hits to the GraphQL Schema.

**Q327: What is Neo4j GraphQL Library and how does it work?**
It translates GraphQL queries directly into Cypher (Neo4j's query language), allowing efficient graph-database fetches without writing manual queries.

**Q328: How do you consume GraphQL in Flutter?**
Use the `graphql_flutter` package. It provides widgets (Query, Mutation) and hooks similar to Apollo Client for React/OCaml.

**Q329: How do you implement GraphQL with FastAPI or Django?**
Use `Strawberry` (Code-first) or `Graphene` (Schema-first). `Strawberry` uses Python type hints to generate the schema, making it very pythonic.

**Q330: How do you use gRPC services inside GraphQL resolvers?**
Import the gRPC client proto. In the resolver, call the gRPC method. Await the response and map it to the GraphQL field.

**Q331: What are pros/cons of wrapping SOAP in GraphQL?**
**Pros**: Modernize the interface for clients. **Cons**: Complexity of parsing XML in resolvers. SOAP is stateful/strict; mapping to flexible GraphQL can be tricky.

**Q332: How do you expose Kafka streams via GraphQL Subscriptions?**
Use a Kafka consumer in your server. When a message arrives, publish it to the GraphQL PubSub system to trigger formatting updates to subscribers.

**Q333: How do you use GraphQL with Firebase?**
Use Cloud Functions to host the GraphQL Server (Apollo/Yoga). Resolvers call Firestore/Realtime DB.

**Q334: What are the trade-offs of using GraphQL with FaunaDB?**
Fauna has a native GraphQL API (upload schema -> get API). **Pros**: No server code needed. **Cons**: Tied to Fauna's capabilities; complex business logic might need UDFs (User Defined Functions).

**Q335: How do you expose GraphQL APIs from an AWS Lambda?**
Use `apollo-server-lambda` or `graphql-yoga`. The handler function helps map API Gateway events to GraphQL execution and back.

**Q336: How do you integrate Auth0 with GraphQL APIs?**
Client gets JWT from Auth0. Sends in Header. GraphQL Context decodes JWT (verifies signature via JWKS). Resolvers check claims.

**Q337: How do you use GraphQL with GitHub’s API?**
GitHub provides a public GraphQL API. You can stitch it into your own schema (Federation/Stitching) or query it directly from your resolvers as a data source.

**Q338: How can GraphQL federate multiple SaaS APIs?**
By treating each SaaS API (Stripe, Contentful) as a subgraph (using Mesh or custom wrappers) and composing them into your supergraph.

**Q339: What is the use of OpenAPI-to-GraphQL converters?**
To quickly bootstrap a GraphQL wrapper around a legacy REST API documented with Swagger/OpenAPI.

**Q340: How do you expose IoT or edge data using GraphQL?**
Usually via MQTT. The GraphQL server subscribes to MQTT topics and exposes the data via Subscriptions to web clients.

---

## 🔹 18. Developer Experience (DX) & CI/CD (341–360)

**Q341: How do you set up a GraphQL monorepo for multiple teams?**
Use Turborepo/Nx. Each team has a package/app (subgraph). A core package holds shared configs (ESLint, configs). A Gateway app composes them.

**Q342: How do you implement GitHub Actions for GraphQL linting?**
Add a step: `npm run lint:schema`. Use `graphql-schema-linter` or `graphql-inspector`. Fail the build if rules (naming, descriptions) are violated.

**Q343: What are good ESLint rules for GraphQL in frontend code?**
`graphql/template-strings` (validate queries against schema), `graphql/named-operations`, `graphql/no-deprecated-fields`.

**Q344: How do you set up Git pre-commit hooks for GraphQL validation?**
Use `husky` + `lint-staged`. Run `graphql-inspector diff` or linter on changed `.graphql` files before commit.

**Q345: How do you build a local schema explorer for devs?**
Host a local storybook or specialized doc site (Docusaurus) that pulls the schema and renders it using `spectaql` or `graphidoc`.

**Q346: What is the best way to mock remote GraphQL schemas for local dev?**
Download SDL. Use `graphql-tools` to add mocks (`addMocksToSchema`). Run a local server serving this mocked schema.

**Q347: How do you auto-document GraphQL queries in frontend PRs?**
Use tools that parse the PR for `.graphql` files and comment with the cost analysis or a visual representation of the query.

**Q348: What CLI tools are essential for GraphQL developers?**
`apollo` (Rover), `graphql-codegen`, `graphql-inspector`.

**Q349: What is `graphql-cli` and how do you use it?**
A command-line tool to manage common GraphQL workflows (init, ping, diff, playground). Note: Ecosystem has largely moved to vendor-specific CLIs (Rover) or dedicated tools (Codegen).

**Q350: How do you use `.graphqlconfig` in multi-project setups?**
It's a standard config file that defines where your schema lives, where operations are, and extensions for IDEs to understand the project structure.

**Q351: How do you use Apollo Rover CLI in CI pipelines?**
`rover subgraph publish` to push schema changes. `rover subgraph check` to validate breaking changes against the registry.

**Q352: How do you test federated services locally?**
Run all subgraphs. Run the Gateway pointing to local ports. Or use `rover dev` which orchestrates this local composition.

**Q353: How do you generate Postman collections from a GraphQL schema?**
Import the Schema (SDL) into Postman. It auto-generates a collection with example queries for every field/mutation.

**Q354: How do you validate that schema changes won’t break frontend apps?**
`graphql-inspector` compares the new schema against the known operations (extracted from client code). If a change affects a used field, it alerts.

**Q355: What is a changelog generator for GraphQL schemas?**
Tools that compare two SDL versions and output a Markdown list of added/removed/changed types.

**Q356: How do you monitor schema drift?**
Automated cron job that fetches the live schema and compares it to the "expected" git-stored schema. Alert on diffs.

**Q357: How do you automatically publish schema docs after every deploy?**
CI step generates static HTML docs (`spectaql`) and pushes them to S3/GitHub Pages.

**Q358: How do you manage secrets for GraphQL environments?**
Standard DevOps: Vault/AWS Secrets Manager. Inject as env vars. Never commit API keys in schema or code.

**Q359: How do you simulate load tests for GraphQL queries?**
`k6` has a graphql module. Write scripts that cycle through different query types, simulating user sessions.

**Q360: How do you apply git tagging/versioning for GraphQL schema releases?**
Tag the repo `schema-v1.2.0`. The registry (Apollo Studio) often uses the git commit hash (`SHAs`) to track variants.

---

## 🔹 19. Observability & Monitoring (361–380)

**Q361: How do you log field-level resolver performance?**
Use Tracing middleware. It records start/end time for every resolver. Aggregates them to find the slow fields.

**Q362: How do you track slow queries in GraphQL production logs?**
Log metadata (duration, query hash) for requests taking > 1s.

**Q363: How do you enable tracing for Apollo Server?**
Set `tracing: true` (deprecated in favor of plugins) or use the Usage Reporting plugin to send traces to Apollo Studio.

**Q364: What are Apollo Studio traces and how do you read them?**
Visual flame graphs showing the execution path. Long horizontal bars = slow resolvers. Verify parallel vs serial execution there.

**Q365: How do you send GraphQL metrics to Prometheus?**
Use `graphql-prometheus` or custom plugins to expose `/metrics` endpoint with counters for requests, errors, and duration histograms.

**Q366: How do you monitor subscriptions for dropped connections?**
Track WebSocket "Connect" and "Disconnect" events. Monitor the number of active connections gauge.

**Q367: How do you detect schema misuse or underutilized fields?**
Usage reporting. If a field hasn't been queried in 30 days (and traffic is normal), it's a candidate for deprecation.

**Q368: What is the role of `apollo-reporting` plugin?**
It asynchronously aggregates traces and sends them to Apollo's cloud, minimizing performance impact on the request itself.

**Q369: How do you create alerting for expensive queries?**
If `QueryComplexity > Threshold`, increment a metric. Alert if this metric spikes (potential DoS attempt).

**Q370: How do you view latency for federated subgraphs?**
The Gateway tracks how long each call to a downstream subgraph takes. The trace shows "Service A took 200ms".

**Q371: What’s the best way to visualize GraphQL usage metrics?**
Grafana (if using Prometheus) or specialized SaaS (Apollo Studio, Graphistry).

**Q372: How do you audit GraphQL requests for security threats?**
Look for: Deeply nested queries, high failure rates (auth brute force), or introspection queries from unknown IPs.

**Q373: What is `apollo-server-plugin-landing-page-graphql-playground`?**
A plugin to enable the "Playground" UI (interactive IDE) on the landing page of your server, usually disabled in prod.

**Q374: How do you attach correlation IDs across resolvers?**
Generate a Request ID at the Gateway. Pass it in headers (`X-Request-ID`) to subgraphs. Add it to logs.

**Q375: What log formats work best for GraphQL telemetry?**
JSON structured logs. `{ "op": "GetUser", "duration": 50, "errors": [...], "user": "..." }`.

**Q376: How do you build a query performance dashboard?**
Widgets: "p95 Latency", "Error Rate", "Requests/sec", "Top 5 Slowest Operations".

**Q377: How do you detect breaking changes via observability?**
Spike in `GRAPHQL_VALIDATION_ERROR` often means a client is sending a query that no longer matches the schema.

**Q378: How can observability inform schema refactoring?**
"Field X is always queried with Field Y". Suggests they should be grouped or merged.

**Q379: How do you set up distributed tracing in a GraphQL app?**
OpenTelemetry. Propagate trace context (headers) from Gateway -> Subgraphs -> DB.

**Q380: How do you identify and fix redundant resolvers?**
If a resolver always returns the parent object unchanged (trivial), it might be unnecessary overhead. Traces showing many tiny bars might indicate this.

---

## 🔹 20. Real-world Strategy (381–400)

**Q381: How would you introduce GraphQL to a legacy backend team?**
Start with a Gateway wrapper around their REST APIs. Show the frontend benefits first. Then gradually move logic.

**Q382: What’s the ROI of migrating from REST to GraphQL?**
Reduced network requests (mobile performance), faster frontend iteration (no waiting for backend endpoints), auto-documentation.

**Q383: When should you avoid using GraphQL?**
Simple server-to-server communication (gRPC is better), simple CRUD apps (REST is fine), or when file binary streaming is the primary use case.

**Q384: How do you pitch GraphQL to a non-technical stakeholder?**
"It's like a smart menu for our data. Instead of ordering a combo meal and throwing away fries (REST), we order exactly what we want. Saves bandwidth, makes the app faster."

**Q385: What types of projects benefit the most from GraphQL?**
Products with complex data relationships (Social Networks, Dashboards, E-commerce) and multi-platform clients (Web, iOS, Android).

**Q386: What’s the difference in client usage pattern between REST and GraphQL?**
REST: "Chatty" (many small requests). GraphQL: "Batchy" (fewer, larger requests).

**Q387: How do you manage schema growth in a large enterprise?**
Federation. Don't let it become a Monolith. Break it down by domain ownership.

**Q388: What’s the best onboarding path for a new dev joining a GraphQL team?**
Have them explore the graph with GraphiQL. Read the SDL comments. Write a simple query and mutation.

**Q389: How do you balance performance vs flexibility in schema design?**
Flexibility first (Graph design). Performance via Optimization (Dataloaders, Lookaheads) and Caching. Don't design ugly RPC-style schemas just for speed prematurely.

**Q390: What business risks come with exposing too much schema?**
Competitors scraping your data. Logic leaks. Maintenance burden of supporting everything forever.

**Q391: How do you sunset a deprecated GraphQL field safely?**
Mark deprecated. Monitor usage (Studio). Contact teams still using it. Brownout (random failures) to get attention. Remove.

**Q392: What are common pitfalls in GraphQL project kickoffs?**
treating it like SQL over HTTP (exposing table structures). ignoring caching strategy. Underestimating N+1 performance hit.

**Q393: How do you align frontend/backend teams on GraphQL practices?**
Joint Schema Design sessions. Frontend drives the requirements ("Schema Driven Development").

**Q394: How do you write GraphQL contracts in cross-team setups?**
SDL is the contract. PR reviews involving both teams are mandatory for schema changes.

**Q395: What are the signs of a well-maintained GraphQL API?**
Consistent naming, good descriptions, no breaking changes, usage of custom scalars, clear deprecation paths.

**Q396: What KPIs should a GraphQL platform team measure?**
Availability, p95 Latency, Breaking Change Rate, % of documented fields, Client Adoption rate.

**Q397: What is the long-term maintenance cost of a GraphQL API?**
Higher than REST initially (tooling setup). Lower later (less versioning hell, fewer endpoint adjustments).

**Q398: How do you support 100+ frontend apps consuming the same GraphQL?**
Strict Persisted Queries (Whitelisting). You must control what queries are allowed to run.

**Q399: How do you document and expose schema examples for 3rd-party devs?**
Public Developer Portal with embedded GraphiQL and tutorial workflows.

**Q400: How do you evaluate GraphQL vendor solutions?**
Check for: Federation support, Security features (WAF/Armor), Observability integration, and licensing costs (User/Seat vs Request volume).

---




