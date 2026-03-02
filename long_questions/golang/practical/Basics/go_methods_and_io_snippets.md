# Go — Method Sets, io Patterns & Functional Options Snippets

> **Format**: Each question is "predict the output / spot the bug / does it compile?" style.
> **Topics**: Value vs pointer receivers · Method sets · Interface satisfaction rules · `io.Reader`/`io.Writer` patterns · Functional options · Channel composition patterns

---

## Section 1: Value Receivers vs Pointer Receivers (Q1–Q15)

### 1. Value Receiver — Called on Value
**Q: Does this compile and what is the output?**
```go
package main
import "fmt"

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

---

### 19. io.Writer Interface
**Q: What is the io.Writer interface?**
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
**A:** `Write` writes `len(p)` bytes from `p`. Returns how many bytes were written and any error. Implementations: `os.Stdout`, `os.File`, `bytes.Buffer`, `http.ResponseWriter`, network conns.

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

---

### 35. io.SectionReader — Read a Portion of a File
**Q: What does this do?**
```go
// Read 20 bytes starting at offset 10 of a file:
sr := io.NewSectionReader(file, 10, 20)
data, _ := io.ReadAll(sr)
```
**A:** `io.SectionReader` implements `io.Reader`, `io.Seeker`, `io.ReaderAt` on a section of a `ReaderAt`. Useful for random-access reads without seeking the underlying resource.

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

---
