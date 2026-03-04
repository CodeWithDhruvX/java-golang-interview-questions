# ⚡ 06 — Design Patterns in Python
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 Must-Know Topics
- Creational: Singleton, Factory, Builder
- Structural: Adapter, Decorator, Proxy
- Behavioral: Observer, Strategy, Command
- Pythonic implementations vs. Java-style
- When not to use patterns

---

## ❓ Most Asked Questions

### Q1. Singleton, Factory, and Builder in Python.

```python
from __future__ import annotations
import threading

# --- SINGLETON ---
# Thread-safe Singleton using Double-checked locking
class DatabasePool:
    _instance: DatabasePool | None = None
    _lock = threading.Lock()

    def __new__(cls, *args, **kwargs):
        if cls._instance is None:
            with cls._lock:               # prevent race condition
                if cls._instance is None: # double check inside lock
                    cls._instance = super().__new__(cls)
        return cls._instance

    def __init__(self, host="localhost", size=10):
        if not hasattr(self, "_initialized"):
            self.host = host
            self.pool_size = size
            self._initialized = True

# --- FACTORY METHOD ---
from abc import ABC, abstractmethod

class Notification(ABC):
    @abstractmethod
    def send(self, message: str, recipient: str) -> bool: ...

class EmailNotification(Notification):
    def send(self, message, recipient):
        print(f"Email to {recipient}: {message}")
        return True

class SMSNotification(Notification):
    def send(self, message, recipient):
        print(f"SMS to {recipient}: {message}")
        return True

class PushNotification(Notification):
    def send(self, message, recipient):
        print(f"Push to device {recipient}: {message}")
        return True

class NotificationFactory:
    _registry: dict[str, type[Notification]] = {}

    @classmethod
    def register(cls, channel: str):
        def decorator(notification_cls):
            cls._registry[channel] = notification_cls
            return notification_cls
        return decorator

    @classmethod
    def create(cls, channel: str) -> Notification:
        if channel not in cls._registry:
            raise ValueError(f"Unknown channel: {channel}")
        return cls._registry[channel]()

NotificationFactory._registry = {
    "email": EmailNotification,
    "sms": SMSNotification,
    "push": PushNotification,
}

n = NotificationFactory.create("email")
n.send("Your order shipped!", "user@example.com")

# --- BUILDER ---
from dataclasses import dataclass, field

@dataclass
class QueryBuilder:
    _table: str = ""
    _conditions: list = field(default_factory=list)
    _columns: list = field(default_factory=lambda: ["*"])
    _limit: int | None = None
    _order_by: str | None = None

    def from_table(self, table: str) -> QueryBuilder:
        self._table = table
        return self   # fluent interface

    def select(self, *columns: str) -> QueryBuilder:
        self._columns = list(columns)
        return self

    def where(self, condition: str) -> QueryBuilder:
        self._conditions.append(condition)
        return self

    def limit(self, n: int) -> QueryBuilder:
        self._limit = n
        return self

    def order_by(self, column: str) -> QueryBuilder:
        self._order_by = column
        return self

    def build(self) -> str:
        cols = ", ".join(self._columns)
        query = f"SELECT {cols} FROM {self._table}"
        if self._conditions:
            query += " WHERE " + " AND ".join(self._conditions)
        if self._order_by:
            query += f" ORDER BY {self._order_by}"
        if self._limit:
            query += f" LIMIT {self._limit}"
        return query

query = (
    QueryBuilder()
    .from_table("users")
    .select("id", "name", "email")
    .where("active = TRUE")
    .where("age > 18")
    .order_by("name")
    .limit(100)
    .build()
)
print(query)
```

---

### Q2. Observer, Strategy, and Command patterns.

```python
from abc import ABC, abstractmethod
from typing import Callable

# --- OBSERVER ---
class EventEmitter:
    def __init__(self):
        self._listeners: dict[str, list[Callable]] = {}

    def on(self, event: str, callback: Callable):
        self._listeners.setdefault(event, []).append(callback)
        return self

    def off(self, event: str, callback: Callable):
        if event in self._listeners:
            self._listeners[event].remove(callback)

    def emit(self, event: str, *args, **kwargs):
        for cb in self._listeners.get(event, []):
            cb(*args, **kwargs)

class OrderService(EventEmitter):
    def place_order(self, order_id: str, amount: float):
        print(f"Order {order_id} placed for ₹{amount:,}")
        self.emit("order_placed", order_id=order_id, amount=amount)

service = OrderService()
service.on("order_placed", lambda **kw: print(f"Email sent for {kw['order_id']}"))
service.on("order_placed", lambda **kw: print(f"Inventory updated for {kw['order_id']}"))
service.place_order("ORD-001", 5000)

# --- STRATEGY ---
from typing import Protocol

class SortStrategy(Protocol):
    def sort(self, data: list) -> list: ...

class BubbleSort:
    def sort(self, data):
        data = data.copy()
        n = len(data)
        for i in range(n):
            for j in range(n - i - 1):
                if data[j] > data[j+1]:
                    data[j], data[j+1] = data[j+1], data[j]
        return data

class QuickSort:
    def sort(self, data):
        if len(data) <= 1:
            return data
        pivot = data[len(data) // 2]
        left  = [x for x in data if x < pivot]
        mid   = [x for x in data if x == pivot]
        right = [x for x in data if x > pivot]
        return self.sort(left) + mid + self.sort(right)

class Sorter:
    def __init__(self, strategy: SortStrategy):
        self._strategy = strategy

    def set_strategy(self, strategy: SortStrategy):
        self._strategy = strategy

    def sort(self, data):
        return self._strategy.sort(data)

sorter = Sorter(QuickSort())
sorter.sort([3, 1, 4, 1, 5, 9])

sorter.set_strategy(BubbleSort())
sorter.sort([3, 1, 4, 1, 5, 9])   # different algorithm, same interface

# --- COMMAND ---
from abc import abstractmethod

class Command(ABC):
    @abstractmethod
    def execute(self): ...

    @abstractmethod
    def undo(self): ...

class AddTextCommand(Command):
    def __init__(self, doc, text, position):
        self.doc = doc
        self.text = text
        self.position = position

    def execute(self):
        self.doc.insert(self.position, self.text)

    def undo(self):
        self.doc.delete(self.position, self.position + len(self.text))

class CommandHistory:
    def __init__(self):
        self._history: list[Command] = []

    def execute(self, command: Command):
        command.execute()
        self._history.append(command)

    def undo(self):
        if self._history:
            command = self._history.pop()
            command.undo()
```

---

### Q3. Adapter and Proxy patterns in Python.

```python
# --- ADAPTER ---
# Make incompatible interfaces work together

class LegacyPaymentGateway:
    def pay_with_card(self, card_number, amount, currency):
        print(f"Legacy: charging {card_number} for {amount} {currency}")
        return {"status": "charged", "ref": "REF123"}

class ModernPaymentInterface(ABC):
    @abstractmethod
    def process_payment(self, payment_method: dict, amount: float) -> dict: ...

class LegacyPaymentAdapter(ModernPaymentInterface):
    def __init__(self, legacy: LegacyPaymentGateway):
        self._legacy = legacy

    def process_payment(self, payment_method, amount):
        card = payment_method.get("card_number")
        currency = payment_method.get("currency", "INR")
        result = self._legacy.pay_with_card(card, amount, currency)
        return {"success": result["status"] == "charged", "reference": result["ref"]}

# Now legacy and new gateways work the same way

# --- PROXY ---
# Control access to another object

class ExpensiveService:
    def fetch_data(self, key: str) -> str:
        import time; time.sleep(0.5)  # simulates expensive operation
        return f"data_for_{key}"

class CachingProxy:
    def __init__(self, service: ExpensiveService):
        self._service = service
        self._cache: dict[str, str] = {}

    def fetch_data(self, key: str) -> str:
        if key not in self._cache:
            print(f"Cache miss for {key} — fetching...")
            self._cache[key] = self._service.fetch_data(key)
        else:
            print(f"Cache hit for {key}")
        return self._cache[key]

real_service = ExpensiveService()
proxy = CachingProxy(real_service)
proxy.fetch_data("config")   # Cache miss
proxy.fetch_data("config")   # Cache hit
```
