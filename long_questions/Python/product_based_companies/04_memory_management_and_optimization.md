# ⚡ 04 — Memory Management & Optimization
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Reference counting and cyclic GC
- `gc` module and generational collection
- `sys.getsizeof`, `memory_profiler`, `tracemalloc`
- Object pools and flyweight
- Generator-based memory efficiency
- NumPy memory layout (C vs Fortran order)

---

## ❓ Most Asked Questions

### Q1. How do you profile and reduce Python memory usage?

```python
import sys
import tracemalloc

# tracemalloc: built-in memory tracing (Python 3.4+)
tracemalloc.start()

# Code to profile
data = [i ** 2 for i in range(100_000)]
nested = [[j for j in range(100)] for _ in range(1000)]

snapshot = tracemalloc.take_snapshot()
top_stats = snapshot.statistics("lineno")

print("=== Top 5 memory consumers ===")
for stat in top_stats[:5]:
    print(stat)

# memory_profiler (pip install memory-profiler)
# @profile decorator on functions — tracks per-line memory

# sys.getsizeof — object size in bytes (shallow!)
import sys
lst = [1, 2, 3, 4, 5]
sys.getsizeof(lst)   # 120 bytes (list headers, not elements!)
sys.getsizeof(lst[0])  # 28 bytes per int

# Deep size calculation
def deep_getsizeof(obj, seen=None):
    """Recursively calculate total size"""
    if seen is None:
        seen = set()
    obj_id = id(obj)
    if obj_id in seen:
        return 0
    seen.add(obj_id)
    size = sys.getsizeof(obj)
    if isinstance(obj, dict):
        size += sum(deep_getsizeof(k, seen) + deep_getsizeof(v, seen)
                    for k, v in obj.items())
    elif hasattr(obj, '__iter__') and not isinstance(obj, (str, bytes)):
        size += sum(deep_getsizeof(item, seen) for item in obj)
    return size

# Practical optimizations:
# 1. Generators vs lists for large data
def get_squares_list(n):
    return [x**2 for x in range(n)]   # allocates all at once

def get_squares_gen(n):
    return (x**2 for x in range(n))   # lazy, O(1) memory

# 2. __slots__ for many instances (see internals file)

# 3. Arrays vs lists for numeric data
import array
lst = list(range(1_000_000))       # ~8M bytes
arr = array.array('i', range(1_000_000))  # ~4M bytes (C ints!)

# 4. NumPy arrays for numeric computation
import numpy as np
np_arr = np.arange(1_000_000, dtype=np.int32)  # ~4M bytes
np_arr.nbytes   # exact byte count
```

---

### Q2. How do you use `gc` to handle circular references and memory leaks?

```python
import gc
import weakref

# Scenario: circular reference creating memory leak
class Node:
    def __init__(self, name):
        self.name = name
        self.children = []
        self.parent = None   # circular: parent → child → parent

    def add_child(self, child):
        child.parent = self
        self.children.append(child)

# Create cycle
root = Node("root")
child = Node("child")
root.add_child(child)
# root → child (via children list), child → root (via parent)

# del root doesn't free memory — cyclic GC must handle it
del root, child
gc.collect()   # triggers cyclic GC

# Better solution: use weakref for back-references
class BetterNode:
    def __init__(self, name):
        self.name = name
        self.children = []
        self._parent = None   # stored as weakref

    @property
    def parent(self):
        if self._parent is not None:
            return self._parent()   # dereference weakref
        return None

    @parent.setter
    def parent(self, node):
        self._parent = weakref.ref(node) if node else None

    def add_child(self, child):
        child.parent = self
        self.children.append(child)

root = BetterNode("root")
child = BetterNode("child")
root.add_child(child)
# No cycle! child._parent is a weakref — reference count not affected

# GC settings
gc.get_threshold()        # (700, 10, 10) — default thresholds
gc.set_threshold(1000, 15, 15)  # tune for your app
gc.disable()   # disable if you manually manage cycles (performance gain)

# Detect unreachable objects
gc.collect()
print(f"Unfreachable: {gc.garbage}")   # objects with __del__ and in cycle

# Context: why disable GC?
# - Applications with no circular refs (pure functional style)
# - High-frequency short-lived object creation (web request handlers)
# Instagram famously disabled GC for ~10% CPU reduction
```

---

### Q3. What are memory-efficient patterns for large data processing?

```python
import csv
import json
from pathlib import Path

# 1. Stream large files — don't load all into memory
def process_large_csv(filename):
    with open(filename, newline='') as f:
        reader = csv.DictReader(f)
        for row in reader:          # one row at a time!
            yield process_row(row)  # generator pipeline

def process_row(row):
    return {k: v.strip() for k, v in row.items()}

# 2. Generator pipeline — compose transformations lazily
def read_lines(path):
    with open(path) as f:
        yield from f               # stream lines

def parse_json_lines(lines):
    for line in lines:
        try:
            yield json.loads(line)
        except json.JSONDecodeError:
            pass

def filter_active(records):
    for r in records:
        if r.get("active", False):
            yield r

def extract_fields(records, fields):
    for r in records:
        yield {k: r[k] for k in fields if k in r}

# Compose: reads 1 line at a time, memory = O(1)
def pipeline(path):
    lines = read_lines(path)
    records = parse_json_lines(lines)
    active = filter_active(records)
    minimal = extract_fields(active, ["id", "name", "email"])
    return minimal

# 3. Chunked processing with itertools.islice
import itertools

def process_in_chunks(iterable, chunk_size=1000):
    iterator = iter(iterable)
    while True:
        chunk = list(itertools.islice(iterator, chunk_size))
        if not chunk:
            break
        yield chunk

# 4. Reuse objects with object pools
class ConnectionPool:
    def __init__(self, factory, size=10):
        self._pool = [factory() for _ in range(size)]
        self._available = list(range(size))
        import threading
        self._lock = threading.Lock()

    def acquire(self):
        with self._lock:
            if self._available:
                idx = self._available.pop()
                return idx, self._pool[idx]
            raise RuntimeError("Pool exhausted")

    def release(self, idx):
        with self._lock:
            self._available.append(idx)
```
