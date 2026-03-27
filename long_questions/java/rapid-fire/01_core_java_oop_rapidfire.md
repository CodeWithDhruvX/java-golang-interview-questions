# ☕ Core Java — OOP & Fundamentals (Rapid-Fire)

> 🔑 **Master Keyword:** **"EPIC-CC"** → Encapsulation, Polymorphism, Inheritance, Class-loading, Constructors, Copy

---

## 🏛️ Section 1: OOP Pillars

### Q1: What are the 4 Pillars of OOP?
🔑 **Keyword: "PEIA"** → Polymorphism, Encapsulation, Inheritance, Abstraction

| Pillar | One-Line | Real Example |
|---|---|---|
| **Encapsulation** | Bundle data + hide it | `Car` hides engine; you call `accelerate()` |
| **Abstraction** | Show essential, hide complexity | TV remote — buttons yes, circuit no |
| **Inheritance** | Child acquires parent properties | `Dog extends Animal` → dog gets `eat()` |
| **Polymorphism** | One name, multiple forms | `speak()` → Bark (Dog), Meow (Cat) |

---

### Q2: Abstract Class vs Interface?
🔑 **Keyword: "SCIV"** → State/Constructor/Is-A vs Capability

| Feature | Abstract Class | Interface |
|---|---|---|
| Methods | Abstract + Concrete | Abstract + `default`/`static` (Java 8+) |
| Variables | Any | `public static final` only |
| Constructor | ✅ YES | ❌ NO |
| Inheritance | Single `extends` | Multiple `implements` |
| Use-case | "Is-A" (Dog is Animal) | "Can-Do" (Dog implements Runnable) |

```java
abstract class Animal { abstract void speak(); void eat() { } }
interface Swimable { void swim(); default void float_() { } }
class Duck extends Animal implements Swimable { ... }
```

---

### Q3: Compile-time vs Runtime Polymorphism?
🔑 **Keyword: "OL-OR"** → Overload=Load-time, Override=Runtime

- **Compile-time (Static) :** Method **Overloading** — same name, different args, resolved at compile time
- **Runtime (Dynamic) :** Method **Overriding** — JVM resolves at runtime based on actual object type

```java
// Overloading (Compile-time)
int add(int a, int b) { return a+b; }
double add(double a, double b) { return a+b; }

// Overriding (Runtime)
Animal a = new Dog();
a.sound(); // Dog's sound is called — not Animal's
```

---

### Q4: Can you override `static` or `private` methods?
🔑 **Keyword: "SH-PV"** → Static=Hides, Private=Vanishes

- **Static methods** → ❌ Cannot override. In child → called **Method Hiding**
- **Private methods** → ❌ Not visible to child, cannot be overridden

---

### Q5: Composition vs Inheritance — Which is better?
🔑 **Keyword: "HALO"** → Has-A = Loose, Is-A = tightly cOupled

- **Inheritance (Is-A):** `Car extends Vehicle` — tightly coupled, Fragile Base Class Problem
- **Composition (Has-A):** `Car has-a Engine` — loosely coupled, change behavior at runtime
- **Verdict:** Always **Prefer Composition** over Inheritance

---

### Q6: What is `super` keyword?
🔑 **Keyword: "PVC"** → Parent Variable/Constructor/method Call

```java
super.variable    // access parent variable
super()           // call parent constructor (must be 1st line)
super.method()    // call parent method (useful when overridden)
```

---

### Q7: What is `this` keyword?
🔑 **Keyword: "SPC"** → Self/Pass/Chaining

```java
this.x = x;      // distinguish field from param
this()            // call another constructor in same class (Constructor Chaining)
method(this)      // pass current object as argument
```

---

### Q8: Overloading vs Overriding?

| Feature | Overloading | Overriding |
|---|---|---|
| Scope | Same class | Parent → Child |
| Signature | Must change | Must be same |
| Return type | Can differ | Same or covariant |
| Binding | Static (compile) | Dynamic (runtime) |

---

## 🔨 Section 2: Constructors & Object Lifecycle

### Q9: Can a constructor be `private`?
🔑 **Keyword: "SF"** → Singleton + Factory

Yes! Used in **Singleton** (prevent external instantiation) and **Factory** patterns (force static method usage).

---

### Q10: Constructor Chaining?
🔑 **Keyword: "TFSP"** → This/First/Super/Parent

- Within same class: `this(args)` — must be 1st statement
- From child to parent: `super(args)` — must be 1st statement
- **Cannot have both `this()` and `super()` in same constructor**

```java
class Employee {
    Employee() { this("Unknown"); }            // chains to next constructor
    Employee(String name) { this.name = name; }
}
```

---

### Q11: Instance Block vs Static Block?
🔑 **Keyword: "IO-SC"** → Instance=Object, Static=ClassLoad

```java
static { /* Runs ONCE when class is loaded */ }
{ /* Runs BEFORE constructor, EVERY object creation */ }
```

---

### Q12: Immutable Class — How to create?
🔑 **Keyword: "FNSD"** → Final-class, No-setters, private-Final-fields, Defensive-copy

1. Declare class `final`
2. All fields `private final`
3. No setters
4. Initialize all in constructor
5. If mutable fields (List/Date) → return **deep copy** in getter

```java
final class Money {
    private final int amount;
    Money(int amount) { this.amount = amount; }
    int getAmount() { return amount; }  // safe
    List<String> getTags() { return new ArrayList<>(tags); } // defensive copy
}
```

---

### Q13: Shallow Copy vs Deep Copy?
🔑 **Keyword: "SRS-DNC"** → Shallow=Reference-Shared, Deep=New-Copy

- **Shallow:** Copies field values. Reference fields **shared** → both objects point to same data
- **Deep:** Creates **new objects** for referenced fields
- Ways to deep copy: Override `clone()`, Serialization/Deserialization, Copy Constructor

---

## 🔧 Section 3: Core Keywords & Types

### Q14: `static` keyword?
🔑 **Keyword: "SVBC"** → Static=Variable/Block/Class belongs to class not instance

| Usage | Meaning |
|---|---|
| `static` variable | Shared across all objects |
| `static` method | No `this`; can't access non-static |
| `static` block | Runs once at class load |
| `static` inner class | No outer instance needed |

---

### Q15: `volatile` keyword?
🔑 **Keyword: "VMCD"** → Visibility+Memory+CPU-cache-bypass, not atomicity

- **Visibility Guarantee:** Reads/Writes → directly to Main Memory (RAM), bypasses CPU cache
- **Prevents reordering** of instructions by JVM
- ❌ Does **NOT** guarantee Atomicity → use `AtomicInteger` for that

---

### Q16: `==` vs `equals()`?
🔑 **Keyword: "RA-LC"** → Reference-Address vs Logical-Content

- `==` → Reference comparison (same memory address?)
- `equals()` → Logical comparison (same content?) — String/Integer/List override this

---

### Q17: Wrapper Classes & Autoboxing?
🔑 **Keyword: "WAN"** → Wrapper-Autobox-NPE

- **Wrapper:** Object version of primitive (`Integer` wraps `int`). Needed for Collections
- **Autoboxing:** `int` → `Integer` (auto)
- **Unboxing:** `Integer` → `int` (auto)
- ⚠️ Unboxing `null` → **NullPointerException**

---

### Q18: Integer Cache (-128 to 127)?
🔑 **Keyword: "128-Cache"**

```java
Integer a = 100, b = 100;
System.out.println(a == b); // TRUE — cached same object

Integer x = 200, y = 200;
System.out.println(x == y); // FALSE — new objects
```

---

### Q19: `BigDecimal` — Why use it?
🔑 **Keyword: "FPE-Money"** → Float-Point-Error, use for Money

```java
// BAD:
0.1 + 0.2 = 0.30000000000000004 (floating point error)

// GOOD:
new BigDecimal("0.1").add(new BigDecimal("0.2")) = 0.3
```
Always use `BigDecimal` for **financial calculations**.

---

## 🎯 Section 4: Generics

### Q20: What is Type Erasure?
🔑 **Keyword: "CER"** → Compile→Erase→Runtime-raw

- Generics = **compile-time safety only**
- Compiler: checks types → inserts casts → **erases** to `Object`
- At runtime: `List<String>` and `List<Integer>` are just `List`
- Consequence: Can't do `instanceof List<String>` or `new T()`

---

### Q21: Wildcards — PECS Rule?
🔑 **Keyword: "PECS"** → Producer-Extends, Consumer-Super

| Wildcard | Can Read? | Can Write? | Use-case |
|---|---|---|---|
| `?` | Object only | ❌ | Read-all |
| `? extends Number` | ✅ Number | ❌ | Producer (read only) |
| `? super Integer` | Object only | ✅ Integer | Consumer (write) |

---

## 🔍 Section 5: Reflection & Annotations

### Q22: What is Reflection?
🔑 **Keyword: "RISC"** → Runtime-Inspect-Self-Code

- Inspect/modify classes, fields, methods at **runtime**
- Uses: IDE, Frameworks (Spring DI, JUnit), Debuggers
- Cons: Slow, breaks encapsulation, security holes

```java
Field f = obj.getClass().getDeclaredField("secret");
f.setAccessible(true); // magic key
Object val = f.get(obj);
```

---

### Q23: Access Modifiers?
🔑 **Keyword: "DPPP-scope"** → from narrowest to widest

| Modifier | Class | Package | Subclass | World |
|---|---|---|---|---|
| `private` | ✅ | ❌ | ❌ | ❌ |
| (default) | ✅ | ✅ | ❌ | ❌ |
| `protected` | ✅ | ✅ | ✅ | ❌ |
| `public` | ✅ | ✅ | ✅ | ✅ |

> ⚠️ `protected` visible via inheritance, NOT via reference of parent type from outside package

---

## 🏷️ Section 6: Enums, Inner Classes & Modern Java

### Q24: Java Enums?
🔑 **Keyword: "ECF-Singleton"** → Enum=Class+Fields+Constructor, best Singleton

```java
enum Status {
    ACTIVE(1), INACTIVE(0);
    int code;
    Status(int code) { this.code = code; } // private by default
}
Status s = Status.ACTIVE;
s.code; // 1
```
- Enums can have **methods, constructors, fields**
- Enum Singleton = best (safe from Reflection + Serialization attacks)

---

### Q25: Inner Class Types?
🔑 **Keyword: "MSLA"** → Member, Static, Local, Anonymous

1. **Member Inner Class** → non-static, needs outer instance
2. **Static Nested Class** → no outer instance, `new Outer.Inner()`
3. **Local Inner Class** → inside method, accesses final local vars
4. **Anonymous Inner Class** → callback/listener, no name: `new Runnable() { ... }`

---

### Q26: Java Records (Java 14+)?
🔑 **Keyword: "RIET"** → Record=Immutable+Equals+ToString auto

```java
record Point(int x, int y) {}
// Auto-generates: constructor, equals, hashCode, toString, getters (x(), y())
```
All fields are `final` → immutable by default.

---

### Q27: Sealed Classes (Java 17+)?
🔑 **Keyword: "SPP"** → Sealed-Permits-Pattern

```java
public sealed class Shape permits Circle, Square {}
public final class Circle extends Shape {}
```
Restricts which classes can extend the sealed class.

---

### Q28: `var` keyword (Java 10+)?
🔑 **Keyword: "LVTI-Local"** → Local-Variable-Type-Inference only

```java
var list = new ArrayList<String>(); // compiler infers ArrayList<String>
// Only for local variables. Must be initialized. Cannot use for fields/params.
```

---

## 📅 Section 7: Date/Time API

### Q29: Legacy vs Modern Date/Time?
🔑 **Keyword: "LMI"** → Legacy=Mutable/Jan=0, Modern=Immutable/Thread-safe

| Legacy (Bad) | Modern Java 8+ (Good) |
|---|---|
| `java.util.Date`, `Calendar` | `java.time.*` |
| Mutable (not thread-safe) | **Immutable** |
| Jan = 0, Dec = 11 | Jan = 1 |

```java
LocalDate birthday = LocalDate.of(1995, 3, 27);
LocalDateTime now = LocalDateTime.now();
ZonedDateTime zdt = ZonedDateTime.now(ZoneId.of("Asia/Kolkata"));
Instant timestamp = Instant.now(); // UTC epoch-based
```

---

## 🗑️ Section 8: Garbage Collection & References

### Q30: Reference Types?
🔑 **Keyword: "SSWP"** → Strong-Soft-Weak-Phantom

| Type | GC Behavior | Use-case |
|---|---|---|
| **Strong** | Never collected if reachable | Normal `Object obj = new Object()` |
| **Soft** | Collected only when memory full | In-memory Cache |
| **Weak** | Collected on next GC | WeakHashMap, metadata |
| **Phantom** | Post-finalization cleanup | Advanced memory management |

---

### Q31: `final`, `finally`, `finalize` — Difference?
🔑 **Keyword: "KBC"** → Keyword/Block/Callback

- `final` → **Keyword**: var=constant, method=no-override, class=no-extend
- `finally` → **Block**: always runs after try-catch (even on exception)
- `finalize` → **Method**: GC calls before destroying object — **Deprecated Java 9+**

---

*End of File — Core Java OOP & Fundamentals*
