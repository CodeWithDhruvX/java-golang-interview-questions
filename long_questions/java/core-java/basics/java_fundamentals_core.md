# ⚙️ Java Core Fundamentals

## 1️⃣ Keywords & Object Methods

### Question 41: Explain `static` keyword in Java.

**Answer:**
Belongs to the class, not the instance.
*   **Variable:** Shared copy for all objects (Global state).
*   **Method:** Can be called without creating object (`Math.pow`). Cannot access `this` or non-static variables.
*   **Block:** Runs once at class loading.
*   **Class:** Only **Nested** classes can be static.

---

### Question 42: What does `volatile` do?

**Answer:**
Thread-safety keyword for variables.
1.  **Visibility Guarantee:** Reads/Writes go directly to Main Memory (RAM), bypassing thread CPU cache.
2.  **Instruction Reordering:** Prevents JVM from reordering instructions involving the variable.
*   *Note:* It does NOT guarantee Atomicity (use `AtomicInteger` for that).

---

### Question 43: Comparing Objects: `==` vs `equals()`.

**Answer:**
*   **`==`**: Reference comparison. Checks if both point to the **same memory address**.
*   **`equals()`**: Logical comparison. Default implementation uses `==`. Classes (String, Integer, List) override it to compare **content**.

---

### Question 44: Common `Object` methods (`toString`, `hashCode`, `equals`).

**Answer:**
Methods present in every class (from `java.lang.Object`):
1.  **`toString()`**: String representation. Default: `ClassName@HexHash`. Always override for logging.
2.  **`equals(Object o)`**: For content comparison.
3.  **`hashCode()`**: Numeric summary used in HashMaps. If `equals` is true, `hashCode` MUST be same.
4.  **`getClass()`**: Returns runtime class metadata.
5.  **`clone()`**: Creates copy (Protected).

---

### Question 45: `finalize()` method - Why is it deprecated?

**Answer:**
*   **Purpose:** Called by GC before deleting object. Thought to be for cleanup.
*   **Problems:** Unpredictable (might never run), Slow, Security risks, Resurrection (can save object from death).
*   **Solution:** Removed/Deprecated since Java 9. Use `try-with-resources` or `Cleaner` API.

---

### Question 46: Wrapper Classes & Autoboxing.

**Answer:**
*   **Wrapper:** Object version of primitive (`Integer` wraps `int`). Used in Collections (`List<Integer>`).
*   **Autoboxing:** Auto-conversion `int` -> `Integer`.
*   **Unboxing:** Auto-conversion `Integer` -> `int`.
*   **Caveat:** `NullPointerException` if unboxing `null`.

---

### Question 47: Integer Cache (-128 to 127).

**Answer:**
Java caches Integer objects from -128 to 127.
```java
Integer a = 100, b = 100;
System.out.println(a == b); // TRUE (Same cached implementation)

Integer x = 200, y = 200;
System.out.println(x == y); // FALSE (New objects)
```

---

### Question 48: `BigInteger` and `BigDecimal`.

**Answer:**
*   **`BigInteger`:** Arbitrary-precision integers. (Crypto, Factorials).
*   **`BigDecimal`:** Arbitrary-precision decimals. **Essential for Money/Finance**. Avoids floating point errors (`0.1 + 0.2 != 0.300000004`).

---

## 2️⃣ Generics

### Question 49: What is Type Erasure?

**Answer:**
Java Generics meant for **Compile-time** safety only.
*   Compiler checks types -> Insert casts -> **Erases** types to `Object` (or bound).
*   Result: `List<String>` and `List<Integer>` become just `List` at runtime.
*   Consequence: You cannot use `instanceof List<String>` or `new T()`.

---

### Question 50: Wildcards in Generics (`?`, `extends`, `super`).

**Answer:**
*   **`List<?>`**: Unbounded. Read-only (conceptually).
*   **`? extends Number` (Upper Bound):** "Producer". Safe to **Read** (you get at least a Number). Cannot Add.
*   **`? super Integer` (Lower Bound):** "Consumer". Safe to **Write** (you can put Integer). Reading gives Object.
*   *PECS Rule:* Producer Extends, Consumer Super.

---

### Question 51: Generic Methods.

**Answer:**
Methods can have their own type parameters independent of the class.
```java
public <T> T print(T item) { ... }
```
Inferred by compiler from arguments.

---

## 3️⃣ Reflection & Annotations

### Question 52: What is Reflection? Pros/Cons.

**Answer:**
Ability of code to inspect/modify itself at runtime.
*   **Uses:** IDEs, Debuggers, Frameworks (Spring DI, JUnit).
*   **Capabilities:** List methods, Access private fields, Create objects dynamically.
*   **Cons:** Slower performance, Breaks Encapsulation, Security holes.

---

### Question 53: How to access a Private Field using Reflection?

**Answer:**
```java
Field f = obj.getClass().getDeclaredField("secret");
f.setAccessible(true); // The magic key
Object value = f.get(obj);
```

---

### Question 54: What is the `Class` class?

**Answer:**
Entry point for Reflection.
*   `String.class` (Compile time).
*   `obj.getClass()` (Runtime).
*   `Class.forName("java.lang.String")` (Dynamic loading).

---

### Question 55: Custom Annotations & Meta-Annotations.

**Answer:**
*   **Annotation:** Interface with `@interface`.
*   **Meta-Annotations:**
    *   `@Retention(RetentionPolicy.RUNTIME)`: Kept at runtime (for Reflection).
    *   `@Target(ElementType.METHOD)`: Where it can be used.
```java
@Retention(RetentionPolicy.RUNTIME)
@interface MyTest { }
```

---

### Question 56: Breaking Singleton using Reflection.

**Answer:**
Reflection can call `private constructor` of Singleton.
*   **Defense:** Throw exception in constructor if `instance` is not null.
*   **Best Defense:** Use **Enum Singleton** (Reflection cannot create Enum instances).

---

## 4️⃣ Access Modifiers & Package Scope

### Question 57: Private vs Default vs Protected vs Public.

**Answer:**
| Modifier | Class | Package | Subclass (diff pkg) | World |
| :--- | :--- | :--- | :--- | :--- |
| **private** | ✅ | ❌ | ❌ | ❌ |
| **(default)**| ✅ | ✅ | ❌ | ❌ |
| **protected**| ✅ | ✅ | ✅ (Inheritance) | ❌ |
| **public** | ✅ | ✅ | ✅ | ✅ |
*   **Note:** `protected` is tricky. Subclass can access it via *inheritance*, but not via *reference* of parent type.

---
