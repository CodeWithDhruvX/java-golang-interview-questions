# ⚡ 01 — Python Internals & CPython
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- How CPython executes Python code
- Bytecode and the `dis` module
- Reference counting and garbage collection
- The Global Interpreter Lock (GIL)
- Object memory model (`id()`, `is`, interning)
- `__slots__` for memory optimization

---

## ❓ Most Asked Questions

### Q1. How does CPython execute Python code?

```python
# Python execution pipeline:
# Source code (.py) → Tokenizer → AST → Bytecode → CPython VM

# Step 1: Tokenizer breaks source into tokens
import tokenize, io
src = "x = 1 + 2"
tokens = list(tokenize.generate_tokens(io.StringIO(src).readline))
# [NAME 'x', OP '=', NUMBER '1', OP '+', NUMBER '2', ...]

# Step 2: AST (Abstract Syntax Tree)
import ast
tree = ast.parse("x = 1 + 2")
print(ast.dump(tree, indent=2))
# Module(body=[Assign(targets=[Name(id='x')], value=BinOp(...))])

# Step 3: Compile to bytecode
code = compile("x = 1 + 2", "<string>", "exec")
print(code.co_consts)    # (1, 2, None)
print(code.co_varnames)  # ('x',)

# Step 4: Disassemble bytecode into human-readable opcodes
import dis

def add(a, b):
    return a + b

dis.dis(add)
# LOAD_FAST   0  (a)
# LOAD_FAST   1  (b)
# BINARY_OP   0  (+)
# RETURN_VALUE

# .pyc files: compiled bytecode cached in __pycache__/
# Skips tokenizer + compiler on subsequent runs → faster startup

# Python code objects
def mystery(x):
    y = x * 2
    return y + 1

code_obj = mystery.__code__
print(code_obj.co_varnames)  # ('x', 'y')
print(code_obj.co_consts)    # (None, 2, 1)
print(code_obj.co_argcount)  # 1
```

---

### Q2. How does Python's memory management and garbage collection work?

```python
import sys, gc

# Primary mechanism: Reference Counting
# Every object has a reference count. When it hits 0 → freed immediately.

a = [1, 2, 3]
sys.getrefcount(a)  # 2 (one for 'a', one passed to getrefcount)

b = a          # refcount → 3
del b          # refcount → 2
a = None       # refcount → 0 → list is freed!

# Problem: Circular references — ref counting alone can't handle them
class Node:
    def __init__(self, val):
        self.val = val
        self.next = None

n1 = Node(1)
n2 = Node(2)
n1.next = n2
n2.next = n1   # cycle! ref counts never drop to 0

# Solution: Cyclic Garbage Collector (gc module)
# Periodically scans for cyclic garbage using generational GC

gc.collect()         # force collection
gc.get_count()       # (gen0_count, gen1_count, gen2_count)
gc.get_threshold()   # (700, 10, 10) — trigger thresholds

# Generations:
# Gen 0: newly created objects (collected most often)
# Gen 1: survived one Gen 0 collection
# Gen 2: survived Gen 1 collection (long-lived objects: classes, modules)

# Disable GC if you manage cycles manually (for performance):
gc.disable()
# ... no circular refs code ...
gc.enable()

# Memory profiling
sys.getsizeof([])         # 56 bytes (empty list)
sys.getsizeof([1,2,3])    # 88 bytes (3 int references)
sys.getsizeof(1)          # 28 bytes (Python int is boxed!)

# __slots__: reduce memory by eliminating __dict__ per instance
class WithDict:
    def __init__(self, x, y):
        self.x = x
        self.y = y

class WithSlots:
    __slots__ = ("x", "y")   # no __dict__, no __weakref__
    def __init__(self, x, y):
        self.x = x
        self.y = y

# WithDict instance: ~240 bytes, WithSlots: ~64 bytes
# Use __slots__ for classes with millions of instances
```

---

### Q3. What is the GIL — deep dive?

```python
# GIL (Global Interpreter Lock): a mutex in CPython that ensures
# only ONE thread executes Python bytecode at a time.

# Why it exists:
# CPython uses reference counting for memory management.
# Without GIL, concurrent refcount updates would corrupt memory.
# GIL simplifies CPython internals — no per-object lock needed.

# GIL effects:
# ✅ Single-threaded code: no performance impact
# ✅ I/O-bound threads: GIL is RELEASED during I/O — threads work well
# ❌ CPU-bound threads: GIL prevents true parallelism on multi-core

# Demonstrating GIL impact:
import threading, time

def cpu_task():
    count = 0
    for _ in range(50_000_000):
        count += 1

# Single threaded: ~2.5s
start = time.perf_counter()
cpu_task()
cpu_task()
print(f"Sequential: {time.perf_counter() - start:.2f}s")

# Two threads: ALSO ~2.5s (GIL prevents true parallelism)
t1 = threading.Thread(target=cpu_task)
t2 = threading.Thread(target=cpu_task)
start = time.perf_counter()
t1.start(); t2.start()
t1.join(); t2.join()
print(f"Threaded (CPU): {time.perf_counter() - start:.2f}s")  # same!

# Fix: multiprocessing — each process has its own GIL
from multiprocessing import Process
p1 = Process(target=cpu_task)
p2 = Process(target=cpu_task)
start = time.perf_counter()
p1.start(); p2.start()
p1.join(); p2.join()
print(f"Multiprocess: {time.perf_counter() - start:.2f}s")  # ~1.25s ✅

# Python 3.13+ Experimental: Free-threaded Python (--disable-gil flag)
# GIL can be disabled, allowing true multi-threaded parallelism

# Bypassing GIL:
# 1. multiprocessing (separate processes)
# 2. C extensions (NumPy releases GIL for array ops)
# 3. Cython with nogil blocks
# 4. PyPy's STM (Software Transactional Memory)
# 5. asyncio (single-threaded, uses event loop)
```

---

### Q4. What is object interning and how does Python optimize small objects?

```python
# CPython interns (caches) certain objects to save memory and speed up ==

# Integer interning: -5 to 256 are pre-allocated
a = 100; b = 100
a is b   # True  — same cached object

a = 1000; b = 1000
a is b   # False — outside cache range (in separate assignments)
a is b   # May be True if Python optimizes the same code block

# String interning: short, identifier-like strings are interned
s1 = "hello"; s2 = "hello"
s1 is s2   # True — interned

s1 = "hello world"; s2 = "hello world"
s1 is s2   # Usually False (spaces prevent auto-interning)

# Force interning with sys.intern()
import sys
s1 = sys.intern("frequently repeated string")
s2 = sys.intern("frequently repeated string")
s1 is s2   # True — forced interning

# None, True, False are always singletons
None is None   # True
True is True   # True

# Practical impact:
# 1. dict key lookups on interned strings: identity check (is) before hash → faster
# 2. At module level, strings are usually interned automatically

# id() reveals identity
x = 42
print(id(x))   # memory address (integer)
print(id(x) == id(42))   # True inside same expression block

# Object creation cost
import timeit
timeit.timeit("a = []; a.append(1)", number=1_000_000)  # list creation cost
```

---

### Q5. How do `__slots__`, descriptors, and `__dict__` relate to each other?

```python
# __dict__: each instance has a dict storing its attributes (flexible, ~200B overhead)

class Normal:
    def __init__(self, x, y):
        self.x = x
        self.y = y

n = Normal(1, 2)
print(n.__dict__)   # {'x': 1, 'y': 2}
n.z = 99            # can add new attributes dynamically!

# __slots__: replaces __dict__ with fixed-size slot descriptors (memory efficient)
class Slotted:
    __slots__ = ("x", "y")
    def __init__(self, x, y):
        self.x = x
        self.y = y

s = Slotted(1, 2)
# s.__dict__        # AttributeError — no __dict__!
# s.z = 99          # AttributeError — only x, y allowed

# Memory comparison (with 1M instances):
# Normal:  ~240 bytes/instance
# Slotted: ~56 bytes/instance → 75% reduction!

# Descriptor Protocol: objects with __get__, __set__, __delete__
# Used internally for: properties, classmethod, staticmethod, functions

class Validator:
    def __set_name__(self, owner, name):
        self.name = name

    def __get__(self, obj, objtype=None):
        if obj is None:
            return self   # accessed on class
        return obj.__dict__.get(self.name)

    def __set__(self, obj, value):
        if not isinstance(value, int):
            raise TypeError(f"{self.name} must be int, got {type(value).__name__}")
        obj.__dict__[self.name] = value

class Config:
    timeout = Validator()
    retries = Validator()

    def __init__(self, timeout, retries):
        self.timeout = timeout   # triggers Validator.__set__
        self.retries = retries

c = Config(30, 3)
c.timeout = "slow"   # TypeError: timeout must be int
```
