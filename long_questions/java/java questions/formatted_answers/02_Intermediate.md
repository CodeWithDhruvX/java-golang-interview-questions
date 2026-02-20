# Intermediate Level Java Interview Questions

## From 03 Exceptions And IO
# 03. Exceptions and IO

**Q: Exception Hierarchy in Java**
> "At the top, we have `Throwable`. It has two main children: `Error` and `Exception`.
>
> **Errors** (like `OutOfMemoryError`) are catastrophic issues that your application usually *cannot* recover from. You generally shouldn't try to catch them.
>
> **Exceptions** are issues your application *might* want to catch. This branch splits further into:
> 1.  **Checked Exceptions** (like `IOException`): These are checked at compile-time. You *must* handle them with try-catch or throws.
> 2.  **Unchecked Exceptions** (like `NullPointerException`): These extend `RuntimeException`. They usually represent programming bugs, and you aren't forced to catch them."

**Indepth:**
> **Stack Tracing**: When an exception occurs, the JVM creates an exception object containing the type, message, and the state of the execution stack (frames) at that moment. This stack trace is expensive to generate (filling in the stack frames), so for high-performance scenarios where control flow is managed via exceptions (anti-pattern but exists), one might override `fillInStackTrace()` to skip this cost.
>
> **Error vs Exception**: Errors usually signify serious problems at the JVM level (like `StackOverflowError` due to infinite recursion) which an application shouldn't attempt to catch. Exceptions are conditions that an application might want to catch (like business logic errors or I/O issues).


---

**Q: Validating try-catch, finally combinations**
> "The rule is simple: a `try` block cannot exist alone. It must be followed by either a `catch` block, a `finally` block, or both.
>
> So:
> *   `try-catch` is valid.
> *   `try-finally` is valid.
> *   `try-catch-finally` is valid.
>
> *However*, you cannot have just `try` by itself. And if you have a `catch` or `finally`, there must be a `try` immediately before it."

**Indepth:**
> **Resource Management**: While `try-finally` ensures cleanup, it doesn't handle exceptions itself. It propagates them up the stack. This was the standard pattern for closing resources (DB connections, streams) before Java 7.
>
> **Best Practice**: Prefer `try-with-resources` over `try-finally` for cleanup, as it handles edge cases like "variable scope" and "exception suppression" (where an exception during close masks the original exception) much better.


---

**Q: throw vs throws**
> "**throw** (singular) is an action. You use it *inside* a method to explicitly throw an exception object. Example: `throw new RuntimeException("Error!");`.
>
> **throws** (plural) is a declaration. You use it in the *method signature* to warn callers that this method *might* throw specific exceptions. Example: `public void readFile() throws IOException`."

**Indepth:**
> **Runtime vs Compile-time**: You are not *required* to declare `throws` for Unchecked Exceptions (`RuntimeException` and its children), but you *must* declare them for Checked Exceptions.
>
> **rethrow**: You can catch an exception, wrap it in a custom exception (or a generic `RuntimeException`), and `throw` that new exception. This is common in layered architectures to abstract implementation details (e.g., hiding `SQLException` behind a `DataAccessException`).


---

**Q: final, finally, finalize**
> "**final** is a modifier. It means 'cannot be changed'. A final variable is a constant; a final class cannot be extended.
>
> **finally** is a block used in exception handling. It executes *always*, whether an exception happens or not. It's perfect for cleanup code.
>
> **finalize** is a deprecated method on the `Object` class. It was designed to run before an object is garbage collected, but it's unpredictable and dangerous. Avoid it."

**Indepth:**
> **Phantom References**: `finalize()` is related to the garbage collection cycle. A better alternative for cleanup actions when an object is reclaimed is using `PhantomReference` with a `ReferenceQueue`.
>
> **Initialization**: `final` variables must be initialized exactly once. This can happen at declaration, in an instance initializer block, or in the constructor. If not initialized by the time the constructor finishes, it's a compile-time error.


---

**Q: What is Try-with-Resources?**
> "Introduced in Java 7, this is a feature that automatically closes resources for you.
>
> Instead of manually closing a file or socket in a `finally` block (which is messy), you declare the resource in parentheses after the `try` keyword: `try (BufferedReader br = new BufferedReader...)`.
>
> When the try block finishes (normally or with an exception), Java automatically calls `.close()` on that resource. The resource just needs to implement the `AutoCloseable` interface."

**Indepth:**
> **Suppressed Exceptions**: If the `try` block throws an exception, and then the resource `.close()` *also* throws an exception, the exception from `.close()` is "suppressed". It is added to the primary exception (from the try block) and can be retrieved via `getSuppressed()`. In standard `try-finally`, the exception from `finally` would overwrite the original one, losing the root cause.
>
> **Order**: Resources are closed in the reverse order of their creation.


---

**Q: Checked vs Unchecked Exception? (When to use?)**
> "**Checked Exceptions** are for 'recoverable' conditions that a reasonable application checks for. For example, `f IleNotFoundException`. If a file is missing, maybe you ask the user to pick another one. The compiler forces you to handle these.
>
> **Unchecked Exceptions** (RuntimeExceptions) usually indicate programming errors, like logic bugs (`IndexOutOfBounds`, `NullPointer`). You generally fix the code rather than trying to catch these.
>
> Use Checked exceptions when the caller *can* do something about the error. Use Unchecked exceptions for coding errors or unrecoverable system failures."

**Indepth:**
> **Controversy**: Many modern languages (Kotlin, C#) use only unchecked exceptions. The argument is that checked exceptions force boilerplate code (`try-catch` blocks that just log and rethrow) and break encapsulation (adding a new exception to a method signature breaks all callers). Spring Framework, for instance, wraps almost all SQL/JPA checked exceptions into unchecked `DataAccessException`.


---

**Q: Custom Exception creation**
> "It's easy. You just create a class and extend either `Exception` (for a checked exception) or `RuntimeException` (for an unchecked exception).
>
> Usually, you'll want to implement constructors that take a message and a cause (another exception), calling `super(message, cause)` so that you preserve the stack trace and error details."

**Indepth:**
> **Inheritance**: Creating a hierarchy of custom exceptions allows for fine-grained error handling. For example, `PaymentException` could have subclasses `InsufficientFundsException` and `PaymentGatewayTimeoutException`, allowing the calling code to handle specific scenarios differently.
>
> **Logging**: Always include the original `cause` when wrapping exceptions. Use constructors that accept `Throwable cause` so the stack trace reveals the root source of the problem.


---

**Q: What happen if you throw exception in finally block?**
> "This is dangerous. If an exception is thrown inside the `finally` block, it will typically *swallow* any exception that was thrown in the `try` block.
>
> For example, if your `try` block throws an 'Error A', but then your `finally` block throws 'Error B', the caller will only see 'Error B'. 'Error A' is completely lost. This makes debugging a nightmare, so be very careful in `finally` blocks."

**Indepth:**
> **Control Flow**: Returning a value from a `finally` block creates a similar issue—it discards any exception thrown in the `try` block and returns normally. Both patterns (throwing or returning in finally) are considered severe anti-patterns because they hide errors.


---

**Q: Exception Propagation**
> "When an exception occurs in a method, if it's not caught there, it drops down (or 'bubbles up') the call stack to the method that called it. This continues until it hits a `catch` block that can handle it.
>
> If it reaches the bottom of the stack (the `main` method) without being caught, the thread terminates and prints the stack trace."

**Indepth:**
> **Uncaught Exception Handler**: You can set a default `UncaughtExceptionHandler` for a Thread (or all threads). This is the last line of defense to log the error before the thread dies, which is crucial for debugging production crashes in background threads.


---

**Q: What is Serialization?**
> "**Serialization** is the process of converting a Java object into a stream of bytes. You do this to save the object to a file, send it over a network, or store it in a database.
>
> **Deserialization** is the reverse: taking that stream of bytes and reconstructing the Java object in memory."

**Indepth:**
> **Security Risks**: Deserialization is a major vector for security vulnerabilities (e.g., remote code execution). If an attacker can manipulate the byte stream, they can force the JVM to instantiate malicious objects. Always validate streams or use formats like JSON/XML instead of native Java serialization for untrusted data.
>
> **Marker Interface**: `Serializable` has no methods. It’s strictly a flag for the JVM.


---

**Q: serialVersionUID significance**
> "`serialVersionUID` is a unique ID number for your class. It acts like a version control stamp.
>
> When you deserialize an object, Java checks if the ID of the incoming bytes matches the ID of the class currently in your code. If they don't match (maybe you added a field to the class since saving the data), Java throws an `InvalidClassException`.
>
> It's best practice to declare this explicitly so you control compatibility."

**Indepth:**
> **Calculation**: If you don't define it, Java calculates it based on class details (methods, fields). Even a small change (like adding a private method) changes the calculated ID, breaking compatibility.
>
> **Strategy**: Always define it as `1L`. This tells Java "I am taking responsibility for version compatibility."


---

**Q: transient keyword**
> "**transient** is used during serialization. It tells Java: 'Ignore this field. Do not save it.'
>
> You use it for sensitive data like passwords (you don't want those written to disk) or for derived fields that you can just recalculate when you load the object back."

**Indepth:**
> **Use Case**: Besides security, `transient` is used for fields that don't make sense to serialize, like a connection to a database or a reference to a thread. These objects are tied to the current JVM execution context and cannot be meaningfully restored in a different JVM.


---

**Q: Externalizable interface vs Serializable**
> "**Serializable** is a marker interface. It's 'automagic'—Java uses reflection to save all non-transient fields for you. It's easy but slow.
>
> **Externalizable** gives you full control. You *must* implement `writeExternal()` and `readExternal()` defining exactly how to save and load the data. It's faster and more compact, but requires more code."

**Indepth:**
> **Performance**: `Externalizable` is often much faster because you write only the specific fields you need, avoiding the overhead of Reflection and the metadata that normal Serialization writes.
>
> **Requirement**: An `Externalizable` class *must* have a public no-arg constructor, because the JVM instantiates it normally before calling `readExternal()`. `Serializable` uses magic to allocate the object without running the constructor.


---

**Q: Byte Stream vs Character Stream**
> "**Byte Streams** (like `InputStream`, `OutputStream`) assume you are working with raw binary data—images, videos, etc. They read/write byte by byte (8 bits).
>
> **Character Streams** (like `Reader`, `Writer`) are designed for text. They handle character encoding (like UTF-8) automatically for you. They read/write character by character (16 bits for Java chars).
>
> Rule of thumb: If it's human-readable text, use Character Streams. For everything else, use Byte Streams."

**Indepth:**
> **Translation**: `InputStreamReader` and `OutputStreamWriter` act as bridges. They convert byte streams to character streams using a specified Charset.
>
> **Buffering**: Raw streams (like `FileInputStream`) are slow because they hit the OS/Disk for every byte. Always wrap them in `BufferedInputStream` or `BufferedReader` to minimize syscalls by reading/writing blocks of data at a time.


---

**Q: Scanner vs BufferedReader**
> "**Scanner** is a text scanner that can parse primitives and strings using regex headers. It's great for reading simple formatted input (like `nextInt()`). However, it has a small buffer and is synchronized (slower).
>
> **BufferedReader** reads text efficiently from a character-input stream. It has a large buffer (8KB default) but only reads simple Strings and chunks. It's faster than Scanner.
>
> Use **Scanner** for parsing user input tokens. Use **BufferedReader** for reading long files efficiently."

**Indepth:**
> **Tokenization**: `Scanner` internally uses regex to split input into tokens, which makes it computationally expensive compared to `BufferedReader`'s simple char-reading logic.
>
> **Memory**: `BufferedReader` reads chunks (8KB) ahead of time, which improves throughput for large files on disk compared to reading byte-by-byte or small tokens.


---

**Q: Handling FileNotFoundException vs IOException**
> "`FileNotFoundException` is a subclass of `IOException`.
>
> `IOException` is the general parent for all input/output errors (disk full, permission denied, etc.).
> `FileNotFoundException` is specific: it means 'I tried to open a file at this path, but it wasn't there.'
>
> You typically catch `FileNotFoundException` separately if you want to explicitly tell the user 'Check your file path,' and then catch `IOException` as a fallback for other errors."

**Indepth:**
> **Specificity**: Exception handling should always go from specific to general.
> ```java
> try { ... }
> catch (FileNotFoundException e) { ... } // Specific
> catch (IOException e) { ... } // General
> catch (Exception e) { ... } // Catch-all (Last resort)
> ```
> If you reverse this order, the general `IOException` would catch the file-not-found case, and the specific block would become unreachable code (compiler error).


---

**Q: File vs Path (NIO.2)**
> "**File** is the old-school class (since Java 1.0) representing a file path. It works, but its API is a bit inconsistent and it lacks comprehensive error handling metadata.
>
> **Path** (introduced in Java 7 with NIO.2) is the modern replacement. It's an interface that represents a path in the file system. It works with the `Files` utility class, which offers much better methods for copying, moving, and reading files, with better exception handling."

**Indepth:**
> **Symbolic Links**: `Path` handles symbolic links correctly, whereas `File` can verify usually resolve them safely or consistently across OSs.
>
> **Agnostic**: `Path` is file-system agnostic. It can point to a file inside a ZIP archive (ZipFileSystem) or just on the local disk. `File` is tied strictly to the default OS filesystem.


---

**Q: Breaking Singleton with Serialization**
> "Even if you make a class a Singleton (private constructor), serialization can break it. If you serialize the instance and then deserialize it, Java creates a *new* instance by default!
>
> To fix this, you must implement the `readResolve()` method in your Singleton class. This method lets you replace the deserialized object with your existing Singleton instance, ensuring only one exists."

**Indepth:**
> **Enum Singleton**: The preferred way to implement Singletons today is using a single-element `enum`.
> ```java
> public enum Singleton { INSTANCE; }
> ```
> Java natively handles serialization for Enums, guaranteeing that `INSTANCE` remains a true singleton even after deserialization, without requiring `readResolve()`.


---

**Q: System.out, System.err, System.in**
> "These are the standard streams provided by the System class.
>
> *   `System.in` is the **Standard Input** stream (usually the keyboard).
> *   `System.out` is the **Standard Output** stream (usually the console/terminal).
> *   `System.err` is the **Standard Error** stream. It also goes to the console by default, but IDEs often color it red, and operating systems allow you to redirect it separately from standard output (e.g., logging errors to a file while printing output to the screen)."

**Indepth:**
> **File Descriptors**: These streams utilize file descriptors (0, 1, 2) provided by the OS shell. They are byte streams (`PrintStream` for out/err, `InputStream` for in).
>
> **System.console()**: For secure password entry (no echo on screen), use `System.console().readPassword()`. Using `System.in` would show the password characters as you type them.


---

**Q: Closeable vs AutoCloseable**
> "**AutoCloseable** is the parent interface (introduced in Java 7) that allows a resource to be used in a `try-with-resources` block. It has a `close()` method that throws `Exception`.
>
> **Closeable** is an older interface (extends AutoCloseable) specifically for I/O streams. Its `close()` method throws `IOException` (a more specific exception).
>
> Practically, both let you use try-with-resources, but `Closeable` is strictly for I/O."

**Indepth:**
> **Idempotency**: The `close()` method is required to be idempotent. Calling it multiple times should have no side effects (it should just return if already closed).
>
> **Legacy**: `Closeable` existed before Java 7. `AutoCloseable` was added later to support try-with-resources. They retrofitted `Closeable` to extend `AutoCloseable` so all existing I/O classes would work immediately with the new feature.


## From 07 OOP Basics Practice
# 07. OOP Basics (Practice)

**Q: What are the 4 Pillars of OOP? Explain with real-world examples.**
> "As a recap:
> 1.  **Encapsulation**: Protecting data. Example: A bank account. You can deposit/withdraw (methods), but you can't manually set the balance (data).
> 2.  **Inheritance**: Reusability. Example: 'Mobile Phone' inherits from 'Phone'. It has all phone features + smart features.
> 3.  **Polymorphism**: Flexibility. Example: 'Speak'. A human speaks, a dog barks, a cat meows. Same action, different implementation.
> 4.  **Abstraction**: Simplicity. Example: Driving a car. You use the steering wheel (interface) without knowing how the engine combustion works (implementation)."

**Indepth:**
> **Key Differentiator**: If asked "Which is most important?", argue for **Polymorphism**. It is the core of flexible, testable, and maintainable code (enabling Mocking, Dependency Injection, and Strategy Patterns). Without it, OOP is just data structures with methods.


---

**Q: Difference between Abstract Class and Interface?**
> "*Practice Perspective*:
> Use an **Abstract Class** when you are building a *framework* and want to provide common functionality that subclasses can re-use (like a `BaseRepository`).
> Use an **Interface** when you want to define a specific *role* that different classes can play (like `Serializable` or `Comparable`)."

**Indepth:**
> **Decision Matrix**:
> *   Is-A relationship? -> Abstract Class.
> *   Can-Do relationship? -> Interface.
> *   Need state? -> Abstract Class.
> *   Need multiple inheritance? -> Interface.


---

**Q: What is Polymorphism? (Compile-time vs Runtime)**
> "In a coding interview, you'd say:
> 'Compile-time is **Overloading**—same method name, different args. The compiler picks the right one.'
> 'Runtime is **Overriding**—same method signature in Parent/Child. The JVM picks the right one based on the actual object.'
> One is static, one is dynamic."

**Indepth:**
> **Memory**:
> *   **Overloading**: Resolved at compile time. Fast. No lookup overhead.
> *   **Overriding**: Resolved at runtime via the vtable (Virtual Method Table). Slight performance cost, but negligible in modern JVMs.


---

**Q: Can you override static or private methods?**
> "No.
> **Static**: They are bound to the class. Re-declaring them hides the parent method (Method Hiding).
> **Private**: They are not visible. Re-declaring them creates a new method."

**Indepth:**
> **Wait, what about shadowing?**
> If you define a variable with the same name in the child class, it *shadows* the parent's variable. This is confusing and bad practice.
> Static methods are also "shadowed" (or hidden), whereas instance methods are "overridden".


---

**Q: What is covariant return type?**
> "It means an overriding method can return a *subclass* of the original return type.
> If `Animal.born()` returns `Animal`, then `Dog.born()` can return `Dog`. This saves you from casting the result."

**Indepth:**
> **Compatibility**: This was introduced in Java 5. Before that, you had to return the exact same type and declare variables more broadly. Covariance allows for cleaner, more specific client code.


---

**Q: Composition vs Inheritance. Which is better?**
> "Composition is usually better. It’s more flexible.
> Inheritance is 'white-box' reuse (you see the internals of the parent).
> Composition is 'black-box' reuse (you just use the public API of the component).
> Changes in a superclass propagate to subclasses (fragile base class problem), but changes in a component class rarely break the wrapper class."

**Indepth:**
> **LSP Violation**: Inheritance forces an "Is-A" relationship. If "Square extends Rectangle", you might break assumptions (setting width changes height). Composition ("Square has a Shape") avoids this semantic trap.


---

**Q: What is the super keyword?**
> "It accesses the parent class.
> Code snippet:
> ```java
> public Dog() {
>     super(); // Calls Animal() constructor
> }
> public void eat() {
>     super.eat(); // Calls Animal.eat()
>     System.out.println("Dog eating");
> }
> ```"

**Indepth:**
> **Constructor Rule**: `super()` must be the *first* statement in a constructor. Why? Because the parent must strictly be fully initialized before the child tries to access any inherited state.


---

**Q: Significance of this keyword?**
> "It refers to *this* instance.
> Mostly used to fix shadowing (when param name == field name), or to pass the current object to another helper method."

**Indepth:**
> **Fluent Interfaces**: Returning `this` at the end of a method (`return this;`) allows for method chaining: `obj.setName("A").setAge(10);`. Crucial for Builder patterns.


---

**Q: Can an Interface extend another Interface?**
> "Yes, using `extends`. And it can extend *multiple* interfaces:
> `interface Hero extends Human, Flyable, Strong { ... }`"

**Indepth:**
> **Why?**: Interfaces have no state and no constructors. Extending multiple interfaces just merges their method contracts. There is no risk of the "Diamond Problem" regarding state initialization.


---

**Q: Difference between Overloading and Overriding?**
> "Overloading = New inputs, same name (Static).
> Overriding = New logic, same signature (Dynamic)."

**Indepth:**
> **Annotation**: Always use `@Override`. It forces the compiler to check your work. If you typo the name, the compiler throws an error instead of silently creating a new method.


---

**Q: Can a constructor be private? Why?**
> "Yes. To stop people from saying `new MyClass()`.
> Mandatory for:
> 1.  Singletons.
> 2.  Utility classes (like `java.util.Collections`) where everything is static."

**Indepth:**
> **Reftection Attack**: Private constructors can still be called via Reflection (`setAccessible(true)`). To be truly safe, throw an exception inside the private constructor if it's called a second time.


---

**Q: Default Constructor vs No-Args Constructor?**
> "Default is the invisible one the compiler gives you.
> No-Args is one you write explicitly (`public Foo() {}`).
> If you write *any* constructor, the Default one is gone."

**Indepth:**
> **Trap**: If you add a parameterized constructor `MyClass(String s)`, the compiler *removes* the default no-args constructor. Any code doing `new MyClass()` will suddenly break.


---

**Q: What is Constructor Chaining?**
> "Calling `this(...)` or `super(...)` as the first line of a constructor.
> It ensures code reuse between constructors and guarantees proper initialization order."

**Indepth:**
> **Output Prediction**: In an interview, if they ask for the output of a chain of constructors, remember: **Parent First**. The Object class constructor finishes first, then the Parent constant, then the Child.


---

**Q: Use of instanceof operator?**
> "Checks type safety before casting.
> ```java
> if (obj instanceof String) {
>     String s = (String) obj; // Safe
> }
> ```"

**Indepth:**
> **Pattern Matching**: Since Java 14, use `if (obj instanceof String s)`. It avoids the separate casting line and scope pollution.


---

**Q: What is an Initialization Block?**
> "Code that runs before the constructor.
> ```java
> {
>     // Instance block
>     System.out.println("Object created");
> }
> ```
> Rarely used, usually we just put this logic in the constructor."

**Indepth:**
> **Order**:
> 1. Static blocks (once, when class loads).
> 2. Instance blocks (every object creation).
> 3. Constructor (every object creation, after instance blocks).


---

**Q: Multiple Inheritance in Java?**
> "Classes: No.
> Interfaces: Yes.
> Why? To avoid the Diamond Problem with state/implementation."

**Indepth:**
> **Default Methods**: With Java 8 default methods, you *can* inherit implementation from multiple interfaces. If two interfaces define the same default method, you **must** override it in your class to resolve the conflict.


---

**Q: What is a Marker Interface?**
> "An empty interface.
> `public interface Safe {}`
> It tells the code something special about the class. Like a sticker on a box saying 'Fragile'."

**Indepth:**
> **Modern**: Annotations (`@Entity`, `@Service`) are the modern replacement for marker interfaces. They carry more metadata (values) and are more flexible.


---

**Q: Can abstract class have constructor?**
> "Yes, to initialize its own fields. It runs when a subclass is created."

**Indepth:**
> **Can you call it?**: No, `new AbstractClass()` is a compile error. But the constructor *exists* and is called via `super()` from the concrete subclass.


---

**Q: Shallow Copy vs Deep Copy?**
> "Shallow: Copies the reference (pointer). Both objects point to the same data.
> Deep: Copies the actual data. Objects are independent."

**Indepth:**
> **Clone vs Copy Constructor**: Prefer Copy Constructors (`public Car(Car c)`) over `clone()`. `clone()` is broken, throws checked exceptions, and bypasses constructors.


---

**Q: Immutable Class - How to create one?**
> "Final class. Private final fields. No setters. Defensive copies for mutable fields.
> Example: `String`, `Integer`, `LocalDate`."

**Indepth:**
> **Benefits**: Immutable objects are thread-safe (no synchronization needed), excellent Map keys (hash code never changes), and failure-atomic (state never gets inconsistent).


## From 08 Exceptions IO Practice
# 08. Exceptions and IO (Practice)

**Q: Exception Hierarchy in Java**
> "Top: `Throwable`.
> Left: `Error` (System crash, don't catch).
> Right: `Exception`.
>    - `RuntimeException` (Unchecked, programming bugs).
>    - Everything else (Checked, external failures like IO/SQL)."

**Indepth:**
> **Trick**: Can you catch `Throwable`? Yes, but don't. It catches `OutOfMemoryError` too, which you can't fix.


---

**Q: Validating try-catch, finally combinations**
> "You need a `try`.
> You need at least one `catch` OR one `finally`.
> `try-catch`, `try-finally`, `try-catch-finally` are all good.
> `try` alone is a syntax error."

**Indepth:**
> **Interview Question**: "Can a try block be followed by nothing?" -> No. Compilation error.


---

**Q: throw vs throws**
> "`throw` = Action. 'I am throwing the ball now.' (`throw new Exception()`)
> `throws` = Warning. 'I might throw the ball.' (`public void foo() throws Exception`)"

**Indepth:**
> **Memory Aid**: `throw` is like a baseball pitcher (action). `throws` is like the warning sign on the fence (declaration).


---

**Q: final, finally, finalize**
> "Three 'F's:
> 1.  **final**: Restricted. (Constant variable, non-overridable method, non-inheritable class).
> 2.  **finally**: Guaranteed execution. (Cleanup code block).
> 3.  **finalize**: Dead. (Old GC cleanup method, don't use it)."

**Indepth:**
> **Gotcha**: Does `finally` run if `System.exit(0)` is called in `try`? -> No. The JVM halts immediately.


---

**Q: What is Try-with-Resources?**
> "Java 7 feature.
> `try (var r = new Resource()) { ... }`.
> It calls `r.close()` for you automatically.
> Replaces the messy `finally { if (r!=null) r.close(); }` pattern."

**Indepth:**
> **Scope**: The variable `br` is only visible *inside* the try block. You can't use it in the `catch` block (it's arguably closed by then) or after the block.


---

**Q: Checked vs Unchecked Exception?**
> "Checked: Compiler forces you to handle it (`IOException`). Use for external problems.
> Unchecked: Compiler ignores it (`NullPointerException`). Use for logic errors."

**Indepth:**
> **Debate**: Some architects say "Checked exceptions are a failed experiment" because they break method signatures when requirements change.


---

**Q: Custom Exception creation**
> "Class `MyException` extends `Exception` (for Checked) or `RuntimeException` (for Unchecked).
> Always call `super(message)`."

**Indepth:**
> **Tip**: Don't just extend `Exception`. Think: "Is this a usage error (Unchecked) or a system failure (Checked)?"


---

**Q: What happens if you throw exception in finally block?**
> "It swallows the original exception. The caller only sees the new one.
> It's a common cause of 'missing' error logs."

**Indepth:**
> **Anti-pattern**: This is called "Exception Masking". It makes debugging impossible because the root cause stack trace is gone.


---

**Q: Exception Propagation**
> "Exceptions float up the stack until someone catches them.
> If nobody catches it, the thread dies."

**Indepth:**
> **Stack Trace Cost**: Creating an exception is slow because filling the stack trace takes time. Don't use exceptions for normal control flow (loops, ifs).


---

**Q: What is Serialization?**
> "Object -> Bytes (Serialization).
> Bytes -> Object (Deserialization).
> Used for caching, networking, or saving state to disk."

**Indepth:**
> **Formats**: Java serialization is native but brittle. JSON (Jackson/Gson) is text-based, cross-language, and standard for web APIs.


---

**Q: serialVersionUID significance**
> "It's a version number for your class.
> If you change the code but load old data, the IDs won't match, and Java throws `InvalidClassException`.
> Always define it manually: `private static final long serialVersionUID = 1L;`."

**Indepth:**
> **Default**: Generates a hash based on class members. Even adding a simple logging method changes the hash and breaks deserialization!


---

**Q: transient keyword**
> "'Don't serialize me.'
> Use it for passwords or derived data (like `age` if you already have `birthDate`)."

**Indepth:**
> **Static**: Static variables are *never* serialized because they belong to the class, not the object instance. You don't need `transient` for statics.


---

**Q: Externalizable vs Serializable**
> "**Serializable**: Automatic. Easy. Slower.
> **Externalizable**: Manual (`writeExternal`). Harder. Faster. You control the exact byte layout."

**Indepth:**
> **Constructor**: `Serializable` uses "magic" (Unsafe) to create objects without constructors. `Externalizable` *calls* the no-arg constructor.


---

**Q: Byte Stream vs Character Stream**
> "Byte (`InputStream`): Raw data (Images, Binary).
> Character (`Reader`): Text (Strings, Files). Handles Unicode/Encoding for you."

**Indepth:**
> **Conversion**: Use `InputStreamReader` to turn Bytes into Characters. You need to specify the Charset (UTF-8).


---

**Q: Scanner vs BufferedReader**
> "**Scanner**: Parses tokens (`nextInt`). Slower. Good for console input.
> **BufferedReader**: Reads references lines (`readLine`). Faster. Good for files."

**Indepth:**
> **Parsing**: Scanner has methods like `hasNextInt()`, making it great for coding competitions or simple CLI tools.


---

**Q: Handling FileNotFoundException vs IOException**
> "Catch `FileNotFoundException` first to tell the user 'check the filename'.
> Catch `IOException` second to handle 'disk full' or 'permission denied'."

**Indepth:**
> **Hierarchy**: `FileNotFoundException extends IOException`. Always catch the child first, parent second.


---

**Q: File vs Path**
> "**File**: Old, legacy class.
> **Path**: New (Java 7 NIO). Part of a better API (`Files` class). Use Path for new code."

**Indepth:**
> **Utilities**: `Files.readAllLines(path)` is a one-liner to read a text file. Much better than the old `BufferedReader` loop boilerplate.


---

**Q: Breaking Singleton with Serialization**
> "Deserialization creates a new object instance.
> Fix it by adding `protected Object readResolve() { return INSTANCE; }` to your Singleton."

**Indepth:**
> **Enum**: Enums handles this automatically. Best Singleton implementation.


---

**Q: System.out, System.err, System.in**
> "Standard Output (Console), Standard Error (Console/Log), Standard Input (Keyboard)."

**Indepth:**
> **Redirection**: You can change these targets using `System.setOut(new PrintStream(...))`. Useful for redirecting console output to a log file.


---

**Q: Closeable vs AutoCloseable**
> "**AutoCloseable**: Generic 'can be closed'. Method throws `Exception`.
> **Closeable**: Specific for I/O. Method throws `IOException`.
> Both work with try-with-resources."

**Indepth:**
> **Idempotent**: `close()` must be safe to call multiple times.


## From 09 Java Fundamentals Practice
# 09. Java Fundamentals (Practice)

**Q: Explain static keyword**
> "'One per class'.
> Static variables are shared. Static methods can be called without an object. Static blocks run once at class load."

**Indepth:**
> **Memory**: Static variables live in the Heap (since Java 8, previously PermGen). They stay alive for the duration of the application.


---

**Q: What does volatile do?**
> "Guarantees *Visibility*.
> Prevents threads from caching variables locally. Forces reads/writes to go to main memory.
> Does *not* guarantee Atomicity."

**Indepth:**
> **CPU Cache**: Without `volatile`, thread A might change a variable in L1 cache, but thread B checks its own L2 cache and sees the old value. `volatile` invalidates local caches.


---

**Q: Comparing Objects: == vs equals()**
> "`==`: Reference check (Same memory address?).
> `.equals()`: Value check (Same content?)."

**Indepth:**
> **Strings**: `String s = "Hello"` uses the pool. `new String("Hello")` skips the pool. `==` fails on the latter.


---

**Q: Common Object methods**
> "`toString()`: Text representation.
> `equals()`: Logical equality.
> `hashCode()`: Bucket address for HashMaps. Note: If equals() is true, hashCode() MUST be same."

**Indepth:**
> **Contract**: If you override `equals()`, you *must* override `hashCode()`. Otherwise, HashMaps won't be able to find your object.


---

**Q: finalize() method**
> "Deprecated. Unreliable. Replaced by `Cleaner` or `try-with-resources`."

**Indepth:**
> **Zombie**: `finalize` can resurrect an object by assigning `this` to a global variable.


---

**Q: Wrapper Classes & Autoboxing**
> "Wrappers (`Integer`) let primitives (`int`) act like Objects.
> Autoboxing is the automatic conversion between them (`int` -> `Integer`)."

**Indepth:**
> **Cost**: Autoboxing is expensive (object creation). Avoid it in tight loops. Use `int[]` instead of `List<Integer>` for heavy math.


---

**Q: Integer Cache**
> "Java pre-creates Integer objects for -128 to 127.
> `Integer a = 127; Integer b = 127;` -> `a == b` is True.
> `Integer a = 128; Integer b = 128;` -> `a == b` is False."

**Indepth:**
> **Config**: `Integer` cache high end (127) can be increased with `-XX:AutoBoxCacheMax`.


---

**Q: BigInteger and BigDecimal**
> "**BigInteger**: For massive integers (crypto, infinite size).
> **BigDecimal**: For money. Handles decimals exactly without floating-point errors."

**Indepth:**
> **Money**: `BigDecimal` constructors: Always use the String constructor `new BigDecimal("0.1")`. The double constructor `new BigDecimal(0.1)` is unpredictable!


---

**Q: What is Type Erasure?**
> "Generics exist only at compile time. At runtime, `List<String>` becomes just `List`.
> This preserves backward compatibility with old Java."

**Indepth:**
> **Reflection**: You can bypass generics with Reflection or raw types (`List l = new ArrayList<String>(); l.add(10);` works at runtime!).


---

**Q: Wildcards in Generics**
> "`<?>`: Anything.
> `<? extends Number>`: Number or children (Read-onlyish).
> `<? super Integer>`: Integer or parents (Write-capable)."

**Indepth:**
> **PECS**: "Producer Extends, Consumer Super". Use `extends` when reading, `super` when writing.


---

**Q: Generic Methods**
> "Defining `<T>` on the method instead of the class.
> `public <T> T pickOne(T a, T b) { ... }`"

**Indepth:**
> **Inference**: Logic mostly inferred by compiler. `Collections.emptyList()` relies on this to know what type of list to return based on the variable you assign it to.


---

**Q: What is Reflection?**
> "Code that looks at itself.
> Used by frameworks (Spring, Hibernate) to inspect classes, fields, and methods at runtime.
> Powerful but slow and unsafe."

**Indepth:**
> **Performance**: Reflection disables JIT optimizations (inlining). It is roughly 2x-50x slower than direct calls depending on the operation/JVM version.


---

**Q: Access Private Field using Reflection?**
> "`field.setAccessible(true);`
> It overrides the private check."

**Indepth:**
> **Modules**: Java 9+ Modules strongly encapsulate internals. `setAccessible` might fail if the module doesn't "open" the package.


---

**Q: What is the Class class?**
> "The metadata object that describes a class. `String.class` contains methods/fields info for String."

**Indepth:**
> **Loading**: `Class.forName("com.example.Foo")` loads the class dynamically. Used in JDBC drivers.


---

**Q: Custom Annotations**
> "`@interface MyTag`.
> Used to add metadata. Processed via Reflection or Compiler."

**Indepth:**
> **Logic**: Annotations are passive. You need a processor (AspectJ, Spring AOP, Reflection) to make them do something.


---

**Q: Breaking Singleton using Reflection**
> "You can access the private constructor and call it.
> Fix: Throw exception in constructor if `instance != null`."

**Indepth:**
> **Enum**: Enums cannot be instantiated via Reflection. `Constructor.newInstance` explicitly throws an exception for Enum classes.


---

**Q: Private vs Default vs Protected vs Public**
> "Private: Class only.
> Default: Package only.
> Protected: Package + Subclasses.
> Public: Everyone."

**Indepth:**
> **Encapsulation**: Using `private` is key to loose coupling. If it's private, you can change it without breaking other classes.


## From 16 Data Structures Collections
# 16. Data Structures (Collections Framework)

**Q: ArrayList vs LinkedList**
> "Think of **ArrayList** as a dynamic array. It forces all elements to sit next to each other in memory.
> *   **Access** is instant (O(1)) because the computer knows exactly where index 5 is.
> *   **Adding/Removing** is slow (O(n)) because if you delete element 0, the computer has to shift every other element one step to the left to fill the gap.
>
> **LinkedList** is a chain of nodes. Each node holds data and a pointer to the next node.
> *   **Access** is slow (O(n)) because to get to the 5th element, you have to start at the head and hop 5 times.
> *   **Adding/Removing** is fast (O(1)) *if you are already at that spot*, because you just change two pointers. No shifting required.
>
> **Verdict**: Use `ArrayList` 99% of the time. Modern CPU caches love contiguous memory, so `ArrayList` is usually faster even for inserts unless the list is massive."

**Indepth:**
> **Resize Strategy**: When ArrayList is full, it creates a new array 50% larger (1.5x) and copies elements. This resizing is costly (O(n)), so always initialize with a capacity if you know the size (`new ArrayList<>(1000)`).


---

**Q: List.of() vs Arrays.asList()**
> "**List.of()** (Java 9+) creates a truly **Immutable List**.
> *   You cannot add/remove elements.
> *   You cannot even set/replace elements.
> *   It doesn't allow `null` elements.
>
> **Arrays.asList()** creates a **Fixed-Size List** backed by an array.
> *   You cannot add/remove (throws exception).
> *   But you **can** replace items (`set()`), and those changes write through to the original array.
> *   It allows `null`."

**Indepth:**
> **Best Practice**: Prefer `List.of()` for constants. Use `Arrays.asList()` only when working with legacy APIs that expect it, or when you need a write-through view of an array.


---

**Q: SubList() caveat**
> "`subList(from, to)` doesn't create a new list. It returns a **View** of the original list.
>
> If you start modifying the *original* list (adding/removing items) after creating a sublist, the sublist becomes undefined and will likely throw a `ConcurrentModificationException` the next time you touch it.
>
> Always treat the sublist as temporary, or wrap it in a `new ArrayList<>(list.subList(...))` to detach it."

**Indepth:**
> **Memory Leak**: In old Java versions, `subList` held a reference to the entire original parent list, preventing GC. New versions copy or are smarter, but referencing sublists is still risky if the parent list is long-lived.


---

**Q: Iterator vs For-Each**
> "For-each loops are syntactic sugar. Under the hood, they use an Iterator.
>
> The big difference is **Modification**.
> If you are looping through a list and try to do `list.remove(item)`, you will crash with a `ConcurrentModificationException`.
> To remove items safely while looping, you **must** use the `Iterator` explicitly and call `iterator.remove()`."

**Indepth:**
> **Java 8**: `Collection.removeIf(filter)` is the modern, thread-safe, and readable way to remove elements. `list.removeIf(s -> s.isEmpty())` creates an iterator internally and handles it correctly.


---

**Q: HashSet vs TreeSet vs LinkedHashSet**
> "It's all about **Order**.
>
> 1.  **HashSet**: The fastest (O(1)). It uses a HashMap internally. It makes **no guarantees** about order. You put items in, they come out in a random jumbled mess.
> 2.  **LinkedHashSet**: Slightly slower. It maintains **Insertion Order**. If you put in [A, B, C], you iterate out [A, B, C]. It uses a Doubly Linked List running through the entries.
> 3.  **TreeSet**: The slowest (O(log n)). It keeps elements **Sorted** (Natural order or custom Comparator). It uses a Red-Black Tree. Useful if you need range queries (like 'give me all numbers greater than 50')."

**Indepth:**
> **LinkedHashSet**: It maintains a doubly-linked list running through all its entries. This adds memory overhead (two extra pointers per entry) but gives predictable iteration order.


---

**Q: HashMap Internal Working**
> "A HashMap is an array of 'Buckets' (Node<K,V>).
>
> 1.  **Put(K, V)**: We calculate `hash(Key)`. This gives us an index (bucket location).
> 2.  **Collision**: If two keys land in the same bucket, we store them as a Linked List (or a Tree if the list gets too long, Java 8+).
> 3.  **Get(K)**: We go to the bucket, walk through the chain, and compare keys using `.equals()`.
>
> This is why `hashCode()` and `equals()` must be consistent. If they disagree, you might put an object in one bucket but look for it in another, getting `null`."

**Indepth:**
> **Load Factor**: The default load factor is 0.75. When the map is 75% full, it resizes (doubles the bucket array). Setting it higher (1.0) saves memory but increases collisions. Setting it lower (0.5) reduces collisions but wastes memory.


---

**Q: HashMap vs TreeMap**
> "Similar to the Set comparison:
>
> *   **HashMap**: Fast (O(1)). Unordered. Uses hashing.
> *   **TreeMap**: Slower (O(log n)). Sorted by Key. Uses a Red-Black Tree.
>
> Use `TreeMap` only if you need to iterate keys in alphabetical order or find 'the key just higher than X'."

**Indepth:**
> **NavigableMap**: TreeMap implements `NavigableMap`, giving you powerful methods like `ceilingKey(K)`, `floorKey(K)`, `higherKey(K)`, `lowerKey(K)`.


---

**Q: computeIfAbsent vs putIfAbsent**
> "Both try to add a value if the key is missing.
>
> **putIfAbsent(key, value)**: You pass the *actual value*. Even if the key exists, that value object is created (and then ignored), which might be wasteful if creating it is expensive.
>
> **computeIfAbsent(key, function)**: You pass a *function*. The function is **only** executed if the key is missing. This is lazy and much more efficient for expensive objects."

**Indepth:**
> **Concurrency**: `computeIfAbsent` is atomic in `ConcurrentHashMap`. It guarantees the computation happens only once, even if multiple threads race to compute the value for the same key.


---

**Q: Queue vs Deque vs Stack**
> "**Queue** is standard FIFO (First-In-First-Out). You join the back, leave from the front.
>
> **Deque** (Double Ended Queue) is the superhero interface. You can add/remove from **both** ends.
>
> **Stack**: The `Stack` class remains in Java for compatibility, but it's considered legacy (it's synchronized and extends Vector). **Do not use Stack.**
> If you need a LIFO Stack, use `ArrayDeque` and call `push()` and `pop()`."

**Indepth:**
> **Speed**: `ArrayDeque` is faster than `Stack` and `LinkedList` because it uses a resizable array and doesn't allocate nodes for every item. It's cache-friendly.


---

**Q: PriorityQueue**
> "A standard Queue orders by arrival time. A **PriorityQueue** orders by **Priority**.
>
> When you call `poll()`, you don't get the oldest item; you get the 'smallest' item (according to `compareTo`).
> Internally, it uses a **Min-Heap**. Accessing the top element is O(1), but adding/removing is O(log n) because the heap has to re-balance."

**Indepth:**
> **Use Case**: This is perfect for task scheduling (high priority tasks first) or Dijkstra's shortest path algorithm (explore cheapest path first).


---

**Q: hashCode() and equals() contract**
> "The rule is simple but strict:
>
> 1.  If `a.equals(b)` is true, then `a.hashCode()` **MUST** be equal to `b.hashCode()`.
> 2.  If this is violated, HashMaps break. You will put an object in, and you won't be able to retrieve it because the map looks in the wrong bucket.
>
> Note: The reverse isn't true. Two different objects *can* describe the same hash code (Collision), and the map handles that."

**Indepth:**
> **Consistent Hashing**: If two objects are equal, they MUST hash to the same bucket. If they hashed to different buckets, the Map would check Bucket A, find nothing, and return null, effectively "losing" your object.


---

**Q: Fail-Fast vs Fail-Safe**
> "**Fail-Fast** iterators (like ArrayList, HashMap) throw `ConcurrentModificationException` immediately if they detect that someone else changed the collection while they were iterating. They'd rather crash than show you inconsistent data.
>
> "**Fail-Safe** iterators (like `ConcurrentHashMap`, `CopyOnWriteArrayList`) work on a snapshot or a weakly consistent view. They allow modifications during iteration and won't throw an exception, but they might not show you the very latest data."

**Indepth:**
> **COW**: `CopyOnWriteArrayList` is expensive for writes (it copies the entire array on every add!), but it's perfect for "Read-Mostly" lists like Event Listeners where iteration is frequent but modification is rare.


## From 18 Java Programs Numbers
# 18. Java Programs (Basic Numbers Logic)

**Q: Fibonacci Series**
> "The Fibonacci series is a sequence where each number is the sum of the two preceding ones: 0, 1, 1, 2, 3, 5, 8...
>
> To print it, you start with `a=0` and `b=1`.
> Inside a loop, you print `a`. Then calculate `next = a + b`.
> Shift everything: `a = b` and `b = next`. Repeat `n` times.
>
> If you need the *nth* number specifically, you can use Recursion (`fib(n-1) + fib(n-2)`), but that's very slow (O(2^n)). Iteration (O(n)) is always better for production code."

**Indepth:**
> **Optimization**: A recursive Fibonacci implementation `fib(n-1) + fib(n-2)` has exponential time complexity O(2^n). It recomputes the same values over and over. Always use Memoization (storing results in a map) or simple Iteration for O(n).


---

**Q: Check Prime Number**
> "A prime number is divisible only by 1 and itself.
>
> To check if `n` is prime:
> 1.  Handle edge cases: if `n <= 1`, return false.
> 2.  Loop from `i = 2` up to `Math.sqrt(n)`.
> 3.  If `n % i == 0`, it's not prime. Return false immediately.
> 4.  If the loop finishes without finding a divisor, return true.
>
> Why `sqrt(n)`? Because if a number has a factor larger than its square root, the *other* factor must be smaller than the square root, so we would have already found it."

**Indepth:**
> **Sieve of Eratosthenes**: If you need to find *all* primes up to N, checking them one by one is slow. The Sieve algorithm creates a boolean array and eliminates multiples of each prime found, which is significantly faster.


---

**Q: Factorial of a Number**
> "Factorial of 5 (written 5!) is `5 * 4 * 3 * 2 * 1 = 120`.
>
> You can do this recursively: `return n * factorial(n-1);` (Base case: if n is 0 or 1, return 1).
> Or iteratively: `result = 1;` loop `i` from 2 to `n`, `result *= i`.
>
> **Watch out**: Factorials grow incredibly fast. `13!` allows overflows a standard `int`. `21!` overflows a `long`. For anything bigger, you **must** use `BigInteger`."

**Indepth:**
> **Recursion Depth**: Recursive factorial is prone to `StackOverflowError` for large inputs (thousands), whereas the iterative version can run until memory runs out (assuming you use `BigInteger`).


---

**Q: Palindrome Number**
> "To check if a number like `121` is a palindrome, you need to reverse it mathematically.
>
> 1.  Store original number in `temp`. Initialize `reversed = 0`.
> 2.  While `temp > 0`:
>     *   Get last digit: `digit = temp % 10`.
>     *   Append to reversed: `reversed = (reversed * 10) + digit`.
>     *   Remove last digit: `temp = temp / 10`.
> 3.  Check if `original == reversed`."

**Indepth:**
> **Strings**: You *could* convert the number to a String (`Integer.toString(n)`) and reverse it using `StringBuilder`. This is easier to write but slower due to memory allocation and parsing overhead.


---

**Q: Armstrong Number**
> "An Armstrong number (like 153) is equal to the sum of its digits raised to the power of the number of digits.
> For 153 (3 digits): `1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153`.
>
> The logic is similar to Palindrome:
> 1.  Count digits first.
> 2.  Loop through the number, extract each digit.
> 3.  Add `Math.pow(digit, count)` to a running sum.
> 4.  Compare sum with original."

**Indepth:**
> **Hardcoding**: Many candidates forget to count the digits first and just assume `Math.pow(digit, 3)`. Armstrong numbers are defined by the power of *number of digits* (Example: 1634 is `1^4 + 6^4 + 3^4 + 4^4`).


---

**Q: Swap Two Numbers without Third Variable**
> "This is a classic 'cool trick' question.
>
> Assume `a = 10`, `b = 20`.
> 1.  `a = a + b;` (a becomes 30)
> 2.  `b = a - b;` (30 - 20 = 10, so b is now original a)
> 3.  `a = a - b;` (30 - 10 = 20, so a is now original b)
>
> It works, but in real life, just use a temporary variable. It's more readable and avoids potential integer overflow issues."

**Indepth:**
> **XOR Swap**: A safer way (avoiding overflow) is using XOR: `a = a ^ b; b = a ^ b; a = a ^ b;`. It works because XOR is its own inverse. However, it's less readable and not necessarily faster on modern CPUs.


---

**Q: Check Leap Year**
> "A year is a leap year if:
> 1.  It is divisible by 4.
> 2.  **EXCEPT** if it's divisible by 100, then it is NOT a leap year.
> 3.  **UNLESS** it is also divisible by 400, then it IS a leap year.
>
> Logic: `(year % 4 == 0 && year % 100 != 0) || (year % 400 == 0)`."

**Indepth:**
> **Why 100/400?**: The Earth takes 365.2425 days to orbit. Adding a day every 4 years (365.25) is slightly too much. Skipping 100 years corrects it, but skipping 400 adds it back key to keep the calendar accurate over centuries.


---

**Q: GCD and LCM**
> "**GCD (Greatest Common Divisor)**: Use Euclid's algorithm.
> Recursive: `gcd(a, b)` -> if `b == 0` return `a`, else return `gcd(b, a % b)`.
>
> **LCM (Least Common Multiple)**: Once you have GCD, LCM is easy.
> Formula: `(a * b) / GCD(a, b)`."

**Indepth:**
> **Euclid**: The Euclidean algorithm (`gcd(b, a % b)`) is one of the oldest known algorithms. It works because the GCD of two numbers also divides their difference.


---

**Q: Perfect Number**
> "A number is Perfect if the sum of its proper divisors equals the number itself.
> Example: 6. Divisors are 1, 2, 3. Sum = 1 + 2 + 3 = 6.
>
> Logic: Loop from 1 to `n/2`. If `n % i == 0`, add `i` to sum. Compare sum to `n`."

**Indepth:**
> **Rarity**: Perfect numbers are extremely rare. The first few are 6, 28, 496, 8128. Don't try to find them by brute-force for large ranges.


---

**Q: Sum of Digits**
> "Very similar to reversing a number.
> Loop while `n > 0`:
> 1.  `sum += n % 10;` (Add last digit)
> 2.  `n /= 10;` (Remove last digit)
>
> If you need the 'recursive sum' (sum digits until you get a single digit), use the modulo 9 trick: `return (n == 0) ? 0 : (n % 9 == 0) ? 9 : n % 9;`."

**Indepth:**
> **Digital Root**: The recursive sum of digits until 1 digit remains is called the "Digital Root". The `n % 9` trick works because of congruences in base-10 arithmetic.


## From 19 Java Programs Arrays
# 19. Java Programs (Arrays Logic)

**Q: Largest/Smallest Element in Array**
> "Initialize `max` (or `min`) to `array[0]`.
> Loop through the array from `i = 1`.
> If `array[i] > max`, update `max = array[i]`.
> If `array[i] < min`, update `min = array[i]`.
>
> Simple O(n) scan. No need to sort the array first (which would be O(n log n))."

**Indepth:**
> **Edge Case**: Always initialize `min/max` to `array[0]` or `Integer.MAX_VALUE / MIN_VALUE`. Initializing to `0` is a bug if the array contains negative numbers (e.g., `[-5, -10]` -> max would wrongly be 0).


---

**Q: Reverse an Array**
> "Don't create a new array unless you have to (that wastes memory). Do it **in-place**.
> Use two pointers: `start = 0` and `end = length - 1`.
> While `start < end`:
> 1.  Swap `array[start]` and `array[end]`.
> 2.  Increment `start`.
> 3.  Decrement `end`."

**Indepth:**
> **Collections**: If you have a `List`, you can just call `Collections.reverse(list)`. It uses the exact same swap login internally but is much more readable.


---

**Q: Bubble Sort Logic**
> "It's the simplest sorting algorithm, but also the slowest (O(n^2)).
> Two nested loops.
> *   Outer loop `i` from 0 to n.
> *   Inner loop `j` from 0 to `n - i - 1`.
> *   If `array[j] > array[j+1]`, swap them.
>
> The largest elements 'bubble' to the top (end) of the array with each pass."

**Indepth:**
> **Optimization**: Bubble sort can be optimized with a boolean flag `swapped`. If a full pass happens with zero swaps, the array is already sorted, and you can break early (best case O(n)).


---

**Q: Linear Search vs Binary Search**
> "**Linear Search**: Scan every element one by one. Works on unsorted arrays. Complexity: O(n).
>
> **Binary Search**: Divide and conquer. **Requires sorted array**.
> 1.  Look at the middle element.
> 2.  If target is smaller, ignore the right half.
> 3.  If target is larger, ignore the left half.
> 4.  Repeat.
> Complexity: O(log n). Much faster for large datasets."

**Indepth:**
> **Overflow**: When calculating mid, using `(low + high) / 2` can overflow if low and high are massive integers. The safer way is `low + (high - low) / 2` (or `(low + high) >>> 1`).


---

**Q: Remove Duplicates from Sorted Array**
> "Since it's **sorted**, duplicates are always adjacent. You can do this in O(n) space or O(1) space.
>
> In-place (O(1) space):
> Use a pointer `j` to track the position of unique elements.
> Loop `i` from 0 to n-1.
> If `array[i] != array[i+1]`, place `array[i]` at `array[j]` and increment `j`.
> Finally, store the last element and return `j` as the new length."

**Indepth:**
> **Sets**: The easiest way to remove duplicates? `new LinkedHashSet<>(list)`. It removes duplicates *and* preserves insertion order.


---

**Q: Second Largest Number**
> "You can sort the array and take the second-to-last element, but that's O(n log n). Better to do it in one pass O(n).
>
> Maintain two variables: `largest` and `secondLargest`.
> Loop through the array:
> 1.  If `current > largest`:
>     *   `secondLargest` becomes `largest`.
>     *   `largest` becomes `current`.
> 2.  Else if `current > secondLargest` AND `current != largest`:
>     *   `secondLargest` becomes `current`.
>
> This handles cases where duplicates exist (like [10, 10, 5] -> second largest is 5)."

**Indepth:**
> **Stream API**: `list.stream().distinct().sorted(Comparator.reverseOrder()).skip(1).findFirst()` is a readable one-liner, but significantly slower than the single-pass loop.


---

**Q: Missing Number in Range 1 to N**
> "Math to the rescue!
> The sum of numbers from 1 to N is `N * (N + 1) / 2`.
>
> 1.  Calculate expected sum using the formula.
> 2.  Iterate through the array and calculate the actual `currentSum`.
> 3.  `Missing Number = expectedSum - currentSum`.
>
> This assumes only *one* number is missing. If multiple are missing, you need a HashSet or a boolean array."

**Indepth:**
> **XOR Method**: A robust way to find the missing number (that avoids integer overflow for large sums) is XOR-ing all indices and all values. `(1^2^...^N) ^ (arr[0]^...^arr[n-1])`. The result is the missing number.


---

**Q: Merge Two Arrays**
> "Create a new array of size `len1 + len2`.
> 1.  Loop through first array and copy items.
> 2.  Loop through second array and copy items (offset index by `len1`).
>
> If you need to merge them **in sorted order** (like Merge Sort):
> Use pointers `i` and `j` for the two arrays. Compare `arr1[i]` and `arr2[j]`. Pick the smaller one, add to result, and move the pointer. Once one array is exhausted, dump the rest of the other array."

**Indepth:**
> **In-Place**: Merging two sorted arrays *in-place* without extra space is a hard problem (requires complex shifting logic or gaps). The standard merge sort uses O(n) auxiliary space.


---

**Q: Find Common Elements (Intersection)**
> "Naive approach: Nested loops. For each element in A, scan B. O(n*m). Slow.
>
> Better approach (use HashSet):
> 1.  Dump array A into a `HashSet`.
> 2.  Loop through array B.
> 3.  If `set.contains(element)`, it's a match! (And remove it from set to avoid duplicates).
> Complexity: O(n + m)."

**Indepth:**
> **RetainAll**: `Set` has a built-in method `setA.retainAll(setB)`, which performs an intersection (keeps only elements present in both sets).


---

**Q: Rotate Array (Left/Right)**
> "To rotate an array by `k` positions efficiently (O(n) time, O(1) space), use the **Value Reversal Algorithm**:
>
> 1.  Reverse the whole array.
> 2.  Reverse the first `k` elements.
> 3.  Reverse the remaining `n - k` elements.
>
> This magically shifts everything into the correct place without needing a temporary array."

**Indepth:**
> **Juggling Algorithm**: Another O(n) approach involves moving elements in cycles computed by the GCD of n and k. It's more complex to implement but conceptually elegant.


## From 20 Java Programs Strings
# 20. Java Programs (String Algorithms)

**Q: Reverse a String (Logic)**
> "Don't use `StringBuilder.reverse()` in an interview unless allowed.
> Convert to char array: `char[] chars = str.toCharArray()`.
> Use two pointers: `left = 0`, `right = length - 1`.
> While `left < right`:
> 1.  Swap `chars[left]` and `chars[right]`.
> 2.  Move `left` forward, `right` backward.
> Return `new String(chars)`."

**Indepth:**
> **Surrogates**: `reverse()` in StringBuilder is actually quite complex because it has to handle Surrogate Pairs (Unicode characters that take 2 chars). If you reverse blindly, you might split an Emoji into two invalid characters. The built-in method handles this.


---

**Q: Check Palindrome String**
> "A palindrome reads the same backwards (e.g., 'MADAM').
> Loop `i` from 0 to `length / 2`.
> If `str.charAt(i) != str.charAt(length - 1 - i)`, return false.
> If you finish the loop, it's true."

**Indepth:**
> **Optimization**: You don't need to check the whole string. `i < length / 2` is enough. Checking `i < length` does double the work (checking everything twice) but gives the same result.


---

**Q: Anagram Check**
> "Anagrams have the same characters in different orders (e.g., 'Silent' vs 'Listen').
>
> **Approach 1 (Sorting)**: Clean strings (lowercase, remove spaces). Convert to char arrays. Sort them. `Arrays.equals(arr1, arr2)`. Complexity: O(n log n).
>
> **Approach 2 (Frequency Array - Faster)**:
> Create an int array `counts` of size 26.
> Loop through string 1: `counts[char - 'a']++`.
> Loop through string 2: `counts[char - 'a']--`.
> Finally, check if every element in `counts` is 0. Complexity: O(n)."

**Indepth:**
> **HashMap vs Array**: Use the array approach (`int[26]`) whenever possible. A HashMap has overhead for hashing, resizing, and boxing integers. The array is pure memory access and is order of magnitude faster.


---

**Q: Count Vowels and Consonants**
> "Iterate through the string.
> Convert char to lowercase.
> If it's 'a', 'e', 'i', 'o', 'u', increment `vowels`.
> Else if it's between 'a' and 'z', increment `consonants`.
> Ignore numbers and symbols."

**Indepth:**
> **Regex**: You can also use `str.replaceAll("[^aeiouAEIOU]", "").length()` to count vowels. It's shorter code but much slower due to compiling the regex.


---

**Q: Find First Non-Repeated Character**
> "This requires two passes.
>
> 1.  **Count Frequencies**: Use a `HashMap<Character, Integer>` (or `int[256]` array for ASCII). Loop through string and fill the map.
> 2.  **Check Order**: Loop through the **String** again (not the map!).
> 3.  The first character where `map.get(char) == 1` is your answer."

**Indepth:**
> **LinkedHashMap**: Alternatively, use a `LinkedHashMap` to store counts. Since it preserves insertion order, you can just iterate through the *Map* afterwards and pick the first one with count 1.


---

**Q: String to Integer (atoi)**
> "Converting '123' to `int` manually.
>
> 1.  Initialize `result = 0`.
> 2.  Loop through digits.
> 3.  `int digit = char - '0';` (ASCII magic).
> 4.  `result = result * 10 + digit;`
>
> **Edge Cases**:
> *   Handle negative sign at index 0.
> *   Check for non-digit characters (throw exception).
> *   Check for integer overflow (if result exceeds `Integer.MAX_VALUE`)."

**Indepth:**
> **Overflow Logic**: To check for overflow *before* it happens: `if (result > Integer.MAX_VALUE / 10 || (result == Integer.MAX_VALUE / 10 && digit > 7))`.


---

**Q: Integer to String**
> "The easy way: `String.valueOf(123)`.
>
> The usage way (if asked logic):
> Use a `StringBuilder`.
> While `num > 0`:
> 1.  `digit = num % 10`.
> 2.  Append digit to builder.
> 3.  `num /= 10`.
> Finally, reverse the builder (because you extracted digits backwards)."

**Indepth:**
> **Log10**: You can determine the number of digits using `Math.log10(num) + 1` to pre-allocate the StringBuilder capacity, avoiding resizing.


---

**Q: Reverse Words in a Sentence**
> "Input: 'Hello World Java'. Output: 'Java World Hello'.
>
> 1.  `String[] words = str.split(\" \");`
> 2.  Use a `StringBuilder`.
> 3.  Loop through `words` array **backwards** (from `len-1` down to 0).
> 4.  Append word + space.
> 5.  Trim the result."

**Indepth:**
> **Whitespace**: `split(" ")` leaves empty strings if there are multiple spaces ("Hello   World"). Use `split("\\s+")` to treat multiple spaces as a single delimiter.


---

**Q: Check for Subsequence**
> "Is 'ace' a subsequence of 'abcde'? Yes.
> Use two pointers: `i` for small string, `j` for big string.
> While `i < smallLen` and `j < bigLen`:
> *   If `small[i] == big[j]`, increment `i` (found a match).
> *   Always increment `j` (keep moving in big string).
>
> If `i == smallLen`, you found all characters in order."

**Indepth:**
> **Iterators**: A cleaner way (if allowed) is to use `Iterator` on the big string and advance it only when a match is found. It's the same logic but more abstract.


---

**Q: Rotation Check**
> "Is 'BCDA' a rotation of 'ABCD'?
>
> The trick:
> 1.  Check if lengths are equal.
> 2.  Concatenate original string with itself: `DoubleStr = "ABCD" + "ABCD" -> "ABCDABCD"`.
> 3.  Check if 'BCDA' is a substring of `DoubleStr`.
> `return (str1.length() == str2.length()) && (str1 + str1).contains(str2);`"

**Indepth:**
> **KMP Algorithm**: `contains()` is naïve (O(n*m)). For massive strings, KMP (Knuth-Morris-Pratt) is O(n+m) because it avoids backtracking by creating a prefix table.


---

**Q: Permutations of a String**
> "This is a recursive backtracking problem.
> Method `permute(String str, String answer)`:
> 1.  **Base Case**: If `str` is empty, print `answer`.
> 2.  **Recursive Step**: Loop `i` from 0 to length.
>     *   Pick character at `i`.
>     *   Rest of string = `substring(0, i) + substring(i+1)`.
>     *   Call `permute(rest, answer + char)`."

**Indepth:**
> **Complexity**: There are `n!` permutations. This algorithm is O(n * n!). Even for a small string like "12characters", this will run forever. It's only feasible for inputs of length <= 10.


## From 21 Java Programs Patterns OOP
# 21. Java Programs (Patterns & OOP Concepts)

**Q: Star Patterns (Triangle, Pivot, Diamond)**
> "Pattern printing is all about nested loops.
>
> 1.  **Right Triangle**:
>     *   Outer loop `i` from 1 to n (rows).
>     *   Inner loop `j` from 1 to `i` (columns).
>     *   Print `*`.
>
> 2.  **Pyramid (Centered)**:
>     *   Outer loop `i`.
>     *   Inner loop 1: Print spaces (`n - i`).
>     *   Inner loop 2: Print stars (`2*i - 1`).
>
> 3.  **Diamond**:
>     *   Print top half (Pyramid).
>     *   Print bottom half (Inverted Pyramid).
>
> **Tip**: Focus on the relation between Row number and Count of stars/spaces. Don't memorize code; derive the formula."

**Indepth:**
> **Formatting**: Use `System.out.printf()` or `String.format()` for complex patterns to align columns perfectly (e.g., %4d) instead of manually calculating spaces.


---

**Q: Floyd's Triangle (0-1 Pattern)**
> "1
> 0 1
> 1 0 1
>
> Logic:
> Loop `i` for rows, `j` for cols.
> If `(i + j)` is Even, print 1.
> If `(i + j)` is Odd, print 0.
> Alternatively, use a boolean flag and flip it `!flag` every time."

**Indepth:**
> **Parity**: This pattern is essentially checking checking the parity (even/odd) of the grid coordinates. It's equivalent to a chessboard coloring problem.


---

**Q: Pascal's Triangle**
> "1
> 1 1
> 1 2 1
>
> Each number is the sum of the two numbers directly above it.
>
> **Efficient Way**:
> Use the Combination formula `nCr`.
> `Value = (Previous_Value * (Row - Col + 1)) / Col`.
> This allows you to calculate the next number in the row using the previous number, without calculating factorials every time."

**Indepth:**
> **Binomial Coefficients**: Each entry is `nCr` (n choose r). The sum of numbers in row `n` is `2^n`. This is a powerful property used in probability theory.


---

**Q: Spiral Pattern (Number Grid)**
> "Input: 3. Output:
> 3 3 3 3 3
> 3 2 2 2 3
> 3 2 1 2 3
> 3 2 2 2 3
> 3 3 3 3 3
>
> This looks hard but has a trick.
> For every cell (i, j), the value is `n - min(distance to any wall)`.
> Distance to Top: `i`
> Distance to Bottom: `size - 1 - i`
> Distance to Left: `j`
> Distance to Right: `size - 1 - j`
> Take the minimum of these 4 distances, and subtract from `n`."

**Indepth:**
> **Generalization**: This "Distance Transform" logic works for any grid filling problem. Instead of simulating the movement (right, down, left, up), you calculate the value based on coordinates `(i, j)`.


---

**Q: Singleton Class (Implementation)**
> "To limit a class to exactly one instance:
>
> 1.  Private Static Variable: `private static Singleton instance;`
> 2.  Private Constructor: `private Singleton() {}` (stops `new S()`).
> 3.  Public Static Method:
> ```java
> public static Singleton getInstance() {
>     if (instance == null) {
>         instance = new Singleton();
>     }
>     return instance;
> }
> ```
> **Note**: For thread safety, use 'Double-Checked Locking' or simpler: `private static final Singleton INSTANCE = new Singleton();` (Eager initialization)."

**Indepth:**
> **Bill Pugh**: The thread-safe, efficient, and clean way is the **Bill Pugh Singleton Implementation**:
> `private static class Holder { static final Singleton INSTANCE = new Singleton(); }`
> `public static Singleton getInstance() { return Holder.INSTANCE; }`
> It uses the classloader guarantees for thread safety and is lazy-loaded.


---

**Q: Immutable Class**
> "How to make a class like `String`:
>
> 1.  Make class `final` (so no one can extend and override it).
> 2.  Make all fields `private` and `final`.
> 3.  No Setters.
> 4.  Initialize via Constructor.
> 5.  **Crucial**: If a field is mutable (like `Date` or `List`), don't return the original reference in the Getter. Return a copy (clone). Otherwise, the caller can modify the internal state."

**Indepth:**
> **Defensive Copying**: Both in the constructor (when accepting a mutable object) and in the getter (when returning it). If you miss either one, immutability is broken.


---

**Q: Method Overloading vs Overriding**
> "**Overloading** happens in the **same class**.
> *   Same method name.
> *   **Different parameters** (type or count).
> *   Return type doesn't matter.
> *   Compile-time polymorphism.
>
> **Overriding** happens in **Child classes**.
> *   Same method name.
> *   **Same parameters**.
> *   Runtime polymorphism.
> *   Used to change behavior inherited from Parent."

**Indepth:**
> **Binding**: Overloading is **Static Binding** (Early Binding) - Compiler decides. Overriding is **Dynamic Binding** (Late Binding) - JVM decides.


---

**Q: Abstract Class vs Interface**
> "**Interface**: Describes **Capabilities** (what it *can* do). 'I can Fly, I can Swim'. A class can implement multiple capabilities. Use for defining contracts.
>
> **Abstract Class**: Describes **Identity** (what it *is*). 'I am a Vehicle'. A class can only be ONE thing (single inheritance). Use when you have shared state (variables) or concrete code that children should reuse."

**Indepth:**
> **Evolution**: Interfaces can now have `default` methods (behavior) and `static` constants (state-ish). Note: They still can't have instance variables (true state). This blurs the line, but the "Is-A" vs "Can-Do" distinction remains.


---

**Q: Deep Copy vs Shallow Copy**
> "**Shallow Copy** (Default `clone()`):
> Copies the object, but the fields still point to the same references.
> `NewObject.list == OldObject.list`. Adding to one affects both.
>
> **Deep Copy**:
> Recursively copies everything.
> `NewObject.list = new ArrayList(OldObject.list)`. The objects are completely independent."

**Indepth:**
> **Copy Constructor**: The best way to implement Deep Copy is using a **Copy Constructor**: `public User(User other)`. It's explicit, doesn't throw exceptions, and avoids the weirdness of `Cloneable`.


---

**Q: Static Block vs Instance Block**
> "**static { ... }**:
> *   Runs **once** when the Class is loaded by ClassLoader.
> *   Used to initialize static variables.
>
> **{ ... }** (Instance Block):
> *   Runs **every time** you create a new Object (`new`).
> *   Runs *before* the constructor.
> *   Rarely used; usually we just put this code in the Constructor."

**Indepth:**
> **Bytecode**: Static blocks are compiled into a special method called `<clinit>` (Class Init). Instance blocks are copied into every `<init>` (Constructor) method.


---

**Q: Comparator vs Comparable**
> "**Comparable** (Internal): The class defines its own sort order (`compareTo`). 'I compare myself to others'.
>
> **Comparator** (External): A separate judge class (`compare`). 'I compare these two objects'.
>
> Use `Comparable` for default sorting (ID, Name). Use `Comparator` for ad-hoc sorting (Sort by Salary, Sort by Age)."

**Indepth:**
> **Contract**: `compareTo` returns negative (less), zero (equal), or positive (greater). A common trick `return this.val - other.val` is dangerous because of integer overflow! Use `Integer.compare(this.val, other.val)` instead.


## From 25 Arrays And Strings Revision
# 25. Data Structures Revision (Arrays & Strings)

**Q: Arrays.copyOf() vs System.arraycopy()**
> "**Arrays.copyOf()** is the easy, high-level way.
> You pass the original array and the new length, and it returns a **new** array. Use it when you just want a copy or need to resize.
>
> **System.arraycopy()** is the low-level, high-performance way.
> You must create the target array **first**. It copies data from source to destination at a specific index. It’s a native method (written in C/C++), making it extremely fast. Use it for complex shifting or merging."

**Indepth:**
> **Under the Hood**: `Arrays.copyOf()` checks the new length. If it's larger, it pads with default values (null/0). If smaller, it truncates. It then calls `System.arraycopy` internally.


---

**Q: Shallow Copy vs Deep Copy (Arrays)**
> "If you have an `int[]` array, a copy is just numbers. Easy.
>
> But if you have an `Employee[]` array:
> *   **Shallow Copy** (Default): The new array points to the **same** Employee objects. If you change Employee #1's name in the copy, the original array sees the change.
> *   **Deep Copy**: You must iterate through the array and manually create `new Employee()` for every single element. It’s the only way to have truly independent data."

**Indepth:**
> **Serialization**: One way to achieve a deep copy without writing manual code is to Serialize the object to a Byte Stream and immediately Deserialize it. It's slower but foolproof for complex graphs of objects.


---

**Q: Arrays.asList() Caveats**
> "Stop treating `Arrays.asList()` like a normal ArrayList.
> It creates a wrapper backed by the **original array**.
>
> 1.  **You cannot add/remove**: It throws `UnsupportedOperationException` because the underlying array size is fixed.
> 2.  **Changes reflect**: If you change an element in the list (`list.set(0, "X")`), the original array **also changes**.
>
> Always wrap it: `new ArrayList<>(Arrays.asList(...))` to be safe."

**Indepth:**
> **View**: Think of `Arrays.asList` as a "View" or "Window" onto the array. It doesn't own the data; the array does. Any change to the view passes through to the backing array.


---

**Q: Arrays.equals() vs ==**
> "Never use `==` on arrays unless you are checking if they are the exact same object in memory.
>
> `int[] a = {1, 2}; int[] b = {1, 2};`
> *   `a == b` is **false**.
> *   `a.equals(b)` is **false** (Arrays don't override `equals` from Object class!).
>
> **The Solution**: Use `Arrays.equals(a, b)` (for 1D arrays) or `Arrays.deepEquals(a, b)` (for 2D arrays)."

**Indepth:**
> **Multidimensional**: `Arrays.equals` only checks the first layer. If you have an array of arrays `int[][]`, the "elements" are array objects. `equals` checks if those array object references are the same (which they aren't). `deepEquals` recursively checks the contents.


---

**Q: String vs StringBuilder vs StringBuffer**
> "**String** is Immutable. Every time you say `str + "A"`, you create a brand new object. Slow for loops.
>
> **StringBuilder** is Mutable. It modifies the existing buffer. It is **Not Thread Safe**, but very fast. Use this 99% of the time.
>
> **StringBuffer** is the legacy version. It is **Synchronized** (Thread Safe), which makes it slower. Only use it if multiple threads are editing the string simultaneously (rare)."

**Indepth:**
> **Capacity**: `StringBuilder` has a capacity. When you append past the limit, it has to resize (usually doubles) which involves copying the old array to a new one. Setting an initial capacity close to expected size improves performance.


---

**Q: String Pool & Immutability**
> "Java saves memory by keeping a 'Pool' of unique strings.
> If you write `String s1 = "Hello"` and `String s2 = "Hello"`, they point to the exact same spot in memory.
>
> **Immutability** is crucial for security and safety.
> If Strings were mutable, I could pass a Database Connection URL to a function, and that function could change the URL to 'MaliciousSite.com', affecting everyone else using that string. Immutability prevents this."

**Indepth:**
> **GC**: String Deduplication (G1GC feature) allows the Garbage Collector to inspect the heap for duplicate strings. If it finds two identical string objects, it makes them share the same underlying `char[]` array to save RAM.


---

**Q: substring() Memory Leak (Historical)**
> "In older Java versions (pre-Java 7), `substring()` didn't create a new string. It just pointed to the original massive string array with a different start/end offset.
> This meant if you loaded a 10MB text file and blindly took a 5-byte substring `str.substring(0, 5)`, Java kept the **entire 10MB** in memory just for those 5 bytes.
>
> Modern Java (JDK 7u6+) fixed this: `substring()` now creates a fresh array copy. But it's good to know the history."

**Indepth:**
> **Offset**: The `String` class used to have `offset` and `count` fields to share `char[]`. This was removed to prevent memory leaks where a tiny substring prevents a massive collection from being GC'd.


---

**Q: equals() vs equalsIgnoreCase()**
> "Simple but standard.
> *   `"A".equals("a")` is **false**.
> *   `"A".equalsIgnoreCase("a")` is **true**.
>
> **Tip**: Always put the constant on the left to avoid NullPointers.
> Don't write `userInput.equals("ADMIN")`.
> Write `"ADMIN".equals(userInput)`. This works even if `userInput` is null."

**Indepth:**
> **Null Safety**: `Objects.equals(a, b)` is the safest way. It handles null checks for both `a` and `b` automatically. `Objects.equals(null, "A")` returns false, not a crash.


---

**Q: split() vs StringTokenizer**
> "**StringTokenizer** is legacy. It was there before Regex existed. Do not use it.
>
> use `str.split(regex)`.
> *   `"a,b,,c".split(",")` gives `["a", "b", "", "c"]`.
> *   **Watch out**: Trailing empty strings are discarded by default. Use `split(regex, -1)` to keep them."

**Indepth:**
> **Performance**: `StringTokenizer` is faster than `split` for simple delimiters because it doesn't use Regex. However, it's considered deprecated for new code due to its limited API and confusing behavior with empty tokens. Use `split` or `Guava Splitter`.


---

**Q: First Non-Repeating Character**
> "The logic:
> 1.  Loop once to populate a Frequency Map (`LinkedHashMap` preserves order, or just `int[256]` for ASCII).
> 2.  Loop through the **string** again.
> 3.  The first character with count 1 is the winner."

**Indepth:**
> **Optimization**: If the string only contains standard ASCII (0-127), a `boolean[128]` array is enough. If Extended ASCII, `boolean[256]`. If Unicode, you effectively need a `HashMap` or a sparse array.


## From 26 Data Structures Intermediate Revision
# 26. Data Structures (Intermediate Revision)

**Q: Regex Validation (matches)**
> "Stop writing complex loops to validate emails. Use Regex.
> `str.matches(\"\\\\d+\")` returns true if the string is all digits.
>
> **Performance Tip**: `String.matches()` re-compiles the pattern every time. If you validate thousands of strings, compile the `Pattern` object once as a `static final` constant and reuse it."

**Indepth:**
> **DOS Attack**: Be careful with Regex. A poorly written regex (nested quantifiers like `(a+)+`) can cause "Catastrophic Backtracking" if the input is malicious, hanging your CPU. This is a common Denial of Service vector.


---

**Q: ArrayList vs LinkedList (Real World)**
> "In textbooks, LinkedList is faster for adding/removing. In the real world, **ArrayList is almost always faster**.
>
> Why? **Cache Locality**.
> ArrayList elements are next to each other in memory. The CPU fetches a chunk of memory and processes it efficiently.
> LinkedList nodes are scattered everywhere. The CPU wastes time waiting for memory fetches (Cache Misses).
> Only use LinkedList if you are building a Queue/Deque or doing heavy splits/merges."

**Indepth:**
> **Memory Overhead**: LinkedList nodes have significant overhead. Each node stores the data object reference + next pointer + previous pointer. That's 24 bytes of overhead per element (on 64-bit JVM) compared to 0 bytes for ArrayList.


---

**Q: List.of() vs Arrays.asList()**
> "**Arrays.asList()** is a bridge to legacy arrays. It allows `set()` but not `add()`, and it passes changes through to the original array.
>
> **List.of()** (Java 9) is the modern standard for constants.
> *   It is **Truly Immutable** (No `set`, no `add`).
> *   It creates a highly optimized internal class (not a standard ArrayList).
> *   It does **not** allow nulls."

**Indepth:**
> **Best Practice**: Use `List.of` generally, but be aware of the "No Nulls" rule. If you need to store nulls (rare), you must use `Arrays.asList` or a standard `ArrayList`.


---

**Q: HashSet vs TreeSet vs LinkedHashSet**
> "You choose based on **Ordering**:
>
> 1.  **HashSet**: 'I don't care about order, just give me speed.' (O(1)).
> 2.  **LinkedHashSet**: 'I want them in the order I inserted them.' (Preserves insertion order).
> 3.  **TreeSet**: 'I want them sorted (A-Z, 1-10).' (O(log n))."

**Indepth:**
> **Consistency**: `TreeSet` uses `compareTo` (or Comparator) to determine equality, *not* `equals()`. If `compareTo` returns 0, TreeSet considers the elements duplicates, even if `equals()` returns false. This can lead to weird bugs.


---

**Q: HashMap vs TreeMap**
> "Same logic as Sets.
> *   **HashMap** for speed (O(1) lookup). Keys are jumbled.
> *   **TreeMap** for sorted keys (O(log n) lookup).
>
> Use `TreeMap` when you need features like `firstKey()`, `lastKey()`, or `subMap()`. For example, getting all events that happened 'between 10 AM and 11 AM'."

**Indepth:**
> **Balancing**: TreeMap uses a Red-Black Tree. This guarantees `log(n)` time for operations, but re-balancing the tree after insertions is more expensive than simply dumping an item into a HashMap bucket.


---

**Q: computeIfAbsent**
> "This method saves you lines of code and enhances performance.
>
> Instead of:
> `if (!map.containsKey(key)) map.put(key, new ArrayList()); return map.get(key);`
>
> You write:
> `return map.computeIfAbsent(key, k -> new ArrayList());`
>
> It's atomic, clean, and ensures the value is created only when needed."

**Indepth:**
> **Concurrent**: `computeIfAbsent` is atomic in `ConcurrentHashMap`. This makes it the perfect tool for implementing local caches without complex `synchronized` blocks.


---

**Q: HashMap Collision Handling**
> "If two keys hash to the same bucket, HashMap starts a **Linked List** in that bucket.
>
> If that list gets too long (more than 8 items), Java 8 automatically transforms it into a **Red-Black Tree**.
> This improves the worst-case performance from O(n) (scanning a long list) to O(log n) (searching a tree)."

**Indepth:**
> **Attack**: This switch to Trees (JEP 180) was done to prevent "Hash Flooding" attacks. Attackers could send thousands of requests with keys that all hash to the same bucket, turning your server into a slow O(n) crawler. The tree protects against this.


---

**Q: Queue vs Deque**
> "**Queue** (First-In-First-Out): Use it for tasks like 'Processing jobs in order'. Methods: `offer()`, `poll()`.
>
> "**Deque** (Double-Ended Queue): You can add/remove from **both** ends.
> Use `ArrayDeque` as your go-to Stack implementation. It is faster than the legacy `Stack` class because it is not synchronized."

**Indepth:**
> **Stack vs Deque**: `Stack` is a class, `Deque` is an interface. `Stack` extends `Vector`, which means every method is synchronized (slow). `ArrayDeque` is the modern, unsynchronized, faster replacement.


---

**Q: PriorityQueue**
> "This isn't a normal queue. It doesn't follow FIFO.
> It keeps elements ordered by **Priority** (Smallest to Largest by default).
>
> Useful for scheduling systems: 'Process the High Priority VIP job before the Regular job, even if the Regular job came first'."

**Indepth:**
> **Implementation**: PriorityQueue uses an array-based **Binary Heap**. It does *not* keep the array sorted. It only guarantees that `array[k] <= array[2*k+1]` and `array[k] <= array[2*k+2]`. This partial ordering is enough for fast polling.


---

**Q: Tree Traversal (Pre, In, Post)**
> "If this is a coding question, remember the position of the **Root**:
>
> 1.  **Pre-Order**: **Root**, Left, Right. (Used for copying trees).
> 2.  **In-Order**: Left, **Root**, Right. (Gives sorted order for BSTs).
> 3.  **Post-Order**: Left, Right, **Root**. (Used for deleting trees/garbage collection)."

**Indepth:**
> **Iterative**: Any recursive traversal can be converted to an iterative one using an explicit `Stack`. In interviews, knowing the iterative approach (especially for Pre-Order) is a big plus.


---

**Q: Streams: map vs flatMap**
> "**map** is for 1-to-1 conversion.
> `Stream<String> -> map(s -> s.length()) -> Stream<Integer>`.
>
> **flatMap** is for flattening nested structures.
> `Stream<List<String>> -> flatMap(List::stream) -> Stream<String>`.
> It 'unwraps' the inner lists so you have one big smooth stream of elements."

**Indepth:**
> **Stream Handling**: `flatMap` is essential when working with `Optional` inside a stream. `stream.map(Optional::stream)` gives you `Stream<Stream<T>>`. `stream.flatMap(Optional::stream)` gives you a clean `Stream<T>`.


---

**Q: Parallel Stream**
> "Don't just add `.parallel()` thinking it makes everything faster.
>
> *   **Good for**: Huge datasets, CPU-intensive tasks (like complex math on every element).
> *   **Bad for**: Small lists, IO operations (network/DB calls), or tasks that need order preservation.
>
> Overhead of splitting tasks and joining threads can often make Parallel Streams *slower* than simple Sequential Streams."

**Indepth:**
> **Spliterator**: Parallel streams effectively rely on the `Spliterator`. If your data source splits poorly (like a LinkedList), parallel performance will be terrible. `ArrayList` splits perfectly.

