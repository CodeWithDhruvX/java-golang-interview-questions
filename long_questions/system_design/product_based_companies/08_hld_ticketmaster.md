# System Design (HLD) - Ticketmaster (High Concurrency Booking)

## Problem Statement
Design a ticketing system like Ticketmaster or BookMyShow. Users can search for events and book tickets.

## 1. Requirements Clarification
### Functional Requirements
*   **Search/Browse:** Users can search for events (concerts, sports, movies).
*   **Booking:** Users can select seats and book them.
*   **Payment:** Users pay for the selected tickets within a short timeframe.

### Non-Functional Requirements
*   **Extremely High Concurrency:** A popular concert (e.g., Taylor Swift) goes on sale, and 1 million people simultaneously try to buy 50,000 tickets at the exact same second.
*   **Strong Consistency:** Absolutely no double booking of the same seat.
*   **Fairness/Availability:** The system shouldn't crash under load.

## 2. High-Level Architecture
This system is defined by its ability to handle sudden, massive spikes (Read-to-Write ratio goes from 100:1 to 1:1 in seconds).

```text
[ Users ] ---> [ CDN (Static assets) ] ---> [ API Gateway ]
                                                |
                               +----------------+----------------+
                               |                                 |
                     [ Search/Browse Service ]          [ Booking Service ]
                               |                                 |
                      [ Elasticsearch ]           +--------------+--------------+
                                                  |              |              |
                                          [ Virtual Queue ]  [ Redis ]   [ PostgreSQL DB ]
```

## 3. The "Taylor Swift" Problem: Virtual Waiting Room
If 1 million users hit the PostgreSQL database at once, the database will die.
*   **Solution: The Virtual Queue API Layer.**
*   Before a user can reach the "Select Seats" page, the API Server places them in a distributed Queue (Kafka or Redis Lists).
*   Users see a "You are in line, your estimated wait time is X minutes" screen.
*   The Booking system only pulls users from the Queue at a rate the PostgreSQL DB can comfortably handle (e.g., 5,000 users/minute).
*   Once 50,000 tickets are gone, the system instantly empties the queue and tells the remaining 950,000 users "Sold Out" without ever letting them touch the main database.

## 4. Seat Locking and Concurrency Control
Once allowed in, User A selects Seat 1A. User B selects Seat 1A a millisecond later.
*   **Redis Reservation (Soft Lock):**
    *   When User A clicks Seat 1A, the backend puts a key in Redis: `SETNX lock:event123:seat1A "User_A_ID" EX 300`. (Set if Not Exists, expires in 5 minutes).
    *   If successful, User A has 5 minutes to enter their credit card.
    *   If User B requests Seat 1A, `SETNX` fails. User B is told "Seat unavailable".
*   **Database Confirmation (Hard Lock / ACID):**
    *   User A successfully pays.
    *   The system now executes a database transaction:
    *   `BEGIN; SELECT * FROM tickets WHERE seat_id = '1A' FOR UPDATE; UPDATE tickets SET status = 'BOOKED', user_id = 'A' WHERE seat_id = '1A'; COMMIT;`
    *   Pessimistic locking ensures that even if Redis completely failed and two users reached the DB, only one transaction commits.

## 5. Follow-up Questions for Candidate
1.  How do you handle a user holding 4 seats in Redis but their internet drops and they never pay? (The Redis TTL expires after 5 minutes, automatically freeing the seats for the next user).
2.  How do you ensure search is fast when new concerts are added daily? (Use Elasticsearch, sync it asynchronously via Kafka when a concert admin adds an event to the main SQL DB).
