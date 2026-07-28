# Advanced Hello World Project

A comprehensive Hello World project demonstrating core Python concepts from Stream 1 of the Python Live Coding series.

## Features

- User input validation
- Random greetings from multiple languages
- Age calculation from birth year
- File operations for logging greetings
- Error handling and graceful degradation
- Type hints for better code documentation
- Docstrings for function documentation
- Main guard pattern for importability

## Installation

### Prerequisites
- Python 3.11 or higher
- Virtual environment (recommended)

### Setup

1. Clone or download this project
2. Create a virtual environment:
```bash
python -m venv venv
```

3. Activate the virtual environment:

**Windows:**
```bash
venv\Scripts\activate
```

**macOS/Linux:**
```bash
source venv/bin/activate
```

4. Install dependencies (optional - uses standard library only):
```bash
pip install -r requirements.txt
```

## Usage

Run the project:
```bash
python main.py
```

### Example Output

```
╔════════════════════════════════════════╗
║     Advanced Hello World - Python      ║
║           Version 1.0.0              ║
╚════════════════════════════════════════╝

Current time: 2024-01-15 10:30:45
Available greetings: 9

Enter your name: Alice

Nice to meet you, Alice!
Enter your birth year (YYYY): 1990
You are 34 years old.

Bonjour, Alice! Welcome to Python programming.

==================================================
Python Facts:
==================================================
- Python version: 3.11.0
- Guido van Rossum created Python in 1991
- Python is named after Monty Python
- Python supports multiple programming paradigms
Greeting saved to greeting_log.txt

==================================================
Thank you for using Advanced Hello World!
Happy coding with Python! 🐍
==================================================
```

## Project Structure

```
stream1/
├── main.py              # Main application code
├── requirements.txt     # Python dependencies
├── README.md           # This file
├── .gitignore          # Git ignore rules
├── config.json         # Configuration file (homework)
├── greeting_log.txt    # Generated log file
└── venv/               # Virtual environment (not committed)
```

## Code Concepts Demonstrated

### 1. Type Hints
```python
def greet_user(name: str, greeting: Optional[str] = None) -> str:
```

### 2. Error Handling
```python
try:
    birth_year = int(input("Enter your birth year (YYYY): "))
except ValueError:
    print("Invalid year format.")
```

### 3. Context Managers
```python
with open(filename, "a", encoding="utf-8") as file:
    file.write(f"[{timestamp}] {message}\n")
```

### 4. F-Strings
```python
return f"{greeting}, {name}! Welcome to Python programming."
```

### 5. Main Guard Pattern
```python
if __name__ == "__main__":
    main()
```

## Homework Enhancements

### 1. Add More Greeting Languages
Add greetings in more languages to the `GREETINGS` list.

### 2. Implement Age-Based Messages
Add conditional messages based on age groups:
- Child (0-12)
- Teen (13-19)
- Adult (20-59)
- Senior (60+)

### 3. Add Color Output
Install colorama and add colored output:
```bash
pip install colorama
```

### 4. Create Configuration File
Create `config.json` for customizable settings:
```json
{
    "version": "1.0.0",
    "greetings": ["Hello", "Hi", "Hey"],
    "log_file": "greeting_log.txt",
    "enable_colors": true
}
```

## Testing

### Test Different Scenarios

1. **Normal input**: Enter valid name and year
2. **Empty name**: Press Enter without typing
3. **Short name**: Enter single character
4. **Invalid year**: Enter non-numeric value
5. **Keyboard interrupt**: Press Ctrl+C during execution

### Run Tests Manually

```bash
# Test with valid input
echo -e "Alice\n1990" | python main.py

# Test with empty name
echo -e "\n1990" | python main.py

# Test with invalid year
echo -e "Bob\nabc" | python main.py
```

## Troubleshooting

### Issue: Python command not found
**Solution**: Ensure Python 3.11+ is installed and added to PATH.

### Issue: Virtual environment activation fails
**Solution**: 
- Windows: Use `venv\Scripts\activate`
- macOS/Linux: Use `source venv/bin/activate`

### Issue: Permission denied writing file
**Solution**: Check directory write permissions or run with appropriate permissions.

### Issue: Encoding errors
**Solution**: Ensure UTF-8 encoding is used in file operations (already implemented).

## Learning Resources

- [Python Official Documentation](https://docs.python.org/3/)
- [Python Tutorial](https://docs.python.org/3/tutorial/)
- [Real Python](https://realpython.com/)
- [VS Code Python Extension](https://code.visualstudio.com/docs/python)

## License

This project is part of the Python Live Coding series and is available for educational purposes.

## Contributing

This is a learning project. Feel free to fork and enhance it as part of your Python learning journey!

## Stream Information

- **Stream**: Stream 1 - Environment Setup & Core Concepts
- **Duration**: 2 hours
- **Repository**: `long_questions/Python/core-python/basics/`
- **Next Stream**: Object-Oriented Programming Deep Dive

---

Happy coding! 🐍
