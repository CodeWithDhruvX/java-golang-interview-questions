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

**Spoken Format:**
"Exception handling is like having different types of first aid kits for different problems.

**Checked exceptions** are like having a bandage kit - they're for expected problems that you can plan for and recover from. If you try to open a file that doesn't exist, you catch `IOException` and can prompt user for correct path.

**Unchecked exceptions** are like having a fire extinguisher - they're for unexpected emergencies that indicate serious system problems. A `NullPointerException` means your code has a fundamental flaw that needs immediate fixing.

The problem with catching generic `Exception` is that it's like using one giant first aid kit for everything. You might put a bandage on a fire, but you won't know if it's a cut, burn, or something more serious.

Always catch specific exceptions so you know exactly what problem you're solving and can provide the right treatment!"

### 119. Why should you avoid catching `Exception`?
"Catching the generic `Exception` is lazy and dangerous.

It swallows *everything*—including `RuntimeException` bugs that *should* crash the app so you notice them.

If I write `try { ... } catch (Exception e)`, I might accidentally catch a `NullPointerException` and just log 'Something went wrong', leaving the system in an inconsistent state.

You should always catch specific exceptions (`catch (IOException e)`) so you know exactly what you are handling and how to recover from it."

**Spoken Format:**
"Catching generic exceptions is dangerous because it hides real problems.

If you write `try { ... } catch (Exception e)`, you might catch a `NullPointerException` and just log 'Something went wrong'. The real issue - why was there a null pointer in the first place? - gets hidden.

Or worse, you might catch `RuntimeException` which could be anything - a network timeout, database error, or logic bug. You'll log it but won't know the root cause.

Instead, catch specific exceptions like `IOException`, `SQLException`, `TimeoutException`. These tell you exactly what went wrong so you can fix the actual problem.

Generic exception handling is like having a first aid kit that only says 'treated for injury' - it doesn't help you prevent injuries in the first place!"

### 120. What is a circuit breaker pattern?
"It prevents cascading failures.

If Service A calls Service B, and Service B is down or slow, Service A's threads will block waiting for a response. Eventually, Service A runs out of threads and crashes too.

A **Circuit Breaker** (like Resilience4j) monitors the calls. If 50% of calls to Service B fail, the circuit 'opens'.

Now, Service A *immediately* fails calls to Service B without waiting, returning a fallback response or error. This gives Service B time to recover. After a while, it allows a few test calls through ('half-open') to see if B is back online."

**Spoken Format:**
"A circuit breaker is like a smart electrical switch that prevents power surges.

Imagine Service A calls Service B repeatedly. If Service B starts failing or responding slowly, Service A's threads start piling up waiting for responses.

The circuit breaker monitors this situation - like watching the current draw. If too many requests fail (high current), it trips the circuit.

When tripped:
- Service A immediately stops calling Service B
- Instead, it returns a fallback response (like cached data or default values)
- This prevents Service A from crashing due to waiting threads

After some time, the circuit breaker allows a few test calls through to see if Service B has recovered

If Service B is working again, the circuit closes and normal operation resumes.

It's like having an automatic circuit breaker that protects your house from electrical overloads - it sacrifices individual calls to save the entire system!"

### 121. What is retry vs timeout?
"**Timeout** is saying 'If you don't answer in 2 seconds, I’m hanging up.' It prevents one slow service from blocking the entire chain.

**Retry** is saying 'You didn't answer? Let me ask again.' It handles transient glitches (like a network blip).

But you have to be careful: **Retries can cause a Retry Storm**. If Service B is overloaded, retrying just adds *more* load, killing it faster. I always use **Exponential Backoff** (wait 1s, then 2s, then 4s) and Jitter (randomness) to prevent this."

**Spoken Format:**
"Retries are like having a persistent friend who keeps asking, but sometimes too insistently.

If a service fails, retrying immediately is like your friend knocking again right away - this can overwhelm the struggling service and make things worse.

**Exponential backoff** is like being smarter about when to ask again:
- First retry: wait 1 second
- Second retry: wait 2 seconds  
- Third retry: wait 4 seconds
- Each time you wait longer, giving the service time to recover

**Jitter** is like adding some randomness - instead of waiting exactly 2 seconds, maybe wait 2.3 seconds to avoid multiple services retrying at exactly the same time.

**Timeouts** are like saying 'I'll ask again, but if you don't answer in 5 seconds, I'll give up.'

The combination protects the struggling service while still being persistent enough to eventually succeed!"

### 122. What is bulkhead pattern?
"It’s inspired by ships. If a ship gets a hole, you close the bulkheads so only one section floods, not the whole ship.

In code, it means isolating resources. I might have separate thread pools for different downstream services.

If the 'Image Processing Service' is slow and uses up all its threads, the 'User Login Service' (which uses a different thread pool) is unaffected. Without bulkheads, one slow dependency can starve the entire application."

**Spoken Format:**
"Bulkheads are like having separate emergency exits for different parts of a building.

Imagine your application is like a building with multiple departments:
- Image processing department
- User authentication department  
- Database operations department
- Email sending department

If the image processing department has a problem (like a fire), you don't want to evacuate the entire building. You just close that department's emergency exit.

This is what bulkheads do - they isolate problems so one failing department doesn't bring down the whole building.

In code terms, each critical service gets its own thread pool and connection limits. If one service is having issues, it doesn't affect others that have their own resources.

It's like having separate circuit breakers for each department - one department's fire alarm doesn't trigger evacuation for the entire building!"

### 123. How do you design fault-tolerant services?
"I assume everything will fail eventually.

1.  **Redundancy**: Run multiple instances of every service.
2.  **Statelessness**: So any instance can handle any request.
3.  **Timeouts/Retries**: Fail fast on slow dependencies.
4.  **Circuit Breakers**: Stop calling dead services.
5.  **Fallbacks**: If the Recommendation Engine is down, show 'Popular Items' (cached) instead of an error page. Graceful degradation."

**Spoken Format:**
"Designing fault-tolerant services is like building a robust house.

You want to make sure it can withstand different types of failures, like a storm or earthquake.

**Redundancy** is like having multiple pillars holding up the house. If one pillar fails, the others can still support the structure.

**Statelessness** is like having a modular design. If one module fails, you can easily replace it without affecting the rest of the house.

**Timeouts/Retries** are like having a smart electrical system. If a circuit is overloaded, it automatically switches to a backup circuit.

**Circuit Breakers** are like having automatic shut-off valves. If a pipe bursts, the valve closes to prevent further damage.

**Fallbacks** are like having a backup generator. If the main power source fails, the generator kicks in to keep the lights on.

By combining these strategies, you can build a robust and resilient system that can withstand failures and keep running smoothly!"

### 124. What is graceful shutdown?
"When a service receives a SIGTERM signal (like when Kubernetes checks a pod), it shouldn't just `kill -9` and drop active connections.

Graceful shutdown means: 1. Stop accepting *new* requests. 2. Allow *existing* requests a grace period (e.g., 30s) to finish processing. 3. Close database connections and file handles cleanly. 4. Then exit. Spring Boot supports this out of the box with `server.shutdown=graceful`."

**Spoken Format:**
"Graceful shutdown is like a store closing properly for the night.

**Stop accepting new requests** - This is like putting a 'Closed' sign on the door. New customers who arrive see the sign and understand not to try entering.

**Grace period for existing requests** - This is like allowing customers already inside to finish their shopping. You give them 30 minutes to complete their purchases before locking up.

**Clean shutdown** - This is like staff properly closing cash registers, cleaning the store, and locking doors. No half-finished transactions or abandoned shopping carts.

**Exit** - This is like turning off the lights and setting the alarm system.

A graceful shutdown ensures customers have a good experience even during system maintenance, and that the system can start up cleanly the next day!"

### 125. How do you handle partial failures in microservices?
"This is tricky. If I succeed in 'Payment' but fail in 'Inventory', I have a data inconsistency.

We use **Sagas** or **Eventual Consistency**.

If a step fails, I trigger a 'Compensating Transaction' (an undo operation). If I charged the card but couldn't reserve the item, I issue a refund. Ideally, I design systems to be idempotent and retryable, so I can just fix the underlying issue and replay the failed message."

**Spoken Format:**
"Partial failures in microservices are like having a team project where one person's mistake affects multiple deliverables.

Imagine you're coordinating a complex online purchase:
1. Charge customer's credit card (Payment Service)
2. Reserve product in inventory (Inventory Service)
3. Send confirmation email (Notification Service)

If step 2 fails (inventory out of stock), you have a problem: customer was charged but product wasn't reserved.

**Sagas** are like having a project manager who coordinates all the undo/redo operations:
- If any step fails, the saga automatically triggers compensating actions
- For payment failure: automatically process refund
- For inventory failure: automatically release the reserved item
- For email failure: automatically retry sending confirmation

This ensures that even if parts of the transaction fail, the system ends up in a consistent state. It's like having an automatic backup plan for complex multi-step operations!"
