# 📘 05 — Functional Programming in Python
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- Lambda functions
- `map()`, `filter()`, `reduce()`
- Comprehensions (list, dict, set, generator)
- Generators and `yield`
- `functools`: `partial`, `lru_cache`, `reduce`
- `itertools`: `chain`, `product`, `groupby`, `islice`

---

## ❓ Most Asked Questions

### Q1. What are generators and how do they differ from lists?

```python
# Generator: produces values lazily, one at a time — saves memory!

# List: eagerly creates ALL values in memory
squares_list = [x**2 for x in range(1_000_000)]  # 8 MB in memory

# Generator: creates values on demand
squares_gen = (x**2 for x in range(1_000_000))   # barely any memory!

# Consume with next()
print(next(squares_gen))   # 0
print(next(squares_gen))   # 1
print(next(squares_gen))   # 4

# Or loop through
for val in squares_gen:
    process(val)

# Generator function using yield
def fibonacci():
    a, b = 0, 1
    while True:                # infinite generator!
        yield a
        a, b = b, a + b

fib = fibonacci()
for _ in range(10):
    print(next(fib), end=" ")  # 0 1 1 2 3 5 8 13 21 34

# yield pauses the function; next() resumes from where it left off

# Generator with return value
def countdown(n):
    while n > 0:
        yield n
        n -= 1
    return "Done!"   # StopIteration carries this value

gen = countdown(3)
next(gen)   # 3
next(gen)   # 2
next(gen)   # 1
# next(gen) → StopIteration: Done!

# yield from — delegate to sub-generator
def chain_gens(*iterables):
    for it in iterables:
        yield from it   # yields all items from each iterable

list(chain_gens([1, 2], [3, 4], [5, 6]))  # [1, 2, 3, 4, 5, 6]

# send() — two-way communication
def accumulator():
    total = 0
    while True:
        value = yield total
        if value is None:
            break
        total += value

acc = accumulator()
next(acc)        # prime (must call next first)
acc.send(10)     # 10
acc.send(25)     # 35
acc.send(5)      # 40
```

---

### Q2. Explain `functools.partial`, `lru_cache`, and `reduce`.

```python
from functools import partial, lru_cache, reduce

# partial: fix some arguments of a function → create specialized version
def power(base, exponent):
    return base ** exponent

square  = partial(power, exponent=2)
cube    = partial(power, exponent=3)
double  = partial(power, exponent=1)  # not useful but demonstrates

print(square(5))   # 25
print(cube(3))     # 27

# Real-world: API base URL
import requests
def api_call(method, endpoint, base_url="https://api.example.com", **kwargs):
    return requests.request(method, base_url + endpoint, **kwargs)

get  = partial(api_call, "GET")
post = partial(api_call, "POST")

get("/users")              # GET https://api.example.com/users
post("/login", json={...}) # POST https://api.example.com/login

# lru_cache: memoize function results (Least Recently Used cache)
@lru_cache(maxsize=128)       # cache up to 128 results
def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)

fibonacci(100)   # extremely fast with caching
fibonacci.cache_info()   # CacheInfo(hits=197, misses=101, ...)
fibonacci.cache_clear()  # clear the cache

# reduce: fold a list into a single value
numbers = [1, 2, 3, 4, 5]

total    = reduce(lambda acc, x: acc + x, numbers)  # 15
product  = reduce(lambda acc, x: acc * x, numbers)  # 120
max_val  = reduce(lambda a, b: a if a > b else b, numbers)  # 5

# reduce with initial value
reduce(lambda acc, x: acc + x, numbers, 100)  # 115  (starts at 100)

# Compose functions using reduce
def compose(*fns):
    return reduce(lambda f, g: lambda x: f(g(x)), fns)

add1  = lambda x: x + 1
double = lambda x: x * 2
square = lambda x: x ** 2

# square → double → add1 (right to left)
transform = compose(add1, double, square)
transform(3)  # add1(double(square(3))) = add1(double(9)) = add1(18) = 19
```

---

### Q3. How do you use `itertools` for efficient iteration?

```python
import itertools

# chain: flatten multiple iterables
a = [1, 2, 3]
b = [4, 5, 6]
c = [7, 8, 9]
list(itertools.chain(a, b, c))      # [1,2,3,4,5,6,7,8,9]
list(itertools.chain.from_iterable([a, b, c]))  # same

# product: Cartesian product
colors = ["red", "blue"]
sizes  = ["S", "M", "L"]
list(itertools.product(colors, sizes))
# [('red','S'),('red','M'),('red','L'),('blue','S'),...]

# combinations and permutations
items = [1, 2, 3, 4]
list(itertools.combinations(items, 2))       # choose 2 (no repeat, no order)
# [(1,2),(1,3),(1,4),(2,3),(2,4),(3,4)]

list(itertools.permutations(items, 2))       # choose 2 (no repeat, with order)
# [(1,2),(1,3),(1,4),(2,1),(2,3),...]

list(itertools.combinations_with_replacement(items, 2))  # allow repeats

# groupby: group consecutive elements
data = [
    {"dept": "Eng", "name": "Alice"},
    {"dept": "Eng", "name": "Bob"},
    {"dept": "HR",  "name": "Carol"},
    {"dept": "HR",  "name": "Dave"},
]
# Must sort first!
data.sort(key=lambda x: x["dept"])
for dept, members in itertools.groupby(data, key=lambda x: x["dept"]):
    print(dept, list(members))

# islice: lazy slicing of iterables
def natural_numbers():
    n = 1
    while True:
        yield n
        n += 1

first_10 = list(itertools.islice(natural_numbers(), 10))   # [1..10]
skip_5   = list(itertools.islice(natural_numbers(), 5, 15)) # [6..15]

# takewhile / dropwhile
data = [1, 3, 5, 2, 7, 8, 4]
list(itertools.takewhile(lambda x: x % 2 != 0, data))  # [1, 3, 5]
list(itertools.dropwhile(lambda x: x % 2 != 0, data))  # [2, 7, 8, 4]

# starmap: map with argument unpacking
points = [(2, 3), (4, 5), (1, 7)]
list(itertools.starmap(pow, points))   # [8, 1024, 1]  (x**y for each)

# cycle and repeat
list(itertools.islice(itertools.cycle([1, 2, 3]), 8))  # [1,2,3,1,2,3,1,2]
list(itertools.repeat(10, 5))                           # [10,10,10,10,10]
```

---

### Q4. What is `zip()`, `enumerate()`, and `map()` — practical uses.

```python
# zip: combine multiple iterables element-wise
names  = ["Alice", "Bob", "Charlie"]
scores = [92, 85, 78]
grades = ["A", "B", "C"]

paired = list(zip(names, scores))
# [("Alice", 92), ("Bob", 85), ("Charlie", 78)]

# Iterate in parallel
for name, score in zip(names, scores):
    print(f"{name}: {score}")

# zip with different lengths — stops at shortest
list(zip([1, 2, 3], [4, 5]))   # [(1, 4), (2, 5)]

# zip_longest: fill missing with fillvalue
from itertools import zip_longest
list(zip_longest([1, 2, 3], [4, 5], fillvalue=0))  # [(1,4),(2,5),(3,0)]

# Unzip — transpose back
zipped = [(1, 4), (2, 5), (3, 6)]
a, b = zip(*zipped)   # a=(1,2,3), b=(4,5,6)

# Create dict from two lists
mapping = dict(zip(names, scores))  # {"Alice": 92, "Bob": 85, ...}

# enumerate: loop with index
for i, name in enumerate(names):
    print(f"{i}: {name}")

for i, name in enumerate(names, start=1):   # start index at 1
    print(f"{i}. {name}")

# map: apply function to each element (returns iterator)
nums = [1, 2, 3, 4, 5]
list(map(str, nums))            # ["1", "2", "3", "4", "5"]
list(map(lambda x: x**2, nums)) # [1, 4, 9, 16, 25]

# map with multiple iterables
list(map(pow, [2, 3, 4], [3, 2, 1]))   # [8, 9, 4]  (2**3, 3**2, 4**1)

# filter: keep elements matching condition
evens = list(filter(lambda x: x % 2 == 0, nums))  # [2, 4]
no_none = list(filter(None, [0, 1, None, "", "a", [], [1]]))
# [1, "a", [1]]  — filter(None, ...) removes falsy values
```

---

### Q5. What are closures and how are they used in Python?

```python
# Closure: inner function that remembers the enclosing scope's variables
# even after the outer function has returned.

# Basic closure
def make_multiplier(n):
    def multiply(x):
        return x * n   # 'n' is captured from enclosing scope
    return multiply

double = make_multiplier(2)
triple = make_multiplier(3)

double(5)   # 10
triple(5)   # 15

# Closure retains its own copy of captured variables
print(double.__closure__[0].cell_contents)  # 2
print(triple.__closure__[0].cell_contents)  # 3

# Real-world: configurable validator
def make_range_validator(min_val, max_val):
    def validate(value):
        if not (min_val <= value <= max_val):
            raise ValueError(f"{value} must be between {min_val} and {max_val}")
        return value
    return validate

validate_age    = make_range_validator(0, 150)
validate_score  = make_range_validator(0, 100)
validate_salary = make_range_validator(10000, 10_000_000)

validate_age(25)      # 25 ✅
validate_score(105)   # ValueError!

# Closure for factory functions
def make_logger(prefix):
    import datetime
    def log(message):
        timestamp = datetime.datetime.now().strftime("%H:%M:%S")
        print(f"[{timestamp}] [{prefix}] {message}")
    return log

info  = make_logger("INFO")
error = make_logger("ERROR")
debug = make_logger("DEBUG")

info("Server started")    # [10:30:15] [INFO] Server started
error("Connection lost")  # [10:30:16] [ERROR] Connection lost

# ⚠️ Classic closure pitfall: late binding
fns = [lambda: i for i in range(5)]
[f() for f in fns]   # [4, 4, 4, 4, 4] — all use final i=4!

# Fix: use default argument to capture current value
fns = [lambda i=i: i for i in range(5)]
[f() for f in fns]   # [0, 1, 2, 3, 4] ✅
```
