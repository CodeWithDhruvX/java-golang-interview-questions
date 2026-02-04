# ⚠️ Java Exceptions & IO Basics

## 1️⃣ Exception Handling

### Question 21: Exception Hierarchy in Java.

**Answer:**
*   **Throwable** (Parent)
    *   **Error:** Serious JVM problems (StackOverflowError, OutOfMemoryError). Cannot be caught/recovered.
    *   **Exception:**
        *   **Checked (Compile-time):** Must be handled (`try-catch` or `throws`). Ex: `IOException`, `SQLException`.
        *   **Unchecked (Runtime):** Programming bugs. Ex: `NullPointerException`, `IndexOutOfBoundsException`.

---

### Question 22: Validating `try-catch`, `finally` combinations.

**Answer:**
*   `try` alone? -> ❌ Compilation Error.
*   `try - catch` -> ✅
*   `try - finally` -> ✅
*   `try - catch - finally` -> ✅
*   **finally:** Always executes (even if exception occurs), unless `System.exit()` is called.

---

### Question 23: `throw` vs `throws`.

**Answer:**
*   **`throw`:** Used **inside** a method to explicitly throw an exception.
    *   `throw new RuntimeException("Error");`
*   **`throws`:** Used in **method signature** to declare potential exceptions.
    *   `public void readFile() throws IOException { ... }`

---

### Question 24: `final`, `finally`, `finalize`.

**Answer:**
*   **`final`:** Keyword. Variable (const), Method (no override), Class (no inheritance).
*   **`finally`:** Block. Cleanup code after try-catch.
*   **`finalize`:** Method. Called by GC before destroying object. **Deprecated** in Java 9. Use `AutoCloseable` instead.

---

### Question 25: What is Try-with-Resources? (Java 7+)

**Answer:**
Automatic Resource Management.
*   Any object implementing `AutoCloseable` can be used.
*   Automatically calls `.close()` at end of block.
```java
try (BufferedReader br = new BufferedReader(new FileReader("path"))) {
    // read file
} catch (IOException e) { ... } // br.close() called automatically
```

---

### Question 26: Checked vs Unchecked Exception? (When to use?)

**Answer:**
*   **Checked:** Use for recoverable conditions outside program control (File missing, DB down). Forces caller to handle.
*   **Unchecked:** Use for programming errors (Null pointer, Bad argument). Caller cannot recover.

---

### Question 27: Custom Exception creation.

**Answer:**
Extend `Exception` (for Checked) or `RuntimeException` (for Unchecked).
```java
public class InsufficientFundsException extends RuntimeException {
    public InsufficientFundsException(String msg) {
        super(msg);
    }
}
```

---

### Question 28: What happen if you throw exception in `finally` block?

**Answer:**
It consumes (masks) the original exception thrown in the variable `try` block. The caller sees only the exception from `finally`.
*   *Best Practice:* Avoid throwing in `finally`.

---

### Question 29: Exception Propagation.

**Answer:**
*   **Unchecked:** Automatically propagates up the call stack until caught or thread dies.
*   **Checked:** Must be explicitly propagated using `throws` at every level.

---

## 2️⃣ Java IO & Serialization

### Question 30: What is Serialization?

**Answer:**
Converting an Object state into a byte stream (to save to file or send over network).
*   Implement `java.io.Serializable` (Marker interface).
*   Use `ObjectOutputStream.writeObject()`.

---

### Question 31: `serialVersionUID` significance.

**Answer:**
Unique ID for the class version.
*   Serializer writes ID. Deserializer checks if local class ID matches.
*   If mismatched (code changed), throws `InvalidClassException`.
*   *Tip:* Always define it manually to avoid issues during partial code updates.

---

### Question 32: `transient` keyword.

**Answer:**
Used in Serialization.
*   Fields marked `transient` are **skipped** during serialization.
*   When deserialized, they get default values (null/0).
*   *UseCase:* Passwords, sensitive data.

---

### Question 33: Externalizable interface vs Serializable.

**Answer:**
*   **Serializable:** Magic. JVM serializes everything automatically.
*   **Externalizable:** You control logic. You MUST implement `writeExternal()` and `readExternal()`. Faster performance because you pick what to save.

---

### Question 34: Byte Stream vs Character Stream.

**Answer:**
*   **Byte Stream (`InputStream`, `OutputStream`):** Raw bytes. Images, Audio, Video.
*   **Character Stream (`Reader`, `Writer`):** Text. Handle Unicode/Encoding automatically.
*   *Bridge:* `InputStreamReader` converts Bytes -> Char.

---

### Question 35: `Scanner` vs `BufferedReader`.

**Answer:**
*   **Scanner:** Parses primitives (`nextInt`, `nextBoolean`). Slower (Regex parsing). Small buffer.
*   **BufferedReader:** Reads lines (`readLine`). Faster. Large buffer (8KB).

---

### Question 36: Handling `FileNotFoundException` vs `IOException`.

**Answer:**
`FileNotFoundException` is a subclass of `IOException`.
Catch specific (FileNotFound) first, then general (IO) if needed.

---

### Question 37: `File` vs `Path` (NIO.2).

**Answer:**
*   **`File` (Legacy):** `java.io.File`. Blocking IO.
*   **`Path` (Modern):** `java.nio.file.Path`. Part of NIO (Non-blocking). Better error handling, works with `Files` utility class (`Files.readAllLines()`).

---

### Question 38: Breaking Singleton with Serialization.

**Answer:**
If you serialize a Singleton and deserialize it, you get a **new instance**.
*   **Fix:** Implement `readResolve()` method in Singleton.
```java
protected Object readResolve() {
    return INSTANCE;
}
```

---

### Question 39: `System.out`, `System.err`, `System.in`.

**Answer:**
*   `out`: Standard Output (PrintStream).
*   `err`: Standard Error (PrintStream). Often red in IDEs.
*   `in`: Standard Input (InputStream).

---

### Question 40: Closeable vs AutoCloseable.

**Answer:**
*   **AutoCloseable (Java 7):** Parent interface. used in try-with-resources. throws `Exception`.
*   **Closeable (Legacy):** Extends AutoCloseable. throws `IOException`.
