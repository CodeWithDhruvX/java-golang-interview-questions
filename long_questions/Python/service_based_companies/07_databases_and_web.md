# 📘 07 — Databases & Web in Python
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- Connecting to databases with `sqlite3` and SQLAlchemy
- ORM basics: models, queries, relationships
- REST APIs with Flask / FastAPI
- JSON handling, status codes, request/response cycle
- Database migrations with Alembic

---

## ❓ Most Asked Questions

### Q1. How do you work with SQLite3 in Python?

```python
import sqlite3

# Connect (creates file if not exists)
conn = sqlite3.connect("app.db")
conn.row_factory = sqlite3.Row   # access columns by name!
cursor = conn.cursor()

# Create table
cursor.execute("""
    CREATE TABLE IF NOT EXISTS users (
        id      INTEGER PRIMARY KEY AUTOINCREMENT,
        name    TEXT    NOT NULL,
        email   TEXT    UNIQUE NOT NULL,
        age     INTEGER,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )
""")
conn.commit()

# INSERT — always use parameterized queries (prevents SQL injection!)
cursor.execute(
    "INSERT INTO users (name, email, age) VALUES (?, ?, ?)",
    ("Rahul", "rahul@example.com", 28)
)
conn.commit()

# INSERT many
users = [
    ("Priya", "priya@example.com", 26),
    ("Anjali", "anjali@example.com", 30),
]
cursor.executemany(
    "INSERT INTO users (name, email, age) VALUES (?, ?, ?)", users
)
conn.commit()

# SELECT
cursor.execute("SELECT * FROM users WHERE age > ?", (25,))
rows = cursor.fetchall()      # all rows
for row in rows:
    print(row["name"], row["email"])  # dict-like access via Row factory

cursor.execute("SELECT * FROM users")
row = cursor.fetchone()       # first row only

# UPDATE
cursor.execute("UPDATE users SET age = ? WHERE email = ?", (29, "rahul@example.com"))
print(cursor.rowcount)        # number of rows affected
conn.commit()

# DELETE
cursor.execute("DELETE FROM users WHERE id = ?", (1,))
conn.commit()

# Context manager (auto-commit or rollback)
with sqlite3.connect("app.db") as conn:
    conn.execute("UPDATE users SET age = age + 1 WHERE id = ?", (1,))
    # auto-commits on success, auto-rollbacks on error

conn.close()
```

---

### Q2. What is SQLAlchemy and how do you use the ORM?

```python
from sqlalchemy import create_engine, Column, Integer, String, ForeignKey, DateTime
from sqlalchemy.orm import declarative_base, relationship, Session
from sqlalchemy.sql import func

# Create engine
engine = create_engine("sqlite:///app.db", echo=True)  # echo=True → log SQL

Base = declarative_base()

# Define models
class Department(Base):
    __tablename__ = "departments"

    id   = Column(Integer, primary_key=True)
    name = Column(String(50), nullable=False, unique=True)

    employees = relationship("Employee", back_populates="department")

    def __repr__(self):
        return f"Department(name={self.name!r})"

class Employee(Base):
    __tablename__ = "employees"

    id         = Column(Integer, primary_key=True)
    name       = Column(String(100), nullable=False)
    email      = Column(String(200), unique=True, nullable=False)
    salary     = Column(Integer, default=50000)
    dept_id    = Column(Integer, ForeignKey("departments.id"), nullable=False)
    created_at = Column(DateTime, server_default=func.now())

    department = relationship("Department", back_populates="employees")

# Create tables
Base.metadata.create_all(engine)

# Create, Read, Update, Delete (CRUD)
with Session(engine) as session:
    # CREATE
    eng_dept = Department(name="Engineering")
    session.add(eng_dept)
    session.flush()  # send SQL but not commit yet — gets dept id

    emp = Employee(name="Alice", email="alice@co.com", salary=90000, dept_id=eng_dept.id)
    session.add(emp)
    session.commit()

    # READ
    all_employees = session.query(Employee).all()
    alice = session.query(Employee).filter_by(email="alice@co.com").first()

    # Filter
    high_paid = session.query(Employee).filter(Employee.salary > 80000).all()
    eng_emp   = session.query(Employee).join(Department).filter(
        Department.name == "Engineering"
    ).all()

    # UPDATE
    alice.salary = 95000
    session.commit()
    # or: session.query(Employee).filter_by(id=alice.id).update({"salary": 95000})

    # DELETE
    session.delete(alice)
    session.commit()
```

---

### Q3. How do you build a REST API with Flask?

```python
from flask import Flask, request, jsonify, abort
from datetime import datetime

app = Flask(__name__)

# In-memory store (use DB in production)
users = {}
next_id = 1

def make_user(name, email):
    global next_id
    user = {
        "id": next_id,
        "name": name,
        "email": email,
        "created_at": datetime.now().isoformat()
    }
    users[next_id] = user
    next_id += 1
    return user

# GET /users — list all users
@app.route("/users", methods=["GET"])
def get_users():
    return jsonify(list(users.values())), 200

# GET /users/<id> — get one user
@app.route("/users/<int:user_id>", methods=["GET"])
def get_user(user_id):
    user = users.get(user_id)
    if not user:
        abort(404, description=f"User {user_id} not found")
    return jsonify(user), 200

# POST /users — create user
@app.route("/users", methods=["POST"])
def create_user():
    data = request.get_json()
    if not data or "name" not in data or "email" not in data:
        abort(400, description="name and email are required")
    user = make_user(data["name"], data["email"])
    return jsonify(user), 201

# PUT /users/<id> — update user
@app.route("/users/<int:user_id>", methods=["PUT"])
def update_user(user_id):
    user = users.get(user_id)
    if not user:
        abort(404)
    data = request.get_json()
    user.update({k: v for k, v in data.items() if k in ("name", "email")})
    return jsonify(user), 200

# DELETE /users/<id> — delete user
@app.route("/users/<int:user_id>", methods=["DELETE"])
def delete_user(user_id):
    if user_id not in users:
        abort(404)
    del users[user_id]
    return "", 204

# Error handler
@app.errorhandler(404)
def not_found(e):
    return jsonify({"error": str(e)}), 404

@app.errorhandler(400)
def bad_request(e):
    return jsonify({"error": str(e)}), 400

if __name__ == "__main__":
    app.run(debug=True, port=5000)
```

---

### Q4. How do you build a REST API with FastAPI?

```python
from fastapi import FastAPI, HTTPException, status
from pydantic import BaseModel, EmailStr
from typing import Optional
from datetime import datetime

app = FastAPI(title="User API", version="1.0")

# Pydantic models — request/response validation
class UserCreate(BaseModel):
    name: str
    email: str
    age: Optional[int] = None

class UserResponse(BaseModel):
    id: int
    name: str
    email: str
    age: Optional[int]
    created_at: datetime

# In-memory storage
db: dict[int, dict] = {}
counter = 1

@app.get("/users", response_model=list[UserResponse])
def list_users():
    return list(db.values())

@app.get("/users/{user_id}", response_model=UserResponse)
def get_user(user_id: int):
    if user_id not in db:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail=f"User {user_id} not found"
        )
    return db[user_id]

@app.post("/users", response_model=UserResponse, status_code=201)
def create_user(user: UserCreate):
    global counter
    new_user = {
        "id": counter,
        "name": user.name,
        "email": user.email,
        "age": user.age,
        "created_at": datetime.now()
    }
    db[counter] = new_user
    counter += 1
    return new_user

@app.delete("/users/{user_id}", status_code=204)
def delete_user(user_id: int):
    if user_id not in db:
        raise HTTPException(status_code=404, detail="Not found")
    del db[user_id]

# Advantages of FastAPI over Flask:
# ✅ Auto-generated interactive docs (Swagger UI at /docs)
# ✅ Automatic request/response validation via Pydantic
# ✅ Type hints → better IDE support and fewer runtime bugs
# ✅ Async support out of the box
# ✅ Much faster (Starlette + uvicorn = near Node.js performance)
```

---

### Q5. How do you make HTTP requests with Python?

```python
import requests

BASE_URL = "https://jsonplaceholder.typicode.com"

# GET request
response = requests.get(f"{BASE_URL}/posts/1")
response.status_code    # 200
response.json()         # parsed JSON as dict
response.text           # raw text
response.headers        # response headers

# POST with JSON body
new_post = {"title": "Hello", "body": "World", "userId": 1}
resp = requests.post(f"{BASE_URL}/posts", json=new_post)
print(resp.status_code)  # 201
print(resp.json()["id"])  # newly created ID

# With headers and auth
headers = {"Authorization": "Bearer my-token", "Content-Type": "application/json"}
resp = requests.get(f"{BASE_URL}/posts", headers=headers, params={"userId": 1})

# Query parameters
resp = requests.get(f"{BASE_URL}/posts", params={"userId": 1, "_limit": 10})

# Error handling
try:
    resp = requests.get("https://api.example.com/data", timeout=5)
    resp.raise_for_status()    # raise HTTPError for 4xx/5xx responses
    data = resp.json()
except requests.exceptions.Timeout:
    print("Request timed out")
except requests.exceptions.ConnectionError:
    print("Network error")
except requests.exceptions.HTTPError as e:
    print(f"HTTP Error: {e.response.status_code}")

# Sessions (reuse TCP connections + cookies)
with requests.Session() as session:
    session.headers.update({"Authorization": "Bearer token"})
    session.get(f"{BASE_URL}/posts")    # reuses connection
    session.get(f"{BASE_URL}/users")    # same session

# Async HTTP — httpx (preferred for async code)
import httpx
import asyncio

async def fetch_data():
    async with httpx.AsyncClient() as client:
        resp = await client.get("https://api.example.com/data")
        return resp.json()

asyncio.run(fetch_data())
```
