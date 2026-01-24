## 🟢 Security, AI & Edge Cases (Questions 451-500)

### Question 451: How do you enable field-level auth logic in federated schemas?

**Answer:**
The Gateway passes the User Context (headers) to the Subgraph.
The Subgraph's resolver performs the check.
If failed, Subgraph returns `null`/error. Gateway aggregates it.

---

### Question 452: How do you implement rate limiting per user in GraphQL?

**Answer:**
Middleware.
Calculate query cost or request count.
Store in Redis with TTL 1 minute.
If count > Limit, throw `GraphQLError("Too Many Requests")`.

---

### Question 453: How do you integrate OAuth2 or OpenID Connect with GraphQL?

**Answer:**
Identity Provider (Auth0/Okta) handles login flow (Redirects).
Client gets Access Token.
GraphQL Server validates Access Token (JWT) on every request.

---

### Question 454: How do you block malicious queries based on behavior?

**Answer:**
Anomaly detection.
"User X is querying `users { id }` 100 times/sec with different variables".
Block User ID.

---

### Question 455: How do you validate request headers for GraphQL auth?

**Answer:**
At the Context creation step.
`const token = req.headers.authorization || '';`
`if (!token.startsWith('Bearer ')) throw new AuthenticationError();`

---

### Question 456: How do you attach session context securely across federated services?

**Answer:**
Gateway validates the public token.
Gateway signs an internal JWT (short lived) with user claims.
Passes internal JWT to Subgraphs. Subgraphs trust Gateway's signature.

---

### Question 457: How do you avoid leaking sensitive data in GraphQL errors?

**Answer:**
Production Mode.
`apollo-server` automatically strips stack traces if `NODE_ENV=production`.
Use generic error codes.

---

### Question 458: How do you implement attribute-based access control (ABAC) in resolvers?

**Answer:**
Dynamic checks.
`canAccess(user, resource)`.
`if (resource.owner === user.id || user.department === resource.dept)`
Complex logic, often delegated to a specialized policy engine like **OPA** (Open Policy Agent).

---

### Question 459: What tools can help scan GraphQL APIs for security risks?

**Answer:**
**InQL** (Scanner).
**Clairvoyance** (Suggests introspection paths).
**GraphQL Map**.

---

### Question 460: How do you verify schema compliance with security policies?

**Answer:**
Schema Linter rules.
"Field `password` must have `@sensitive` directive".
"Query type must typically require auth".

---

### Question 461: How can GraphQL be used to expose ML model predictions?

**Answer:**
`type Prediction { output: String, probability: Float }`.
Resolver makes RPC call to TensorFlow Serving container.

---

### Question 462: How do you integrate GraphQL with vector databases like Pinecone?

**Answer:**
Query `search(text: String)`.
Resolver:
1.  Call OpenAI to get Embedding Vector.
2.  Query Pinecone with Vector.
3.  Return matched IDs.

---

### Question 463: How do you handle large language model APIs via GraphQL?

**Answer:**
LLMs are slow.
Use `@defer` or Subscriptions.
`subscription { chatResponse(prompt: "...") }`.
Stream tokens as they arrive.

---

### Question 464: How can GraphQL support real-time analytics dashboards?

**Answer:**
Subscriptions.
`subscription { dashboardStats { activeUsers, revenue } }`.
Server publishes updates every 10s or on event.

---

### Question 465: How do you serve model inference pipelines via GraphQL?

**Answer:**
Async Job.
Mutation `startInference` -> returns Job ID.
Client subscribes to Job ID for result.

---

### Question 466: How do you expose time-series anomaly detection in GraphQL APIs?

**Answer:**
`type Anomaly { time: String, severity: Float }`.
Resolver queries Time-Series DB (InfluxDB) for points exceeding threshold.

---

### Question 467: How do you cache expensive ML predictions in resolvers?

**Answer:**
Redis Cache.
Key: `predict:${hash(inputParams)}`.
ML models are often pure functions (Same Input -> Same Output).

---

### Question 468: How do you use GraphQL to fetch contextual embeddings for AI tasks?

**Answer:**
Batch resolver.
Take list of Text.
Call Embedding API in batch (OpenAI allows batch inputs).
Return list of Vectors.

---

### Question 469: How do you structure input types for prompt engineering via GraphQL?

**Answer:**
`input PromptOptions { temperature: Float, stopSequences: [String], maxTokens: Int }`.
Pass directly to LLM provider.

---

### Question 470: How do you paginate through ML-generated recommendations?

**Answer:**
The Recommendation Engine returns a ranked list of Item IDs.
GraphQL Resolver implements slicing/cursor logic on that list.

---

### Question 471: How do you expose tuning parameters via a GraphQL interface?

**Answer:**
Admin Mutations.
`updateModelConfig(threshold: 0.8)`.
Useful for tweaking live algorithms without redeploying code.

---

### Question 472: What’s the best way to log model feedback via GraphQL mutations?

**Answer:**
`mutation submitFeedback(predictionId: ID, wasUseful: Boolean)`.
Write to Data Lake for retraining.

---

### Question 473: How do you manage experiment tracking data with GraphQL?

**Answer:**
`type ExperimentRun { parameters: JSON, metrics: JSON }`.
Flexible JSON type is often acceptable here because params vary wildly between experiments.

---

### Question 474: How do you expose real-time metrics with GraphQL Subscriptions?

**Answer:**
Feed from StatsD/Prometheus.
Publish to GraphQL subscription channel "metrics".

---

### Question 475: How can GraphQL help in AI-powered document search systems?

**Answer:**
Unified Search.
`search(text: "...")`.
Resolver queries DB (Keyword match) + Vector DB (Semantic match).
Merges and ranks results.

---

### Question 476: How do you integrate OpenAI APIs into GraphQL resolvers?

**Answer:**
Standard HTTP call.
Use `openai` npm package.
Map `completion.choices[0].text` to GraphQL field.

---

### Question 477: How do you throttle usage of AI-based features via GraphQL?

**Answer:**
Token Bucket.
AI features consume 10x more tokens than DB reads.
Reject if bucket empty.

---

### Question 478: How do you return explanations (XAI) via GraphQL queries?

**Answer:**
`type Explanation { feature: String, weight: Float }`.
`prediction { result, explanation }`.

---

### Question 479: How do you model semantic search layers in GraphQL schema?

**Answer:**
`interface SearchResult`.
Resolvers fetch IDs from Vector DB, then hydrate full objects from SQL DB.

---

### Question 480: How can you use GraphQL to build an AI assistant interface?

**Answer:**
`type Message { role: String, content: String }`.
Mutation `sendMessage(history: [MessageInput])`.
Returns `Message` (Bot response).

---

### Question 481: What are the dangers of exposing internal IDs in schema?

**Answer:**
If you change DB (Postgres -> Mongo), ID format changes (Int -> ObjectId). Breaking change.
**Fix:** Use Global Opaque IDs (base64).

---

### Question 482: How do you handle orphaned fields in evolving schemas?

**Answer:**
Fields declared but not used by any client.
Detect with Studio.
Remove them to keep schema clean.

---

### Question 483: Why should you avoid dynamic field names in GraphQL?

**Answer:**
`{ user_1: ..., user_2: ... }`.
Can't be typed in Schema.
Use Lists: `users { id, ... }`.

---

### Question 484: What are the risks of resolver-level side effects?

**Answer:**
Queries executed in parallel. Order not guaranteed.
If `query { deleteUser }` exists also `query { fetchUser }`... unpredictable.
Use Mutations for side effects.

---

### Question 485: What’s wrong with returning booleans for mutation success?

**Answer:**
`deleteUser: Boolean`.
If false, why? No error info.
If true, client might want the deleted ID to remove from UI list.
Always return an Object.

---

### Question 486: What are problems with auto-generated schemas from databases?

**Answer:**
Exposes implementation details.
Hard to refactor DB without breaking API.
Lack of business-logic layer.

---

### Question 487: Why is it bad to have a single “root” mutation field?

**Answer:**
`mutation { do(action: "create", entity: "user") }`.
Loses all type safety and discovery (Docs).
Use explicit mutations: `createUser`.

---

### Question 488: What are the issues with overusing interfaces and unions?**

**Answer:**
Client has to spread fragments for every case.
`... on A ... on B ... on C`.
If you add `D`, client must update code to see `D`.

---

### Question 489: What problems arise from using scalar `JSON` types everywhere?

**Answer:**
"Mystery Meat" data.
Client doesn't know structure.
No type validation.
Valid only for truly dynamic data (configs, logs).

---

### Question 490: Why is schema nesting depth a scalability concern?

**Answer:**
Complexity grows exponentially.
`friends(first:10) { friends(first:10) { ... } }`.
User can crash server easily.

---

### Question 491: What’s wrong with exposing file uploads as simple strings?

**Answer:**
Base64 encoding increases memory usage and payload size significantly.
Blocks Node event loop during decoding.

---

### Question 492: What happens when too many clients share the same fragment?

**Answer:**
Fragile.
Mobile changes fragment to add `avatar`.
Web now downloads `avatar` (unused) and gets slower.
Co-locate fragments with components.

---

### Question 493: What causes memory leaks in GraphQL subscriptions?

**Answer:**
Closure retaining context.
Or failing to unsubscribe from the PubSub engine on client disconnect.

---

### Question 494: What are the issues with exposing `__typename` to clients?

**Answer:**
If your internal class names are `"LegacyUserWrapperImpl"`, it looks bad.
Keep Type names semantic (`User`).

---

### Question 495: Why can using custom scalars lead to inconsistencies?**

**Answer:**
Different teams implement `Date`.
One expects milliseconds. Another ISO string.
Must use a shared library (`graphql-scalars`).

---

### Question 496: What happens when you re-use input types across unrelated mutations?**

**Answer:**
Coupling.
You add `age` to `UserInput` for Registration.
But `UpdateUser` doesn't support changing age. Now it looks like you can update it.
Separate `CreateUserInput` and `UpdateUserInput`.

---

### Question 497: How can poor nullability design lead to frontend bugs?

**Answer:**
`name: String!`.
DB has one record with null name (bad migration).
Query fails entirely.
UI shows "Something went wrong" instead of just hiding the name.

---

### Question 498: What’s the danger of relying only on introspection for docs?

**Answer:**
Introspection has no context.
"What does `score` mean? Is high good or bad?".
You need human-written descriptions.

---

### Question 499: Why should you avoid overly generic field names like `data` or `info`?

**Answer:**
Hard to grep.
Ambiguous meaning.
Future collisions.

---

### Question 500: What is schema fatigue and how do you prevent it?

**Answer:**
Too huge schema.
Prevent by modularizing.
Use "Viewer" pattern to scope logged-in user fields.
Use Federation to split domains.
