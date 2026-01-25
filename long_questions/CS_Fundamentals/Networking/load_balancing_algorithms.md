# Load Balancing Algorithms

## 1. Static Algorithms
These algorithms follow a fixed set of rules and do not consider the current state of the servers.

### Round Robin
*   **Mechanism**: Requests are distributed sequentially to the list of servers.
    *   Req 1 -> Server A
    *   Req 2 -> Server B
    *   Req 3 -> Server A (repeat)
*   **Pros**: Simple, stateless.
*   **Cons**: Assumes all servers have equal capacity.

### Weighted Round Robin
*   **Mechanism**: Same as RR, but servers are assigned a "weight" based on their capacity.
    *   Server A (Weight 2), Server B (Weight 1).
    *   Req 1 -> A, Req 2 -> A, Req 3 -> B.
*   **Pros**: Handles servers with unequal specs.

### IP Hash
*   **Mechanism**: Hash(Client IP) % NumberOfServers.
*   **Pros**: Sticky Session (requests from same user go to same server).
*   **Cons**: Uneven distribution if user pool is behind a proxy.

## 2. Dynamic Algorithms
These algorithms check the server's current health and load before routing.

### Least Connections
*   **Mechanism**: Route traffic to the server with the fewest active connections.
*   **Pros**: Excellent for long-lived connections (e.g., WebSocket).
*   **Cons**: Connections count doesn't always reflect actual load (CPU/RAM).

### Least Response Time
*   **Mechanism**: Route to the server with the lowest latency and fewest connections.
*   **Pros**: Most responsive user experience.
*   **Cons**: Complex to calculate.

### Resource-Based (Adaptive)
*   **Mechanism**: Agent on the server reports actual CPU/Memory usage to LB. LB routes to least loaded.
*   **Pros**: Intelligent.
*   **Cons**: Requires installing agents on backend.

## 3. L4 vs L7 Load Balancing

### L4 (Transport Layer)
*   **Data**: Inspects IP and TCP Port. (Does NOT see HTTP headers).
*   **Speed**: Extremely fast (packet forwarding).
*   **Example**: TCP Load Balancing.

### L7 (Application Layer)
*   **Data**: Inspects HTTP Headers, Cookies, URL path.
*   **Flexibility**: Can route `/api/*` to Service A and `/static/*` to Service B.
*   **Example**: NGINX, AWS ALB.
*   **Cost**: Slower (needs to decrypt HTTPS, inspect, re-encrypt).

## 4. Consistent Hashing
Used in Distributed Systems (like Cassandra, DynamoDB) to distribute data/requests.
*   **Problem with Modulo (Hash % N)**: If you add/remove a server, N changes, and nearly all keys get remapped. Massive cache miss storm.
*   **Consistent Hashing Solution**:
    *   Servers and Keys are mapped to a "Ring" (0 to 360 degrees).
    *   Key K is stored on the first Server found moving clockwise on the ring.
    *   **Result**: Adding/Removing a server only affects keys in its immediate vicinity.

## 5. Interview Questions
1.  **When to use Least Connections vs Round Robin?**
    *   *Ans*: Use Round Robin for short, stateless HTTP requests where processing time is uniform. Use Least Connections for long-lived sessions (WebSockets, Database connections) where one connection might be active for hours.
2.  **How does a Load Balancer handle a server failure?**
    *   *Ans*: Health Checks (Heartbeats). The LB periodically pings an endpoint (e.g., `/health`). If it fails N times, the server is marked "Unhealthy" and removed from rotation.
