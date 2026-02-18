# 22. Java Programs (Collections & Advanced Concepts)

**Q: Iterate ArrayList**
> "You have 3 main ways, and interviewers look for the last one.
>
> 1.  **Old School**: `for (int i = 0; i < list.size(); i++)`. Good if you need the index.
> 2.  **Enhanced For-Loop**: `for (String s : list)`. Cleanest syntax.
> 3.  **Java 8 Streams/ForEach**: `list.forEach(System.out::println);`. This shows you know modern Java."

**Indepth:**
> **Performance**: An index-based loop (`for(i=0)`) is actually extremely slow for a `LinkedList` (O(n^2)) because `get(i)` scans from the start every time. `forEach` or `Iterator` handles the linked structure correctly (O(n)).


---

**Q: Convert Array to ArrayList**
> "Be careful here.
> `Arrays.asList(arr)` returns a **fixed-size** list. You can't add to it.
>
> To get a fully modifiable `ArrayList`, you must wrap it:
> `List<String> list = new ArrayList<>(Arrays.asList(arr));`.
>
> For primitives (`int[]`), `Arrays.asList` doesn't work well (it creates a List of arrays). You need `IntStream`: `IntStream.of(arr).boxed().collect(Collectors.toList());`."

**Indepth:**
> **Generics**: `Arrays.asList` returns a `List<T>`. If you pass `int[]`, T becomes `int[]`, so you get a `List<int[]>` (a list of arrays). `Integer[]` works fine because `T` becomes `Integer`.


---

**Q: Iterate and Modify (Remove)**
> "If you try to remove an item inside a `for-each` loop, you get `ConcurrentModificationException`.
>
> **The Fix**: Use an **Iterator**.
> ```java
> Iterator<String> it = list.iterator();
> while(it.hasNext()) {
>     if (it.next().equals(\"DeleteMe\")) {
>         it.remove(); // Safe!
>     }
> }
> ```
> In Java 8+, simply use: `list.removeIf(s -> s.equals("DeleteMe"));`."

**Indepth:**
> **Internal**: When you call `next()`, the iterator checks a `modCount` variable. If the collection's `modCount` doesn't match the iterator's expected `modCount` (meaning someone else changed the list), it explodes instantly. `iterator.remove()` updates both counts safely.


---

**Q: Sort HashMap by Value**
> "HashMaps are unsorted. To sort by value, you need a List.
>
> 1.  Get the Entry Set: `list = new ArrayList<>(map.entrySet());`
> 2.  Sort the list with a custom Comparator: `list.sort(Map.Entry.comparingByValue());`
> 3.  (Optional) Put it back into a `LinkedHashMap` to preserve that sorted order."

**Indepth:**
> **Complexity**: Sorting a Map by value is O(n log n) because you must extract all entries into a list and sort the list. There is no way to make the Map itself sort by value permanently without custom data structures.


---

**Q: Merge Two Lists**
> "The simple way:
> `list1.addAll(list2);`
>
> The Stream way (if you want a new list without modifying originals):
> `Stream.concat(list1.stream(), list2.stream()).collect(Collectors.toList());`
>
> **Intersection** (Common elements):
> `list1.retainAll(list2);` (Modifies list1 to keep only matches)."

**Indepth:**
> **Mutability**: `List.addAll` modifies the first list in place. `Stream.concat` creates a new stream (and eventually a new list), leaving the originals untouched. Functional programming prefers immutability.


---

**Q: Functional Interface & Lambda**
> "A Functional Interface has **exactly one abstract method**. Examples: `Runnable`, `Callable`, `Comparator`.
>
> A **Lambda** is just a shortcut to implement that interface without writing a bulky anonymous class.
> Instead of `new Runnable() { public void run() { ... } }`, you write `() -> { ... }`.
> It makes code concise and enables functional programming patterns."

**Indepth:**
> **SAM**: Functional Interfaces are also called SAM types (Single Abstract Method). The `@FunctionalInterface` annotation is optional but recommended—it stops colleagues from accidentally adding a second abstract method and breaking your lambdas.


---

**Q: Stream API: Filter, Map, Reduce**
> "The Holy Trinity of Streams:
>
> 1.  **Filter**: Logic to say 'Keep this, throw that away'. Returns a boolean.
>     *   `stream.filter(n -> n % 2 == 0)` (Keep evens).
> 2.  **Map**: Transform data. Input Type -> Output Type.
>     *   `stream.map(n -> n * n)` (Square each number).
> 3.  **Reduce/Collect**: Aggregate results.
>     *   `.collect(Collectors.toList())` or `.reduce(0, Integer::sum)`."

**Indepth:**
> **Lazy Evaluation**: Streams are lazy. `stream.filter().map()` doesn't actually do anything until you call a terminal operation like `.collect()`. This allows optimization (loop fusion).


---

**Q: GroupingBy (Stream)**
> "How do you group a list of Strings by their length?
>
> `Map<Integer, List<String>> groups = list.stream().collect(Collectors.groupingBy(String::length));`
>
> This one line of code replaces about 10 lines of old-school loops and if-checks. It is extremely powerful for report generation or data analysis."

**Indepth:**
> **Under the Hood**: `groupingBy` uses a `HashMap` (or `TreeMap` if requested) to store the groups. It iterates the stream once, applies the classifier function to each element, and adds it to the corresponding list bucket.


---

**Q: Producer-Consumer Problem**
> "This is a concurrency pattern where one thread (Producer) keeps adding work to a buffer, and another (Consumer) keeps taking it.
>
> The challenge is coordination:
> *   If buffer is full, Producer must wait.
> *   If buffer is empty, Consumer must wait.
>
> **Modern Solution**: Don't use `wait()` and `notify()`. Use a `BlockingQueue` (like `ArrayBlockingQueue`).
> *   Producer calls `queue.put()` (blocks if full).
> *   Consumer calls `queue.take()` (blocks if empty).
> It handles all the thread safety and locking for you."

**Indepth:**
> **Blocking**: `BlockingQueue` uses `ReentrantLock` and `Condition` variables (`notFull`, `notEmpty`) internally. It puts the `put()` thread to sleep if the queue is full, and wakes it up only when space becomes available.


---

**Q: Deadlock Scenario**
> "Deadlock happens when two threads hold locks the other one wants.
>
> Thread 1: Holds Lock A, waits for Lock B.
> Thread 2: Holds Lock B, waits for Lock A.
> Result: They wait forever.
>
> **Prevention**: Always acquire locks in a consistent order (e.g., Always Lock A before Lock B)."

**Indepth:**
> **Analysis**: If your app hangs, take a Thread Dump (jstack). Look for "Found one Java-level deadlock". It will show you exactly which threads are holding which locks.


---

**Q: Optional Class**
> "`Optional` is a wrapper that might contain a value or might be empty. It was created to stop `NullPointerException`.
>
> Instead of: `if (user != null) { print(user.getName()); }`
> You do: `userVal.ifPresent(u -> print(u.getName()));`
>
> Best practice: Never use `Optional.get()` without checking. Use `.orElse("Default")` or `.orElseThrow()`."

**Indepth:**
> **Primitive Optionals**: Java also has `OptionalInt`, `OptionalDouble`, etc. checking to avoid boxing overhead when working with primitives.

