# 💻 Live Coding + Architecture — Product-Based Companies

> **Level:** 🔴 Senior / Staff
> **Asked at:** Uber, Amazon, Flipkart, Swiggy, Razorpay
> **Format:** "Machine Coding" or "Craftsmanship" rounds where you write executable code that demonstrates architectural patterns.

---

## The "Live Coding Architecture" Interview Format

At companies like Flipkart, Swiggy, and Uber, the **Machine Coding round (2 hours)** is effectively an architecture round where you must write executable code. You aren't expected to write a Spring Boot app with real DB connections. 

Instead, you write **in-memory, executable core logic** that demonstrates distributed systems patterns (like concurrency control, event-driven design, and SOLID principles) using pure Java or Go.

Requirements generally include:
1. Working, executable code.
2. In-memory data structures (no DB).
3. Concurrency handling (thread safety).
4. Good domain modeling (Classes/Structs).
5. Clean architecture (separation of presentation, business, and storage).

---

## Q1. Design and Implement an In-Memory Rate Limiter (Token Bucket)

**Scenario:** "Implement a rate limiter for our API gateway. It should support 100 requests per minute per user. Write executable code that handles concurrent requests correctly."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Amazon (SDE-2/3), Razorpay, Zepto

#### Architectural Focus
- Thread safety (concurrency)
- Memory management (evicting stale users)
- Low latency (no blocking operations unnecessarily)

#### Java Implementation (Core Logic)
```java
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicInteger;

public class RateLimiter {
    private final int maxTokens;
    private final long refillIntervalMs;
    
    // Thread-safe map storing buckets for each user
    private final ConcurrentHashMap<String, TokenBucket> userBuckets = new ConcurrentHashMap<>();

    public RateLimiter(int maxTokens, long refillIntervalMs) {
        this.maxTokens = maxTokens;
        this.refillIntervalMs = refillIntervalMs;
        
        // Optional: Background thread to clean up inactive users to prevent memory leak
        startCleanupTask();
    }

    public boolean allowRequest(String userId) {
        // computeIfAbsent is atomic in ConcurrentHashMap
        TokenBucket bucket = userBuckets.computeIfAbsent(userId, k -> new TokenBucket(maxTokens));
        return bucket.tryConsume();
    }

    private class TokenBucket {
        private final AtomicInteger tokens;
        private volatile long lastRefillTime;

        public TokenBucket(int maxTokens) {
            this.tokens = new AtomicInteger(maxTokens);
            this.lastRefillTime = System.currentTimeMillis();
        }

        public boolean tryConsume() {
            refill(); // Lazy refill on every request
            
            // Lock-free decrement
            while (true) {
                int currentTokens = tokens.get();
                if (currentTokens == 0) {
                    return false; // Rate limited
                }
                if (tokens.compareAndSet(currentTokens, currentTokens - 1)) {
                    return true; // Successfully consumed a token
                }
                // If compareAndSet fails, another thread modified tokens; loop and try again
            }
        }

        private void refill() {
            long now = System.currentTimeMillis();
            long timeElapsed = now - lastRefillTime;
            
            if (timeElapsed > refillIntervalMs) {
                // Calculate how many tokens to add based on time elapsed
                int tokensToAdd = (int) (timeElapsed / refillIntervalMs) * maxTokens; // Simplified for bursting
                
                // Only refill if we actually need to
                if (tokensToAdd > 0) {
                    // Update timestamp (risk of slight race condition here, but acceptable for rate limiting precision)
                    lastRefillTime = now;
                    
                    // Lock-free increment up to maxTokens
                    int current;
                    do {
                        current = tokens.get();
                        if (current >= maxTokens) break;
                    } while (!tokens.compareAndSet(current, Math.min(maxTokens, current + tokensToAdd)));
                }
            }
        }
    }
    
    private void startCleanupTask() {
        ScheduledExecutorService scheduler = Executors.newScheduledThreadPool(1);
        scheduler.scheduleAtFixedRate(() -> {
            long now = System.currentTimeMillis();
            // Remove buckets not accessed in the last 10 minutes
            userBuckets.entrySet().removeIf(entry -> 
                (now - entry.getValue().lastRefillTime) > TimeUnit.MINUTES.toMillis(10));
        }, 10, 10, TimeUnit.MINUTES);
    }
}
```

---

## Q2. Implement an In-Memory Task Scheduler (Delayed Queue)

**Scenario:** "Design a task scheduler that can accept tasks with a scheduled execution time in the future. Execute them as close to their scheduled time as possible."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Uber, Swiggy, PhonePe

#### Architectural Focus
- Priority scheduling (Min-Heap / PriorityQueue)
- Producer-Consumer pattern
- Thread signaling (`wait() / notify()` or `Condition` variables)

#### Go Implementation
```go
package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

// --- Domain Models ---
type Task struct {
	ID        string
	ExecuteAt time.Time
	Payload   func()
	index     int // For heap interface
}

// --- Priority Queue (Min-Heap by ExecuteAt) ---
type TaskQueue []*Task

func (pq TaskQueue) Len() int { return len(pq) }
func (pq TaskQueue) Less(i, j int) bool {
	return pq[i].ExecuteAt.Before(pq[j].ExecuteAt)
}
func (pq TaskQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *TaskQueue) Push(x interface{}) {
	n := len(*pq)
	task := x.(*Task)
	task.index = n
	*pq = append(*pq, task)
}
func (pq *TaskQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	task := old[n-1]
	task.index = -1
	*pq = old[0 : n-1]
	return task
}

// --- Scheduler Architecture ---
type Scheduler struct {
	mu        sync.Mutex
	pq        TaskQueue
	cond      *sync.Cond
	isStopped bool
}

func NewScheduler() *Scheduler {
	s := &Scheduler{
		pq: make(TaskQueue, 0),
	}
	s.cond = sync.NewCond(&s.mu)
	heap.Init(&s.pq)
	return s
}

// Producer
func (s *Scheduler) Schedule(id string, delay time.Duration, run func()) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := &Task{
		ID:        id,
		ExecuteAt: time.Now().Add(delay),
		Payload:   run,
	}
	heap.Push(&s.pq, task)
	
	// Wake up consumer thread since a new task arrived
	// (It might be earlier than what consumer is currently waiting for)
	s.cond.Signal() 
}

// Consumer
func (s *Scheduler) Start() {
	go func() {
		for {
			s.mu.Lock()

			if s.isStopped {
				s.mu.Unlock()
				return
			}

			if len(s.pq) == 0 {
				// Queue empty, wait for signal
				s.cond.Wait()
				s.mu.Unlock()
				continue
			}

			// Peek at the soonest task
			now := time.Now()
			nextTask := s.pq[0]

			if now.After(nextTask.ExecuteAt) {
				// Task is due! Remove from heap and execute
				taskToRun := heap.Pop(&s.pq).(*Task)
				s.mu.Unlock()

				// Execute outside the lock to prevent blocking the scheduler
				go func(t *Task) {
					fmt.Printf("[%s] Executing task %s\n", time.Now().Format("15:04:05.000"), t.ID)
					t.Payload()
				}(taskToRun)
				
				continue // Check next task immediately
			}

			// Task is not due yet. Wait until it is, or until a new task arrives
			waitDuration := nextTask.ExecuteAt.Sub(now)
			
			// Start a timer. We use a channel to wake up from wait.
			timer := time.NewTimer(waitDuration)
			
			// We must drop the lock while waiting, but sync.Cond.Wait() doesn't take a timeout.
			// Workaround: custom wait loop or letting another goroutine signal us.
			s.mu.Unlock()

			// Block until timer fires OR a new task is scheduled (signaled over a channel)
			// *Simplified for interview context*: In a real system, you'd use a channel `wakeup` mapped to `cond`.
			<-timer.C
		}
	}()
}
```

---

## Q3. Implement an In-Memory Pub/Sub Event Broker

**Scenario:** "Design a message queue (like a lightweight Kafka). Consumers can subscribe to topics. Producers publish to topics. Handle concurrency and ensure consumers process messages in the background."

#### Company Context & Level
**Level:** 🔴 Senior | **Asked at:** Flipkart, MakeMyTrip, Swiggy

#### Architectural Focus
- Observer / Pub-Sub Pattern
- Decoupling publishers from subscribers
- Background processing (thread pools or goroutines)

#### Java Implementation
```java
import java.util.*;
import java.util.concurrent.*;

// Message Domain
record Message(String id, String payload, long timestamp) {}

// Subscriber Interface
interface Subscriber {
    void onMessage(Message topic);
    String getId();
}

public class EventBroker {
    // Topic -> List of Subscribers
    // Using CopyOnWriteArrayList so we can iterate safely while others modify it
    private final Map<String, List<Subscriber>> topicSubscribers = new ConcurrentHashMap<>();
    
    // Thread pool to deliver messages asynchronously (don't block the Publisher)
    private final ExecutorService deliveryExecutor = Executors.newFixedThreadPool(10);

    public void subscribe(String topic, Subscriber subscriber) {
        topicSubscribers
            .computeIfAbsent(topic, k -> new CopyOnWriteArrayList<>())
            .add(subscriber);
        System.out.println("Subscribed " + subscriber.getId() + " to " + topic);
    }

    public void publish(String topic, Message message) {
        List<Subscriber> subscribers = topicSubscribers.get(topic);
        if (subscribers == null || subscribers.isEmpty()) {
            return; // No one listening
        }

        // Fan-out: deliver to all subscribers asynchronously
        for (Subscriber sub : subscribers) {
            deliveryExecutor.submit(() -> {
                try {
                    sub.onMessage(message);
                } catch (Exception e) {
                    System.err.println("Subscriber " + sub.getId() + " failed: " + e.getMessage());
                    // In a real system: DLQ (Dead Letter Queue) or Retry logic here
                }
            });
        }
    }
    
    public void shutdown() {
        deliveryExecutor.shutdown();
    }
}

// --- Usage Example ---
class EmailService implements Subscriber {
    public void onMessage(Message msg) {
        System.out.println("[EmailService] Sending email for: " + msg.payload());
    }
    public String getId() { return "email-svc-1"; }
}

class AnalyticsService implements Subscriber {
    public void onMessage(Message msg) {
        System.out.println("[AnalyticsService] Tracking event: " + msg.payload());
    }
    public String getId() { return "analytics-svc-1"; }
}

// In main():
// EventBroker broker = new EventBroker();
// broker.subscribe("ORDER_PLACED", new EmailService());
// broker.subscribe("ORDER_PLACED", new AnalyticsService());
// broker.publish("ORDER_PLACED", new Message("1", "Order details", System.currentTimeMillis()));
```

### Why this format matters
If you are asked to "Design splitwise" or "Design a Cab Booking system" in a 2-hour coding round:
1. Don't write uncompilable pseudo-code.
2. Use strong types, interfaces, and separation of concerns.
3. Your lock management (Mutex/synchronized) and thread-safe collections (`ConcurrentHashMap`) are precisely what they evaluate to see if you understand real-world application architecture.
