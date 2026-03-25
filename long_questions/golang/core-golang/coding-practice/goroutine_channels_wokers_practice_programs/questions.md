Perfect—here’s your **FINAL, CLEAN, PRACTICAL-ONLY Go Concurrency Question Bank**
(no theory, no explanations — only **hands-on coding problems** like interviews expect)

---

# 🚀 FINAL GO CONCURRENCY PRACTICE SET

---

## 🟢 1. Ordering & Coordination (VERY IMPORTANT)

### 1.

Print numbers 1–20 using 2 goroutines (odd/even)
👉 Output must be strictly ordered

---

### 2.

Print:

```
1 A 2 B 3 C 4 D ...
```

👉 Use 2 goroutines

---

### 3.

3 goroutines print:

* G1 → 1,4,7...
* G2 → 2,5,8...
* G3 → 3,6,9...
  👉 Output must be in order

---

### 4.

Print 1–50 using 5 goroutines
👉 Maintain order

---

### 5.

Two goroutines alternate execution using channels
👉 Strict turn-based printing

---

### 6.

Token passing:

* Only one goroutine prints at a time
* Pass token via channel

---

### 7.

Process numbers concurrently but print results in order

---

---

## 🟡 2. Channels (Buffered & Unbuffered)

---

### 8.

Send 5 values to an unbuffered channel
👉 Show blocking behavior

---

### 9.

Repeat using buffered channel (size 3)

---

### 10.

Producer sends 10 jobs
Consumer processes slowly
👉 Use buffered channel

---

### 11.

Multiple producers → single channel → one consumer

---

### 12.

Fill buffered channel completely
👉 Handle extra send safely

---

---

## 🔵 3. Producer–Consumer Problems

---

### 13.

1 producer, 1 consumer
👉 Process 10 jobs

---

### 14.

1 producer, 2 consumers
👉 Ensure all jobs processed

---

### 15.

Multiple consumers but maintain output order

---

### 16.

Bounded buffer:

* Limit queue size
* Block producer when full

---

### 17.

Graceful shutdown:

* Stop consumers after producer finishes

---

---

## 🟣 4. Worker Pool (MOST ASKED)

---

### 18.

Create worker pool:

* 3 workers
* 10 jobs

---

### 19.

Worker pool with results collection

---

### 20.

Maintain ordered output in worker pool

---

### 21.

Limit max 3 goroutines running at a time

---

### 22.

Worker pool with retry (max 3 retries)

---

---

## 🟤 5. Fan-in / Fan-out

---

### 23.

Fan-out:

* Distribute tasks to multiple workers

---

### 24.

Fan-in:

* Merge multiple channels into one

---

### 25.

Merge 3 channels
👉 Output must be sorted

---

### 26.

Split one channel:

* Even → ch1
* Odd → ch2

---

---

## 🔴 6. select & Timeout

---

### 27.

Read from 2 channels using `select`

---

### 28.

Add timeout (2 seconds)

---

### 29.

Non-blocking read using `select`

---

### 30.

Prioritize one channel over another

---

---

## ⚫ 7. sync Package (CRITICAL)

---

### 31.

Use `WaitGroup` to wait for 5 goroutines

---

### 32.

Concurrent counter using `Mutex`

---

### 33.

Protect shared map using `Mutex`

---

### 34.

Use `RWMutex`:

* Multiple readers
* Single writer

---

### 35.

Initialize resource once using `sync.Once`

---

---

## ⚙️ 8. Context (REAL-WORLD)

---

### 36.

Cancel a goroutine using `context.WithCancel`

---

### 37.

Timeout a task using `context.WithTimeout`

---

### 38.

Pass context through multiple goroutines (pipeline)

---

---

## 🧠 9. Advanced Patterns

---

### 39.

Pipeline:

* Stage1 → generate
* Stage2 → process
* Stage3 → output

---

### 40.

Pipeline with cancellation on error

---

### 41.

Rate limiter using channel

---

### 42.

Semaphore using buffered channel

---

### 43.

Barrier:

* All goroutines complete phase 1 before phase 2

---

---

## 💣 10. Edge Cases (VERY IMPORTANT)

---

### 44.

Create and fix a deadlock

---

### 45.

Read from closed channel

---

### 46.

Write to closed channel (handle panic)

---

### 47.

Detect and fix goroutine leak

---

### 48.

Nil channel behavior:

* Read/write

---

---

## 🧪 11. Real-World Simulations

---

### 49.

Task scheduler (jobs at intervals)

---

### 50.

Parallel API calls
👉 Maintain response order

---

### 51.

Log processor:

* Multiple writers
* Single ordered output

---

### 52.

Chat system:

* Multiple senders
* Ordered processing

---

### 53.

File pipeline:

* Read → process → write

---

---

# 🎯 FINAL NOTE

If you can **implement ~80% of these without help**, you are:

✅ Ready for Go backend interviews
✅ Strong in concurrency
✅ Above average candidate

---

If you want next:

* 💻 I can give **solutions one-by-one (interview style)**
* 🧪 Or take a **mock interview**
* 🔥 Or give **top 10 hardest problems only**

Just tell me 👍
