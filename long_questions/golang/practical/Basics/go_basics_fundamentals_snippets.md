# 100 Pure Code Snippet Interview Questions: Go Basics & Fundamentals

*Each question is a "predict the output / spot the bug / does it compile?" style question.*
*Topics: Variables, Types, Control Flow, Functions, Closures, Pointers, Strings, Arrays, Slices, Maps, Structs, Interfaces, Error Handling, Defer, Goroutines (basic), Type System.*

---

## 📋 Reading Progress

> Mark each section `[x]` when done. Use `🔖` to note where you left off.

- [ ] **Section 1:** Variables, Constants & Types (Q1–Q15)
- [ ] **Section 2:** Control Flow (Q16–Q28)
- [ ] **Section 3:** Functions, Closures & Defer (Q29–Q44)
- [ ] **Section 4:** Pointers (Q45–Q52)
- [ ] **Section 5:** Strings & Runes (Q53–Q62)
- [ ] **Section 6:** Arrays, Slices & Maps (Q63–Q78)
- [ ] **Section 7:** Structs & Interfaces (Q79–Q91)
- [ ] **Section 8:** Error Handling (Q92–Q96)
- [ ] **Section 9:** Goroutines Basics & Misc (Q97–Q100)

> 🔖 **Last read:** <!-- e.g. Q15 · Section 1 done -->

---

## Section 1: Variables, Constants & Types (Q1–Q15)

### 1. Short Variable Declaration Outside Function
**Q: Does this compile?**
```go
package main

:= 10
fmt.Println(x)
```
**A:** **Compile Error.** Short variable declaration (`:=`) is not allowed at the package level. Use `var x = 10` or `var x int = 10`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's wrong with this code and why doesn't it compile?
**Your Response:** This code doesn't compile because the short variable declaration operator `:=` can only be used inside functions, not at the package level. At the package level, Go requires explicit variable declaration using the `var` keyword. So instead of `:= 10`, we need to write `var x = 10` or `var x int = 10`. This is a fundamental Go syntax rule that ensures package-level variables are explicitly declared with their types.

---

### 2. Multiple Return Values
**Q: What is the output?**
```go
package main
import "fmt"

func minMax(a, b int) (int, int) {
    if a < b {
        return a, b
    }
    return b, a
}

func main() {
    lo, hi := minMax(7, 3)
    fmt.Println(lo, hi)
}
```
**A:** `3 7`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this code output and how does it work?
**Your Response:** This code outputs `3 7`. The function `minMax` returns two values - the minimum and maximum of the two input parameters. Inside the function, there's an if statement that checks if `a < b`. If true, it returns `a` first then `b`, otherwise it returns `b` first then `a`. When we call `minMax(7, 3)`, since 7 is not less than 3, the else branch executes, returning `3` first and `7` second. Go supports multiple return values, which is a clean way to return related data without needing to wrap it in a struct.

---

### 3. Zero Values
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```
**A:** `0 0 false ""`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What gets printed here and why?
**Your Response:** This outputs `0 0 false ""`. In Go, every variable has a zero value when declared but not explicitly initialized. For numeric types like `int` and `float64`, the zero value is `0`. For boolean type, it's `false`. For strings, it's an empty string `""`. This is a fundamental Go design principle that ensures variables are never in an undefined state, which helps prevent bugs and makes the code more predictable.

---

### 4. Untyped Constants
**Q: Does this compile?**
```go
package main
import "fmt"

const x = 1e300

func main() {
    var f float32 = x
    fmt.Println(f)
}
```
**A:** **Compile Error.** Constant `1e300` overflows `float32`. Untyped constants have arbitrary precision but must fit the target type at assignment.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why does this code fail to compile?
**Your Response:** This fails to compile because we're trying to assign the constant `1e300` (which is 1 followed by 300 zeros) to a `float32` variable, but `float32` can only hold values up to about 3.4e38. While untyped constants in Go have arbitrary precision, when you assign them to a typed variable, they must fit within that type's limits. The constant `1e300` is way too large for `float32`, so the compiler gives an overflow error. If we used `float64` instead, it would work since `float64` can handle much larger values.

---

### 5. Integer Overflow (Wrapping)
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "math"
)

func main() {
    var x int8 = math.MaxInt8
    x++
    fmt.Println(x)
}
```
**A:** `-128`. Integer overflow wraps around silently in Go at runtime.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What happens when we increment the maximum int8 value?
**Your Response:** This outputs `-128` because Go integers wrap around when they overflow. `int8` can hold values from -128 to 127. When we have the maximum value 127 and increment it, it wraps around to the minimum value -128. Unlike some languages that might panic or throw an exception on overflow, Go's integer arithmetic wraps silently at runtime. This is important to keep in mind when working with fixed-size integer types, especially in performance-critical code or when dealing with counters.

---

### 6. Multiple Assignment
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x, y := 1, 2
    x, y = y, x
    fmt.Println(x, y)
}
```
**A:** `2 1`. Go evaluates all right-hand side expressions before assigning.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does this swap work and what's the output?
**Your Response:** This outputs `2 1`. The magic here is that Go evaluates all the right-hand side expressions first before doing any assignments. So when we write `x, y = y, x`, Go first evaluates `y` (which is 2) and `x` (which is 1), stores these values temporarily, and then assigns them to `x` and `y` respectively. This is why we don't need a temporary variable like in some other languages. It's a clean, idiomatic way to swap values in Go.

---

### 7. Type Alias vs New Type
**Q: Does this compile?**
```go
package main
import "fmt"

type Celsius float64
type Fahrenheit float64

func main() {
    var c Celsius = 100
    var f Fahrenheit = c
    fmt.Println(f)
}
```
**A:** **Compile Error.** `Celsius` and `Fahrenheit` are distinct named types. You need an explicit conversion: `var f Fahrenheit = Fahrenheit(c)`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why doesn't this compile and how would you fix it?
**Your Response:** This doesn't compile because `Celsius` and `Fahrenheit` are different named types, even though they're both based on `float64`. In Go, named types are distinct and cannot be directly assigned to each other, even if they have the same underlying type. This is a type safety feature that prevents accidental mixing of incompatible units. To fix this, we need an explicit type conversion: `var f Fahrenheit = Fahrenheit(c)`. This makes the conversion intentional and clear in the code.

---

### 8. Iota with Bit Shift
**Q: What are the values of KB, MB, GB?**
```go
package main
import "fmt"

const (
    _  = iota
    KB = 1 << (10 * iota)
    MB
    GB
)

func main() {
    fmt.Println(KB, MB, GB)
}
```
**A:** `1024 1048576 1073741824`
In Go, `iota` is a special identifier used in `const` blocks. It starts at **0** and increments by **1** for every line in the block.

- Line 1 (`_`): `iota` is **0**. (The underscore `_` discards this value).
    
- Line 2 (`KB`): `iota` is **1**.
    
- Line 3 (`MB`): `iota` is **2**.
    
- Line 4 (`GB`): `iota` is **3**.
#### **Bit Shifting (`<<`)**

The expression `1 << n` means "take the binary number 1 and move it to the left $n$ times." In mathematics, this is equivalent to $2^{n}$.

- $1 \ll 10$ is $2^{10} = 1,024$.
    
- $1 \ll 20$ is $2^{20} = 1,048,576$.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can you explain how iota and bit shifting work in this code?
**Your Response:** This code defines storage units using `iota` and bit shifting. `iota` starts at 0 and increments by 1 for each line in the const block. We skip the first line with `_`, so KB gets iota=1, MB gets iota=2, and GB gets iota=3. The bit shifting `1 << (10 * iota)` calculates powers of 2. For KB: 1 << 10 = 2^10 = 1024. For MB: 1 << 20 = 2^20 = 1,048,576. For GB: 1 << 30 = 2^30 = 1,073,741,824. This is a common Go idiom for creating constants that follow mathematical patterns.

---

### 9. Blank Identifier
**Q: Does this compile?**
```go
package main
import "fmt"

func twoVals() (int, string) { return 1, "a" }

func main() {
    _, s := twoVals()
    fmt.Println(s)
}
```
**A:** **Yes, compiles and prints** `a`. The blank identifier `_` discards the first return value.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does the blank identifier do here?
**Your Response:** This compiles and prints `a`. The blank identifier `_` is a special placeholder in Go that tells the compiler "I don't care about this value." When `twoVals()` returns two values `(1, "a")`, we use `_` to ignore the first value and only capture the second value in the variable `s`. This is useful when a function returns multiple values but you only need some of them. It's better than creating a variable you'll never use, which would cause a compile error in Go.

---

### 10. var vs :=
**Q: What is the output?**
```go
package main
import "fmt"

var x = 10

func main() {
    x := 20
    fmt.Println(x)
}
```
**A:** `20`. The `:=` inside `main` creates a new local variable `x` that shadows the package-level `x`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What gets printed and why?
**Your Response:** This prints `20` because of variable shadowing. There are actually two different `x` variables here - one at the package level with value 10, and another local to the `main` function with value 20. The `:=` operator declares a new variable, and since there's already an `x` in the outer scope, the new local `x` shadows the package-level one. When we print `x` inside `main`, we're printing the local variable, not the package-level one. This is an important concept to understand to avoid subtle bugs.

---

### 11. Typed nil
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var p *int
    fmt.Println(p == nil)
    fmt.Println(p)
}
```
**A:**
```
true
<nil>
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this code demonstrate about nil pointers?
**Your Response:** This prints `true` and `<nil>`. The first print shows that the pointer `p` is indeed `nil` when compared with `nil`. The second print shows the actual value of the pointer, which is `<nil>`. This demonstrates that in Go, pointers have a zero value of `nil`, and you can safely check if a pointer is `nil` before dereferencing it. Dereferencing a `nil` pointer would cause a panic, so this nil check is a common pattern for safety.

---

### 12. Constants Cannot Be Addressed
**Q: Does this compile?**
```go
package main

const x = 42

func main() {
    p := &x
    _ = p
}
```
**A:** **Compile Error.** You cannot take the address of a constant in Go.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why doesn't this compile?
**Your Response:** This doesn't compile because you can't take the address of a constant in Go. Constants are compile-time values that don't have a memory address - they're embedded directly into the code where they're used. Only variables have memory addresses. If you need a pointer to the value 42, you'd need to first store it in a variable like `x := 42` and then take its address with `&x`. This is a fundamental difference between constants and variables in Go's memory model.

---

### 13. Declaring Multiple Variables
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var (
        a = 1
        b = 2
        c = a + b
    )
    fmt.Println(c)
}
```
**A:** `3`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this code output?
**Your Response:** This outputs `3`. The code declares three variables using the `var` keyword with parentheses syntax. Variables `a` and `b` are initialized to 1 and 2 respectively, and then `c` is calculated as `a + b`, which equals 3. This parenthesized `var` syntax is a clean way to declare multiple related variables at once, especially when some depend on others. It's commonly used for initializing configuration variables or related constants.

---

### 14. Iota Reset Per const Block
**Q: What are the values of X and Y?**
```go
package main
import "fmt"

const A = iota
const (
    X = iota
    Y
)

func main() {
    fmt.Println(A, X, Y)
}
```
**A:** `0 0 1`. `iota` resets to `0` at the start of each new `const` block.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why do we get these values for iota?
**Your Response:** This outputs `0 0 1` because `iota` resets to 0 at the start of each const block. The first const block has just one line `const A = iota`, so A gets 0. Then we have a completely new const block with parentheses, where iota starts fresh at 0 again. So X gets 0 and Y gets 1. This demonstrates that `iota`'s counter is scoped to individual const blocks, not global across all const declarations. This reset behavior is important to understand when using iota for creating enumerated constants.

---

### 15. Unused Variables
**Q: Does this compile?**
```go
package main

func main() {
    x := 5
}
```
**A:** **Compile Error.** `x declared and not used`. Go enforces that every declared local variable must be used.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why doesn't this code compile?
**Your Response:** This doesn't compile because Go has a strict rule that every local variable declared must be used. Here we declare `x := 5` but never use it, so the compiler gives an error. This is actually a helpful feature that prevents code clutter and potential bugs from unused variables. It encourages clean, maintainable code. If you really need to declare a variable you won't use, you can use the blank identifier `_ := 5`, though that's uncommon. This rule applies to local variables but not to package-level variables or function parameters.

---

## Section 2: Control Flow (Q16–Q28)

### 16. For as While Loop
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    i := 0
    for i < 3 {
        fmt.Print(i)
        i++
    }
}
```
**A:** `012`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this code output and how does the for loop work?
**Your Response:** This outputs `012`. This is Go's way of creating a while loop. The for statement has only a condition `i < 3`, without any initialization or post statement. It starts with `i = 0`, checks if `i < 3` (true), prints `i` (0), increments `i`, and repeats. This continues until `i` becomes 3, at which point the condition fails and the loop exits. Go doesn't have a separate `while` keyword - the `for` loop serves this purpose when you only provide a condition.

---

### 17. Infinite Loop with Break
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    i := 0
    for {
        if i == 3 {
            break
        }
        fmt.Print(i)
        i++
    }
}
```
**A:** `012`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does this infinite loop work and what gets printed?
**Your Response:** This outputs `012`. This is an infinite loop written as `for {}` with no condition at all. Inside the loop, we check if `i == 3`, and when that's true, we break out of the loop. The loop starts with `i = 0`, prints 0, increments to 1, prints 1, increments to 2, prints 2, increments to 3, then the break statement executes and we exit the loop. This pattern is common in Go when you need to loop until some condition is met, but you don't know ahead of time how many iterations it will take.

---

### 18. Switch Without Condition
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x := 15
    switch {
    case x < 10:
        fmt.Println("small")
    case x < 20:
        fmt.Println("medium")
    default:
        fmt.Println("large")
    }
}
```
**A:** `medium`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this switch statement do and why doesn't it need a condition?
**Your Response:** This prints `medium`. This is a switch statement without an expression, which acts like a series of if-else if statements. Each case contains a condition that's evaluated in order. When `x = 15`, it first checks `x < 10` (false), then `x < 20` (true), so it executes that case and prints `medium`. This is a clean alternative to long if-else chains in Go, making the code more readable when you have multiple conditions to check.

---

### 19. Switch Type
**Q: What is the output?**
```go
package main
import "fmt"

func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("int: %d\n", v)
    case string:
        fmt.Printf("string: %s\n", v)
    default:
        fmt.Printf("other: %T\n", v)
    }
}

func main() {
    describe(42)
    describe("hi")
    describe(3.14)
}
```
**A:**
```
int: 42
string: hi
other: float64
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does this type switch work?
**Your Response:** This demonstrates a type switch in Go. The `switch v := i.(type)` statement checks the actual type of the interface value `i`. When we pass `42`, it's an `int`, so the first case matches and prints `int: 42`. When we pass `"hi"`, it's a `string`, so the second case matches and prints `string: hi`. When we pass `3.14`, it's a `float64`, which doesn't match any specific case, so the default case executes and prints `other: float64`. The `%T` verb prints the type of the value. Type switches are a powerful Go feature for working with interface values when you need to know their concrete type.

---

### 20. Continue in For Loop
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Print(i)
    }
}
```
**A:** `13`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What gets printed and how does the continue statement work?
**Your Response:** This outputs `13`. The loop iterates from 0 to 4. When `i` is 0, `i % 2 == 0` is true, so `continue` skips to the next iteration without printing. When `i` is 1, the condition is false, so it prints 1. When `i` is 2, continue executes again. When `i` is 3, it prints 3. When `i` is 4, continue executes again. So only the odd numbers 1 and 3 get printed. The `continue` statement skips the rest of the current iteration and moves to the next one.

---

### 21. Labeled Break
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if j == 1 {
                break outer
            }
            fmt.Print(i, j, " ")
        }
    }
}
```
**A:** `00 `. The labeled `break outer` exits the outer loop immediately when `j == 1` on the first iteration.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the labeled break work here?
**Your Response:** This prints `00 `. The `outer:` label marks the outer for loop. Inside the nested loops, when `j == 1` on the first iteration of the outer loop, `break outer` executes. Unlike a regular break which would only exit the inner loop, the labeled break exits both loops immediately. So we only get one iteration of the outer loop where `i = 0`, and within that, only one iteration of the inner loop where `j = 0`, printing `00 `. Labeled breaks are useful when you need to break out of nested loops from within the inner loop.

---

### 22. Switch Fallthrough Does Not Check Condition
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x := 1
    switch x {
    case 1:
        fmt.Println("one")
        fallthrough
    case 2:
        fmt.Println("two")
        fallthrough
    case 3:
        fmt.Println("three")
    }
}
```
**A:**
```
one
two
three
```
`fallthrough` bypasses the case condition check entirely.
**No Condition Check**: When you "fall through," Go **does not check the condition** of the next case. It simply executes the next block of code.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does fallthrough work in switch statements?
**Your Response:** This prints all three lines: `one`, `two`, and `three`. When `x = 1`, the first case matches and prints `one`. The `fallthrough` keyword then causes execution to continue to the next case, printing `two`, even though `x` doesn't equal 2. The second `fallthrough` continues to the third case, printing `three`. Unlike other languages where cases fall through by default, in Go you need explicit `fallthrough`. Importantly, fallthrough bypasses the condition check - it just executes the next case's code regardless of whether the condition matches.
---

### 23. goto Statement
**Q: Does this compile and what is the output?**
```go
package main
import "fmt"

func main() {
    i := 0
loop:
    if i < 3 {
        fmt.Print(i)
        i++
        goto loop
    }
}
```
**A:** **Compiles and prints** `012`. `goto` is valid in Go but discouraged.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Is goto valid in Go and what does this code do?
**Your Response:** This compiles and prints `012`. Yes, `goto` is valid in Go, though it's generally discouraged except in specific situations like error handling cleanup. The `loop:` label marks a location, and `goto loop` jumps execution to that label. This creates a loop that continues while `i < 3`, incrementing `i` each time. While this works, Go provides better structured alternatives like for loops that are more readable and maintainable. Goto is typically only used in Go for breaking out of nested loops or centralized error handling.

---

### 24. Range Over String
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    for i, c := range "Go" {
        fmt.Println(i, c)
    }
}
```
**A:**
```
0 71
1 111
```
`range` over a string yields byte index and Unicode code point (`rune`). `G`=71, `o`=111.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does ranging over a string give you?
**Your Response:** This prints the byte index and Unicode code point for each character. For the string `"Go"`, it prints `0 71` and `1 111`. The first value is the byte index (0 and 1), and the second value is the rune (Unicode code point) - 71 for 'G' and 111 for 'o'. Unlike some languages that might give you characters, Go gives you runes, which are Unicode code points. This is important because Go strings are UTF-8 encoded, and a single character might be multiple bytes, but range handles this correctly by giving you the actual Unicode characters.

---

### 25. Index-Only Range
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := []int{10, 20, 30}
    for i := range s {
        fmt.Print(i)
    }
}
```
**A:** `012`. When you use a single variable in `range`, you get only the index.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why does this only print indices?
**Your Response:** This prints `012` because when you use a single variable with range, you only get the index, not the value. The slice `s` has 3 elements at indices 0, 1, and 2. The range loop iterates through each index and prints it. If we wanted both index and value, we'd write `for i, v := range s`. If we only wanted the values, we'd use the blank identifier: `for _, v := range s`. This is a common Go pattern when you only need the indices of a slice or array.

---

### 26. Switch with Multiple Values Per Case
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    day := "Saturday"
    switch day {
    case "Saturday", "Sunday":
        fmt.Println("Weekend")
    default:
        fmt.Println("Weekday")
    }
}
```
**A:** `Weekend`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does this switch case with multiple values work?
**Your Response:** This prints `Weekend`. The switch statement allows multiple values in a single case, separated by commas. When `day = "Saturday"`, it matches the first case `"Saturday", "Sunday"`, so it prints `Weekend`. This is equivalent to having separate cases for each value but more concise. If the day were any other value, it would fall through to the default case and print `Weekday`. This feature makes switch statements cleaner when you need the same action for multiple different values.

---

### 27. Defer in Loop Execution Order
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        defer fmt.Print(i)
    }
}
```
**A:** `210`. Deferred calls execute in LIFO (last-in, first-out) order after the surrounding function returns.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why does this print in reverse order?
**Your Response:** This prints `210` because deferred function calls execute in LIFO (last-in, first-out) order. In the loop, we defer three print statements: first defer `Print(0)`, then defer `Print(1)`, then defer `Print(2)`. When `main` function ends, these deferred calls execute in reverse order - last one deferred executes first. So it prints 2, then 1, then 0. This LIFO behavior is fundamental to how defer works in Go and is useful for cleanup operations that need to happen in reverse order of acquisition.

---

### 28. If with Initialization Statement
**Q: What is the output?**
```go
package main
import "fmt"

func getValue() int { return 42 }

func main() {
    if v := getValue(); v > 10 {
        fmt.Println("big:", v)
    }
    // fmt.Println(v) // v is not in scope here
}
```
**A:** `big: 42`. The `v` initialized in the `if` statement is scoped only to the `if-else` block.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the if statement with initialization work?
**Your Response:** This prints `big: 42`. The if statement has an initialization clause `v := getValue()` that declares and initializes `v` before checking the condition. The scope of `v` is limited to the if statement and any else blocks. Here, `getValue()` returns 42, which is greater than 10, so the if body executes and prints `big: 42`. If we tried to use `v` after the if statement, we'd get a compile error because it's out of scope. This pattern is commonly used in Go for error handling: `if err := someFunction(); err != nil { ... }`.

---

## Section 3: Functions, Closures & Defer (Q29–Q44)

### 29. Function as First-Class Value
**Q: What is the output?**
```go
package main
import "fmt"

func apply(f func(int) int, x int) int {
    return f(x)
}

func double(n int) int { return n * 2 }

func main() {
    fmt.Println(apply(double, 5))
}
```
**A:** `10`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does this higher-order function work?
**Your Response:** This outputs `10`. The `apply` function is a higher-order function that takes another function `f` as a parameter and an integer `x`. It calls the function `f` with `x` and returns the result. We pass the `double` function to `apply`, which doubles any number. When we call `apply(double, 5)`, it calls `double(5)` which returns `10`. This demonstrates that functions are first-class citizens in Go - they can be passed as arguments, returned from other functions, and assigned to variables. This enables powerful functional programming patterns.

---

### 30. Closure Captures Variable by Reference
**Q: What is the output?**
```go
package main
import "fmt"

func makeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    c := makeCounter()
    fmt.Println(c())
    fmt.Println(c())
    fmt.Println(c())
}
```
**A:** `1 2 3`. The closure captures the variable `count` by reference — each call increments the same `count`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does this closure work and why does it produce an increasing sequence?
**Your Response:** This outputs `1 2 3`. The `makeCounter` function returns a closure - an anonymous function that captures the `count` variable from its outer scope. The key thing is that the closure captures `count` by reference, not by value. So each time we call the returned function, it's incrementing the same `count` variable. The first call increments from 0 to 1 and returns 1, the second increments from 1 to 2 and returns 2, and so on. This is how you create stateful functions in Go, and it's fundamental to understanding how closures work.

---

### 31. Multiple Defers Order
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    defer fmt.Println("first")
    defer fmt.Println("second")
    defer fmt.Println("third")
    fmt.Println("start")
}
```
**A:**
```
start
third
second
first
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why do the deferred statements print in reverse order?
**Your Response:** This prints "start" first, then "third", "second", and "first". The regular `fmt.Println("start")` executes immediately. But the three deferred statements are pushed onto a stack and execute in LIFO (last-in, first-out) order when the function returns. The last defer statement `defer fmt.Println("third")` was added last, so it executes first among the deferred calls. This LIFO behavior is fundamental to Go's defer mechanism and is useful for cleanup operations that need to happen in reverse order.

---

### 32. Variadic Function
**Q: What is the output?**
```go
package main
import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))
    s := []int{4, 5, 6}
    fmt.Println(sum(s...))
}
```
**A:**
```
6
15
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do variadic functions work in Go?
**Your Response:** This prints `6` and `15`. The `sum` function is variadic - it can accept any number of integer arguments. The `...int` syntax means `nums` is a slice of integers inside the function. When we call `sum(1, 2, 3)`, Go packs these into a slice `[1, 2, 3]` and sums them to get 6. The second call `sum(s...)` uses the `...` operator to unpack the slice `s` into individual arguments. This is called variadic unpacking and it's the reverse of what happens when you call the function. Variadic functions are Go's way of implementing functions with variable numbers of arguments.

---

### 33. Named Return Values
**Q: What is the output?**
```go
package main
import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}
```
**A:** `7 10`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do named return values work in Go?
**Your Response:** This outputs `7 10`. The `split` function has named return values `(x, y int)`. When you use named return values, Go automatically creates these variables inside the function. The `return` statement without any values returns the current values of `x` and `y`. When we call `split(17)`, it calculates `x = 17 * 4 / 9 = 7` (integer division), then `y = 17 - 7 = 10`. The naked `return` returns these values. Named returns are useful for clarity and when you need to modify return values in deferred functions, as they're accessible throughout the function scope.

---

### 34. Defer and Panic Recovery
**Q: What is the output?**
```go
package main
import "fmt"

func safeDiv(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)
        }
    }()
    result = a / b
    return
}

func main() {
    r, err := safeDiv(10, 0)
    fmt.Println(r, err)
}
```
**A:** `0 recovered: runtime error: integer divide by zero`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does panic and recover work together with defer?
**Your Response:** This outputs `0 recovered: runtime error: integer divide by zero`. The `safeDiv` function demonstrates Go's panic/recover mechanism. When we divide by zero, Go panics with a runtime error. However, we have a deferred function that calls `recover()`. The `recover()` function catches the panic value and allows the program to continue gracefully instead of crashing. The deferred function runs during the panic, sets the error variable, and then the function returns normally with result 0 and the error message. This pattern is commonly used in Go for error handling in functions that might panic, allowing them to return errors instead of crashing the program.

---

### 35. Closure in Goroutine (Classic Bug)
**Q: What is the typical output (pre Go 1.22)?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Print(i)
        }()
    }
    wg.Wait()
}
```
**A:** Usually `333`. The goroutines capture the same variable `i`; by the time they run, the loop has finished and `i` is `3`.
**Fix:** `go func(n int) { ... }(i)`

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why does this code typically print 333 and how do you fix it?
**Your Response:** This typically prints `333` because of a classic Go closure bug. All goroutines capture the same variable `i` by reference, not by value. The loop finishes quickly, setting `i` to 3, then the goroutines run and all see `i` as 3. The fix is to pass `i` as a parameter to the goroutine: `go func(n int) { fmt.Print(n) }(i)`. This creates a new copy of `i` for each goroutine. This is one of the most common Go concurrency bugs and understanding it shows you know how closures work with goroutines. In Go 1.22+, this behavior changed, but this pattern is still important to understand.

## The Channel-Based Solution (The "Fan-In" Pattern)

```go
package main

import (
    "fmt"
)

func main() {
    count := 3
    results := make(chan int)

    for i := 0; i < count; i++ {
        go func(val int) {
            // Perform work...
            results <- val // Send result to channel
        }(i)
    }

    // Instead of wg.Wait(), we just loop the exact number of times
    // This blocks the main thread until 'count' items are received
    for i := 0; i < count; i++ {
        fmt.Println(<-results) 
    }
}
```
## using waitgroup to solve above problem

```go 
package main

import (
    "fmt"
    "sync"
)

func main() {
    const count = 3
    // Pre-allocate the exact size needed
    results := make([]int, count)
    var wg sync.WaitGroup

    for i := 0; i < count; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            // Optimization: Each goroutine writes to its own index
            // No sorting needed later because the index 'id' is already ordered
            results[id] = id 
        }(i)
    }

    wg.Wait()

    // The data is already "sorted" by the time we get here
    for _, val := range results {
        fmt.Println(val)
    }
}
```



---

### 36. Function Type as Struct Field
**Q: What is the output?**
```go
package main
import "fmt"

type Greeter struct {
    Greet func(name string) string
}

func main() {
    g := Greeter{
        Greet: func(name string) string {
            return "Hello, " + name
        },
    }
    fmt.Println(g.Greet("Go"))
}
```
**A:** `Hello, Go`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does this struct with a function field work?
**Your Response:** This prints `Hello, Go`. The `Greeter` struct has a field `Greet` that's of function type `func(name string) string`. When we create a `Greeter` instance, we assign an anonymous function to this field that returns a greeting string. When we call `g.Greet("Go")`, we're calling the function stored in that field. This demonstrates that functions are first-class citizens in Go - they can be stored in struct fields, passed around, and called just like any other value. This pattern is useful for implementing strategies, callbacks, or dependency injection.

---

### 37. Deferred Function Sees Updated Return Value
**Q: What is returned?**
```go
package main
import "fmt"

func f() (n int) {
    defer func() {
        n++
    }()
    return 1
}

func main() {
    fmt.Println(f())
}
```
**A:** `2`. `return 1` sets named return `n = 1`, then the deferred closure increments it to `2`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why does this function return 2 instead of 1?
**Your Response:** This returns `2` because of how named return values work with defer. The function has a named return value `n`. When `return 1` executes, it sets `n = 1`. However, deferred functions execute after the return statement but before the function actually returns to the caller. The deferred closure increments `n` from 1 to 2. Since the deferred function has access to the named return variable, it can modify the final return value. This is a powerful Go feature that's useful for things like logging return values or ensuring cleanup happens even when returning early.

---

### 38. Anonymous Function Immediately Invoked
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    result := func(x, y int) int {
        return x + y
    }(3, 4)
    fmt.Println(result)
}
```
**A:** `7`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's happening in this code with the immediate function invocation?
**Your Response:** This outputs `7`. This demonstrates an immediately invoked function expression (IIFE) in Go. We define an anonymous function `func(x, y int) int { return x + y }` and immediately call it with arguments `(3, 4)`. The result `7` is assigned to the `result` variable. This pattern is useful when you need to create a self-contained scope or perform a one-time calculation without defining a named function. It's commonly used for initialization logic or when you want to limit the scope of variables to just that computation.

---

### 39. Defer with Panic (No Recover)
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    defer fmt.Println("deferred")
    panic("crash!")
}
```
**A:**
```
deferred
panic: crash!
...
```
Deferred functions run even during a panic before the program prints the panic message and exits.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What happens when defer and panic are used together?
**Your Response:** This prints "deferred" first, then shows the panic message. Even though `panic("crash!")` is called, the deferred function still executes before the program terminates. This demonstrates that deferred functions run during a panic, which is crucial for cleanup operations. The defer stack is unwound and all deferred functions execute in LIFO order, then the panic message is printed and the program exits. This is why defer is often used for resource cleanup - it guarantees cleanup happens even if something unexpected goes wrong and the function panics.

---

### 40. Recursive Closure
**Q: Does this compile?**
```go
package main
import "fmt"

func main() {
    var fib func(n int) int
    fib = func(n int) int {
        if n <= 1 {
            return n
        }
        return fib(n-1) + fib(n-2)
    }
    fmt.Println(fib(7))
}
```
**A:** `2`

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this recursive closure demonstrate?
**Your Response:** This compiles and prints `13`. The key here is that we first declare `var fib func(n int) int` to create a function variable, then assign an anonymous function to it. This allows the anonymous function to refer to itself by name `fib` recursively. Without the initial declaration, the function couldn't refer to itself because it wouldn't be in scope yet. This pattern is Go's way of creating recursive anonymous functions. The function calculates the nth Fibonacci number - when called with 7, it returns 13 (0, 1, 1, 2, 3, 5, 8, 13).

---

### 41. Panic in Deferred Function
**Q: What happens?**
```go
package main
import "fmt"

func main() {
    defer func() {
        fmt.Println("defer 1")
        panic("panic in defer")
    }()
    defer fmt.Println("defer 2")
    fmt.Println("main")
}
```
**A:**
```
main
defer 2
defer 1
panic: panic in defer
```
Deferred functions still run in LIFO order. A panic inside a deferred function propagates and can be recovered by another deferred function.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What happens when a panic occurs inside a deferred function?
**Your Response:** This shows how panic and defer interact. The code prints "main" first, then "defer 2", then "defer 1", then panics. Even though the first deferred function panics, all deferred functions still execute in LIFO order. When a deferred function panics, it replaces any existing panic, but the defer stack continues to unwind. This means if you have multiple deferred functions and one panics, the others will still execute. This is important for cleanup - even if cleanup code panics, other cleanup operations will still run.

---

### 42. Function With No Parameters or Returns
**Q: Does this compile?**
```go
package main
import "fmt"

func sayHi() {
    fmt.Println("Hi!")
}

func main() {
    f := sayHi
    f()
}
```
**A:** **Yes, compiles and prints** `Hi!`. Functions are first-class values and can be assigned to variables.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How can functions be assigned to variables in Go?
**Your Response:** This prints `Hi!` and demonstrates that functions are first-class citizens in Go. We can assign the `sayHi` function to the variable `f` using `f := sayHi`. After this assignment, `f` and `sayHi` refer to the same function. When we call `f()`, it's identical to calling `sayHi()`. This ability to treat functions as values allows for higher-order functions, callbacks, and functional programming patterns in Go. You can store functions in variables, pass them as arguments, and return them from other functions.

---

### 43. Unused Function Parameter
**Q: Does this compile?**
```go
package main
import "fmt"

func greet(name string) {
    fmt.Println("Hello!")
}

func main() {
    greet("Alice")
}
```
**A:** **Yes.** Unlike local variables, unused function parameters do **not** cause a compile error.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why doesn't this code have a compile error despite the unused parameter?
**Your Response:** This compiles successfully even though the `name` parameter is never used. In Go, while unused local variables cause compile errors, unused function parameters are allowed. This is because function signatures often need to maintain compatibility with interfaces or other code, even if a particular implementation doesn't use all parameters. The blank identifier `_` would be clearer if the parameter is truly unused, but it's not required. This design choice makes Go more flexible for implementing interfaces where you might not need all the parameters in every implementation.

---

### 44. defer + os.Exit
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "os"
)

func main() {
    defer fmt.Println("deferred")
    os.Exit(0)
}
```
**A:** Nothing is printed. `os.Exit` immediately terminates the process; deferred functions are **not** called.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why doesn't the deferred statement execute when using os.Exit?
**Your Response:** This prints nothing because `os.Exit(0)` terminates the program immediately without running deferred functions. Unlike a normal return or panic, `os.Exit` bypasses Go's normal defer mechanism entirely. It calls the underlying system exit function directly. This is important to understand - if you need cleanup to happen, don't use `os.Exit`. Instead, return an error and handle it in the calling code, or use panic/recover if appropriate. The defer mechanism only works with Go's normal function return flow, not with abrupt process termination.

---

## Section 4: Pointers (Q45–Q52)

### 45. Pointer Basics
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x := 42
    p := &x
    *p = 100
    fmt.Println(x)
}
```
**A:** `100`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do pointers work in this basic example?
**Your Response:** This outputs `100`. Here we create a variable `x` with value 42, then create a pointer `p` that points to the memory address of `x` using `&x`. The `*p = 100` dereferences the pointer and modifies the value at that memory address from 42 to 100. Since `p` points to `x`, changing `*p` actually changes `x`. So when we print `x`, it shows the updated value 100. This demonstrates the fundamental pointer concept: pointers reference memory locations, and dereferencing allows you to modify the original value through the pointer.

---

### 46. Pointer to Pointer
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    x := 10
    p := &x
    pp := &p
    **pp = 20
    fmt.Println(x)
}
```
**A:** `20`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does a pointer to a pointer work?
**Your Response:** This outputs `20`. Here we have a double pointer or pointer-to-pointer. `x` is 10, `p` is a pointer to `x` (contains `&x`), and `pp` is a pointer to `p` (contains `&p`). When we do `**pp = 20`, we're double-dereferencing: first `*pp` gives us `p` (the pointer to `x`), then `**pp` gives us `x` (the original variable). So `**pp = 20` is equivalent to `x = 20`. This demonstrates that you can have pointers to pointers, and each `*` dereferences one level. Double pointers are commonly used in Go for functions that need to modify pointers passed to them.

---

### 47. new() Builtin
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    p := new(int)
    fmt.Println(*p)
    *p = 7
    fmt.Println(*p)
}
```
**A:**
```
0
7
```
`new(T)` allocates a zeroed T and returns a pointer to it.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the new() builtin function work?
**Your Response:** This prints `0` then `7`. The `new(int)` function allocates memory for an integer, initializes it to zero value (which is 0 for int), and returns a pointer to that memory. When we first print `*p`, we get 0 because that's the zero value. Then we assign `*p = 7`, which updates the value at that memory location, and the second print shows 7. The `new()` function is Go's built-in way to allocate memory and get a pointer. It's different from `&` which takes the address of an existing variable - `new()` creates a new zeroed variable and returns its address.

---

### 48. Nil Pointer Dereference
**Q: What is the output or error?**
```go
package main
import "fmt"

type Node struct{ val int }

func main() {
    var n *Node
    fmt.Println(n.val)
}
```
**A:** **Panic.** `runtime error: invalid memory address or nil pointer dereference`. You cannot dereference a nil pointer.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What happens when you try to access a field of a nil struct pointer?
**Your Response:** This panics with a runtime error because we're trying to dereference a nil pointer. The variable `n` is declared as `*Node` but not initialized, so it's `nil`. When we try to access `n.val`, Go tries to dereference the nil pointer to get to the struct, which causes a panic. This is one of the most common runtime errors in Go. The fix is to either initialize the pointer with `new(Node)` or `&Node{val: 42}`, or check for nil before dereferencing: `if n != nil { fmt.Println(n.val) }`. This demonstrates why nil checks are important when working with pointers.

---

### 49. Passing Pointer to Function
**Q: What is the output?**
```go
package main
import "fmt"

func increment(x *int) {
    *x++
}

func main() {
    n := 5
    increment(&n)
    fmt.Println(n)
}
```
**A:** `6`

---

### 50. Pointer Comparison
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a, b := 1, 1
    p1, p2 := &a, &b
    p3 := &a
    fmt.Println(p1 == p2)
    fmt.Println(p1 == p3)
}
```
**A:**
```
false
true
```
Pointers are compared by address, not by the value they point to.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How are pointers compared in Go?
**Your Response:** This prints `false` then `true`. The first comparison `p1 == p2` is false because `p1` points to variable `a` and `p2` points to variable `b` - they have different memory addresses even though both variables contain the value 1. The second comparison `p1 == p3` is true because both pointers point to the same variable `a`, so they have the same memory address. This shows that pointer comparison in Go is about comparing memory addresses, not the values they point to. This is useful for checking if two pointers refer to the same data structure.

---

### 51. Returning Local Variable Pointer
**Q: Is this safe in Go?**
```go
package main
import "fmt"

func newInt() *int {
    x := 42
    return &x
}

func main() {
    p := newInt()
    fmt.Println(*p)
}
```
**A:** **Yes, it is safe.** Go's compiler detects that `x` escapes to the heap and allocates it there. This is called escape analysis. Output: `42`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Is it safe to return a pointer to a local variable in Go?
**Your Response:** This is safe in Go and prints `42`. In many languages like C, returning a pointer to a local variable would be dangerous because the local variable's memory would be freed when the function returns. But Go's compiler performs escape analysis - it detects that the variable `x` "escapes" the function scope (because we return its address), so it allocates `x` on the heap instead of the stack. The heap-allocated memory persists after the function returns, making the pointer valid. This is a key Go safety feature that prevents common memory bugs while still allowing efficient memory management through garbage collection.

---

### 52. Modifying Struct Through Pointer
**Q: What is the output?**
```go
package main
import "fmt"

type Point struct{ X, Y int }

func moveRight(p *Point) {
    p.X += 10
}

func main() {
    pt := Point{1, 2}
    moveRight(&pt)
    fmt.Println(pt)
}
```
**A:** `{11 2}`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does modifying a struct through a pointer work?
**Your Response:** This prints `{11 2}`. We create a `Point` struct with values `{1, 2}`, then pass a pointer to it to the `moveRight` function using `&pt`. Inside the function, `p.X += 10` accesses the `X` field through the pointer - Go automatically dereferences the pointer when accessing struct fields. This modifies the original `pt` struct, changing its `X` value from 1 to 11. The `Y` value remains 2. This demonstrates how pointers allow functions to modify structs directly, which is more efficient than passing large structs by value and enables in-place modifications.

---

## Section 5: Strings & Runes (Q53–Q62)

### 53. len() on a String
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := "café"
    fmt.Println(len(s))
}
```
**A:** `5`. `len()` counts bytes, not characters. The `é` character is 2 bytes in UTF-8.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why does len() return 5 for "café" when it only has 4 characters?
**Your Response:** This prints `5` because `len()` counts bytes, not characters. The string "café" contains 4 characters, but the `é` character is encoded as 2 bytes in UTF-8 (Go's string encoding). The other characters `c`, `a`, `f` are each 1 byte, so the total is 1+1+1+2 = 5 bytes. This is a fundamental concept in Go - strings are UTF-8 encoded byte sequences, and `len()` gives you the byte count. If you want the character count, you need to convert to a rune slice or use the `utf8.RuneCountInString()` function. This distinction is important when working with Unicode text.

---

### 54. String Concatenation in Loop (Performance)
**Q: What is the inefficiency?**
```go
package main
import "fmt"

func main() {
    s := ""
    for i := 0; i < 5; i++ {
        s += fmt.Sprintf("%d", i)
    }
    fmt.Println(s)
}
```
**A:** Output is `01234`, but each `+=` allocates a new string (strings are immutable). For large loops, use `strings.Builder`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's inefficient about this string concatenation in a loop?
**Your Response:** This prints `01234`, but it's inefficient because strings in Go are immutable. Each time we do `s += fmt.Sprintf("%d", i)`, Go creates a completely new string by copying all existing characters plus the new ones. For 5 iterations this is fine, but for large loops this creates O(n²) memory allocation and copying. The idiomatic Go solution is to use `strings.Builder`, which pre-allocates a buffer and appends efficiently, or create a slice of strings and join them once. This demonstrates an important Go performance principle - avoid repeated string concatenation in loops.

---

### 55. Byte Slice to String Conversion
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    b := []byte{72, 101, 108, 108, 111}
    s := string(b)
    fmt.Println(s)
}
```
**A:** `Hello`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does byte slice to string conversion work?
**Your Response:** This prints `Hello`. We have a byte slice containing the ASCII values for the letters H, e, l, l, o (72, 101, 108, 108, 111). When we convert this byte slice to a string using `string(b)`, Go interprets these bytes as UTF-8 encoded text. Since these are valid ASCII characters, they form the string "Hello". This conversion creates a new string - it doesn't modify the original byte slice. This is useful when working with network data, file I/O, or other byte-oriented operations that need to be converted to human-readable text.

---

### 56. String to Rune Slice
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := "café"
    r := []rune(s)
    fmt.Println(len(r))
}
```
**A:** `4`. Converting to `[]rune` counts Unicode code points, not bytes.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why does converting to []rune give us 4 instead of 5?
**Your Response:** This prints `4` because converting a string to `[]rune` counts Unicode code points (characters), not bytes. While `len()` gave us 5 bytes for "café" (because `é` is 2 bytes), `[]rune(s)` properly handles UTF-8 encoding and gives us an array where each element represents one character. So we get 4 runes: 'c', 'a', 'f', 'é'. This is the correct way to count characters in Go when dealing with Unicode text. The difference between byte count and character count is crucial for internationalization and text processing.

---

### 57. Comparing Strings
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := "hello"
    b := "hello"
    fmt.Println(a == b)
}
```
**A:** `true`. String comparison in Go is value-based (compares contents byte by byte).

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does string comparison work in Go?
**Your Response:** This prints `true` because string comparison in Go is value-based, not reference-based. Even though `a` and `b` are different variables, Go compares their contents byte by byte. Since both contain the exact same sequence of bytes representing "hello", the comparison returns true. This is different from some languages where string comparison might check if two variables reference the same string object. In Go, strings are immutable, so the runtime can safely compare by content. This makes string comparison intuitive and predictable.

---

### 58. String Indexing Returns Byte
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := "Go"
    fmt.Printf("%T %v\n", s[0], s[0])
}
```
**A:** `uint8 71`. Indexing a string returns a `byte` (`uint8`), not a `rune`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does string indexing return in Go?
**Your Response:** This prints `uint8 71`. When you index a string in Go with `s[0]`, you get a byte (which has type `uint8`), not a character or rune. The value 71 is the ASCII/UTF-8 byte value for the character 'G'. This is important because string indexing gives you raw bytes, not Unicode characters. For multi-byte characters like 'é', indexing would give you the first byte of its UTF-8 encoding. If you want to work with characters properly, you should use range over the string or convert to `[]rune`. This distinction is fundamental to understanding how strings work in Go.

---

### 59. String Contains Check
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Contains("seafood", "foo"))
    fmt.Println(strings.Contains("seafood", "bar"))
}
```
**A:**
```
true
false
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the strings.Contains function work?
**Your Response:** This prints `true` then `false`. The `strings.Contains` function checks if one string is a substring of another. When we check `strings.Contains("seafood", "foo")`, it returns true because "foo" appears within "seafood". When we check `strings.Contains("seafood", "bar")`, it returns false because "bar" doesn't appear in "seafood". This is a common string operation used for searching within text. It's case-sensitive and works with UTF-8 encoded strings. For case-insensitive searching, you'd use `strings.ToLower()` first or use other string matching functions.

---

### 60. Rune Literal
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    r := 'A'
    fmt.Printf("%T %v\n", r, r)
}
```
**A:** `int32 65`. Rune literals have type `rune`, which is an alias for `int32`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What type is a rune literal in Go?
**Your Response:** This prints `int32 65`. In Go, the `rune` type is actually just an alias for `int32`. When we write `r := 'A'`, we're creating a rune literal containing the Unicode code point for the character 'A', which is 65. The `%T` format verb shows the actual type, which is `int32`, and `%v` shows the value 65. Runes are Go's way of representing individual Unicode characters. They're essentially integers that represent Unicode code points. This is why you can do arithmetic with runes and why they're useful for text processing at the character level.

---

### 61. String Builder
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    var sb strings.Builder
    for i := 0; i < 3; i++ {
        sb.WriteString("Go")
    }
    fmt.Println(sb.String())
}
```
**A:** `GoGoGo`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does strings.Builder work?
**Your Response:** This prints `GoGoGo`. The `strings.Builder` type provides an efficient way to build strings by concatenation. Unlike the `+=` operator which creates a new string each time, `Builder` maintains an internal buffer and efficiently appends to it. We call `sb.WriteString("Go")` three times in the loop, and each call appends to the same buffer. Finally, `sb.String()` converts the buffer to a string. This is the idiomatic way to build strings in loops in Go because it avoids the O(n²) performance of repeated string concatenation. `Builder` is pre-allocated with a reasonable starting capacity and grows as needed.

---

### 62. Strings Are Not Byte Slices
**Q: Does this compile?**
```go
package main

func main() {
    s := "hello"
    b := []byte(s)
    s2 := string(b)
    _ = s2
    b[0] = 'H'
    _ = s
}
```
**A:** **Yes.** Mutation of `b` does not affect `s`. `string(b)` copies; strings and byte slices are independent.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Are strings and byte slices independent in Go?
**Your Response:** This compiles and demonstrates that strings and byte slices are independent. We start with string `s`, convert it to byte slice `b` using `[]byte(s)`, then convert it back to string `s2`. The key point is that these conversions create copies - they don't share underlying memory. When we modify `b[0] = 'H'`, it only affects the byte slice, not the original string `s`. This shows that Go strings are immutable - once created, they can't be changed. Converting between strings and byte slices always creates new data, which is why modifying the byte slice doesn't affect the original string.

---

## Section 6: Arrays, Slices & Maps (Q63–Q78)

### 63. Array Is a Value Type
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := [3]int{1, 2, 3}
    b := a
    b[0] = 99
    fmt.Println(a[0], b[0])
}
```
**A:** `1 99`. Arrays are value types; `b` is a complete copy of `a`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why does modifying array b not affect array a?
**Your Response:** This prints `1 99` because arrays in Go are value types, not reference types. When we do `b := a`, Go creates a complete copy of the array `a` and assigns it to `b`. They are two separate arrays with the same initial values. When we modify `b[0] = 99`, it only changes the copy, not the original. This is different from slices, which are reference types. Arrays being value types means they're copied when passed to functions or assigned to new variables. This is why arrays are rarely used directly in Go - slices are more flexible and efficient for most use cases.

---

### 64. Slice Length and Capacity
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := make([]int, 3, 5)
    fmt.Println(len(s), cap(s))
    s = append(s, 1, 2)
    fmt.Println(len(s), cap(s))
}
```
**A:**
```
3 5
5 5
```
### 1. Initialization

`s := make([]int, 3, 5)` You are creating a slice of integers with:

- **Length (`len`): 3** — The number of elements currently in the slice. These are initialized to their zero value (`0, 0, 0`).
    
- **Capacity (`cap`): 5** — The total number of elements the underlying array can hold before Go needs to allocate a new, larger array.
    

### 2. First Print Statement

`fmt.Println(len(s), cap(s))`

- **Output:** `3 5`
    
- At this point, the slice looks like this: `[0, 0, 0]` (with 2 empty slots waiting in the background).
    

---

### 3. Appending Elements

`s = append(s, 1, 2)` The `append` function adds elements to the end of the slice. Since the current length is 3 and the capacity is 5, there is enough room to add two more numbers without "overflowing" the underlying array.

- The slice now becomes: `[0, 0, 0, 1, 2]`.
    
- The **length** increases to 5.
    
- The **capacity** remains 5.
    

### 4. Second Print Statement

`fmt.Println(len(s), cap(s))`

- **Output:** `5 5`
---

### 65. Append Grows Capacity
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var s []int
    for i := 0; i < 5; i++ {
        s = append(s, i)
        fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
    }
}
```
**A:** Len grows by 1 each time; capacity doubles when exceeded (e.g., `1 1 → 2 2 → 3 4 → 4 4 → 5 8`). Exact values may vary by runtime.

| **Iteration** | **Element Added** | **len** | **cap** | **Why?**                                                        |
| ------------- | ----------------- | ------- | ------- | --------------------------------------------------------------- |
| **1**         | `0`               | 1       | 1       | The nil slice was empty; Go allocated an array of size 1.       |
| **2**         | `1`               | 2       | 2       | Capacity was full, so Go doubled it ($1 \times 2 = 2$).         |
| **3**         | `2`               | 3       | 4       | Capacity was full, so Go doubled it again ($2 \times 2 = 4$).   |
| **4**         | `3`               | 4       | 4       | There was a spare slot in the array of 4, so no new allocation. |
| **5**         | `4`               | 5       | 8       | Capacity was full, so Go doubled it again ($4 \times 2 = 8$).   |

---

### 66. Nil Slice vs Empty Slice
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var s1 []int
    s2 := []int{}
    fmt.Println(s1 == nil)
    fmt.Println(s2 == nil)
    fmt.Println(len(s1), len(s2))
}
```
**A:**
```
true
false
0 0
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between nil slice and empty slice?
**Your Response:** This shows the difference between nil and empty slices. `s1` is a nil slice (declared with `var`), so `s1 == nil` is true. `s2` is an empty slice (created with `[]int{}`), so `s2 == nil` is false. However, both have length 0, which is why `len(s1), len(s2)` prints `0 0`. This distinction is important because a nil slice doesn't have an underlying array, while an empty slice does have an underlying array (just with zero length). In practice, they behave similarly for most operations, but the nil check can be useful to distinguish between "no slice" and "empty slice".

---

### 67. Slice of Slice (Shared Backing Array)
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5}
    b := a[1:3]
    b[0] = 99
    fmt.Println(a)
}
```
**A:** `[1 99 3 4 5]`. Slicing does not copy the underlying array; `b[0]` is the same memory as `a[1]`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why does modifying slice b affect array a?
**Your Response:** This prints `[1 99 3 4 5]` because slices share the same underlying array. When we create `b := a[1:3]`, we're creating a slice that points to elements 1 and 2 of the original array `a`. The slice `b` doesn't copy the data - it just creates a new view into the same underlying array. So when we modify `b[0] = 99`, we're actually modifying `a[1]` because they refer to the same memory location. This is a fundamental Go concept - slices are reference types that provide windows into underlying arrays. This sharing is efficient but can lead to unexpected modifications if you're not careful.

---

### 68. Map Key Existence
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    m := map[string]int{"a": 1}
    v, ok := m["a"]
    fmt.Println(v, ok)
    v2, ok2 := m["b"]
    fmt.Println(v2, ok2)
}
```
**A:**
```
1 true
0 false
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the two-value assignment from a map work?
**Your Response:** This prints `1 true` then `0 false`. When accessing a map in Go with two values, the first is the value and the second is a boolean indicating whether the key exists. For `v, ok := m["a"]`, we get value 1 and `true` because key "a" exists. For `v2, ok2 := m["b"]`, we get the zero value for int (0) and `false` because key "b" doesn't exist. This is the idiomatic way to check if a key exists in a map without causing a panic. The zero value is returned for non-existent keys, so the boolean tells you whether the key was actually found.

---

### 69. Deleting from Map
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    m := map[string]int{"x": 10, "y": 20}
    delete(m, "x")
    delete(m, "z") // deleting non-existent key: no-op
    fmt.Println(len(m))
}
```
**A:** `1`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the delete function work with maps?
**Your Response:** This prints `1`. The `delete` function removes key-value pairs from a map. We start with map `m` containing keys "x" and "y". After `delete(m, "x")`, the key "x" is removed, leaving only "y". When we try to delete a non-existent key with `delete(m, "z")`, nothing happens - it's just a no-op. The `len(m)` returns 1 because only one key remains. This shows that `delete` is safe to use even with non-existent keys - it won't panic or cause errors. Maps in Go are dynamic collections that can grow and shrink at runtime.

---

### 70. Map Cannot Be Compared with ==
**Q: Does this compile?**
```go
package main

func main() {
    m1 := map[string]int{"a": 1}
    m2 := map[string]int{"a": 1}
    _ = m1 == m2
}
```
**A:** **Compile Error.** Maps are not comparable with `==` in Go (only maps can be compared to `nil`). Use `reflect.DeepEqual` instead.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why can't you compare maps with == in Go?
**Your Response:** This doesn't compile because maps cannot be compared using the `==` operator in Go, except for comparison with `nil`. The reason is that maps are reference types and their comparison would be ambiguous - should it compare whether they reference the same underlying data structure, or whether they contain the same key-value pairs? Go avoids this ambiguity by disallowing map comparison entirely. If you need to compare maps for content equality, you should use `reflect.DeepEqual` or iterate through the maps and compare key-value pairs manually. This design choice forces developers to be explicit about what kind of comparison they want.

---

### 71. 2D Slice
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    matrix := [][]int{
        {1, 2, 3},
        {4, 5, 6},
    }
**A:** `2`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do 2D slices work in Go?
**Your Response:** This prints `6`. We create a 2D slice (slice of slices) called `matrix` with two rows: `[1, 2, 3]` and `[4, 5, 6]`. When we access `matrix[1][2]`, we're accessing the second row (index 1) and the third element (index 2) within that row, which is 6. 2D slices in Go are just slices of slices - each row is an independent slice. This is different from 2D arrays where all rows have the same length. With slices, you can have ragged arrays where different rows have different lengths. This flexibility makes 2D slices useful for representing matrices, grids, or other tabular data.

---

### 72. copy() Function
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    src := []int{1, 2, 3}
    dst := make([]int, 2)
    n := copy(dst, src)
    fmt.Println(n, dst)
}
```
**A:** `2 [1 2]`. `copy` copies `min(len(dst), len(src))` elements.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the copy function work with slices?
**Your Response:** This prints `2 [1 2]`. The `copy` function copies elements from a source slice to a destination slice. It copies the minimum of the lengths of both slices - in this case, `min(len(dst), len(src))` which is `min(2, 3) = 2`. So only the first 2 elements from `src` (`[1, 2]`) are copied to `dst`. The function returns the number of elements copied. This is useful when you need to copy a specific number of elements or when the slices might have different lengths. Unlike assignment, `copy` creates an actual copy of the elements, not just another reference to the same underlying array.

---

### 73. Range Over Map Returns Copies
**Q: Does modifying the value in range affect the map?**
```go
package main
import "fmt"

func main() {
    m := map[string]int{"a": 1, "b": 2}
    for k, v := range m {
        v += 10
        _ = k
    }
    fmt.Println(m)
}
```
**A:** `map[a:1 b:2]`. The `v` in range is a copy; modifying it does not affect the original map.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why doesn't modifying the value in a range loop affect the map?
**Your Response:** This prints `map[a:1 b:2]` unchanged. When we range over a map, the value `v` is a copy of the actual value stored in the map, not a reference to it. When we do `v += 10`, we're only modifying the local copy, not the value in the map. To actually modify the map, we would need to access it by key: `m[k] += 10`. This is an important Go concept - range over maps gives you copies of values. This design prevents accidental modification and makes the code clearer about when you're actually modifying the map versus just reading values.

---

### 74. Slice as Map Value
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    m := map[string][]int{}
    m["evens"] = append(m["evens"], 2, 4, 6)
    fmt.Println(m["evens"])
}
```
**A:** `[2 4 6]`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does using a slice as a map value work?
**Your Response:** This prints `[2 4 6]`. We have a map where the values are slices of integers. The key `"evens"` initially has no slice (nil slice). When we do `m["evens"] = append(m["evens"], 2, 4, 6)`, we're appending to the slice at that key. Since `m["evens"]` is initially nil, `append` creates a new slice `[2, 4, 6]` and assigns it to that key. This demonstrates that maps can hold any type as values, including slices, which are themselves reference types. This pattern is useful for grouping related data or building more complex data structures.

---

### 75. Iterating Map Order
**Q: What is guaranteed about the output?**
```go
package main
import "fmt"

func main() {
    m := map[int]string{1: "a", 2: "b", 3: "c"}
    for k, v := range m {
        fmt.Println(k, v)
    }
}
```
**A:** **Nothing is guaranteed.** Map iteration order is intentionally randomized in Go.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why can't you rely on map iteration order in Go?
**Your Response:** The output could be any order and might even be different each time you run the program. Go intentionally randomizes map iteration order to prevent developers from relying on any specific ordering. This design choice was made to avoid subtle bugs where code accidentally depends on implementation details. If you need a specific order, you should extract the keys into a slice, sort them, and then iterate in that order. This randomization applies to all maps - even maps with integer keys won't iterate in numerical order. This forces Go developers to be explicit about ordering when it matters.

---

### 76. Append to Nil Slice
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var s []string
    s = append(s, "hello", "world")
    fmt.Println(s, len(s))
}
```
**A:** `[hello world] 2`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does append work with nil slices?
**Your Response:** This prints `[hello world] 2`. We start with a nil slice `s` (declared with `var`). When we use `append`, Go automatically allocates a new underlying array and creates a proper slice. The `append` function works with nil slices - it just creates a new slice with the appended elements. This is why we don't need to explicitly initialize the slice before appending to it. The resulting slice has length 2 containing the two strings we appended. This is a common Go pattern - start with a nil slice and build it up with append, which is efficient and idiomatic.

---

### 77. Slice Tricks: Remove Element
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5}
    i := 2 // remove index 2
    s = append(s[:i], s[i+1:]...)
    fmt.Println(s)
}
```
**A:** `[1 2 4 5]`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you remove an element from a slice in Go?
**Your Response:** This prints `[1 2 4 5]`. The code removes the element at index 2 (value 3) from the slice. The trick is `s = append(s[:i], s[i+1:]...)`. Here, `s[:i]` gives us elements before index 2 (`[1, 2]`), and `s[i+1:]` gives us elements after index 2 (`[4, 5]`). The `...` operator unpacks the second slice, and `append` combines them. Go doesn't have a built-in `remove` function for slices, so this append-with-two-slices pattern is the idiomatic way to remove elements. It's efficient because it only copies the elements that need to shift.

---

### 78. Array vs Slice Type System
**Q: Does this compile?**
```go
package main
import "fmt"

func printSlice(s []int) { fmt.Println(s) }

func main() {
    arr := [3]int{1, 2, 3}
    printSlice(arr)
}
```
**A:** **Compile Error.** An array `[3]int` cannot be passed where a slice `[]int` is expected. Use `printSlice(arr[:])`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why can't you pass an array to a function expecting a slice?
**Your Response:** This doesn't compile because arrays and slices are different types in Go. An array `[3]int` has a fixed size known at compile time, while a slice `[]int` has a dynamic size determined at runtime. They are incompatible types. The fix is to convert the array to a slice using `arr[:]`, which creates a slice view of the entire array. This distinction is important because arrays are value types (copied when passed) while slices are reference types. Go's type system is strict about this to prevent bugs. Most Go code uses slices rather than arrays because of their flexibility.

---

## Section 7: Structs & Interfaces (Q79–Q91)

### 79. Struct Embedding
**Q: What is the output?**
```go
package main
import "fmt"

type Animal struct{ Name string }
func (a Animal) Speak() string { return a.Name + " speaks" }

type Dog struct{ Animal }

func main() {
    d := Dog{Animal{"Rex"}}
    fmt.Println(d.Speak())
    fmt.Println(d.Name)
}
```
**A:**
```
Rex speaks
Rex
```
Promoted methods and fields are accessible directly on the embedding struct.

---

### 80. Interface Satisfaction
**Q: Does this compile?**
```go
package main
import "fmt"

type Stringer interface {
    String() string
}

type Person struct{ Name string }
func (p Person) String() string { return "Person: " + p.Name }

func main() {
    var s Stringer = Person{"Alice"}
    fmt.Println(s.String())
}
```
**A:** **Yes, compiles and prints** `Person: Alice`. Go uses structural (implicit) interface satisfaction.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does interface satisfaction work in Go?
**Your Response:** This prints `Person: Alice` and demonstrates Go's structural typing for interfaces. The `Stringer` interface defines a method `String() string`. The `Person` struct has a method with exactly that signature, so Go automatically determines that `Person` satisfies the `Stringer` interface. This is called implicit or structural interface satisfaction - you don't need to explicitly declare that a type implements an interface like in some other languages. This makes Go interfaces very flexible and decoupled. The variable `s` of type `Stringer` can hold any value that has a `String()` method, making this a powerful form of polymorphism.

---

### 81. Interface Nil Check
**Q: What is the output?**
```go
package main
import "fmt"

type MyError struct{}
func (e *MyError) Error() string { return "error" }

func getErr() error {
    var e *MyError
    return e
}

func main() {
    err := getErr()
    fmt.Println(err == nil)
}
```
**A:** `false`. The interface value is non-nil because it holds a type descriptor (`*MyError`), even though the underlying pointer is nil.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why is this interface value not nil even though the pointer is nil?
**Your Response:** This prints `false` because of how nil works with interfaces in Go. An interface value has two parts: a type descriptor and a value. In `getErr()`, we declare `var e *MyError` which is a nil pointer, but when we return it as an `error` interface, the interface gets the type descriptor `*MyError` and a nil value. An interface is only nil if both the type and value are nil. Here, the type is non-nil (`*MyError`), so the interface is not nil even though the underlying pointer is nil. This is a subtle Go concept that's important for error handling and interface comparisons.

---

### 82. Struct Tag
**Q: What does this print?**
```go
package main
import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age,omitempty"`
}

func main() {
    u := User{Name: "Bob"}
    b, _ := json.Marshal(u)
    fmt.Println(string(b))
}
```
**A:** `{"name":"Bob"}`. `omitempty` causes the `Age` field (zero value) to be omitted.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do struct tags work in Go?
**Your Response:** This prints `{"name":"Bob"}`. The struct tags `json:"name"` and `json:"age,omitempty"` provide metadata to the `json` package. The `name` tag tells JSON to use "name" as the field name instead of "Name". The `omitempty` tag on Age tells JSON to omit this field entirely if it has its zero value (0 for int). This is why Age doesn't appear in the output. Struct tags are a powerful Go feature for providing metadata to packages. They're commonly used for JSON/XML encoding, validation, database mapping, and more. The tag syntax is backtick strings with key:"value" pairs.

---

### 83. Empty Interface
**Q: What is the output?**
```go
package main
import "fmt"

func printAny(v interface{}) {
    fmt.Printf("%T: %v\n", v, v)
}

func main() {
    printAny(42)
    printAny("hello")
    printAny([]int{1, 2})
}
```
**A:**
```
int: 42
string: hello
[]int: [1 2]
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the empty interface and how does it work?
**Your Response:** This prints the type and value of each argument. The `interface{}` type is the empty interface - it has no methods, so every type in Go satisfies it. This makes it a universal container that can hold any value. When we pass values to `printAny`, they're automatically boxed into interface values. The `%T` verb prints the actual type, and `%v` prints the value. The empty interface is Go's way of handling generic programming before Go introduced generics. It's commonly used in functions like `fmt.Printf` or containers that need to hold different types.

---

### 84. Struct Anonymous Fields Cannot Duplicate
**Q: Does this compile?**
```go
package main

type A struct{ ID int }
type B struct{ ID int }
type C struct {
    A
    B
}

func main() {
    c := C{}
    _ = c.ID
}
```
**A:** **Compile Error.** `c.ID` is ambiguous — both `A.ID` and `B.ID` are promoted. You must use `c.A.ID` or `c.B.ID`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why does this code have a compile error?
**Your Response:** This doesn't compile because of ambiguous field promotion. Both embedded structs `A` and `B` have a field named `ID`. When we embed them in struct `C`, both fields get promoted to `C`, creating ambiguity. The compiler doesn't know whether `c.ID` refers to `A.ID` or `B.ID`, so it gives an error. The fix is to be explicit: use `c.A.ID` or `c.B.ID`. This shows that Go's field promotion is convenient but can create conflicts when multiple embedded types have fields with the same name. It's a design trade-off that encourages explicit code when there's potential ambiguity.

---

### 85. Value vs Pointer Method Sets
**Q: Does this compile?**
```go
package main
import "fmt"

type Counter struct{ n int }
func (c *Counter) Inc() { c.n++ }

type Incrementer interface{ Inc() }

func main() {
    c := Counter{}
    var i Incrementer = c // not addressable
    i.Inc()
    fmt.Println(c.n)
}
```
**A:** **Compile Error.** `Counter` (value) does not implement `Incrementer` because `Inc` has a pointer receiver. Use `var i Incrementer = &c`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why doesn't this value type implement the interface?
**Your Response:** This doesn't compile because of method sets in Go. The `Inc` method has a pointer receiver `(c *Counter)`, which means it belongs to the pointer type `*Counter`, not the value type `Counter`. When we try to assign a `Counter` value to an `Incrementer` interface, Go checks if the value type's method set includes all required methods - but it doesn't, because `Inc` is only on the pointer type. The fix is to use a pointer: `var i Incrementer = &c`. This is a fundamental Go concept - methods with pointer receivers aren't available on values, only on pointers. This affects interface satisfaction and is a common source of confusion for Go developers.

---

### 86. Struct Initialization Positional
**Q: What is the output?**
```go
package main
import "fmt"

type Point struct{ X, Y, Z int }

func main() {
    p := Point{1, 2, 3}
    fmt.Println(p.Y)
}
```
**A:** `2`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does struct positional initialization work?
**Your Response:** This prints `2`. We're initializing a `Point` struct using positional arguments: `Point{1, 2, 3}`. The values are assigned to fields in the order they're declared in the struct definition: `X=1`, `Y=2`, `Z=3`. When we print `p.Y`, we get 2. Positional initialization is concise but can be brittle - if you add or reorder fields in the struct, the initialization code might break. For more robust code, Go recommends using named field initialization like `Point{X: 1, Y: 2, Z: 3}`. However, positional initialization is commonly used for small, simple structs where the field order is stable.

---

### 87. Interfaces Can Hold nil Values
**Q: What is the output?**
```go
package main
import "fmt"

type Describer interface{ Describe() }

func process(d Describer) {
    if d != nil {
        fmt.Println("not nil interface")
    }
}

func main() {
    process(nil)
}
```
**A:** Nothing is printed. Passing a literal `nil` to an interface parameter gives a nil interface (both type and value are nil), so the check `d != nil` is false.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why doesn't anything get printed in this code?
**Your Response:** Nothing prints because we pass `nil` directly to the `process` function, which expects a `Describer` interface. When you pass `nil` directly to an interface parameter, you get a truly nil interface where both the type and value are nil. Inside `process`, the check `d != nil` is false, so the print statement doesn't execute. This is different from the earlier example where we returned a nil pointer as an interface - there the interface had a type but nil value. Here, both type and value are nil, making it a completely nil interface.

---

### 88. Type Assertion Safe Form
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var i interface{} = "hello"
    s, ok := i.(string)
    n, ok2 := i.(int)
    fmt.Println(s, ok)
    fmt.Println(n, ok2)
}
```
**A:**
```
hello true
0 false
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the safe type assertion work in Go?
**Your Response:** This prints `hello true` then `0 false`. The safe type assertion `value, ok := i.(Type)` returns two values: the asserted value and a boolean indicating success. For `s, ok := i.(string)`, since `i` holds a string, it succeeds and returns the string "hello" and `true`. For `n, ok2 := i.(int)`, since `i` doesn't hold an int, it returns the zero value for int (0) and `false`. This is the idiomatic way to do type assertions in Go when you're not sure about the type - it avoids panics that would occur with the single-value form. The boolean tells you whether the assertion succeeded, so you can handle both cases safely.

---

### 89. Struct Pointer Shorthand
**Q: What is the output?**
```go
package main
import "fmt"

type Node struct{ Val int }

func main() {
    n := &Node{Val: 5}
    fmt.Println(n.Val) // auto-dereferenced
}
```
**A:** `5`. Go auto-dereferences struct pointers when accessing fields.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does field access work with struct pointers?
**Your Response:** This prints `5`. We create a pointer to a `Node` struct using `&Node{Val: 5}`. When we access `n.Val`, Go automatically dereferences the pointer for us. We don't need to write `(*n).Val` - Go handles the dereference implicitly. This syntactic sugar makes working with struct pointers much cleaner. It's one of Go's quality-of-life features that makes the language more pleasant to use. The auto-dereference only works for accessing fields and methods - you still need explicit dereference with `*` if you want the whole struct value.

---

### 90. Interface Wrapping
**Q: What is the output?**
```go
package main
import "fmt"

type Inner struct{}
func (Inner) Hello() string { return "inner" }

type Outer struct{ Inner }
func (Outer) Hello() string { return "outer" }

func main() {
    o := Outer{}
    fmt.Println(o.Hello())
    fmt.Println(o.Inner.Hello())
}
```
**A:**
```
outer
inner
```
The `Outer` method `Hello` overrides the promoted `Inner.Hello`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does method overriding work with struct embedding?
**Your Response:** This prints "outer" then "inner". The `Outer` struct embeds `Inner` and also defines its own `Hello` method. When we call `o.Hello()`, it calls `Outer`'s own method, printing "outer". When we call `o.Inner.Hello()`, we're explicitly calling the embedded struct's method, which prints "inner". This shows that a struct's own methods take precedence over promoted methods from embedded types. However, the embedded methods are still accessible if you qualify them with the embedded type name. This is Go's way of providing method overriding while still allowing access to the original methods.

---

### 91. Comparing Nil Interfaces
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var a, b interface{}
    fmt.Println(a == b)
    a = 1
    fmt.Println(a == b)
}
```
**A:**
```
true
false
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How are interface values compared in Go?
**Your Response:** This prints `true` then `false`. When both interface values are nil (both type and value are nil), they are equal, so `a == b` is true. After we assign `a = 1`, `a` now has a type (`int`) and a value (1), while `b` is still nil. Since one has a type and the other doesn't, they're not equal, so `a == b` is false. Interface comparison in Go is only meaningful when both interfaces have the same type - then it compares the underlying values. If the types differ, the comparison is always false (except when both are nil). This is why interface comparison is limited in Go - you can only meaningfully compare interfaces of the same type.

---

## Section 8: Error Handling (Q92–Q96)

### 92. Error Wrapping with fmt.Errorf
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func findUser(id int) error {
    return fmt.Errorf("findUser %d: %w", id, ErrNotFound)
}

func main() {
    err := findUser(42)
    fmt.Println(err)
    fmt.Println(errors.Is(err, ErrNotFound))
}
```
**A:**
```
findUser 42: not found
true
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does error wrapping work with fmt.Errorf?
**Your Response:** This prints the wrapped error message and `true`. The `fmt.Errorf` function with the `%w` verb wraps an existing error with additional context. When we call `findUser(42)`, it wraps `ErrNotFound` with the message "findUser 42: not found". The `errors.Is` function checks if an error (or any wrapped error) matches a specific sentinel error. Since our wrapped error contains `ErrNotFound`, `errors.Is(err, ErrNotFound)` returns true. This is Go's modern error wrapping mechanism introduced in Go 1.13, allowing you to add context while preserving the ability to check for specific errors.

---

### 93. errors.As
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
)

type ValidationError struct{ Field string }
func (e *ValidationError) Error() string { return "invalid: " + e.Field }

func validate() error {
    return fmt.Errorf("wrap: %w", &ValidationError{Field: "email"})
}

func main() {
    err := validate()
    var ve *ValidationError
    if errors.As(err, &ve) {
        fmt.Println("field:", ve.Field)
    }
}
```
**A:** `field: email`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does the errors.As function work?
**Your Response:** This prints `field: email`. The `errors.As` function attempts to unwrap an error and match it to a specific type. We have a `ValidationError` type and we wrap it in `validate()`. In `main`, we declare `var ve *ValidationError` and pass a pointer to it to `errors.As`. The function unwraps the error and, if it finds a `*ValidationError` anywhere in the wrapping chain, it sets `ve` to that error and returns true. This allows us to access the original error type and its fields. Unlike `errors.Is` which checks for exact error matches, `errors.As` checks for type compatibility and gives you access to the underlying error.

---

### 94. Ignoring Errors
**Q: What is the bug?**
```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    n, _ := strconv.Atoi("abc")
    fmt.Println(n * 2)
}
```
**A:** Output is `0`. There's no panic, but the error from `Atoi` is silently ignored. `n` defaults to `0` on failure, leading to silently wrong results.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's wrong with ignoring errors in Go?
**Your Response:** This prints `0` but demonstrates a serious bug. The `strconv.Atoi("abc")` fails because "abc" isn't a valid number, returning an error and 0 as the default value. By using the blank identifier `_` to ignore the error, we silently continue with the wrong result. The program then calculates `0 * 2 = 0`. This is why Go's philosophy is "errors are values" - you should always handle errors explicitly. Ignoring errors can lead to subtle bugs that are hard to diagnose. The fix is to check the error and handle it appropriately, either by returning it, logging it, or handling the failure case.

---

### 95. Custom Error Type
**Q: What is the output?**
```go
package main
import "fmt"

type AppError struct {
    Code    int
    Message string
}

func (e AppError) Error() string {
    return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func riskyOp() error {
    return AppError{Code: 404, Message: "resource not found"}
}

func main() {
    err := riskyOp()
    fmt.Println(err)
}
```
**A:** `[404] resource not found`

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do custom error types work in Go?
**Your Response:** This prints `[404] resource not found`. We define a custom error type `AppError` with `Code` and `Message` fields. By implementing the `Error() string` method, it satisfies the `error` interface. When we return `AppError{Code: 404, Message: "resource not found"}` from `riskyOp()`, it can be used as any other error. The `Error()` method formats the error message with the code and message. Custom error types are useful when you want to provide structured error information that callers can inspect programmatically, like checking error codes or accessing additional context. This is more expressive than just returning string errors.

---

### 96. Sentinel Error Comparison
**Q: What is the output?**
```go
package main
import (
    "errors"
    "fmt"
    "io"
)

func main() {
    err := io.EOF
    fmt.Println(err == io.EOF)
    fmt.Println(errors.Is(err, io.EOF))
}
```
**A:**
```
true
true
```

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do sentinel errors work in Go?
**Your Response:** This prints `true` twice. Sentinel errors are predefined error values that are used to indicate specific error conditions. `io.EOF` is a classic example - it represents the end of a file. The first comparison `err == io.EOF` works because we're comparing the exact same error value. The second comparison `errors.Is(err, io.EOF)` also returns true because `errors.Is` checks if the error matches the sentinel, including through wrapping. Sentinel errors are useful for checking specific error conditions, but they should be used sparingly and only for errors that are truly exceptional and need to be checked specifically. For most cases, error wrapping is preferred.

---

## Section 9: Goroutines Basics & Misc (Q97–Q100)

### 97. Goroutine Without Wait
**Q: Why might this print nothing?**
```go
package main
import "fmt"

func main() {
    go func() {
        fmt.Println("hello from goroutine")
    }()
}
```
**A:** `main` returns before the goroutine gets to execute. Since `main` exiting terminates the program, the goroutine may never run. Use `time.Sleep` or `sync.WaitGroup` to wait.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Why might this goroutine code print nothing?
**Your Response:** This might print nothing because the `main` function can exit before the goroutine gets a chance to run. When we launch a goroutine with `go func()`, it runs concurrently. But the `main` function doesn't wait for it - it just continues and exits immediately. When `main` exits, the entire program terminates, killing any running goroutines. This is a common Go beginner mistake. The fix is to use synchronization primitives like `sync.WaitGroup` to wait for the goroutine to finish, or use `time.Sleep` for simple cases. This demonstrates that goroutines don't make the main program wait - you need explicit coordination.

---

### 98. Goroutine Count
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println(runtime.NumGoroutine())
}
```
**A:** `1`. After `wg.Wait()`, all 5 goroutines have finished, leaving only the main goroutine running.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How does runtime.NumGoroutine work?
**Your Response:** This prints `1`. We launch 5 goroutines, each of which calls `wg.Done()` and then exits. The `wg.Wait()` blocks until all goroutines have finished. After `wg.Wait()` returns, all 5 goroutines are no longer running - they've completed and been cleaned up. Only the main goroutine is still running, so `runtime.NumGoroutine()` returns 1. This function is useful for debugging and monitoring goroutine leaks - if you see the goroutine count growing over time, you might have goroutines that aren't exiting properly. It's a handy tool for understanding the concurrent behavior of your Go programs.

---

### 99. GOMAXPROCS
**Q: What does this print and what does it mean?**
```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.GOMAXPROCS(0))
}
```
**A:** Prints the number of logical CPUs available (e.g., `8`). Passing `0` queries the current value without changing it. By default, Go uses all available CPUs.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does GOMAXPROCS do in Go?
**Your Response:** This prints the number of logical CPUs available to the Go runtime (like 8 on a modern machine). `GOMAXPROCS` controls how many operating system threads can execute Go code simultaneously. Passing `0` queries the current setting without changing it. By default, Go sets `GOMAXPROCS` to the number of available CPU cores, allowing your program to use all CPU power for parallel execution. In older Go versions, you had to manually set this, but since Go 1.5, it defaults to all available CPUs. You might adjust this for specific performance tuning, but the default is usually optimal for most applications.

---

### 100. Channel Direction in Function Signature
**Q: Does this compile, and what is the benefit?**
```go
package main
import "fmt"

func producer(ch chan<- int) {
    ch <- 42
}

func consumer(ch <-chan int) {
    fmt.Println(<-ch)
}

func main() {
    ch := make(chan int, 1)
    producer(ch)
    consumer(ch)
}
```
**A:** **Yes, compiles and prints** `42`. Directional channels (`chan<-` send-only, `<-chan` receive-only) restrict how a channel is used in a function, improving type safety and self-documenting intent.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do directional channels work in Go?
**Your Response:** This prints `42` and demonstrates directional channels. The `producer` function takes a `chan<- int` (send-only channel), so it can only send to the channel, not receive from it. The `consumer` function takes a `<-chan int` (receive-only channel), so it can only receive from it, not send. This provides type safety - the compiler will prevent you from accidentally receiving in the producer or sending in the consumer. It also serves as documentation - you can immediately see how each function uses the channel. This is especially useful in complex concurrent programs where clear communication patterns are important. Directional channels make your concurrent code safer and more readable.

---

*End of 100 Go Basics & Fundamentals Code Snippet Questions*
