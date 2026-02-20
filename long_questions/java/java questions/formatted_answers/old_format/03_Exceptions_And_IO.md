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

