# Memory Management

## 1. Virtual Memory
**Virtual Memory** is a memory management technique that provides an idealized abstraction of the storage resources that are actually available on a given machine. It creates the illusion of a large, contiguous memory for each process, even if physical memory (RAM) is fragmented or smaller than the process needs.

*   **How it works**: The OS maps virtual addresses (seen by the application) to physical addresses (in RAM) using a **Page Table**.
*   **MMU (Memory Management Unit)**: Hardware component that handles this mapping translation at runtime.

## 2. Paging
**Paging** is the method of dividing Virtual Memory into fixed-size blocks called **Pages** and Physical Memory into blocks of the same size called **Frames**.

*   **Common Page Size**: 4KB.
*   **Page Table**: A data structure stored in RAM that maps Page Number -> Frame Number.
*   **TLB (Translation Lookaside Buffer)**: A hardware cache in the CPU that stores recent translations (Virtual -> Physical) to speed up access.

## 3. Segmentation
**Segmentation** divides memory into variable-sized segments based on logical divisions (Code, Stack, Heap).
*   *Contrast*: Paging is physical division (fixed size); Segmentation is logical division (variable size).

## 4. Page Faults & Thrashing
*   **Page Fault**: Happens when a program accesses a page that is mapped in address space but not currently loaded in physical RAM.
    *   *Action*: OS interrupts, finds the page on disk (Swap/Swapfile), loads it into an empty frame in RAM, updates Page Table, and resumes execution.
*   **Thrashing**: Occurs when the system spends more time swapping pages in and out (handling Page Faults) than executing actual instructions. Happens when RAM is full and processes actively need more memory than available.

## 5. Page Replacement Algorithms
When RAM is full and a new page needs to be loaded, the OS must evict an existing page. Which one?

1.  **FIFO (First In First Out)**: Evict oldest page.
    *   *Issue*: Belady's Anomaly (more frames can cause more faults).
2.  **LRU (Least Recently Used)**: Evict the page that hasn't been used for the longest time.
    *   *Pros*: Good performance, approximates optimal.
    *   *Implementation*: Doubly Linked List + Hash Map (O(1)).
3.  **Optimal**: Evict page that will not be used for the longest time in future. (Impossible to implement in real-time, used as benchmark).

## 6. Interview Questions
1.  **What is the difference between Virtual Memory and Physical Memory?**
    *   *Ans*: Virtual is logical addressing (0 to 4GB per process on 32-bit), physical is actual RAM slots. Virtual allows isolation and swapping.
2.  **What happens during a context switch regarding memory?**
    *   *Ans*: The OS switches the Page Directory Base Register (CR3 on x86) to point to the new process's Page Table. This effectively switches the entire address space. The TLB is usually flushed (making initial access slower).
3.  **Why is the Stack faster than the Heap?**
    *   *Ans*:
        *   **Allocation**: Stack is just moving a pointer (add/sub). Heap requires searching for a free block, updating free lists.
        *   **Locality**: Stack is contiguous, better cache locality. Heap is fragmented.
        *   **Deallocation**: Stack is automatic (function return). Heap requires GC or manual free.
