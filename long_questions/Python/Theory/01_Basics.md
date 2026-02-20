# Basic Level Python Interview Questions

## From Python Basics

# 🟢 **1–20: Basics & Control Flow**

### 1. What is Python?
"Python is a high-level, interpreted, and dynamically-typed programming language. I use it primarily because of its incredible readability and its massive ecosystem of libraries. 

At its core, it emphasizes developer productivity over raw execution speed. For example, I don't need to declare variable types or manage memory manually, which makes prototyping APIs or writing automation scripts blazingly fast. What takes 100 lines in Java usually takes 20 lines in Python."

#### Indepth
Python is typically compiled to bytecode (in `.pyc` files) which is then executed by the Python Virtual Machine (PVM). The reference implementation is CPython, written in C. Because of the Global Interpreter Lock (GIL) in CPython, true parallel thread execution is restricted, making it better suited for I/O-bound rather than CPU-bound concurrency.

---

### 2. What are the key features of Python?
"For me, the standout features are its simplicity, dynamic typing, and the 'batteries-included' standard library.

It's object-oriented but doesn't force you into strict class hierarchies if you just want to write a functional script. It has built-in high-level data structures like lists, dictionaries, and sets which are incredibly optimized. I also appreciate its cross-platform nature; a script I write on my Mac runs flawlessly on a Linux server without recompilation."

#### Indepth
The "batteries-included" philosophy means the standard library comes with modules for regex, networking, thread management, and file I/O out of the box. Additionally, everything in Python is an object, including functions and classes, which enables powerful metaprogramming and higher-order functions.

---

### 3. What is the difference between `list` and `tuple`?
"A **list** is mutable, meaning I can add, remove, or change items after it's created. A **tuple** is immutable; once I define it, its contents cannot be altered.

I use lists when I expect the collection of data to change, like a queue of active tasks. I use tuples for fixed collections of heterogeneous data, like coordinates `(x, y)` or returning multiple values from a function. Because tuples are immutable, they can be used as keys in a dictionary, whereas lists cannot."

#### Indepth
Under the hood, both are arrays of object pointers. However, tuples are more memory-efficient. CPython over-allocates memory for lists to make `append()` operations O(1) amortized. Tuples allocate exactly the memory they need. Also, Python caches small tuples to reuse them, making tuple creation slightly faster.

---

### 4. How is memory managed in Python?
"I don't manually allocate or free memory in Python; the language does it for me. It primarily uses **Reference Counting**. 

Every object keeps track of how many variables or data structures point to it. When I reassign a variable or it goes out of scope, the reference count drops. When the count hits zero, Python immediately reclaims the memory. This makes memory management deterministic for the most part."

#### Indepth
Reference counting can't handle circular references (e.g., Object A references Object B, and B references A). To solve this, CPython includes a **Generational Garbage Collector** that periodically runs to detect and collect these unreachable reference cycles. You can tune this using the `gc` module.

---

### 5. What is a dictionary in Python?
"A dictionary (`dict`) is a built-in data structure that stores key-value pairs. It provides O(1) average time complexity for lookups, insertions, and deletions.

I use dictionaries constantly. Whenever I need to map unique identifiers to data—like mapping user IDs to user objects, or parsing JSON responses from an API—dictionaries are the perfect tool. I usually create them using curly braces `{'key': 'value'}`."

#### Indepth
Dictionaries in Python 3.6+ maintain the insertion order of their keys. Under the hood, they are implemented as hash tables. The keys must be `hashable` (immutable types like strings, numbers, or tuples). If a hash collision occurs, CPython resolves it using open addressing with quadratic probing.

---

### 6. What is the difference between `is` and `==`?
"`==` checks for **value equality**, while `is` checks for **identity** (whether they are the exact same object in memory).

If I have two separate lists with the identical items `a = [1, 2]` and `b = [1, 2]`, `a == b` is True because their contents match. But `a is b` is False because they live at different memory addresses. I almost entirely use `==` for comparisons, except when checking against Singletons like `None`, where I always use `if x is None:`."

#### Indepth
The `is` operator mathematically translates to comparing the memory addresses using the `id()` function (i.e., `id(a) == id(b)`). Because Python interns small integers (from -5 to 256) and some short strings, `a = 10; b = 10` will actually result in `a is b` being True, which can sometimes trip up junior developers.

---

### 7. What are local and global variables?
"A **local variable** is defined inside a function and can only be accessed there. A **global variable** is defined at the module level and can be read by any function in that module.

I try to avoid global variables for mutable state because they make debugging difficult and break thread safety. If a function absolutely needs to modify a global variable, I have to explicitly declare it using the `global` keyword inside the function, although I'd usually prefer passing it as an argument and returning the modified value."

#### Indepth
Python resolves variable scope using the **LEGB rule**: Local, Enclosing, Global, and Built-in. When you reference a variable, Python checks the local function scope first. If not found, it checks enclosing functions (closures), then global module scope, and finally built-in names (like `len()` or `print()`).

---

### 8. What is the difference between `for` and `while` loops?
"I use a **`for` loop** when I have a known sequence or collection to iterate over, like items in a list, keys in a dictionary, or a specific range of numbers. 

I use a **`while` loop** when the iteration depends on a condition being true, and I don't know in advance how many times it will run. For example, listening continuously to a server socket or polling an API until a specific status is returned."

#### Indepth
Python's `for` loops are actually "foreach" loops that work on Iterators. When you call `for x in my_list:`, Python internally calls `iter(my_list)` to get an iterator and then repeatedly calls `next()` on it until a `StopIteration` exception is raised. 

---

### 9. What are `break`, `continue`, and `pass`?
"`break` completely exits the innermost loop immediately. I use it when I've found what I'm looking for and don't need to check remaining items.

`continue` skips the rest of the current loop iteration and moves directly to the next one. I use this to skip invalid data early in a loop.

`pass` does absolutely nothing. It's just a syntactic placeholder. I use it when I'm writing structural boilerplate (like setting up an empty class or an `except` block) where the syntax requires a statement but I haven't written the logic yet."

#### Indepth
A unique feature of Python is the `for...else` construct. The `else` block executes *only* if the loop runs to completion without hitting a `break` statement. This is highly useful for search algorithms, negating the need for a boolean "found" flag.

---

### 10. How does Python handle switch-case?
"Historically, Python didn't have a `switch-case` statement. We would use a chain of `if-elif-else` statements or a dictionary mapping keys to functions to simulate it.

However, starting in Python 3.10, we now have **Structural Pattern Matching** using the `match` and `case` keywords. I use this extensively now; it's much more powerful than a traditional switch because it allows unpacking data structures and conditional guards rather than just checking for equality."

#### Indepth
Under the hood, structural pattern matching avoids evaluating the subject repeatedly. It is highly optimized and can match against classes, dicts, lists, and perform variable binding simultaneously: `case Point(x, y) if x == y:`.

---

### 11. What are `*args` and `**kwargs`?
"They allow me to define functions that accept a variable number of arguments. 

`*args` collects extra positional arguments into a **tuple**. 
`**kwargs` collects extra keyword arguments into a **dictionary**.

I use them constantly for wrapper functions or decorators where my function doesn't know (or care) what arguments the inner function takes. I just write `def wrapper(*args, **kwargs): return inner(*args, **kwargs)` to pass everything smoothly."

#### Indepth
The names `args` and `kwargs` are just conventions; the real magic is the `*` and `**` unpacking operators. You can also use these operators to unpack iterables and dictionaries when calling a function, e.g., `my_func(*[1, 2], **{'a': 3})`, or to merge dictionaries: `z = {**x, **y}`.

---

### 12. What are default arguments and their pitfall?
"Default arguments let me provide a fallback value for a function parameter if the caller doesn't provide one: `def connect(timeout=30):`.

The biggest pitfall is using **mutable default arguments**, like empty lists or dictionaries: `def add_item(item, basket=[]):`.
The default list is instantiated exactly *once* when the function is defined, not every time it's called. So multiple calls will unintentionally share the same list. I always use `None` as the default instead and initialize the list inside the function."

#### Indepth
Because function definitions are executed at module load time to create a function object, default arguments are attached to the function object's `__defaults__` attribute. This is why the mutable object persists across successive calls. 

---

### 13. What is a lambda function?
"A lambda is a small, anonymous, single-expression function. I define it using the `lambda` keyword. 

`lambda x, y: x + y`

I use them for short, throwaway operations where defining a full `def` function feels too verbose. They are most commonly used as the `key` argument in sorting operations, or with higher-order functions like `map()` and `filter()`."

#### Indepth
Lambdas in Python are strictly limited to a single expression. They cannot contain statements (like `if` statements, `assignment`, or `return`). They return the evaluation of that single expression implicitly. Because of their limitations, PEP 8 recommends using a standard `def` if the lambda logic becomes barely readable.

---

### 14. What is the difference between `str.strip()` and `str.lstrip()`?
"`strip()` removes whitespace (or specified characters) from **both** the beginning and the end of a string.
`lstrip()` removes them only from the **left** (start) side.

I use `strip()` virtually every time I parse raw text from a database or a user input field to ensure invisible trailing spaces don't break string comparisons or validations."

#### Indepth
Both methods can take a string of characters to strip, not just whitespace. A common mistake is thinking it strips matched substrings. `strip("abc")` strips *any combinations* of 'a', 'b', or 'c' from the ends, effectively acting as a character set, not a sequence.

---

### 15. What is a list comprehension?
"List comprehension is an elegant, concise way to create lists from existing iterables. 

Instead of writing a 3-line `for` loop with an `.append()`, I can write it in one line: `[x*2 for x in nums if x > 0]`.

I use this all the time because it's more readable and often slightly faster than a standard loop. However, if the logic inside gets too complex with multiple nested loops, I revert to a standard `for` loop so my team can actually read it."

#### Indepth
List comprehensions execute closer to C speed because they avoid the overhead of calling the list's `append` method in every iteration. Python also supports dictionary comprehensions `{k:v for ...}` and set comprehensions `{x for ...}`, but using parentheses `(x for ...)` creates a generator expression rather than a tuple comprehension.

---

### 16. How do you remove elements from a list?
"I have several options depending on the need:
1. `list.remove(value)` removes the *first occurrence* of a specific value.
2. `list.pop(index)` removes the item at a specific index and *returns* it. If I don't provide an index, it pops the last item.
3. `del list[index]` explicitly deletes the item at an index or a slice.

I use `pop()` when I'm treating the list like a stack or queue and need the value. I use `remove()` when I know the value but not its position."

#### Indepth
If you use `remove(value)`, Python must perform a linear search (O(N)) to find the element, and then shift all subsequent elements left (O(N)), making it inefficient for large lists. Popping from the end of a list is O(1), but popping from the beginning is O(N). For a true double-ended queue, `collections.deque` is the correct structure.

---

### 17. How do you loop through a dictionary?
"By default, if I do `for key in my_dict:`, it loops through the keys. 

But most of the time, I need both the key and the value. For that, I use `for key, value in my_dict.items():`. 
This gives me a tuple of the key and value on every iteration, which I unpack immediately. It's clean and idiomatically Pythonic."

#### Indepth
`dict.keys()`, `dict.values()`, and `dict.items()` return dictionary view objects, not lists (unlike Python 2). View objects dynamically reflect changes to the dictionary and allow for fast set-like operations (e.g., finding the intersection of keys between two dicts using bitwise `&`).

---

### 18. What is the difference between `dict.get('key')` and `dict['key']`?
"When I use square brackets `dict['key']`, Python will raise a `KeyError` if the key doesn't exist. It's a strict approach.

When I use `dict.get('key')`, it returns the value if the key exists, but gracefully returns `None` (or a default value I specify) if the key is missing.

I heavily prefer `.get()` when dealing with external payloads like JSON from a web request, where fields might be unexpectedly missing. It prevents my application from crashing over a missing optional field."

#### Indepth
The `.get(key, default)` method is an atomic lookup. An alternative for dealing with missing keys is the `collections.defaultdict` class, which automatically instantiates a default value using a factory function when an unknown key is queried via square brackets.

---

### 19. What is a set in Python?
"A set is an unordered collection of unique elements. 

I use sets primarily for two things: rapidly removing duplicates from a list (`list(set(my_list))`), and performing mathematical set operations like unions, intersections, and differences. 

Because sets are implemented as hash tables, checking membership (`if item in my_set:`) is blazingly fast (O(1)) compared to a list (O(N))."

#### Indepth
Sets require their elements to be hashable (immutable), so you cannot put lists or dictionaries inside a set. Since sets are mutable, they themselves cannot be used as dictionary keys or nested within other sets. If you need an immutable set, Python provides the `frozenset` type.

---

### 20. What is list slicing?
"Slicing allows me to extract a portion of a list (or string) using the syntax `[start:stop:step]`.

If I want the first 3 items: `my_list[:3]`.
If I want the last item or to count from the end: `my_list[-1]`.
If I want to reverse the list: `my_list[::-1]`.

It's an incredibly expressive syntax. Every slice operation creates a completely new list (a shallow copy) in memory, which is useful when I need to clone a list safely without modifying the original: `new_list = old_list[:]`."

#### Indepth
Underneath, slicing creates a `slice` object (`slice(start, stop, step)`) which is passed to the list's `__getitem__` method. While slicing creates copies, if you assign *to* a slice (`my_list[1:3] = [8, 9]`), it modifies the original list in place, effectively replacing that segment with the new iterable.

## From Collections & Exceptions

# 🟡 **21–40: Advanced Data Structures & Error Handling**

### 21. How do you append to a list and how do you extend it?
"I use `append(item)` when I want to add a *single* element to the end of a list. If I append another list `[4, 5]`, it goes in as one single sublist element.

I use `extend(iterable)` when I have another sequence (like a list, tuple, or set) and I want to add *each individual item* from that sequence to the end of my current list. It merges them. `listA.extend(listB)` modifies `listA` in place and is faster than a `for` loop with `append`."

#### Indepth
`extend()` takes any iterable and effectively calls CPython's `list_extend` C function, which pre-allocates the exact memory needed for the new elements and bulk-copies them in, bypassing iterating through byte code entirely, making it highly optimized.

---

### 22. What is tuple unpacking?
"Tuple unpacking lets me assign the individual elements of a tuple (or any iterable) to a series of variables in one clean line.

If I have a tuple `point = (10, 20)`, I can write `x, y = point`. 

I use this all the time for functions that return multiple values. It makes the code incredibly readable. I can even use the `*` operator to gather remaining elements: `first, *middle, last = [1, 2, 3, 4, 5]`, which assigns `[2, 3, 4]` to `middle`."

#### Indepth
Unpacking evaluates the entire right-hand side before any assignments occur on the left-hand side. This is why the classic swap idiom `a, b = b, a` works without a temporary variable; Python builds a temporary tuple on the C stack and then unpacks it.

---

### 23. What are dictionary comprehensions?
"Similar to list comprehensions, dict comprehensions allow me to dynamically generate dictionaries in a single readable line. 

`{x: x**2 for x in (2, 4, 6)}` generates `{2: 4, 4: 16, 6: 36}`.

I frequently use them to invert dictionaries (swapping keys and values): `{v: k for k, v in my_dict.items()}`, assuming values are unique and hashable. It replaces messy loops with a declarative expression."

#### Indepth
Since Python 3.8, dictionary comprehensions execute slightly slower than calling `dict()` on a generator expression if the generator is simple, due to dictionary initialization overhead, but for any complex logic, comprehensions are faster and definitely more standard in the Python ecosystem.

---

### 24. What are exceptions in Python?
"Exceptions are events triggered when Python encounters a runtime error. Instead of crashing immediately, Python stops normal execution and 'raises' an exception object.

I use `try/except` blocks to catch and handle these gracefully. For example, catching a `NetworkError` to trigger a retry mechanism rather than letting the web server crash. It separates the 'happy path' business logic from error-handling logic."

#### Indepth
Exceptions map to a class hierarchy inheriting from `BaseException` (like `KeyboardInterrupt`) and `Exception` (standard application errors like `ValueError` and `KeyError`). You should rarely catch `BaseException` or bare `except:`, as this catches system-exiting signals like `sys.exit()` unexpectedly.

---

### 25. How do you handle exceptions in Python?
"I wrap the risky code in a `try` block and catch specific errors in `except` blocks.

`try: \n  val = int(user_input) \nexcept ValueError: \n  print('Invalid number')`

I never use a bare `except:` clause because it swallows *everything*, including typos in my code (like `NameError`) or user cancellations (`KeyboardInterrupt`). It's considered an anti-pattern. I always catch the tightest specific error possible."

#### Indepth
Python supports an `else` clause at the end of a `try/except` block. The `else` block runs *only if* no exceptions were raised. This is excellent for keeping the `try` block small (only covering the code that might fail) and putting the dependent success logic in the `else` block.

---

### 26. What is the use of `finally`?
"The `finally` block executes *no matter what* happens in the `try` block—whether it ran perfectly, crashed with an exception, or even executed a `return` or `break` statement.

I use `finally` for strictly guaranteed cleanup operations. If I open a file, network socket, or database connection inside the `try`, I use `finally` to ensure `.close()` is called so I don't bleed resources."

#### Indepth
While `finally` is standard, modern Python heavily prefers **Context Managers** (`with` statements) for resource cleanup because it embeds the setup and teardown logic directly into the object, making it impossible to forget to write the `finally` block.

---

### 27. How does Python manage scopes (LEGB)?
"Python evaluates scope strictly following the LEGB rule, moving outwards.

1. **L**ocal: Inside the current function.
2. **E**nclosing: Inside any functions wrapping the current function (Closures).
3. **G**lobal: At the top level of the module file.
4. **B**uilt-in: Python's pre-defined names like `len` or `Exception`.

When I use an unknown variable, Python queries LEGB in order. If it hits the end without finding it, it raises a `NameError`."

#### Indepth
If you assign to a variable anywhere in a function, Python statically analyzes that variable as Local for the *entire* function. If you try to print it before the assignment, you don't get the Global variable; you get an `UnboundLocalError`.

---

### 28. What is the difference between `remove()`, `del`, and `pop()`?
"`remove(value)` removes the *first occurrence* of a specific value. I don't know the index.
`del list[index]` explicitly deletes the element at a specific index or slice without returning it.
`pop(index)` removes the element at an index *and returns it* to me so I can assign it to a variable.

I use `del` for cleaning up dictionaries or slicing lists (`del arr[2:5]`), `pop` for queues, and `remove` when the value is my only reference."

#### Indepth
`del` is not a function; it is a keyword and statement. It effectively drops the reference from the current namespace scope and decrements the object's reference count. If the reference count drops to zero, the Garbage Collector picks it up.

---

### 29. How do you implement a stack in Python?
"I use a standard Python `list`.

Lists are implemented efficiently for operations at the end of the array. I use `append()` to push items onto the top of the stack, and `pop()` (with no arguments) to remove them from the top. Both operations are O(1). 

I don't need custom classes or node pointers; the built-in list is perfect for stacks."

#### Indepth
You can optionally use `collections.deque` if you prefer queue-oriented semantics (`append` and `pop`), but a `list` actually has better cache locality due to sequential memory allocation, making lists slightly faster for pure Stack usage than deques.

---

### 30. How do you format a string using f-strings?
"F-strings (formatted string literals) were introduced in Python 3.6, and I use them exclusively now over `.format()` or `%` substitution.

I prefix the string with an `f` and put variables directly inside curly braces: `f'My name is {name} and I am {age}'`.

I use them because they are faster at runtime (evaluated in C directly) and much easier to read than matching positional arguments at the end of a long string."

#### Indepth
F-strings can execute arbitrary Python expressions inside the braces: `f"Result: {math.sqrt(x):.2f}"` (also formatting to 2 decimal places). Python 3.8 introduced a debug syntax: `f"{user_id=}"` which outputs `"user_id=123"`, perfect for quick `print()` debugging.

---

### 31. What is the difference between a module and a package?
"A **module** is simply a single `.py` file containing Python definitions and statements. If I create `math_tools.py`, it's a module.

A **package** is a directory that contains multiple modules, and crucially, an `__init__.py` file (though in modern Python this is optional for namespace packages). 

I use packages to organize a large project into a hierarchy, like `auth/login.py` and `auth/roles.py`, allowing me to use dotted imports like `from auth.login import authenticate`."

#### Indepth
The `__init__.py` file initializes the package. If I place code in `__init__.py`, it runs exactly once the first time the package or any of its submodules are imported. It is commonly used to expose a clean public API for the package by selectively importing internal submodule classes.

---

### 32. What does `if __name__ == '__main__':` do?
"This idiom checks if the current script is being run directly from the terminal or if it's being imported into another file.

If I run `python script.py`, the special variable `__name__` is set to the string `'__main__'`. If another file imports my script, `__name__` is set to `'script'`.

I use this at the bottom of almost every utility file. It allows me to include unit tests or an example execution block that *only* runs locally, but stays completely quiet when imported as a library by the main server."

#### Indepth
This structure acts as the program entry point (the pseudo `main()` function) in a language that otherwise executes scripts sequentially from top to bottom. It enforces module reusability.

---

### 33. What is `with open()` used for?
"It invokes a Context Manager to handle file operations safely.

`with open('data.txt', 'r') as file:` handles opening the file, and guarantees that the file is automatically closed when the indented block finishes, *even if an exception occurs inside the block*.

I use it religiously. Before Context Managers, developers (including me) often forgot to write `file.close()` in an explicit `finally` block, leading to file handle leaks and 'Too many open files' crashes on servers."

#### Indepth
The `with` statement works through the `__enter__()` and `__exit__()` magic methods. When `open` is bound by `with`, `__enter__` returns the file object, and `__exit__` automatically calls `close` when the block terminates, cleanly abstracting away resource management.

---

### 34. What is the difference between `read()` and `readlines()`?
"`file.read()` loads the entire file contents into memory as a single, massive string. I only use it for small configuration files.

`file.readlines()` loads the entire file into memory but splits it into a list of strings, one string per line.

For logs or large data files, I don't use either. I iterate directly over the file object: `for line in file:`. This streams the file lazily one line at a time, using almost zero memory regardless of whether the file is 1MB or 50GB."

#### Indepth
`file.read(size)` can take an optional byte parameter to read the file in chunks. But direct iteration (`for line in file`) is implemented under the hood using an internal buffer read, making it not just memory-efficient but exceptionally fast due to standard library C-optimizations.

---

### 35. What does the `zip()` function do?
"The `zip()` function takes two or more iterables (like lists) and pairs their elements together side-by-side into tuples.

If I have `names = ['A', 'B']` and `ages = [20, 30]`, `zip(names, ages)` creates `[('A', 20), ('B', 30)]`. 

I use it mostly in `for` loops when I need to iterate through two corresponding lists simultaneously without explicitly managing an integer index variable."

#### Indepth
`zip()` stops lazily when the *shortest* input iterable is exhausted. If you have unequal lists and don't want to lose data, you must use `itertools.zip_longest()`, which fills the gaps in the shorter lists with `None` or a custom fill value.

---

### 36. What is `enumerate()`?
"`enumerate()` allows me to loop over an iterable while simultaneously keeping track of the current index counter.

Instead of initializing `i = 0` and doing `i += 1` inside a `for item in items:` loop, I write `for i, item in enumerate(items):`.

It’s vastly more Pythonic and less error-prone. I often pass the `start` parameter, like `enumerate(items, start=1)`, if I am generating human-readable rows or outputs."

#### Indepth
`enumerate` returns an iterator generator under the hood, not a list of tuples. Therefore, it is instantaneous to create and takes almost no memory, regardless of how large the wrapped iterable might be.

---

### 37. What are `any()` and `all()`?
"These are built-in reduction logic functions. 

`all()` returns True only if *every* element in an iterable evaluates to True. 
`any()` returns True if *at least one* element evaluates to True.

I use these heavily with Generator Expressions for validation. For instance, `if any(bad_word in text for bad_word in blacklist):` is an elegant, single-line way to check for profanity in a text stream."

#### Indepth
Both functions **short-circuit**. As soon as `all()` sees a single False item, it returns False immediately without checking the rest of the list. If `any()` sees a single True item, it returns True immediately. This makes them highly performant on large data sets if the breaking condition is near the beginning.

---

### 38. How to flatten a nested list?
"If I have a moderately nested list like `[[1,2], [3,4]]`, my go-to is a flat list comprehension: `[item for sublist in nested_list for item in sublist]`.

However, if it's deeply nested (like nested JSON structures), I write a recursive function (or stack generator) to traverse and yield elements dynamically. Itertools `chain.from_iterable` is also a great standard library option for flattening one level."

#### Indepth
List comprehensions that flatten are often hard for newcomers to read because the `for` loops are evaluated strictly from left to right: `<outer loop> <inner loop> <expression>`. Using `itertools.chain(*nested_list)` is generally faster since it relies strictly on C implementation.

---

### 39. Are Python sets ordered?
"Historically, no. They are completely unordered collections backed by hash tables, just like old dictionaries.

When I run a `for` loop over a standard set, I might get the elements in any random order based on memory hashes, which can even change between Python executions. 

If I require an ordered set, I use a dictionary structure, because dicts *are* ordered (since Python 3.7), acting as `dict.fromkeys(elements)` to simulate ordered uniqueness."

#### Indepth
In CPython 3.6+, dictionaries maintain order not through the hash table buckets directly, but via a dense array of indices pointing to a compact array that holds the KV pairs chronologically. Sets do not have this dual-array architecture, so they remain permanently, mathematically unordered to save memory.

---

### 40. How do you swap two variables without a temp variable?
"I use Python’s built-in tuple packing and unpacking syntax: 
`a, b = b, a`

I love this feature. In languages like C or Java, swapping requires a third temporary temporary variable (`temp = a; a = b; b = temp`). In Python, the right side is completely evaluated into an unnamed tuple in memory *before* being bound to the variables on the left side."

#### Indepth
At the bytecode level, Python employs specific opcodes (`ROT_TWO`, `ROT_THREE`) to rotate items immediately on the internal stack, eliminating the overhead of actually creating and destroying the tuple object in memory, making it incredibly fast.


## From File Handling & OS

# 🟡 **41–60: File I/O, OS, and Environment**

### 41. How do you read a file line by line?
"To read a file efficiently line by line, I simply iterate over the file object itself.

`with open('large_file.txt', 'r') as file:`
`    for line in file:`
`        process(line.strip())`

This is the standard, most memory-efficient approach. It reads the file chunk by chunk (usually 4KB or 8KB at a time) and yields lines lazily, meaning I can process a 100GB text file on a machine with 512MB of RAM without crashing."

#### Indepth
Behind the scenes, the file object returned by `open()` is an iterable that uses an internal readline buffer. Calling `file.readlines()` does the opposite: it loads the *entire* file into a list of strings in RAM, which can result in an `MemoryError` for large files.

---

### 42. How do you append to a file?
"To append data without overwriting the existing content, I open the file in `append` mode using `'a'` or `'a+'`.

`with open('log.txt', 'a') as file:`
`    file.write('New log entry\n')`

The file pointer is automatically placed at the end of the file. If the file doesn't exist, Python will seamlessly create it for me."

#### Indepth
When opening files in `'a'` mode, seeking backwards (e.g., `file.seek(0)`) is technically allowed, but any subsequent `write()` call will automatically jump back to the end of the file before writing, depending on the OS implementation (POSIX append mode).

---

### 43. What is the difference between `w+` and `a+` modes?
"Both allow reading and writing, but they handle existing data entirely differently.

`'w+'` (write and read) **truncates** the file to zero length immediately upon opening. All existing data is instantly destroyed.
`'a+'` (append and read) preserves existing data and places the write pointer at the **end** of the file.

If I need to update a config file without deleting it, I use `r+`. I almost never use `w+` unless my explicit intent is to completely wipe the file first."

#### Indepth
In `'a+'` mode, while the write pointer is locked to the end, the *read* pointer defaults to the end too. If you want to read from the beginning, you must explicitly call `file.seek(0)` before calling `file.read()`, otherwise you will get an empty string.

---

### 44. How do you check if a file exists?
"I use the `pathlib` module, which is the modern, object-oriented way to handle paths in Python 3.

`from pathlib import Path`
`if Path('data.csv').exists():`
`    # do something`

Alternatively, the legacy way is `os.path.exists('data.csv')`. But `pathlib` is safer because it explicitly allows me to check `Path('data.csv').is_file()` to ensure it's actually a file and not a directory."

#### Indepth
Checking if a file exists before opening it can cause a **Race Condition** (TOCTOU: Time of check to time of use). A safer practice (EAFP: Easier to Ask for Forgiveness than Permission) is to just `try` to open the file and catch the `FileNotFoundError`.

---

### 45. What does `os.listdir()` do?
"`os.listdir(path)` returns a list of all files and directories inside the specified path. 

However, it only returns the string names (e.g., `['file.txt', 'img.png']`), not the full paths, and it doesn't distinguish between files and folders.

Today, I heavily prefer `os.scandir(path)` or `Path(path).iterdir()`. Both return iterator objects with rich attributes, allowing me to instantly check `.is_file()` or get `.stat()` metadata without making extra expensive system calls."

#### Indepth
`os.scandir()` was introduced in Python 3.5 to dramatically speed up directory traversal (like `os.walk()`). It caches file attributes (like whether it is a directory) during the initial fetch, which on Windows can result in a 5x to 20x performance improvement over `listdir + isdir`.

---

### 46. How do you create a new directory?
"I use `os.mkdir('new_folder')` if I'm just creating a single directory.

But if I need to create a nested path, like `logs/2026/january/`, I exclusively use `os.makedirs()`. It creates all necessary intermediate directories automatically. 

With `pathlib`, it’s exactly the same concept: `Path('logs/2026/january/').mkdir(parents=True, exist_ok=True)`. The `exist_ok=True` is fantastic because it prevents an error if the directory is already there."

#### Indepth
`os.makedirs(name, mode=0o777, exist_ok=False)` also accepts a permission mode (octal). Note that the mode is subject to the current umask of the process, which usually clears the group/world write permissions by default.

---

### 47. How do you delete a file and a directory?
"To delete a file, I use `os.remove('file.txt')` or `Path('file.txt').unlink()`.

To delete an empty directory, I use `os.rmdir('folder')`. 

If the directory contains files, `rmdir` will fail for safety reasons. To forcefully delete a directory and everything inside it, I import `shutil` and use `shutil.rmtree('folder')`. I have to be extremely careful with `rmtree` because it permanently destroys data with no recovery."

#### Indepth
Both `os.remove()` and `os.unlink()` do exactly the same thing. The name `unlink` comes from UNIX terminology, as deleting a file actually just removes a hard link; the data is only freed by the OS when the link count reaches zero.

---

### 48. What is `os.path.join()` used for?
"`os.path.join()` intelligently glues path segments together using the correct directory separator for the current operating system.

Instead of hardcoding `folder + '/' + filename`, which crashes on Windows, I use `os.path.join('folder', 'filename')`. If I'm on Linux, it outputs `folder/filename`. On Windows, it outputs `folder\filename`.

In modern code, `pathlib` handles this even more elegantly with the slash operator: `Path('folder') / 'filename'`."

#### Indepth
`os.path.join` has a slightly confusing edge case: if any component is an absolute path (starts with `/` on Linux or `C:\` on Windows), all previous components are discarded and joining continues from the absolute path component.

---

### 49. How do you get the current working directory?
"I use `os.getcwd()`. 

This returns a string representing the absolute path of the directory from which the Python script was *executed*, not necessarily where the script file lives.

If I explicitly need the directory where the current Python file sits (e.g., to load a relative config file), I compute it dynamically using `os.path.dirname(os.path.abspath(__file__))`. In `pathlib`, that is `Path(__file__).parent.resolve()`."

#### Indepth
Using `getcwd()` is dangerous in larger applications because changing it (`os.chdir()`) changes global state for the entire process, including all threads. Relying on `__file__` makes your modules portable and independent of the execution context.

---

### 50. What does `os.path.exists()` check?
"It checks if a path (file or directory) exists physically on the disk.

If it exists, it returns `True`. If it doesn't, or if the process lacks the operating system permissions to access the path (like a protected root folder), it safely returns `False`.

However, as per the LEAP/EAFP principle, for purely boolean existence checks, it's fine. For file operations, blindly opening and catching `FileNotFoundError` is practically superior."

#### Indepth
`os.path.exists()` follows symbolic links (symlinks). If a symlink exists but points to a deleted file (a broken link), `exists()` returns `False`. If you explicitly want to check if the link itself exists, broken or not, use `os.path.lexists()`.

---

### 51. What is a virtual environment?
"A virtual environment (`venv`) is a self-contained directory tree that isolates Python installations and packages.

I use virtual environments for every single project. It ensures that if Project A needs `requests==2.10` and Project B needs `requests==3.0`, they don't conflict. 

It keeps the global system Python installation completely clean and drastically reduces 'it works on my machine' deployment bugs."

#### Indepth
Virtual environments don't copy the entire Python binary. They create a lightweight structure including a `pyvenv.cfg` file and symlinks (or hard links on Windows) to the base Python executable. The activation script just prepends the environment's `bin` or `Scripts` directory to the shell's `PATH`.

---

### 52. How do you create and activate a virtual environment?
"I run this command in my terminal to create it: `python -m venv .venv`. (The `.venv` is just a convention for the hidden folder name).

To activate it:
On macOS/Linux: `source .venv/bin/activate`
On Windows: `.venv\Scripts\activate`

Once activated, my terminal prompt changes, and any `pip install` I run will securely place the libraries inside the `.venv` folder."

#### Indepth
You don't *strictly* have to activate a venv to use it. You can simply invoke the Python executable directly inside the environment (`/path/to/venv/bin/python script.py`), and it will automatically comprehend its isolated library paths relative to that isolated binary.

---

### 53. What is `pip`?
"`pip` stands for 'Pip Installs Packages'. It is the standard package manager for Python.

I use it to download, install, and manage third-party libraries from the Python Package Index (PyPI). 

`pip install requests` will download the library and its dependencies. I also frequently use `pip freeze > requirements.txt` to capture the exact versions of all installed packages so my deployment environment can replicate them using `pip install -r requirements.txt`."

#### Indepth
`pip` does not implement true dependency resolution in the way `npm` or `cargo` does. If Package A needs `requests<2` and Package B needs `requests>3`, traditional `pip` would just overwrite the installation depending on the install order. Modern `pip` (v20.3+) includes a strict dependency resolver that throws an error in these conflicts.

---

### 54. What is `__pycache__`?
"When Python imports a module, it compiles the human-readable `.py` source code into intermediate bytecode and caches it as a `.pyc` file inside the `__pycache__` directory.

This makes strictly the *import process* faster on subsequent runs because Python doesn't need to recompile the file unless the source code timestamp has changed. 

I entirely ignore it. I always add `__pycache__/` to my `.gitignore` to prevent committing compiled, machine-specific binaries to source control."

#### Indepth
The compiled bytecode is largely platform-independent (it targets the Python VM), but it *is* strictly tied to the Python version. A `pyc` generated by Python 3.9 cannot be loaded by Python 3.10. The naming convention reflects this: `script.cpython-310.pyc`.

---

### 55. What is `requirements.txt` vs `setup.py`?
"`requirements.txt` is an exact, pinned manifest for an application's deployment. It lists every library and its specific version (e.g., `requests==2.28.1`) ensuring the app runs identically in production.

`setup.py` (or modern `pyproject.toml`) is for *distributing libraries*. It defines abstract dependencies (e.g., `requests>=2.0`) to ensure maximum compatibility if someone installs my library alongside other tools.

If I'm building an app, I use `requirements.txt`. If I'm building a library for PyPI, I use `setup.py`."

#### Indepth
A common advanced workflow uses both. `setup.py` or `pyproject.toml` defines the loose dependencies. Tools like `pip-tools` or `Poetry` read those loose dependencies, solve the entire dependency graph, and generate a strictly pinned `requirements.txt` (or lockfile) for deployment.

---

### 56. What happens when you type `python script.py`?
"Python executes in a sequence of distinct phases:
1. **Parsing:** The CPython compiler parses the source code into an Abstract Syntax Tree (AST).
2. **Compilation:** It compiles the AST into specialized bytecode instructions.
3. **Execution:** The Python Virtual Machine (PVM)—a massive C loop—iterates over the bytecode instructions and executes them one by one.

If `script.py` imports modules, those modules go through the same process and are cached as `.pyc` files for future speedups."

#### Indepth
The Python Virtual Machine is a stack-based machine. Almost every instruction operates by pushing variables onto an evaluation stack, doing work (like adding the top two items), and pushing the result back onto the stack.

---

### 57. Can Python be used for mobile apps?
"Yes, but it is not the native or standard choice. 

Frameworks like **Kivy** or **BeeWare** allow me to write Python code and package it into native iOS or Android applications. 

However, because iOS and Android are heavily optimized for Swift/Kotlin respectively, Python apps carry a massive footprint (they embed the Python interpreter) and can suffer performance and UI/UX issues. For serious mobile development, I prefer native languages or React Native/Flutter over Python."

#### Indepth
The main challenge is the `GIL` (Global Interpreter Lock). Mobile OS UI threads demand instant responsiveness (60+ FPS). Dealing with the GIL in highly concurrent, event-driven UI environments is notoriously difficult.

---

### 58. How to run a Python script with command-line arguments?
"I pass arguments simply by appending them in the terminal: `python script.py --env prod port=8080`.

Inside the script, I can access these arguments natively via the `sys.argv` list. `sys.argv[0]` is the script name, and subsequent indices are the arguments as raw strings.

However, handling strings manually is tedious. I almost universally use the built-in `argparse` module, which automatically handles parsing, type checking, default values, and automatically generates beautiful `--help` documentation."

#### Indepth
For modern, complex CLIs, libraries like `Click` (used by Flask) or `Typer` (which uses type hints) heavily abstract `argparse` and make building robust command-line tools a matter of writing single decorated functions.

---

### 59. What does the `math` module provide?
"The `math` module is the standard library for advanced mathematical operations. 

It provides C-optimized functions for trigonometry (`sin`, `cos`), logarithms (`log`, `log10`), and foundational constants (`math.pi`, `math.e`). 

It also includes highly robust functions like `math.isclose()` which I use to safely compare floating-point numbers instead of using `==` to avoid precision errors, and `math.floor/ceil` for accurate rounding."

#### Indepth
The `math` module is strictly for scalar numbers. If you attempt to pass a list of numbers to `math.sqrt()`, it crashes with a TypeError. For high-performance, vectorized mathematical operations across arrays or matrices, the third-party framework `NumPy` is universally used.

---

### 60. How do you check if a Python script is running interactively?
"To check if I'm running inside the REPL or an interactive shell, I can check if the `sys` module possesses the `ps1` attribute (which represents the `>>>` prompt).

`import sys`
`is_interactive = hasattr(sys, 'ps1')`

I occasionally use this to modify logging behavior. If I'm running interactively, I might just print warnings directly. If running as a daemon script, I ensure everything routes cleanly to the rotating log file."

#### Indepth
Another method is to use the `isatty()` module method on output streams. `sys.stdout.isatty()` returns True if the output is connected to a terminal, and False if the output is currently being piped or redirected to a text file (e.g., `python script.py > output.txt`).

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
