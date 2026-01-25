# REST vs gRPC vs GraphQL

## 1. REST (Representational State Transfer)
The standard architecture for Web APIs.
*   **Format**: JSON (usually).
*   **Transport**: HTTP/1.1 (Text based).
*   **Resource Based**: URLs represent resources (`/users/1`).
*   **Pros**: Simple, Browser compatible, easy to debug (curl), heavy ecosystem.
*   **Cons**:
    *   **Over-fetching**: You want just user name, but get full user object.
    *   **Under-fetching**: You need user + posts. Requires 2 HTTP calls (N+1 problem).

## 2. gRPC (Google Remote Procedure Call)
High performance framework.
*   **Format**: Protocol Buffers (Binary).
*   **Transport**: HTTP/2 (Multiplexing, Streaming).
*   **Action Based**: Methods (`GetUser()`), not Resources.
*   **Pros**:
    *   **Fast**: Smaller payload (binary), faster parsing.
    *   **Contract First**: Strict `.proto` definition prevents breakage.
    *   **Streaming**: Supports Bi-directional streaming.
*   **Cons**: Not browser compatible (needs gRPC-Web proxy). Hard to debug (binary blob).

## 3. GraphQL (Graph Query Language)
A query language for APIs (Facebook).
*   **Format**: JSON.
*   **Transport**: HTTP (Single endpoint usually `POST /graphql`).
*   **Client Driven**: Client asks for exactly what it needs.
    *   `query { user(id: 1) { name, posts { title } } }`
*   **Pros**: Solves Over-fetching and Under-fetching. Great for Mobile/Frontend.
*   **Cons**:
    *   **Complexity**: Backend logic is complex (Resolvers).
    *   **Caching**: HTTP Caching is hard (everything is POST).

## 4. Comparison Table

| Feature | REST | gRPC | GraphQL |
| :--- | :--- | :--- | :--- |
| **Protocol** | HTTP/1.1 | HTTP/2 | HTTP |
| **Data Format** | JSON (Text) | Protobuf (Binary) | JSON (Text) |
| **Parsing** | Slow (Text) | Fast (Binary) | Medium |
| **Browser Support** | Native | Poor | Native |
| **Use Case** | Public APIs, CRUD | Microservices (Internal) | Frontend/Mobile |

## 5. Interview Questions
1.  **Why use gRPC for Microservices?**
    *   *Ans*: Microservices talk to each other thousands of times. JSON parsing overhead adds up. gRPC (Binary) is 10x faster and strictly typed, preventing integration bugs.
2.  **How does GraphQL solve N+1 problem?**
    *   *Ans*: It doesn't solve it automatically (backend might still fire 100 queries). You need to use **DataLoader** pattern (Batching) to combine requests on the server side.
