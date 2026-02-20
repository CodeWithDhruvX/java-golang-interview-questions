import os

content = """
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

`try: \\n  val = int(user_input) \\nexcept ValueError: \\n  print('Invalid number')`

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

"""

file_path = r'g:\My Drive\All Documents\java-golang-interview-questions\long_questions\Python\Theory\01_Basics.md'

with open(file_path, 'a', encoding='utf-8') as f:
    f.write(content)
