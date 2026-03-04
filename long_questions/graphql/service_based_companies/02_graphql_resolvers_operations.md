# 🟣 GraphQL Resolvers & Operations — Interview Questions (Service-Based Companies)

This document covers GraphQL resolver patterns and operational questions commonly tested in service-based company interviews. Focus on practical implementation for 2–5 years experience rounds.

---

### Q1: What is the N+1 problem in GraphQL and how do you solve it with DataLoader?

**Answer:**
The **N+1 problem** is a performance anti-pattern where fetching a list of N items causes N additional database queries for a related field.

**The Problem:**
```graphql
query {
  posts {       # 1 query to get 10 posts
    title
    author {    # 10 queries — one for each post's author!
      name
    }
  }
}
```
This generates **11 SQL queries** instead of 2. It's not visible in the schema but is a resolver-level bug.

**Root Cause:**
```javascript
const resolvers = {
  Post: {
    // Called once per post — causes N separate DB calls!
    author: (post, args, context) => {
      return context.db.query(`SELECT * FROM users WHERE id = ${post.authorId}`);
    }
  }
};
```

**Solution — DataLoader (request-level batching):**
DataLoader batches all IDs collected in a single tick and fires one bulk query.

```javascript
const DataLoader = require('dataloader');

// Create a DataLoader that batches user IDs
const createUserLoader = (db) => new DataLoader(async (userIds) => {
  // Called ONCE with all collected IDs: [1, 2, 3, ...]
  const users = await db.query(`SELECT * FROM users WHERE id IN (${userIds.join(',')})`);
  
  // IMPORTANT: Result order must match input order
  const userMap = users.reduce((map, user) => {
    map[user.id] = user;
    return map;
  }, {});
  
  return userIds.map(id => userMap[id]);
});

// Add loader to context (recreate per request!)
const context = ({ req }) => ({
  db,
  userLoader: createUserLoader(db)     // ← fresh loader per request
});

// Resolver — now uses the loader
const resolvers = {
  Post: {
    author: (post, args, context) => {
      // Queued and batched automatically
      return context.userLoader.load(post.authorId);   // No N+1!
    }
  }
};
```
Result: **2 SQL queries** regardless of how many posts are fetched.

---

### Q2: How do you handle errors in GraphQL?

**Answer:**
GraphQL's error handling follows a unique pattern compared to REST:
- Success and partial errors both return **HTTP 200**
- Errors are included in an `errors` array in the response body
- Partial data can be returned even when errors occur

**Response structure:**
```json
{
  "data": {
    "user": null,
    "posts": [{ "id": "1", "title": "Hello" }]
  },
  "errors": [
    {
      "message": "User not found",
      "locations": [{ "line": 2, "column": 3 }],
      "path": ["user"],
      "extensions": {
        "code": "USER_NOT_FOUND",
        "statusCode": 404
      }
    }
  ]
}
```

**Throwing errors in resolvers:**
```javascript
const { ApolloError, AuthenticationError, ForbiddenError, UserInputError } = require('apollo-server');

const resolvers = {
  Query: {
    user: async (parent, { id }, context) => {
      // Authentication check
      if (!context.user) {
        throw new AuthenticationError('Must be logged in');
      }

      const user = await context.db.findUser(id);

      // Not found
      if (!user) {
        throw new ApolloError('User not found', 'USER_NOT_FOUND', { userId: id });
      }

      // Authorization check
      if (context.user.id !== id && context.user.role !== 'ADMIN') {
        throw new ForbiddenError('Cannot view other users');
      }

      return user;
    }
  },

  Mutation: {
    createUser: async (parent, args) => {
      // Validation error
      if (!args.email.includes('@')) {
        throw new UserInputError('Invalid email format', {
          invalidArgs: ['email']
        });
      }
      // ...
    }
  }
};
```

**Custom error formatting (global):**
```javascript
const server = new ApolloServer({
  typeDefs,
  resolvers,
  formatError: (err) => {
    console.error('GraphQL Error:', err);
    // Don't expose internal error details in production
    if (err.extensions?.code === 'INTERNAL_SERVER_ERROR') {
      return new Error('Internal Server Error');
    }
    return err;
  }
});
```

---

### Q3: What is schema first vs code first approach in GraphQL?

**Answer:**

| Aspect | Schema-First | Code-First |
|---|---|---|
| Definition | Write SDL file first, then implement resolvers | Write code (classes/decorators), SDL is generated |
| Tools | Apollo Server, graphql-tools | TypeGraphQL, Nexus, Pothos |
| Type safety | SDL and resolvers can drift | Single source of truth (TS types = GraphQL types) |
| Developer experience | Clear SDL visible to all | Better IDE support, autocompletion |
| Flexibility | Easy to read/understand schema | More programmatic control |

---

**Schema-First (Apollo):**
```graphql
# schema.graphql
type User {
  id: ID!
  name: String!
  email: String!
}

type Query {
  user(id: ID!): User
}
```
```javascript
// resolvers.js
const resolvers = {
  Query: {
    user: (_, { id }) => db.findUser(id)
  }
};
const server = new ApolloServer({ typeDefs, resolvers });
```

**Code-First (TypeGraphQL):**
```typescript
// user.resolver.ts
@ObjectType()
class User {
  @Field(() => ID)
  id: string;

  @Field()
  name: string;

  @Field()
  email: string;
}

@Resolver()
class UserResolver {
  @Query(() => User, { nullable: true })
  async user(@Arg("id") id: string): Promise<User | undefined> {
    return db.findUser(id);
  }
}
// SDL is auto-generated from this code
```

**Service-based company preference:** Schema-first is more common and easier to understand in team environments. Code-first is preferred in mature TypeScript projects.

---

### Q4: How does pagination work in GraphQL? Explain cursor-based vs offset-based.

**Answer:**
GraphQL doesn't mandate a pagination approach but has common conventions.

**Offset/Limit Pagination (simple):**
```graphql
type Query {
  users(limit: Int!, offset: Int!): [User!]!
}
```
```javascript
const resolvers = {
  Query: {
    users: (_, { limit, offset }) =>
      db.query('SELECT * FROM users LIMIT $1 OFFSET $2', [limit, offset])
  }
};
```
- ✅ Simple to implement
- ❌ Unstable (new items shift pages), poor for large datasets

**Cursor-Based Pagination (Relay Connection Spec):**
```graphql
type UserConnection {
  edges: [UserEdge!]!
  pageInfo: PageInfo!
  totalCount: Int!
}

type UserEdge {
  node: User!
  cursor: String!    # opaque pointer to this item in the list
}

type PageInfo {
  hasNextPage: Boolean!
  hasPreviousPage: Boolean!
  startCursor: String
  endCursor: String
}

type Query {
  users(first: Int, after: String, last: Int, before: String): UserConnection!
}
```
```javascript
// Resolver implementation (simplified)
const resolvers = {
  Query: {
    users: async (_, { first = 10, after }) => {
      const cursor = after ? Buffer.from(after, 'base64').toString() : null;
      const users = await db.query(
        `SELECT * FROM users ${cursor ? 'WHERE id > $1' : ''} ORDER BY id LIMIT $2`,
        cursor ? [cursor, first + 1] : [first + 1]
      );

      const hasNextPage = users.length > first;
      const edges = users.slice(0, first).map(user => ({
        cursor: Buffer.from(user.id.toString()).toString('base64'),
        node: user
      }));

      return {
        edges,
        totalCount: db.count('users'),
        pageInfo: {
          hasNextPage,
          hasPreviousPage: !!after,
          startCursor: edges[0]?.cursor,
          endCursor: edges[edges.length - 1]?.cursor
        }
      };
    }
  }
};
```
- ✅ Stable pagination (works even with new inserts)
- ✅ Efficient for large datasets
- ❌ More complex to implement

---

### Q5: How do you handle file uploads in GraphQL?

**Answer:**
GraphQL itself uses JSON for transport, so file uploads require a special multipart request spec.

**Using `graphql-upload` (Apollo Server):**

**Schema:**
```graphql
scalar Upload

type Mutation {
  uploadAvatar(userId: ID!, file: Upload!): String!
  uploadFiles(files: [Upload!]!): [String!]!
}
```

**Setup in Apollo Server v3+:**
```javascript
const { graphqlUploadExpress } = require('graphql-upload');
const app = express();

// Add middleware BEFORE Apollo
app.use(graphqlUploadExpress({ maxFileSize: 10_000_000, maxFiles: 5 }));
```

**Resolver:**
```javascript
const resolvers = {
  Mutation: {
    uploadAvatar: async (parent, { userId, file }) => {
      const { createReadStream, filename, mimetype, encoding } = await file.promise;

      // Validate file type
      if (!mimetype.startsWith('image/')) {
        throw new Error('Only image files allowed');
      }

      // Save to disk or S3
      const filePath = `./uploads/${userId}-${filename}`;
      await new Promise((resolve, reject) => {
        createReadStream()
          .pipe(fs.createWriteStream(filePath))
          .on('finish', resolve)
          .on('error', reject);
      });

      // Or upload to S3
      // await s3.upload({ Bucket: 'my-bucket', Key: filename, Body: stream }).promise();

      return `https://cdn.example.com/${filename}`;
    }
  }
};
```

**Important:** Apollo Server v4 dropped built-in `graphql-upload` support. You need to add the middleware manually or use alternative approaches.

---

### Q6: What is the difference between `type` and `input` in GraphQL? Why can't you reuse types as inputs?

**Answer:**
This is one of the most common GraphQL schema design questions.

**Output Types (`type`):**
- Used as **return types** from queries and mutations
- Fields can have **resolvers** with arguments
- Can contain **interfaces**, **unions**, and **computed fields**

**Input Types (`input`):**
- Used as **arguments** passed to queries and mutations
- Fields cannot have arguments
- Cannot reference output types (Object types, interfaces, unions)
- All fields must be scalar or other input types

```graphql
# OUTPUT TYPE — returned from queries
type User {
  id: ID!
  name: String!
  email: String!
  fullName: String  @deprecated(reason: "Use name")
  posts: [Post!]!   # Can reference other types
}

# INPUT TYPE — used as argument
input CreateUserInput {
  name: String!
  email: String!
  role: UserRole!   # Can reference Enum
  # posts: [Post!]! ← ERROR! Input can't reference output types
}

type Mutation {
  createUser(data: CreateUserInput!): User!
  # createUser(user: User!): User! ← ERROR! Can't use output type as input
}
```

**Why the limitation?** Output types can have circular references and resolver logic. If used as inputs, it would create ambiguous validation logic. The separation is by design to ensure clear boundaries.

---

### Q7: How do you test GraphQL APIs?

**Answer:**

**1. Unit Testing Resolvers:**
Test individual resolver functions in isolation.
```javascript
const { createUser } = require('./resolvers/mutation');

describe('createUser resolver', () => {
  const mockContext = {
    db: {
      createUser: jest.fn()
    }
  };

  it('should create a user successfully', async () => {
    const user = { id: '1', name: 'John', email: 'john@test.com' };
    mockContext.db.createUser.mockResolvedValue(user);

    const result = await createUser(null, { name: 'John', email: 'john@test.com' }, mockContext);

    expect(result).toEqual(user);
    expect(mockContext.db.createUser).toHaveBeenCalledWith({ name: 'John', email: 'john@test.com' });
  });

  it('should throw for invalid email', async () => {
    await expect(
      createUser(null, { name: 'John', email: 'invalid' }, mockContext)
    ).rejects.toThrow('Invalid email');
  });
});
```

**2. Integration Testing with Apollo Server:**
Use `ApolloServer.executeOperation()` for in-memory testing without HTTP.
```javascript
const { ApolloServer } = require('@apollo/server');
const { typeDefs, resolvers } = require('./schema');

describe('GraphQL Integration', () => {
  let server;
  beforeAll(() => {
    server = new ApolloServer({ typeDefs, resolvers });
  });

  it('should query a user', async () => {
    const result = await server.executeOperation({
      query: `query GetUser($id: ID!) { user(id: $id) { name email } }`,
      variables: { id: '1' }
    });

    expect(result.body.singleResult.errors).toBeUndefined();
    expect(result.body.singleResult.data.user.name).toBe('John');
  });
});
```

**3. End-to-End Testing with `supertest`:**
```javascript
const request = require('supertest');
const app = require('./app');  // Express app with Apollo middleware

it('should create and query a user', async () => {
  const res = await request(app)
    .post('/graphql')
    .send({
      query: `mutation { createUser(name: "Test", email: "t@t.com") { id name } }`
    });

  expect(res.body.data.createUser.name).toBe('Test');
});
```

**Tools:** Jest, Vitest, Mocha, `@apollo/client/testing` (MockedProvider for React), Postman/Insomnia (manual), Bruno.

---

*Prepared for technical screening rounds at service-based companies (TCS, Infosys, Wipro, Capgemini, HCL, Cognizant).*
