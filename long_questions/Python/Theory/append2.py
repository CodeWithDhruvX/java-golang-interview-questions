import os

content = """
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
`    file.write('New log entry\\n')`

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

Instead of hardcoding `folder + '/' + filename`, which crashes on Windows, I use `os.path.join('folder', 'filename')`. If I'm on Linux, it outputs `folder/filename`. On Windows, it outputs `folder\\filename`.

In modern code, `pathlib` handles this even more elegantly with the slash operator: `Path('folder') / 'filename'`."

#### Indepth
`os.path.join` has a slightly confusing edge case: if any component is an absolute path (starts with `/` on Linux or `C:\\` on Windows), all previous components are discarded and joining continues from the absolute path component.

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
On Windows: `.venv\\Scripts\\activate`

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
"""

file_path = r'g:\My Drive\All Documents\java-golang-interview-questions\long_questions\Python\Theory\01_Basics.md'

with open(file_path, 'a', encoding='utf-8') as f:
    f.write(content)
