# Design Uber Backend

## 1. Requirements

### Functional
*   **Driver**: Report location updates (every 5 seconds). Receive trip requests.
*   **Rider**: See nearby drivers. Request a ride. Match with a driver.
*   **ETA**: Calculate cost and time.

### Non-Functional
*   **Low Latency**: Location updates must be real-time.
*   **High Availability**: System should never be down.
*   **Consistency**: A driver cannot be matched to two riders at the same time.

## 2. Core Components

### A. Location Service (Driver Location)
*   **Challenge**: 100k active drivers sending updates every 5s = 20k writes/sec.
*   **Database**: Redis (Geospatial) or Cassandra (Write-heavy). NOT SQL.
*   **Algorithm**: **QuadTree** (or Google S2 Geometry).
    *   Divide the map into small squares.
    *   Driver belongs to one square.
    *   To find "Nearby Drivers", query the current square + 8 neighbors.

### B. Matching Service
*   **Task**: Match Rider R with Driver D.
*   **Logic**:
    1.  Rider requests ride.
    2.  Find K nearest drivers (from Location Service).
    3.  Filter (Exclude busy drivers).
    4.  Send request to Driver D1.
    5.  If D1 rejects/timeouts, send to D2.
*   **Concurrency**: Use Distributed Locking (Redis Lock) on the Driver ID so D1 doesn't get two requests.

### C. Trip Management
*   **State Machine**: `REQUESTED` -> `MATCHED` -> `STARTED` -> `COMPLETED` -> `CANCELLED`.
*   **Database**: SQL (Postgres/MySQL) because Trip data is structured and needs ACID (Billing).

## 3. Communication

### Driver <-> Server
*   **WebSockets**: Bi-directional. Server needs to push "New Ride Request" to Driver instantly. Polling is too slow.

### Server <-> Server
*   **gRPC**: Low latency internal communication (Location Service -> Matching Service).
*   **Kafka**: For Analytics, Fraud Detection, and Billing (Async processing).

## 4. Scale & Partitioning
*   **Sharding by City**: New York and Mumbai data don't overlap.
    *   *Shard Key*: `CityID`.
    *   *Geo-Sharding*: Use S2 Cell ID as key.

## 5. Interview Questions
1.  **Why QuadTree over Grid?**
    *   *Ans*: Grid is fixed size. QuadTree is adaptive (dense areas have smaller squares, rural areas have larger squares). Efficient for unequal distribution.
2.  **How to handle location drift (GPS error)?**
    *   *Ans*: Map Matching algorithms (Snap to road).
