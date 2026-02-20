# 12. Extra Concepts (Practice)

**Q: When would you use Optional, and when should you avoid it?**
> "**Optional** is a container object that might (or might not) contain a value. It was introduced in Java 8 to avoid `NullPointerException`.
>
> You **should** use it as a *return type* for methods that might not find a result. Like `findUserById()`. It forces the caller to think: 'What if the user isn't found?' and handle it gracefully using methods like `.orElse()` or `.ifPresent()`.
>
> You **should generally avoid** using it for:
> 1.  Field variables (it's not serializable).
> 2.  Method parameters (it just makes calling the method annoying).
> 3.  Collections (never put Optional in a List, just have an empty List)."

**Indepth:**
> **Performance**: `Optional` is an object. Creating it adds overhead. Using it deeply in tight loops or for every single field in a massive data structure will hurt performance and GC.


---

**Q: Why are generics invariant in Java?**
> "This is a tricky one. Invariance means `List<String>` is **not** a subtype of `List<Object>`.
>
> Why? Because of type safety.
> If Java allowed you to treat a `List<String>` as a `List<Object>`, you could add an `Integer` to it!
>
> ```java
> List<String> strings = new ArrayList<>();
> List<Object> objects = strings; // If this were allowed...
> objects.add(10); // You just put an int into a list of strings!
> ```
> When you try to read that 'int' back as a String, your program would crash. So Java prevents this at compile time by making generics invariant."

**Indepth:**
> **Covariance**: Generics *can* be covariant using wildcards (`List<? extends Number>`). This allows reading (you know everything inside is at least a Number) but prevents writing (you don't know if it's meant to hold Integers or Doubles).


---

**Q: Strategy Pattern real-world use case?**
> "The **Strategy Pattern** is about swapping algorithms at runtime.
>
> Think of a Payment System on an e-commerce site. You have a `pay()` method.
> But the user might want to pay with **Credit Card**, **PayPal**, or **Bitcoin**.
>
> Instead of writing one giant `if-else` block inside the `pay()` method, you define a `PaymentStrategy` interface. Then you create classes `CreditCardStrategy`, `PayPalStrategy`, etc.
>
> You pass the chosen strategy to the payment processor. This makes it super easy to add a new payment method later (like Apple Pay) without touching the existing code."

**Indepth:**
> **Open/Closed Principle**: This is the textbook example of OCP. Classes should be open for extension (adding new Strategies) but closed for modification (not touching the `pay()` method).


---

**Q: Abstract Factory vs Factory Method?**
> "They both create objects, but strictly speaking:
>
> **Factory Method** uses *inheritance*. You have a method `createAnimal()` in a base class, and subclasses override it to return a `Dog` or `Cat`. It creates *one* product.
>
> **Abstract Factory** uses *composition*. It's a factory *of factories*. It creates *families* of related products.
> Like a `GUIFactory` that creates a `Button`, `Checkbox`, and `Scrollbar`. You might have a `WindowsFactory` that returns Windows-style buttons and checkboxes, and a `MacFactory` that returns Mac-style ones. You ensure that all components match the same theme."

**Indepth:**
> **Dependency Inversion**: Abstract Factory allows the client code to be completely decoupled from concrete classes. It only knows about the interfaces (`Button`, `Window`). This makes cross-platform UI toolkits possible.


---

**Q: StackOverflowError Simulation**
> "A `StackOverflowError` happens when the call stack gets too deep, usually due to **infinite recursion**.
>
> To simulate it, just write a method that calls itself without a breaking condition:
>
> ```java
> public void recursive() {
>     recursive();
> }
> ```
> Run that, and boom—StackOverflow."

**Indepth:**
> **Tail Call Optimization**: Java does *not* support tail call optimization (yet). So even if the recursive call is the very last thing, it still consumes a stack frame.


---

**Q: OutOfMemoryError Simulation**
> "An `OutOfMemoryError` (OOM) happens when the **Heap** is full.
>
> To simulate it, just keep creating objects and holding onto them so the Garbage Collector can't delete them.
>
> ```java
> List<byte[]> list = new ArrayList<>();
> while (true) {
>     list.add(new byte[1024 * 1024]); // Add 1MB chunks continuously
> }
> ```
> Eventually, the heap fills up, and you crash."

**Indepth:**
> **Analysis**: When OOM happens, you need a Heap Dump. Tools like Eclipse MAT or VisualVM can analyze this dump to find the "Leak Suspects" (which objects are consuming the most RAM).

