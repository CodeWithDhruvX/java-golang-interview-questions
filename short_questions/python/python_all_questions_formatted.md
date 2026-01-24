# Python Interview Questions & Answers

## 🔹 1. General & Basics (Questions 1-10)

**Q1: What is Python and what are its key features?**
Python is a high-level, interpreted, general-purpose programming language known for its readability and simplicity. Key features include:
*   **Interpreted:** Code is executed line-by-line.
*   **Dynamically Typed:** Type checking happens at runtime.
*   **Object-Oriented:** Everything is an object.
*   **Extensive Standard Library:** "Batteries included" philosophy.
*   **Platform Independent:** Runs on Windows, Mac, Linux, etc.

**Q2: How is Python an interpreted language?**
Python code (.py) is compiled into bytecode (.pyc) which is then executed by the Python Virtual Machine (PVM). Unlike C/C++, it is not compiled directly to machine code before execution.

**Q3: What is the difference between Python 2 and Python 3?**
*   **Print:** `print "hello"` (Py2) vs `print("hello")` (Py3).
*   **Integer Division:** `5/2 = 2` (Py2) vs `5/2 = 2.5` (Py3).
*   **Unicode:** Strings are ASCII by default in Py2, Unicode in Py3.
*   **Range:** `xrange()` (Py2) vs `range()` (Py3).

**Q4: What is PEP 8 and why is it important?**
PEP 8 (Python Enhancement Proposal 8) is the style guide for Python code. It ensures code consistency and readability (e.g., using 4 spaces for indentation, snake_case for functions).

**Q5: How is memory managed in Python?**
Python uses a private heap to manage memory.
*   **Memory Manager:** Allocates heap space for objects.
*   **Garbage Collector:** Recycles unused memory using reference counting and a cycle detector.

**Q6: What are Python namespaces?**
A namespace is a mapping from names to objects (like a dictionary). Examples include:
*   **Built-in:** `print()`, `id()`
*   **Global:** Defined at the module level.
*   **Local:** Defined inside a function.

**Q7: What are local and global variables in Python?**
*   **Local:** Declared inside a function; accessible only within that function.
*   **Global:** Declared outside any function; accessible throughout the module. Use the `global` keyword to modify them inside a function.

**Q8: What is the difference between lists and tuples?**
*   **List:** Mutable (`[]`), slower, consumes more memory.
*   **Tuple:** Immutable (`()`), faster, consumes less memory.

**Q9: What are the built-in data types in Python?**
*   **Numeric:** `int`, `float`, `complex`
*   **Sequence:** `list`, `tuple`, `range`, `str`
*   **Mapping:** `dict`
*   **Set:** `set`, `frozenset`
*   **Boolean:** `bool`
*   **Binary:** `bytes`, `bytearray`

**Q10: What is type conversion in Python?**
The process of converting one data type to another. 
*   **Implicit:** Automatic (e.g., `int` + `float` -> `float`).
*   **Explicit:** Manual using functions like `int()`, `str()`, `list()`.

---

## 🔹 2. Data Types & Variables (Questions 11-20)

**Q11: What is the difference between `is` and `==`?**
*   `==` checks for **value equality** (do they hold the same data?).
*   `is` checks for **reference equality** (do they point to the exact same object in memory?).

**Q12: What are mutable and immutable data types?**
*   **Mutable:** Can be changed after creation (List, Dictionary, Set).
*   **Immutable:** Cannot be changed after creation (Int, Float, String, Tuple, Frozenset).

**Q13: How do you comment code in Python?**
*   **Single-line:** Using `#`.
*   **Multi-line:** Using triple quotes `''' ... '''` or `""" ... """` (technically multiline strings, often used as comments).

**Q14: What are docstrings in Python?**
Documentation strings used to explain modules, classes, or functions. They are enclosed in triple quotes and stored in the `__doc__` attribute.
```python
def my_func():
    """This is a docstring."""
    pass
```

**Q15: What is the `pass` statement?**
A null statement used as a placeholder. It does nothing when executed but avoids syntax errors in empty blocks (loops, functions, classes).

**Q16: What is the `break`, `continue`, and `pass` difference?**
*   `break`: Exits the loop immediately.
*   `continue`: Skips the rest of the current iteration and moves to the next.
*   `pass`: Does nothing (placeholder).

**Q17: What is slicing in Python?**
Extracting a portion of a sequence (string, list, tuple) using the syntax `[start:stop:step]`.
```python
lst = [1, 2, 3, 4]
print(lst[1:3]) # Output: [2, 3]
```

**Q18: How do you reverse a string in Python?**
Using slicing: `string[::-1]`.
```python
s = "hello"
print(s[::-1]) # "olleh"
```

**Q19: What are negative indexes and why are they used?**
They allow accessing elements from the end of a sequence. `-1` is the last element, `-2` is the second to last.

**Q20: What is a dictionary in Python?**
An unordered, mutable collection of key-value pairs (`{key: value}`). Keys must be unique and immutable.

---

## 🔹 3. Lists & Dictionaries (Questions 21-30)

**Q21: How do you remove duplicates from a list?**
Convert it to a set and back to a list (order may be lost).
```python
lst = [1, 2, 2, 3]
lst = list(set(lst)) # [1, 2, 3]
```

**Q22: What is list comprehension? Give an example.**
A concise way to create lists.
```python
sq = [x**2 for x in range(5)] # [0, 1, 4, 9, 16]
```

**Q23: What is dictionary comprehension?**
Similar to list comprehension but for dictionaries.
```python
sq_dict = {x: x**2 for x in range(5)} # {0:0, 1:1, ...}
```

**Q24: How do you merge two dictionaries?**
*   **Python 3.9+:** `d1 | d2`
*   **Older:** `{**d1, **d2}` or `d1.update(d2)`

**Q25: What is the difference between `append()` and `extend()`?**
*   `append()`: Adds the element as a single item at the end.
*   `extend()`: Iterates over the element and adds each item to the list.

**Q26: What is a lambda function?**
An anonymous, single-line function defined using the `lambda` keyword.
```python
add = lambda x, y: x + y
print(add(2, 3)) # 5
```

**Q27: What is the difference between `remove()`, `del`, and `pop()`?**
*   `remove(val)`: Removes the first occurrence of a value.
*   `del list[i]`: Deletes the item at index `i`.
*   `pop(i)`: Removes and *returns* the item at index `i` (default last).

**Q28: How do you sort a dictionary by value?**
Using `sorted()` with a lambda key.
```python
d = {'a': 2, 'b': 1}
d_sorted = dict(sorted(d.items(), key=lambda item: item[1]))
```

**Q29: What is the purpose of `zip()` function?**
Aggregates elements from two or more iterables into tuples (paired by index).
```python
names = ['A', 'B']
ages = [20, 30]
print(list(zip(names, ages))) # [('A', 20), ('B', 30)]
```

**Q30: What is the `enumerate()` function?**
Adds a counter to an iterable and returns it as an enumerate object (index, value).
```python
for i, val in enumerate(['a', 'b']):
    print(i, val)
```

---

## 🔹 4. Sets & Advanced Data Structures (Questions 31-40)

**Q31: How does `set` work internally in Python?**
It uses a hash table. The elements acts as keys in a dictionary with dummy values. This allows O(1) average time complexity for lookups.

**Q32: What is the difference between a set and a frozenset?**
*   **set:** Mutable (can add/remove items).
*   **frozenset:** Immutable (cannot change after creation), hashable (can be a dict key).

**Q33: How do you count the occurrences of items in a list?**
Using `collections.Counter`.
```python
from collections import Counter
c = Counter(['a', 'a', 'b']) # Counter({'a': 2, 'b': 1})
```

**Q34: What is the usage of `*args` and `**kwargs`?**
*   `*args`: Passes variable number of positional arguments (tuple).
*   `**kwargs`: Passes variable number of keyword arguments (dictionary).

**Q35: How do you deep copy an object in Python?**
Using the `copy` module.
```python
import copy
new_list = copy.deepcopy(old_list)
```

**Q36: What is the difference between shallow copy and deep copy?**
*   **Shallow Copy:** Creates a new object but inserts references to the original child objects.
*   **Deep Copy:** Creates a new object and recursively copies all child objects (independent copy).

**Q37: How can you implement a stack and queue in Python?**
*   **Stack (LIFO):** Use `list` with `append()` and `pop()`.
*   **Queue (FIFO):** Use `collections.deque` with `append()` and `popleft()`.

**Q38: What is the usage of `map()`, `filter()`, and `reduce()`?**
*   `map(func, iter)`: Applies function to all items.
*   `filter(func, iter)`: Keeps items where function returns True.
*   `reduce(func, iter)`: Reduces iterable to a single value (needs `functools`).

**Q39: How do you find the intersection of two sets?**
Using `&` operator or `intersection()` method.
```python
s1 = {1, 2}; s2 = {2, 3}
print(s1 & s2) # {2}
```

**Q40: How do you use the `collections` module? (Counter, defaultdict)**
*   `Counter`: Counts hashable objects.
*   `defaultdict`: Dictionary that calls a factory function to supply missing values.
*   `namedtuple`: Tuple subclass with named fields.

---

## 🔹 5. Object-Oriented Programming (OOP) (Questions 41-50)

**Q41: What is a Class in Python?**
A blueprint for creating objects. It defines the properties (attributes) and behaviors (methods) that the objects created from the class will have.

**Q42: What is `self` in Python?**
A reference to the current instance of the class. It is used to access variables and methods associated with the current object.

**Q43: What is the `__init__` method?**
The constructor method in Python. It is automatically called when a new object exists to initialize its attributes.

**Q44: How do you handle inheritance in Python?**
By passing the parent class as an argument to the child class definition.
```python
class Child(Parent):
    pass
```

**Q45: What is Multiple Inheritance?**
A feature where a class can inherit attributes and methods from more than one parent class. Python supports this.

**Q46: What is hierarchical inheritance?**
When multiple child classes inherit from a single parent class.

**Q47: How does Python handle Method Overloading?**
Python does **not** support method overloading (same method name, different parameters) by default. If defined twice, the latest one overwrites the previous. It is achieved using variable arguments or default values.

**Q48: What is Method Overriding?**
When a child class provides a specific implementation of a method that is already defined in its parent class.

**Q49: What is Encapsulation in Python?**
Bundling data (variables) and methods together and restricting direct access. In Python, it's a convention (using `_` or `__` prefix) rather than strict enforcement.

**Q50: How do you make a variable private in Python?**
Prefix the variable name with double underscores `__`. Python performs name mangling (changes it to `_ClassName__variable`) to make it harder (but not impossible) to access from outside.

---

## 🔹 6. OOP - Advanced (Questions 51-60)

**Q51: What is Data Abstraction?**
Hiding the implementation details and showing only the necessary functionality to the user. In Python, this is achieved using Abstract Base Classes (ABCs).

**Q52: What are Python Mixins?**
A class used to provide specific functionality to other classes through multiple inheritance, but not intended to stand on its own (e.g., logging mixin).

**Q53: What is the difference between a class method and a static method?**
*   **@classmethod:** Takes `cls` as first argument. Can modify class state.
*   **@staticmethod:** Takes no implicit first argument. behaves like a regular function but belongs to the class namespace.

**Q54: What is the `@property` decorator?**
Allows you to define a method that can be accessed like an attribute (getter), often used to compute values or validate setting values (setter).

**Q55: What is the MRO (Method Resolution Order)?**
The order in which Python searches for a method in a hierarchy of classes. Python uses the **C3 Linearization** algorithm. You can check it via `ClassName.__mro__`.

**Q56: What are "magic methods" or "dunder methods"?**
Special methods with double underscores (e.g., `__init__`, `__len__`, `__add__`) that allow you to define how objects behave with built-in operations (operator overloading).

**Q57: What is the difference between `__str__` and `__repr__`?**
*   `__str__`: Readable string for end-users (`print()`).
*   `__repr__`: Unambiguous string for developers (debugging/console).

**Q58: What is inheritance vs composition?**
*   **Inheritance:** "Is-a" relationship (Dog is an Animal). Code reuse via subclassing.
*   **Composition:** "Has-a" relationship (Car has an Engine). Code reuse by containing objects of other classes.

**Q59: What are abstract base classes (ABC)?**
Classes that cannot be instantiated and are used to define a common interface for subclasses. Defined using `abc` module.

**Q60: How do you create an abstract class?**
Inherit from `ABC` and use `@abstractmethod`.
```python
from abc import ABC, abstractmethod
class Shape(ABC):
    @abstractmethod
    def area(self): pass
```

---

## 🔹 7. Functions & Generators (Questions 61-70)

**Q61: What is a recursive function?**
A function that calls itself to solve a problem by breaking it down into smaller sub-problems (e.g., factorial, fibonacci).

**Q62: What are default arguments in functions?**
Parameters that assume a default value if no argument is provided during the function call.
*   **Warning:** Do not use mutable default arguments (like lists).

**Q63: What is the scope of variables in Python?**
The region where a variable is recognized. LEGB Rule: **L**ocal, **E**nclosing, **G**lobal, **B**uilt-in.

**Q64: What is a closure in Python?**
A nested function that captures and remembers values from its enclosing scope even after the outer function has finished execution.

**Q65: What is a decorator? Explain with an example.**
A function that takes another function and extends its behavior without modifying it explicitly.
```python
def my_decorator(func):
    def wrapper():
        print("Before")
        func()
        print("After")
    return wrapper
```

**Q66: What are generators in Python?**
Functions that return an iterable set of items, one at a time, using the `yield` keyword. They are memory efficient.

**Q67: What is the `yield` keyword?**
Used in a function to pause execution and produce a value to the caller, saving the state for the next call.

**Q68: What is the difference between `yield` and `return`?**
*   `return`: Returns a value and terminates the function.
*   `yield`: Returns a value and pauses the function, resuming from there next time.

**Q69: What involves in Python package vs module?**
*   **Module:** A single `.py` file.
*   **Package:** A directory containing multiple modules and an `__init__.py` file.

**Q70: How do you create a Python module?**
Simply create a file with `.py` extension (e.g., `mymodule.py`) and write functions/classes in it. Import it using `import mymodule`.

---

## 🔹 8. Modules & Exceptions (Questions 71-80)

**Q71: What is the `__name__` variable?**
A built-in variable that evaluates to the name of the current module. If run directly, it is `"__main__"`.

**Q72: What does `if __name__ == "__main__":` do?**
Ensures that the block of code inside it runs only if the script is executed directly, not when imported as a module.

**Q73: How do you install external packages?**
Using `pip`, the Python package installer.
`pip install package-name`

**Q74: What is PYTHONPATH?**
An environment variable that adds additional directories where Python looks for modules and packages.

**Q75: What are iterators in Python?**
Objects that implement the iterator protocol (`__iter__` and `__next__`). They allow you to traverse through all the elements of a collection.

**Q76: What is difference between iterator and iterable?**
*   **Iterable:** An object you can loop over (has `__iter__`, e.g., list).
*   **Iterator:** The object that keeps state of traversal (has `__next__`).

**Q77: How to create a custom exception?**
Create a class that inherits from `Exception`.
```python
class MyError(Exception):
    pass
```

**Q78: What are `try`, `except`, `else` and `finally` blocks?**
*   `try`: Code that might raise exception.
*   `except`: Handle the exception.
*   `else`: Run if no exception occurred.
*   `finally`: Always run (cleanup).

**Q79: How do you handle multiple exceptions?**
By specifying multiple `except` blocks or a tuple of exceptions.
```python
except (TypeError, ValueError):
    pass
```

**Q80: What is raising an exception?**
Manually triggering an error using the `raise` keyword.
`raise ValueError("Invalid value")`

---

## 🔹 9. File Handling (Questions 81-90)

**Q81: How do you open and close a file in Python?**
Using `open()` function.
`f = open('file.txt'); ...; f.close()`

**Q82: What are the different file modes in Python?**
*   `'r'`: Read (default).
*   `'w'`: Write (overwrites).
*   `'a'`: Append.
*   `'b'`: Binary mode (e.g., `'rb'`).
*   `'+'`: Read and Write.

**Q83: How do you read a file line by line?**
Using a loop or `readline()`.
```python
with open('file.txt') as f:
    for line in f:
        print(line)
```

**Q84: What is the `with` statement used for in file handling?**
It is a context manager that ensures the file is automatically closed after the block ends, even if an exception occurs.

**Q85: How do you write data to a JSON file?**
Using the `json` module.
```python
import json
json.dump(data, open('file.json', 'w'))
```

**Q86: How to check if a file exists?**
Using `os.path.exists()` or `pathlib.Path.exists()`.

**Q87: What is pickling and unpickling?**
*   **Pickling:** Serializing a Python object structure into a byte stream (`pickle.dump`).
*   **Unpickling:** Deserializing the byte stream back into an object (`pickle.load`).

**Q88: How to rename a file in Python?**
`os.rename('old.txt', 'new.txt')`

**Q89: How to delete a file in Python?**
`os.remove('file.txt')`

**Q90: How to list all files in a directory?**
`os.listdir()` or `os.walk()` or `glob.glob('*')`.

---

## 🔹 10. Testing & Debugging (Questions 91-100)

**Q91: What is unit testing in Python?**
Testing individual units of source code (functions, methods) to determine they are fit for use.

**Q92: What is the `unittest` module?**
Python's built-in standard library for creating unit tests. Based on JUnit.

**Q93: What is `pytest`?**
A popular third-party testing framework that is more Pythonic, less boilerplate, and supports fixtures/plugins.

**Q94: How do you debug a Python program?**
Using `print()` statements, logging, IDE debuggers, or the `pdb` module.

**Q95: What is `pdb`?**
 The Python Debugger. You can invoke it with `import pdb; pdb.set_trace()` (or `breakpoint()` in Py3.7+) to step through code.

**Q96: How do you profile Python code?**
Using `cProfile` module to measure where time is being spent.
`python -m cProfile script.py`

**Q97: What is usage of assertions?**
Debug tools used to check if a condition is True. If False, raises `AssertionError`.
`assert x > 0, "x must be positive"`

**Q98: How do you mock data in tests?**
Using `unittest.mock` to replace real objects with mock objects (simulated behavior) during testing, useful for API calls or DB access.

**Q99: What are fixtures in testing?**
Setup and teardown code (e.g., creating a temp DB) that runs before/after tests. Common in `pytest`.

**Q100: How do you measure execution time of a code block?**
Using `time` module or context managers.
```python
import time
start = time.time()
# code...
print(time.time() - start)
```

---

## 🔹 11. Concurrency & Performance (Questions 101-110)

**Q101: What is the Global Interpreter Lock (GIL)?**
A mutex that protects access to Python objects, preventing multiple threads from executing Python bytecodes at once. This effectively limits Python programs to running on a single processor core.

**Q102: How does GIL impact multithreading?**
It prevents true parallelism in CPU-bound tasks. However, it does not affect I/O-bound tasks significantly, as the lock is released during I/O operations.

**Q103: What is the difference between multithreading and multiprocessing in Python?**
*   **Multithreading:** Uses threads (shared memory), affected by GIL. Best for I/O-bound tasks.
*   **Multiprocessing:** Uses processes (separate memory), bypasses GIL. Best for CPU-bound tasks.

**Q104: How does garbage collection work in Python?**
Primarily uses Reference Counting. When an object's reference count drops to zero, it is deallocated. It also uses a Cyclic Garbage Collector to detect and clean up reference cycles.

**Q105: What is Reference Counting?**
A memory management technique where each object keeps a count of how many references point to it. Returns memory when count hits zero.

**Q106: What are circular references and how are they handled?**
When two objects reference each other, creating a cycle. The Reference Counter cannot handle this (count never hits zero). The Cyclic GC periodically scans for and collects these.

**Q107: What are metaclasses?**
Classes of classes. They define how classes behave. A class is an instance of a metaclass (default is `type`). Used for API enforcement or code injection.

**Q108: What is the usage of `__new__` vs `__init__`?**
*   `__new__`: A static method that creates the instance. Rarely used (mostly for immutable subclasses like tuple).
*   `__init__`: An instance method that initializes the already created instance.

**Q109: What are slots (`__slots__`) in Python classes?**
A mechanism to restrict the dynamic creation of attributes. It saves memory by preventing the creation of `__dict__` for each instance.
```python
class Point:
    __slots__ = ['x', 'y']
```

**Q110: What is Monkey Patching?**
The practice of dynamically modifying or extending modules or classes at runtime (e.g., replacing a method during testing).

---

## 🔹 12. Advanced Python Concepts (Questions 111-120)

**Q111: What is Duck Typing?**
"If it walks like a duck and quacks like a duck, it is a duck." Python assumes objects are compatible if they have the required methods, regardless of their type.

**Q112: How do you implement a Singleton pattern in Python?**
By overriding `__new__` or using a decorator/metaclass to ensure only one instance of the class exists.
```python
class Singleton:
    _instance = None
    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
        return cls._instance
```

**Q113: What are context managers and how to create a custom one?**
Objects used with `with` statement. Implement `__enter__` and `__exit__`.
```python
class MyContext:
    def __enter__(self): print("Enter"); return self
    def __exit__(self, type, value, tb): print("Exit")
```

**Q114: What are descriptors in Python?**
Objects that define how attributes are accessed/set (implement `__get__`, `__set__`, `__delete__`). `@property` is a built-in descriptor.

**Q115: What is the difference between `.py` and `.pyc` files?**
*   `.py`: Human-readable source code.
*   `.pyc`: Compiled bytecode (platform-independent) created by the Python interpreter for faster loading.

**Q116: How does Python import system work?**
It searches for the module in `sys.modules`, then built-ins, then directories in `sys.path`. If found, it executes the module code.

**Q117: What is introspection in Python?**
The ability of a program to examine the type or properties of an object at runtime (e.g., `type()`, `dir()`, `id()`, `hasattr()`).

**Q118: What is the `dis` module used for?**
The Disassembler module. It converts Python bytecode into a human-readable format, useful for performance analysis.

**Q119: How do you optimize Python code performance?**
Use built-in functions, list comprehensions (faster than loops), `map`, generators, or compile critical parts with Cython/NumPy.

**Q120: What is Cython?**
A programming language that makes writing C extensions for Python as easy as Python itself. Used for high-performance applications.

---

## 🔹 13. Web Frameworks (Django & Flask) (Questions 121-130)

**Q121: What is WSGI?**
Web Server Gateway Interface. A standard interface between web servers and Python web applications (e.g., Gunicorn -> Flask).

**Q122: What is ASGI?**
Asynchronous Server Gateway Interface. The spiritual successor to WSGI, supporting async/await applications (handling WebSockets, long polling).

**Q123: Difference between Django and Flask?**
*   **Django:** "Batteries-included" (built-in ORM, Admin, Auth), monolithic, opinionated.
*   **Flask:** Micro-framework (minimal core), flexible, requires extensions for ORM/Auth.

**Q124: What are Django Signals?**
A dispatcher system that allows decoupled applications to get notified when certain actions occur (e.g., `post_save` signal after a model is saved).

**Q125: What is Django ORM?**
Object-Relational Mapping. Allows interacting with the database using Python objects instead of raw SQL queries (e.g., `User.objects.all()`).

**Q126: Explain the MVT architecture in Django.**
*   **Model:** Data structure (Database).
*   **View:** Business logic (Controllers in MVC).
*   **Template:** Presentation layer (HTML).

**Q127: What is Middleware in Django?**
Hooks into Django’s request/response processing cycle. Used for global functionality like Authentication, CSRF protection, logging.

**Q128: How do you manage database migrations in Django?**
*   `makemigrations`: Creates new migration files based on model changes.
*   `migrate`: Applies the changes to the database.

**Q129: What is a Blueprint in Flask?**
A way to organize a Flask application into modules or components. useful for large applications.

**Q130: How do you handle sessions in Flask?**
Using `session` object (dictionary-like). It signs the cookie cryptographically to store data on the client side securely.

---

## 🔹 14. REST & Security (Questions 131-140)

**Q131: What is a RESTful API?**
Representational State Transfer. An architectural style for designing networked applications. Uses HTTP methods (GET, POST, PUT, DELETE) to manipulate resources.

**Q132: How do you secure a Python web application?**
Input validation, SQL injection prevention (ORM), CSRF protection, XSS headers, using HTTPS, and secure password hashing (bcrypt/Argon2).

**Q133: What is CSRF and how to prevent it?**
Cross-Site Request Forgery. Prevented using CSRF tokens (unique secret included in forms) which the server validates.

**Q134: What are template engines (Jinja2)?**
Libraries that combine templates (HTML with placeholders) with data to produce the final document. Jinja2 is default for Flask.

**Q135: How do you implement authentication in Flask?**
Using extensions like `Flask-Login` (session-based) or `Flask-JWT-Extended` (token-based).

**Q136: What is Gunicorn?**
Green Unicorn. A production-grade WSGI HTTP Server for UNIX. It sits behind Nginx and serves the Python app.

**Q137: How to deploy a Python web app?**
Common stack: [Client] -> [Nginx (Reverse Proxy)] -> [Gunicorn (WSGI Server)] -> [Flask/Django App].

**Q138: What is GraphQL integration in Python?**
Using libraries like **Graphene** or **Ariadne** to build GraphQL schemas and resolvers in Python.

**Q139: How do you handle CORS in Python APIs?**
Using middleware (e.g., `django-cors-headers` or `Flask-CORS`) to add `Access-Control-Allow-Origin` headers.

**Q140: What is Celery and why is it used?**
A distributed task queue for handling asynchronous/background jobs (sending emails, heavy processing) using a message broker (RabbitMQ/Redis).

---

## 🔹 15. Data Science & Pandas (Questions 141-150)

**Q141: What is the difference between methods in list and NumPy array?**
NumPy arrays are homogeneous (same type), fixed-size, and support vectorized operations (faster). Lists are heterogeneous and dynamic.

**Q142: Why is NumPy faster than lists?**
It uses contiguous memory blocks (locality of reference) and C-optimized algorithms, avoiding Python's type checking overhead during iteration.

**Q143: What is a DataFrame in Pandas?**
A 2-dimensional labeled data structure (like a SQL table or Excel sheet). Columns can be different types.

**Q144: How do you handle missing data in Pandas?**
*   `dropna()`: Drop rows/cols with nulls.
*   `fillna(value)`: Fill nulls with a value (mean, median, 0).
*   `isna()`: Check for nulls.

**Q145: What is broadcasting in NumPy?**
A mechanism that allows arithmetic operations on arrays of different shapes (e.g., adding a scalar `5` to a matrix).

**Q146: How do you merge Datarames in Pandas?**
Using `pd.merge(df1, df2, on='key', how='inner')`. similar to SQL joins.

**Q147: What is `groupby` in Pandas?**
Used to split data into groups based on some criteria, apply a function (sum, mean), and combine results. `df.groupby('col').mean()`.

**Q148: How do you read a CSV file with Pandas?**
`df = pd.read_csv('file.csv')`.

**Q149: What is `loc` vs `iloc`?**
*   `loc`: Label-based indexing (`df.loc['row_label', 'col_label']`).
*   `iloc`: Integer-position based indexing (`df.iloc[0, 1]`).

**Q150: What is data cleaning?**
The process of detecting and correcting (or removing) corrupt or inaccurate records from a dataset (handling nulls, duplicates, bad formats).

---

## 🔹 16. Data Analysis & Tools (Questions 151-160)

**Q151: What is the difference between Series and DataFrame?**
*   **Series:** 1D labeled array (like a column).
*   **DataFrame:** 2D data structure (collection of Series).

**Q152: How do you filter rows in a DataFrame?**
Using boolean indexing.
`filtered_df = df[df['age'] > 25]`

**Q153: What is Matplotlib used for?**
A comprehensive library for creating static, animated, and interactive visualizations (plots, charts) in Python.

**Q154: What is Scikit-learn?**
A robust library for machine learning (classification, regression, clustering) built on NumPy, SciPy, and Matplotlib.

**Q155: How do you prevent overfitting?**
Cross-validation, regularization (L1/L2), pruning (trees), early stopping, or using more training data.

**Q156: What is a virtual environment and why use it?**
An isolated environment for a Python project. It ensures project dependencies (versions) don't conflict with system-wide packages or other projects.

**Q157: What is the difference between `pip` and `conda`?**
*   `pip`: Python package installer (PyPI).
*   `conda`: Cross-language package and environment manager (handles binary libraries especially for Data Science).

**Q158: What is `requirements.txt`?**
A file listing all dependencies for a project.
`pip install -r requirements.txt`

**Q159: How do you create a virtual environment?**
`python -m venv myenv`

**Q160: What is data serialization?**
Converting structured data (objects) into a format that can be stored or transmitted (JSON, XML, Pickle).

---

## 🔹 17. Modern Python Features (Questions 161-170)

**Q161: What are f-strings?**
Formatted string literals (Python 3.6+). Concise and fast.
`name="John"; print(f"Hello {name}")`

**Q162: What is type hinting?**
Adding type annotations to function arguments and return values (Python 3.5+). Checked by static tools like `mypy` (not enforced at runtime).
`def add(x: int, y: int) -> int:`

**Q163: What are Data Classes (`@dataclass`)?**
A decorator (Python 3.7+) to automatically generate special methods like `__init__`, `__repr__`, `__eq__` for classes that primarily store data.

**Q164: What is the walrus operator (`:=`)?**
Assignment expression (Python 3.8+). Allows assigning a value to a variable as part of a larger expression.
`if (n := len(data)) > 10: ...`

**Q165: What is `asyncio`?**
A library to write concurrent code using the `async`/`await` syntax. Used for high-performance I/O-bound network/web servers.

**Q166: What are `async` and `await` keywords?**
*   `async def`: Defines a coroutine.
*   `await`: Pauses execution of the coroutine until the awaitable completes.

**Q167: What is an Event Loop?**
The core of `asyncio`. It runs asynchronous tasks and callbacks, performing network I/O operations, and managing subprocesses.

**Q168: How do you read/write files asynchronously?**
Using libraries like `aiofiles`, since standard `open()` is blocking.

**Q169: What is pattern matching (match-case) in Python 3.10?**
Structural pattern matching. Similar to switch-case but more powerful (can match types, shapes).
```python
match status:
    case 200: print("OK")
    case 404: print("Not Found")
```

**Q170: What are positional-only arguments?**
Arguments that can only be passed by position, not by keyword. Defined using `/`.
`def func(x, /, y):`

---

## 🔹 18. Best Practices & Tools (Questions 171-180)

**Q171: How do you document python code effectively?**
Use Docstrings (Google/NumPy style), Type Hinting, and descriptive variable/function names.

**Q172: What are Python linters (Pylint, Flake8)?**
Tools that statically analyze code to check for errors, coding standards (PEP 8), and bad smells.

**Q173: What is Black formatter?**
An uncompromising code formatter. It reformats entire files to a consistent style, removing manual debate over style.

**Q174: How do you structure a large Python project?**
Separate concerns: `src/` (code), `tests/`, `docs/`, `requirements.txt`. Use packages/modules to organize logic logically.

**Q175: What is dependency injection in Python?**
Passing dependencies (objects/services) into a class/function rather than creating them internally. Improves testability.

**Q176: What is The Twelve-Factor App methodology?**
A set of best practices for building modern, scalable, cloud-native applications (e.g., config in env, stateless processes).

**Q177: How do you manage configuration/secrets?**
Environment variables (`os.environ`). Don't hardcode secrets. Use `.env` files (python-dotenv).

**Q178: What is setup.py?**
A script (legacy) used for packaging and distributing Python modules. Defines metadata (name, version, dependencies).

**Q179: What is a wheel file?**
A built-package format (`.whl`). It installs faster than source distributions because it doesn't require compilation.

**Q180: How do you publish a package to PyPI?**
Build the package (`build` tool), then upload using `twine`.

---

## 🔹 19. Tricky & Conceptual Questions (Questions 181-190)

**Q181: What is the output of `0.1 + 0.2 == 0.3`?**
`False`. Due to floating-point precision issues in binary representation, it equals `0.30000000000000004`.

**Q182: Why do we need `if __name__ == "__main__"`?**
To allow a file to be both imported as a module (reusable code) and run as a standalone script (main execution).

**Q183: How do you reverse a list without using `reverse()`?**
`lst[::-1]` or `list(reversed(lst))`.

**Q184: Can you modify a tuple?**
No, tuples are immutable. You must create a new one.

**Q185: What happens if you modify a list while iterating over it?**
Unexpected behavior (skipping elements) because the list size/indexes shift. Best practice: Iterate over a copy (`list[:]`).

**Q186: How can a function return multiple values?**
It returns them as a **tuple**.
`return x, y` is effectively `return (x, y)`.

**Q187: How to call a C function from Python?**
Using `ctypes`, `cffi`, or writing a C-extension (Python C-API).

**Q188: What is the purpose of `_` (underscore) in interpreter?**
Stores the result of the last executed expression in the interactive shell (REPL).

**Q189: Can you handle an error without `try/except`?**
Generally no, but you can use conditional checks (`if os.path.exists`) to prevent some, or `contextlib.suppress`.

**Q190: What is the Maximum Recursion Depth?**
The limit on the stack depth (default 1000). Prevent infinite recursion crashes. `sys.setrecursionlimit()` can change it.

---

## 🔹 20. Advanced Syntax (Questions 191-200)

**Q191: How to define a constant in Python?**
Python doesn't have strict constants. By convention, use UPPERCASE names (`MAX_SIZE = 100`).

**Q192: What are callable objects?**
Any object that can be called like a function (has `__call__` method). Functions, Classes, and methods are callables.

**Q193: How does `isinstance()` differ from `type()`?**
*   `type(obj) == Class`: Strictly checks class.
*   `isinstance(obj, Class)`: Checks class **AND** subclasses (supports inheritance).

**Q194: What is the use of `assert` statement?**
Used for debugging to test if a condition is true. If false, raises `AssertionError`.
`assert len(x) > 0`

**Q195: How do you create a generic function?**
Using `functools.singledispatch` (overloading by first arg type) or `typing.Generic`.

**Q196: What is function annotation?**
Metadata attached to function parameters/return. Commonly used for type hints. stored in `__annotations__`.

**Q197: How do you create a command-line interface (CLI) in Python?**
Using `argparse` (built-in) or libraries like `Click` or `Typer`.

**Q198: What is the difference between `json.dump` and `json.dumps`?**
*   `dump`: Writes to a **file** object.
*   `dumps`: Returns a **string** (S stands for String).

**Q199: How do you swap two variables in one line?**
`a, b = b, a` (Tuple unpacking).

**Q200: Why is Python not good for mobile development?**
It consumes higher and memory power (runtime overhead). Native languages (Swift/Kotlin/Java) or frameworks like Flutter/React Native are preferred, though Kivy/BeeWare exist.

---

## 🔹 21. FastAPI & Modern Web (Questions 201-210)

**Q201: What is FastAPI and what are its key features?**
A modern, fast (high-performance) web framework for building APIs with Python 3.6+ based on standard Python type hints. Key features: Fast (Starlette), Data Validation (Pydantic), Automatic Docs (Swagger UI).

**Q202: What is Pydantic and how is it used in FastAPI?**
A data validation library. FastAPI uses it to validate request bodies/responses. You define a class inheriting from `BaseModel`, and Pydantic ensures data types are correct.

**Q203: Explain Dependency Injection in FastAPI.**
A system where you declare things your path operation function needs (like DB sessions, auth tokens), and FastAPI provides them at runtime. defined using `Depends()`.

**Q204: How does FastAPI handle asynchronous programming (async/await)?**
It is built on ASGI (Asynchronous Server Gateway Interface) and natively supports `async def` route handlers, allowing non-blocking I/O operations (DB calls, external APIs).

**Q205: What is the difference between Flask and FastAPI?**
*   **Flask:** WSGI, synchronous by default, mature ecosystem, looser typing.
*   **FastAPI:** ASGI, async native, faster performance, strict content validation via Type Hints.

**Q206: How do you define a path parameter in FastAPI?**
By declaring it in the path string and the function argument.
```python
@app.get("/items/{item_id}")
def read_item(item_id: int): ...
```

**Q207: What are Query Parameters and how do you use them?**
Arguments declared in the function but not in the path string.
```python
@app.get("/items/")
def read_item(skip: int = 0, limit: int = 10): ...
```

**Q208: How do you validate request bodies in FastAPI?**
By creating a Pydantic model and declaring it as a function parameter.
```python
class Item(BaseModel):
    name: str
    price: float
@app.post("/items/")
def create_item(item: Item): ...
```

**Q209: What is OpenAPI (Swagger UI) and how is it generated?**
A standard specification for APIs. FastAPI automatically generates an openapi.json schema from your code and serves interactive docs at `/docs` (Swagger UI) and `/redoc`.

**Q210: How do you handle HTTP exceptions in FastAPI?**
Using `HTTPException`.
```python
from fastapi import HTTPException
raise HTTPException(status_code=404, detail="Item not found")
```

---

## 🔹 22. FastAPI Advanced (Questions 211-220)

**Q211: What is Middleware in FastAPI?**
Code that runs before and after every request. Used for CORS, GZip compression, or Session management. Added via `app.add_middleware()`.

**Q212: How do you implement OAuth2 authentication with JWT?**
FastAPI provides `OAuth2PasswordBearer` for the flow. You generate a JWT token upon login and validate it in a dependency function (`get_current_user`) for protected routes.

**Q213: What is Uvicorn and why is it needed?**
An ASGI web server implementation. FastAPI (the framework) needs an ASGI server (Uvicorn/Hypercorn) to run the application.

**Q214: How do you implement background tasks in FastAPI?**
Using `BackgroundTasks` parameter.
```python
def send_email(email: str): ...
@app.post("/send-notification/")
async def send_notification(email: str, background_tasks: BackgroundTasks):
    background_tasks.add_task(send_email, email)
```

**Q215: How do you write unit tests for FastAPI using TestClient?**
FastAPI provides `TestClient` (based on Requests) to make calls to your app without running a server.
```python
from fastapi.testclient import TestClient
client = TestClient(app)
def test_read_main():
    response = client.get("/")
    assert response.status_code == 200
```

**Q216: What is the purpose of `Depends` in FastAPI?**
It declares a dependency. It allows re-using logic (like database connection or security checks) across multiple path operations.

**Q217: How do you handle CORS (Cross-Origin Resource Sharing)?**
Using `CORSMiddleware`.
```python
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["*"],
)
```

**Q218: What is the difference between `APIRouter` and `FastAPI` instance?**
*   `FastAPI`: The main application class.
*   `APIRouter`: A mini-application used to split the API into multiple files/modules. You include routers into the main app.

**Q219: How do you support WebSockets in FastAPI?**
Using `WebSocket` endpoint.
```python
@app.websocket("/ws")
async def websocket_endpoint(websocket: WebSocket):
    await websocket.accept()
    await websocket.send_text("Hello")
```

**Q220: How do you validate environment variables using Pydantic?**
Using `BaseSettings`.
```python
from pydantic import BaseSettings
class Settings(BaseSettings):
    app_name: str = "My App"
    admin_email: str
settings = Settings()
```

---

## 🔹 23. Python Refresher & Concepts (Questions 221-230)

**Q221: What is the difference between list and tuple?**
(Refresher) List is mutable, Tuple is immutable.

**Q222: What is Python?**
(Refresher) An interpreted, high-level, general-purpose programming language.

**Q223: What are the key features of Python?**
Easy to learn, interpreted, dynamically typed, object-oriented, huge library support.

**Q224: How is Python interpreted?**
Source -> Bytecode -> PVM (Python Virtual Machine) -> Machine Code.

**Q225: What are Python’s data types?**
Standard types: int, float, complex, bool, str, list, tuple, range, dict, set, frozenset, bytes.

**Q226: What is the difference between is and ==?**
`is`: Identity operator (checks memory address). `==`: Equality operator (checks value).

**Q227: What are Python’s standard data types?**
See Q225. (Duplicate coverage in source, answering for completeness).

**Q228: What is the use of if, elif, and else?**
Conditional branching. `if` starts, `elif` adds another condition, `else` catches anything else.

**Q229: What is a loop in Python?**
A control flow statement for repeated execution. `for` (iteration over sequence) and `while` (repeat until condition false).

**Q230: Difference between for and while loops?**
*   `for`: Used when number of iterations is known (or iterating a collection).
*   `while`: Used when iterations depend on a condition.

---

## 🔹 24. Syntax & Basic Logic (Questions 231-240)

**Q231: What are break, continue, and pass?**
Loop controls. `break` (exit loop), `continue` (skip to next iteration), `pass` (do nothing).

**Q232: How does Python handle switch-case?**
Pre-3.10: Use dictionary mapping or if-elif chains. 3.10+: Use `match-case` statement.

**Q233: How do you define a function in Python?**
Using `def` keyword.
`def my_func(arg): return arg`

**Q234: What is *args and **kwargs?**
Variable length arguments. `*args` (tuple of positional args), `**kwargs` (dict of keyword args).

**Q235: What are default arguments?**
Arguments with a preset value if not provided. `def func(a=10):`.

**Q236: What is recursion?**
A function calling itself. Must have a base case to stop infinite loops.

**Q237: What is a string in Python?**
A sequence of Unicode characters. Enclosed in single or double quotes. Immutable.

**Q238: Are Python strings mutable?**
No. You cannot change a character in place. `s[0] = 'a'` throws TypeError.

**Q239: What are string slicing and indexing?**
Accessing parts of a string. `s[index]` or `s[start:end:step]`.

**Q240: Difference between str.strip() and str.lstrip()?**
*   `strip()`: Removes whitespace from **both** ends.
*   `lstrip()`: Removes whitespace from **left** end only.

---

## 🔹 25. Collections & Iterables (Questions 241-250)

**Q241: What is a list in Python?**
An ordered, mutable, heterogeneous collection of items so `[1, "a", True]`.

**Q242: How do you append and extend lists?**
`append(x)` adds x. `extend(iterable)` adds all elements of iterable to the list.

**Q243: What is list comprehension?**
A concise syntax for creating lists. `[expr for item in iterable if condition]`.

**Q244: How do you remove elements from a list?**
`remove(val)`, `pop(index)`, or `del list[index]`.

**Q245: How do you sort a list?**
*   `lst.sort()`: In-place sort.
*   `sorted(lst)`: Returns a new sorted list.

**Q246: What is a tuple?**
An ordered, immutable collection. `(1, 2)`.

**Q247: How are tuples different from lists?**
Immutability (Tuples) vs Mutability (Lists). Tuples are hashable (can be dict keys).

**Q248: What is tuple unpacking?**
Assigning tuple elements to variables. `x, y = (1, 2)`.

**Q249: When would you use a tuple over a list?**
For fixed collections of items, dictionary keys, or performance optimization (slight edge).

**Q250: What is a dictionary?**
A collection of key-value pairs. Unordered (insertion ordered since 3.7+).

---

## 🔹 26. Sets & Dictionaries (Questions 251-260)

**Q251: How do you add/update values in a dictionary?**
`d['key'] = value` or `d.update({'key': value})`.

**Q252: How do you loop through a dictionary?**
`for key, val in d.items():` loops through key-value pairs.

**Q253: What is the difference between get() and []?**
*   `d['key']`: Raises KeyError if missing.
*   `d.get('key')`: Returns `None` (or default) if missing.

**Q254: What are dictionary comprehensions?**
`{k:v for item in iterable}`.

**Q255: What is a set?**
An unordered collection of unique elements. `{1, 2, 3}`.

**Q256: How do you add and remove elements in a set?**
`s.add(val)` and `s.remove(val)` (raises Error if missing) or `s.discard(val)` (no error).

**Q257: What is the difference between set and list?**
Set: Unique, unordered, O(1) lookup. List: Duplicates allowed, ordered, O(n) lookup.

**Q258: How do you find union and intersection?**
*   Union: `s1 | s2`
*   Intersection: `s1 & s2`

**Q259: Are sets ordered in Python?**
No. They do not record element position. (Unlike Dictionary in 3.7+).

**Q260: What is OOP?**
Object-Oriented Programming. Paradigm based on "objects" containing data and code (methods).

---

## 🔹 27. OOP & Exceptions (Questions 261-270)

**Q261: What are classes and objects?**
Class: Template/Blueprint. Object: Instance of the class.

**Q262: What is __init__() method?**
The constructor. Initializes the object's state.

**Q263: What is inheritance?**
Mechanism where a new class derives properties and behavior from an existing class.

**Q264: What is polymorphism?**
ability to present the same interface for differing underlying forms (e.g., `len()` works on list and str).

**Q265: What are exceptions?**
Events causing the program to stop normal execution (Errors).

**Q266: How do you handle exceptions in Python?**
Using `try...except` blocks.

**Q267: What is the use of finally?**
A block that runs regardless of whether an exception occurred (for cleanup).

**Q268: What is raise used for?**
To explicitly trigger an exception.

**Q269: What are custom exceptions?**
User-defined error classes inheriting from `Exception`.

**Q270: What is a module in Python?**
A file containing Python code.

---

## 🔹 28. Modules & I/O (Questions 271-280)

**Q271: How do you import a module?**
`import math` or `from math import sqrt`.

**Q272: What is __name__ == '__main__'?**
Checks if the script is running directly (not imported).

**Q273: What is the difference between a module and a package?**
Package is a directory of modules with an `__init__.py`. Module is a file.

**Q274: What is pip?**
Python's package installer.

**Q275: How do you read a file in Python?**
`f.read()` reads entire content.

**Q276: How do you write to a file?**
`f.write("text")` in 'w' or 'a' mode.

**Q277: Difference between read() and readlines()?**
`read()`: String of full content. `readlines()`: List of strings (one per line).

**Q278: What is with open() used for?**
Context manager for safe file handling (auto-close).

**Q279: How do you check if a file exists?**
`os.path.exists(path)`.

**Q280: What is map()?**
Applies a function to all items in an input list.

---

## 🔹 29. Functional & Advanced (Questions 281-290)

**Q281: What is filter()?**
Creates a list of elements for which a function returns true.

**Q282: What is zip()?**
Iterates multiple iterables in parallel, returning tuples.

**Q283: What is enumerate()?**
Returns iterator of (index, value) pairs.

**Q284: What is any() and all()?**
*   `any()`: True if at least one element is True.
*   `all()`: True if all elements are True.

**Q285: What is set comprehension?**
`{expr for item in iterable}`. Creates a set.

**Q286: When to use comprehensions?**
For concise, readable transformations of collections.

**Q287: How to add conditions inside a comprehension?**
`[x for x in data if x > 10]`.

**Q288: What is the difference between deep copy and shallow copy?**
(Refresher) Deep copy copies objects recursively. Shallow copy copies references.

**Q289: What is slicing?**
Extracting sub-parts of a sequence.

**Q290: What is Python’s GIL?**
(Refresher) Global Interpreter Lock. Prevents multiple native threads from executing Python bytecodes at once.

---

## 🔹 30. Algorithms & Logic (Questions 291-300)

**Q291: What are decorators?**
(Refresher) Functions that modify other functions.

**Q292: What is the use of yield?**
To convert a function into a generator.

**Q293: What is a stack? How do you implement it in Python?**
LIFO structure. Use List: `append()` to push, `pop()` to pop.

**Q294: What is a queue? How do you implement it?**
FIFO structure. Use `collections.deque`: `append()`, `popleft()`.

**Q295: What is a linked list?**
A linear data structure where elements are not stored in contiguous memory. Python list is NOT a linked list (it's a dynamic array).

**Q296: How do you reverse a linked list?**
Iterate, changing `next` pointer of current node to `prev`.

**Q297: What is a binary tree?**
Tree data structure where each node has at most two children.

**Q298: How do you implement binary search?**
Divide and conquer on a sorted list. O(log n).

**Q299: How do you implement bubble sort?**
Nested loops, swapping adjacent elements if wrong order. O(n^2).

**Q300: How do you check if a string is a palindrome?**
`s == s[::-1]`.

---

## 🔹 31. CS Theory & Algorithms (Questions 301-310)

**Q301: How do you find the factorial of a number?**
Recursive: `return 1 if n==0 else n * fact(n-1)`. Iterative: `math.factorial(n)`.

**Q302: How to count occurrences of words in a string?**
`collections.Counter(string.split())`.

**Q303: How to swap two variables without using a temp variable?**
`a, b = b, a`.

**Q304: How to remove duplicates from a list?**
`list(set(lst))` (unordered) or `list(dict.fromkeys(lst))` (ordered, Py3.7+).

**Q305: How do you check if two lists have any common elements?**
`not set(list1).isdisjoint(list2)`.

**Q306: How to flatten a nested list?**
List comprehension: `[item for sublist in lst for item in sublist]` or `itertools.chain(*lst)`.

**Q307: What is Python bytecode?**
The internal representation of a Python program (low-level instructions) executed by the PVM. stored in `.pyc`.

**Q308: What is a namespace?**
A context in which a name (variable) exists.

**Q309: What are Python’s memory management techniques?**
Reference counting and Garbage Collection (Generational Cyclic GC).

**Q310: What are Python iterators and generators?**
(Refresher) Iterators traverse collections. Generators simplify creating iterators using `yield`.

---

## 🔹 32. Basics & formatting (Questions 311-320)

**Q311: What happens when you execute a Python script?**
It is compiled to bytecode and then executed by the PVM.

**Q312: What is indentation in Python and why is it important?**
It defines code blocks (scope). It is mandatory, replacing braces `{}` found in C/Java.

**Q313: Can Python run without semicolons?**
Yes. Semicolons are optional and used only to separate multiple statements on one line.

**Q314: How do you comment multiple lines in Python?**
Use `#` on each line. Triple quotes `'''` are often used but technically create strings.

**Q315: What is a docstring in Python?**
(Refresher) String literal as the first statement in a module/function/class defining its documentation.

**Q316: How do you write a multiline string?**
Using triple quotes: `"""Line 1\nLine 2"""`.

**Q317: What is the difference between int and float?**
`int`: Whole numbers. `float`: Decimal numbers (floating point).

**Q318: What is the type of True and False in Python?**
`bool` (which is a subclass of `int`).

**Q319: What does bool(0) and bool('') return?**
`False`. (0, empty sequences, and None are falsy).

**Q320: What is the output of 5 // 2?**
`2` (Floor division).

---

## 🔹 33. Operators & Casting (Questions 321-330)

**Q321: What does divmod() do?**
Returns a tuple `(quotient, remainder)`. `divmod(5, 2)` -> `(2, 1)`.

**Q322: What is the difference between == and is?**
(Refresher) `==` value equality. `is` reference identity.

**Q323: What is the in operator used for?**
Membership test. `x in y` checks if x is present in y.

**Q324: What is a bitwise operator?**
Operators processing bits: `&` (AND), `|` (OR), `^` (XOR), `~` (NOT), `<<`, `>>`.

**Q325: How does the not operator work?**
Inverts boolean value. `not True` is `False`.

**Q326: What is the output of True + True?**
`2`. (True is treated as 1 in arithmetic context).

**Q327: How do you convert a string to an integer?**
`int("123")`.

**Q328: What does int("10.5") do?**
Raises `ValueError`. String must represent an integer. You must do `int(float("10.5"))`.

**Q329: How do you safely cast a float to int?**
`int(10.9)` truncates to `10`. `round(10.9)` rounds to `11`.

**Q330: How do you convert a list to a tuple?**
`tuple([1, 2])`.

---

## 🔹 34. String Manipulation (Questions 331-340)

**Q331: What happens if you try to cast a string with letters to an int?**
`ValueError`.

**Q332: What does str.find() return if the substring is not found?**
Returns `-1`. (Unlike `index()` which raises ValueError).

**Q333: How do you check if a string is numeric?**
`s.isdigit()` or `s.isnumeric()`.

**Q334: What is the difference between replace() and translate()?**
*   `replace()`: Replaces occurrences of a substring.
*   `translate()`: Maps characters one-to-one using a translation table (faster for char replacement).

**Q335: How to capitalize the first letter of every word?**
`s.title()`.

**Q336: How to remove all whitespace from a string?**
`s.replace(" ", "")` or using regex.

**Q337: What is the difference between list.remove(x) and list.pop(i)?**
*   `remove(x)`: Removes by value.
*   `pop(i)`: Removes by index.

**Q338: What does list.count(x) do?**
Returns the number of times `x` appears in the list.

**Q339: How do you reverse a list without using .reverse()?**
`lst[::-1]`.

**Q340: How to find the index of an element in a list?**
`lst.index(x)`.

---

## 🔹 35. Tuples & Dictionaries (Questions 341-350)

**Q341: How to remove all duplicates from a list?**
(Refresher) `list(set(lst))`.

**Q342: How do you convert a tuple to a list?**
`list((1, 2))`.

**Q343: How do you create a single-element tuple?**
`(1,)` (Need the comma). `(1)` is just an integer in parenthesis.

**Q344: Can you delete a tuple?**
You can delete the variable reference `del t`, but not elements inside it.

**Q345: Are tuples faster than lists? Why?**
Yes. Immutability allows optimization by the runtime (simpler memory allocation).

**Q346: Can tuples have mutable elements?**
Yes. `t = ([1], 2)`. You can modify the list inside the tuple.

**Q347: How to get all keys from a dictionary?**
`d.keys()` (returns a view object).

**Q348: How to merge two dictionaries?**
(Refresher) `d1 | d2` or `d1.update(d2)`.

**Q349: What is dict.items() used for?**
Returns a view of `(key, value)` pairs.

**Q350: What happens when you access a non-existent key?**
`KeyError`.

---

## 🔹 36. Sets & Functions (Questions 351-360)

**Q351: What does setdefault() do?**
Returns value if key exists. If not, inserts key with specified value and returns it.

**Q352: How do you check if a value exists in a set?**
`x in s`.

**Q353: How to remove duplicates from a list using a set?**
(Refresher) `list(set(lst))`.

**Q354: What is the difference between discard() and remove()?**
*   `remove(x)`: Raises KeyError if x not found.
*   `discard(x)`: Does nothing if x not found.

**Q355: How do you clear a set?**
`s.clear()`.

**Q356: Can sets have nested lists as elements?**
No. Set elements must be hashable (immutable), and lists are mutable.

**Q357: What is a function signature?**
The declaration of a function (name, parameters, return type).

**Q358: What does return without a value return?**
`None`.

**Q359: Can a function return multiple values?**
Yes, as a tuple.

**Q360: What is the purpose of nonlocal?**
Declares that a variable belongs to an enclosing (nested) scope, not local or global.

---

## 🔹 37. Scopes & Loops (Questions 361-370)

**Q361: What is the LEGB rule in Python?**
(Refresher) Local, Enclosing, Global, Built-in variable scope resolution order.

**Q362: What is a global variable shadowing?**
When a local variable has the same name as a global one, hiding the global variable within the local scope.

**Q363: What is the difference between global and nonlocal?**
*   `global`: Refers to module-level variables.
*   `nonlocal`: Refers to enclosing function variables.

**Q364: What is a namespace collision?**
When two distinct variables/identifiers have the same name in the same scope.

**Q365: What does for else mean in Python?**
The `else` block runs after the loop completes normally, but NOT if the loop was exited via `break`.

**Q366: How to iterate through two lists simultaneously?**
`zip(l1, l2)`.

**Q367: What is range() and how does it work?**
Generates an immutable sequence of numbers. `range(start, stop, step)`.

**Q368: What is the output of range(10, 0, -1)?**
`10, 9, 8, ..., 1`.

**Q369: Can you break a for loop from inside an if?**
Yes. `if condition: break`.

**Q370: What is the difference between SyntaxError and IndentationError?**
IndentationError is a specific subclass of SyntaxError related to incorrect spacing/tabs.

---

## 🔹 38. Exceptions & System (Questions 371-380)

**Q371: How do you catch multiple exceptions?**
`except (Error1, Error2):`.

**Q372: What is the base class for all exceptions?**
`BaseException` (SystemExit, KeyboardInterrupt inherit from it). Standard code errors inherit from `Exception`.

**Q373: What does try-except-finally do?**
Tries code, catches errors, and ensures cleanup logic runs.

**Q374: How to raise a specific custom exception?**
`raise MyException("Message")`.

**Q375: How to check if a file exists using Python?**
`os.path.exists()`.

**Q376: What does os.listdir() do?**
Returns a list of names of entries in the directory.

**Q377: How to get the current working directory?**
`os.getcwd()`.

**Q378: How to delete a file?**
`os.remove()`.

**Q379: What does reversed() do?**
Returns a reverse iterator for a sequence.

**Q380: How does sorted() differ from .sort()?**
`sorted()` returns a new list. `.sort()` modifies the list in-place.

---

## 🔹 39. Generators & Magic Methods (Questions 381-390)

**Q381: What does chr() and ord() do?**
*   `chr(97)` -> `'a'` (ASCII/Unicode to char).
*   `ord('a')` -> `97` (Char to integer code).

**Q382: What does isinstance() check?**
Checks if object is an instance of class or subclass.

**Q383: What is the use of id()?**
Returns the unique memory identity (address) of the object.

**Q384: What is a generator?**
(Refresher) Function using `yield` that maintains state.

**Q385: What is the difference between a generator and an iterator?**
All generators are iterators, but generators are created using functions/yield, while custom iterators use classes with `__iter__`.

**Q386: How do you create a generator function?**
Use `yield` keyword inside a `def`.

**Q387: What is next() used for?**
Retrieves the next item from an iterator/generator.

**Q388: What is the use case of yield from?**
Delegates part of a generator's operations to another generator.

**Q389: What is the meaning of “Pythonic” code?**
Code that follows Python's conventions and best practices (readable, concise, using built-features).

**Q390: What does a, b = b, a do?**
Swaps values of `a` and `b`.

---

## 🔹 40. Environment & CLI (Questions 391-400)

**Q391: How do you check if a list is empty?**
`if not lst:` (Empty list is Falsy).

**Q392: How do you use with statement?**
`with open(...) as f:` ensures setup/teardown (like closing files).

**Q393: What does any() and all() do?**
(Refresher) `any`: True if >=1 true. `all`: True if all true.

**Q394: How to run a Python file from command line?**
`python script.py`.

**Q395: What is the use of input()?**
Reads a line from standard input as a string.

**Q396: How to pass command-line arguments to a script?**
Using `sys.argv` or `argparse` library.

**Q397: What is the difference between python and python -i?**
`python -i` runs the script and then drops into the interactive shell (REPL) instead of exiting.

**Q398: What does exit() do in REPL?**
Exits the interactive shell.

**Q399: How do you get current time in Python?**
`datetime.datetime.now()`.

**Q400: What is the difference between datetime and time modules?**
*   `datetime`: High-level dates/times manipulation.
*   `time`: Low-level system time (Unix timestamps).

---

## 🔹 41. CLI & Date/Time (Questions 401-410)

**Q401: What does time.sleep() do?**
Pauses the execution of the program for the given number of seconds.

**Q402: How do you format a date string?**
`date_obj.strftime("%Y-%m-%d")`.

**Q403: What is the use of the math module?**
Provides standard mathematical functions like `sqrt`, `sin`, `cos`, `pi`, `ceil`, `floor`.

**Q404: What is __pycache__ folder?**
Directory where Python stores compiled bytecode files (`.pyc`) to speed up subsequent imports.

**Q405: What is the .pyc file?**
(Refresher) Bytecode file generated by CPython.

**Q406: What is the difference between del, remove(), and pop()?**
(Refresher) `del` (keyword) deletes reference/index. `remove()` delete first value. `pop()` deletes and return index.

**Q407: What are Python keywords? How to list them?**
Reserved words (`if`, `else`, `while`). List via `import keyword; keyword.kwlist`.

**Q408: Can Python be used for mobile apps?**
Yes, via frameworks like Kivy or BeeWare, but not native.

**Q409: What happens if Python file ends without a newline?**
Nothing bad. Pylint might warn (PEP 8 recommends it), but code runs fine.

**Q410: How does Python handle indentation errors?**
Raises `IndentationError` before execution starts (during parsing).

---

## 🔹 42. Basics & Variables (Questions 411-420)

**Q411: What is a Python shell?**
Interactive command-line interface (REPL) to type and execute Python code immediately.

**Q412: How do you check the Python version in your system?**
`python --version` (CLI) or `sys.version` (script).

**Q413: What is the difference between a script and a module?**
(Refresher) Script: Intended to be run (`__name__=="__main__"`). Module: Intended to be imported.

**Q414: Can a variable name start with an underscore?**
Yes. `_var` (convention for internal use), `__var` (private clashing).

**Q415: What are dynamic types in Python?**
Variables are bound to objects, not types. A variable can hold an int, then a string later.

**Q416: Can you swap two variables without a third variable?**
Yes. `a, b = b, a`.

**Q417: What is unpacking in Python?**
Extracting values from iterable (list/tuple) into individual variables. `a, b = [1, 2]`.

**Q418: What does a = b = c = 10 mean?**
Chained assignment. All three variables point to the same integer object `10`.

**Q419: What is the difference between is not and !=?**
*   `is not`: Checks if memory address is different.
*   `!=`: Checks if values are unequal.

**Q420: What is an identity operator?**
operators `is` and `is not`.

---

## 🔹 43. Operators & Strings (Questions 421-430)

**Q421: How do logical operators work with non-boolean values?**
They return one of the operands (short-circuiting). `5 and 10` -> `10`. `0 or 5` -> `5`.

**Q422: What is operator precedence in Python?**
The order in which operations are performed (PEMDAS rule applies). `**` is highest.

**Q423: What is the result of not (True and False)?**
`True` (not False).

**Q424: How do you check for a substring in a string?**
`if "sub" in string:`.

**Q425: What does string.join(list) do?**
Concatenates list elements into a string using the separator. `",".join(['a', 'b'])` -> `"a,b"`.

**Q426: What is the difference between find() and index()?**
*   `find()`: Returns -1 if missing.
*   `index()`: Raises ValueError if missing.

**Q427: What does rjust() and ljust() do?**
Pads string with spaces (or char) to designated length. `ljust` (align left), `rjust` (align right).

**Q428: What is the use of zfill() in strings?**
Pads string with zeros on the left. `"5".zfill(3)` -> `"005"`.

**Q429: What is a shallow copy of a list?**
`lst[:]` or `list(lst)` or `copy.copy()`. Copies references of elements.

**Q430: What is the difference between slicing and indexing?**
*   Indexing (`a[0]`): Returns element. Error if out of bounds.
*   Slicing (`a[0:1]`): Returns new list. Safe if out of bounds (empty list).

---

## 🔹 44. Lists & Tuples (Questions 431-440)

**Q431: How does list.clear() work?**
Removes all items from the list in-place. `lst[:] = []`.

**Q432: Can lists contain other lists?**
Yes (Nested lists).

**Q433: What is the output of list * 2?**
Replicates the list. `[1] * 2` -> `[1, 1]`.

**Q434: Can a tuple contain another tuple?**
Yes.

**Q435: How do you slice a tuple?**
Same as list. `t[1:3]`. Returns a new tuple.

**Q436: Are tuples hashable?**
Yes, if they contain only immutable elements. Can be used as dictionary keys.

**Q437: How to loop through a tuple?**
`for item in t:`.

**Q438: What is the use of namedtuple?**
Creates tuple subclasses with named fields. Memory efficient objects.
`Point = namedtuple('Point', ['x', 'y'])`.

**Q439: How to update multiple key-values at once?**
`d.update({'k1': v1, 'k2': v2})`.

**Q440: What does dict.popitem() do?**
Removes and returns the last inserted key-value pair (LIFO) in Python 3.7+.

---

## 🔹 45. Dicts & Sets (Questions 441-450)

**Q441: What is dictionary unpacking using **?**
Merges dictionary into another context. `func(**d)` passes dict items as kwargs.

**Q442: How to invert keys and values in a dictionary?**
`{v: k for k, v in d.items()}`. (Warning: Values must be unique hashable).

**Q443: Can dictionary keys be mutable?**
No. Keys must be hashable (immutable). Lists cannot be keys. Tuples can.

**Q444: What is symmetric difference in sets?**
Elements present in either set, but not both. `s1 ^ s2`.

**Q445: What is the difference between subset and superset?**
*   `A <= B` (Subset): All elements of A are in B.
*   `A >= B` (Superset): All elements of B are in A.

**Q446: How to find difference between two sets?**
`s1 - s2` (Elements in s1 but not s2).

**Q447: Can you store different data types in a set?**
Yes, as long as they are hashable strings, ints, tuples mixed together.

**Q448: What is a frozen set?**
(Refresher) Immutable set.

**Q449: What is a higher-order function?**
A function that takes a function as argument (e.g., `map`, `filter`) or returns a function (decorators).

**Q450: Can functions be nested?**
Yes. Inner functions can access outer scope variables (closures).

---

## 🔹 46. Functions & Lambdas (Questions 451-460)

**Q451: What is a callback function?**
A function passed to another function to be executed later (e.g., in async operations or event handling).

**Q452: What does locals() return?**
A dictionary of local variables in the current scope.

**Q453: What does globals() return?**
A dictionary of global variables in the module.

**Q454: How do you sort a list of tuples using lambda?**
`lst.sort(key=lambda x: x[1])` (sort by second element).

**Q455: Can lambda functions have multiple expressions?**
No. Only one single expression (returned value).

**Q456: What are the limitations of lambda functions?**
Single line, no statements (assignments, loop, print), limited readability.

**Q457: How do you use lambda with map()?**
`list(map(lambda x: x*2, lst))`.

**Q458: How do you use lambda with filter()?**
`list(filter(lambda x: x>0, lst))`.

**Q459: How do you use list comprehension with if-else?**
`[x if x > 0 else 0 for x in lst]`.

**Q460: Can list comprehension have multiple loops?**
Yes. `[x+y for x in l1 for y in l2]` (Nested loops result in cartesian product).

---

## 🔹 47. Comprehensions & Loops (Questions 461-470)

**Q461: How to use list comprehension to flatten a list?**
`[item for sublist in lst for item in sublist]`.

**Q462: What is the equivalent for dictionary comprehension?**
`{k:v for k,v in iterable}`.

**Q463: Can set comprehension be used with conditions?**
Yes. `{x for x in lst if x > 0}`.

**Q464: How does continue work inside nested loops?**
It only skips the current iteration of the **inner-most** loop.

**Q465: Can you modify a list while iterating over it?**
Unsafe. See Q185.

**Q466: What does for x in range(len(list)) do?**
Iterates over indices of the list.

**Q467: Why should you avoid modifying a collection while looping?**
Causes index shifting, leading to skipped items.

**Q468: What is the use of reversed(range())?**
`reversed(range(5))` -> `4, 3, 2, 1, 0`.

**Q469: What error do you get when dividing by zero?**
`ZeroDivisionError`.

**Q470: What is a KeyError?**
Raised when accessing a dictionary key that does not exist.

---

## 🔹 48. Exceptions & Classes (Questions 471-480)

**Q471: What is AttributeError?**
Raised when attribute reference or assignment fails (e.g., `x.missing_method()`).

**Q472: What is the difference between TypeError and ValueError?**
*   `TypeError`: Operation applied to object of inappropriate type (`"1" + 2`).
*   `ValueError`: Operation applied to correct type but inappropriate value (`int("abc")`).

**Q473: What happens when no exception is caught?**
The program terminates and prints a traceback.

**Q474: How do you define a simple class?**
`class MyClass: pass`.

**Q475: What is self in Python?**
(Refresher) Reference to the instance.

**Q476: What is the purpose of __str__() method?**
Returns string representation for `str()` and `print()`.

**Q477: What is encapsulation?**
(Refresher) Bundling data/methods.

**Q478: What are instance variables?**
Variables defined inside `__init__` attached to `self`. Unique to each object.

**Q479: What is a decorator in Python?**
(Refresher) Wrapper function.

**Q480: How do you apply a decorator to a function?**
Using `@decorator_name` above definition.

---

## 🔹 49. Modules & Files (Questions 481-490)

**Q481: Can a decorator take arguments?**
Yes, requires an extra layer of nesting (wrapper returning wrapper).

**Q482: What is @staticmethod?**
(Refresher) Defines a static method.

**Q483: What is @classmethod?**
(Refresher) Defines a class method.

**Q484: What is the difference between import and from import?**
*   `import m`: Access via `m.func()`.
*   `from m import func`: Access via `func()` (namespace pollution risk).

**Q485: What happens when you import a module?**
The code in the module file is executed once.

**Q486: How to import a module under an alias?**
`import pandas as pd`.

**Q487: How to reload a module?**
`import importlib; importlib.reload(module)`.

**Q488: What is __init__.py used for?**
Marks a directory as a package.

**Q489: How to read a file line by line?**
`for line in f:` (File object is an iterator).

**Q490: How to count words in a file?**
`len(f.read().split())`.

---

## 🔹 50. File OS Operations (Questions 491-500)

**Q491: How do you append to a file?**
Open with mode `'a'`.

**Q492: What’s the difference between a+ and w+ modes?**
*   `w+`: Truncates (deletes content) file before reading/writing.
*   `a+`: Appends (preserves content) and allows reading.

**Q493: How do you read a specific number of bytes?**
`f.read(10)`.

**Q494: What does os.path.exists() do?**
Returns True if path exists.

**Q495: How to create a new folder in Python?**
`os.mkdir()` or `os.makedirs()` (recursive).

**Q496: How to list all .txt files in a directory?**
`glob.glob("*.txt")` or list comp with `endswith(".txt")`.

**Q497: How to delete a directory using Python?**
`os.rmdir()` (empty only) or `shutil.rmtree()` (recursive delete).

**Q498: What is os.path.join() used for?**
Joins path components intelligently (handles OS separators `/` or `\`).

**Q499: What is the difference between mutable and immutable types?**
(Refresher) Mutables can change in-place. Immutables cannot.

**Q500: What is a magic method?**
(Refresher) Dunder methods like `__init__`, `__str__` called implicitly.

---

## 🔹 51. Deployment & Virtualization (Questions 501-510)

**Q501: What is `__repr__()` method?**
Returns an unambiguous string representation of the object, used for debugging.

**Q502: What is the difference between `==` and `__eq__()`?**
They are effectively the same. `a == b` calls `a.__eq__(b)`.

**Q503: How to install a Python package?**
`pip install package_name`.

**Q504: What is a virtual environment?**
Isolated Python environment to manage dependencies locally for a specific project.

**Q505: How do you activate a virtual environment?**
*   Windows: `venv\Scripts\activate`
*   Mac/Linux: `source venv/bin/activate`

**Q506: What is the difference between requirements.txt and setup.py?**
*   `requirements.txt`: List of dependencies for **deploying** an app (reproducibility).
*   `setup.py`: Script for **distributing** a package (defines metadata).

**Q507: 📚 All 300+ questions combined in a PDF or Notion format?**
(Meta-question: Skipping).

**Q508: 🔁 Flashcards for revision?**
(Meta-question: Skipping).

**Q509: ✅ Full answers and code examples?**
(Meta-question: Skipping).

**Q510: What is an interactive Python session?**
The REPL (Read-Eval-Print Loop) where you type code and get immediate results.

---

## 🔹 52. Python Core (Questions 511-520)

**Q511: What is the difference between compiling and interpreting?**
*   **Compiler:** Translates entire code to machine code before execution (faster run, slower build).
*   **Interpreter:** Translates code line-by-line during execution (slower run, instant start). Python does both (compiles to bytecode, interprets bytecode).

**Q512: Can Python run on all operating systems?**
Yes, it is cross-platform (Windows, macOS, Linux, Unix).

**Q513: What is the REPL in Python?**
Read-Eval-Print Loop. The interactive shell.

**Q514: What is the Python `__main__` function used for?**
Python doesn't have a rigid `main()` function like C/Java. The top-level script environment is called `__main__`.

**Q515: How to declare a constant in Python?**
Use ALL_CAPS notation (`PI = 3.14`). It's a convention, not enforced (variables can still be changed).

**Q516: What is variable shadowing in Python?**
When a variable in an inner scope (function) has the same name as an outer variable, "shadowing" or hiding the outer one.

**Q517: What happens when you assign None to a variable?**
The variable points to the `None` object (singleton). It represents "no value".

**Q518: How do you assign multiple variables in one line?**
`x, y, z = 1, 2, 3`.

**Q519: Can Python variables be declared without assignment?**
No. A variable must be assigned a value to exist.

**Q520: What is a ternary conditional operator in Python?**
`value_if_true if condition else value_if_false`.

---

## 🔹 53. Operators & Logic (Questions 521-530)

**Q521: What is short-circuit evaluation?**
In boolean logic (`and`, `or`), Python stops evaluating as soon as the result is determined. `False and x` -> `False` (x never evaluated).

**Q522: How does Python handle chained comparisons like 3 < x < 10?**
It evaluates as `(3 < x) and (x < 10)`.

**Q523: Can you overload operators in Python?**
Yes, by defining magic methods (`__add__` for `+`, `__sub__` for `-`).

**Q524: What is the difference between arithmetic and assignment operators?**
*   Arithmetic: `+`, `-`, `*` (Return result).
*   Assignment: `=`, `+=` (Assign value).

**Q525: What does "abc" * 3 return?**
`"abcabcabc"`.

**Q526: How do you check if a string is a valid identifier?**
`s.isidentifier()`.

**Q527: How to format strings using f-strings?**
`f"Value: {val}"`.

**Q528: What is the difference between split() and rsplit()?**
*   `split()`: Splits from left.
*   `rsplit()`: Splits from right (useful only with `maxsplit`).

**Q529: How to remove a prefix/suffix from a string?**
*   `s.removeprefix('pre')` (Py3.9+)
*   `s.removesuffix('suf')` (Py3.9+)

**Q530: How to get the last element of a list?**
`lst[-1]`.

---

## 🔹 54. List Manipulation (Questions 531-540)

**Q531: What does list[-1::-1] do?**
Reverses the list. Identical to `list[::-1]`.

**Q532: What is the difference between list() and [x for x in iterable]?**
*   `list(iter)`: Uses constructor.
*   `[x for x in iter]`: Uses list comprehension (usually slightly faster).

**Q533: How do you copy a list using slicing?**
`new_list = old_list[:]`.

**Q534: What happens when you multiply a list by 0?**
Returns an empty list `[]`.

**Q535: How can you convert a tuple of tuples to a flat tuple?**
`tuple(sum(tup_of_tups, ()))` or using itertools.

**Q536: Are empty tuples equal: () == tuple()?**
Yes. Since tuples are immutable, Python reuses the same empty tuple object.

**Q537: How do you slice a tuple in reverse?**
`tup[::-1]`.

**Q538: Can you loop over a tuple using indices?**
Yes. `for i in range(len(tup)): print(tup[i])`.

**Q539: Can you pass tuples as function arguments?**
Yes. Pass as standard arg `func(t)` or unpack `func(*t)`.

**Q540: What happens when you use a list as a key in a dictionary?**
`TypeError: unhashable type: 'list'`. Lists are mutable and cannot be keys.

---

## 🔹 55. Dictionary & Set Tricks (Questions 541-550)

**Q541: How do you check if two dictionaries are equal?**
`d1 == d2` (Checks if keys and values match).

**Q542: How to merge two dictionaries in Python 3.9+?**
`d3 = d1 | d2`.

**Q543: What does dict.fromkeys() do?**
Creates a new dictionary with specified keys and a single default value. `dict.fromkeys(['a','b'], 0) -> {'a':0, 'b':0}`.

**Q544: How do you sort a dictionary by values?**
`dict(sorted(d.items(), key=lambda x: x[1]))`.

**Q545: Can sets contain mixed types (int, string, tuple)?**
Yes, as long as they are immutable elements.

**Q546: How do you convert a string to a set?**
`set("hello")` -> `{'h','e','l','o'}`.

**Q547: What happens if you try to add a list to a set?**
`TypeError`. Lists are unhashable.

**Q548: How do you get a sorted list from a set?**
`sorted(s)`.

**Q549: How to remove all common elements between two sets?**
`s1.symmetric_difference(s2)` keeps unique elements from both.
`s1.difference(s2)` keeps elements only in s1.

**Q550: Can functions be defined inside loops?**
Yes, but a new function object is created in each iteration (inefficient).

---

## 🔹 56. Functions & Lambdas II (Questions 551-560)

**Q551: What are positional-only parameters (Python 3.8+)?**
Defined using `/`. Arguments before `/` cannot be passed by keyword.
`def f(x, /): ...`

**Q552: What are keyword-only parameters?**
Defined using `*`. Arguments after `*` must be passed by keyword.
`def f(*, y): ...`

**Q553: What is function caching?**
Storing results of expensive function calls to return the cached result when the same inputs occur again.

**Q554: What does functools.lru_cache do?**
Decorator that implements memoization (Least Recently Used cache).
`@lru_cache(maxsize=None)`

**Q555: How to write a lambda that squares a number?**
`sq = lambda x: x**2`.

**Q556: How to use filter() to remove empty strings?**
`filter(None, ["a", "", "b"])`. (Passing None filters falsy values).

**Q557: What is the return type of map() in Python 3?**
A map object (iterator). You must wrap it in `list()` to see values.

**Q558: How does reduce() work in Python?**
Applies a rolling computation. `reduce(fn, [1,2,3])` -> `fn(fn(1,2), 3)`.

**Q559: What happens if map() and filter() return empty?**
They return an empty iterator. `list()` on it yields `[]`.

**Q560: Can comprehensions be nested?**
Yes. `[[j for j in range(2)] for i in range(3)]`.

---

## 🔹 57. Comprehensions & Iteration (Questions 561-570)

**Q561: How to use list comprehension with zip()?**
`[x+y for x, y in zip(l1, l2)]`.

**Q562: What’s the difference between a generator expression and a list comprehension?**
*   List Comp: `[...]` - Creates entire list in memory.
*   Gen Exp: `(...)` - Returns generator (lazy evaluation).

**Q563: Can you use ternary inside comprehension?**
Yes. `[x if x>0 else 0 for x in lst]`.

**Q564: Can comprehensions be used with sets and dicts?**
Yes. `{x for x in list}` (Set) and `{k:v for k,v in pairs}` (Dict).

**Q565: What is the difference between for i in list vs for i in range(len(list))?**
*   `in list`: Helper loop (Pythonic), gives values.
*   `range(len)`: Index loop (C-style), gives indices.

**Q566: How to iterate over a dictionary’s keys and values at the same time?**
`for k, v in d.items():`.

**Q567: Can while loop have an else clause?**
Yes. Runs if loop completes without `break`.

**Q568: How does enumerate() help in loops?**
Gives both index and value cleaner than `range(len())`.

**Q569: Can you break out of nested loops?**
`break` only exits the innermost loop. Use flags or `return` or exception to break all.

**Q570: What happens if finally block has a return statement?**
It overrides any return/exception from the try block. (Dangerous).

---

## 🔹 58. Advanced Exception Handling (Questions 571-580)

**Q571: Can try block exist without except?**
Yes, if it has `finally`.

**Q572: What is the hierarchy of exceptions?**
`BaseException` -> `Exception` -> Standard Errors (ValueError, TypeError, etc.).

**Q573: Can you raise exceptions manually?**
Yes, `raise ValueError`.

**Q574: What is the difference between Exception and BaseException?**
*   `BaseException`: Root class. Includes system exits.
*   `Exception`: Inherit this for user code. Excludes system exits.

**Q575: Can a class exist without methods?**
Yes. `class Data: pass`.

**Q576: What is the purpose of __del__()?**
Destructor. Called when object is garbage collected. (Rarely used/reliable).

**Q577: How do you override methods in Python?**
Define a method with the same name in the child class.

**Q578: What are mixins?**
Small classes designed to add functionality to others via inheritance, not meant to be instantiated alone.

**Q579: What does super() do?**
Returns a proxy object that delegates calls to a parent/sibling class (used to call parent methods).

**Q580: Can a closure access updated variables?**
Yes. It references the variable itself, not the value at definition time.

---

## 🔹 59. Decorators & Files (Questions 581-590)

**Q581: What is a practical use case of a decorator?**
Logging, timing execution, authentication checks, caching.

**Q582: Can you chain multiple decorators?**
Yes. `@d1 @d2`. They apply bottom-up (d2, then d1).

**Q583: What are function wrappers?**
The inner function returned by a decorator that "wraps" the original function.

**Q584: Can a decorator return a different function?**
Yes. It can return entirely new logic, replacing the original function completely.

**Q585: What happens if you open a file in r+ mode?**
Read and Write. The pointer starts at the beginning.

**Q586: What is binary file mode (rb, wb)?**
Reads/Writes bytes objects instead of strings. No encoding/decoding performed.

**Q587: How do you write multiple lines to a file?**
`f.writelines(list_of_strings)`.

**Q588: How to get the file pointer position?**
`f.tell()`.

**Q589: What does seek() do in file handling?**
Moves the file cursor to a specific byte position. `f.seek(0)`.

**Q590: What does the random module do?**
Generates pseudo-random numbers (`random.random`, `randint`, `choice`).

---

## 🔹 60. Libraries & Modules (Questions 591-600)

**Q591: What is collections.Counter()?**
A Dict subclass for counting hashable objects. `Counter('aabbc')` -> `{'a':2, 'b':2, 'c':1}`.

**Q592: What is itertools.product()?**
Cartesian product of input iterables. Like nested for-loops.

**Q593: What is the purpose of math.isclose()?**
Compares floating point numbers for equality within a tolerance (avoids precision errors).

**Q594: What is pathlib.Path()?**
Object-oriented filesystem paths (Python 3.4+). Easier than `os.path`.

**Q595: How do you upgrade a package using pip?**
`pip install --upgrade package_name`.

**Q596: How do you uninstall a package using pip?**
`pip uninstall package_name`.

**Q597: How to check all installed packages?**
`pip list` or `pip freeze`.

**Q598: What is pyproject.toml?**
New standard for Python configuration (build system requirements, tool configs like Black/Ruff).

**Q599: What is the difference between venv and virtualenv?**
*   `venv`: Built-in (Python 3.3+).
*   `virtualenv`: Third-party, supports older Python versions, more features.

**Q600: How do you check the type of an object at runtime?**
`type(obj)` or `isinstance(obj, Type)`.

---

## 🔹 61. Introspection & Typing (Questions 601-610)

**Q601: What does dir() return?**
A list of valid attributes and methods for an object (or the current scope if no argument).

**Q602: What is hasattr() used for?**
Checks if an object has a given attribute. `if hasattr(obj, 'prop'): ...`.

**Q603: What does getattr() do?**
Returns the value of a named attribute. `getattr(obj, 'x', default_value)`.

**Q604: What is monkey patching in Python?**
(Refresher) Runtime modification of a class or module.

**Q605: What is the use of assert?**
(Refresher) Debugging check.

**Q606: How to document Python code properly?**
(Refresher) Docstrings and type hints.

**Q607: How to convert a function to a string using inspect?**
`inspect.getsource(func)`.

**Q608: What does __slots__ do in Python classes?**
(Refresher) Optimizes memory by preventing `__dict__` creation.

**Q609: How does Python handle indentation internally?**
The parser generates `INDENT` and `DEDENT` tokens based on whitespace.

**Q610: What is a Python bytecode file?**
(Refresher) `.pyc` file containing compiled instructions.

---

## 🔹 62. Python Interpreter (Questions 611-620)

**Q611: What does the Python interpreter actually do?**
Parses source code, compiles it to bytecode, and executes it on the Virtual Machine.

**Q612: How is Python different from other scripting languages like Bash or Perl?**
Python emphasizes readability ("The Zen of Python") and general-purpose usage (OOP), while Bash/Perl focus more on shell scripting/text processing.

**Q613: What is interactive mode vs script mode?**
*   **Interactive:** REPL (One command at a time).
*   **Script:** Running a `.py` file (Batch execution).

**Q614: What is duck typing in Python?**
(Refresher) Behavior over type.

**Q615: Can a variable change its type during execution?**
Yes (Dynamic typing).

**Q616: What is type hinting? Is it enforced?**
(Refresher) Checking at static analysis time, not runtime.

**Q617: How do you check the type of a variable?**
`type(var)`.

**Q618: What is the difference between type() and isinstance()?**
(Refresher) `isinstance()` supports inheritance checks.

**Q619: What is the difference between float('nan') and None?**
*   `NaN`: Not a Number (a numeric float value).
*   `None`: Absence of value.

**Q620: How do you check for infinity in Python?**
`math.isinf(x)` or `x == float('inf')`.

---

## 🔹 63. Numbers & Math (Questions 621-630)

**Q621: What is the result of dividing two integers using / and //?**
*   `/`: Float (`5/2=2.5`).
*   `//`: Int (Floor) (`5//2=2`).

**Q622: What is a complex number in Python?**
Numbers with a real and imaginary part (`j`). `3 + 4j`.

**Q623: How do you round a float to 2 decimal places?**
`round(3.14159, 2)` -> `3.14`.

**Q624: What is the boolean value of empty structures?**
`False`.

**Q625: Is None a keyword or object?**
It is a singleton object of `NoneType`. (In Python 3 it is a keyword-like constant).

**Q626: What is the difference between None and False?**
`None` is nothing. `False` is a specific boolean value.

**Q627: How do you check if a value is None?**
`if x is None:`.

**Q628: What is the type of None?**
`<class 'NoneType'>`.

**Q629: How do you escape characters in a string?**
Using backslash `\`. `\'` for quote, `\n` for newline.

**Q630: What does ord() and chr() do?**
(Refresher) Char <-> Int conversion.

---

## 🔹 64. List Advanced (Questions 631-640)

**Q631: How do you reverse a string using slicing?**
`s[::-1]`.

**Q632: How to remove punctuation from a string?**
`s.translate(str.maketrans('', '', string.punctuation))`.

**Q633: What happens if you append to a list while iterating?**
You create an infinite loop (or crash memory).

**Q634: How do you flatten a nested list without using recursion?**
Using a stack/queue or `itertools.chain`.

**Q635: What is the difference between clear() and del list[:]?**
Both remove all elements. `clear()` is the specific method (more readable).

**Q636: What is list.__getitem__(index)?**
The dunder method called by `list[index]`.

**Q637: How do you rotate a list in Python?**
`lst = lst[-k:] + lst[:-k]` or using `collections.deque.rotate()`.

**Q638: How do you make a tuple with one element?**
`(x,)`.

**Q639: Why would you use zip(*args)?**
To "unzip" a list of tuples back into separate lists.

**Q640: Can tuples be sorted?**
You can't sort them in-place, but `sorted(tup)` returns a sorted list.

---

## 🔹 65. Advanced Collections (Questions 641-650)

**Q641: What’s the use of collections.namedtuple()?**
(Refresher) Tuple with field names.

**Q642: Are all tuple elements required to be immutable?**
No. Just the tuple container itself.

**Q643: What’s the output of dict.get('missing_key', 'default')?**
`'default'`.

**Q644: How do you remove a key safely?**
`d.pop('key', None)`.

**Q645: What is the use of collections.defaultdict()?**
Automatically creates values for missing keys.

**Q646: How do you create a dictionary from two lists?**
`dict(zip(keys, values))`.

**Q647: What is dictionary comprehension with conditions?**
`{k:v for k,v in d.items() if v > 10}`.

**Q648: What is the difference between union() and update()?**
*   `union()`: Returns new set.
*   `update()`: Modifies set in-place.

**Q649: How do you check if one set is a subset of another?**
`s1.issubset(s2)` or `s1 <= s2`.

**Q650: Can a set have duplicate elements?**
No. They are automatically removed.

---

## 🔹 66. Functional & Scope (Questions 651-660)

**Q651: How to iterate over a set?**
`for item in s:`. (Note: Order is undefined).

**Q652: What does frozenset do?**
(Refresher) Immutable set.

**Q653: Can an if block exist without else?**
Yes.

**Q654: What is the result of if []:?**
False (Empty list is falsy).

**Q655: Can for loop be used with else?**
Yes.

**Q656: How do you simulate switch-case in Python?**
Match-case (3.10+) or Dict mapping.

**Q657: What is the difference between pass and continue?**
*   `pass`: Does nothing.
*   `continue`: Jumps to next iteration.

**Q658: Can a function be assigned to a variable?**
Yes (First-class object). `f = print; f("Hi")`.

**Q659: What happens if a function returns nothing?**
It returns `None` implicitly.

**Q660: What does function.__doc__ do?**
Returns the docstring.

---

## 🔹 67. Lambda & Maps (Questions 661-670)

**Q661: How do you return multiple values from a function?**
Tuple packing.

**Q662: What does callable(obj) check?**
If object implements `__call__` (Functions, classes).

**Q663: Can lambda contain an if expression?**
Yes. `lambda x: "Yes" if x else "No"`.

**Q664: Can you use lambda inside a class?**
Yes, but usually not recommended (use `def` method).

**Q665: What is a lambda returning another lambda?**
Currying / Closure. `lambda x: (lambda y: x+y)`.

**Q666: What happens if you call a lambda without parameters?**
`lambda: print("Hi")`. It runs like a normal function.

**Q667: Can a lambda modify external variables?**
Only if they are mutable or strictly via side-effects (bad practice).

**Q668: How do you apply map() to a dictionary?**
It iterates over Keys by default.

**Q669: What happens when filter() returns an empty iterator?**
It yields nothing.

**Q670: Can you combine map() and lambda()?**
Yes. `map(lambda x: x+1, lst)`.

---

## 🔹 68. Advanced Comprehensions (Questions 671-680)

**Q671: Can reduce() be used to calculate factorial?**
`reduce(lambda x,y: x*y, range(1, n+1))`.

**Q672: What’s the difference between filter() and list comprehension?**
Comprehension is generally faster and more readable (Pythonic). `filter` returns an iterator (memory efficient).

**Q673: Can comprehensions be used with conditions and else?**
Yes. `[x if x else y for x in lst]`.

**Q674: Can you create a matrix using list comprehension?**
Yes. `[[0]*3 for _ in range(3)]`.

**Q675: Can you use list comprehension with enumerate()?**
`[i for i, x in enumerate(lst) if x > 0]`.

**Q676: How to create a dict from a list using comprehension?**
`{i: lst[i] for i in range(len(lst))}`.

**Q677: Are comprehensions faster than loops?**
Generally yes, because the loop is executed in C.

**Q678: How does next() work on an iterator?**
Returns next item. Raises `StopIteration` if exhausted.

**Q679: What is the difference between iterator and iterable?**
(Refresher) Iterable has `__iter__`. Iterator has `__next__`.

**Q680: Can you create a custom iterator?**
Yes. Implement `__iter__` and `__next__`.

---

## 🔹 69. Generators & Exceptions (Questions 681-690)

**Q681: What is the output of zip() if lists have different lengths?**
Stops at the shortest list. Use `zip_longest` to keep going.

**Q682: What is a generator expression?**
`(x for x in iter)`.

**Q683: What error occurs when accessing an undefined variable?**
`NameError`.

**Q684: What happens if finally has an error?**
The original exception (if any) is lost, and the new error propagates.

**Q685: Can a try block exist without a catch?**
Yes, with `finally`.

**Q686: What’s the purpose of catching Exception vs specific errors?**
Catching `Exception` is broad (catches almost everything). Specific is better practice.

**Q687: How do you define your own exception class?**
`class MyError(Exception): pass`.

**Q688: What is the difference between a class and an instance?**
Class is definition. Instance is the concrete object.

**Q689: What is an instance attribute?**
Attribute attached to `self`.

**Q690: Can methods be added dynamically to an object?**
Yes. `obj.method = func`.

---

## 🔹 70. IO & Imports (Questions 691-700)

**Q691: What is method overloading? Does Python support it?**
Not directly. See Q47.

**Q692: How do you read large files efficiently in Python?**
Iterate line by line (`for line in f`) to keep memory usage low.

**Q693: How do you write a list to a file?**
`f.writelines(list)`.

**Q694: What happens if you open a file that doesn’t exist in r mode?**
`FileNotFoundError`.

**Q695: What is the default encoding used in open()?**
Platform dependent (usually UTF-8 or CP1252). Best to specify `encoding='utf-8'`.

**Q696: How to read a file line-by-line using a generator?**
File objects are generators themselves.

**Q697: What is the difference between absolute and relative imports?**
*   Absolute: `from pkg.mod import f` (Full path).
*   Relative: `from .mod import f` (Relative to current file).

**Q698: What is a namespace package?**
A package spread across multiple directories (no `__init__.py` required in Py3.3+).

**Q699: How do you dynamically import a module?**
`importlib.import_module("mod_name")`.

**Q700: What is __package__?**
Attribute containing the name of the package the module belongs to.

---

## 🔹 71. Advanced Imports & Memory (Questions 701-710)

**Q701: What happens on import module behind the scenes?**
Python checks `sys.modules`, finds the file, compiles to bytecode (if needed), executes the module body, and creates a module object.

**Q702: What does shutil module do?**
High-level file operations: copying, moving, archiving, and directory tree management.

**Q703: How is tempfile useful in Python scripts?**
Creates temporary files/directories that are automatically cleaned up. Secure and unique names.

**Q704: What is uuid module used for?**
Generating universally unique identifiers (UUIDs), essentially random 128-bit numbers.

**Q705: How to work with command-line arguments using argparse?**
Define parser, add arguments (`add_argument`), and parse (`parse_args()`).

**Q706: How do you create a timer or stopwatch using time module?**
Record `start = time.time()` and then `elapsed = time.time() - start`.

**Q707: A PDF or Excel sheet of all 500?**
(Meta-question: Skipping).

**Q708: What is the default return value of a function with no return statement?**
`None`.

**Q709: What is the difference between exit() and sys.exit()?**
*   `sys.exit()`: Raises `SystemExit` exception (proper way for scripts).
*   `exit()`: Helper for interactive shell (should not be used in scripts).

**Q710: What happens when you reassign a built-in function (e.g., len = 5)?**
You shadow the built-in function in the current namespace. Calling `len()` afterwards will fail.

---

## 🔹 72. Internals & Comparisons (Questions 711-720)

**Q711: Why is Python called a “batteries included” language?**
Because it comes with a rich standard library that covers almost everything (Email, HTTP, CSV, JSON, Threads, XML, etc.) without needing external installations.

**Q712: How can you tell if a script is being run directly or imported?**
Check `if __name__ == "__main__":`.

**Q713: How do you check if two variables point to the same object?**
`x is y`.

**Q714: Can different types compare as equal in Python?**
Yes. `1 == 1.0` (Int and Float can be equal).

**Q715: What is interning in Python (for strings/ints)?**
Python caches small integers (-5 to 256) and some strings to save memory. Reuses the same object.

**Q716: What is Ellipsis (...) in Python?**
A singleton object used in slicing (NumPy) or as a placeholder in stubs (Type Hinting).

**Q717: What is a sentinel value?**
A unique object used to indicate "missing" or "end" states, distinguishable from any valid data (e.g., `object()`).

**Q718: What is the result of "5" + str(5)?**
`"55"`.

**Q719: How do you split a string on multiple delimiters?**
Use `re.split('[;,]', string)`.

**Q720: How to count all vowels in a string?**
`sum(1 for char in s if char.lower() in 'aeiou')`.

---

## 🔹 73. Algorithms & Collections (Questions 721-730)

**Q721: How do you check if all characters are alphanumeric?**
`s.isalnum()`.

**Q722: What is str.casefold() used for?**
Aggressive lowercasing for caseless matching (handles German 'ß' -> 'ss'). Better than `lower()`.

**Q723: How to filter out negative numbers from a list?**
`[x for x in lst if x >= 0]`.

**Q724: How to merge two lists element-wise?**
`[a+b for a,b in zip(l1, l2)]`.

**Q725: What does *list do in argument passing?**
Unpacks the list into positional arguments.

**Q726: How to count how many times each element appears in a list?**
`collections.Counter(lst)`.

**Q727: How to remove elements at even indices?**
`lst[1::2]` (Keep odd indices) or `del lst[::2]`.

**Q728: Why are tuples hashable but lists are not?**
Tuples are immutable; their content doesn't change, so their hash is stable. Lists can change.

**Q729: How can tuples be used as keys in dictionaries?**
Since they are hashable, `d = {(1, 2): "value"}` works fine.

**Q730: How do you convert a tuple to a CSV row?**
`",".join(map(str, tup))`.

---

## 🔹 74. Dictionary & Set Deep Dive (Questions 731-740)

**Q731: Can a tuple contain a list? Is it still hashable?**
Yes, it can contain a list `([1], 2)`. But it becomes **unhashable** and cannot be a dict key.

**Q732: What is unpacking with * in tuple assignment?**
`a, *b, c = (1, 2, 3, 4)`. `a=1`, `b=[2, 3]`, `c=4`.

**Q733: What happens if you del a key that doesn’t exist?**
`KeyError`.

**Q734: How do you use dict.setdefault() to group data?**
`d.setdefault(key, []).append(val)`. Useful for grouping items into lists.

**Q735: How can you use a dictionary to simulate a switch-case?**
`func = cases.get(value, default_func); func()`.

**Q736: How to get the N largest values in a dictionary?**
`heapq.nlargest(N, d, key=d.get)`.

**Q737: What is a ChainMap in collections?**
A wrapper that treats multiple dictionaries as a single mapping (lookups search all, writes go to first).

**Q738: Can you convert a set back to a list while preserving uniqueness?**
Yes, but order is undefined. `list(set_obj)`.

**Q739: What happens if you add a duplicate to a set?**
Nothing. The set remains unchanged (no errors).

**Q740: How to get only elements that appear once from a list?**
Use `Counter` and filter where count is 1. Or `set` operations if uniqueness logic differs.

---

## 🔹 75. Functional Programming (Questions 741-750)

**Q741: What’s the time complexity of set lookup?**
O(1) on average.

**Q742: Can sets be used for membership testing in strings?**
Yes. `char in set_of_chars` is faster than `char in string` for many lookups.

**Q743: What is a function attribute in Python?**
Functions are objects; you can attach arbitrary attributes to them. `func.my_attr = 5`.

**Q744: What happens if you pass fewer arguments than expected?**
`TypeError: missing required positional argument`.

**Q745: What is a partial function?**
A new function with some arguments of the original function pre-filled (`functools.partial`).

**Q746: What is currying in Python?**
Transforming a function `f(a, b)` into `f(a)(b)`. Python doesn't support this natively but it can be implemented with nested lambdas/functions.

**Q747: Can a function reference itself recursively?**
Yes.

**Q748: Can lambda functions use default arguments?**
Yes. `lambda x, y=1: x+y`.

**Q749: How to use lambda with sorted()?**
`sorted(data, key=lambda x: x['age'])`.

**Q750: Can you assign a lambda to a variable name?**
Yes. `add = lambda x, y: x+y`. (Though `def` is preferred for named functions).

---

## 🔹 76. Comprehensions & Errors (Questions 751-760)

**Q751: How to mimic ternary operator using lambda?**
`lambda x: 'yes' if x else 'no'`.

**Q752: Why are lambda functions limited to single expressions?**
To enforce readability and prevent complex logic from being buried in one-liners.

**Q753: How to transpose a matrix using list comprehension?**
`[[row[i] for row in matrix] for i in range(len(matrix[0]))]`.

**Q754: Can you use nested if in list comprehension?**
Yes. `[x for x in data if cond1 if cond2]`.

**Q755: How to apply a function to every element of a 2D list?**
`[[func(x) for x in row] for row in matrix]`.

**Q756: Can you use else if in list comprehension?**
Yes, in the expression part. `[a if c1 else b if c2 else c for ...]` (Ternary chaining).

**Q757: How to use list comprehension to get only uppercase strings?**
`[s for s in lst if s.isupper()]`.

**Q758: How to iterate through a string in reverse?**
`for char in reversed(s):`.

**Q759: What is loop unrolling?**
Optimization technique (manual or compiler) to decrease loop overhead by executing multiple iterations at once. Not common in standard Python code.

**Q760: How to flatten a list of lists using a loop?**
`flat = []; for sub in lst: flat.extend(sub)`.

---

## 🔹 77. Exception & OOP Advanced (Questions 761-770)

**Q761: Can you use for...else with break correctly?**
Yes. `break` skips the `else`. Use it to find an item; `else` handles "not found".

**Q762: How to use loop with both index and value?**
`enumerate()`.

**Q763: How to catch any exception?**
`except Exception:` (Catches all standard errors) or `except:` (Catches everything including exit signals - dangerous).

**Q764: What is exception chaining in Python?**
When raising a new exception while handling another, Python attaches the original trace (`__context__`). `raise NewErr from OldErr`.

**Q765: Can you write multiple except clauses for the same block?**
Yes.

**Q766: How to log exceptions instead of printing?**
`logging.exception("Message")` (automatically captures traceback).

**Q767: What is the performance impact of exception-heavy code?**
Exceptions are slightly slower than `if` checks if raised frequently. "Look Before You Leap" (LBYL) can be faster than "Easier to Ask Forgiveness" (EAFP) in tight loops if failures are common.

**Q768: What is class variable vs instance variable?**
*   Class var: Shared by all instances (defined in class body).
*   Instance var: Unique to instance (defined in `__init__`).

**Q769: How to count the number of objects of a class?**
Increment a class variable in `__init__` and decrement in `__del__`.

**Q770: What is __slots__ used for?**
(Refresher) Restricting attributes to save memory.

---

## 🔹 78. Decorators & Files II (Questions 771-780)

**Q771: What is the output of object.__class__?**
Returns the class (type) of the object.

**Q772: Can a class have no attributes or methods?**
Yes. `class Empty: pass`.

**Q773: How to write a timing decorator?**
Record time before/after func call and print diff.

**Q774: Can you decorate class methods?**
Yes.

**Q775: What’s the purpose of functools.wraps()?**
Copies metadata (name, docstring) from the original function to the wrapper logic in a decorator.

**Q776: Can a decorator modify function arguments?**
Yes, by intercepting `*args` in the wrapper before calling the function.

**Q777: Can a decorator be class-based?**
Yes. Implement `__init__` (to accept function) and `__call__` (to handle execution).

**Q778: How to detect file encoding?**
Use libraries like `chardet`. Python cannot definitively guess encoding without reading content first.

**Q779: How to safely read large files without memory issues?**
Read line by line or in chunks (`read(size)`).

**Q780: How do you append data to an existing file?**
Mode `'a'`.

---

## 🔹 79. Randomness & Datetime (Questions 781-790)

**Q781: How to read only first N lines of a file?**
`head = [next(f) for _ in range(N)]`.

**Q782: What is tell() used for in file reading?**
Returns current cursor position (byte offset).

**Q783: How to shuffle a list?**
`random.shuffle(lst)` (in-place).

**Q784: How to get a random float between 1 and 5?**
`random.uniform(1, 5)`.

**Q785: How to seed the random generator?**
`random.seed(value)`. Ensures reproducibility.

**Q786: What is the difference between random.choice() and random.sample()?**
*   `choice`: Picks one item (replacement allowed if called multiple times).
*   `sample`: Picks k unique items (no replacement).

**Q787: How to generate random strings?**
`''.join(random.choices(string.ascii_letters, k=10))`.

**Q788: How to get current timestamp in Python?**
`time.time()` (Unix epoch) or `datetime.now().timestamp()`.

**Q789: How to convert timestamp to datetime?**
`datetime.fromtimestamp(ts)`.

**Q790: How to find the difference between two dates?**
Subtract them. Returns a `timedelta` object.

---

## 🔹 80. Environment & Modules II (Questions 791-800)

**Q791: How to get today's date?**
`date.today()`.

**Q792: How to add days to a datetime object?**
`dt + timedelta(days=5)`.

**Q793: How to get environment variables in Python?**
`os.environ.get('VAR_NAME')`.

**Q794: How to exit a script gracefully?**
`sys.exit(0)`.

**Q795: What’s the use of pdb module?**
Interactive source code debugger.

**Q796: How to get the path of the current file?**
`__file__` (inside a script).

**Q797: How to handle keyboard interrupt in Python?**
Catch `KeyboardInterrupt` (Ctrl+C).

**Q798: What is a context manager?**
Object defining runtime context (setup/teardown). Used in `with` statements.

**Q799: What does with statement do internally?**
Calls `__enter__` on start and `__exit__` on end/error.

**Q800: What is a weak reference?**
Reference that doesn't prevent the object from being garbage collected (`weakref` module).

---

## 🔹 81. Python Deep Dive & Internals (Questions 801-810)

**Q801: What’s the difference between yield and return?**
(Refresher) Yield pauses function (generator). Return exits it.

**Q802: What is a module-level dunder variable?**
Variables like `__name__`, `__file__`, `__doc__`, `__package__` defined globally in a module.

**Q803: What is memoization?**
Caching the return value of a function based on its input arguments to avoid repeating calculations.

**Q804: How to avoid circular imports?**
Import inside functions/methods (lazy import) or restructure code to remove the cycle.

**Q805: What is the purpose of assert?**
(Refresher) Development-time debugging check.

**Q806: What are f-strings and why are they preferred?**
(Refresher) Fast, readable string interpolation.

**Q807: 🧠 Flashcards for spaced repetition?**
(Meta-question: Skipping).

**Q808: 📄 PDF/CSV/Excel format of all 600?**
(Meta-question: Skipping).

**Q809: 🚀 Continue to Batch 7: 100 more questions?**
(Meta-question: Skipping).

**Q810: Why is indentation preferred over braces in Python?**
It forces readability and eliminates "brace wars". Code structure must match visual structure.

---

## 🔹 82. Interpreters & Scope (Questions 811-820)

**Q811: What are some limitations of Python as a language?**
*   **Speed:** Slower than C/C++.
*   **Mobile:** Weak native mobile apps.
*   **Threading:** GIL limits CPU-bound concurrency.

**Q812: What is the role of the Python Virtual Machine (PVM)?**
The engine that iterates over bytecode instructions and executes them.

**Q813: What is the lifecycle of a Python script?**
Parser -> AST -> Compiler -> Bytecode -> PVM -> Execution.

**Q814: What is a token in Python?**
The smallest unit of the language (identifier, keyword, literal, operator) produced by the tokenizer.

**Q815: What happens if you use a variable before assigning a value?**
`NameError: name 'x' is not defined`.

**Q816: Can you assign values to multiple variables in one line?**
Yes. `a, b = 1, 2` or `a = b = 1`.

**Q817: What is variable hoisting? Does Python support it?**
Python does not support hoisting (unlike JS). Variables must be assigned before use.

**Q818: Is Python pass-by-value or pass-by-reference?**
It is **Pass-by-Object-Reference** (or Pass-by-Assignment).
*   Mutables: Changes affect caller.
*   Immutables: Changes create new local object (don't affect caller).

**Q819: What does unpacking a, *b, c = [1,2,3,4,5] do?**
`a=1`, `c=5`, `b=[2,3,4]` (b absorbs the middle).

**Q820: How is range different from list(range())?**
`range()` is a generator-like object (lazy, O(1) memory). `list(range())` stores all numbers in memory (O(n)).

---

## 🔹 83. Details of Data Types (Questions 821-830)

**Q821: What is the output of None == False?**
`False`.

**Q822: What happens when you compare a string with an integer?**
In Python 3: `TypeError` for ordered comparison (`<`, `>`). `False` for equality (`==`).

**Q823: Are custom objects hashable by default?**
Yes. They use their `id()` as hash. (Unless `__eq__` is overridden without `__hash__`).

**Q824: Can you sort a list of mixed types?**
No (in Python 3). `[1, "a"].sort()` raises `TypeError`.

**Q825: What is string interning?**
(Refresher) Caching unique string objects.

**Q826: What is the use of str.format_map()?**
Similar to `format()` but takes a mapping (dict) directly. `"{name}".format_map({'name': 'John'})`.

**Q827: What happens if you multiply a string by a negative number?**
Returns empty string `""`.

**Q828: Can strings be concatenated using join() on numbers?**
No. `join` expects strings. Must use `map(str, numbers)`.

**Q829: How to find all occurrences of a substring?**
Use `re.finditer()` or a loop with `str.find()`.

**Q830: How to chunk a list into N-sized groups?**
`[lst[i:i + n] for i in range(0, len(lst), n)]`.

---

## 🔹 84. Performance & Collections (Questions 831-840)

**Q831: What happens if you modify a list while looping over it?**
(Refresher) Unsafe/Unpredictable.

**Q832: How does list.index() behave on duplicates?**
Returns iteration index of the **first** match.

**Q833: How do you find the difference between two lists?**
`list(set(l1) - set(l2))` (Order lost) or list comprehension (Order kept).

**Q834: What is the output of [[]] * 3?**
`[[], [], []]`. All three internal lists refer to the **same** object. modifying one modifies all.

**Q835: Are tuples always faster than lists?**
Usually yes, due to simpler allocation and no need for resizing logic.

**Q836: Can you convert a list of tuples to a dictionary?**
Yes. `dict([(k, v), ...])`.

**Q837: What does tuple(x for x in iterable) do?**
Creates a tuple from a generator.

**Q838: Can tuple slicing produce a new object?**
Yes, unless it slices the whole tuple `t[:]` (returns same object).

**Q839: How to reverse a tuple?**
`t[::-1]` (Returns new tuple).

**Q840: How to invert a dictionary with duplicate values?**
Result needs lists of keys. `{v: [k for k in d if d[k]==v] for v in set(d.values())}`.

---

## 🔹 85. Sets & Efficiency (Questions 841-850)

**Q841: What is the fastest way to check if a key exists?**
`key in dict` (O(1)).

**Q842: Can you use a class instance as a dictionary key?**
Yes, if it is hashable.

**Q843: What happens if you update a dictionary with None?**
`d.update(None)` might raise `TypeError` depending on arg. `d[key] = None` is valid.

**Q844: What is dict comprehension with filter?**
`{k:v for k,v in d.items() if v > 0}`.

**Q845: How does set handle case-sensitive strings?**
"Apple" and "apple" are distinct different elements.

**Q846: What is the output of set("banana")?**
`{'b', 'a', 'n'}` (Unique chars).

**Q847: How do you test if two sets are disjoint?**
`s1.isdisjoint(s2)`.

**Q848: Can sets be compared using ==, <, >?**
Yes. `==` (Equality), `<` (Proper subset), `>` (Proper superset).

**Q849: How to remove multiple elements from a set at once?**
`s.difference_update(elements)`.

**Q850: What happens if you return inside a loop in a function?**
The function terminates immediately.

---

## 🔹 86. Closures & Lambdas (Questions 851-860)

**Q851: Can a function change a mutable argument?**
Yes. `def f(l): l.append(1)` changes the list passed by caller.

**Q852: Can a function return another function?**
Yes (Closures/Decorators).

**Q853: What is a closure variable?**
A variable in the enclosing scope that is referenced by the inner function.

**Q854: How does function scoping work in Python?**
LEGB (Local, Enclosing, Global, Builtin).

**Q855: Can lambdas use external variables?**
Yes (Closures). But be careful with loops (late binding).

**Q856: Can lambdas contain function calls?**
Yes. `lambda x: print(x)`.

**Q857: Why are lambdas limited to one expression?**
To force code clarity.

**Q858: What is the purpose of lambda in sorting?**
To modify the sort key (e.g., sort by length, or second item).

**Q859: When should you avoid using lambda?**
When the logic is complex or needs checking. Use a named function instead.

**Q860: What does zip(*list_of_tuples) do?**
Unzips the list (Transposes matrix).

---

## 🔹 87. Iterators & Exceptions II (Questions 861-870)

**Q861: What’s the difference between zip() and enumerate()?**
*   `zip`: Combines multiple lists.
*   `enumerate`: Adds index to one list.

**Q862: How do you use map() with multiple iterables?**
`map(lambda x,y: x+y, l1, l2)`.

**Q863: Can filter() return None?**
No, it returns an iterator.

**Q864: What happens when you zip iterables of unequal length?**
It truncates to the shortest.

**Q865: What’s the difference between iterating with index vs value?**
*   `for x in lst`: Cleaner.
*   `for i in range(len(lst))`: Needed if you need to modify list index.

**Q866: How to loop through two lists at once?**
`zip(l1, l2)`.

**Q867: How does range(len(list)) compare to enumerate()?**
`enumerate()` is more Pythonic and readable.

**Q868: Can you break only the inner loop of a nested structure?**
Yes, using `break`. To break both, use a flag or return.

**Q869: How do you create a progress bar while looping?**
Use `tqdm` library or print `\r` characters.

**Q870: What is the difference between try-except and try-finally?**
*   `except`: Handles error.
*   `finally`: Cleans up (runs always).

---

## 🔹 88. Classes & OOP Deep Dive (Questions 871-880)

**Q871: What is the result of an uncaught exception in Python?**
Program crashes with traceback.

**Q872: What does except Exception as e: give you?**
The exception object instance `e`.

**Q873: How to raise exceptions with a custom message?**
`raise ValueError("Custom Message")`.

**Q874: How do you catch multiple exception types together?**
`except (E1, E2):`.

**Q875: Can you instantiate a class without __init__()?**
Yes, if `__init__` is not defined (default one used) or via `__new__`.

**Q876: How to check if an object is an instance of a subclass?**
`isinstance(obj, ParentClass)` returns True for subclass instances too.

**Q877: How to override the equality operator?**
Define `__eq__(self, other)`.

**Q878: What’s the purpose of __new__()?**
To create the instance.

**Q879: What is the difference between classmethod and staticmethod?**
(Refresher) `cls` access vs no access.

**Q880: How to apply a decorator to a class?**
`@decorator class MyClass:`.

---

## 🔹 89. Files & Decorators (Questions 881-890)

**Q881: What is the difference between @property and a method?**
`@property` is accessed like a variable `obj.prop`. Method needs calls `obj.method()`.

**Q882: Can decorators access and change arguments?**
Yes. Use `args` in the wrapper.

**Q883: What is a memoization decorator?**
Caches results. `functools.lru_cache`.

**Q884: How do you undo a decorator effect?**
Usually impossible unless the decorator attaches the original function (e.g., `func.__wrapped__`).

**Q885: How to list all files in a directory with .txt extension?**
`glob.glob("*.txt")`.

**Q886: What’s the difference between read() and readline()?**
*   `read()`: All content.
*   `readline()`: Next line.

**Q887: How do you write binary data to a file?**
`open('file', 'wb').write(b'data')`.

**Q888: What happens when a file is opened in write mode but already exists?**
Content is truncated (erased).

**Q889: How to safely delete a file?**
`if os.path.exists(f): os.remove(f)`.

**Q890: What does from module import * do?**
Imports all public names from module into current namespace. (Generally discouraged).

---

## 🔹 90. System & Stdlib (Questions 891-900)

**Q891: What’s the difference between __import__() and import?**
`import`: Statement. `__import__()`: Function (used for dynamic imports, low-level).

**Q892: Can you reload a module at runtime?**
`importlib.reload()`.

**Q893: How do you make a directory a package?**
Add `__init__.py`.

**Q894: What does math.ceil() do?**
Rounds up to nearest integer. `3.1` -> `4`.

**Q895: What’s the use of collections.OrderedDict()?**
(Historical) Dict that remembers keys. Default `dict` does this now in Py3.7+.

**Q896: What is the difference between os and sys modules?**
*   `os`: Operating System interaction (files, env).
*   `sys`: Python Runtime/Interpreter interaction (modules, path, argv).

**Q897: How can itertools.cycle() be used?**
Iterates infinite cycle `[1,2,1,2...]`. Good for "round robin".

**Q898: What does shutil.rmtree() do?**
Deletes directory and all contents recursively.

**Q899: How to add a simple logger to your Python script?**
`logging.basicConfig(level=logging.INFO); logging.info("Msg")`.

**Q900: What does __traceback__ provide in exceptions?**
Access to the stack trace object attached to the exception.

---

## 🔹 91. Profiling & Performance II (Questions 901-910)

**Q901: What’s the use of %timeit in Python?**
A magic command in IPython/Jupyter to time the execution of a statement/expression (runs multiple loops for accuracy).

**Q902: How to check memory usage of a variable?**
`sys.getsizeof(var)`.

**Q903: What are common performance bottlenecks in Python scripts?**
Excessive loops, incorrect data structures (list vs set lookup), global variable lookups, and IO blocking.

**Q904: How to convert a datetime to a string?**
`dt.strftime(format_string)`.

**Q905: How to convert a string to a datetime object?**
`datetime.strptime(date_string, format_string)`.

**Q906: How to get UTC time in Python?**
`datetime.now(timezone.utc)`.

**Q907: What is the difference between datetime.now() and datetime.utcnow()?**
*   `now()`: Local time.
*   `utcnow()`: UTC time (Naive, no timezone info). *Deprecated in Py3.12+*, prefer `now(utc)`.

**Q908: How to get the start of the day using datetime?**
`dt.replace(hour=0, minute=0, second=0, microsecond=0)` or `date.today()`.

**Q909: 🔄 Batch 8 with 100 more?**
(Meta-question: Skipping).

**Q910: 📘 All 700 in a structured format?**
(Meta-question: Skipping).

---

## 🔹 92. Concepts & Internals III (Questions 911-920)

**Q911: 💡 Begin answer+explanation?**
(Meta-question: Skipping).

**Q912: What does "everything is an object" mean in Python?**
Functions, classes, modules, and basic types (int, str) are all objects (instances of classes) that can be assigned to variables or passed as arguments.

**Q913: What is dynamic name binding?**
Variables are just names pointing to objects. You can rebind a name to a different object (even different type) at any time.

**Q914: What is the purpose of the Python GIL?**
(Refresher) To protect internal memory structures from concurrent access by multiple threads.

**Q915: How are Python functions first-class objects?**
They can be assigned to variables, passed as arguments, returned from other functions, and stored in data structures.

**Q916: What does Python mean by being an "interpreted" language?**
Source code is not compiled to native binary code. It is executed by a specific runtime (interpreter).

**Q917: What are valid characters in a Python variable name?**
Letters (a-z, A-Z), numbers (0-9), and underscore (_). Cannot start with a number.

**Q918: Can variable names be the same as keywords?**
No. `class = 1` is a SyntaxError.

**Q919: What does a trailing underscore in a variable name indicate (e.g. class_)?**
Used to avoid conflict with Python keywords (e.g., `class_`, `id_`).

**Q920: What is the difference between _var, __var, and __var__?**
*   `_var`: Internal use convention.
*   `__var`: Name mangled (private-ish).
*   `__var__`: System-defined names (Magic methods).

---

## 🔹 93. Syntax & Strings II (Questions 921-930)

**Q921: What is name mangling in Python?**
Interpreter rewrites `__var` to `_ClassName__var` to prevent accidental overriding in subclasses.

**Q922: What is the difference between int, float, and Decimal?**
*   `int`: Arbitrary precision integers.
*   `float`: IEEE 754 double precision (approximate).
*   `Decimal`: Exact decimal representation (financial).

**Q923: How does Python handle very large integers?**
Automatically promotes them to arbitrary-precision integers (limited only by memory).

**Q924: What’s the result of -5 // 2?**
`-3`. (Floor division moves towards negative infinity).

**Q925: How to convert a float to an integer safely?**
`math.floor()`, `math.ceil()`, or `int()` (truncates).

**Q926: What does string[::-1] do?**
Reverses the string.

**Q927: How do you remove all whitespace from a string?**
`"".join(s.split())` or regex.

**Q928: What is the translate() method used for?**
Mapping characters to other characters (or None to delete). Efficient replacement.

**Q929: How do you replace multiple substrings in one pass?**
Use regex `re.sub()` or chain `replace()`.

**Q930: What is the purpose of the maketrans() function?**
Creates the translation table for `translate()`.

---

## 🔹 94. Sorting & Tuples (Questions 931-940)

**Q931: How to sort a list of tuples by the second value?**
`lst.sort(key=lambda x: x[1])`.

**Q932: How does the sort() method behave with custom objects?**
Uses comparison operators (`<`). You must define `__lt__` in your class.

**Q933: What’s the difference between x = y.copy() and x = y[:]?**
For lists, they are identical (shallow copy). `.copy()` is more readable (introduced Py3.3).

**Q934: How to get unique items from a list while preserving order?**
`list(dict.fromkeys(lst))`.

**Q935: How to get the frequency of each item in a list?**
`collections.Counter(lst)`.

**Q936: How are tuples used in function return values?**
To return multiple values. Python packs them into a single tuple automatically.

**Q937: What is the difference between tuple() and ()?**
Both create an empty tuple. `()` is syntax literal (faster).

**Q938: Are tuples safer than lists?**
Yes, "write-protected" data prevents accidental modification.

**Q939: What is the tuple packing and unpacking concept?**
*   Packing: `t = 1, 2`.
*   Unpacking: `x, y = t`.

**Q940: Can you pass a tuple as arguments using *?**
Yes. `func(*tup)` unpacks tuple into positional args.

---

## 🔹 95. Dicts & Sets III (Questions 941-950)

**Q941: What is dictionary unpacking using ** in function calls?**
`func(**d)` unpacks dict into keyword args.

**Q942: How to count the frequency of words using a dictionary?**
Loop and increment: `d[word] = d.get(word, 0) + 1`.

**Q943: How does dict.get() differ from bracket-based access?**
`get()` handles missing keys gracefully (returns None/default). Brackets raise Error.

**Q944: Can keys be complex objects like tuples or frozensets?**
Yes, because they are immutable and hashable.

**Q945: What’s the purpose of dict.items()?**
Iterates over key-value pairs.

**Q946: How does Python remove duplicates from a list using sets?**
`list(set(lst))`.

**Q947: What is the result of set([1, 2, 2, 3])?**
`{1, 2, 3}`.

**Q948: How to check if one set is a strict subset of another?**
`s1 < s2` (Subset AND not equal).

**Q949: How to find common elements across multiple sets?**
`set.intersection(*list_of_sets)`.

**Q950: How do sets affect performance when checking membership?**
O(1) (Constant time) vs List O(n) (Linear time). Huge speedup for large collections.

---

## 🔹 96. Functions & Docstrings (Questions 951-960)

**Q951: What are *args and **kwargs used for?**
(Refresher) Variable positional and keyword arguments.

**Q952: What is a docstring and how is it used?**
Documentation at the top of function/class. Accessed via `help(obj)` or `obj.__doc__`.

**Q953: How to provide a default value for a function parameter?**
`def func(arg=default):`.

**Q954: Can default values be mutable? Why should you avoid that?**
Yes, but they are evaluated only once at definition. Modifying them persists across calls (the "mutable default argument" trap).

**Q955: How do you annotate function arguments and return types?**
`def func(a: int) -> str:`.

**Q956: What is the limitation of lambda for multiline logic?**
Syntax prohibits it. Lambdas are strictly single expressions.

**Q957: Can you store a lambda in a dictionary?**
Yes. `actions = {'add': lambda x,y: x+y}`. Allows data-driven logic dispatch.

**Q958: What happens if you use a loop variable inside a lambda?**
It captures the variable *reference*, not value. All lambdas might use the final loop value. Fix: `lambda x=x: x`.

**Q959: Can you write a lambda inside another function?**
Yes.

**Q960: What’s the practical difference between a named function and a lambda?**
Name (debugging), docstrings, capability (multiline). Lambda is just syntax sugar for simple cases.

---

## 🔹 97. Control Flow & Loops (Questions 961-970)

**Q961: What is the result of using break in nested loops?**
Only the inner loop breaks.

**Q962: How do you use continue inside list comprehensions?**
You can't use `continue` directly. Use filtering condition `if`.

**Q963: How can you create an infinite loop?**
`while True:`.

**Q964: How do you iterate with a custom step size?**
`range(start, stop, step)`.

**Q965: What is the output of for _ in range(3):?**
Loop runs 3 times. `_` indicates loop variable is ignored.

**Q966: What is the difference between syntax error and runtime error?**
*   **Syntax:** Code violates rules (parsing fails). Script doesn't run.
*   **Runtime:** Error during execution (e.g., ZeroDivision).

**Q967: How to write a try-except-else block?**
`try: code; except: handle; else: run_if_no_error`.

**Q968: Can you catch multiple exceptions in one line?**
`except (E1, E2):`.

**Q969: How to define a custom exception with arguments?**
Pass args to `super().__init__(args)`.

**Q970: What is the use of exception chaining?**
(Refresher) Preserving original cause context.

---

## 🔹 98. Classes & Class Magic (Questions 971-980)

**Q971: What is the default value of __str__() if not defined?**
Typically the object type and memory address.

**Q972: What happens if __init__() is missing?**
Parent class `__init__` is called (ultimately `object.__init__` which does nothing).

**Q973: How to define a class with private attributes?**
Use `__attr`.

**Q974: How can you make a class iterable?**
Implement `__iter__()`.

**Q975: How to create a singleton class in Python?**
Overrides `__new__` to return specific instance.

**Q976: Can you apply more than one decorator to a function?**
Yes.

**Q977: What does @property do in Python?**
Turns a method into a read-only attribute getter.

**Q978: How to write a class-based decorator?**
Class with `__call__` method.

**Q979: Can decorators accept optional arguments?**
Yes, but logic is complex (needs to check if called with function or arguments).

**Q980: What is a common use case for logging decorators?**
Tracing entry/exit of functions automatically.

---

## 🔹 99. Files & Resources (Questions 981-990)

**Q981: What is the context manager behavior in file operations?**
Closes file automatically when block exits.

**Q982: What happens if you forget to close a file?**
Resource leak. It might be closed by GC eventually, but not guaranteed immediately.

**Q983: How to read binary content as hexadecimal?**
`f.read().hex()`.

**Q984: What is the difference between a+ and r+ modes?**
(Refresher) `a+` appends. `r+` reads/writes from start (overwriting).

**Q985: How to create a temporary file that deletes automatically?**
`tempfile.NamedTemporaryFile()`.

**Q986: How to calculate days between two dates?**
`(d2 - d1).days`.

**Q987: How to generate a date range?**
Loop adding `timedelta` or using pandas `date_range`.

**Q988: How do you localize a datetime object?**
`dt.replace(tzinfo=timezone.utc)` or `pytz.localize`.

**Q989: How to convert a string timestamp to a datetime object?**
`datetime.fromisoformat()` or `strptime()`.

**Q990: What is timezone-aware datetime?**
Date object that includes timezone info, allowing unambiguous comparison.

---

## 🔹 100. Introspection & Modules (Questions 991-1000)

**Q991: What is the use of id() function?**
Memory address identity.

**Q992: How to get all attributes of an object?**
`dir(obj)` or `obj.__dict__` (for instance attributes).

**Q993: What does __dict__ show?**
Dictionary containing the object's (writable) attributes.

**Q994: What’s the difference between __str__ and __repr__?**
(Refresher) Readable vs Unambiguous.

**Q995: How to inspect the call stack?**
`traceback.print_stack()` or `inspect.stack()`.

**Q996: What is the use of the importlib module?**
Dynamic importing mechanism.

**Q997: What is lazy import?**
Importing a module only when it is needed (inside a function), to speed up initial startup.

**Q998: How do you organize a Python project into modules?**
Foldering structure with `__init__.py`.

**Q999: What is relative import and how is it done?**
Importing from current package. `from . import module`.

**Q1000: What is the difference between module and package?**
Package is a special module (directory) that can contain other modules.

---

## 🔹 101. Modules, Packages & Clean Code (Questions 1001-1010)

**Q1001: What is lazy import?**
Importing a module inside a function so it's only loaded when the function is called, not at startup. Useful for reducing startup time.

**Q1002: How do you organize a Python project into modules?**
Group related functions into `.py` files. Group related files into directories (Packages). Use `__init__.py`.

**Q1003: What is relative import and how is it done?**
Use dot notation to import from current package. `from . import utils` (same dir) or `from .. import config` (parent dir).

**Q1004: What is the difference between module and package?**
*   **Module:** File containing Python code (e.g., `math.py`).
*   **Package:** Directory containing modules and `__init__.py`.

**Q1005: What are Pythonic ways to swap values?**
`a, b = b, a`.

**Q1006: What is the EAFP vs LBYL philosophy in Python?**
*   **EAFP:** Easier to Ask Forgiveness than Permission (Try/Except). *Preferred in Python*.
*   **LBYL:** Look Before You Leap (If/Else checks).

**Q1007: What is the Zen of Python?**
A collection of 19 guiding principles for writing computer programs in Python. access via `import this`.

**Q1008: How to avoid mutable default arguments?**
Use `None` as default. `def f(x=None): if x is None: x=[]`.

**Q1009: How to prevent circular imports?**
Refactor code to separate concerns or use delayed imports inside functions.

**Q1010: What is pip freeze used for?**
Lists installed packages and versions in requirements format. `pip freeze > requirements.txt`.

---

## 🔹 102. Dependency Management (Questions 1011-1020)

**Q1011: How to lock dependencies using pip-tools or poetry?**
*   **pip-tools:** `pip-compile`.
*   **Poetry:** `poetry lock`.
Generates a lock file with exact versions/hashes for reproducibility.

**Q1012: What is the purpose of a requirements.txt file?**
To list all dependencies so environments can be replicated.

**Q1013: How do you publish a Python package?**
Register on PyPI, build distribution (`sdist`/`wheel`), upload using `twine`.

**Q1014: How to install a package directly from GitHub?**
`pip install git+https://github.com/user/repo.git`.

**Q1015: 🗂 All batches compiled into one Notion table or PDF?**
(Meta-question: Skipping).

**Q1016: 🧠 Flashcard versions for spaced revision?**
(Meta-question: Skipping).

**Q1017: 🎯 Daily challenge mode: 5 questions + explanations?**
(Meta-question: Skipping).

**Q1018: What are Python's soft keywords?**
Keywords like `match`, `case`, `type` (in 3.12) that are context-sensitive. They can still be used as variable names in other contexts.

**Q1019: What is the purpose of the global keyword?**
Allows a function to modify a variable defined at the module level.

**Q1020: What is the effect of nonlocal inside nested functions?**
Allows modification of a variable in the nearest enclosing scope (not global).

---

## 🔹 103. Advanced Scope (Questions 1021-1030)

**Q1021: Can return and yield exist in the same function?**
Yes (in modern Python). Return implies the generator stops (raising StopIteration with value).

**Q1022: What is a compound statement in Python?**
A statement that contains other statements (e.g., `if`, `while`, `def`, `class`, `with`).

**Q1023: What happens when a variable is declared in both global and local scopes?**
The local variable shadows the global one inside the function.

**Q1024: Can variables inside a loop leak into the outer scope?**
Yes (in Python 3, loop variables in `for`-loops differ from comprehensions).
`for i in range(3): pass; print(i)` prints `2`.

**Q1025: How does Python resolve variable scope with closures?**
The inner function keeps a reference to the variables in the outer function's stack frame (cell objects).

**Q1026: What is late binding in Python functions?**
Closures look up variables at call time, not definition time. Common issue in loops with lambdas.

**Q1027: What is variable aliasing?**
Two different names pointing to the same object in memory. `a = []; b = a`.

**Q1028: What is the result of "hello".capitalize()?**
`"Hello"` (First char upper, rest lower).

**Q1029: How does str.zfill(width) work?**
Pads string with zeros on the left.

**Q1030: What does "123".isdecimal() return?**
`True`. (Strictly checking if characters are decimals 0-9).

---

## 🔹 104. String tricks (Questions 1031-1040)

**Q1031: How to safely concatenate potentially None strings?**
`f"{str(a or '')}{str(b or '')}"`.

**Q1032: What’s the use of .partition()?**
Splits string at first occurrence of separator into 3-tuple `(before, sep, after)`. Safe if separator is missing.

**Q1033: How to remove duplicates without changing order?**
`list(dict.fromkeys(lst))`.

**Q1034: What’s the difference between reversed(list) and list[::-1]?**
*   `reversed()`: Returns an iterator (memory efficient).
*   `[::-1]`: Creates a new reversed list in memory.

**Q1035: What happens if you slice a list beyond its length?**
Nothing bad. It returns whatever exists (or empty list). No IndexError.

**Q1036: How do you insert an item in the middle of a list?**
`lst.insert(len(lst)//2, val)`.

**Q1037: How to update list values in-place?**
Loop with index `lst[i] = new_val` or slicing `lst[:] = new_values`.

**Q1038: How can tuples be used in dictionaries?**
As keys (since they are hashable).

**Q1039: What does *args, = mytuple mean?**
Unpacking everything into a list called `args`. (Equivalent to `args = list(mytuple)`).

**Q1040: What is the tuple result of splitting a string?**
`tuple("a b".split())` -> `('a', 'b')`.

---

## 🔹 105. Operations (Questions 1041-1050)

**Q1041: How do you “zip” and “unzip” a tuple of tuples?**
*   Zip: `zip(l1, l2)`
*   Unzip: `zip(*zipped_list)`

**Q1042: Can you use tuple as a stack?**
Technically yes (immutable stack?), but practically no because you can't push/pop. Use `list`.

**Q1043: How to safely get a nested key in a dict?**
`d.get('k1', {}).get('k2')`.

**Q1044: How to convert a list of pairs into a dictionary?**
`dict(list_of_pairs)`.

**Q1045: What happens when dictionary keys are not unique?**
The last inserted value overwrites previous ones.

**Q1046: How does dictionary hashing work?**
Keys are hashed (`hash(key)`). The hash determines the bucket (slot) in the internal hash table.

**Q1047: How to find symmetric difference between sets?**
`s1 ^ s2`.

**Q1048: What’s the output of set('abcabc')?**
`{'a', 'b', 'c'}`.

**Q1049: Can you compare sets with >= or <=?**
Yes. Checks superset/subset relationship.

**Q1050: What does set.discard() return?**
`None`. (Removes element if present, no error if missing).

---

## 🔹 106. Advanced Functions (Questions 1051-1060)

**Q1051: How to perform set operations on lists?**
Convert to sets `set(l1) & set(l2)`, output to list.

**Q1052: How does Python resolve ambiguity in function overloading?**
It doesn't. Last definition wins.

**Q1053: What happens if you mutate a parameter default value?**
It persists. The default object is created once at definition time.

**Q1054: How to document functions using annotations?**
Type hints `def f(x: int):` are stored in `f.__annotations__`.

**Q1055: Can a function modify a global variable?**
Yes, if declared `global var_name` first.

**Q1056: How do you use a lambda to sort strings by length?**
`s.sort(key=lambda x: len(x))`.

**Q1057: How can lambdas help in functional composition?**
Passing small logic chunks to higher-order functions like `map`, `reduce`.

**Q1058: Can you simulate switch-case using lambda and dict?**
`cases = {'a': lambda: 1, 'b': lambda: 2}; cases.get(val, lambda: 0)()`.

**Q1059: How do lambdas behave inside list comprehensions?**
Normal function behavior.

**Q1060: Can lambda expressions be used with map/filter inside class methods?**
Yes.

---

## 🔹 107. Comprehensions Deep Dive (Questions 1061-1070)

**Q1061: How to use dictionary comprehension to invert a dict?**
`{v: k for k,v in d.items()}`.

**Q1062: How to write a set comprehension to get unique vowels?**
`{c for c in string if c in 'aeiou'}`.

**Q1063: Can you create nested list comprehensions?**
Yes. `[[x for x in sub] for sub in matrix]`.

**Q1064: How to apply condition inside a set comprehension?**
`{x for x in s if predicate(x)}`.

**Q1065: How to convert list of dicts into one dict using comprehension?**
`{k:v for d in list_of_dicts for k,v in d.items()}`.

**Q1066: What is the result of looping over a dict.keys()?**
Iterates keys (Same as looping over the dict object itself).

**Q1067: What happens if you break an infinite loop?**
Execution continues after the loop.

**Q1068: How to iterate over a dictionary in sorted key order?**
`for k in sorted(d):`.

**Q1069: How to simultaneously loop over two lists with index?**
`for i, (a, b) in enumerate(zip(l1, l2)):`.

**Q1070: How does Python treat else in loops?**
Executes if loop finishes **without** hitting `break`.

---

## 🔹 108. Exception Nuances (Questions 1071-1080)

**Q1071: What is the difference between raise and raise e?**
*   `raise`: Re-raises the *active* exception (preserving context).
*   `raise e`: Raises a specific exception object `e` (traceback starts here).

**Q1072: What’s the use of contextlib.suppress()?**
Gracefully ignores specified exceptions. `with suppress(FileNotFoundError): os.remove(f)`.

**Q1073: What happens if exception occurs in finally block?**
It overrides previous exceptions/returns. The new exception propagates.

**Q1074: Can exception be re-raised?**
Yes, using `raise`.

**Q1075: How to write a fail-safe try-except block for file I/O?**
Use `try...except IOError...finally` or just `with open()`.

**Q1076: What is class inheritance and how is it defined?**
`class Child(Parent):`. Mechanism for code reuse/specialization.

**Q1077: How can you override a method in a subclass?**
Define method with same name.

**Q1078: What is super().__init__() and why use it?**
Calls parent constructor to ensure parent initialization logic runs.

**Q1079: Can you have multiple constructors in Python?**
No. But `__init__` can handle vary args, or use `@classmethod` factory methods.

**Q1080: What is the use of __call__() in a class?**
Values instances callable like functions. `obj()`.

---

## 🔹 109. Decorator and File IO (Questions 1081-1090)

**Q1081: What’s the benefit of using a logging decorator?**
Separation of concerns. Don't clutter business logic with logging calls.

**Q1082: Can a decorator modify function metadata?**
Yes, but `functools.wraps` prevents that (restores generic medata).

**Q1083: How to apply a decorator to all methods of a class?**
Class decorator that iterates attributes and wraps callables, or a metaclass.

**Q1084: What’s the difference between @classmethod and @staticmethod?**
*   Classmethod: receives `cls`. Factory methods.
*   Staticmethod: plain function in class namespace. Utilities.

**Q1085: How to debug behavior inside a decorator?**
Print/Log inside the wrapper, or use debugger stepping into the wrapper.

**Q1086: How to read JSON from a file?**
`json.load(f)`.

**Q1087: How to read a file and skip blank lines?**
`[line for line in f if line.strip()]`.

**Q1088: What’s the difference between seek() and tell()?**
*   `seek(pos)`: Move cursor.
*   `tell()`: Get cursor position.

**Q1089: How to check if a file exists before opening?**
`os.path.exists()`.

**Q1090: How to write a list of strings to a file line-by-line?**
`f.writelines(s + '\n' for s in lst)`.

---

## 🔹 110. Date, Time & Meta (Questions 1091-1100)

**Q1091: How do you get the current time in milliseconds?**
`time.time() * 1000`.

**Q1092: What is strftime() used for?**
Format date to string.

**Q1093: How do you convert a timestamp to human-readable format?**
`datetime.fromtimestamp(ts).ctime()`.

**Q1094: How to get weekday from a date?**
`dt.weekday()` (0=Mon) or `dt.strftime("%A")`.

**Q1095: How to add hours or minutes to current time?**
`dt + timedelta(hours=5)`.

**Q1096: What is the __module__ attribute?**
Name of the module where the class/function is defined.

**Q1097: What is the difference between hasattr() and getattr()?**
*   `has`: Check existence (Bool).
*   `get`: Retrieve value (Object).

**Q1098: What is the purpose of inspect.signature()?**
Get parameter details (names, defaults, annotations) of a callable.

**Q1099: How do you check the number of arguments a function accepts?**
`len(inspect.signature(func).parameters)`.

**Q1100: Can you access docstrings dynamically at runtime?**
Yes. `obj.__doc__`.

---

## 🔹 111. Code Patterns & Control (Questions 1101-1110)

**Q1101: What does pass do in Python?**
(Refresher) Null operation placeholder.

**Q1102: How to emulate a switch-case in Python?**
Dictionary mapping or `match-case` (Py3.10).

**Q1103: What is the enumerate() function useful for?**
(Refresher) Index and value in loops.

**Q1104: How to create a constant in Python?**
(Refresher) Uppercase convention.

**Q1105: How to implement a retry loop?**
Loop with `try-except` and a `break` on success (or max retries).

**Q1106: What is the use of uuid module?**
(Refresher) Unique ID generation.

**Q1107: What does pathlib.Path().exists() check?**
Checks if the path points to an existing file or directory.

**Q1108: How to zip and unzip files using zipfile?**
`zipfile.ZipFile('file.zip', 'w')` to write. `'r'` to read/extract.

**Q1109: What is the purpose of argparse?**
(Refresher) CLI argument parsing.

**Q1110: What does subprocess.run() return?**
A `CompletedProcess` object containing args, returncode, stdout, and stderr.

---

## 🔹 112. Best Practices & Quality (Questions 1111-1120)

**Q1111: How to handle configuration in a Python app?**
Environment variables, `.ini`/`.toml` files, or dedicated classes (`pydantic`).

**Q1112: How to avoid race conditions in file writing?**
File locking (`fcntl` or third-party `portalocker`) or using atomic writes (write to temp, then rename).

**Q1113: What is the importance of testing your Python code?**
Ensures correctness, prevents regressions, and documents behavior.

**Q1114: What is a linter and how does it help in Python?**
(Refresher) Static code analysis.

**Q1115: What’s the importance of writing docstrings?**
Auto-generated documentation and IDE support.

**Q1116: What does Python compile .py files into?**
Bytecode (`.pyc`).

**Q1117: What’s the difference between .pyc and .pyo files?**
*   `.pyc`: Normal bytecode.
*   `.pyo`: Optimized bytecode (removed asserts, docstrings) generated with `-O` flag. (Removed in Py3.5, now `.pyc` handles it).

**Q1118: What is the role of PYTHONPATH?**
(Refresher) Search path for modules.

**Q1119: What is CPython, and how is it different from Jython or PyPy?**
*   **CPython:** Standard C implementation.
*   **Jython:** Runs on Java VM.
*   **PyPy:** JIT compilation (Faster).

**Q1120: What is the Abstract Syntax Tree (AST) in Python?**
Tree representation of the source code structure. Accessible via `ast` module.

---

## 🔹 113. Strings & Memory (Questions 1121-1130)

**Q1121: What happens when you assign one list to another?**
Copies reference (Aliasing). Both point to the same list.

**Q1122: How does Python handle memory allocation for integers?**
Small integers are interned/cached. Large integers use arbitrary precision implementation (arrays of digits).

**Q1123: What is a memory leak in Python?**
Unreferenced objects that GC fails to clean (usually due to circular references or global caches growing indefinitely).

**Q1124: What is garbage collection?**
(Refresher) Automated memory management.

**Q1125: How to check if a string is a valid identifier?**
`s.isidentifier()`.

**Q1126: What is the result of "hello".find("x")?**
`-1`.

**Q1127: How to count overlapping substrings?**
Regex `re.findall('(?=(sub))', s)`. Standard `count()` does not overlap.

**Q1128: What’s the difference between str.title() and str.capitalize()?**
*   `title`: "Hello World" (Every word).
*   `capitalize`: "Hello world" (First char only).

**Q1129: What is Unicode and how is it handled in Python strings?**
Python 3 strings are Unicode by default (UTF-8). `ord()` returns code point.

**Q1130: What is list slicing with negative step?**
Reverses operation. `[start:end:-1]`.

---

## 🔹 114. List & Tuple Logic (Questions 1131-1140)

**Q1131: How to swap two elements in a list?**
`lst[i], lst[j] = lst[j], lst[i]`.

**Q1132: How does list multiplication affect references?**
`[[]] * 3` copies references. Inner lists are shared.

**Q1133: How to filter a list by length of elements?**
`[s for s in lst if len(s) > n]`.

**Q1134: What is the result of list.append(list)?**
Adds the list itself as an element (Recursive reference if appending to self).

**Q1135: How are tuples useful in database queries?**
Parameter passing (safe from SQL injection as arguments) and returning rows (immutable records).

**Q1136: Can you use slicing with tuples like with lists?**
Yes. Returns a new tuple.

**Q1137: What is the performance benefit of using tuples?**
Slightly faster creation/indexing, lower memory overhead.

**Q1138: Can tuples store mutable types?**
Yes (e.g., list inside tuple).

**Q1139: How does tuple hashing work?**
Combines hashes of elements. Fails if any element is unhashable.

**Q1140: What happens when you merge two dictionaries with duplicate keys?**
Later dictionary's values overwrite earlier ones.

---

## 🔹 115. Dict & Sets Deep Dive (Questions 1141-1150)

**Q1141: How to update only missing keys in a dict?**
`d.setdefault(key, val)` or `d = default | d`.

**Q1142: What is dict.fromkeys() used for?**
(Refresher) batch creation with default value.

**Q1143: What’s the output of dict(zip(...)) with unequal lengths?**
Truncates to shortest.

**Q1144: What is OrderedDict and when to use it?**
Preserves insertion order (Built-in dict does this now too, but OrderedDict has equality check sensitive to order).

**Q1145: What is the difference between isdisjoint() and intersection()?**
*   `isdisjoint`: Bool (True if no common).
*   `intersection`: Set (Returns common).

**Q1146: How do you remove duplicates from a string using a set?**
`"".join(set(s))`. (Order lost).

**Q1147: What is the use of frozenset() as a key?**
Allows using a set of items as a dictionary key.

**Q1148: Can you update a set with another iterable?**
`s.update(iterable)`.

**Q1149: How to generate all combinations of set elements?**
`itertools.combinations(s, r)`.

**Q1150: What happens if you return a lambda from a function?**
You get a closure (function object).

---

## 🔹 116. Functional Concepts (Questions 1151-1160)

**Q1151: Can you return a function defined inside another?**
Yes.

**Q1152: How does Python handle function argument unpacking?**
Matches positional args first, then keywords.

**Q1153: How do you document a function properly?**
Docstrings (`"""..."""`) and type hints.

**Q1154: How do you chain multiple lambda expressions?**
`(lambda x: (lambda y: y+x))(1)(2)`. (Confusing, avoid).

**Q1155: What’s the difference between reduce() and a loop?**
`reduce` is functional (folding). Loop is imperative.

**Q1156: How to use map() on nested data?**
Nested map or list comprehension (better).

**Q1157: What are some limitations of filter()?**
Only allows strictly True/False filtering (no transformation).

**Q1158: What’s a real-world use case for reduce()?**
Cumulative product, flattening lists (sometimes), running totals.

**Q1159: How do you use list comprehension with nested loops?**
`[x for inner in outer for x in inner]`.

**Q1160: Can generator expressions be passed to sum()?**
Yes. `sum(x*2 for x in data)` (No extra brackets needed).

---

## 🔹 117. Generator & Loop Nuances (Questions 1161-1170)

**Q1161: What’s the benefit of using generator expressions over list ones?**
Memory efficiency (Lazy evaluation).

**Q1162: Can you use continue in list comprehension?**
No.

**Q1163: How to build a generator that returns infinite numbers?**
`def gen(): i=0; while True: yield i; i+=1`.

**Q1164: What’s the use of else in while loops?**
Runs if condition becomes False (normal exit), not if `break`.

**Q1165: What does for x in reversed(range(3)) output?**
`2, 1, 0`.

**Q1166: How to simulate a do-while loop in Python?**
`while True: ... if condition: break`.

**Q1167: What is a sentinel-controlled loop?**
Loop that runs until a special value (sentinel) is processed.

**Q1168: What’s the difference between while True and while 1?**
Historically `1` was faster. Now (Py3) they are identical (compiler optimizes `True`).

**Q1169: How do you suppress an exception without using try-except?**
`contextlib.suppress(Error)`.

**Q1170: How do you raise an exception manually?**
`raise Error`.

---

## 🔹 118. Advanced Class Concepts (Questions 1171-1180)

**Q1171: What does assert do and when should you use it?**
(Refresher) Dev checks.

**Q1172: What’s a real-world use of finally?**
Closing DB connections or files.

**Q1173: What happens if you return from finally block?**
Swallows exceptions and overrides previous returns.

**Q1174: What’s the difference between instance and class variables?**
(Refresher) Scope (Object vs Class).

**Q1175: How do you override __eq__() and why?**
To allow object comparison (`obj1 == obj2`) based on internal data.

**Q1176: What is method resolution order (MRO)?**
(Refresher) C3 linearization for inheritance.

**Q1177: How do you prevent subclassing?**
Raise error in `__init_subclass__` or use `final` decorator (Py3.8+ typing only).

**Q1178: What does @staticmethod do inside a class?**
(Refresher) No `self` access.

**Q1179: How to write a decorator that repeats a function N times?**
Decorator accepting args (3 levels deep). Loop inside wrapper.

**Q1180: How can you debug decorated functions?**
`func.__wrapped__` allows access to original (if `wraps` used).

---

## 🔹 119. File & IO Tricks (Questions 1181-1190)

**Q1181: How can you chain multiple decorators?**
(Refresher) Stack them.

**Q1182: How to apply decorators conditionally?**
Apply normally, but inside the decorator check a condition and return original `func` if needed.

**Q1183: How to rename a file using Python?**
`os.rename(src, dst)`.

**Q1184: How to read a CSV file without using pandas?**
`csv` module. `csv.reader(f)`.

**Q1185: How to append JSON to an existing file?**
Read list, append, write back. (JSON structure doesn't support simple appending).

**Q1186: How to copy a file using Python?**
`shutil.copy(src, dst)`.

**Q1187: What is os.path.exists() used for?**
(Refresher) Check path.

**Q1188: How to get current time in different timezones?**
`datetime.now(pytz.timezone('US/Eastern'))`.

**Q1189: How to add 1 week to current date?**
`dt + timedelta(weeks=1)`.

**Q1190: How to convert datetime to timestamp?**
`dt.timestamp()`.

---

## 🔹 120. Inspection & Runtime (Questions 1191-1200)

**Q1191: How to parse a date string with a custom format?**
`strptime`.

**Q1192: How to find the last day of the month?**
Calculate 1st day of next month - 1 day.

**Q1193: What does type(obj) return?**
The class object.

**Q1194: What does dir() show?**
(Refresher) Attributes list.

**Q1195: What is __annotations__ in a function?**
Dict of type hints.

**Q1196: How to list all methods of a class?**
Filter `dir(Class)` where attribute is callable.

**Q1197: How do you inspect default values of function arguments?**
`inspect.signature(func).parameters['arg'].default`.

**Q1198: What’s the difference between is and == for strings?**
(Refresher) Identity vs Equality. Short strings may intersect (True/True), long strings usually don't (False/True).

**Q1199: Can two empty lists be == and not is?**
Yes. `[] == []` (True). `[] is []` (False).

**Q1200: Why should you avoid using eval()?**
Security risk (Arbitrary code execution).

---

## 🔹 121. Advanced Typing & Async (Questions 1201-1210)

**Q1201: What is the difference between input() and raw_input()?**
*   `raw_input()`: Python 2 (returns string).
*   `input()`: Python 3 (returns string). In Python 2, `input()` evaluated the input.

**Q1202: What does breakpoint() do in Python 3.7+?**
Drops into the debugger (`pdb` by default). Replaces `import pdb; pdb.set_trace()`.

**Q1203: How do you reload a module in runtime?**
`importlib.reload(module)`.

**Q1204: How to use tempfile.TemporaryDirectory()?**
Creates a temp dir cleaned up on exit.
`with tempfile.TemporaryDirectory() as tmpdir: ...`.

**Q1205: What is shutil.move() used for?**
Moves a file or directory.

**Q1206: How to list installed packages programmatically?**
`pkg_resources` (legacy) or `importlib.metadata.distributions()` (modern).

**Q1207: Why avoid from module import *?**
Pollutes namespace, risks variable shadowing, makes code hard to read (where did `func` come from?).

**Q1208: How to ensure your code is forward compatible with Python 4?**
Follow deprecation warnings, use `__future__` imports, avoid removed features.

**Q1209: What is the difference between unit test and integration test?**
*   Unit: Tests smallest part (function) in isolation.
*   Integration: Tests how parts work together (API + DB).

**Q1210: How to organize code in packages/modules?**
Logical structure (Domain-driven), `src` layout, `__init__.py` for exports.

---

## 🔹 122. Bytecode & Execution (Questions 1211-1220)

**Q1211: ✅ Start detailed answers per batch?**
(Meta-question: Skipping).

**Q1212: 🧠 Practice with mock MCQs per topic?**
(Meta-question: Skipping).

**Q1213: 📦 Export all batches into one structured Excel/PDF/Notion table?**
(Meta-question: Skipping).

**Q1214: What happens when a Python script is run?**
Compilation to Bytecode -> PVM Execution.

**Q1215: How does Python convert source code to bytecode?**
The Parser generates an AST, which the Compiler transforms into bytecode instructions (.pyc).

**Q1216: What is the significance of indentation levels in Python blocks?**
Determines scope nesting.

**Q1217: How does Python handle operator precedence?**
PEMDAS + logic rules. `not` > `and` > `or`.

**Q1218: What are compound assignment operators in Python?**
`+=`, `-=`, `*=`, `/=`. Update variable in-place (if mutable) or rebind (if immutable).

**Q1219: What is a namespace in Python?**
Mapping of names to objects (Dict-based).

**Q1220: What is a symbol table?**
Internal table created by compiler to track variable scope (local vs global) before execution.

---

## 🔹 123. Garbage Collection & Strings (Questions 1221-1230)

**Q1221: What is dynamic typing and how is it implemented?**
Variables are references to objects. The object carries the type info, not the variable.

**Q1222: What’s the effect of del on an object?**
Decrements reference count. If count hits 0, object is destroyed.

**Q1223: How can you manually trigger garbage collection?**
`gc.collect()`.

**Q1224: How does Python internally represent Unicode strings?**
Flexible String Representation (PEP 393). Dependent on max char width: ASCII (1 byte), UCS-2 (2 bytes), or UCS-4 (4 bytes) per char.

**Q1225: How to efficiently check if a string is a palindrome?**
`s == s[::-1]`.

**Q1226: How to extract only numeric characters from a string?**
`"".join(filter(str.isdigit, s))`.

**Q1227: How do escape sequences work in Python strings?**
Backslash `\` interpreted by parser. `\n` -> Line feed byte.

**Q1228: How to center-align text using Python?**
`s.center(width, fillchar)`.

**Q1229: What is the output of list("abc")?**
`['a', 'b', 'c']`.

**Q1230: What happens if you extend a list with itself?**
It doubles in length with repeated elements.

---

## 🔹 124. Tuples & Dictionaries (Questions 1231-1240)

**Q1231: How to flatten a deeply nested list?**
Recursion or `deepflatten` from libraries like `iteration_utilities`.

**Q1232: How to find duplicates in a list?**
`seen=set(); dups=[x for x in lst if x in seen or seen.add(x)]`.

**Q1233: Can you convert a generator to a tuple?**
Yes. `tuple(gen)`.

**Q1234: What’s the difference between (1) and (1,)?**
`(1)` is int `1`. `(1,)` is a `tuple`.

**Q1235: How can a tuple be used in recursion?**
As an immutable state carrier passed down call stack.

**Q1236: How to index into a tuple of tuples?**
`t[i][j]`.

**Q1237: What is tuple immutability and how does it affect performance?**
Allows caching, optimization, and safe sharing across threads.

**Q1238: How to remove all dictionary items conditionally?**
`{k:v for k,v in d.items() if check(v)}`.

**Q1239: What’s the role of dict.update()?**
Merges another dict (or iterable pairs) into current dict. Overwrites existing keys.

**Q1240: How to iterate through nested dictionaries?**
Recursion.

---

## 🔹 125. Sets & Functions II (Questions 1241-1250)

**Q1241: What happens when you use + on two dictionaries?**
`TypeError` (before Py3.9). Use `|` in 3.9+.

**Q1242: Can dictionary values be functions?**
Yes. Useful for dispatch tables (replacing switch-case).

**Q1243: How to filter values that are not in a set?**
`[x for x in lst if x not in s]`.

**Q1244: What happens when sets are created with duplicate elements?**
Duplicates are discarded immediately.

**Q1245: Can sets be directly serialized to JSON?**
No. JSON has no set type. Must convert to list first.

**Q1246: What does set.symmetric_difference() return?**
Elements in A or B, but not both.

**Q1247: How does Python internally store sets?**
Hash table with dummy values.

**Q1248: Can you define functions inside functions?**
Yes (Nested functions).

**Q1249: What’s the use of globals() in functions?**
Access/Modify global symbol table dynamically.

**Q1250: Can default arguments be overridden in a decorator?**
Yes, by modifying `args`/`kwargs` in the wrapper.

---

## 🔹 126. Lambdas & Comprehensions II (Questions 1251-1260)

**Q1251: How to make functions pure and testable?**
Avoid side effects (globals, I/O) and ensure output depends only on input.

**Q1252: How do Python functions simulate optional parameters?**
Default arguments `def f(a=None)`.

**Q1253: Can lambdas use conditional expressions?**
Yes. `x if c else y`.

**Q1254: How to use lambda in filtering JSON-like data?**
`filter(lambda item: item['id'] == 5, data)`.

**Q1255: Can you assign a lambda inside a class?**
Yes, acts like a method.

**Q1256: How to use lambda inside reduce()?**
`reduce(lambda acc, x: acc+x, lst)`.

**Q1257: Why would you wrap a lambda in a regular function?**
To give it a name and docstring (Best practice: just use `def`).

**Q1258: Can you use a function call inside list comprehension?**
Yes. `[func(x) for x in list]`.

**Q1259: How to reverse the logic of a dictionary comprehension?**
`{v:k for k,v in d.items()}`.

**Q1260: How do you flatten a matrix using list comprehension?**
(Refresher) `[col for row in mat for col in row]`.

---

## 🔹 127. Loops & Exceptions III (Questions 1261-1270)

**Q1261: How to create a comprehension with two filters?**
`[x for x in list if cond1 if cond2]`.

**Q1262: What’s the performance benefit of comprehension?**
Faster than `for` loop logic because iteration happens at C level.

**Q1263: How do you use enumerate() for reverse indexing?**
`reversed(list(enumerate(lst)))` or calculating index manually.

**Q1264: How to loop over multiple ranges simultaneously?**
`zip(range(n), range(m))`.

**Q1265: What does zip(*matrix) achieve in a loop?**
Transposes the matrix (loops over columns).

**Q1266: How does unpacking work in loop variables?**
`for a, b in list_of_pairs:`.

**Q1267: How to detect infinite loop conditions?**
Static analysis (hard) or timeout guards / iteration limits.

**Q1268: What’s the use of raise from syntax?**
(Refresher) Exception chaining.

**Q1269: What does try/except/finally look like when all used together?**
`try: ... except: ... finally: ...`.

**Q1270: How do you suppress expected exceptions during logging?**
`try: ... except Error: pass` or `logging.debug("Ignored error", exc_info=True)`.

---

## 🔹 128. Classes & OOP III (Questions 1271-1280)

**Q1271: What’s the difference between catching Exception and BaseException?**
(Refresher) `BaseException` catches system exit signals too. `Exception` is safe for logic.

**Q1272: Can you re-raise the last caught exception?**
`raise` (bare).

**Q1273: Can you override __init__() in a subclass?**
Yes. Usually call `super().__init__()` too.

**Q1274: How to restrict instance creation in Python?**
Override `__new__` (Singleton) or raise Error in `__init__`.

**Q1275: What is the role of __del__() destructor method?**
Cleanup when object is destroyed.

**Q1276: What does __getattr__() do?**
Called when attribute lookup **fails**. (Fallback).

**Q1277: How to implement duck typing in a class?**
Just implement the required methods (e.g., `__len__`, `__getitem__`).

**Q1278: What is a decorator factory?**
A function that takes arguments and returns a decorator.

**Q1279: How to write a decorator that times a function?**
(Refresher) Measure start/end time around func call.

**Q1280: How to make a decorator that retries on failure?**
Loop inside wrapper `try: func(); return; except: continue`.

---

## 🔹 129. Decorators & File OS III (Questions 1281-1290)

**Q1281: Can a decorator prevent function execution?**
Yes, if wrapper doesn't call `func()`.

**Q1282: Can you apply decorators to properties?**
Yes. `@property` is itself a decorator.

**Q1283: What does os.walk() do?**
Generates file names in a directory tree (recursive travel). Yields `(dirpath, dirnames, filenames)`.

**Q1284: How to write data to CSV without using pandas?**
`csv.writer(f).writerows(data)`.

**Q1285: What’s the use of os.chmod()?**
Change file permissions (Read/Write/Exec).

**Q1286: How to append binary content to a file?**
`open('file', 'ab')`.

**Q1287: How to read file in reverse line order?**
Read all lines and `reversed()`, or seek from end (complex).

**Q1288: How to convert naive datetime to aware datetime?**
`dt.replace(tzinfo=...)` or `timezone.localize(dt)`.

**Q1289: How to get date N days from today?**
(Refresher) `timedelta`.

**Q1290: How to calculate difference in months between two dates?**
`(d2.year - d1.year)*12 + d2.month - d1.month`.

---

## 🔹 130. Meta & Introspection (Questions 1291-1300)

**Q1291: How to round datetime to nearest hour?**
Round minutes/seconds. If >= 30, add hour.

**Q1292: What is timedelta.total_seconds() used for?**
Get total duration in seconds (float), ignoring day divisions.

**Q1293: What does inspect.getmembers() return?**
List of `(name, value)` pairs for all members of an object.

**Q1294: What is the use of __slots__ and how does it affect attributes?**
Prevents creation of new attributes not listed in slots. Faster access.

**Q1295: How to dynamically modify function behavior using reflection?**
Monkey patching or decorators.

**Q1296: What’s the difference between __class__ and type()?**
`obj.__class__` is the attribute pointing to the class. `type(obj)` returns it.

**Q1297: What’s the use of callable()?**
Checks if object can be called `()`.

**Q1298: What is a contextlib context manager?**
Helper `@contextmanager` to create context managers using generators (yield).

**Q1299: How to implement a simple command-line interface (CLI)?**
`if len(sys.argv) > 1: ...` or `argparse`.

**Q1300: What is a generator vs coroutine?**
*   **Generator:** Produces data (`yield`).
*   **Coroutine:** Consumes data (yield used as expression `x = yield`). Async/Await are specialized coroutines.

---

## 🔹 131. Text & Strings IV (Questions 1301-1310)

**Q1301: How to use textwrap for formatting paragraphs?**
`textwrap.fill(text, width=50)`.

**Q1302: What does uuid.uuid4() generate?**
A random UUID (Version 4). `uuid1` is time/host-based.

**Q1303: How to get system information using platform module?**
`platform.system()`, `platform.release()`.

**Q1304: What’s the role of json.loads() vs json.load()?**
*   `loads`: Load string (S).
*   `load`: Load file.

**Q1305: How does argparse support optional flags?**
Prefix with `-` or `--`. `parser.add_argument('--verbose')`.

**Q1306: Why avoid mutable default values?**
They are shared across all calls. if you append to a default list, it grows forever.

**Q1307: How to handle bad user input gracefully?**
Try-except blocks around conversion logic (`int(input())`).

**Q1308: What is docstring and how should it be written?**
Triple-quoted string at start of block. Description -> Args -> Returns.

**Q1309: Why is logging preferred over print() in production?**
Configurable levels (INFO/ERROR), destinations (File/Stream), and formatting.

**Q1310: What is the benefit of writing unit tests with pytest or unittest?**
Automated verification, refactoring safety, code documentation.

---

## 🔹 132. Advanced Caching & Compilation (Questions 1311-1320)

**Q1311: 🎓 Get answer explanations + real examples?**
(Meta-question: Skipping).

**Q1312: 🧠 Practice flashcards or MCQs?**
(Meta-question: Skipping).

**Q1313: What is the .pyc file and when is it generated?**
(Refresher) Bytecode cache. Generated on import.

**Q1314: How to disable bytecode generation in Python?**
`PYTHONDONTWRITEBYTECODE=1` env var or `python -B`.

**Q1315: What is the role of the __pycache__ folder?**
Stores versioned `.pyc` files (e.g., `script.cpython-39.pyc`).

**Q1316: What happens during the import of a module?**
Find -> Compile (if needed) -> Run -> Cache in `sys.modules`.

**Q1317: What is the difference between interpreter and compiler in Python’s context?**
Python compiles to bytecode (Compiler step), then interprets bytecode (Interpreter step).

**Q1318: How does locals() behave inside a function?**
Returns a dictionary representing the current local symbol table. Updates to it may not affect actual variables.

**Q1319: What is the vars() function used for?**
Returns `__dict__` attribute of an object. If no arg, behaves like `locals()`.

**Q1320: Can two variables reference the same object in memory?**
Yes. `a = b = []`.

---

## 🔹 133. Identity & Strings Deep Dive (Questions 1321-1330)

**Q1321: What is object identity and how is it checked?**
The memory address. Checked with `id(obj)` or `is`.

**Q1322: What does isinstance(obj, (list, tuple)) do?**
Checks if obj is *either* a list *or* a tuple.

**Q1323: What does "hello".startswith("h") return?**
`True`.

**Q1324: How to convert camelCase to snake_case in Python?**
Regex `re.sub(r'(?<!^)(?=[A-Z])', '_', name).lower()`.

**Q1325: What’s the difference between .split() and .rsplit()?**
Same unless `maxsplit` is defined. `rsplit` starts from right.

**Q1326: What does str.casefold() do?**
(Refresher) Aggressive lowercase.

**Q1327: How do you count words in a sentence using split?**
`len(text.split())`.

**Q1328: How to filter all None values from a list?**
`list(filter(None, lst))` or `[x for x in lst if x is not None]`.

**Q1329: What is list unpacking and how is it used?**
`head, *tail = lst`.

**Q1330: How do you slice from the end of the list?**
`lst[-n:]`.

---

## 🔹 134. List & Tuple Tricks (Questions 1331-1340)

**Q1331: Can lists hold functions as elements?**
Yes. `[print, len][0]("Hi")`.

**Q1332: What happens when a list is sorted in-place?**
It returns `None` and modifies the original list.

**Q1333: How do tuples support destructuring assignments?**
`x, y = (1, 2)`.

**Q1334: Can a tuple contain itself?**
Yes. `t = ([],); t[0].append(t)`. (Not directly possible with immutable tuple content, but possible if tuple contains mutable container).

**Q1335: What is the result of multiplying a tuple by 0?**
`()`.

**Q1336: What is the role of tuple nesting in pattern matching?**
Allows matching complex structures. `case (x, (y, z)):`.

**Q1337: How do you convert tuple of keys and values into a dict?**
`dict(tuple_of_pairs)`.

**Q1338: What does dict.setdefault() do?**
(Refresher) Get value or set default.

**Q1339: How do you check if a value (not key) exists in a dict?**
`val in d.values()`. (O(n) time).

**Q1340: How does dict.clear() behave?**
Empties content in-place. Other references see empty dict.

---

## 🔹 135. Set & Dict Mechanics (Questions 1341-1350)

**Q1341: Can dict values be other dictionaries?**
Yes (Nested dicts).

**Q1342: What is the result of len({})?**
`0`.

**Q1343: What does set.update() do?**
Adds elements from an iterable to the set.

**Q1344: Can sets contain other sets?**
No. Sets are mutable and unhashable. Use `frozenset`.

**Q1345: What is the hashability requirement for set elements?**
Elements must be immutable (implement `__hash__`).

**Q1346: How does set subtraction work?**
`s1 - s2` (Items in s1 not in s2).

**Q1347: What does __doc__ attribute store?**
Docstring.

**Q1348: What’s the use of function.__name__?**
The name of the function.

**Q1349: Can functions be renamed dynamically?**
Yes. `func.__name__ = "new_name"`. (Debug tools use this).

**Q1350: How to bind function parameters with functools.partial()?**
`new_f = partial(f, arg1=val)`.

---

## 🔹 136. Lambda & Functional III (Questions 1351-1360)

**Q1351: What is a keyword-only argument?**
(Refresher) Arg that must be named. `def f(*, a):`.

**Q1352: What does lambda x=10: x+5 mean?**
Lambda with default argument. Returns 15 if called as `f()`.

**Q1353: Can lambdas have default values?**
Yes.

**Q1354: How to use lambda to sort dicts by value?**
`sorted(d.items(), key=lambda i: i[1])`.

**Q1355: What is the return type of map() with lambda?**
Map object (iterator).

**Q1356: Can a lambda return another lambda?**
Yes.

**Q1357: Can list comprehensions use ternary operators?**
(Refresher) Yes.

**Q1358: How do nested comprehensions work for matrices?**
Outer loop first.

**Q1359: Can comprehensions include function calls?**
Yes.

**Q1360: Can you use enumerate() in a comprehension?**
Yes. `[i for i, x in enumerate(lst)]`.

---

## 🔹 137. Loop & Logic (Questions 1361-1370)

**Q1361: What happens if you use a comprehension to build a dict with duplicate keys?**
Last one wins.

**Q1362: Can for loop work on generators?**
Yes. Generators are iterables.

**Q1363: How to skip every 2nd element in a loop?**
`range(0, len(lst), 2)`.

**Q1364: What is the result of break vs return in loops?**
*   `break`: Exits loop, continues function.
*   `return`: Exits function entirely.

**Q1365: How to run two loops in parallel?**
`zip()`.

**Q1366: What’s the output of for x in []:?**
Loop body never runs.

**Q1367: What is an exception hierarchy in Python?**
(Refresher) Tree of exception classes.

**Q1368: What happens if you don’t catch a raised exception?**
Crash / Termination.

**Q1369: How does Python traceback work?**
Stack is unwound, printing file/line info for each frame.

**Q1370: What is the use of raise AssertionError?**
Usually triggered by `assert False`.

---

## 🔹 138. Exception & Class Attributes (Questions 1371-1380)

**Q1371: How do exceptions behave with context managers?**
`__exit__` receives the exception types/values. Can suppress or propagate them.

**Q1372: Can two instances of the same class have different attributes?**
Yes. Instance attributes are dynamic. `obj1.x = 1` doesn't affect `obj2`.

**Q1373: What is dynamic attribute assignment?**
Adding fields at runtime. `obj.new_field = value`.

**Q1374: Can methods be added to classes dynamically?**
Yes. `Class.method = func`.

**Q1375: What is method shadowing?**
Subclass method overrides parent method found in MRO.

**Q1376: What’s the difference between instance and class method in behavior?**
Instance receives `self` (access state). Class receives `cls`.

**Q1377: What is a common use of a class method decorator?**
Alternative constructors.

**Q1378: How do decorators affect function signatures?**
They can hide them. `functools.wraps` helps fix this.

**Q1379: What’s the benefit of using functools.wraps()?**
(Refresher) Metadata preservation.

**Q1380: Can decorators have return values?**
The wrapper function returns the value.

---

## 🔹 139. Metaclasses & I/O IV (Questions 1381-1390)

**Q1381: What is a metaclass-level decorator?**
Not common terminology. Usually refers to decorating the metaclass methods.

**Q1382: What’s the effect of opening a file in 'w+' mode?**
Read/Write, but truncates (clears) file first.

**Q1383: How to write UTF-8 encoded text?**
`open(f, 'w', encoding='utf-8')`.

**Q1384: How to read all lines of a file into a list?**
`f.readlines()`.

**Q1385: What is the file cursor and how is it used?**
Internal pointer to current read/write position.

**Q1386: What is the effect of calling .flush()?**
Forces write buffer content to disk immediately.

**Q1387: How to calculate age from a birthdate?**
`(today - birthdate).days // 365`.

**Q1388: How to find the first Monday of the month?**
Iterate days 1-7 using `datetime`, check `weekday() == 0`.

**Q1389: How to convert date to Unix timestamp?**
(Refresher) `.timestamp()`.

**Q1390: How to parse ISO 8601 dates?**
`datetime.fromisoformat()`.

---

## 🔹 140. Reflection & Context (Questions 1391-1400)

**Q1391: What is the difference between datetime and date objects?**
*   `date`: Y-M-D.
*   `datetime`: Y-M-D H:M:S.

**Q1392: How to check if an object has a method?**
`hasattr(obj, 'method')` and `callable(getattr(obj, 'method'))`.

**Q1393: How do you get the source code of a function?**
`inspect.getsource(func)`.

**Q1394: What is dynamic dispatch?**
Runtime decision of which method to call (Polymorphism).

**Q1395: What’s the purpose of globals() and locals()?**
(Refresher) Scope inspection.

**Q1396: How do you remove leading/trailing characters from a string?**
`strip('chars')`.

**Q1397: How do you convert a list to a CSV line?**
`",".join(list)`.

**Q1398: What’s the use of zip_longest()?**
`itertools.zip_longest`: Zips to longest iter, filling missing values with `None`/fillvalue.

**Q1399: How to filter out falsy values from a list?**
`list(filter(None, lst))`.

**Q1400: How to merge two JSON strings?**
Parse `loads` to dicts, merge dicts, `dumps` back.

---

## 🔹 141. Types & Inspection (Questions 1401-1410)

**Q1401: What does type() return for old-style classes?**
(Python 2 only) Returns `<type 'instance'>`. In Python 3, all classes are new-style.

**Q1402: How to force garbage collection?**
`gc.collect()`.

**Q1403: What does the weakref module do?**
Creates weak references to objects, allowing them to be garbage collected.

**Q1404: How to check if a class is a subclass of another?**
`issubclass(Child, Parent)`.

**Q1405: What is the method resolution order (MRO) for multiple inheritance?**
C3 Linearization. Left-to-right, depth-first (but preserving order).

**Q1406: How to create a dynamic class at runtime?**
`type('ClassName', (Parent,), {'attr': val})`.

**Q1407: What is the descriptor protocol?**
Objects with `__get__`, `__set__`, or `__delete__` methods.

**Q1408: What’s the difference between __getattribute__ and __getattr__?**
*   `__getattribute__`: Called for *all* attribute access.
*   `__getattr__`: Called only if attribute lookup fails.

**Q1409: How to list all imported modules?**
`sys.modules.keys()`.

**Q1410: What is __all__ in a module?**
List of strings defining what `from module import *` imports.

---

## 🔹 142. Performance & Memory (Questions 1411-1420)

**Q1411: How to measure memory usage of an object?**
`sys.getsizeof(obj)`.

**Q1412: What is interning?**
(Refresher) Reusing objects for optimization.

**Q1413: How does Python manage memory for small integers?**
Pre-allocates range -5 to 256.

**Q1414: What is the purpose of gc.disable()?**
Pauses automatic garbage collection (improves performance in critical sections/bulk allocations).

**Q1415: How to profile a Python script?**
`cProfile.run('main()')` or `python -m cProfile script.py`.

**Q1416: What is a memory view?**
`memoryview(obj)`. Safe handling of buffer protocol (bytes) without copying.

**Q1417: How to find reference count of an object?**
`sys.getrefcount(obj)`.

**Q1418: What are slots?**
(Refresher) `__dict__` optimization.

**Q1419: How to optimize loop performance?**
Move calculations out, use local variables, use built-ins (`map`/`filter`), use compiled libs (`numpy`).

**Q1420: What is the benefit of using local variables over global?**
Faster access (STORE_FAST vs STORE_GLOBAL opcodes).

---

## 🔹 143. Functional & Iterators (Questions 1421-1430)

**Q1421: How to make an object iterable?**
Implement `__iter__`.

**Q1422: How to make an object an iterator?**
Implement `__iter__` and `__next__`.

**Q1423: What is the difference between an iterable and a sequence?**
*   Iterable: Can be looped over.
*   Sequence: Iterable + Indexing + Length (e.g., list, tuple, str).

**Q1424: How does iter() work with a sentinel?**
`iter(func, sentinel)`: Calls func until it returns sentinel.

**Q1425: What is the StopIteration exception?**
Signal that an iterator is exhausted.

**Q1426: How to chain multiple iterators?**
`itertools.chain(i1, i2)`.

**Q1427: How to slice an iterator/generator?**
`itertools.islice(gen, start, stop)`.

**Q1428: What is itertools.takewhile()?**
Returns elements as long as predicate is true.

**Q1429: What is itertools.dropwhile()?**
Drops elements as long as predicate is true, then returns the rest.

**Q1430: How to generate Cartesian product of lists?**
`itertools.product(l1, l2)`.

---

## 🔹 144. Async & Concurrency (Questions 1431-1440)

**Q1431: What is the asyncio module?**
Library for writing concurrent code using async/await syntax (Single-threaded).

**Q1432: What is an event loop?**
Central executor that runs async tasks and callbacks.

**Q1433: How to define a coroutine?**
`async def my_func():`.

**Q1434: How to run a coroutine?**
`asyncio.run(my_func())` or await it inside another coroutine.

**Q1435: What is await?**
Pauses the coroutine until the awaitable (Future/Task) completes.

**Q1436: How to run multiple coroutines concurrently?**
`asyncio.gather(coro1(), coro2())`.

**Q1437: What is the difference between threading vs asyncio?**
*   Threading: Preemptive multitasking (OS managed), good for blocking IO.
*   Asyncio: Cooperative multitasking (App managed), specialized for network IO.

**Q1438: Can you mix synchronous and asynchronous code?**
Yes, but blocking calls in async functions freeze the loop. Use `run_in_executor`.

**Q1439: What is a Future?**
An object representing a result that hasn't happened yet.

**Q1440: How to create a task in asyncio?**
`asyncio.create_task(coro())`.

---

## 🔹 145. Packaging & Distribution (Questions 1441-1450)

**Q1441: What is setuptools?**
Library to facilitate packaging Python projects.

**Q1442: What is MANIFEST.in?**
File listing extra files (non-code) to include in the source distribution.

**Q1443: What is a .whl (Wheel) file?**
Built distribution format (ZIP archive) ready to unpack.

**Q1444: What is twine used for?**
Securely uploading packages to PyPI.

**Q1445: How to specify dependency versions in setup.py?**
`install_requires=['pkg>=1.0']`.

**Q1446: What is a virtualenv?**
(Refresher) Isolated environment.

**Q1447: What does pip install -e . do?**
Installs package in "editable" mode (symlink to source).

**Q1448: What is semantic versioning?**
Format `Major.Minor.Patch`.

**Q1449: How to distribute a Python app as an executable?**
`PyInstaller` or `cx_Freeze`.

**Q1450: What is pyproject.toml?**
(Refresher) Build configuration file (PEP 518).

---

## 🔹 146. Python 2 vs 3 (Questions 1451-1460)

**Q1451: What happened to print statement in Python 3?**
Became a function `print()`.

**Q1452: How did division change in Python 3?**
`/` returns float. `//` behaves like old `/` (floor).

**Q1453: What happened to xrange()?**
Renamed to `range()` (old `range` removed).

**Q1454: Are strings unicode by default in Python 3?**
Yes.

**Q1455: What happened to long integer type?**
Unified with `int`.

**Q1456: How do you port Python 2 code to 3?**
`2to3` tool or `modernize`.

**Q1457: What is __future__ module?**
Allows using new features in older versions.

**Q1458: What is the input() change?**
`raw_input()` -> `input()`. `input()` (eval) removed.

**Q1459: Can you mix tabs and spaces in Python 3?**
No. It raises `TabError`.

**Q1460: What happened to file() built-in?**
Removed. Use `open()`.

---

## 🔹 147. Testing (Questions 1461-1470)

**Q1461: What is assertions in testing?**
Verifying that a result matches expectation.

**Q1462: How to organize tests?**
`tests/` folder matching source structure.

**Q1463: What needs to be mocked in tests?**
External dependencies (APIs, DBs, Files, Time).

**Q1464: How to check test coverage?**
`coverage.py` tool.

**Q1465: What is tox?**
Generic virtualenv management and test command line tool (Test against multiple python versions).

**Q1466: What is TDD?**
Test Driven Development (Write test -> Fail -> Write code -> Pass).

**Q1467: How to skip a test in unittest?**
`@unittest.skip("reason")`.

**Q1468: What is setup and teardown?**
Pre-test initialization and post-test cleanup.

**Q1469: What is parameterization in tests?**
Running same test logic with different inputs. `@pytest.mark.parametrize`.

**Q1470: What is a fixture in pytest?**
Reusable setup logic injected as arguments.

---

## 🔹 148. Web & Frameworks II (Questions 1471-1480)

**Q1471: What is Flask’s app context?**
Keeps track of application-level data.

**Q1472: What is request context in Flask?**
Keeps track of request-level data (`request`, `session`).

**Q1473: How does Django handle request lifecycle?**
Middleware -> URL Conf -> View -> Response -> Middleware.

**Q1474: What is Django Admin?**
Automatic admin interface for models.

**Q1475: What is WSGI middleware?**
Wraps application to process requests/responses (e.g., Logging, Auth).

**Q1476: What is g object in Flask?**
Global namespace for holding data during a single request context.

**Q1477: How to handle file uploads in Flask?**
`request.files['file']`.

**Q1478: What is Django Q object?**
Used for complex database lookups (OR/AND conditions).

**Q1479: What is Django F object?**
Represents the value of a model field (for DB-side operations).

**Q1480: How to prevent SQL injection in Python?**
Use parameterized queries (offered by DB-API drivers and ORMs).

---

## 🔹 149. Data Science Basics (Questions 1481-1490)

**Q1481: What is a NumPy array?**
Values of same type, indexed by tuple of non-negative integers.

**Q1482: Difference between List and Array?**
Array: Compact, homogeneous, fast math. List: Flexible, heterogeneous.

**Q1483: What is broadcasting?**
(Refresher) Arithmetic on arrays of different shapes.

**Q1484: How to select column in Pandas?**
`df['col']`.

**Q1485: How to handle NaN in Pandas?**
`dropna()`, `fillna()`.

**Q1486: What is apply() in Pandas?**
Apply function along axis of DataFrame.

**Q1487: What is matplotlib.pyplot?**
State-based interface to matplotlib (similar to MATLAB).

**Q1488: What is scikit-learn fit/predict?**
`fit`: Train model. `predict`: Use model.

**Q1489: What is Jupyter Notebook?**
Web application for live code, equations, visualizations.

**Q1490: How to read Excel file in Pandas?**
`pd.read_excel()`.

---

## 🔹 150. Miscellaneous (Questions 1491-1500)

**Q1491: How to play sound in Python?**
`playsound`, `pygame`, `winsound`.

**Q1492: How to send email with Python?**
`smtplib` module.

**Q1493: How to scrape a website?**
`BeautifulSoup`, `Scrapy`, `requests`.

**Q1494: How to parse XML?**
`xml.etree.ElementTree`.

**Q1495: How to automate mouse/keyboard?**
`pyautogui`.

**Q1496: How to create a GUI?**
`Tkinter`, `PyQt`, `Kivy`.

**Q1497: How to generate QR code?**
`qrcode` library.

**Q1498: How to work with images?**
`Pillow` (PIL fork).

**Q1499: How to connect to SQLite?**
`sqlite3` module (built-in).

**Q1500: Can Python interface with C/C++?**
Yes. `ctypes`, `swig`, `cffi`, Python C-API.

---

## 🔹 151. Threading & Multiprocessing (Questions 1501-1510)

**Q1501: What is a Daemon Thread?**
A background thread that automatically exits when the main program exits (e.g., garbage collection, heartbeat).

**Q1502: How to create a thread in Python?**
`threading.Thread(target=func).start()`.

**Q1503: What does join() do in threading?**
Waits for the thread to complete execution before proceeding.

**Q1504: What is a Lock (Mutex) in Python?**
Primitive to prevent race conditions. `lock.acquire()` / `lock.release()`.

**Q1505: How to use a ThreadPoolExecutor?**
`with concurrent.futures.ThreadPoolExecutor() as ex: ex.submit(func, arg)`.

**Q1506: What is a race condition?**
When multiple threads access shared data concurrently, leading to unpredictable results.

**Q1507: What is a Deadlock?**
When two threads wait for each other to release resources, freezing forever.

**Q1508: How does multiprocessing differ from threading?**
(Refresher) Processes have separate memory (Bypasses GIL). Threads share memory.

**Q1509: How to share data between processes?**
`multiprocessing.Queue`, `Pipe`, or `Value`/`Array` (Shared Memory).

**Q1510: What is the purpose of if __name__ == '__main__' in multiprocessing?**
Required on Windows to prevent recursive spawning of subprocesses.

---

## 🔹 152. Advanced Libraries (Questions 1511-1520)

**Q1511: What is NumPy mainly used for?**
Numerical computing (High-performance arrays/matrices).

**Q1512: What is Pandas Series vs DataFrame?**
(Refresher) 1D vs 2D.

**Q1513: What does matplotlib.pyplot.plot() do?**
Draws lines/markers.

**Q1514: What is SQLAlchemy?**
SQL Toolkit and ORM for Python.

**Q1515: What is Requests library used for?**
Making HTTP requests (`get`, `post`) simply.

**Q1516: What is Beautiful Soup?**
HTML/XML parser for web scraping.

**Q1517: What does PyTest offer over Unittest?**
Simpler syntax (no classes needed), powerful fixtures, plugins.

**Q1518: What is Celery used for?**
Distributed task queue (Async jobs).

**Q1519: What is Pillow (PIL)?**
Image processing library (Resize, Crop, Filter).

**Q1520: What is Scrapy?**
Framework for extracting data from websites (Large scaler scraping).

---

## 🔹 153. Quality & Security (Questions 1521-1530)

**Q1521: What is PEP 8?**
(Refresher) Style guide.

**Q1522: What is pylint?**
(Refresher) Linter.

**Q1523: What is type checking in Python?**
Static analysis (`mypy`) using type hints.

**Q1524: How to secure passwords in Python?**
Hash them (bcrypt/argon2). Never store plain text.

**Q1525: What is SQL Injection prevention?**
Use parameterized queries (`?` or `%s`).

**Q1526: What represents "False" in Python?**
`False`, `None`, `0`, `""`, `[]`, `{}`, `set()`.

**Q1527: How to check complexity of code?**
Cyclomatic complexity tools (`radon`, `xenon`).

**Q1528: What is code coverage?**
Percentage of code lines executed during tests.

**Q1529: What is Monkey Patching in logic terms?**
Changing behavior at runtime (often for testing/mocking).

**Q1530: How to prevent importing a module?**
`sys.modules['mod'] = None`.

---

## 🔹 154. Python 3 Improvements (Questions 1531-1540)

**Q1531: What is the walrus operator?**
(Refresher) `:=`.

**Q1532: What are f-strings?**
(Refresher) `f"{var}"`.

**Q1533: What is asyncio?**
(Refresher) Concurrency lib.

**Q1534: What are data classes?**
(Refresher) `@dataclass`.

**Q1535: What is positional-only argument syntax?**
(Refresher) `/`.

**Q1536: What sort algorithm does Python use?**
Timsort (Hybrid Merge/Insertion sort). O(n log n).

**Q1537: What is the new dict merge operator?**
(Refresher) `|`.

**Q1538: What pattern matching did Python 3.10 add?**
`match case` (Structural pattern matching).

**Q1539: What is type union operator?**
`int | str` (Py3.10) instead of `Union[int, str]`.

**Q1540: What happened to collections.Mapping?**
Moved to `collections.abc.Mapping`.

---

## 🔹 155. Functional & Logic IV (Questions 1541-1550)

**Q1541: Can you pass a class as an argument?**
Yes. Classes are objects. Use for Dependency Injection / Factory.

**Q1542: What is memoization with decorators?**
(Refresher) Caching results.

**Q1543: What is tail recursion?**
Recursion where the last action is the call (Python does NOT optimize this to prevent stack overflow).

**Q1544: How to check if a number is prime?**
Loop 2 to sqrt(n).

**Q1545: How to find factorial recursively?**
`fact(n) = n * fact(n-1)`.

**Q1546: How to reverse a number mathematically?**
Modulo 10 loop. `rev = rev*10 + digit`.

**Q1547: How to find Fibonacci series?**
`a, b = 0, 1; while... a, b = b, a+b`.

**Q1548: How to count set bits in an integer?**
`bin(n).count('1')`.

**Q1549: How to implement a Stack?**
(Refresher) List `append`/`pop`.

**Q1550: How to implement a Queue?**
(Refresher) `deque` `append`/`popleft`.

---

## 🔹 156. Web & Network (Questions 1551-1560)

**Q1551: What is socket programming?**
Low-level networking. `socket.socket()`.

**Q1552: How to make a TCP client?**
Connect to (IP, Port). Send/Recv bytes.

**Q1553: How to make a simple HTTP server?**
`python -m http.server`.

**Q1554: What is JSON Web Token (JWT)?**
Stateless authentication token.

**Q1555: What is basic auth?**
Username/Password encoded in HTTP Header.

**Q1556: What is a Cookie?**
Small data stored in browser by server.

**Q1557: What is REST endpoint?**
URL where API service can be accessed.

**Q1558: What is GraphQL?**
Query language for APIs (single endpoint, flexible queries).

**Q1559: What is WebScraping legality?**
Check `robots.txt` and Terms of Service.

**Q1560: What is Selenium?**
Browser automation tool (for scraping dynamic JS sites).

---

## 🔹 157. Data Structures & Algo II (Questions 1561-1570)

**Q1561: What is a Linked List?**
Not built-in. Nodes pointing to next node.

**Q1562: What is a Hash Table?**
Dict/Set implementation.

**Q1563: What is Binary Search Tree?**
Sorted tree structure.

**Q1564: What is DFS vs BFS?**
*   DFS: Depth First Search (Stack).
*   BFS: Breadth First Search (Queue).

**Q1565: How to find max depth of a tree?**
Recursion. `1 + max(left, right)`.

**Q1566: How to detect cycle in linked list?**
Floyd’s Cycle-Finding (Tortoise and Hare).

**Q1567: What is Dynamic Programming?**
Breaking problem into subproblems + Memoization.

**Q1568: What is Big O notation?**
Measure of algorithm complexity (Time/Space).

**Q1569: What is the complexity of Python sort?**
O(n log n).

**Q1570: What is the complexity of list access?**
O(1).

---

## 🔹 158. Exceptions & Logging II (Questions 1571-1580)

**Q1571: How to get line number of error?**
`traceback` module or `sys.exc_info()`.

**Q1572: How to rotate log files?**
`RotatingFileHandler` (Size/Time based).

**Q1573: How to log to console and file?**
Add two handlers to the logger.

**Q1574: What is log level?**
DEBUG < INFO < WARNING < ERROR < CRITICAL.

**Q1575: How to format log messages?**
`Formatter('%(asctime)s - %(message)s')`.

**Q1576: What is assert used for in tests?**
Validating condition is True.

**Q1577: Can you nest try-except blocks?**
Yes.

**Q1578: What happens with unhandled exception in thread?**
Thread dies. Main program might stay alive (depends on implementation).

**Q1579: How to catch keyboard interrupt?**
`except KeyboardInterrupt`.

**Q1580: How to create a custom error class?**
(Refresher) Inherit `Exception`.

---

## 🔹 159. Advanced IO (Questions 1581-1590)

**Q1581: What is standard input/output?**
`sys.stdin`, `sys.stdout`.

**Q1582: How to redirect stdout to a file?**
`sys.stdout = open('log.txt', 'w')`.

**Q1583: How to read environment variables?**
(Refresher) `os.environ`.

**Q1584: How to execute shell command?**
`subprocess.run()`.

**Q1585: How to pipe commands?**
`subprocess.Popen(..., stdout=PIPE)`.

**Q1586: What is a file descriptor?**
Integer handle for open file in OS.

**Q1587: How to lock a file?**
`fcntl.flock` (Unix) or `msvcrt` (Windows).

**Q1588: How to watch a directory for changes?**
`watchdog` library.

**Q1589: How to serialize object to bytes?**
`pickle.dumps(obj)`.

**Q1590: How to compress a file?**
`gzip` or `zipfile` module.

---

## 🔹 160. Logic Puzzles (Questions 1591-1600)

**Q1591: How to check if two strings are anagrams?**
`Counter(s1) == Counter(s2)`.

**Q1592: How to find missing number in 1..N?**
Sum(1..N) - Sum(List).

**Q1593: How to rotate array by K?**
(Refresher) Slicing.

**Q1594: How to find intersection of two arrays?**
Set intersection.

**Q1595: How to reverse words in a sentence?**
`" ".join(s.split()[::-1])`.

**Q1596: How to check balanced parentheses?**
Use a stack. Push `(`, pop on `)`.

**Q1597: How does Python determine truthiness of objects?**
Calls `__bool__` or `__len__`.

**Q1598: Can you use eval() safely?**
No. It executes arbitrary code. Use `ast.literal_eval` for data.

**Q1599: What is the difference between copy.copy() and copy.deepcopy()?**
(Refresher) Shallow vs Recursive copy.

**Q1600: What’s the result of None == False?**
(Refresher) False.

---

## 🔹 161. Memory & Internals IV (Questions 1601-1610)

**Q1601: What are slots and how do they save memory?**
`__slots__` tells Python to allocate a static amount of memory for a known set of attributes, avoiding the creation of a dynamic `__dict__` for each instance.

**Q1602: What is bisect and when is it used?**
Library for maintaining a sorted list without having to sort each time. `bisect.insort(list, item)` inserts element in correct position.

**Q1603: How does functools.lru_cache() work?**
Decorator that wraps a function with a memoizing callable that helps in saving time when an expensive or I/O bound function is periodically called with the same arguments.

**Q1604: What is the use of itertools.permutations()?**
Returns successive r-length permutations of elements in the iterable. `permutations('ABCD', 2)` -> AB AC AD ...

**Q1605: How to serialize objects using pickle?**
`pickle.dump(obj, file)`. (Be careful: Unpickling untrusted data is a security risk).

**Q1606: What does pathlib.Path.glob() do?**
Iterates over the subtree rooted at the path yielding objects matching the pattern properly (like `**/*.py`).

**Q1607: What is the "EAFP" principle in Python?**
"Easier to Ask for Forgiveness than Permission". Code assumes keys/attributes exist and catches exceptions if they don't (Try/Except methodology).

**Q1608: What does if __debug__: do?**
Runs code only if Python is not started with `-O` optimization flag. `__debug__` is True by default.

**Q1609: How to avoid deeply nested if-statements?**
Guard clauses (Return early), Dictionary Dispatch, or extracting methods.

**Q1610: How to enforce type safety with mypy?**
Run `mypy script.py` to check type hints statically.

---

## 🔹 162. Testing & Environment II (Questions 1611-1620)

**Q1611: How to write readable, testable functions?**
Small, single responsibility, pure functions (no side effects), clear naming.

**Q1612: 🎯 Batch 15 (1,400 ➝ 1,500)?**
(Meta-question: Skipping).

**Q1613: 📤 Export to JSON/Excel/Notion?**
(Meta-question: Skipping).

**Q1614: 💡 Begin batch-wise answers with code + explanations?**
(Meta-question: Skipping).

**Q1615: What does python -m venv do?**
Creates a lightweight "virtual environment" with its own site directories, optionally isolated from system site directories.

**Q1616: How to check Python’s memory usage during runtime?**
`tracemalloc` module or `psutil`.

**Q1617: What is the PYTHONSTARTUP environment variable?**
File path to a script run when the interactive interpreter starts up (like bashrc for Python).

**Q1618: How to get the current recursion limit in Python?**
`sys.getrecursionlimit()`.

**Q1619: What does python -i script.py enable?**
Runs the script and enters interactive mode immediately after script terminates. Great for debugging variables.

**Q1620: How do you access a variable defined in an outer function?**
Read-only: directly. Modify: Use `nonlocal`.

---

## 🔹 163. Scope & Strings (Questions 1621-1630)

**Q1621: What is LEGB rule in Python scope resolution?**
(Refresher) Scope search order: Local -> Enclosing -> Global -> Built-in.

**Q1622: What’s the effect of using nonlocal?**
Marks a variable as belonging to the nested (enclosing) scope, allowing modification.

**Q1623: What is the difference between local, global, and built-in scope?**
*   Local: Inside function.
*   Global: Module level.
*   Built-in: Python internals (`len`, `str`).

**Q1624: What happens when you shadow a built-in function?**
The local name overrides the built-in. `list = [1,2]; list([3])` raises Error because `list` is now a list object, not the class.

**Q1625: How does str.encode() work?**
Converts Unicode string to bytes using specified encoding (default UTF-8).

**Q1626: How do you remove trailing digits from a string?**
`s.rstrip('0123456789')`.

**Q1627: What does string.punctuation contain?**
String of all punctuation characters (`!"#$%&\'()*+,-./:;<=>?@[\\]^_{|}~`).

**Q1628: How to align multiple string columns in tabular format?**
`f"{val:<10} {val2:>5}"` or `str.ljust()/rjust()`.

**Q1629: What is the effect of str.partition()?**
(Refresher) Splits at first occurrence into (head, sep, tail).

**Q1630: How to remove duplicates while preserving order in a list?**
`list(dict.fromkeys(lst))` (Fastest in Py3.7+).

---

## 🔹 164. Lists & Tuples Logic (Questions 1631-1640)

**Q1631: Why is list.pop() from front slower than back?**
Pop from back is O(1). Pop from front is O(n) because all subsequent elements must be shifted in memory.

**Q1632: What is the use of reversed(list) vs list[::-1]?**
(Refresher) Iterator vs New List.

**Q1633: How to perform element-wise operations on two lists?**
List comprehension `[a+b for a,b in zip(l1, l2)]`.

**Q1634: What’s the result of list slicing with a step of 0?**
`ValueError: slice step cannot be zero`.

**Q1635: Can tuples contain mutable objects?**
Yes. `t = ([1],)`. The tuple is immutable (cannot replace list), but list content can change.

**Q1636: What happens when you sort a list of tuples?**
Sorts by first element; if equal, uses second element (Lexicographical sort).

**Q1637: How to use tuple unpacking in function return values?**
`def f(): return 1, 2`. `a, b = f()`.

**Q1638: Can tuples be used in JSON serialization?**
Yes, but they are converted to JSON Arrays (same as lists). Upon decoding, they usually return as lists.

**Q1639: What does tuple.count() return?**
Number of occurrences of a value.

**Q1640: How do you iterate through a dict safely while modifying it?**
Iterate over a copy directly or list of keys. `for k in list(d.keys()): ...`.

---

## 🔹 165. Dicts & Sets Advanced (Questions 1641-1650)

**Q1641: What happens when you assign a new value to an existing key?**
Overwrites value.

**Q1642: How to invert a dictionary with non-unique values?**
Values become keys, original keys become list of values. `d_inv = {}; [d_inv.setdefault(v, []).append(k) for k, v in d.items()]`.

**Q1643: What’s the output of comparing two dicts?**
True if they have same (key, value) pairs regardless of order.

**Q1644: What is the default type of dict.items()?**
`dict_items` view object.

**Q1645: How does set.pop() behave?**
Removes and returns an arbitrary element. Raises KeyError if empty.

**Q1646: What happens if you pass a dict to the set() constructor?**
Creates a set of the dictionary's **keys**.

**Q1647: Can sets be used as keys in another set?**
No (Sets are mutable/unhashable). Use `frozenset`.

**Q1648: How to intersect multiple sets efficiently?**
`set.intersection(s1, s2, s3, ...)` or `s1 & s2 & s3`.

**Q1649: What’s the result of chaining set.union()?**
Union of all sets involved.

**Q1650: What is the result of calling print(print("Hi"))?**
Prints "Hi", then prints "None" (return value of inner print).

---

## 🔹 166. Functions & Lambdas IV (Questions 1651-1660)

**Q1651: How can you define a function with no arguments or return?**
`def f(): pass`.

**Q1652: Can you override a built-in function name?**
Yes (Shadowing).

**Q1653: How does Python treat callables differently than other objects?**
They implement `__call__`. `callable()` checks this.

**Q1654: What does lambda *args: args return?**
A tuple containing all passed arguments.

**Q1655: Can you use lambdas inside class methods?**
Yes.

**Q1656: Why do lambdas capture variables, not values?**
(Refresher) Late binding closures.

**Q1657: Can lambdas be serialized?**
Not easily with `pickle`. Libraries like `dill` can handle it.

**Q1658: How to sort using lambda with custom keys?**
`sort(key=lambda x: (x.prop1, -x.prop2))`.

**Q1659: Can a list comprehension create a list of dictionaries?**
Yes. `[{'i': i} for i in range(5)]`.

**Q1660: What is the result of comprehension on an empty list?**
Empty list.

---

## 🔹 167. Loops & Logic IV (Questions 1661-1670)

**Q1661: Can you use a comprehension to generate a 2D matrix?**
Yes. `[[0]*w for _ in range(h)]`.

**Q1662: What is the result of [x for x in range(3) if x > 3]?**
`[]` (Condition never met).

**Q1663: How to write nested comprehensions using multiple for?**
(Refresher) Flat list from nested structure.

**Q1664: How to break out of deeply nested loops?**
Raise exception or return from function. Flag variables are clunky.

**Q1665: How to maintain loop counter manually in while loop?**
`i=0; while... i+=1`.

**Q1666: What is loop unrolling and is it relevant in Python?**
Manual unrolling rarely helps in Python (interpreter overhead dominates).

**Q1667: What happens if you use break in a finally block?**
Not possible/SyntaxError in older Python. In newer, it silently cancels the exception if one was active. (Bad practice).

**Q1668: How to use enumerate() with a custom start index?**
`enumerate(iterable, start=1)`.

**Q1669: How to catch only file-related exceptions?**
`except (OSError, IOError)`.

**Q1670: How to retry a function if an exception occurs?**
Loop with try-except blocks.

---

## 🔹 168. Exceptions & Classes IV (Questions 1671-1680)

**Q1671: How to raise a warning without halting execution?**
`warnings.warn("Message")`.

**Q1672: What’s the benefit of custom exception classes?**
More granular error handling (`except MySpecificError`). Clarity.

**Q1673: What is the result of catching an exception and re-raising it?**
Propagates up the stack to next handler.

**Q1674: How does Python implement operator overloading?**
Magic methods (`__add__`, `__mul__`, etc).

**Q1675: What is a descriptor?**
(Refresher) Object managing attribute access (`__get__`).

**Q1676: What is the role of __call__()?**
(Refresher) Makes instance callable.

**Q1677: How does Python support interface-like behavior?**
Abstract Base Classes (`abc.ABC`).

**Q1678: What does object.__repr__() return?**
`<module.Class object at 0xId>`.

**Q1679: Can you use decorators to cache function results?**
Yes (`lru_cache`).

**Q1680: What happens if a decorator does not return a function?**
It causes `TypeError` when you try to call specific decorated function later (unless it returns a callable object).

---

## 🔹 169. Decorators & Files V (Questions 1681-1690)

**Q1681: How to unit test a decorated function?**
Access `func.__wrapped__` if available, or test the wrapper performance/logic.

**Q1682: Can you apply a decorator to a lambda?**
Syntax limits this. `decorator(lambda...)` works functionally but `@` syntax doesn't.

**Q1683: How can decorators modify class behavior?**
Class decorators receive the `class` object and can modify `cls.__dict__` or generic attributes.

**Q1684: How to read file contents into a list of words?**
`f.read().split()`.

**Q1685: How to check if a file exists and is writable?**
`os.access(path, os.W_OK)`.

**Q1686: What’s the difference between 'rb' and 'r'?**
(Refresher) Bytes vs String.

**Q1687: What is the use of os.path.join()?**
(Refresher) Cross-platform path concat.

**Q1688: What happens when you open the same file twice?**
OS allows it (usually). You get two independent file handles/cursors.

**Q1689: How to get the current weekday as string?**
`dt.strftime("%A")`.

**Q1690: How to create a datetime from a timestamp?**
(Refresher) `fromtimestamp()`.

---

## 🔹 170. DateTime & Meta II (Questions 1691-1700)

**Q1691: How to set timezone on naive datetime?**
(Refresher) `dt.replace(tzinfo=...)`.

**Q1692: What’s the use of calendar.isleap()?**
Checks if year is leap year.

**Q1693: How to add N months to a date safely?**
`dateutil.relativedelta(months=+N)` handles roll-over correctly (e.g. Jan 31 + 1 month -> Feb 28).

**Q1694: How to inspect arguments passed to a class constructor?**
Override `__new__` or `__init__` and inspect locals/args.

**Q1695: What is __annotations__ and how is it useful?**
(Refresher) Stores type hints. Useful for data validation libs (Pydantic).

**Q1696: How to track object creation using metaclass?**
Override `__call__` in metaclass.

**Q1697: What is the use of inspect.signature()?**
(Refresher) Parameter introspection.

**Q1698: What is the output of globals().keys()?**
List of global variable names.

**Q1699: What is the difference between and vs & in conditionals?**
*   `and`: Logical AND (Short-circuiting boolean).
*   `&`: Bitwise AND (Integers) or Set Intersection.

**Q1700: What is the behavior of bool([]) vs bool([[]])?**
*   `bool([])`: False (Empty).
*   `bool([[]])`: True (List containing one empty list is not empty).

---

## 🔹 171. Advanced Testing (Questions 1701-1710)

**Q1701: How to mock a database connection in pytest?**
Use `unittest.mock.patch` or a pytest fixture yielding a fake connection object.

**Q1702: What is monkeypatch fixture in pytest?**
Helper to safely set attributes/env vars for a test and restore them afterwards.

**Q1703: How to test async functions with pytest?**
Install `pytest-asyncio` plugin and mark tests `@pytest.mark.asyncio`.

**Q1704: What is property-based testing?**
Testing with randomly generated inputs satisfying properties (e.g., using `hypothesis` library).

**Q1705: How to check if a function was called N times?**
`mock_obj.assert_called_with(...)` or `mock_obj.call_count == N`.

**Q1706: What is a side_effect in Mock?**
Allows the mock to raise an exception or return different values sequentially.

**Q1707: How to spy on a real object?**
`wraps=real_obj` in Mock constructor.

**Q1708: What’s the difference between Mock and MagicMock?**
`MagicMock` pre-implements all magic methods (`__len__`, `__getitem__` etc). `Mock` does not.

**Q1709: How to measure test execution time?**
`pytest --durations=N`.

**Q1710: How to group tests in pytest?**
Use markers (`@pytest.mark.groupname`) or classes.

---

## 🔹 172. Performance & Optimizations (Questions 1711-1720)

**Q1711: What is Cython?**
(Refresher) Static compiler for Python.

**Q1712: How to use Numba for JIT compilation?**
`@numba.jit` decorator accelerates numeric code.

**Q1713: What is PyPy’s main advantage?**
JIT compiler makes pure Python loops much faster.

**Q1714: How to optimize memory for millions of objects?**
Use `__slots__` or Arrays (NumPy/Structs).

**Q1715: What is intern() for strings?**
(Refresher) `sys.intern()`.

**Q1716: How to use multiprocessing.Pool?**
`with Pool(n) as p: p.map(func, data)`.

**Q1717: What is the cost of context switching in threads?**
Time wasted saving/restoring thread state (registers/stack). High overhead in Python threads due to GIL flickering.

**Q1718: How to preallocate a list of size N?**
`lst = [None] * N`.

**Q1719: Why is list comprehension faster than for-loop append?**
Opcode optimization. `append` is a method lookup; comprehension has dedicated bytecode (`LIST_APPEND`).

**Q1720: How to lazily load a large dataset?**
Generators or memory mapping (`mmap`).

---

## 🔹 173. Web Frameworks III (Questions 1721-1730)

**Q1721: What is a Blueprint in Flask?**
(Refresher) Module for app organization.

**Q1722: How does Django handle transactions?**
`transaction.atomic()` context manager.

**Q1723: What is FastAPI’s main advantage over Flask?**
(Refresher) Async native, type-safe, auto-docs.

**Q1724: What is WSGI vs ASGI?**
(Refresher) Synchronous vs Asynchronous interface.

**Q1725: How to handle CORS in Django?**
`django-cors-headers` middleware.

**Q1726: What is a ViewSet in Django REST Framework?**
Class that combines logic for listing, creating, retrieving, updating, and destroying objects.

**Q1727: How to run Flask in production?**
Use Gunicorn or uWSGI behind Nginx. Not `flask run`.

**Q1728: What is Celery Beat?**
Scheduler for periodic tasks.

**Q1729: What is dependency injection in FastAPI?**
(Refresher) `Depends()`.

**Q1730: How to stream response in Flask?**
Generator function passed to `Response()`.

---

## 🔹 174. Packaging & Setup II (Questions 1731-1740)

**Q1731: What is setup.cfg?**
Declarative configuration for setuptools (standardized replacement for setup.py arguments).

**Q1732: What is poetry?**
Modern dependency management and packaging tool (replaces pip/setuptools/venv).

**Q1733: How to pin dependencies?**
`pip freeze > requirements.txt` with exact versions `==1.0.0`.

**Q1734: What is a wheel cache?**
Pip caches downloaded wheels to speed up reinstalls.

**Q1735: How to create a namespace package?**
(Refresher) Omit `__init__.py`.

**Q1736: What is __main__.py used for in a package?**
Allows package to be executed as a script `python -m pkg`.

**Q1737: How to include data files in a package?**
`package_data` arg in setup or `MANIFEST.in`.

**Q1738: What is editable install?**
(Refresher) `pip install -e .`.

**Q1739: How to build a Docker image for Python app?**
`FROM python:slim`, `COPY reqs`, `RUN pip install`, `CMD [...]`.

**Q1740: What is .dockerignore?**
Files to exclude from build context (like `.git`, `venv`, `__pycache__`).

---

## 🔹 175. Security & Cryptography (Questions 1741-1750)

**Q1741: How to hash a password with salt?**
`hashlib.pbkdf2_hmac` or `bcrypt.hashpw(pass, salt)`.

**Q1742: What is HMAC?**
Hash-based Message Authentication Code. Verifies integrity + authenticity.

**Q1743: How to generate cryptographically strong random numbers?**
`secrets` module (`secrets.token_hex`).

**Q1744: What’s the danger of random.random() for security?**
It's pseudo-random (deterministic). Predictable if seed is known.

**Q1745: How to encrypt data symmetrically in Python?**
`cryptography.fernet`.

**Q1746: What is a timing attack?**
Deducing secrets by measuring time taken to verify (e.g., string comparison `==` returns early). Use `hmac.compare_digest`.

**Q1747: How to validate an email address properly?**
`email-validator` library (regex is often insufficient).

**Q1748: How to sanitize inputs against XSS?**
Use templating engines (Jinja2) which auto-escape HTML, or `bleach`.

**Q1749: What is clickjacking?**
Tricking user to click hidden UI. Prevent with `X-Frame-Options` header.

**Q1750: How to manage secrets securely?**
Environment variables + Secret Managers (Vault/AWS Secrets), never hardcoded.

---

## 🔹 176. Networking & Sockets II (Questions 1751-1760)

**Q1751: What is a blocking vs non-blocking socket?**
*   Blocking: Waits until data ready.
*   Non-blocking: Returns error immediately if not ready.

**Q1752: How to get IP address of hostname?**
`socket.gethostbyname('google.com')`.

**Q1753: How to send a UDP packet?**
`sock.sendto(bytes, (ip, port))`.

**Q1754: What is selection/polling in sockets?**
`select` module. Monitors multiple sockets for I/O readiness.

**Q1755: How to parse a URL?**
`urllib.parse.urlparse()`.

**Q1756: How to make a HEAD request?**
`requests.head(url)`.

**Q1757: What is chunked transfer encoding?**
Sending data in chunks (for streaming large content). `Response(gen())`.

**Q1758: How to create a WebSocket client?**
Use `websockets` library (async) or `websocket-client`.

**Q1759: What is TCP keepalive?**
Heartbeat packets to keep connection open through firewalls. `socket.SO_KEEPALIVE`.

**Q1760: How do you perform DNS lookup in Python?**
`socket.getaddrinfo()`.

---

## 🔹 177. Advanced Concurrency (Questions 1761-1770)

**Q1761: What is a semaphore?**
Counter-based lock allowing N concurrent accesses.

**Q1762: What is an Event object in threading?**
Simple communication mechanism. One thread signals an event, others wait for it.

**Q1763: What’s the difference between Thread and Process pool?**
(Refresher) Shared memory vs Isolation.

**Q1764: How to terminate a thread safely?**
Set a `stop_event` flag that the thread checks periodically. Killing threads is unsafe.

**Q1765: What is asyncio.Queue?**
Queue designed for async/await producers and consumers.

**Q1766: How to bridge sync and async code?**
`asgiref.sync.async_to_sync` or wrapping blocking code in `loop.run_in_executor`.

**Q1767: What does loop.run_until_complete() do?**
Runs event loop until the future completes.

**Q1768: What is contextvars module?**
Manages context-local state (like thread-local but works with async tasks).

**Q1769: How to protect shared data in asyncio?**
Usually not needed (single thread), but `asyncio.Lock` needed if `await` yields during critical section.

**Q1770: What is a Barrier in threading?**
Primitives where threads wait until N threads reach the barrier point.

---

## 🔹 178. Data Handling & ETL (Questions 1771-1780)

**Q1771: How to read large CSV in chunks?**
`pd.read_csv(chunksize=N)`.

**Q1772: How to merge two DataFrames on a key?**
`pd.merge(df1, df2, on='key')`.

**Q1773: What is pivot table in Pandas?**
Reshapes data. `df.pivot_table()`.

**Q1774: How to vectorize a function in NumPy?**
`np.vectorize(func)`.

**Q1775: What is Apache Arrow and how does Python use it?**
Values stored in columnar memory format. `pyarrow` facilitates zero-copy transfers between Pandas/Spark.

**Q1776: How to read a Parquet file?**
`pd.read_parquet()`.

**Q1777: What is regex capture group?**
Parentheses `(pattern)` capture part of the match.

**Q1778: How to replace string using regex group reference?**
`re.sub(r'(\d+)', r'<\1>', s)`.

**Q1779: What does re.compile() do?**
Pre-compiles regex pattern to object for faster reuse.

**Q1780: How to validate JSON schema?**
`jsonschema.validate()`.

---

## 🔹 179. System & OS Advanced (Questions 1781-1790)

**Q1781: How to get disk usage via Python?**
`shutil.disk_usage("/")`.

**Q1782: How to walk directory tree bottom-up?**
`os.walk(path, topdown=False)`.

**Q1783: How to change file owner?**
`shutil.chown()` or `os.chown()`.

**Q1784: How to create a hard link?**
`os.link(src, dst)`.

**Q1785: What is memory mapping (mmap)?**
Accessing file on disk as if it were in memory (byte array). Good for huge files.

**Q1786: How to send a POSIX signal?**
`os.kill(pid, signal.SIGTERM)`.

**Q1787: How to get current user ID?**
`os.getuid()` (Unix).

**Q1788: What is a named pipe (FIFO)?**
IPC mechanism. `os.mkfifo()`.

**Q1789: How to get terminal size?**
`shutil.get_terminal_size()`.

**Q1790: How to daemonize a process?**
Double-fork technique (or use `python-daemon`).

---

## 🔹 180. Miscellaneous II (Questions 1791-1800)

**Q1791: How to format currency in Python?**
`locale.currency()`.

**Q1792: What is the difflib module?**
Compares sequences/files. `SequenceMatcher`.

**Q1793: How to parse arguments with subcommands (git style)?**
`argparse` subparsers.

**Q1794: What is textwrap.dedent()?**
Removes common leading indentation from multiline strings.

**Q1795: How to open a URL in default browser?**
`webbrowser.open(url)`.

**Q1796: What is a UUID5?**
Deterministic UUID based on namespace + name (SHA-1 hash).

**Q1797: How to find the MRO of a class programmatically?**
`Class.mro()`.

**Q1798: How many keywords are there in Python?**
Around 35. `len(keyword.kwlist)`.

**Q1799: What does help() function invoke?**
The built-in help system (pydoc).

**Q1800: How to inspect all local variables at a crash?**
Use `cgitb` (Traceback manager) or `pdb.pm()` (Post-Mortem).

---

## 🔹 181. Deployment & Optimization (Questions 1801-1810)

**Q1801: How to use WSGI server gunicorn?**
`gunicorn app:app -w 4`.

**Q1802: What is the benefit of using -O flag?**
Optimizes bytecode (removes asserts and `__debug__` blocks).

**Q1803: How to profile memory allocation over time?**
`mprof run script.py` (memory_profiler).

**Q1804: What is PyInstaller?**
Bundles Python app + dependencies into a standalone executable.

**Q1805: How to securely store API keys?**
Environment variables or Secret Store. Not in code.

**Q1806: What is a requirements.lock?**
File defining exact versions (hashes) of dependency tree (e.g., from Pipenv/Poetry).

**Q1807: How to optimize Docker image size for Python?**
Multi-stage builds, use `alpine` or `slim` tags, don't install caching files (`--no-cache-dir`).

**Q1808: What is uWSGI?**
Another high-performance WSGI server (C-based).

**Q1809: How to run a script on a schedule?**
Cron jobs (OS level) or `schedule` library (Python level).

**Q1810: What is the purpose of Procfile?**
Used by Heroku/Dokku to determine startup commands. `web: gunicorn app:app`.

---

## 🔹 182. Advanced Coding (Questions 1811-1820)

**Q1811: How to implement a Singleton in Python?**
Metaclass, Decorator, or module-level global instance.

**Q1812: What is the Factory Pattern?**
Creational pattern. A function/method creates objects without specifying exact class.

**Q1813: What is Dependency Injection?**
Passing dependencies (objects) to clients (functions/classes) rather than having them build them.

**Q1814: How to implement Observer Pattern?**
Subject maintains list of Observers and notifies them on state change.

**Q1815: What is a mixin class?**
Class providing methods to other classes but not designed to stand alone.

**Q1816: How to use ABCs (Abstract Base Classes)?**
Inherit `ABC` and use `@abstractmethod`. Enforces implementation in subclasses.

**Q1817: What is the difference between composition and inheritance?**
*   Composition: "Has-a" (Flexible).
*   Inheritance: "Is-a" (Rigid).

**Q1818: What is polymorphism in Python?**
Using objects of different types through a uniform interface (Duck Typing).

**Q1819: How to override new operator?**
`__new__`.

**Q1820: What is method chaining?**
Methods returning `self` to allow contiguous calls. `obj.step1().step2()`.

---

## 🔹 183. Asyncio Deep Dive (Questions 1821-1830)

**Q1821: How to handle exceptions in asyncio.gather?**
`return_exceptions=True` allows other tasks to finish even if one fails.

**Q1822: What is an asyncio Future?**
Low-level object representing an eventual result. `Task` wraps `Future`.

**Q1823: How to run blocking code in asyncio?**
`loop.run_in_executor(None, func)`. Points to ThreadPoolExecutor by default.

**Q1824: What is asyncio.sleep()?**
Non-blocking sleep. Yields control to event loop.

**Q1825: How to limit concurrency in asyncio?**
Use `asyncio.Semaphore(N)`.

**Q1826: What is a coroutine chain?**
Coroutines awaiting other coroutines.

**Q1827: How to debug an asyncio hang?**
`PYTHONASYNCIODEBUG=1`. `loop.set_debug(True)`.

**Q1828: What is a TaskGroup (Python 3.11+)?**
Context manager for managing a group of tasks. Safer than `gather`.

**Q1829: How to run asyncio code from a synchronous function?**
`asyncio.run()`.

**Q1830: How to implement a periodic async task?**
`while True: await task(); await asyncio.sleep(interval)`.

---

## 🔹 184. Networking & Web (Questions 1831-1840)

**Q1831: What is HTTP persistent connection?**
Keep-Alive. Reusing TCP connection for multiple requests.

**Q1832: What is the requests.Session object?**
Persists cookies and connection pooling across requests.

**Q1833: How to handle redirects in requests?**
`allow_redirects=False` to stop auto-redirect.

**Q1834: How to verify SSL certificate?**
`verify=True` (default) or path to CA bundle.

**Q1835: What is multipart/form-data?**
Encoding for uploading files via HTTP POST.

**Q1836: How to parse HTML with BeautifulSoup?**
`Soup(html, 'html.parser')`.

**Q1837: What is rate limiting?**
Restricting number of requests per client.

**Q1838: How to mock an API response?**
`responses` library or `unittest.mock`.

**Q1839: What is a webhook?**
HTTP callback. Server calls your URL when event happens.

**Q1840: How to encode URL parameters?**
`urllib.parse.urlencode()`.

---

## 🔹 185. NumPy & Data (Questions 1841-1850)

**Q1841: How to create a NumPy array of zeros?**
`np.zeros((rows, cols))`.

**Q1842: How to reshape an array?**
`arr.reshape((new_r, new_c))`.

**Q1843: What is slicing in NumPy?**
`arr[start:stop, start:stop]`. Returns a **view** (no copy).

**Q1844: How to perform matrix multiplication?**
`np.dot(a, b)` or `a @ b`.

**Q1845: What is a boolean mask in NumPy?**
Array of booleans used to filter elements. `arr[arr > 5]`.

**Q1846: How to stack arrays vertically?**
`np.vstack((a, b))`.

**Q1847: How to calculate mean and std dev?**
`np.mean()`, `np.std()`.

**Q1848: How to save/load NumPy array?**
`np.save('file.npy', arr)`, `np.load()`.

**Q1849: What is the advantage of NumPy over valid Python lists?**
Contiguous memory, vectorization (SIMD), smaller overhead.

**Q1850: How to generate random numbers in NumPy?**
`np.random.rand()`.

---

## 🔹 186. Pandas & Analytics (Questions 1851-1860)

**Q1851: How to read a CSV with specific columns?**
`pd.read_csv(f, usecols=['a', 'b'])`.

**Q1852: How to filter DataFrame rows?**
`df[df['col'] > 0]`.

**Q1853: How to group by multiple columns?**
`df.groupby(['c1', 'c2']).sum()`.

**Q1854: How to handle duplicate rows?**
`df.drop_duplicates()`.

**Q1855: How to convert column to datetime?**
`pd.to_datetime(df['col'])`.

**Q1856: How to merge DataFrames?**
`pd.merge()` (Join) or `pd.concat()` (Stack).

**Q1857: What is iloc vs loc?**
(Refresher) Integer-based vs Label-based.

**Q1858: How to apply a function to a column?**
`df['col'].apply(func)`.

**Q1859: How to get simple statistics of data?**
`df.describe()`.

**Q1860: How to write DataFrame to JSON?**
`df.to_json()`.

---

## 🔹 187. System & Shell (Questions 1861-1870)

**Q1861: How to get hostname of machine?**
`socket.gethostname()`.

**Q1862: How to find path of the python executable?**
`sys.executable`.

**Q1863: How to check if a process is running?**
`psutil.pid_exists(pid)`.

**Q1864: How to schedule a shutdown?**
`os.system("shutdown /s /t 1")` (Windows).

**Q1865: How to expand user home directory ~?**
`os.path.expanduser("~/file")`.

**Q1866: How to check available disk space?**
`shutil.disk_usage(".")`.

**Q1867: How to get current thread ID?**
`threading.get_ident()`.

**Q1868: How to copy permission bits only?**
`shutil.copymode(src, dst)`.

**Q1869: How to get file creation time?**
`os.path.getctime()`.

**Q1870: How to perform glob matching recursively?**
`glob.glob("**/*.py", recursive=True)`.

---

## 🔹 188. Regex & Parsing (Questions 1871-1880)

**Q1871: How to match an email address (simple)?**
`r"[^@]+@[^@]+\.[^@]+"`.

**Q1872: What is basic vs greedy matching?**
*   `*`: Greedy (Matches max possible).
*   `*?`: Non-greedy (Matches min possible).

**Q1873: How to split string by regex?**
`re.split(pattern, string)`.

**Q1874: How to replace only first occurrence?**
`re.sub(pat, repl, str, count=1)`.

**Q1875: What is regex flag re.IGNORECASE?**
Makes matching case-insensitive.

**Q1876: How to extract all numbers from text?**
`re.findall(r"\d+", text)`.

**Q1877: What is re.match() vs re.search()?**
*   `match`: Checks START of string only.
*   `search`: Checks ANYWHERE in string.

**Q1878: How to name a regex group?**
`(?P<name>...)`.

**Q1879: How to parse CLI flags manually?**
Use `sys.argv`. Iterate and check elements.

**Q1880: How to parse JSON with comments?**
Standard `json` module doesn't support it. Strip comments or use `json5`.

---

## 🔹 189. Type Hinting & MyPy (Questions 1881-1890)

**Q1881: How to type a list of integers?**
`List[int]` (old) or `list[int]` (Py3.9+).

**Q1882: What is Optional[T]?**
Same as `Union[T, None]` or `T | None`.

**Q1883: How to define a TypedDict?**
`class Point(TypedDict): x: int; y: int`.

**Q1884: What is Any type?**
Disables type checking for that variable. compatibility escape hatch.

**Q1885: What is Callable[[int], str]?**
Function taking int and returning str.

**Q1886: How to type hint `*args`?**
Hint the type of *one* element. `def f(*args: int)` means all args are ints.

**Q1887: What is cast() in typing?**
`cast(Type, val)`. Tells static checker to treat val as Type (Runtime no-op).

**Q1888: What is Final type?**
Variable shouldn't be reassigned. `x: Final = 1`.

**Q1889: What is Literal type?**
Restricts value to specific literals. `mode: Literal['r', 'w']`.

**Q1890: Where to put type stubs?**
`.pyi` files.

---

## 🔹 190. Debugging & Tools (Questions 1891-1900)

**Q1891: How to profile memory of a single function?**
`@profile` decorator from `memory_profiler`.

**Q1892: What is PDB post-mortem?**
Debugging after a crash. `python -m pdb -c continue script.py` or inside script `pdb.post_mortem()`.

**Q1893: How to enable verbose logging?**
`logging.getLogger().setLevel(logging.DEBUG)`.

**Q1894: How to pretty print JSON?**
`json.dumps(obj, indent=4)`.

**Q1895: What is black?**
Opinionated code formatter.

**Q1896: What is isort?**
Import sorter.

**Q1897: What is flake8?**
Linter combining pyflakes, pycodestyle, and mccabe.

**Q1898: How to list outdated packages?**
`pip list --outdated`.

**Q1899: What is pre-commit hook?**
Git hook script running checks (linting/testing) before commit.

**Q1900: How to measure execution time of a block?**
`timeit.default_timer()` or `contextlib.contextmanager` timer.

---

## 🔹 191. Architecture & Design (Questions 1901-1910)

**Q1901: What is the Singleton Pattern logic?**
Ensure a class has only one instance and provide a global point of access to it.

**Q1902: What is the Adapter Pattern?**
Wraps an interface to match another interface a client expects.

**Q1903: What is the Decorator Pattern in logic?**
Dynamically adds responsibilities to an object (using aggregation) without inheritance.

**Q1904: What is MVC architecture?**
Model (Data), View (UI), Controller (Logic). Python web frameworks often use MTV (Django) or MVC concepts.

**Q1905: What is Microservices Architecture?**
Splitting app into small, independent services communicating via APIs.

**Q1906: What is a REST API?**
(Refresher) Stateless architecture using HTTP verbs.

**Q1907: What is SOLID?**
SRP, OCP, LSP, ISP, DIP. (Design Principles).

**Q1908: What is Dependency Inversion Principle (DIP)?**
Depend on abstractions, not concretions.

**Q1909: What is Dry Principle?**
Don't Repeat Yourself.

**Q1910: What is KISS Principle?**
Keep It Simple, Stupid.

---

## 🔹 192. Advanced Python Features (Questions 1911-1920)

**Q1911: What is __init_subclass__?**
Hook method in parent class called when a subclass is defined. Easier than metaclasses.

**Q1912: What is __mro__ attribute?**
Tuple showing the method resolution order of a class.

**Q1913: What is a data descriptor?**
A descriptor that defines both `__get__` and `__set__`.

**Q1914: What is __slots__ caveat?**
Cannot add new attributes; inheritance merging slots can be tricky; mostly for memory.

**Q1915: What is the purpose of __del__?**
(Refresher) Finalizer.

**Q1916: What is __enter__ and __exit__?**
(Refresher) Context Manager protocol.

**Q1917: What does functools.singledispatch do?**
Creates a generic function that dispatches based on the type of the first argument.

**Q1918: What is yield from used for?**
Delegating to a sub-generator.

**Q1919: What is the purpose of coroutine decorator?**
(Historical) `asyncio.coroutine` used before `async def`.

**Q1920: What is f-string debugging syntax?**
`f"{var=}"` prints `var=value`.

---

## 🔹 193. Deployment & Cloud (Questions 1921-1930)

**Q1921: What is Serverless Python?**
Running functions (AWS Lambda, Azure Functions) without managing servers.

**Q1922: What is AWS Boto3?**
Python SDK for Amazon Web Services.

**Q1923: What is a WSGI container?**
Server (Gunicorn/uWSGI) that runs Python web apps.

**Q1924: What is Nginx used for with Python?**
Reverse proxy, SSL termination, static file serving.

**Q1925: What is CI/CD for Python?**
Continuous Integration/Deployment (GitHub Actions, GitLab CI) running tests/lints/deploy.

**Q1926: What is Blue-Green Deployment?**
Running two identical environments (Live/Idle) to switch traffic instantly.

**Q1927: What is Canary Deployment?**
Rolling out update to a small subset of users first.

**Q1928: What is Infrastructure as Code (IaC)?**
Managing infra via code (Terraform, CloudFormation).

**Q1929: What is 12-Factor App?**
(Refresher) Methodology for SaaS apps.

**Q1930: What is Kubernetes (K8s)?**
Container orchestration platform.

---

## 🔹 194. Database & ORM (Questions 1931-1940)

**Q1931: What is an ORM?**
Object-Relational Mapper (Maps Classes to Tables).

**Q1932: What is N+1 problem?**
Fetching related data in a loop (1 query for parent + N queries for children). Fix: Eager loading `select_related`.

**Q1933: What is a migration?**
Types of version control for database schemas.

**Q1934: What is a transaction?**
Atomic unit of work (Commit or Rollback).

**Q1935: What is indexing in DB?**
Data structure to speed up data retrieval.

**Q1936: What is NoSQL?**
Non-relational databases (MongoDB, Redis, Cassandra).

**Q1937: What is Redis used for?**
Caching, Message Broker, Session Store.

**Q1938: What is replication?**
Copying data to multiple servers for redundancy.

**Q1939: What is sharding?**
Splitting database horizontally across servers.

**Q1940: What is ACID?**
Atomicity, Consistency, Isolation, Durability.

---

## 🔹 195. Security & Auth (Questions 1941-1950)

**Q1941: What is OAuth2?**
Authorization framework (Access tokens).

**Q1942: What is OpenID Connect?**
Authentication layer on top of OAuth2 (Identity).

**Q1943: What is CORS?**
(Refresher) Browser security mechanism.

**Q1944: What is CSRF?**
(Refresher) Forging requests.

**Q1945: What is XSS?**
(Refresher) Injecting scripts.

**Q1946: What is a rainbow table?**
Precomputed table for cracking password hashes. Avoid with Salt.

**Q1947: What is Man-in-the-Middle attack?**
Attacker intercepts communication. Prevent with HTTPS.

**Q1948: What is Least Privilege Principle?**
Granting only minimum necessary permissions.

**Q1949: What is Zero Trust?**
Strict identity verification for every person/device trying to access resources.

**Q1950: What is OWASP Top 10?**
List of most critical web security risks.

---

## 🔹 196. Modern Python (Questions 1951-1960)

**Q1951: What is match-case?**
(Refresher) Structural pattern matching (Py3.10).

**Q1952: What is the zoneinfo module?**
Standard library support for IANA time zones (Py3.9).

**Q1953: What is assignment expression?**
(Refresher) `:=`.

**Q1954: What is __future__ annotations?**
Postponed evaluation of annotations (Py3.7+).

**Q1955: What is PEP 572?**
Assignment Expressions (`:=`).

**Q1956: What is PEP 484?**
Type Hints.

**Q1957: What is PEP 8?**
(Refresher) Style Guide.

**Q1958: What is PEP 20?**
The Zen of Python.

**Q1959: What is graphlib?**
Module for topological sort of graphs (Py3.9).

**Q1960: What is tomllib?**
Standard library for parsing TOML files (Py3.11).

---

## 🔹 197. Interview logic (Questions 1961-1970)

**Q1961: How to swap two numbers without temp?**
`a, b = b, a`.

**Q1962: How to check for palindrome number?**
`str(n) == str(n)[::-1]`.

**Q1963: How to find factorial?**
`math.factorial(n)`.

**Q1964: How to check for Armstrong number?**
Sum of digits raised to power of num_digits equals number.

**Q1965: How to print Fibonacci series?**
Loop adding previous two.

**Q1966: How to find prime factors?**
Divide by 2, then odd numbers until n became 1.

**Q1967: How to check for perfect number?**
Sum of divisors equals number.

**Q1968: How to find GCD of two numbers?**
`math.gcd(a, b)`.

**Q1969: How to check if two strings are anagrams?**
Sort and compare. `sorted(s1) == sorted(s2)`.

**Q1970: How to count occurrences of a char in string?**
`s.count(char)`.

---

## 🔹 198. Coding Challenges (Questions 1971-1980)

**Q1971: How to find longest word in a sentence?**
`max(s.split(), key=len)`.

**Q1972: How to remove vowels from string?**
Using regex or list comp.

**Q1973: How to find second largest number in list?**
`sorted(set(lst))[-2]`.

**Q1974: How to check substring presence?**
`sub in s`.

**Q1975: How to convert list to string?**
`"".join(lst)`.

**Q1976: How to sort dictionary by key?**
`dict(sorted(d.items()))`.

**Q1977: How to merge two lists?**
`l1 + l2` or `l1.extend(l2)`.

**Q1978: How to find common elements in two lists?**
`set(l1) & set(l2)`.

**Q1979: How to reverse a number?**
`int(str(n)[::-1])`.

**Q1980: How to check leap year?**
`calendar.isleap(y)`.

---

## 🔹 199. Specific Scenarios (Questions 1981-1990)

**Q1981: How to read a large file?**
Line by line generator.

**Q1982: How to copy a file?**
`shutil.copy()`.

**Q1983: How to delete a file?**
`os.remove()`.

**Q1984: How to get list of files in dir?**
`os.listdir()`.

**Q1985: How to create a directory?**
`os.makedirs()`.

**Q1986: How to check file existence?**
`os.path.exists()`.

**Q1987: How to get file size?**
`os.path.getsize()`.

**Q1988: How to get absolute path?**
`os.path.abspath()`.

**Q1989: How to execute system command?**
`subprocess.run()`.

**Q1990: How to get current directory?**
`os.getcwd()`.

---

## 🔹 200. Final Review (Questions 1991-2000)

**Q1991: What is the most important Python feature?**
Readability.

**Q1992: Why use Python for AI?**
Libraries (TensorFlow, PyTorch), simplicity, community.

**Q1993: Why use Python for Web?**
Fast development (Django/Flask), mature ecosystem.

**Q1994: Why use Python for Scripting?**
Cross-platform, standard library.

**Q1995: What is the main drawback of Python?**
Execution speed (interpreted).

**Q1996: How to improve Python speed?**
PyPy, Cython, NumPy, Asyncio.

**Q1997: What is the future of Python?**
Continued growth in AI/Data, speed improvements (Faster CPython).

**Q1998: Is Python strongly typed?**
Yes (Strongly but Dynamically typed). `1 + "1"` fails.

**Q1999: Is Python compiled?**
Yes, to bytecode.

**Q2000: Who created Python?**
Guido van Rossum.

---

## 🔹 201. Advanced Python Logic (Questions 2001-2010)

**Q2001: What is the purpose of the __bool__ method?**
Defines the truthiness of an object. If not defined, `__len__` is utilized.

**Q2002: How do you overload the + operator?**
Implement `__add__(self, other)`.

**Q2003: What implies a function is a generator?**
Usage of the `yield` keyword.

**Q2004: How can we make a class context manager?**
Implement `__enter__` and `__exit__`.

**Q2005: Can we access private variables outside the class?**
Yes, via name mangling `_ClassName__var`.

**Q2006: What is method overriding?**
Redefining a method in a child class that exists in the parent class.

**Q2007: What is method overloading?**
(Refresher) Multiple methods with same name but different args. Not natively supported.

**Q2008: What are mixins?**
(Refresher) Helper classes for inheritance.

**Q2009: What is the MRO?**
(Refresher) Method Resolution Order (C3).

**Q2010: What is the super() function?**
(Refresher) Access parent class methods.

---

## 🔹 202. Data Structures III (Questions 2011-2020)

**Q2011: What is a set comprehension?**
`{x for x in iterable}`.

**Q2012: What is the complexity of set operations?**
Avg O(1).

**Q2013: How to freeze a dictionary?**
Use `types.MappingProxyType(d)` or convert items to tuple.

**Q2014: How to implement a LRU Cache?**
`collections.OrderedDict` or `functools.lru_cache`.

**Q2015: What is a heap?**
Example of priority queue. `heapq` module implements min-heap.

**Q2016: How to find k-largest elements?**
`heapq.nlargest(k, iterable)`.

**Q2017: How to implement a Trie?**
Nested dictionaries. `trie = {'c': {'a': {'t': {'_end': True}}}}`.

**Q2018: What is a deque?**
Double-ended queue. O(1) appends/pops from both ends.

**Q2019: How to rotate a list efficiently?**
Using `deque.rotate()`.

**Q2020: What is a namedtuple?**
(Refresher) Tuple with named fields.

---

## 🔹 203. Functional Python II (Questions 2021-2030)

**Q2021: What is a pure function?**
Function where return value depends only on arguments and has no side effects.

**Q2022: What is a first-class function?**
(Refresher) Functions treated as objects.

**Q2023: What is a higher-order function?**
(Refresher) Takes/Returns function.

**Q2024: What is partial application?**
Fixing some arguments of a function. `functools.partial`.

**Q2025: What is closures?**
(Refresher) Function capturing outer scope.

**Q2026: What is a decorator?**
(Refresher) Wrapper modifying function behavior.

**Q2027: How to chain decorators?**
Stacking them.

**Q2028: What is map-filter-reduce?**
Functional primitives for processing iterables.

**Q2029: What is lazy evaluation?**
Evaluation only when needed (Generators).

**Q2030: What is immutability?**
State cannot change after creation. Key concept in functional programming.

---

## 🔹 204. Internals & GIL (Questions 2031-2040)

**Q2031: What is GIL?**
(Refresher) Global Interpreter Lock.

**Q2032: Does GIL affect I/O bound tasks?**
No. It is released during I/O.

**Q2033: Does GIL affect CPU bound tasks?**
Yes. Prevents true parallelism on threads.

**Q2034: How to bypass GIL?**
Multiprocessing or C-extensions releasing GIL.

**Q2035: What is reference counting?**
(Refresher) Memory management mechanism.

**Q2036: What is garbage collection?**
(Refresher) Cyclic GC.

**Q2037: What is __slots__?**
(Refresher) RAM optimization.

**Q2038: What is bytecode?**
(Refresher) `.pyc` instructions.

**Q2039: What is PVM?**
(Refresher) Python Virtual Machine.

**Q2040: What is dynamic typing?**
(Refresher) Runtime type checking.

---

## 🔹 205. Testing & Debugging III (Questions 2041-2050)

**Q2041: What is unittest?**
(Refresher) Standard testing lib.

**Q2042: What is pytest?**
(Refresher) Popular 3rd party testing lib.

**Q2043: What is coverage?**
(Refresher) Code execution metric.

**Q2044: What is mocking?**
(Refresher) Faking dependencies.

**Q2045: What is TDD?**
(Refresher) Test Driven Development.

**Q2046: What is pdb?**
(Refresher) Python Debugger.

**Q2047: How to debug memory usage?**
`tracemalloc`.

**Q2048: How to profile code?**
`cProfile`.

**Q2049: What is static analysis?**
Checking code without running (MyPy, Pylint).

**Q2050: What is integration testing?**
(Refresher) Testing combined units.

---

## 🔹 206. Web & Async II (Questions 2051-2060)

**Q2051: What is Django?**
(Refresher) Full stack web framework.

**Q2052: What is Flask?**
(Refresher) Micro framework.

**Q2053: What is FastAPI?**
(Refresher) Modern Async framework.

**Q2054: What is an ORM?**
(Refresher) Object Relational Mapper.

**Q2055: What is WSGI?**
(Refresher) Web Server Gateway Interface.

**Q2056: What is ASGI?**
(Refresher) Async Server Gateway Interface.

**Q2057: What is REST?**
(Refresher) Architectural style.

**Q2058: What is GraphQL?**
(Refresher) Query language.

**Q2059: What is a Coroutine?**
(Refresher) Pausable function.

**Q2060: What is Event Loop?**
(Refresher) Async scheduler.

---

## 🔹 207. Data Science & ML (Questions 2061-2070)

**Q2061: What is Pandas?**
Data Analysis library.

**Q2062: What is NumPy?**
Numerical Computing library.

**Q2063: What is Scikit-Learn?**
Machine Learning library.

**Q2064: What is Matplotlib?**
Plotting library.

**Q2065: What is DataFrame?**
(Refresher) Table data structure.

**Q2066: What is Series?**
(Refresher) Column data structure.

**Q2067: What is vectorization?**
(Refresher) Batch operation on arrays.

**Q2068: What is broadcasting?**
(Refresher) Arithmetic on different shapes.

**Q2069: What is TensorFlow?**
Deep Learning framework by Google.

**Q2070: What is PyTorch?**
Deep Learning framework by Meta.

---

## 🔹 208. Libraries & Tools (Questions 2071-2080)

**Q2071: What is requests?**
HTTP library.

**Q2072: What is beautifulsoup?**
HTML parser.

**Q2073: What is selenium?**
Browser automation.

**Q2074: What is pillow?**
Image processing.

**Q2075: What is sqlalchemy?**
Database toolkit.

**Q2076: What is celery?**
Task queue.

**Q2077: What is click?**
CLI creation kit.

**Q2078: What is virtualenv?**
(Refresher) Environment isolation.

**Q2079: What is pip?**
(Refresher) Package manager.

**Q2080: What is docker?**
Containerization platform.

---

## 🔹 209. System & Files (Questions 2081-2090)

**Q2081: What is os module?**
OS interaction.

**Q2082: What is sys module?**
Interpreter interaction.

**Q2083: What is subprocess?**
Running shell commands.

**Q2084: What is shutil?**
File operations.

**Q2085: How to read file?**
`open(f).read()`.

**Q2086: How to write file?**
`open(f, 'w').write()`.

**Q2087: What is json module?**
JSON parsing/generation.

**Q2088: What is csv module?**
CSV parsing.

**Q2089: What is threading module?**
Threads.

**Q2090: What is multiprocessing module?**
Processes.

---

## 🔹 210. Python History & Meta (Questions 2091-2100)

**Q2091: Who is BDFL?**
Guido van Rossum (Benevolent Dictator For Life - retired).

**Q2092: When was Python released?**
1991.

**Q2093: Why the name Python?**
Monty Python's Flying Circus.

**Q2094: What is Python 2 end of life?**
Jan 1, 2020.

**Q2095: What is PEP?**
Python Enhancement Proposal.

**Q2096: What is PyPI?**
Python Package Index.

**Q2097: What is PSF?**
Python Software Foundation.

**Q2098: Is Python open source?**
Yes.

**Q2099: What license does Python use?**
PSFL (Python Software Foundation License).

**Q2100: What is the current stable version?**
(Continually updating) 3.12+ (as of late 2024).

---

## 🔹 211. Specific Python Versions & History (Questions 2101-2110)

**Q2101: What was new in Python 3.5?**
Async/Await syntax, Type Hints.

**Q2102: What was new in Python 3.6?**
f-strings, underscores in numbers, async generators.

**Q2103: What was new in Python 3.7?**
Data Classes, `breakpoint()`, contextvars.

**Q2104: What was new in Python 3.8?**
Walrus operator `:=`, Positional-only params `/`.

**Q2105: What was new in Python 3.9?**
Dict merge `|`, String prefix/suffix removal methods.

**Q2106: What was new in Python 3.10?**
Structural Pattern Matching (`match-case`), `int | str` union types.

**Q2107: What was new in Python 3.11?**
Performance boost (Specializing Adaptive Interpreter), `TaskGroup`.

**Q2108: What is the Python Governance Model?**
Steering Council (elected members) manages decisions after BDFL retirement.

**Q2109: What is the Zen of Python's author?**
Tim Peters.

**Q2110: How many principles are in Zen of Python?**
19 written, 1 unwritten.

---

## 🔹 212. Comparisons (Questions 2111-2120)

**Q2111: Python vs Java?**
Python: Interpreted, dynamic, concise. Java: Compiled, static, verbose.

**Q2112: Python vs C++?**
Python: Slower but safer/easier. C++: Performance-critical systems.

**Q2113: Python vs JavaScript?**
Python: Backend/AI/Data. JS: Frontend/Fullstack (Node).

**Q2114: Python vs Go?**
Go: Concurrency-native, compiled static binary. Python: Ease of use, ecosystem.

**Q2115: Python vs R?**
Python: General purpose + ML. R: Statistics specialized.

**Q2116: Static vs Dynamic Typing?**
Static checks types at compile time. Dynamic checks at runtime.

**Q2117: Strong vs Weak Typing?**
Strong doesn't allow implicit type coercion (Python). Weak does (JS).

**Q2118: Compiled vs Interpreted?**
Executed by CPU directly vs Executed by Interpreter.

**Q2119: What is JIT?**
Just-In-Time compilation (compiling hot paths to machine code at runtime).

**Q2120: What is AOT?**
Ahead-Of-Time compilation.

---

## 🔹 213. Final Mixed Bag (Questions 2121-2130)

**Q2121: How to get list of running threads?**
`threading.enumerate()`.

**Q2122: How to check main thread?**
`threading.main_thread()`.

**Q2123: What is stack overflow in Python?**
Recursion limit exceeded (`RecursionError`).

**Q2124: How to increase recursion limit?**
`sys.setrecursionlimit(N)`.

**Q2125: What happens if you import a package that raises an error?**
ImportError. Module is not loaded.

**Q2126: How to handle circular buffer?**
`collections.deque(maxlen=N)`.

**Q2127: What is a bloom filter?**
Probabilistic data structure for set membership. (Not built-in).

**Q2128: How to implement a graph?**
Adjacency list (Dict of Lists).

**Q2129: What is topological sort?**
Linear ordering of vertices in DAG.

**Q2130: What is Dijkstra's algorithm?**
Shortest path in graph with non-negative weights.

---

## 🔹 214. Miscellaneous & Trivia (Questions 2131-2140)

**Q2131: What is antigravity module?**
`import antigravity` opens an XKCD comic. Easter egg.

**Q2132: What is this module?**
`import this` displays Zen of Python.

**Q2133: What is braces import?**
`from __future__ import braces` raises SyntaxError "not a chance".

**Q2134: How to open a file with default app?**
`os.startfile()` (Windows).

**Q2135: What is the fastest Python implementation?**
Generally PyPy or CPython 3.11+.

**Q2136: How to create a Windows service with Python?**
`pywin32` library.

**Q2137: How to access Windows Registry?**
`winreg` module.

**Q2138: How to play a beep sound?**
`print('\a')` or `winsound.Beep()`.

**Q2139: How to hide console window?**
Save as `.pyw`.

**Q2140: How to package resources inside executable?**
PyInstaller's `--add-data`.

---

## 🔹 215. Deep Internals (Questions 2141-2150)

**Q2141: What is a stack frame?**
Data structure containing state of a function call (locals, return address).

**Q2142: How to inspect stack frames?**
`sys._getframe()` or `inspect`.

**Q2143: What is reflection?**
Ability of a program to examine and modify its own structure/behavior (introspection).

**Q2144: How to change class of an object at runtime?**
Assign to `__class__` (Dangerous).

**Q2145: What is method wrapper?**
Bound method object.

**Q2146: How does Python handle method resolution?**
Search in instance dict -> Search in Class dict -> Search in Base Classes (MRO).

**Q2147: What is __dict__ proxy?**
`mappingproxy` object returned by `Class.__dict__` to prevent direct modification.

**Q2148: What is code object?**
Compiled bytecode (`func.__code__`).

**Q2149: How to dynamically execute bytecode?**
`exec(code_obj)`.

**Q2150: What is dis module?**
Disassembler for Python bytecode.

---

## 🔹 216. Security & Ethics (Questions 2151-2160)

**Q2151: What is obfuscation?**
Making code hard to read to protect IP.

**Q2152: Is Python secure?**
Depends on code quality. Language itself is memory safe (no buffer overflows usually) but vulnerable to logic bugs/injections.

**Q2153: What is pickle bomb?**
Malicious pickle data that executes code upon loading.

**Q2154: How to scan dependencies for vulnerabilities?**
`safety check` or `pip-audit`.

**Q2155: What is Typosquatting?**
Uploading packages with similar names to popular ones (e.g., `reqests`) to trick users.

**Q2156: How to verify package integrity?**
Check hashes in lockfiles.

**Q2157: What is Zip Slip?**
Directory traversal vulnerability when extracting zip files.

**Q2158: What is exponential regex?**
Regex that takes exponential time on certain inputs (ReDoS).

**Q2159: How to sanitize HTML?**
`bleach` library.

**Q2160: What is bandit?**
Security linter for Python.

---

## 🔹 217. Career & Roles (Questions 2161-2170)

**Q2161: What is a Python Developer role?**
Backend, Scripting, Automation.

**Q2162: What is a Data Scientist role?**
Insights from data using Python/Pandas/ML.

**Q2163: What is a Data Engineer role?**
Pipelines, ETL, Big Data.

**Q2164: What is DevOps with Python?**
Infrastructure automation (Ansible/Fabric), CI/CD.

**Q2165: What is Full Stack Python?**
Django/Flask backend + JS Frontend (or HTMX).

**Q2166: What is QA Automation Engineer?**
Writing tests with Selenium/Pytest.

**Q2167: What is salary of Python dev?**
Varies. Usually high demand.

**Q2168: How to contribute to Python?**
bugs.python.org, GitHub PRs to CPython.

**Q2169: What is a core developer?**
Someone with commit rights to CPython.

**Q2170: What is Python Software Foundation Membership?**
Recognition for community contribution.

---

## 🔹 218. Community & Learning (Questions 2171-2180)

**Q2171: What is PyCon?**
Largest annual gathering for the Python community.

**Q2172: What is a Sprint?**
Coding session (usually at conferences) to contribute to open source.

**Q2173: What is Real Python?**
Popular learning resource.

**Q2174: What is Talk Python To Me?**
Popular podcast.

**Q2175: What is Stack Overflow?**
Q&A site (Invaluable resource).

**Q2176: How to read PEPs?**
python.org/dev/peps.

**Q2177: What is the Python Weekly?**
Newsletter.

**Q2178: How to find local Python meetups?**
Meetup.com or python.org user groups.

**Q2179: What is discord.py?**
Library for Discord bots (now active forks exist).

**Q2180: What is Kaggle?**
Data Science competitions platform.

---

## 🔹 219. Future & Trends (Questions 2181-2190)

**Q2181: What is Python in the browser?**
PyScript / Pyodide (WASM).

**Q2182: What is Mojo?**
Superset of Python designed for AI hardware speed.

**Q2183: What is No-GIL Python?**
PEP 703 (Making the GIL optional in future CPython).

**Q2184: What is AI interaction?**
LLM integration (LangChain) via Python.

**Q2185: What is functional python trend?**
Increasing features (Pattern matching, types).

**Q2186: What is MicroPython?**
Python for microcontrollers.

**Q2187: What is CircuitPython?**
Adafruit's fork of MicroPython.

**Q2188: What is Embedded Python?**
Running Python inside C++ apps.

**Q2189: What is Rust + Python?**
Using PyO3 to write extensions in Rust for speed.

**Q2190: What is Polars?**
Fast DataFrame library (Rust-based) rivaling Pandas.

---

## 🔹 220. The End (Questions 2191-2200)

**Q2191: How to say goodbye in Python?**
`print("Goodbye World")`.

**Q2192: What does quit() do?**
Exits interpreter.

**Q2193: How to clean up resources?**
`atexit` module registers cleanup functions.

**Q2194: What is memory leak debugging?**
`objgraph` library.

**Q2195: How to check active children threads?**
`multiprocessing.active_children()`.

**Q2196: What is a zombie process?**
Finished process but parent hasn't read exit status.

**Q2197: How to orphan a process?**
Parent dies, child adopted by init.

**Q2198: What is the max integer in Python?**
Memory limited (no hard limit).

**Q2199: What is a nanosecond resolution?**
`time.time_ns()`.

**Q2200: What is the ultimate answer?**
`42` (Douglas Adams).

---

## 🔹 221. Bonus & Conclusion (Questions 2201-2212)

**Q2201: Can you use semicolons in Python?**
Yes, but unpythonic.

**Q2202: What is the Walrus?**
`:=`.

**Q2203: What is the Ellipsis?**
`...`.

**Q2204: What is the Zen?**
`import this`.

**Q2205: What is Antigravity?**
`import antigravity`.

**Q2206: What is Hello World?**
`print("Hello World")`.

**Q2207: How to swap variables?**
`a, b = b, a`.

**Q2208: How to reverse string?**
`s[::-1]`.

**Q2209: How to get help?**
`help()`.

**Q2210: How to exit?**
`exit()`.

**Q2211: Is Python great?**
`True`.

**Q2212: Are we done?**
`print("Yes, all 2212 questions answered.")`.
