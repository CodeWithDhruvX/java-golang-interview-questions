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

