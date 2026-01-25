# Process vs Thread

## 1. Definitions

### Process
A **Process** is an instance of a computer program that is being executed. It contains the program code and its current activity. Each process has its own isolated memory space.
*   **Analogy**: A detailed spreadsheet file open in Excel. Excel is the program, but *that specific open file* window is a process.

### Thread
A **Thread** is the smallest unit of processing that can be performed in an OS. It exists *within* a process. A process can contain multiple threads.
*   **Analogy**: The formulas calculating automatically in that spreadsheet, while you type in another cell, and the auto-save running in the background. All these are threads within the Excel process.

## 2. Key Differences Table

| Feature | Process | Thread |
| :--- | :--- | :--- |
| **Memory** | Isolated. Does not share memory with other processes. | Shared. Shares memory (Layout, Heap, Code) with other threads in the same process. |
| **Creation Cost** | High (Heavyweight). Requires OS to allocate new memory segments. | Low (Lightweight). Uses existing process resources. |
| **Communication** | Slow. Needs Inter-Process Communication (IPC) like Pipes, Sockets, Files. | Fast. Can read/write same variables directly. |
| **Context Switching** | Slow. OS must save/reload distinct memory maps and registers. | Fast. Only registers and stack need saving. |
| **Failure Scope** | If a process crashes, others are usually unaffected. | If a thread crashes (e.g., segfault), it can crash the entire process. |

## 3. Shared vs Private Resources

### Shared by all Threads in a Process
*   **Heap Memory**: Dynamic memory allocation.
*   **Global Variables**: Static data.
*   **Code Segment**: The instructions being executed.
*   **File Descriptors**: Open files.

### Private to Each Thread
*   **Stack**: Local variables and function call history.
*   **Registers**: CPU register values (Program Counter, Stack Pointer).
*   **Thread Local Storage (TLS)**: Data specific to the thread.

## 4. Context Switching using Stacks
When the CPU switches from Thread A to Thread B:
1.  Saves Thread A's **Program Counter** and **Registers** to Thread A's control block (TCB).
2.  Loads Thread B's registers from its TCB.
3.  Execution resumes where Thread B left off.

## 5. Interview Questions
1.  **Why is creating a thread cheaper than a process?**
    *   *Ans*: Threads share the same address space. You don't need to create new page tables or duplicate resources like file descriptors.
2.  **What happens if one thread crashes?**
    *   *Ans*: Usually limits the whole process because they share the memory address space. An illegal operation (segfault) by one thread triggers a signal that terminates the process.
3.  **Chrome uses a "Process per Tab" model. Why?**
    *   *Ans*: Stability and Security. If one tab (webpage) crashes, it's an isolated process, so it doesn't bring down the whole browser. Also, it prevents malicious sites from scanning memory of other tabs.
