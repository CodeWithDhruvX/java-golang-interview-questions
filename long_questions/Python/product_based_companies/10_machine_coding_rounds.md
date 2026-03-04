# ⚡ 10 — Machine Coding Rounds
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 What to Expect
- Design and implement a mini-application end-to-end
- Evaluated on: code structure, OOP design, edge cases, extensibility, tests
- Typical: 60–90 minute timed challenge

---

## 🧩 Coding Challenges

### Challenge 1: Implement a Thread-Safe LRU Cache

```python
from collections import OrderedDict
import threading

class LRUCache:
    """
    Thread-safe LRU Cache with O(1) get and put.
    Uses OrderedDict (doubly linked list + hash map) internally.
    """

    def __init__(self, capacity: int):
        if capacity <= 0:
            raise ValueError("Capacity must be positive")
        self._capacity = capacity
        self._cache: OrderedDict[int, int] = OrderedDict()
        self._lock = threading.RLock()    # reentrant lock for thread safety

    def get(self, key: int) -> int:
        with self._lock:
            if key not in self._cache:
                return -1
            self._cache.move_to_end(key)   # mark as most recently used
            return self._cache[key]

    def put(self, key: int, value: int) -> None:
        with self._lock:
            if key in self._cache:
                self._cache.move_to_end(key)
            self._cache[key] = value
            if len(self._cache) > self._capacity:
                self._cache.popitem(last=False)   # evict LRU (first item)

    def __len__(self):
        return len(self._cache)

    def __repr__(self):
        return f"LRUCache(capacity={self._capacity}, size={len(self)}, items={list(self._cache.keys())})"

# Test
cache = LRUCache(3)
cache.put(1, 10)
cache.put(2, 20)
cache.put(3, 30)
cache.get(1)       # 10 — 1 is now MRU
cache.put(4, 40)   # evicts 2 (LRU)
cache.get(2)       # -1 (evicted)
print(cache)       # LRUCache(capacity=3, size=3, items=[3, 1, 4])

# Pytest tests
import pytest

def test_lru_basic():
    c = LRUCache(2)
    c.put(1, 1)
    c.put(2, 2)
    assert c.get(1) == 1
    c.put(3, 3)          # evicts key 2
    assert c.get(2) == -1
    assert c.get(3) == 3

def test_lru_update_existing():
    c = LRUCache(2)
    c.put(1, 1)
    c.put(1, 100)        # update
    assert c.get(1) == 100

def test_lru_invalid_capacity():
    with pytest.raises(ValueError):
        LRUCache(0)
```

---

### Challenge 2: Design a Parking Lot System

```python
from enum import Enum, auto
from dataclasses import dataclass, field
from datetime import datetime
from abc import ABC, abstractmethod
import threading

class VehicleType(Enum):
    MOTORCYCLE = auto()
    CAR = auto()
    TRUCK = auto()

@dataclass
class Vehicle:
    license_plate: str
    vehicle_type: VehicleType

@dataclass
class ParkingSpot:
    spot_id: str
    spot_type: VehicleType
    is_occupied: bool = False
    current_vehicle: Vehicle | None = None

    def park(self, vehicle: Vehicle):
        if self.is_occupied:
            raise RuntimeError(f"Spot {self.spot_id} already occupied")
        self.is_occupied = True
        self.current_vehicle = vehicle

    def release(self):
        vehicle = self.current_vehicle
        self.is_occupied = False
        self.current_vehicle = None
        return vehicle

@dataclass
class Ticket:
    ticket_id: str
    vehicle: Vehicle
    spot: ParkingSpot
    entry_time: datetime = field(default_factory=datetime.now)

class PricingStrategy(ABC):
    @abstractmethod
    def calculate(self, hours: float, vehicle_type: VehicleType) -> float: ...

class HourlyPricing(PricingStrategy):
    RATES = {
        VehicleType.MOTORCYCLE: 20,
        VehicleType.CAR: 40,
        VehicleType.TRUCK: 80,
    }

    def calculate(self, hours, vehicle_type):
        import math
        return math.ceil(hours) * self.RATES[vehicle_type]

class ParkingLot:
    def __init__(self, name: str, spots: list[ParkingSpot], pricing: PricingStrategy):
        self.name = name
        self._spots = spots
        self._pricing = pricing
        self._active_tickets: dict[str, Ticket] = {}
        self._lock = threading.Lock()
        self._ticket_counter = 0

    def _find_spot(self, vehicle_type: VehicleType) -> ParkingSpot | None:
        # Type hierarchy: motorcycle can use any spot, truck needs truck spot
        compatible = {
            VehicleType.MOTORCYCLE: [VehicleType.MOTORCYCLE, VehicleType.CAR, VehicleType.TRUCK],
            VehicleType.CAR: [VehicleType.CAR, VehicleType.TRUCK],
            VehicleType.TRUCK: [VehicleType.TRUCK],
        }
        for spot_type in compatible[vehicle_type]:
            for spot in self._spots:
                if spot.spot_type == spot_type and not spot.is_occupied:
                    return spot
        return None

    def park(self, vehicle: Vehicle) -> Ticket:
        with self._lock:
            spot = self._find_spot(vehicle.vehicle_type)
            if not spot:
                raise RuntimeError(f"No spots available for {vehicle.vehicle_type.name}")
            spot.park(vehicle)
            self._ticket_counter += 1
            ticket_id = f"TKT-{self._ticket_counter:06d}"
            ticket = Ticket(ticket_id=ticket_id, vehicle=vehicle, spot=spot)
            self._active_tickets[ticket_id] = ticket
            return ticket

    def exit(self, ticket_id: str) -> float:
        with self._lock:
            ticket = self._active_tickets.pop(ticket_id, None)
            if not ticket:
                raise ValueError(f"Ticket {ticket_id} not found")
            ticket.spot.release()
            hours = (datetime.now() - ticket.entry_time).seconds / 3600
            return self._pricing.calculate(hours, ticket.vehicle.vehicle_type)

    @property
    def available_spots(self) -> dict[VehicleType, int]:
        with self._lock:
            result = {t: 0 for t in VehicleType}
            for spot in self._spots:
                if not spot.is_occupied:
                    result[spot.spot_type] += 1
            return result
```

---

### Challenge 3: Event-Driven Notification System

```python
from typing import Callable, Any
from dataclasses import dataclass, field
from datetime import datetime
from enum import Enum
import threading
import queue

class EventType(Enum):
    ORDER_PLACED    = "order.placed"
    ORDER_SHIPPED   = "order.shipped"
    PAYMENT_SUCCESS = "payment.success"
    PAYMENT_FAILED  = "payment.failed"
    USER_SIGNUP     = "user.signup"

@dataclass
class Event:
    type: EventType
    payload: dict
    created_at: datetime = field(default_factory=datetime.now)
    event_id: str = field(default_factory=lambda: f"evt-{id(object())}")

Handler = Callable[[Event], None]

class EventBus:
    """Async event bus with background worker thread"""

    def __init__(self):
        self._handlers: dict[EventType, list[Handler]] = {}
        self._queue: queue.Queue[Event | None] = queue.Queue()
        self._lock = threading.RLock()
        self._worker = threading.Thread(target=self._process_events, daemon=True)
        self._worker.start()

    def subscribe(self, event_type: EventType, handler: Handler) -> None:
        with self._lock:
            self._handlers.setdefault(event_type, []).append(handler)

    def unsubscribe(self, event_type: EventType, handler: Handler) -> None:
        with self._lock:
            if event_type in self._handlers:
                self._handlers[event_type].remove(handler)

    def publish(self, event: Event) -> None:
        self._queue.put(event)   # non-blocking — returns immediately

    def _process_events(self):
        while True:
            event = self._queue.get()
            if event is None:
                break   # shutdown signal
            with self._lock:
                handlers = list(self._handlers.get(event.type, []))
            for handler in handlers:
                try:
                    handler(event)
                except Exception as e:
                    print(f"Handler error for {event.type}: {e}")

    def shutdown(self):
        self._queue.put(None)
        self._worker.join()

# Usage
bus = EventBus()

def send_confirmation_email(event: Event):
    print(f"Email: Order {event.payload['order_id']} confirmed!")

def update_inventory(event: Event):
    print(f"Inventory: Deducting stock for {event.payload['items']}")

def notify_warehouse(event: Event):
    print(f"Warehouse: New order {event.payload['order_id']} received")

bus.subscribe(EventType.ORDER_PLACED, send_confirmation_email)
bus.subscribe(EventType.ORDER_PLACED, update_inventory)
bus.subscribe(EventType.ORDER_PLACED, notify_warehouse)

bus.publish(Event(
    type=EventType.ORDER_PLACED,
    payload={"order_id": "ORD-001", "items": ["Laptop", "Mouse"], "amount": 55000}
))

import time; time.sleep(0.1)   # let worker process
bus.shutdown()
```
