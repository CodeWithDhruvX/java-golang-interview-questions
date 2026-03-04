# 📘 02 — OOP & Classes in Python
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- Defining classes, `__init__`, `self`
- Inheritance and `super()`
- Polymorphism and method overriding
- Encapsulation (private/protected)
- Dunder/magic methods
- `@property`, `@staticmethod`, `@classmethod`
- `isinstance()` and `issubclass()`

---

## ❓ Most Asked Questions

### Q1. How do you define a class in Python and what is `self`?

```python
class Employee:
    # Class variable: shared across all instances
    company = "TechCorp"
    headcount = 0

    def __init__(self, name, salary):
        # Instance variables: unique to each object
        self.name = name
        self.salary = salary
        Employee.headcount += 1   # modify class variable

    def introduce(self):
        # self refers to the calling instance
        return f"I'm {self.name}, earning ₹{self.salary:,}"

    def __repr__(self):
        return f"Employee(name={self.name!r}, salary={self.salary})"

    def __str__(self):
        return self.introduce()

e1 = Employee("Rahul", 80000)
e2 = Employee("Priya", 95000)

print(e1.introduce())      # I'm Rahul, earning ₹80,000
print(e2)                  # I'm Priya, earning ₹95,000
print(Employee.headcount)  # 2
print(e1.company)          # TechCorp (accessed via instance too)

# Instance dict vs class dict
print(e1.__dict__)         # {'name': 'Rahul', 'salary': 80000}
print(Employee.__dict__)   # includes company, headcount, methods
```

---

### Q2. How does inheritance work in Python?

```python
class Animal:
    def __init__(self, name, sound):
        self.name = name
        self.sound = sound

    def speak(self):
        return f"{self.name} says {self.sound}"

    def __repr__(self):
        return f"{self.__class__.__name__}({self.name!r})"

# Single inheritance
class Dog(Animal):
    def __init__(self, name):
        super().__init__(name, "Woof")  # call parent __init__

    def fetch(self, item):
        return f"{self.name} fetches the {item}!"

class Cat(Animal):
    def __init__(self, name):
        super().__init__(name, "Meow")

    def speak(self):  # override parent method
        return f"{self.name} says {self.sound}... and ignores you."

dog = Dog("Bruno")
cat = Cat("Whiskers")

print(dog.speak())   # Bruno says Woof
print(cat.speak())   # Whiskers says Meow... and ignores you.
print(dog.fetch("ball"))  # Bruno fetches the ball!

# isinstance and issubclass
isinstance(dog, Dog)     # True
isinstance(dog, Animal)  # True — Dog IS-A Animal
issubclass(Dog, Animal)  # True
issubclass(Dog, Cat)     # False

# Multiple inheritance
class Flyable:
    def fly(self):
        return f"{self.name} is flying!"

class FlyingDog(Dog, Flyable):
    pass

fd = FlyingDog("AirBud")
print(fd.speak())   # AirBud says Woof
print(fd.fly())     # AirBud is flying!
print(FlyingDog.__mro__)  # Method Resolution Order
```

---

### Q3. What are dunder/magic methods? Give key examples.

```python
class Vector:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    # String representations
    def __repr__(self):
        return f"Vector({self.x}, {self.y})"   # for debugging

    def __str__(self):
        return f"({self.x}, {self.y})"          # for print()

    # Arithmetic operators
    def __add__(self, other):
        return Vector(self.x + other.x, self.y + other.y)

    def __sub__(self, other):
        return Vector(self.x - other.x, self.y - other.y)

    def __mul__(self, scalar):
        return Vector(self.x * scalar, self.y * scalar)

    def __rmul__(self, scalar):  # scalar * vector
        return self.__mul__(scalar)

    # Comparison
    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

    def __lt__(self, other):
        return (self.x**2 + self.y**2) < (other.x**2 + other.y**2)

    # Container protocol
    def __len__(self):
        return 2  # 2D vector has 2 components

    def __getitem__(self, index):
        return (self.x, self.y)[index]

    # Context manager protocol
    def __enter__(self):
        print("Entering vector context")
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        print("Exiting vector context")
        return False  # don't suppress exceptions

v1 = Vector(1, 2)
v2 = Vector(3, 4)

print(v1 + v2)    # (4, 6)
print(v1 * 3)     # (3, 6)
print(3 * v1)     # (3, 6)
print(v1 == v2)   # False
print(len(v1))    # 2
print(v1[0])      # 1

with Vector(1, 2) as v:
    print(v)      # (1, 2)
```

---

### Q4. What is encapsulation and how do you achieve it in Python?

```python
class BankAccount:
    def __init__(self, owner, balance):
        self.owner = owner           # public
        self._account_type = "savings"  # protected (convention)
        self.__balance = balance     # private (name-mangled)

    # Property: controlled access
    @property
    def balance(self):
        return self.__balance

    @balance.setter
    def balance(self, amount):
        if amount < 0:
            raise ValueError("Balance cannot be negative")
        self.__balance = amount

    def deposit(self, amount):
        if amount <= 0:
            raise ValueError("Deposit must be positive")
        self.__balance += amount
        return self.__balance

    def withdraw(self, amount):
        if amount > self.__balance:
            raise ValueError("Insufficient funds")
        self.__balance -= amount
        return self.__balance

acc = BankAccount("Anjali", 10000)

# Accessing public attribute
print(acc.owner)    # "Anjali"

# Accessing via property
print(acc.balance)  # 10000

# Setting via property setter
acc.balance = 15000  # allowed
# acc.balance = -1  # raises ValueError

# Private name mangling: __attr → _ClassName__attr
# print(acc.__balance)         # AttributeError
print(acc._BankAccount__balance)  # 15000 (bypass — not recommended)

acc.deposit(5000)   # 20000
acc.withdraw(3000)  # 17000
```

---

### Q5. How do `@classmethod` and `@staticmethod` differ?

```python
class Date:
    def __init__(self, day, month, year):
        self.day = day
        self.month = month
        self.year = year

    def __repr__(self):
        return f"Date({self.day:02d}/{self.month:02d}/{self.year})"

    # classmethod: receives class as first argument (cls)
    # Used for alternative constructors, factory methods
    @classmethod
    def from_string(cls, date_string):
        # "dd-mm-yyyy"
        day, month, year = map(int, date_string.split("-"))
        return cls(day, month, year)  # cls() works for subclasses too!

    @classmethod
    def today(cls):
        import datetime
        d = datetime.date.today()
        return cls(d.day, d.month, d.year)

    # staticmethod: no self or cls — pure utility function
    # Logically belongs to class but doesn't need class/instance data
    @staticmethod
    def is_valid_date(day, month, year):
        if month < 1 or month > 12:
            return False
        if day < 1 or day > 31:
            return False
        return True

    @staticmethod
    def is_leap_year(year):
        return year % 4 == 0 and (year % 100 != 0 or year % 400 == 0)

# Usage
d1 = Date(15, 8, 1947)
d2 = Date.from_string("26-01-1950")  # classmethod
d3 = Date.today()                     # classmethod

Date.is_valid_date(30, 2, 2024)   # False — staticmethod
Date.is_leap_year(2024)           # True — staticmethod

# Key differences:
# instance method: needs instance (self) → access/modify instance
# classmethod:  needs class (cls) → factory methods, class state
# staticmethod: needs neither → utility, helper functions
```

---

### Q6. Explain polymorphism with a real-world example.

```python
# Polymorphism: same interface, different behaviours

from abc import ABC, abstractmethod

class Shape(ABC):
    @abstractmethod
    def area(self) -> float:
        pass

    @abstractmethod
    def perimeter(self) -> float:
        pass

    def describe(self):
        # Works for ALL subclasses — uses their area() and perimeter()
        return f"{self.__class__.__name__}: area={self.area():.2f}, perimeter={self.perimeter():.2f}"

class Circle(Shape):
    def __init__(self, radius):
        self.radius = radius

    def area(self):
        return 3.14159 * self.radius ** 2

    def perimeter(self):
        return 2 * 3.14159 * self.radius

class Rectangle(Shape):
    def __init__(self, width, height):
        self.width = width
        self.height = height

    def area(self):
        return self.width * self.height

    def perimeter(self):
        return 2 * (self.width + self.height)

class Triangle(Shape):
    def __init__(self, a, b, c):
        self.a, self.b, self.c = a, b, c

    def area(self):
        s = (self.a + self.b + self.c) / 2
        return (s * (s - self.a) * (s - self.b) * (s - self.c)) ** 0.5

    def perimeter(self):
        return self.a + self.b + self.c

# Polymorphic function — works with any Shape
def print_areas(shapes):
    for shape in shapes:
        print(shape.describe())  # each shape responds differently!

shapes = [Circle(5), Rectangle(4, 6), Triangle(3, 4, 5)]
print_areas(shapes)
# Circle: area=78.54, perimeter=31.42
# Rectangle: area=24.00, perimeter=20.00
# Triangle: area=6.00, perimeter=12.00

# Can't instantiate abstract class
# Shape()  # TypeError: Can't instantiate abstract class
```

---

### Q7. What is the difference between `__repr__` and `__str__`?

```python
class Product:
    def __init__(self, name, price):
        self.name = name
        self.price = price

    def __repr__(self):
        # Goal: unambiguous, ideally eval()-able representation
        # Used by: repr(), in REPL, in containers (lists, dicts)
        return f"Product(name={self.name!r}, price={self.price})"

    def __str__(self):
        # Goal: human-readable, for end users
        # Used by: print(), str(), f-strings
        return f"{self.name} — ₹{self.price:,}"

p = Product("Laptop", 75000)

repr(p)   # "Product(name='Laptop', price=75000)"  ← __repr__
str(p)    # "Laptop — ₹75,000"                     ← __str__
print(p)  # Laptop — ₹75,000                       ← __str__

# In a list, __repr__ is used:
products = [p, Product("Phone", 25000)]
print(products)
# [Product(name='Laptop', price=75000), Product(name='Phone', price=25000)]

# f-string uses __str__ by default, !r forces __repr__
print(f"{p}")    # Laptop — ₹75,000
print(f"{p!r}")  # Product(name='Laptop', price=75000)

# If only __repr__ is defined, it's also used as __str__
# If only __str__ is defined, __repr__ falls back to default '<...object...>'
# Best practice: always define __repr__; optionally add __str__ for user display
```

---

### Q8. What is Method Resolution Order (MRO)?

```python
# MRO: the order in which Python searches for methods in inheritance hierarchy
# Python uses C3 Linearization algorithm

class A:
    def hello(self):
        return "A"

class B(A):
    def hello(self):
        return "B"

class C(A):
    def hello(self):
        return "C"

class D(B, C):
    pass

# MRO for D:
print(D.__mro__)
# (<class 'D'>, <class 'B'>, <class 'C'>, <class 'A'>, <class 'object'>)

print(D().hello())  # "B" — B comes before C in MRO

# super() follows MRO
class X:
    def method(self):
        print("X")

class Y(X):
    def method(self):
        print("Y")
        super().method()   # goes to next in MRO

class Z(X):
    def method(self):
        print("Z")
        super().method()

class W(Y, Z):
    def method(self):
        print("W")
        super().method()

# MRO: W → Y → Z → X → object
W().method()
# W
# Y
# Z
# X

# Diamond problem is handled by MRO:
#     A
#    / \
#   B   C
#    \ /
#     D
# MRO: D, B, C, A — each class appears only ONCE
```
