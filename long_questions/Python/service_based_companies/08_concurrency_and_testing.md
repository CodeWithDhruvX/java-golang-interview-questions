# 📘 08 — Concurrency & Testing in Python
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- Threading vs Multiprocessing vs asyncio
- The GIL and when to use what
- `pytest` fundamentals
- Mocking with `unittest.mock`
- Test fixtures and parameterization

---

## ❓ Most Asked Questions

### Q1. What is the GIL and when do you use threading vs multiprocessing?

```python
# GIL (Global Interpreter Lock): CPython mutex that allows only ONE thread
# to execute Python bytecode at a time — even on multi-core systems.

# ✅ Use threading for: I/O-bound work (network calls, file reads, DB queries)
# ✅ Use multiprocessing for: CPU-bound work (data crunching, image processing)
# ✅ Use asyncio for: many concurrent I/O operations (high-performance servers)

# --- THREADING (I/O bound) ---
import threading
import time

def download_file(url):
    print(f"Downloading {url}")
    time.sleep(2)   # simulates I/O wait
    print(f"Done: {url}")

urls = ["file1.zip", "file2.zip", "file3.zip"]

threads = [threading.Thread(target=download_file, args=(url,)) for url in urls]
for t in threads:
    t.start()
for t in threads:
    t.join()   # wait for all to complete

# Thread-safe shared state with Lock
counter = 0
lock = threading.Lock()

def safe_increment():
    global counter
    for _ in range(100_000):
        with lock:   # prevents race condition
            counter += 1

# --- MULTIPROCESSING (CPU bound) ---
from multiprocessing import Pool

def cpu_work(n):
    return sum(i * i for i in range(n))

with Pool(processes=4) as pool:
    results = pool.map(cpu_work, [1_000_000] * 8)

# concurrent.futures: cleaner API for both
from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor

# I/O bound — threads
with ThreadPoolExecutor(max_workers=10) as ex:
    futures = [ex.submit(download_file, url) for url in urls]
    results = [f.result() for f in futures]

# CPU bound — processes
with ProcessPoolExecutor(max_workers=4) as ex:
    results = list(ex.map(cpu_work, [1_000_000] * 8))
```

---

### Q2. How does `asyncio` work in Python?

```python
import asyncio
import httpx

# async/await: non-blocking I/O — single-threaded event loop

async def fetch_user(client, user_id):
    resp = await client.get(f"https://api.example.com/users/{user_id}")
    return resp.json()

async def main():
    async with httpx.AsyncClient() as client:
        # Concurrent — all requests in parallel
        tasks = [fetch_user(client, uid) for uid in range(1, 6)]
        users = await asyncio.gather(*tasks)
    return users

asyncio.run(main())

# asyncio Queue (producer-consumer pattern)
async def producer(queue):
    for i in range(10):
        await queue.put(i)
        await asyncio.sleep(0.1)

async def consumer(name, queue):
    while True:
        item = await queue.get()
        print(f"{name}: processing {item}")
        await asyncio.sleep(0.2)
        queue.task_done()

async def run():
    q = asyncio.Queue(maxsize=5)
    await asyncio.gather(
        producer(q),
        consumer("Worker-1", q),
        consumer("Worker-2", q),
    )

# ⚠️ Don't block the event loop with synchronous calls!
# time.sleep(5)           # BAD — blocks all coroutines
# await asyncio.sleep(5)  # GOOD — yields to event loop
```

---

### Q3. How do you write unit tests with `pytest`?

```python
# Function under test
def add(a, b):
    return a + b

def divide(a, b):
    if b == 0:
        raise ZeroDivisionError("Cannot divide by zero")
    return a / b

# test_calculator.py
import pytest

def test_add_two_positives():
    assert add(2, 3) == 5

def test_divide_normal():
    assert divide(10, 2) == 5.0

def test_divide_by_zero():
    with pytest.raises(ZeroDivisionError, match="Cannot divide by zero"):
        divide(5, 0)

# Parameterized tests — test many cases concisely
@pytest.mark.parametrize("a, b, expected", [
    (1, 2, 3),
    (0, 0, 0),
    (-1, 1, 0),
    (100, -50, 50),
])
def test_add_parametrized(a, b, expected):
    assert add(a, b) == expected

# Fixtures — reusable setup
@pytest.fixture
def sample_user():
    return {"name": "Alice", "email": "alice@co.com", "age": 30}

def test_user_name(sample_user):
    assert sample_user["name"] == "Alice"

@pytest.fixture
def temp_file(tmp_path):
    f = tmp_path / "test.txt"
    f.write_text("Hello!")
    yield f
    # tmp_path auto-cleans up

def test_file_content(temp_file):
    assert temp_file.read_text() == "Hello!"

# Run tests:
# pytest -v                → verbose
# pytest -k "add"          → filter by name
# pytest --cov=src         → coverage report
```

---

### Q4. How do you use mocking in Python tests?

```python
from unittest.mock import Mock, patch
import pytest

# Mock: fake object for testing
m = Mock()
m.get_user.return_value = {"name": "Rahul", "email": "rahul@co.com"}

result = m.get_user(1)
assert result["name"] == "Rahul"
m.get_user.assert_called_once_with(1)

# patch: temporarily replace real objects
# Code in services.py:
# def process_user(user_id):
#     user = get_user_from_db(user_id)
#     return user["name"].upper()

@patch("services.get_user_from_db")
def test_process_user(mock_get):
    mock_get.return_value = {"name": "priya"}
    from services import process_user
    result = process_user(42)
    assert result == "PRIYA"
    mock_get.assert_called_with(42)

# side_effect — simulate exceptions
@patch("services.get_user_from_db")
def test_user_not_found(mock_get):
    mock_get.side_effect = LookupError("User not found")
    with pytest.raises(LookupError):
        from services import process_user
        process_user(999)

# patch as context manager
def test_api_call():
    with patch("requests.get") as mock_get:
        mock_get.return_value.json.return_value = {"status": "ok"}
        mock_get.return_value.status_code = 200
        # call your function that uses requests.get
        # result = call_my_api()
        # assert result["status"] == "ok"

# Best practices:
# 1. Arrange-Act-Assert pattern
# 2. One concept per test
# 3. Meaningful test names: test_what_when_expected
# 4. Test edge cases (empty inputs, boundaries, errors)
# 5. Aim for 80%+ code coverage
```
