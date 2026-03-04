# 🐍 Python Interview Questions — Service-Based Companies

> Questions tailored for **TCS, Infosys, Wipro, Cognizant, HCL, Capgemini, Accenture** and similar service-based organizations.
> Difficulty: 🟢 Easy–Medium | Focus: Fundamentals, practical usage, common patterns

---

## 📂 Topics Covered

| # | File | Topics |
|---|------|--------|
| 01 | [Python Basics & Core](./01_python_basics_and_core.md) | Data types, mutability, `==` vs `is`, LEGB scope, f-strings, comprehensions, `*args`/`**kwargs`, decorators |
| 02 | [OOP & Classes](./02_oop_and_classes.md) | Classes, `self`, inheritance, `super()`, dunder methods, encapsulation, `@property`, polymorphism, MRO |
| 03 | [Data Structures & Algorithms](./03_data_structures_and_algorithms.md) | Lists, dicts, sets, `collections` module, stack/queue, sorting, binary search, two pointers, sliding window |
| 04 | [File Handling & Exceptions](./04_file_handling_and_exceptions.md) | File read/write, `with` statement, `try/except/else/finally`, custom exceptions, `pathlib` |
| 05 | [Functional Programming](./05_functional_programming.md) | Generators, `yield`, `functools`, `itertools`, closures, `zip`/`enumerate`/`map`/`filter` |
| 06 | [Modules, Packages & Env](./06_modules_packages_and_env.md) | Imports, `__init__.py`, virtual environments, `pip`, `os`/`sys`/`json`/`re`, `.env` files |
| 07 | [Databases & Web](./07_databases_and_web.md) | `sqlite3`, SQLAlchemy ORM, REST APIs with Flask/FastAPI, `requests`, `httpx` |
| 08 | [Concurrency & Testing](./08_concurrency_and_testing.md) | GIL, threading, multiprocessing, `asyncio`, `pytest`, mocking |

---

## 🎯 Interview Focus Areas

### Frequently Asked in Rounds 1 & 2:
- Python data types and difference between mutable/immutable
- OOP concepts: inheritance, polymorphism, encapsulation
- List comprehensions, generators, decorators
- Exception handling patterns
- Basic sorting and searching

### Frequently Asked in Technical Rounds:
- Write a REST API using Flask or FastAPI
- Explain threading vs multiprocessing vs asyncio
- SQLAlchemy ORM usage
- Writing unit tests with pytest
- File handling and context managers

---

## 📝 Quick Tips
- Always use `is` for `None` checks (`if x is None:`)
- Use `with` statement for all file operations
- Prefer `pathlib.Path` over `os.path` in modern Python
- Use `pytest` fixtures to avoid test code duplication
- For I/O-bound: use threads; for CPU-bound: use processes
