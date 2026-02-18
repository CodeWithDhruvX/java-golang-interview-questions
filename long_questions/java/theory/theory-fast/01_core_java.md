# Core Java Interview Questions (1-25)

## Java Basics & Language Features

### 1. What are the key features introduced in Java 8?
"Java 8 was a massive shift because it introduced functional programming concepts to Java. The biggest headline feature was **Lambda expressions**, which let us treat functionality as a method argument, making code much more concise.

Along with that, we got the **Stream API**, which completely changed how we process collections—allowing us to do things like filter, map, and reduce in a declarative way instead of writing long `for` loops.

Other important additions were the `Optional` class to handle nulls better, **default methods** in interfaces so we can evolve interfaces without breaking implementations, and the new `CompletableFuture` for asynchronous programming.

Have you worked much with the Stream API, or do you stick to loops?"

### 2. How do lambda expressions work internally?
"So, even though lambdas look like syntactic sugar for anonymous inner classes, they’re actually quite different under the hood. When you compile a lambda, the compiler, creates a bytecode instruction called `invokedynamic`.

Instead of generating a separate `.class` file for every lambda (which would bloat the jar), the JVM creates the instance at runtime using a `MethodHandle`. This makes it more efficient in terms of memory and loading performance compared to the old anonymous class approach.

It’s basically the JVM saying, 'I’ll figure out how to represent this function when I actually need it.'"

### 3. What is the difference between `map`, `flatMap`, and `filter` in streams?
"These are the bread and butter of Stream processing.

`filter` is strictly for conditional checks. It takes an element, applies a boolean logic, and if it’s true, keeps it; otherwise, it discards it. The count of elements might go down, but the type stays the same.

`map` is for transformation. It takes an element and transforms it into something else—like converting a list of `User` objects into a list of just their `String` usernames. It’s always a one-to-one mapping.

`flatMap` is interesting—it flattens the structure. If you have a `List<List<String>>` and you use `map`, you’d still have nested lists. But `flatMap` flattens that out into a single continuous stream of strings. It’s a one-to-many mapping.

I use `flatMap` a lot when dealing with nested collections or Optional inside Optional."

### 4. When would you use `Optional`, and when should you avoid it?
"`Optional` is designed specifically to represent a return type that might return 'nothing'—eliminating the need for those dreaded null checks everywhere. I use it primarily as a return type for methods where 'no result' is a valid business case, like `findUserById`.

However, you should avoid using it as a field in a class or a parameter in a method. Using it as a field adds unnecessary memory overhead and makes serialization a pain since `Optional` isn’t serializable. And passing `Optional` as a parameter just makes the calling code clunky; it's better to just pass the value or null and handle it inside."

### 5. Difference between `==` and `equals()`?
"This is a classic trap. `==` checks for **reference equality**. It verifies if both variables point to the exact same object in memory location.

`equals()`, on the other hand, checks for **value equality** or logical equality. By default, in the `Object` class, `equals()` behaves just like `==`, but classes like `String` or `Integer` override it to compare the actual data inside.

So if I have two different String objects with the text 'hello', `==` will return false, but `equals()` will return true. That’s why strictly always using `.equals()` for object comparison is a best practice."

### 6. Why must `hashCode()` be overridden when `equals()` is overridden?
"It comes down to the contract used by hash-based collections like `HashMap` or `HashSet`.

These collections use the hash code to figure out which 'bucket' a key belongs to. If you override `equals()` to say two objects are logically the same, but you don't override `hashCode()`, they will generate different hash codes.

This means the HashMap might put them in different buckets, and you’ll never find the object again when you try to `get()` it. Essentially, if two objects are equal according to `equals()`, they **must** have the same `hashCode`."

### 7. Difference between `String`, `StringBuilder`, and `StringBuffer`?
"Everything revolves around mutability and thread safety here.

`String` is immutable. Every time you modify a string, like concatenating it, you’re essentially destroying the old object and creating a brand new one. This is great for caching and security but terrible for performance in a loop.

`StringBuilder` is mutable. You modify the existing object directly, which is much faster for string manipulation. However, it’s not thread-safe.

`StringBuffer` is the legacy version of StringBuilder. It’s also mutable but all its methods are `synchronized`, making it thread-safe. But that synchronization adds overhead, so in 99% of modern code where the string is local to a method, we just use `StringBuilder`."

### 8. What are immutable objects? Why are they important?
"An immutable object is one whose state cannot be changed after it’s created—like the `String` class.

They are incredibly important for concurrency. because if an object can't change, you don't need any synchronization to read it from multiple threads—it’s inherently thread-safe.

They also make great Map keys. If a Map key were mutable and its hashCode changed after you put it in the map, you’d essentially lose that entry forever. Immutability prevents that entire class of bugs."

### 9. What happens when you make a class immutable?
"To make a class truly immutable, you need to do a few things. First, make the class `final` so it can’t be subclassed and behavior altered.

Make all fields `private` and `final`, so they are set once via the constructor. Crucially, don’t provide any setters.

And if your class has mutable fields—like a List—you must return a deep copy or use `Collections.unmodifiableList` in the getter. Otherwise, someone could modify the list contents from outside, breaking the immutability."

## Collections & Generics

### 10. Difference between `ArrayList` and `LinkedList`?
"It’s mostly about memory layout and access patterns.

`ArrayList` is backed by a dynamic array. Random access—like `get(5)`—is super fast, O(1), because it just calculates the memory address. But adding or removing elements in the middle is slow because it has to shift all the subsequent elements.

`LinkedList` is a doubly linked list. Inserting or deleting is fast, O(1), because you just change the pointers. But finding an element is slow, O(n), because you have to traverse the chain from the start.

Honestly, in modern hardware, `ArrayList` is almost always faster due to cache locality, unless you're doing heavy insertions at the beginning of the list."

### 11. Difference between `HashMap` and `ConcurrentHashMap`?
"A regular `HashMap` isn't thread-safe. If multiple threads try to modify it at the same time, you can end up with data corruption or, in older Java versions, an infinite loop during resizing.

`ConcurrentHashMap` solves this. In older versions, it used 'segment locking'—locking only a small part of the map. In Java 8+, it’s even smarter; it uses `CAS` (Compare-And-Swap) operations and synchronized blocks only on the specific node being modified.

So, `ConcurrentHashMap` allows concurrent reads without locking and supports high-concurrency writes, making it much more scalable than wrapping a HashMap in `Collections.synchronizedMap`."

### 12. How does `HashMap` work internally?
"It’s an array of buckets (Node<K,V>). When you put a key-value pair, it calculates the hash of the key and maps it to an index in that array.

If that bucket is empty, it sits there. If there's already something there—a collision—it chains the new node using a LinkedList.

The cool part is the Java 8 improvement: if that list gets too long (more than 8 nodes), it converts it into a Balanced Tree (Red-Black Tree). This changes the worst-case lookup from O(n) to O(log n), preventing performance degradation during high collision scenarios."

### 13. What happens when two keys have the same hashcode?
"This is called a **collision**. The HashMap handles this by storing both items in the same bucket.

Initially, it uses a LinkedList. So the first entry points to the second entry. When you try to `get()` one of them, the HashMap finds the bucket using the hash, and then walks through the chain calling `.equals()` on the keys until it finds the exact match.

As I mentioned, if this collision chain gets too long, it upgrades to a Tree structure for performance."

### 14. Difference between `Comparable` and `Comparator`?
"Think of `Comparable` as the 'natural' ordering of an object. You implement it inside the class itself using the `compareTo` method. For example, a `Student` object might naturally sort by ID.

`Comparator` is for 'custom' ordering. It’s an external class or lambda. Maybe in one specific report, I want to sort Students by 'Grade', not ID. I wouldn't change the Student class; I’d just pass a generic `Comparator` to the sort method.

So `Comparable` is 'I compare myself', and `Comparator` is 'I compare these two things'."

### 15. Why are generics invariant in Java?
"Invariance means that a `List<String>` is **not** a subtype of `List<Object>`, even though String is a subtype of Object.

Java does this for type safety. If it allowed that, you could pass a `List<String>` into a method that accepts `List<Object>`, and that method could theoretically add an `Integer` to it. The compiler wouldn't catch it, but you'd get a nice runtime crash later when you tried to read that 'String'.

Invariance forces us to catch these errors at compile time."

### 16. What are wildcards (`? extends`, `? super`)?
"We use wildcards to gain flexibility back from invariance. This follows the **PECS** principle: Producer Extends, Consumer Super.

If you want to read data *from* a collection (it Produces data), use `? extends T`. This guarantees everything coming out is at least a T.

If you want to write data *into* a collection (it Consumes data), use `? super T`. This guarantees the collection can hold a T.

It’s confusing at first, but it allows us to create generic utilities that work across hierarchies."

### 17. Difference between `Set`, `List`, and `Map`?
"These are the core interfaces.

`List` is an ordered collection that allows duplicates. You care about the sequence, like a playlist.

`Set` is a collection of unique elements. It doesn't allow duplicates. Order typically isn't guaranteed (unless it's a LinkedHashSet or TreeSet). It’s great for things like filtering out duplicate user IDs.

`Map` isn't technically a Collection (doesn't extend Collection interface) but it maps Keys to Values. Keys must be unique, values can be duplicates. It’s essentially a dictionary."

### 18. When would you use `EnumMap` or `EnumSet`?
"You should use them whenever your keys (for Map) or elements (for Set) are Enums.

Internally, they are highly optimized. They don't use hashing. Instead, they use a simple array or a bit-vector based on the ordinal value of the Enum constants. This makes them extremely fast—much faster than a generic HashSet or HashMap—and very memory efficient."

## Multithreading & Concurrency

### 19. Difference between `Thread` and `Runnable`?
"`Thread` is a class you extend, and `Runnable` is an interface you implement.

You should almost always use `Runnable`. Extending `Thread` restricts your design because Java doesn't support multiple inheritance—if you extend Thread, you can't extend anything else.

Plus, `Runnable` separates the 'task' from the 'runner'. You can pass the same Runnable to a Thread, or an ExecutorService, or a thread pool. It’s much more flexible."

### 20. What is the Java Memory Model?
"The Java Memory Model (JMM) defines how threads interact with memory. The key concept is that every thread has its own local cache (stack), and there is a shared 'Main Memory' (heap).

The JMM specifies rules—like 'happens-before' relationships—that guarantee when changes made by one thread become visible to others. Without these rules (like using `volatile` or `synchronized`), one thread might update a variable, but another thread might keep reading a stale value from its local cache forever."

### 21. Difference between `volatile` and `synchronized`?
"`synchronized` does two things: it guarantees **atomicity** (only one thread enters the block at a time) and **visibility** (changes are flushed to main memory).

`volatile` is a lighter mechanism that only guarantees **visibility**. It tells the compiler, 'Don't cache this variable locally, always read/write from main memory.'

However, `volatile` doesn't guarantee atomicity. If you have `count++` with volatile, you can still have race conditions because read-modify-write is not a single atomic step. You’d need `AtomicInteger` or `synchronized` for that."

### 22. What is a deadlock? How do you prevent it?
"A deadlock happens when two threads are stuck waiting for each other forever. Thread A holds Resource 1 and wants Resource 2. Thread B holds Resource 2 and wants Resource 1. Neither can proceed.

To prevent it, the most common strategy is **consistent lock ordering**. Always acquire resources in the same global execution order. If everyone grabs Lock A before Lock B, a deadlock is impossible.

Another way is using `tryLock()` with a timeout, so a thread gives up if it can't get a lock, preventing the infinite wait."

### 23. What is `ExecutorService`?
"It’s a higher-level replacement for manually creating Threads. It provides a pool of threads and an API to submit tasks for execution.

It manages the lifecycle of threads for you. You don't create `new Thread()`; you just say `executor.submit(task)`.

It’s great because creating threads is expensive. ExecutorService reuses existing threads, which improves stability and performance. You typically use `Executors.newFixedThreadPool()` or `CachedThreadPool` depending on the workload."

### 24. Difference between `Callable` and `Runnable`?
"`Runnable` is the old interface—its `run()` method returns void and cannot throw checked exceptions. It’s strictly 'fire and forget'.

`Callable` is the newer sibling. Its `call()` method returns a result (via Generics) and can throw exceptions.

When you submit a `Callable` to an ExecutorService, you get back a `Future` object, which you can use to retrieve the result later."

### 25. What are atomic classes (`AtomicInteger`)?
"Atomic classes like `AtomicInteger` or `AtomicBoolean` provide a way to perform thread-safe operations on single variables without using `synchronized`.

They use low-level CPU instructions (CAS - Compare and Swap) to update the value.

For example, `incrementAndGet()` is atomic. It’s much faster than wrapping `count++` in a synchronized block because it avoids the overhead of context switching and locking. I use them religiously for counters and metrics."
