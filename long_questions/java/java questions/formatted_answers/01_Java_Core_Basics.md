# 01. Java Core Basics

**Q: Difference between JDK, JRE, JVM**
> "Think of the **JDK**, or Java Development Kit, as your complete toolbox. It’s what we developers use because it has everything needed to *write* and *compile* code, including the compiler (`javac`) and debuggers.
>
> Inside that toolbox is the **JRE**, the Java Runtime Environment. This is for the *users* who just want to run the app. It contains the libraries and the JVM, but not the development tools.
>
> Finally, the **JVM**, or Java Virtual Machine, is the engine inside the JRE. It’s the brain that actually executes the bytecode line-by-line.
>
> So simply put: JDK is for developing, JRE is for running, and JVM is the engine that does the work."

**Indepth:**
> The JVM acts as the abstraction layer between the compiled bytecode and the underlying hardware. One critical component inside the JVM is the **JIT (Just-In-Time) Compiler**. Instead of interpreting bytecode line-by-line every time (which is slow), the JIT compiler identifies "hot spots" (frequently executed code) and compiles them into native machine code for performance optimization.
>
> Additionally, the JDK includes the **javac** compiler which produces bytecode (.class files). This bytecode is platform-agnostic, fulfilling Java's "Write Once, Run Anywhere" promise. The JVM then translates this bytecode to the specific machine instructions of the host OS.


---

**Q: Abstract class vs Interface**
> "This is a classic one. Think of an **Abstract Class** as a template where you can provide some default behavior but still force subclasses to implement specific methods. It’s roughly saying, 'These classes are a *type of* this thing.' You can have state (variables) and constructors here.
> 
> An **Interface**, on the other hand, is like a contract or a capability. It says, 'I can *do* this behavior.' Since Java 8, interfaces can have default methods, which blurs the line a bit, but the key rule remains: a class can implement *multiple* interfaces but extend only *one* abstract class.
>
> So, use an abstract class for shared state and identity, and use an interface for defining capabilities or roles."

**Indepth:**
> **Java 8+ Evolution**: Prior to Java 8, interfaces could only have abstract methods. Java 8 introduced **default** and **static** methods, allowing interfaces to have implementation logic without breaking existing implementations. Java 9 added **private** methods for code reusability within the interface itself. Despite these changes, interfaces still cannot hold state (instance variables), which remains the key differentiator from abstract classes.
>
> **Design Intent**: Use an Abstract Class when you want to share code among several closely related classes (an "is-a" relationship) and define a common base with state. Use an Interface to define a contract for classes that may not be related (a "can-do" capability) or when you need to simulate multiple inheritance.


---

**Q: String vs StringBuilder vs StringBuffer**
> "**String** is immutable. That means every time you modify it—like concatenating "Hello" and "World"—you aren't changing the original string; you're actually creating a brand new object in memory. That’s safe but slow if you do it a lot.
>
> **StringBuilder** solves this. It’s mutable, so you can append and change characters without creating new objects. It’s much faster for heavy string manipulation.
>
> **StringBuffer** is the older brother of StringBuilder. It does the same thing but is *synchronized*, meaning it’s thread-safe. However, that synchronization adds overhead.
>
> In short: Use **String** for constants. Use **StringBuilder** for almost all string manipulation. Use **StringBuffer** only if you strictly need thread safety in a legacy context."

**Indepth:**
> **String Pool**: Because Strings are immutable, Java optimizes memory using the **String Constant Pool** (a special area in the Heap). If you create a String literal (e.g., `String s = "hello"`), the JVM checks the pool first. If "hello" exists, it returns a reference to the same instance. Using `new String("hello")` bypasses this check and forces a new object creation, which is inefficient.
>
> **Performance**: `StringBuilder` and `StringBuffer` use an internal `char[]` array that resizes dynamically (usually doubling capacity). This resizing involves copying the old array to a new one, so initializing them with an estimated capacity can significantly improve performance in tight loops.


---

**Q: equals() vs ==**
> "The `==` operator checks for **reference equality**. It asks: 'Are these two pointing to the exact *same object* in memory?'
>
> The `.equals()` method checks for **value equality** (or logical equality). It asks: 'Do these two objects contain the *same data*?'
>
> For example, if you create two new String objects with the text 'hello', `==` will return `false` because they are different memory addresses, but `.equals()` will return `true` because the content is identical. Always use `.equals()` for comparing object values."

**Indepth:**
> **Contract**: When you override `equals()`, you **must** also override `hashCode()`. The contract states that if two objects are equal, they must have the same hash code. Violating this breaks hash-based collections like `HashMap` and `HashSet`, where objects might get "lost" because they fall into the, wrong bucket.
>
> **Interning**: A rare case where `==` returns true for different String variables is when they are **interned**. `s1.intern() == s2.intern()` will be true if `s1.equals(s2)`, because `intern()` forces the String into the pool and returns the canonical representation.


---

**Q: final vs finally vs finalize**
> "They sound similar but do very different things.
>
> **final** is a keyword to restrict modification. If a variable is final, it’s a constant. If a method is final, it can’t be overridden. If a class is final, it can’t be inherited.
>
> **finally** is a block used in exception handling. It guarantees that code runs—like closing a file or database connection—regardless of whether an exception occurred or not.
>
> **finalize** is a method on the `Object` class used for garbage collection cleanup. But be careful—it’s deprecated and unpredictable. You should use `try-with-resources` or `Cleaner` instead."

**Indepth:**
> **Memory Model (final)**: In the Java Memory Model (JMM), `final` fields have special initialization safety guarantees. Once a constructor finishes, any thread that sees the object is guaranteed to see the correctly initialized values of its `final` fields, without needing explicit synchronization.
>
> **Deprecation (finalize)**: `finalize()` is deprecated because it is unreliable (no guarantee when strings run), can resurrect objects, and adds significant GC overhead. Java 9 introduced `java.lang.ref.Cleaner` as a safer alternative, though `try-with-resources` is the preferred standard.


---

**Q: Collections vs Arrays**
> "**Arrays** are the most basic container. They are fixed in size—once you create an array of size 10, it stays size 10. They are fast and can hold primitives, but not very flexible.
>
> **Collections** (like ArrayList, HashSet) are dynamic. They can grow and shrink as needed. They also come with powerful utility methods for sorting, searching, and shuffling. However, they can only hold Objects, not primitives (though autoboxing handles that for us).
>
> So, use Arrays for performance or fixed data, and Collections for flexibility."

**Indepth:**
> **Generic Type Erasure**: Collections use Generics which provide compile-time type safety. However, at runtime, due to **Type Erasure**, the JVM sees only the raw type. This is why you can't check `if (list instanceof List<String>)` or create generic arrays like `new List<String>[10]`.
>
> **Covariance**: Arrays are covariant (`Integer[]` is a subtype of `Number[]`), which can lead to runtime `ArrayStoreException`. Collections are invariant (`List<Integer>` is NOT a subtype of `List<Number>`), preventing these errors at compile time.


---

**Q: List vs Set vs Map**
> "Think of a **List** like a shopping list. Order matters, and duplicates are fine (you can have 'Milk' twice).
>
> A **Set** is like a bag of unique items. Order usually doesn't matter (unless it's a LinkedHashSet), but duplicates are strictly forbidden. If you try to add 'Milk' twice, the second one is ignored.
>
> A **Map** is a dictionary. It holds Key-Value pairs. You look up a value using its key, like looking up a definition using a word. Keys must be unique, but values can be duplicated."

**Indepth:**
> **Complexity**: `ArrayList` offers O(1) access but O(n) search/remove. `HashSet` offers O(1) add/contains. `HashMap` offers O(1) get/put. `TreeMap` ensures O(log n) for operations due to sorting.
>
> **Choosing**: Use `Set` for uniqueness and membership checks. Use `List` for ordered collections with duplicates. Use `Map` for key-value associations.


---

**Q: ArrayList vs LinkedList**
> "**ArrayList** is backed by a dynamic array. It’s super fast for retrieving elements (`get(i)`) because it uses index-based access. But if you add or remove elements in the middle, it’s slow because it has to shift all the other elements around.
>
> **LinkedList** is a doubly-linked list. Adding or removing elements is fast—you just change the pointers. But retrieving an element is slow because you have to traverse the list from the start node to find it.
>
> Stick with **ArrayList** 99% of the time, unless you are frequently inserting/deleting from the middle of a large list."

**Indepth:**
> **Cache Locality**: `ArrayList` is significantly faster in practice because its backing array is stored in contiguous memory. This allows the CPU to leverage **spatial locality** and prefetch data into the L1/L2 cache.
>
> **Memory Overhead**: `LinkedList` consumes more memory because every element is wrapped in a `Node` object storing references to previous and next elements (overhead of 2 extra references per item). This pointer chasing causes frequent cache misses.


---

**Q: HashMap vs TreeMap vs LinkedHashMap**
> "**HashMap** is your go-to. It makes no guarantees about order—keys might be scrambled—but it’s incredibly fast (O(1)) for lookups and insertions.
>
> **LinkedHashMap** preserves insertion order. If you iterate through it, keys come out in the exact order you put them in. It’s slightly slower than HashMap but useful for things like caches.
>
> **TreeMap** sorts the keys according to their natural order or a Comparator. So if your keys are names, they will be stored alphabetically. It’s slower (O(log n)) but essential if you need sorted data."

**Indepth:**
> **Internal Working (HashMap)**: Keys are hashed to find a "bucket" index. If multiple keys land in the same bucket (collision), they are stored as a linked list. Since Java 8, if a bucket gets too populated (threshold of 8), the list transforms into a **Red-Black Tree**, improving worst-case performance from O(n) to O(log n).
>
> **TreeMap**: Always uses a Red-Black tree, keeping keys sorted. This enables operations like `subMap()` or `tailMap()` which `HashMap` cannot provide.


---

**Q: Fail-fast vs Fail-safe**
> "**Fail-fast** iterators throw a `ConcurrentModificationException` immediately if you try to modify the collection (add/remove) while iterating over it. Common examples are `ArrayList` and `HashMap` iterators. They yell at you right away to prevent bugs.
>
> **Fail-safe** iterators work on a *copy* of the collection or use specific concurrency mechanisms. They won't throw an exception if the collection changes during iteration, but they might not reflect the latest data. Examples are `ConcurrentHashMap` and `CopyOnWriteArrayList`."

**Indepth:**
> **Mechanism**: Fail-fast iterators maintain a `modCount`. Every time the collection is structurally modified, `modCount` increments. The iterator checks if `expectedModCount == modCount` on every `next()` call. If not, it throws `ConcurrentModificationException`.
>
> **Weakly Consistent**: "Fail-safe" usually refers to **weakly consistent** iterators (like in `ConcurrentHashMap`). They guarantee traversal without throwing exceptions but may show a state of the collection strictly at the time of iterator creation or merely some consistent state, potentially missing later updates.


---

**Q: map() vs flatMap()**
> "Both are used in Streams to transform data.
>
> `map()` applies a function/operation to each element and returns *one* result for each input. One goes in, one comes out. Like mapping a list of numbers to their squares.
>
> `flatMap()` is used when each element produces a *stream* of values, and you want to flatten all those streams into a single list. Think of it as 'map then flatten'. If you have a list of sentences and want a single list of *all words* in those sentences, you’d use `flatMap()`."

**Indepth:**
> **Dimensionality Reduction**: `flatMap` is essentially `map` + `flatten`. Use `map` for 1-to-1 transformations (`Stream<Object>` -> `Stream<OtherObject>`). Use `flatMap` when your function returns a Stream/List itself (`Stream<List<Object>>`), and you want a single smooth `Stream<Object>` as the result.
>
> **Optional**: `flatMap` is crucial in `Optional` to avoid nested structures like `Optional<Optional<String>>`.


---

**Q: Thread vs Runnable**
> "**Runnable** is an interface that represents a task—something to be done. It just has a `run()` method.
>
> **Thread** is a class that manages the actual execution of that task on the CPU.
>
> You typically implement `Runnable` (or usage a Lambda) to define *what* to do, and then pass it to a `Thread` (or Executor) to define *how* to run it. Implementing `Runnable` is preferred because you can still extend another class, whereas extending `Thread` limits your inheritance options."

**Indepth:**
> **Decoupling**: Using `Runnable` separates the *task* from the *thread*. This is vital for `ExecutorService` (Thread Pools). You can pass the same `Runnable` to a Cached Pool, Fixed Pool, or virtual thread.
>
> **Resources**: `Thread` creation is expensive (stack memory, OS thread). Using `Runnable` allows reusing threads via pools, preventing memory exhaustion.


---

**Q: @Component vs @Service vs @Repository**
> "Technically, they are all the same! `@Service` and `@Repository` are just specializations of `@Component`. They all mark a class as a Spring Bean.
>
> But we distinguish them for semantic clarity:
> *   **@Component** is a generic bean.
> *   **@Service** marks the business logic layer.
> *   **@Repository** marks the data access layer (DAO) and enables automatic exception translation for database errors.
>
> Using the specific annotations helps you (and Spring) understand the architecture of your app."

**Indepth:**
> **Aspect-Oriented Programming (AOP)**: These annotations act as pointcuts. For `@Repository`, Spring automatically registers a `PersistenceExceptionTranslationPostProcessor` which intercepts platform-specific database exceptions (SQL/Hibernate) and re-throws them as Spring’s unified, unchecked `DataAccessException` hierarchy.
>
> **Scanning**: While functionally similar, using the correct annotation acts as self-documentation and allows for layer-specific processing in the future.

