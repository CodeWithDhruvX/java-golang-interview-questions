# Basic Level Java Interview Questions

## From 01 Java Core Basics
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


## From 04 Java Fundamentals Core
# 04. Java Fundamentals (Core)

**Q: Explain static keyword in Java**
> "**static** means 'belongs to the class, not the instance'.
>
> If a variable is static, there is only *one copy* of it in memory, shared by all objects of that class.
> If a method is static, you can call it without creating an object (like `Math.abs()`).
>
> It’s great for constants and utility functions, but be careful—overusing static can make your code hard to test and thread-unsafe."

**Indepth:**
> **Static Initializers**: You can use a `static { }` block to complex static variable initialization (like loading a config file or handling exceptions during init). This block runs once when the class is loaded by the ClassLoader.
>
> **Method Hiding**: Static methods cannot be *overridden* (because they don't dispatch on the instance type at runtime), they can only be *hidden*. If a subclass defines the same static method, the version called depends on the reference type, not the object type.


---

**Q: What does volatile do?**
> "**volatile** is all about visibility between threads.
>
> In a multi-threaded app, CPU cores might cache a variable's value to run faster. **volatile** tells the JVM: 'Never cache this. Always read and write this directly to main memory.'
>
> This guarantees that if Thread A changes a flag, Thread B sees it immediately. However, it does *not* make compound operations (like `i++`) atomic. For that, you need `AtomicInteger` or `synchronized`."

**Indepth:**
> **Hardware Level**: `volatile` works by emitting a "memory barrier" (fence) instruction to the CPU. This prevents the CPU from reordering instructions across the barrier and forces a cache flush/reload.
>
> **Limitations**: It does NOT guarantee atomicity. `volatile int count = 0; count++` is unsafe because `count++` is three steps (read, add, write). If two threads do this, they might overwrite each other's work.


---

**Q: Comparing Objects: == vs equals()**
> "We touched on this, but it's worth repeating.
>
> `==` checks if two references point to the *same memory address*. It’s like asking, 'Are these two keys for the exact same house?'
>
> `.equals()` checks if two objects have the *same content*. It’s like asking, 'Do these two houses look identical?'
>
> Always use `.equals()` for objects, especially Strings."

**Indepth:**
> **String Interning**: String literals are interned (stored in a pool), so `==` works for them often unexpectedly. `String s1 = "a"; String s2 = "a";` makes `s1 == s2` true. But `new String("a")` creates a heap object, making `==` false. This inconsistency is why `equals()` is mandatory.


---

**Q: Common Object methods (toString, hashCode, equals)**
> "Every class in Java inherits from `Object`, which gives you three key methods:
> 1.  `toString()`: Returns a string representation of the object. You override this to make logging and debugging easier (e.g., printing `User{id=1, name='Bob'}` instead of `User@1a2b3c`).
> 2.  `equals()`: See above. You override this to define what makes two instances 'equal'.
> 3.  `hashCode()`: Returns an integer hash. This is crucial for HashMaps. The rule is: if `equals()` says two objects are equal, their `hashCode()` *must* be the same."

**Indepth:**
> **Hash Collision**: When two different objects produce the same `hashCode`, they land in the same bucket in a HashMap. This is a collision. If many objects collide, performance degrades from O(1) to O(n) (or O(log n) in Java 8+). A good `hashCode()` algorithm spreads keys uniformly across buckets to minimize this.


---

**Q: finalize() method - Why is it deprecated?**
> "It was meant to be a cleanup method called before Garbage Collection. But it’s a disaster in practice.
>
> 1.  You don't know *when* (or if) it will run.
> 2.  It can resurrect objects (stop them from being collected).
> 3.  If it throws an exception, the GC ignores it.
>
> Just don't use it. Use `try-with-resources` or `Cleaner` instead."

**Indepth:**
> **Zombie Objects**: If you override `finalize()` and assign `this` to a static variable, you "resurrect" the object, preventing GC. The GC tracks this and won't call finalize again, leaving the object in a weird state.
>
> **Performance**: Objects with finalizers take at least two GC cycles to be reclaimed (one to mark, one to run finalize and reclaim), delaying memory release.


---

**Q: Wrapper Classes & Autoboxing**
> "Java has primitives (int, boolean) for speed, and Objects (Integer, Boolean) for flexibility (like putting them in Lists).
>
> **Wrapper Classes** allow primitives to be treated as Objects.
> **Autoboxing** is the automatic conversion the compiler does for you. When you say `List<Integer> list = new ArrayList<>(); list.add(5);`, Java automatically converts that primitive `5` into an `Integer` object.
> **Unboxing** is the reverse."

**Indepth:**
> **Performance Overhead**: Autoboxing creates objects. In a tight loop (like adding 1 million ints to a `List<Integer>`), this creates 1 million Integer objects, causing massive GC pressure. For high performance, use primitive arrays or specialized libraries like *Eclipse Collections* (`IntList`) to avoid boxing.
>
> **Nulls**: Unboxing a `null` Integer throws a `NullPointerException`. Always check for null before unboxing if the wrapper might be null.


---

**Q: Integer Cache (-128 to 127)**
> "This is a cool optimization. Java caches `Integer` objects for values between -128 and 127.
>
> So if you say `Integer a = 100; Integer b = 100;`, `a == b` is actually **true** because they point to the exact same cached object.
> But if you say `Integer a = 200; Integer b = 200;`, `a == b` is **false** because they are outside the cache range and are new objects.
>
> This is a classic 'gotcha' interview question!"

**Indepth:**
> **Configurable**: This cache range is technically configurable via JVM flags (`-XX:AutoBoxCacheMax=<size>`).
>
> **Flyweight Pattern**: This is a classic implementation of the **Flyweight Design Pattern**, where common object instances are shared to save memory.


---

**Q: BigInteger and BigDecimal**
> "**BigInteger** is for integers that are too large to fit in a `long` (64-bit). It can theoretically handle numbers as large as your RAM allows.
>
> **BigDecimal** is for precise decimal arithmetic. Floating point numbers (`float`, `double`) introduce rounding errors (0.1 + 0.2 is not exactly 0.3). BigDecimal is exact, so it's mandatory for financial calculations involving money."

**Indepth:**
> **Immutability**: Both classes are immutable. Operations like `a.add(b)` return a *new* BigInteger; they do not modify `a`.
>
> **Performance**: They are significantly slower than primitives because they use `int[]` arrays internally to simulate large numbers, and math operations are done in software, not directly by the CPU ALU.


---

**Q: What is Type Erasure?**
> "Generics in Java were added in Java 5 to provide compile-time type safety. But to stay compatible with older Java versions, they implemented it using **Type Erasure**.
>
> This means that generic type information (like `<String>` in `List<String>`) is removed (erased) by the compiler. At runtime, the JVM just sees a raw `List`.
>
> This is why you can't check `if (list instanceof List<String>)`—because at runtime, that `<String>` info is gone."

**Indepth:**
> **Bridge Methods**: To make polymorphism work with erasure, the compiler sometimes generates synthetic "bridge methods".
>
> **Heap Pollution**: Erasure can lead to "Heap Pollution" if you mix raw types and generics, where a variable of a parameterized type refers to an object that isn't of that type. This usually triggers a warning.


---

**Q: Wildcards in Generics (?, extends, super)**
> "These give you flexibility.
>
> 1.  `<?>` (Unbounded): 'I accept a list of *anything*.' Read-only basically.
> 2.  `<? extends Number>` (Upper Bound): 'I accept a list of Number or its subclasses (Integer, Double).' You can read Numbers from it, but you can't add to it (because you don't know if it's really a list of Integers or Doubles).
> 3.  `<? super Integer>` (Lower Bound): 'I accept a list of Integer or its parents (Number, Object).' You can add Integers to it."

**Indepth:**
> **PECS Mnemonic**: "Producer Extends, Consumer Super".
> *   If you want to **Produce** items from the list (read them), use `? extends`.
> *   If you want to **Consume** items into the list (write them), use `? super`.


---

**Q: Generic Methods**
> "You can make a single method generic without making the whole class generic. You declare the type parameter `<T>` before the return type.
>
> Example: `public <T> void printArray(T[] array)`.
>
> This lets you write one method that can print an array of Integers, Strings, or Dogs, ensuring type safety for that specific call."

**Indepth:**
> **Type Inference**: You usually don't need to specify the type witness (e.g., `Utils.<String>printArray(arr)`). The compiler infers it from the arguments.
>
> **Constructors**: Constructors can also be generic, allowing class instantiation to infer types based on arguments (the "diamond operator" `<>` does this in modern Java).


---

**Q: What is Reflection? Pros/Cons**
> "**Reflection** allows code to inspect and modify itself at runtime. You can list methods of a class, call private methods, or change private fields dynamically. Frameworks like Spring and Hibernate live on this.
>
> Pros: Extremely powerful and flexible. Enables frameworks to work.
> Cons: Slower performance. Breaks encapsulation (accessing private stuff). Can lead to security issues and fragile code."

**Indepth:**
> **JIT De-optimization**: Extensive use of reflection obscures the control flow from the JIT compiler, preventing optimizations like inlining.
>
> **Usage**: It matches strings (names of methods/classes) to code. It's heavily used in dependency injection containers (Spring), serialization libraries (Jackson), and testing frameworks (JUnit).


---

**Q: How to access a Private Field using Reflection?**
> "You get the `Field` object from the class, call `setAccessible(true)` on it, and then call `get(instance)`.
>
> `setAccessible(true)` suppresses Java language access checking. It’s like a master key. But modern Java (modules) makes this harder to do."

**Indepth:**
> **Modules (Java 9+)**: The Java Module System (Jigsaw) restricts strict encapsulation. Even with Reflection, you cannot access private members of a module unless that module explicitly `opens` the package to your code (or to everyone). This forces better architectural boundaries.


---

**Q: What is the Class class?**
> "It’s the entry point for all reflection. Every type in Java is associated with a `java.lang.Class` object. It holds the metadata: the name of the class, its methods, fields, constructors, etc.
>
> You get it via `MyClass.class` or `obj.getClass()`."

**Indepth:**
> **Class Loaders**: The `Class` object is loaded by a `ClassLoader`. If you have the same class file loaded by two different ClassLoaders (e.g., in a container like Tomcat), they are treated as *different* classes. `A.class != A.class` can happen!
>
> **Literals**: `int.class` exists even for primitives, though they don't have fields. It's a placeholder for the type.


---

**Q: Custom Annotations & Meta-Annotations**
> "An **Annotation** (`@MyAnnotation`) gives metadata about code.
>
> To create one, you use `@interface`.
>
> **Meta-Annotations** are annotations *for* your annotation:
> *   `@Target`: Where can I use this? (Field, Method, Class?)
> *   `@Retention`: How long does it last? (Source only? Compile time? Runtime?)
> *   `@Documented`: Should it show up in Javadoc?"

**Indepth:**
> **Processing**: Annotations do nothing by themselves. They need a processor.
> 1.  **Compile-time**: Annotation Processors (like Lombok) scan source code and generate new code.
> 2.  **Runtime**: Frameworks (Spring) use Reflection to inspect annotations on loaded classes and change behavior (e.g., wrap a method in a transaction if `@Transactional` is found).


---

**Q: Breaking Singleton using Reflection**
> "Even with a private constructor, you can use Reflection to call `setAccessible(true)` on that constructor and create a new instance!
>
> To prevent this, you can throw an exception inside your private constructor if `instance` is not null. Or, even better, use an **Enum** to implement your Singleton—Enums are safe from reflection attacks by design."

**Indepth:**
> **Defense**: If using a class-based Singleton, adding a check in the constructor `if (instance != null) throw new RuntimeException();` protects against Reflection but *not* against serialization or cloning attacks (unless you handle those too).
>
> **Enums**: Enums are compiled as classes extending `java.lang.Enum`, which strictly prevents instantiation via reflection (`Constructor.newInstance()` throws exception for Enums).


---

**Q: Private vs Default vs Protected vs Public**
> "These are the access modifiers:
>
> 1.  **private**: Visible only within the *same class*.
> 2.  **default** (no keyword): Visible within the *same package*.
> 3.  **protected**: Visible within the *same package* AND subclasses in *other packages*.
> 4.  **public**: Visible everywhere.
>
> The Golden Rule: Start with private and open up only what is necessary."

**Indepth:**
> **Package-Private**: The default access (no keyword) is often underused. It’s excellent for creating a "component" where multiple classes work together in the same package but only the main "Public API" class is exposed to the rest of the world.


## From 06 OOP Basics
# 06. OOP Basics

**Q: What are the 4 Pillars of OOP? Explain with real-world examples.**
> "The four pillars are essentially the rules for good object-oriented design:
>
> 1.  **Encapsulation**: 'Keep your secrets.' Bundling data and methods together and hiding the internal details.
>     *   *Real world:* A capsule. You know it's medicine, but you don't know the chemical formula inside.
>
> 2.  **Inheritance**: 'Parent and Child.' Creating new classes based on existing ones to reuse code.
>     *   *Real world:* A Child inherits traits (eye color) from their Parent.
>
> 3.  **Polymorphism**: 'Many forms.' One interface, multiple implementations.
>     *   *Real world:* A 'Person' can behave like a 'Student' in school, an 'Employee' at work, and a 'Customer' at a shop. Same person, different roles.
>
> 4.  **Abstraction**: 'Hiding complexity.' Showing only the essential features of an object.
>     *   *Real world:* Using a TV remote. You press 'Power' and it works. You don't need to know the circuit logic inside."

**Indepth:**
> **Cohesion vs Coupling**: Encapsulation increases Cohesion (classes do one thing well) and reduces Coupling (change in one class doesn't break another).
>
> **Liskov Substitution Principle**: Inheritance should follow LSP. A subclass must be substitutable for its superclass without breaking the application. If you override a method and throw a new checked exception or change its behavior drastically, you violate this principle.


---

**Q: Difference between Abstract Class and Interface?**
> "We covered this earlier, but to reinforce it:
>
> *   **Abstract Class**: Defines *identity* ('is-a'). Use it when classes share a common core but need specific implementations (e.g., `Dog` is an `Animal`). Can have state (fields) and constructors.
> *   **Interface**: Defines *capability* ('can-do'). Use it to define a contract that unrelated classes can implement (e.g., `Dog` and `Car` both implement `Moveable`). No state (until recently with static constants), multi-inheritance support."

**Indepth:**
> **Functional Interfaces**: An interface with exactly one abstract method is a Functional Interface and can be implemented using Lambda expressions (`() -> ...`). Abstract classes cannot be instantiated with Lambdas.
>
> **Constructors**: Abstract classes have constructors (called via `super()`), interfaces don't.


---

**Q: What is Polymorphism? (Compile-time vs Runtime)**
> "Polymorphism means 'many forms'.
>
> **Compile-time Polymorphism** (Static Binding) is **Method Overloading**.
> *   Same method name, different parameters. The compiler decides *which* method to call based on the arguments you pass.
>
> **Runtime Polymorphism** (Dynamic Binding) is **Method Overriding**.
> *   Same method signature in Parent and Child. The JVM decides *at runtime* which version to call based on the actual object type, not the reference type (`Animal a = new Dog(); a.sound()` calls Dog's sound)."

**Indepth:**
> **Dynamic Dispatch**: At runtime, the JVM uses the `vtable` (virtual method table) to look up the correct method to call. Overloaded methods are linked at compile time (Static Dispatch), so they are faster (no lookup).
>
> **Upcasting**: Polymorphism relies on Upcasting (`Parent p = new Child()`). You can access only the methods defined in `Parent`, but if `Child` overrides them, the `Child`'s version executes.


---

**Q: Can you override static or private methods?**
> "**No.**
>
> *   **static methods** belong to the *class*, not the instance. If you define a static method with the same signature in a subclass, it's called **Method Hiding**, not overriding. The parent's version remains if you access it via the parent class reference.
> *   **private methods** are invisible to the subclass. You can't override what you can't see. If you define a method with the same name in the subclass, it's just a completely new, unrelated method."

**Indepth:**
> **Virtual Methods**: In Java, all non-static, non-private, non-final methods are "virtual" by default, meaning they can be overridden.
>
> **Static Binding**: Static methods are bound at compile time based on the reference type. If you have `Animal a = new Dog(); a.staticMethod()`, it calls Animal's static method, not Dog's.


---

**Q: What is covariant return type?**
> "It sounds fancy, but it just means: When you override a method, the return type doesn't have to be *exactly* the same—it can be a **subclass** of the original return type.
>
> Example:
> Parent has `Animal produce()`.
> Child can override it with `Dog produce()`.
>
> This is allowed because a Dog *is* an Animal. It helps avoid type casting in client code."

**Indepth:**
> **Bridge Methods**: To support covariant return types (introduced in Java 5), the compiler generates a synthetic "bridge method" in the subclass with the original return type, which internally calls the new method. This maintains binary compatibility with older bytecode.


---

**Q: Composition vs Inheritance. Which is better?**
> "**Composition** is generally preferred over Inheritance ('Favor Composition over Inheritance').
>
> **Inheritance** (`extends`) creates a tight coupling. If the Parent class changes, the Child might break. It's an 'Is-A' relationship.
>
> **Composition** (having a private field of another class) is looser. It allows you to change behavior at runtime (by swapping the object) and doesn't expose the internal details. It's a 'Has-A' relationship."

**Indepth:**
> **Dependency Injection**: Composition is the basis of Dependency Injection (DI). Instead of hardcoding a dependency (`this.engine = new GasEngine()`), you pass it in (`this.engine = someEngine`). This makes testing easier because you can swap in a mock engine.
>
> **Fragile Base Class**: Heavy inheritance leads to the Fragile Base Class problem, where a change in a superclass causes unexpected bugs in subclasses. Composition avoids this.


---

**Q: What is the super keyword?**
> "**super** refers to the immediate parent class object.
>
> You use it to:
> 1.  Call the parent's constructor: `super(name)`.
> 2.  Call a parent's method that you have overridden: `super.printInfo()`.
> 3.  Access a parent's variable (rarely needed if encapsulated properly)."

**Indepth:**
> **Constructor Rule**: If you write a constructor in a child class, the very first line *must* be a call to `super()` or `this()`. If you don't write it, the compiler inserts `super()` (unexpectedly calling the no-arg parent constructor).


---

**Q: Significance of this keyword?**
> "**this** refers to the *current* object instance.
>
> You use it to:
> 1.  Distinguish instance variables from local variables (e.g., `this.name = name`).
> 2.  Call another constructor in the same class: `this(name, 0)`.
> 3.  Pass the current object as a parameter to another method."

**Indepth:**
> **Builder Pattern**: `this` is crucial for method chaining (Fluent Interface) where methods return `this`.
>
> **Inner Classes**: To access the outer class instance from an inner class, you use `OuterClassName.this`.


---

**Q: Can an Interface extend another Interface?**
> "**Yes.** An interface can extend multiple other interfaces.
>
> `public interface Robot extends Machine, Intelligent { ... }`
>
> This allows you to build complex contracts from smaller ones."

**Indepth:**
> **Interface Segregation Principle**: This feature supports ISP. Instead of creating one massive interface, you create small, specific interfaces. complex classes can then implement multiple of them, or a new interface can extend several of them to bundle capabilities.


---

**Q: Difference between Overloading and Overriding?**
> "**Overloading**: Same method name, *different* parameters (signature). Happens in the *same* class. Resolved at Compile-time.
>
> **Overriding**: Same method name, *same* parameters. Happens in *Parent-Child* classes. Resolved at Runtime."

**Indepth:**
> **Return Type**: You cannot overload a method *only* by changing the return type. The parameter list must change.
>
> **Exceptions**: Overridden methods cannot throw new or broader checked exceptions than the parent method (Liskov principle), but they can throw fewer or narrower exceptions. Overloaded methods have no such restrictions.


---

**Q: Can a constructor be private? Why?**
> "**Yes.**
>
> You make a constructor private to:
> 1.  Prevent instantiation from outside (Singleton Pattern).
> 2.  Prevent subclassing (you can't extend a class if you can't call its constructor).
> 3.  Force usage of static factory methods (`MyClass.create()`)."

**Indepth:**
> **Utility Classes**: `java.lang.Math` is a classic example. It has a private constructor because it only contains static methods. There is no point in creating an instance of `Math`.
>
> **Subclassing**: If a class has *only* private constructors, it cannot be subclassed (referenced in `extends`), effectively making it final.


---

**Q: Default Constructor vs No-Args Constructor?**
> "A **No-Args Constructor** is simply a constructor that takes no parameters. You can write one yourself.
>
> A **Default Constructor** is the no-args constructor that the **compiler** automatically inserts for you *only if* you haven't defined *any* other constructors.
>
> Once you write `public MyClass(int x)`, the default constructor vanishes. If you still want a no-args one, you must type it manually."

**Indepth:**
> **Serialization Issue**: If a Serializable class extends a non-Serializable parent, the parent *must* have a no-args constructor (visible to the subclass), otherwise serialization fails with `InvalidClassException`. The JVM needs it to initialize the parent's fields during deserialization.


---

**Q: What is Constructor Chaining?**
> "It's the process of calling one constructor from another.
>
> *   `this()` calls another constructor in the *same* class.
> *   `super()` calls a constructor in the *parent* class.
>
> This happens automatically (the compiler inserts `super()` if you don't), ensuring that the Object is fully initialized from the top of the hierarchy (Object) down to the specific subclass."

**Indepth:**
> **Recursive Execution**: Constructors execute top-down (Parent first, then Child).
>
> **Risk**: Avoid calling overridable methods inside a constructor! If Parent constructor calls `print()`, and Child overrides `print()`, the Child's method runs *before* the Child's constructor variables are initialized, leading to bugs (e.g., accessing null fields).


---

**Q: Use of instanceof operator?**
> "It checks if an object is an instance of a specific class (or one of its subclasses/interfaces).
>
> `if (animal instanceof Dog)` returns true if `animal` is actually a Dog.
>
> In modern Java (14+), you can use 'Pattern Matching for instanceof' to check and cast in one step:
> `if (animal instanceof Dog d) { d.bark(); }`."

**Indepth:**
> **Implementation**: `instanceof` usage is generally a code smell indicating poor polymorphism. Instead of `if (obj instanceof Dog) bark() else if (obj instanceof Cat) meow()`, you should just call `obj.makeSound()` and let polymorphism handle it.
>
> **Null**: `null instanceof Anything` always returns `false`.


---

**Q: What is an Initialization Block?**
> "It's a block of code inside a class (surrounded by `{}`). It runs every time an object is created, *before* the constructor.
>
> There is also a **Static Initialization Block** (`static {}`), which runs only *once* when the class is loaded by the ClassLoader."

**Indepth:**
> **Order of Execution**:
> 1.  Static Blocks (Class Load time)
> 2.  Parent Constructor
> 3.  Instance Initializer Blocks (in order of appearance)
> 4.  Constructor Body
>
> **Copying**: The compiler actually copies the code from instance initializer blocks into *every* constructor (right after the `super()` call).


---

**Q: Multiple Inheritance in Java?**
> "Java does **not** support multiple inheritance of *classes* (A extends B, C) to avoid the 'Diamond Problem' (ambiguity if both B and C have the same method).
>
> However, Java **does** support multiple inheritance of *interfaces* (A implements B, C). Since Java 8 (default methods), the Diamond Problem can occur with interfaces too, but the compiler forces you to override the conflicting method to resolve the ambiguity."

**Indepth:**
> **Diamond Problem Solution**: If interface B and C both define `default void run()`, and A implements both, the compiler forces A to override `run()`. Inside A, you can choose which one to call using `B.super.run()`.


---

**Q: What is a Marker Interface?**
> "It’s an empty interface with no methods (like `Serializable`, `Cloneable`, `Remote`).
>
> It serves as a 'tag' or metadata. The JVM checks for this tag to enable special behavior (like allowing serialization).
>
> In modern Java, **Annotations** are often used instead of Marker Interfaces."

**Indepth:**
> **Performance**: `instanceof` checks against marker interfaces are extremely fast.
>
> **Modern replacement**: Use Annotations implies using Reflection, which is slower. Marker interfaces are part of the type system, allowing compile-time checks (e.g., a method taking `Serializable` argument).


---

**Q: Can abstract class have constructor?**
> "**Yes.**
>
> Even though you can't create an instance of an abstract class directly (`new AbstractClass()` is illegal), the constructor is still needed to initialize the fields defined in the abstract class. It is called by the subclass constructor using `super()`."

**Indepth:**
> **Chain of Responsibility**: When you instantiate a `Child`, the runtime allocates memory for the *entire* object, including Parent fields. The `super()` call initializes those parent fields. Without it, the object would be partially uninitialized logic-wise.


---

**Q: Shallow Copy vs Deep Copy (Object Cloning)?**
> "Refers to the `clone()` method.
>
> The default `Object.clone()` performs a **Shallow Copy** (copies primitives, but copies *references* for objects).
>
> To get a **Deep Copy** (independent copy of everything), you must override `clone()` and manually clone the mutable fields, or use serialization/library tools."

**Indepth:**
> **Cloneable Interface**: To use `clone()`, your class *must* implement `Cloneable`, otherwise it throws `CloneNotSupportedException`. This is a weird design decision in Java (methods on Object, but requires interface).
>
> **Alternatives**: Copy Constructors (`public Car(Car other)`) are generally preferred over `clone()` because they are simpler to implement and don't require handling checked exceptions.


---

**Q: Immutable Class - How to create one?**
> "To make a class immutable (like String):
> 1.  Make the class `final` (so no subclasses can mess it up).
> 2.  Make all fields `private` and `final`.
> 3.  Provide only Getters, no Setters.
> 4.  If a field is a mutable object (like a Date or List), return a *copy* of it in the getter, not the original reference ('Defensive Copying')."

**Indepth:**
> **Thread Safety**: Immutable objects are inherently thread-safe. They require no synchronization.
>
> **Hash Keys**: Immutable objects make excellent HashMap keys because their `hashCode()` never changes. If a key object changes after being put in a Map, you won't be able to retrieve the value.


## From 14 SOLID Arrays Basics
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


## From 15 Strings Basics
# 15. String Manipulation Basics

**Q: Difference between String, StringBuilder, and StringBuffer**
> "This is the classic Java interview question.
>
> **String** is **Immutable**. Once you create `"Hello"`, it can never change. If you do `"Hello" + " World"`, you aren't changing the original string; you are creating a brand new object in memory. This is safe but slow if you do it inside a loop.
>
> **StringBuilder** is **Mutable**. You can modify it (`.append()`) without creating new objects. It's fast and efficient. It is **not** thread-safe, but that's usually what you want for local variables.
>
> **StringBuffer** is the old, legacy version of StringBuilder. It is also mutable, but it is **Synchronized** (thread-safe). This makes it slower. You almost never need it anymore unless you are sharing a string builder between threads (which is rare)."

**Indepth:**
>

---

**Q: String Immutability — Why is it important?**
> "Java made Strings immutable for several key reasons:
>
> 1.  **Security**: Strings are used for everything—database URLs, passwords, file paths. If I pass a filename string to a method, I need to be 100% sure that method can't modify my string and trick me into writing to the wrong file.
> 2.  **Caching (String Pool)**: Because they can't change, Java can safely store one copy of `"Hello"` and let 100 different variables point to it. This saves massive amounts of memory.
> 3.  **Thread Safety**: Immutable objects are automatically thread-safe. You can share a String across threads without any locking."

**Indepth:**
>

---

**Q: Reverse, Palindrome, Anagrams (Logic)**
> "These are the bread-and-butter of coding rounds.
>
> *   **Reverse**: You can use `StringBuilder.reverse()`, but interviewers usually want you to do it manually. Convert to `char[]`, then swap start/end pointers until they meet.
> *   **Palindrome**: A string that reads the same forwards and backwards. Just reuse your reverse logic: does `str.equals(reverse(str))`? Or check pointers: `charAt(0) == charAt(len-1)`, etc.
> *   **Anagrams**: Two strings with the same characters in different orders (e.g., 'listen' and 'silent'). The easiest way? Sort both strings and check if they are equal. The faster way? Use a frequency map (or int[26] array) to count character occurrences."

**Indepth:**
> **Thread Safety**: `StringBuffer` is synchronized (thread-safe) but slow. `StringBuilder` is not synchronized but fast. Since Java 5, `StringBuilder` is the default choice.


---

**Q: First non-repeating character**
> "To find the first unique character (like 'l' in 'google'), you need two passes.
>
> Pass 1: Loop through the string and build a frequency map (`Map<Character, Integer>`) counting how many times each char appears.
> Pass 2: Loop through the string *again* (not the map, because order matters). Check the map for each char. The first one with a count of 1 is your winner."

**Indepth:**
> **UTF-16**: Java Strings use UTF-16 encoding internally. Most characters take 2 bytes, but some rare Unicode characters (emojis) take 4 bytes (surrogate pairs). `length()` returns the number of 2-byte code units, not the number of actual characters!


---

**Q: equals() vs equalsIgnoreCase() vs ==**
> "Never use `==` for Strings!
> `==` checks **Reference Equality**—are these two variables pointing to the exact same object in heap memory? Even if both contain "hello", `==` might return false if one was created with `new String()`.
>
> Always use `.equals()` for **Content Equality**. It checks if the actual characters are the same.
>
> `.equalsIgnoreCase()` is just a convenience wrapper that ignores casing, so 'JAVA' equals 'java'."

**Indepth:**
>

---

**Q: substring() vs subSequence()**
> "Functionally, they do almost the same thing.
>
> `substring()` returns a **String**. It’s what you use 99% of the time.
> `subSequence()` works on the `CharSequence` interface (which `String`, `StringBuilder`, and `StringBuffer` all implement). You usually only see this when using generic APIs that accept any char sequence.
>
> Historically (pre-Java 7), `substring` caused memory leaks because it shared the underlying char array. Modern Java doesn't do that—it copies the data, so it's safe."

**Indepth:**
>

---

**Q: trim() vs strip() (Java 11)**
> "For years, we used `trim()` to remove whitespace. But `trim()` is old—it only removes ASCII characters (space, tab, newline). It doesn't understand newer Unicode whitespace standards.
>
> Java 11 introduced `strip()`. It is 'Unicode-aware'. It removes all kinds of weird whitespace characters that `trim()` misses.
>
> Use `strip()` for modern applications. It also comes with `stripLeading()` and `stripTrailing()` for more control."

**Indepth:**
>

---

**Q: String Constant Pool**
> "This is a special area in the Heap memory.
>
> When you type `String s = "Hello";` (a literal), Java checks the pool. If "Hello" exists, it returns a reference to the existing one. If not, it adds it.
>
> When you type `String s = new String("Hello");`, you force Java to create a **new** object on the heap, bypassing the pool checks (though the internal char array might still be shared). This is generally wasteful and discouraged."

**Indepth:**
>

---

**Q: intern() method**
> "This method manually puts a String into the String Pool.
>
> If you have a String object on the heap (maybe read from a file), calling `.intern()` checks the pool.
> If the pool already has that value, it returns the pool's reference.
> If not, it adds your string to the pool and returns it.
>
> It’s a way to deduplicate strings in memory manually."

**Indepth:**
> **Conversion**: `String.valueOf(10)` vs `Integer.toString(10)`. They do the same thing. `"" + 10` also works but generates extra StringBuilder garbage.


## From 31 Spring Boot Basics Revision
# 31. Spring Boot (Basics & Testing Revision)

**Q: Spring Boot vs Spring**
> "**Spring** is the Engine. It provides Dependency Injection, MVC, Transaction Management, etc. But it requires a lot of setup (XML or huge Java Config classes).
>
> **Spring Boot** is the Car. It wraps the Engine with 'Auto Configuration', an Embedded Server, and 'Starters'. It lets you just turn the key (Run the main method) and drive, without assembling the parts yourself."

**Indepth:**
> **Convention over Configuration**: This is the core philosophy. Spring Boot assumes reasonable defaults (convention) so you don't have to configure them, unless you *want* to different (configuration).


---

**Q: @SpringBootApplication**
> "It's a convenience annotation that combines three others:
> 1.  `@Configuration`: Defines this as a source of bean definitions.
> 2.  `@EnableAutoConfiguration`: Tells Boot to start adding beans based on classpath settings.
> 3.  `@ComponentScan`: Tells Boot to look for other components (`@Service`, `@Controller`) in the current package and sub-packages."

**Indepth:**
> **Entry Point**: This annotation acts as the blueprint for the application. It triggers the scanning process that finds all your Beans, Controllers, and Services to wire them together in the ApplicationContext.


---

**Q: Auto Configuration**
> "It's Spring Boot's 'Opinionated' logic.
> At startup, Boot checks your Classpath.
> *   Do you have `h2.jar`? -> I'll create an in-memory DataSource.
> *   Do you have `spring-webmvc.jar`? -> I'll configure a DispatcherServlet and Tomcat.
>
> You can override any of these 'opinions' by defining your own bean."

**Indepth:**
> **Backing Off**: Conditional annotations like `@ConditionalOnMissingBean` allow you to "back off" the default. "If the user defined their own `DataSource`, don't create my default embedded one."


---

**Q: Spring Boot Starter**
> "A Starter is a 'Bill of Materials'. It's a single dependency in your `pom.xml` that brings in all the necessary jars for a feature.
>
> Instead of adding `spring-web`, `jackson`, `tomcat-embed`, and `validation-api` separately, you just add **`spring-boot-starter-web`**, and it pulls them all in with compatible versions."

**Indepth:**
> **Version Management**: Starters also act as a parent pom, managing transitive dependencies. You don't need to specify version numbers for individual libraries (like logging or jackson); the Starter ensures they are compatible.


---

**Q: Embedded Tomcat**
> "In the old days, you installed a separate Tomcat server, built a `.war` file, and deployed it.
>
> In Spring Boot, Tomcat is just a library (a JAR) *inside* your application.
> When you run your App, it starts Tomcat programmatically. This means your app is self-contained and portable."

**Indepth:**
> **Switching**: Boot also supports Jetty and Undertow. You can exclude `spring-boot-starter-tomcat` and add `spring-boot-starter-jetty` if you prefer a different servlet container engine.


---

**Q: JAR vs WAR**
> "**JAR (Java Archive)**: The default for Boot. It includes the embedded server. You run it with `java -jar app.jar`. Best for Microservices and Containers.
>
> "**WAR (Web Archive)**: Legacy format. Used if you *must* deploy to an external, existing Tomcat/Wildfly server. Boot supports it, but it's less common now."

**Indepth:**
> **Cloud Native**: JAR is the standard for Docker and Kubernetes. It aligns with the "12 Factor App" methodology where configuration and runtime are bundled together.


---

**Q: TestContainers**
> "Stop mocking your database in integration tests. And stop using H2 if you use Postgres in production.
>
> **TestContainers** spins up a *real* Docker container (e.g., a real Postgres DB) for your test, runs the test against it, and shuts it down.
> It ensures your code works with the actual database technology you use in Prod."

**Indepth:**
> **Transient**: The containers are ephemeral. They start fresh for the test suite and are destroyed afterwards. No more "dirty database" issues causing flaky tests.


---

**Q: @MockBean vs @SpyBean**
> "**@MockBean**: Replaces the real bean with a hollow shell (Mockito mock).
> *   `userService.getUser()` returns `null` unless you define `when(...).thenReturn(...)`.
> *   Use this to isolate the component you are testing.
>
> "**@SpyBean**: Wraps the *real* bean.
> *   Methods run normally, but you can verify them (`verify(bean).called()`) or stub specific methods.
> *   Use this when you want integration testing but need to spy on internal calls."

**Indepth:**
> **Mockito**: Both of these integrate natively with Mockito. `@MockBean` is essentially `Mockito.mock()`, and `@SpyBean` is `Mockito.spy()`, but they are automatically injected into the Spring ApplicationContext.


---

**Q: @WebMvcTest vs @SpringBootTest**
> "**@SpringBootTest**: Starts the **whole application context**. Connecting to DB, loading all services. Slow. Good for full Integration Tests.
>
> "**@WebMvcTest**: Slices the context. Only loads the Controller layer. It does **not** load `@Service` or `@Repository` beans.
> Fast. Use it for testing Unit Testing controllers (URL mapping, JSON serialization)."

**Indepth:**
> **Performance**: `@WebMvcTest` is dramatically faster than `@SpringBootTest` because it doesn't initialize the database or business layer. Use it for checking HTTP status codes and JSON formatting.


---

**Q: MockMvc**
> "This allows you to test your Controllers without starting a real HTTP Server.
> It simulates incoming HTTP requests.
>
> `mockMvc.perform(get("/api/users")).andExpect(status().isOk())`
>
> It tests the *web layer logic* (routing, headers, cookies) but calls the Java methods directly, skipping the network stack."

**Indepth:**
> **Integration**: MockMvc is usually used with `@WebMvcTest`. It allows checking the *content* of the response (body, headers) using a fluent assertion API.

