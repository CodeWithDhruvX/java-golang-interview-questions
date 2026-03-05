# High-Level Design (HLD): Resiliency and Fault Tolerance

In distributed systems, failures are guaranteed. A third-party API will time out, a database will restart, or a network link will drop packets. Resiliency is the ability of the system to gracefully handle these failures without crashing completely.

## 1. What is the Circuit Breaker Pattern?
**Answer:**
Imagine Service A depends heavily on a legacy, slow Service B. Suddenly, Service B starts timing out (taking 30 seconds to fail). If Service A receives 10,000 requests, all 10,000 threads in Service A will be blocked waiting for Service B, causing Service A to crash entirely.
*   **The Circuit Breaker:**
    *   **Closed State:** Normal operation. Requests flow through. The breaker monitors responses.
    *   **Open State:** If the failure rate (e.g., 50% errors/timeouts in the last 10 seconds) breaches a threshold, the circuit "trips". Now, any request sent to Service B immediately fails fast (returning an error or fallback data) *without actually making the network call*. This protects Service B from being hammered while it's down and prevents Service A's threads from blocking.
    *   **Half-Open State:** After a timeout (e.g., 30 seconds), the breaker allows a *single* test request through. If it succeeds, the circuit closes. If it fails, it remains Open.
    *   *(Popular Implementation: Resilience4j in Java).*

## 2. Why are Retries with Exponential Backoff and Jitter important?
**Answer:**
When a network call fails, simply retrying immediately in a tight loop is a bad idea. If the destination server is already overloaded, 10,000 clients retrying instantly will cause a Denial of Service (DoS) attack against your own services.
*   **Exponential Backoff:** The client waits 1 second, then retries. Fails. Waits 2s. Fails. Waits 4s, 8s, 16s. This gives the overloaded server time to recover.
*   **Jitter:** If 1,000 clients all start their exponential backoff at the exact same millisecond, they will all retry simultaneously at t=1, t=2, t=4. This still creates giant spikes of traffic. "Jitter" introduces a random variable (e.g., +/- 20%) to the wait time. So Client 1 might wait 1.1s, 2.3s, 3.8s, spreading the retry load smoothly.

## 3. What is the Bulkhead Pattern?
**Answer:**
Named after the watertight compartments in a ship. If a ship's hull is breached, only one compartment floods, preventing the entire ship from sinking.
*   **Software Application:** If your API Gateway route `/generate-pdf` consumes massive CPU and suddenly gets spammed, it will consume 100% of the server's thread pool. Now, simple requests to `/health-check` or `/get-user` will timeout because there are no threads left.
*   **Implementation:** Isolate resources. Give `/generate-pdf` a dedicated, isolated thread pool of a maximum of 10 threads, and give the rest of the application 100 threads. If `/generate-pdf` gets overwhelmed, only those 10 threads queue up. The other 100 threads remain freely available to process the rest of the system's traffic.

## 4. Rate Limiting vs Load Shedding
**Answer:**
*   **Rate Limiting:** Protects the system from a *specific greedy user* or tenant. "User John can only make 10 requests per second." Implemented at the API Gateway level (often backed by Redis).
*   **Load Shedding:** Protects the system from *overall global overload*, regardless of who the users are. If the main database CPU hits 95%, the system deliberately starts rejecting 20% of ALL incoming requests (returning HTTP 503 Service Unavailable) so the system can survive and process the remaining 80%. It's a survival mechanism when scaling up is too slow.
