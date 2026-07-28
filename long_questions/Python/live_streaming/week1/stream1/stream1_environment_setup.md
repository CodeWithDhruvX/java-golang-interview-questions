# Stream 1: Environment Setup & Core Concepts - Complete Script

**Duration**: 2 hours  
**Project**: Hello World with Advanced Features  
**Repository**: `long_questions/Python/core-python/basics/`

---

## Stream Overview

**Target Audience**: Beginners to Python development  
**Prerequisites**: None (complete beginners welcome)  
**Learning Outcomes**:
- Install Python 3.11+ and set up virtual environments
- Configure VS Code/PyCharm for Python development
- Understand Python execution model and REPL
- Build an advanced Hello World project demonstrating core concepts
- Learn Python syntax basics and best practices

---

## Timeline Breakdown

| Segment | Duration | Description |
|---------|----------|-------------|
| Introduction | 10 mins | Welcome, Python overview, objectives |
| Python Installation | 20 mins | Python 3.11+ installation, verification |
| Virtual Environments | 15 mins | venv/conda setup and management |
| IDE Configuration | 20 mins | VS Code/PyCharm setup, extensions |
| Python Execution & REPL | 15 mins | Python execution model, REPL usage |
| Hello World Project | 25 mins | Complete project implementation |
| Testing & Demo | 15 mins | Run project, demonstrate features |
| Q&A | 15 mins | Viewer questions and clarifications |
| Summary | 5 mins | Recap, homework, next stream preview |

---

## Complete Script

### 1. Introduction (10 minutes)

**[0:00-0:10] Welcome & Python Overview**

"Welcome to Stream 1 of our Python Live Coding series! Today we're setting up your Python development environment from scratch and building your first advanced Hello World project.

**Today's Objectives:**
- Install Python 3.11+ on your system
- Set up virtual environments for project isolation
- Configure a professional IDE (VS Code or PyCharm)
- Understand how Python executes code
- Build an advanced Hello World demonstrating core Python concepts

**Prerequisites Check:**
- No prior programming experience required
- Administrator access for installation
- Internet connection for downloads
- 5GB free disk space

**What You'll Achieve:**
By the end of this stream, you'll have a fully functional Python development environment and understand the fundamentals of Python syntax. You'll build a project that goes beyond simple print statements to demonstrate variables, data types, functions, and user interaction.

**Why Python Matters:**
Python is one of the most popular programming languages in the world, used by:
- Google, Netflix, Instagram for web development
- NASA, SpaceX for scientific computing
- JPMorgan, Goldman Sachs for financial analysis
- AI/ML researchers for machine learning

Understanding Python fundamentals opens doors to:
- Web development (Django, Flask, FastAPI)
- Data science (Pandas, NumPy)
- Machine learning (TensorFlow, PyTorch)
- Automation and scripting
- DevOps and cloud engineering

Let's get started!"

---

### 2. Python Installation (20 minutes)

**[0:10-0:30] Installing Python 3.11+**

"First, let's install Python on your system. We'll use Python 3.11+ for the latest features and performance improvements.

**Step 1: Download Python**

For Windows:
- Visit python.org/downloads
- Download Python 3.11.x or 3.12.x
- **CRITICAL**: Check 'Add Python to PATH' during installation

For macOS:
- Visit python.org/downloads
- Download macOS installer (pkg or universal2)
- Or use Homebrew: `brew install python@3.11`

For Linux:
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install python3.11 python3.11-venv

# Fedora
sudo dnf install python3.11
```

**Step 2: Verify Installation**

Open terminal/command prompt and run:

```bash
python --version
# or
python3 --version

# Expected output: Python 3.11.x or 3.12.x
```

If you see an error, Python might not be in your PATH. We'll fix that.

**Step 3: Verify pip (Package Manager)**

```bash
pip --version
# or
pip3 --version

# Expected output: pip 23.x.x from ...
```

**Common Installation Issues:**

1. **'python' command not found**:
   - Try `python3` instead
   - Add Python to PATH manually
   - Windows: Add to System Environment Variables

2. **Permission denied**:
   - Run installer as administrator
   - Use user-space installation on Linux

3. **Multiple Python versions**:
   - Use py launcher on Windows: `py --list`
   - Use version managers: pyenv (macOS/Linux), pywin (Windows)

**Live Demo - Installation Verification:**

```bash
# Check Python version
python --version

# Check pip
pip --version

# Check Python location
where python  # Windows
which python  # macOS/Linux

# Test Python in interactive mode
python
>>> print("Hello, Python!")
>>> exit()
```

**Best Practices:**
- Always use the latest stable version (3.11+)
- Keep Python updated for security patches
- Use virtual environments for each project
- Never use system Python for development"

---

### 3. Virtual Environments (15 minutes)

**[0:30-0:45] Setting Up Virtual Environments**

"Virtual environments are essential for Python development. They isolate project dependencies and prevent conflicts between different projects.

**What are Virtual Environments?**
- Isolated Python environments for each project
- Separate package installations per project
- Prevents dependency conflicts
- Makes project reproducible

**Creating a Virtual Environment:**

Using built-in `venv` (recommended):

```bash
# Navigate to your project directory
cd path/to/your/project

# Create virtual environment
python -m venv venv

# or specify Python version
python3.11 -m venv venv

# Activate virtual environment
# Windows:
venv\Scripts\activate

# macOS/Linux:
source venv/bin/activate
```

**Using conda (alternative for data science):**

```bash
# Create conda environment
conda create --name myproject python=3.11

# Activate conda environment
conda activate myproject
```

**Live Demo - Virtual Environment Setup:**

```bash
# Create project directory
mkdir python_hello_world
cd python_hello_world

# Create virtual environment
python -m venv venv

# Activate (Windows)
venv\Scripts\activate

# Verify activation (should see (venv) in prompt)
python --version

# Install a package
pip install requests

# Check installed packages
pip list

# Deactivate
deactivate
```

**Virtual Environment Best Practices:**

1. **Always use virtual environments** for projects
2. **Never commit venv/ folder** to Git
3. **Use requirements.txt** for dependency tracking
4. **Name environments descriptively** (e.g., `projectname-dev`)

**Creating requirements.txt:**

```bash
# Export current dependencies
pip freeze > requirements.txt

# Install from requirements.txt
pip install -r requirements.txt
```

**requirements.txt example:**

```
requests==2.31.0
numpy==1.24.3
pandas==2.0.3
```

**Git .gitignore for Python:**

```
# Virtual environments
venv/
env/
.venv/

# Python cache
__pycache__/
*.py[cod]
*$py.class

# IDE
.vscode/
.idea/
*.swp
```

**Common Virtual Environment Issues:**

1. **Activation script not found**:
   - Check venv folder exists
   - Verify Python installation
   - Recreate virtual environment

2. **Package not found after activation**:
   - Verify activation worked (check prompt)
   - Install package in active environment
   - Check pip is from venv: `which pip`"

---

### 4. IDE Configuration (20 minutes)

**[0:45-1:05] Setting Up VS Code/PyCharm**

"A good IDE significantly boosts productivity. We'll set up VS Code (free, lightweight) or PyCharm (feature-rich, free Community Edition).

**Option 1: VS Code Setup**

**Step 1: Install VS Code**
- Download from code.visualstudio.com
- Install with default settings

**Step 2: Install Python Extension**
- Open VS Code
- Go to Extensions (Ctrl+Shift+X)
- Search "Python" by Microsoft
- Install extension

**Step 3: Configure Python Interpreter**

```json
// settings.json
{
    "python.defaultInterpreterPath": "${workspaceFolder}/venv/Scripts/python.exe",
    "python.linting.enabled": true,
    "python.linting.pylintEnabled": true,
    "python.formatting.provider": "black",
    "python.testing.pytestEnabled": true
}
```

**Step 4: Install Recommended Extensions**

```
- Python (Microsoft)
- Pylance (Microsoft)
- Black Formatter (Microsoft)
- Python Test Explorer (LittleFoxTeam)
- GitLens (GitKraken)
- Material Icon Theme (Philipp Kief)
```

**Step 5: Configure VS Code Settings**

```json
{
    "editor.formatOnSave": true,
    "editor.tabSize": 4,
    "editor.insertSpaces": true,
    "python.analysis.autoImportCompletions": true,
    "python.analysis.typeCheckingMode": "basic"
}
```

**Option 2: PyCharm Setup**

**Step 1: Install PyCharm Community**
- Download from jetbrains.com/pycharm
- Install Community Edition (free)

**Step 2: Configure Project Interpreter**
- File → Settings → Project → Python Interpreter
- Click "Add Interpreter" → "Add Local Interpreter"
- Select "Existing environment"
- Browse to your venv/Scripts/python.exe

**Step 3: Configure Code Style**
- File → Settings → Editor → Code Style → Python
- Set indentation to 4 spaces
- Enable "Show line numbers"

**Live Demo - VS Code Setup:**

```bash
# Open VS Code in project directory
code .

# Create first Python file
# File → New File → hello.py

# Type: print("Hello, VS Code!")
# Save and run with F5 or play button
```

**VS Code Keyboard Shortcuts:**

```
Ctrl+`      - Toggle terminal
Ctrl+Shift+P - Command palette
Ctrl+P      - Quick open file
Ctrl+D      - Duplicate line
Alt+Up/Down - Move line
Ctrl+/      - Comment/uncomment
Ctrl+S      - Save
F5          - Run/debug
```

**PyCharm Keyboard Shortcuts:**

```
Alt+F12     - Toggle terminal
Ctrl+Shift+A - Find action
Ctrl+Shift+N - Quick open file
Ctrl+D      - Duplicate line
Alt+Shift+Up/Down - Move line
Ctrl+/      - Comment/uncomment
Ctrl+S      - Save
Shift+F10   - Run
```

**IDE Best Practices:**

1. **Use integrated terminal** for consistency
2. **Enable auto-format on save** (Black)
3. **Configure linting** (Pylint/Flake8)
4. **Use code snippets** for common patterns
5. **Customize theme** for eye comfort

**Live Demo - First Python File in IDE:**

```python
# hello.py
"""
This is our first Python file
Demonstrating basic Python syntax
"""

# Simple print statement
print("Hello, World!")

# String with f-string (Python 3.6+)
name = "Python"
print(f"Hello, {name}!")

# Multi-line string
message = """
This is a multi-line string.
It can span multiple lines.
"""
print(message)
```

**Run the file:**
- VS Code: Press F5 or click play button
- PyCharm: Right-click → Run 'hello'
- Terminal: `python hello.py`"

---

### 5. Python Execution & REPL (15 minutes)

**[1:05-1:20] Understanding Python Execution**

"Let's understand how Python executes code and how to use the REPL (Read-Eval-Print Loop) for interactive programming.

**Python Execution Model:**

1. **Source Code (.py files)** → Python interpreter
2. **Compilation to bytecode** (.pyc files)
3. **Python Virtual Machine (PVM)** executes bytecode
4. **Output** to console

**Running Python Code:**

**Method 1: Script execution**
```bash
python script.py
python3 script.py
```

**Method 2: Interactive REPL**
```bash
python
>>> print("Hello")
Hello
>>> 2 + 2
4
>>> exit()
```

**Method 3: IPython (enhanced REPL)**
```bash
pip install ipython
ipython
```

**Live Demo - REPL Usage:**

```python
# Start Python REPL
python

# Basic arithmetic
>>> 2 + 3
5
>>> 10 / 3
3.3333333333333335
>>> 10 // 3  # Floor division
3
>>> 10 % 3   # Modulus
1
>>> 2 ** 10  # Exponentiation
1024

# Variables
>>> x = 10
>>> y = 20
>>> x + y
30

# Strings
>>> name = "Python"
>>> name.upper()
'PYTHON'
>>> name * 3
'PythonPythonPython'

# Lists
>>> numbers = [1, 2, 3, 4, 5]
>>> numbers[0]
1
>>> numbers[-1]
5
>>> numbers[1:3]
[2, 3]

# Built-in functions
>>> len(numbers)
5
>>> sum(numbers)
15
>>> max(numbers)
5
>>> min(numbers)
1

# Type checking
>>> type(x)
<class 'int'>
>>> type(name)
<class 'str'>

# Help system
>>> help(print)
>>> help(len)

# Exit REPL
>>> exit()
```

**Python Script Structure:**

```python
#!/usr/bin/env python3
"""
This is a docstring - describes the file
Author: Your Name
Date: Current date
"""

# Imports go at the top
import sys
import os

# Constants
PI = 3.14159

# Functions
def greet(name):
    """Greet the user"""
    return f"Hello, {name}!"

# Main execution
if __name__ == "__main__":
    # This code runs only when script is executed directly
    print(greet("World"))
```

**Understanding `if __name__ == "__main__":`:**

```python
# module.py
def function():
    print("Function called")

if __name__ == "__main__":
    # Runs when executed: python module.py
    function()
    print("Script executed directly")
else:
    # Runs when imported: import module
    print("Module imported")
```

**Live Demo - Script vs Import:**

```bash
# Create script.py
python script.py  # Runs main block

# In Python REPL
>>> import script  # Does NOT run main block
>>> script.function()  # Can call functions
```

**Python Execution Best Practices:**

1. **Use shebang** for executable scripts: `#!/usr/bin/env python3`
2. **Add docstrings** to modules and functions
3. **Use `if __name__ == "__main__":`** for main logic
4. **Keep imports at the top** of files
5. **Use meaningful variable names**

**Common Execution Issues:**

1. **IndentationError**: Python uses indentation, not braces
2. **SyntaxError**: Check for typos and missing colons
3. **NameError**: Variable used before definition
4. **ModuleNotFoundError**: Package not installed

**Python -c for one-liners:**

```bash
python -c "print('Hello from command line')"
python -c "import math; print(math.pi)"
```

**Python -m for running modules:**

```bash
python -m http.server 8000  # Start simple HTTP server
python -m pip install package  # Run pip as module
```"

---

### 6. Hello World Project (25 minutes)

**[1:20-1:45] Building Advanced Hello World**

"Now let's build an advanced Hello World project that demonstrates core Python concepts beyond simple print statements.

**Project Features:**
- User input and output
- Variables and data types
- String formatting and manipulation
- Functions and control flow
- Error handling
- File operations

**Step 1: Create Project Structure**

```bash
mkdir hello_world_project
cd hello_world_project
python -m venv venv
venv\Scripts\activate  # Windows
# source venv/bin/activate  # macOS/Linux
```

**Step 2: Create main.py**

```python
#!/usr/bin/env python3
"""
Advanced Hello World Project
Demonstrates core Python concepts
Author: Your Name
"""

import random
import datetime
from typing import Optional

# Constants
VERSION = "1.0.0"
GREETINGS = [
    "Hello", "Hi", "Hey", "Greetings", "Welcome",
    "Bonjour", "Hola", "Ciao", "Namaste"
]

def get_current_time() -> str:
    """Get current formatted time"""
    now = datetime.datetime.now()
    return now.strftime("%Y-%m-%d %H:%M:%S")

def get_random_greeting() -> str:
    """Get a random greeting from the list"""
    return random.choice(GREETINGS)

def greet_user(name: str, greeting: Optional[str] = None) -> str:
    """
    Greet the user with a personalized message
    
    Args:
        name: User's name
        greeting: Optional custom greeting
    
    Returns:
        Formatted greeting message
    """
    if greeting is None:
        greeting = get_random_greeting()
    
    return f"{greeting}, {name}! Welcome to Python programming."

def get_user_name() -> str:
    """Get user name with input validation"""
    while True:
        try:
            name = input("Enter your name: ").strip()
            if not name:
                print("Name cannot be empty. Please try again.")
                continue
            if len(name) < 2:
                print("Name must be at least 2 characters. Please try again.")
                continue
            return name
        except KeyboardInterrupt:
            print("\nOperation cancelled by user.", end=" ")
            return "Guest"
        except Exception as e:
            print(f"An error occurred: {e}")
            return "Guest"

def display_banner():
    """Display a welcome banner"""
    banner = """
╔════════════════════════════════════════╗
║     Advanced Hello World - Python      ║
║           Version {}              ║
╚════════════════════════════════════════╝
    """.format(VERSION)
    print(banner)

def calculate_age(birth_year: int) -> int:
    """Calculate age from birth year"""
    current_year = datetime.datetime.now().year
    age = current_year - birth_year
    return max(0, age)  # Ensure non-negative

def get_user_info():
    """Collect and display user information"""
    name = get_user_name()
    
    print(f"\nNice to meet you, {name}!")
    
    try:
        birth_year = int(input("Enter your birth year (YYYY): "))
        age = calculate_age(birth_year)
        print(f"You are {age} years old.")
    except ValueError:
        print("Invalid year format. Skipping age calculation.")
    except Exception as e:
        print(f"Error calculating age: {e}")
    
    return name

def save_greeting_to_file(name: str, message: str):
    """Save greeting to a file"""
    filename = "greeting_log.txt"
    timestamp = get_current_time()
    
    try:
        with open(filename, "a", encoding="utf-8") as file:
            file.write(f"[{timestamp}] {message}\n")
        print(f"Greeting saved to {filename}")
    except IOError as e:
        print(f"Error saving to file: {e}")

def main():
    """Main program execution"""
    display_banner()
    
    print(f"Current time: {get_current_time()}")
    print(f"Available greetings: {len(GREETINGS)}")
    print()
    
    # Get user information
    name = get_user_info()
    
    # Generate and display greeting
    greeting = greet_user(name)
    print(f"\n{greeting}")
    
    # Display some Python facts
    print("\n" + "="*50)
    print("Python Facts:")
    print("="*50)
    print(f"- Python version: {random.__version__}")
    print(f"- Guido van Rossum created Python in 1991")
    print(f"- Python is named after Monty Python")
    print(f"- Python supports multiple programming paradigms")
    
    # Save greeting to file
    save_greeting_to_file(name, greeting)
    
    # Farewell message
    print("\n" + "="*50)
    print("Thank you for using Advanced Hello World!")
    print("Happy coding with Python! 🐍")
    print("="*50)

if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("\n\nProgram interrupted by user. Goodbye!")
    except Exception as e:
        print(f"\nAn unexpected error occurred: {e}")
```

**Step 3: Create requirements.txt**

```
# No external packages needed for this project
# Using only Python standard library
```

**Step 4: Create README.md**

```markdown
# Advanced Hello World Project

A comprehensive Hello World project demonstrating core Python concepts.

## Features
- User input validation
- Random greetings
- Age calculation
- File operations
- Error handling
- Type hints

## Installation
```bash
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
```

## Usage
```bash
python main.py
```

## Project Structure
```
hello_world_project/
├── main.py              # Main application
├── requirements.txt     # Dependencies
├── README.md           # Documentation
├── greeting_log.txt    # Generated log file
└── venv/               # Virtual environment (not committed)
```
```

**Step 5: Create .gitignore**

```
# Virtual environment
venv/
env/
.venv/

# Python cache
__pycache__/
*.py[cod]
*$py.class

# IDE
.vscode/
.idea/
*.swp

# Generated files
greeting_log.txt

# OS
.DS_Store
Thumbs.db
```

**Live Demo - Running the Project:**

```bash
# Run the project
python main.py

# Expected output:
# Welcome banner
# User input prompts
# Personalized greeting
# Python facts
# File save confirmation
```

**Testing Different Scenarios:**

1. **Normal input**: Enter valid name and year
2. **Empty name**: Test validation
3. **Short name**: Test minimum length validation
4. **Invalid year**: Test error handling
5. **Keyboard interrupt**: Test Ctrl+C handling

**Code Explanation - Key Concepts:**

1. **Type Hints**: `def greet_user(name: str) -> str`
   - Improves code documentation
   - Enables IDE autocomplete
   - Supports static analysis tools

2. **Optional Parameters**: `greeting: Optional[str] = None`
   - Provides default values
   - Makes parameters optional

3. **Context Managers**: `with open(filename, "a") as file:`
   - Automatic resource cleanup
   - Exception-safe file handling

4. **F-strings**: `f"{greeting}, {name}!"`
   - Modern string formatting
   - Readable and efficient

5. **Error Handling**: `try-except` blocks
   - Graceful error recovery
   - User-friendly error messages

6. **Constants**: Uppercase variable names
   - Convention for immutable values
   - Improves code readability"

---

### 7. Testing & Demo (15 minutes)

**[1:45-2:00] Running and Demonstrating**

"Let's run our project and demonstrate all the features we've built.

**Live Demo - Complete Project Run:**

```bash
# Ensure virtual environment is activated
venv\Scripts\activate

# Run the project
python main.py

# Test with various inputs:
# 1. Normal name: "Alice"
# 2. Empty input: just press Enter
# 3. Short name: "A"
# 4. Valid year: "1990"
# 5. Invalid year: "abc"
```

**Demonstrating File Operations:**

```bash
# Check if log file was created
cat greeting_log.txt  # Linux/macOS
type greeting_log.txt  # Windows

# Run multiple times to see log accumulation
python main.py
python main.py
python main.py

# View log file again
type greeting_log.txt
```

**Demonstrating Python REPL Integration:**

```python
# Import our module in REPL
python
>>> import main
>>> main.greet_user("Alice")
>>> main.get_random_greeting()
>>> main.VERSION
>>> exit()
```

**Demonstrating IDE Features:**

1. **Code Completion**: Start typing and let IDE suggest
2. **Go to Definition**: Ctrl+Click on function names
3. **Rename Symbol**: F2 to rename variables
4. **Format Document**: Shift+Alt+F to format code
5. **Run Tests**: Use test explorer (if using pytest)

**Performance Check:**

```python
# Add timing to main.py
import time

start_time = time.time()
main()
end_time = time.time()

print(f"\nExecution time: {end_time - start_time:.2f} seconds")
```

**Code Quality Checks:**

```bash
# Install linting tools
pip install pylint black flake8

# Run pylint
pylint main.py

# Run black formatter
black main.py

# Run flake8
flake8 main.py
```

**Demonstrating Different Python Versions:**

```bash
# Check Python version
python --version

# Run with specific version
python3.11 main.py

# Check compatibility
python -m py_compile main.py
```

**Common Issues and Solutions:**

1. **Module not found error**:
   - Ensure virtual environment is activated
   - Check Python interpreter in IDE

2. **Permission denied writing file**:
   - Check directory permissions
   - Run with appropriate permissions

3. **Encoding issues**:
   - Use UTF-8 encoding in file operations
   - Specify encoding in `open()` call

**Best Practices Demonstrated:**

1. **Virtual environment isolation**
2. **Type hints for clarity**
3. **Error handling for robustness**
4. **Docstrings for documentation**
5. **Constants for configuration**
6. **Functions for modularity**
7. **Main guard for importability**"

---

### 8. Q&A Session (15 minutes)

**[2:00-2:15] Viewer Questions**

"Let's address common questions about Python setup and fundamentals.

**Q1: Why use virtual environments instead of system Python?**
A: Virtual environments prevent dependency conflicts between projects. Different projects may require different versions of the same package. System Python should remain clean and stable.

**Q2: Should I use VS Code or PyCharm?**
A: Both are excellent choices. VS Code is lightweight and highly customizable. PyCharm has more built-in features for Python specifically. Start with VS Code for simplicity, try PyCharm for advanced features.

**Q3: What's the difference between `python` and `python3`?**
A: On some systems, `python` refers to Python 2.x (deprecated), while `python3` refers to Python 3.x. On modern systems, `python` often points to Python 3. Always use `python3` or configure aliases for clarity.

**Q4: Do I need to learn Python 2?**
A: No. Python 2 reached end-of-life in 2020. Focus entirely on Python 3.6+ (preferably 3.11+ for latest features).

**Q5: How do I keep Python packages updated?**
A: Use `pip install --upgrade package_name` or `pip install --upgrade -r requirements.txt` to update all packages. Regular updates ensure security patches and new features.

**Q6: What's the difference between pip and conda?**
A: pip is Python's default package manager. conda is a cross-language package manager (popular in data science). Use pip for general Python, conda for data science workflows.

**Q7: Why does Python use indentation instead of braces?**
A: Python's design philosophy emphasizes readability. Mandatory indentation makes code structure visually apparent and reduces braces-related errors. It's a feature, not a limitation.

**Q8: How do I debug Python code?**
A: Use IDE debuggers (VS Code debugger, PyCharm debugger), print statements for simple cases, or `pdb` (Python debugger) for command-line debugging. We'll cover debugging in detail in future streams.

**Q9: What are .pyc files?**
A: .pyc files are compiled bytecode files created by Python for faster execution. Python automatically compiles .py files to .pyc and caches them. You can safely delete them; Python will regenerate them.

**Q10: How do I share my Python project with others?**
A: Include requirements.txt, README.md, and .gitignore. Use Git for version control. Others can recreate your environment with `pip install -r requirements.txt`."

---

### 9. Summary & Next Steps (5 minutes)

**[2:15-2:20] Recap and Preview**

**Today's Achievements:**
✅ Installed Python 3.11+ on your system
✅ Set up virtual environments for project isolation
✅ Configured VS Code/PyCharm for Python development
✅ Understood Python execution model and REPL
✅ Built an advanced Hello World project with:
   - User input and validation
   - Variables and data types
   - Functions and control flow
   - Error handling
   - File operations
   - Type hints and documentation

**Homework Assignments:**

1. **Enhance the Hello World project**:
   - Add more greeting languages
   - Implement age-based messages (child, teen, adult, senior)
   - Add color output using colorama package
   - Create a configuration file (config.json)

2. **Practice Python basics**:
   - Write a calculator that performs basic operations
   - Create a program to convert temperature units
   - Build a simple quiz game with score tracking

3. **Explore Python standard library**:
   - Read about `os`, `sys`, `datetime`, `random` modules
   - Try using `math` module for calculations
   - Experiment with `json` module for data handling

**Next Stream Preview:**

Stream 2: Object-Oriented Programming Deep Dive
- Classes, objects, and constructors
- Inheritance, polymorphism, encapsulation
- Abstract classes vs interfaces (ABC module)
- **Project**: Employee Management System

**Resources:**
- Official Python documentation: docs.python.org
- Python tutorial: docs.python.org/3/tutorial
- VS Code Python extension: code.visualstudio.com/docs/python
- PyCharm guide: jetbrains.com/pycharm

**Community:**
- Join our Discord for Q&A
- GitHub repository for code: [link]
- Twitter for updates: @[handle]

**Thank you for joining Stream 1!**
You now have a solid foundation for Python development. See you in the next stream where we'll dive deep into Object-Oriented Programming!

Happy coding! 🐍"

---

## Additional Resources

### Python Installation Links
- Windows: python.org/downloads/windows
- macOS: python.org/downloads/macos
- Linux: Use package manager (apt, dnf, yum)

### IDE Download Links
- VS Code: code.visualstudio.com
- PyCharm Community: jetbrains.com/pycharm/download

### Useful Python Packages
```bash
# Development tools
pip install black pylint flake8 mypy

# Data science (preview)
pip install numpy pandas matplotlib

# Web development (preview)
pip install flask django requests

# Testing (preview)
pip install pytest pytest-cov
```

### Learning Resources
- Python.org Official Tutorial
- Real Python (realpython.com)
- Python for Beginners (youtube.com/@ProgrammingwithMosh)
- Corey Schafer (youtube.com/@coreyms)

### Troubleshooting Common Issues

**Issue: pip not found**
```bash
python -m ensurepip --upgrade
python -m pip install --upgrade pip
```

**Issue: SSL certificate errors**
```bash
pip install --trusted-host pypi.org --trusted-host files.pythonhosted.org package_name
```

**Issue: Permission denied**
```bash
# Use user directory
pip install --user package_name

# Or use virtual environment (recommended)
python -m venv venv
```

---

## Stream Checklist

- [x] Introduction and objectives
- [x] Python installation demonstration
- [x] Virtual environment setup
- [x] IDE configuration (VS Code/PyCharm)
- [x] Python execution and REPL
- [x] Hello World project build
- [x] Testing and demonstration
- [x] Q&A session
- [x] Summary and next steps

**Stream Status**: ✅ Complete
**Duration**: 2 hours
**Project Files**: main.py, requirements.txt, README.md, .gitignore
