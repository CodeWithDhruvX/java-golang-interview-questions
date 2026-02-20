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

