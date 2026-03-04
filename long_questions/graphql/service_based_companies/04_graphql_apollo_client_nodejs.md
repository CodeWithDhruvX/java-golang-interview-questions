# 🟣 GraphQL — Apollo Client (React) & Node.js Setup (Service-Based Companies)

This document covers Apollo Client integration with React and setting up a GraphQL server with Node.js/Express — commonly tested in service-based company frontend/full-stack roles (2–5 years experience).

---

### Q1: How do you set up Apollo Client in a React application?

**Answer:**
Apollo Client is the most popular GraphQL client for React. It manages data fetching, caching, and state in a single library.

**1. Install dependencies:**
```bash
npm install @apollo/client graphql
```

**2. Configure and wrap your app:**
```javascript
// src/apolloClient.js
import { ApolloClient, InMemoryCache, createHttpLink, ApolloProvider } from '@apollo/client';
import { setContext } from '@apollo/client/link/context';

// HTTP link — points to GraphQL endpoint
const httpLink = createHttpLink({
  uri: 'http://localhost:4000/graphql'
});

// Auth link — adds JWT token to every request header
const authLink = setContext((_, { headers }) => {
  const token = localStorage.getItem('authToken');
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : ''
    }
  };
});

// Create the client
const client = new ApolloClient({
  link: authLink.concat(httpLink),   // Chain: auth → http
  cache: new InMemoryCache()         // Normalized in-memory cache
});

// src/index.jsx — Wrap the app
function App() {
  return (
    <ApolloProvider client={client}>
      <Router>
        <Routes>...</Routes>
      </Router>
    </ApolloProvider>
  );
}
```

---

### Q2: Explain `useQuery`, `useMutation`, and `useSubscription` hooks.

**Answer:**
These hooks are the primary way to interact with GraphQL in React components.

**`useQuery` — Fetching data:**
```javascript
import { useQuery, gql } from '@apollo/client';

const GET_USERS = gql`
  query GetUsers($role: String) {
    users(role: $role) {
      id
      name
      email
      avatar { url }
    }
  }
`;

function UserList() {
  const { loading, error, data, refetch, fetchMore } = useQuery(GET_USERS, {
    variables: { role: 'USER' },
    fetchPolicy: 'cache-and-network',   // Return cache immediately, update from network
    pollInterval: 30000,                // Refetch every 30 seconds
    skip: !isLoggedIn,                  // Skip query if not logged in
    onCompleted: (data) => console.log('Done!', data),
    onError: (err) => console.error('Error:', err)
  });

  if (loading) return <Spinner />;
  if (error) return <ErrorMessage error={error} />;

  return (
    <ul>
      {data.users.map(user => (
        <li key={user.id}>{user.name} — {user.email}</li>
      ))}
    </ul>
  );
}
```

**`useMutation` — Modifying data:**
```javascript
import { useMutation, gql } from '@apollo/client';

const CREATE_USER = gql`
  mutation CreateUser($name: String!, $email: String!) {
    createUser(name: $name, email: $email) {
      id
      name
      email
    }
  }
`;

function CreateUserForm() {
  const [createUser, { loading, error, data }] = useMutation(CREATE_USER, {
    // Automatically update the cache after mutation
    update(cache, { data: { createUser } }) {
      cache.modify({
        fields: {
          users(existingUsers = []) {
            const newUserRef = cache.writeFragment({
              data: createUser,
              fragment: gql`fragment NewUser on User { id name email }`
            });
            return [...existingUsers, newUserRef];
          }
        }
      });
    },
    onCompleted: (data) => console.log('Created:', data.createUser.id),
    onError: (err) => console.error('Failed:', err)
  });

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await createUser({
        variables: { name: 'John', email: 'john@test.com' }
      });
    } catch (err) {
      // Errors also available in the `error` destructured above
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <button type="submit" disabled={loading}>
        {loading ? 'Creating...' : 'Create User'}
      </button>
      {error && <p>Error: {error.message}</p>}
    </form>
  );
}
```

**`useSubscription` — Real-time data:**
```javascript
import { useSubscription, gql } from '@apollo/client';

const ON_MESSAGE = gql`
  subscription OnMessage($channelId: ID!) {
    messageSent(channelId: $channelId) {
      id
      text
      sender { name }
      createdAt
    }
  }
`;

function ChatMessages({ channelId }) {
  const [messages, setMessages] = useState([]);

  useSubscription(ON_MESSAGE, {
    variables: { channelId },
    onData: ({ data: { data } }) => {
      setMessages(prev => [...prev, data.messageSent]);
    }
  });

  return <ul>{messages.map(m => <li key={m.id}>{m.text}</li>)}</ul>;
}
```

---

### Q3: What are Apollo Client fetch policies and when do you use them?

**Answer:**
Fetch policies control how Apollo Client uses the cache vs. the network.

| Policy | Cache? | Network? | Use Case |
|---|---|---|---|
| `cache-first` (default) | ✅ Use if available | Only if not in cache | Read-heavy data that changes rarely |
| `cache-and-network` | ✅ Return immediately | Always fetch & update | Pages where you want fast + fresh data |
| `network-only` | ❌ Always skip | Always fetch | Critical real-time data (cart, balance) |
| `cache-only` | ✅ Only cache | Never | Offline mode, pre-seeded data |
| `no-cache` | ❌ Skip | Always fetch, don't store | Sensitive queries (OTP, one-time codes) |
| `standby` | ✅ Use if available | No, but triggers on refetch | Background/secondary queries |

```javascript
// Cache-first: good for product catalog (changes rarely)
useQuery(GET_PRODUCTS, { fetchPolicy: 'cache-first' });

// Network-only: good for order status (must be fresh)
useQuery(GET_ORDER_STATUS, { fetchPolicy: 'network-only' });

// Cache-and-network: good for dashboard (fast load + auto-refresh)
useQuery(GET_DASHBOARD, { fetchPolicy: 'cache-and-network' });
```

---

### Q4: How do you handle Apollo Client cache updates after mutations?

**Answer:**
After a mutation, Apollo Client needs to know how to update the cache. There are several approaches:

**1. `refetchQueries` — Re-run affected queries (simple):**
```javascript
const [deleteUser] = useMutation(DELETE_USER, {
  refetchQueries: [
    { query: GET_USERS },             // Re-run the users list query
    'GetUserCount'                    // Or by operation name
  ],
  awaitRefetchQueries: true           // Wait before calling onCompleted
});
```
- ✅ Simple — always correct
- ❌ Makes extra network request

**2. `cache.modify` — Manual cache update (efficient):**
```javascript
const [addPost] = useMutation(ADD_POST, {
  update(cache, { data: { addPost } }) {
    // Directly write to the cache — no extra network call
    cache.modify({
      fields: {
        posts(existingPosts = []) {
          const newPostRef = cache.writeFragment({
            data: addPost,
            fragment: gql`fragment NewPost on Post { id title createdAt }`
          });
          return [newPostRef, ...existingPosts];  // Prepend new post
        }
      }
    });
  }
});
```

**3. `cache.evict` — Remove stale entries:**
```javascript
const [deletePost] = useMutation(DELETE_POST, {
  update(cache, { data: { deletePost } }) {
    if (deletePost.success) {
      // Remove from cache by ID
      cache.evict({ id: cache.identify({ __typename: 'Post', id: postId }) });
      cache.gc();  // Garbage collect dangling references
    }
  }
});
```

---

### Q5: How do you set up a GraphQL server with Node.js and Express (Apollo Server)?

**Answer:**

**1. Install:**
```bash
npm install @apollo/server graphql express @as-integrations/express
```

**2. Basic setup:**
```javascript
// server.js
const express = require('express');
const { ApolloServer } = require('@apollo/server');
const { expressMiddleware } = require('@as-integrations/express');
const cors = require('cors');
const bodyParser = require('body-parser');

const typeDefs = `
  type User {
    id: ID!
    name: String!
    email: String!
  }

  type Query {
    users: [User!]!
    user(id: ID!): User
  }

  type Mutation {
    createUser(name: String!, email: String!): User!
  }
`;

// In-memory store for demo
let users = [
  { id: '1', name: 'Alice', email: 'alice@test.com' },
  { id: '2', name: 'Bob', email: 'bob@test.com' }
];

const resolvers = {
  Query: {
    users: () => users,
    user: (_, { id }) => users.find(u => u.id === id)
  },
  Mutation: {
    createUser: (_, { name, email }) => {
      const user = { id: Date.now().toString(), name, email };
      users.push(user);
      return user;
    }
  }
};

async function startServer() {
  const app = express();
  const server = new ApolloServer({ typeDefs, resolvers });

  await server.start();

  app.use(
    '/graphql',
    cors(),
    bodyParser.json(),
    expressMiddleware(server, {
      context: async ({ req }) => {
        // Add auth, DB, etc. to context
        const token = req.headers.authorization?.replace('Bearer ', '');
        const user = token ? await getUserFromToken(token) : null;
        return { user, db };
      }
    })
  );

  app.listen(4000, () => {
    console.log('🚀 GraphQL server ready at http://localhost:4000/graphql');
  });
}

startServer();
```

**3. Test with cURL:**
```bash
# Query
curl -X POST http://localhost:4000/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ users { id name email } }"}'

# Mutation
curl -X POST http://localhost:4000/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "mutation { createUser(name: \"Charlie\", email: \"c@c.com\") { id name } }"}'
```

---

### Q6: How do you set up Apollo Client for WebSocket subscriptions?

**Answer:**
Subscriptions require a WebSocket link in addition to the HTTP link.

```bash
npm install graphql-ws @apollo/client
```

```javascript
import {
  ApolloClient,
  InMemoryCache,
  createHttpLink,
  split
} from '@apollo/client';
import { GraphQLWsLink } from '@apollo/client/link/subscriptions';
import { createClient } from 'graphql-ws';
import { getMainDefinition } from '@apollo/client/utilities';

// HTTP link for queries and mutations
const httpLink = createHttpLink({ uri: 'http://localhost:4000/graphql' });

// WebSocket link for subscriptions
const wsLink = new GraphQLWsLink(
  createClient({
    url: 'ws://localhost:4000/graphql',
    connectionParams: {
      authToken: localStorage.getItem('authToken')
    }
  })
);

// Route: subscriptions → WebSocket, everything else → HTTP
const splitLink = split(
  ({ query }) => {
    const definition = getMainDefinition(query);
    return (
      definition.kind === 'OperationDefinition' &&
      definition.operation === 'subscription'
    );
  },
  wsLink,     // If subscription
  httpLink    // If query or mutation
);

const client = new ApolloClient({
  link: splitLink,
  cache: new InMemoryCache()
});
```

---

*Prepared for frontend/full-stack technical rounds at service-based companies (TCS, Infosys, Wipro, HCL, Cognizant).*
