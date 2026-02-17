# 🟡 Go Theory Questions: 121–141 Generics, Type System, and Advanced Types

## 121. What is type inference in Go?

**Answer:**
Type inference is the compiler's ability to figure out what type a variable is without you spelling it out.

When you write `x := 10`, you didn't say `int`, but the compiler knows `10` is an integer, so `x` becomes an `int`.

However, Go's inference is intentionally limited. It only works inside functions, and it doesn't work for function arguments. You can't write `func(x)`—you must write `func(x int)`. This strikes a balance: it saves you typing inside logic blocks but preserves explicit documentation at API boundaries.

---

## 122. What are generics in Go?

**Answer:**
Generics, introduced in Go 1.18, allow you to write functions and data structures that work with *any* type, not just one specific type.

Before generics, if you wanted a `Min()` function, you had to write `MinInt`, `MinFloat`, etc. Now you write `func Min[T Ordered](a, b T) T`.

Under the hood, Go uses **Monomorphization** (or Stenciling). When you compile your code, Go effectively copy-pastes your generic function for every concrete type you use. If you use `Min` with ints and floats, the compiler generates two optimized versions of the function. This means generics have **zero runtime cost**—they are as fast as hand-written code.

---

## 123. How do you use generics with struct types?

**Answer:**
You define a struct with a type parameter, like `type List[T any] struct { items []T }`.

This creates a blueprint. When you create a list, you specify the type: `List[int]` or `List[string]`.

This is massive for data structures. We used to rely on `interface{}` to build linked lists or trees, which was dangerous because you could accidentally put a cat in a list of dogs. Generics give us compile-time type safety—a `List[int]` will refuse to accept a string, preventing bugs before you even run the code.

---

## 124. Can you restrict generic types using constraints?

**Answer:**
Yes, and you have to. If you just say `[T any]`, you can't do much with `T`—you can't add it or compare it, because not all types support `+` or `<`.

So we use **Constraints**. A constraint is just an interface. You can say `[T constraints.Ordered]`.

This tells the compiler: "T can be anything, as long as it supports greater-than and less-than operators." This allows you to write math functions effectively while still blocking invalid types like structs or maps.

---

## 125. How to create reusable generic containers (e.g., Stack)?

**Answer:**
Refactoring a container to be generic is straightforward. You replace specific types with a parameter `T`.

`type Stack[T any] struct { data []T }`.

Then, all your methods—Push, Pop—use `T`. `func (s *Stack[T]) Push(v T)`. This one generic implementation replaces the need for `IntStack`, `StringStack`, and `UserStack`. It’s cleaner, safer, and just as fast.

---

## 126. What is the difference between `any` and `interface{}`?

**Answer:**
Mechanically? None. Zero.

`any` is literally a built-in alias: `type any = interface{}`.

It was introduced because `interface{}` is annoying to type and hard to read when you have complex generic signatures like `func Map[K comparable, V interface{}]`. Writing `[K comparable, V any]` is just much easier on the eyes. Use `any` everywhere unless you're maintaining 5-year-old legacy code.

---

## 127. Can you have multiple constraints in a generic function?

**Answer:**
Yes, you can have as many as you need.

A standard map function needs two types: `func Keys[K comparable, V any](m map[K]V) []K`.

Here, `K` must be **comparable** (because map keys must be hashable), but `V` can be **any**. You can list them out just like function arguments.

---

## 128. Can interfaces be used in generics?

**Answer:**
In Go, interfaces play a dual role. They are used as **Constraints**.

In the old days, interfaces only defined methods. Now, interfaces can also define **Type Sets**. You can write an interface that says "I can be an `int` OR a `float`."

When used as a generic constraint, this interface limits what types are allowed. But be careful—you can't use these "type set" interfaces as regular variables, only as generic constraints.

---

## 129. What is type embedding and how does it differ from inheritance?

**Answer:**
Embedding allows you to put one struct inside another anonymously.

If `Truck` embeds `Vehicle`, `Truck` automatically gets all of `Vehicle`'s methods. You can call `truck.StartEngine()`.

However, this is **NOT inheritance**. `Truck` is not a `Vehicle`. You cannot pass a `Truck` to a function expecting a `Vehicle` (unless you use interfaces). It is purely a mechanism for composition—syntactic sugar to forward method calls—not for building a class hierarchy.

---

## 130. How does Go perform type conversion vs. type assertion?

**Answer:**
**Type Conversion** changes the data. `float64(10)` takes the integer bits and rewrites them into floating-point format. It changes the representation.

**Type Assertion** checks the data. `x.(int)` looks at an interface variable and asks, "Is the thing inside you an int?" It doesn't change the memory; it just unwraps it.

Conversion happens between compatible types (numbers). Assertion happens between Interfaces and Concrete types.

---

## 131. What are tagged unions and how can you simulate them in Go?

**Answer:**
A Tagged Union (or Sum Type) is a type that can be one of several specific things—like "Result is either Success OR Failure".

Go doesn't have native support for this (like Rust's `enum`), but we simulate it using **Sealed Interfaces**.

You define an interface with a private method: `isResult()`. Then you define `Success` and `Failure` structs that implement it. Since the method is private, nobody outside your package can add new types to it. This gives you a closed set of possible types, simulating a union.

---

## 132. What is the use of `iota` in Go?

**Answer:**
`iota` is Go’s counter for creating enums.

Inside a `const` block, `iota` starts at 0 and increments by 1 on every line.
`const ( Red = iota; Blue; Green )`. Red is 0, Blue is 1, Green is 2.

It’s convenient, but dangerous. If you delete "Blue", "Green" shifts down to 1. If you stored these values in a database, your data is now corrupted. We often manually assign values (`Green = 2`) to avoid this drift.

---

## 133. How are custom types different from type aliases?

**Answer:**
`type MyInt int` creates a **New Type**. `MyInt` and `int` are incompatible. You can't mix them. This is what we use 99% of the time for domain safety (like `type Password string`).

`type MyInt = int` creates an **Alias**. They are the same thing.

We only use aliases for refactoring—when we move a type to a new package but want to keep the old name available for backward compatibility.

---

## 134. What are type sets in Go 1.18+?

**Answer:**
Type Sets are a new concept added for Generics.

Previously, an interface defined a set of **methods**. Now, an interface can define a set of **types** using the pipe operator: `type Number interface { int | float64 }`.

This interface doesn't mean "has methods," it means "is one of these types." This allows valid operations. If you know T is `int | float`, the compiler lets you use `+` and `*` on it.

---

## 135. Can generic types implement interfaces?

**Answer:**
Yes. A generic struct like `List[T]` can implement the `Stringer` interface.

You just define `func (l List[T]) String() string`.

However, the interface *itself* cannot have type parameters. You can't define `type Printer[T] interface`. Interfaces describe behavior, and behavior shouldn't depend on the specific generic type used instantiation.

---

## 136. How do you handle constraints with operations like +, -, *?

**Answer:**
You cannot use operators on generic types by default. `func Add[T any](a, b T) T { return a + b }` fails to compile because `T` might be a struct.

To fix this, you use constraints that include primitive types. `type Addable interface { int | float64 | string }`.

If you constrain `T` to `Addable`, the compiler verifies that *every* type in that set supports the `+` operator.

---

## 137. What is structural typing?

**Answer:**
Structural typing means compatibility is determined by the **structure** (methods), not the name.

In Java, you explicitly say `class A implements I`. In Go, if your struct happens to have the right method signatures, it inadvertently implements the interface.

This is mostly good—it decouples code—but it allows for **accidental implementation**. If you have `pop()` method for a balloon, you might accidentally satisfy a `stack` interface, leading to confusion.

---

## 138. Explain the difference between concrete and abstract types.

**Answer:**
Concrete types (structs, ints) tell the computer **how much memory** to allocate. You can create instances of them.

Abstract types (interfaces) tell the compiler **what operations** are allowed. You cannot instantiate an interface directly; you can only store a concrete instance *inside* it.

Programming in Go is about defining behavior with abstract types and implementing it with concrete types.

---

## 139. What are phantom types and are they used in Go?

**Answer:**
A phantom type is a weird generic trick: `type ID[T any] string`.

The struct holds a string, but it also has a type tag `T` that isn't actually used in any field.

We use this for **extra type safety**. We can distinguish `ID[User]` from `ID[Post]` at compile time. They are both just strings at runtime, so it costs nothing, but the compiler ensures you never pass a Post ID to a function expecting a User ID.

---

## 140. How would you implement an enum pattern in Go?

**Answer:**
Go doesn't have a native Enum type.

We simulate it using a custom integer type and constants. `type Role int`. `const ( Admin Role = iota; User )`.

The trade-off is safety. Nothing stops a user from casting `Role(99)` and passing it to your function. Unlike Rust or Swift, Go enums are not "closed sets," so you always need a `default` case in your switch statements to handle invalid values.

---

## 141. How can you implement optional values in Go idiomatically?

**Answer:**
We don't have `Optional<T>`.

The most common way is **Pointers**. `*int`. If it’s `nil`, it’s empty. If it has a value, it’s set. The risk is accidental nil-pointer dereferences.

The safer, more verbose way is the **Ok-Check struct**: `sql.NullString` has `{String: "foo", Valid: true}`. We generally prefer pointers for simplicity, but use Null types for database interactions where "NULL" is a specific concept distinct from "Zero".
