# 📘 06 — Modules, Packages & Environment
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- Modules, packages, `__init__.py`
- Import styles and `__name__ == "__main__"`
- Virtual environments
- `pip` and `requirements.txt`
- `sys.path` and how Python finds modules
- Standard library essentials: `os`, `sys`, `datetime`, `json`, `re`

---

## ❓ Most Asked Questions

### Q1. What is a module and how do Python imports work?

```python
# A module is simply a .py file.
# A package is a directory containing __init__.py (or just a dir in Python 3.3+)

# -- mypackage/
# --   __init__.py    (makes it a package)
# --   utils.py
# --   models/
# --     __init__.py
# --     user.py

# Import styles:
import os                          # import whole module
from os import path, getcwd        # import specific names
from os.path import join, exists   # from submodule
import numpy as np                 # alias

# Absolute vs relative imports:
# From inside mypackage/models/user.py:
from mypackage.utils import helper_fn   # absolute import ✅
from ..utils import helper_fn           # relative import ✅ (2 dots = up 2 levels)
from .base import BaseModel             # relative (same package)

# __init__.py controls what 'from package import *' exports
# mypackage/__init__.py:
# from .utils import format_date, slugify
# __all__ = ["format_date", "slugify"]   # explicit public API

# How Python finds modules (sys.path):
import sys
print(sys.path)
# ['', '/usr/lib/python3.12', '/usr/lib/python3.12/lib-dynload', ...]
# '' = current directory (searched first!)

# Check module location
import json
print(json.__file__)   # /usr/lib/python3.12/json/__init__.py

# Reload a modified module (useful in interactive sessions)
import importlib
importlib.reload(my_module)

# __name__ == "__main__": run code only when script is executed directly
# Not when imported as a module
if __name__ == "__main__":
    print("Script executed directly")
    main()
```

---

### Q2. How do you manage virtual environments and dependencies?

```python
# Virtual environment: isolated Python environment per project
# Prevents dependency conflicts between projects

# Create virtual environment
# python -m venv venv           # creates venv/ directory
# python -m venv .venv          # convention: .venv (hidden)

# Activate:
# Windows:  venv\Scripts\activate
# macOS/Linux: source venv/bin/activate

# Deactivate:
# deactivate

# After activation, pip installs to the venv — not globally
# pip install requests Flask SQLAlchemy

# requirements.txt: pin exact versions for reproducibility
# Generate current environment:
# pip freeze > requirements.txt

# requirements.txt example:
# Flask==3.0.2
# SQLAlchemy==2.0.28
# requests==2.31.0
# python-dotenv==1.0.1

# Install from requirements.txt:
# pip install -r requirements.txt

# pyproject.toml (modern, PEP 517/518)
# [project]
# name = "myapp"
# version = "1.0.0"
# dependencies = ["flask>=3.0", "sqlalchemy>=2.0"]

# Useful pip commands:
# pip install package            # install latest
# pip install package==1.2.3    # specific version
# pip install "package>=1.0,<2" # range
# pip install -e .              # editable install for dev
# pip list                      # list installed packages
# pip show flask                # details about a package
# pip uninstall package         # remove

# Check for security vulnerabilities:
# pip audit

# Python version management:
# pyenv — manage multiple Python versions
# conda — popular in data science
```

---

### Q3. What are the most essential standard library modules?

```python
import os, sys, datetime, json, re, math, random, time

# os — operating system interface
os.getcwd()                      # current working directory
os.listdir(".")                  # list files in directory
os.path.join("a", "b", "c.txt") # "/a/b/c.txt"
os.path.exists("file.txt")       # check file exists
os.environ.get("PATH", "")       # read environment variable
os.makedirs("a/b/c", exist_ok=True)  # create dirs recursively
os.remove("file.txt")            # delete file

# sys — system-specific parameters
sys.argv          # command-line arguments (list of strings)
sys.exit(0)       # exit with return code
sys.stdin         # standard input stream
sys.stdout        # standard output stream
sys.version       # Python version string

# datetime — dates and times
from datetime import datetime, date, timedelta

now = datetime.now()            # current date and time
today = date.today()            # current date only
tomorrow = today + timedelta(days=1)

formatted = now.strftime("%Y-%m-%d %H:%M:%S")   # "2024-08-15 10:30:00"
parsed = datetime.strptime("15-08-2024", "%d-%m-%Y")

# json — JSON serialization
data = {"name": "Alice", "scores": [95, 87, 92]}
json_str = json.dumps(data, indent=2)      # dict → JSON string
back = json.loads(json_str)                # JSON string → dict

with open("data.json", "w") as f:
    json.dump(data, f, indent=2)           # dict → JSON file

with open("data.json") as f:
    loaded = json.load(f)                  # JSON file → dict

# re — regular expressions
import re
pattern = r"\b\d{10}\b"                    # 10-digit number
re.search(pattern, "Call 9876543210 now")  # Match object
re.findall(r"\d+", "123 abc 456 def")     # ['123', '456']
re.sub(r"\s+", "_", "hello world")        # "hello_world"
re.split(r"[,;|]", "a,b;c|d")             # ['a','b','c','d']

# Compile for repeated use
phone_re = re.compile(r"\b\d{10}\b")
phone_re.findall("Call 9876543210 or 8123456789")

# math and random
import math
math.sqrt(16)    # 4.0
math.ceil(3.1)   # 4
math.floor(3.9)  # 3
math.pi          # 3.14159...
math.log(100, 10)  # 2.0

import random
random.randint(1, 100)              # random int from 1 to 100
random.choice(["a", "b", "c"])     # random element
random.shuffle([1, 2, 3, 4, 5])    # in-place shuffle
random.sample(range(100), 10)      # 10 unique random items
```

---

### Q4. How do you work with environment variables and `.env` files?

```python
import os
from pathlib import Path

# Reading environment variables
DB_HOST = os.environ.get("DB_HOST", "localhost")   # with default
DB_PORT = int(os.environ.get("DB_PORT", "5432"))   # convert type
SECRET_KEY = os.environ["SECRET_KEY"]              # raises KeyError if missing

# Setting environment variables (runtime only)
os.environ["TEMP_VAR"] = "hello"

# .env file (never commit to git!)
# DB_HOST=localhost
# DB_PORT=5432
# SECRET_KEY=abc123super-secret

# python-dotenv: load .env file
from dotenv import load_dotenv

load_dotenv()               # loads .env from current dir
load_dotenv(".env.local")   # specific file
load_dotenv(override=True)  # override existing env vars

# After load_dotenv, os.environ.get() works as normal
DB_HOST = os.environ.get("DB_HOST")

# Configuration class pattern (common in Flask/Django):
class Config:
    DEBUG       = os.environ.get("DEBUG", "False").lower() == "true"
    DB_URL      = os.environ.get("DATABASE_URL", "sqlite:///dev.db")
    SECRET_KEY  = os.environ.get("SECRET_KEY", "dev-secret-change-in-prod")
    REDIS_URL   = os.environ.get("REDIS_URL", "redis://localhost:6379")
    MAX_WORKERS = int(os.environ.get("MAX_WORKERS", "4"))

class DevelopmentConfig(Config):
    DEBUG = True
    DB_URL = "sqlite:///dev.db"

class ProductionConfig(Config):
    DEBUG = False

# Using dataclasses for typed config
from dataclasses import dataclass, field

@dataclass
class AppConfig:
    host: str = field(default_factory=lambda: os.environ.get("HOST", "0.0.0.0"))
    port: int = field(default_factory=lambda: int(os.environ.get("PORT", "8000")))
    debug: bool = field(default_factory=lambda: os.environ.get("DEBUG") == "1")

config = AppConfig()
```
