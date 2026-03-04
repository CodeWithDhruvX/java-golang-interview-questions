# 🟣 GraphQL Basics & Schema — Interview Questions (Service-Based Companies)

This document covers foundational GraphQL concepts commonly tested in service-based company interviews, including TCS, Infosys, Wipro, Capgemini, HCL, and similar. These questions appear in 1–5 years experience rounds.

---

### Q1: What is GraphQL and how does it differ from REST?

**Answer:**
GraphQL is a **query language for APIs** and a runtime for executing queries, developed by Facebook in 2012 and open-sourced in 2015. Unlike REST, which exposes data through multiple fixed endpoints, GraphQL exposes a **single endpoint** and lets the client ask for the exact data it needs.

**Key Differences:**

| Feature | REST | GraphQL |
|---|---|---|
| Endpoints | Multiple (`/users`, `/posts`) | Single (`/graphql`) |
| Data Fetching | Fixed response shape | Client-defined response |
| Over-fetching | Common (extra fields returned) | Eliminated |
| Under-fetching | Common (multiple calls needed) | Eliminated |
| Versioning | `/v1`, `/v2` endpoints | Schema evolution (`@deprecated`) |
| Error Handling | HTTP status codes (404, 500) | `200 OK` + `errors` array |
| Documentation | Swagger/OpenAPI | Introspection / GraphiQL |

**Example — REST vs GraphQL:**
```bash
# REST (requires 2 calls)
GET /users/123          → returns ALL user fields
GET /users/123/posts    → returns ALL post fields

# GraphQL (1 call, exact fields)
query {
  user(id: "123") {
    name
    email
    posts {
      title
    }
  }
}
```

**When to prefer GraphQL:**
- Mobile apps (data bandwidth sensitive)
- Aggregating data from multiple microservices
- Rapidly evolving APIs

---

### Q2: Explain the GraphQL schema and its core components.

**Answer:**
The **schema** is the heart of GraphQL — it's the contract between the **client and server**, written in Schema Definition Language (SDL). It defines all available types, queries, mutations, and subscriptions.

**Core Components:**

**1. Object Types (most common):**
```graphql
type User {
  id: ID!
  name: String!
  email: String!
  age: Int
  posts: [Post!]!
}
```

**2. Root Types (entry points):**
```graphql
type Query {
  user(id: ID!): User
  users: [User!]!
}

type Mutation {
  createUser(name: String!, email: String!): User!
  updateUser(id: ID!, name: String): User
  deleteUser(id: ID!): Boolean!
}

type Subscription {
  userCreated: User!
}
```

**3. Scalar Types (leaf values):**
- Built-in: `Int`, `Float`, `String`, `Boolean`, `ID`
- Custom: `scalar Date`, `scalar JSON`, `scalar Upload`

**4. Enum Types:**
```graphql
enum UserRole {
  ADMIN
  USER
  GUEST
}
```

**5. Input Types (for mutations):**
```graphql
input CreateUserInput {
  name: String!
  email: String!
  role: UserRole!
}
```

**6. Interface and Union:**
```graphql
interface Animal { name: String! }
type Dog implements Animal { name: String!, breed: String! }

union SearchResult = User | Post | Product
```

---

### Q3: What are resolvers in GraphQL and how do they work?

**Answer:**
A **resolver** is a function that returns the value for a specific field in the schema. GraphQL calls the appropriate resolver for each field in a query.

**Resolver Signature:**
```javascript
fieldName(parent, args, context, info) { ... }
```
- **parent**: The result from the parent field's resolver
- **args**: Arguments passed to the field (e.g., `id: "123"`)
- **context**: Shared request-level data (auth user, DB connection, dataloaders)
- **info**: Metadata about the query (field name, AST, path, etc.)

**Example:**
```javascript
const resolvers = {
  Query: {
    // parent = undefined for root queries
    user: async (parent, args, context) => {
      return await context.db.findUserById(args.id);
    },

    users: async (parent, args, context) => {
      return await context.db.getAllUsers();
    }
  },

  User: {
    // parent = User object returned by Query.user
    posts: async (parent, args, context) => {
      return await context.db.getPostsByUserId(parent.id);
    }
  },

  Mutation: {
    createUser: async (parent, args, context) => {
      const { name, email } = args;
      const user = await context.db.createUser({ name, email });
      return user;
    }
  }
};
```

**Default Resolvers:** If you don't define a resolver for a field, GraphQL uses a **default resolver** that simply returns `parent[fieldName]`. So if your DB returns an object `{ id: 1, name: "John" }`, the `User.name` field resolver is automatically handled.

---

### Q4: What are queries, mutations, and subscriptions?

**Answer:**

**Query (Read):** Fetch data without side effects. Analogous to `GET` in REST.
```graphql
query GetUser($id: ID!) {
  user(id: $id) {
    id
    name
    email
    posts {
      title
      createdAt
    }
  }
}
# Variables: { "id": "123" }
```

**Mutation (Write):** Modify data (Create, Update, Delete). Analogous to `POST/PUT/DELETE`. Mutations execute **sequentially** (unlike queries which run in parallel).
```graphql
mutation CreatePost($title: String!, $authorId: ID!) {
  createPost(title: $title, authorId: $authorId) {
    id
    title
    author {
      name
    }
  }
}
```

**Subscription (Real-time):** Maintain a long-lived connection (WebSocket) where the server pushes data when an event occurs.
```graphql
subscription OnNewMessage($roomId: ID!) {
  messageAdded(roomId: $roomId) {
    id
    text
    sender {
      name
    }
  }
}
```

**Key difference — Mutations are sequential:**
```graphql
# These run one after the other (sequential)
mutation {
  addItem(name: "A")
  addItem(name: "B")
}
```

---

### Q5: What are fragments and how do you use them?

**Answer:**
**Fragments** are reusable units of a query. They help reduce repetition and keep queries DRY (Don't Repeat Yourself).

**Named Fragments:**
```graphql
fragment UserBasicFields on User {
  id
  name
  email
}

fragment PostFields on Post {
  id
  title
  createdAt
}

query {
  user(id: "1") {
    ...UserBasicFields
    posts {
      ...PostFields
    }
  }
  admin: user(id: "2") {
    ...UserBasicFields
  }
}
```

**Inline Fragments** (used with Unions/Interfaces):
```graphql
union SearchResult = User | Post | Product

query Search($term: String!) {
  search(query: $term) {
    ... on User {
      name
      email
    }
    ... on Post {
      title
      author { name }
    }
    ... on Product {
      name
      price
    }
  }
}
```

**Benefits:**
- Avoid repetition in large queries
- Colocation: each component defines its own fragment (used in Relay/Apollo)
- Easier to maintain large query files

---

### Q6: Explain GraphQL variables, directives, and aliases.

**Answer:**

**Variables — Dynamic query arguments:**
Instead of hardcoding values in queries:
```graphql
# Without variables (hardcoded — bad practice)
query { user(id: "123") { name } }

# With variables (safe, reusable)
query GetUser($id: ID!, $includeEmail: Boolean = false) {
  user(id: $id) {
    name
    email @include(if: $includeEmail)
  }
}
# Variables JSON: { "id": "123", "includeEmail": true }
```

**Built-in Directives:**
- `@include(if: Boolean)` — Include field only if condition is true
- `@skip(if: Boolean)` — Skip field if condition is true
- `@deprecated(reason: String)` — Mark field as obsolete

```graphql
query GetProfile($showAddress: Boolean!) {
  user(id: "1") {
    name
    address @include(if: $showAddress) {
      city
      country
    }
  }
}
```

**Aliases — Multiple calls to the same field:**
```graphql
query {
  adminUser: user(id: "1") { name, role }
  guestUser: user(id: "2") { name, role }
}
```
Without aliases, this would be a naming conflict. Aliases let you rename the result key.

---

### Q7: What is GraphQL introspection and how is it used?

**Answer:**
**Introspection** is GraphQL's built-in capability to query the schema about itself. It allows tooling and developers to discover what types, queries, mutations, and fields are available.

**How it works:**
```graphql
# Query all types in the schema
query {
  __schema {
    types {
      name
      kind
    }
  }
}

# Query available queries
query {
  __schema {
    queryType {
      fields {
        name
        description
        args { name, type { name } }
      }
    }
  }
}

# Inspect a specific type
query {
  __type(name: "User") {
    fields {
      name
      type { name, kind }
      isDeprecated
      deprecationReason
    }
  }
}
```

**`__typename` meta-field:**
```graphql
query {
  search(query: "flight") {
    __typename    # Returns "User" or "Post" etc.
    ... on User { name }
    ... on Post { title }
  }
}
```

**Practical uses:**
- Powers **GraphiQL** / **Apollo Explorer** auto-complete
- Client code generation (TypeScript types from the schema)
- Schema documentation generation

> ⚠️ **Security:** Disable introspection in production APIs to prevent schema enumeration by attackers (`introspection: false` in Apollo Server config).

---

### Q8: What is a GraphQL context and how is it used for authentication?

**Answer:**
**Context** is a shared object passed to every resolver function. It's created once per request and is the standard way to share:
- Authenticated user identity
- Database connections
- DataLoader instances
- Request headers

**Setting up context in Apollo Server:**
```javascript
const server = new ApolloServer({
  typeDefs,
  resolvers,
  context: async ({ req }) => {
    // Extract JWT from header
    const token = req.headers.authorization?.replace('Bearer ', '');
    let user = null;

    if (token) {
      try {
        user = jwt.verify(token, process.env.JWT_SECRET);
      } catch (err) {
        // Invalid token — let resolvers handle authorization
      }
    }

    return {
      user,          // authenticated user or null
      db,            // database instances
      dataloaders    // DataLoader instances
    };
  }
});
```

**Using context in resolvers:**
```javascript
const resolvers = {
  Query: {
    myProfile: (parent, args, context) => {
      // Check authentication
      if (!context.user) {
        throw new Error('You must be logged in!');
      }
      return context.db.findUser(context.user.id);
    }
  },

  Mutation: {
    deletePost: (parent, { id }, context) => {
      if (!context.user) throw new Error('Unauthorized');
      if (context.user.role !== 'ADMIN') throw new Error('Forbidden');
      return context.db.deletePost(id);
    }
  }
};
```

---

*Prepared for technical screening rounds at service-based companies (TCS, Infosys, Wipro, Capgemini, HCL, Cognizant).*
