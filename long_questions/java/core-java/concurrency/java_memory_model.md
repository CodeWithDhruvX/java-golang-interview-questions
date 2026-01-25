# Java Memory Model (JMM)

## 1. Overview
The JMM defines how threads interact through memory. It details how and when changes made to variables by one thread become visible to other threads.

## 2. Key Concepts

### Atomicity
*   Operations that are indivisible.
*   Reading/writing `int`, `boolean` is atomic.
*   `long` and `double` (64-bit) are *not* guaranteed to be atomic on 32-bit systems (unless declared `volatile`).
*   `i++` is **NOT** atomic (Read -> Modify -> Write).

### Visibility
*   Changes made by one thread might only exist in CPU Cache and not be visible to others.
*   **Solution**: Use `volatile` or `synchronized`.

### Ordering (Happens-Before Relationship)
*   Compilers and CPUs reorder instructions for performance. JMM guarantees specific ordering rules.
*   **Happens-Before Rule**: If Action A happens-before Action B, then B sees the results of A.

## 3. The `volatile` Keyword
*   **Guarantees Visibility**: Reads/Writes go directly to Main Memory (RAM), bypassing CPU Cache.
*   **Guarantees Ordering**: Prevents instruction reordering around the variable.
*   **Does NOT guarantee Atomicity**: `volatile int count; count++` is still unsafe.

## 4. `synchronized` vs `volatile`

| Feature | volatile | synchronized |
| :--- | :--- | :--- |
| **Visibility** | Yes | Yes |
| **Atomicity** | No | Yes |
| **Blocking** | Non-blocking | Blocks other threads |
| **Overhead** | Low | High |

## 5. Singleton Pattern (Double Checked Locking)
Classic JMM application.

```java
class Singleton {
    private static volatile Singleton instance; // volatile is crucial!

    public static Singleton getInstance() {
        if (instance == null) {
            synchronized (Singleton.class) {
                if (instance == null) {
                    instance = new Singleton();
                }
            }
        }
        return instance;
    }
}
```
*Why `volatile`?* Without it, `instance = new Singleton()` can be reordered. The reference might be assigned *before* the constructor finishes, causing other threads to see a partially initialized object.

## 6. Interview Questions
1.  **Can we use `volatile` for a counter variable?**
    *   *Ans*: No. `count++` is not atomic. Use `AtomicInteger`.
2.  **What is the "Happens-Before" relationship?**
    *   *Ans*: A guarantee that memory writes by one specific statement are visible to another specific statement. Examples: `thread.start()` happens-before `thread.run()`. Write to `volatile` happens-before Read of `volatile`.
3.  **What is False Sharing?**
    *   *Ans*: When two threads write to different variables that happen to be on the same CPU Cache Line. This causes "ping-ponging" of the cache line between cores, degrading performance. Solution: Padding variables.
