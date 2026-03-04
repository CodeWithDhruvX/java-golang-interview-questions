# ⚡ 05 — Decorators, Generators & Context Managers
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Decorators with state and arguments
- Class-based decorators
- Generator-based pipelines
- Coroutine decorators / async decorators
- `contextlib`: `contextmanager`, `AsyncContextManager`, `ExitStack`
- `functools.wraps` internals

---

## ❓ Most Asked Questions

### Q1. Build a class-based decorator with `__call__`.

```python
import functools
import time

# Class-based decorator: cleaner than closures for decorators with state

class Retry:
    """Retry a function N times on failure with exponential backoff"""

    def __init__(self, max_retries=3, delay=1.0, exceptions=(Exception,)):
        self.max_retries = max_retries
        self.delay = delay
        self.exceptions = exceptions

    def __call__(self, func):
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            last_exc = None
            for attempt in range(1, self.max_retries + 1):
                try:
                    return func(*args, **kwargs)
                except self.exceptions as e:
                    last_exc = e
                    wait = self.delay * (2 ** (attempt - 1))   # exponential
                    print(f"Attempt {attempt} failed: {e}. Retrying in {wait}s...")
                    time.sleep(wait)
            raise RuntimeError(f"Failed after {self.max_retries} retries") from last_exc
        return wrapper

@Retry(max_retries=3, delay=0.5, exceptions=(ConnectionError, TimeoutError))
def call_external_api(url):
    import random
    if random.random() < 0.7:
        raise ConnectionError("Network timeout")
    return {"data": "success"}

try:
    result = call_external_api("https://api.example.com")
except RuntimeError as e:
    print(e)

# Timer decorator with statistics tracking
class TimingDecorator:
    def __init__(self, func):
        functools.update_wrapper(self, func)
        self.func = func
        self.call_count = 0
        self.total_time = 0.0

    def __call__(self, *args, **kwargs):
        start = time.perf_counter()
        result = self.func(*args, **kwargs)
        elapsed = time.perf_counter() - start
        self.call_count += 1
        self.total_time += elapsed
        return result

    @property
    def avg_time(self):
        return self.total_time / self.call_count if self.call_count else 0

@TimingDecorator
def sort_data(lst):
    return sorted(lst)

sort_data([3,1,2])
print(sort_data.call_count)   # 1
print(sort_data.avg_time)     # ~0.0001
```

---

### Q2. Build an advanced generator pipeline with bidirectional communication.

```python
import json
from typing import Generator

# Generator as pipeline stage: transform → forward
def read_lines(filename):
    with open(filename) as f:
        yield from f

def parse_json(source):
    for line in source:
        try:
            yield json.loads(line.strip())
        except json.JSONDecodeError:
            continue

def filter_records(source, key, value):
    for record in source:
        if record.get(key) == value:
            yield record

def transform(source, mapping_fn):
    for record in source:
        yield mapping_fn(record)

# Compose pipeline (zero intermediate lists!)
def process_log_file(filename):
    lines   = read_lines(filename)
    records = parse_json(lines)
    errors  = filter_records(records, "level", "ERROR")
    minimal = transform(errors, lambda r: {"msg": r["message"], "ts": r["timestamp"]})
    return minimal

# Bidirectional generator (send + receive)
def running_average():
    total = count = 0
    average = None
    while True:
        value = yield average      # receive via send(), yield current average
        if value is None:
            break
        total += value
        count += 1
        average = total / count

gen = running_average()
next(gen)            # prime the generator
gen.send(10)         # avg = 10.0
gen.send(20)         # avg = 15.0
gen.send(30)         # avg = 20.0
gen.close()          # send GeneratorExit

# Generator with throw() for error injection
def resilient_processor():
    while True:
        try:
            data = yield
            print(f"Processing: {data}")
        except ValueError as e:
            print(f"Skipping invalid data: {e}")
            # Generator continues!

proc = resilient_processor()
next(proc)
proc.send("valid data 1")
proc.throw(ValueError, "bad data")   # handled internally
proc.send("valid data 2")            # continues after error!
```

---

### Q3. Implement custom context managers with `contextlib`.

```python
from contextlib import contextmanager, asynccontextmanager, ExitStack
import contextlib
import time
import sqlite3

# contextmanager: generator-based context manager
@contextmanager
def timer(name=""):
    start = time.perf_counter()
    try:
        yield    # body of 'with' block runs here
    finally:
        elapsed = time.perf_counter() - start
        print(f"{name}: {elapsed:.4f}s")

with timer("data processing"):
    result = sorted(range(1_000_000))

# Contextmanager with value
@contextmanager
def temp_directory(prefix="tmp"):
    import tempfile, shutil
    tmpdir = tempfile.mkdtemp(prefix=prefix)
    try:
        yield tmpdir           # 'as' variable = tmpdir
    finally:
        shutil.rmtree(tmpdir)  # always cleanup

with temp_directory() as d:
    print(f"Working in {d}")
    # files created here are deleted on exit

# ExitStack: dynamically compose multiple context managers
def process_files(filenames):
    with ExitStack() as stack:
        files = [
            stack.enter_context(open(f)) for f in filenames
        ]
        # all files open in 'files'; all closed on exit even if one fails
        return [f.read() for f in files]

# ExitStack for conditional cleanup
def setup_resources(use_cache=True):
    stack = ExitStack()
    db = stack.enter_context(get_db_connection())
    if use_cache:
        cache = stack.enter_context(get_redis_connection())
    return stack, db   # caller owns stack lifetime

# Async context manager
@asynccontextmanager
async def managed_connection(url):
    import httpx
    async with httpx.AsyncClient() as client:
        try:
            yield client
        finally:
            pass  # httpx cleanup handled by its own __aexit__

async def fetch_data():
    async with managed_connection("https://api.example.com") as client:
        resp = await client.get("/data")
        return resp.json()
```

---

### Q4. Implement async decorators and decorator stacking strategies.

```python
import functools
import asyncio
import time

# Async-compatible decorator
def rate_limited(calls_per_second: float):
    min_interval = 1.0 / calls_per_second
    last_call = [0.0]

    def decorator(func):
        @functools.wraps(func)
        async def async_wrapper(*args, **kwargs):
            now = time.monotonic()
            elapsed = now - last_call[0]
            if elapsed < min_interval:
                await asyncio.sleep(min_interval - elapsed)
            last_call[0] = time.monotonic()
            return await func(*args, **kwargs)

        @functools.wraps(func)
        def sync_wrapper(*args, **kwargs):
            now = time.monotonic()
            elapsed = now - last_call[0]
            if elapsed < min_interval:
                time.sleep(min_interval - elapsed)
            last_call[0] = time.monotonic()
            return func(*args, **kwargs)

        # Return appropriate wrapper
        if asyncio.iscoroutinefunction(func):
            return async_wrapper
        return sync_wrapper
    return decorator

@rate_limited(calls_per_second=10)
async def call_api(endpoint):
    return f"Response from {endpoint}"

# Decorator that preserves type signatures (advanced)
from typing import TypeVar, Callable, Any
F = TypeVar("F", bound=Callable[..., Any])

def logged(func: F) -> F:
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        print(f"Calling {func.__name__}({args!r}, {kwargs!r})")
        result = func(*args, **kwargs)
        print(f"{func.__name__} returned {result!r}")
        return result
    return wrapper   # type: ignore[return-value]

# Stacking order matters!
@logged   # applied second (outer wrapper)
@rate_limited(10)    # applied first (inner wrapper)
def my_function(x):
    return x * 2

# Execution order: logged.__call__ → rate_limited.__call__ → my_function
```
