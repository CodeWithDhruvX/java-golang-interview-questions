# Coding: Real World & System Design - Interview Answers

> 🎯 **Focus:** These demonstrate you can build actual systems features.

### 1. Implement an LRU (Least Recently Used) Cache
"I don't need to invent this from scratch. Java's `LinkedHashMap` has a constructor specifically for this!
I just override `removeEldestEntry`. If the size exceeds capacity, it automatically kicks out the oldest item."

```java
class LRUCache<K, V> extends LinkedHashMap<K, V> {
    private final int capacity;

    public LRUCache(int capacity) {
        // true = access order (LRU), false = insertion order
        super(capacity, 0.75f, true); 
        this.capacity = capacity;
    }

    @Override
    protected boolean removeEldestEntry(Map.Entry<K, V> eldest) {
        return size() > capacity;
    }
}
```

---

### 2. Create an Immutable Class
"Immutability gives thread safety for free.
1. Make the class `final` so it can't be extended.
2. Make fields `private final`.
3. No Setters.
4. Initialize everything in constructor.
5. If a field is mutable (like a List), return a copy in the getter, not the original reference."

```java
public final class ImmutableUser {
    private final String name;
    private final List<String> roles;

    public ImmutableUser(String name, List<String> roles) {
        this.name = name;
        // Deep copy
        this.roles = new ArrayList<>(roles);
    }

    public String getName() { return name; }

    public List<String> getRoles() {
        // Return copy to protect internal state
        return new ArrayList<>(roles);
    }
}
```

---

### 3. Implement a Simple Rate Limiter
"I’ll use a Token Bucket algorithm concept.
I record the `lastRequestTime`. If the time elapsed is less than my allowed interval (e.g., 1000ms), I reject the request."

```java
class RateLimiter {
    private long lastRequestTime = 0;
    private final long intervalMillis;

    public RateLimiter(long intervalMillis) {
        this.intervalMillis = intervalMillis;
    }

    public synchronized boolean allowRequest() {
        long now = System.currentTimeMillis();
        if (now - lastRequestTime > intervalMillis) {
            lastRequestTime = now;
            return true;
        }
        return false;
    }
}
// Usage: RateLimiter limit = new RateLimiter(1000); // 1 req per sec
```

---

### 4. Flatten a Nested List (Deep Flat)
"I'll use recursion. If I encounter an element that is just an integer, I add it. If I encounter a List, I make a recursive call to flatten that sub-list."

```java
public List<Integer> flatten(List<Object> nested) {
    List<Integer> flat = new ArrayList<>();
    for (Object obj : nested) {
        if (obj instanceof Integer) {
            flat.add((Integer) obj);
        } else if (obj instanceof List) {
            flat.addAll(flatten((List<Object>) obj));
        }
    }
    return flat;
}
```
