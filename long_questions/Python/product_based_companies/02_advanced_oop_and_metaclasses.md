# ⚡ 02 — Advanced OOP & Metaclasses
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Metaclasses and `type`
- `__new__` vs `__init__`
- Descriptors (`__get__`, `__set__`, `__delete__`)
- Abstract Base Classes (ABCs)
- Multiple inheritance and C3 MRO
- Mixin pattern
- `dataclasses` and `attrs`

---

## ❓ Most Asked Questions

### Q1. What is a metaclass and how do you use it?

```python
# Everything in Python is an object — including classes!
# A metaclass is the "class of a class". Controls how classes are created.

# type is the default metaclass
print(type(int))         # <class 'type'>
print(type(list))        # <class 'type'>
print(type(type))        # <class 'type'>  — type is its own metaclass

# type() with 3 args: creates a class dynamically
MyClass = type("MyClass", (object,), {
    "x": 42,
    "greet": lambda self: f"Hello from {self.__class__.__name__}"
})
obj = MyClass()
print(obj.greet())   # Hello from MyClass

# Custom metaclass
class SingletonMeta(type):
    """Metaclass that enforces Singleton pattern"""
    _instances = {}

    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            cls._instances[cls] = super().__call__(*args, **kwargs)
        return cls._instances[cls]

class Database(metaclass=SingletonMeta):
    def __init__(self, host):
        self.host = host

db1 = Database("localhost")
db2 = Database("remote-server")  # returns SAME instance!
db1 is db2   # True

# Registry metaclass: auto-register all subclasses
class PluginMeta(type):
    registry = {}

    def __new__(mcs, name, bases, namespace):
        cls = super().__new__(mcs, name, bases, namespace)
        if bases:   # don't register base class itself
            mcs.registry[name] = cls
        return cls

class Plugin(metaclass=PluginMeta):
    pass

class CSVPlugin(Plugin):
    pass

class JSONPlugin(Plugin):
    pass

print(PluginMeta.registry)
# {'CSVPlugin': <class 'CSVPlugin'>, 'JSONPlugin': <class 'JSONPlugin'>}

# Auto-instantiate plugins by name:
PluginMeta.registry["CSVPlugin"]()
```

---

### Q2. What is the difference between `__new__` and `__init__`?

```python
# __new__: creates the instance (allocates memory) — called before __init__
# __init__: initializes the already-created instance

class MyClass:
    def __new__(cls, *args, **kwargs):
        print(f"__new__ called — creating instance of {cls.__name__}")
        instance = super().__new__(cls)   # allocate memory
        return instance                   # MUST return instance!

    def __init__(self, value):
        print(f"__init__ called — initializing with {value}")
        self.value = value

obj = MyClass(42)
# __new__ called — creating instance of MyClass
# __init__ called — initializing with 42

# Use __new__ for:
# 1. Singleton pattern
# 2. Immutable types (str, int, tuple — can't change in __init__)
# 3. Metaclass customization

# Singleton via __new__
class Logger:
    _instance = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
            cls._instance.logs = []
        return cls._instance

l1 = Logger()
l2 = Logger()
l1 is l2   # True

# Immutable subclass (must use __new__)
class PositiveInt(int):
    def __new__(cls, value):
        if value <= 0:
            raise ValueError(f"Value must be positive, got {value}")
        return super().__new__(cls, value)

x = PositiveInt(5)    # 5
PositiveInt(-3)       # ValueError

# __init__ on immutable: too late (value already set)
class PositiveIntWrong(int):
    def __init__(self, value):       # this won't work!
        if value <= 0:
            raise ValueError(...)
        super().__init__()  # value already locked in!
```

---

### Q3. Explain Mixins and how they're used in Python.

```python
# Mixin: a class designed to add specific behaviour to other classes
# via multiple inheritance, without being a standalone base class.

# Problem: we want to add JSON serialization + logging to various models
# without duplicating code in each one.

import json
import logging
from datetime import datetime

class JSONMixin:
    """Adds to_json and from_json capabilities"""
    def to_json(self):
        return json.dumps(self.__dict__, default=str)

    @classmethod
    def from_json(cls, json_str):
        data = json.loads(json_str)
        obj = cls.__new__(cls)
        obj.__dict__.update(data)
        return obj

class TimestampMixin:
    """Adds created_at and updated_at tracking"""
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.created_at = datetime.now().isoformat()
        self.updated_at = self.created_at

    def touch(self):
        self.updated_at = datetime.now().isoformat()

class ValidationMixin:
    """Adds validate() method that subclasses override"""
    def validate(self):
        for field, expected_type in getattr(self, "_field_types", {}).items():
            val = getattr(self, field, None)
            if val is not None and not isinstance(val, expected_type):
                raise TypeError(f"{field} must be {expected_type.__name__}")

# Compose mixins with actual model
class User(JSONMixin, TimestampMixin, ValidationMixin):
    _field_types = {"name": str, "age": int}

    def __init__(self, name, age):
        super().__init__()   # TimestampMixin.__init__ called
        self.name = name
        self.age = age

u = User("Rahul", 28)
u.validate()           # ✅ no error
print(u.to_json())     # {"name": "Rahul", "age": 28, ...}
u.touch()              # updates updated_at

# MRO with mixins:
print(User.__mro__)
# User → JSONMixin → TimestampMixin → ValidationMixin → object

# Django uses this pattern extensively:
# class MyView(LoginRequiredMixin, PermissionRequiredMixin, DetailView):
#     pass
```

---

### Q4. How do Abstract Base Classes (ABCs) work?

```python
from abc import ABC, abstractmethod, abstractproperty
from typing import Protocol

# ABC: formally declares an interface — subclasses MUST implement abstract methods

class Storage(ABC):
    @abstractmethod
    def read(self, key: str) -> bytes:
        """Read data by key"""
        ...

    @abstractmethod
    def write(self, key: str, data: bytes) -> None:
        """Write data for key"""
        ...

    @abstractmethod
    def delete(self, key: str) -> bool:
        """Delete key, return True if existed"""
        ...

    def exists(self, key: str) -> bool:
        """Default implementation using read"""
        try:
            self.read(key)
            return True
        except KeyError:
            return False

# Storage()   # TypeError: Can't instantiate abstract class

class LocalStorage(Storage):
    def __init__(self, base_dir):
        self.base_dir = base_dir
        self._store = {}

    def read(self, key):
        if key not in self._store:
            raise KeyError(key)
        return self._store[key]

    def write(self, key, data):
        self._store[key] = data

    def delete(self, key):
        return self._store.pop(key, None) is not None

local = LocalStorage("/tmp")
local.write("config", b"{'debug': True}")
local.exists("config")   # True (uses default impl)

# Protocol (structural subtyping / duck typing)
class Readable(Protocol):
    def read(self, key: str) -> bytes: ...

def fetch_data(storage: Readable, key: str) -> bytes:
    return storage.read(key)    # works with ANY class that has read()!

# No need to inherit — duck typing!
class RedisClient:     # doesn't inherit Readable
    def read(self, key):
        return b"data from redis"

fetch_data(RedisClient(), "key")   # works fine!
```

---

### Q5. How do `dataclasses` work and when should you use them?

```python
from dataclasses import dataclass, field, KW_ONLY, asdict, astuple
from typing import ClassVar

@dataclass
class Point:
    x: float
    y: float

p = Point(3.0, 4.0)
print(p)           # Point(x=3.0, y=4.0)  — __repr__ auto-generated
p1 = Point(3, 4)
p2 = Point(3, 4)
p1 == p2           # True — __eq__ auto-generated!

# Dataclass features
@dataclass(frozen=True, order=True)   # frozen=immutable, order=comparisons
class Version:
    major: int
    minor: int
    patch: int

    def __str__(self):
        return f"{self.major}.{self.minor}.{self.patch}"

v1 = Version(2, 0, 1)
v2 = Version(1, 9, 9)
v1 > v2   # True — order=True generates __lt__, __le__, etc.
# v1.major = 3  # FrozenInstanceError

@dataclass
class Employee:
    name: str
    salary: float
    dept: str = "Engineering"          # default value
    skills: list = field(default_factory=list)  # mutable default!
    _instance_count: ClassVar[int] = 0          # class variable (not in __init__)

    def __post_init__(self):
        # Called after auto-generated __init__
        Employee._instance_count += 1
        if self.salary < 0:
            raise ValueError("Salary cannot be negative")
        self.name = self.name.strip().title()   # normalize name

e1 = Employee("rahul sharma", 80000)
e2 = Employee("priya", 90000, skills=["Python", "Django"])

# Convert to dict or tuple
asdict(e1)    # {"name": "Rahul Sharma", "salary": 80000, "dept": "Engineering", "skills": []}
astuple(e1)   # ("Rahul Sharma", 80000, "Engineering", [])

# dataclass fields inspection
from dataclasses import fields
for f in fields(Employee):
    print(f.name, f.type, f.default)
```
