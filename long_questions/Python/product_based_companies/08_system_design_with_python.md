# ⚡ 08 — System Design with Python
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Building scalable REST APIs (FastAPI)
- Caching patterns: Redis, `lru_cache`, in-process
- Rate limiting in production
- Background workers and task queues (Celery/RQ)
- Circuit breaker and bulkhead patterns
- gRPC and microservice communication

---

## ❓ Most Asked Questions

### Q1. How do you implement caching in a Python service?

```python
import time
import redis
import json
from functools import lru_cache, wraps

# --- 1. In-process LRU Cache (single instance) ---
@lru_cache(maxsize=1000)
def get_user_from_db(user_id: int):
    """Cached for 1000 unique user IDs"""
    # simulate DB query
    return {"id": user_id, "name": f"User-{user_id}"}

# Cache time-to-live (TTL) with timestamp
def ttl_cache(seconds: int):
    cache = {}

    def decorator(func):
        @wraps(func)
        def wrapper(*args):
            now = time.time()
            if args in cache:
                result, expiry = cache[args]
                if now < expiry:
                    return result
                del cache[args]   # expired
            result = func(*args)
            cache[args] = (result, now + seconds)
            return result
        wrapper.cache = cache
        return wrapper
    return decorator

@ttl_cache(seconds=300)   # cache for 5 minutes
def get_config(key: str):
    return {"key": key, "value": "some_value"}

# --- 2. Redis Cache (distributed, multi-instance) ---
class RedisCache:
    def __init__(self, host="localhost", port=6379, db=0):
        self.client = redis.Redis(host=host, port=port, db=db, decode_responses=True)

    def get(self, key: str):
        value = self.client.get(key)
        return json.loads(value) if value else None

    def set(self, key: str, value, ttl: int = 300):
        self.client.setex(key, ttl, json.dumps(value, default=str))

    def delete(self, key: str):
        self.client.delete(key)

    def cache(self, ttl: int = 300, key_prefix: str = ""):
        """Decorator to cache function results in Redis"""
        def decorator(func):
            @wraps(func)
            def wrapper(*args, **kwargs):
                cache_key = f"{key_prefix}:{func.__name__}:{args}:{sorted(kwargs.items())}"
                cached = self.get(cache_key)
                if cached is not None:
                    return cached
                result = func(*args, **kwargs)
                self.set(cache_key, result, ttl)
                return result
            return wrapper
        return decorator

cache = RedisCache()

@cache.cache(ttl=600, key_prefix="user")
def fetch_user_profile(user_id: int):
    # expensive DB call
    return {"id": user_id, "profile": "..."}

# Cache invalidation strategies:
# 1. TTL-based: set expiry → eventual consistency
# 2. Write-through: update cache when DB updates
# 3. Write-behind: update DB asynchronously after caching
# 4. Cache-aside: app checks cache, falls back to DB on miss
```

---

### Q2. How do you implement a production-grade rate limiter?

```python
import time
import redis
from functools import wraps
from fastapi import Request, HTTPException

# Token Bucket Rate Limiter using Redis (distributed)
class RedisRateLimiter:
    def __init__(self, redis_client: redis.Redis):
        self.redis = redis_client

    def is_allowed(self, key: str, max_requests: int, window_seconds: int) -> tuple[bool, dict]:
        """Sliding window rate limiter using Redis sorted sets"""
        now = time.time()
        window_start = now - window_seconds
        pipe = self.redis.pipeline()

        # Remove old entries outside the window
        pipe.zremrangebyscore(key, 0, window_start)
        # Count requests in window
        pipe.zcard(key)
        # Add current request timestamp
        pipe.zadd(key, {str(now): now})
        # Set expiry to auto-cleanup
        pipe.expire(key, window_seconds * 2)

        _, count, _, _ = pipe.execute()

        allowed = count < max_requests
        remaining = max(0, max_requests - count - 1)
        reset_at = int(window_start + window_seconds)

        return allowed, {
            "limit": max_requests,
            "remaining": remaining,
            "reset": reset_at,
        }

# FastAPI middleware
from fastapi import FastAPI
from starlette.middleware.base import BaseHTTPMiddleware

class RateLimitMiddleware(BaseHTTPMiddleware):
    def __init__(self, app, limiter: RedisRateLimiter, max_req=100, window=60):
        super().__init__(app)
        self.limiter = limiter
        self.max_req = max_req
        self.window = window

    async def dispatch(self, request: Request, call_next):
        client_ip = request.client.host
        key = f"rate_limit:{client_ip}"

        allowed, info = self.limiter.is_allowed(key, self.max_req, self.window)

        if not allowed:
            from fastapi.responses import JSONResponse
            return JSONResponse(
                status_code=429,
                content={"error": "Too Many Requests", **info},
                headers={
                    "X-RateLimit-Limit": str(info["limit"]),
                    "X-RateLimit-Remaining": str(info["remaining"]),
                    "Retry-After": str(info["reset"]),
                }
            )

        response = await call_next(request)
        response.headers["X-RateLimit-Limit"] = str(info["limit"])
        response.headers["X-RateLimit-Remaining"] = str(info["remaining"])
        return response
```

---

### Q3. How do you implement a Circuit Breaker pattern?

```python
import time
from enum import Enum
from dataclasses import dataclass, field
from typing import Callable

class CircuitState(Enum):
    CLOSED = "closed"       # working normally
    OPEN = "open"           # rejecting requests (failing)
    HALF_OPEN = "half_open" # testing if service recovered

@dataclass
class CircuitBreaker:
    failure_threshold: int = 5
    recovery_timeout: float = 30.0
    half_open_max_calls: int = 3

    _state: CircuitState = field(default=CircuitState.CLOSED, init=False)
    _failure_count: int = field(default=0, init=False)
    _last_failure_time: float = field(default=0.0, init=False)
    _half_open_calls: int = field(default=0, init=False)

    @property
    def state(self) -> CircuitState:
        if self._state == CircuitState.OPEN:
            if time.time() - self._last_failure_time > self.recovery_timeout:
                self._state = CircuitState.HALF_OPEN
                self._half_open_calls = 0
        return self._state

    def call(self, func: Callable, *args, **kwargs):
        state = self.state

        if state == CircuitState.OPEN:
            raise RuntimeError(f"Circuit breaker OPEN — service unavailable")

        if state == CircuitState.HALF_OPEN:
            if self._half_open_calls >= self.half_open_max_calls:
                raise RuntimeError("Half-open call limit reached")
            self._half_open_calls += 1

        try:
            result = func(*args, **kwargs)
            self._on_success()
            return result
        except Exception as e:
            self._on_failure()
            raise

    def _on_success(self):
        self._failure_count = 0
        self._state = CircuitState.CLOSED

    def _on_failure(self):
        self._failure_count += 1
        self._last_failure_time = time.time()
        if self._failure_count >= self.failure_threshold:
            self._state = CircuitState.OPEN

# Usage
import requests

cb = CircuitBreaker(failure_threshold=3, recovery_timeout=10.0)

def make_api_call():
    resp = requests.get("https://unreliable-service.com/api", timeout=2)
    resp.raise_for_status()
    return resp.json()

for i in range(10):
    try:
        result = cb.call(make_api_call)
        print(f"Success: {result}")
    except RuntimeError as e:
        print(f"Circuit Breaker blocked: {e}")
    except Exception as e:
        print(f"Service error: {e}")
    time.sleep(1)
```

---

### Q4. How do you use Celery for background task queues?

```python
# Celery: distributed task queue — offload heavy work to background workers

# celery_app.py
from celery import Celery
from celery.utils.log import get_task_logger

app = Celery(
    "tasks",
    broker="redis://localhost:6379/0",   # message broker
    backend="redis://localhost:6379/1",  # result storage
)

app.conf.update(
    task_serializer="json",
    accept_content=["json"],
    result_expires=3600,          # results expire in 1 hour
    worker_max_tasks_per_child=100,  # restart worker after 100 tasks (memory leak protection)
)

logger = get_task_logger(__name__)

@app.task(bind=True, max_retries=3, default_retry_delay=60)
def send_email(self, recipient: str, subject: str, body: str):
    try:
        logger.info(f"Sending email to {recipient}")
        result = email_service.send(recipient, subject, body)
        return {"status": "sent", "id": result.id}
    except Exception as exc:
        logger.error(f"Email failed: {exc}")
        raise self.retry(exc=exc)   # retry with exponential backoff

@app.task
def generate_report(report_id: int, filters: dict):
    # Long-running computation — safe to run in background
    data = fetch_data(filters)
    report = build_report(data)
    save_report(report_id, report)
    notify_user_report_ready(report_id)

# In your FastAPI/Flask endpoint:
@app.post("/reports")
async def create_report(filters: dict, background_tasks):
    report_id = create_report_record()
    task = generate_report.delay(report_id, filters)   # non-blocking!
    return {"report_id": report_id, "task_id": task.id}

@app.get("/reports/{report_id}/status")
async def report_status(task_id: str):
    task = app.AsyncResult(task_id)
    return {"status": task.status, "result": task.result}

# Periodic tasks (like cron):
from celery.schedules import crontab

app.conf.beat_schedule = {
    "send-daily-digest": {
        "task": "tasks.send_email",
        "schedule": crontab(hour=8, minute=0),   # every day at 8am
        "args": ("digest@company.com", "Daily Digest", "..."),
    },
}
```
