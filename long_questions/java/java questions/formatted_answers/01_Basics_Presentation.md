## 🏗️ 01. Java Core Basics

### **Q: Difference between JDK, JRE, JVM?**
*   **JVM:** Engine that runs bytecode. Platform-dependent.
*   **JRE:** JVM + Runtime Libraries. For **running** apps.
*   **JDK:** JRE + Dev Tools (javac). For **building** apps.
*   **Takeaway:** Java is 'Write Once, Run Anywhere' because of the JVM.

**🏢 Asked in:** All companies (TCS, Infosys, Wipro, Cognizant, Accenture, HCL, Tech Mahindra, Capgemini) - Fundamental concept

---

### **Q: Abstract Class vs Interface?**
*   **Abstract Class:** An "is-a" relationship. Can have state (fields) and constructors.
*   **Interface:** A "can-do" capability. No instance state. Multiple inheritance support.
*   **Decision:** Use Abstract Class for identity; Interface for behavior contracts.

**🏢 Asked in:** Service-based companies (TCS, Infosys, Wipro, Cognizant) + Product companies (Amazon, Flipkart, Microsoft) - OOP fundamentals

---

### **Q: String vs StringBuilder vs StringBuffer?**
*   **String:** Immutable. Thread-safe but slow for heavy edits.
*   **StringBuilder:** Mutable. Fast but **not** thread-safe.
*   **StringBuffer:** Mutable. Thread-safe but slow due to synchronization.
*   **Takeaway:** 99% of the time, use `StringBuilder` for local edits.

**🏢 Asked in:** All companies - Critical for performance optimization and memory management

---

### **Q: equals() vs ==?**
*   **==:** Reference equality (same memory address).
*   **equals():** Value equality (logical content check).
*   **Warning:** Always override `hashCode()` if you override `equals()`.

**🏢 Asked in:** All companies - Fundamental concept, often asked with HashMap implementation questions

---

### **Q: final vs finally vs finalize?**
*   **final:** Constant (var), Non-overridable (method), Non-inheritable (class).
*   **finally:** Block that **always** runs after try-catch (cleanup).
*   **finalize:** Deprecated. Unreliable GC cleanup. DON'T USE.

**🏢 Asked in:** Service-based companies (TCS, Infosys, Wipro, HCL) - Exception handling fundamentals

---

### **Q: Collections vs Arrays?**
*   **Arrays:** Fixed size. Fast. Can hold primitives and objects.
*   **Collections:** Dynamic size. Rich API. Only holds objects (uses Autoboxing for primitives).

**🏢 Asked in:** All companies - Data structure selection is crucial for performance

---

### **Q: List vs Set vs Map?**
*   **List:** Ordered, allows duplicates. (ArrayList).
*   **Set:** Unordered (usually), no duplicates. (HashSet).
*   **Map:** Key-Value pairs. Unique keys. (HashMap).

**🏢 Asked in:** All companies - Core collection framework, essential for daily coding

---

### **Q: ArrayList vs LinkedList?**
*   **ArrayList:** Array-based. O(1) random access. Fast for reading.
*   **LinkedList:** Node-based. O(1) add/remove at ends. Slow random access.
*   **Fact:** ArrayList is usually faster due to CPU cache efficiency.

**🏢 Asked in:** Product companies (Amazon, Microsoft, Google) + Service companies - Performance analysis

---

### **Q: HashMap vs TreeMap vs LinkedHashMap?**
*   **HashMap:** O(1) performance. No guaranteed order.
*   **LinkedHashMap:** Preserves **insertion order**.
*   **TreeMap:** Keys are **sorted**. O(log n) performance.

**🏢 Asked in:** Product companies (Amazon, Flipkart, Paytm) - Critical for caching and data processing

---

### **Q: map() vs flatMap()?**
*   **map():** 1-to-1 transformation. (Value -> Value).
*   **flatMap():** 1-to-many transformation. (Value -> Stream). Flattens nests.

**🏢 Asked in:** Product companies (Swiggy, Zomato, Razorpay) + Modern service companies - Java 8+ streams

---

### **Q: Thread vs Runnable?**
*   **Runnable:** An interface representing a **task**. Preferred (Composition).
*   **Thread:** A class representing a **worker**. (Inheritance).
*   **Tip:** Always prefer `Runnable` to keep your class available for other inheritance.

**🏢 Asked in:** Product companies (Amazon, Microsoft, Adobe) - Concurrency fundamentals

---

### **Q: @Component vs @Service vs @Repository?**
*   **@Component:** Generic Spring bean.
*   **@Service:** Business logic layer.
*   **@Repository:** Data access layer. Enables **Exception Translation** (SQL -> Spring Exceptions).

**🏢 Asked in:** All companies using Spring (TCS, Infosys, Cognizant, Accenture, HCL, Persistent) - Spring framework essentials

---

## 🧩 04. Java Fundamentals (Deep Dive)

### **Q: static keyword in Java?**
*   Belongs to the **class**, not the object.
*   Shared across all instances.
*   Static methods can be called without creating an object.

**🏢 Asked in:** All companies - Core OOP concept, frequently tested

---

### **Q: Wrapper Classes & Autoboxing?**
*   **Wrapper:** Object version of primitives (Integer, Boolean).
*   **Autoboxing:** Automatic conversion from `int` to `Integer`.
*   **Risk:** Autoboxing in tight loops causes high GC pressure.

**🏢 Asked in:** Product companies (Amazon, Microsoft, Google) + Service companies - Performance optimization

---

### **Q: Integer Cache (-128 to 127)?**
*   Java caches Integer objects in this range.
*   `Integer a = 127; Integer b = 127; (a == b)` is **true**.
*   `Integer a = 128; Integer b = 128; (a == b)` is **false**.

**🏢 Asked in:** Product companies (Google, Microsoft, Amazon) - Deep Java knowledge test

---

### **Q: BigInteger and BigDecimal?**
*   **BigInteger:** Unlimited size integers.
*   **BigDecimal:** Exact precision decimals. **Use for Money.**

**🏢 Asked in:** Financial companies (Paytm, PhonePe, Razorpay, Banks) + E-commerce (Amazon, Flipkart) - Critical for financial applications

---

### **Q: Type Erasure?**
*   Generics are removed after compilation for backward compatibility.
*   `List<String>` becomes `List` at runtime.

**🏢 Asked in:** Product companies (Microsoft, Google, Amazon) - Advanced Java concept

---

### **Q: Wildcards (?, extends, super)?**
*   **<?>**: Unbounded (anything).
*   **<? extends T>**: Upper bound (T or its children). **Producer (Read).**
*   **<? super T>**: Lower bound (T or its parents). **Consumer (Write).**
*   **Rule:** PECS (Producer Extends, Consumer Super).

**🏢 Asked in:** Product companies (Amazon, Microsoft, Google) - Generic programming expertise

---

### **Q: Reflection (Pros/Cons)?**
*   **Pros:** Dynamic inspecting/modifying of code. Basis of Spring/Hibernate.
*   **Cons:** Slower, breaks encapsulation, Sarder to debug.

**🏢 Asked in:** Product companies (Microsoft, Google, Oracle) + Framework development - Advanced Java

---

### **Q: Custom Annotations & Meta-Annotations?**
*   Created using `@interface`.
*   **@Target:** Where it can be used (Field, Method).
*   **@Retention:** How long it lasts (Source, Class, Runtime).

**🏢 Asked in:** Product companies (Spring framework users, Google, Microsoft) - Framework development expertise

---

## 🏛️ 06. OOP Basics

### **Q: The 4 Pillars of OOP?**
1.  **Encapsulation:** Protect state with access modifiers.
2.  **Inheritance:** "Is-A" relation for code reuse.
3.  **Polymorphism:** Method Overloading (Static) & Overriding (Dynamic).
4.  **Abstraction:** Focus on "what" an object does, not "how".

**🏢 Asked in:** All companies - Fundamental OOP concept, must-know for every Java developer

---

### **Q: Composition vs Inheritance?**
*   **Inheritance:** Tight coupling. Harder to change.
*   **Composition:** Looser coupling. Greater flexibility (Swap parts at runtime).
*   **Verdict:** Favor Composition.

**🏢 Asked in:** Product companies (Amazon, Microsoft, Google) + Service companies - Design pattern knowledge

---

### **Q: super vs this?**
*   **this:** Refers to **current** instance.
*   **super:** Refers to **immediate parent** instance.

**🏢 Asked in:** Service-based companies (TCS, Infosys, Wipro, HCL) - Basic OOP concept

---

### **Q: Overloading vs Overriding?**
*   **Overloading:** Same name, different args. Compile-time. Same class.
*   **Overriding:** Same signature. Runtime. Parent-Child relationship.

**🏢 Asked in:** All companies - Core polymorphism concept, frequently tested

---

### **Q: Shallow Copy vs Deep Copy?**
*   **Shallow:** Copies references. Both original and copy share nested objects.
*   **Deep:** Copies values and nested objects. Complely independent.

**🏢 Asked in:** Product companies (Amazon, Microsoft, Google) + Service companies - Object cloning and immutability

---

### **Q: Immutable Class - How?**
1.  Class is `final`.
2.  All fields are `private` and `final`.
3.  No setters.
4.  **Defensive Copying** for mutable field getters.

**🏢 Asked in:** Product companies (Amazon, Microsoft, Google) - Thread-safe design patterns

---

## 🛠️ 14 & 15. SOLID, Arrays & Strings

### **Q: SOLID - Open/Closed Principle (OCP)?**
*   "Open for extension, closed for modification."
*   Add new features by adding new classes, not modifying existing ones.

**🏢 Asked in:** Product companies (Amazon, Microsoft, Google) + Service companies - Design principles

---

### **Q: Comparable vs Comparator?**
*   **Comparable:** Natural ordering. Implemented inside the class (`compareTo`).
*   **Comparator:** Custom ordering. External class or Lambda (`compare`).

**🏢 Asked in:** All companies - Sorting and collections, frequently asked with TreeSet/TreeMap

---

### **Q: Arrays.asList() Caveats?**
*   Fixed-size wrapper around an array.
*   No add/remove allowed.
*   Changes affect the original array.

**🏢 Asked in:** Service-based companies (TCS, Infosys, Wipro, Cognizant) - Common pitfalls

---

### **Q: String Constant Pool & intern()?**
*   **Pool:** Memory area in Heap for literals.
*   **intern():** Manually moves a runtime string into the pool to save memory.

**🏢 Asked in:** Product companies (Amazon, Microsoft, Google) - Memory optimization

---

### **Q: String Immutability Importance?**
1.  **Security:** File paths/URLs can't be changed after validation.
2.  **Performance:** Enables String Pool.
3.  **Thread-safety:** Shared safely without locks.

**🏢 Asked in:** Product companies (Amazon, Microsoft, Google) + Service companies - Core Java concept

---

## 🍃 31. Spring Boot Basics

### **Q: Spring vs Spring Boot?**
*   Spring Boot = Spring + Auto-Configuration + Embedded Server + Starters.
*   It eliminates XML/boilerplate config.

**🏢 Asked in:** All companies using Spring (TCS, Infosys, Cognizant, Accenture, HCL, Persistent, Capgemini) - Framework fundamentals

---

### **Q: @SpringBootApplication?**
Combines `@Configuration`, `@EnableAutoConfiguration`, and `@ComponentScan`.

**🏢 Asked in:** All companies using Spring Boot - Core annotation understanding

---

### **Q: JAR vs WAR?**
*   **JAR:** Contains embedded Tomcat. "Just Run It" (`java -jar`). Modern standard.
*   **WAR:** Needs an external server. Legacy/Enterprise standard.

**🏢 Asked in:** Service-based companies (TCS, Infosys, Wipro, HCL) + Enterprise companies - Deployment knowledge

---

### **Q: @WebMvcTest vs @SpringBootTest?**
*   **@SpringBootTest:** Full app context. Slow.
*   **@WebMvcTest:** Controller layer only. Fast.

**🏢 Asked in:** Product companies (Amazon, Microsoft, Google) + Service companies - Testing best practices

---

### **Q: TestContainers?**
Starts **real Docker containers** (Postgres, Kafka) for integration tests. Ensures production parity.

**🏢 Asked in:** Product companies (Amazon, Microsoft, Swiggy, Zomato) - Modern testing practices

---

### **Q: @MockBean vs @SpyBean?**
*   **@MockBean:** Replaces bean with a hollow Mockito mock.
*   **@SpyBean:** Wraps a **real** bean, allowing you to track calls or stub specific methods.

**🏢 Asked in:** Product companies (Amazon, Microsoft, Google) + Service companies - Advanced testing concepts
