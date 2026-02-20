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

