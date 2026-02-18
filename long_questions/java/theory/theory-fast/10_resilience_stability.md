# Exception Handling & Resilience Interview Questions (118-125)

## Stability Patterns & Error Handling

### 118. Difference between checked and unchecked exceptions in real systems?
"**Checked exceptions** (like `IOException`) force you to handle them at compile time. They are meant for recoverable errors. If a file is missing, maybe prompt the user for a new path.

**Unchecked exceptions** (like `NullPointerException`, `IllegalArgumentException`) indicate programmer errors or unrecoverable system failures. You can't really 'recover' from a null pointer in 99% of cases.

In modern backend development (Spring), we almost exclusively use **Unchecked Exceptions**. Using checked exceptions leads to cluttered code with empty `catch` blocks or method signatures that declare `throws Exception` all the way up the stack. We wrap checked exceptions in RuntimeExceptions (like `DataAccessException`) and handle them globally."

### 119. Why should you avoid catching `Exception`?
"Catching the generic `Exception` is lazy and dangerous.

It swallows *everything*—including `RuntimeException` bugs that *should* crash the app so you notice them.

If I write `try { ... } catch (Exception e)`, I might accidentally catch a `NullPointerException` and just log 'Something went wrong', leaving the system in an inconsistent state.

You should always catch specific exceptions (`catch (IOException e)`) so you know exactly what you are handling and how to recover from it."

### 120. What is a circuit breaker pattern?
"It prevents cascading failures.

If Service A calls Service B, and Service B is down or slow, Service A's threads will block waiting for a response. Eventually, Service A runs out of threads and crashes too.

A **Circuit Breaker** (like Resilience4j) monitors the calls. If 50% of calls to Service B fail, the circuit 'opens'.

Now, Service A *immediately* fails calls to Service B without waiting, returning a fallback response or error. This gives Service B time to recover. After a while, it allows a few test calls through ('half-open') to see if B is back online."

### 121. What is retry vs timeout?
"**Timeout** is saying 'If you don't answer in 2 seconds, I’m hanging up.' It prevents one slow service from blocking the entire chain.

**Retry** is saying 'You didn't answer? Let me ask again.' It handles transient glitches (like a network blip).

But you have to be careful: **Retries can cause a Retry Storm**. If Service B is overloaded, retrying just adds *more* load, killing it faster. I always use **Exponential Backoff** (wait 1s, then 2s, then 4s) and Jitter (randomness) to prevent this."

### 122. What is bulkhead pattern?
"It’s inspired by ships. If a ship gets a hole, you close the bulkheads so only one section floods, not the whole ship.

In code, it means isolating resources. I might have separate thread pools for different downstream services.

If the 'Image Processing Service' is slow and uses up all its threads, the 'User Login Service' (which uses a different thread pool) is unaffected. Without bulkheads, one slow dependency can starve the entire application."

### 123. How do you design fault-tolerant services?
"I assume everything will fail eventually.

1.  **Redundancy**: Run multiple instances of every service.
2.  **Statelessness**: So any instance can handle any request.
3.  **Timeouts/Retries**: Fail fast on slow dependencies.
4.  **Circuit Breakers**: Stop calling dead services.
5.  **Fallbacks**: If the Recommendation Engine is down, show 'Popular Items' (cached) instead of an error page. Graceful degradation."

### 124. What is graceful shutdown?
"When a service receives a SIGTERM signal (like when Kubernetes checks a pod), it shouldn't just `kill -9` and drop active connections.

Graceful shutdown means:
1.  Stop accepting *new* requests.
2.  Allow *existing* requests a grace period (e.g., 30s) to finish processing.
3.  Close database connections and file handles cleanly.
4.  Then exit.

Spring Boot supports this out of the box with `server.shutdown=graceful`."

### 125. How do you handle partial failures in microservices?
"This is tricky. If I succeed in 'Payment' but fail in 'Inventory', I have a data inconsistency.

We use **Sagas** or **Eventual Consistency**.

If a step fails, I trigger a 'Compensating Transaction' (an undo operation). If I charged the card but couldn't reserve the item, I issue a refund.

Ideally, I design systems to be idempotent and retryable, so I can just fix the underlying issue and replay the failed message."
