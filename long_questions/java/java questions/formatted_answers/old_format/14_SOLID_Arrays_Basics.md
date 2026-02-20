# 14. SOLID, Design Patterns, and Arrays

**Q: SOLID - Open/Closed Principle (OCP)**
> "The Open/Closed Principle states that software entities (classes, modules, functions) should be **open for extension, but closed for modification**.
>
> Ideally, when you need to add a new feature, you shouldn't have to touch the existing, working code (risking bugs). Instead, you should be able to extend the existing code by creating a new class.
>
> For example: If you have a `NotificationService` that sends Emails, and you want to add SMS, you shouldn't modify the `NotificationService` class. You should have an interface `Notification` and create a new class `SMSNotification` that implements it. The original code remains untouched."

**Indepth:**
>

---

**Q: Comparable vs Comparator implementation?**
> "Both are interfaces used for sorting, but they answer different questions.
>
> **Comparable** is about **natural ordering**. If a class implements `Comparable` (like `String` or `Integer`), it means instances of that class 'know' how to sort themselves. You implement `compareTo(Object o)`.
>
> **Comparator** is about **custom ordering**. If you want to sort objects in a way that isn't their 'natural' order (like sorting Strings by length instead of alphabetically), you create a separate class (or lambda) that implements `Comparator`. You implement `compare(Object o1, Object o2)`.
>
> Use `Comparable` for the default sort. Use `Comparator` for special cases."

**Indepth:**
> **Performance**: Since arrays use contiguous memory locations, they are extremely cache-friendly (spatial locality). Accessing `arr[i]` is O(1) and very fast due to CPU prefetching.


---

**Q: Reference Types: Strong, Soft, Weak, Phantom**
> "In Java, not all references are equal. The Garbage Collector treats them differently.
>
> 1.  **Strong Reference**: The default. `Dog d = new Dog();`. As long as you hold this, the GC will **never** collect it.
> 2.  **Soft Reference**: Use this for memory-sensitive caches. The GC will only collect these objects if the JVM is running out of memory.
> 3.  **Weak Reference**: Use this for metadata (like `WeakHashMap`). The GC collects these as soon as it sees them, provided no strong references exist.
> 4.  **Phantom Reference**: The weakest link. You use this to track when an object has been literally removed from memory, usually to perform some post-mortem cleanup. You rarely need this."

**Indepth:**
>

---

**Q: How to use ShutdownHook?**
> "A **ShutdownHook** is a thread that the JVM runs just before it shuts down (whether normally or via Ctrl+C).
>
> It's your last chance to say goodbye. You use it to close database connections, save state, or release resources gracefully.
>
> You register it like this: `Runtime.getRuntime().addShutdownHook(new Thread(() -> { ... }));`. But be careful—you can't rely on it running if the JVM crashes hard (like `kill -9`)."

**Indepth:**
> **Real World**: OCP is why plugins work. An IDE like IntelliJ allows you to add plugins (extension) without rewriting the core IDE code (modification).


---

**Q: Dependency Injection (Manual Implementation)**
> "Dependency Injection (DI) sounds complex, but it's just passing variables.
>
> Without DI:
> ```java
> class Car {
>     private Engine engine = new V8Engine(); // Car is hardcoded to V8Engine
> }
> ```
>
> With DI (Manual):
> ```java
> class Car {
>     private Engine engine;
>     public Car(Engine engine) { // You pass the engine in
>         this.engine = engine;
>     }
> }
> // Usage
> Car car = new Car(new ElectricEngine());
> ```
> By passing the dependency in (via constructor), you decouple the classes. Basic DI is just using constructors properly."

**Indepth:**
>

---

**Q: How do you declare, initialize, and copy an array?**
> "Declaration is simple: `int[] numbers;`.
> Initialization can be static (`{1, 2, 3}`) or dynamic (`new int[5]`).
>
> Copying is where people trip up.
> `int[] b = a;` is **NOT** a copy. It's just a new reference to the *same* array.
>
> To actually copy the data, you use `Arrays.copyOf(original, newLength)` or `System.arraycopy()`. These create a fresh array in memory with independent data."

**Indepth:**
>

---

**Q: Arrays.copyOf() vs System.arraycopy()**
> "**Arrays.copyOf()** is the readable, developer-friendly way. It creates a new array for you and returns it. It's great for readability.
>
> **System.arraycopy()** is the low-level, high-performance way. You have to create the destination array yourself first. It looks scary (`src, srcPos, dest, destPos, length`), but it allows you to copy into the *middle* of an existing array, which `copyOf` can't do.
>
> Under the hood? `Arrays.copyOf` actually calls `System.arraycopy`."

**Indepth:**
>

---

**Q: Shallow Copy vs Deep Copy of Arrays**
> "If you have an array of primitives (`int[]`), a standard copy is a 'deep' copy because the values are just numbers.
>
> But if you have an array of objects (`Person[]`), a standard copy (`clone()` or `copyOf()`) is a **Shallow Copy**.
> Use `clone()` on `Person[]`, and you get a new array, but it's filled with references to the **same** Person objects. If you change a Person's name in one array, it changes in the other too!
>
> To do a **Deep Copy**, you must loop through the array and manually `new Person()` for every single element."

**Indepth:**
>

---

**Q: Max/Min/Reverse/Rotate/Duplicates in Arrays**
> "These are classic logic problems.
>
> *   **Max/Min**: Initialize `max` to the first element. Loop through the rest. If `current > max`, update `max`.
> *   **Reverse**: Use two pointers. One at start (`0`), one at end (`length-1`). Swap them, move pointers towards the center until they meet.
> *   **Remove Duplicates**: If sorted, it's easy—just check if `current == previous`. If not sorted, simpler to dump everything into a `HashSet`."

**Indepth:**
>

---

**Q: Arrays.sort() vs Collections.sort()**
> "**Arrays.sort()** works on arrays (`int[]`, `String[]`). It uses a Dual-Pivot Quicksort for primitives (fast but unstable) and Timsort (MergeSort variant) for Objects (stable).
>
> **Collections.sort()** works on Lists (`ArrayList`). Internally, it actually dumps the List into an array, calls `Arrays.sort()`, and then dumps it back into the List! So they use the same engine."

**Indepth:**
>

---

**Q: Arrays.binarySearch()**
> "Binary Search is super fast—O(log n)—but it has one golden rule: **The array must be sorted first!**
>
> If you run `binarySearch()` on an unsorted array, the result is undefined (garbage).
> It returns the index if found. If not found, it returns a negative number `-(insertionPoint) - 1`, telling you exactly where the element *would* go if you wanted to insert it while keeping the order."

**Indepth:**
>

---

**Q: Arrays.asList() caveats**
> "`Arrays.asList()` is a handy bridge between Arrays and Collections, but it's a trap.
>
> It returns a **fixed-size list** backed by the original array.
>
> 1.  You **cannot add or remove** elements. Calling `.add()` throws `UnsupportedOperationException`.
> 2.  Changes strictly write-through. If you set an element in the List, the original Array changes too.
>
> If you want a normal, modifiable ArrayList, you must wrap it: `new ArrayList<>(Arrays.asList(...))`."

**Indepth:**
>

---

**Q: 2D Arrays (Declaration, Traversal, Logic)**
> "A 2D array in Java is really just an 'array of arrays'.
> `int[][] matrix = new int[3][3];`.
>
> Since each row is its own object, you can actually have 'jagged arrays' where row 0 has length 5 and row 1 has length 2.
>
> Traversal is standard nested loops: Outer loop for rows (`i`), inner loop for columns (`j`).
> Common interview ops like **Rotation** usually involve:
> 1.  Transposing the matrix (swapping `[i][j]` with `[j][i]`).
> 2.  Reversing each row."

**Indepth:**
> **Searching**: `Arrays.binarySearch()` requires the array to be sorted first. If it's not sorted, the result is undefined.

