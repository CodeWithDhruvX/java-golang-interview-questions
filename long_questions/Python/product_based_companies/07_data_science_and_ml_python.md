# ⚡ 07 — Data Science & ML Python Patterns
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- NumPy array operations and broadcasting
- Pandas DataFrames: groupby, merge, pivot
- Data pipeline patterns
- Serialization: JSON, Pickle, Parquet
- Type hints and `pydantic` for data validation
- Memory-efficient data loading

---

## ❓ Most Asked Questions

### Q1. Essential NumPy operations for interviews.

```python
import numpy as np

# Array creation
arr = np.array([1, 2, 3, 4, 5])
zeros = np.zeros((3, 4))
ones  = np.ones((2, 3), dtype=np.float32)
eye   = np.eye(3)            # identity matrix
rng   = np.arange(0, 10, 2) # [0, 2, 4, 6, 8]
lin   = np.linspace(0, 1, 5) # [0.0, 0.25, 0.5, 0.75, 1.0]

# Reshape and indexing
arr = np.arange(24).reshape(4, 6)  # 4x6 matrix
arr.shape    # (4, 6)
arr.ndim     # 2
arr.dtype    # int64

arr[2, 3]       # element (row=2, col=3)
arr[1:3, 2:5]   # slice: rows 1-2, cols 2-4
arr[:, 0]       # first column
arr[-1, :]      # last row

# Boolean indexing
data = np.array([10, -5, 3, -8, 7, -2])
positive = data[data > 0]   # [10, 3, 7]
data[data < 0] = 0          # zero out negatives → [10, 0, 3, 0, 7, 0]

# Broadcasting: operations on different shapes
a = np.array([[1], [2], [3]])    # shape (3, 1)
b = np.array([10, 20, 30, 40])  # shape (4,)
a + b   # broadcasts to (3, 4):
# [[11, 21, 31, 41],
#  [12, 22, 32, 42],
#  [13, 23, 33, 43]]

# Vectorized operations (much faster than loops!)
arr = np.random.randn(1_000_000)
# Slow: for loop calculating mean
mean_loop = sum(arr) / len(arr)  # O(n) Python loop

# Fast: vectorized
mean_np = arr.mean()     # C-level operation, ~50x faster
arr.std()
arr.min(), arr.max()
np.percentile(arr, [25, 50, 75])

# Matrix operations
A = np.array([[1, 2], [3, 4]])
B = np.array([[5, 6], [7, 8]])

A @ B     # matrix multiplication
A.T       # transpose
np.linalg.det(A)   # determinant
np.linalg.inv(A)   # inverse
np.dot(A, B)       # same as @

# Memory layout (important for performance)
# C-order (row-major): default — row operations are fast
# F-order (column-major): column operations are fast
c_arr = np.zeros((1000, 1000), order='C')
f_arr = np.zeros((1000, 1000), order='F')
```

---

### Q2. Essential Pandas for data manipulation.

```python
import pandas as pd
import numpy as np

# Create DataFrame
df = pd.DataFrame({
    "name":    ["Alice", "Bob", "Carol", "Dave", "Eve"],
    "dept":    ["Eng", "HR", "Eng", "Finance", "HR"],
    "salary":  [90000, 65000, 85000, 75000, 70000],
    "years":   [3, 5, 2, 8, 4],
    "active":  [True, True, True, False, True]
})

# Basic exploration
df.head(3)
df.info()
df.describe()
df.dtypes

# Filtering
eng_staff = df[df["dept"] == "Eng"]
high_paid = df[df["salary"] > 80000]
active_eng = df[(df["dept"] == "Eng") & (df["active"])]

# Selecting columns
df[["name", "salary"]]
df.loc[:, "salary":"years"]   # slice by label
df.iloc[1:3, 0:2]            # slice by index

# Groupby
dept_stats = df.groupby("dept").agg(
    avg_salary=("salary", "mean"),
    total       =("name", "count"),
    max_years   =("years", "max")
).reset_index()

# Apply: row-wise / column-wise transformation
df["salary_band"] = df["salary"].apply(
    lambda s: "High" if s > 80000 else "Mid" if s > 65000 else "Low"
)
df["bonus"] = df.apply(
    lambda row: row["salary"] * 0.15 if row["active"] else 0,
    axis=1  # row-wise
)

# Merge: SQL-style joins
departments = pd.DataFrame({
    "dept": ["Eng", "HR", "Finance"],
    "location": ["Bangalore", "Delhi", "Mumbai"]
})
merged = df.merge(departments, on="dept", how="left")

# Pivot table
pivot = df.pivot_table(
    values="salary",
    index="dept",
    aggfunc=["mean", "count"]
)

# Handle missing data
df["bonus"].fillna(0, inplace=True)
df.dropna(subset=["name"], inplace=True)
df.isnull().sum()   # count nulls per column

# Efficient reading
df_csv     = pd.read_csv("data.csv", usecols=["id", "name"], dtype={"id": np.int32})
df_parquet = pd.read_parquet("data.parquet", columns=["id", "name"])
df.to_parquet("output.parquet", index=False)   # ~5x smaller than CSV
```

---

### Q3. Data serialization formats — JSON, Pickle, Parquet.

```python
import json, pickle
import datetime
from pathlib import Path

# JSON: text-based, human-readable, language-agnostic
# Limitation: only handles str, int, float, list, dict, bool, None

class DateTimeEncoder(json.JSONEncoder):
    def default(self, obj):
        if isinstance(obj, datetime.datetime):
            return obj.isoformat()      # serialize datetime as ISO string
        if isinstance(obj, set):
            return list(obj)            # serialize set as list
        return super().default(obj)

data = {
    "name": "Priya",
    "joined": datetime.datetime.now(),
    "skills": {"Python", "Django"}
}

# Serialize
json_str = json.dumps(data, cls=DateTimeEncoder, indent=2)

# Deserialize (need custom hook for datetime)
def decode_dates(d):
    for k, v in d.items():
        if isinstance(v, str):
            try:
                d[k] = datetime.datetime.fromisoformat(v)
            except ValueError:
                pass
    return d

parsed = json.loads(json_str, object_hook=decode_dates)

# Pickle: Python-only, handles any Python object
# ⚠️ NEVER unpickle from untrusted sources — arbitrary code execution!
data = {"model": some_ml_model, "timestamp": datetime.datetime.now()}

with open("model.pkl", "wb") as f:
    pickle.dump(data, f, protocol=pickle.HIGHEST_PROTOCOL)

with open("model.pkl", "rb") as f:
    loaded = pickle.load(f)

# Parquet: columnar binary format, ideal for DataFrames
import pandas as pd

df = pd.DataFrame({"a": range(1_000_000), "b": range(1_000_000)})

# Comparison:
# CSV:     ~8 MB, slow read, no types
# Parquet: ~2 MB, fast read (columnar), preserves types
df.to_parquet("data.parquet", compression="snappy")    # compressed!
df_back = pd.read_parquet("data.parquet")

# MessagePack: faster than JSON, binary format
# import msgpack
# packed = msgpack.packb(data)
# unpacked = msgpack.unpackb(packed, raw=False)
```

---

### Q4. Type hints, `pydantic`, and data validation best practices.

```python
from pydantic import BaseModel, Field, validator, root_validator
from typing import Optional, List
from datetime import datetime

# Pydantic: data validation + serialization using Python type hints

class Address(BaseModel):
    street: str
    city: str
    pincode: str = Field(..., regex=r"^\d{6}$")   # must be 6 digits

class User(BaseModel):
    id: int
    name: str = Field(..., min_length=2, max_length=100)
    email: str = Field(..., regex=r"^[^@]+@[^@]+\.[^@]+$")
    age: Optional[int] = Field(None, ge=0, le=150)
    salary: float = Field(..., gt=0)
    skills: List[str] = []
    address: Optional[Address] = None
    created_at: datetime = Field(default_factory=datetime.now)

    @validator("name")
    def normalize_name(cls, v):
        return v.strip().title()

    @validator("email")
    def lowercase_email(cls, v):
        return v.lower()

    @root_validator
    def senior_employee_must_have_address(cls, values):
        salary = values.get("salary", 0)
        address = values.get("address")
        if salary > 100000 and address is None:
            raise ValueError("Senior employees must have an address on file")
        return values

    class Config:
        # Allow extra fields to be ignored
        extra = "ignore"
        # Validate on assignment
        validate_assignment = True

# Usage
user = User(
    id=1,
    name="  rahul sharma  ",
    email="RAHUL@EXAMPLE.COM",
    age=28,
    salary=75000,
    skills=["Python", "FastAPI"]
)
print(user.name)    # "Rahul Sharma" (normalized)
print(user.email)   # "rahul@example.com" (lowercased)

# Validation errors are detailed
try:
    bad = User(id="not-an-int", name="X", email="bad", salary=-100)
except ValueError as e:
    print(e.errors())   # list of validation errors with field paths

# Serialize
user.dict()   # → Python dict
user.json()   # → JSON string
User.parse_raw(user.json())  # → User from JSON
User.parse_obj({"id": 1, "name": "Rahul", ...})  # → User from dict
```
