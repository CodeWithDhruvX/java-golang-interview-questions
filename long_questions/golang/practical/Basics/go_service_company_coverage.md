# Go — Remaining 15% Topics for Service Company 100% Coverage

> Covers: `fmt` verbs · `strings` & `strconv` stdlib · `math` · `os` / file I/O · JSON · `time` · `net/http` basics · `bufio` · `sort` · `log` · basic `testing`

---

## Section 1: fmt Package Deep Dives (Q1–Q12)

### 1. Printf Verb %T
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    fmt.Printf("%T\n", 42)
    fmt.Printf("%T\n", 3.14)
    fmt.Printf("%T\n", "hello")
    fmt.Printf("%T\n", true)
    fmt.Printf("%T\n", []int{1,2})
}
```
**A:**
```
int
float64
string
bool
[]int
```

---

### 2. Printf %v vs %+v vs %#v on Struct
**Q: What is the output?**
```go
package main
import "fmt"

type Point struct{ X, Y int }

func main() {
    p := Point{1, 2}
    fmt.Printf("%v\n", p)
    fmt.Printf("%+v\n", p)
    fmt.Printf("%#v\n", p)
}
```
**A:**
```
{1 2}
{X:1 Y:2}
main.Point{X:1, Y:2}
```

---

### 3. Sprintf Returns String
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    s := fmt.Sprintf("Name: %s, Age: %d", "Alice", 30)
    fmt.Println(s)
    fmt.Printf("%T\n", s)
}
```
**A:**
```
Name: Alice, Age: 30
string
```

---

### 4. Fprintf to a Writer
**Q: Does this compile and what does it print?**
```go
package main
import (
    "fmt"
    "os"
)

func main() {
    fmt.Fprintf(os.Stdout, "Hello, %s!\n", "Go")
}
```
**A:** **Yes.** Prints `Hello, Go!`. `fmt.Fprintf` writes formatted output to any `io.Writer`, including `os.Stdout`, files, or buffers.

---

### 5. Sscanf Parsing
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    var name string
    var age int
    fmt.Sscanf("Alice 30", "%s %d", &name, &age)
    fmt.Println(name, age)
}
```
**A:** `Alice 30`

---

### 6. %b %o %x for Integers
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    n := 255
    fmt.Printf("bin: %b\n", n)
    fmt.Printf("oct: %o\n", n)
    fmt.Printf("hex: %x\n", n)
    fmt.Printf("HEX: %X\n", n)
}
```
**A:**
```
bin: 11111111
oct: 377
hex: ff
HEX: FF
```

---

### 7. Width and Padding
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    fmt.Printf("|%10s|\n", "Go")
    fmt.Printf("|%-10s|\n", "Go")
    fmt.Printf("|%010d|\n", 42)
}
```
**A:**
```
|        Go|
|Go        |
|0000000042|
```

---

### 8. fmt.Errorf
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    err := fmt.Errorf("user %d not found", 42)
    fmt.Println(err)
    fmt.Printf("%T\n", err)
}
```
**A:**
```
user 42 not found
*errors.errorString
```

---

### 9. Stringer Interface with fmt.Println
**Q: What is the output?**
```go
package main
import "fmt"

type Direction int

const (
    North Direction = iota
    South
    East
    West
)

func (d Direction) String() string {
    return []string{"North", "South", "East", "West"}[d]
}

func main() {
    fmt.Println(North, East)
}
```
**A:** `North East`. `fmt.Println` calls `String()` automatically.

---

### 10. fmt.Sprint vs fmt.Sprintln
**Q: What is the output?**
```go
package main
import "fmt"

func main() {
    a := fmt.Sprint("Go", "Lang")
    b := fmt.Sprintln("Go", "Lang")
    fmt.Printf("[%s]\n", a)
    fmt.Printf("[%s]", b)
}
```
**A:**
```
[GoLang]
[Go Lang
]
```
`Sprint` concatenates without spaces (for non-string args it adds spaces). `Sprintln` adds spaces between args and a trailing newline.

---

### 11. %p for Pointer Address
**Q: Does this compile?**
```go
package main
import "fmt"

func main() {
    x := 42
    fmt.Printf("%p\n", &x)
}
```
**A:** **Yes.** Prints a hex memory address like `0xc000014090`. Output varies by run.

---

### 12. fmt.Scan Reading from Stdin (Concept)
**Q: What is the difference?**
```go
// fmt.Scan reads space-separated tokens
fmt.Scan(&a, &b)

// fmt.Scanln reads until newline
fmt.Scanln(&a, &b)

// fmt.Scanf reads with format string
fmt.Scanf("%d %s", &n, &s)
```
**A:** `Scan` treats whitespace/newlines as separators. `Scanln` stops at newline. `Scanf` uses format directives — useful for structured input in service-company coding tests.

---

## Section 2: strings & strconv Package (Q13–Q24)

### 13. strings.ToUpper / ToLower
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    s := "Hello, Go!"
    fmt.Println(strings.ToUpper(s))
    fmt.Println(strings.ToLower(s))
}
```
**A:**
```
HELLO, GO!
hello, go!
```

---

### 14. strings.TrimPrefix / TrimSuffix
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    s := "GoLang"
    fmt.Println(strings.TrimPrefix(s, "Go"))
    fmt.Println(strings.TrimSuffix(s, "Lang"))
    fmt.Println(strings.TrimPrefix(s, "Java")) // no match
}
```
**A:**
```
Lang
Go
GoLang
```

---

### 15. strings.Index / LastIndex
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    s := "go gopher"
    fmt.Println(strings.Index(s, "go"))
    fmt.Println(strings.LastIndex(s, "go"))
    fmt.Println(strings.Index(s, "java"))
}
```
**A:**
```
0
3
-1
```

---

### 16. strings.Count
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Count("cheese", "e"))
    fmt.Println(strings.Count("five", ""))
}
```
**A:**
```
3
5
```
Counting empty string returns `len(s)+1` (between each rune and at edges).

---

### 17. strings.Join
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    words := []string{"Go", "is", "awesome"}
    fmt.Println(strings.Join(words, " "))
    fmt.Println(strings.Join(words, "-"))
}
```
**A:**
```
Go is awesome
Go-is-awesome
```

---

### 18. strings.HasPrefix / HasSuffix
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    s := "https://golang.org"
    fmt.Println(strings.HasPrefix(s, "https"))
    fmt.Println(strings.HasSuffix(s, ".org"))
    fmt.Println(strings.HasPrefix(s, "http://"))
}
```
**A:**
```
true
true
false
```

---

### 19. strconv.Atoi and Itoa
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    n, err := strconv.Atoi("123")
    fmt.Println(n, err)

    _, err2 := strconv.Atoi("abc")
    fmt.Println(err2)

    s := strconv.Itoa(456)
    fmt.Println(s, fmt.Sprintf("%T", s))
}
```
**A:**
```
123 <nil>
strconv.Atoi: parsing "abc": invalid syntax
456 string
```

---

### 20. strconv.ParseFloat / ParseBool
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    f, _ := strconv.ParseFloat("3.14", 64)
    fmt.Printf("%.2f\n", f)

    b, _ := strconv.ParseBool("true")
    fmt.Println(b)

    b2, err := strconv.ParseBool("yes")
    fmt.Println(b2, err)
}
```
**A:**
```
3.14
true
false strconv.ParseBool: parsing "yes": invalid syntax
```

---

### 21. strconv.FormatInt with Base
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println(strconv.FormatInt(255, 2))
    fmt.Println(strconv.FormatInt(255, 16))
    fmt.Println(strconv.FormatInt(255, 10))
}
```
**A:**
```
11111111
ff
255
```

---

### 22. strings.Repeat
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Repeat("ab", 3))
    fmt.Println(strings.Repeat("-", 10))
}
```
**A:**
```
ababab
----------
```

---

### 23. strings.ContainsAny
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.ContainsAny("hello", "aeiou"))
    fmt.Println(strings.ContainsAny("rhythm", "aeiou"))
}
```
**A:**
```
true
false
```

---

### 24. strings.Map
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    rot13 := func(r rune) rune {
        if r >= 'a' && r <= 'z' {
            return 'a' + (r-'a'+13)%26
        }
        return r
    }
    fmt.Println(strings.Map(rot13, "hello"))
}
```
**A:** `uryyb`

---

## Section 3: math Package (Q25–Q29)

### 25. math.Abs / math.Sqrt
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Abs(-7.5))
    fmt.Println(math.Sqrt(16))
    fmt.Println(math.Pow(2, 10))
}
```
**A:**
```
7.5
4
1024
```

---

### 26. math.Max / math.Min
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Max(3.5, 7.2))
    fmt.Println(math.Min(3.5, 7.2))
}
```
**A:**
```
7.2
3.5
```

---

### 27. math.Floor / Ceil / Round
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Floor(2.9))
    fmt.Println(math.Ceil(2.1))
    fmt.Println(math.Round(2.5))
    fmt.Println(math.Round(2.4))
}
```
**A:**
```
2
3
3
2
```

---

### 28. math.MaxInt / MinInt (Go 1.17+)
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.MaxInt8)
    fmt.Println(math.MinInt8)
    fmt.Println(math.MaxInt32)
}
```
**A:**
```
127
-128
2147483647
```

---

### 29. math.Inf and math.IsNaN
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.IsInf(math.Inf(1), 1))
    fmt.Println(math.IsNaN(math.NaN()))
    fmt.Println(math.NaN() == math.NaN())
}
```
**A:**
```
true
true
false
```
NaN is never equal to itself — always use `math.IsNaN()`.

---

## Section 4: os Package & File I/O (Q30–Q38)

### 30. os.Args
**Q: What does os.Args contain?**
```go
package main
import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.Args[0])       // program name
    fmt.Println(len(os.Args))     // total args including program
}
```
**A:** `os.Args[0]` is the executable path. `len(os.Args)` is 1 when run without args.

---

### 31. os.Getenv / os.Setenv
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "os"
)

func main() {
    os.Setenv("APP_MODE", "production")
    mode := os.Getenv("APP_MODE")
    fmt.Println(mode)

    missing := os.Getenv("NOT_SET")
    fmt.Printf("[%s]\n", missing) // empty string, no panic
}
```
**A:**
```
production
[]
```

---

### 32. os.Exit Does Not Run Defers
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "os"
)

func main() {
    defer fmt.Println("this will NOT print")
    fmt.Println("before exit")
    os.Exit(1)
}
```
**A:** Only `before exit` is printed. `os.Exit` terminates immediately without running deferred functions.

---

### 33. Writing to a File
**Q: What does this do step by step?**
```go
package main
import (
    "log"
    "os"
)

func main() {
    f, err := os.Create("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    f.WriteString("Hello, File!")
}
```
**A:** Creates (or truncates) `test.txt`, writes `Hello, File!` to it, then closes the file when `main` returns via `defer`. Always defer `Close()` immediately after checking the error from `Create/Open`.

---

### 34. Reading a File with os.ReadFile (Go 1.16+)
**Q: What is the idiomatic way?**
```go
package main
import (
    "fmt"
    "log"
    "os"
)

func main() {
    data, err := os.ReadFile("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(data))
}
```
**A:** Reads entire file into a `[]byte`. `os.ReadFile` is the modern replacement for `ioutil.ReadFile` (deprecated). Output: contents of `test.txt`.

---

### 35. bufio.Scanner for Line-by-Line Reading
**Q: What does this pattern do?**
```go
package main
import (
    "bufio"
    "fmt"
    "strings"
)

func main() {
    input := "line1\nline2\nline3"
    scanner := bufio.NewScanner(strings.NewReader(input))
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}
```
**A:**
```
line1
line2
line3
```
`bufio.Scanner` is the standard pattern for reading line-by-line from files or stdin.

---

### 36. bufio.Writer Buffered Write
**Q: Why is Flush() necessary?**
```go
package main
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    w := bufio.NewWriter(os.Stdout)
    fmt.Fprintln(w, "buffered")
    w.Flush() // Without this, output may not appear
}
```
**A:** `bufio.Writer` accumulates data in a buffer and writes in bulk. `Flush()` forces any buffered data to be sent to the underlying writer. Without it, output may be lost.

---

### 37. os.Stat — Check if File Exists
**Q: What is the idiomatic file existence check?**
```go
package main
import (
    "errors"
    "fmt"
    "os"
)

func fileExists(path string) bool {
    _, err := os.Stat(path)
    return !errors.Is(err, os.ErrNotExist)
}

func main() {
    fmt.Println(fileExists("test.txt"))
    fmt.Println(fileExists("nope.txt"))
}
```
**A:** Returns `true`/`false` depending on disk state. `os.ErrNotExist` is the sentinel error for missing files — never string-match error messages.

---

### 38. os.MkdirAll vs os.Mkdir
**Q: What is the difference?**
```go
// os.Mkdir: creates exactly ONE directory; fails if parent missing
os.Mkdir("a/b/c", 0755) // error if a/b doesn't exist

// os.MkdirAll: creates all intermediate dirs; no error if already exists
os.MkdirAll("a/b/c", 0755) // safe, idiomatic
```
**A:** `os.MkdirAll` is the safe, idiomatic choice when you don't know if parent directories exist — equivalent to `mkdir -p` in shell.

---

## Section 5: JSON Encoding / Decoding (Q39–Q46)

### 39. json.Marshal Basic Types
**Q: What is the output?**
```go
package main
import (
    "encoding/json"
    "fmt"
)

func main() {
    data, _ := json.Marshal(map[string]interface{}{
        "name": "Alice",
        "age":  30,
        "active": true,
    })
    fmt.Println(string(data))
}
```
**A:** `{"active":true,"age":30,"name":"Alice"}` *(map keys sorted alphabetically)*

---

### 40. json.Unmarshal into Struct
**Q: What is the output?**
```go
package main
import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    raw := `{"name":"Bob","age":25}`
    var p Person
    json.Unmarshal([]byte(raw), &p)
    fmt.Println(p.Name, p.Age)
}
```
**A:** `Bob 25`

---

### 41. json.MarshalIndent
**Q: What is the output?**
```go
package main
import (
    "encoding/json"
    "fmt"
)

func main() {
    m := map[string]int{"a": 1, "b": 2}
    data, _ := json.MarshalIndent(m, "", "  ")
    fmt.Println(string(data))
}
```
**A:**
```json
{
  "a": 1,
  "b": 2
}
```

---

### 42. Unknown JSON Fields Are Ignored
**Q: What is p.Name after Unmarshal?**
```go
package main
import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
}

func main() {
    raw := `{"name":"Carol","city":"NYC","age":22}`
    var p Person
    json.Unmarshal([]byte(raw), &p)
    fmt.Println(p.Name)
}
```
**A:** `Carol`. Unknown fields (`city`, `age`) are silently ignored during unmarshal by default.

---

### 43. json.Decoder for Streaming
**Q: What does this pattern do?**
```go
package main
import (
    "encoding/json"
    "fmt"
    "strings"
)

func main() {
    r := strings.NewReader(`{"name":"Dave"}`)
    var p struct{ Name string }
    json.NewDecoder(r).Decode(&p)
    fmt.Println(p.Name)
}
```
**A:** `Dave`. `json.NewDecoder` is preferred over `json.Unmarshal` when reading from `io.Reader` (HTTP request body, files) as it avoids loading the full JSON into memory.

---

### 44. omitempty Tag
**Q: What is the output?**
```go
package main
import (
    "encoding/json"
    "fmt"
)

type Config struct {
    Host    string `json:"host"`
    Port    int    `json:"port,omitempty"`
    Debug   bool   `json:"debug,omitempty"`
    Timeout int    `json:"timeout,omitempty"`
}

func main() {
    c := Config{Host: "localhost"}
    data, _ := json.Marshal(c)
    fmt.Println(string(data))
}
```
**A:** `{"host":"localhost"}`. Zero-value fields with `omitempty` are excluded.

---

### 45. json.RawMessage — Delay Parsing
**Q: What is the output?**
```go
package main
import (
    "encoding/json"
    "fmt"
)

type Envelope struct {
    Type    string          `json:"type"`
    Payload json.RawMessage `json:"payload"`
}

func main() {
    raw := `{"type":"user","payload":{"name":"Eve"}}`
    var env Envelope
    json.Unmarshal([]byte(raw), &env)
    fmt.Println(env.Type)
    fmt.Println(string(env.Payload))
}
```
**A:**
```
user
{"name":"Eve"}
```
`json.RawMessage` stores the raw JSON bytes for deferred/conditional parsing.

---

### 46. JSON Number to interface{} is float64
**Q: What is the output?**
```go
package main
import (
    "encoding/json"
    "fmt"
)

func main() {
    var m map[string]interface{}
    json.Unmarshal([]byte(`{"id":100}`), &m)
    v := m["id"]
    fmt.Printf("%T %v\n", v, v)
}
```
**A:** `float64 100`. All JSON numbers unmarshal as `float64` when target is `interface{}`. Use `json.Number` or a typed struct to avoid this.

---

## Section 6: time Package (Q47–Q54)

### 47. time.Now and time.Since
**Q: What does this code demonstrate?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    time.Sleep(100 * time.Millisecond)
    elapsed := time.Since(start)
    fmt.Printf("%.0fms\n", elapsed.Seconds()*1000)
}
```
**A:** Prints approximately `100ms`. `time.Since(t)` is shorthand for `time.Now().Sub(t)`.

---

### 48. time.Duration Arithmetic
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    d := 2*time.Hour + 30*time.Minute + 15*time.Second
    fmt.Println(d)
    fmt.Println(d.Hours())
    fmt.Println(d.Minutes())
    fmt.Println(d.Seconds())
}
```
**A:**
```
2h30m15s
2.504166...
150.25
9015
```

---

### 49. time.Parse and Format
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    // Go's reference time: Mon Jan 2 15:04:05 MST 2006
    t, _ := time.Parse("2006-01-02", "2025-03-02")
    fmt.Println(t.Format("02 Jan 2006"))
    fmt.Println(t.Year(), t.Month(), t.Day())
}
```
**A:**
```
02 Mar 2025
2025 March 2
```
Go uses a **specific reference time** (`2006-01-02 15:04:05`) as the format template — not `YYYY-MM-DD`.

---

### 50. time.After and time.Tick (Concept)
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    ch := time.After(50 * time.Millisecond)
    <-ch
    fmt.Println("50ms passed")
}
```
**A:** `50ms passed` (after ~50ms). `time.After` returns a channel that receives after the duration.

---

### 51. time.Unix Conversion
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    t := time.Unix(0, 0).UTC()
    fmt.Println(t)
}
```
**A:** `1970-01-01 00:00:00 +0000 UTC` — the Unix epoch.

---

### 52. Adding Duration to Time
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    t := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
    future := t.Add(24 * time.Hour)
    fmt.Println(future.Format("2006-01-02"))
}
```
**A:** `2025-01-02`

---

### 53. time.Before and After
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    t1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
    t2 := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
    fmt.Println(t1.Before(t2))
    fmt.Println(t1.After(t2))
    fmt.Println(t1.Equal(t2))
}
```
**A:**
```
true
false
false
```

---

### 54. time.Ticker
**Q: What is the pattern?**
```go
package main
import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(100 * time.Millisecond)
    defer ticker.Stop()

    count := 0
    for range ticker.C {
        count++
        fmt.Println("tick", count)
        if count == 3 {
            break
        }
    }
}
```
**A:** Prints `tick 1`, `tick 2`, `tick 3` — one every 100ms. Unlike `time.After`, `Ticker` fires repeatedly. Always call `Stop()` to free resources.

---

## Section 7: sort Package (Q55–Q59)

### 55. sort.Slice
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sort"
)

func main() {
    nums := []int{5, 2, 8, 1, 9}
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    fmt.Println(nums)
}
```
**A:** `[1 2 5 8 9]`

---

### 56. sort.Strings / sort.Ints
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sort"
)

func main() {
    words := []string{"banana", "apple", "cherry"}
    sort.Strings(words)
    fmt.Println(words)

    nums := []int{3, 1, 4, 1, 5}
    sort.Ints(nums)
    fmt.Println(nums)
}
```
**A:**
```
[apple banana cherry]
[1 1 3 4 5]
```

---

### 57. sort.Slice Descending
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sort"
)

func main() {
    s := []int{3, 1, 4, 1, 5, 9}
    sort.Slice(s, func(i, j int) bool {
        return s[i] > s[j] // reverse
    })
    fmt.Println(s)
}
```
**A:** `[9 5 4 3 1 1]`

---

### 58. sort.Search (Binary Search)
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "sort"
)

func main() {
    s := []int{1, 3, 5, 7, 9}
    target := 5
    i := sort.Search(len(s), func(i int) bool {
        return s[i] >= target
    })
    fmt.Println(i, s[i])
}
```
**A:** `2 5`. `sort.Search` returns the smallest index where the condition is true.

---

### 59. sort.SliceStable
**Q: When should you use SliceStable vs Slice?**
```go
package main
import (
    "fmt"
    "sort"
)

type Item struct {
    Name     string
    Priority int
}

func main() {
    items := []Item{{"b", 1}, {"a", 1}, {"c", 2}}
    sort.SliceStable(items, func(i, j int) bool {
        return items[i].Priority < items[j].Priority
    })
    for _, it := range items {
        fmt.Println(it.Name, it.Priority)
    }
}
```
**A:**
```
b 1
a 1
c 2
```
`SliceStable` preserves the original order of equal elements (`b` before `a` since both have priority 1). `sort.Slice` may not.

---

## Section 8: log Package (Q60–Q63)

### 60. log vs fmt for Errors
**Q: What is the difference?**
```go
package main
import (
    "fmt"
    "log"
)

func main() {
    fmt.Println("info message")
    log.Println("log message") // adds timestamp prefix
}
```
**A:** `log.Println` automatically adds a timestamp and writes to `os.Stderr` by default. `fmt.Println` writes to `os.Stdout` with no decoration.

---

### 61. log.Fatal
**Q: What is the output?**
```go
package main
import (
    "fmt"
    "log"
)

func main() {
    defer fmt.Println("deferred")
    log.Fatal("fatal error occurred")
    fmt.Println("after fatal") // unreachable
}
```
**A:** Logs the fatal message then calls `os.Exit(1)`. `deferred` is **NOT** printed — `log.Fatal` exits without running deferred functions.

---

### 62. log.SetPrefix
**Q: What is the output?**
```go
package main
import "log"

func main() {
    log.SetPrefix("[APP] ")
    log.SetFlags(0) // remove timestamp
    log.Println("Server started")
    log.Println("Listening on :8080")
}
```
**A:**
```
[APP] Server started
[APP] Listening on :8080
```
`log.SetPrefix` and `log.SetFlags(0)` are commonly used to customize log output in service applications.

---

### 63. Custom Logger
**Q: What is the output?**
```go
package main
import (
    "log"
    "os"
)

func main() {
    logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
    logger.Println("application started")
}
```
**A:** `INFO: 2025/03/02 16:20:31 application started` (date/time varies). `log.New` creates a custom logger — standard pattern in Go services.

---

## Section 9: Basic net/http (Q64–Q70)

### 64. Simple HTTP Server
**Q: What does this create?**
```go
package main
import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Go!")
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)
}
```
**A:** A basic HTTP server listening on port 8080. Any request to `/` responds `Hello, Go!`. `nil` uses the default ServeMux. `http.ResponseWriter` implements `io.Writer`.

---

### 65. Reading Query Parameters
**Q: What is the output for request `/search?q=golang`?**
```go
func search(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    fmt.Fprintf(w, "Searching for: %s", q)
}
```
**A:** `Searching for: golang`. `r.URL.Query()` returns a `url.Values` map; `.Get()` returns the first value for the key.

---

### 66. Writing HTTP Status Codes
**Q: What is the correct order?**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    // CORRECT: WriteHeader before Write
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintln(w, "not found")

    // WRONG:
    // fmt.Fprintln(w, "not found")
    // w.WriteHeader(http.StatusNotFound) // too late; headers already sent
}
```
**A:** `WriteHeader` must be called before any `Write`. Once you call `Write`, Go auto-sends `200 OK` and headers are locked.

---

### 67. http.StatusXxx Constants
**Q: What are the values?**
```go
package main
import (
    "fmt"
    "net/http"
)

func main() {
    fmt.Println(http.StatusOK)           // 200
    fmt.Println(http.StatusCreated)      // 201
    fmt.Println(http.StatusBadRequest)   // 400
    fmt.Println(http.StatusUnauthorized) // 401
    fmt.Println(http.StatusNotFound)     // 404
    fmt.Println(http.StatusInternalServerError) // 500
}
```
**A:** `200 201 400 401 404 500`. Always use named constants over raw numbers.

---

### 68. Reading Request Body
**Q: What is the correct pattern?**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()             // always close
    body, err := io.ReadAll(r.Body)  // Go 1.16+
    if err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    fmt.Println(string(body))
}
```
**A:** Always `defer r.Body.Close()` to prevent resource leaks. `io.ReadAll` is the modern replacement for `ioutil.ReadAll`.

---

### 69. http.Error Helper
**Q: What does this send to the client?**
```go
http.Error(w, "something went wrong", http.StatusInternalServerError)
```
**A:** Sets `Content-Type: text/plain; charset=utf-8`, writes the HTTP status code `500`, and writes the message body `something went wrong\n`. Cleaner than manually calling `WriteHeader` + `Write`.

---

### 70. JSON Response Pattern
**Q: What is the correct pattern to send JSON from a handler?**
```go
func jsonHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    resp := map[string]string{"status": "ok"}
    json.NewEncoder(w).Encode(resp)
}
```
**A:** Always set `Content-Type` header **before** writing body. Use `json.NewEncoder(w)` to stream JSON directly to the `ResponseWriter` without a temporary buffer.

---

## Section 10: Basic Testing Patterns (Q71–Q75)

### 71. Basic Unit Test Structure
**Q: What is the standard Go test file pattern?**
```go
// add.go
package math

func Add(a, b int) int { return a + b }

// add_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2,3) = %d; want %d", got, want)
    }
}
```
**A:** Test files end in `_test.go`. Functions start with `Test`. Run with `go test ./...`. `t.Errorf` marks the test as failed but continues; `t.Fatalf` marks as failed and stops immediately.

---

### 72. Table-Driven Tests
**Q: What is the table-driven test pattern?**
```go
package main

import "testing"

func Multiply(a, b int) int { return a * b }

func TestMultiply(t *testing.T) {
    tests := []struct {
        a, b, want int
    }{
        {2, 3, 6},
        {0, 5, 0},
        {-1, 4, -4},
    }
    for _, tc := range tests {
        got := Multiply(tc.a, tc.b)
        if got != tc.want {
            t.Errorf("Multiply(%d,%d) = %d; want %d", tc.a, tc.b, got, tc.want)
        }
    }
}
```
**A:** The standard Go idiom for testing multiple cases cleanly. Easy to add new cases without new functions.

---

### 73. t.Fatal vs t.Error
**Q: What is the difference?**
```go
func TestSomething(t *testing.T) {
    result := riskyOperation()
    if result == nil {
        t.Fatal("result must not be nil") // stops test immediately
    }
    // safe to use result below because Fatal stopped if nil
    if result.Value != 42 {
        t.Errorf("expected 42, got %d", result.Value) // continues
    }
}
```
**A:** `t.Fatal` (and `t.Fatalf`) stops the test immediately — use when subsequent assertions would panic or make no sense. `t.Error` (and `t.Errorf`) marks failure but continues collecting errors.

---

### 74. go test -v and -run Flags
**Q: What do these commands do?**
```bash
go test ./...              # run all tests in all packages
go test -v ./...           # verbose: show each test name and PASS/FAIL
go test -run TestAdd ./... # run only tests matching "TestAdd"
go test -cover ./...       # show code coverage percentage
```
**A:** These are the four most common `go test` commands used daily. `-cover` shows what % of code is exercised by tests.

---

### 75. Benchmark Function
**Q: What is the benchmark pattern?**
```go
package main

import "testing"

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = 1 + 2
    }
}
```
**A:** Run with `go test -bench=. ./...`. `b.N` is automatically adjusted by the testing framework to get a stable measurement. Benchmarks don't run with `go test` alone — you must pass `-bench`.

---

*End — This file covers the remaining 15% topics for 100% Service Company Interview Coverage.*
*Combined with `go_basics_fundamentals_snippets.md` and `go_basics_indepth_snippets.md`, you have full coverage.*
