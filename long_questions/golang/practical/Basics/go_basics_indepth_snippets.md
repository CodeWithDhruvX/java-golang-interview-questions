# 100 In-Depth Go Basics & Fundamentals — Pure Code Snippet Questions

> **Format**: Each question is "predict the output / spot the bug / does it compile?" style.
> **Focus**: Subtle edge cases, gotchas, and deeper mechanics of Go's type system, memory model, and runtime.

---

## 📋 Reading Progress

> Mark each section `[x]` when done. Use `🔖` to note where you left off.

- [ ] **Section 1:** Variables, Scope & Type System Deep Dives (Q1–Q18)
- [ ] **Section 2:** Control Flow Deep Dives (Q19–Q30)
- [ ] **Section 3:** Functions & Closures Deep Dives (Q31–Q45)
- [ ] **Section 4:** Pointers Deep Dives (Q46–Q56)
- [ ] **Section 5:** Strings & Runes Deep Dives (Q57–Q67)
- [ ] **Section 6:** Slices & Maps Deep Dives (Q68–Q83)
- [ ] **Section 7:** Structs & Interfaces Deep Dives (Q84–Q95)
- [ ] **Section 8:** Error Handling Deep Dives (Q96–Q100)

> 🔖 **Last read:** <!-- e.g. Q30 · Section 2 done -->

---

## Section 1: Variables, Scope & Type System Deep Dives (Q1–Q18)

### 1. Short Declaration Reuse in Same Scope
**Q: Does this compile?**
```go
package main
import "fmt"

func main() {
    x, y := 1, 2
    x, z := 3, 4
    fmt.Println(x, y, z)
}
```
**A:** **Yes, compiles and prints** `3 2 4`. `:=` is valid when at least one variable on the left is **new** (`z`). It reassigns `x` rather than redeclaring it.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this code do and why?
**Your Response:** This code compiles and prints `3 2 4`. The short variable declaration `:=` works here because at least one variable (`z`) is new. When `:=` is used with existing variables, it reassigns them rather than redeclaring. So `x` gets reassigned to `3`, `y` stays `2`, and `z` is declared as `4`.

---

### 2. Redeclaration in Inner Scope Gets Separate Variable
**Q: What exactly is printed and why?**
```go
package main
import "fmt"

func main() {
    x := 1
    {
        x, y := 2, 3
        fmt.Println(x, y)
    }
    fmt.Println(x)
}
```
**A:**
```
2 3
1
```
The inner block's `x, y := 2, 3` creates a **brand-new** `x` (because it is a new scope), shadowing the outer one. The outer `x` remains `1`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What gets printed and why do we get two different values for x?
**Your Response:** This prints `2 3` then `1`. The inner block creates a completely new variable `x` because it's a different scope. This is called variable shadowing - the inner `x` shadows the outer one. When we exit the inner block, we're back to the original outer `x` which still has value `1`.

---

### 3. Constant Expression Evaluation
**Q: What is the output?**
```go
package main
import "fmt"

const (
    a = 2
    b = 3
    c = a * b + 1
)

func main() {
    fmt.Println(c)
}
```
**A:** `7`. Constant expressions are evaluated at compile time with arbitrary-precision arithmetic.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the output and how does this work?
**Your Response:** This prints `7`. Go evaluates constant expressions at compile time using arbitrary-precision arithmetic. So `a * b + 1` becomes `2 * 3 + 1` which is `7`. This happens during compilation, not runtime.

---

### 4. Typed vs Untyped Constant Assignment
**Q: Does this compile?**
```go
package main
import "fmt"

func main() {
    const x int = 5
    var y float64 = x
    fmt.Println(y)
}
```
**A:** **Compile Error.** `x` is a **typed** constant of type `int`, so it cannot be assigned to `float64` without an explicit conversion. Compare: an **untyped** `const x = 5` would work fine.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and why or why not?
**Your Response:** This doesn't compile. The issue is that `x` is a typed constant (`const x int = 5`). When you have a typed constant, you can't assign it to a different type without explicit conversion. If it were an untyped constant like `const x = 5`, Go would automatically convert it to `float64`. You'd need to write `var y float64 = float64(x)` to make this work.

---

### 5. Type Alias Shares Method Set
**Q: Does this compile?**
```go
package main
import "fmt"

type MyString = string  // alias, not new type

func main() {
    var s MyString = "hello"
    var t string = s
    fmt.Println(t)
}
```
**A:** **Yes, compiles and prints** `hello`. A type alias (`=`) is the exact same type, so no conversion is needed.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and what's the difference between type alias and new type?
**Your Response:** Yes, this compiles and prints `hello`. Here `MyString` is a type alias, not a new type. The `=` syntax means `MyString` is exactly the same type as `string` - just another name for it. That's why we can assign between them without conversion. If we used `type MyString string` (without `=`), it would create a new type and require explicit conversion.

---

### 6. Named Type Does NOT Share Method Set With Underlying Type
**Q: Does this compile?**
```go
package main

type MyInt int

func double(n int) int { return n * 2 }

func main() {
    x := MyInt(5)
    _ = double(x)
}
```
**A:** **Compile Error.** `MyInt` is a distinct type from `int`. You must convert: `double(int(x))`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why doesn't this compile?
**Your Response:** This doesn't compile because `MyInt` is a completely different type from `int`, even though its underlying type is `int`. Go's type system is strict - you can't pass a `MyInt` to a function expecting an `int` without explicit conversion. You'd need to call `double(int(x))` to convert it first. This is different from type aliases which are the same type.

---

### 7. Multiple Return Ignore with Blank
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    n, _ := strconv.Atoi("42")
    fmt.Printf("%T %v\n", n, n)
}
```
**A:** `int 42`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What gets printed here?
**Your Response:** This prints `int 42`. The `strconv.Atoi` function returns two values - the integer and an error. We're using the blank identifier `_` to ignore the error value. The `fmt.Printf` with `%T` shows the type is `int` and `%v` shows the value is `42`.

---

### 8. Three-Component for and Post Statement
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    for i := 0; i < 3; i += 2 {
        fmt.Print(i, " ")
    }
}
```
**A:** `0 2 `. The post-statement is `i += 2`, so it goes `0 → 2 → 4` (stops because 4 ≥ 3).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why?
**Your Response:** This prints `0 2`. The for loop has a custom post-statement `i += 2` instead of the usual `i++`. So the sequence is: start at `0`, check if `0 < 3` (true), print `0`, then increment by `2` to get `2`. Next iteration: check if `2 < 3` (true), print `2`, increment to `4`. Next iteration: check if `4 < 3` (false), exit. So we only see `0` and `2`.

---

### 9. Zero Value of a Function Type
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var f func()
    fmt.Println(f == nil)
}
```
**A:** `true`. The zero value of a function type is `nil`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `true`. When we declare a variable of function type like `var f func()` without initializing it, it gets the zero value for that type, which is `nil`. So the comparison `f == nil` evaluates to `true`.

---

### 10. Comparing Function Types
**Q: Does this compile?**
```go
package main

func main() {
    f := func() {}
    g := func() {}
    _ = f == g
}
```
**A:** **Compile Error.** Function values are not comparable with `==` in Go (except to `nil`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. In Go, you can't compare function values with `==` except when comparing to `nil`. Functions are reference types, and Go doesn't define equality for them. This is a design choice to avoid complexity around function identity and closure comparison.

---

### 11. Integer Literal Representations
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    dec := 42
    hex := 0x2A
    oct := 0o52
    bin := 0b101010
    fmt.Println(dec, hex, oct, bin)
}
```
**A:** `42 42 42 42`. All four literals represent the same value.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `42 42 42 42`. Go supports multiple integer literal formats: decimal (`42`), hexadecimal (`0x2A`), octal (`0o52`), and binary (`0b101010`). They all represent the same value, just written in different number bases.

---

### 12. Floating-Point Precision Comparison
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := 0.1 + 0.2
    fmt.Println(a == 0.3)
    fmt.Printf("%.17f\n", a)
}
```
**A:**
```
false
0.30000000000000004
```
IEEE 754 floating-point arithmetic is imprecise. Never compare floats with `==`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the output and why is the comparison false?
**Your Response:** This prints `false` then `0.30000000000000004`. The issue is floating-point precision - `0.1 + 0.2` doesn't exactly equal `0.3` in binary floating-point representation. This is a fundamental limitation of IEEE 754. The lesson is: never compare floating-point numbers with `==`. Use tolerance-based comparison instead.

---

### 13. var Block Dependency Order
**Q: Does this compile?**
```go
package main
import "fmt"

var (
    x = y + 1
    y = 2
)

func main() {
    fmt.Println(x, y)
}
```
**A:** **Yes, compiles and prints** `3 2`. Package-level `var` declarations are resolved by the compiler regardless of order.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile? What about the dependency order?
**Your Response:** Yes, this compiles and prints `3 2`. At package level, Go allows forward references in variable declarations. The compiler resolves dependencies across the entire package, so `x = y + 1` works even though `y` is declared after. This is different from inside functions where you must declare before use.

---

### 14. := in if Condition and Scope
**Q: Is `err` accessible after the if-else block?**
```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    if n, err := strconv.Atoi("bad"); err != nil {
        fmt.Println("error:", err)
    } else {
        fmt.Println("value:", n)
    }
    // fmt.Println(n) // Compile error: n undefined
}
```
**A:** No. Both `n` and `err` are scoped to the `if-else` block. The print of `error: strconv...` executes. Uncommenting the last line would cause a compile error.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Is `err` accessible after the if-else block?
**Your Response:** No, `err` is not accessible after the if-else block. When you use the `if` statement with initialization (`if n, err := ...`), both variables are scoped only to that if-else block. This is intentional design to keep variables tightly scoped. If you need `err` outside, you'd have to declare it before the if statement.

---

### 15. Constant Converted to String
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    const n = 65
    s := string(n)
    fmt.Println(s)
}
```
**A:** `A`. `string(65)` treats the integer as a Unicode code point (rune), producing the character `A`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `A`. When you convert an integer to a string using `string()`, Go treats the integer as a Unicode code point (rune). Since 65 is the ASCII/Unicode code point for 'A', it produces that character. This is different from converting a number to its string representation, which would require `strconv.Itoa`.

---

### 16. Shadowing Built-in Identifier
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    len := 42
    fmt.Println(len)
}
```
**A:** `42`. Go allows shadowing built-in identifiers like `len`, `make`, `new`, etc. (this is generally a bad practice but is valid).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `42`. Go allows you to shadow built-in identifiers like `len`, `make`, `new`, etc. When you declare `len := 42`, you're creating a local variable that shadows the built-in `len` function. While valid, this is generally considered bad practice as it can confuse readers of your code.

---

### 17. Untyped Numeric Constant Default Type
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    const x = 3
    fmt.Printf("%T\n", x)
}
```
**A:** `int`. An untyped integer constant defaults to `int` when its type needs to be determined (e.g., as an argument to `%T`).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What type is printed and why?
**Your Response:** This prints `int`. The constant `x = 3` is untyped, so when Go needs to determine its type (like when we use `%T` in Printf), it defaults to `int`. This is Go's default type for untyped integer constants. If we used it in a float context, it would become `float64`.

---

### 18. Package-Level Init Ordering
**Q: What is the output?**
```go
package main
import "fmt"

var (
    msg   = greet()
    hello = "Hello"
)

func greet() string { return hello + " World" }

func main() {
    fmt.Println(msg)
}
```
**A:** ` World`. `hello` is `""` (zero value) when `greet()` runs because `msg = greet()` is initialized first. Package-level variables are initialized in declaration order, not dependency order for cross-variable references here.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What gets printed and why isn't it "Hello World"?
**Your Response:** This prints ` World`. The issue is initialization order. Package-level variables are initialized in declaration order, not dependency order. So `msg` gets initialized first by calling `greet()`, but at that point `hello` is still its zero value (empty string). Then `hello` gets set to "Hello". This shows why you should be careful with initialization dependencies at package level.

---

## Section 2: Control Flow Deep Dives (Q19–Q30)

### 19. Labeled Continue
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                continue outer
            }
            fmt.Printf("(%d,%d) ", i, j)
        }
    }
}
```
**A:** `(0,0) (1,0) (2,0) `. When `j == 1`, `continue outer` skips the rest of the inner loop AND the rest of the outer loop body, going straight to the next outer iteration.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and how does the labeled continue work?
**Your Response:** This prints `(0,0) (1,0) (2,0)`. The `continue outer` statement is powerful - when `j == 1`, it doesn't just continue the inner loop, it jumps to the next iteration of the outer loop labeled `outer`. This means it skips both the rest of the inner loop body AND any remaining code in the outer loop body for that iteration.

---

### 20. Switch Case with Initialization
**Q: What is the output?**
```go
package main
import "fmt"

func val() int { return 2 }

func main() {
    switch v := val(); {
    case v == 1:
        fmt.Println("one")
    case v == 2:
        fmt.Println("two")
    default:
        fmt.Println("other")
    }
}
```
**A:** `two`. The `switch v := val();` form initializes `v` before the switch. The trailing `;` with no expression means `switch true`, making each case a boolean expression.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and what's this switch syntax?
**Your Response:** This prints `two`. This is a switch with initialization. The `switch v := val();` first declares and initializes `v`, then the trailing semicolon with no expression makes it equivalent to `switch true`. Each case then becomes a boolean expression that's evaluated in order. This is a clean alternative to if-else chains.

---

### 21. For Range Over Channel
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)

    for v := range ch {
        fmt.Print(v, " ")
    }
}
```
**A:** `1 2 3 `. `for range` over a channel receives until the channel is closed and drained.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `1 2 3`. The `for range` over a channel automatically receives values until the channel is closed and drained. We put 3 values in the buffered channel, closed it, then the range loop receives each one. When the channel is empty and closed, the range loop exits automatically.

---

### 22. Fallthrough Cannot Be Last in Case
**Q: Does this compile?**
```go
package main
import "fmt"

func main() {
    switch 1 {
    case 1:
        fallthrough
    case 2:
        fmt.Println("matched")
    case 3:
        fmt.Println("three")
        fallthrough
    }
}
```
**A:** **Compile Error.** `fallthrough` in the last case (`case 3`) is not allowed. There is no next case to fall into.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. The issue is `fallthrough` in the last case (`case 3`). `fallthrough` can only be used when there's a next case to fall into. In the last case, there's nowhere to fall, so it's a compile error. This prevents logical errors in your switch statements.

---

### 23. Break Out of Select Inside For
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    ch := make(chan int, 1)
    ch <- 42
    for {
        select {
        case v := <-ch:
            fmt.Println(v)
            break // breaks the select, NOT the for
        }
        break // this breaks the for
    }
}
```
**A:** `42`. The first `break` exits only the `select`. The second `break` exits the `for` loop. This is a subtle trap: `break` inside `select` only breaks the `select`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `42`. The key insight here is that `break` inside a `select` statement only breaks the `select`, not the surrounding `for` loop. So the first `break` exits the `select`, then the second `break` exits the `for` loop. This is a common gotcha - many people expect the first break to exit the for loop.

---

### 24. Blank Switch (switch true)
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x := 25
    switch {
    case x > 100:
        fmt.Println("huge")
    case x > 10:
        fmt.Println("big")
    case x > 0:
        fmt.Println("positive")
    }
}
```
**A:** `big`. Cases are evaluated top-to-bottom; the first true one wins.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `big`. This is a blank switch (equivalent to `switch true`). Cases are evaluated from top to bottom, and the first one that evaluates to true wins. Since `25 > 100` is false but `25 > 10` is true, it prints `big` and doesn't check the remaining cases.

---

### 25. Range Loop Variable Reuse
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    funcs := make([]func(), 3)
    for i := 0; i < 3; i++ {
        funcs[i] = func() { fmt.Print(i) }
    }
    for _, f := range funcs {
        f()
    }
}
```
**A:** `333`. All closures capture the same `i`; after the loop, `i == 3`.  
**Fix:** `j := i; funcs[i] = func() { fmt.Print(j) }`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why?
**Your Response:** This prints `333`. This is a classic Go closure gotcha. All the closures capture the same variable `i` by reference, not by value. When the closures finally execute, they all see the final value of `i` which is `3`. The fix is to create a local copy: `j := i` inside the loop, so each closure captures its own value.

---

### 26. For with Post Increment Inside Body
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    i := 0
    for i < 5 {
        i++
        if i == 3 {
            continue
        }
        fmt.Print(i, " ")
    }
}
```
**A:** `1 2 4 5 `. When `i == 3`, `continue` skips the `fmt.Print` but `i` was already incremented.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `1 2 4 5`. The key here is understanding when `continue` takes effect. The `i++` happens before the `if i == 3` check, so when `i` becomes `3`, we increment it first, then check if it equals `3`. Since it does, we `continue` to the next iteration, skipping the `fmt.Print`. But `i` was already incremented to `4`, so the next iteration prints `4`, then `5`.

---

### 27. Nested Switch Break
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        switch i {
        case 1:
            fmt.Println("one")
            break // breaks switch, not for
        }
        fmt.Println("after switch:", i)
    }
}
```
**A:**
```
after switch: 0
one
after switch: 1
after switch: 2
```
`break` inside `switch` exits only the switch, not the surrounding `for`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints the three lines showing `after switch: 0`, then `one` and `after switch: 1`, then `after switch: 2`. The important point is that `break` inside a `switch` only breaks out of the switch, not the surrounding `for` loop. So when `i == 1`, we print "one", break the switch, but continue with the for loop and print "after switch: 1".

---

### 28. Range Index on Map
**Q: Is this valid and what does it print?**
```go
package main
import "fmt"

func main() {
    m := map[string]int{"a": 1, "b": 2}
    for k := range m {
        fmt.Println(k)
    }
}
```
**A:** **Valid.** Prints `a` and `b` (order not guaranteed). When ranging a map, a single variable gets the key.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Is this valid and what does it print?
**Your Response:** Yes, this is valid code. It prints `a` and `b` (though the order isn't guaranteed since map iteration order is random). When you range over a map with a single variable, you get just the keys. If you used two variables (`for k, v := range m`), you'd get both keys and values.

---

### 29. for init; ; post — Missing Condition
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    for i := 0; ; i++ {
        if i == 3 {
            break
        }
        fmt.Print(i)
    }
}
```
**A:** `012`. Omitting the condition in a three-part `for` creates an infinite loop equivalent to `for true`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `012`. When you omit the condition in a three-part for loop (`for i := 0; ; i++`), it creates an infinite loop equivalent to `for true`. The only way out is the explicit `break` when `i == 3`. So it prints `0`, `1`, `2`, then when `i` becomes `3`, it breaks out of the loop.

---

### 30. Defer Runs After Return Value Set
**Q: What does `run()` return?**
```go
package main
import "fmt"

func run() int {
    x := 0
    defer func() { x = 100 }()
    return x
}

func main() {
    fmt.Println(run())
}
```
**A:** `0`. Unlike named returns, the deferred closure modifies a local copy `x`. The `return x` already captured `0`. Contrast with a named return where the defer CAN modify the result.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does `run()` return?
**Your Response:** This returns `0`. The key insight is how defer works with regular vs named return values. Here we have a regular return, so `return x` captures the current value of `x` (which is `0`) and stores it for return. The deferred function runs after this capture, but it's modifying the local variable `x`, not the return value that was already captured. If we used a named return like `func run() (result int)`, the defer could modify the result.

---

## Section 3: Functions & Closures Deep Dives (Q31–Q45)

### 31. Method Value vs Method Expression
**Q: What is the output?**
```go
package main
import "fmt"

type Adder struct{ base int }
func (a Adder) Add(n int) int { return a.base + n }

func main() {
    adder := Adder{base: 10}

    // Method value: bound to adder instance
    addFn := adder.Add
    fmt.Println(addFn(5))

    // Method expression: unbound, receiver passed explicitly
    addExpr := Adder.Add
    fmt.Println(addExpr(adder, 5))
}
```
**A:**
```
15
15
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and what's the difference?
**Your Response:** This prints `15` twice. The difference is between method values and method expressions. `addFn := adder.Add` creates a method value - it's bound to the specific `adder` instance, so you just pass the missing parameter. `addExpr := Adder.Add` creates a method expression - it's unbound, so you must pass both the receiver and the parameter. Both approaches call the same method but with different syntax.

---

### 32. Passing Variadic to Variadic
**Q: What is the output?**
```go
package main
import "fmt"

func inner(nums ...int) {
    fmt.Println(nums)
}

func outer(nums ...int) {
    inner(nums...)
}

func main() {
    outer(1, 2, 3)
}
```
**A:** `[1 2 3]`. `nums...` unpacks the slice and passes it directly without creating a new slice.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `[1 2 3]`. When you have a variadic function and want to pass its variadic parameter to another variadic function, you use the `...` syntax to unpack the slice. `inner(nums...)` takes the slice `nums` and unpacks it so that each element becomes a separate argument to `inner`. This is efficient because it doesn't create a new slice - it just passes the elements directly.

---

### 33. Defer Captures Value vs Reference
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x := 1
    defer func() { fmt.Println("ref:", x) }()  // captures by reference
    defer fmt.Println("val:", x)                  // arg evaluated NOW (value)
    x = 99
}
```
**A:**
```
val: 1
ref: 99
```
The second `defer` line evaluates `x` immediately (value=1). The first closure sees the final value of `x` (99) when it runs.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why?
**Your Response:** This prints `val: 1` then `ref: 99`. The key difference is how defer handles arguments vs closures. When you defer a function call like `defer fmt.Println("val:", x)`, the arguments are evaluated immediately, so `x` is captured as `1`. But when you defer a closure like `defer func() { fmt.Println("ref:", x) }()`, the closure captures the variable by reference, so it sees the final value of `x` (99) when it executes.

---

### 34. Recursive Anonymous Function
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var fact func(n int) int
    fact = func(n int) int {
        if n == 0 {
            return 1
        }
        return n * fact(n-1)
    }
    fmt.Println(fact(5))
}
```
**A:** `120`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `120`. This is a recursive anonymous function that calculates factorial. We declare a variable `fact` of function type, then assign a recursive function to it. The recursion works because after the assignment, `fact` refers to itself. This pattern is useful when you need a recursive function but don't want to give it a name in the global scope.

---

### 35. Functions Returning Functions (Currying)
**Q: What is the output?**
```go
package main
import "fmt"

func multiply(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    triple := multiply(3)
    fmt.Println(triple(4))
    fmt.Println(triple(10))
}
```
**A:**
```
12
30
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `12` then `30`. This demonstrates function currying - a function that returns another function. `multiply(3)` returns a new function that multiplies its argument by 3. So `triple` becomes a function that multiplies by 3. When we call `triple(4)`, we get `12`, and `triple(10)` gives `30`. This is useful for creating specialized functions from more general ones.

---

### 36. Function Signature Mismatch
**Q: Does this compile?**
```go
package main
import "fmt"

func greet(name string) string {
    return "Hello, " + name
}

func main() {
    var f func(string) int = greet
    fmt.Println(f("Go"))
}
```
**A:** **Compile Error.** `greet` returns `string` but `f` expects a function returning `int`. Function types must match exactly.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. The issue is a function signature mismatch. `greet` has signature `func(string) string`, but variable `f` is declared as `func(string) int`. Go's type system is strict - function types must match exactly, including all parameter types and return types. You can't assign a function returning `string` to a variable expecting a function returning `int`.

---

### 37. Multiple Defers and Named Return Interaction
**Q: What does `g()` return?**
```go
package main
import "fmt"

func g() (result int) {
    defer func() { result++ }()
    defer func() { result += 10 }()
    return 1
}

func main() {
    fmt.Println(g())
}
```
**A:** `12`. `return 1` sets `result = 1`. Defers run in LIFO: first `result += 10` → `11`, then `result++` → `12`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does `g()` return?
**Your Response:** This returns `12`. With named return values, `return 1` sets `result = 1`. The deferred functions run in LIFO order (Last In, First Out). So the second defer (`result += 10`) runs first, making `result = 11`, then the first defer (`result++`) runs, making it `12`. This shows how defer can modify named return values.

---

### 38. Panicking in a Called Function
**Q: What is the output?**
```go
package main
import "fmt"

func risky() {
    panic("boom")
}

func safe() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered:", r)
        }
    }()
    risky()
    fmt.Println("after risky") // unreachable
}

func main() {
    safe()
    fmt.Println("main continues")
}
```
**A:**
```
recovered: boom
main continues
```
`recover()` stops the panic, `safe()` returns normally, and `main` continues.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `recovered: boom` then `main continues`. The `risky()` function panics, but the `safe()` function has a deferred function that calls `recover()`. `recover()` catches the panic and returns its value (`"boom"`), stopping the panic from propagating. So `safe()` returns normally instead of crashing, and `main` continues execution.

---

### 39. init() Function
**Q: What is the output?**
```go
package main
import "fmt"

var x int

func init() {
    x = 42
    fmt.Println("init ran, x =", x)
}

func main() {
    fmt.Println("main ran, x =", x)
}
```
**A:**
```
init ran, x = 42
main ran, x = 42
```
`init()` always runs before `main()`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `init ran, x = 42` then `main ran, x = 42`. In Go, `init()` functions always run before `main()`. The `init()` function sets `x = 42` and prints it, then when `main()` runs, `x` already has the value `42`. This is guaranteed ordering by the Go runtime - all init functions complete before main begins.

---

### 40. Multiple init() Functions in One File
**Q: Does this compile and what is the output?**
```go
package main
import "fmt"

func init() { fmt.Println("init 1") }
func init() { fmt.Println("init 2") }

func main() { fmt.Println("main") }
```
**A:** **Yes, compiles.** Output:
```
init 1
init 2
main
```
A package can have multiple `init()` functions; they run in source order.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and what does it print?
**Your Response:** Yes, this compiles and prints `init 1`, `init 2`, then `main`. Go allows multiple `init()` functions in the same package, and they run in source code order (top to bottom). This is useful for complex initialization that you want to split into logical pieces. All init functions must complete before main starts.

---

### 41. Blank Identifier to Suppress Unused Import
**Q: Does this compile?**
```go
package main
import _ "fmt"

func main() {}
```
**A:** **Yes.** The blank import `_ "fmt"` imports the package solely for its side effects (running `init()`). The package's exported names are not accessible.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** Yes, this compiles. The blank import `_ "fmt"` imports the package only for its side effects, typically running its `init()` functions. You can't access the package's exported names because of the blank identifier, but the package is still initialized. This is commonly used for database drivers, encoding packages, or any package that registers itself in an `init()` function.

---

### 42. Named Return with Early Non-Naked Return
**Q: What is printed?**
```go
package main
import "fmt"

func compute(x int) (result int) {
    if x < 0 {
        return -1 // explicit return, ignores named result
    }
    result = x * x
    return // naked return
}

func main() {
    fmt.Println(compute(-5))
    fmt.Println(compute(4))
}
```
**A:**
```
-1
16
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `-1` then `16`. The key insight is how named returns work with explicit vs naked returns. When `x < 0`, we use `return -1` which is an explicit return - it ignores the named `result` variable and returns `-1` directly. When `x >= 0`, we set `result = x * x` and use a naked `return`, which returns the current value of the named variable `result`.

---

### 43. Closure Over Loop — Range Version
**Q: What is the output (pre Go 1.22)?**
```go
package main
import "fmt"

func main() {
    fns := make([]func(), 0)
    for _, v := range []string{"a", "b", "c"} {
        fns = append(fns, func() { fmt.Print(v) })
    }
    for _, f := range fns {
        f()
    }
}
```
**A:** `ccc`. All closures share the single loop variable `v`, which is `c` after the loop ends.  
**Fix (pre 1.22):** `v := v` inside the loop body.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print (pre Go 1.22)?
**Your Response:** This prints `ccc`. This is the classic closure capture problem with range loops. All closures capture the same variable `v` by reference. After the loop finishes, `v` has the value `c` (the last element), so all closures print `c` when executed. The fix is to create a local copy: `v := v` inside the loop so each closure captures its own value. Note: Go 1.22 fixed this issue.

---

### 44. Panic with Non-String Value
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    defer func() {
        r := recover()
        fmt.Printf("%T: %v\n", r, r)
    }()
    panic(42)
}
```
**A:** `int: 42`. `panic` can take any value, not just strings or errors.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `int: 42`. Many people think `panic` only accepts strings or errors, but it can take any value. Here we panic with an integer `42`. The `recover()` function returns whatever value was passed to `panic`, so we get back the integer `42`. The `%T` format specifier shows it's type `int`, and `%v` shows its value `42`.

---

### 45. init Cannot Be Called Explicitly
**Q: Does this compile?**
```go
package main
import "fmt"

func init() { fmt.Println("init") }

func main() {
    init()
}
```
**A:** **Compile Error.** `init` functions cannot be called explicitly; they are invoked automatically by the runtime.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. `init` functions are special - they can only be called automatically by the Go runtime, not explicitly by your code. The runtime ensures all `init` functions run before `main()` starts. This restriction prevents accidental multiple calls to initialization code which could cause problems.

---

## Section 4: Pointers Deep Dives (Q46–Q56)

### 46. Struct Field via Pointer Without Explicit Dereference
**Q: What is the output?**
```go
package main
import "fmt"

type Rect struct{ W, H int }

func area(r *Rect) int { return r.W * r.H }

func main() {
    r := Rect{3, 4}
    fmt.Println(area(&r))
}
```
**A:** `12`. Go auto-dereferences pointer receivers/arguments when accessing fields.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `12`. Go automatically dereferences pointers when accessing struct fields. Even though `area` expects a `*Rect` pointer and we pass `&r`, inside the function we can access `r.W` and `r.H` directly without writing `(*r).W`. Go handles the dereference for us, making the code cleaner and more readable.

---

### 47. Pointer to Interface Is Almost Always Wrong
**Q: Does this produce expected behavior?**
```go
package main
import "fmt"

type Stringer interface{ String() string }

type T struct{ Name string }
func (t T) String() string { return t.Name }

func printIt(s *Stringer) {
    fmt.Println((*s).String())
}

func main() {
    var s Stringer = T{"Hello"}
    printIt(&s)
}
```
**A:** Prints `Hello` but passing `*Stringer` is almost always a design mistake. Interfaces should be passed by value; they already contain an internal pointer.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this produce expected behavior?
**Your Response:** This prints `Hello`, but passing `*Stringer` is almost always a design mistake. Interfaces in Go are already reference types - they contain a pointer to the underlying value and type information. When you pass a pointer to an interface (`*Stringer`), you're creating a pointer to a pointer, which is unnecessary and confusing. The idiomatic approach is to pass interfaces by value directly.

---

### 48. Nil Pointer in Method Call (on Pointer Receiver)
**Q: What is the output?**
```go
package main
import "fmt"

type Node struct{ val int }

func (n *Node) GetVal() int {
    if n == nil {
        return -1
    }
    return n.val
}

func main() {
    var n *Node
    fmt.Println(n.GetVal())
}
```
**A:** `-1`. Calling a method on a nil pointer is **valid** in Go as long as the method explicitly handles the nil case. The method receives `nil` as the receiver.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `-1`. In Go, calling a method on a nil pointer is valid as long as the method handles the nil case. The method receives `nil` as its receiver, and our code checks `if n == nil` and returns `-1`. This is actually a useful pattern - it allows methods to be called safely on nil values, similar to how some standard library methods work.

---

### 49. Unsafe Pointer Arithmetic (Concept)
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "unsafe"
)

func main() {
    arr := [3]int{10, 20, 30}
    p := unsafe.Pointer(&arr[0])
    p2 := unsafe.Pointer(uintptr(p) + unsafe.Sizeof(arr[0]))
    second := *(*int)(p2)
    fmt.Println(second)
}
```
**A:** `20`. Using `unsafe.Pointer` and `uintptr` to manually advance through array memory. This is fragile and unsafe — only the combination of uintptr+Pointer in a single expression is safe from GC movement.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `20`. This uses unsafe pointer arithmetic to manually access the second element of the array. We convert `&arr[0]` to `unsafe.Pointer`, then to `uintptr`, add the size of an int, then convert back to `int*` and dereference. This is extremely dangerous and should never be used in production code - the garbage collector can move memory, making your pointers invalid. Only use unsafe for very specific, well-understood cases.

---

### 50. Pointer Receiver on Value Type Variable
**Q: Does this compile?**
```go
package main
import "fmt"

type Counter struct{ n int }
func (c *Counter) Inc() { c.n++ }

func main() {
    c := Counter{}
    c.Inc()        // addressable — Go auto-takes &c
    fmt.Println(c.n)
}
```
**A:** **Yes, compiles and prints** `1`. When a variable is addressable, Go automatically takes its address for a pointer-receiver method call.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** Yes, this compiles and prints `1`. Even though `Inc()` has a pointer receiver (`*Counter`), we can call it on a value variable (`c`). Go automatically takes the address of addressable variables when calling pointer-receiver methods. This is a convenience - you don't need to write `(&c).Inc()`. This works because `c` is a variable, not a temporary value.

---

### 51. Map Is Already a Pointer Under the Hood
**Q: What is the output?**
```go
package main
import "fmt"

func addKey(m map[string]int) {
    m["x"] = 99
}

func main() {
    m := map[string]int{}
    addKey(m)
    fmt.Println(m["x"])
}
```
**A:** `99`. Maps are reference types (internally a pointer to a hash table). Passing a map to a function shares the same underlying data.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `99`. Maps are reference types in Go - they're internally pointers to a hash table structure. When you pass a map to a function, you're passing a copy of that pointer, so both the caller and callee are looking at the same underlying data. That's why `addKey` can modify the map and the main function sees the change. This is different from slices, where passing by value doesn't update the caller's slice header.

---

### 52. Function Modifying Slice Header vs Contents
**Q: What is the output?**
```go
package main
import "fmt"

func appendElem(s []int) {
    s = append(s, 99) // modifies local header only
}

func modifyElem(s []int) {
    s[0] = 99 // modifies shared underlying array
}

func main() {
    a := []int{1, 2, 3}
    appendElem(a)
    fmt.Println(a)

    modifyElem(a)
    fmt.Println(a)
}
```
**A:**
```
[1 2 3]
[99 2 3]
```
`append` may allocate a new array; the caller's slice header is not updated. Direct element modification affects shared memory.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `[1 2 3]` then `[99 2 3]`. This shows the difference between modifying slice contents vs the slice header. In `appendElem`, `append` might allocate a new array, but we're only modifying the local slice header, so the caller doesn't see the change. In `modifyElem`, we're directly modifying the shared underlying array, so the caller sees the change. This is why you need to pass a pointer to slice if you want to modify its length/capacity.

---

### 53. Pointer to Slice Can Update Header
**Q: What is the output?**
```go
package main
import "fmt"

func grow(s *[]int) {
    *s = append(*s, 42)
}

func main() {
    s := []int{1, 2}
    grow(&s)
    fmt.Println(s)
}
```
**A:** `[1 2 42]`. Passing `*[]int` lets the function update the caller's slice header (len, cap, ptr).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `[1 2 42]`. By passing a pointer to the slice (`*[]int`), the `grow` function can modify the caller's slice header. The slice header contains the pointer, length, and capacity. When we do `*s = append(*s, 42)`, we're updating the caller's slice header to point to the new array (if reallocation happened). This is the correct way to write functions that modify slices.

---

### 54. Stack vs Heap — Escape Analysis
**Q: Which variable escapes to the heap?**
```go
package main
import "fmt"

func stackVar() int {
    x := 42
    return x // x stays on stack; value is copied
}

func heapVar() *int {
    x := 42
    return &x // x escapes to heap; its address is returned
}

func main() {
    fmt.Println(stackVar())
    fmt.Println(*heapVar())
}
```
**A:** `42 42`. `x` in `heapVar` escapes to the heap because its address is returned. This is safe; Go's GC manages it.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which variable escapes to the heap?
**Your Response:** `x` in `heapVar` escapes to the heap. In `stackVar`, `x` stays on the stack because we only return its value. In `heapVar`, we return the address of `x`, so the compiler must move `x` to the heap to keep it alive after the function returns. This is called escape analysis - Go automatically determines which variables need to be on the heap vs stack. You don't need to manage this manually.

---

### 55. Struct Copy vs Pointer Copy
**Q: What is the output?**
```go
package main
import "fmt"

type Config struct{ Debug bool }

func disable(c Config) {
    c.Debug = false
}

func disablePtr(c *Config) {
    c.Debug = false
}

func main() {
    cfg := Config{Debug: true}
    disable(cfg)
    fmt.Println(cfg.Debug)
    disablePtr(&cfg)
    fmt.Println(cfg.Debug)
}
```
**A:**
```
true
false
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `true` then `false`. When we pass `cfg` by value to `disable`, we're making a copy of the struct, so changes to the copy don't affect the original. When we pass a pointer `&cfg` to `disablePtr`, we're modifying the original struct. This shows the fundamental difference between value and pointer semantics - pass by value for immutability, pass by pointer for mutability.

---

### 56. Nil Map Read Is Safe, Write Is Not
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m["missing"]) // read from nil map: returns zero value
    m["key"] = 1              // write to nil map: PANIC
}
```
**A:** `0` then **panic: assignment to entry in nil map**. Reading from a nil map is safe; writing panics.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `0` then panics. Reading from a nil map is safe - it returns the zero value for the key type. But writing to a nil map causes a panic because there's no underlying hash table to write to. This is a common gotcha - always initialize maps with `make()` before writing to them. You can check if a map is nil before writing, but it's better to ensure proper initialization.

---

## Section 5: Strings & Runes Deep Dives (Q57–Q67)

### 57. Iterating Bytes vs Runes
**Q: What is the difference in output?**
```go
package main
import "fmt"

func main() {
    s := "Go😀"
    fmt.Println("byte len:", len(s))
    fmt.Println("rune len:", len([]rune(s)))

    for i := 0; i < len(s); i++ {
        fmt.Printf("%d:%x ", i, s[i])
    }
}
```
**A:** `byte len: 6`, `rune len: 3`. The emoji 😀 is 4 bytes. Byte iteration: `0:47 1:6f 2:f0 3:9f 4:98 5:80`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and what's the difference between byte and rune iteration?
**Your Response:** This prints the byte length as 6 and rune length as 3, then shows the hex values of each byte. The key difference is that `len()` on a string gives bytes, while `len([]rune(s))` gives actual Unicode characters. The emoji takes 4 bytes but is 1 rune. When iterating bytes, you see the raw UTF-8 encoding, while iterating runes gives you the actual characters. This is important when working with Unicode text - always use runes for character-level operations.

---

### 58. strings.Builder vs bytes.Buffer
**Q: What is printed and what is the key difference in design?**
```go
package main
import (
    "bytes"
    "fmt"
    "strings"
)

func main() {
    var sb strings.Builder
    sb.WriteString("Go")
    sb.WriteByte('!')
    fmt.Println(sb.String())

    var bb bytes.Buffer
    bb.WriteString("Go")
    bb.WriteByte('!')
    fmt.Println(bb.String())
}
```
**A:** Both print `Go!`. Key difference: `strings.Builder` is write-only and optimized for building strings (no read methods). `bytes.Buffer` supports both reading and writing.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the output and what's the key difference between these two?
**Your Response:** Both print `Go!`, but they serve different purposes. `strings.Builder` is specifically designed for building strings efficiently - it's write-only and optimized for that use case. `bytes.Buffer` is more general-purpose - it can both read and write, making it useful when you need to manipulate the buffer contents. Use `strings.Builder` when you're just concatenating strings, and `bytes.Buffer` when you need to read from the buffer as well.

---

### 59. Multi-line Raw String Literal
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := `line1
line2
line3`
    fmt.Println(len(s))
}
```
**A:** `17`. Raw string literals (backtick) include literal newlines (`\n`) as-is. `line1\nline2\nline3` = 5+1+5+1+5 = 17 bytes.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why?
**Your Response:** This prints 17. Raw string literals in Go use backticks and include everything exactly as written, including newlines. So the string contains `line1` followed by a literal newline, then `line2`, another newline, and `line3`. Each line is 5 characters plus 1 newline character, totaling 17 bytes. This is different from regular string literals where you'd need to escape newlines with `\n`.

---

### 60. strings.Split Behavior
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    parts := strings.Split("a,,b", ",")
    fmt.Println(len(parts), parts)
}
```
**A:** `3 [a  b]`. `Split` keeps the empty string between consecutive delimiters.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why?
**Your Response:** This prints `3 [a  b]`. The `strings.Split` function preserves empty strings between consecutive delimiters. When you split `"a,,b"` by comma, you get three parts: the first `"a"`, then an empty string from the two consecutive commas, then `"b"`. This behavior is important to remember - if you want to ignore empty strings, you need to filter them out yourself.

---

### 61. String Formatting %v vs %s vs %q
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := "Go\nLang"
    fmt.Printf("%v\n", s)
    fmt.Printf("%s\n", s)
    fmt.Printf("%q\n", s)
}
```
**A:**
```
Go
Lang
Go
Lang
"Go\nLang"
```
`%q` quotes the string and escapes special characters.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the output and what's the difference between these format verbs?
**Your Response:** The output shows the string printed with three different format verbs. `%v` prints the default format, which for strings is just the string content. `%s` also prints the string content. `%q` quotes the string and escapes special characters - so it wraps it in quotes and converts `\n` to the literal `\\n`. This is useful when you need to see the exact string representation, including escape sequences.

---

### 62. Converting int to string Pitfall
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    n := 65
    fmt.Println(string(n))           // rune conversion
    fmt.Println(strconv.Itoa(n))     // integer to string
}
```
**A:**
```
A
65
```
`string(n)` treats `n` as a Unicode code point. Use `strconv.Itoa` for numeric string conversion.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and what's the common mistake here?
**Your Response:** This prints `A` then `65`. The first line `string(65)` treats 65 as a Unicode code point, which corresponds to the character 'A'. This is a common pitfall - developers expect it to convert the number to the string "65". The correct way to convert an integer to its string representation is `strconv.Itoa()`. This is important because `string()` conversion only works for Unicode code points, not general numeric conversion.

---

### 63. strings.TrimSpace
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    s := "  hello world  "
    fmt.Printf("[%s]\n", strings.TrimSpace(s))
}
```
**A:** `[hello world]`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `[hello world]`. The `strings.TrimSpace` function removes leading and trailing whitespace characters from a string. It handles spaces, tabs, newlines, and other whitespace characters. This is commonly used when cleaning up user input or parsing data where extra whitespace might be present.

---

### 64. Byte Mutation via []byte
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := "hello"
    b := []byte(s)
    b[0] = 'H'
    fmt.Println(string(b))
    fmt.Println(s) // original unchanged
}
```
**A:**
```
Hello
hello
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why doesn't the original string change?
**Your Response:** This prints `Hello` then `hello`. The key concept here is that strings in Go are immutable. When we convert the string to a byte slice and modify the first byte, we're working on a copy. The original string remains unchanged. This demonstrates Go's string immutability - if you need to modify a string, you must create a new one. This is actually a good thing for performance and safety in concurrent programs.

---

### 65. strings.Replace vs strings.ReplaceAll
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    s := "aabbaa"
    fmt.Println(strings.Replace(s, "a", "x", 2))    // replace first 2
    fmt.Println(strings.ReplaceAll(s, "a", "x"))     // replace all
}
```
**A:**
```
xxbbaa
xxbbxx
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and what's the difference between these two functions?
**Your Response:** This prints `xxbbaa` then `xxbbxx`. `strings.Replace` replaces only the first N occurrences - here we specify 2, so only the first two 'a's are replaced. `strings.ReplaceAll` replaces all occurrences. This is important to remember - if you want to replace all instances, use `ReplaceAll`. If you want more control, use `Replace` with the count parameter.

---

### 66. strings.Fields vs strings.Split
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    s := "  foo   bar  baz  "
    fmt.Println(strings.Fields(s))
    fmt.Println(strings.Split(s, " "))
}
```
**A:**
```
[foo bar baz]
[ foo   bar  baz  ]
```
`Fields` splits by whitespace and ignores leading/trailing/multiple spaces. `Split` is literal.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between Fields and Split here?
**Your Response:** The output shows `Fields` produces `[foo bar baz]` while `Split` produces `[ foo   bar  baz  ]`. `Fields` splits on any whitespace and automatically trims and collapses multiple spaces - it's ideal for parsing text where you want words. `Split` splits on the exact delimiter provided, so it preserves empty strings and extra spaces. Use `Fields` when you want words from natural text, and `Split` when you need precise control over the splitting behavior.

---

### 67. Rune Arithmetic
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    r := 'A'
    for i := 0; i < 5; i++ {
        fmt.Printf("%c", r+rune(i))
    }
}
```
**A:** `ABCDE`. Rune arithmetic works on Unicode code points.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and how does it work?
**Your Response:** This prints `ABCDE`. The code starts with rune 'A' and adds 0, 1, 2, 3, 4 to it in the loop. Since runes are just Unicode code points (which are integers), we can do arithmetic on them. 'A' + 0 = 'A', 'A' + 1 = 'B', and so on. This works because the Unicode code points for uppercase letters are consecutive. This technique is useful for character manipulation, but be careful - it only works reliably for consecutive character sets like ASCII letters.

---

## Section 6: Slices & Maps Deep Dives (Q68–Q83)

### 68. Three-Index Slice (Full Slice Expression)
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5}
    b := a[1:3:4] // len=2, cap=3
    fmt.Println(len(b), cap(b))
}
```
**A:** `2 3`. The three-index form `a[low:high:max]` sets `cap = max - low`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and what does the third index do?
**Your Response:** This prints `2 3`. The three-index slice expression `a[1:3:4]` creates a slice with length 2 (3-1) and capacity 3 (4-1). The third index controls the capacity of the new slice, which is useful when you want to limit how much the slice can grow before triggering a reallocation. This prevents accidental growth and can be a performance optimization when you know the maximum size needed.

---

### 69. append Does Not Modify Original If Capacity Exceeded
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := make([]int, 3, 3)
    a[0], a[1], a[2] = 1, 2, 3
    b := append(a, 4) // triggers reallocation
    b[0] = 99
    fmt.Println(a[0]) // original unaffected
    fmt.Println(b[0])
}
```
**A:**
```
1
99
```
When `append` exceeds capacity, it creates a new backing array. `a` and `b` no longer share memory.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why doesn't changing b affect a?
**Your Response:** This prints `1` then `99`. The key concept is that when `append` exceeds a slice's capacity, it allocates a new backing array. Initially, both slices share the same underlying array, but when we append to `b` and exceed its capacity of 3, Go creates a new array. After that, `a` and `b` point to different arrays, so modifying `b[0]` doesn't affect `a[0]`. This is crucial for understanding slice behavior and avoiding unexpected side effects.

---

### 70. Slice Tricks: Insert at Index
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := []int{1, 2, 4, 5}
    i := 2
    s = append(s[:i+1], s[i:]...)
    s[i] = 3
    fmt.Println(s)
}
```
**A:** `[1 2 3 4 5]`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and how does this insertion trick work?
**Your Response:** This prints `[1 2 3 4 5]`. This is a clever slice insertion trick. The line `s = append(s[:i+1], s[i:]...)` effectively makes room for a new element at index `i`. It works by taking the slice up to and including position `i`, then appending everything from position `i` onwards. This shifts all elements from position `i` one step to the right, creating space. Then we set `s[i] = 3`. This is more efficient than creating a new slice and copying elements individually.

---

### 71. Map with Struct Values (Non-Addressable)
**Q: Does this compile?**
```go
package main

type Point struct{ X, Y int }

func main() {
    m := map[string]Point{"a": {1, 2}}
    m["a"].X = 10 // cannot assign to struct field in a map
}
```
**A:** **Compile Error.** Struct values stored in maps are **not addressable**. You must retrieve the whole struct, modify, and put it back:
```go
p := m["a"]; p.X = 10; m["a"] = p
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and why?
**Your Response:** This doesn't compile - it gives a compile error. The issue is that struct values stored in maps are not addressable. When you access `m["a"]`, you get a copy of the struct, not a reference to it. Go doesn't allow you to take the address of a map element because the underlying hash table might need to relocate elements. To modify a struct in a map, you must retrieve the whole struct, modify it, and assign it back to the map.

---

### 72. Slice of Pointers vs Slice of Values
**Q: What is the output?**
```go
package main
import "fmt"

type Item struct{ V int }

func main() {
    items := []Item{{1}, {2}, {3}}
    for i := range items {
        items[i].V *= 10
    }
    for _, it := range items {
        fmt.Print(it.V, " ")
    }
}
```
**A:** `10 20 30 `. Using `items[i]` gives a pointer to the element; modifying via range value (`it`) would fail since `it` is a copy.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why does this work?
**Your Response:** This prints `10 20 30 `. The key is using `items[i]` instead of the range variable `it`. When you use `range items`, the variable `it` is a copy of each element, so modifying it wouldn't affect the original slice. But `items[i]` directly accesses the element in the slice, so `items[i].V *= 10` modifies the actual element. This demonstrates the difference between range values (copies) and direct slice access (references to elements).

---

### 73. Map Concurrent Access
**Q: What is the bug?**
```go
package main
import "sync"

func main() {
    m := map[int]int{}
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            m[n] = n // DATA RACE
        }(i)
    }
    wg.Wait()
}
```
**A:** **Data race.** Concurrent writes to a map without synchronization cause undefined behavior and will be detected by `-race`. Fix: use `sync.Mutex` or `sync.Map`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the bug in this code?
**Your Response:** This has a data race bug. Multiple goroutines are writing to the same map concurrently without any synchronization. Maps in Go are not safe for concurrent access - this can cause memory corruption, crashes, or incorrect results. The race detector would catch this. To fix it, you need to either use a `sync.Mutex` to protect map access, or use `sync.Map` which is designed for concurrent use. This is a fundamental Go concurrency concept.

---

### 74. Delete During Iteration
**Q: Is this safe?**
```go
package main
import "fmt"

func main() {
    m := map[int]string{1: "a", 2: "b", 3: "c"}
    for k := range m {
        if k == 2 {
            delete(m, k)
        }
    }
    fmt.Println(len(m))
}
```
**A:** **Yes, safe.** It is safe to delete map keys during a `range` loop. Output: `2`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Is it safe to delete map keys while iterating?
**Your Response:** Yes, this is safe in Go. It's safe to delete map keys during a range loop. The iteration continues normally and deleted keys won't be visited if they haven't been visited yet. The output is 2 because we delete the key with value 2. This is different from some languages where modifying a collection during iteration causes issues. Go's map iteration is designed to handle deletions safely, though adding new keys during iteration might not show up in the current iteration.

---

### 75. make vs literal for Map
**Q: What difference exists between these two maps?**
```go
package main
import "fmt"

func main() {
    m1 := make(map[string]int)
    m2 := map[string]int{}
    fmt.Println(m1 == nil, m2 == nil)
}
```
**A:** `false false`. Both are initialized (non-nil) maps. `make` with a capacity hint is preferred for performance when size is known in advance.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between these two maps?
**Your Response:** Both print `false false` - neither map is nil. Both create initialized, ready-to-use maps. The difference is mostly stylistic and performance-related. `make(map[string]int)` creates an empty map, while the literal `map[string]int{}` also creates an empty map. The advantage of `make` is that you can provide a capacity hint like `make(map[string]int, 100)` to pre-allocate space, which improves performance when you know approximately how many entries you'll have.

---

### 76. Slice Contains (Linear Search)
**Q: What is the output?**
```go
package main
import "fmt"

func contains(s []string, target string) bool {
    for _, v := range s {
        if v == target {
            return true
        }
    }
    return false
}

func main() {
    fmt.Println(contains([]string{"a", "b", "c"}, "b"))
    fmt.Println(contains([]string{"a", "b", "c"}, "z"))
}
```
**A:**
```
true
false
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `true` then `false`. The `contains` function does a linear search through the slice to find the target string. It returns `true` for "b" because it's in the slice, and `false` for "z" because it's not. This is the standard way to check if a slice contains an element in Go. For better performance with many lookups, you'd typically convert the slice to a map for O(1) lookup time instead of O(n) linear search.

---

### 77. Copy Between Different Length Slices
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    src := []int{1, 2, 3, 4, 5}
    dst := make([]int, 3)
    n := copy(dst, src)
    fmt.Println(n, dst)
}
```
**A:** `3 [1 2 3]`. `copy` copies `min(len(dst), len(src))` elements.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and how does copy work?
**Your Response:** This prints `3 [1 2 3]`. The `copy` function copies elements from the source slice to the destination slice. It copies the minimum of the two slice lengths - here the destination has length 3 and source has length 5, so only 3 elements are copied. The function returns the number of elements copied. This is a built-in function that's efficient and handles overlapping slices correctly.

---

### 78. Map of Slices: Append Pattern
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    m := map[string][]int{}
    for i := 0; i < 3; i++ {
        m["nums"] = append(m["nums"], i)
    }
    fmt.Println(m["nums"])
}
```
**A:** `[0 1 2]`. The zero value for a missing map key is `nil` slice, and appending to `nil` is valid.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why does this work?
**Your Response:** This prints `[0 1 2]`. The key insight is that accessing a missing map key returns the zero value for that type. For slices, the zero value is `nil`. The magic is that appending to a `nil` slice in Go is perfectly valid - it automatically allocates a new slice. This pattern is commonly used for building collections in maps without needing to explicitly initialize each slice.

---

### 79. Slice of nil vs Empty Slice in JSON
**Q: What is the output?**
```go
package main
import (
    "encoding/json"
    "fmt"
)

func main() {
    var a []int
    b := []int{}

    ja, _ := json.Marshal(a)
    jb, _ := json.Marshal(b)
    fmt.Println(string(ja))
    fmt.Println(string(jb))
}
```
**A:**
```
null
[]
```
A nil slice marshals to JSON `null`; an empty slice marshals to `[]`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference in JSON output?
**Your Response:** This prints `null` then `[]`. The difference between a nil slice and an empty slice matters when serializing to JSON. A nil slice marshals to `null`, while an empty slice marshals to `[]`. This can be important in APIs where you want to distinguish between "no values" versus "an empty collection". In most cases, you'll want to initialize slices as empty rather than leave them nil to avoid null values in JSON.

---

### 80. Nested Map Initialization
**Q: Does this panic?**
```go
package main
import "fmt"

func main() {
    m := map[string]map[string]int{}
    m["a"]["x"] = 1
    fmt.Println(m)
}
```
**A:** **Panic.** `m["a"]` returns `nil` (the zero value for `map[string]int`). You must initialize the inner map: `m["a"] = map[string]int{}; m["a"]["x"] = 1`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this panic and why?
**Your Response:** Yes, this panics. The issue is that `m["a"]` returns `nil` because the inner map doesn't exist yet. When you try to assign to `m["a"]["x"]`, you're trying to write to a nil map, which causes a panic. You need to initialize the inner map first: `m["a"] = map[string]int{}` and then you can assign values to it. This is a common gotcha with nested maps - you must initialize each level before using it.

---

### 81. Slice as Set (Using map[T]struct{})
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    set := map[string]struct{}{}
    words := []string{"go", "is", "fun", "go"}
    for _, w := range words {
        set[w] = struct{}{}
    }
    fmt.Println(len(set))
}
```
**A:** `3`. Duplicate `"go"` is overwritten. The empty struct uses zero memory as the value type.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why use empty struct?
**Your Response:** This prints `3`. The duplicate "go" is overwritten because map keys must be unique. The interesting part is using `struct{}` as the value type - this is a common Go idiom for implementing sets. The empty struct occupies zero memory, so the map is essentially storing just the keys. This is more memory-efficient than using `bool` or `int` as values when you only need to track existence.

---

### 82. Modifying Slice Backing Array via Two Slices
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    base := make([]int, 5)
    a := base[0:3]
    b := base[2:5]
    a[2] = 99
    fmt.Println(b[0])
}
```
**A:** `99`. `a[2]` and `b[0]` both point to `base[2]` — the same underlying array element.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why?
**Your Response:** This prints `99`. The key concept is that slices are views into an underlying array. Both slice `a` and slice `b` are created from the same `base` array. `a[2]` and `b[0]` both refer to the same element `base[2]`. When we modify `a[2] = 99`, we're changing the underlying array, so `b[0]` also sees that change. This demonstrates how slices share memory and why modifying one slice can affect others that view the same data.

---

### 83. Map Key Types Must Be Comparable
**Q: Does this compile?**
```go
package main

func main() {
    m := map[[]int]string{}
    _ = m
}
```
**A:** **Compile Error.** Slices are not comparable and therefore cannot be used as map keys. Valid key types include: bool, numeric, string, pointer, channel, interface, array, struct (if all fields are comparable).

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and why?
**Your Response:** No, this doesn't compile. Slices are not comparable in Go, so they can't be used as map keys. The reason is that slices are references to underlying arrays, and their equality would be ambiguous - should it compare the pointers or the contents? Only types that have a well-defined equality can be map keys. Valid key types include basic types like strings and numbers, pointers, channels, and structs where all fields are comparable. If you need to use a slice-like key, you could convert it to an array or string first.

---

## Section 7: Structs & Interfaces Deep Dives (Q84–Q95)

### 84. Struct Promoted Method Conflict Resolution
**Q: What is the output?**
```go
package main
import "fmt"

type Base struct{}
func (Base) Describe() string { return "Base" }

type Child struct{ Base }
func (Child) Describe() string { return "Child" }

func main() {
    c := Child{}
    fmt.Println(c.Describe())
    fmt.Println(c.Base.Describe())
}
```
**A:**
```
Child
Base
```
`Child.Describe()` overrides the promoted method. You can still access `Base.Describe` explicitly.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and how does method overriding work with embedding?
**Your Response:** This prints `Child` then `Base`. When you embed a struct, its methods are promoted to the outer struct. However, if the outer struct defines a method with the same name, it overrides the promoted method. So `c.Describe()` calls `Child.Describe()`, but you can still access the base method explicitly with `c.Base.Describe()`. This is Go's way of achieving something similar to inheritance while keeping composition explicit.

---

### 85. Interface with Multiple Types Implementing
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "math"
)

type Shape interface{ Area() float64 }

type Circle struct{ R float64 }
type Rect struct{ W, H float64 }

func (c Circle) Area() float64 { return math.Pi * c.R * c.R }
func (r Rect) Area() float64   { return r.W * r.H }

func printArea(s Shape) { fmt.Printf("%.2f\n", s.Area()) }

func main() {
    printArea(Circle{R: 1})
    printArea(Rect{W: 3, H: 4})
}
```
**A:**
```
3.14
12.00
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and how does this work?
**Your Response:** This prints `3.14` then `12.00`. Both `Circle` and `Rect` satisfy the `Shape` interface because they both have an `Area()` method with the same signature. The `printArea` function accepts any type that implements `Shape`. This demonstrates Go's structural typing - you don't need to explicitly declare that a type implements an interface. If it has the methods, it satisfies the interface. This makes Go interfaces very flexible and decoupled.

---

### 86. Embedding Interface in Struct
**Q: Does this compile and does it satisfy the interface?**
```go
package main
import "fmt"

type Speaker interface{ Speak() string }

type Robot struct{ Speaker }

func main() {
    r := Robot{}
    fmt.Println(r.Speak())
}
```
**A:** **Compiles but panics at runtime.** `Robot.Speaker` is `nil` (zero value). Calling `r.Speak()` is a nil pointer dereference on the embedded interface.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and what happens?
**Your Response:** This compiles fine but panics at runtime. The issue is that when you embed an interface in a struct, the embedded field gets the zero value for that interface type, which is `nil`. When you call `r.Speak()`, you're trying to call a method on a nil interface, which causes a panic. You need to initialize the embedded interface with a concrete type that implements `Speaker`. This is a subtle gotcha with interface embedding.

---

### 87. Empty Interface as Function Parameter
**Q: Can an `interface{}` argument modify the original value?**
```go
package main
import "fmt"

func tryChange(v interface{}) {
    v = 999
}

func main() {
    x := 42
    tryChange(x)
    fmt.Println(x)
}
```
**A:** `42`. `interface{}` wraps the value; rebinding `v` inside the function does not affect the original.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can the function modify the original value?
**Your Response:** No, this prints `42`. The original value is unchanged. When you pass a value to an `interface{}` parameter, Go creates a copy of the value and wraps it in the interface. Inside the function, `v` is a new variable holding that copy. When you assign `v = 999`, you're only changing what `v` points to, not the original value. To modify the original, you'd need to pass a pointer. This demonstrates how interface parameters work with value semantics.

---

### 88. Stringer Interface Auto-Used by fmt
**Q: What is the output?**
```go
package main
import "fmt"

type Color int

const (
    Red Color = iota
    Green
    Blue
)

func (c Color) String() string {
    return []string{"Red", "Green", "Blue"}[c]
}

func main() {
    fmt.Println(Red)
    fmt.Println(Green)
    fmt.Println(Blue)
}
```
**A:**
```
Red
Green
Blue
```
`fmt` package automatically calls `String()` if a type implements the `fmt.Stringer` interface.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why?
**Your Response:** This prints the color names `Red`, `Green`, `Blue`. The `fmt` package automatically looks for and calls the `String()` method if a type implements the `fmt.Stringer` interface. This is a powerful Go feature - you can control how your types are displayed by simply implementing a `String() string` method. This is much cleaner than having to remember to call a specific formatting function everywhere.

---

### 89. Struct Tags Are Metadata Only
**Q: Does modifying a tag affect runtime behavior of non-json usage?**
```go
package main
import (
    "fmt"
    "reflect"
)

type T struct {
    Name string `mytag:"custom"`
}

func main() {
    t := T{Name: "Go"}
    field, _ := reflect.TypeOf(t).FieldByName("Name")
    fmt.Println(field.Tag.Get("mytag"))
    fmt.Println(t.Name)
}
```
**A:**
```
custom
Go
```
Struct tags are only accessible via reflection; they don't affect normal field access.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and what are struct tags used for?
**Your Response:** This prints `custom` then `Go`. Struct tags are metadata that you can attach to struct fields. They don't affect normal field access - that's why `t.Name` still prints `Go`. But you can access them via reflection, as shown with `field.Tag.Get("mytag")`. Tags are commonly used by libraries for things like JSON marshaling, database mapping, validation rules, etc. They provide a way to add extra information to fields without changing the field names or types.

---

### 90. Nil Receiver of Concrete Type
**Q: What is the output?**
```go
package main
import "fmt"

type T struct{ val int }

func (t *T) Print() {
    if t == nil {
        fmt.Println("nil receiver")
        return
    }
    fmt.Println(t.val)
}

func main() {
    var t *T
    t.Print()
}
```
**A:** `nil receiver`. Calling a method on a nil pointer is valid if the method handles it.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and is this safe?
**Your Response:** This prints `nil receiver`. Yes, this is safe in Go! You can call a method on a nil pointer receiver, as long as the method checks for nil before trying to access any fields. The method receives `nil` as its receiver value and can handle it appropriately. This is actually a useful pattern - for example, you might have a method that works on both nil and non-nil receivers, providing different behavior in each case. It's more explicit than having separate methods.

---

### 91. Interface Comparison — Different Dynamic Types
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var a interface{} = 1
    var b interface{} = 1
    var c interface{} = "1"
    fmt.Println(a == b)
    fmt.Println(a == c)
}
```
**A:**
```
true
false
```
Interface values are equal only if both their dynamic type AND dynamic value are equal.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why?
**Your Response:** This prints `true` then `false`. Interface equality in Go requires both the dynamic type and dynamic value to be equal. `a` and `b` both hold the integer `1`, so they're equal. But `c` holds the string `"1"`, which has a different dynamic type, so `a == c` is false even though they look similar. This is important to understand - interface comparison checks both what type is stored AND what value is stored.

---

### 92. Type Switch vs Type Assertion
**Q: What is the output?**
```go
package main
import "fmt"

func check(i interface{}) {
    switch v := i.(type) {
    case int, float64:
        fmt.Printf("numeric: %v\n", v)
    case string:
        fmt.Printf("string: %s\n", v)
    }
}

func main() {
    check(3.14)
    check("hello")
    check(true)
}
```
**A:**
```
numeric: 3.14
string: hello
```
`true` matches no case and falls through silently (no default).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and what about the boolean?
**Your Response:** This prints `numeric: 3.14` then `string: hello`. The boolean `true` doesn't match any case and there's no default case, so it's silently ignored. The type switch matches both `int` and `float64` with the same case using a comma-separated list. This is a clean way to handle multiple types that should be processed the same way. The `v` variable in the type switch has the type of the matching case, which is why we can use it directly in the Printf.

---

### 93. Embedding Promotes Fields, Not Methods When Overridden
**Q: What is the output?**
```go
package main
import "fmt"

type Animal struct{ Name string }
func (a Animal) Sound() string { return "..." }

type Cat struct {
    Animal
    Sound string // field named Sound shadows promoted method
}

func main() {
    c := Cat{Animal{"Whiskers"}, "Meow"}
    fmt.Println(c.Sound)       // field
    fmt.Println(c.Animal.Sound()) // method
}
```
**A:**
```
Meow
...
```
The `Sound` field in `Cat` shadows the promoted `Sound()` method from `Animal`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why?
**Your Response:** This prints `Meow` then `...`. The interesting thing here is that the `Sound` field in the `Cat` struct shadows the promoted `Sound()` method from the embedded `Animal` struct. When you access `c.Sound`, you get the field value, not the method. But you can still access the method explicitly with `c.Animal.Sound()`. This demonstrates how Go's embedding works - fields can shadow promoted methods, and you need to be explicit when you want the method.

---

### 94. Interface Slice Type Assertion
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var i interface{} = []int{1, 2, 3}
    s, ok := i.([]int)
    fmt.Println(s, ok)
}
```
**A:** `[1 2 3] true`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print?
**Your Response:** This prints `[1 2 3] true`. The type assertion `i.([]int)` succeeds because the interface actually holds a slice of integers. The `ok` value is `true` indicating the assertion succeeded. This is the safe way to do type assertions - you get both the value and a boolean indicating success. If the assertion failed, `s` would be `nil` and `ok` would be `false`.

---

### 95. Structural Typing — Duck Typing
**Q: Does this compile?**
```go
package main
import "fmt"

type Walker interface{ Walk() }

type Person struct{}
func (p Person) Walk() { fmt.Println("walking") }

type Robot struct{}
func (r Robot) Walk() { fmt.Println("rolling") }

func makeWalk(w Walker) { w.Walk() }

func main() {
    makeWalk(Person{})
    makeWalk(Robot{})
}
```
**A:** **Yes.** Go uses structural typing — any type with the `Walk()` method satisfies `Walker` without explicitly declaring it.
```
walking
rolling
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and how does this work?
**Your Response:** Yes, this compiles and prints `walking` then `rolling`. This demonstrates Go's structural typing - also called duck typing. Neither `Person` nor `Robot` explicitly implements the `Walker` interface, but they both satisfy it because they have a `Walk()` method with the correct signature. Go doesn't care about the declared type, only about whether the required methods are present. This makes interfaces very flexible and decoupled - you can create new types that work with existing interfaces without modifying those interfaces.

---

## Section 8: Error Handling Deep Dives (Q96–Q100)

### 96. errors.New vs fmt.Errorf
**Q: What is the difference?**
```go
package main
import (
    "errors"
    "fmt"
)

var ErrA = errors.New("error A")

func main() {
    err1 := errors.New("error A") // new sentinel each time
    err2 := fmt.Errorf("context: %w", ErrA) // wraps ErrA

    fmt.Println(err1 == ErrA)           // false: different pointers
    fmt.Println(errors.Is(err2, ErrA))  // true: unwraps to find ErrA
}
```
**A:**
```
false
true
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and what's the difference?
**Your Response:** This prints `false` then `true`. The key difference is that `errors.New("error A")` creates a new error value each time, so `err1` is not equal to `ErrA` even though they have the same message. But `fmt.Errorf` with `%w` wraps `ErrA`, and `errors.Is` can unwrap the error to find the original. This is the modern Go error handling pattern - use sentinel errors for comparison and wrap them with context using `%w`.

---

### 97. Wrapping Errors Multiple Levels
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

var ErrRoot = errors.New("root cause")

func level2() error { return fmt.Errorf("L2: %w", ErrRoot) }
func level1() error { return fmt.Errorf("L1: %w", level2()) }

func main() {
    err := level1()
    fmt.Println(err)
    fmt.Println(errors.Is(err, ErrRoot))
}
```
**A:**
```
L1: L2: root cause
true
```
`errors.Is` unwraps through the chain to find `ErrRoot`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and how does error wrapping work?
**Your Response:** This prints the full error chain `L1: L2: root cause` then `true`. The `fmt.Errorf` with `%w` creates a wrapped error that preserves the original error. `errors.Is` can recursively unwrap through the chain to find if a specific error is present anywhere in the chain. This is powerful for error handling - you can add context at each level while still being able to check for specific root causes.

---

### 98. Panic vs Error Return Convention
**Q: What is the idiomatic Go approach?**
```go
package main
import (
    "errors"
    "fmt"
)

// Idiomatic: return error for expected failures
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(result)
}
```
**A:** `Error: division by zero`. Idiomatic Go: return errors for expected/recoverable failures; use `panic` only for truly unrecoverable programming errors.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the idiomatic Go approach to error handling?
**Your Response:** This prints `Error: division by zero`. The idiomatic Go approach is to return errors for expected failures like division by zero, not panic. Panics should be reserved for truly unrecoverable programming errors. Go's explicit error handling with multiple return values makes error handling clear and mandatory. This is different from exceptions in other languages - Go prefers explicit error checking over implicit error handling.

---

### 99. Error Type Switch
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "os"
)

func main() {
    _, err := os.Open("/nonexistent/file")
    if err != nil {
        switch e := err.(type) {
        case *os.PathError:
            fmt.Println("PathError:", e.Op, e.Path)
        default:
            fmt.Println("other:", e)
        }
    }
}
```
**A:** `PathError: open /nonexistent/file`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and how does error type switching work?
**Your Response:** This prints `PathError: open /nonexistent/file`. The code uses a type switch to check the specific type of error. `os.Open` returns an `*os.PathError` when a file doesn't exist, which contains additional information about the operation and path. Type switching on errors allows you to handle different error types differently - you might retry on network errors but fail immediately on permission errors. This is more precise than just checking error strings.

---

### 100. Sentinel Error Pattern
**Q: What is the output and why is `var` (not `const`) used?**
```go
package main
import (
    "errors"
    "fmt"
)

var (
    ErrNotFound   = errors.New("not found")
    ErrPermission = errors.New("permission denied")
)

func fetch(id int) error {
    if id < 0 {
        return ErrPermission
    }
    if id == 0 {
        return ErrNotFound
    }
    return nil
}

func main() {
    for _, id := range []int{0, -1, 1} {
        err := fetch(id)
        fmt.Printf("id=%d: errors.Is(ErrNotFound)=%v errors.Is(ErrPermission)=%v\n",
            id,
            errors.Is(err, ErrNotFound),
            errors.Is(err, ErrPermission))
    }
}
```
**A:**
```
id=0: errors.Is(ErrNotFound)=true errors.Is(ErrPermission)=false
id=-1: errors.Is(ErrNotFound)=false errors.Is(ErrPermission)=true
id=1: errors.Is(ErrNotFound)=false errors.Is(ErrPermission)=false
```
`var` is used because `errors.New` returns a pointer and **identity** is what uniquely identifies sentinel errors. Constants cannot hold pointer values.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this print and why use var instead of const?
**Your Response:** This shows the results for each ID. The key insight is why we use `var` for sentinel errors instead of `const`. `errors.New` returns a pointer, and pointer identity is what uniquely identifies sentinel errors for comparison with `errors.Is`. Constants can't hold pointer values in Go, so we must use variables. This pattern allows for precise error checking - you can distinguish between "not found" and "permission denied" errors rather than just checking if any error occurred.

---

*End of 100 In-Depth Go Basics & Fundamentals Code Snippet Questions*
