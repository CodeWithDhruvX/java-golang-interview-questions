# System Design (HLD) - Ride Sharing App (Uber/Lyft)

## Problem Statement
Design a ride-sharing service where passengers can book cars and drivers are matched with them based on proximity.

## 1. Requirements Clarification
### Functional Requirements
*   **Driver Matching:** Passengers should see nearby drivers.
*   **Ride Request:** Passengers can request a ride.
*   **Driver Notification:** Nearby drivers receive a notification and can accept/decline the request.
*   **Location Tracking:** Both driver and passenger locations are tracked in real-time.

### Non-Functional Requirements
*   **Low Latency Matching:** The system must quickly match a driver and passenger.
*   **High Scalability:** Should support millions of concurrent users globally, particularly during peak hours or events.
*   **Consistency vs Availability:** The matching system must firmly avoid double-booking a single driver for two divergent passengers (Consistency is crucial during the transaction).

## 2. High-Level Design (Architecture)

```text
[Passenger App] <--(WebSockets)--> [ API Gateway / Load Balancer ] <--(WebSockets)--> [Driver App]
                                             |
                                +------------+------------+
                                |                         |
                   [ Driver Location Service ]     [ Matching Service ]
                                |                         |
                      [ Redis / GeoHash DB ]       [ Ride Transactions DB ]
```

## 3. Core Component Design: Driver Location Service
The primary challenge is querying the locations of moving objects efficiently.
*   **Naive Approach:** Store `(driver_id, lat, long)` in a SQL database. To find nearby drivers, we calculate the Euclidean distance for every driver. *This is $O(N)$ and scales terribly.*
*   **Better Approach (GeoHashing):** We divide the world into a grid. Each cell in the grid has a unique string ID (e.g., `9q8yy`). If two points share the same prefix (e.g., `9q8yyA` and `9q8yyZ`), they are in the same grid.
*   **Implementation:** Use a specialized spatial database like **Redis Geospatial** or **Elasticsearch**. 
    *   Drivers send location updates every 3-5 seconds to the WebSockets.
    *   The Location Service updates the `(DriverID, Lat, Long)` in Redis.
    *   When a passenger opens the app, the backend queries Redis: "Give me all drivers within a 3km radius of this coordinate." Redis uses GeoHashing to return this instantly.

## 4. The Matching Flow
1. Passenger selects a destination and taps "Book".
2. The Request hits the **Matching Service**.
3. Matching Service queries the **Location Service** to get the top 5 nearest Available drivers.
4. Matching Service pushes a notification to Driver #1 (via WebSockets or FCM/APNs).
5. If Driver #1 declines or times out (e.g., 10 seconds), the Matching Service pushes the request to Driver #2.
6. When Driver #2 accepts, the database is updated, and the Passenger receives confirmation.

## 5. Dealing with High Read/Write Volume
*   Updating millions of driver locations every 5 seconds creates immense write pressure.
*   We cannot perform this in a standard SQL database. We must use an in-memory datastore (Redis) paired with a Write-Ahead Log or period snapshots if persistence is strictly required. Often, precise historical tracking of drivers is sent asynchronously via Kafka to a cheaper Storage layer (like Cassandra/S3) for analytics/billing, keeping the "live" grid small and fast.

## 6. Follow-up Questions for Candidate
1.  How do you calculate ETAs and routes? (Usually delegated to third-party routing engines like Google Maps APIs/OSRM, but requires caching common routes).
2.  How do you scale Geohashing globally? (Sharding by regions: North America servers handle NA drivers, Europe handles EU drivers).
3.  How do you handle the race condition where two passengers simultaneously get assigned the same driver? (Database transactions or Redis atomic SETNX locks).
