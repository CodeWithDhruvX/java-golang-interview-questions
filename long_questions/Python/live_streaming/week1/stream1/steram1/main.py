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
