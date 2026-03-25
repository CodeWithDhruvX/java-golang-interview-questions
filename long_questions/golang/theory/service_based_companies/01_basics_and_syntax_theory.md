# 🗣️ Theory — Go Basics & Syntax
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What is Go and why was it created?"

> *"Go, also called Golang, is a statically typed, compiled programming language created by Google in 2009 — by Robert Griesemer, Rob Pike, and Ken Thompson. It was built because Google was frustrated with the slowness of C++ compilation, the complexity of Java, and the lack of performance in dynamic languages like Python. So they created Go to have fast compile times, built-in concurrency support, a simple syntax, and strong performance — almost like if C and Python had a baby that was designed for the modern web."*

---

## Q: "What is the difference between `var`, `:=`, and `const`?"

> *"So, `var` is the explicit way to declare a variable. You can use it inside a function or at the package level. `var name string = 'Go'`. The `:=` — called the short variable declaration — is the shorthand. It only works inside functions, and it automatically infers the type. So `name := 'Go'` is the same thing but shorter. And `const` is for values that cannot change — they're fixed at compile time. Like `const Pi = 3.14`. The key distinction: `var` and `:=` are mutable, `const` is immutable, and `:=` cannot be used at package level."*

---

## Q: "What are zero values in Go?"

> *"Zero values are Go's way of ensuring every variable has a safe default — there's no such thing as an uninitialized variable in Go. For numeric types like `int` and `float64`, the zero value is `0`. For `bool`, it's `false`. For a `string`, it's an empty string. For pointers, slices, maps, channels, and functions, the zero value is `nil`. This is actually a really clean design decision — it prevents a whole class of bugs you'd see in C where reading uninitialized memory gives you garbage."*

---

## Q: "Explain how `defer` works."

> *"Defer is a statement that schedules a function call to run right before the surrounding function returns — think of it like 'run this on the way out'. The most common use is cleanup: `defer f.Close()` right after you open a file, so you never forget to close it. What's interesting is that if you have multiple defer statements, they run in LIFO order — last in, first out — like a stack. Also, defer runs even if the function panics. And here's a subtle point — if you defer a function call, the arguments to that function are evaluated immediately, not when the defer actually runs."*

---

## Q: "What is the difference between `new()` and `make()`?"

> *"Both allocate memory but for very different purposes. `new(T)` allocates memory for any type, zeroes it out, and returns a pointer to it — so `new(int)` gives you a `*int` pointing to zero. `make()`, on the other hand, is only for three special types — slices, maps, and channels — and it actually initializes their internal structure, not just allocates raw memory. So `make([]int, 5)` creates a ready-to-use slice with length 5, whereas `new([]int)` just gives you a pointer to a nil slice. In practice, you use `make` far more than `new`."*

---

## Q: "What is a variadic function?"

> *"A variadic function is one that can accept a variable number of arguments of the same type. You declare it with `...` before the type — like `func sum(nums ...int) int`. Inside the function, `nums` is just a regular slice. To call it you can pass arguments directly — `sum(1, 2, 3)` — or you can spread a slice with the `...` operator — `sum(mySlice...)`. The most famous example in Go is `fmt.Println()` which uses `...interface{}` to accept any number of any type."*

---

## Q: "What is the blank identifier `_` used for?"

> *"The blank identifier is Go's way of saying 'I know this value exists but I don't want it'. It's used in three main scenarios: first, when a function returns multiple values and you only need some of them — like `value, _ := strconv.Atoi('42')`, where you ignore the error. Second, in for-range loops when you don't care about the index or the value. Third — and this is a really neat Go trick — you use it to verify that a type implements an interface at compile time: `var _ io.Writer = (*MyType)(nil)`. If `MyType` doesn't implement `io.Writer`, you get a compile error right there."*

---

## Q: "How does the `for` loop work in Go? Are there while loops?"

> *"Go has only one looping keyword — `for` — but it works in three ways. First, the traditional C-style: `for i := 0; i < 5; i++`. Second, the while-style — you just write `for condition { }` with no init or post. Third, the infinite loop: `for { }` — just `for` with nothing, which you break out of with `break` or `return`. And then there's `for range` for iterating over slices, maps, strings, and channels. So Go doesn't have a separate `while` keyword — `for` does it all. People coming from Java or Python find this strange at first but it actually simplifies the language."*

---

## Q: "What is a type alias vs a type definition?"

> *"A type alias with `type A = B` means A and B are literally the same type — interchangeable everywhere. A type definition with `type A B` creates a brand new distinct type that happens to have the same underlying representation as B. The distinction matters for method sets and type safety. For example, `type Celsius float64` creates a distinct type — you can't accidentally mix it with a plain `float64` without an explicit conversion. This is how Go achieves type safety without having full inheritance. You'd use aliases mostly for compatibility when refactoring packages."*

---

## Q: "What is `init()` and when does it run?"

> *"The `init()` function is a special function in Go that runs automatically before `main()`. Its job is package initialization — setting up global variables, registering things, loading config, that kind of thing. Every package can have one or more `init()` functions, and they all run in the order the files are processed. The `_` import — `import _ 'github.com/lib/pq'` — is specifically to trigger a package's `init()` for side effects like registering a database driver, without actually using the package's exported symbols."*

---


