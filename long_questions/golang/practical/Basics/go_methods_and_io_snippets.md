# Go — Method Sets, io Patterns & Functional Options Snippets

> **Format**: Each question is "predict the output / spot the bug / does it compile?" style.
> **Topics**: Value vs pointer receivers · Method sets · Interface satisfaction rules · `io.Reader`/`io.Writer` patterns · Functional options · Channel composition patterns

---

## 📋 Reading Progress

> Mark each section `[x]` when done. Use `🔖` to note where you left off.

- [ ] **Section 1:** Value Receivers vs Pointer Receivers (Q1–Q15)
- [ ] **Section 2:** io.Reader / io.Writer Patterns (Q16–Q38)
- [ ] **Section 3:** Functional Options Pattern (Q39–Q50)
- [ ] **Section 4:** Channel Composition Patterns (Q51–Q65)

> 🔖 **Last read:** <!-- e.g. Q15 · Section 1 done -->

---

## Section 1: Value Receivers vs Pointer Receivers (Q1–Q15)

### 1. Value Receiver — Called on Value
**Q: Does this compile and what is the output?**
```go
package main
import "fmt"
==
type Point struct{ X, Y int }

func (p Point) String() string {
    return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func main() {
    p := Point{3, 4}
    fmt.Println(p.String())
}
```
**A:** **Yes.** `(3,4)`. Value receiver — called on a value. Clean: the method receives a copy.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile and what is the output?
**Your Response:** Yes, this compiles and prints `(3,4)`. This is a value receiver method - it receives a copy of the Point value. Since we're calling it on a Point value, it works perfectly. Value receivers are used when the method doesn't need to modify the original value.

---

### 2. Pointer Receiver — Modifies the Original
**Q: What is the output?**
```go
package main
import "fmt"

type Counter struct{ n int }

func (c *Counter) Inc() { c.n++ }
func (c Counter) Val() int { return c.n }

func main() {
    c := Counter{}
    c.Inc()
    c.Inc()
    fmt.Println(c.Val())
}
```
**A:** `2`. `Inc` has a pointer receiver — it modifies `c` in place. `Val` has a value receiver — it reads a copy.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `2`. The `Inc` method has a pointer receiver, so it modifies the original Counter struct. The `Val` method has a value receiver but that's fine since it's just reading the value. This shows how you can mix receiver types - use pointer receivers when you need to modify, value receivers for read-only operations.

---

### 3. Value Receiver Does NOT Modify the Original
**Q: What is the output?**
```go
package main
import "fmt"

type Counter struct{ n int }

func (c Counter) Inc() { c.n++ } // value receiver — modifies a copy

func main() {
    c := Counter{}
    c.Inc()
    c.Inc()
    fmt.Println(c.n)
}
```
**A:** `0`. `Inc` operates on a copy of `c`. The original is unchanged.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `0`. The `Inc` method has a value receiver, which means it receives a copy of the Counter. When it increments `c.n`, it's modifying the copy, not the original. This is why `c.n` is still 0 after calling `Inc` twice. Value receivers can't modify the original value.

---

### 4. Auto-Dereference: Pointer Method on Addressable Value
**Q: Does this compile?**
```go
package main
import "fmt"

type T struct{ v int }
func (t *T) Double() { t.v *= 2 }

func main() {
    t := T{v: 5}
    t.Double() // t is addressable, Go auto-takes &t
    fmt.Println(t.v)
}
```
**A:** **Yes.** Output: `10`. Go automatically rewrites `t.Double()` as `(&t).Double()` if `t` is addressable.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** Yes, this compiles and prints `10`. Go has automatic dereferencing - when you call a pointer receiver method on an addressable value, Go automatically takes the address of the value for you. So `t.Double()` is automatically rewritten as `(&t).Double()`. This makes the code cleaner and more intuitive.

---

### 5. Pointer Method on Non-Addressable — Compile Error
**Q: Does this compile?**
```go
package main

type T struct{ v int }
func (t *T) Double() { t.v *= 2 }

func main() {
    T{v: 5}.Double() // T{v:5} is a non-addressable temporary
}
```
**A:** **Compile Error.** `cannot take the address of T{...}`. Composite literals used directly as expressions are not addressable. Assign to a variable first.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. The error is `cannot take the address of T{...}`. Composite literals like `T{v: 5}` are not addressable when used directly as expressions. Go can't take the address of a temporary value. You need to assign it to a variable first, then call the method on the variable.

---

### 6. The Method Set Rule
**Q: Which types satisfy the `Stringer` interface?**
```go
type Stringer interface{ String() string }

type T struct{}
func (t T) String() string  { return "value" }  // value receiver

type U struct{}
func (u *U) String() string { return "pointer" } // pointer receiver
```
**A:**
- `T` satisfies `Stringer` ✅ (value receiver — both `T` and `*T` satisfy)
- `*T` satisfies `Stringer` ✅
- `U` does **NOT** satisfy `Stringer` ❌ (pointer receiver — only `*U` satisfies)
- `*U` satisfies `Stringer` ✅

**Rule:** A type `T`'s method set contains only value-receiver methods. `*T`'s method set contains both.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which types satisfy the Stringer interface?
**Your Response:** Both `T` and `*T` satisfy the Stringer interface because `T` has a value receiver method. Only `*U` satisfies Stringer, not `U` itself, because `U` has a pointer receiver method. The rule is: the method set of a type includes only value receiver methods, while the method set of a pointer to that type includes both value and pointer receiver methods.

---

### 7. Pointer Receiver — Only *T Satisfies the Interface
**Q: Does this compile?**
```go
package main

type Saver interface{ Save() error }

type DB struct{}
func (d *DB) Save() error { return nil }

func process(s Saver) {}

func main() {
    db := DB{}
    process(db)  // passing value, not pointer
}
```
**A:** **Compile Error.** `DB does not implement Saver (Save method has pointer receiver)`. Fix: `process(&db)`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. The error is `DB does not implement Saver`. Even though `DB` has a `Save` method, it has a pointer receiver, so only `*DB` satisfies the interface, not `DB` itself. When we pass `db` (a value) to `process`, it doesn't satisfy the interface. The fix is to pass `&db` instead.

---

### 8. Value Receiver — Both T and *T Satisfy
**Q: Does this compile?**
```go
package main

type Printer interface{ Print() }

type Doc struct{ text string }
func (d Doc) Print() { println(d.text) }

func show(p Printer) { p.Print() }

func main() {
    d := Doc{text: "hello"}
    show(d)   // value — OK
    show(&d)  // pointer — also OK
}
```
**A:** **Yes, both compile.** Value receivers are in the method set of both `Doc` and `*Doc`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** Yes, both calls compile. When a method has a value receiver, it's included in the method set of both the value type and the pointer type. So both `Doc` and `*Doc` satisfy the `Printer` interface. This is why you can call `show(d)` with a value and `show(&d)` with a pointer - both work.

---

### 9. Interface Holding a Non-Pointer Value — Cannot Call Pointer Method
**Q: Does this compile?**
```go
package main

type Mover interface{ Move() }

type Car struct{ speed int }
func (c *Car) Move() { c.speed++ }

func main() {
    var m Mover = Car{speed: 0} // store value, not pointer
    m.Move()
}
```
**A:** **Compile Error.** `Car does not implement Mover (Move method has pointer receiver)`. An interface value holding a `Car` (value) cannot auto-take its address, because the interface slot doesn't expose an addressable memory location. Fix: `var m Mover = &Car{speed: 0}`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. The issue is that we're storing a `Car` value (not a pointer) in an interface variable, but the `Move` method has a pointer receiver. Once a value is stored in an interface, Go can't take its address anymore. The fix is to store a pointer: `var m Mover = &Car{speed: 0}`.

---

### 10. Pointer to Interface — Almost Never Correct
**Q: Does this compile?**
```go
package main

type Doer interface{ Do() }
type Task struct{}
func (t Task) Do() { println("done") }

func run(d *Doer) { (*d).Do() } // *Doer is a pointer to interface

func main() {
    var d Doer = Task{}
    run(&d)
}
```
**A:** **Compiles but is a code smell.** A `*Doer` (pointer to interface) is almost never what you want in Go. Pass the interface by value directly — interfaces are already reference types.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** Yes, this compiles but it's a code smell. You're taking a pointer to an interface (`*Doer`), which is almost never needed in Go. Interfaces are already reference types under the hood - they contain a pointer to the underlying value and type information. Just pass interfaces by value, not by pointer.

---

### 11. Method on Map Value — Not Addressable
**Q: Does this compile?**
```go
package main

type Counter struct{ n int }
func (c *Counter) Inc() { c.n++ }

func main() {
    m := map[string]Counter{"a": {}}
    m["a"].Inc() // can we call pointer method on map value?
}
```
**A:** **Compile Error.** `cannot take the address of m["a"]`. Map values are not addressable. Fix: use `map[string]*Counter` or copy the value, modify, and reassign.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** No, this doesn't compile. Map values are not addressable in Go, so you can't take their address to call pointer receiver methods on them. The error is `cannot take the address of m["a"]`. You either need to store pointers in the map (`map[string]*Counter`) or copy the value out, modify it, and put it back in the map.

---

### 12. Receiver Name Convention
**Q: What is the output, and what is the convention?**
```go
package main
import "fmt"

type Server struct{ addr string }

func (s *Server) Start() { fmt.Println("starting:", s.addr) }

func main() {
    srv := &Server{addr: ":8080"}
    srv.Start()
}
```
**A:** `starting: :8080`. Convention: receiver name should be 1-2 letters abbreviating the type name (e.g., `s` for `Server`, `c` for `Client`). Never `self` or `this` — those are not idiomatic Go.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output, and what is the convention?
**Your Response:** This prints `starting: :8080`. The Go convention for receiver names is to use 1-2 letter abbreviations of the type name - like `s` for Server or `c` for Client. Never use `self` or `this` like in other languages - these are not idiomatic in Go. The receiver name should be consistent and short.

---

### 13. When to Use Pointer Receiver
**Q: Which of these MUST use a pointer receiver?**
```go
// Case A: large struct — avoid copy overhead
func (b BigStruct) Process() { ... }

// Case B: method modifies the receiver
func (c *Counter) Reset() { c.n = 0 }

// Case C: mutex field — must NOT copy
type SafeMap struct { mu sync.Mutex; ... }
func (s *SafeMap) Get() { ... }

// Case D: implements an interface where other methods are pointer receivers
// (consistency rule: if any method is pointer, all should be)
```
**A:** B is *required* (must modify). C is *required* (mutexes must never be copied). A and D are *strongly recommended*. Rule of thumb: if in doubt with a struct, use pointer receivers.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Which of these MUST use a pointer receiver?
**Your Response:** Cases B and C are required. Case B needs a pointer receiver to modify the struct. Case C is critical because mutexes contain internal state and must never be copied - copying a mutex can cause deadlocks. Cases A and D are strongly recommended for performance and consistency. The rule of thumb: when in doubt with structs, use pointer receivers.

---

### 14. Nil Pointer Receiver — Valid Call (No Auto-Deref Needed)
**Q: What is the output?**
```go
package main
import "fmt"

type Node struct{ val int }

func (n *Node) Val() int {
    if n == nil { return 0 }
    return n.val
}

func main() {
    var n *Node
    fmt.Println(n.Val())
}
```
**A:** `0`. Calling a method on a nil pointer is valid if the method handles nil explicitly. This is a useful pattern for linked-list/tree node types.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `0`. In Go, you can call methods on nil pointers, and it's valid as long as the method handles the nil case. This pattern is commonly used for linked lists and trees where a nil node represents an empty list or leaf. The method checks `if n == nil` and returns a default value instead of panicking.

---

### 15. Consistency Rule — Mix of Pointer and Value Receivers
**Q: What is the issue?**
```go
type File struct{ name string }

func (f File) Name() string   { return f.name }      // value
func (f *File) Close() error  { return nil }          // pointer

var _ io.Closer = File{}  // File doesn't have Close in its method set!
```
**A:** `File` does not satisfy `io.Closer` — only `*File` does (pointer receiver). Mixing value and pointer receivers on the same type means only `*File` has the full method set. **Recommendation:** be consistent — if any method uses a pointer receiver, use pointer receivers for all.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the issue?
**Your Response:** The issue is that `File` doesn't satisfy the `io.Closer` interface because `Close` has a pointer receiver but `Name` has a value receiver. This mixing means only `*File` has the complete method set. The recommendation is to be consistent - if any method needs a pointer receiver, use pointer receivers for all methods of that type.

---

## Section 2: io.Reader / io.Writer Patterns (Q16–Q38)

### 16. io.Reader Interface
**Q: What is the io.Reader interface?**
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```
**A:** `Read` fills `p` with up to `len(p)` bytes, returns how many were read (`n`) and an error. When the stream ends, it returns `n=0, err=io.EOF`. Anything can implement `io.Reader`: files, network connections, byte buffers, strings.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the io.Reader interface?
**Your Response:** The `io.Reader` interface has just one method: `Read(p []byte) (n int, err error)`. It fills the byte slice `p` with data from the source, returning how many bytes were read and any error. When there's no more data, it returns `io.EOF`. Anything can implement this interface - files, network connections, strings, byte buffers - making it a very flexible abstraction for reading data.

---

### 17. Basic io.Reader — Reading a String
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Hello, Go!")
    buf := make([]byte, 5)
    for {
        n, err := r.Read(buf)
        fmt.Printf("%q ", buf[:n])
        if err == io.EOF { break }
    }
}
```
**A:** `"Hello" ", Go!" ""` *(last read may return 0 bytes + EOF)*
Actually: `"Hello" ", Go!"`. `Read` returns bytes until exhausted, then `io.EOF`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `"Hello" ", Go!"`. `strings.NewReader` creates an `io.Reader` from a string. We read in a loop, each time reading up to 5 bytes into our buffer. The loop continues until `Read` returns an error. When the string is exhausted, `Read` returns `io.EOF` and we break. The last read might return fewer bytes than the buffer size.

---

### 18. io.ReadAll — Read Entire Reader
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Hello, io.ReadAll!")
    data, err := io.ReadAll(r)
    fmt.Println(string(data), err)
}
```
**A:** `Hello, io.ReadAll! <nil>`. `io.ReadAll` reads until `io.EOF` and returns all data as a `[]byte`. It never returns `io.EOF` as an error — only real errors.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `Hello, io.ReadAll! <nil>`. `io.ReadAll` is a convenience function that reads from an `io.Reader` until `io.EOF` and returns all the data as a byte slice. It handles the reading loop for you and never returns `io.EOF` as an error - only actual read errors. This is perfect for when you need to read an entire response or file into memory.

---

### 19. io.Writer Interface
**Q: What is the io.Writer interface?**
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
**A:** `Write` writes `len(p)` bytes from `p`. Returns how many bytes were written and any error. Implementations: `os.Stdout`, `os.File`, `bytes.Buffer`, `http.ResponseWriter`, network conns.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the io.Writer interface?
**Your Response:** The `io.Writer` interface has one method: `Write(p []byte) (n int, err error)`. It writes the bytes from slice `p` to the destination, returning how many bytes were actually written and any error. Many types implement this - files, network connections, HTTP responses, byte buffers. It's the counterpart to `io.Reader` for writing data.

---

### 20. io.Copy — From Reader to Writer
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "io"
    "os"
    "strings"
)

func main() {
    src := strings.NewReader("Copying data!")
    n, err := io.Copy(os.Stdout, src)
    fmt.Println()
    fmt.Println(n, err)
}
```
**A:**
```
Copying data!
13 <nil>
```
`io.Copy` reads from `src` and writes to `dst` in a loop using an internal buffer. Returns total bytes copied.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `Copying data!` then `13 <nil>`. `io.Copy` efficiently copies data from a reader to a writer. It uses an internal 32KB buffer and handles the read/write loop for you. It returns the total number of bytes copied and any error. This is the idiomatic way to copy streams in Go - much simpler than writing your own loop.

---

### 21. io.TeeReader — Read and Simultaneously Write
**Q: What is the output?**
```go
package main
import (
    "bytes"
    "fmt"
    "io"
    "strings"
)

func main() {
    src := strings.NewReader("Hello Tee!")
    var log bytes.Buffer
    tee := io.TeeReader(src, &log)

    data, _ := io.ReadAll(tee)
    fmt.Println("read:", string(data))
    fmt.Println("log:", log.String())
}
```
**A:**
```
read: Hello Tee!
log: Hello Tee!
```
`TeeReader` wraps a reader: every byte read from `tee` is also written to `&log`. Like a T-junction pipe.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `read: Hello Tee!` then `log: Hello Tee!`. `io.TeeReader` creates a reader that, when read from, also writes the data to a writer. It's like a T-junction in plumbing - data flows both to the reader and to the writer. This is perfect for when you need to read data but also log or save it simultaneously.

---

### 22. io.MultiWriter — Write to Multiple Destinations
**Q: What is the output?**
```go
package main
import (
    "bytes"
    "fmt"
    "io"
    "os"
)

func main() {
    var buf bytes.Buffer
    mw := io.MultiWriter(os.Stdout, &buf)
    fmt.Fprintln(mw, "Hello MultiWriter!")
    fmt.Println("buf:", buf.String())
}
```
**A:**
```
Hello MultiWriter!
buf: Hello MultiWriter!

```
`MultiWriter` fans out writes to all given writers simultaneously — perfect for logging to both stdout and a file.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `Hello MultiWriter!` twice - once to stdout and once showing in the buffer. `io.MultiWriter` creates a writer that duplicates all writes to multiple destination writers. It's perfect for scenarios like logging to both stdout and a file simultaneously, or writing to multiple network connections at once.

---

### 23. io.LimitReader — Read at Most N Bytes
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Hello, World!")
    limited := io.LimitReader(r, 5)
    data, _ := io.ReadAll(limited)
    fmt.Println(string(data))
}
```
**A:** `Hello`. `LimitReader` wraps a reader and returns at most N bytes — useful for safely reading untrusted/large input.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `Hello`. `io.LimitReader` wraps another reader and limits how many bytes can be read from it. Even though the original string has 13 characters, we limited it to 5 bytes, so only "Hello" is read. This is great for safely reading from potentially large or untrusted sources where you want to enforce a maximum size.

---

### 24. io.MultiReader — Concatenate Readers
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r1 := strings.NewReader("Hello, ")
    r2 := strings.NewReader("World!")
    mr := io.MultiReader(r1, r2)
    data, _ := io.ReadAll(mr)
    fmt.Println(string(data))
}
```
**A:** `Hello, World!`. `MultiReader` chains multiple readers into one — reading from them sequentially.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `Hello, World!`. `io.MultiReader` combines multiple readers into a single reader that reads them sequentially. When you read from the multi-reader, it reads from the first reader until it's exhausted, then automatically moves to the second, and so on. This is useful for concatenating multiple streams without copying all the data into memory first.

---

### 25. io.Pipe — Synchronous In-Memory Pipe
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "io"
)

func main() {
    pr, pw := io.Pipe()

    go func() {
        fmt.Fprintln(pw, "data through pipe")
        pw.Close()
    }()

    data, _ := io.ReadAll(pr)
    fmt.Print(string(data))
}
```
**A:** `data through pipe`. `io.Pipe` creates a synchronous, in-memory pipe connecting a writer and reader. The writer blocks until the reader consumes. Used to adapt APIs that need a Writer with those that need a Reader.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `data through pipe`. `io.Pipe` creates a synchronous in-memory pipe with a reader and writer end. When you write to the pipe writer, it blocks until someone reads from the pipe reader. This is perfect for connecting APIs - when one function expects to write data and another expects to read it, you can connect them with a pipe.

---

### 26. io.ReadFull — Exactly N Bytes or Error
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Hi")
    buf := make([]byte, 5)
    n, err := io.ReadFull(r, buf)
    fmt.Println(n, err)
}
```
**A:** `2 unexpected EOF`. `ReadFull` reads exactly `len(buf)` bytes. If fewer are available, it returns `io.ErrUnexpectedEOF` (not `io.EOF`). Returns `io.EOF` only if zero bytes were read.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `2 unexpected EOF`. `io.ReadFull` tries to read exactly the number of bytes specified by the buffer length. Since the string only has 2 bytes but we're trying to read 5, it returns `io.ErrUnexpectedEOF` after reading the available bytes. This is different from `io.EOF` which is returned only when zero bytes are read.

---

### 27. io.Discard — Write and Throw Away
**Q: What does this do?**
```go
package main
import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("throw this away")
    n, _ := io.Copy(io.Discard, r)
    fmt.Println("discarded", n, "bytes")
}
```
**A:** `discarded 15 bytes`. `io.Discard` is a writer that silently discards all data (like `/dev/null`). Useful for draining a response body or benchmarking read throughput.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this do?
**Your Response:** This prints `discarded 15 bytes`. `io.Discard` is an `io.Writer` that silently discards everything written to it, like `/dev/null` on Unix systems. It's useful when you need to read data from a source but don't want to keep it - like draining an HTTP response body or benchmarking how fast you can read from a source.

---

### 28. bufio.NewReader — Buffered Reading
**Q: What is the output?**
```go
package main
import (
    "bufio"
    "fmt"
    "strings"
)

func main() {
    r := bufio.NewReader(strings.NewReader("Hello\nWorld\n"))
    line1, _ := r.ReadString('\n')
    line2, _ := r.ReadString('\n')
    fmt.Printf("%q %q", line1, line2)
}
```
**A:** `"Hello\n" "World\n"`. `bufio.NewReader` wraps any `io.Reader` with a buffer. `ReadString(delim)` reads until (and including) the delimiter.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `"Hello\n" "World\n"`. `bufio.NewReader` adds buffering to any reader, which makes reading more efficient. `ReadString('\n')` reads until it finds a newline character, including the newline in the result. This is much more efficient than reading byte by byte, especially for network streams or files.

---

### 29. bufio.Writer — Must Call Flush
**Q: What is the bug?**
```go
package main
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    w := bufio.NewWriter(os.Stdout)
    fmt.Fprintln(w, "Buffered write")
    // BUG: forgot w.Flush() — data may never reach os.Stdout
}
```
**A:** Output may not appear or may be incomplete. `bufio.Writer` accumulates writes in a buffer for efficiency. **Always call `w.Flush()`** (or `defer w.Flush()`) to ensure buffered data is written to the underlying writer.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the bug?
**Your Response:** The bug is that we forgot to call `w.Flush()`. `bufio.Writer` buffers writes in memory for efficiency, but the buffered data isn't actually written to the underlying writer (stdout in this case) until you call `Flush()`. Without flushing, the data might never appear or might be incomplete. Always call `defer w.Flush()` when using buffered writers.

---

### 30. Implementing a Custom io.Reader
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "io"
)

type RepeatReader struct {
    data []byte
    pos  int
    max  int
}

func (r *RepeatReader) Read(p []byte) (int, error) {
    if r.pos >= r.max { return 0, io.EOF }
    n := copy(p, r.data[r.pos%len(r.data):])
    r.pos += n
    return n, nil
}

func main() {
    rr := &RepeatReader{data: []byte("AB"), max: 6}
    data, _ := io.ReadAll(rr)
    fmt.Println(string(data))
}
```
**A:** `ABABAB`. Implementing `io.Reader` requires only one method — `Read`. The struct tracks position to return repeatng data.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `ABABAB`. We've implemented a custom `io.Reader` that repeats the pattern "AB". The `Read` method copies data from our internal slice, advances a position counter, and returns `io.EOF` when we've read the maximum number of bytes. This shows how easy it is to create custom readers - just implement the `Read` method.

---

### 31. http.Response.Body — io.ReadCloser
**Q: What is the correct pattern?**
```go
resp, err := http.Get("https://example.com")
if err != nil {
    return err
}
defer resp.Body.Close()           // ← always close
body, err := io.ReadAll(resp.Body)
if err != nil {
    return err
}
fmt.Println(string(body))
```
**A:** `resp.Body` is an `io.ReadCloser` (implements both `io.Reader` and `io.Closer`). Always `Close()` it regardless of whether you read it — otherwise goroutines and TCP connections leak.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the correct pattern?
**Your Response:** The correct pattern is to always call `defer resp.Body.Close()` immediately after checking for errors. `resp.Body` is an `io.ReadCloser` - it's both a reader and has a `Close` method. If you don't close it, you leak both the TCP connection and a goroutine that manages the connection. Even if you don't read the body, you should still close it.

---

### 32. io.ReadAtLeast
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Hi")
    buf := make([]byte, 10)
    n, err := io.ReadAtLeast(r, buf, 5)
    fmt.Println(n, err)
}
```
**A:** `2 short buffer`. `ReadAtLeast` reads at least `min` bytes. If fewer bytes are available, it returns `io.ErrUnexpectedEOF`. If `len(buf) < min`, it returns `io.ErrShortBuffer`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `2 short buffer`. `io.ReadAtLeast` tries to read at least a minimum number of bytes. Here we're asking for at least 5 bytes, but the buffer is only 10 bytes and the source only has 2 bytes. Since the buffer is too small for the minimum, it returns `io.ErrShortBuffer`. This is different from `ReadFull` which would return `io.ErrUnexpectedEOF`.

---

### 33. Wrapping an io.Writer — Middleware Pattern
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "io"
    "os"
    "strings"
)

type UpperWriter struct{ w io.Writer }

func (u UpperWriter) Write(p []byte) (int, error) {
    return u.w.Write([]byte(strings.ToUpper(string(p))))
}

func main() {
    uw := UpperWriter{w: os.Stdout}
    fmt.Fprintln(uw, "hello world")
}
```
**A:** `HELLO WORLD`. Wrapping an `io.Writer` is the standard middleware pattern for transforming a stream — used in compression, encryption, logging, etc.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `HELLO WORLD`. We've created a custom writer that transforms text to uppercase before passing it to the underlying writer. This is the middleware pattern - wrapping an `io.Writer` to add functionality. It's commonly used for compression, encryption, logging, or any transformation of data being written. The pattern is powerful because it works with any writer.

---

### 34. io.Copy with io.Writer — Efficient File Copy
**Q: What is the pattern for efficient file-to-file copy?**
```go
func copyFile(dst, src string) error {
    in, err := os.Open(src)
    if err != nil { return err }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil { return err }
    defer out.Close()

    _, err = io.Copy(out, in)
    return err
}
```
**A:** `io.Copy` uses an internal 32KB buffer — no manual chunking needed. It calls `WriteTo`/`ReadFrom` optimised paths if available (e.g., `sendfile` syscall for files on Linux).

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the pattern for efficient file-to-file copy?
**Your Response:** `io.Copy` is the most efficient way to copy files. It uses a 32KB internal buffer and automatically uses optimized system calls like `sendfile` on Linux when available. You don't need to manually chunk the data or worry about buffer sizes - `io.Copy` handles all of that for you and is highly optimized.

---

### 35. io.SectionReader — Read a Portion of a File
**Q: What does this do?**
```go
// Read 20 bytes starting at offset 10 of a file:
sr := io.NewSectionReader(file, 10, 20)
data, _ := io.ReadAll(sr)
```
**A:** `io.SectionReader` implements `io.Reader`, `io.Seeker`, `io.ReaderAt` on a section of a `ReaderAt`. Useful for random-access reads without seeking the underlying resource.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does this do?
**Your Response:** `io.SectionReader` lets you read a specific portion of a larger file or reader as if it were a separate file. It implements multiple interfaces (`Reader`, `Seeker`, `ReaderAt`) on just that section. This is perfect for reading parts of large files without having to seek the underlying file, or for implementing file formats where different sections contain different data.

---

### 36. io.WriteString — Optimised String Writing
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "io"
    "os"
)

func main() {
    n, err := io.WriteString(os.Stdout, "Hello WriteString!\n")
    fmt.Println(n, err)
}
```
**A:**
```
Hello WriteString!
18 <nil>
```
`io.WriteString` checks if the writer implements `io.StringWriter` (which has `WriteString(string)`) and calls it directly — avoiding a `[]byte` conversion. More efficient than `w.Write([]byte(s))`.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `Hello WriteString!` then `18 <nil>`. `io.WriteString` is an optimized way to write strings to writers. It checks if the writer implements the `io.StringWriter` interface and calls the more efficient `WriteString` method if available, avoiding the overhead of converting the string to a byte slice. This is why it's preferred over `w.Write([]byte(s))`.

---

### 37. bufio.Scanner — Line-by-Line, With Custom Split
**Q: What is the output?**
```go
package main
import (
    "bufio"
    "fmt"
    "strings"
)

func main() {
    r := strings.NewReader("one two three")
    sc := bufio.NewScanner(r)
    sc.Split(bufio.ScanWords)
    for sc.Scan() {
        fmt.Println(sc.Text())
    }
}
```
**A:**
```
one
two
three
```
`bufio.Scanner` defaults to line splitting. `Split(bufio.ScanWords)` changes it to word splitting. Other options: `ScanBytes`, `ScanRunes`, or a custom split function.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `one`, `two`, `three` on separate lines. `bufio.Scanner` provides a convenient way to read data token by token. By default it splits on lines, but `scanner.Split(bufio.ScanWords)` changes it to split on words instead. This is much easier than manually reading and parsing text - the scanner handles all the buffering and tokenization for you.

---

### 38. io.Closer and defer
**Q: What is the correct defer pattern for Closer?**
```go
// Pattern 1 — common but swallows errors
func readFile(path string) error {
    f, err := os.Open(path)
    if err != nil { return err }
    defer f.Close()
    // ... read
    return nil
}

// Pattern 2 — captures Close error (for writes)
func writeFile(path string) (err error) {
    f, err := os.Create(path)
    if err != nil { return err }
    defer func() {
        if cerr := f.Close(); cerr != nil && err == nil {
            err = cerr
        }
    }()
    // ... write
    return nil
}
```
**A:** Pattern 1 is fine for reads (close errors are rarely meaningful). Pattern 2 is important for writes — `Close` flushes OS buffers, and ignoring its error can silently lose data.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the correct defer pattern for Closer?
**Your Response:** For reading files, Pattern 1 is usually fine - if closing fails during a read, it's rarely critical. But for writing files, Pattern 2 is important because `Close` flushes buffered data to the OS. If `Close` fails during a write, you might lose data. Pattern 2 captures the close error and returns it only if there wasn't already an error from the write operation.

---

## Section 3: Functional Options Pattern (Q39–Q50)

### 39. Simple Options Struct Pattern
**Q: What is the output?**
```go
package main
import "fmt"

type ServerConfig struct {
    host    string
    port    int
    timeout int
}

func NewServer(host string, port int, cfg ServerConfig) string {
    return fmt.Sprintf("%s:%d (timeout=%ds)", host, port, cfg.timeout)
}

func main() {
    cfg := ServerConfig{timeout: 30}
    fmt.Println(NewServer("localhost", 8080, cfg))
}
```
**A:** `localhost:8080 (timeout=30s)`. The options struct pattern groups optional parameters — but requires callers to construct a struct even for defaults.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `localhost:8080 (timeout=30s)`. This is the simple options struct pattern where you pass a configuration struct to your constructor. The downside is that callers must always create a struct, even for defaults. It's better than positional parameters but still requires some boilerplate from the caller.

---

### 40. Functional Options Pattern — Basic
**Q: What is the output?**
```go
package main
import "fmt"

type Server struct {
    host    string
    port    int
    timeout int
}

type Option func(*Server)

func WithHost(h string) Option    { return func(s *Server) { s.host = h } }
func WithPort(p int) Option       { return func(s *Server) { s.port = p } }
func WithTimeout(t int) Option    { return func(s *Server) { s.timeout = t } }

func NewServer(opts ...Option) *Server {
    s := &Server{host: "localhost", port: 8080, timeout: 30} // defaults
    for _, opt := range opts {
        opt(s)
    }
    return s
}

func main() {
    s := NewServer(WithPort(9090), WithTimeout(60))
    fmt.Printf("%s:%d timeout=%d\n", s.host, s.port, s.timeout)
}
```
**A:** `localhost:9090 timeout=60`. The functional options pattern: defaults are set at construction; callers override only what they need. Discoverable via IDE autocomplete, easily extensible without breaking API.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `localhost:9090 timeout=60`. This is the functional options pattern - each option is a function that modifies the struct. We set sensible defaults in the constructor, then callers can override only what they need. The benefits are: defaults are explicit, options are discoverable via IDE autocomplete, and you can add new options without breaking existing code.

---

### 41. Functional Options — Why Better Than Variadic Config Structs
**Q: What are the advantages?**
```go
// Hard to add new options without breaking callers:
func NewServer(host string, port int, timeout int, maxConns int) *Server

// Functional options — add WithMaxConns later, existing calls unaffected:
func NewServer(opts ...Option) *Server
```
**A:**
- **Backward compatible** — new options don't break existing callers
- **Self-documenting** — `WithTimeout(30)` is clearer than a positional `30`
- **Zero values don't mask intent** — omitting `WithTLS()` clearly means no TLS; whereas `false` in a struct is ambiguous
- **Testing** — easy to inject mock options

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are the advantages?
**Your Response:** The functional options pattern has several advantages over positional parameters or config structs. It's backward compatible - you can add new options without breaking existing code. It's self-documenting - `WithTimeout(30)` is clearer than just `30`. Zero values don't mask intent - omitting an option clearly means you want the default. And it's great for testing since you can easily inject mock options.

---

### 42. Functional Options — With Validation
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "errors"
)

type Client struct{ timeout int }
type Option func(*Client) error

func WithTimeout(t int) Option {
    return func(c *Client) error {
        if t <= 0 { return errors.New("timeout must be positive") }
        c.timeout = t
        return nil
    }
}

func NewClient(opts ...Option) (*Client, error) {
    c := &Client{timeout: 10}
    for _, opt := range opts {
        if err := opt(c); err != nil {
            return nil, err
        }
    }
    return c, nil
}

func main() {
    _, err := NewClient(WithTimeout(-1))
    fmt.Println(err)
}
```
**A:** `timeout must be positive`. Returning errors from options allows validation at construction time.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `timeout must be positive`. This shows how to add validation to the functional options pattern. Instead of returning `func(*Client)`, the option functions return `func(*Client) error`. The constructor checks each option's error and returns early if validation fails. This catches configuration errors at construction time rather than during use.

---

### 43. Rob Pike's Option Type — func(*T) returns Option
**Q: Does this pattern compile?**
```go
package main
import "fmt"

type Option func(*Builder)

type Builder struct {
    name  string
    debug bool
}

func Name(n string) Option  { return func(b *Builder) { b.name = n } }
func Debug() Option         { return func(b *Builder) { b.debug = true } }

func Build(opts ...Option) Builder {
    var b Builder
    for _, o := range opts { o(&b) }
    return b
}

func main() {
    b := Build(Name("myapp"), Debug())
    fmt.Println(b.name, b.debug)
}
```
**A:** `myapp true`. This is the canonical Go functional options pattern, described by Rob Pike in his 2014 talk.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this pattern compile?
**Your Response:** Yes, this prints `myapp true`. This is the canonical functional options pattern described by Rob Pike. The `Option` type is a function that modifies a `Builder`. Each option function returns an `Option` that makes specific changes. This pattern is elegant, type-safe, and widely used in Go libraries for configurable constructors.

---

### 44. Builder Pattern in Go
**Q: What is the output?**
```go
package main
import "fmt"

type QueryBuilder struct {
    table  string
    where  string
    limit  int
}

func NewQuery(table string) *QueryBuilder { return &QueryBuilder{table: table} }
func (q *QueryBuilder) Where(cond string) *QueryBuilder { q.where = cond; return q }
func (q *QueryBuilder) Limit(n int) *QueryBuilder { q.limit = n; return q }
func (q *QueryBuilder) Build() string {
    return fmt.Sprintf("SELECT * FROM %s WHERE %s LIMIT %d", q.table, q.where, q.limit)
}

func main() {
    sql := NewQuery("users").Where("age > 18").Limit(10).Build()
    fmt.Println(sql)
}
```
**A:** `SELECT * FROM users WHERE age > 18 LIMIT 10`. The builder pattern chains method calls — each method returns `*QueryBuilder`, enabling a fluent API.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `SELECT * FROM users WHERE age > 18 LIMIT 10`. This is the builder pattern, different from functional options. Each method returns the builder itself, allowing method chaining. It creates a fluent API that's great for building complex objects like SQL queries. The builder pattern is more verbose but very readable for complex constructions.

---

### 45. Singleton with sync.Once (Options Applied Once)
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

type Config struct{ dsn string }
var (
    cfg     *Config
    once    sync.Once
)

func GetConfig() *Config {
    once.Do(func() {
        cfg = &Config{dsn: "postgres://localhost/db"}
        fmt.Println("config initialised")
    })
    return cfg
}

func main() {
    GetConfig()
    GetConfig()
    fmt.Println(GetConfig().dsn)
}
```
**A:**
```
config initialised
postgres://localhost/db
```
Initialisation prints only once regardless of how many times `GetConfig` is called.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `config initialised` then `postgres://localhost/db`. `sync.Once` ensures that the initialization function runs exactly once, no matter how many times `GetConfig` is called. This is perfect for singleton patterns where you need expensive one-time initialization that's thread-safe. The first call triggers initialization, subsequent calls just return the cached value.

---

## Section 4: Channel Composition Patterns (Q46–Q60)

### 46. Generator Pattern — Function Returns a Channel
**Q: What is the output?**
```go
package main
import "fmt"

func naturals(n int) <-chan int {
    ch := make(chan int)
    go func() {
        for i := 1; i <= n; i++ {
            ch <- i
        }
        close(ch)
    }()
    return ch
}

func main() {
    for v := range naturals(5) {
        fmt.Print(v, " ")
    }
}
```
**A:** `1 2 3 4 5 `. Returning a receive-only channel from a generator function is idiomatic Go for lazy, infinite, or async sequences.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `1 2 3 4 5 `. This is the generator pattern - a function that returns a channel. The goroutine generates numbers and sends them to the channel, then closes it. The caller can range over the channel to receive values. This pattern is great for lazy evaluation, infinite sequences, or when you want to produce values asynchronously.

---

### 47. Pipeline — Stage 1: Generate, Stage 2: Transform
**Q: What is the output?**
```go
package main
import "fmt"

func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums { out <- n }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in { out <- n * n }
        close(out)
    }()
    return out
}

func main() {
    c := generate(2, 3, 4)
    sq := square(c)
    for v := range sq {
        fmt.Print(v, " ")
    }
}
```
**A:** `4 9 16 `. The pipeline pattern: each stage reads from an upstream channel and writes to a downstream one. Stages are composable and concurrent.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `4 9 16`. This demonstrates the pipeline pattern where each stage processes data and passes it to the next. The `generate` function produces numbers, the `square` function squares them, and each runs in its own goroutine. Pipelines are composable - you can chain multiple stages, and they run concurrently for better performance.

---

### 48. Fan-In — Merge Multiple Channels
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

func merge(cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    for _, c := range cs {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for v := range ch { out <- v }
        }(c)
    }
    go func() { wg.Wait(); close(out) }()
    return out
}

func gen(vals ...int) <-chan int {
    c := make(chan int, len(vals))
    for _, v := range vals { c <- v }
    close(c)
    return c
}

func main() {
    merged := merge(gen(1, 2), gen(3, 4))
    var results []int
    for v := range merged { results = append(results, v) }
    fmt.Println(len(results))
}
```
**A:** `4`. All 4 values are merged into one channel (order non-deterministic). Fan-in collects results from multiple concurrent producers.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `4`. This is the fan-in pattern - merging multiple channels into one. We have two generators producing values, and the `merge` function combines them into a single output channel. The order is non-deterministic because the goroutines run concurrently, but all 4 values are collected. Fan-in is perfect for aggregating results from multiple concurrent operations.

---

### 49. Done Channel — Early Termination of Pipeline
**Q: What is the pattern?**
```go
func generate(done <-chan struct{}, nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            select {
            case out <- n:
            case <-done: // abort if upstream cancelled
                return
            }
        }
    }()
    return out
}
```
**A:** The `done` channel pattern allows any stage in a pipeline to be cancelled. When `done` is closed, all goroutines watching it return immediately — preventing goroutine leaks when downstream consumers stop early.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the pattern?
**Your Response:** This shows the done channel pattern for graceful shutdown. The `done` channel is passed to each stage, and they check it in a select statement. When you close `done`, all stages exit immediately instead of blocking forever. This prevents goroutine leaks when downstream consumers stop early and is essential for building robust concurrent systems.

---

### 50. Or-Done Channel — Clean Range over Cancellable Channel
**Q: What does the orDone helper do?**
```go
func orDone(done, c <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for {
            select {
            case <-done:
                return
            case v, ok := <-c:
                if !ok { return }
                select {
                case out <- v:
                case <-done:
                    return
                }
            }
        }
    }()
    return out
}
```
**A:** `orDone` wraps a channel so that ranging over `orDone(done, c)` automatically stops when `done` is closed. This lets you write `for v := range orDone(done, myChan)` instead of nested `select` statements everywhere.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What does the orDone helper do?
**Your Response:** The `orDone` helper eliminates boilerplate code when ranging over channels that might be cancelled. Instead of writing the same select pattern everywhere, you wrap the channel with `orDone`. It returns a channel that closes when either the original channel closes or the `done` channel is closed, making the code much cleaner and less error-prone.

---

### 51. Repeat Pattern — Infinite Channel of a Value
**Q: What is the output?**
```go
package main
import "fmt"

func repeat(done <-chan struct{}, values ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for {
            for _, v := range values {
                select {
                case <-done:
                    return
                case out <- v:
                }
            }
        }
    }()
    return out
}

func take(done <-chan struct{}, in <-chan int, n int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for i := 0; i < n; i++ {
            select {
            case <-done:
                return
            case out <- <-in:
            }
        }
    }()
    return out
}

func main() {
    done := make(chan struct{})
    defer close(done)
    for v := range take(done, repeat(done, 1, 2, 3), 7) {
        fmt.Print(v, " ")
    }
}
```
**A:** `1 2 3 1 2 3 1 `. `repeat` produces 1,2,3 forever; `take` stops after 7 values. Composable pipeline stages.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `1 2 3 1 2 3 1`. The `repeat` function creates an infinite stream of values, while `take` limits it to a specific number. This shows how pipeline stages are composable - you can combine them in different ways. The `repeat` function loops forever sending values, and `take` stops after receiving 7 values, then closes the output channel.

---

### 52. Rate Limiter with time.Tick
**Q: What is the pattern?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    requests := []int{1, 2, 3, 4, 5}
    limiter := time.Tick(200 * time.Millisecond)

    for _, req := range requests {
        <-limiter // block until next tick
        fmt.Println("request", req, "at", time.Now().Format("15:04:05.000"))
    }
}
```
**A:** One request processed every 200ms. `time.Tick` returns a channel that sends at the specified interval — a simple rate limiter. **Note:** `time.Tick` leaks the ticker if the goroutine exits; use `time.NewTicker` and call `Stop()` in production.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the pattern?
**Your Response:** This processes one request every 200ms using `time.Tick`. `time.Tick` returns a channel that sends the current time at regular intervals. By reading from this channel before each request, we rate limit the processing. Note that `time.Tick` can leak resources if the goroutine exits - in production, use `time.NewTicker` with `defer ticker.Stop()` instead.

---

### 53. time.NewTicker — Production Rate Limiter
**Q: What is the correct pattern?**
```go
ticker := time.NewTicker(200 * time.Millisecond)
defer ticker.Stop() // ← always stop to release resources

for {
    select {
    case <-ticker.C:
        doWork()
    case <-done:
        return
    }
}
```
**A:** `time.NewTicker` + `defer Stop()` is the production-safe version. The goroutine exits cleanly when `done` closes, and `Stop()` releases the ticker goroutine.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the correct pattern?
**Your Response:** The production-safe pattern uses `time.NewTicker()` with `defer ticker.Stop()`. Unlike `time.Tick`, this doesn't leak resources. The ticker runs in its own goroutine, and `Stop()` cleans it up. The select statement allows clean shutdown when `done` is closed, preventing goroutine leaks.

---

### 54. Burst Rate Limiter with Buffered Channel
**Q: What is the pattern?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    // Allow bursts of up to 3, then 1 per 200ms
    burstyLimiter := make(chan time.Time, 3)
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now() // pre-fill burst tokens
    }

    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

    for i := 1; i <= 5; i++ {
        <-burstyLimiter
        fmt.Println("request", i)
    }
}
```
**A:** Requests 1–3 fire immediately (burst). Requests 4–5 wait 200ms each. Buffered channel as token bucket — a classic Go pattern.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the pattern?
**Your Response:** This implements a burst rate limiter using a buffered channel as a token bucket. We pre-fill the channel with 3 tokens, allowing 3 immediate requests (the burst). After that, a goroutine adds a new token every 200ms, so subsequent requests wait. This pattern lets you handle bursts while maintaining an overall rate limit.

---

### 55. Worker Pool — Bounded Concurrency
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

func workerPool(jobs <-chan int, results chan<- int, workers int) {
    var wg sync.WaitGroup
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := range jobs {
                results <- j * j
            }
        }()
    }
    wg.Wait()
    close(results)
}

func main() {
    jobs := make(chan int, 10)
    results := make(chan int, 10)

    go workerPool(jobs, results, 3)

    for i := 1; i <= 5; i++ { jobs <- i }
    close(jobs)

    var total int
    for r := range results { total += r }
    fmt.Println("sum of squares:", total)
}
```
**A:** `sum of squares: 55` (1+4+9+16+25). Worker pool with 3 goroutines processes 5 jobs concurrently. Results are aggregated after all work is done.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `sum of squares: 55`. The worker pool pattern limits concurrency to 3 goroutines, processing 5 jobs concurrently. Each worker squares numbers and sends results to the results channel. After all jobs are sent and the jobs channel is closed, workers finish and we close the results channel. Then we sum all results. This pattern prevents creating too many goroutines while still getting concurrency benefits.

---

### 56. Timeout on a Slow Operation — select + goroutine
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "time"
)

func slowOp() <-chan string {
    ch := make(chan string, 1)
    go func() {
        time.Sleep(200 * time.Millisecond)
        ch <- "result"
    }()
    return ch
}

func main() {
    select {
    case res := <-slowOp():
        fmt.Println(res)
    case <-time.After(100 * time.Millisecond):
        fmt.Println("timeout")
    }
}
```
**A:** `timeout`. The operation takes 200ms but the timeout fires at 100ms. Note: `slowOp` goroutine still runs to completion — use context cancellation to truly cancel it.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `timeout`. The slow operation takes 200ms but we only wait 100ms. The `select` statement races between the operation result and the timeout. Since the timeout fires first, we print "timeout". Note that the `slowOp` goroutine continues running - for true cancellation, you'd need to use context.Context to signal the goroutine to stop.

---

### 57. Retry with Backoff
**Q: What is the pattern?**
```go
package main
import (
    "errors"
    "fmt"
    "time"
)

func retry(attempts int, sleep time.Duration, f func() error) error {
    for i := 0; i < attempts; i++ {
        if err := f(); err == nil {
            return nil
        }
        fmt.Printf("attempt %d failed, retrying in %v\n", i+1, sleep)
        time.Sleep(sleep)
        sleep *= 2 // exponential backoff
    }
    return errors.New("all attempts failed")
}
```
**A:** Classic retry with exponential backoff. Each failed attempt doubles the sleep duration. Pass `errors.New(...)` convertible context for the final error.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the pattern?
**Your Response:** This is a retry pattern with exponential backoff. It attempts the operation up to `attempts` times. After each failure, it sleeps for the specified duration, then doubles the sleep time for the next attempt. Exponential backoff is great for handling temporary failures without overwhelming the system. After all attempts fail, it returns an error indicating all attempts failed.

---

### 58. Channel Direction Enforced in Pipeline
**Q: Does this compile?**
```go
package main

func produce(out chan<- int) {
    out <- 42
    // <-out // compile error: cannot receive from send-only channel
}

func consume(in <-chan int) int {
    return <-in
    // in <- 1  // compile error: cannot send to receive-only channel
}

func main() {
    ch := make(chan int, 1)
    produce(ch)
    v := consume(ch)
    _ = v
}
```
**A:** **Yes.** Channel direction types (`chan<-`, `<-chan`) enforce producer/consumer roles at compile time — preventing accidental misuse in pipelines.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Does this compile?
**Your Response:** Yes, this compiles. Channel direction types enforce roles at compile time. `chan<- int` is a send-only channel - you can only send to it, not receive from it. `<-chan int` is receive-only - you can only receive from it, not send to it. This prevents bugs in pipelines where you might accidentally try to receive from a channel that should only send data.

---

### 59. Closing Channels to Signal Completion
**Q: What is the output?**
```go
package main
import "fmt"

func countdown(n int) <-chan int {
    ch := make(chan int)
    go func() {
        for i := n; i >= 0; i-- {
            ch <- i
        }
        close(ch) // signal completion
    }()
    return ch
}

func main() {
    for v := range countdown(3) {
        fmt.Print(v, " ")
    }
}
```
**A:** `3 2 1 0 `. Closing a channel is the idiomatic way to signal "no more data". `range` over a channel exits cleanly when it's closed.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `3 2 1 0`. Closing a channel is the idiomatic way to signal that there's no more data coming. The `range` keyword over a channel automatically exits when the channel is closed, making it perfect for consuming all values from a producer. This is much cleaner than manually checking for a special sentinel value.

---

### 60. sync.WaitGroup + Channel Results Collection
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    results := make(chan int, 5)

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            results <- n * 10
        }(i)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    var total int
    for r := range results {
        total += r
    }
    fmt.Println(total)
}
```
**A:** `150` (10+20+30+40+50). Pattern: goroutines send results to a buffered channel; a sentinel goroutine closes the channel after all workers finish; main ranges to collect results.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the output?
**Your Response:** This prints `150`. This pattern combines `sync.WaitGroup` with channels for result collection. The main goroutine starts 5 workers, each sending their result to a buffered channel. A sentinel goroutine waits for all workers to finish (using `WaitGroup`), then closes the results channel. The main goroutine ranges over the results channel to collect all values. This ensures we collect all results before proceeding.

---
