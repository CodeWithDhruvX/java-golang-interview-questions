# Concurrency, Race Conditions, and Deadlocks

## 1. Race Conditions
A **race condition** occurs when two or more threads access shared data and try to change it at the same time. Because the thread scheduling algorithm can swap between threads at any time, you don't know the order in which the threads will attempt to access the shared data.

### Example (Data Race)
```go
counter = 0

// Thread 1
counter++

// Thread 2
counter++
```
Steps taken by CPU: `Read -> Modify -> Write`. If Thread 1 reads `0`, pauses, Thread 2 reads `0`, both write `1`. Counter is 1, should be 2.

## 2. Deadlocks
A **Deadlock** happens when two or more threads are blocked forever, waiting for each other.

### The 4 Necessary Conditions (Coffman Conditions)
For a deadlock to occur, **ALL** four must be true simultaneously:
1.  **Mutual Exclusion**: At least one resource must be held in a non-shareable mode.
2.  **Hold and Wait**: A process holds a resource while waiting for another.
3.  **No Preemption**: Resources cannot be forcibly taken from a process.
4.  **Circular Wait**: A waits for B, B waits for C, ..., Z waits for A.

### Prevention vs Avoidance
*   **Prevention**: Break one of the 4 conditions. removing Mutual Exclusion (lock-free structures) or preventing Circular Wait (always order locks: Lock A before Lock B).
*   **Avoidance**: Use algorithms like **Banker's Algorithm** to check if a state is safe before granting a resource.

## 3. Concurrency Primitives

### Mutex (Mutual Exclusion)
*   A lock that allows only one thread to access a critical section.
*   **Usage**: Protecting a shared map or counter.

### Semaphore
*   A variable or abstract data type used to control access to a common resource by multiple processes in a concurrent system such as a multitasking operating system.
*   **Counting Semaphore**: Allows `N` threads to access. (e.g., Connection Pool with 10 connections).
*   **Binary Semaphore**: Same as Mutex (0 or 1).

### Monitor
*   High-level synchronization construct. It encapsulates data and methods, ensuring only one method can be active at a time. Java `synchronized` is based on Monitors.

## 4. Producer-Consumer Problem
Standard concurrency problem.
*   **Producer**: Adds items to a buffer.
*   **Consumer**: Removes items.
*   **Constraints**:
    *   Producer sleeps if buffer full.
    *   Consumer sleeps if buffer empty.
*   **Solution**: Use Semaphores (EmptySlots, FullSlots) or Condition Variables.

## 5. Interview Questions
1.  **Difference between Mutex and Semaphore?**
    *   *Ans*: Mutex is "ownership" model (only the thread that locked it can unlock it). Semaphore is a signaling mechanism (any thread can signal).
2.  **How to detect a Deadlock?**
    *   *Ans*: Build a resource allocation graph. If there is a cycle, there is a deadlock.
3.  **What is a Livelock?**
    *   *Ans*: Threads are not blocked, but they are constantly changing state in response to each other without doing useful work (e.g., two people trying to pass each other in a corridor and stepping left/right in sync).
