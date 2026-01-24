## 🟢 Basics, Schema & Type System (Questions 1-50)

### Question 1: What is GraphQL?

**Answer:**
GraphQL is a query language for APIs and a runtime for fulfilling those queries with your existing data. It provides a complete and understandable description of the data in your API, gives clients the power to ask for exactly what they need and nothing more.

**Code:**
```graphql
query {
  user(id: "1") {
    name
    email
  }
}
```

---

### Question 2: How does GraphQL differ from REST?

**Answer:**
*   **Data Fetching:** REST often requires multiple endpoints to fetch related data (over-fetching/under-fetching). GraphQL uses a single endpoint to fetch the exact data tree required.
*   **Versioning:** REST uses v1/v2 endpoints. GraphQL evolves the schema (adding fields, deprecating old ones).
*   **Error Handling:** REST uses HTTP Status Codes (404, 500). GraphQL usually returns 200 OK with an `errors` array.

---

### Question 3: What are the main features of GraphQL?

**Answer:**
1.  **Declarative Data Fetching:** Client dictates the response shape.
2.  **Strongly Typed:** Schema validates queries before execution.
3.  **Single Endpoint:** Simplifies network logic.
4.  **Real-time:** Subscriptions via WebSockets.

---

### Question 4: Explain the GraphQL type system.

**Answer:**
The Type System is the core of GraphQL.
*   **Scalars:** `Int`, `Float`, `String`, `Boolean`, `ID`.
*   **Object Types:** Custom objects (`User`, `Post`).
*   **Enums:** Fixed set of values.
*   **Interfaces/Unions:** Polymorphism.
*   **Input Types:** For passing objects to mutations.

---

### Question 5: What is a GraphQL schema?

**Answer:**
The contract between client and server, defined in Schema Definition Language (SDL). It defines all possible queries, mutations, and types.

**Code:**
```graphql
type Query {
  hello: String
}
```

---

### Question 6: What are queries in GraphQL?

**Answer:**
Read-only operations to fetch data. Equivalent to `GET` in REST.

**Code:**
```graphql
query GetUser {
  me {
    name
  }
}
```

---

### Question 7: What is a mutation in GraphQL?

**Answer:**
Write operations to modify data (Create, Update, Delete). Equivalent to `POST/PUT/DELETE` in REST. Mutations run sequentially.

**Code:**
```graphql
mutation CreatePost {
  addPost(title: "New") {
    id
    title
  }
}
```

---

### Question 8: What is a subscription in GraphQL?

**Answer:**
A long-lived connection (usually WebSocket) where the server pushes updates to the client when a specific event occurs.

**Code:**
```graphql
subscription OnMessage {
  messageAdded {
    text
    user
  }
}
```

---

### Question 9: How does GraphQL handle versioning?

**Answer:**
It avoids versioning entirely. New fields are added to the schema. Old fields are marked with `@deprecated` directive but remain functional until removed.

---

### Question 10: What are resolvers in GraphQL?

**Answer:**
Functions that provide the logic for filling data into a field. Steps:
1.  Query field.
2.  Runtime calls Resolver.
3.  Resolver fetches DB/API.
4.  Returns value.

**Code:**
```javascript
const resolvers = {
  Query: {
    user: (parent, args) => db.findUser(args.id)
  }
};
```

---

### Question 11: What is introspection in GraphQL?

**Answer:**
The ability to query the schema about itself.
`__schema` and `__type` fields allow tools like GraphiQL to auto-complete and show documentation.

**Code:**
```graphql
query {
  __schema {
    types {
      name
    }
  }
}
```

---

### Question 12: What is the role of `__typename` in GraphQL?

**Answer:**
A meta-field available on every object. It returns the name of the object type. Essential for client-side caching (Apollo) and handling Union/Interface fragments.

---

### Question 13: How do you define custom scalar types in GraphQL?

**Answer:**
Define in Schema: `scalar Date`.
Implement in Resolver: define `serialize`, `parseValue`, and `parseLiteral`.

**Code:**
```javascript
const dateScalar = new GraphQLScalarType({
  name: 'Date',
  serialize(value) { return value.toISOString(); }
});
```

---

### Question 14: What is the difference between input and output types?

**Answer:**
*   **Output Types (`type`):** Can contain fields and arguments. returned by queries.
*   **Input Types (`input`):** Passed *argument* to mutations. Cannot contain arguments.

**Code:**
```graphql
input UserInput { name: String! }
type Mutation { createUser(data: UserInput): User }
```

---

### Question 15: How are enums used in GraphQL?

**Answer:**
Restrict a field to a specific set of constants.

**Code:**
```graphql
enum Role { ADMIN, USER }
type User { role: Role }
```

---

### Question 16: Explain the concept of nullability in GraphQL.

**Answer:**
Fields are nullable by default (`String`).
`!` makes them non-nullable (`String!`).
If a non-nullable field errors, the error propagates up to the nearest nullable parent.

---

### Question 17: What is the purpose of fragments?

**Answer:**
Reusable chunks of logic in queries. Helps avoid code duplication on the client.

**Code:**
```graphql
fragment UserFields on User {
  id
  name
}
query {
  users { ...UserFields }
}
```

---

### Question 18: What is the difference between inline fragments and named fragments?

**Answer:**
*   **Named:** `fragment Name on Type { ... }`. Reusable.
*   **Inline:** `... on Type { ... }`. Used inside the query, typically for Unions (switching based on type).

---

### Question 19: Can you explain how aliases work in GraphQL?

**Answer:**
Rename the result key of a field. Essential when querying the same field twice with different arguments.

**Code:**
```graphql
query {
  admin: user(id: 1) { name }
  guest: user(id: 2) { name }
}
```

---

### Question 20: What are variables in GraphQL?

**Answer:**
Dynamic values passed separately from the query string. Prevents string concatenation/injection.

**Code:**
```graphql
query GetUser($id: ID!) {
  user(id: $id) { name }
}
# Variables: { "id": "123" }
```

---

### Question 21: What is a root type in GraphQL?

**Answer:**
The top-level types that serve as entry points: `Query`, `Mutation`, `Subscription`.

---

### Question 22: How do you structure a GraphQL schema?

**Answer:**
Usually split into domain-specific modules (User, Product).
Combined using `makeExecutableSchema` or `mergeTypeDefs`.

---

### Question 23: What are object types in GraphQL?

**Answer:**
The most common type. Represents a shape with fields.
`type Book { title: String, author: Author }`.

---

### Question 24: How do you create relationships between types?

**Answer:**
By defining a field on Type A that returns Type B. The resolver for that field handles the "Join".

**Code:**
```graphql
type Author {
  posts: [Post] 
  # Resolver for 'posts' will fetch posts where author_id = parent.id
}
```

---

### Question 25: How are interfaces used in GraphQL?

**Answer:**
Abstract type. Defines fields that implementing types *must* support.

**Code:**
```graphql
interface Character { name: String! }
type Human implements Character { name: String! height: Float }
```

---

### Question 26: What are unions in GraphQL?

**Answer:**
Returns one of several Object types. No shared fields required.

**Code:**
```graphql
union Result = User | Error
```

---

### Question 27: How does GraphQL handle recursive types?

**Answer:**
Natively supported. A type can reference itself.

**Code:**
```graphql
type Comment {
  text: String
  replies: [Comment]
}
```

---

### Question 28: What are directives in GraphQL?

**Answer:**
Instructions preceded by `@` that modify execution or schema behavior.
Built-in: `@include`, `@skip`, `@deprecated`.

---

### Question 29: What is the `@include` and `@skip` directive?

**Answer:**
Dynamic conditional fetching.
`field @include(if: $bool)`: fetch only if true.
`field @skip(if: $bool)`: skip if true.

---

### Question 30: What is the `@deprecated` directive and how do you use it?

**Answer:**
Marks a field as obsolete. It still works but tools warn the developer.

**Code:**
```graphql
type User {
  fullname: String @deprecated(reason: "Use name instead")
}
```

---

### Question 31: How do you write custom directives?

**Answer:**
1.  Define in Schema: `directive @upper on FIELD_DEFINITION`.
2.  Implement logic: Transformer function that wraps the field resolver to convert result to Uppercase.

---

### Question 32: How do you ensure schema validation?

**Answer:**
GraphQL engine validates Schema at startup (e.g., ensuring implementing types have all interface fields).
At runtime, it validates incoming queries against the schema.

---

### Question 33: How do you modularize a GraphQL schema?

**Answer:**
Split `typeDefs` into `.graphql` files.
Use `graphql-tools` to merge them.
Or use **GraphQL Modules** library.

---

### Question 34: What is schema stitching?**

**Answer:**
Combining multiple remote/local GraphQL schemas into one Gateway schema manually. Deprecated in favor of Federation.

---

### Question 35: What is schema delegation?

**Answer:**
Forwarding a part of a query from the Gateway to a Subschema (another service) to resolve.

---

### Question 36: What is the difference between schema-first and code-first approaches?

**Answer:**
*   **Schema-First:** Write `schema.graphql` (SDL). Implement Resolvers. (Apollo Default).
*   **Code-First:** Write TS Classes/Decorators. Generate SDL from code. (Nexus, TypeGraphQL). Better Type safety.

---

### Question 37: What tools are used to build GraphQL schemas?**

**Answer:**
`Apollo Server`, `graphql.js`, `TypeGraphQL`, `Nexus`, `Hasura` (No-code).

---

### Question 38: How do you document a GraphQL schema?

**Answer:**
Use Markdown in SDL (Triple quotes).

**Code:**
```graphql
"""
Represents a customer in the system
"""
type User { ... }
```

---

### Question 39: What is the difference between SDL and introspection schema?

**Answer:**
*   **SDL:** Human readable text (`type User { ... }`).
*   **Introspection:** JSON result identifying types/fields for tooling.

---

### Question 40: How can you expose multiple versions of a schema?

**Answer:**
Since GraphQL has 1 endpoint, you can mount different schemas on different paths:
`/graphql/v1` -> Schema 1
`/graphql/v2` -> Schema 2

---

### Question 41: How does a resolver function work?

**Answer:**
It populates data for a specific field.
`Query.user` resolver finds the user.
`User.posts` resolver finds posts for that user.

---

### Question 42: What is the resolver signature?

**Answer:**
`(parent, args, context, info)`.
*   **Parent:** Result of previous resolver.
*   **Args:** Query arguments.
*   **Context:** Global state (Auth, DB connection).
*   **Info:** AST of the specific field query.

---

### Question 43: What is the difference between field-level and type-level resolvers?

**Answer:**
*   **Field-level:** Specifically resolves `User.email`.
*   **Type-level:** `isTypeOf` check (rarely used). Or sometimes "Root Resolver" refers to Query/Mutation fields.

---

### Question 44: What is the parent (or root) argument in resolvers?

**Answer:**
The object returned by the parent field.
For `Query.me`, parent is `undefined`.
For `User.address`, parent is the `User` object.

---

### Question 45: How can you resolve nested fields?

**Answer:**
GraphQL does this automatically by chaining resolvers.
You just write the resolver for the child field, assuming `parent` contains the foreign key needed to fetch the child.

---

### Question 46: How can you handle errors in resolvers?

**Answer:**
Throw an `Error`.
`if (!user) throw new Error("Not found");`
GraphQL catches it, nulls the field, and adds to `errors` array.

---

### Question 47: What are resolver chains and how do they work?**

**Answer:**
Execution flows depth-first.
Query -> User Resolver -> (waits) -> returns User.
User -> Address Resolver -> returns Address.

---

### Question 48: How do you access context in resolvers?**

**Answer:**
It's the 3rd argument.
`const { user, db } = context;`
Context is created once per request (e.g. extracting token from headers).

---

### Question 49: How do you secure resolver logic?

**Answer:**
Check `context.user`.
`if (!context.user.isAdmin) throw new ForbiddenError()`.

---

### Question 50: How do you batch resolvers?

**Answer:**
**DataLoader**.
Instead of running SQL query in resolver, call `loader.load(id)`.
Loader waits a tick, gathers all IDs, runs one `SELECT * FROM table WHERE id IN (...)`, distributes results.
