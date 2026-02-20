import os

content = """
## From Python Internals & Advanced Concepts

# 🟢 **81–100: Memory, Internals, and Execution**

### 81. How is memory managed in Python?
"Memory is managed entirely by the Python Memory Manager.

It allocates memory dynamically on a private heap. To free it, Python uses **Reference Counting**—every time an object is referenced, a counter goes up; when the reference is removed, it goes down. When it hits zero, the memory is instantly reclaimed.

As a fallback for objects referencing each other (circular references), Python periodically runs a **Generational Garbage Collector** to sweep them up."

#### Indepth
The memory manager allocates memory in chunks called `Arenas` (256KB), which are broken down into `Pools` (4KB), which are broken down into `Blocks`. This tiered allocation system is designed specifically to optimize the creation of small, short-lived objects (ints, small strings) which Python generates by the millions.

---

### 82. What is Python's GIL (Global Interpreter Lock)?
"The GIL is a mutex (lock) in the standard CPython implementation.

It strictly allows only **one thread** to execute Python bytecode at any given time. Even if my server has 64 CPU cores, a multi-threaded Python program will only ever utilize 1 mathematical core for Python execution.

I rely on multi-threading purely for **I/O bound** tasks (like making 100 simultaneous network requests), where threads yield the GIL while waiting. For CPU-bound math, I *must* use `multiprocessing` to spawn totally separate processes, each with its own memory and its own GIL."

#### Indepth
The GIL exists because CPython's fundamentally pervasive memory management (Reference Counting) is not natively thread-safe. Historically, attempting to implement fine-grained locking on every object slowed down single-threaded Python scripts unacceptably, so the single "Global" lock was implemented instead.

---

### 83. What is a shallow copy and a deep copy?
"A **shallow copy** (`copy.copy(list)`) creates a new top-level object, but inserts *references* into it pointing to the original child objects. If the list contains a dictionary, modifying the dictionary in the copy will modify the original.

A **deep copy** (`copy.deepcopy(list)`) recursively walks the entire data structure, literally cloning every single nested object found. The new list and the old list share absolutely zero physical memory.

I only use deepcopy when dealing with highly nested state trees that I need to intentionally mutate without side-effecting the master state."

#### Indepth
`deepcopy` is incredibly slow and resource-intensive because it has to track already-copied items in a "memo" dictionary to prevent infinite recursive loops when encountering circular references (like an object containing a reference to itself).

---

### 84. What is the difference between `is` and `==`?
"`==` compares the **Value Equality** (Do these two objects have the same data?).
`is` compares the **Identity Equality** (Do these two variables point to the exact same memory address?).

`[1, 2] == [1, 2]` is True. `[1, 2] is [1, 2]` is False because they are distinct list objects.

I overwhelmingly use `==` for business logic, and strictly reserve `is` for singleton checks: `if var is None:` or `if var is True:`."

#### Indepth
To optimize performance, Python statically "interns" small integers (-5 to 256) and some small strings. This means `a = 100` and `b = 100` will actually evaluate `a is b` as True, because they point to the exact same cached integer object in the CPython runtime.

---

### 85. What are decorators?
"A decorator is a function that takes another function, extends its behavior without explicitly modifying it, and returns the modified wrapper function.

I use them extensively to pull cross-cutting concerns (like logging, authentication, or retry logic) out of my core business functions. 

Typing `@require_admin` above my `delete_user()` function is clean, expressive, and keeps my domain logic entirely separated from my routing/security logic."

#### Indepth
When writing a decorator, it's critical to use `@functools.wraps(func)` on the inner wrapper function. Without this, the heavily introspection-reliant tools in Python lose sequence of the original function's `__name__` and `__doc__` string, because the wrapper function technically replaced it.

---

### 86. How and why to use `@staticmethod`?
"I use the `@staticmethod` decorator to place a standard utility function strictly inside a class namespace purely for organizational code clarity.

It takes neither `self` (the instance) nor `cls` (the class) as its first argument. 

If I have a `DateFormatter` class, a function like `@staticmethod def is_leap_year(year):` doesn't need to read any instance state. It’s conceptually related to the class, so putting it there makes the API cleaner instead of leaving it drifting as a module global."

#### Indepth
Because static methods have no access to the class or instance state, they are incredibly easy to write comprehensive unit tests for. They are pure, deterministic functions (given X input, ALWAYS produce Y output).

---

### 87. What is `@classmethod`?
"A `@classmethod` receives the class itself (`cls`) as the implicit first argument, rather than the instance (`self`).

I use it almost universally for providing **Alternative Constructors** (the Factory Pattern). 

If an `Employee` defaults to initializing from a username string, I write `@classmethod def from_json(cls, json_payload):` to parse the dictionary and return `cls(parsed_data)`. If someone subclasses `Manager(Employee)` and calls `Manager.from_json`, the `cls` automatically creates a `Manager` instead of an `Employee`."

#### Indepth
If a class method tries to call `cls()` but the `__init__` signature in a child class changes, it will crash dynamically. This highlights why Python's dynamic typing requires robust unit testing for inheritance paths.

---

### 88. What is Python bytecode?
"It is the platform-independent intermediate language that the CPython compiler generates.

When I run `script.py`, the interpreter compiles the human-readable text into a `.pyc` file containing simplified operations (opcodes like `LOAD_FAST`, `BINARY_ADD`).

The Python Virtual Machine (PVM) is a massive C program that strictly iterates over these opcodes one-by-one and executes them. This is the exact reason Python is slower than compiled C; the PVM is constantly interpreting opcodes dynamically."

#### Indepth
You can completely expose and read this underlying bytecode matrix for any function using the standard `dis` module. Running `dis.dis(my_function)` prints the precise assembly-like instructions Python runs, which is invaluable for hyper-optimizing inner loops.

---

### 89. What happens when you execute a Python script?
"1. **Parsing:** The compiler reads the `.py` source text and converts it into an Abstract Syntax Tree (AST).
2. **Compilation:** The AST is compiled down into raw Python bytecode (caching it as `.pyc` in `__pycache__` if it's an imported module).
3. **Execution:** The Python Virtual Machine (PVM) spins up. It takes the bytecode and executes it opcode-by-opcode using a giant C `switch` statement."

#### Indepth
Python execution is strictly sequential from the top of the file down. This means `def function()` literally executes a definition command at runtime, creating a function object in memory and binding it to that variable name.

---

### 90. What is a namespace?
"A namespace is practically a mapping (usually implemented as a dictionary under the hood) tying variable, function, or class names to their actual underlying physical objects.

Python strictly isolates them. If I `import math`, I use `math.pi`. The `pi` lives safely inside the `math` namespace. My own variable `pi = 3` lives safely inside my Local/Global namespace. 

This isolation completely stops function and variable names from clashing across huge codebases."

#### Indepth
You can directly inspect namespaces! The `locals()` built-in returns a dictionary representing the current local namespace, and `globals()` returns the module-level namespace. Dynamically mutating `globals()` works, but mutating `locals()` is highly undefined and usually fails gracefully.

---

### 91. What is the difference between global and nonlocal?
"The `global` keyword forces a variable lookup/assignment straight to the top module level.

The `nonlocal` keyword forces the lookup exactly *one step outwards* to the immediately enclosing function scope (a Closure).

I use `nonlocal` mostly in decorators or nested factory functions when an inner wrapper function needs to increment a counter or modify state maintained by the outer wrapper factory, without polluting the global scope."

#### Indepth
Before `nonlocal` was introduced in Python 3, developers had to use a mutability hack to maintain state in closures: wrapping the state in a list `counter = [0]` in the outer function, and modifying `counter[0] += 1` in the inner function, because mutating an object bypasses scope reassignment rules.

---

### 92. What does `for else` mean in Python?
"It's a very unique construct. The `else` block executes *only if* the `for` loop completed naturally without hitting a `break` statement.

I use this predominantly in Search algorithms. 

`for user in users:`
`    if user.name == 'Admin':`
`        print('Found!')`
`        break`
`else:`
`    print('Admin distinctly not found.')`

This brilliantly removes the need to maintain an external boolean `found = False` flag purely to trigger "Not found" logic."

#### Indepth
The keyword `else` here is famously counter-intuitive conceptually; prominent core developers have suggested that historically naming it `nobreak` would have saved millions of hours of confusion.

---

### 93. How do you raise a specific custom exception?
"To build a custom exception, I inherit straight from the base `Exception` class.

`class InvalidTokenError(Exception):`
`    pass`

Inside my auth logic, I run `raise InvalidTokenError("Token expired")`. 

This makes my application incredibly robust because a high-level API wrapper can now easily `except InvalidTokenError:` specifically, separating domain-specific logical failures from raw `AttributeError` parsing crashes."

#### Indepth
If you are catching an error, applying logic, and then throwing a *different* error, you should use Python 3's advanced `raise new_err from old_err` syntax. This explicitly chains the stack traces together, preventing the nightmare of "I overwrote the original failure context."

---

### 94. What is duck typing?
"'If it walks like a duck and quacks like a duck, I assume it’s a duck.'

In Python, I don’t care what explicit class an object inherits from. If my function just calls `animal.speak()`, I'll happily accept a `Dog` object, a `Cat` object, or an `Alien` object, as long as they all uniquely implement `.speak()`. 

This dynamic philosophy allows massive decoupling and makes mocking for unit tests trivial."

#### Indepth
While pure duck typing is the Python ideal, modern enormous codebases rely heavily on Type Hinting (`def process(animal: Animal):`) to allow static analysis tools (like `mypy`) to catch interface errors *before* runtime execution hits the `.speak()` line and violently crashes.

---

### 95. What is the difference between mutable and immutable types?
"**Mutable** objects (Lists, Dicts, Sets) can be physically changed in memory after creation. Appending to a list keeps the exact identical `id()` memory address.

**Immutable** objects (Strings, Tuples, Integers) can completely *never* change. If I do `string += "!"`, Python destroys the old string and secretly allocates brand new memory for the concatenated result.

I specifically use immutable tuples for dictionary keys because their hashes are mathematically guaranteed to remain identical forever."

#### Indepth
A notorious trap: A tuple is immutable in structure. `t = (1, [2, 3])` means `t` must always have an integer at index 0 and a list at index 1. But the list *itself* is deeply mutable! `t[1].append(4)` is perfectly valid and completely mutates the tuple's deeply enclosed data.

---

### 96. What is a magic method / dunder method?
"Any method surrounded by double underscores (`__like_this__`) is a magic method. They are Python's native way of doing Operator Overloading.

I use `__str__` to dictate exactly how my object uniquely renders when passed to `print()`. I use `__eq__` to intercept the `==` operator and dictate my own exact mathematical logic for judging if two objects are equivalent.

They embed native Python syntax natively into my completely custom classes."

#### Indepth
If you write `__repr__` instead of `__str__`, it acts as the raw fallback representation, heavily used in interactive debugging. A solid rule: `__str__` is for end-users formatting; `__repr__` is strictly for developers debugging code.

---

### 97. What is the use of `zip()`?
"The `zip()` function takes two or more iterables and perfectly interleaves them together side-by-side into tuples.

If `names = ['Alice', 'Bob']` and `grades = [90, 85]`, doing `zip(names, grades)` returns `[('Alice', 90), ('Bob', 85)]`.

I use it heavily to map two parallel lists into a dictionary instantaneously: `dict(zip(names, grades))` is one of my favorite elegant, highly optimized Python one-liners."

#### Indepth
`zip()` inherently stops lazily when the absolute *shortest* sequence runs out. Data actively loss can occur! If the sequences must be definitively exhaustive, `itertools.zip_longest()` must be fundamentally utilized instead.

---

### 98. How do you sort a list of dictionaries by a specific key?
"I universally use the `sort()` method or `sorted()` function and inject a custom `key` argument using a quick `lambda`.

`sorted(users, key=lambda user: user['age'])`

This iterates over my dictionary objects, secretly extracts just the `age` integer for every user, violently sorts them solely based on those integers, and returns the strictly ordered original dictionary objects."

#### Indepth
A faster, more C-optimized alternative to typing raw `lambda x: x['age']` is utilizing the `operator.itemgetter('age')` method. It is physically faster because it entirely avoids spinning up a Python bytecode function scope for every single iterated list item.

---

### 99. How to find the most frequent element in a list?
"I import `Counter` from the highly optimized `collections` module. 

`from collections import Counter`
`most_common = Counter(my_list).most_common(1)`

Writing a raw `for` loop and storing integer frequencies in a dictionary manually is slow and error-prone. The `Counter` heavily leverages highly tuned native C dictionary implementation paths to crunch frequencies exponentially faster."

#### Indepth
The `.most_common(n)` method secretly runs a hyper-optimized `heapq` (heap queue) priority algorithm under the hood if `n` is vastly smaller than the total unique items in the list, performing substantially faster than a full array sort.

---

### 100. What is a Python virtual environment entirely used for?
"I use virtual environments (`venv`) strictly to create tightly isolated, project-specific installations of Python and external libraries.

If Project A violently requires `Django==3.0` and Project B violently requires `Django==4.0`, installing globally will permanently break one of them. The `venv` abstracts a cleanly separate library namespace for each distinct project folder.

I religiously never run `pip install` globally natively; every single project strictly gets a dedicated `requirements.txt` and a tightly localized `.venv` folder."

#### Indepth
The sheer nightmare of dependency conflicts in Python—often notoriously called "Dependency Hell"—is the literal primary reason `Docker` explicitly exploded in popularity. Virtual environments solve Python library isolation locally, while Docker fundamentally solves OS-level isolation globally.
"""

file_path = r'g:\My Drive\All Documents\java-golang-interview-questions\long_questions\Python\Theory\01_Basics.md'

with open(file_path, 'a', encoding='utf-8') as f:
    f.write(content)
