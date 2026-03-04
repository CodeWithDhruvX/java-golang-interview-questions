# 📘 04 — File Handling & Exception Handling
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- Reading and writing files
- Context managers (`with` statement)
- Exception types and hierarchy
- `try/except/else/finally` patterns
- Custom exceptions
- `pathlib` for modern file operations

---

## ❓ Most Asked Questions

### Q1. How do you read and write files in Python?

```python
# Writing to a file
with open("data.txt", "w") as f:    # 'w' = write (truncates existing)
    f.write("Hello, World!\n")
    f.write("Second line\n")

# Appending to file
with open("data.txt", "a") as f:    # 'a' = append
    f.write("Third line\n")

# Reading entire file
with open("data.txt", "r") as f:    # 'r' = read (default)
    content = f.read()              # entire file as string

# Reading line by line (memory efficient for large files)
with open("data.txt", "r") as f:
    for line in f:                  # iterates line by line — lazy!
        print(line.strip())

# Reading all lines into a list
with open("data.txt", "r") as f:
    lines = f.readlines()           # list of strings with '\n'

with open("data.txt", "r") as f:
    lines = f.read().splitlines()   # list of strings WITHOUT '\n'

# Reading first N bytes
with open("data.txt", "r") as f:
    chunk = f.read(100)             # read 100 bytes

# File modes:
# "r"  → read (default)
# "w"  → write (truncate)
# "a"  → append
# "x"  → create (fails if exists)
# "b"  → binary mode (combine: "rb", "wb")
# "+"  → read+write ("r+", "w+")

# Binary files (images, PDFs, etc.)
with open("image.png", "rb") as f:
    data = f.read()

with open("copy.png", "wb") as f:
    f.write(data)

# Encoding (always specify for text files)
with open("unicode.txt", "w", encoding="utf-8") as f:
    f.write("नमस्ते दुनिया")
```

---

### Q2. Explain exception handling with `try/except/else/finally`.

```python
# Basic exception handling
try:
    result = 10 / 0
except ZeroDivisionError:
    print("Cannot divide by zero!")

# Catching multiple exceptions
def safe_convert(value):
    try:
        return int(value)
    except (ValueError, TypeError) as e:
        print(f"Conversion failed: {e}")
        return None

# else: runs if NO exception was raised
def read_file(filename):
    try:
        f = open(filename)
    except FileNotFoundError:
        print(f"File not found: {filename}")
        return None
    else:
        # Only runs if open() succeeded
        content = f.read()
        f.close()
        return content

# finally: ALWAYS runs (cleanup code)
def process_file(filename):
    f = None
    try:
        f = open(filename, "r")
        data = f.read()
        return data.upper()
    except FileNotFoundError:
        print("File not found!")
        return ""
    except PermissionError:
        print("No permission to read file!")
        return ""
    finally:
        if f:
            f.close()   # always close, even if exception occurred

# Catching all exceptions (use sparingly!)
try:
    risky_operation()
except Exception as e:
    print(f"Unexpected error: {type(e).__name__}: {e}")
    raise   # re-raise to not swallow the error

# Exception chaining
def connect_db():
    try:
        return open("db_config.txt")
    except FileNotFoundError as e:
        raise RuntimeError("Cannot start: DB config missing") from e
```

---

### Q3. How do you create and use custom exceptions?

```python
# Custom exceptions — inherit from Exception (or more specific base)

class AppError(Exception):
    """Base exception for our app"""
    pass

class ValidationError(AppError):
    """Raised when input validation fails"""
    def __init__(self, field, message):
        self.field = field
        self.message = message
        super().__init__(f"Validation error on '{field}': {message}")

class AuthenticationError(AppError):
    """Raised when authentication fails"""
    def __init__(self, username):
        self.username = username
        super().__init__(f"Authentication failed for user: {username}")

class InsufficientFundsError(AppError):
    """Raised when withdrawal exceeds balance"""
    def __init__(self, balance, amount):
        self.balance = balance
        self.amount = amount
        super().__init__(
            f"Insufficient funds: attempted ₹{amount:,}, available ₹{balance:,}"
        )

# Using custom exceptions
def validate_age(age):
    if not isinstance(age, int):
        raise ValidationError("age", "must be an integer")
    if age < 0 or age > 150:
        raise ValidationError("age", f"must be between 0 and 150, got {age}")
    return age

def withdraw(balance, amount):
    if amount > balance:
        raise InsufficientFundsError(balance, amount)
    return balance - amount

# Catching custom exceptions
try:
    validate_age("twenty")
except ValidationError as e:
    print(f"Field: {e.field}")
    print(f"Error: {e.message}")
except AppError as e:
    print(f"App error: {e}")

try:
    new_balance = withdraw(1000, 5000)
except InsufficientFundsError as e:
    print(e)  # Insufficient funds: attempted ₹5,000, available ₹1,000
```

---

### Q4. What is the `with` statement and how do context managers work?

```python
# Context manager: guarantees cleanup using __enter__ and __exit__
# 'with' statement calls __enter__ on enter, __exit__ on exit (even on error)

# File example — most common use case
with open("file.txt", "r") as f:
    data = f.read()
# f.close() called automatically — even if exception occurs!

# Writing a custom context manager using a class
class Timer:
    def __init__(self, name):
        self.name = name

    def __enter__(self):
        import time
        self.start = time.perf_counter()
        print(f"Started: {self.name}")
        return self   # returned as 'as' variable

    def __exit__(self, exc_type, exc_val, exc_tb):
        import time
        elapsed = time.perf_counter() - self.start
        print(f"Finished: {self.name} in {elapsed:.4f}s")
        return False  # False = don't suppress exceptions

with Timer("database query") as t:
    result = [x**2 for x in range(1_000_000)]
# Finished: database query in 0.0843s

# Using contextlib.contextmanager (generator-based — simpler!)
from contextlib import contextmanager

@contextmanager
def managed_resource(name):
    print(f"Acquiring: {name}")
    try:
        yield name   # 'as' variable = whatever is yielded
    finally:
        print(f"Releasing: {name}")

with managed_resource("DB connection") as res:
    print(f"Using {res}")

# Multiple context managers
with open("input.txt") as fin, open("output.txt", "w") as fout:
    fout.write(fin.read().upper())

# contextlib.suppress — swallow specific exceptions
from contextlib import suppress

with suppress(FileNotFoundError):
    import os
    os.remove("temp.txt")  # no error if file doesn't exist
```

---

### Q5. How do you use `pathlib` for modern file operations?

```python
from pathlib import Path

# Create paths
p = Path("data/reports/2024")
p = Path.home() / "Documents" / "projects"  # / operator joins paths!

# Check existence and type
p.exists()    # True/False
p.is_file()   # True if it's a file
p.is_dir()    # True if it's a directory

# File properties
f = Path("report.pdf")
f.name        # "report.pdf"
f.stem        # "report"
f.suffix      # ".pdf"
f.parent      # Path(".")
f.parts       # ('report.pdf',)

# Create directories
Path("data/output").mkdir(parents=True, exist_ok=True)

# Read/write (no need to open/close!)
p = Path("hello.txt")
p.write_text("Hello, World!", encoding="utf-8")
content = p.read_text(encoding="utf-8")

p.write_bytes(b"\x89PNG\r\n")  # binary
data = p.read_bytes()

# List directory contents
for item in Path(".").iterdir():
    print(item)

# Glob patterns
for py_file in Path(".").glob("**/*.py"):  # recursive
    print(py_file)

for md_file in Path("docs").glob("*.md"):  # non-recursive
    print(md_file)

# File operations
src = Path("source.txt")
dst = Path("destination.txt")
src.rename(dst)          # rename (move)
import shutil
shutil.copy(src, dst)    # copy file

# Delete
Path("temp.txt").unlink(missing_ok=True)   # delete file
Path("empty_dir").rmdir()                  # delete empty dir
shutil.rmtree("nonempty_dir")              # delete dir and all contents

# Path comparisons
print(Path("a/b/c") == Path("a/b/c"))   # True
```
