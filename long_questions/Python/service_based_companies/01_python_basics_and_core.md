# 📘 01 — Python Basics & Core Concepts
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- Data types, variables, type conversion
- Mutability vs immutability
- String operations and formatting
- Operators and expressions
- `None`, truthiness, and identity checks
- Scope: LEGB rule, `global`, `nonlocal`

---

## ❓ Most Asked Questions

### Q1. What are Python's built-in data types?

```python
# Python has several built-in data types:

# Numeric
x_int   = 42          # int
x_float = 3.14        # float
x_cplx  = 2 + 3j      # complex

# Sequential
x_str   = "hello"     # str (immutable)
x_list  = [1, 2, 3]   # list (mutable)
x_tuple = (1, 2, 3)   # tuple (immutable)
x_range = range(10)   # range

# Mapping
x_dict  = {"a": 1}    # dict (mutable, ordered since 3.7)

# Set types
x_set        = {1, 2, 3}        # set (mutable, unordered, unique)
x_frozenset  = frozenset({1, 2}) # frozenset (immutable)

# Boolean
x_bool = True         # bool (subclass of int: True==1, False==0)

# None
x_none = None         # NoneType (singleton)

# Checking types
type(42)              # <class 'int'>
isinstance(42, int)   # True  ← preferred (handles inheritance)
isinstance(True, int) # True  ← bool IS a subclass of int!
```

---

### Q2. What is the difference between mutable and immutable types?

```python
# Immutable: cannot be changed after creation
# Mutable: can be changed in-place

# Immutable: int, float, str, tuple, frozenset, bool, bytes
s = "hello"
# s[0] = "H"  # ❌ TypeError: str object does not support item assignment
s = "Hello"   # ✅ creates a NEW string object; old one may be GC'd

t = (1, 2, 3)
# t[0] = 10   # ❌ TypeError

# Mutable: list, dict, set, bytearray
lst = [1, 2, 3]
lst[0] = 10   # ✅ modifies in-place
lst.append(4) # ✅

# ⚠️ Common trap: mutable default argument
def append_to(item, lst=[]):   # lst is shared across calls!
    lst.append(item)
    return lst

append_to(1)  # [1]
append_to(2)  # [1, 2]  ← unexpected!

# Fix: use None as default
def append_to_safe(item, lst=None):
    if lst is None:
        lst = []
    lst.append(item)
    return lst

# Why mutability matters: when passed to functions
def modify(lst):
    lst.append(99)   # modifies original list

original = [1, 2, 3]
modify(original)
print(original)  # [1, 2, 3, 99]  ← changed!
```

---

### Q3. How does Python handle `==` vs `is`?

```python
# == : checks VALUE equality (calls __eq__)
# is : checks IDENTITY (same object in memory, same id())

a = [1, 2, 3]
b = [1, 2, 3]

a == b   # True  — same values
a is b   # False — different objects in memory

# is with None — always use is, never ==
def check(x):
    if x is None:   # ✅ correct
        return "empty"
    if x == None:   # ❌ avoid — could be overridden by __eq__

# Integer caching (CPython implementation detail)
x = 256
y = 256
x is y   # True — CPython caches small ints (-5 to 256)

x = 257
y = 257
x is y   # False — outside cache range (usually)

# String interning
a = "hello"
b = "hello"
a is b   # True  — short strings are interned by CPython

a = "hello world"
b = "hello world"
a is b   # May be False — not guaranteed for longer strings

# Best practice: use is only for None, True, False singletons
result = None
if result is None:  # ✅
    print("No result")
```

---

### Q4. Explain the LEGB rule and variable scoping in Python.

```python
# LEGB: Local → Enclosing → Global → Built-in

x = "global"  # G

def outer():
    x = "enclosing"  # E

    def inner():
        x = "local"  # L
        print(x)     # "local"

    inner()
    print(x)         # "enclosing"

outer()
print(x)             # "global"

# global keyword: modify a global variable from inside a function
counter = 0

def increment():
    global counter
    counter += 1   # without 'global', would raise UnboundLocalError

increment()
print(counter)  # 1

# nonlocal: modify enclosing scope variable
def make_counter():
    count = 0

    def increment():
        nonlocal count   # refers to make_counter's count
        count += 1
        return count

    return increment

c = make_counter()
c()  # 1
c()  # 2
c()  # 3

# Built-in scope — Python's built-in names
print(len([1, 2, 3]))  # len is in the built-in scope
# If you redefine: len = 10  → shadows the built-in!
```

---

### Q5. Explain f-strings and string formatting techniques.

```python
name = "Priya"
score = 95.678
items = ["apple", "mango"]

# f-strings (Python 3.6+) — fastest, most readable
msg = f"Hello, {name}! Your score is {score:.2f}"
# "Hello, Priya! Your score is 95.68"

# Expressions inside f-strings
print(f"{2 + 2}")                # "4"
print(f"{name.upper()}")         # "PRIYA"
print(f"Items: {', '.join(items)}")  # "Items: apple, mango"

# Alignment and padding
print(f"{'left':<10}|")   # "left      |"
print(f"{'right':>10}|")  # "      right|"
print(f"{'center':^10}|") # "  center  |"
print(f"{42:05d}")         # "00042"
print(f"{3.14:08.3f}")     # "0003.140"

# Debug mode (Python 3.8+): variable = value
x = 42
print(f"{x=}")  # "x=42"

# .format() (older style)
"Hello, {}! Score: {:.2f}".format(name, score)

# %-formatting (oldest, avoid in new code)
"Hello, %s! Score: %.2f" % (name, score)

# String methods
s = "  Hello World  "
s.strip()        # "Hello World"
s.lower()        # "  hello world  "
s.upper()        # "  HELLO WORLD  "
s.split()        # ["Hello", "World"]
" ".join(["a", "b", "c"])  # "a b c"
"hello".replace("l", "r")  # "herro"
"hello".startswith("he")   # True
"world" in "hello world"   # True
```

---

### Q6. What is list comprehension and how does it compare to `map`/`filter`?

```python
numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

# List comprehension: [expression for item in iterable if condition]
squares      = [x**2 for x in numbers]
even_squares = [x**2 for x in numbers if x % 2 == 0]
# [4, 16, 36, 64, 100]

# With else (ternary)
labelled = ["even" if x % 2 == 0 else "odd" for x in numbers]

# Nested comprehension (flatten 2D list)
matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
flat = [item for row in matrix for item in row]
# [1, 2, 3, 4, 5, 6, 7, 8, 9]

# map(): apply function to every element → returns iterator
squares_map = list(map(lambda x: x**2, numbers))

# filter(): keep elements satisfying condition → returns iterator
evens_filter = list(filter(lambda x: x % 2 == 0, numbers))

# Dictionary comprehension
word_lengths = {word: len(word) for word in ["apple", "mango", "kiwi"]}
# {"apple": 5, "mango": 5, "kiwi": 4}

# Set comprehension
unique_chars = {char.lower() for char in "Hello World"}

# Generator expression (lazy, memory-efficient)
gen = (x**2 for x in range(1_000_000))  # no memory allocated yet
next(gen)  # compute one at a time

# ✅ Prefer list comprehension for readability
# ✅ Use generator expressions for large data sets
# ✅ map/filter are fine when using built-in functions
squares_builtin = list(map(pow, numbers, [2]*len(numbers)))
```

---

### Q7. How does `*args` and `**kwargs` work?

```python
# *args: collects positional arguments into a tuple
def greet(*args):
    for name in args:
        print(f"Hello, {name}!")

greet("Alice", "Bob", "Charlie")

# **kwargs: collects keyword arguments into a dict
def display(**kwargs):
    for key, value in kwargs.items():
        print(f"{key}: {value}")

display(name="Rahul", city="Mumbai", score=95)

# Combining: positional, *args, keyword-only, **kwargs
def full_example(a, b, *args, keyword_only="default", **kwargs):
    print(f"a={a}, b={b}")
    print(f"args={args}")           # tuple of extra positional args
    print(f"keyword_only={keyword_only}")
    print(f"kwargs={kwargs}")       # dict of extra keyword args

full_example(1, 2, 3, 4, keyword_only="custom", x=10, y=20)
# a=1, b=2
# args=(3, 4)
# keyword_only=custom
# kwargs={'x': 10, 'y': 20}

# Unpacking with * and **
def add(a, b, c):
    return a + b + c

nums = [1, 2, 3]
add(*nums)   # equivalent to add(1, 2, 3)

config = {"a": 1, "b": 2, "c": 3}
add(**config)  # equivalent to add(a=1, b=2, c=3)

# Forwarding arguments (very common in decorators)
def decorator(func):
    def wrapper(*args, **kwargs):
        print("Before")
        result = func(*args, **kwargs)  # pass everything through
        print("After")
        return result
    return wrapper
```

---

### Q8. What are Python decorators and how do you write one?

```python
import functools

# A decorator is a function that takes a function and returns a new function.

# Simple decorator
def log_call(func):
    @functools.wraps(func)  # preserves __name__, __doc__, etc.
    def wrapper(*args, **kwargs):
        print(f"Calling {func.__name__}")
        result = func(*args, **kwargs)
        print(f"{func.__name__} returned {result}")
        return result
    return wrapper

@log_call
def add(a, b):
    return a + b

add(2, 3)
# Calling add
# add returned 5

# Decorator with arguments
def repeat(n):
    def decorator(func):
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            for _ in range(n):
                result = func(*args, **kwargs)
            return result
        return wrapper
    return decorator

@repeat(3)
def say_hello():
    print("Hello!")

say_hello()  # prints "Hello!" 3 times

# Stacking decorators (applied bottom-up)
@log_call
@repeat(2)
def greet(name):
    print(f"Hi, {name}!")

# Common built-in decorators
class Circle:
    def __init__(self, radius):
        self.radius = radius

    @property        # getter — access like attribute
    def area(self):
        return 3.14 * self.radius ** 2

    @staticmethod    # no self, no cls — pure utility
    def validate_radius(r):
        return r > 0

    @classmethod     # receives class (cls), not instance
    def from_diameter(cls, diameter):
        return cls(diameter / 2)

c = Circle(5)
c.area           # 78.5 — called without ()
Circle.from_diameter(10)   # creates Circle(5)
Circle.validate_radius(3)  # True
```

---

### Q9. What is `None` and how does Python handle truthiness?

```python
# None is a singleton — the only instance of NoneType
x = None
x is None   # True ← correct way to check
x == None   # True but discouraged

# Falsy values in Python:
# None, False, 0, 0.0, 0j, "", [], {}, set(), range(0)

# Everything else is truthy!
bool(None)   # False
bool(0)      # False
bool("")     # False
bool([])     # False
bool({})     # False
bool(0.0)    # False

bool(1)      # True
bool("a")    # True
bool([0])    # True  ← list with one element is truthy!
bool(-1)     # True  ← any non-zero int is truthy

# Practical use of truthiness
def process(data):
    if data:              # checks if data is non-empty/non-None/non-zero
        return handle(data)
    return default_value

# Short-circuit evaluation
name = None
display_name = name or "Anonymous"  # "Anonymous"

user = {"name": "Raj"}
city = user.get("city") or "Unknown"  # "Unknown"

# Walrus operator (Python 3.8+) — assign and test
import re
if m := re.search(r"\d+", "hello123"):
    print(m.group())  # "123"
```

---

### Q10. What is the difference between `deepcopy` and `shallow copy`?

```python
import copy

# Shallow copy: new object, but nested objects are still shared
original = [[1, 2], [3, 4], [5, 6]]

shallow = original.copy()          # list method
shallow = list(original)           # list constructor
shallow = original[:]              # slicing
shallow = copy.copy(original)      # copy module

shallow[0][0] = 99   # changes BOTH original and shallow!
print(original[0])   # [99, 2] ← affected

# Deep copy: recursively copies all nested objects
original = [[1, 2], [3, 4], [5, 6]]
deep = copy.deepcopy(original)

deep[0][0] = 99
print(original[0])   # [1, 2] ← NOT affected
print(deep[0])       # [99, 2]

# Why it matters with dicts
config = {"db": {"host": "localhost", "port": 5432}}
copy1 = config.copy()          # shallow
copy1["db"]["port"] = 9999
print(config["db"]["port"])    # 9999 ← original affected!

config = {"db": {"host": "localhost", "port": 5432}}
copy2 = copy.deepcopy(config)  # deep
copy2["db"]["port"] = 9999
print(config["db"]["port"])    # 5432 ← original safe

# When to use each:
# shallow copy: flat data structures, performance matters
# deep copy: nested data, need true isolation
```
