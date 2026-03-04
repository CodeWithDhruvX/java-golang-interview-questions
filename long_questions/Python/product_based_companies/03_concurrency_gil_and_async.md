# ⚡ 03 — Concurrency, GIL & Async Python
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Thread safety, Race conditions, Locks
- asyncio internals: event loop, coroutines, tasks
- `asyncio.gather`, `asyncio.wait`, `asyncio.Queue`
- Semaphores and rate limiting
- `trio` and `anyio` alternatives
- Integrating sync and async code

---

## ❓ Most Asked Questions

### Q1. Explain `asyncio` internals — event loop, coroutines, and tasks.

```python
import asyncio

# Coroutine: a function defined with async def
# It doesn't run immediately — creates a coroutine object
async def greet(name, delay):
    await asyncio.sleep(delay)   # yield control to event loop
    print(f"Hello, {name}!")
    return f"greeted {name}"

# Running the event loop
asyncio.run(greet("Rahul", 1))   # runs until coroutine completes

# Task: wraps a coroutine for concurrent scheduling
async def main():
    # create_task schedules immediately (doesn't block here)
    t1 = asyncio.create_task(greet("Alice", 2))
    t2 = asyncio.create_task(greet("Bob", 1))
    t3 = asyncio.create_task(greet("Carol", 3))

    # Wait for all tasks
    results = await asyncio.gather(t1, t2, t3)
    # Bob greets first (1s), then Alice (2s), then Carol (3s)
    # Total time: ~3s, not 6s!
    print(results)

asyncio.run(main())

# asyncio.wait — more control than gather
async def fetch(n):
    await asyncio.sleep(n * 0.1)
    if n == 3:
        raise ValueError(f"Error on {n}")
    return n

async def with_wait():
    tasks = [asyncio.create_task(fetch(i)) for i in range(1, 5)]

    # Wait until ALL done (including failures)
    done, pending = await asyncio.wait(tasks, return_when=asyncio.ALL_COMPLETED)

    for task in done:
        if task.exception():
            print(f"Failed: {task.exception()}")
        else:
            print(f"Result: {task.result()}")

    # Or: return when FIRST completes
    done, pending = await asyncio.wait(tasks, return_when=asyncio.FIRST_COMPLETED)
    for p in pending:
        p.cancel()   # cancel remaining tasks

# asyncio.timeout (Python 3.11+)
async def with_timeout():
    try:
        async with asyncio.timeout(2.0):
            await asyncio.sleep(5)   # too slow!
    except asyncio.TimeoutError:
        print("Timed out!")
```

---

### Q2. Implement a rate limiter and semaphore with asyncio.

```python
import asyncio
import time

# Semaphore: limit concurrent coroutines
async def limited_fetch(semaphore, url):
    async with semaphore:   # only N coroutines enter at a time
        await asyncio.sleep(0.1)  # simulate I/O
        return f"data from {url}"

async def fetch_all_limited(urls, max_concurrent=5):
    semaphore = asyncio.Semaphore(max_concurrent)
    tasks = [limited_fetch(semaphore, url) for url in urls]
    return await asyncio.gather(*tasks)

# Rate limiter: N requests per second
class AsyncRateLimiter:
    def __init__(self, rate: int):
        self.rate = rate            # requests per second
        self.tokens = rate
        self.last_refill = time.monotonic()
        self._lock = asyncio.Lock()

    async def acquire(self):
        async with self._lock:
            now = time.monotonic()
            elapsed = now - self.last_refill
            self.tokens = min(self.rate, self.tokens + elapsed * self.rate)
            self.last_refill = now

            if self.tokens < 1:
                wait_time = (1 - self.tokens) / self.rate
                await asyncio.sleep(wait_time)
                self.tokens = 0
            else:
                self.tokens -= 1

    async def __aenter__(self):
        await self.acquire()
        return self

    async def __aexit__(self, *args):
        pass

async def make_api_call(limiter, endpoint):
    async with limiter:
        return f"Response from {endpoint}"

async def main():
    limiter = AsyncRateLimiter(rate=10)  # 10 req/sec
    tasks = [make_api_call(limiter, f"/endpoint/{i}") for i in range(50)]
    results = await asyncio.gather(*tasks)
```

---

### Q3. How do you bridge sync and async code?

```python
import asyncio
from concurrent.futures import ThreadPoolExecutor

# Problem 1: calling blocking sync code from async
def blocking_io():
    import time
    time.sleep(1)   # BLOCKS the event loop if called with await!
    return "slow result"

async def async_with_blocking():
    loop = asyncio.get_event_loop()

    # Run in thread pool → doesn't block event loop
    result = await loop.run_in_executor(None, blocking_io)
    print(result)

    # With specific executor
    with ThreadPoolExecutor(max_workers=4) as pool:
        result = await loop.run_in_executor(pool, blocking_io)

# asyncio.to_thread (Python 3.9+) — cleaner API
async def use_to_thread():
    result = await asyncio.to_thread(blocking_io)  # same thing, cleaner
    print(result)

# Problem 2: calling async code from sync context
def sync_caller():
    result = asyncio.run(async_function())   # creates new event loop
    return result

# Problem 3: sharing state between threads and asyncio
# asyncio.run_coroutine_threadsafe — call coroutine from another thread
import threading

async def async_work(value):
    await asyncio.sleep(0.1)
    return value * 2

def thread_func(loop):
    future = asyncio.run_coroutine_threadsafe(async_work(21), loop)
    result = future.result(timeout=5)    # blocks thread until done
    print(result)   # 42

async def main():
    loop = asyncio.get_event_loop()
    t = threading.Thread(target=thread_func, args=(loop,))
    t.start()
    await asyncio.sleep(1)
    t.join()
```

---

### Q4. What are common asyncio pitfalls and how do you avoid them?

```python
import asyncio

# Pitfall 1: forgetting await (fire and forget)
async def bad():
    asyncio.sleep(1)   # ❌ returns coroutine object — NOT awaited!
    print("Done")      # runs immediately!

async def good():
    await asyncio.sleep(1)   # ✅ actually waits 1 second

# Pitfall 2: blocking the event loop
async def blocking_bad():
    import time
    time.sleep(5)         # ❌ blocks ALL coroutines for 5 seconds!

async def blocking_good():
    await asyncio.sleep(5)   # ✅ or use run_in_executor for real blocking I/O

# Pitfall 3: uncaught task exceptions are silently lost
async def failing_task():
    await asyncio.sleep(0.1)
    raise ValueError("Unexpected error!")

async def main_bad():
    t = asyncio.create_task(failing_task())
    await asyncio.sleep(1)
    # t's exception is silently ignored! (prints warning but app continues)

async def main_good():
    t = asyncio.create_task(failing_task())
    try:
        await t
    except ValueError as e:
        print(f"Task failed: {e}")

    # Or add exception handler:
    def handle_exception(task):
        if not task.cancelled():
            exc = task.exception()
            if exc:
                print(f"Task {task.get_name()} failed: {exc}")

    t = asyncio.create_task(failing_task())
    t.add_done_callback(handle_exception)

# Pitfall 4: not cancelling pending tasks on shutdown
async def main():
    tasks = [asyncio.create_task(work(i)) for i in range(10)]
    try:
        await asyncio.gather(*tasks)
    except asyncio.CancelledError:
        for t in tasks:
            t.cancel()
        await asyncio.gather(*tasks, return_exceptions=True)
```
