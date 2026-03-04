# 📘 03 — Data Structures & Algorithms in Python
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- Lists, tuples, sets, dicts — operations and complexity
- Sorting and searching
- Stack, queue using Python collections
- `collections` module: `Counter`, `defaultdict`, `deque`, `OrderedDict`
- Common algorithmic patterns

---

## ❓ Most Asked Questions

### Q1. What are the time complexities of common list operations?

```python
# Python list is a dynamic array (like ArrayList in Java)

lst = [1, 2, 3, 4, 5]

# O(1) operations
lst.append(6)     # append to end
lst.pop()         # remove from end
lst[-1]           # access last element
len(lst)          # length

# O(n) operations
lst.insert(0, 0)  # insert at beginning — shifts all elements
lst.pop(0)        # remove from beginning — shifts all elements
5 in lst          # search — linear scan
lst.index(3)      # find index — linear scan
lst.remove(3)     # find and remove — linear scan
lst.reverse()     # in-place reverse
lst.copy()        # copy

# O(n log n) operations
lst.sort()           # in-place sort (Timsort)
sorted(lst)          # returns new sorted list

# ⚠️ Use deque for O(1) front operations
from collections import deque

dq = deque([1, 2, 3])
dq.appendleft(0)  # O(1)
dq.popleft()      # O(1)

# List vs deque
# List:  append/pop from RIGHT = O(1), from LEFT = O(n)
# Deque: append/pop from BOTH ends = O(1), random access = O(n)
```

---

### Q2. How do dictionaries work and what are their time complexities?

```python
# Python dict is a hash map — average O(1) for get/set/delete

d = {"name": "Rahul", "age": 25, "city": "Delhi"}

# O(1) operations
d["name"]         # access by key
d["email"] = "rahul@example.com"  # insert/update
del d["city"]     # delete
"age" in d        # membership check (by key)
len(d)            # length

# Safe access
d.get("missing", "default")  # returns "default" — no KeyError!
d.setdefault("score", 0)     # sets key only if not present

# Iteration
d.keys()    # dict_keys(['name', 'age'])
d.values()  # dict_values(['Rahul', 25])
d.items()   # dict_items([('name', 'Rahul'), ('age', 25)])

for key, value in d.items():
    print(f"{key}: {value}")

# Dict comprehension
squares = {x: x**2 for x in range(1, 6)}
# {1: 1, 2: 4, 3: 9, 4: 16, 5: 25}

# Merging dicts
a = {"x": 1, "y": 2}
b = {"y": 99, "z": 3}
merged = {**a, **b}  # b values override a
# {"x": 1, "y": 99, "z": 3}

# Python 3.9+:
merged = a | b   # same result

# Counter: count occurrences
from collections import Counter

words = ["apple", "banana", "apple", "cherry", "banana", "apple"]
count = Counter(words)
# Counter({'apple': 3, 'banana': 2, 'cherry': 1})

count.most_common(2)    # [('apple', 3), ('banana', 2)]
count["grape"]          # 0 (no KeyError!)

# defaultdict: default value for missing keys
from collections import defaultdict

dd = defaultdict(list)
dd["fruits"].append("apple")   # no need to check if key exists!
dd["fruits"].append("mango")
# {"fruits": ["apple", "mango"]}
```

---

### Q3. How do you implement a stack and queue in Python?

```python
# STACK → LIFO (Last In, First Out)
# Use: undo/redo, parsing, DFS

# Using list (simplest)
stack = []
stack.append(1)   # push
stack.append(2)
stack.append(3)
stack.pop()       # pop → 3 (O(1))
stack[-1]         # peek → 2 (O(1))

# Using deque (thread-safe alternative)
from collections import deque

stack_dq = deque()
stack_dq.append(1)
stack_dq.append(2)
stack_dq.pop()    # → 2

# QUEUE → FIFO (First In, First Out)
# Use: BFS, task scheduling, print queue

# ❌ Don't use list for queue (popleft is O(n))
# ✅ Use deque for O(1) enqueue and dequeue
queue = deque()
queue.append(1)      # enqueue
queue.append(2)
queue.append(3)
queue.popleft()      # dequeue → 1 (O(1))
queue[0]             # peek → 2

# For thread-safe queue (multi-threading):
from queue import Queue

q = Queue()
q.put("task1")
q.put("task2")
item = q.get()     # dequeues, blocks if empty

# Priority Queue (min-heap)
import heapq

pq = []
heapq.heappush(pq, (3, "low priority"))
heapq.heappush(pq, (1, "high priority"))
heapq.heappush(pq, (2, "medium priority"))

heapq.heappop(pq)   # (1, "high priority") — lowest priority first
```

---

### Q4. Explain sorting in Python — `sort()` vs `sorted()`.

```python
# list.sort(): in-place sort, returns None, modifies list
lst = [3, 1, 4, 1, 5, 9]
lst.sort()               # [1, 1, 3, 4, 5, 9]
lst.sort(reverse=True)   # [9, 5, 4, 3, 1, 1]

# sorted(): returns new sorted list, works on any iterable
lst = [3, 1, 4, 1, 5]
new_lst = sorted(lst)          # [1, 1, 3, 4, 5] — lst unchanged!
sorted(lst, reverse=True)      # [5, 4, 3, 1, 1]

# Sorting with key function
employees = [
    {"name": "Priya",  "salary": 80000, "age": 28},
    {"name": "Rahul",  "salary": 65000, "age": 32},
    {"name": "Anjali", "salary": 90000, "age": 25},
]

# Sort by salary
by_salary = sorted(employees, key=lambda e: e["salary"])

# Sort by name
by_name = sorted(employees, key=lambda e: e["name"])

# Sort by multiple fields: age ascending, salary descending
from operator import itemgetter
by_multi = sorted(employees, key=lambda e: (e["age"], -e["salary"]))

# Sort objects using attrgetter
from operator import attrgetter

class Student:
    def __init__(self, name, grade):
        self.name = name
        self.grade = grade

students = [Student("A", 85), Student("B", 92), Student("C", 78)]
sorted(students, key=attrgetter("grade"), reverse=True)

# Python's sort is STABLE (preserves relative order of equal elements)
# Algorithm: Timsort — O(n log n) worst case

# Counting sort / bucket sort (when input range is known):
def counting_sort(arr, max_val):
    count = [0] * (max_val + 1)
    for x in arr:
        count[x] += 1
    return [x for x, c in enumerate(count) for _ in range(c)]
```

---

### Q5. How do you work with sets and when should you use them?

```python
# Set: unordered, unique elements, O(1) average lookup

a = {1, 2, 3, 4, 5}
b = {3, 4, 5, 6, 7}

# Set operations
a | b   # union: {1, 2, 3, 4, 5, 6, 7}
a & b   # intersection: {3, 4, 5}
a - b   # difference: {1, 2}
b - a   # difference: {6, 7}
a ^ b   # symmetric difference: {1, 2, 6, 7}

# Methods
a.union(b)
a.intersection(b)
a.difference(b)
a.symmetric_difference(b)

# Subset/superset
{1, 2}.issubset({1, 2, 3})    # True
{1, 2, 3}.issuperset({1, 2})  # True
{1, 2}.isdisjoint({3, 4})     # True (no common elements)

# Add/remove
a.add(6)
a.discard(6)   # no error if not found
a.remove(6)    # raises KeyError if not found
a.clear()

# Use cases
# 1. Remove duplicates (order not preserved)
lst = [3, 1, 4, 1, 5, 9, 2, 6, 5, 3]
unique = list(set(lst))   # [1, 2, 3, 4, 5, 6, 9] — order may vary!

# 2. Fast membership check
valid_codes = {"ENG", "HR", "FIN", "IT"}
dept = "HR"
if dept in valid_codes:   # O(1)
    print("Valid department")

# 3. Finding common/different items
set1 = {"alice", "bob", "charlie"}
set2 = {"bob", "diana", "eve"}
common = set1 & set2     # {"bob"}
only_set1 = set1 - set2  # {"alice", "charlie"}

# frozenset: immutable set (can be used as dict key or set element)
fs = frozenset([1, 2, 3])
d = {fs: "immutable"}
```

---

### Q6. What is the `collections` module and its most-used tools?

```python
from collections import Counter, defaultdict, deque, OrderedDict, namedtuple

# Counter — frequency counting
text = "hello world hello python hello"
word_count = Counter(text.split())
# Counter({'hello': 3, 'world': 1, 'python': 1})

most_common = word_count.most_common(2)   # top 2
word_count["hello"] -= 1                  # subtract
word_count.update(["hello", "java"])      # add more
total = sum(word_count.values())

# defaultdict — auto-creates missing keys
graph = defaultdict(list)
graph["A"].append("B")    # no KeyError even if "A" didn't exist
graph["A"].append("C")
# {"A": ["B", "C"]}

word_groups = defaultdict(list)
words = ["cat", "car", "bar", "can", "ban"]
for w in words:
    word_groups[w[0]].append(w)
# {"c": ["cat", "car", "can"], "b": ["bar", "ban"]}

# deque — double-ended queue
dq = deque([1, 2, 3, 4, 5], maxlen=3)  # maxlen: auto-removes oldest
dq.append(6)      # [4, 5, 6]  ← 4 because maxlen=3 drops from left
dq.appendleft(0)  # [0, 4, 5]
dq.rotate(1)      # rotate right: [5, 0, 4]
dq.rotate(-1)     # rotate left: [0, 4, 5]

# namedtuple — lightweight immutable record
Point = namedtuple("Point", ["x", "y"])
p = Point(3, 4)
print(p.x, p.y)    # 3 4
print(p[0], p[1])  # 3 4 — indexable like tuple
x, y = p           # unpackable

Employee = namedtuple("Employee", ["name", "salary", "dept"])
emp = Employee("Rahul", 80000, "Engineering")
emp._asdict()   # OrderedDict form
emp._replace(salary=90000)  # returns new namedtuple
```

---

### Q7. Explain binary search and when to use it.

```python
# Binary search: find element in SORTED list in O(log n)

def binary_search(arr, target):
    left, right = 0, len(arr) - 1

    while left <= right:
        mid = (left + right) // 2    # avoid overflow (Python ints don't overflow)

        if arr[mid] == target:
            return mid               # found!
        elif arr[mid] < target:
            left = mid + 1           # search right half
        else:
            right = mid - 1          # search left half

    return -1  # not found

arr = [2, 5, 8, 12, 16, 23, 38, 56, 72, 91]
binary_search(arr, 23)   # 5
binary_search(arr, 10)   # -1

# Using built-in bisect module
import bisect

arr = [1, 3, 5, 7, 9]
bisect.bisect_left(arr, 5)   # 2 — leftmost position to insert 5
bisect.bisect_right(arr, 5)  # 3 — rightmost position
bisect.insort(arr, 6)        # insert 6 in sorted order: [1,3,5,6,7,9]

# Find first occurrence of target
def first_occurrence(arr, target):
    pos = bisect.bisect_left(arr, target)
    if pos < len(arr) and arr[pos] == target:
        return pos
    return -1

# Common interview variations:
# 1. Search in rotated sorted array
# 2. Find first/last occurrence
# 3. Find peak element
# 4. Find minimum in rotated array
def search_rotated(arr, target):
    left, right = 0, len(arr) - 1
    while left <= right:
        mid = (left + right) // 2
        if arr[mid] == target:
            return mid
        if arr[left] <= arr[mid]:  # left half is sorted
            if arr[left] <= target < arr[mid]:
                right = mid - 1
            else:
                left = mid + 1
        else:                       # right half is sorted
            if arr[mid] < target <= arr[right]:
                left = mid + 1
            else:
                right = mid - 1
    return -1
```

---

### Q8. Common algorithmic patterns you must know.

```python
# 1. Two Pointers — find pair with target sum in sorted array
def two_sum_sorted(arr, target):
    left, right = 0, len(arr) - 1
    while left < right:
        s = arr[left] + arr[right]
        if s == target:
            return (left, right)
        elif s < target:
            left += 1
        else:
            right -= 1
    return None

# 2. Sliding Window — max sum of k consecutive elements
def max_sum_k(arr, k):
    window_sum = sum(arr[:k])
    max_sum = window_sum
    for i in range(k, len(arr)):
        window_sum += arr[i] - arr[i - k]  # slide: add new, remove old
        max_sum = max(max_sum, window_sum)
    return max_sum

# 3. Fast and Slow Pointer — detect cycle in linked list
class ListNode:
    def __init__(self, val, next=None):
        self.val = val
        self.next = next

def has_cycle(head):
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            return True
    return False

# 4. Prefix Sum — range sum queries
def prefix_sum(arr):
    prefix = [0] * (len(arr) + 1)
    for i, x in enumerate(arr):
        prefix[i + 1] = prefix[i] + x
    return prefix

# range_sum(l, r) = prefix[r+1] - prefix[l]
arr = [1, 2, 3, 4, 5]
prefix = prefix_sum(arr)
print(prefix[2] - prefix[0])   # sum(arr[0:2]) = 3
print(prefix[5] - prefix[2])   # sum(arr[2:5]) = 12

# 5. Hash Map for O(1) lookup
def two_sum(nums, target):
    seen = {}
    for i, num in enumerate(nums):
        complement = target - num
        if complement in seen:
            return [seen[complement], i]
        seen[num] = i
```
