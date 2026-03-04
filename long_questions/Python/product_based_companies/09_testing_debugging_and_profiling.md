# ⚡ 09 — Testing, Debugging & Profiling
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Advanced `pytest`: fixtures, parametrize, markers, plugins
- Property-based testing with `hypothesis`
- Profiling: `cProfile`, `line_profiler`, `py-spy`
- Debugging with `pdb`, `ipdb`
- Mutation testing
- Contract testing

---

## ❓ Most Asked Questions

### Q1. Advanced `pytest` patterns for production test suites.

```python
import pytest
from unittest.mock import AsyncMock, patch
from fastapi.testclient import TestClient

# conftest.py — shared fixtures across test files
# pytest auto-discovers conftest.py at each directory level

# Scope: function(default), class, module, session
@pytest.fixture(scope="session")
def db_engine():
    """Session-scoped: created once for entire test run"""
    from sqlalchemy import create_engine
    from myapp.models import Base

    engine = create_engine("sqlite:///:memory:")
    Base.metadata.create_all(engine)
    yield engine
    engine.dispose()

@pytest.fixture
def db_session(db_engine):
    """Function-scoped: fresh transaction per test"""
    from sqlalchemy.orm import Session
    with Session(db_engine) as session:
        session.begin_nested()   # savepoint — allows rollback
        yield session
        session.rollback()       # undo: test isolation!

@pytest.fixture
def test_client(db_session):
    """TestClient with overridden DB dependency"""
    from myapp.main import app
    from myapp.dependencies import get_db

    app.dependency_overrides[get_db] = lambda: db_session
    with TestClient(app) as client:
        yield client
    app.dependency_overrides.clear()

# Custom pytest markers
@pytest.mark.slow
@pytest.mark.integration
def test_full_user_registration_flow(test_client, db_session):
    # Register
    resp = test_client.post("/users", json={"name": "Alice", "email": "a@b.com"})
    assert resp.status_code == 201
    user_id = resp.json()["id"]

    # Login
    resp = test_client.post("/auth/login", json={"email": "a@b.com"})
    assert resp.status_code == 200
    token = resp.json()["token"]

    # Access protected resource
    resp = test_client.get(f"/users/{user_id}", headers={"Authorization": f"Bearer {token}"})
    assert resp.json()["name"] == "Alice"

# pytest.ini or pyproject.toml:
# [pytest]
# markers =
#     slow: marks tests as slow
#     integration: marks integration tests
# Run: pytest -m "not slow"  (skip slow tests in CI)

# Async tests (pytest-asyncio)
import pytest
import asyncio

@pytest.mark.asyncio
async def test_async_service():
    from myapp.services import fetch_user

    with patch("myapp.services.httpx.AsyncClient") as MockClient:
        mock_instance = MockClient.return_value.__aenter__.return_value
        mock_instance.get = AsyncMock(return_value=...)
        result = await fetch_user(1)
        assert result["id"] == 1
```

---

### Q2. Property-based testing with `hypothesis`.

```python
from hypothesis import given, settings, assume, example
from hypothesis import strategies as st

# Property-based testing: generate random inputs to find edge cases

# Traditional test: specific inputs
def test_add_specific():
    assert add(2, 3) == 5

# Property test: works for ALL integers
@given(st.integers(), st.integers())
def test_add_commutative(a, b):
    assert add(a, b) == add(b, a)   # commutativity

@given(st.integers())
def test_add_identity(n):
    assert add(n, 0) == n   # identity element

# String operations
@given(st.text())
def test_reverse_involution(s):
    assert reverse(reverse(s)) == s    # reversing twice = original

@given(st.lists(st.integers()))
def test_sort_idempotent(lst):
    sorted_once = sorted(lst)
    sorted_twice = sorted(sorted_once)
    assert sorted_once == sorted_twice

# Custom strategies
employee_strategy = st.fixed_dictionaries({
    "name": st.text(min_size=2, max_size=50, alphabet=st.characters(whitelist_categories=("L",))),
    "salary": st.floats(min_value=10000, max_value=1_000_000, allow_nan=False),
    "dept": st.sampled_from(["Eng", "HR", "Finance", "Marketing"]),
})

@given(employee_strategy)
def test_employee_creation(emp_data):
    emp = Employee(**emp_data)
    assert emp.salary > 0
    assert len(emp.name) >= 2

# Specific example to always test
@example(lst=[])
@example(lst=[1])
@given(st.lists(st.integers()))
def test_max_is_at_least_each_element(lst):
    assume(len(lst) > 0)   # skip empty lists
    m = max(lst)
    for x in lst:
        assert m >= x

# Settings: control how hard hypothesis tries
@settings(max_examples=1000, deadline=None)
@given(st.integers(min_value=0, max_value=1000))
def test_factorial(n):
    result = factorial(n)
    assert result >= 1
    if n > 0:
        assert result % n == 0
```

---

### Q3. How do you profile Python code and fix bottlenecks?

```python
import cProfile
import pstats
import io
import time
from line_profiler import LineProfiler

# --- cProfile: function-level profiling ---

def profile(func):
    """Decorator to profile a function"""
    def wrapper(*args, **kwargs):
        pr = cProfile.Profile()
        pr.enable()
        result = func(*args, **kwargs)
        pr.disable()

        s = io.StringIO()
        ps = pstats.Stats(pr, stream=s).sort_stats("cumulative")
        ps.print_stats(20)   # top 20 functions
        print(s.getvalue())
        return result
    return wrapper

@profile
def expensive_function():
    return sorted([i**2 for i in range(100_000)])

# Command line profiling:
# python -m cProfile -s cumulative script.py | head -40
# python -m cProfile -o output.prof script.py
# snakeviz output.prof    # visual flame graph!

# --- line_profiler: line-by-line timing ---
def slow_function(n):
    result = []
    for i in range(n):          # ← which line is slow?
        result.append(i ** 2)
    return sorted(result)

lp = LineProfiler()
lp.add_function(slow_function)
lp.run("slow_function(100_000)")
lp.print_stats()

# --- timeit: benchmark small snippets ---
import timeit

# Compare list vs generator for sum
list_time = timeit.timeit(
    "sum([x**2 for x in range(10000)])",
    number=1000
)
gen_time = timeit.timeit(
    "sum(x**2 for x in range(10000))",
    number=1000
)
print(f"List: {list_time:.3f}s, Gen: {gen_time:.3f}s")

# Common Python optimization tips:
# 1. Use local variable aliases in tight loops
def fast_sum(lst):
    local_sum = 0  # access local var is faster than global
    append = local_sum  # local alias for method
    for x in lst:
        local_sum += x
    return local_sum

# 2. Use built-ins (C-speed): sum(), min(), max(), map(), filter()
# 3. Avoid repeated dict lookups in loops
# 4. Use sets for O(1) membership checks vs O(n) list lookup
# 5. Disable GC during large batch operations
# 6. Use __slots__ for classes with many instances
# 7. Vectorize with NumPy for numeric operations
```

---

### Q4. How do you debug Python code effectively?

```python
import pdb
import logging

# --- 1. pdb: Python Debugger ---
def buggy_function(data):
    result = []
    for i, item in enumerate(data):
        # breakpoint()   # Python 3.7+ shorthand for pdb.set_trace()
        pdb.set_trace()  # drops into interactive debugger here
        result.append(process(item))
    return result

# pdb commands:
# n (next)     → execute next line
# s (step)     → step into function
# c (continue) → run until next breakpoint
# l (list)     → show surrounding code
# p expr       → print expression value
# pp expr      → pretty-print
# w (where)    → show call stack
# b line_no    → set breakpoint
# q (quit)     → exit debugger

# --- 2. Logging (production-grade debugging) ---
logging.basicConfig(
    level=logging.DEBUG,
    format="%(asctime)s %(name)s %(levelname)s %(message)s",
    handlers=[
        logging.FileHandler("app.log"),
        logging.StreamHandler()
    ]
)
logger = logging.getLogger(__name__)

def process_order(order_id, amount):
    logger.debug(f"Processing order {order_id} for ₹{amount:,}")
    try:
        result = charge_payment(order_id, amount)
        logger.info(f"Order {order_id} completed: {result}")
        return result
    except Exception as e:
        logger.exception(f"Order {order_id} failed: {e}")  # includes traceback!
        raise

# Structured logging (JSON) for log aggregation (ELK, CloudWatch):
import json

class JSONFormatter(logging.Formatter):
    def format(self, record):
        log = {
            "timestamp": self.formatTime(record),
            "level": record.levelname,
            "message": record.getMessage(),
            "module": record.module,
        }
        if record.exc_info:
            log["exception"] = self.formatException(record.exc_info)
        return json.dumps(log)

# --- 3. Rich tracebacks (pip install rich) ---
from rich.traceback import install
install(show_locals=True)   # beautiful tracebacks with local variable values!
```
