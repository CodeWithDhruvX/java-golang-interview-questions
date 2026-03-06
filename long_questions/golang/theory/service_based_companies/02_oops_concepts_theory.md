# 🗣️ Theory — OOP Concepts in Go
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "Does Go support Object-Oriented Programming?"

> *"Go doesn't have classes, inheritance, or the traditional OOP model from Java or C++. But it absolutely supports object-oriented design — just through a different mechanism. Instead of classes, you use structs with methods. Instead of inheritance, you use composition through struct embedding. And instead of interfaces being explicitly declared, they're implicitly satisfied — a type just needs to have the right methods. So people say Go prefers composition over inheritance, which is actually a well-known best practice even in the OOP world. It's a cleaner, more flexible model."*

---

## Q: "What are structs in Go?"

> *"Structs are Go's way of grouping related data together — similar to a class but without methods baked in. You define one with `type Employee struct { Name string; Age int }`. Then you create instances using a struct literal: `e := Employee{Name: 'Alice', Age: 30}`. Methods are defined separately but associated with the struct using a receiver. The key thing about Go structs is that there's no constructor — you create factory functions like `NewEmployee()` that return a struct, which is the idiomatic Go pattern."*

---

## Q: "What is the difference between value receivers and pointer receivers?"

> *"When you define a method on a struct, you choose between a value receiver and a pointer receiver. A value receiver — `func (c Circle) Area()` — works on a copy of the struct. So the method can read the data but any changes it makes don't affect the original. A pointer receiver — `func (c *Circle) Scale()` — works on the actual struct, so it can modify it. The rule of thumb is: use a pointer receiver when you need to mutate the struct, or when the struct is large and you don't want to copy it every time. Also, for consistency, if your type has ANY pointer receiver methods, make ALL your methods pointer receivers."*

---

## Q: "How do interfaces work in Go? What does 'implicit satisfaction' mean?"

> *"An interface in Go is just a set of method signatures. Any type that has all those methods automatically implements the interface — there's no `implements` keyword. This is called implicit or duck typing. So if I have an interface `Shape` with `Area() float64`, and my `Circle` struct has that method, then `Circle` implements `Shape` — automatically, without me declaring it. This is powerful because you can define an interface long after a type was created, and existing types will satisfy it. It also means Go interfaces are typically very small — often just one method."*

---

## Q: "What is the empty interface? What is `any`?"

> *"The empty interface `interface{}` has zero method requirements, which means every type in Go satisfies it. It's Go's equivalent of `Object` in Java — a way to pass any value without caring about its type. Since Go 1.18, `any` is just an alias for `interface{}`, so they're identical. The downside is you lose type safety — to actually use a value stored as `interface{}`, you need a type assertion to get the concrete type back. It's used in generic containers, JSON decoding into unknown structures, and function parameters that accept anything."*

---

## Q: "What is type assertion? What's the difference between safe and unsafe assertions?"

> *"Type assertion is how you extract the concrete type from an interface value. If I have `var i interface{} = 'hello'`, I can do `s := i.(string)` to get the string out. But this panics if `i` is not actually a string. The safe version uses the two-return form: `s, ok := i.(string)` — if `ok` is true, `s` is the string; if false, `s` is the zero value and no panic. In practice, always use the two-return form — single-return assertions are only appropriate when you're absolutely certain of the type."*

---

## Q: "What is struct embedding? How is it different from inheritance?"

> *"Struct embedding is when you include one struct type inside another without a field name — like `type Dog struct { Animal }`. This 'promotes' all of Animal's fields and methods to Dog. So `d.Name` works even though Name is defined on Animal. This looks like inheritance but it's fundamentally different — there's no IS-A relationship. Dog doesn't 'extend' Animal. Dog simply reuses Animal by containing it. And you can always shadow a promoted method by defining your own with the same name on Dog. Multiple embedding is also possible, giving you mixin-like behavior."*

---

## Q: "What is polymorphism in Go and how is it achieved?"

> *"Polymorphism in Go is achieved through interfaces. If multiple types all implement the same interface, you can write functions that work with any of them uniformly, without knowing the concrete type at compile time. For example, if `EmailNotifier`, `SMSNotifier`, and `PushNotifier` all implement the `Notifier` interface, you can have a `sendAlert(n Notifier)` function that works with all three. That's runtime polymorphism — the right `Send()` method gets called based on the actual type in the interface variable. There's no `virtual` keyword needed — Go's interfaces are always polymorphic."*

---

## Q: "What is the nil interface gotcha in Go?"

> *"This is a classic Go interview question. A nil interface and an interface holding a nil pointer are not the same thing. An interface value has two components: a type and a value. A nil interface has both set to nil. But if you assign a typed nil pointer — like `var d *Dog = nil; var a Animal = d` — the interface `a` has the type `*Dog` but the value nil. So `a == nil` is false, even though the underlying value is nil. This catches a lot of people off guard, especially when returning errors from functions that use concrete types."*
