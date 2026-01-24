## 🟢 Client-Side & Error Handling (Questions 101-150)

### Question 101: How do you integrate GraphQL with React?

**Answer:**
Use **Apollo Client**.
1.  Wrap App in `ApolloProvider`.
2.  Use `useQuery` hook.

**Code:**
```javascript
const { loading, error, data } = useQuery(GET_DOGS);
if (loading) return <p>Loading...</p>;
```

---

### Question 102: How does Apollo Client cache data?

**Answer:**
**Normalization.**
Splits response into objects.
Keys them by `__typename` + `id` (e.g., `User:1`).
Stores them in a flat lookup table.
Updates automatically if Mutation returns same ID.

---

### Question 103: How do you invalidate or evict Apollo Client cache?

**Answer:**
`cache.evict({ id: 'User:1' })`.
`cache.gc()` (Garbage Collect).
Or `refetchQueries` option in Mutation to force network reload.

---

### Question 104: What is optimistic UI in Apollo Client?

**Answer:**
Assume success. Update UI immediately.
If server fails later, roll back.

**Code:**
```javascript
useMutation(UPDATE_COMMENT, {
  optimisticResponse: {
    __typename: "Mutation",
    updateComment: { id: 1, content: "New Content", __typename: "Comment" }
  }
})
```

---

### Question 105: How do you handle loading and error states with Apollo Client?

**Answer:**
`useQuery` returns flags.
`const { loading, error } = useQuery(...)`.
Conditional rendering logic.

---

### Question 106: How do you write local-only fields in Apollo Client?

**Answer:**
Directive `@client`.
`query { user { isLoggedIn @client } }`.
Apollo resolves this from local Cache/State, not Network.

---

### Question 107: How do you manage state with Apollo Client?

**Answer:**
Reactive variables (`makeVar`).
Link them to schema policies.
Use them as a global store (Alternative to Redux).

---

### Question 108: What is Apollo Link and how does it work?

**Answer:**
Middleware chain for Client.
Request -> Link 1 (Auth) -> Link 2 (Error) -> Link 3 (Http) -> Server.
Composable network stack.

---

### Question 109: How do you use fragments with Apollo Client?

**Answer:**
Define fragment.
Spread in Query.

```javascript
const NAME_FRAGMENT = gql`fragment NameParts on User { first last }`;
const QUERY = gql`
  query { user { ...NameParts } }
  ${NAME_FRAGMENT}
`;
```

---

### Question 110: How do you execute a query manually using Apollo Client?

**Answer:**
`useLazyQuery`.
`const [getDog, { data }] = useLazyQuery(GET_DOGS);`
Call `getDog()` when button clicked.

---

### Question 111: How do you handle polling in GraphQL queries?

**Answer:**
`useQuery(QUERY, { pollInterval: 500 })`.
Auto-refetches every 0.5s.

---

### Question 112: What is `useQuery` vs `useLazyQuery`?

**Answer:**
*   `useQuery`: Runs on Mount.
*   `useLazyQuery`: Runs on Event (Click).

---

### Question 113: How can you cancel a GraphQL request?**

**Answer:**
Use `AbortController`.
Pass signal to context.
Apollo also has internal deduping that might cancel previous if new one starts.

---

### Question 114: What’s the difference between Apollo Client and Relay?

**Answer:**
*   **Apollo:** Flexible, easy hooks, works with any schema.
*   **Relay:** Strict, requires Compiler, High Perf, requires Schema adherence (Global ID).

---

### Question 115: How do you use GraphQL in Angular applications?

**Answer:**
`apollo-angular`.
Uses **RxJS** Observables.
`this.apollo.watchQuery({...}).valueChanges.subscribe()`.

---

### Question 116: How do you integrate GraphQL with Vue.js?

**Answer:**
`v4` (Vue Apollo).
Composition API: `const { result } = useQuery(...)`.

---

### Question 117: How do you manage multiple GraphQL endpoints in a frontend app?

**Answer:**
`ApolloLink.split`.
Router logic in Link chain.
If query name starts with `Auth`, route to Auth Server. Else Main Server.

---

### Question 118: How do you debug client-side GraphQL issues?

**Answer:**
**Apollo Client DevTools** (Chrome Extension).
See Cache state. See active queries.

---

### Question 119: How do you persist Apollo cache across page reloads?

**Answer:**
`apollo3-cache-persist`.
Syncs Cache <-> localStorage.
App starts with data already loaded.

---

### Question 120: What is the purpose of Apollo DevTools?

**Answer:**
Inspection. Mutation testing (Run mutation from tool). Cache visualizations.

---

### Question 121: How do you define custom error codes in GraphQL?

**Answer:**
Server returns `extensions`.
`throw new GraphQLError('Msg', { extensions: { code: 'BAD_INPUT' } })`.

---

### Question 122: What is the format of a GraphQL error response?

**Answer:**
```json
{
  "errors": [
     { "message": "...", "locations": [], "path": ["user"], "extensions": {} }
  ],
  "data": null // or partial data
}
```

---

### Question 123: How do you distinguish between GraphQL and network errors?

**Answer:**
*   **Network:** 400/500 status. fetch throws.
*   **GraphQL:** 200 OK. `result.errors` array exists.

---

### Question 124: What is the best way to log GraphQL errors?

**Answer:**
DidEncounterErrors plugin (Apollo).
Send full error + Path to Sentry/Datadog.

---

### Question 125: How do you mask internal errors from clients?

**Answer:**
`formatError` function.
If error is "DatabaseConnectionFailed", return "Internal Error".
Keep original log on server.

---

### Question 126: How can you use extensions field in error responses?

**Answer:**
Validation details.
`extensions: { validationErrors: { email: "Invalid" } }`.
Client uses this to show red borders on inputs.

---

### Question 127: How do you create global error handlers for GraphQL?

**Answer:**
**Apollo Link onError**.
`onError(({ graphQLErrors }) => { if (code === 'UNAUTHENTICATED') logout(); })`.

---

### Question 128: What is GraphQL Shield and how is it used for auth/errors?

**Answer:**
Middleware.
Define "Rule Tree".
If rule fails, it throws localized error before resolver runs.

---

### Question 129: How do you debug circular references in resolvers?

**Answer:**
Logs showing repeated path `user.friends.user.friends...`.
Fix with Max Depth limit.

---

### Question 130: How do you catch field-level errors in resolvers?

**Answer:**
Try/Catch inside resolver.
Return `null`.
Main query succeeds (200), but that one field is null + verified error in `errors` list.

---

### Question 131: How do you report validation errors in input objects?

**Answer:**
Check all fields.
Collect errors.
Throw `UserInputError` (Apollo) with array of failures.

---

### Question 132: How do you track slow queries in production?

**Answer:**
Apollo Studio Tracing.
Shows "Flame graph" of resolver execution time.

---

### Question 133: How do you debug failed subscriptions?

**Answer:**
Check "Connection Ack".
WS logic is tricky.
Verify Heartbeat interval.

---

### Question 134: How can error boundaries be implemented in a GraphQL frontend?

**Answer:**
React internal feature (`componentDidCatch`).
Wrap Query component. Cases where `data` is undefined due to error.

---

### Question 135: How do you test error scenarios in GraphQL?

**Answer:**
`MockedProvider` with `error: new Error(...)`.
Assert UI shows "Error" banner.

---

### Question 136: How can you log queries with sensitive data excluded?

**Answer:**
Strip `variables.password`.
Or enable "Private Variables" in Apollo Studio setting.

---

### Question 137: How do you validate incoming queries for correctness?

**Answer:**
`validate(schema, parse(query))`.
Engine does this automatically. Status 400 if syntax invalid.

---

### Question 138: How do you detect unused types or fields in a schema?

**Answer:**
Schema Coverage tools.
Analyze traffic logs for 30 days.
If field `Address.zip` count is 0, safe to deprecate.

---

### Question 139: How do you simulate backend errors during frontend testing?

**Answer:**
(See Q135). Mock response with `errors` array.

---

### Question 140: What is Apollo Server’s `formatError` function used for?

**Answer:**
Final transformation of error object before JSON serialization.
Remove stack trace (`delete err.stack`).
Translate message.

---

### Question 141: What is GraphiQL and how is it different from Playground?

**Answer:**
*   **GraphiQL:** Reference React component. Simple.
*   **Playground:** Fancy wrapper (Prisma/Apollo). Headers tab, multiple tabs.
*   **Apollo Explorer:** Modern standard.

---

### Question 142: What is GraphQL Voyager?

**Answer:**
Visualization tool.
Shows Schema as a Node Graph (ER Diagram).
Good for understanding relationships.

---

### Question 143: What is GraphQL Inspector used for?

**Answer:**
CI Tool.
Diffs Schema (PR vs Master).
Detects Breaking Changes.

---

### Question 144: What is GraphQL Code Generator?

**Answer:**
Best tool.
Schema + Query -> TypeScript Interfaces.
`user.name` is typed. `user.nmae` errors at compile time.

---

### Question 145: How do you generate TypeScript types from a schema?

**Answer:**
`graphql-codegen`.
Run it in watch mode.

---

### Question 146: What is `graphql-tag` used for?**

**Answer:**
It provides the `gql` tag.
`const query = gql'{ ... }'`.
Parses string to AST Object.

---

### Question 147: How do you use Postman with GraphQL?

**Answer:**
Postman supports GraphQL Body type.
Can fetch Schema (Introspection) to enable Autocomplete.

---

### Question 148: What is `graphql-tools` and how is it useful?

**Answer:**
Utilities for:
Stitching.
Mocking (`addMocksToSchema`).
Directives.

---

### Question 149: What is GraphQL Mesh used for?

**Answer:**
Gateway.
Input: OpenAPI / SOAP / gRPC.
Output: GraphQL.
No code (Config based).

---

### Question 150: What are some popular GraphQL servers in Node.js?

**Answer:**
1.  **Apollo Server**: Standard.
2.  **GraphQL Yoga**: Modern, lighter.
3.  **Mercurius**: Fastify plugin (Fast).
