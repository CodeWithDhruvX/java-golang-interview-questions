import os

content = """
## From Object-Oriented Programming & Advanced Features

# 🟢 **61–80: OOP, Iterators, and Decorators**

### 61. What is Object-Oriented Programming (OOP) in Python?
"OOP is a programming paradigm where we organize code around data (objects) and behavior (methods) rather than just writing sequential functions.

Python supports OOP natively. I use classes to bundle related state and logic together. If I'm building a game with 10 enemies, I write one `Enemy` class, and then create 10 distinct `Enemy` objects in memory. It makes code highly reusable and much easier to reason about as systems grow."

#### Indepth
Unlike Java, Python does not enforce strict access modifiers. "Private" variables (`_var`) are just conventions, and "mangled" variables (`__var`) are designed to prevent accidental shadowing in subclasses, not to enforce cryptographic security. The community relies entirely on developer discipline.

---

### 62. What is the `__init__` method?
"The `__init__` method is the constructor in Python. It's the very first function that gets called automatically when I instantiate an object from a class.

I use it exclusively to set up the initial state of my object. 

`def __init__(self, username, role):`
`    self.username = username`
`    self.role = role`

It guarantees that my object has all the required attributes *before* any other methods in the class attempt to interact with it."

#### Indepth
`__init__` is an *initializer*, not technically the instantiator. The actual underlying memory allocation is performed by the `__new__` magic method, which returns the blank object instance, and *then* Python automatically passes that blank instance into your `__init__` method.

---

### 63. What does `self` mean in Python?
"`self` refers explicitly to the specific object instance that is currently executing the method.

When I create `user A` and call `user_A.login()`, Python silently translates that to `User.login(user_A)`. The `self` parameter catches that object instance. 

I must explicitly type `self.username` everywhere inside the class to access instance variables. Unlike C++ or Java (`this`), it's not a reserved keyword, but it is a rigid community convention that almost no one ever breaks."

#### Indepth
You can legitimately name it anything you want (like `me` or `this`), but doing so violates PEP 8 and heavily confuses other developers and IDE linters. `self` is explicitly required in Python to enforce the philosophy that "explicit is better than implicit" when dealing with state mutations.

---

### 64. What is inheritance?
"Inheritance allows a new class (the child) to automatically inherit attributes and methods from an existing class (the parent).

`class Admin(User):` means `Admin` instantly gains every method `User` has. 

I use inheritance to avoid duplicating code. If both `Admin` and `Guest` need a `login()` method, I write it once in the `User` class. If `Admin` needs a specialized `login()`, I simply override the method in the child class."

#### Indepth
Python supports Multiple Inheritance (`class Child(Mother, Father):`), which is notoriously complex due to the "Diamond Problem". Python resolves method conflicts mathematically using the C3 Linearization algorithm to determine the Method Resolution Order (MRO). The MRO can be inspected via `Child.__mro__`.

---

### 65. What is the `super()` function?
"The `super()` function returns a temporary object of the parent class, allowing me to call its methods.

I use it almost constantly inside an overridden `__init__` method child class. 

`def __init__(self, name, admin_level):`
`    super().__init__(name)`  # Calls User.__init__
`    self.admin_level = admin_level`

This ensures that the parent class fully initializes its core state before my child class adds its specialized setup."

#### Indepth
In Python 2, `super()` required explicitly passing the current class and instance: `super(Admin, self).__init__(name)`. In Python 3, calling `super()` with no arguments automatically binds to the correct class and instance implicitly, drastically cleaning up the syntax.

---

### 66. What is polymorphism?
"Polymorphism allows different child classes to be treated exactly like their parent class, while still responding differently.

I might have a `Notification` interface, with `Email` and `SMS` subclasses. Both have a `.send(message)` method.

My core logic just loops through a list: `for notif in list: notif.send('Hi')`. It doesn't care if it's an Email or an SMS. It simply asks the object to execute its specialized version of `.send()`, keeping my core code extremely decoupled."

#### Indepth
Because Python is dynamically typed ("Duck Typing"), polymorphism doesn't actually strictly require inheritance at all! If a completely unrelated class implements a `.send(message)` method, Python will happily execute it inside that loop without any compilation complaints.

---

### 67. What are Class Methods and Static Methods?
"A normal instance method takes `self` and manipulates a specific object instance.

A **Class Method** takes `cls` (the class itself) rather than an instance. I use it mostly to build alternative constructors (Factory Pattern), like `@classmethod def from_json(cls, data):`.

A **Static Method** takes neither `self` nor `cls`. It's just a normal utility function completely disconnected from class state, but grouped inside the class namespace for organizational logic: `@staticmethod def is_valid_email(address):`."

#### Indepth
The `@classmethod` dynamically receives whatever child class it was actually called on. If `Admin` inherits `User.from_json()`, calling `Admin.from_json()` passes the `Admin` class into `cls`, ensuring an `Admin` object is returned, not a `User`.

---

### 68. What are magic (dunder) methods?
"Magic methods start and end with double underscores (e.g., `__init__`, `__str__`, `__add__`). We call them 'dunder' methods.

They allow me to define how my custom objects interact with Python's built-in syntax. 

If I define `__str__(self)`, my object returns a clean string when I call `print(obj)`. If I define `__len__(self)`, I can use Python's built-in `len(obj)` function on my completely custom data structure."

#### Indepth
Python operators are entirely syntactic sugar for dunder methods. Writing `objA + objB` literally executes `objA.__add__(objB)`. Overriding these methods is the fundamental way Operator Overloading is implemented in Python.

---

### 69. What is a property decorator (`@property`)?
"The `@property` decorator is brilliant for Encapsulation. It allows me to define a method but access it like a direct attribute.

If I have a `Circle` class with a `radius`, I can define `@property def area(self): return 3.14 * (self.radius ** 2)`.

Now, consumers just type `circle.area`, not `circle.area()`. To them, it looks like a variable. But underneath, a dynamically calculated function is running, guaranteeing the area is always precisely up to date if the radius changes."

#### Indepth
Properties drastically reduce the need for explicit "Getter" and "Setter" methods like Java's `getArea()`. By pairing `@property` with an `@attribute.setter` decorator, you can transparently add validation logic to attribute assignment: `circle.radius = -5` could trigger a ValueError smoothly behind the scenes.

---

### 70. What is an iterator?
"An iterator is an object representing a stream of data. It respects the Iterator Protocol by implementing exactly two methods: `__iter__()` and `__next__()`.

When I use a `for` loop, Python secretly calls `__iter__()` to get the iterator, and then repeatedly calls `next()` until the iterator throws a `StopIteration` exception.

Iterators are highly memory efficient because they only ever hold one item in memory at a time, generating the next item lazily upon request."

#### Indepth
Every iterator is technically an iterable, but not every iterable is an iterator. A `list` is an iterable (it implements `__iter__` returning a ListIterator), but it does not implement `__next__`. This split design allows you to iterate over the same list multiple times independently.

---

### 71. What is a generator?
"A generator is the absolute easiest way to create an iterator in Python. It's simply a function that contains a `yield` statement instead of a `return`.

Instead of returning an entire massive list on line 1, `yield` returns a single value, *pauses* the function exactly where it is, saves its local state/variables, and yields control back to the caller. 

I use generators heavily for reading endless data streams, paginating massive database tables, or parsing log files line-by-line where loading the whole file into RAM would crash the server."

#### Indepth
Under the hood, generators are implemented as state machines. Python saves the execution frame, the instruction pointer, and all local variables into heap-allocated state structure for that generator object, enabling it to resume precisely on the next `next()` call.

---

### 72. What is the difference between `yield` and `return`?
"`return` terminates a function completely and returns a final value. All local variables in the function are instantly destroyed.

`yield` returns a value but *suspends* the function's execution state. When the generator is called again, it resumes directly after the `yield` statement with all its local variables perfectly preserved.

I use `return` for calculating an answer, and `yield` for generating an ongoing sequence."

#### Indepth
Since Python 3.3, you actually can use `return` *inside* a generator! However, it doesn't return a value to the `for` loop. Instead, it aggressively raises the `StopIteration` exception early and attaches the returned value to the exception object.

---

### 73. What are generator expressions?
"A generator expression is the lazy, memory-efficient sibling of list comprehension. The syntax is identical except we use parentheses `()` instead of brackets `[]`.

`huge_gen = (x ** 2 for x in range(10_000_000))`

If I used a list comprehension here, Python would allocate memory for 10 million integers instantly, freezing my laptop. The generator expression takes practically zero bytes because it merely creates a blueprint to generate those numbers only when asked."

#### Indepth
If you pass a generator expression as the solitary argument to a function, like `sum(x**2 for x in range(100))`, you don't even need double parentheses `sum((...))`, Python parses the single set as the generator, creating incredibly readable, highly optimized aggregation chains.

---

### 74. What is a decorator?
"A decorator is a powerful structural pattern. It's essentially a function that takes *another function* as input, wraps some extra behavior around it, and returns a new function.

I use them everywhere in web frameworks. Writing `@app.route('/login')` or `@require_auth` completely abstracts routing and security logic away from the core business function underneath. 

It keeps my business functions pure and drastically reduces boilerplate code repetition (DRY concept)."

#### Indepth
Decorators run exactly once, when the module loads, replacing the target function name with the wrapper function. Because the original function name and docstring are lost to the wrapper, you must use `functools.wraps(func)` inside the decorator to copy that crucial metadata back onto the wrapper.

---

### 75. How do you pass arguments to a decorator?
"Because decorators return functions, building a decorator that itself accepts arguments requires nesting *three* layers deep!

1. The outer function takes the arguments (e.g., `@retry(max_attempts=3)`).
2. The middle function is the actual decorator taking the target `func`.
3. The innermost function is the `wrapper` that accepts `(*args, **kwargs)` and uses the outer parameter (`max_attempts`).

It looks intimidating, but it's an extremely elegant metaprogramming tool once you wrap your head around closures."

#### Indepth
A more modern, readable approach to complex decorators is bypassing nested functions entirely and implementing the decorator using a Class. You define the `__init__` method to catch the `max_attempts` arguments, and the `__call__` dunder method acts as the actual wrapper intercepting the target function execution.

---

### 76. What is the `map()` function?
"The `map(function, iterable)` built-in applies a given function across every single item in an iterable.

If I need to convert a list of string numbers `['1', '2']` into real integers, I use `map(int, strings)`. 

In Python 3, `map` returns a lazy iterator, not a list. To see everything immediately or serialize it, I must explicitly cast it: `list(map(int, strings))`."

#### Indepth
Historically, `map` combined with `lambda` was huge. Modern Python strongly prefers List Comprehensions: `[int(x) for x in strings]`. Comprehensions execute slightly faster than map+lambda (avoiding the function call overhead) and are generally considered more readable idiomatically.

---

### 77. What is the `filter()` function?
"The `filter(function, iterable)` built-in selectively keeps elements based on a condition function.

If I write `filter(lambda x: x > 5, numbers)`, it returns an iterator containing solely the numbers greater than 5. If the condition function returns `True`, the element survives. 

Like `map`, it evaluates lazily, which is fantastic for chaining together long functional pipelines over massive datasets."

#### Indepth
If you pass `None` as the first argument instead of a custom boolean function, `filter(None, iterable)` defaults to using Python's "Truthiness" rules. It explicitly strips out all falsy values (`0`, `False`, `""`, `None`) from the iterable in a highly optimized C-loop. 

---

### 78. What does `isinstance()` do?
"The `isinstance(object, class)` built-in checks if an object is derived from a specific class (or a tuple of possible classes).

`isinstance(5, int)` is True.

I use it heavily for highly dynamic code or input validation, especially if a function accepts the `any` generic type. If it's a string, do X; if it's a list, do Y."

#### Indepth
Always use `isinstance()` rather than `type(obj) == class`. `type()` acts rigidly and strictly checks for exact class matching. `isinstance()` supports inheritance and polymorphism cleanly (e.g., `isinstance(admin, User)` is True, whereas `type(admin) == User` is False).

---

### 79. What does the `id()` function do?
"The `id(obj)` built-in returns a unique integer guaranteeing the object's identity for its entire lifetime. 

In CPython, this number literally maps entirely directly to the object's physical memory address in RAM.

I exclusively use it for incredibly deep debugging—for instance, verifying if my function accidentally created a new dictionary copy or successfully modified the exact original dictionary passed to it."

#### Indepth
The memory address isn't mathematically guaranteed to be permanently unique *over time*. As soon as the object is completely garbage collected, CPython will recycle that specific memory block/address. Two objects with non-overlapping lifetimes can absolutely have the same `id()`.

---

### 80. What are Python iterators and generators (Summary)?
"To summarize, Iterators and Generators abstract data streams.
- **Iterable:** Anything you can loop over (list, dict).
- **Iterator:** The engine powering the loop, keeping track of the current state (`__next__`).
- **Generator:** An elegant, function-based Iterator written using `yield`.

They embody Python’s memory efficiency. I rely on Generators whenever I am doing file parsing, web requests streams, or big data pagination, ensuring I never consume excessive server RAM despite infinite data."

#### Indepth
Generators are technically single-pass. Once they throw `StopIteration`, they are exhausted forever. If you need to loop over the data again, you must call the generator function *again* to launch a fresh iterator instance mathematically from scratch.
"""

file_path = r'g:\My Drive\All Documents\java-golang-interview-questions\long_questions\Python\Theory\01_Basics.md'

with open(file_path, 'a', encoding='utf-8') as f:
    f.write(content)
